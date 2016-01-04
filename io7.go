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

	PostFlg hwio.Reg8 `hwio:"rwmask=1"`

	Dummy8 hwio.Reg8 `hwio:"rwmask=0"`
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
	Dma    [4]*HwDmaChannel
	Wifi   *HwWifi

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

	m.TableLo.MapReg8(0x4000300, &m.misc.PostFlg)
	m.TableLo.MapReg16(0x4000504, &m.misc.SndBias)
	m.TableLo.MapBank(0x4000000, m.Lcd, 0)
	m.TableLo.MapBank(0x40000B0, m.Dma[0], 0)
	m.TableLo.MapBank(0x40000BC, m.Dma[1], 0)
	m.TableLo.MapBank(0x40000C8, m.Dma[2], 0)
	m.TableLo.MapBank(0x40000D4, m.Dma[3], 0)
	m.TableLo.MapBank(0x4000100, &m.Timers.Timers[0], 0)
	m.TableLo.MapBank(0x4000104, &m.Timers.Timers[1], 0)
	m.TableLo.MapBank(0x4000108, &m.Timers.Timers[2], 0)
	m.TableLo.MapBank(0x400010C, &m.Timers.Timers[3], 0)
	m.TableLo.MapReg16(0x4000134, &m.misc.Rcnt)
	m.TableLo.MapReg8(0x4000138, &m.Rtc.Serial)
	m.TableLo.MapReg8(0x4000139, &m.misc.Dummy8)
	m.TableLo.MapBank(0x4000180, m.Ipc, 2)
	m.TableLo.MapBank(0x40001A0, m.Card, 0)
	m.TableLo.MapBank(0x40001C0, m.Spi, 0)
	m.TableLo.MapBank(0x4000200, m.Irq, 0)
	m.TableLo.MapReg16(0x4000204, &m.Mc.ExMemStat)
	m.TableLo.MapBank(0x4000240, m.Mc, 1)

	m.TableHi.MapBank(0x4100000, m.Ipc, 3)
	m.TableHi.MapBank(0x4100010, m.Card, 1)

	m.TableWifi.MapBank(0x4800000, m.Wifi, 0)
	m.TableWifi.MapBank(0x4801000, m.Wifi, 0)
	m.TableWifi.MapBank(0x4804000, m.Wifi, 1)
	m.TableWifi.MapBank(0x4806000, m.Wifi, 0)
	m.TableWifi.MapBank(0x4807000, m.Wifi, 0)

	m.TableWifi.MapBank(0x4808000, m.Wifi, 0)
	m.TableWifi.MapBank(0x4809000, m.Wifi, 0)
	m.TableWifi.MapBank(0x480C000, m.Wifi, 1)
	m.TableWifi.MapBank(0x480E000, m.Wifi, 0)
	m.TableWifi.MapBank(0x480F000, m.Wifi, 0)
}

func (m *NDS7IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	default:
		return m.TableLo.Read8(addr)
	}
}

func (m *NDS7IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	case 0x0301:
		nds7.Cpu.SetLine(arm.LineHalt, true)
	default:
		m.TableLo.Write8(addr, val)
	}
}

func (m *NDS7IOMap) Read16(addr uint32) uint16 {
	switch addr & 0xFFFF {
	case 0x0130:
		// log.Warn("[IO7] read KEYINPUT")
		return 0x3FF
	case 0x0136:
		// log.Warn("[IO7] read EXTKEYIN")
		return (1 << 0) | (1 << 1) | (1 << 3) | (1 << 6)
	default:
		return m.TableLo.Read16(addr)
	}
}

func (m *NDS7IOMap) Write16(addr uint32, val uint16) {
	switch addr & 0xFFFF {
	default:
		m.TableLo.Write16(addr, val)
	}
}

func (m *NDS7IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFF {
	default:
		return m.TableLo.Read32(addr)
	}
}

func (m *NDS7IOMap) Write32(addr uint32, val uint32) {
	switch addr & 0xFFFF {
	default:
		m.TableLo.Write32(addr, val)
	}
}
