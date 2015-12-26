package emu

import (
	"errors"
	"sort"
)

// A subsystem is a frequency-based emulation component. It can be anything: a
// CPU, a video engine, a sound engine, a hardware timer, a serial port, etc.
//
// Each subsystem declares the frequency at which it wants to run, through the
// Frequency() method. It then must implement a Run() method that advances
// the emulation up until the specified instant.
type Subsystem interface {
	Frequency() Fixed8

	// Do a hardware reset
	Reset()

	// Return the current internal timer. Notice that, for some subsystems
	// (eg: CPUs), it will be common for this function to be called while the
	// subsytem is running (that is, from the callees of a Run() call). For
	// accurate emulation, Cycles() should thus return also updated values
	// while Run() is executing.
	Cycles() int64

	// Run the emulation of the subsystem until reaching the specified
	// absolute number of cycles since last reset. Do nothing if the
	// target time is in the past.
	// It is acceptable that a subsystem might be unable to emulate to the
	// specific instant, and might stop emulation a little bit before or
	// after that instant.
	Run(targetCycles int64)
}

// syncSubsystem wraps a subsystem saving its frequency scaler, and overrides
// the Cycles() and Run() calls so that they automatically convert the cycles
// count to the correct frequency
type syncSubsystem struct {
	Subsystem
	scaler Fixed8
}

func (s syncSubsystem) Cycles() int64 {
	return s.scaler.Mul(s.Subsystem.Cycles()).ToInt64()
}

func (s syncSubsystem) Run(target int64) {
	t := NewFixed8(target).DivFixed(s.scaler)
	s.Subsystem.Run(t.ToInt64())
}

// SyncConfig contains the configuration for the syncing system
type SyncConfig struct {
	// Main oscillator. This is the basic for subsystem timing
	MainClock int64

	// Dot clock divider. This is the frequency at which each dot is drawn
	DotClockDivider int

	// Number of horizontal and vertical dots of the screen
	HDots, VDots int

	// The horizontal positions at which we want to generate an emulation sync
	// (across all subsystems). Each element of this slice is expected to be
	// between 0 and HDots.
	//
	// For instance, if the slice is [0], the emulation is synced once per line,
	// at the beginning of that line. If the slice is empty, the emulation is
	// not synced every line (that is, frame-based emulation is performed).
	// Putting more values allow to sub-line syncing accuracy.
	HSyncs []int

	// The vertical positions at which we want to generate an emulation sync
	// (across all subsystems). See HSyncs for some examples.
	//
	// The VSync is always generated at the beginning of the line. If HSyncs
	// contains [0], the VSync event will be generated before the HSync event.
	VSyncs []int

	// If non-nil, this function will be called at each HSync.
	HSync func(x, y int)

	// If non-nil, this function will be called at each VSync.
	// NOTE: x will always be zero, as VSync events are generated at the
	// beginning of each specified line. It is still passed to use the
	// same signature of HSync and thus allow reusing the same function.
	VSync func(x, y int)
}

type syncEventType int

const (
	eventTypeVSync syncEventType = iota
	eventTypeHSync
)

type syncEvent struct {
	Cycles int64
	Type   syncEventType
	X, Y   int
}

type sortByCycles []syncEvent

func (s sortByCycles) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortByCycles) Less(i, j int) bool { return s[i].Cycles < s[j].Cycles }
func (s sortByCycles) Len() int           { return len(s) }

type Sync struct {
	cfg         SyncConfig
	mainClock   Fixed8
	lineCycles  int64
	frameCycles int64
	frameSyncs  []syncEvent
	runningSub  *syncSubsystem
	subs        []syncSubsystem
	cycles      int64
}

func NewSync(cfg SyncConfig) (*Sync, error) {
	// Since configuration errors are programmer's errors, let's just
	// panic on those
	for _, h := range cfg.HSyncs {
		if h >= cfg.HDots {
			return nil, errors.New("invalid hsync point bigger than screen width")
		}
	}
	for _, v := range cfg.VSyncs {
		if v >= cfg.VDots {
			return nil, errors.New("invalid vsync point bigger than screen height")
		}
	}

	sync := &Sync{
		cfg:       cfg,
		mainClock: NewFixed8(cfg.MainClock),
	}
	sync.calc()
	return sync, nil
}

