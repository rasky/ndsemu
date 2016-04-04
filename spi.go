package main

import (
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modSpi = log.NewModule("spi")

type SpiStatus int

const (
	SpiFinish SpiStatus = iota
	SpiContinue
)

type SpiDevice interface {
	SpiBegin()
	SpiTransfer(data []byte) ([]byte, SpiStatus)
	SpiEnd()
}

type HwSpiBus struct {
	devs  [4]SpiDevice
	tdev  SpiDevice
	buf   []byte
	reply []byte

	SpiCnt  hwio.Reg16 `hwio:"offset=0x0,rwmask=0xCF03,wcb"`
	SpiData hwio.Reg8  `hwio:"offset=0x2,wcb"`
	Dummy   hwio.Reg8  `hwio:"offset=0x3,rwmask=0"` // disable logging
}

func NewHwSpiBus() *HwSpiBus {
	spi := new(HwSpiBus)
	spi.buf = make([]byte, 0, 16)
	hwio.MustInitRegs(spi)
	return spi
}

func (spi *HwSpiBus) AddDevice(addr int, dev SpiDevice) {
	if spi.devs[addr] != nil {
		panic("spi: address already assigned")
	}
	spi.devs[addr] = dev
}

func (spi *HwSpiBus) WriteSPICNT(_, val uint16) {
	// log.Infof("control=%04x (%04x)", spi.control, val)

	if spi.SpiCnt.Value&(1<<15) != 0 {
		// Begin transfer
		didx := (spi.SpiCnt.Value >> 8) & 3
		if spi.tdev != nil {
			if spi.tdev != spi.devs[didx] {
				modSpi.Warnf("wrong new device=%d", didx)
				// panic("SPI changed device during transfer")
			} else {
				return
			}
		}
		spi.tdev = spi.devs[didx]
		if spi.tdev == nil {
			modSpi.Fatalf("SPI device %d not implemented", didx)
		}
		modSpi.Infof("begin transfer device=%d (%T)", didx, spi.tdev)
		spi.tdev.SpiBegin()
		spi.buf = spi.buf[:0]
		spi.reply = nil

		if spi.SpiCnt.Value&(1<<14) != 0 {
			panic("SPI IRQ not implemented")
		}
	}
}

func (spi *HwSpiBus) WriteSPIDATA(_, val uint8) {
	if spi.tdev == nil {
		modSpi.Warn("SPIDATA written but no transfer")
		return
	}

	if len(spi.reply) == 0 {
		spi.buf = append(spi.buf, val)
		retval, stat := spi.tdev.SpiTransfer(spi.buf)
		if stat == SpiFinish {
			spi.buf = spi.buf[:0]
		}
		spi.reply = retval
	}

	if len(spi.reply) == 0 {
		spi.SpiData.Value = 0
	} else {
		spi.SpiData.Value = spi.reply[0]
		spi.reply = spi.reply[1:]
	}

	if spi.SpiCnt.Value&(1<<11) == 0 {
		spi.tdev.SpiEnd()
		spi.tdev = nil
		modSpi.Info("end of transfer")
	}
}
