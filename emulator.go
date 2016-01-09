package main

import (
	"ndsemu/emu"
	"ndsemu/emu/debugger"
	"ndsemu/emu/gfx"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDSMemory struct {
	Ram  [4 * 1024 * 1024]byte
	Vram [656 * 1024]byte
}

type NDSHardware struct {
	E2d  [2]*HwEngine2d
	Lcd9 *HwLcd
	Lcd7 *HwLcd
	Mc   *HwMemoryController
	Ipc  *HwIpc
	Div  *HwDivisor
	Rtc  *HwRtc
	Wifi *HwWifi
	Spi  *HwSpiBus
	Gc   *Gamecard
	Tsc  *HwTouchScreen
	Key  *HwKey
}

type NDSEmulator struct {
	Mem  *NDSMemory
	Hw   *NDSHardware
	Sync *emu.Sync

	dbg    *debugger.Debugger
	screen gfx.Buffer
}

var Emu *NDSEmulator

func NewNDSHardware(mem *NDSMemory) *NDSHardware {
	hw := new(NDSHardware)

	nds9 = NewNDS9(mem.Ram[:])
	nds7 = NewNDS7(mem.Ram[:])
	hw.Mc = NewMemoryController(nds9, nds7, mem.Vram[:])
	hw.E2d[0] = NewHwEngine2d(0, hw.Mc)
	hw.E2d[1] = NewHwEngine2d(1, hw.Mc)
	hw.Lcd9 = NewHwLcd(nds9.Irq)
	hw.Lcd7 = NewHwLcd(nds7.Irq)
	hw.Ipc = NewHwIpc(nds9.Irq, nds7.Irq)
	hw.Div = NewHwDivisor()
	hw.Rtc = NewHwRtc()
	hw.Wifi = NewHwWifi()
	hw.Gc = NewGamecard(nds7.Irq, "bios/biosnds7.rom")
	hw.Tsc = NewHwTouchScreen()
	hw.Key = NewHwKey()

	hw.Spi = NewHwSpiBus()
	hw.Spi.AddDevice(0, NewHwPowerMan())
	hw.Spi.AddDevice(1, NewHwFirmwareFlash("bios/firmware.bin"))
	hw.Spi.AddDevice(2, hw.Tsc)

	return hw
}

func NewNDSEmulator() *NDSEmulator {
	mem := new(NDSMemory)
	hw := NewNDSHardware(mem)

	sync, err := emu.NewSync(SyncConfig)
	if err != nil {
		panic(err)
	}

	// Register the syncer's logger as global logging function,
	// so that everything will also log the current subsystem
	// status (eg: CPU program counter)
	emu.Log = sync.Log

	e := &NDSEmulator{
		Mem:  mem,
		Hw:   hw,
		Sync: sync,
	}
	e.Sync.SetHSyncCallback(e.hsync)

	return e
}

func (emu *NDSEmulator) StartDebugger() {
	emu.dbg = debugger.New([]debugger.Cpu{nds7.Cpu, nds9.Cpu}, emu.Sync)
	emu.dbg.AddBreakpoint(0x02305796)

	go emu.dbg.Run()
}

func (emu *NDSEmulator) DebugBreak(msg string) {
	if emu.dbg != nil {
		emu.dbg.Break(msg)
	} else {
		log.Error(msg)
		log.Fatal("debugging breakpoint, aborting")
	}
}

func (emu *NDSEmulator) hsync(x, y int) {
	emu.Hw.Lcd9.SyncEvent(x, y)
	emu.Hw.Lcd7.SyncEvent(x, y)

	if y == 0 && x == 0 {
		emu.Hw.E2d[0].BeginFrame()
		emu.Hw.E2d[1].BeginFrame()
	}

	if y < 192 {
		if x == 0 {
			emu.beginLine(y)
		} else if x == cHBlankFirstDot {
			emu.endLine(y)
		}
	}

	if y == 192 && x == 0 {
		emu.Hw.E2d[0].EndFrame()
		emu.Hw.E2d[1].EndFrame()
	}
}

func (emu *NDSEmulator) RunOneFrame(screen gfx.Buffer) {
	emu.screen = screen
	emu.Sync.RunOneFrame()
	log.Infof("chip id: %08x %08x", nds9.Bus.Read32(0x27FF800), nds9.Bus.Read32(0x27FF804))
}

func (emu *NDSEmulator) beginLine(y int) {
	ya := y + 192 + 90
	emu.Hw.E2d[0].BeginLine(emu.screen.Line(ya))

	yb := y
	emu.Hw.E2d[1].BeginLine(emu.screen.Line(yb))
}

func (emu *NDSEmulator) endLine(y int) {
	emu.Hw.E2d[0].EndLine()
	emu.Hw.E2d[1].EndLine()
}
