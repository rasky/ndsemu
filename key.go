package main

import (
	"ndsemu/emu"
	"ndsemu/emu/hw"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

type HwKey struct {
	KeyIn    hwio.Reg16 `hwio:"bank=0,offset=0x0,reset=0x3FF,readonly,rcb"`
	KeyCnt   hwio.Reg16 `hwio:"bank=0,offset=0x2,wcb"`
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

func (key *HwKey) WriteKEYCNT(_, val uint16) {
	if val&(1<<14) != 0 {
		log.ModInput.Fatal("key interrupt not implemented")
	}
}

func (key *HwKey) ReadKEYIN(val uint16) uint16 {
	if KeyState[hw.SCANCODE_Z] != 0 {
		val &^= 1 << 0
	}
	if KeyState[hw.SCANCODE_X] != 0 {
		val &^= 1 << 1
	}
	if KeyState[hw.SCANCODE_RSHIFT] != 0 {
		val &^= 1 << 2
	}
	if KeyState[hw.SCANCODE_RETURN] != 0 {
		val &^= 1 << 3
	}
	if KeyState[hw.SCANCODE_RIGHT] != 0 {
		val &^= 1 << 4
	}
	if KeyState[hw.SCANCODE_LEFT] != 0 {
		val &^= 1 << 5
	}
	if KeyState[hw.SCANCODE_UP] != 0 {
		val &^= 1 << 6
	}
	if KeyState[hw.SCANCODE_DOWN] != 0 {
		val &^= 1 << 7
	}
	if KeyState[hw.SCANCODE_A] != 0 {
		val &^= 1 << 8
	}
	if KeyState[hw.SCANCODE_S] != 0 {
		val &^= 1 << 9
	}
	return val
}

func (key *HwKey) ReadEXTKEYIN(val uint16) uint16 {

	if KeyState[hw.SCANCODE_D] != 0 {
		val &^= 1 << 0
	}
	if KeyState[hw.SCANCODE_C] != 0 {
		val &^= 1 << 1
	}
	if key.penDown {
		val &^= 1 << 6
	}
	log.ModInput.WithField("val", emu.Hex16(val)).Info("read EXTKEYIN")
	return val
}
