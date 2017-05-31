package e2d

import (
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hw"
)

/************************************************
 * Display Modes: basic dispatch
 ************************************************/

func (e2d *HwEngine2d) BeginFrame() {
	if gKeyState == nil {
		gKeyState = hw.GetKeyboardState()
	}

	// Read current display mode once per frame (do not switch between
	// display modes within a frame)
	e2d.dispmode = int((e2d.DispCnt.Value >> 16) & 3)
	if e2d.B() {
		e2d.dispmode &= 1
	}
	modLcd.DebugZ("begin frame").String("e2d", string(e2d.Name())).Int("mode", e2d.dispmode).End()

	e2d.capture_BeginFrame()
	e2d.layers_BeginFrame()
}

func (e2d *HwEngine2d) EndFrame() {
	e2d.layers_EndFrame()
	e2d.capture_EndFrame()
}

func (e2d *HwEngine2d) BeginLine(y int, screen gfx.Line) {
	if e2d.masterBrightChanged {
		e2d.updateMasterBrightTable()
		e2d.masterBrightChanged = false
	}
	if e2d.specialEffectsChanged {
		e2d.updateSpecialEffectsTable()
		e2d.specialEffectsChanged = false
	}

	e2d.curline = y
	e2d.curscreen = screen

	e2d.capture_BeginLine(y, screen)
	e2d.layers_BeginLine(y, screen)
}

func (e2d *HwEngine2d) EndLine(y int) {
	e2d.layers_EndLine(y)
	e2d.capture_EndLine(y)

	// Final output.
	// curscreen now contains the mixer output (bg/obj layers)
	screen := e2d.curscreen
	switch e2d.dispmode {
	case 0:
		// Display off -> output white
		for x := 0; x < cScreenWidth; x++ {
			screen.Set32(x, 0xFFFFFF)
		}

	case 1:
		// Apply master brightness to the screen output
		for i := 0; i < cScreenWidth; i++ {
			pix := screen.Get32(i)
			r := uint8(pix) & 0x1F
			g := uint8(pix>>5) & 0x1F
			b := uint8(pix>>10) & 0x1F
			screen.Set32(i, e2d.masterBrightR[r]|e2d.masterBrightG[g]|e2d.masterBrightB[b])
		}

	case 2:
		// VRAM display
		block := (e2d.DispCnt.Value >> 18) & 3
		vram := e2d.mc.VramLcdcBank(int(block))
		if vram == nil {
			vram = zero[:]
		} else {
			vram = vram[y*cScreenWidth*2:]
		}
		for x := 0; x < cScreenWidth; x++ {
			pix := emu.Read16LE(vram[x*2:])
			r := uint8(pix) & 0x1F
			g := uint8(pix>>5) & 0x1F
			b := uint8(pix>>10) & 0x1F
			screen.Set32(x, e2d.masterBrightR[r]|e2d.masterBrightG[g]|e2d.masterBrightB[b])
		}

	case 3:
		panic("mode 3 not implemented")
	}
}
