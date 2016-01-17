package main

import (
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modGx = log.NewModule("gx")

type HwGeometry struct {
	GxStat hwio.Reg32 `hwio:"bank=1,offset=0,rwmask=0xC0008000,rcb,wcb"`
}

func NewHwGeometry() *HwGeometry {
	g := new(HwGeometry)
	hwio.MustInitRegs(g)
	return g
}

func (g *HwGeometry) ReadGXSTAT(val uint32) uint32 {
	val |= (1 << 25) // less-than-half-full
	val |= (1 << 26) // empty
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
	if val&0x80000000 != 0 {
		nds9.Irq.Raise(1 << 21)
	}
}
