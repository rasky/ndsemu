package main

import (
	"io/ioutil"
	"ndsemu/arm"
	"unsafe"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDS9 struct {
	Cpu *arm.Cpu
	Bus *BankedBus

	MainRam []byte
}

const cItcmPhysicalSize = 32 * 1024
const cDtcmPhysicalSize = 16 * 1024

func NewNDS9(ram []byte) *NDS9 {
	bios9, err := ioutil.ReadFile("bios/biosnds9.rom")
	if err != nil {
		log.Fatal(err)
	}

	bus := BankedBus{}

	cpu := arm.NewCpu(arm.ARMv5, &bus)
	cp15 := cpu.EnableCp15()
	cp15.ConfigureTcm(cItcmPhysicalSize, cDtcmPhysicalSize)
	cp15.ConfigureControlReg(0x2078, 0x00FF085)

	nds9 := &NDS9{
		Cpu:     cpu,
		Bus:     &bus,
		MainRam: ram,
	}

	bus.MapMemory(0x02000000, 0x02FFFFFF, unsafe.Pointer(&nds9.MainRam[0]), len(nds9.MainRam), false)
	bus.MapMemory(0x0FFF0000, 0x0FFF7FFF, unsafe.Pointer(&bios9[0]), len(bios9), true)

	return nds9
}

func (n *NDS9) Frequency() fixed8 {
	return NewFixed8(cNds9Clock)
}

func (n *NDS9) Reset() {
	n.Cpu.Reset()
}

func (n *NDS9) Cycles() int64 {
	return n.Cpu.Clock
}

func (n *NDS9) Run(target int64) {
	n.Cpu.Run(target)
}
