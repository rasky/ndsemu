package main

import (
	"ndsemu/emu/hwio"
)

type HwSoundChannel struct {
	SndCnt hwio.Reg32 `hwio:"offset=0x00"`
	SndSad hwio.Reg32 `hwio:"offset=0x04,rwmask=0x07FFFFFF"`
	SndTmr hwio.Reg16 `hwio:"offset=0x08"`
	SndPnt hwio.Reg16 `hwio:"offset=0x0A"`
	SndLen hwio.Reg32 `hwio:"offset=0x0C,rwmask=0x001FFFFF"`
}

type HwSound struct {
	Ch [16]HwSoundChannel

	SndCnt hwio.Reg32 `hwio:"bank=1,offset=0x0"`
	// The NDS7 BIOS brings this register to 0x200 at boot, with a slow loop
	// with delay that takes ~1 second. If we reset it at 0x200, it will just
	// skip everything and the emulator will boot faster.
	SndBias    hwio.Reg32 `hwio:"bank=1,offset=0x4,reset=0x200,rwmask=0x3FF"`
	SndCap0Cnt hwio.Reg8  `hwio:"bank=1,offset=0x8,rwmask=0x8F"`
	SndCap1Cnt hwio.Reg8  `hwio:"bank=1,offset=0x9,rwmask=0x8F"`
}

func NewHwSound() *HwSound {
	snd := new(HwSound)
	for i := 0; i < 16; i++ {
		hwio.MustInitRegs(&snd.Ch[i])
	}
	hwio.MustInitRegs(snd)
	return snd
}
