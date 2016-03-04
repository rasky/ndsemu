package raster3d

import (
	"fmt"
	"ndsemu/emu"
)

// Linear interpolator for a triangle.
type lerp struct {
	cur   int32
	delta [2]int32
	start int32
}

func newLerp(start emu.Fixed12, d0 emu.Fixed12, d1 emu.Fixed12) lerp {
	return lerp{start: start.V, delta: [2]int32{d0.V, d1.V}}
}

func newLerpFromInt(start int32, d0 int32, d1 int32) lerp {
	return lerp{start: start, delta: [2]int32{d0, d1}}
}

func (l *lerp) Reset() {
	l.cur = l.start
}

func (l *lerp) Cur() emu.Fixed12 {
	return emu.Fixed12{V: l.cur}
}

func (l *lerp) CurAsInt() int32 {
	return l.cur
}

func (l *lerp) Next(didx int) {
	l.cur = l.cur + l.delta[didx]
}

func (l lerp) String() string {
	return fmt.Sprintf("lerp(%v (%v,%v) [%v])",
		emu.Fixed12{V: l.cur}, emu.Fixed12{V: l.delta[0]}, emu.Fixed12{V: l.delta[1]}, emu.Fixed12{V: l.start})
}
