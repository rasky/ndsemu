package e2d

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/gfx"
)

const (
	objPixModeNormal = iota
	objPixModeAlpha
	objPixModeWindow
	objPixModeBitmap
)

const (
	objModeNormal = iota
	objModeAffine
	objModeHidden
	objModeAffineDouble
)

var objWidth = []struct{ w, h int }{
	// square
	{1, 1}, {2, 2}, {4, 4}, {8, 8},
	// horizontal
	{2, 1}, {4, 1}, {4, 2}, {8, 4},
	// vertical
	{1, 2}, {1, 4}, {2, 4}, {4, 8},
}

func objBitmap_CalcAddress_2D128(tilenum int) int {
	return int((tilenum&0xF)*0x10 + (tilenum & ^0xF)*0x80)
}

func objBitmap_CalcAddress_2D256(tilenum int) int {
	return int((tilenum&0x1F)*0x10 + (tilenum & ^0x1F)*0x80)
}

func objBitmap_CalcAddress_1D128(tilenum int) int {
	return int(tilenum) * 128
}

func objBitmap_CalcAddress_1D256(tilenum int) int {
	return int(tilenum) * 256
}

func (e2d *HwEngine2d) DrawOBJ(lidx int) func(gfx.Line) {
	oam := e2d.mc.VramOAM(e2d.Idx)
	tiles := e2d.mc.VramLinearBank(e2d.Idx, VramLinearOAM, 0)

	mapping1d := (e2d.DispCnt.Value>>4)&1 != 0
	boundary := 32
	if mapping1d {
		boundary <<= (e2d.DispCnt.Value >> 20) & 3
	}

	var vramBitmapCalcAddress func(int) int
	var objPitch int
	bitmapMapping1d := (e2d.DispCnt.Value>>6)&1 != 0
	if !bitmapMapping1d {
		// OBJ Bitmap 2D mapping
		if (e2d.DispCnt.Value>>5)&1 == 0 {
			vramBitmapCalcAddress = objBitmap_CalcAddress_2D128
			objPitch = 16
		} else {
			vramBitmapCalcAddress = objBitmap_CalcAddress_2D256
			objPitch = 32
		}
	} else {
		// OBJ Bitmap 1D mapping
		if (e2d.DispCnt.Value>>22)&1 == 0 {
			vramBitmapCalcAddress = objBitmap_CalcAddress_1D128
		} else {
			vramBitmapCalcAddress = objBitmap_CalcAddress_1D256
		}
	}

	if false {
		for i := 127; i >= 0; i-- {
			a0, a1, a2 := emu.Read16LE(oam[i*8:]), emu.Read16LE(oam[i*8+2:]), emu.Read16LE(oam[i*8+4:])
			mode := (a0 >> 8) & 3
			if mode == objModeHidden {
				continue
			}
			pixmode := (a0 >> 10) & 3
			const XMask = 0x1FF
			const YMask = 0xFF

			x := int(a1 & XMask)
			y := int(a0 & YMask)
			if y >= 129 {
				continue
			}
			widthChars := ((a0 >> 14) << 2) | (a1 >> 14)
			tilenum := int(a2 & 1023)
			depth256 := (a0>>13)&1 != 0
			hflip := (a1>>12)&1 != 0 && mode == objModeNormal // hflip not available in affine mode
			vflip := (a1>>13)&1 != 0 && mode == objModeNormal // vflip not available in affine mode
			pri := (a2 >> 10) & 3
			pal := (a2 >> 12) & 0xF

			fmt.Printf("%cOBJ %3d | %3d,%3d/%2d | %d,%d | %d/%d/%v | %v,%v | %d\n",
				e2d.Name(), i, x, y, widthChars, mode, pixmode, tilenum, pal, depth256, hflip, vflip, pri)
		}
	}

	objwindow := e2d.DispCnt.Value&(1<<15) != 0

	sy := 0
	return func(line gfx.Line) {
		// If sprites are globally disabled, nothing to do
		if e2d.DispCnt.Value&(1<<12) == 0 {
			sy++
			return
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

			// Sprite mode: 0=normal, 1=affine, 2=hidden, 3=affine double
			// We immediately skip hidden sprites, so from this point onward,
			// affine!=0 means affine mode.
			mode := (a0 >> 8) & 3
			if mode == objModeHidden {
				continue
			}

			// Sprite pixel mode: 0=normal, 1=semi-transparent, 2=window, 3=bitmap
			pixmode := (a0 >> 10) & 3

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
			tws, ths := tw, th

			if mode == objModeAffineDouble {
				tws *= 2
				ths *= 2
			}

			// If the sprite is visible
			// FIXME: this doesn't handle wrapping yet
			if sy >= y && sy < (y+ths*8) && (x < cScreenWidth && (x+tws*8) >= 0) {
				tilenum := int(a2 & 1023)
				depth256 := (a0>>13)&1 != 0
				hflip := (a1>>12)&1 != 0 && mode == objModeNormal // hflip not available in affine mode
				vflip := (a1>>13)&1 != 0 && mode == objModeNormal // vflip not available in affine mode
				pri := (a2 >> 10) & 3
				pal := (a2 >> 12) & 0xF

				// Size of a char (in byte), depending on the color setting
				charSize := 32
				if depth256 {
					charSize = 64
				}

				// Compute the offset within VRAM of the current object (for
				// now, its top-left pixel)
				var vramOffset int
				if pixmode == objPixModeBitmap {
					vramOffset = vramBitmapCalcAddress(tilenum)
				} else {
					vramOffset = tilenum * boundary
				}

				// Compute the line being drawn *within* the current object.
				// This must also handle vertical flip (in which the whole
				// object is flipped, not just the single chars)
				y0 := (sy - y)
				if vflip {
					y0 = ths*8 - y0 - 1
				}

				// Calculate the pitch of a sprite, expressed in number of chars)
				// This depends on the 1D vs 2D tile mapping in VRAM; 1D
				// mapping means that tiles are arranged linearly in memory so the
				// pitch is just the size of the sprite (in tiles).
				// In 2D mapping, tiles are arranged in a 2D grid with a fixed size
				// depending on the BPP, and thus not
				pitch := tw
				if pixmode == objPixModeBitmap {
					if !bitmapMapping1d {
						pitch = objPitch
					}
				} else {
					if !mapping1d {
						if depth256 {
							pitch = 16
						} else {
							pitch = 32
						}
					}
				}

				// See if we need to draw in affine mode
				if mode != objModeNormal {
					if pixmode == objPixModeBitmap {
						panic("bitmap mode not supported in affine")
					}
					if pixmode == objPixModeWindow {
						panic("window mode not supported in affine")
					}

					parms := ((a1>>9)&0x1F)*0x20 + 0x6
					dx := int(int16(emu.Read16LE(oam[parms:])))
					dmx := int(int16(emu.Read16LE(oam[parms+8:])))
					dy := int(int16(emu.Read16LE(oam[parms+16:])))
					dmy := int(int16(emu.Read16LE(oam[parms+24:])))

					sx := (tw*8/2)<<8 - (tws*8/2)*dx - (ths*8/2)*dmx + y0*dmx
					sy := (th*8/2)<<8 - (tws*8/2)*dy - (ths*8/2)*dmy + y0*dmy

					src := tiles.FetchPointer(vramOffset)
					dst := line

					attrs := uint32(pri) << 29
					attrs |= (4 << 26) // layer=4 -> obj
					if depth256 {
						if useExtPal {
							attrs |= uint32(pal<<8) | (1 << 12)
						}
					} else {
						attrs |= uint32(pal << 4)
						if useExtPal {
							attrs |= (1 << 12)
						}
					}

					for j := 0; j < tws*8; j++ {
						if x >= 0 && x < cScreenWidth {
							isx, isy := sx>>8, sy>>8
							if isx >= 0 && isx < tw*8 && isy >= 0 && isy < th*8 {
								ty := isy / 8
								off := (pitch * charSize) * ty
								isy &= 7

								tx := isx / 8
								off += charSize * tx
								isx &= 7

								if depth256 {
									pix := uint32(src[off+isy*8+isx])
									if pix != 0 {
										dst.Set32(x, pix|attrs)
									}
								} else {
									pix := uint32(src[off+isy*4+isx/2])
									pix >>= 4 * uint(isx&1)
									pix &= 0xF
									if pix != 0 {
										dst.Set32(x, pix|attrs)
									}
								}
							}
						}

						sx += dx
						sy += dy
						x++
					}

				} else {
					if pixmode == objPixModeWindow && objwindow {
						panic("window mode not supported for obj")
					}
					if pixmode == objPixModeBitmap {
						if hflip {
							modLcd.FatalZ("horizontal flip in obj bitmap").End()
						}

						vramOffset += (pitch * 8 * y0) * 2
						src := tiles.FetchPointer(vramOffset)
						dst := line

						attrs := (uint32(pri) << 29) | (4 << 26) | 0x80000000
						for j := 0; j < tw*8; j++ {
							if x >= 0 && x < cScreenWidth {
								px := uint32(emu.Read16LE(src[j*2:]))
								if px&0x8000 != 0 {
									dst.Set32(x, px|attrs)
								}
							}
							x++
						}

					} else {
						// Calculate the char row being drawn.
						ty := y0 / 8

						// Adjust the offset to the beginning of the correct char row
						// within the object.
						vramOffset += (pitch * charSize) * ty

						// Now calculate the line being drawn within the current char row
						y0 &= 7

						// Prepare initial src/dst pointer for drawing
						dst := line
						dst.Add32(x)

						attrs := (uint32(pri) << 29) | (4 << 26)
						for j := 0; j < tw; j++ {
							var tsrc []byte
							if !hflip {
								tsrc = tiles.FetchPointer(vramOffset + charSize*j)
							} else {
								tsrc = tiles.FetchPointer(vramOffset + charSize*(tw-j-1))
							}

							if x > -8 && x < cScreenWidth {
								if depth256 {
									if !useExtPal {
										pal = 0
									}
									e2d.drawChar256(y0, tsrc, dst, hflip, attrs, pal, useExtPal)
								} else {
									e2d.drawChar16(y0, tsrc, dst, hflip, attrs, pal, false)
								}
							}
							dst.Add32(8)
							x += 8
						}
					}
				}
			}
		}
		sy++
	}
}
