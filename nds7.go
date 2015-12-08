package main

import (
	"io/ioutil"
	"log"
	"ndsemu/arm"
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

	bus := BankedBus{}

	cpu := arm.NewCpu(arm.ARMv4, &bus)

	nds7 := &NDS7{
		Cpu:     cpu,
		Bus:     &bus,
		MainRam: ram,
	}

	bus.MapMemorySlice(0x00000000, 0x00003FFF, bios7, true)
	bus.MapMemorySlice(0x02000000, 0x02FFFFFF, nds7.MainRam, false)
	bus.MapMemorySlice(0x03800000, 0x03FFFFFF, nds7.WRam[:], false)

	return nds7
}
