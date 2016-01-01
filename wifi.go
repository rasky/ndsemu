package main

import (
	"math/rand"
	"ndsemu/emu/hwio"
)

type HwWifi struct {
	bbRegWritable [256]bool
	bbRegs        [256]uint8
	BaseBandCnt   hwio.Reg16 `hwio:"offset=0x158,wcb"`
	BaseBandWrite hwio.Reg16 `hwio:"offset=0x15A,writeonly"`
	BaseBandRead  hwio.Reg16 `hwio:"offset=0x15C,readonly"`
	BaseBandBusy  hwio.Reg16 `hwio:"offset=0x15E,readonly"`
	BaseBandMode  hwio.Reg16 `hwio:"offset=0x160"`
	BaseBandPower hwio.Reg16 `hwio:"offset=0x168"`

	Random hwio.Reg16 `hwio:"offset=0x044,readonly,rcb"`
	rand   *rand.Rand
}

func NewHwWifi() *HwWifi {
	wf := new(HwWifi)
	hwio.MustInitRegs(wf)
	wf.rand = rand.New(rand.NewSource(0))
	wf.bbInit()
	return wf
}

func (wf *HwWifi) bbInit() {
	// Initialize baseband registers
	wf.bbRegs[0x00] = 0x6D // Chip ID
	wf.bbRegs[0x5D] = 0x1

	for idx := range wf.bbRegWritable {
		if (idx >= 0x1 && idx <= 0xC) || (idx >= 0x13 && idx <= 0x15) ||
			(idx >= 0x1B && idx <= 0x26) || (idx >= 0x28 && idx <= 0x4C) ||
			(idx >= 0x4E && idx <= 0x5C) || (idx >= 0x62 && idx <= 0x63) ||
			idx == 0x65 || idx == 0x67 || idx == 0x68 {
			wf.bbRegWritable[idx] = true
		}
	}

}

func (wf *HwWifi) WriteBASEBANDCNT(_, val uint16) {
	idx := val & 0xFF

	switch val >> 12 {
	case 5:
		// Write to regs
		if wf.bbRegWritable[idx] {
			wf.bbRegs[idx] = uint8(wf.BaseBandWrite.Value & 0xFF)
			Emu.Log().Infof("[wifi] BB write reg %02x: %02x", idx, wf.bbRegs[idx])
		} else {
			Emu.Log().Warnf("[wifi] BB write ignored to reg %02x", idx)
		}
	case 6:
		// Read regs
		wf.BaseBandRead.Value = uint16(wf.bbRegs[idx])
		Emu.Log().Infof("[wifi] BB read reg %02x: %02x", idx, wf.bbRegs[idx])

	default:
		Emu.Log().Errorf("[wifi] invalid BB control: %04x", val)
	}
}

func (wf *HwWifi) ReadRANDOM(_ uint16) uint16 {
	return uint16(wf.rand.Uint32()) & 0x3FF
}
