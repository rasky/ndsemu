package main

import (
	"flag"
	"fmt"
	"ndsemu/gamecard"
	"os"
)

type CpuNum int

const (
	CpuNds9 CpuNum = 0
	CpuNds7 CpuNum = 1
)

/*
 * NDS9: ARM946E-S, architecture ARMv5TE, 66Mhz
 * NDS7: ARM7TDMI, architecture ARMv4T, 33Mhz
 *
 */

var (
	skipBiosArg = flag.Bool(
		"s",
		false,
		"skip bios and run immediately",
	)
)

func main() {
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("game card file is required")
		return
	}

	var ram [4 * 1024 * 1024]byte

	gc := gamecard.NewGamecard()
	gc.MapCartFile(os.Args[1])

	nds9 := NewNDS9(ram[:])
	nds7 := NewNDS7(ram[:])

	timers9 := new(HwTimers)
	timers7 := new(HwTimers)
	ipc := new(HwIpc)
	mc := &HwMemoryController{
		Nds9: nds9,
		Nds7: nds7,
	}

	iomap9 := NDS9IOMap{
		GetPC:  func() uint32 { return uint32(nds9.Cpu.GetPC()) },
		Card:   gc,
		Ipc:    ipc,
		Mc:     mc,
		Timers: timers9,
	}
	iomap9.Reset()

	iomap7 := NDS7IOMap{
		GetPC:  func() uint32 { return uint32(nds7.Cpu.GetPC()) },
		Ipc:    ipc,
		Mc:     mc,
		Timers: timers7,
	}
	iomap7.Reset()

	nds9.Bus.MapIORegs(0x04000000, 0x04FFFFFF, &iomap9)
	nds9.Cpu.Reset() // trigger reset exception

	nds7.Bus.MapIORegs(0x04000000, 0x04FFFFFF, &iomap7)
	nds7.Cpu.Reset() // trigger reset exception

	if *skipBiosArg {
		if err := InjectGamecard(gc, nds9, nds7); err != nil {
			fmt.Println(err)
			return
		}
	}

	sync := SyncEmu{}
	sync.AddSubsystem(nds9)
	sync.AddSubsystem(nds7)
	sync.AddSubsystem(timers9)
	sync.AddSubsystem(timers7)

	clock := int64(0)
	for {
		clock += 100000
		sync.Run(clock)
	}
}
