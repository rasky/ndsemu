package main

import (
	"flag"
	"fmt"
	"ndsemu/emu/hw"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"

	log "gopkg.in/Sirupsen/logrus.v0"
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
	skipBiosArg = flag.Bool("s", false, "skip bios and run immediately")
	debug       = flag.Bool("debug", false, "run with debugger")
	cpuprofile  = flag.String("cpuprofile", "", "write cpu profile to file")
	quiet       = flag.Bool("quiet", false, "low logging verbosity")

	nds7 *NDS7
	nds9 *NDS9
)

func main() {
	// Required by go-sdl2, to be run at the beginning of main
	runtime.LockOSThread()

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
	timers9.SetName("t9")
	timers7 := &HwTimers{
		Irq: irq7,
	}
	timers7.SetName("t7")
	ipc := new(HwIpc)
	ipc.HwIrq[CpuNds9] = irq9
	ipc.HwIrq[CpuNds7] = irq7
	mc := NewMemoryController(nds9, nds7)
	gc := NewGamecard(irq7, "bios/biosnds7.rom")
	if err := gc.MapCartFile(flag.Arg(0)); err != nil {
		panic(err)
	}
	lcd := NewHwLcd(irq9, irq7)
	div := NewHwDivisor()

	var dma9 [4]*HwDmaChannel
	var dma7 [4]*HwDmaChannel
	for i := 0; i < 4; i++ {
		dma9[i] = NewHwDmaChannel(CpuNds9, i, nds9.Bus, irq9)
		dma7[i] = NewHwDmaChannel(CpuNds7, i, nds7.Bus, irq7)
	}

	iocommon := &NDSIOCommon{}

	iomap9 := NDS9IOMap{
		Common: iocommon,
		GetPC:  func() uint32 { return uint32(nds9.Cpu.GetPC()) },
		Card:   gc,
		Ipc:    ipc,
		Mc:     mc,
		Timers: timers9,
		Irq:    irq9,
		Lcd:    lcd,
		Div:    div,
		Dma:    dma9,
	}
	iomap9.Reset()

	rtc := NewHwRtc()

	spi := new(HwSpiBus)
	spi.AddDevice(0, NewHwPowerMan())
	spi.AddDevice(1, NewHwFirmwareFlash("bios/firmware.bin"))
	spi.AddDevice(2, NewHwTouchScreen())

	iomap7 := NDS7IOMap{
		Common: iocommon,
		GetPC:  func() uint32 { return uint32(nds7.Cpu.GetPC()) },
		Card:   gc,
		Ipc:    ipc,
		Mc:     mc,
		Timers: timers7,
		Spi:    spi,
		Irq:    irq7,
		Rtc:    rtc,
		Lcd:    lcd,
		Dma:    dma7,
	}
	iomap7.Reset()

	nds9.Bus.MapIORegs(0x04000000, 0x04FFFFFF, &iomap9)
	nds9.Cpu.Reset() // trigger reset exception

	nds7.Bus.MapIORegs(0x04000000, 0x04FFFFFF, &iomap7)
	nds7.Cpu.Reset() // trigger reset exception

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
		if *cpuprofile != "" {
			pprof.StopCPUProfile()
		}
		os.Exit(1)
	}()

	SyncConfig.HSync = lcd.SyncEvent

	Emu = NewNDSEmulator()
	Emu.Sync.AddCpu(nds9)
	Emu.Sync.AddCpu(nds7)
	Emu.Sync.AddSubsystem(timers9)
	Emu.Sync.AddSubsystem(timers7)

	if *skipBiosArg {
		if err := InjectGamecard(gc, nds9, nds7); err != nil {
			fmt.Println(err)
			return
		}
		mc.WriteWRAMCNT(3)
		iocommon.postflg = 1
	}

	if *debug {
		Emu.StartDebugger()
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *quiet {
		log.SetLevel(log.FatalLevel)
	}

	hwout := hw.NewOutput(hw.OutputConfig{
		Title:  "NDSEmu - Nintendo DS Emulator",
		Width:  256,
		Height: 192 + 90 + 192,
	})
	hwout.EnableVideo(true)

	for nf := 0; ; nf++ {
		log.Infof("Begin frame: %d", nf)

		if !hwout.Poll() {
			break
		}

		hwout.BeginFrame()
		Emu.Sync.RunOneFrame()
		hwout.EndFrame()
	}
}
