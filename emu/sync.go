package emu

import (
	"errors"
	"sort"

	log "gopkg.in/Sirupsen/logrus.v0"
)

// A CPU subsystem. This is the basic interface that the syncing engine uses
// to communicate with CPU cores.
type Cpu interface {
	// The frequency at which the CPU core runs
	Frequency() Fixed8

	// Do a hardware reset
	Reset()

	// Run a single instruction step of emulation. Returns the new internal
	// clock after the instruction was performed.
	Step() int64
}

// A non-CPU subsystem is a frequency-based emulation component. It can be
// anything: a CPU, a video engine, a sound engine, a hardware timer,
// a serial port, etc.
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

type subsystemCpu struct {
	Cpu
	target int64
	cycles int64
}

func (cpu *subsystemCpu) Retarget(target int64) {
	if cpu.target > target {
		cpu.target = target
	}
}

func (cpu *subsystemCpu) Run(target int64) {
	cpu.target = target
	for cpu.cycles < cpu.target {
		cpu.cycles = cpu.Cpu.Step()
	}
}

func (cpu *subsystemCpu) Cycles() int64 {
	return cpu.cycles
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

func (s syncSubsystem) Retarget(target int64) {
	if cpu, ok := s.Subsystem.(*subsystemCpu); ok {
		t := NewFixed8(target).DivFixed(s.scaler)
		cpu.Retarget(t.ToInt64())
	}
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
	subCpus     []syncSubsystem
	subOthers   []syncSubsystem
	reqSyncs    []int64
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

func (s *Sync) AddCpu(cpu Cpu) {
	s.subCpus = append(s.subCpus, syncSubsystem{
		Subsystem: &subsystemCpu{Cpu: cpu},
		scaler:    s.mainClock.DivFixed(cpu.Frequency()),
	})
}

func (s *Sync) AddSubsystem(sub Subsystem) {
	s.subOthers = append(s.subOthers, syncSubsystem{
		Subsystem: sub,
		scaler:    s.mainClock.DivFixed(sub.Frequency()),
	})
}

func (s *Sync) Reset() {
	for _, sub := range append(s.subCpus, s.subOthers...) {
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

// Schedule a new one-shot sync point in the future. This can be useful to make
// sure all subsystems will be synced at this specific point, possibly aligned
// with a IRQ generation or a similar event.
func (s *Sync) ScheduleSync(when int64) {
	if when < s.Cycles() {
		log.Info("sync in the past:", when, s.Cycles())
		panic("scheduling sync in the past")
	}

	for i := range s.reqSyncs {
		if s.reqSyncs[i] > when {
			s.reqSyncs = append(s.reqSyncs, 0)
			copy(s.reqSyncs[i+1:], s.reqSyncs[i:])
			s.reqSyncs[i] = when
			if i == 0 && s.runningSub != nil {
				// If this is the earliest sync point to date, and there is
				// a subsytem running, retarget it to make sure it doesn't
				// skip the sync point.
				s.runningSub.Retarget(when)
			}
			return
		} else if s.reqSyncs[i] == when {
			return
		}
	}
	s.reqSyncs = append(s.reqSyncs, when)
	if len(s.reqSyncs) == 1 && s.runningSub != nil {
		// As above: if it's the earliest, retarget the running subsystem
		s.runningSub.Retarget(when)
	}
}

func (s *Sync) CancelSync(when int64) {
	for i := range s.reqSyncs {
		if s.reqSyncs[i] == when {
			s.reqSyncs = append(s.reqSyncs[:i], s.reqSyncs[i+1:]...)
			return
		} else if s.reqSyncs[i] > when {
			return
		}
	}
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
	if s.runningSub != nil {
		panic("reentrancy")
	}

	// First go through CPUs
	next := int64(0)
	for next < target {
		next = target

		// See if there are additional one-shot sync requests scheduled
		for len(s.reqSyncs) > 0 && next > s.reqSyncs[0] {
			next = s.reqSyncs[0]
			s.reqSyncs = s.reqSyncs[1:]
		}

		for idx := range s.subCpus {
			s.runningSub = &s.subCpus[idx]
			s.runningSub.Run(next)
			cycles := s.runningSub.Cycles()
			if next > cycles {
				next = cycles
			}
			s.runningSub = nil
		}

		for idx := range s.subOthers {
			s.runningSub = &s.subOthers[idx]
			s.runningSub.Run(next)
			s.runningSub = nil
		}
	}

	s.runningSub = nil
	s.cycles = target
}

func (s *Sync) CurrentSubsystem() Subsystem {
	if s.runningSub == nil {
		return nil
	}
	return s.runningSub.Subsystem
}

func (s *Sync) CurrentCpu() Cpu {
	if s.runningSub == nil {
		return nil
	}
	if cpu, ok := s.runningSub.Subsystem.(*subsystemCpu); ok {
		return cpu.Cpu
	}
	return nil
}
