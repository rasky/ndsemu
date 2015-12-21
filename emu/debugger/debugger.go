package debugger

import (
	"ndsemu/emu"

	ui "github.com/gizak/termui"
)

type Cpu interface {
	GetRegNames() []string
	GetRegs() []uint32
	SetReg(idx int, val uint32)

	GetSpecialRegNames() []string
	GetSpecialRegs() []string

	GetPc() uint32
	Disasm(pc uint32) (string, []byte)

	Step()
}

type Debugger struct {
	sync   *emu.Sync
	cpus   []Cpu
	curcpu Cpu

	userBkps []uint32
	ourBkps  []uint32

	running   bool
	runch     chan bool
	focusline int
	pcline    int
	lines     []string
	linepc    []uint32
	uiCode    *ui.List
	uiRegs    *ui.List
}

func New(cpus []Cpu, sync *emu.Sync) *Debugger {
	return &Debugger{
		sync:      sync,
		cpus:      cpus,
		focusline: -1,
		runch:     make(chan bool),
	}
}

func (dbg *Debugger) runMonitored() {
	defer func() { dbg.running = false }()

	i := 0
	for {
		select {
		case <-dbg.runch:
			return
		default:
			dbg.curcpu.Step()
			pc := dbg.curcpu.GetPc()
			for _, b := range dbg.userBkps {
				if b == pc {
					dbg.refreshUi()
					return
				}
			}
			for idx, b := range dbg.ourBkps {
				if b == pc {
					dbg.ourBkps = append(dbg.ourBkps[:idx], dbg.ourBkps[idx+1:]...)
					return
				}
			}
			i++
			if i&300 == 0 {
				dbg.sync.Sync()
			}
		}
	}
}

func (dbg *Debugger) stopMonitored() {
	if dbg.running {
		dbg.runch <- false
	}
}

func (dbg *Debugger) Run() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	dbg.curcpu = dbg.cpus[0]

	dbg.initUi()
	defer ui.Close()

	run := func() {
		par := ui.NewPar("Running....\nPress SPACE to break")
		par.Width = 50
		par.Height = 4
		par.Align()
		par.SetX((ui.TermWidth() - par.Width) / 2)
		par.SetY((ui.TermHeight() - par.Height) / 2)
		ui.Render(par)

		dbg.running = true
		go func() {
			dbg.runMonitored()
			dbg.refreshUi()
		}()
	}

	runto := func(stop uint32) {
		dbg.ourBkps = append(dbg.ourBkps, stop)
		dbg.focusline = -1
		run()
	}

	stop := func() {
		dbg.stopMonitored()
	}

	switchcpu := func(idx int) {
		dbg.curcpu = dbg.cpus[idx]
		dbg.focusline = -1
		dbg.linepc = nil
		dbg.uiCode.Items = nil
		dbg.uiRegs.Items = nil
		dbg.refreshUi()
	}

	ui.Handle("/sys/kbd/<space>", func(ui.Event) {
		if !dbg.running {
			run()
		} else {
			stop()
		}
	})

	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		if !dbg.running && dbg.focusline >= 0 {
			pc := dbg.linepc[dbg.focusline]
			runto(pc)
		}
	})

	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		if !dbg.running {
			dbg.focusline--
			dbg.refreshUi()
		}
	})
	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		if !dbg.running {
			dbg.focusline++
			dbg.refreshUi()
		}
	})

	ui.Handle("/sys/kbd/r", func(ui.Event) {
		if !dbg.running {
			// force refresh of disasm screen
			dbg.linepc = nil
			dbg.lines = nil
			dbg.refreshUi()
		}
	})

	ui.Handle("/sys/kbd/s", func(ui.Event) {
		if !dbg.running {
			dbg.curcpu.Step()
			dbg.refreshUi()
		}
	})

	ui.Handle("/sys/kbd/n", func(ui.Event) {
		if !dbg.running {
			pc := dbg.curcpu.GetPc()
			if pc != dbg.linepc[dbg.pcline] {
				panic("inconsistent pc")
			}
			dbg.curcpu.Step()
			if pc != dbg.linepc[dbg.pcline+1] {
				runto(dbg.linepc[dbg.pcline+1])
			} else {
				dbg.refreshUi()
			}
		}
	})

	ui.Handle("/sys/kbd/1", func(ui.Event) {
		if !dbg.running {
			switchcpu(0)
		}
	})

	ui.Handle("/sys/kbd/2", func(ui.Event) {
		if !dbg.running {
			switchcpu(1)
		}
	})

	ui.Handle("/sys/kbd/q", func(ui.Event) {
		dbg.stopMonitored()
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		dbg.stopMonitored()
		ui.StopLoop()
	})

	dbg.refreshUi()
	ui.Loop()
}

func (dbg *Debugger) AddBreakpoint(pc uint32) {
	dbg.userBkps = append(dbg.userBkps, pc)
}
