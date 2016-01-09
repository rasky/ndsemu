package main

import (
	"io/ioutil"
	"ndsemu/arm"
	"ndsemu/emu"
	"unsafe"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDS9 struct {
	Cpu     *arm.Cpu
	Bus     *emu.BankedBus
	Irq     *HwIrq
	Timers  *HwTimers
	Dma     [4]*HwDmaChannel
	DmaFill *HwDmaFill

	MainRam    []byte
	PaletteRam [16384]byte // FIXME: make 2k long
	OamRam     [16384]byte // FIXME: make 2k long
}

const cItcmPhysicalSize = 32 * 1024
const cDtcmPhysicalSize = 16 * 1024

func NewNDS9(ram []byte) *NDS9 {
	bios9, err := ioutil.ReadFile("bios/biosnds9.rom")
	if err != nil {
		log.Fatal(err)
	}

	bus := emu.BankedBus{
		NumWaitStates: 7, // this should be 3, but the bus goes at 33Mhz vs 66Mhz of CPU
	}

	cpu := arm.NewCpu(arm.ARMv5, &bus)
	cp15 := cpu.EnableCp15()
	cp15.ConfigureTcm(cItcmPhysicalSize, cDtcmPhysicalSize)
	cp15.ConfigureControlReg(0x2078, 0x00FF085)

	nds9 := &NDS9{
		Cpu:     cpu,
		Bus:     &bus,
		MainRam: ram,
	}

	nds9.Irq = NewHwIrq("irq9", cpu)
	nds9.Timers = NewHWTimers("t9", nds9.Irq)
	for i := 0; i < 4; i++ {
		nds9.Dma[i] = NewHwDmaChannel(CpuNds9, i, nds9.Bus, nds9.Irq)
	}
	nds9.DmaFill = NewHwDmaFill()

	bus.MapMemory(0x02000000, 0x02FFFFFF, unsafe.Pointer(&nds9.MainRam[0]), len(nds9.MainRam), false)
	bus.MapMemory(0x0FFF0000, 0x0FFF7FFF, unsafe.Pointer(&bios9[0]), len(bios9), true)
	bus.MapMemory(0x05000000, 0x05FFFFFF, unsafe.Pointer(&nds9.PaletteRam[0]), len(nds9.PaletteRam), false)
	bus.MapMemory(0x07000000, 0x07FFFFFF, unsafe.Pointer(&nds9.OamRam[0]), len(nds9.OamRam), false)

	var zero [16]byte
	bus.MapMemorySlice(0x08000000, 0x09FFFFFF, zero[:], true)
	bus.MapMemorySlice(0x0A000000, 0x0AFFFFFF, zero[:], true)

	return nds9
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
