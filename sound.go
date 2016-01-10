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
}

func NewHwSound() *HwSound {
	snd := new(HwSound)
	for i := 0; i < 16; i++ {
		hwio.MustInitRegs(&snd.Ch[i])
	}
	return snd
}
