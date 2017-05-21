package main

import (
	"fmt"
	"io/ioutil"
	"ndsemu/e2d"
	"ndsemu/emu"
	"ndsemu/emu/debugger"
	"ndsemu/emu/gfx"
	log "ndsemu/emu/logger"
	"ndsemu/raster3d"
	"os"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
)

type NDSMemory struct {
	Ram        [4 * 1024 * 1024]byte // main RAM
	Vram       [656 * 1024]byte      // video RAM
	Wram       [64 * 1024]byte       // work RAM (nds7)
	PaletteRam [2048]byte            // pallette RAM
	OamRam     [2048]byte            // object attribute RAM
}

type NDSRom struct {
	Bios9 []byte
	Bios7 []byte
}

type NDSHardware struct {
	E2d  [2]*e2d.HwEngine2d
	E3d  *raster3d.HwEngine3d
	Lcd9 *HwLcd
	Lcd7 *HwLcd
	Mc   *HwMemoryController
	Ipc  *HwIpc
	Div  *HwDivisor
	Rtc  *HwRtc
	Wifi *HwWifi
	Spi  *HwSpiBus
	Gc   *Gamecard
	Ff   *HwFirmwareFlash
	Tsc  *HwTouchScreen
	Pow  *HwPowerMan
	Key  *HwKey
	Snd  *HwSound
	Geom *HwGeometry
	Bkp  *HwBackupRam
	Sl2  *HwSlot2
}

type NDSEmulator struct {
	Mem  *NDSMemory
	Rom  *NDSRom
	Hw   *NDSHardware
	Sync *emu.Sync

	dbg        *debugger.Debugger
	screen     gfx.Buffer
	audio      []int16
	framecount int
	powcnt     uint32
}

var Emu *NDSEmulator

func NewNDSHardware(mem *NDSMemory, firmware string) *NDSHardware {
	hw := new(NDSHardware)
	bindir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	nds9 = NewNDS9()
	nds7 = NewNDS7()
	hw.Mc = NewMemoryController(nds9, nds7, mem.Vram[:])
	hw.E3d = raster3d.NewHwEngine3d()
	hw.E2d[0] = e2d.NewHwEngine2d(0, hw.Mc, gfx.LayerFunc{Func: hw.E3d.Draw3D})
	hw.E2d[1] = e2d.NewHwEngine2d(1, hw.Mc, nil)
	hw.Lcd9 = NewHwLcd(nds9.Irq)
	hw.Lcd7 = NewHwLcd(nds7.Irq)
	hw.Ipc = NewHwIpc(nds9.Irq, nds7.Irq)
	hw.Div = NewHwDivisor()
	hw.Rtc = NewHwRtc()
	hw.Wifi = NewHwWifi()
	hw.Bkp = NewHwBackupRam()
	hw.Gc = NewGamecard(filepath.Join(bindir, "bios/biosnds7.rom"), hw.Bkp)
	hw.Tsc = NewHwTouchScreen()
	hw.Key = NewHwKey()
	hw.Snd = NewHwSound(nds7.Bus)
	hw.Geom = NewHwGeometry(nds9.Irq, hw.E3d)
	hw.Sl2 = NewHwSlot2()

	hw.Spi = NewHwSpiBus()
	hw.Ff = NewHwFirmwareFlash()
	hw.Pow = NewHwPowerMan()
	hw.Spi.AddDevice(0, hw.Pow)
	hw.Spi.AddDevice(1, hw.Ff)
	hw.Spi.AddDevice(2, hw.Tsc)

	// Pass bg scrolling regs to 3D engine for final 2D compositing pass
	hw.E3d.SetBgRegs(&hw.E2d[0].DispCnt.Value,
		&hw.E2d[0].Bg0Cnt.Value, &hw.E2d[0].Bg0XOfs.Value)

	return hw
}

func NewNDSRom() *NDSRom {
	rom := new(NDSRom)
	bindir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	bios9, err := ioutil.ReadFile(filepath.Join(bindir, "bios/biosnds9.rom"))
	if err != nil {
		log.ModEmu.FatalZ("error loading rom").Error("err", err).End()
	}
	rom.Bios9 = bios9

	bios7, err := ioutil.ReadFile(filepath.Join(bindir, "bios/biosnds7.rom"))
	if err != nil {
		log.ModEmu.FatalZ("error loading rom").Error("err", err).End()
	}
	rom.Bios7 = bios7

	return rom
}

func NewNDSEmulator(firmware string) *NDSEmulator {
	mem := new(NDSMemory)
	rom := NewNDSRom()
	hw := NewNDSHardware(mem, firmware)

	// Initialize syncing system
	sync, err := emu.NewSync(SyncConfig)
	if err != nil {
		panic(err)
	}
	sync.AddCpu(nds9, "arm9")
	sync.AddCpu(nds7, "arm7")
	sync.AddSubsystem(nds9.Timers, "timers9")
	sync.AddSubsystem(nds7.Timers, "timers7")
	sync.AddSubsystem(hw.Geom, "gx")

	e := &NDSEmulator{
		Mem:  mem,
		Hw:   hw,
		Rom:  rom,
		Sync: sync,
	}

	// Set the hsync callback to this instance's function
	e.Sync.SetHSyncCallback(e.hsync)

	// Register the syncer's logger as global logging function,
	// so that everything will also log the current subsystem
	// status (eg: CPU program counter)
	log.AddContext(e.Sync)

	emu.BreakFunc = e.DebugBreak

	// Initialize the memory map and reset the CPUs
	nds9.InitBus(e)
	nds7.InitBus(e)
	nds9.Reset()
	nds7.Reset()

	return e
}

