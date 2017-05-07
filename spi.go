package main

import (
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"ndsemu/emu/spi"
)

type HwSpiBus struct {
	spi.Bus

	SpiCnt  hwio.Reg16 `hwio:"offset=0x0,rwmask=0xCF03,wcb"`
	SpiData hwio.Reg8  `hwio:"offset=0x2,wcb"`
	Dummy   hwio.Reg8  `hwio:"offset=0x3,rwmask=0"` // disable logging
}

func NewHwSpiBus() *HwSpiBus {
	spi := &HwSpiBus{}
	spi.SpiBusName = "SpiMain"
	hwio.MustInitRegs(spi)
	spi.SpiData.WriteCb = spi.WriteSPIDATA // for speed
	return spi
}

func (spi *HwSpiBus) WriteSPICNT(_, val uint16) {
	// log.Infof("control=%04x (%04x)", spi.control, val)
	log.ModSpi.InfoZ("SpiMain: write SPICNT").Hex16("val", val).End()
	if val&(1<<14) != 0 {
		panic("SPI IRQ not implemented")
	}
}

func (spi *HwSpiBus) WriteSPIDATA(_, val uint8) {

	if spi.SpiCnt.Value&(1<<15) != 0 {
		devaddr := (spi.SpiCnt.Value >> 8) & 3
		spi.BeginTransfer(int(devaddr))
	}

	// Transfer the byte through SPI, and set the value
	// read back into the register, making it available
	// for the CPU.
	spi.SpiData.Value = spi.Transfer(val)

	// Bit 11 is the "auto-reselect CS line". When 1, the CS
	// line is kept high at the end of the current byte transfer,
	// so basically the transfer continue. When 0, the CS line
	// goes down after the current byte is written.
	if spi.SpiCnt.Value&(1<<11) == 0 {
		spi.EndTransfer()
	}
}
