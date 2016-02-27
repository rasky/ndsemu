package main

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hw"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

const (
	cScreenWidth  = 256
	cScreenHeight = 192
)

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
	MBright  hwio.Reg32 `hwio:"offset=0x6C,rwmask=0xC01F,wcb"`

	// bank 1: registers available only on display A
	DispCapCnt   hwio.Reg32 `hwio:"bank=1,offset=0x64"`
	DispMMemFifo hwio.Reg32 `hwio:"bank=1,offset=0x68,wcb"`

	bgregs    [4]bgRegs
	bgmodes   [4]BgMode
	mc        *HwMemoryController
	lineBuf   [4 * (cScreenWidth + 16)]byte
	lm        gfx.LayerManager
	dispmode  int
	curline   int
	curscreen gfx.Line
	modeTable [4]struct {
		BeginFrame func()
		EndFrame   func()
		BeginLine  func(y int, screen gfx.Line)
		EndLine    func(y int)
	}
	dispcap struct {
		Enabled       bool
		Bank          int
		Offset        uint32
		Width, Height int
	}

	// Master brightness conversion. Depending on the master brightness
	// register, we precalculate highlighted/shadowed colors for the 32
	// shades; moreover, to save time on the mixer, we also already shift
	// R,G,B of the correct amount, so that the mixer can just OR them
	// together.
	masterBrightChanged bool
	masterBrightR       [32]uint32
	masterBrightG       [32]uint32
	masterBrightB       [32]uint32

	bgPal     []byte
	objPal    []byte
	bgExtPals [4][]byte
	objExtPal []byte
}

