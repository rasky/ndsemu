package main

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/debugger"
	"ndsemu/emu/gfx"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDSMemory struct {
	Vram [656 * 1024]byte
}

type NDSHardware struct {
	E2d  [2]*HwEngine2d
	Lcd  *HwLcd
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

	hw.Mc = NewMemoryController(nds9, nds7, mem.Vram[:])
	hw.E2d[0] = NewHwEngine2d(0, hw.Mc)
	hw.E2d[1] = NewHwEngine2d(1, hw.Mc)
	hw.Ipc = NewHwIpc(nds9.Irq, nds7.Irq)
	hw.Lcd = NewHwLcd(nds9.Irq, nds7.Irq)
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

	emu := &NDSEmulator{
		Mem:  mem,
		Hw:   hw,
		Sync: sync,
	}
	emu.Sync.SetHSyncCallback(emu.hsync)
	return emu
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

func (emu *NDSEmulator) Log() *log.Entry {
	sub := emu.Sync.CurrentCpu()
	if c7, ok := sub.(*NDS7); ok {
		return log.WithField("pc7", c7.Cpu.GetPC())
	}
	if c9, ok := sub.(*NDS9); ok {
		return log.WithField("pc9", c9.Cpu.GetPC())
	}
	return log.WithField("sub", fmt.Sprintf("%T", emu.Sync.CurrentSubsystem()))
}

func (emu *NDSEmulator) hsync(x, y int) {
	emu.Hw.Lcd.SyncEvent(x, y)

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
