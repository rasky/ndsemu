package main

import (
	"io/ioutil"
	"log"
	"ndsemu/arm"
	"unsafe"
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

	bus.MapMemory(0x00000000, 0x00003FFF, unsafe.Pointer(&bios7[0]), len(bios7), true)
	bus.MapMemory(0x02000000, 0x02FFFFFF, unsafe.Pointer(&nds7.MainRam[0]), len(nds7.MainRam), false)

	return nds7
}
