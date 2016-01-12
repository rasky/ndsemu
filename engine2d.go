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

func (r *bgRegs) priority() uint16 { return (*r.Cnt & 3) }
func (r *bgRegs) depth256() bool   { return (*r.Cnt>>7)&1 != 0 }

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

	bgregs    [4]bgRegs
	mc        *HwMemoryController
	lineBuf   [4 * (cScreenWidth + 16)]byte
	lm        gfx.LayerManager
	dispmode  int
	modeTable [4]struct {
		BeginFrame func()
		EndFrame   func()
		BeginLine  func(y int, screen gfx.Line)
		EndLine    func(y int)
	}
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

	// Initialize the mode table, used to implement the four different
	// display modes
	e2d.modeTable[0].BeginFrame = e2d.Mode0_BeginFrame
	e2d.modeTable[0].EndFrame = e2d.Mode0_EndFrame
	e2d.modeTable[0].BeginLine = e2d.Mode0_BeginLine
	e2d.modeTable[0].EndLine = e2d.Mode0_EndLine

	e2d.modeTable[1].BeginFrame = e2d.Mode1_BeginFrame
	e2d.modeTable[1].EndFrame = e2d.Mode1_EndFrame
	e2d.modeTable[1].BeginLine = e2d.Mode1_BeginLine
	e2d.modeTable[1].EndLine = e2d.Mode1_EndLine

	e2d.modeTable[2].BeginFrame = e2d.Mode2_BeginFrame
	e2d.modeTable[2].EndFrame = e2d.Mode2_EndFrame
	e2d.modeTable[2].BeginLine = e2d.Mode2_BeginLine
	e2d.modeTable[2].EndLine = e2d.Mode2_EndLine

	e2d.modeTable[3].BeginFrame = e2d.Mode3_BeginFrame
	e2d.modeTable[3].EndFrame = e2d.Mode3_EndFrame
	e2d.modeTable[3].BeginLine = e2d.Mode3_BeginLine
	e2d.modeTable[3].EndLine = e2d.Mode3_EndLine

	// Initialize layer manager (used in mode1)
	e2d.lm.Cfg = gfx.LayerManagerConfig{
		Width:          cScreenWidth,
		Height:         cScreenHeight,
		ScreenBpp:      4,
		LayerBpp:       2,
		OverflowPixels: 8,
		Mixer:          e2dMixer_Normal,
	}

	// Background layers
	e2d.lm.AddLayer(gfx.LayerFunc{Func: e2d.DrawBG})
	e2d.lm.AddLayer(gfx.LayerFunc{Func: e2d.DrawBG})
	e2d.lm.AddLayer(gfx.LayerFunc{Func: e2d.DrawBG})
	e2d.lm.AddLayer(gfx.LayerFunc{Func: e2d.DrawBG})

	// Sprites layer
	e2d.lm.AddLayer(gfx.LayerFunc{Func: e2d.DrawOBJ})
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

// A pixel in a layer of the layer manager. It is composed as follows:
//   Bits 0-7: color index in the palette
//   Bits 8-13: palette number
//   Bits 14-15: priority
type LayerPixel uint16

func (p LayerPixel) Color() uint16     { return uint16(p & 0xFF) }
func (p LayerPixel) Palette() uint16   { return uint16(p>>8) & 0x3F }
func (p LayerPixel) Priority() uint16  { return uint16(p >> 14) }
func (p LayerPixel) Transparent() bool { return p.Color() == 0 }

func (e2d *HwEngine2d) drawChar16(y int, src []byte, dst gfx.Line, hflip bool, pri uint16) {
	src = src[y*4:]
	pri <<= 14
	if !hflip {
		for x := 0; x < 4; x++ {
			p0, p1 := uint16(src[x]&0xF), uint16(src[x]>>4)
			if p0 != 0 {
				p0 = (p0 << 4) | p0
				dst.Set16(0, p0|pri)
			}
			if p1 != 0 {
				p1 = (p1 << 4) | p1
				dst.Set16(1, p1|pri)
			}
			dst.Add16(2)
		}
	} else {
		for x := 3; x >= 0; x-- {
			p1, p0 := uint16(src[x]&0xF), uint16(src[x]>>4)
			if p0 != 0 {
				p0 = (p0 << 4) | p0
				dst.Set16(0, p0|pri)
			}
			if p1 != 0 {
				p1 = (p1 << 4) | p1
				dst.Set16(1, p1|pri)
			}
			dst.Add16(2)
		}
	}
}

func (e2d *HwEngine2d) drawChar256(y int, src []byte, dst gfx.Line, hflip bool, pri uint16) {
	src = src[y*8:]
	pri <<= 14
	if !hflip {
		for x := 0; x < 8; x++ {
			p0 := uint16(src[x])
			if p0 != 0 {
				dst.Set16(x, p0|pri)
			}
		}
	} else {
		for x := 7; x >= 0; x-- {
			p0 := uint16(src[x])
			if p0 != 0 {
				dst.Set16(x, p0|pri)
			}
		}
	}
}