func NewHwEngine2d(idx int, mc *HwMemoryController) *HwEngine2d {
	e2d := new(HwEngine2d)
	hwio.MustInitRegs(e2d)
	e2d.Idx = idx
	e2d.mc = mc
	e2d.masterBrightChanged = true // force initial table calculation

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
		MixerCtx:       e2d,
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

func (e2d *HwEngine2d) WriteDISPMMEMFIFO(old, val uint32) {
	if val != 0 {
		modLcd.Fatalf("unimplemented DISP MMEM FIFO")
	}
}

func (e2d *HwEngine2d) WriteMBRIGHT(old, val uint32) {
	if old != val {
		e2d.masterBrightChanged = true
	}
}

func (e2d *HwEngine2d) updateMasterBrightTable() {
	// Setup master brightness lookup tables. Do this for every line just to be safe
	brightMode := (e2d.MBright.Value >> 14) & 3
	brightFactor := int32(e2d.MBright.Value & 0x1F)
	if brightFactor > 16 {
		brightFactor = 16
	}

	for i := int32(0); i < int32(32); i++ {
		// Expand to 6-bit color using the hardware formula.
		// It looks like internally the mixer handle 6-bit colors;
		// the 3D layer also outputs 6-bit colors, but we currently
		// drop the lower bit, so the mixer always receives 5-bit colors.
		c := i
		if c > 0 {
			c = c*2 + 1
		}

		switch brightMode {
		case 1: // brightness up
			c += (63 - c) * brightFactor / 16
			if c > 63 {
				c = 63
			}
		case 2: // brightness down
			c -= c * brightFactor / 16
			if c < 0 {
				c = 0
			}
		}

		// Expand from 6-bits to 8-bits
		c = c<<2 | c>>4

		// Fill up the table with the 3 masks
		e2d.masterBrightR[i] = uint32(c)
		e2d.masterBrightG[i] = uint32(c) << 8
		e2d.masterBrightB[i] = uint32(c) << 16
	}
}

// A pixel in a layer of the layer manager. It is composed as follows:
//   Bits 0-11: color index in the palette
//   Bit 12: unused
//   Bit 13: set if the pixel uses the extended palette for its layer (either obj or bg)
//   Bits 14-15: priority

//   Bits 0-11: color index in the palette
//   Bit 12: set if the pixel uses the extended palette for its layer (either obj or bg)
//   Bit 13-14: priority
//   Bit 15: direct color
type LayerPixel uint16

func (p LayerPixel) Color() uint16       { return uint16(p & 0xFFF) }
func (p LayerPixel) ExtPal() bool        { return uint16(p&(1<<12)) != 0 }
func (p LayerPixel) Priority() uint16    { return uint16(p>>13) & 3 }
func (p LayerPixel) Direct() bool        { return int16(p) < 0 }
func (p LayerPixel) DirectColor() uint16 { return uint16(p & 0x7FFF) }
func (p LayerPixel) Transparent() bool   { return p == 0 }

func (e2d *HwEngine2d) drawChar16(y int, src []byte, dst gfx.Line, hflip bool, pri uint16, pal uint16, extpal bool) {
	src = src[y*4:]
	attrs := pri<<13 | pal<<4
	if extpal {
		attrs |= (1 << 12)
	}
	if !hflip {
		for x := 0; x < 4; x++ {
			p0, p1 := uint16(src[x]&0xF), uint16(src[x]>>4)
			if p0 != 0 {
				// p0 = (p0 << 4) | p0
				dst.Set16(0, p0|attrs)
			}
			if p1 != 0 {
				// p1 = (p1 << 4) | p1
				dst.Set16(1, p1|attrs)
			}
			dst.Add16(2)
		}
	} else {
		for x := 3; x >= 0; x-- {
			p1, p0 := uint16(src[x]&0xF), uint16(src[x]>>4)
			if p0 != 0 {
				// p0 = (p0 << 4) | p0
				dst.Set16(0, p0|attrs)
			}
			if p1 != 0 {
				// p1 = (p1 << 4) | p1
				dst.Set16(1, p1|attrs)
			}
			dst.Add16(2)
		}
	}
}

func (e2d *HwEngine2d) drawChar256(y int, src []byte, dst gfx.Line, hflip bool, pri uint16, pal uint16, extpal bool) {
	src = src[y*8:]
	attrs := pri<<13 | pal<<8
	if extpal {
		attrs |= (1 << 12)
	}
	if !hflip {
		for x := 0; x < 8; x++ {
			p0 := uint16(src[x])
			if p0 != 0 {
				dst.Set16(0, p0|attrs)
			}
			dst.Add16(1)
		}
	} else {
		for x := 7; x >= 0; x-- {
			p0 := uint16(src[x])
			if p0 != 0 {
				dst.Set16(0, p0|attrs)
			}
			dst.Add16(1)
		}
	}
}

func (e2d *HwEngine2d) DrawBGAffine(ctx *gfx.LayerCtx, lidx int, y int) {
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

	if y != 0 {
		panic("unimplemented initial line not zero on affine plane")
	}

	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		if e2d.DispCnt.Value&onmask == 0 || KeyState[hw.SCANCODE_1+lidx] != 0 {
			y++
			continue
		}
		if (e2d.A() && KeyState[hw.SCANCODE_9] != 0) || (e2d.B() && KeyState[hw.SCANCODE_0] != 0) {
			y++
			continue
		}

		// Check if we are in extended palette mode (more palettes available for
		// 256-color tiles).
		useExtPal := (e2d.DispCnt.Value & (1 << 30)) != 0

		pri := regs.priority()

		size := 128 << ((*regs.Cnt >> 14) & 3)

		mapx := int32(*regs.PX<<4) >> 4
		mapy := int32(*regs.PY<<4) >> 4

		dx := int32(*regs.PA<<4) >> 4
		dy := int32(*regs.PC<<4) >> 4

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

		// Update the mapx/mapy register for next line. We write the value back
		// to the register, and we re-read it at beginning of next line (after
		// hblank), so that we allow CPU to mess with it in the blank period.
		mapx = int32(*regs.PX<<4) >> 4
		mapy = int32(*regs.PY<<4) >> 4
		dmx := int32(*regs.PB<<4) >> 4
		dmy := int32(*regs.PD<<4) >> 4
		*regs.PX = uint32(mapx + dmx)
		*regs.PY = uint32(mapy + dmy)

		y++
	}
}

