package main

import (
	"fmt"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDS7IOMap struct {
	Ipc *HwIpc
	Mc  *HwMemoryController
}

func (m *NDS7IOMap) Reset() {

}

func (m *NDS7IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x0241:
		return m.Mc.ReadWRAMCNT()
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS7 I/O Read8")
		return 0x00
	}
}

func (m *NDS7IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%02x", val),
		}).Error("invalid NDS7 I/O Write8")
	}
}

func (m *NDS7IOMap) Read16(addr uint32) uint16 {
	switch addr & 0xFFFF {
	case 0x0180:
		return m.Ipc.ReadIPCSYNC(CpuNds7)
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS7 I/O Read16")
		return 0x0000
	}
}

func (m *NDS7IOMap) Write16(addr uint32, val uint16) {
	switch addr & 0xFFFF {
	case 0x0180:
		m.Ipc.WriteIPCSYNC(CpuNds7, val)
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%04x", val),
		}).Error("invalid NDS7 I/O Write16")
	}
}

func (m *NDS7IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFF {
	case 0x0180:
		return uint32(m.Ipc.ReadIPCSYNC(CpuNds7))
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS7 I/O Read32")
		return 0x00000000
	}
}

func (m *NDS7IOMap) Write32(addr uint32, val uint32) {
	switch addr & 0xFFFF {
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%08x", val),
		}).Error("invalid NDS7 I/O Write32")
	}
}
