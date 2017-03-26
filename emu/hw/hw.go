package hw

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"

	"ndsemu/emu/gfx"
	log "ndsemu/emu/logger"

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
	AudioFrequency    int    // Audio frequency in hertz
	AudioChannels     int    // Number of output channels (1 or 2)
	AudioSampleSigned bool   // True if samples are signed, False if unsigned
}

type Output struct {
	cfg OutputConfig

	screen      *sdl.Window
	renderer    *sdl.Renderer
	frame       *sdl.Texture
	framebuf    [2][]byte
	framebufidx int

	videoEnabled bool
	audioEnabled bool
	framecounter int
	fpscounter   int
	fpsclock     uint32

	audiocounter int32 // atomic
	aindexw      int32 // atomic
	aindexr      int32 // atomic
	audiobuf     [kHwAudioBuffers]AudioBuffer
}

func NewOutput(cfg OutputConfig) *Output {
	if sdl.WasInit(sdl.INIT_VIDEO|sdl.INIT_AUDIO) == 0 {
		sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO)
	}

	return &Output{
		cfg: cfg,
		framebuf: [2][]byte{
			make([]byte, cfg.Width*cfg.Height*4),
			make([]byte, cfg.Width*cfg.Height*4),
		},
	}
}

func (out *Output) EnableVideo(enable bool) {
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
}

func (out *Output) EnableAudio(enable bool) {
	if !enable {
		panic("unimplemented")
	}

	var format sdl.AudioFormat
	if out.cfg.AudioSampleSigned {
		format = sdl.AUDIO_S16
	} else {
		format = sdl.AUDIO_U16
	}

	if out.cfg.AudioFrequency%out.cfg.FramePerSecond != 0 {
		panic("audio frequency must be a multiple of frames-per-second")
	}
	samplesPerFrame := out.cfg.AudioFrequency / out.cfg.FramePerSecond

	for i := range out.audiobuf {
		out.audiobuf[i] = make(AudioBuffer, samplesPerFrame*out.cfg.AudioChannels)
	}

	spec := sdl.AudioSpec{
		Freq:     int32(out.cfg.AudioFrequency),
		Format:   format,
		Channels: uint8(out.cfg.AudioChannels),
		Samples:  uint16(samplesPerFrame),
	}
	out.audioSpecSetCallback(&spec)
	if dev, err := sdl.OpenAudioDevice("", false, &spec, nil, 0); err != nil {
		panic(err)
	} else {
		sdl.PauseAudioDevice(dev, false)
	}

	out.audioEnabled = enable
}

func (out *Output) BeginFrame() (gfx.Buffer, AudioBuffer) {
	out.framebufidx = 1 - out.framebufidx
	fbuf := gfx.NewBuffer(unsafe.Pointer(&out.framebuf[out.framebufidx][0]),
		out.cfg.Width, out.cfg.Height, out.cfg.Width*4)

	aindexw := atomic.LoadInt32(&out.aindexw)
	if aindexw >= atomic.LoadInt32(&out.aindexr)+kHwAudioBuffers {
		if out.cfg.EnforceSpeed {
			log.ModHw.WithFields(log.Fields{
				"fc": fmt.Sprintf("%04d", out.framecounter),
				"ar": fmt.Sprintf("%04d", atomic.LoadInt32(&out.aindexr)),
				"aw": fmt.Sprintf("%04d", atomic.LoadInt32(&out.aindexw)),
			}).Warn("overflow audio buffer (producing too fast)")
		}
	}
	abuf := out.audiobuf[aindexw%kHwAudioBuffers]

	return fbuf, abuf
}

func (out *Output) EndFrame(screen gfx.Buffer, audio AudioBuffer) {
	if out.audioEnabled {
		atomic.AddInt32(&out.aindexw, 1)
	}

	if out.videoEnabled {
		if int(atomic.LoadInt32(&out.audiocounter)) < out.framecounter {
			out.frame.Update(nil, screen.Pointer(), out.cfg.Width*4)
			out.renderer.Clear()
			out.renderer.Copy(out.frame, nil, nil)
			out.renderer.Present()
			out.fpscounter++

			if out.cfg.EnforceSpeed {
				// Wait until audio catches up; this is where we slow down emulation
				// to match the desired framerate (but we do that syncing with audio
				// rathern than a timer).
				for int(atomic.LoadInt32(&out.audiocounter)) < out.framecounter {
					time.Sleep(1 * time.Millisecond)
				}
			}
		}

		if out.fpsclock+1000 < sdl.GetTicks() {
			out.screen.SetTitle(fmt.Sprintf("%s - %d FPS", out.cfg.Title, out.fpscounter))
			out.fpscounter = 0
			out.fpsclock += 1000
		}
	}

	out.framecounter++
}

func (out *Output) audioCallback(outbuf []int16) {

	aindexr := atomic.LoadInt32(&out.aindexr)

	if out.aindexr == atomic.LoadInt32(&out.aindexw) {
		fmt.Println("audio underflow: no audio generated, silencing")
		for i := range outbuf {
			outbuf[i] = 0
		}
		return
	}

	buf := out.audiobuf[aindexr%kHwAudioBuffers]
	if len(buf) != len(outbuf) {
		fmt.Println(len(buf), len(outbuf))
		panic("invalid audio buffer size")
	}
	copy(outbuf, buf)

	atomic.AddInt32(&out.aindexr, 1)
	atomic.AddInt32(&out.audiocounter, 1)
}

type MouseButtons int

const (
	MouseButtonLeft MouseButtons = 1 << iota
	MouseButtonMiddle
	MouseButtonRight
)

func (out *Output) GetMouseState() (x, y int, buttons MouseButtons) {
	x, y, state := sdl.GetMouseState()

	// Scale back to logical size
	w, h := out.screen.GetSize()
	x = x * out.cfg.Width / w
	y = y * out.cfg.Height / h

	if state&sdl.BUTTON_LEFT != 0 {
		buttons |= MouseButtonLeft
	}
	if state&sdl.BUTTON_RIGHT != 0 {
		buttons |= MouseButtonRight
	}
	if state&sdl.BUTTON_MIDDLE != 0 {
		buttons |= MouseButtonMiddle
	}

	return x, y, buttons
}

func GetKeyboardState() []uint8 {
	return sdl.GetKeyboardState()
}

func (out *Output) Poll() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyDownEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				return false
			}
		}
	}
	return true
}

func (out *Output) Screenshot(fn string) error {
	surf, err := sdl.CreateRGBSurfaceFrom(
		unsafe.Pointer(&out.framebuf[0]),
		out.cfg.Width, out.cfg.Height, 32, out.cfg.Width*4,
		0x00000FF, 0x0000FF00, 0x00FF0000, 0)
	if err != nil {
		return err
	}
	surf.SaveBMP(fn)
	surf.Free()
	return nil
}
