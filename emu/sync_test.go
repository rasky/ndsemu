package emu

import (
	"reflect"
	"testing"
)

type testSubsystem struct {
	Freq    int64
	targets []int64
}

func (ts *testSubsystem) Frequency() Fixed8 { return NewFixed8(ts.Freq) }
func (ts *testSubsystem) Reset()            { ts.targets = nil }

func (ts *testSubsystem) Cycles() int64 {
	if len(ts.targets) > 0 {
		return ts.targets[len(ts.targets)-1]
	}
	return 0
}

func (ts *testSubsystem) Run(target int64) {
	ts.targets = append(ts.targets, target)
}

type dotpos struct {
	X, Y int
}

type testSyncs struct {
	syncs []dotpos
}

func (ts *testSyncs) SyncEvent(x, y int) {
	ts.syncs = append(ts.syncs, dotpos{x, y})
}

func TestHVSyncs(t *testing.T) {
	tsub := testSubsystem{Freq: 400}
	hsyncs := testSyncs{}
	vsyncs := testSyncs{}

	var sync *Sync
	sync, err := NewSync(SyncConfig{
		MainClock:       200,
		DotClockDivider: 2,
		HDots:           10,
		VDots:           5,
		HSyncs:          []int{5},
		VSyncs:          []int{3, 1},

		HSync: func(x, y int) {
			sx, sy := sync.DotPos()
			if sx != x || sy != y {
				t.Errorf("invalid sync dotpos at hsync: got:%v,%v, want:%v,%v",
					sx, sy, x, y)
			}
			hsyncs.SyncEvent(x, y)
		},
		VSync: func(x, y int) {
			sx, sy := sync.DotPos()
			if sx != x || sy != y {
				t.Errorf("invalid sync dotpos at vsync: got:%v,%v, want:%v,%v",
					sx, sy, x, y)
			}
			vsyncs.SyncEvent(x, y)
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	sync.AddSubsystem(&tsub)
	sync.RunOneFrame()

	expHsyncs := []dotpos{{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}}
	if !reflect.DeepEqual(hsyncs.syncs, expHsyncs) {
		t.Errorf("wrong hsyncs: got:%v, want:%v", hsyncs.syncs, expHsyncs)
	}
	expVsyncs := []dotpos{{0, 1}, {0, 3}}
	if !reflect.DeepEqual(vsyncs.syncs, expVsyncs) {
		t.Errorf("wrong vsyncs: got:%v, want:%v", vsyncs.syncs, expVsyncs)
	}

	expTargets := []int64{5 * 4, 10 * 4, 15 * 4, 25 * 4, 30 * 4, 35 * 4, 45 * 4, 50 * 4}
	if !reflect.DeepEqual(tsub.targets, expTargets) {
		t.Errorf("wrong sub targets: got:%v, want:%v", tsub.targets, expTargets)
	}
}
