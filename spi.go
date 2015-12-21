package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
)

type SpiDevice interface {

	// Begin a new SPI data transfer
	BeginTransfer() chan uint8
}

type HwSpiBus struct {
	devs    [4]SpiDevice
	control uint16
	tdev    SpiDevice
	ch      chan uint8
	data    uint8
}

func (spi *HwSpiBus) AddDevice(addr int, dev SpiDevice) {
	if spi.devs[addr] != nil {
		panic("spi: address already assigned")
	}
	spi.devs[addr] = dev
}

func (spi *HwSpiBus) WriteSPICNT(val uint16) {
	const rwmask = 0xCF03
	spi.control = (spi.control &^ rwmask) | (val & rwmask)
	// log.Infof("[SPI] control=%04x (%04x)", spi.control, val)

	if spi.control&(1<<15) != 0 {
		// Begin transfer
		didx := (spi.control >> 8) & 3
		if spi.tdev != nil {
			if spi.tdev != spi.devs[didx] {
				panic("SPI changed device during transfer")
			}
		} else {
			spi.tdev = spi.devs[didx]
			if spi.tdev == nil {
				log.Fatalf("SPI device %d not implemented", didx)
			}
			spi.ch = spi.tdev.BeginTransfer()
		}

		if spi.control&(1<<14) != 0 {
			panic("SPI IRQ not implemented")
		}
	}
}

func (spi *HwSpiBus) ReadSPICNT() uint16 {
	cntrl := spi.control
	return cntrl
}

func (spi *HwSpiBus) WriteSPIDATA(val uint8) {
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
	case spi.data = <-spi.ch:
	case <-time.After(1 * time.Second):
		panic("deadlock in SPI channel reading")
	}
	if spi.control&(1<<11) == 0 {
		close(spi.ch)
		spi.tdev = nil
		log.Info("[SPI] end of transfer")
	}
}

func (spi *HwSpiBus) ReadSPIDATA() uint8 {
	// log.WithField("PC", nds7.Cpu.GetPC()).Infof("[SPI] reading %02x", spi.data)
	return spi.data
}
