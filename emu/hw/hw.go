package hw

import (
	"fmt"
	"sync"
	"time"
	"unsafe"

	"ndsemu/emu/gfx"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	kHwAudioBuffers = 3
)

type OutputConfig struct {
	Title             string // Name of the window (displayed in titlebar)
	Width, Height     int    // Size of the window in pixels
	FramePerSecond    int    // Number of frames per second when running at full speed
	EnforceSpeed      bool   // True if we want to block to enforce the requested FramePerSecond / Audio.Frequency
	NumBackBuffers    int    // Number of back buffers used; more buffers means smoother but laggier (default=2)
	AudioFrequency    int    // Audio frequency in hertz
	AudioChannels     int    // Number of output channels (1 or 2)
	AudioSampleSigned bool   // True if samples are signed, False if unsigned
}

type frame struct {
	video gfx.Buffer
	audio AudioBuffer
}

type Output struct {
	cfg OutputConfig

	quit    bool
	framech chan frame

	mouse struct {
		x, y    int
		buttons MouseButtons
	}

	screen      *sdl.Window
	renderer    *sdl.Renderer
	frame       *sdl.Texture
	framebuf    [][]byte
	framebufidx int

	videoEnabled bool
	audioEnabled bool
	framecounter int
	fpscounter   int
	fpsclock     uint32
	fpsticks     []time.Time
	fpsticksidx  int

	audioDev sdl.AudioDeviceID
	audiobuf []AudioBuffer
}

func NewOutput(cfg OutputConfig) *Output {
	sdl.Do(func() {
		if sdl.WasInit(sdl.INIT_VIDEO|sdl.INIT_AUDIO) == 0 {
			sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO)
		}
	})

	if cfg.NumBackBuffers == 0 {
		cfg.NumBackBuffers = 2
	}

	framebuf := make([][]byte, cfg.NumBackBuffers)
	for i := range framebuf {
		framebuf[i] = make([]byte, cfg.Width*cfg.Height*4)
	}
	audiobuf := make([]AudioBuffer, cfg.NumBackBuffers)
	samplesPerFrame := cfg.AudioFrequency/cfg.FramePerSecond + 1 // round up
	for i := range audiobuf {
		audiobuf[i] = make(AudioBuffer, samplesPerFrame*cfg.AudioChannels)
	}

	out := &Output{
		cfg:      cfg,
		framebuf: framebuf,
		audiobuf: audiobuf,
		framech:  make(chan frame, cfg.NumBackBuffers-2),
		fpsticks: make([]time.Time, cfg.FramePerSecond),
	}
	go out.render()
	go out.poll()
	return out
}

func (out *Output) EnableVideo(enable bool) {
	sdl.Do(func() {
		if enable && !out.videoEnabled {
			var err error

			out.screen, err = sdl.CreateWindow(out.cfg.Title,
				sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
				out.cfg.Width*4/2, out.cfg.Height*4/2, sdl.WINDOW_RESIZABLE)
			if err != nil {
				panic(err)
			}

			// Create a renderer than never sync with vsync.
			// Syncing is always done with audio, not vsync,
			out.renderer, err = sdl.CreateRenderer(out.screen, -1, 0)
			if err != nil {
				panic(err)
			}
			out.renderer.SetLogicalSize(out.cfg.Width, out.cfg.Height)

			// make the scaled rendering look smoother.
			sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "nearest")

			out.frame, err = out.renderer.CreateTexture(
				sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING,
				out.cfg.Width, out.cfg.Height)
			if err != nil {
				panic(err)
			}
		} else {
			out.frame.Destroy()
			out.frame = nil
			out.renderer.Destroy()
			out.renderer = nil
			out.screen.Destroy()
			out.screen = nil
		}

		out.videoEnabled = enable
	})
}

func (out *Output) EnableAudio(enable bool) {
	if !enable {
		panic("unimplemented")
	}

	sdl.Do(func() {
		var format sdl.AudioFormat
		if out.cfg.AudioSampleSigned {
			format = sdl.AUDIO_S16
		} else {
			format = sdl.AUDIO_U16
		}
		samplesPerFrame := out.cfg.AudioFrequency / out.cfg.FramePerSecond

		spec := sdl.AudioSpec{
			Freq:     int32(out.cfg.AudioFrequency),
			Format:   format,
			Channels: uint8(out.cfg.AudioChannels),
			Samples:  uint16(samplesPerFrame),
		}
		if dev, err := sdl.OpenAudioDevice("", false, &spec, nil, 0); err != nil {
			panic(err)
		} else {
			sdl.PauseAudioDevice(dev, false)
			out.audioDev = dev
		}

		out.audioEnabled = enable
	})
}

func (out *Output) BeginFrame() (gfx.Buffer, AudioBuffer) {
	out.framebufidx++
	if out.framebufidx == out.cfg.NumBackBuffers {
		out.framebufidx = 0
	}
	fbuf := gfx.NewBuffer(unsafe.Pointer(&out.framebuf[out.framebufidx][0]),
		out.cfg.Width, out.cfg.Height, out.cfg.Width*4)
	abuf := out.audiobuf[out.framebufidx]

	fc := out.framecounter % out.cfg.FramePerSecond
	nsamples0 := (out.cfg.AudioFrequency * fc) / out.cfg.FramePerSecond
	nsamples1 := (out.cfg.AudioFrequency * (fc + 1)) / out.cfg.FramePerSecond
	ns := nsamples1 - nsamples0
	return fbuf, abuf[:ns*out.cfg.AudioChannels]
}

