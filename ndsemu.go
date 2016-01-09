package main

import (
	"flag"
	"fmt"
	"ndsemu/emu/hw"
	log "ndsemu/emu/logger"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"
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
	flagLogging = flag.String("log", "", "enable logging for specified modules")

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

	Emu = NewNDSEmulator()
	Emu.Sync.AddCpu(nds9, "arm9")
	Emu.Sync.AddCpu(nds7, "arm7")
	Emu.Sync.AddSubsystem(nds9.Timers, "timers9")
	Emu.Sync.AddSubsystem(nds7.Timers, "timers7")

	if err := Emu.Hw.Gc.MapCartFile(flag.Arg(0)); err != nil {
		panic(err)
	}

	iomap9 := NDS9IOMap{
		GetPC:   func() uint32 { return uint32(nds9.Cpu.GetPC()) },
		Card:    Emu.Hw.Gc,
		Ipc:     Emu.Hw.Ipc,
		Mc:      Emu.Hw.Mc,
		Timers:  nds9.Timers,
		Irq:     nds9.Irq,
		Lcd:     Emu.Hw.Lcd9,
		Div:     Emu.Hw.Div,
		Dma:     nds9.Dma,
		E2d:     Emu.Hw.E2d,
		DmaFill: nds9.DmaFill,
		Key:     Emu.Hw.Key,
	}
	iomap9.Reset()

	iomap7 := NDS7IOMap{
		GetPC:  func() uint32 { return uint32(nds7.Cpu.GetPC()) },
		Card:   Emu.Hw.Gc,
		Ipc:    Emu.Hw.Ipc,
		Mc:     Emu.Hw.Mc,
		Timers: nds7.Timers,
		Spi:    Emu.Hw.Spi,
		Irq:    nds7.Irq,
		Rtc:    Emu.Hw.Rtc,
		Lcd:    Emu.Hw.Lcd7,
		Dma:    nds7.Dma,
		Wifi:   Emu.Hw.Wifi,
		Key:    Emu.Hw.Key,
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
			f.Write(Emu.Hw.Mc.wram[:])
			f.Write(nds7.WRam[:])
			f.Close()
		}
		f, err = os.Create("vram.dump")
		if err == nil {
			f.Write(Emu.Mem.Vram[:])
			f.Close()
		}
		f, err = os.Create("oam.dump")
		if err == nil {
			f.Write(nds9.OamRam[:])
			f.Close()
		}
		if *cpuprofile != "" {
			pprof.StopCPUProfile()
		}
		os.Exit(1)
	}()

	if *skipBiosArg {
		if err := InjectGamecard(Emu.Hw.Gc, nds9, nds7); err != nil {
			fmt.Println(err)
			return
		}

		// Shared wram: map everything to ARM7
		Emu.Hw.Mc.WramCnt.Write8(0, 3)

		// Set post-boot flag to 1
		iomap9.misc.PostFlg.Value = 1
		iomap7.misc.PostFlg.Value = 1

		nds9.Irq.Ime.Value = 0x1
		nds7.Irq.Ime.Value = 0x1
		nds9.Irq.Ie.Value = uint32(IrqIpcRecvFifo | IrqTimers | IrqVBlank)
		nds7.Irq.Ie.Value = uint32(IrqIpcRecvFifo | IrqTimers | IrqVBlank)

		// VRAM: map everything in "LCDC mode"
		Emu.Hw.Mc.VramCntA.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntB.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntC.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntD.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntE.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntF.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntG.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntH.Write8(0, 0x80)
		Emu.Hw.Mc.VramCntI.Write8(0, 0x80)

		// Gamecard: skip directly to key2 status
		Emu.Hw.Gc.stat = gcStatusKey2

		nds9.Cp15.ConfigureControlReg(0x52078, 0x00FF085)
	}

	if *debug {
		Emu.StartDebugger()
	}

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.ModEmu.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *flagLogging != "" {
		var modmask log.ModuleMask
		for _, modname := range strings.Split(*flagLogging, ",") {
			if modname == "all" {
				modmask |= log.ModuleMaskAll
			} else if m, found := log.ModuleByName(modname); found {
				modmask |= 1 << m
			} else {
				log.ModEmu.Fatal("invalid module name:", modname)
			}
		}
		log.EnableDebugModules(modmask)
	}

	hwout := hw.NewOutput(hw.OutputConfig{
		Title:  "NDSEmu - Nintendo DS Emulator",
		Width:  256,
		Height: 192 + 90 + 192,
	})
	hwout.EnableVideo(true)

	for {
		if !hwout.Poll() {
			break
		}

		x, y, btn := hwout.GetMouseState()
		pendown := btn&hw.MouseButtonLeft != 0
		Emu.Hw.Key.SetPenDown(pendown)
		Emu.Hw.Tsc.SetPen(pendown, x, y)

		screen := hwout.BeginFrame()
		Emu.RunOneFrame(screen)
		hwout.EndFrame()
	}
}