func (e2d *HwEngine2d) DrawBG(ctx *gfx.LayerCtx, lidx int, y int) {
	regs := &e2d.bgregs[lidx]

	mapBase := int((*regs.Cnt>>8)&0x1F) * 2 * 1024
	charBase := int((*regs.Cnt>>2)&0xF) * 16 * 1024

	if e2d.A() {
		mapBase += int((e2d.DispCnt.Value>>27)&7) * 64 * 1024
		charBase += int((e2d.DispCnt.Value>>24)&7) * 64 * 1024
	}
	tmap := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, mapBase)
	chars := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBG, charBase)
	onmask := uint32(1 << uint(8+lidx))

	// true if this is layer BG0 in engine A
	bg0a := lidx == 0 && e2d.A()

	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		if e2d.DispCnt.Value&onmask == 0 {
			y++
			continue
		}

		// If this is layer BG0 in engine A, and 3D is activated, we shouldn't
		// draw anything here because the layer is replaced by the 3D layer
		if bg0a && (e2d.DispCnt.Value>>3)&1 != 0 {
			y++
			continue
		}

		pri := regs.priority()
		depth256 := regs.depth256()
		mapx := int(*regs.XOfs)
		mapy := (y + int(*regs.YOfs)) & 255

		line.Add16(-(mapx & 7))
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
				e2d.drawChar256(ty, ch, line, hflip, pri)
			} else {
				ch := chars.FetchPointer(tnum * 32)
				e2d.drawChar16(ty, ch, line, hflip, pri)
			}
			line.Add16(8)

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
			if a0&0x300 == 0x200 {
				continue
			}

			const XMask = 0x1FF
			const YMask = 0xFF

			x := int(a1 & XMask)
			y := int(a0 & YMask)
			if x >= cScreenWidth {
				x -= XMask + 1
			}

			// Get the object size. The size is expressed in number of chars,
			// not pixels.
			sz := objWidth[((a0>>14)<<2)|(a1>>14)]
			tw, th := sz.w, sz.h

			// If the sprite is visible
			// FIXME: this doesn't handle wrapping yet
			if sy >= y && sy < (y+th*8)&YMask && (x < cScreenWidth && (x+tw*8) >= 0) {
				tilenum := int(a2 & 1023)
				depth256 := (a0>>13)&1 != 0
				hflip := (a1>>12)&1 != 0
				vflip := (a1>>13)&1 != 0
				pri := (a2 >> 10) & 3

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
							e2d.drawChar256(y0, tsrc, dst, hflip, pri)
						} else {
							e2d.drawChar16(y0, tsrc, dst, hflip, pri)
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

/************************************************
 * Display Modes: basic dispatch
 ************************************************/

func (e2d *HwEngine2d) BeginFrame() {
	// Read current display mode once per frame (do not switch between
	// display modes within a frame)
	e2d.dispmode = int((e2d.DispCnt.Value >> 16) & 3)
	if e2d.B() {
		e2d.dispmode &= 1
	}
	e2d.modeTable[e2d.dispmode].BeginFrame()
}
func (e2d *HwEngine2d) EndFrame()                   { e2d.modeTable[e2d.dispmode].EndFrame() }
func (e2d *HwEngine2d) BeginLine(y int, s gfx.Line) { e2d.modeTable[e2d.dispmode].BeginLine(y, s) }
func (e2d *HwEngine2d) EndLine(y int)               { e2d.modeTable[e2d.dispmode].EndLine(y) }

/************************************************
 * Display Mode 0: display off
 ************************************************/

func (e2d *HwEngine2d) Mode0_BeginFrame() {}
func (e2d *HwEngine2d) Mode0_EndFrame()   {}

func (e2d *HwEngine2d) Mode0_BeginLine(y int, screen gfx.Line) {
	for x := 0; x < cScreenWidth; x++ {
		// Display off -> draw white
		screen.Set32(x, 0xFFFFFF)
	}
}
func (e2d *HwEngine2d) Mode0_EndLine(y int) {}

/************************************************
 * Display Mode 1: BG & OBJ layers
 ************************************************/

func (e2d *HwEngine2d) Mode1_BeginFrame() {
	// Set the 4 BG layer priorities
	for i := 0; i < 4; i++ {
		pri := uint(e2d.bgregs[i].priority())
		e2d.lm.SetLayerPriority(i, pri)
	}
	e2d.lm.SetLayerPriority(4, 100) // put sprites always last in the mixer

	e2d.lm.BeginFrame()

	bgmode := e2d.DispCnt.Value & 7
	bg3d := (e2d.DispCnt.Value >> 3) & 1
	bg0on := (e2d.DispCnt.Value >> 8) & 1
	bg1on := (e2d.DispCnt.Value >> 9) & 1
	bg2on := (e2d.DispCnt.Value >> 10) & 1
	bg3on := (e2d.DispCnt.Value >> 11) & 1
	objon := (e2d.DispCnt.Value >> 12) & 1
	win0on := (e2d.DispCnt.Value >> 13) & 1
	win1on := (e2d.DispCnt.Value >> 14) & 1
	objwinon := (e2d.DispCnt.Value >> 15) & 1

	modLcd.Infof("%s: mode=%d bg=[%d,%d,%d,%d] obj=%d win=[%d,%d,%d] 3d=%d",
		string('A'+e2d.Idx), bgmode, bg0on, bg1on, bg2on, bg3on, objon, win0on, win1on, objwinon, bg3d)
	modLcd.Infof("%s: scroll0=[%d,%d] scroll1=[%d,%d] scroll2=[%d,%d] scroll3=[%d,%d] size0=%d size3=%d",
		string('A'+e2d.Idx),
		e2d.Bg0XOfs.Value, e2d.Bg0YOfs.Value,
		e2d.Bg1XOfs.Value, e2d.Bg1YOfs.Value,
		e2d.Bg2XOfs.Value, e2d.Bg2YOfs.Value,
		e2d.Bg3XOfs.Value, e2d.Bg3YOfs.Value,
		e2d.Bg0Cnt.Value>>14, e2d.Bg3Cnt.Value>>13)
}

func (e2d *HwEngine2d) Mode1_EndFrame() {
	e2d.lm.EndFrame()
}

func (e2d *HwEngine2d) Mode1_BeginLine(y int, screen gfx.Line) {
	e2d.lm.BeginLine(screen)
}

func (e2d *HwEngine2d) Mode1_EndLine(y int) {
	e2d.lm.EndLine()
}

func e2dMixer_Normal(layers []uint32, ctx interface{}) (res uint32) {
	var pix, objpix, bgpix LayerPixel

	// Extract the layers. They've been already sorted in priority order,
	// so the first layer with a non-transparent pixel is the one that gets
	// drawn.
	bgpix = LayerPixel(layers[0])
	if !bgpix.Transparent() {
		goto checkobj
	}
	bgpix = LayerPixel(layers[1])
	if !bgpix.Transparent() {
		goto checkobj
	}
	bgpix = LayerPixel(layers[2])
	if !bgpix.Transparent() {
		goto checkobj
	}
	bgpix = LayerPixel(layers[3])
	if !bgpix.Transparent() {
		goto checkobj
	}

	// No bglayer was drawn here, so draw the obj directly
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() {
		pix = objpix
		goto draw
	}

	// objlayer is also transparent. Draw backdrop (for now, green)
	return 0x00FF00

checkobj:
	// Check if there is an object pixel here, and if it's priority is higher
	// (or equal, because with equal priorities, objects win)
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() && objpix.Priority() <= bgpix.Priority() {
		pix = objpix
		goto draw
	} else {
		pix = bgpix
		goto draw
	}

draw:
	c := uint32(pix.Color())
	return c<<16 | c<<8 | c
}

/************************************************
 * Display Mode 2: VRAM display
 ************************************************/

func (e2d *HwEngine2d) Mode2_BeginFrame() {}
func (e2d *HwEngine2d) Mode2_EndFrame()   {}

func (e2d *HwEngine2d) Mode2_BeginLine(y int, screen gfx.Line) {
	block := (e2d.DispCnt.Value >> 18) & 3
	vram := e2d.mc.vram[block][y*cScreenWidth*2:]
	for x := 0; x < cScreenWidth; x++ {
		pix := le16(vram[x*2:])
		r := uint32(pix & 0x1F)
		g := uint32((pix >> 5) & 0x1F)
		b := uint32((pix >> 10) & 0x1F)
		r = (r << 3) | (r >> 2)
		g = (g << 3) | (g >> 2)
		b = (b << 3) | (b >> 2)
		screen.Set32(x, r|g<<8|b<<16)
	}
}
func (e2d *HwEngine2d) Mode2_EndLine(y int) {}

/************************************************
 * Display Mode 3: Main memory display
 ************************************************/

func (e2d *HwEngine2d) Mode3_BeginFrame() {
	modLcd.Warn("mode 3 not implemented")
}
func (e2d *HwEngine2d) Mode3_EndFrame() {}

func (e2d *HwEngine2d) Mode3_BeginLine(y int, screen gfx.Line) {
	for x := 0; x < cScreenWidth; x++ {
		screen.Set32(x, 0xFF0000)
	}
}
func (e2d *HwEngine2d) Mode3_EndLine(y int) {}