func (out *Output) EndFrame(screen gfx.Buffer, audio AudioBuffer) {
	out.framecounter++
	// Send the frame to the render() goroutine; this normally avoids blocking unless we're going
	// too fast, in which case the channel buffer would be full and the call would block.
	out.framech <- frame{screen, audio}
}

func (out *Output) render() {
	for f := range out.framech {
		sdl.Do(func() {
			if out.videoEnabled {
				out.renderVideo(f.video)
			}

			// When audio is enabled, we use it to enforce the correct speed
			// instead of using a timer. This avoids sound cracks.
			if out.audioEnabled {
				if out.cfg.EnforceSpeed {
					// Wait until there is no more audio buffered. This means that all the audio
					// has been sent to the kernel and is being played.
					for sdl.GetQueuedAudioSize(out.audioDev) > uint32(len(f.audio)*2) {
						time.Sleep(100 * time.Microsecond)
					}
					out.renderAudio(f.audio)
				} else {
					// If speed is not enforced, we want to avoid queueing too much audio
					// as it would desync from video. So send audio only if
					// there's less than a frame's worth in the buffer.
					if sdl.GetQueuedAudioSize(out.audioDev) < uint32(len(f.audio)*2) {
						out.renderAudio(f.audio)
					}
				}
			} else {
				// If there's no audio, enforce speed using timers. We save the time at which
				// we rendered each frame in the last second, so that we sleep only averaging
				// the frame rate over a window of one second (it's smoother).
				if out.cfg.EnforceSpeed {
					since := time.Since(out.fpsticks[out.fpsticksidx])
					if since < time.Second {
						time.Sleep(time.Second - since)
					}
					out.fpsticks[out.fpsticksidx] = time.Now()
					out.fpsticksidx++
					if out.fpsticksidx == out.cfg.FramePerSecond {
						out.fpsticksidx = 0
					}
				}
			}

			// Update FPS counter in title bar
			out.fpscounter++
			if out.fpsclock+1000 < sdl.GetTicks() {
				out.screen.SetTitle(fmt.Sprintf("%s - %d FPS", out.cfg.Title, out.fpscounter))
				out.fpscounter = 0
				out.fpsclock += 1000
			}
		})

	}
}

func (out *Output) renderVideo(video gfx.Buffer) {
	out.frame.Update(nil, video.Pointer(), out.cfg.Width*4)
	out.renderer.Clear()
	out.renderer.Copy(out.frame, nil, nil)
	out.renderer.Present()
}

func (out *Output) renderAudio(audio AudioBuffer) {
	buf := (*[100000]uint8)(unsafe.Pointer(&audio[0]))
	sdl.QueueAudio(out.audioDev, (*buf)[:len(audio)*2])
}

type MouseButtons int

const (
	MouseButtonLeft MouseButtons = 1 << iota
	MouseButtonMiddle
	MouseButtonRight
)

func (out *Output) GetMouseState() (int, int, MouseButtons) {
	return out.mouse.x, out.mouse.y, out.mouse.buttons
}

var kstate []uint8
var kstateOnce sync.Once

func GetKeyboardState() []uint8 {
	kstateOnce.Do(func() {
		sdl.Do(func() {
			kstate = sdl.GetKeyboardState()
		})
	})
	return kstate
}

func (out *Output) poll() {
	for !out.quit {
		time.Sleep(16 * time.Millisecond)

		sdl.Do(func() {
			x, y, state := sdl.GetMouseState()

			// Scale back to logical size
			w, h := out.screen.GetSize()
			x = x * out.cfg.Width / w
			y = y * out.cfg.Height / h

			var buttons MouseButtons
			if state&sdl.BUTTON_LEFT != 0 {
				buttons |= MouseButtonLeft
			}
			if state&sdl.BUTTON_RIGHT != 0 {
				buttons |= MouseButtonRight
			}
			if state&sdl.BUTTON_MIDDLE != 0 {
				buttons |= MouseButtonMiddle
			}

			out.mouse.x = x
			out.mouse.y = y
			out.mouse.buttons = buttons

			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch t := event.(type) {
				case *sdl.QuitEvent:
					out.quit = true
					return
				case *sdl.KeyDownEvent:
					if t.Keysym.Sym == sdl.K_ESCAPE {
						out.quit = true
						return
					}
				}
			}
		})
	}
}

func (out *Output) Poll() bool {
	return !out.quit
}

func (out *Output) Screenshot(fn string) (gerr error) {
	sdl.Do(func() {
		surf, err := sdl.CreateRGBSurfaceFrom(
			unsafe.Pointer(&out.framebuf[0]),
			out.cfg.Width, out.cfg.Height, 32, out.cfg.Width*4,
			0x00000FF, 0x0000FF00, 0x00FF0000, 0)
		if err != nil {
			gerr = err
			return
		}
		surf.SaveBMP(fn)
		surf.Free()
	})
	return
}
