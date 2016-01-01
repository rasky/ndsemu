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
	E2d [2]*HwEngine2d
	Lcd *HwLcd
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

	hw.E2d[0] = NewHwEngine2d(0, mem.Vram[:])
	hw.E2d[1] = NewHwEngine2d(1, mem.Vram[:])

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
	// emu.dbg.AddBreakpoint(0x0202ae82)
	// emu.dbg.AddBreakpoint(0x0200095c)
	// emu.dbg.AddWatchpoint(0x02042b80)
	emu.dbg.AddBreakpoint(0x02042590)
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

	if x == 0 && y < 192 {
		emu.drawLine(y)
	}
}

func (emu *NDSEmulator) RunOneFrame(screen gfx.Buffer) {
	emu.screen = screen
	emu.Sync.RunOneFrame()
}

func (emu *NDSEmulator) drawLine(y int) {
	emu.Hw.E2d[0].DrawLine(y, emu.screen.LineAsSlice(y))

	yb := y + 192 + 90
	emu.Hw.E2d[1].DrawLine(y, emu.screen.LineAsSlice(yb))
}
