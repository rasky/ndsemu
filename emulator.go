package main

import (
	"io/ioutil"
	"ndsemu/emu"
	"ndsemu/emu/debugger"
	"ndsemu/emu/gfx"
	log "ndsemu/emu/logger"
	"ndsemu/raster3d"
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
	E2d  [2]*HwEngine2d
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
	Key  *HwKey
	Snd  *HwSound
	Geom *HwGeometry
}

type NDSEmulator struct {
	Mem  *NDSMemory
	Rom  *NDSRom
	Hw   *NDSHardware
	Sync *emu.Sync

	dbg        *debugger.Debugger
	screen     gfx.Buffer
	framecount int
	powcnt     uint32
}

var Emu *NDSEmulator

func NewNDSHardware(mem *NDSMemory, firmware string) *NDSHardware {
	hw := new(NDSHardware)

	nds9 = NewNDS9()
	nds7 = NewNDS7()
	hw.Mc = NewMemoryController(nds9, nds7, mem.Vram[:])
	hw.E2d[0] = NewHwEngine2d(0, hw.Mc)
	hw.E2d[1] = NewHwEngine2d(1, hw.Mc)
	hw.E3d = raster3d.NewHwEngine3d()
	hw.Lcd9 = NewHwLcd(nds9.Irq)
	hw.Lcd7 = NewHwLcd(nds7.Irq)
	hw.Ipc = NewHwIpc(nds9.Irq, nds7.Irq)
	hw.Div = NewHwDivisor()
	hw.Rtc = NewHwRtc()
	hw.Wifi = NewHwWifi()
	hw.Gc = NewGamecard("bios/biosnds7.rom")
	hw.Tsc = NewHwTouchScreen()
	hw.Key = NewHwKey()
	hw.Snd = NewHwSound()
	hw.Geom = NewHwGeometry(nds9.Irq, hw.E3d)

	hw.Spi = NewHwSpiBus()
	hw.Ff = NewHwFirmwareFlash()
	hw.Spi.AddDevice(0, NewHwPowerMan())
	hw.Spi.AddDevice(1, hw.Ff)
	hw.Spi.AddDevice(2, hw.Tsc)

	return hw
}

func NewNDSRom() *NDSRom {
	rom := new(NDSRom)

	bios9, err := ioutil.ReadFile("bios/biosnds9.rom")
	if err != nil {
		log.ModEmu.Fatal("error loading rom:", err)
	}
	rom.Bios9 = bios9

	bios7, err := ioutil.ReadFile("bios/biosnds7.rom")
	if err != nil {
		log.ModEmu.Fatal("error loading rom:", err)
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

	// Initialize the memory map and reset the CPUs
	nds9.InitBus(e)
	nds7.InitBus(e)
	nds9.Reset()
	nds7.Reset()

	return e
}

func (emu *NDSEmulator) StartDebugger() {
	emu.dbg = debugger.New([]debugger.Cpu{nds7.Cpu, nds9.Cpu}, emu.Sync)
	// emu.dbg.AddBreakpoint(0x0200d574)
	// emu.dbg.AddBreakpoint(0x0200e362)
	// emu.dbg.AddBreakpoint(0x0200dc18)
	// emu.dbg.AddBreakpoint(0x02000c30)
	// emu.dbg.AddBreakpoint(0x0202c9a0)
	// emu.dbg.AddBreakpoint(0x0202c864)
	// emu.dbg.AddBreakpoint(0x02112c94)
	// emu.dbg.AddBreakpoint(0x0202d494)
	// emu.dbg.AddBreakpoint(0x0202d322)
	// emu.dbg.AddBreakpoint(0x020406dc)
	// emu.dbg.AddBreakpoint(0x0020057f4)
	// emu.dbg.AddBreakpoint(0x02005aa0)
	//02041f2c
	// emu.dbg.AddBreakpoint(0x0200dd84)
	// emu.dbg.AddBreakpoint(0x0200dd04)
	// emu.dbg.AddBreakpoint(0x0200dd8c)
	// emu.dbg.AddBreakpoint(0x0200dd8E)
	// emu.dbg.AddBreakpoint(0x0200dd90)
	// emu.dbg.AddBreakpoint(0x03805fd8)
	// emu.dbg.AddBreakpoint(0x038060a4)
	// emu.dbg.AddBreakpoint(0x01FFA4C4)
	// emu.dbg.AddBreakpoint(0x0232D5D6)
	// emu.dbg.AddWatchpoint(0x27FFE00 + 0x38)
	// emu.dbg.AddBreakpoint(0x130a)
	go emu.dbg.Run()
}

func (emu *NDSEmulator) DebugBreak(msg string) {
	if emu.dbg != nil {
		emu.dbg.Break(msg)
	} else {
		log.ModEmu.Error(msg)
		log.ModEmu.Fatal("debugging breakpoint, aborting")
	}
}

func (emu *NDSEmulator) eaOn() bool       { return emu.powcnt&(1<<1) != 0 }
func (emu *NDSEmulator) ebOn() bool       { return emu.powcnt&(1<<9) != 0 }
func (emu *NDSEmulator) lcdSwapped() bool { return emu.powcnt&(1<<15) != 0 }

func (emu *NDSEmulator) hsync(x, y int) {
	emu.Hw.Lcd9.SyncEvent(x, y)
	emu.Hw.Lcd7.SyncEvent(x, y)

	if y == 0 && x == 0 {
		// NOTE: we must call BeginFrame on 3D before 2D,
		// so that the engine is then ready if the 2D BeginFrame
		// needs it.
		emu.Hw.E3d.BeginFrame()
		emu.Hw.E3d.SetVram(emu.Hw.Mc.VramTextureBank(), emu.Hw.Mc.VramTexturePaletteBank())
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

	if y == 192 && x == 0 {
		if emu.eaOn() {
			emu.Hw.E2d[0].EndFrame()
		}
		if emu.ebOn() {
			emu.Hw.E2d[1].EndFrame()
		}
		emu.Hw.E3d.EndFrame()
	}
}

func (emu *NDSEmulator) RunOneFrame(screen gfx.Buffer) {
	log.ModEmu.Infof("Begin frame: %d", emu.framecount)
	emu.framecount++
	// Save powcnt for this frame; letting it change within a frame isn't
	// really necessary and it's hard to handle with our parallel system
	emu.powcnt = nds9.misc.PowCnt.Value

	emu.screen = screen
	emu.Sync.RunOneFrame()
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
