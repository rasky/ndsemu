package main

import (
	"ndsemu/gamecard"
	"os"
)

type CpuNum int

const (
	CpuNds9 CpuNum = 0
	CpuNds7 CpuNum = 1
)

func main() {
	gc := gamecard.NewGamecard()
	gc.MapCartFile(os.Args[1])

	ipc := new(HwIpc)

	iomap9 := NDS9IOMap{
		Card: gc,
		Ipc:  ipc,
	}
	iomap9.Reset()

	iomap7 := NDS7IOMap{
		Ipc: ipc,
	}
	iomap7.Reset()

	nds9 := NewNDS9()
	nds9.Bus.MapIORegs(0x04000000, &iomap9)
	nds9.Cpu.Reset() // trigger reset exception

	nds7 := NewNDS7()
	nds7.Bus.MapIORegs(0x04000000, &iomap7)
	nds7.Cpu.Reset() // trigger reset exception

	clock := int64(0)
	for {
		clock += 100000
		nds7.Cpu.Run(clock)
		nds9.Cpu.Run(clock)
	}
}
