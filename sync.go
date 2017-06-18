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

const (
	// Audio Frequency. In most games the exact frequency does not
	// really matter (within some reasonable bounds), but a few games
	// stream musics into a circular buffer, using VMatch IRQ to syncrhonize
	// with audio; in those case, we need a perfect audio frequency
	// otherwise we can hear statics in audio. Examples: Animal Crossing
	// (title screen), Who Wants To Be A Millionaire (voices).

	// This value is probably not precise, but it works for now
	// (at least until we don't understand why the formula below
	// requires a magic adjustment)
	cAudioFreq = 32768

	// Calculate how much the audio timers increment for every
	// sound tick. Theoretically, the formula below is correct
	// (modulo truncation) but we need to add a magic constant to mostly fix
	// statics in some games, and it's still not perfect, so we're not
	// fully understanding this yet.
	cAudioBugAdjust     = 2
	cTimerStepPerSample = uint32((cEmuClock / 2 / cAudioFreq) + cAudioBugAdjust)
)

// NDS SYNC

var NdsLcdConfig = HwLcdConfig{
	VBlankFirstLine: 192,
	VBlankLastLine:  261,
	HBlankFirstDot:  268,
}

var NdsSyncConfig = &emu.SyncConfig{
	MainClock:       cBusClock,
	DotClockDivider: 6, // Dot Clock = 5.585664 MHz
	HDots:           355,
	VDots:           263,

	// Sync at the beginning of each line, and at hblank
	HSyncs: []int{0, NdsLcdConfig.HBlankFirstDot},
}

// GBA SYNC

var GbaLcdConfig = HwLcdConfig{
	VBlankFirstLine: 160,
	VBlankLastLine:  226,
	HBlankFirstDot:  251,
}

var GbaSyncConfig = &emu.SyncConfig{
	MainClock:       cBusClock / 2, // this is what happens on NDS hardware
	DotClockDivider: 4,
	HDots:           308,
	VDots:           228,

	// Sync at the beginning of each line, and at hblank
	HSyncs: []int{0, GbaLcdConfig.HBlankFirstDot},
}
