package main

import (
	"fmt"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDS7IOMap struct {
	Ipc    *HwIpc
	Mc     *HwMemoryController
	Timers *HwTimers
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
	case 0x0100:
		return m.Timers.Timers[0].ReadCounter()
	case 0x0102:
		return m.Timers.Timers[0].ReadControl()
	case 0x0104:
		return m.Timers.Timers[1].ReadCounter()
	case 0x0106:
		return m.Timers.Timers[1].ReadControl()
	case 0x0108:
		return m.Timers.Timers[2].ReadCounter()
	case 0x010A:
		return m.Timers.Timers[2].ReadControl()
	case 0x010C:
		return m.Timers.Timers[3].ReadCounter()
	case 0x010E:
		return m.Timers.Timers[3].ReadControl()
	case 0x0180:
		return m.Ipc.ReadIPCSYNC(CpuNds7)
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS7 I/O Read16")
		return 0x0000
	}
}

func (m *NDS7IOMap) Write16(addr uint32, val uint16) {
	switch addr & 0xFFFF {
	case 0x0100:
		m.Timers.Timers[0].WriteReload(val)
	case 0x0102:
		m.Timers.Timers[0].WriteControl(val)
	case 0x0104:
		m.Timers.Timers[1].WriteReload(val)
	case 0x0106:
		m.Timers.Timers[1].WriteControl(val)
	case 0x0108:
		m.Timers.Timers[2].WriteReload(val)
	case 0x010A:
		m.Timers.Timers[2].WriteControl(val)
	case 0x010C:
		m.Timers.Timers[3].WriteReload(val)
	case 0x010E:
		m.Timers.Timers[3].WriteControl(val)
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
	case 0x0100:
		m.Timers.Timers[0].WriteReload(uint16(val))
		m.Timers.Timers[0].WriteControl(uint16(val >> 16))
	case 0x0104:
		m.Timers.Timers[1].WriteReload(uint16(val))
		m.Timers.Timers[1].WriteControl(uint16(val >> 16))
	case 0x0108:
		m.Timers.Timers[2].WriteReload(uint16(val))
		m.Timers.Timers[2].WriteControl(uint16(val >> 16))
	case 0x010C:
		m.Timers.Timers[3].WriteReload(uint16(val))
		m.Timers.Timers[3].WriteControl(uint16(val >> 16))
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%08x", val),
		}).Error("invalid NDS7 I/O Write32")
	}
}