func (e2d *HwEngine2d) DrawBG(ctx *gfx.LayerCtx, lidx int, y int) {
	regs := &e2d.bgregs[lidx]

	mapBase := int((*regs.Cnt>>8)&0x1F) * 2 * 1024
	charBase := int((*regs.Cnt>>2)&0xF) * 16 * 1024

	if lidx == 2 && e2d.B() {
		modLcd.Infof("B2: map:%x, char:%x, x:%x, y:%x, 8bpp:%v", mapBase, charBase, *regs.XOfs, *regs.YOfs, regs.depth256())
	}

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

	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		if e2d.DispCnt.Value&onmask == 0 || KeyState[hw.SCANCODE_1+lidx] != 0 {
			y++
			continue
		}
		if (e2d.A() && KeyState[hw.SCANCODE_9] != 0) || (e2d.B() && KeyState[hw.SCANCODE_0] != 0) {
			y++
			continue
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

		line.Add16(-(mapx & 7))
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

			if depth256 {
				ch := chars.FetchPointer(tnum * 64)
				// 256-color tiles only have one palette in normal (GBA) mode, but
				// can have multiple palettes in extended palette mode.
				// So we ignore the palette number if extended palette is disabled
				// (it should be already zero, but better safe than sorry)
				if !useExtPal {
					pal = 0
				}
				e2d.drawChar256(ty, ch, line, hflip, pri, pal, useExtPal)
			} else {
				ch := chars.FetchPointer(tnum * 32)
				// 16-color tiles don't use extended palettes, so we always pass false
				// to the drawChar16() function
				e2d.drawChar16(ty, ch, line, hflip, pri, pal, false)
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

/************************************************
 * Display Modes: basic dispatch
 ************************************************/

func (e2d *HwEngine2d) BeginFrame() {

	// Check if display capture is activated
	if e2d.DispCapCnt.Value&(1<<31) != 0 {
		source := (e2d.DispCapCnt.Value >> 29) & 3
		srca := (e2d.DispCapCnt.Value >> 24) & 1
		srcb := (e2d.DispCapCnt.Value >> 25) & 1

		modLcd.WithDelayedFields(func() log.Fields {
			return log.Fields{
				"src": source,
				"sa":  srca,
				"sb":  srcb,
			}
		}).Infof("Capture activated")
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
		vram := Emu.Hw.Mc.vram[e2d.dispcap.Bank]
		vram = vram[e2d.dispcap.Offset:]
		capbuf := gfx.NewLine(vram)
		for ; i < e2d.dispcap.Width; i++ {
			pix := screen.Get32(i)
			capbuf.Set16(i, uint16(pix))
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

type BgMode int

//go:generate stringer -type BgMode
const (
	BgModeText BgMode = iota
	BgModeAffine
	BgModeAffineMap16
	BgModeAffineBitmap
	BgModeAffineBitmapDirect
	BgModeLarge
	BgMode3D
)

func (e2d *HwEngine2d) Mode1_BeginFrame() {
	// Set the 4 BG layer priorities
	for i := 0; i < 4; i++ {
		pri := uint(e2d.bgregs[i].priority())
		e2d.lm.SetLayerPriority(i, pri)
	}
	e2d.lm.SetLayerPriority(4, 100) // put sprites always last in the mixer

	bgmode := e2d.DispCnt.Value & 7
	bg3d := (e2d.DispCnt.Value>>3)&1 != 0

	// Switch bg0 between text and 3d layer, depending on setting in DISPCNT
	// Notice also that in bgmode 6, layer 0 is always 3d
	if e2d.A() {
		if bg3d || bgmode == 6 {
			e2d.Mode1_setBgMode(0, BgMode3D)
		} else {
			e2d.Mode1_setBgMode(0, BgModeText)
		}
	}

	// Bg1 is always text
	// e2d.Mode1_setBgMode(1, BgModeText)

	// Change bg2/bg3 depending on mode
	for idx := 2; idx <= 3; idx++ {
		switch bgmode {
		case 0, 1, 3:
			e2d.Mode1_setBgMode(idx, BgModeText)
		case 2, 4:
			e2d.Mode1_setBgMode(idx, BgModeAffine)
		case 5:
			// Extended mode: need to check other bits
			if *e2d.bgregs[idx].Cnt&(1<<7) == 0 {
				e2d.Mode1_setBgMode(idx, BgModeAffineMap16)
			} else if *e2d.bgregs[idx].Cnt&(1<<2) == 0 {
				e2d.Mode1_setBgMode(idx, BgModeAffineBitmap)
			} else {
				e2d.Mode1_setBgMode(idx, BgModeAffineBitmapDirect)
			}
		case 6:
			e2d.Mode1_setBgMode(idx, BgModeLarge)
		}
	}
	e2d.lm.BeginFrame()

	bg0on := (e2d.DispCnt.Value >> 8) & 1
	bg1on := (e2d.DispCnt.Value >> 9) & 1
	bg2on := (e2d.DispCnt.Value >> 10) & 1
	bg3on := (e2d.DispCnt.Value >> 11) & 1

	objon := (e2d.DispCnt.Value >> 12) & 1
	win0on := (e2d.DispCnt.Value >> 13) & 1
	win1on := (e2d.DispCnt.Value >> 14) & 1
	objwinon := (e2d.DispCnt.Value >> 15) & 1

	modLcd.Infof("%s: modes=%v bg=[%d,%d,%d,%d] obj=%d win=[%d,%d,%d]",
		string('A'+e2d.Idx), e2d.bgmodes, bg0on, bg1on, bg2on, bg3on, objon, win0on, win1on, objwinon)
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
	e2d.bgPal = Emu.Mem.PaletteRam[e2d.Idx*1024 : e2d.Idx*1024+512]
	e2d.objPal = Emu.Mem.PaletteRam[e2d.Idx*1024+512 : e2d.Idx*1024+512+512]

	// Fetch direct pointers to extended palettes (both bg and obj). These
	// will be used by the mixer function to access the actual colors.
	//
	// Extended palettes are active if bit 30 (bg) and/or 31 (obj) in
	// DISPCNT are set, and of course the appropriate VRAM banks need to
	// be mapped. In any case, VramLinearBank() returnes an empty memory
	// area if the banks are unmapped, and we can even try to use it
	// (and display all black). This is possibly consistent with what NDS
	// does if the extended palettes are used without being properly
	// configured.
	bgextpal := Emu.Hw.Mc.VramLinearBank(e2d.Idx, VramLinearBGExtPal, 0)
	for i := range e2d.bgExtPals {
		// Compute the BG Ext Palette slot used by each bg layer. Normally,
		// BG0 uses Slot 0, BG3 uses Slot 3, etc. but BG0 and BG1 can optionally
		// use a different slot (depending on bit 13 of BGxCNT register)
		slotnum := i
		if i == 0 && e2d.Bg0Cnt.Value&(1<<13) != 0 {
			slotnum = 2
		}
		if i == 1 && e2d.Bg1Cnt.Value&(1<<13) != 0 {
			slotnum = 3
		}

		e2d.bgExtPals[i] = bgextpal.FetchPointer(8 * 1024 * slotnum)
	}

	objextpal := Emu.Hw.Mc.VramLinearBank(e2d.Idx, VramLinearOBJExtPal, 0)
	e2d.objExtPal = objextpal.FetchPointer(0)

	e2d.lm.BeginLine(screen)
}

func (e2d *HwEngine2d) Mode1_EndLine(y int) {
	e2d.lm.EndLine()
}

func (e2d *HwEngine2d) Mode1_setBgMode(lidx int, mode BgMode) {
	if e2d.bgmodes[lidx] == mode {
		return
	}

	e2d.bgmodes[lidx] = mode
	switch mode {
	case BgModeText:
		e2d.lm.ChangeLayer(lidx, gfx.LayerFunc{Func: e2d.DrawBG})
	case BgMode3D:
		e2d.lm.ChangeLayer(lidx, gfx.LayerFunc{Func: Emu.Hw.E3d.Draw3D})
	case BgModeAffineMap16, BgModeAffineBitmap, BgModeAffineBitmapDirect:
		e2d.lm.ChangeLayer(lidx, gfx.LayerFunc{Func: e2d.DrawBGAffine})
	default:
		panic(fmt.Errorf("bgmode %v not implemented", mode))
	}
}

func e2dMixer_Normal(layers []uint32, ctx interface{}) (res uint32) {
	var objpix, bgpix LayerPixel
	var pix uint16
	var cram []uint8
	var c16 uint16
	e2d := ctx.(*HwEngine2d)

	// Extract the layers. They've been already sorted in priority order,
	// so the first layer with a non-transparent pixel is the one that gets
	// drawn.
	bgpix = LayerPixel(layers[0])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[0]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}
	bgpix = LayerPixel(layers[1])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[1]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}
	bgpix = LayerPixel(layers[2])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[2]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}
	bgpix = LayerPixel(layers[3])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[3]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}

	// No bglayer was drawn here, so see if there is at least an obj, in which
	// case we draw it directly
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() {
		pix = objpix.Color()
		if objpix.ExtPal() {
			cram = e2d.objExtPal
		} else {
			cram = e2d.objPal
		}
		goto lookup
	}

	// No objlayer, and no bglayer. Draw the backdrop.
	// pix = 0
	cram = e2d.bgPal
	goto lookup

checkobj:
	// We found a bg pixel; now check if there is an object pixel here: if so,
	// we need to check the priority to choose between bg and obj which pixel
	// to draw (if the priorities are equal, objects win)
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() && objpix.Priority() <= bgpix.Priority() {
		pix = objpix.Color()
		if objpix.ExtPal() {
			cram = e2d.objExtPal
		} else {
			cram = e2d.objPal
		}
	} else {
		if bgpix.Direct() {
			c16 = bgpix.DirectColor()
			goto draw
		}
		pix = bgpix.Color()
	}

lookup:
	c16 = emu.Read16LE(cram[pix*2:])
draw:
	// Just return the 16-bit value for now, the post-processing
	// function will take care of the last step
	return uint32(c16)
}

/************************************************
 * Display Mode 2: VRAM display
 ************************************************/

func (e2d *HwEngine2d) Mode2_BeginFrame() {
	block := (e2d.DispCnt.Value >> 18) & 3
	modLcd.Infof("%s: mode=VRAM-Display bank:%s",
		string('A'+e2d.Idx), string('A'+block))

}
func (e2d *HwEngine2d) Mode2_EndFrame() {}

func (e2d *HwEngine2d) Mode2_BeginLine(y int, screen gfx.Line) {
	block := (e2d.DispCnt.Value >> 18) & 3
	vram := e2d.mc.vram[block][y*cScreenWidth*2:]
	for x := 0; x < cScreenWidth; x++ {
		pix := emu.Read16LE(vram[x*2:])
		screen.Set32(x, uint32(pix))
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
		screen.Set32(x, 0x0000FF)
	}
}
func (e2d *HwEngine2d) Mode3_EndLine(y int) {}
