package main

import (
	"encoding/binary"
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

const (
	cScreenWidth  = 256
	cScreenHeight = 192
)

func le16(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
}

type bgRegs struct {
	Cnt        *uint16
	XOfs, YOfs *uint16
	PA, PB     *uint16
	PC, PD     *uint16
	PX, PY     *uint32
}

type HwEngine2d struct {
	Idx      int
	DispCnt  hwio.Reg32 `hwio:"offset=0x00,wcb"`
	Bg0Cnt   hwio.Reg16 `hwio:"offset=0x08"`
	Bg1Cnt   hwio.Reg16 `hwio:"offset=0x0A"`
	Bg2Cnt   hwio.Reg16 `hwio:"offset=0x0C"`
	Bg3Cnt   hwio.Reg16 `hwio:"offset=0x0E"`
	Bg0XOfs  hwio.Reg16 `hwio:"offset=0x10,writeonly"`
	Bg0YOfs  hwio.Reg16 `hwio:"offset=0x12,writeonly"`
	Bg1XOfs  hwio.Reg16 `hwio:"offset=0x14,writeonly"`
	Bg1YOfs  hwio.Reg16 `hwio:"offset=0x16,writeonly"`
	Bg2XOfs  hwio.Reg16 `hwio:"offset=0x18,writeonly"`
	Bg2YOfs  hwio.Reg16 `hwio:"offset=0x1A,writeonly"`
	Bg3XOfs  hwio.Reg16 `hwio:"offset=0x1C,writeonly"`
	Bg3YOfs  hwio.Reg16 `hwio:"offset=0x1E,writeonly"`
	Bg2PA    hwio.Reg16 `hwio:"offset=0x20,writeonly"`
	Bg2PB    hwio.Reg16 `hwio:"offset=0x22,writeonly"`
	Bg2PC    hwio.Reg16 `hwio:"offset=0x24,writeonly"`
	Bg2PD    hwio.Reg16 `hwio:"offset=0x26,writeonly"`
	Bg2PX    hwio.Reg32 `hwio:"offset=0x28,writeonly"`
	Bg2PY    hwio.Reg32 `hwio:"offset=0x2C,writeonly"`
	Bg3PA    hwio.Reg16 `hwio:"offset=0x30,writeonly"`
	Bg3PB    hwio.Reg16 `hwio:"offset=0x32,writeonly"`
	Bg3PC    hwio.Reg16 `hwio:"offset=0x34,writeonly"`
	Bg3PD    hwio.Reg16 `hwio:"offset=0x36,writeonly"`
	Bg3PX    hwio.Reg32 `hwio:"offset=0x38,writeonly"`
	Bg3PY    hwio.Reg32 `hwio:"offset=0x3C,writeonly"`
	Win0X    hwio.Reg16 `hwio:"offset=0x40,writeonly"`
	Win1X    hwio.Reg16 `hwio:"offset=0x42,writeonly"`
	Win0Y    hwio.Reg16 `hwio:"offset=0x44,writeonly"`
	Win1Y    hwio.Reg16 `hwio:"offset=0x46,writeonly"`
	WinIn    hwio.Reg16 `hwio:"offset=0x48"`
	WinOut   hwio.Reg16 `hwio:"offset=0x4A"`
	Mosaic   hwio.Reg16 `hwio:"offset=0x4C,writeonly"`
	BldCnt   hwio.Reg16 `hwio:"offset=0x50"`
	BldAlpha hwio.Reg16 `hwio:"offset=0x52"`
	BldY     hwio.Reg32 `hwio:"offset=0x54"`

	bgregs  [4]bgRegs
	mc      *HwMemoryController
	lineBuf [4 * (cScreenWidth + 16)]byte
	lm      gfx.LayerManager
}

func NewHwEngine2d(idx int, mc *HwMemoryController) *HwEngine2d {
	e2d := new(HwEngine2d)
	hwio.MustInitRegs(e2d)
	e2d.Idx = idx
	e2d.mc = mc

	// Initialize bgregs data structure which is easier to index
	// compared to the raw registers
	e2d.bgregs[0].Cnt = &e2d.Bg0Cnt.Value
	e2d.bgregs[0].XOfs = &e2d.Bg0XOfs.Value
	e2d.bgregs[0].YOfs = &e2d.Bg0YOfs.Value

	e2d.bgregs[1].Cnt = &e2d.Bg1Cnt.Value
	e2d.bgregs[1].XOfs = &e2d.Bg1XOfs.Value
	e2d.bgregs[1].YOfs = &e2d.Bg1YOfs.Value

	e2d.bgregs[2].Cnt = &e2d.Bg2Cnt.Value
	e2d.bgregs[2].XOfs = &e2d.Bg2XOfs.Value
	e2d.bgregs[2].YOfs = &e2d.Bg2YOfs.Value
	e2d.bgregs[2].PA = &e2d.Bg2PA.Value
	e2d.bgregs[2].PB = &e2d.Bg2PB.Value
	e2d.bgregs[2].PC = &e2d.Bg2PC.Value
	e2d.bgregs[2].PD = &e2d.Bg2PD.Value
	e2d.bgregs[2].PX = &e2d.Bg2PX.Value
	e2d.bgregs[2].PY = &e2d.Bg2PY.Value

	e2d.bgregs[3].Cnt = &e2d.Bg3Cnt.Value
	e2d.bgregs[3].XOfs = &e2d.Bg3XOfs.Value
	e2d.bgregs[3].YOfs = &e2d.Bg3YOfs.Value
	e2d.bgregs[3].PA = &e2d.Bg3PA.Value
	e2d.bgregs[3].PB = &e2d.Bg3PB.Value
	e2d.bgregs[3].PC = &e2d.Bg3PC.Value
	e2d.bgregs[3].PD = &e2d.Bg3PD.Value
	e2d.bgregs[3].PX = &e2d.Bg3PX.Value
	e2d.bgregs[3].PY = &e2d.Bg3PY.Value

	// Initialize layer manager
	e2d.lm.Cfg = gfx.LayerManagerConfig{
		Width:          cScreenWidth,
		Height:         cScreenHeight,
		ScreenBpp:      4,
		LayerBpp:       1,
		OverflowPixels: 8,
		Mixer:          e2dMixer_DisplayOff,
	}

	// Background layers
	e2d.lm.AddLayer(gfx.LayerFunc{e2d.DrawBG})
	e2d.lm.AddLayer(gfx.LayerFunc{e2d.DrawBG})
	e2d.lm.AddLayer(gfx.LayerFunc{e2d.DrawBG})
	e2d.lm.AddLayer(gfx.LayerFunc{e2d.DrawBG})

	// Sprites layer
	e2d.lm.AddLayer(gfx.LayerFunc{e2d.DrawOBJ})
	return e2d
}

func (e2d *HwEngine2d) A() bool    { return e2d.Idx == 0 }
func (e2d *HwEngine2d) B() bool    { return e2d.Idx != 0 }
func (e2d *HwEngine2d) Name() byte { return 'A' + byte(e2d.Idx) }

func (e2d *HwEngine2d) WriteDISPCNT(old, val uint32) {
	modLcd.WithFields(log.Fields{
		"name": string('A' + e2d.Idx),
		"val":  emu.Hex32(val),
	}).Info("write dispcnt")
}

func (e2d *HwEngine2d) drawChar16(y int, src []byte, dst gfx.Line, hflip bool) {
	src = src[y*4:]

	if !hflip {
		for x := 0; x < 4; x++ {
			p0, p1 := src[x]&0xF, src[x]>>4
			if p0 != 0 {
				p0 = (p0 << 4) | p0
				dst.Set8(0, p0)
			}
			if p1 != 0 {
				p1 = (p1 << 4) | p1
				dst.Set8(1, p1)
			}
			dst.Add8(2)
		}
	} else {
		for x := 3; x >= 0; x-- {
			p1, p0 := src[x]&0xF, src[x]>>4
			if p0 != 0 {
				p0 = (p0 << 4) | p0
				dst.Set8(0, p0)
			}
			if p1 != 0 {
				p1 = (p1 << 4) | p1
				dst.Set8(1, p1)
			}
			dst.Add8(2)
		}
	}
}

func (e2d *HwEngine2d) drawChar256(y int, src []byte, dst gfx.Line, hflip bool) {
	src = src[y*8:]
	if !hflip {
		for x := 0; x < 8; x++ {
			p0 := src[x]
			if p0 != 0 {
				dst.Set8(x, p0)
			}
		}
	} else {
		for x := 7; x >= 0; x-- {
			p0 := src[x]
			if p0 != 0 {
				dst.Set8(x, p0)
			}
		}
	}
}

func (e2d *HwEngine2d) DrawBG(ctx *gfx.LayerCtx, lidx int, y int) {
	regs := &e2d.bgregs[lidx]

	mapBase := int((*regs.Cnt>>8)&0xF) * 2 * 1024
	charBase := int((*regs.Cnt>>2)&0xF) * 16 * 1024
	if e2d.A() {
		mapBase += int((e2d.DispCnt.Value>>27)&7) * 64 * 1024
		charBase += int((e2d.DispCnt.Value>>24)&7) * 64 * 1024
	}
	tmap := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, mapBase)
	chars := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, charBase)
	onmask := uint32(1 << uint(8+lidx))

	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		if e2d.DispCnt.Value&onmask == 0 {
			y++
			continue
		}

		depth256 := (*regs.Cnt>>7)&1 != 0
		mapx := int(*regs.XOfs)
		mapy := (y + int(*regs.YOfs)) & 255

		mapYOff := 32 * (mapy / 8)
		for x := 0; x < cScreenWidth/8; x++ {
			mapx &= 255
			tile := tmap.Get16(mapYOff + (mapx / 8))

			// Decode tile
			tnum := int(tile & 1023)
			hflip := (tile>>10)&1 != 0
			vflip := (tile>>11)&1 != 0
			if tnum == 0 {
				continue
			}

			// Calculate tile line (and apply vertical flip)
			ty := mapy & 7
			if vflip {
				ty = 7 - ty
			}

			if depth256 {
				ch := chars.FetchPointer(tnum * 64)
				e2d.drawChar256(ty, ch, line, hflip)
			} else {
				ch := chars.FetchPointer(tnum * 32)
				e2d.drawChar16(ty, ch, line, hflip)
			}
			line.Add8(8)

			mapx += 8
		}

		y++
	}
}

