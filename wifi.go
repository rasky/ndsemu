package main

import (
	"encoding/binary"
	"math/rand"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modWifi = log.NewModule("wifi")

type HwWifi struct {
	WRxBufBegin  hwio.Reg16 `hwio:"offset=0x50"`
	WRxBufEnd    hwio.Reg16 `hwio:"offset=0x52"`
	WRxBufRdAddr hwio.Reg16 `hwio:"offset=0x58,rwmask=0x1FFF"`
	WRxBufRdData hwio.Reg16 `hwio:"offset=0x60,readonly,rcb"`

	WTxBufWrAddr  hwio.Reg16 `hwio:"offset=0x68,rwmask=0x1FFF"`
	WTxBufWrData  hwio.Reg16 `hwio:"offset=0x70,writeonly,wcb"`
	WTxBufGapTop  hwio.Reg16 `hwio:"offset=0x74,rwmask=0x1FFF"`
	WTxBufGapDisp hwio.Reg16 `hwio:"offset=0x76,rwmask=0xFFF"`

	BaseBandCnt   hwio.Reg16 `hwio:"offset=0x158,wcb"`
	BaseBandWrite hwio.Reg16 `hwio:"offset=0x15A,writeonly"`
	BaseBandRead  hwio.Reg16 `hwio:"offset=0x15C,readonly"`
	BaseBandBusy  hwio.Reg16 `hwio:"offset=0x15E,readonly"`
	BaseBandMode  hwio.Reg16 `hwio:"offset=0x160"`
	BaseBandPower hwio.Reg16 `hwio:"offset=0x168"`
	bbRegWritable [256]bool
	bbRegs        [256]uint8

	Random hwio.Reg16 `hwio:"offset=0x044,readonly,rcb"`
	rand   *rand.Rand

	WifiRam hwio.Mem `hwio:"bank=1,offset=0,size=0x2000,rw8=off,rw16,rw32"`
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
			modWifi.InfoZ("BB write").Hex8("reg", uint8(idx)).Hex8("val", wf.bbRegs[idx]).End()
		} else {
			modWifi.WarnZ("BB write ignored").Hex8("reg", uint8(idx)).End()
		}
	case 6:
		// Read regs
		wf.BaseBandRead.Value = uint16(wf.bbRegs[idx])
		modWifi.InfoZ("BB read").Hex8("reg", uint8(idx)).Hex8("val", wf.bbRegs[idx]).End()

	default:
		modWifi.ErrorZ("invalid BB control").Hex16("val", val).End()
	}
}

func (wf *HwWifi) ReadRANDOM(_ uint16) uint16 {
	return uint16(wf.rand.Uint32()) & 0x3FF
}

func (wf *HwWifi) WriteWTXBUFWRDATA(_, val uint16) {
	off := wf.WTxBufWrAddr.Value
	binary.LittleEndian.PutUint16(wf.WifiRam.Data[off:off+2], val)
	off += 2
	if off == wf.WTxBufGapTop.Value {
		off += wf.WTxBufGapDisp.Value * 2
	}
	off &= 0x1FFF
	wf.WTxBufWrAddr.Value = off
}

func (wf *HwWifi) ReadWRXBUFRDDATA(_ uint16) uint16 {
	off := wf.WRxBufRdAddr.Value
	val := binary.LittleEndian.Uint16(wf.WifiRam.Data[off : off+2])
	off += 2
	if off == wf.WRxBufEnd.Value&0x1FFF {
		off = wf.WRxBufBegin.Value
	}
	off &= 0x1FFF
	wf.WRxBufRdAddr.Value = off
	return val
}
