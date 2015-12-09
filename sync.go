package main

const (
	cBusClock  = 0x1FF61FE // 33.513982 Mhz
	cNds7Clock = cBusClock
	cNds9Clock = cBusClock * 2

	cEmuClock = cBusClock
)

type fixed8 int64

func NewFixed8(val int64) fixed8 {
	return fixed8(val << 8)
}

func (f fixed8) ToInt64() int64 {
	return (int64(f) + 0x80) >> 8
}

func (f fixed8) Div(den int64) fixed8 {
	return f / fixed8(den)
}

func (f fixed8) DivFixed(den fixed8) fixed8 {
	return (f << 8) / den
}

type Subsystem interface {
	Frequency() fixed8

	// Do a hardware reset
	Reset()

	// Return the current internal timer
	Cycles() int64

	// Run the emulation of the subsystem until reaching the specified
	// absolute number of cycles since last reset. Do nothing if the
	// target time was already emulated.
	Run(targetCycles int64)
}

type SyncSubsystem struct {
	Subsystem
	scaler fixed8
}

type SyncEmu struct {
	subs   []SyncSubsystem
	cycles int64
}

func (s *SyncEmu) AddSubsystem(sub Subsystem) {
	s.subs = append(s.subs, SyncSubsystem{
		Subsystem: sub,
		scaler:    NewFixed8(cEmuClock).DivFixed(sub.Frequency()),
	})
}

func (s *SyncEmu) Reset() {
	for _, sub := range s.subs {
		sub.Reset()
	}
	s.cycles = 0
}

func (s *SyncEmu) Cycles() int64 {
	return s.cycles
}

func (s *SyncEmu) Run(target int64) {
	if target < s.cycles {
		panic("assert: invalid target cycles")
	}

	for _, sub := range s.subs {
		t := NewFixed8(target).DivFixed(sub.scaler)
		sub.Run(t.ToInt64())
	}

	s.cycles = target
}
