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

		if srca != 0 || srcb != 0 {
			modLcd.Fatalf("unimplemented display capture source=%d srca=%d srb=%d", source, srca, srcb)
		}

		// Begin capturing this frame
		e2d.dispcap.Enabled = true
		e2d.dispcap.Mode = int(source)
		e2d.dispcap.WBank = int((e2d.DispCapCnt.Value >> 16) & 3)
		e2d.dispcap.WOffset = ((e2d.DispCapCnt.Value >> 18) & 3) * 0x8000
		e2d.dispcap.RBank = int((e2d.DispCnt.Value >> 18) & 3)
		e2d.dispcap.ROffset = ((e2d.DispCapCnt.Value >> 26) & 3) * 0x8000
		e2d.dispcap.AlphaA = e2d.DispCapCnt.Value & 0x1F
		e2d.dispcap.AlphaB = (e2d.DispCapCnt.Value >> 8) & 0x1F
		if e2d.dispcap.AlphaA > 16 {
			e2d.dispcap.AlphaA = 16
		}
		if e2d.dispcap.AlphaB > 16 {
			e2d.dispcap.AlphaB = 16
		}

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
				"wbank": string(e2d.dispcap.WBank + 'A'),
				"woff":  e2d.dispcap.WOffset,
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

	screen := e2d.curscreen

	// If capture is enabled, capture the screen output
	// and since we go through the pixels, also apply the
	// master brightness (which must be applied AFTER capturing)
	if e2d.dispcap.Enabled && e2d.curline < e2d.dispcap.Height {
		vram := e2d.mc.VramRawBank(e2d.dispcap.WBank)
		vram = vram[e2d.dispcap.WOffset:]
		capbuf := gfx.NewLine(vram)

		vram = e2d.mc.VramRawBank(e2d.dispcap.RBank)
		vram = vram[e2d.dispcap.ROffset:]
		readbuf := gfx.NewLine(vram)

		switch e2d.dispcap.Mode {
		case 0:
			for i := 0; i < e2d.dispcap.Width; i++ {
				pix := screen.Get32(i)
				capbuf.Set16(i, uint16(pix)|0x8000)
			}
		case 1:
			for i := 0; i < e2d.dispcap.Width; i++ {
				pix := readbuf.Get16(i)
				capbuf.Set16(i, uint16(pix)|0x8000)
			}
		case 2, 3:
			eva := e2d.dispcap.AlphaA
			evb := e2d.dispcap.AlphaB
			for i := 0; i < e2d.dispcap.Width; i++ {
				pix1 := uint16(screen.Get32(i))
				pix2 := uint16(readbuf.Get16(i))
				r1, g1, b1 := (pix1 & 0x1F), ((pix1 >> 5) & 0x1F), ((pix1 >> 10) & 0x1F)
				r2, g2, b2 := (pix2 & 0x1F), ((pix2 >> 5) & 0x1F), ((pix2 >> 10) & 0x1F)

				r := (uint32(r1)*eva + uint32(r2)*evb) >> 4
				g := (uint32(g1)*eva + uint32(g2)*evb) >> 4
				b := (uint32(b1)*eva + uint32(b2)*evb) >> 4

				capbuf.Set16(i, 0x8000|uint16(r)|uint16(g)<<5|uint16(b)<<10)
			}
		}

		e2d.dispcap.ROffset += uint32(e2d.dispcap.Width * 2)
		if e2d.dispcap.ROffset == 128*1024 {
			e2d.dispcap.ROffset = 0
		}

		e2d.dispcap.WOffset += uint32(e2d.dispcap.Width * 2)
		if e2d.dispcap.WOffset == 128*1024 {
			e2d.dispcap.WOffset = 0
		}
	}

	// Apply master brightness and output to the screen
	for i := 0; i < 256; i++ {
		pix := screen.Get32(i)
		r := uint8(pix) & 0x1F
		g := uint8(pix>>5) & 0x1F
		b := uint8(pix>>10) & 0x1F
		screen.Set32(i, e2d.masterBrightR[r]|e2d.masterBrightG[g]|e2d.masterBrightB[b])
	}
}