func (s *Sync) calc() {
	s.lineCycles = int64(s.cfg.DotClockDivider * s.cfg.HDots)
	s.frameCycles = s.lineCycles * int64(s.cfg.VDots)

	s.frameSyncs = nil
	for _, y := range s.cfg.VSyncs {
		evt := syncEvent{
			Cycles: s.lineCycles * int64(y),
			Type:   eventTypeVSync,
			X:      0,
			Y:      y,
		}
		s.frameSyncs = append(s.frameSyncs, evt)
	}

	for _, x := range s.cfg.HSyncs {
		for y := 0; y < s.cfg.VDots; y++ {
			evt := syncEvent{
				Cycles: s.lineCycles*int64(y) + int64(x*s.cfg.DotClockDivider),
				Type:   eventTypeHSync,
				X:      x,
				Y:      y,
			}
			s.frameSyncs = append(s.frameSyncs, evt)
		}
	}

	// Use stable stort so that vsyncs are generated before hsyncs (as they
	// are created before in the slice)
	sort.Stable(sortByCycles(s.frameSyncs))
}

// Returns the number of frames per second at which the emulation runs. This can
// be computed by the sync configuration (that is, given the master clock, the
// dot clock and the screen resolution). It is returned as a fixed value, but
// most users will want to round it.
func (s *Sync) Fps() Fixed8 {
	return s.mainClock.Div(s.frameCycles)
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

// Return the current clock. If this function is called from within a subsystem,
// it returns that subsystem's vision of the current timing. Especially for CPUs,
// it is thus important that Subsytem.Cycles() calls within Subsystem.Run()
// calls are updated correctly.
func (s *Sync) Cycles() int64 {
	if s.runningSub != nil {
		return s.runningSub.Cycles()
	}
	return s.cycles
}

// Return the (x,y) dot position of the emulation in the current frame. The
// returned values are calculated from the result of Sync.Cycles() and is thus
// always updated.
func (s *Sync) DotPos() (int, int) {
	clk := s.Cycles() % s.frameCycles
	y := clk / int64(s.lineCycles)
	x := clk % int64(s.lineCycles) / int64(s.cfg.DotClockDivider)
	return int(x), int(y)
}

// Align all subsystems to the one that was emulated further.
func (s *Sync) Sync() {
	target := s.subs[0].Cycles()
	for _, sub := range s.subs[1:] {
		t2 := sub.Cycles()
		if target < t2 {
			target = t2
		}
	}
	s.RunUntil(target)
}

// Advance the emulation of exactly one frame. This will panic if the emulation
// is not exactly at a frame boundary, because otherwise DotPos() would return
// an invalid value.
func (s *Sync) RunOneFrame() {
	baseclk := s.cycles
	if baseclk%s.frameCycles != 0 {
		panic("RunOneFrame called while not a frame boundary")
	}

	// Go through all sync events that have been configured for a frame.
	// The events are already sorted in chronological order.
	for _, evt := range s.frameSyncs {
		s.RunUntil(baseclk + evt.Cycles)

		switch evt.Type {
		case eventTypeVSync:
			if s.cfg.VSync != nil {
				s.cfg.VSync(evt.X, evt.Y)
			}
		case eventTypeHSync:
			if s.cfg.HSync != nil {
				s.cfg.HSync(evt.X, evt.Y)
			}
		default:
			panic("unreachable")
		}
	}

	// Complete the frame
	s.RunUntil(baseclk + s.frameCycles)
}

func (s *Sync) RunUntil(target int64) {
	if target < s.cycles {
		panic("assert: invalid target cycles")
	}

	// Main loop across subsystems. Preserve the reference to the previously
	// running subsystem to allow for full reentrancy.
	oldrs := s.runningSub
	for idx := range s.subs {
		s.runningSub = &s.subs[idx]
		// While reentring, avoid running the same subsystem. This is just
		// a one-level fix to avoid triggering gratuitious bugs in
		// non-reentrant subsystem code. Obviously, it cannot protect
		// against aggressive/bugged used of reentrancy.
		if s.runningSub != oldrs {
			s.runningSub.Run(target)
		}
	}
	s.runningSub = oldrs

	s.cycles = target
}
