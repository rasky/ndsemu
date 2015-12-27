package main

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/debugger"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type NDSEmulator struct {
	Sync *emu.Sync

	dbg *debugger.Debugger
}

var Emu *NDSEmulator

func NewNDSEmulator() *NDSEmulator {
	var err error
	e := new(NDSEmulator)
	e.Sync, err = emu.NewSync(SyncConfig)
	if err != nil {
		panic(err)
	}
	return e
}

func (emu *NDSEmulator) StartDebugger() {
	emu.dbg = debugger.New([]debugger.Cpu{nds7.Cpu, nds9.Cpu}, emu.Sync)
	emu.dbg.AddBreakpoint(0x03803FCC)
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
