package raster3d

import (
	"fmt"
	"ndsemu/emu"
)

// Linear interpolator for a triangle.
type lerp struct {
	cur   emu.Fixed12
	delta [2]emu.Fixed12
	start emu.Fixed12
}

func newLerp(start emu.Fixed12, d0 emu.Fixed12, d1 emu.Fixed12) lerp {
	return lerp{start: start, delta: [2]emu.Fixed12{d0, d1}}
}

func (l *lerp) Reset() {
	l.cur = l.start
}

func (l *lerp) Cur() emu.Fixed12 {
	return l.cur
}

func (l *lerp) Next(didx int) {
	l.cur = l.cur.AddFixed(l.delta[didx])
}

func (l lerp) String() string {
	return fmt.Sprintf("lerp(%v (%v,%v) [%v])", l.cur, l.delta[0], l.delta[1], l.start)
}
