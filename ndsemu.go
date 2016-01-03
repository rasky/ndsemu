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
	log.SetOutput(os.Stdout)

	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println("game card file is required")
		return
	}

	var ram [4 * 1024 * 1024]byte

	nds9 = NewNDS9(ram[:])
	nds7 = NewNDS7(ram[:])

	Emu = NewNDSEmulator()
	Emu.Sync.AddCpu(nds9)
	Emu.Sync.AddCpu(nds7)
	Emu.Sync.AddSubsystem(nds9.Timers)
	Emu.Sync.AddSubsystem(nds7.Timers)

	ipc := NewHwIpc(nds9.Irq, nds7.Irq)
	mc := NewMemoryController(nds9, nds7, Emu.Mem.Vram[:])
	gc := NewGamecard(nds7.Irq, "bios/biosnds7.rom")
	if err := gc.MapCartFile(flag.Arg(0)); err != nil {
		panic(err)
	}
	lcd := NewHwLcd(nds9.Irq, nds7.Irq)
	div := NewHwDivisor()

	dmafill := NewHwDmaFill()

	iomap9 := NDS9IOMap{
		GetPC:   func() uint32 { return uint32(nds9.Cpu.GetPC()) },
		Card:    gc,
		Ipc:     ipc,
		Mc:      mc,
		Timers:  nds9.Timers,
		Irq:     nds9.Irq,
		Lcd:     lcd,
		Div:     div,
		Dma:     nds9.Dma,
		E2d:     Emu.Hw.E2d,
		DmaFill: dmafill,
	}
	iomap9.Reset()

	rtc := NewHwRtc()
	wifi := NewHwWifi()

	spi := new(HwSpiBus)
	spi.AddDevice(0, NewHwPowerMan())
	spi.AddDevice(1, NewHwFirmwareFlash("bios/firmware.bin"))
	spi.AddDevice(2, NewHwTouchScreen())

	iomap7 := NDS7IOMap{
		GetPC:  func() uint32 { return uint32(nds7.Cpu.GetPC()) },
		Card:   gc,
		Ipc:    ipc,
		Mc:     mc,
		Timers: nds7.Timers,
		Spi:    spi,
		Irq:    nds7.Irq,
		Rtc:    rtc,
		Lcd:    lcd,
		Dma:    nds7.Dma,
		Wifi:   wifi,
	}
	iomap7.Reset()

	nds9.Bus.MapIORegs(0x04000000, 0x0400FFFF, &iomap9)
	nds9.Bus.MapIORegs(0x04100000, 0x0410FFFF, &iomap9.TableHi)
	nds9.Cpu.Reset() // trigger reset exception

	nds7.Bus.MapIORegs(0x04000000, 0x0400FFFF, &iomap7)
	nds7.Bus.MapIORegs(0x04100000, 0x0410FFFF, &iomap7.TableHi)
	nds7.Bus.MapIORegs(0x04800000, 0x0480FFFF, &iomap7.TableWifi)
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

	if *skipBiosArg {
		if err := InjectGamecard(gc, nds9, nds7); err != nil {
			fmt.Println(err)
			return
		}
		mc.WramCnt.Write8(0, 3)
		iomap9.misc.PostFlg.Value = 1
		iomap7.misc.PostFlg.Value = 1
		gc.stat = gcStatusKey2
		mc.VramCntA.Write8(0, 0x80)
		mc.VramCntB.Write8(0, 0x80)
		mc.VramCntC.Write8(0, 0x80)
		mc.VramCntD.Write8(0, 0x80)
		mc.VramCntE.Write8(0, 0x80)
		mc.VramCntF.Write8(0, 0x80)
		mc.VramCntG.Write8(0, 0x80)
		mc.VramCntH.Write8(0, 0x80)
		mc.VramCntI.Write8(0, 0x80)
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

	// FIXME
	Emu.Hw.Lcd = lcd

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

		screen := hwout.BeginFrame()
		Emu.RunOneFrame(screen)
		hwout.EndFrame()
	}
}
