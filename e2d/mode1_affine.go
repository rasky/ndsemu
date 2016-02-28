package e2d

import (
	"ndsemu/emu/gfx"
	"ndsemu/emu/hw"
)

var bmpSize = []struct{ w, h int }{
	{128, 128}, {256, 256}, {512, 256}, {512, 512},
}

func (e2d *HwEngine2d) DrawBGAffine(ctx *gfx.LayerCtx, lidx int, y int) {
	regs := &e2d.bgregs[lidx]
	bgmode := e2d.bgmodes[lidx]

	var mapBase, charBase int

	switch bgmode {
	case BgModeLargeBitmap:
		// In large bitmap mode, the whole 512K vram is used for a single large bitmap,
		// so there is no offset
	case BgModeAffineBitmap, BgModeAffineBitmapDirect:
		mapBase = int((*regs.Cnt>>8)&0x1F) * 16 * 1024
		// charbase is obviously not used in bitmap modes, as there are no tiles
	default:
		mapBase = int((*regs.Cnt>>8)&0x1F) * 2 * 1024
		charBase = int((*regs.Cnt>>2)&0xF) * 16 * 1024
		if e2d.A() {
			mapBase += int((e2d.DispCnt.Value>>27)&7) * 64 * 1024
			charBase += int((e2d.DispCnt.Value>>24)&7) * 64 * 1024
		}
	}

	tmap := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, mapBase)
	chars := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, charBase)
	onmask := uint32(1 << uint(8+lidx))
	startx := int32(*regs.PX<<4) >> 4
	starty := int32(*regs.PY<<4) >> 4

	if e2d.DispCnt.Value&onmask != 0 {
		ch := string('A' + e2d.Idx)
		dx := int32(int16(*regs.PA))
		dy := int32(int16(*regs.PC))
		dmx := int32(int16(*regs.PB))
		dmy := int32(int16(*regs.PD))
		modLcd.Infof("%s%d: %v pos=(%x,%x), dx=(%x,%x), dy=(%x,%x) map=%x", ch, lidx, bgmode,
			startx, starty, dx, dy, dmx, dmy, mapBase)
	}

	if y != 0 {
		panic("unimplemented initial line not zero on affine plane")
	}

	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		if e2d.DispCnt.Value&onmask == 0 || gKeyState[hw.SCANCODE_1+lidx] != 0 {
			y++
			continue
		}
		if (e2d.A() && gKeyState[hw.SCANCODE_9] != 0) || (e2d.B() && gKeyState[hw.SCANCODE_8] != 0) {
			y++
			continue
		}

		pri := regs.priority()

		mapx := startx
		mapy := starty

		dx := int32(int16(*regs.PA))
		dy := int32(int16(*regs.PC))

		switch bgmode {
		case BgModeAffineBitmap:
			size := bmpSize[((*regs.Cnt >> 14) & 3)]

			for x := 0; x < cScreenWidth; x++ {
				px := int(mapx>>8) & (size.w - 1)
				py := int(mapy>>8) & (size.h - 1)

				// 8-bit bitmap layers don't use extended palettes, so
				// create a layer pixel without ext pal number
				// We only encode the priority bit
				col := uint16(tmap.Get8(py*size.w + px))
				col |= pri << 13
				line.Set16(x, col)
				mapx += dx
				mapy += dy
			}

		case BgModeAffineBitmapDirect:
			size := bmpSize[((*regs.Cnt >> 14) & 3)]

			for x := 0; x < cScreenWidth; x++ {
				px := int(mapx>>8) & (size.w - 1)
				py := int(mapy>>8) & (size.h - 1)

				// In Direct Color Bitmaps, bit 15 is used as a transparency
				// bit, so if not set, the pixel is not displayed.
				col := tmap.Get16(py*size.w + px)
				if col&0x8000 != 0 {
					// NOTE: in our output layer, we must mark direct colors
					// with bit 15 = 1, which matches exactly the format of
					// direct color bitmaps, so leave it as-is.
					line.Set16(x, col)
				}
				mapx += dx
				mapy += dy
			}

		case BgModeAffineMap16:
			// Check if we are in extended palette mode (more palettes available for
			// 256-color tiles).
			useExtPal := (e2d.DispCnt.Value & (1 << 30)) != 0

			size := 128 << ((*regs.Cnt >> 14) & 3)

			for x := 0; x < cScreenWidth; x++ {
				px := int(mapx>>8) & (size - 1)
				py := int(mapy>>8) & (size - 1)

				tx := px / 8
				ty := py / 8
				tile := tmap.Get16(ty*size/8 + tx)

				// Decode tile
				tnum := int(tile & 1023)
				hflip := (tile>>10)&1 != 0
				vflip := (tile>>11)&1 != 0
				pal := (tile >> 12) & 0xF

				// Calculate tile line (and apply vertical flip)
				ty = py & 7
				if vflip {
					ty = 7 - ty
				}

				ch := chars.FetchPointer(tnum*64 + ty*8)
				// 256-color tiles only have one palette in normal (GBA) mode, but
				// can have multiple palettes in extended palette mode.
				// So we ignore the palette number if extended palette is disabled
				// (it should be already zero, but better safe than sorry)
				attrs := pri << 13
				if useExtPal {
					attrs |= (pal << 8) | (1 << 12)
				}

				if !hflip {
					p0 := uint16(ch[px&7])
					if p0 != 0 {
						line.Set16(0, p0|attrs)
					}
				} else {
					p0 := uint16(ch[7-(px&7)])
					if p0 != 0 {
						line.Set16(0, p0|attrs)
					}
				}

				line.Add16(1)
				mapx += dx
				mapy += dy
			}

		default:
			panic("unimplemented")
		}

		dmx := int32(int16(*regs.PB))
		dmy := int32(int16(*regs.PD))
		startx += dmx
		starty += dmy

		y++
	}
}
