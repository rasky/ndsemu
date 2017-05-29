package e2d

import "ndsemu/emu/gfx"

var zero [1024]byte

func (e2d *HwEngine2d) capture_BeginFrame() {
	// Check if display capture is activated
	if e2d.A() && e2d.DispCapCnt.Value&(1<<31) != 0 {
		source := (e2d.DispCapCnt.Value >> 29) & 3
		srca := (e2d.DispCapCnt.Value >> 24) & 1
		srcb := (e2d.DispCapCnt.Value >> 25) & 1

		if srcb != 0 {
			modLcd.FatalZ("unimplemented display capture").
				Uint32("source", source).Uint32("srca", srca).Uint32("srcb", srcb).
				End()
		}

		// Begin capturing this frame
		e2d.dispcap.Enabled = true
		e2d.dispcap.Mode = int(source)
		e2d.dispcap.SrcA = int(srca)
		e2d.dispcap.SrcB = int(srcb)
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

		modLcd.InfoZ("capture activated").
			Uint32("src", source).
			Uint32("sa", srca).
			Uint32("sb", srcb).
			String("wbank", string(e2d.dispcap.WBank+'A')).
			Uint32("woff", e2d.dispcap.WOffset).
			String("rbank", string(e2d.dispcap.RBank+'A')).
			Uint32("roff", e2d.dispcap.ROffset).
			Int("w", e2d.dispcap.Width).
			Int("h", e2d.dispcap.Height).
			End()
	}
}

func (e2d *HwEngine2d) capture_EndFrame() {
	// If we were capturing, stop it
	if e2d.dispcap.Enabled {
		e2d.dispcap.Enabled = false
		e2d.DispCapCnt.Value &^= 1 << 31
	}
}

func (e2d *HwEngine2d) capture_BeginLine(y int, screen gfx.Line) {

}

func (e2d *HwEngine2d) capture(y int) {
	screen := e2d.curscreen

	vram := e2d.mc.VramLcdcBank(e2d.dispcap.WBank)
	if vram == nil {
		// If destination bank is not allocated to LCDC,
		// skip capture. This is confirmed by mariokart which
		// otherwise crashes: in fact, bank D is used for
		// capture, but at some point it is allocated as ARM7
		// and used to run code from it.
		return
	}
	vram = vram[e2d.dispcap.WOffset:]
	capbuf := gfx.NewLine(vram)

	vram = e2d.mc.VramLcdcBank(e2d.dispcap.RBank)
	if vram == nil {
		// If source bank is not allocated to LCDC,
		// capture zero bytes.
		vram = zero[:]
	} else {
		vram = vram[e2d.dispcap.ROffset:]
	}
	readbuf := gfx.NewLine(vram)

	var srca gfx.Line
	if e2d.dispcap.SrcA == 0 {
		// Source A is the final mixer output
		srca = e2d.curscreen
	} else {
		// Source A is the 3D layer only. See in which
		// layer it is enabled.
		if e2d.l3dIdx < 0 {
			panic("capturing 3D but 3D is disabled")
		}
		srca = e2d.lm.LayerBuffer(e2d.l3dIdx)
	}

	switch e2d.dispcap.Mode {
	case 0:
		for i := 0; i < e2d.dispcap.Width; i++ {
			pix := srca.Get32(i)
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
}

func (e2d *HwEngine2d) capture_EndLine(y int) {
	// If capture is enabled, capture the screen output
	// and since we go through the pixels, also apply the
	// master brightness (which must be applied AFTER capturing)
	if e2d.dispcap.Enabled && e2d.curline < e2d.dispcap.Height {
		e2d.capture(y)

		e2d.dispcap.ROffset += uint32(e2d.dispcap.Width * 2)
		if e2d.dispcap.ROffset == 128*1024 {
			e2d.dispcap.ROffset = 0
		}

		e2d.dispcap.WOffset += uint32(e2d.dispcap.Width * 2)
		if e2d.dispcap.WOffset == 128*1024 {
			e2d.dispcap.WOffset = 0
		}
	}

}
