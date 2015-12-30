package main

import (
	"ndsemu/arm"
	"ndsemu/emu/hwio"
)

type miscRegs7 struct {
	Rcnt hwio.Reg16 `hwio:"rwmask=0x8000"`

	// The NDS7 BIOS brings this register to 0x200 at boot, with a slow loop
	// with delay that takes ~1 second. If we reset it at 0x200, it will just
	// skip everything and the emulator will boot faster.
	SndBias hwio.Reg16 `hwio:"reset=0x200,rwmask=0x1FF"`
}

type NDS7IOMap struct {
	TableLo   hwio.Table
	TableHi   hwio.Table
	TableWifi hwio.Table

	GetPC  func() uint32
	Card   *Gamecard
	Irq    *HwIrq
	Ipc    *HwIpc
	Mc     *HwMemoryController
	Timers *HwTimers
	Spi    *HwSpiBus
	Rtc    *HwRtc
	Lcd    *HwLcd
	Common *NDSIOCommon
	Dma    [4]*HwDmaChannel

	misc miscRegs7
}

func (m *NDS7IOMap) Reset() {
	m.TableLo.Name = "io7"
	m.TableLo.Reset()
	m.TableHi.Name = "io7-hi"
	m.TableHi.Reset()
	m.TableWifi.Name = "io7-wifi"
	m.TableWifi.Reset()

	hwio.MustInitRegs(&m.misc)

	m.TableLo.MapReg16(0x4000134, &m.misc.Rcnt)
	m.TableLo.MapReg16(0x4000504, &m.misc.SndBias)
	m.TableLo.MapBank(0x40000B0, m.Dma[0], 0)
	m.TableLo.MapBank(0x40000BC, m.Dma[1], 0)
	m.TableLo.MapBank(0x40000C8, m.Dma[2], 0)
	m.TableLo.MapBank(0x40000D4, m.Dma[3], 0)
	m.TableLo.MapBank(0x4000180, m.Ipc, 2)

	m.TableHi.MapBank(0x4100000, m.Ipc, 3)
	m.TableHi.MapBank(0x4100010, m.Card, 1)
}

func (m *NDS7IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x0138:
		return m.Rtc.ReadSERIAL()
	case 0x01C2:
		return m.Spi.ReadSPIDATA()
	case 0x0241:
		return m.Mc.ReadWRAMCNT()
	default:
		return m.TableLo.Read8(addr)
	}
}

func (m *NDS7IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	case 0x138:
		m.Rtc.WriteSERIAL(val)
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
	case 0x0208:
		m.Irq.WriteIME(uint16(val))
	case 0x0301:
		nds7.Cpu.SetLine(arm.LineHalt, true)
	default:
		m.TableLo.Write8(addr, val)
	}
}

func (m *NDS7IOMap) Read16(addr uint32) uint16 {
	switch addr & 0xFFFF {
	case 0x0004:
		return m.Lcd.ReadDISPSTAT()
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
	case 0x0130:
		// log.Warn("[IO7] read KEYINPUT")
		return 0x3FF
	case 0x0136:
		// log.Warn("[IO7] read EXTKEYIN")
		return (1 << 0) | (1 << 1) | (1 << 3) | (1 << 6)
	case 0x0138:
		return uint16(m.Rtc.ReadSERIAL())
	case 0x01C0:
		return m.Spi.ReadSPICNT()
	case 0x01C2:
		return uint16(m.Spi.ReadSPIDATA())
	case 0x0208:
		return m.Irq.ReadIME()
	default:
		return m.TableLo.Read16(addr)
	}
}

func (m *NDS7IOMap) Write16(addr uint32, val uint16) {
	switch addr & 0xFFFF {
	case 0x0004:
		m.Lcd.WriteDISPSTAT(val)
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
	case 0x0138:
		m.Rtc.WriteSERIAL(uint8(val))
	case 0x01C0:
		m.Spi.WriteSPICNT(val)
	case 0x01C2:
		m.Spi.WriteSPIDATA(uint8(val))
	case 0x0208:
		m.Irq.WriteIME(val)
	default:
		m.TableLo.Write16(addr, val)
	}
}

func (m *NDS7IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFF {
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
	default:
		return m.TableLo.Read32(addr)
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
		m.TableLo.Write32(addr, val)
	}
}
