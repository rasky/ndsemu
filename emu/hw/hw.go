package hw

import (
	"fmt"
	"unsafe"

	"ndsemu/emu/gfx"

	"github.com/veandco/go-sdl2/sdl"
)

type OutputConfig struct {
	Title         string
	Width, Height int
	WaitVSync     bool
}

type Output struct {
	cfg OutputConfig

	screen   *sdl.Window
	renderer *sdl.Renderer
	frame    *sdl.Texture
	framebuf []byte

	videoEnabled bool
	audioEnabled bool
	framecounter int
	fpscounter   int
	fpsclock     uint32
}

func NewOutput(cfg OutputConfig) *Output {
	if sdl.WasInit(sdl.INIT_VIDEO) == 0 {
		sdl.Init(sdl.INIT_VIDEO)
	}

	return &Output{
		cfg:      cfg,
		framebuf: make([]byte, cfg.Width*cfg.Height*4),
	}
}

func (out *Output) EnableVideo(enable bool) {
	if enable && !out.videoEnabled {
		var err error

		out.screen, err = sdl.CreateWindow(out.cfg.Title,
			sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			out.cfg.Width*3/2, out.cfg.Height*3/2, sdl.WINDOW_RESIZABLE)
		if err != nil {
			panic(err)
		}

		rflags := uint32(0)
		if out.cfg.WaitVSync {
			rflags |= sdl.RENDERER_PRESENTVSYNC
		}
		out.renderer, err = sdl.CreateRenderer(
			out.screen, -1, rflags)
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

func (out *Output) BeginFrame() gfx.Buffer {
	return gfx.NewBuffer(unsafe.Pointer(&out.framebuf[0]),
		out.cfg.Width, out.cfg.Height, out.cfg.Width*4)
}

func (out *Output) EndFrame() {
	if out.videoEnabled {
		out.frame.Update(nil, unsafe.Pointer(&out.framebuf[0]), out.cfg.Width*4)
		out.renderer.Clear()
		out.renderer.Copy(out.frame, nil, nil)
		out.renderer.Present()
		out.fpscounter++

		if out.fpsclock+1000 < sdl.GetTicks() {
			out.screen.SetTitle(fmt.Sprintf("%s - %d FPS", out.cfg.Title, out.fpscounter))
			out.fpscounter = 0
			out.fpsclock += 1000
		}
	}

	out.framecounter++
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

func (out *Output) GetKeyboardState() []uint8 {
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

func init() {
}
