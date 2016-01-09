package main

import (
	"io/ioutil"
	"log"
	"ndsemu/arm"
	"ndsemu/emu"
)

type NDS7 struct {
	Cpu    *arm.Cpu
	Bus    *emu.BankedBus
	Irq    *HwIrq
	Timers *HwTimers
	Dma    [4]*HwDmaChannel

	WRam    [64 * 1024]byte
	MainRam []byte
}

func NewNDS7(ram []byte) *NDS7 {
	bios7, err := ioutil.ReadFile("bios/biosnds7.rom")
	if err != nil {
		log.Fatal(err)
	}

	bus := emu.BankedBus{
		NumWaitStates: 0,
	}

	cpu := arm.NewCpu(arm.ARMv4, &bus)

	nds7 := &NDS7{
		Cpu:     cpu,
		Bus:     &bus,
		MainRam: ram,
	}

	nds7.Irq = NewHwIrq("irq7", cpu)
	nds7.Timers = NewHWTimers("t7", nds7.Irq)
	for i := 0; i < 4; i++ {
		nds7.Dma[i] = NewHwDmaChannel(CpuNds7, i, nds7.Bus, nds7.Irq)
	}

	var zero [16]byte
	bus.MapMemorySlice(0x00000000, 0x00003FFF, bios7, true)
	bus.MapMemorySlice(0x02000000, 0x02FFFFFF, nds7.MainRam, false)
	bus.MapMemorySlice(0x03800000, 0x03FFFFFF, nds7.WRam[:], false)
	bus.MapMemorySlice(0x08000000, 0x09FFFFFF, zero[:], true)

	return nds7
}

func (n *NDS7) Frequency() emu.Fixed8 {
	return emu.NewFixed8(cNds7Clock)
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