var objWidth = []struct{ w, h int }{
	// square
	{1, 1}, {2, 2}, {4, 4}, {8, 8},
	// horizontal
	{2, 1}, {4, 1}, {4, 2}, {8, 4},
	// vertical
	{1, 2}, {1, 4}, {2, 4}, {4, 8},
}

func (e2d *HwEngine2d) DrawOBJ(ctx *gfx.LayerCtx, lidx int, sy int) {
	oam := Emu.Mem.OamRam[0x400*e2d.Idx : 0x400+0x400*e2d.Idx]
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

		// Go through the sprite list in reverse order, because an object with
		// lower index has HIGHER priority (so it gets drawn in front of all).
		// FIXME: This is actually a temporary hack because it will fail once we
		// emulate the sprite line limits; the correct solution would be
		// to go through in the correct order, but avoiding writing pixels
		// that have been already written to.
		for i := 127; i >= 0; i-- {
			a0, a1, a2 := le16(oam[i*8:]), le16(oam[i*8+2:]), le16(oam[i*8+4:])
			if a0&0x18 == 0x10 {
				continue
			}

			y := int(a0 & 0xff)
			x := int(a1 & 0x1ff)

			// Get the object size. The size is expressed in number of chars,
			// not pixels.
			sz := objWidth[((a0>>14)<<2)|(a1>>14)]
			tw, th := sz.w, sz.h

			// If the sprite is visible
			// FIXME: this doesn't handle wrapping yet
			if sy >= y && sy < y+th*8 && x < cScreenWidth {
				tilenum := int(a2 & 1023)
				depth256 := (a0>>13)&1 != 0
				hflip := (a1>>12)&1 != 0
				vflip := (a1>>13)&1 != 0

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

				// Compute the offset within VRAM of the current object (for
				// now, it's top-left pixel)
				vramOffset := tilenum * boundary

				// Adjust the offset to the beginning of the correct char row
				// within the object. ty is the number of char row being drawn.
				// This depends on the 1D vs 2D tile mapping in VRAM; 1D
				// mapping is what you would expect (tiles are arranged
				// linearly in memory), while 2D mapping means that tiles are
				// arrange in a grid.
				ty := y0 / 8
				if mapping1d {
					vramOffset += (tw * ty) * charSize
				} else {
					if depth256 {
						vramOffset += (16 * ty) * charSize
					} else {
						vramOffset += (32 * ty) * charSize
					}
				}

				// Now calculate the line being draw within the current char row
				y0 &= 7

				// Prepare initial src/dst pointer for drawing
				src := tiles.FetchPointer(vramOffset)
				dst := line
				dst.Add8(x)

				for j := 0; j < tw && x < cScreenWidth; j++ {
					tsrc := src[charSize*j:]
					if hflip {
						tsrc = src[charSize*(tw-j-1):]
					}

					if depth256 {
						e2d.drawChar256(y0, tsrc, dst, hflip)
					} else {
						e2d.drawChar16(y0, tsrc, dst, hflip)
					}
					dst.Add8(8)
					x += 8
				}
			}
		}
		sy++
	}
}

