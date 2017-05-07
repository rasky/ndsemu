package raster3d

import (
	"fmt"
	"ndsemu/emu/fixed"
)

// Linear interpolator for a triangle.
type lerp struct {
	cur   int64
	delta [2]int64
	start int64
}

func newLerp(start fixed.F32, d0 fixed.F32, d1 fixed.F32) lerp {
	return lerp{start: start.V, delta: [2]int64{d0.V, d1.V}}
}

func newLerp12(start fixed.F12, d0 fixed.F12, d1 fixed.F12) lerp {
	return lerp{start: int64(start.V), delta: [2]int64{int64(d0.V), int64(d1.V)}}
}

func newLerpFromInt(start int64, d0 int64, d1 int64) lerp {
	return lerp{start: start, delta: [2]int64{d0, d1}}
}

func (l *lerp) Reset() {
	l.cur = l.start
}

func (l *lerp) Cur() fixed.F32 {
	return fixed.F32{V: l.cur}
}

func (l *lerp) Cur12() fixed.F12 {
	return fixed.F12{V: int32(l.cur)}
}

func (l *lerp) CurAsInt() int64 {
	return l.cur
}

func (l *lerp) Next(didx int) {
	l.cur = l.cur + l.delta[didx]
}

func (l lerp) String() string {
	return fmt.Sprintf("lerp(%v (%v,%v) [%v])",
		fixed.F32{V: l.cur}, fixed.F32{V: l.delta[0]}, fixed.F32{V: l.delta[1]}, fixed.F32{V: l.start})
}
