package e2d

import (
	"ndsemu/emu/gfx"
	"ndsemu/emu/hw"
)

func (e2d *HwEngine2d) drawChar16(y int, src []byte, dst gfx.Line, hflip bool, attrs uint32, pal uint16, extpal bool) {
	src = src[y*4:]
	attrs |= uint32(pal) << 4
	if extpal {
		attrs |= (1 << 12)
	}
	if !hflip {
		for x := 0; x < 4; x++ {
			p0, p1 := uint32(src[x]&0xF), uint32(src[x]>>4)
			if p0 != 0 {
				// p0 = (p0 << 4) | p0
				dst.Set32(0, p0|attrs)
			}
			if p1 != 0 {
				// p1 = (p1 << 4) | p1
				dst.Set32(1, p1|attrs)
			}
			dst.Add32(2)
		}
	} else {
		for x := 3; x >= 0; x-- {
			p1, p0 := uint32(src[x]&0xF), uint32(src[x]>>4)
			if p0 != 0 {
				dst.Set32(0, p0|attrs)
			}
			if p1 != 0 {
				dst.Set32(1, p1|attrs)
			}
			dst.Add32(2)
		}
	}
}

func (e2d *HwEngine2d) drawChar256(y int, src []byte, dst gfx.Line, hflip bool, attrs uint32, pal uint16, extpal bool) {
	src = src[y*8:]
	attrs |= uint32(pal) << 8
	if extpal {
		attrs |= (1 << 12)
	}
	if !hflip {
		for x := 0; x < 8; x++ {
			p0 := uint32(src[x])
			if p0 != 0 {
				dst.Set32(0, p0|attrs)
			}
			dst.Add32(1)
		}
	} else {
		for x := 7; x >= 0; x-- {
			p0 := uint32(src[x])
			if p0 != 0 {
				dst.Set32(0, p0|attrs)
			}
			dst.Add32(1)
		}
	}
}

func (e2d *HwEngine2d) DrawBG(lidx int) func(gfx.Line) {
	regs := &e2d.bgregs[lidx]

	mapBase := int((*regs.Cnt>>8)&0x1F) * 2 * 1024
	charBase := int((*regs.Cnt>>2)&0xF) * 16 * 1024

	if e2d.A() {
		mapBase += int((e2d.DispCnt.Value>>27)&7) * 64 * 1024
		charBase += int((e2d.DispCnt.Value>>24)&7) * 64 * 1024
	}

	var tmaps [4]VramLinearBank
	for i := 0; i < 4; i++ {
		tmaps[i] = e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, mapBase+2048*i)
	}
	chars := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, charBase)
	onmask := uint32(1 << uint(8+lidx))

	y := 0
	return func(line gfx.Line) {
		if e2d.DispCnt.Value&onmask == 0 || gKeyState[hw.SCANCODE_1+lidx] != 0 {
			y++
			return
		}
		if (e2d.A() && gKeyState[hw.SCANCODE_9] != 0) || (e2d.B() && gKeyState[hw.SCANCODE_8] != 0) {
			y++
			return
		}

		// Check if we are in extended palette mode (more palettes available for
		// 256-color tiles).
		useExtPal := (e2d.DispCnt.Value & (1 << 30)) != 0

		pri := regs.priority()
		depth256 := regs.depth256()

		doubleh := (*regs.Cnt>>14)&0x1 != 0
		doublev := (*regs.Cnt>>15)&0x1 != 0
		mapx := int(*regs.XOfs)
		mapy := (y + int(*regs.YOfs))
		tmapidx := 0

		if doublev {
			mapy &= 511
		} else {
			mapy &= 255
		}
		mapYOff := 32 * ((mapy & 255) / 8)
		if mapy >= 256 {
			if doubleh {
				tmapidx += 2
			} else {
				tmapidx += 1
			}
		}

		line.Add32(-(mapx & 7))
		for x := 0; x <= cScreenWidth/8; x++ {
			if doubleh {
				mapx &= 511
			} else {
				mapx &= 255
			}
			tile := tmaps[tmapidx+mapx/256].Get16(mapYOff + ((mapx & 255) / 8))

			// Decode tile
			tnum := int(tile & 1023)
			hflip := (tile>>10)&1 != 0
			vflip := (tile>>11)&1 != 0
			pal := (tile >> 12) & 0xF

			// Calculate tile line (and apply vertical flip)
			ty := mapy & 7
			if vflip {
				ty = 7 - ty
			}

			attrs := uint32(pri) << 29
			if depth256 {
				ch := chars.FetchPointer(tnum * 64)
				// 256-color tiles only have one palette in normal (GBA) mode, but
				// can have multiple palettes in extended palette mode.
				// So we ignore the palette number if extended palette is disabled
				// (it should be already zero, but better safe than sorry)
				if !useExtPal {
					pal = 0
				}
				e2d.drawChar256(ty, ch, line, hflip, attrs, pal, useExtPal)
			} else {
				ch := chars.FetchPointer(tnum * 32)
				// 16-color tiles don't use extended palettes, so we always pass false
				// to the drawChar16() function
				e2d.drawChar16(ty, ch, line, hflip, attrs, pal, false)
			}
			line.Add32(8)

			mapx += 8
		}

		y++
	}
}
