package main

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/hwio"

	"gopkg.in/Sirupsen/logrus.v0"
)

const cTimerClock = cBusClock

type HwTimer struct {
	Reload  hwio.Reg16 `hwio:"offset=0x0,rcb,wcb"`
	Control hwio.Reg16 `hwio:"offset=0x2,rwmask=0xC7,wcb"`
	counter uint16

	name   string
	cycles int64
	next   *HwTimer
	irqt   bool
	sync   int64
}

func (t *HwTimer) running() bool { return t.Control.Value&0x80 != 0 }
func (t *HwTimer) irq() bool     { return t.Control.Value&0x40 != 0 }
func (t *HwTimer) countup() bool { return t.Control.Value&0x04 != 0 }
func (t *HwTimer) scaler() int {
	switch t.Control.Value & 3 {
	case 0:
		return 1
	case 1:
		return 64
	case 2:
		return 256
	case 3:
		return 1024
	default:
		panic("unreachable")
	}
}

func (t *HwTimer) log() *logrus.Entry {
	return emu.Log().WithField("name", t.name)
}

func (t *HwTimer) reschedule() {
	if t.sync != 0 {
		Emu.Sync.CancelSync(t.sync)
		t.sync = 0
	}

	if t.running() && t.irq() {
		ticks := 0x10000 - int(t.counter)
		t.sync = t.cycles + int64(ticks*t.scaler())
		// t.log().WithFields(logrus.Fields{
		// 	"now":  Emu.Sync.Cycles(),
		// 	"now2": t.cycles,
		// 	"when": t.sync,
		// }).Infof("scheduling sync")
		Emu.Sync.ScheduleSync(t.sync)
	}
}

func (t *HwTimer) WriteRELOAD(_, val uint16) {
	t.log().WithField("val", fmt.Sprintf("%04x", val)).Info("[timers] write reload")
}

func (t *HwTimer) WriteCONTROL(old, val uint16) {
	t.Control.Value = old
	wasrunning := t.running()

	t.Run(Emu.Sync.Cycles())

	t.Control.Value = val
	if !wasrunning && t.running() {
		// 0->1 transition: reload the counter value
		t.counter = t.Reload.Value
	}
	t.reschedule()
}

func (t *HwTimer) ReadRELOAD(_ uint16) uint16 {
	// Reading reload actually accesses the current counter
	t.Run(Emu.Sync.Cycles())
	return t.counter
}

// Handle an overflow event
func (t *HwTimer) overflow() {
	t.counter = t.Reload.Value
	if t.next != nil && t.next.countup() {
		t.next.up()
	}
	if t.irq() {
		if t.irqt {
			t.log().Warnf("double timer reload=%04x scaler=%d", t.Reload.Value, t.scaler())
			panic("double timer overflow")
		} else {
			// t.log().WithField("cycles", Emu.Sync.Cycles()).Infof("[timers] overflow")
		}
		t.irqt = true
	}
	t.reschedule()
}

// Increment the timer by one; this is meant to be used only
// on countup timers
func (t *HwTimer) up() {
	if !t.running() {
		// we don't know whether a countup timer should respect
		// the start/stop bit. Probably it does, but we need to
		// check.
		panic("investigate: up called on stopped timer")
	}
	if !t.countup() {
		panic("assert: up called on wrong timer")
	}
	t.counter++
	if t.counter == 0 {
		t.overflow()
	}
}

func (t *HwTimer) Run(target int64) {
	// (Small) negative timeline is possible because of small rounding errors. For
	// instance, CPUs might stop a little bit earlier than requested, and this
	// in turn would cause timers to be called with negative offsets. Nothing to
	// worry about.
	if target < t.cycles {
		return
	}

	if !t.running() {
		t.cycles = target
		return
	}
	// countup timers are emulated in the context of the base timer
	if t.countup() {
		t.cycles = target
		return
	}

	scaler := int64(t.scaler())
	elapsed := target - t.cycles
	if elapsed < 0 {
		elapsed = 0
	}
	elapsed /= scaler

	for elapsed > 0 {
		span := 0x10000 - int64(t.counter)
		if span > elapsed {
			t.cycles += elapsed * scaler
			t.counter += uint16(elapsed)
			elapsed = 0
		} else {
			t.cycles += span * scaler
			t.overflow()
			elapsed -= span
		}
	}
}

type HwTimers struct {
	Irq    *HwIrq
	Timers [4]HwTimer
}

func NewHWTimers(name string, irq *HwIrq) *HwTimers {
	t := &HwTimers{
		Irq: irq,
	}
	t.Reset()
	t.SetName(name)
	return t
}

func (t *HwTimers) SetName(prefix string) {
	for i := range t.Timers {
		t.Timers[i].name = fmt.Sprintf("%s-%d", prefix, i)
	}
}

func (t *HwTimers) Reset() {
	for i := range t.Timers {
		t.Timers[i] = HwTimer{}
		hwio.MustInitRegs(&t.Timers[i])
		if i != 3 {
			t.Timers[i].next = &t.Timers[i+1]
		}
	}
}

func (t *HwTimers) Frequency() emu.Fixed8 {
	return emu.NewFixed8(cTimerClock)
}

func (t *HwTimers) Cycles() int64 {
	// All timers are (mostly) aligned, so just use any to return the
	// current time. The only possible misalignment could be because
	// of rounding errors
	return t.Timers[0].cycles
}

func (t *HwTimers) Run(target int64) {
	for i := 0; i < 4; i++ {
		t.Timers[i].Run(target)
		if t.Timers[i].irqt {
			t.Timers[i].irqt = false
			t.Irq.Raise(IrqTimer0 << uint(i))
		}
	}
}
