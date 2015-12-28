package main

import (
	"fmt"
	"ndsemu/emu"

	"gopkg.in/Sirupsen/logrus.v0"
)

const cTimerClock = cBusClock

type HwTimer struct {
	counter uint16
	reload  uint16
	control uint16

	name   string
	cycles int64
	next   *HwTimer
	irqt   bool
	sync   int64
}

func (t *HwTimer) running() bool { return t.control&0x80 != 0 }
func (t *HwTimer) irq() bool     { return t.control&0x40 != 0 }
func (t *HwTimer) countup() bool { return t.control&0x04 != 0 }
func (t *HwTimer) scaler() int {
	switch t.control & 3 {
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
	return Emu.Log().WithField("name", t.name)
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

func (t *HwTimer) WriteReload(val uint16) {
	t.reload = val
	if t.reload == 0xFFFF || t.reload == 0xFFFE {
		// t.log().WithField("val", fmt.Sprintf("%04x", val)).Info("[timers] write reload")
		// Emu.DebugBreak("reload 0xFFFF")
	}
}

func (t *HwTimer) WriteControl(val uint16) {
	wasrunning := t.running()
	// t.log().WithFields(logrus.Fields{
	// 	"val":    fmt.Sprintf("%04x", val),
	// 	"target": Emu.Sync.Cycles(),
	// }).Infof("[timers] write control, syncing")
	t.Run(Emu.Sync.Cycles())
	// t.log().WithFields(logrus.Fields{
	// 	"cycles": t.cycles,
	// }).Info("[timers] end syncing")
	t.control = val
	if !wasrunning && t.running() {
		// 0->1 transition: reload the counter value
		t.counter = t.reload
	}
	t.reschedule()
}

func (t *HwTimer) ReadCounter() uint16 {
	// t.log().WithFields(logrus.Fields{
	// 	"cycles": Emu.Sync.Cycles(),
	// }).Infof("[timers] read counter, syncing")
	t.Run(Emu.Sync.Cycles())
	// t.log().WithFields(logrus.Fields{
	// 	"cycles": Emu.Sync.Cycles(),
	// 	"val":    fmt.Sprintf("%04x", t.counter),
	// }).Infof("[timers] read counter")
	return t.counter
}

func (t *HwTimer) ReadControl() uint16 {
	return t.control
}

// Handle an overflow event
func (t *HwTimer) overflow() {
	t.counter = t.reload
	if t.next != nil && t.next.countup() {
		t.next.up()
	}
	if t.irq() {
		if t.irqt {
			t.log().Warnf("double timer reload=%04x scaler=%d", t.reload, t.scaler())
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
	if !t.running() {
		t.cycles = target
		return
	}
	// countup timers are emulated in the context of the base timer
	if t.countup() {
		t.cycles = target
		return
	}

	if target < t.cycles {
		t.log().WithFields(logrus.Fields{
			"target": target,
			"cycles": t.cycles,
		}).Info("[timers] negative timeline")
		panic("negative timeline for timers")
	}

	// if t.name == "t7-1" {
	// 	t.log().WithFields(logrus.Fields{
	// 		"target": target,
	// 		"cycles": t.cycles,
	// 	}).Info("[timers] run")
	// }

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

func (t *HwTimers) SetName(prefix string) {
	for i := range t.Timers {
		t.Timers[i].name = fmt.Sprintf("%s-%d", prefix, i)
	}
}

func (t *HwTimers) Reset() {
	for i := range t.Timers {
		t.Timers[i] = HwTimer{}
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
