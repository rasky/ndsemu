package main

import (
	"ndsemu/arm"
	"ndsemu/emu"
	"ndsemu/emu/hwio"
)

type NDS9 struct {
	Cpu     *arm.Cpu
	Bus     *emu.BankedBus
	Irq     *HwIrq
	Timers  *HwTimers
	Dma     [4]*HwDmaChannel
	DmaFill *HwDmaFill
	Cp15    *arm.Cp15
	misc    miscRegs9
}

const cItcmPhysicalSize = 32 * 1024
const cDtcmPhysicalSize = 16 * 1024

func NewNDS9() *NDS9 {
	bus := &emu.BankedBus{
		NumWaitStates: 7, // this should be 3, but the bus goes at 33Mhz vs 66Mhz of CPU
	}

	cpu := arm.NewCpu(arm.ARMv5, bus)
	cp15 := cpu.EnableCp15()
	cp15.ConfigureTcm(cItcmPhysicalSize, cDtcmPhysicalSize)
	cp15.ConfigureControlReg(0x2078, 0x00FF085)

	nds9 := &NDS9{
		Cpu:  cpu,
		Bus:  bus,
		Cp15: cp15,
	}

	nds9.Irq = NewHwIrq("irq9", cpu)
	nds9.Timers = NewHWTimers("t9", nds9.Irq)
	for i := 0; i < 4; i++ {
		nds9.Dma[i] = NewHwDmaChannel(CpuNds9, i, nds9.Bus, nds9.Irq)
	}
	nds9.DmaFill = NewHwDmaFill()
	hwio.MustInitRegs(&nds9.misc)

	return nds9
}

func (n *NDS9) InitBus(emu *NDSEmulator) {
	var zero [16]byte
	tableLo := hwio.NewTable("io9")
	tableHi := hwio.NewTable("io9-hi")

	n.Bus.MapMemorySlice(0x02000000, 0x02FFFFFF, emu.Mem.Ram[:], false)
	n.Bus.MapIORegs(0x04000000, 0x0400FFFF, tableLo)
	n.Bus.MapIORegs(0x04100000, 0x0410FFFF, tableHi)
	n.Bus.MapMemorySlice(0x05000000, 0x05FFFFFF, emu.Mem.PaletteRam[:], false)
	n.Bus.MapMemorySlice(0x07000000, 0x07FFFFFF, emu.Mem.OamRam[:], false)
	n.Bus.MapMemorySlice(0x08000000, 0x09FFFFFF, zero[:], true)
	n.Bus.MapMemorySlice(0x0A000000, 0x0AFFFFFF, zero[:], true)
	n.Bus.MapMemorySlice(0x0FFF0000, 0x0FFF7FFF, emu.Rom.Bios9, true)

	tableLo.MapReg8(0x4000300, &n.misc.PostFlg)
	tableLo.MapReg32(0x4000304, &n.misc.PowCnt)
	tableLo.MapBank(0x4000000, emu.Hw.Lcd9, 0)
	tableLo.MapBank(0x4000000, emu.Hw.E2d[0], 0)
	tableLo.MapBank(0x4000100, &n.Timers.Timers[0], 0)
	tableLo.MapBank(0x4000104, &n.Timers.Timers[1], 0)
	tableLo.MapBank(0x4000108, &n.Timers.Timers[2], 0)
	tableLo.MapBank(0x400010C, &n.Timers.Timers[3], 0)
	tableLo.MapBank(0x4000130, emu.Hw.Key, 0)
	tableLo.MapBank(0x40001A0, emu.Hw.Gc, 0)
	tableLo.MapReg16(0x4000204, &emu.Hw.Mc.ExMemCnt)
	tableLo.MapBank(0x4000200, n.Irq, 0)
	tableLo.MapBank(0x4000240, emu.Hw.Mc, 0)
	tableLo.MapBank(0x4000280, emu.Hw.Div, 0)
	tableLo.MapBank(0x40000B0, n.Dma[0], 0)
	tableLo.MapBank(0x40000BC, n.Dma[1], 0)
	tableLo.MapBank(0x40000C8, n.Dma[2], 0)
	tableLo.MapBank(0x40000D4, n.Dma[3], 0)
	tableLo.MapBank(0x40000E0, n.DmaFill, 0)
	tableLo.MapBank(0x4000180, emu.Hw.Ipc, 0)
	tableLo.MapBank(0x4001000, emu.Hw.E2d[1], 0)

	tableHi.MapBank(0x4100000, emu.Hw.Ipc, 1)
	tableHi.MapBank(0x4100010, emu.Hw.Gc, 1)
}

func (n *NDS9) Frequency() emu.Fixed8 {
	return emu.NewFixed8(cNds9Clock)
}

func (n *NDS9) GetPC() uint32 {
	return uint32(n.Cpu.GetPC())
}

func (n *NDS9) Reset() {
	n.Cpu.Reset()
}

func (n *NDS9) Cycles() int64 {
	return n.Cpu.Clock
}

func (n *NDS9) Run(targetCycles int64) {
	n.Cpu.Run(targetCycles)
}

func (n *NDS9) Retarget(targetCycles int64) {
	n.Cpu.Retarget(targetCycles)
}

func (n *NDS9) TriggerDmaEvent(event DmaEvent) {
	for i := 0; i < 4; i++ {
		n.Dma[i].TriggerEvent(event)
	}
}

type miscRegs9 struct {
	PostFlg hwio.Reg8  `hwio:"rwmask=3"`
	PowCnt  hwio.Reg32 `hwio:""`
}
