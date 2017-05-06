package raster3d

import (
	"fmt"
	"ndsemu/emu/fixed"
)

// Linear interpolator for a triangle.
type lerp struct {
	cur   int32
	delta [2]int32
	start int32
}

func newLerp(start fixed.F22, d0 fixed.F22, d1 fixed.F22) lerp {
	return lerp{start: start.V, delta: [2]int32{d0.V, d1.V}}
}

func newLerp12(start fixed.F12, d0 fixed.F12, d1 fixed.F12) lerp {
	return lerp{start: start.V, delta: [2]int32{d0.V, d1.V}}
}

func newLerpFromInt(start int32, d0 int32, d1 int32) lerp {
	return lerp{start: start, delta: [2]int32{d0, d1}}
}

func (l *lerp) Reset() {
	l.cur = l.start
}

func (l *lerp) Cur() fixed.F22 {
	return fixed.F22{V: l.cur}
}

func (l *lerp) Cur12() fixed.F12 {
	return fixed.F12{V: l.cur}
}

func (l *lerp) CurAsInt() int32 {
	return l.cur
}

func (l *lerp) Next(didx int) {
	l.cur = l.cur + l.delta[didx]
}

func (l lerp) String() string {
	return fmt.Sprintf("lerp(%v (%v,%v) [%v])",
		fixed.F22{V: l.cur}, fixed.F22{V: l.delta[0]}, fixed.F22{V: l.delta[1]}, fixed.F22{V: l.start})
}
