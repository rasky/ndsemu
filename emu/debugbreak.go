package emu

import (
	log "ndsemu/emu/logger"
)

var BreakFunc func(msg string)

func init() {
	BreakFunc = func(msg string) {
		log.ModEmu.FatalZ(msg).End()
	}
}

// DebugBreak is meant to be called in all cases the emulator
// wants to notify a breaking condition that needs debugging.
//
// The function does nothing more than call BreakFunc, if
// it is not nil. This allows emulators to configure different
// ways to handle this kind of conditions without creating
// interdependency.
func DebugBreak(msg string) {
	if BreakFunc != nil {
		BreakFunc(msg)
	}
}
