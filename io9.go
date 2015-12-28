package main

import "ndsemu/emu/hwio"

type NDSIOCommon struct {
	postflg uint8
}

type NDS9IOMap struct {
	TableLo hwio.Table
	TableHi hwio.Table

	GetPC  func() uint32
	Card   *Gamecard
	Ipc    *HwIpc
	Mc     *HwMemoryController
	Timers *HwTimers
	Irq    *HwIrq
	Common *NDSIOCommon
	Lcd    *HwLcd
	Div    *HwDivisor
	Dma    [4]*HwDmaChannel
}

func (m *NDS9IOMap) Reset() {
	m.TableLo.Name = "io9"
	m.TableLo.Reset()
	m.TableHi.Name = "io9-hi"
	m.TableHi.Reset()

	m.TableLo.MapBank(0x4000280, m.Div, 0)
	m.TableLo.MapBank(0x40000B0, m.Dma[0], 0)
	m.TableLo.MapBank(0x40000BC, m.Dma[1], 0)
	m.TableLo.MapBank(0x40000C8, m.Dma[2], 0)
	m.TableLo.MapBank(0x40000D4, m.Dma[3], 0)
}

func (m *NDS9IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x0247:
		return m.Mc.ReadWRAMCNT()
	case 0x0300:
		return m.Common.postflg
	default:
		if addr < 0x4100000 {
			return m.TableLo.Read8(addr)
		} else {
			return m.TableHi.Read8(addr)
		}
	}
}

func (m *NDS9IOMap) Write8(addr uint32, val uint8) {
	switch addr & 0xFFFF {
	case 0x0240:
		m.Mc.WriteVRAMCNTA(val)
	case 0x0241:
		m.Mc.WriteVRAMCNTB(val)
	case 0x0242:
		m.Mc.WriteVRAMCNTC(val)
	case 0x0243:
		m.Mc.WriteVRAMCNTD(val)
	case 0x0244:
		m.Mc.WriteVRAMCNTE(val)
	case 0x0245:
		m.Mc.WriteVRAMCNTF(val)
	case 0x0246:
		m.Mc.WriteVRAMCNTG(val)
	case 0x0247:
		m.Mc.WriteWRAMCNT(val)
	case 0x0248:
		m.Mc.WriteVRAMCNTH(val)
	case 0x0249:
		m.Mc.WriteVRAMCNTI(val)
	default:
		if addr < 0x4100000 {
			m.TableLo.Write8(addr, val)
		} else {
			m.TableHi.Write8(addr, val)
		}
	}
}

func (m *NDS9IOMap) Read16(addr uint32) uint16 {
	switch addr & 0xFFFF {
	case 0x0004:
		return m.Lcd.ReadDISPSTAT()
	case 0x0006:
		return m.Lcd.ReadVCOUNT()
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
		return m.Ipc.ReadIPCSYNC(CpuNds9)
	case 0x0184:
		return m.Ipc.ReadIPCFIFOCNT(CpuNds9)
	case 0x0208:
		return m.Irq.ReadIME()
	default:
		if addr < 0x4100000 {
			return m.TableLo.Read16(addr)
		} else {
			return m.TableHi.Read16(addr)
		}
	}
}

func (m *NDS9IOMap) Write16(addr uint32, val uint16) {
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
	case 0x0180:
		m.Ipc.WriteIPCSYNC(CpuNds9, val)
	case 0x0184:
		m.Ipc.WriteIPCFIFOCNT(CpuNds9, val)
	case 0x0208:
		m.Irq.WriteIME(val)
	default:
		if addr < 0x4100000 {
			m.TableLo.Write16(addr, val)
		} else {
			m.TableHi.Write16(addr, val)
		}
	}
}

func (m *NDS9IOMap) Read32(addr uint32) uint32 {
	switch addr & 0xFFFFFF {
	case 0x01A0:
		return uint32(m.Card.ReadAUXSPICNT()) | (uint32(m.Card.ReadAUXSPIDATA()) << 16)
	case 0x01A4:
		return m.Card.ReadROMCTL()
	case 0x0208:
		return uint32(m.Irq.ReadIME())
	case 0x0210:
		return m.Irq.ReadIE()
	case 0x0214:
		return m.Irq.ReadIF()
	case 0x100000:
		return m.Ipc.ReadIPCFIFORECV(CpuNds9)
	default:
		if addr < 0x4100000 {
			return m.TableLo.Read32(addr)
		} else {
			return m.TableHi.Read32(addr)
		}
	}
}

func (m *NDS9IOMap) Write32(addr uint32, val uint32) {
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
	case 0x0188:
		m.Ipc.WriteIPCFIFOSEND(CpuNds9, val)
	case 0x01A0:
		m.Card.WriteAUXSPICNT(uint16(val & 0xFFFF))
		m.Card.WriteAUXSPIDATA(uint16(val >> 16))
	case 0x01A4:
		m.Card.WriteROMCTL(val)
	case 0x0208:
		m.Irq.WriteIME(uint16(val))
	case 0x0210:
		m.Irq.WriteIE(val)
	case 0x0214:
		m.Irq.WriteIF(val)
	default:
		if addr < 0x4100000 {
			m.TableLo.Write32(addr, val)
		} else {
			m.TableHi.Write32(addr, val)
		}
	}
}
