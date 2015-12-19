package emu

type Fixed8 int64

func NewFixed8(val int64) Fixed8 {
	return Fixed8(val << 8)
}

func (f Fixed8) ToInt64() int64 {
	return (int64(f) + 0x80) >> 8
}

func (f Fixed8) Mul(mul int64) Fixed8 {
	return f * Fixed8(mul)
}

func (f Fixed8) Div(den int64) Fixed8 {
	return f / Fixed8(den)
}

func (f Fixed8) DivFixed(den Fixed8) Fixed8 {
	return (f << 8) / den
}

func (f Fixed8) MulFixed(mul Fixed8) Fixed8 {
	return (f * mul) >> 8
}

type Subsystem interface {
	Frequency() Fixed8

	// Do a hardware reset
	Reset()

	// Return the current internal timer
	Cycles() int64

	// Run the emulation of the subsystem until reaching the specified
	// absolute number of cycles since last reset. Do nothing if the
	// target time was already emulated.
	Run(targetCycles int64)
}

type syncSubsystem struct {
	Subsystem
	scaler Fixed8
}

type Sync struct {
	mainClock Fixed8
	subs      []syncSubsystem
	cycles    int64
}

func NewSync(mainClock int64) *Sync {
	return &Sync{
		mainClock: NewFixed8(mainClock),
	}
}

func (s *Sync) AddSubsystem(sub Subsystem) {
	s.subs = append(s.subs, syncSubsystem{
		Subsystem: sub,
		scaler:    s.mainClock.DivFixed(sub.Frequency()),
	})
}

func (s *Sync) Reset() {
	for _, sub := range s.subs {
		sub.Reset()
	}
	s.cycles = 0
}

func (s *Sync) Cycles() int64 {
	return s.cycles
}

func (s *Sync) Sync() {
	target := s.subs[0].scaler.Mul(s.subs[0].Cycles())
	for _, sub := range s.subs[1:] {
		t2 := sub.scaler.Mul(sub.Cycles())
		if target < t2 {
			target = t2
		}
	}
	s.Run(target.ToInt64())
}

func (s *Sync) Run(target int64) {
	if target < s.cycles {
		panic("assert: invalid target cycles")
	}

	for _, sub := range s.subs {
		t := NewFixed8(target).DivFixed(sub.scaler)
		sub.Run(t.ToInt64())
	}

	s.cycles = target
}
