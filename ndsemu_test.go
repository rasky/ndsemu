package main

import (
	"ndsemu/emu/gfx"
	log "ndsemu/emu/logger"

	"testing"
)

func BenchmarkCpuSpeed(b *testing.B) {
	screen := gfx.NewBufferMem(256, 192+90+192)
	log.Disable()

	for i := 0; i < b.N; i++ {
		Emu = NewNDSEmulator()
		Emu.Hw.Gc.MapCartFile("roms/phoenixwright.nds")

		for j := 0; j < 300; j++ {
			Emu.RunOneFrame(screen)
		}
	}
}
