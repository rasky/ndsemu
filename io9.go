package main

import (
	"fmt"
	"ndsemu/gamecard"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDS9IOMap struct {
	Card *gamecard.Gamecard
	Ipc  *HwIpc
	Mc   *HwMemoryController

	postflg uint8
}

func (m *NDS9IOMap) Reset() {

}

func (m *NDS9IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x0247:
		return m.Mc.ReadWRAMCNT()
	case 0x0300:
		log.Warn("Read8 POSTFLG")
		return m.postflg
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS9 I/O Read8")
		return 0x00
	}
}

func (m *NDS9IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	case 0x0247:
		m.Mc.WriteWRAMCNT(val)
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%02x", val),
		}).Error("invalid NDS9 I/O Write8")
	}
}

func (m *NDS9IOMap) Read16(addr uint32) uint16 {
	switch addr & 0xFFFF {
	case 0x0180:
		return m.Ipc.ReadIPCSYNC(CpuNds9)
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS9 I/O Read16")
		return 0x0000
	}
}

func (m *NDS9IOMap) Write16(addr uint32, val uint16) {
	switch addr & 0xFFFF {
	case 0x0180:
		m.Ipc.WriteIPCSYNC(CpuNds9, val)
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%04x", val),
		}).Error("invalid NDS9 I/O Write16")
	}
}

func (m *NDS9IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFF {
	case 0x01A0:
		return uint32(m.Card.ReadAUXSPICNT()) | (uint32(m.Card.ReadAUXSPIDATA()) << 16)
	case 0x01A4:
		return m.Card.ReadROMCTL()
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS9 I/O Read32")
		return 0x00000000
	}
}

func (m *NDS9IOMap) Write32(addr uint32, val uint32) {
	switch addr & 0xFFFF {
	case 0x01A0:
		m.Card.WriteAUXSPICNT(uint16(val & 0xFFFF))
		m.Card.WriteAUXSPIDATA(uint16(val >> 16))
	case 0x01A4:
		m.Card.WriteROMCTL(val)
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%08x", val),
		}).Error("invalid NDS9 I/O Write32")
	}
}