func (e2d *HwEngine2d) BeginFrame() {

	dispmode := (e2d.DispCnt.Value >> 16) & 3
	if e2d.B() {
		dispmode &= 1
	}

	switch dispmode {
	case 0:
		e2d.lm.Cfg.Mixer = e2dMixer_DisplayOff
	case 1:
		e2d.lm.Cfg.Mixer = e2dMixer_Normal
	default:
		modLcd.Fatalf("display mode not supported: %d", dispmode)
	}

	// Set the 4 BG layer priorities
	for i := 0; i < 4; i++ {
		pri := uint(*e2d.bgregs[i].Cnt & 3)
		e2d.lm.SetLayerPriority(i, pri)
	}
	e2d.lm.SetLayerPriority(4, 100) // put sprites always last in the mixer

	e2d.lm.BeginFrame()

	bgmode := e2d.DispCnt.Value & 7
	bg0on := (e2d.DispCnt.Value >> 8) & 1
	bg1on := (e2d.DispCnt.Value >> 9) & 1
	bg2on := (e2d.DispCnt.Value >> 10) & 1
	bg3on := (e2d.DispCnt.Value >> 11) & 1
	objon := (e2d.DispCnt.Value >> 12) & 1
	win0on := (e2d.DispCnt.Value >> 13) & 1
	win1on := (e2d.DispCnt.Value >> 14) & 1
	objwinon := (e2d.DispCnt.Value >> 15) & 1

	modLcd.Infof("%s: mode=%d bg=[%d,%d,%d,%d] obj=%d win=[%d,%d,%d]",
		string('A'+e2d.Idx), bgmode, bg0on, bg1on, bg2on, bg3on, objon, win0on, win1on, objwinon)
	modLcd.Infof("%s: scroll0=[%d,%d] scroll1=[%d,%d] scroll2=[%d,%d] scroll3=[%d,%d] size0=%d size3=%d",
		string('A'+e2d.Idx),
		e2d.Bg0XOfs.Value, e2d.Bg0YOfs.Value,
		e2d.Bg1XOfs.Value, e2d.Bg1YOfs.Value,
		e2d.Bg2XOfs.Value, e2d.Bg2YOfs.Value,
		e2d.Bg3XOfs.Value, e2d.Bg3YOfs.Value,
		e2d.Bg0Cnt.Value>>14, e2d.Bg3Cnt.Value>>13)
}

