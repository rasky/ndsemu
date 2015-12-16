package main

import (
	"fmt"
	"ndsemu/arm"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDS7IOMap struct {
	GetPC  func() uint32
	Card   *Gamecard
	Irq    *HwIrq
	Ipc    *HwIpc
	Mc     *HwMemoryController
	Timers *HwTimers
	Spi    *HwSpiBus
}

func (m *NDS7IOMap) Reset() {

}

func (m *NDS7IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x01C2:
		return m.Spi.ReadSPIDATA()
	case 0x0241:
		return m.Mc.ReadWRAMCNT()
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"pc":   fmt.Sprintf("%08x", m.GetPC()),
		}).Error("invalid NDS7 I/O Read8")
		return 0x00
	}
}

func (m *NDS7IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	case 0x1A0:
		data := m.Card.ReadAUXSPICNT()
		data = (data & 0xFF00) | uint16(val)
		m.Card.WriteAUXSPICNT(data)
	case 0x1A1:
		data := m.Card.ReadAUXSPICNT()
		data = (data & 0x00FF) | (uint16(val) << 8)
		m.Card.WriteAUXSPICNT(data)
	case 0x1A8, 0x1A9, 0x1AA, 0x1AB, 0x1AC, 0x1AD, 0x1AE, 0x1AF:
		m.Card.WriteCommand(addr, val)
	case 0x01C2:
		m.Spi.WriteSPIDATA(uint8(val))
	case 0x0301:
		nds7.Cpu.SetLine(arm.LineHalt, true)
	default:
		log.WithFields(log.Fields{
			"pc":   fmt.Sprintf("%08x", m.GetPC()),
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
	case 0x01C0:
		return m.Spi.ReadSPICNT()
	case 0x01C2:
		return uint16(m.Spi.ReadSPIDATA())
	case 0x0208:
		return m.Irq.ReadIME()
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"pc":   fmt.Sprintf("%08x", m.GetPC()),
		}).Error("invalid NDS7 I/O Read16")
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
	case 0x01C0:
		m.Spi.WriteSPICNT(val)
	case 0x01C2:
		m.Spi.WriteSPIDATA(uint8(val))
	case 0x0208:
		m.Irq.WriteIME(val)
	default:
		log.WithFields(log.Fields{
			"pc":   fmt.Sprintf("%08x", m.GetPC()),
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%04x", val),
		}).Error("invalid NDS7 I/O Write16")
	}
}

func (m *NDS7IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFFFF {
	case 0x0180:
		return uint32(m.Ipc.ReadIPCSYNC(CpuNds7))
	case 0x01A4:
		return m.Card.ReadROMCTL()
	case 0x01C0:
		w1 := m.Spi.ReadSPICNT()
		w2 := m.Spi.ReadSPIDATA()
		return (uint32(w2) << 16) | uint32(w1)
	case 0x0208:
		return uint32(m.Irq.ReadIME())
	case 0x0210:
		return m.Irq.ReadIE()
	case 0x0214:
		return m.Irq.ReadIF()
	case 0x100010:
		return m.Card.ReadData()
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"pc":   fmt.Sprintf("%08x", m.GetPC()),
		}).Error("invalid NDS7 I/O Read32")
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
	case 0x01A4:
		m.Card.WriteROMCTL(val)
	case 0x0208:
		m.Irq.WriteIME(uint16(val))
	case 0x0210:
		m.Irq.WriteIE(val)
	case 0x0214:
		m.Irq.WriteIF(val)
	default:
		log.WithFields(log.Fields{
			"pc":   fmt.Sprintf("%08x", m.GetPC()),
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%08x", val),
		}).Error("invalid NDS7 I/O Write32")
	}
}
