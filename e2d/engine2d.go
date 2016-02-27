package e2d

import (
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
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
	mc        MemoryController
	lineBuf   [4 * (cScreenWidth + 16)]byte
	lm        gfx.LayerManager
	l3d       gfx.Layer
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

func NewHwEngine2d(idx int, mc MemoryController, l3d gfx.Layer) *HwEngine2d {
	e2d := new(HwEngine2d)
	hwio.MustInitRegs(e2d)
	e2d.Idx = idx
	e2d.mc = mc
	e2d.l3d = l3d
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
