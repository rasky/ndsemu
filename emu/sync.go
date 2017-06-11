package emu

import (
	"errors"
	"fmt"
	"sort"

	"ndsemu/emu/fixed"
	log "ndsemu/emu/logger"
)

// A non-CPU subsystem is a frequency-based emulation component. It can be
// anything: a CPU, a video engine, a sound engine, a hardware timer,
// a serial port, etc.
//
// Each subsystem declares the frequency at which it wants to run, through the
// Frequency() method. It then must implement a Run() method that advances
// the emulation up until the specified instant.
type Subsystem interface {
	Frequency() fixed.F8

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

// A CPU subsystem. This is the basic interface that the syncing engine uses
// to communicate with CPU cores.
type Cpu interface {
	Subsystem

	// Change the current running target to the new specified amount.
	//
	// This function can be invoked through reentrancy while the cpu is
	// running (so within the context of a Run()) and is used to notify
	// the cpu that it should change the previously specified target
	// cycles. A common case for this to happen is when the cpu triggered
	// a condition that requires a closer sync that what was initially thought;
	// for instance, a timer might need trigger a critical IRQ that must be
	// served immediately.
	Retarget(newTarget int64)

	// Get the program counter. This function is only used for the purpose
	// of logging, so that Sync can expose a logger that always the log
	// the current CPU's program counter, which is of course very useful.
	GetPC() uint32
}

// syncSubsystem wraps a subsystem saving its frequency scaler, and overrides
// the Cycles() and Run() calls so that they automatically convert the cycles
// count to the correct frequency
type syncSubsystem struct {
	Subsystem
	scaler fixed.F8
	name   string
}

func (s syncSubsystem) Cycles() int64 {
	return s.scaler.Mul(s.Subsystem.Cycles()).ToInt64()
}

func (s syncSubsystem) Retarget(target int64) {
	if cpu, ok := s.Subsystem.(Cpu); ok {
		t := fixed.NewF8(target).DivFixed(s.scaler)
		cpu.Retarget(t.ToInt64())
	}
}

func (s syncSubsystem) Run(target int64) {
	t := fixed.NewF8(target).DivFixed(s.scaler)
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

type syncPointType int

const (
	pointTypeVSync syncPointType = iota
	pointTypeHSync
)

type syncPoint struct {
	Cycles int64
	Type   syncPointType
	X, Y   int
}

type syncEvent struct {
	When int64
	Cb   func()
}

type Sync struct {
	cfg         *SyncConfig
	mainClock   fixed.F8
	lineCycles  int64
	frameCycles int64
	frameSyncs  []syncPoint
	runningSub  *syncSubsystem
	subCpus     []syncSubsystem
	subOthers   []syncSubsystem
	events      []syncEvent
	cycles      int64
	frames      int64
}

func NewSync(cfg *SyncConfig) (*Sync, error) {
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
		mainClock: fixed.NewF8(cfg.MainClock),
	}
	sync.calc()
	return sync, nil
}

func (s *Sync) SetConfig(cfg *SyncConfig) {
	s.cfg = cfg
	s.calc()
}

func (s *Sync) calc() {
	s.lineCycles = int64(s.cfg.DotClockDivider * s.cfg.HDots)
	s.frameCycles = s.lineCycles * int64(s.cfg.VDots)

	s.frameSyncs = nil
	for _, y := range s.cfg.VSyncs {
		pt := syncPoint{
			Cycles: s.lineCycles * int64(y),
			Type:   pointTypeVSync,
			X:      0,
			Y:      y,
		}
		s.frameSyncs = append(s.frameSyncs, pt)
	}

	for _, x := range s.cfg.HSyncs {
		for y := 0; y < s.cfg.VDots; y++ {
			pt := syncPoint{
				Cycles: s.lineCycles*int64(y) + int64(x*s.cfg.DotClockDivider),
				Type:   pointTypeHSync,
				X:      x,
				Y:      y,
			}
			s.frameSyncs = append(s.frameSyncs, pt)
		}
	}

	// Use stable stort so that vsyncs are generated before hsyncs (as they
	// are created before in the slice)
	sort.SliceStable(s.frameSyncs, func(i, j int) bool {
		return s.frameSyncs[i].Cycles < s.frameSyncs[j].Cycles
	})
}

// Returns the number of frames per second at which the emulation runs. This can
// be computed by the sync configuration (that is, given the master clock, the
// dot clock and the screen resolution). It is returned as a fixed value, but
// most users will want to round it.
func (s *Sync) Fps() fixed.F8 {
	return s.mainClock.Div(s.frameCycles)
}

func (s *Sync) AddCpu(cpu Cpu, name string) {
	s.subCpus = append(s.subCpus, syncSubsystem{
		Subsystem: cpu,
		name:      name,
		scaler:    s.mainClock.DivFixed(cpu.Frequency()),
	})
}

func (s *Sync) AddSubsystem(sub Subsystem, name string) {
	s.subOthers = append(s.subOthers, syncSubsystem{
		Subsystem: sub,
		name:      name,
		scaler:    s.mainClock.DivFixed(sub.Frequency()),
	})
}

func (s *Sync) SetHSyncCallback(cb func(int, int)) {
	s.cfg.HSync = cb
}

func (s *Sync) SetVSyncCallback(cb func(int, int)) {
	s.cfg.VSync = cb
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

// Return the number of cycles missing to reach the (x,y) dot position. Just
// like DotPos(), it is using Sync.Cycles() so it should always be correct.
func (s *Sync) DotPosDistance(x, y int) int64 {
	cx, cy := s.DotPos()

	t := y*s.cfg.HDots + x
	ct := cy*s.cfg.HDots + cx

	dist := int64((t - ct) * s.cfg.DotClockDivider)
	if dist < 0 {
		dist += s.frameCycles
	}
	return dist
}

// Schedule a new one-shot sync point in the future. This can be useful to make
// sure all subsystems will be synced at this specific point, possibly aligned
// with a IRQ generation or a similar event.
func (s *Sync) ScheduleEvent(when int64, cb func()) {
	if when < s.Cycles() {
		log.ModEmu.PanicZ("sync in the past").Int64("when", when).Int64("cycles", s.Cycles()).End()
	}

	for i := range s.events {
		if s.events[i].When > when {
			s.events = append(s.events, syncEvent{})
			copy(s.events[i+1:], s.events[i:])
			s.events[i].When = when
			s.events[i].Cb = cb
			if i == 0 && s.runningSub != nil {
				// If this is the earliest sync point to date, and there is
				// a subsytem running, retarget it to make sure it doesn't
				// skip the sync point.
				s.runningSub.Retarget(when)
			}
			return
		} else if s.events[i].When == when && false {
			return
		}
	}
	s.events = append(s.events, syncEvent{When: when, Cb: cb})
	if len(s.events) == 1 && s.runningSub != nil {
		// As above: if it's the earliest, retarget the running subsystem
		s.runningSub.Retarget(when)
	}
}

func (s *Sync) ScheduleSync(when int64) {
	s.ScheduleEvent(when, nil)
}

func (s *Sync) CancelSync(when int64) {
	for i := range s.events {
		if s.events[i].When == when && s.events[i].Cb == nil {
			s.events = append(s.events[:i], s.events[i+1:]...)
			return
		} else if s.events[i].When > when {
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
	for _, pt := range s.frameSyncs {
		s.RunUntil(baseclk + pt.Cycles)

		switch pt.Type {
		case pointTypeVSync:
			if s.cfg.VSync != nil {
				s.cfg.VSync(pt.X, pt.Y)
			}
		case pointTypeHSync:
			if s.cfg.HSync != nil {
				s.cfg.HSync(pt.X, pt.Y)
			}
		default:
			panic("unreachable")
		}
	}

	// Complete the frame
	s.RunUntil(baseclk + s.frameCycles)
	s.frames++
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

		for idx := range s.subCpus {
			if len(s.events) > 0 && next > s.events[0].When {
				next = s.events[0].When
			}

			s.runningSub = &s.subCpus[idx]
			s.runningSub.Run(next)
			s.runningSub = nil
		}

		for idx := range s.subOthers {
			if len(s.events) > 0 && next > s.events[0].When {
				next = s.events[0].When
			}

			s.runningSub = &s.subOthers[idx]
			s.runningSub.Run(next)
			s.runningSub = nil
		}

		for len(s.events) > 0 && next >= s.events[0].When {
			evt := s.events[0]
			s.events = s.events[1:]
			if evt.Cb != nil {
				evt.Cb()
			}
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
	if cpu, ok := s.runningSub.Subsystem.(Cpu); ok {
		return cpu
	}
	return nil
}

// Implement logger.LogContextAdders
func (s *Sync) AddLogContext(entry *log.EntryZ) {
	entry.Int64("_frame", s.frames)
	if cur := s.runningSub; cur != nil {
		if cpu, ok := cur.Subsystem.(Cpu); ok {
			entry.Hex32("pc-"+cur.name, cpu.GetPC())
		} else {
			entry.String("sub", fmt.Sprintf("%T", cur.Subsystem))
		}
	}
}
