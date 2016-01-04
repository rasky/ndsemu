package main

import (
	"ndsemu/emu/hwio"
	"time"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type SpiDevice interface {
	// Begin a new SPI data transfer
	BeginTransfer() chan uint8
}

type HwSpiBus struct {
	devs [4]SpiDevice
	tdev SpiDevice
	ch   chan uint8

	SpiCnt  hwio.Reg16 `hwio:"offset=0x0,rwmask=0xCF03,wcb"`
	SpiData hwio.Reg8  `hwio:"offset=0x2,wcb"`
	Dummy   hwio.Reg8  `hwio:"offset=0x3,rwmask=0"` // disable logging
}

func NewHwSpiBus() *HwSpiBus {
	spi := new(HwSpiBus)
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
	// log.Infof("[SPI] control=%04x (%04x)", spi.control, val)

	if spi.SpiCnt.Value&(1<<15) != 0 {
		// Begin transfer
		didx := (spi.SpiCnt.Value >> 8) & 3
		if spi.tdev != nil {
			if spi.tdev != spi.devs[didx] {
				log.Warnf("[SPI] wrong new device=%d", didx)
				// panic("SPI changed device during transfer")
				close(spi.ch)
			} else {
				return
			}
		}
		spi.tdev = spi.devs[didx]
		if spi.tdev == nil {
			log.Fatalf("SPI device %d not implemented", didx)
		}
		spi.ch = spi.tdev.BeginTransfer()
		log.Infof("[SPI] begin transfer device=%d", didx)

		if spi.SpiCnt.Value&(1<<14) != 0 {
			panic("SPI IRQ not implemented")
		}
	}
}

func (spi *HwSpiBus) WriteSPIDATA(_, val uint8) {
	if spi.tdev == nil {
		log.Warn("SPIDATA written but no transfer")
		return
	}

	// log.WithField("PC", nds7.Cpu.GetPC()).Infof("[SPI] writing %02x", val)
	select {
	case spi.ch <- val:
	case <-time.After(1 * time.Second):
		panic("deadlock in SPI channel writing")
	}
	select {
	case spi.SpiData.Value = <-spi.ch:
	case <-time.After(1 * time.Second):
		panic("deadlock in SPI channel reading")
	}
	if spi.SpiCnt.Value&(1<<11) == 0 {
		close(spi.ch)
		spi.tdev = nil
		log.Info("[SPI] end of transfer")
	}
}
