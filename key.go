package main

import (
	"ndsemu/emu/hwio"
)

type HwKey struct {
	KeyIn    hwio.Reg16 `hwio:"bank=0,offset=0x0,reset=0x1FF,readonly"`
	KeyCnt   hwio.Reg16 `hwio:"bank=0,offset=0x2"`
	ExtKeyIn hwio.Reg16 `hwio:"bank=1,offset=0x6,reset=0x7F,readonly,rcb"`

	penDown bool
}

func NewHwKey() *HwKey {
	key := new(HwKey)
	hwio.MustInitRegs(key)
	return key
}

func (key *HwKey) SetPenDown(value bool) {
	key.penDown = value
}

func (key *HwKey) ReadEXTKEYIN(val uint16) uint16 {
	if key.penDown {
		val &^= 1 << 6
	}
	return val
}
