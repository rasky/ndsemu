package main

import (
	"fmt"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDSIOCommon struct {
	postflg uint8
}

type NDS9IOMap struct {
	GetPC  func() uint32
	Card   *Gamecard
	Ipc    *HwIpc
	Mc     *HwMemoryController
	Timers *HwTimers
	Irq    *HwIrq
	Common *NDSIOCommon
	Lcd    *HwLcd
	Div    *HwDivisor
}

func (m *NDS9IOMap) Reset() {
}

func (m *NDS9IOMap) Read8(addr uint32) uint8 {
	switch addr & 0xFFFF {
	case 0x0247:
		return m.Mc.ReadWRAMCNT()
	case 0x0300:
		return m.Common.postflg
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS9 I/O Read8")
		return 0x00
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
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%02x", val),
		}).Error("invalid NDS9 I/O Write8")
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
	case 0x0280:
		return m.Div.ReadDIVCNT()
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS9 I/O Read16")
		return 0x0000
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
	case 0x0280:
		m.Div.WriteDIVCNT(val)
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%04x", val),
		}).Error("invalid NDS9 I/O Write16")
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
	case 0x02A0:
		return m.Div.ReadDIVRESULT_LO()
	case 0x02A4:
		return m.Div.ReadDIVRESULT_HI()
	case 0x02A8:
		return m.Div.ReadDIVREM_LO()
	case 0x02AC:
		return m.Div.ReadDIVREM_HI()
	case 0x100000:
		return m.Ipc.ReadIPCFIFORECV(CpuNds9)
	default:
		log.WithField("addr", fmt.Sprintf("%08x", addr)).Error("invalid NDS9 I/O Read32")
		return 0x00000000
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
	case 0x0290:
		m.Div.WriteDIVNUMER_LO(val)
	case 0x0294:
		m.Div.WriteDIVNUMER_HI(val)
	case 0x0298:
		m.Div.WriteDIVDENOM_LO(val)
	case 0x029C:
		m.Div.WriteDIVDENOM_HI(val)
	default:
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%08x", addr),
			"val":  fmt.Sprintf("%08x", val),
		}).Error("invalid NDS9 I/O Write32")
	}
}
