package main

import (
	"ndsemu/emu"
	"ndsemu/emu/hwio"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwEngine2d struct {
	Idx     int
	DispCnt hwio.Reg32 `hwio:"offset=0x0,wcb"`
}

func NewHwEngine2d(idx int) *HwEngine2d {
	e2d := new(HwEngine2d)
	e2d.Idx = idx
	hwio.MustInitRegs(e2d)
	return e2d
}

func (e2d *HwEngine2d) A() bool { return e2d.Idx == 0 }
func (e2d *HwEngine2d) B() bool { return e2d.Idx != 0 }

func (e2d *HwEngine2d) WriteDISPCNT(old, val uint32) {
	log.WithFields(log.Fields{
		"name": string('A' + e2d.Idx),
		"val":  emu.Hex32(val),
	}).Info("[lcd] write dispcnt")
}
