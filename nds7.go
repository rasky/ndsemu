package main

import (
	"io/ioutil"
	"log"
	"ndsemu/arm"
	"ndsemu/emu"
)

type NDS7 struct {
	Cpu *arm.Cpu
	Bus *BankedBus

	WRam    [64 * 1024]byte
	MainRam []byte
}

func NewNDS7(ram []byte) *NDS7 {
	bios7, err := ioutil.ReadFile("bios/biosnds7.rom")
	if err != nil {
		log.Fatal(err)
	}

	bus := BankedBus{
		NumWaitStates: 0,
	}

	cpu := arm.NewCpu(arm.ARMv4, &bus)

	nds7 := &NDS7{
		Cpu:     cpu,
		Bus:     &bus,
		MainRam: ram,
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

func (n *NDS7) Reset() {
	n.Cpu.Reset()
}

func (n *NDS7) Cycles() int64 {
	return n.Cpu.Clock
}

func (n *NDS7) Run(target int64) {
	n.Cpu.Run(target)
}