func (e2d *HwEngine2d) EndFrame() {
	e2d.lm.EndFrame()
}

func (e2d *HwEngine2d) BeginLine(screen gfx.Line) {
	e2d.lm.BeginLine(screen)
}

func (e2d *HwEngine2d) EndLine() {
	e2d.lm.EndLine()
}

func e2dMixer_DisplayOff(layers []uint32, ctx interface{}) uint32 {
	// When the display is off, the screen is white
	return 0xFFFFFF
}

func e2dMixer_Normal(layers []uint32, ctx interface{}) (res uint32) {
	l0 := uint8(layers[0])
	l1 := uint8(layers[1])
	l2 := uint8(layers[2])
	l3 := uint8(layers[3])
	s := uint8(layers[4])

	if s != 0 {
		return uint32(s) | uint32(s)<<8 | uint32(s)<<16
	}
	if l0 != 0 {
		return uint32(l0) | uint32(l0)<<8 | uint32(l0)<<16
	}
	if l1 != 0 {
		return uint32(l1) | uint32(l1)<<8 | uint32(l1)<<16
	}
	if l2 != 0 {
		return uint32(l2) | uint32(l2)<<8 | uint32(l2)<<16
	}
	if l3 != 0 {
		return uint32(l3) | uint32(l3)<<8 | uint32(l3)<<16
	}

	return 0
}
