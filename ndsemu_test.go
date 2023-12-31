package main

import (
	"ndsemu/emu/gfx"
	log "ndsemu/emu/logger"
	"os"

	"testing"
)

func BenchmarkCpuSpeed(b *testing.B) {
	screen := gfx.NewBufferMem(256, 192+90+192)
	log.Disable()

	f, err := os.CreateTemp("", "")
	if err != nil {
		b.Fatal(err)
	}
	f.Close()
	defer os.Remove(f.Name())

	for i := 0; i < b.N; i++ {
		Emu = NewNDSEmulator(f.Name())
		Emu.Hw.Gc.MapCartFile("roms/phoenixwright.nds")
		Emu.Hw.Ff.MapFirmwareFile("bios/firmware.bin")
		Emu.Hw.Rtc.ResetDefaults()

		for j := 0; j < 300; j++ {
			Emu.RunOneFrame(screen)
		}
	}
}
