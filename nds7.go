package main

import (
	"ndsemu/arm"
	"ndsemu/emu/fixed"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

type NDS7 struct {
	Cpu    *arm.Cpu
	Bus    *hwio.Table
	Irq    *HwIrq
	Timers *HwTimers
	Dma    [4]*HwDmaChannel

	misc7   miscRegs7
	miscgba miscRegsGba
}

func NewNDS7(dojit bool) *NDS7 {
	bus := hwio.NewTable("bus7")
	bus.SetWaitStates(0)

	cpu := arm.NewCpu(arm.ARMv4, bus, dojit)

	nds7 := &NDS7{
		Cpu: cpu,
		Bus: bus,
	}

	nds7.Irq = NewHwIrq("irq7", cpu)
	nds7.Timers = NewHWTimers("t7", nds7.Irq)
	for i := 0; i < 4; i++ {
		nds7.Dma[i] = NewHwDmaChannel(CpuNds7, i, nds7.Bus, nds7.Irq)
	}
	hwio.MustInitRegs(&nds7.misc7)
	hwio.MustInitRegs(&nds7.miscgba)

	return nds7
}

func (n *NDS7) InitBus(emu *NDSEmulator) {

	n.Bus.MapMemorySlice(0x00000000, 0x00003FFF, emu.Rom.Bios7, true)
	n.Bus.MapMemorySlice(0x02000000, 0x02FFFFFF, emu.Mem.Ram[:], false)
	n.Bus.MapMemorySlice(0x03800000, 0x03FFFFFF, emu.Mem.Wram[:], false)

	n.Bus.MapReg8(0x4000300, &n.misc7.PostFlg)
	n.Bus.MapBank(0x4000000, emu.Hw.Lcd7, 0)
	n.Bus.MapBank(0x40000B0, n.Dma[0], 0)
	n.Bus.MapBank(0x40000BC, n.Dma[1], 0)
	n.Bus.MapBank(0x40000C8, n.Dma[2], 0)
	n.Bus.MapBank(0x40000D4, n.Dma[3], 0)
	n.Bus.MapBank(0x4000100, &n.Timers.Timers[0], 0)
	n.Bus.MapBank(0x4000104, &n.Timers.Timers[1], 0)
	n.Bus.MapBank(0x4000108, &n.Timers.Timers[2], 0)
	n.Bus.MapBank(0x400010C, &n.Timers.Timers[3], 0)
	n.Bus.MapBank(0x4000130, emu.Hw.Key, 0)
	n.Bus.MapBank(0x4000130, emu.Hw.Key, 1)
	n.Bus.MapReg16(0x4000134, &n.misc7.Rcnt)
	n.Bus.MapReg8(0x4000138, &emu.Hw.Rtc.Serial)
	n.Bus.MapReg8(0x4000139, &n.misc7.Dummy8)
	n.Bus.MapBank(0x4000180, emu.Hw.Ipc, 2)
	// n.Bus.MapBank(0x40001A0, emu.Hw.Gc, 0)  mapped by memcnt
	n.Bus.MapBank(0x40001C0, emu.Hw.Spi, 0)
	n.Bus.MapBank(0x4000200, n.Irq, 0)
	n.Bus.MapReg16(0x4000204, &emu.Hw.Mc.ExMemStat)
	n.Bus.MapBank(0x4000240, emu.Hw.Mc, 1)
	n.Bus.MapReg8(0x4000301, &n.misc7.Halt7)
	for i := 0; i < 16; i++ {
		n.Bus.MapBank(0x4000400+uint32(i)*0x10, &emu.Hw.Snd.Ch[i], 0)
	}
	n.Bus.MapBank(0x4000500, emu.Hw.Snd, 1)
	n.Bus.MapBank(0x4100000, emu.Hw.Ipc, 3)
	// n.Bus.MapBank(0x4100010, emu.Hw.Gc, 1)  mapped by memcnt

	// Setup all wifi mirrors
	n.Bus.MapBank(0x4800000, emu.Hw.Wifi, 0)
	n.Bus.MapBank(0x4801000, emu.Hw.Wifi, 0)
	n.Bus.MapBank(0x4804000, emu.Hw.Wifi, 1)
	n.Bus.MapBank(0x4806000, emu.Hw.Wifi, 0)
	n.Bus.MapBank(0x4807000, emu.Hw.Wifi, 0)

	n.Bus.MapBank(0x4808000, emu.Hw.Wifi, 0)
	n.Bus.MapBank(0x4809000, emu.Hw.Wifi, 0)
	n.Bus.MapBank(0x480C000, emu.Hw.Wifi, 1)
	n.Bus.MapBank(0x480E000, emu.Hw.Wifi, 0)
	n.Bus.MapBank(0x480F000, emu.Hw.Wifi, 0)
}

func (n *NDS7) InitBusGba(emu *NDSEmulator) {
	n.Bus.Unmap(0x0, 0xFFFFFFFF)

	n.Bus.MapMemorySlice(0x00000000, 0x00003FFF, emu.Rom.BiosGba, true)
	n.Bus.MapMemorySlice(0x02000000, 0x0203FFFF, emu.Mem.Ram[:], false)
	n.Bus.MapMemorySlice(0x03000000, 0x03FFFFFF, emu.Mem.Wram[:32*1024], false)
	n.Bus.MapMemorySlice(0x05000000, 0x050003FF, emu.Mem.PaletteRam[:], false)
	n.Bus.MapMemorySlice(0x06000000, 0x06017FFF, emu.Mem.Vram[256*1024:256*1024+128*1024], false)
	n.Bus.MapMemorySlice(0x07000000, 0x070003FF, emu.Mem.OamRam[:], false)
	n.Bus.MapMemorySlice(0x08000000, 0x09FFFFFF, Emu.Hw.Sl2.Rom[:], true)
	n.Bus.MapMemorySlice(0x0A000000, 0x0BFFFFFF, Emu.Hw.Sl2.Rom[:], true)
	n.Bus.MapMemorySlice(0x0C000000, 0x0DFFFFFF, Emu.Hw.Sl2.Rom[:], true)

	n.Bus.MapBank(0x4000000, emu.Hw.Lcd7, 0)
	n.Bus.MapBank(0x4000000, emu.Hw.E2d[1], 0)

	n.Bus.MapReg32(0x4000088, &n.miscgba.SndBias)
	n.Bus.MapBank(0x40000B0, n.Dma[0], 0)
	n.Bus.MapBank(0x40000BC, n.Dma[1], 0)
	n.Bus.MapBank(0x40000C8, n.Dma[2], 0)
	n.Bus.MapBank(0x40000D4, n.Dma[3], 0)
	n.Bus.MapBank(0x4000100, &n.Timers.Timers[0], 0)
	n.Bus.MapBank(0x4000104, &n.Timers.Timers[1], 0)
	n.Bus.MapBank(0x4000108, &n.Timers.Timers[2], 0)
	n.Bus.MapBank(0x400010C, &n.Timers.Timers[3], 0)
	n.Bus.MapBank(0x4000130, emu.Hw.Key, 0)
	n.Bus.MapBank(0x4000200, n.Irq, 0)
	n.Bus.MapReg8(0x4000301, &n.miscgba.HaltCnt)
}

func (n *NDS7) Frequency() fixed.F8 {
	return fixed.NewF8(cNds7Clock)
}

func (n *NDS7) GetPC() uint32 {
	return uint32(n.Cpu.GetPC())
}

func (n *NDS7) Reset() {
	n.Cpu.Reset()
}

func (n *NDS7) Cycles() int64 {
	return n.Cpu.Clock
}

func (n *NDS7) Run(targetCycles int64) {
	n.Cpu.Run(targetCycles)
}

func (n *NDS7) Retarget(targetCycles int64) {
	n.Cpu.Retarget(targetCycles)
}

func (n *NDS7) TriggerDmaEvent(event DmaEvent) {
	for i := 0; i < 4; i++ {
		n.Dma[i].TriggerEvent(event)
	}
}

type miscRegs7 struct {
	Rcnt hwio.Reg16 `hwio:"rwmask=0x8000"`

	PostFlg hwio.Reg8 `hwio:"rwmask=1"`

	Dummy8 hwio.Reg8 `hwio:"rwmask=0"`

	Halt7 hwio.Reg8 `hwio:"wcb"`
}

func (m *miscRegs7) WriteHALT7(_, val uint8) {
	mode := val >> 6
	switch mode {
	case 0: // ignored

	case 1: // Switch GBA
		// Halt the CPU. NDS9 has been already halted by BIOS,
		// so we quickly get to end of frame
		nds7.Cpu.SetLine(arm.LineHalt, true)
		Emu.SwitchToGba()

	case 2: // halt
		nds7.Cpu.SetLine(arm.LineHalt, true)

	case 3: // sleep
		log.ModEmu.FatalZ("unimplemented sleep mode").End()
	}
}

type miscRegsGba struct {
	SndBias hwio.Reg32 // FIXME: remove once we emulate gba sound
	HaltCnt hwio.Reg8  `hwio:"wcb"`
}

func (m *miscRegsGba) WriteHALTCNT(_, _ uint8) {
	nds7.Cpu.SetLine(arm.LineHalt, true)
}