func (emu *NDSEmulator) StartDebugger() {
	emu.dbg = debugger.New([]debugger.Cpu{nds7.Cpu, nds9.Cpu}, emu.Sync)

	type DebugConfig struct {
		Breakpoints []string
		Watchpoints []string
	}

	cfg := &DebugConfig{}
	if _, err := toml.DecodeFile("debug.ini", cfg); err != nil {
		log.ModEmu.WithField("error", err).Warnf("error loading debug.ini")
	} else {
		for _, bkp := range cfg.Breakpoints {
			if b, err := strconv.ParseUint(bkp, 0, 32); err != nil {
				log.ModEmu.WithField("error", err).Fatalf("invalid breakpoint %q", bkp)
			} else {
				emu.dbg.AddBreakpoint(uint32(b))
				log.ModEmu.WithField("break", fmt.Sprintf("0x%08x", uint32(b))).Warnf("add breakpoint")
			}
		}
		for _, bkp := range cfg.Watchpoints {
			if b, err := strconv.ParseUint(bkp, 0, 32); err != nil {
				log.ModEmu.WithField("error", err).Fatalf("invalid watchpoint %q", bkp)
			} else {
				emu.dbg.AddWatchpoint(uint32(b))
				log.ModEmu.WithField("watch", fmt.Sprintf("0x%08x", uint32(b))).Warnf("add watchpoint")
			}
		}
	}

	go emu.dbg.Run()
}

func (emu *NDSEmulator) DebugBreak(msg string) {
	if emu.dbg != nil {
		emu.dbg.Break(msg)
	} else {
		log.ModEmu.ErrorZ(msg).End()
		log.ModEmu.PanicZ("debugging breakpoint, aborting").End()
	}
}

func (emu *NDSEmulator) eaOn() bool       { return emu.powcnt&(1<<1) != 0 }
func (emu *NDSEmulator) ebOn() bool       { return emu.powcnt&(1<<9) != 0 }
func (emu *NDSEmulator) lcdSwapped() bool { return emu.powcnt&(1<<15) != 0 }

func (emu *NDSEmulator) hsync(x, y int) {
	emu.Hw.Lcd9.SyncEvent(x, y)
	emu.Hw.Lcd7.SyncEvent(x, y)

	if y == 0 && x == 0 {
		if emu.eaOn() {
			emu.Hw.E2d[0].BeginFrame()
		}
		if emu.ebOn() {
			emu.Hw.E2d[1].BeginFrame()
		}
	}

	if y < 192 {
		if x == 0 {
			emu.beginLine(y)
		} else if x == cHBlankFirstDot {
			emu.endLine(y)
			// Trigger the DMA hblank event (only in visible part of screen)
			for _, dmach := range nds9.Dma {
				dmach.TriggerEvent(DmaEventHBlank)
			}
		}
	}

	// Vblank
	if y == 192 && x == 0 {
		if emu.eaOn() {
			emu.Hw.E2d[0].EndFrame()
		}
		if emu.ebOn() {
			emu.Hw.E2d[1].EndFrame()
		}
		emu.Hw.E3d.EndFrame()
	}

	// 3D starts at scanline 214, before VBlank end. This is useful for us too, as we
	// need some time to prepare the first line (computations, texture decompression, etc.)
	if y == 214 && x == 0 {
		emu.Hw.E3d.SetVram(emu.Hw.Mc.VramTextureBank(), emu.Hw.Mc.VramTexturePaletteBank())
		emu.Hw.E3d.BeginFrame()
	}

	// Per-line audio emulation
	if x == 0 && emu.Hw.Pow.AudioEnabled() {
		nsamples := len(emu.audio) / 2
		n0 := (nsamples * y) / 263
		n1 := (nsamples * (y + 1)) / 263
		emu.Hw.Snd.RunOneFrame(emu.audio[n0*2 : n1*2])
	}
}

func (emu *NDSEmulator) RunOneFrame(screen gfx.Buffer, audio []int16) {
	// Save powcnt for this frame; letting it change within a frame isn't
	// really necessary and it's hard to handle with our parallel system
	emu.powcnt = nds9.misc.PowCnt.Value

	up, down := "B", "A"
	if emu.lcdSwapped() {
		up, down = down, up
	}

	log.ModGfx.InfoZ("begin frame").String("up", up).String("down", down).End()

	emu.screen = screen
	emu.audio = audio
	emu.Sync.RunOneFrame()
	emu.audio = nil
	emu.framecount++
}

func (emu *NDSEmulator) beginLine(y int) {
	ya := y + 192 + 90
	yb := y
	if emu.lcdSwapped() {
		ya, yb = yb, ya
	}

	if emu.eaOn() {
		emu.Hw.E2d[0].BeginLine(y, emu.screen.Line(ya))
	}
	if emu.ebOn() {
		emu.Hw.E2d[1].BeginLine(y, emu.screen.Line(yb))
	}
}

func (emu *NDSEmulator) endLine(y int) {
	if emu.eaOn() {
		emu.Hw.E2d[0].EndLine(y)
	}
	if emu.ebOn() {
		emu.Hw.E2d[1].EndLine(y)
	}
}
