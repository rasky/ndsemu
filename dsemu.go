package main

import (
	"dsemu/gamecard"
	"os"
)

func main() {
	gc := gamecard.NewGamecard()
	gc.MapCartFile(os.Args[1])

	iomap9 := NDS9IOMap{
		Card: gc,
	}
	iomap9.Reset()

	nds9 := NewNDS9()
	nds9.Bus.MapIORegs(0x04000000, &iomap9)
	nds9.Cpu.Reset() // trigger reset exception

	nds7 := NewNDS7()
	nds7.Cpu.Reset() // trigger reset exception

	nds7.Cpu.Run(1000000)
}
