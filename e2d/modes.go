package e2d

import (
	"ndsemu/emu/gfx"
	"ndsemu/emu/hw"
	log "ndsemu/emu/logger"
)

/************************************************
 * Display Modes: basic dispatch
 ************************************************/

func (e2d *HwEngine2d) BeginFrame() {

	if gKeyState == nil {
		gKeyState = hw.GetKeyboardState()
	}

	// Check if display capture is activated
	if e2d.DispCapCnt.Value&(1<<31) != 0 {
		source := (e2d.DispCapCnt.Value >> 29) & 3
		srca := (e2d.DispCapCnt.Value >> 24) & 1
		srcb := (e2d.DispCapCnt.Value >> 25) & 1

		if source != 0 {
			modLcd.Fatalf("unimplemented display capture source=%d", source)
		}
		if srca != 0 {
			modLcd.Fatalf("unimplemented display capture srca=%d", srca)
		}

		// Begin capturing this frame
		e2d.dispcap.Enabled = true
		e2d.dispcap.Bank = int((e2d.DispCapCnt.Value >> 16) & 3)
		e2d.dispcap.Offset = ((e2d.DispCapCnt.Value >> 18) & 3) * 0x8000

		switch (e2d.DispCapCnt.Value >> 20) & 3 {
		case 0:
			e2d.dispcap.Width = 128
			e2d.dispcap.Height = 128
		case 1:
			e2d.dispcap.Width = 256
			e2d.dispcap.Height = 64
		case 2:
			e2d.dispcap.Width = 256
			e2d.dispcap.Height = 128
		case 3:
			e2d.dispcap.Width = 256
			e2d.dispcap.Height = 192
		}

		modLcd.WithDelayedFields(func() log.Fields {
			return log.Fields{
				"src":   source,
				"sa":    srca,
				"sb":    srcb,
				"wbank": string(e2d.dispcap.Bank + 'A'),
				"woff":  e2d.dispcap.Offset,
				"w":     e2d.dispcap.Width,
				"h":     e2d.dispcap.Height,
			}
		}).Infof("Capture activated")

	}

	// Read current display mode once per frame (do not switch between
	// display modes within a frame)
	e2d.dispmode = int((e2d.DispCnt.Value >> 16) & 3)
	if e2d.B() {
		e2d.dispmode &= 1
	}
	e2d.modeTable[e2d.dispmode].BeginFrame()
}

func (e2d *HwEngine2d) EndFrame() {
	e2d.modeTable[e2d.dispmode].EndFrame()

	// If we were capturing, stop it
	if e2d.dispcap.Enabled {
		e2d.dispcap.Enabled = false
		e2d.DispCapCnt.Value &^= 1 << 31
	}
}

func (e2d *HwEngine2d) BeginLine(y int, screen gfx.Line) {
	if e2d.masterBrightChanged {
		e2d.updateMasterBrightTable()
		e2d.masterBrightChanged = false
	}

	e2d.curline = y
	e2d.curscreen = screen
	e2d.modeTable[e2d.dispmode].BeginLine(y, screen)
}

func (e2d *HwEngine2d) EndLine(y int) {
	e2d.modeTable[e2d.dispmode].EndLine(y)

	i := 0
	screen := e2d.curscreen

	// If capture is enabled, capture the screen output
	// and since we go through the pixels, also apply the
	// master brightness (which must be applied AFTER capturing)
	if e2d.dispcap.Enabled && e2d.curline < e2d.dispcap.Height {
		vram := e2d.mc.VramRawBank(e2d.dispcap.Bank)
		vram = vram[e2d.dispcap.Offset:]
		capbuf := gfx.NewLine(vram)
		for ; i < e2d.dispcap.Width; i++ {
			pix := screen.Get32(i)
			// Set the output alpha bit to 1
			capbuf.Set16(i, uint16(pix)|0x8000)
			r := uint8(pix) & 0x1F
			g := uint8(pix>>5) & 0x1F
			b := uint8(pix>>10) & 0x1F
			screen.Set32(i, e2d.masterBrightR[r]|e2d.masterBrightG[g]|e2d.masterBrightB[b])
		}
		e2d.dispcap.Offset += uint32(e2d.dispcap.Width * 2)
		if e2d.dispcap.Offset == 128*1024 {
			e2d.dispcap.Offset = 0
		}
	}

	// Apply master brightness on the remaining pixels
	// (if capture is disabled, this will be the whole line)
	for ; i < 256; i++ {
		pix := screen.Get32(i)
		r := uint8(pix) & 0x1F
		g := uint8(pix>>5) & 0x1F
		b := uint8(pix>>10) & 0x1F
		screen.Set32(i, e2d.masterBrightR[r]|e2d.masterBrightG[g]|e2d.masterBrightB[b])
	}
}
