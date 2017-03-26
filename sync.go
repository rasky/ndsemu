package main

import (
	"ndsemu/emu"
)

const (
	cBusClock  = int64(0x1FF61FE) // 33.513982 Mhz
	cNds7Clock = cBusClock
	cNds9Clock = cBusClock * 2

	cEmuClock  = cBusClock
	cAudioFreq = 32760 // should be 32768, but we need a multiple of FPS
)

var SyncConfig = emu.SyncConfig{
	MainClock:       cBusClock,
	DotClockDivider: 6, // Dot Clock = 5.585664 MHz
	HDots:           355,
	VDots:           263,

	// Sync at the beginning of each line, and at hblank
	HSyncs: []int{0, cHBlankFirstDot},
}
