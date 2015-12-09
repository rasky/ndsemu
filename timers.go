package main

const cTimerClock = cBusClock

type HwTimer struct {
	counter uint16
	reload  uint16
	control uint16

	cycles int64
	next   *HwTimer
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

func (t *HwTimer) WriteReload(val uint16) {
	t.reload = val
}

func (t *HwTimer) WriteControl(val uint16) {
	wasrunning := t.running()
	t.control = val
	if !wasrunning && t.running() {
		// 0->1 transition: reload the counter value
		t.counter = t.reload
	}
}

func (t *HwTimer) ReadCounter() uint16 {
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
		panic("IRQ on timer not implemented")
	}
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

	scaler := int64(t.scaler())
	elapsed := target - t.cycles
	elapsed /= scaler
	t.cycles = target + elapsed*scaler

	for elapsed > 0 {
		span := 0x10000 - int64(t.counter)
		if span > elapsed {
			span = elapsed
			t.counter += uint16(span)
		} else {
			t.overflow()
		}
		elapsed -= span
	}
}

type HwTimers struct {
	Timers [4]HwTimer
}

func (t *HwTimers) Reset() {
	for i := range t.Timers {
		t.Timers[i] = HwTimer{}
		if i != 3 {
			t.Timers[i].next = &t.Timers[i+1]
		}
	}
}

func (t *HwTimers) Frequency() fixed8 {
	return NewFixed8(cTimerClock)
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
	}
}
