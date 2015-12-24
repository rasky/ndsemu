package main

import "ndsemu/emu"

type NDSEmulator struct {
	Sync *emu.Sync
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
