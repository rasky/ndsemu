package main

import (
	"ndsemu/emu"
)

const (
	cBusClock  = int64(0x1FF61FE) // 33.513982 Mhz
	cNds7Clock = cBusClock
	cNds9Clock = cBusClock * 2

	cEmuClock = cBusClock
)

var SyncConfig = emu.SyncConfig{
	MainClock:       cBusClock,
	DotClockDivider: 6, // Dot Clock = 5.585664 MHz
	HDots:           355,
	VDots:           263,

	// We sync multiple times per line for now; revisit this once we have more
	// accurate timing in ARM core
	HSyncs: []int{0, 355 / 4, 355 * 2 / 4, 355 * 3 / 4},
}
