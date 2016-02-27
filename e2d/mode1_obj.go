package e2d

import (
	"ndsemu/emu"
	"ndsemu/emu/gfx"
)

var objWidth = []struct{ w, h int }{
	// square
	{1, 1}, {2, 2}, {4, 4}, {8, 8},
	// horizontal
	{2, 1}, {4, 1}, {4, 2}, {8, 4},
	// vertical
	{1, 2}, {1, 4}, {2, 4}, {4, 8},
}

func (e2d *HwEngine2d) DrawOBJ(ctx *gfx.LayerCtx, lidx int, sy int) {
	oam := e2d.mc.VramOAM(e2d.Idx)
	tiles := e2d.mc.VramLinearBank(e2d.Idx, VramLinearOAM, 0)

	mapping1d := (e2d.DispCnt.Value>>4)&1 != 0
	boundary := 32
	if mapping1d {
		boundary <<= (e2d.DispCnt.Value >> 20) & 3
	}

	/*
		for i := 0; i < 128; i++ {
			a0, a1, _ := le16(oam[i*8:]), le16(oam[i*8+2:]), le16(oam[i*8+4:])
			y := int(a0 & 0xff)
			x := int(a1 & 0x1ff)
			if y != 0xc0 {
				Emu.Log().Infof("oam=%d: pos=%d,%d", i, x, y)
			}
		}
	*/

	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		// If sprites are globally disabled, nothing to do
		if e2d.DispCnt.Value&(1<<12) == 0 {
			sy++
			continue
		}

		useExtPal := (e2d.DispCnt.Value & (1 << 31)) != 0

		// Go through the sprite list in reverse order, because an object with
		// lower index has HIGHER priority (so it gets drawn in front of all).
		// FIXME: This is actually a temporary hack because it will fail once we
		// emulate the sprite line limits; the correct solution would be
		// to go through in the correct order, but avoiding writing pixels
		// that have been already written to.
		for i := 127; i >= 0; i-- {
			a0, a1, a2 := emu.Read16LE(oam[i*8:]), emu.Read16LE(oam[i*8+2:]), emu.Read16LE(oam[i*8+4:])
			if a0&0x300 == 0x200 {
				continue
			}

			// Sprite mode: 0=normal, 1=semi-transparent, 2=window, 3=bitmap
			// mode := (a0 >> 10) & 3

			const XMask = 0x1FF
			const YMask = 0xFF

			x := int(a1 & XMask)
			y := int(a0 & YMask)
			if x >= cScreenWidth {
				x -= XMask + 1
			}
			if y >= cScreenHeight {
				y -= YMask + 1
			}

			// Get the object size. The size is expressed in number of chars,
			// not pixels.
			sz := objWidth[((a0>>14)<<2)|(a1>>14)]
			tw, th := sz.w, sz.h

			// If the sprite is visible
			// FIXME: this doesn't handle wrapping yet
			if sy >= y && sy < (y+th*8) && (x < cScreenWidth && (x+tw*8) >= 0) {
				tilenum := int(a2 & 1023)
				depth256 := (a0>>13)&1 != 0
				hflip := (a1>>12)&1 != 0
				vflip := (a1>>13)&1 != 0
				pri := (a2 >> 10) & 3
				pal := (a2 >> 12) & 0xF

				// Size of a char (in byte), depending on the color setting
				charSize := 32
				if depth256 {
					charSize = 64
				}

				// Compute the line being drawn *within* the current object.
				// This must also handle vertical flip (in which the whole
				// object is flipped, not just the single chars)
				y0 := (sy - y)
				if vflip {
					y0 = th*8 - y0 - 1
				}

				// Calculate the char row being drawn.
				ty := y0 / 8

				// Compute the offset within VRAM of the current object (for
				// now, its top-left pixel)
				vramOffset := tilenum * boundary

				// Adjust the offset to the beginning of the correct char row
				// within the object.
				// This depends on the 1D vs 2D tile mapping in VRAM; 1D
				// mapping means that tiles are arranged linearly in memory,
				// while 2D mapping means that tiles are arranged in a 2D grid
				// with a fixed width
				if mapping1d {
					vramOffset += (tw * charSize) * ty
				} else {
					if depth256 {
						vramOffset += (16 * charSize) * ty
					} else {
						vramOffset += (32 * charSize) * ty
					}
				}

				// Now calculate the line being drawn within the current char row
				y0 &= 7

				// Prepare initial src/dst pointer for drawing
				src := tiles.FetchPointer(vramOffset)
				dst := line
				dst.Add16(x)

				for j := 0; j < tw; j++ {
					tsrc := src[charSize*j:]
					if hflip {
						tsrc = src[charSize*(tw-j-1):]
					}

					if x > -8 && x < cScreenWidth {
						if depth256 {
							if !useExtPal {
								pal = 0
							}
							e2d.drawChar256(y0, tsrc, dst, hflip, pri, pal, useExtPal)
						} else {
							e2d.drawChar16(y0, tsrc, dst, hflip, pri, pal, false)
						}
					}
					dst.Add16(8)
					x += 8
				}
			}
		}
		sy++
	}
}
