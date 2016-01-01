package main

import "ndsemu/emu/hwio"

type NDSIOCommon struct {
	postflg uint8
}

type NDS9IOMap struct {
	TableLo hwio.Table
	TableHi hwio.Table

	GetPC   func() uint32
	Card    *Gamecard
	Ipc     *HwIpc
	Mc      *HwMemoryController
	Timers  *HwTimers
	Irq     *HwIrq
	Common  *NDSIOCommon
	Lcd     *HwLcd
	Div     *HwDivisor
	Dma     [4]*HwDmaChannel
	DmaFill *HwDmaFill
	E2d     [2]*HwEngine2d
}

func (m *NDS9IOMap) Reset() {
	m.TableLo.Name = "io9"
	m.TableLo.Reset()
	m.TableHi.Name = "io9-hi"
	m.TableHi.Reset()

	m.TableLo.MapBank(0x4000000, m.E2d[0], 0)
	m.TableLo.MapBank(0x4000100, &m.Timers.Timers[0], 0)
	m.TableLo.MapBank(0x4000104, &m.Timers.Timers[1], 0)
	m.TableLo.MapBank(0x4000108, &m.Timers.Timers[2], 0)
	m.TableLo.MapBank(0x400010C, &m.Timers.Timers[3], 0)
	m.TableLo.MapBank(0x40001A0, m.Card, 0)
	m.TableLo.MapReg16(0x4000204, &m.Mc.ExMemCnt)
	m.TableLo.MapBank(0x4000200, m.Irq, 0)
	m.TableLo.MapBank(0x4000240, m.Mc, 0)
	m.TableLo.MapBank(0x4000280, m.Div, 0)
	m.TableLo.MapBank(0x40000B0, m.Dma[0], 0)
	m.TableLo.MapBank(0x40000BC, m.Dma[1], 0)
	m.TableLo.MapBank(0x40000C8, m.Dma[2], 0)
	m.TableLo.MapBank(0x40000D4, m.Dma[3], 0)
	m.TableLo.MapBank(0x40000E0, m.DmaFill, 0)
	m.TableLo.MapBank(0x4000180, m.Ipc, 0)
	m.TableLo.MapBank(0x4001000, m.E2d[1], 0)

	m.TableHi.MapBank(0x4100000, m.Ipc, 1)
	m.TableHi.MapBank(0x4100010, m.Card, 1)
}

func (m *NDS9IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x0300:
		return m.Common.postflg
	default:
		return m.TableLo.Read8(addr)
	}
}

func (m *NDS9IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	default:
		m.TableLo.Write8(addr, val)
	}
}

func (m *NDS9IOMap) Read16(addr uint32) uint16 {
	switch addr & 0xFFFF {
	case 0x0004:
		return m.Lcd.ReadDISPSTAT()
	case 0x0006:
		return m.Lcd.ReadVCOUNT()
	case 0x0130:
		// log.Warn("[IO7] read KEYINPUT")
		return 0x3FF
	default:
		return m.TableLo.Read16(addr)
	}
}

func (m *NDS9IOMap) Write16(addr uint32, val uint16) {
	switch addr & 0xFFFF {
	case 0x0004:
		m.Lcd.WriteDISPSTAT(val)
	default:
		m.TableLo.Write16(addr, val)
	}
}

func (m *NDS9IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFF {
	default:
		return m.TableLo.Read32(addr)
	}
}

func (m *NDS9IOMap) Write32(addr uint32, val uint32) {
	switch addr & 0xFFFF {
	default:
		m.TableLo.Write32(addr, val)
	}
}
