package main

import (
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modGx = log.NewModule("gx")

type GxCmd struct {
	code uint8
	parm uint32
}

type HwGeometry struct {
	GxFifo0 hwio.Reg32 `hwio:"bank=0,offset=0x0,readonly,wcb=WriteGXFIFO"`
	GxFifo1 hwio.Reg32 `hwio:"bank=0,offset=0x4,readonly,wcb=WriteGXFIFO"`
	GxFifo2 hwio.Reg32 `hwio:"bank=0,offset=0x8,readonly,wcb=WriteGXFIFO"`
	GxFifo3 hwio.Reg32 `hwio:"bank=0,offset=0xc,readonly,wcb=WriteGXFIFO"`
	GxFifo4 hwio.Reg32 `hwio:"bank=0,offset=0x10,readonly,wcb=WriteGXFIFO"`
	GxFifo5 hwio.Reg32 `hwio:"bank=0,offset=0x14,readonly,wcb=WriteGXFIFO"`
	GxFifo6 hwio.Reg32 `hwio:"bank=0,offset=0x18,readonly,wcb=WriteGXFIFO"`
	GxFifo7 hwio.Reg32 `hwio:"bank=0,offset=0x1c,readonly,wcb=WriteGXFIFO"`
	GxFifo8 hwio.Reg32 `hwio:"bank=0,offset=0x20,readonly,wcb=WriteGXFIFO"`
	GxFifo9 hwio.Reg32 `hwio:"bank=0,offset=0x24,readonly,wcb=WriteGXFIFO"`
	GxFifo10 hwio.Reg32 `hwio:"bank=0,offset=0x28,readonly,wcb=WriteGXFIFO"`
	GxFifo11 hwio.Reg32 `hwio:"bank=0,offset=0x2c,readonly,wcb=WriteGXFIFO"`
	GxFifo12 hwio.Reg32 `hwio:"bank=0,offset=0x30,readonly,wcb=WriteGXFIFO"`
	GxFifo13 hwio.Reg32 `hwio:"bank=0,offset=0x34,readonly,wcb=WriteGXFIFO"`
	GxFifo14 hwio.Reg32 `hwio:"bank=0,offset=0x38,readonly,wcb=WriteGXFIFO"`
	GxFifo15 hwio.Reg32 `hwio:"bank=0,offset=0x3c,readonly,wcb=WriteGXFIFO"`
	GxCmd10 hwio.Reg32 `hwio:"bank=0,offset=0x40,writeonly,wcb"`
	GxCmd16 hwio.Reg32 `hwio:"bank=0,offset=0x58,writeonly,wcb"`
	GxCmd50 hwio.Reg32 `hwio:"bank=0,offset=0x140,writeonly,wcb"`
	GxStat hwio.Reg32 `hwio:"bank=1,offset=0,rwmask=0xC0008000,rcb,wcb"`

	fifo []GxCmd
}

func NewHwGeometry() *HwGeometry {
	g := new(HwGeometry)
	hwio.MustInitRegs(g)
	return g
}

func (g *HwGeometry) ReadGXSTAT(val uint32) uint32 {
	if len(g.fifo) < 128 {
		val |= (1 << 25) // less-than-half-full
	}
	if len(g.fifo) == 0 {
		val |= (1 << 26) // empty
	}
	val |= (uint32(len(g.fifo)) & 0x1ff) << 16
	if len(g.fifo) != 0 {
		val |= (1 << 27) // busy bit
	}
	modGx.Infof("read GXSTAT: %08x", val)
	return val
}

func (g *HwGeometry) WriteGXSTAT(old, val uint32) {
	g.GxStat.Value &^= 0x8000
	if val&0x8000 != 0 {
		old &^= 0x8000
		// TODO: also reset matrix stat pointer
	}
	g.GxStat.Value |= old & 0x8000
	modGx.Infof("write GXSTAT: %08x", val)
	if val&0xc0000000 != 0 {
		modGx.Fatal("IRQ GX FIFO not emulated")
	}
}

func (g *HwGeometry) WriteGXFIFO(_, val uint32) {
	// TODO
}

func (g *HwGeometry) fifoPush(code uint8, parm uint32) {
	if len(g.fifo) < 256 {
		cmd := GxCmd{
			code: code,
			parm: parm,
		}
		g.fifo = append(g.fifo, cmd)
	} else {
		modGx.Errorf("attempt to push to full GX FIFO")
	}
}

func (g *HwGeometry) WriteGXCMD10(_, val uint32) {
	modGx.Infof("write GXCMD10: %08x", val)
	g.fifoPush(0x10, val)
}

func (g *HwGeometry) WriteGXCMD16(_, val uint32) {
	modGx.Infof("write GXCMD16: %08x", val)
	g.fifoPush(0x16, val)
}

func (g *HwGeometry) WriteGXCMD50(_, val uint32) {
	modGx.Infof("!!!!!! SWAP BUFFERS!!!!! write GXCMD50: %08x", val)
	g.fifo = nil
}
