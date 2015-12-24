package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
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

	nds7 *NDS7
	nds9 *NDS9
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("game card file is required")
		return
	}

	var ram [4 * 1024 * 1024]byte

	nds9 = NewNDS9(ram[:])
	nds7 = NewNDS7(ram[:])

	irq9 := &HwIrq{
		Cpu: nds9.Cpu,
	}
	irq7 := &HwIrq{
		Cpu: nds7.Cpu,
	}
	timers9 := &HwTimers{
		Irq: irq9,
	}
	timers7 := &HwTimers{
		Irq: irq7,
	}
	ipc := new(HwIpc)
	mc := &HwMemoryController{
		Nds9: nds9,
		Nds7: nds7,
	}

	gc := NewGamecard(irq7, "bios/biosnds7.rom")
	if err := gc.MapCartFile(flag.Arg(0)); err != nil {
		panic(err)
	}

	iomap9 := NDS9IOMap{
		GetPC:  func() uint32 { return uint32(nds9.Cpu.GetPC()) },
		Card:   gc,
		Ipc:    ipc,
		Mc:     mc,
		Timers: timers9,
		Irq:    irq9,
	}
	iomap9.Reset()

	rtc := NewHwRtc()

	spi := new(HwSpiBus)
	spi.AddDevice(0, NewHwPowerMan())
	spi.AddDevice(1, NewHwFirmwareFlash("bios/firmware.bin"))

	iomap7 := NDS7IOMap{
		GetPC:  func() uint32 { return uint32(nds7.Cpu.GetPC()) },
		Card:   gc,
		Ipc:    ipc,
		Mc:     mc,
		Timers: timers7,
		Spi:    spi,
		Irq:    irq7,
		Rtc:    rtc,
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
		mc.WriteWRAMCNT(3)
		iomap9.postflg = 1
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		f, err := os.Create("ram.dump")
		if err == nil {
			f.Write(nds9.MainRam)
			f.Close()
		}
		f, err = os.Create("wram.dump")
		if err == nil {
			f.Write(mc.wram[:])
			f.Write(nds7.WRam[:])
			f.Close()
		}
		os.Exit(1)
	}()

	sync := SyncEmu{}
	sync.AddSubsystem(nds9)
	sync.AddSubsystem(nds7)
	sync.AddSubsystem(timers9)
	sync.AddSubsystem(timers7)

	clock := int64(0)
	for {
		clock += 10000
		sync.Run(clock)
	}
}
