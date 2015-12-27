package debugger

import (
	"fmt"
	"ndsemu/emu"

	ui "github.com/gizak/termui"
)

// This is the interface that can be passed to a CPU core to make it interoperate
// with the debugger.
type CpuDebugger interface {
	// Trace() must be called before each opcode is executed. This is the main
	// entry point for debugging activity, as the debug can stop the CPU
	// execution by making this function blocking until user interaction
	// finishes.
	Trace(pc uint32)

	// WatchRead/WatchWrite must be called before each memory access. They can
	// be used by the debugger to implement watchpoints and thus intercept
	// memory accesses
	// WatchRead(addr uint32)
	// WatchWrite(addr uint32, val uint32)

	// Break() can be called by the CPU core to force breaking into the debugger.
	// It can be used in situations such as invalid opcodes
	Break(msg string)
}

type Cpu interface {
	// SetDebugger is used to install a debugger into the CPU core. Notice that,
	// to maximize performance.
	SetDebugger(dbg CpuDebugger)

	GetRegNames() []string
	GetRegs() []uint32
	SetReg(idx int, val uint32)

	GetSpecialRegNames() []string
	GetSpecialRegs() []string

	GetPc() uint32
	Disasm(pc uint32) (string, []byte)
}

type Debugger struct {
	sync   *emu.Sync
	cpus   []Cpu
	curcpu int

	userBkps []uint32
	ourBkps  []uint32

	running   []bool
	focusline int
	pcline    int
	lines     []string
	linepc    []uint32
	uiCode    *ui.List
	uiRegs    *ui.List
	uiLog     *ui.List
	uiCalls   *ui.List
	pcchain   [][]uint32

	breakch chan string

	log *logReader
}

type dbgForCpu struct {
	*Debugger
	cpuidx int
}

func New(cpus []Cpu, sync *emu.Sync) *Debugger {
	dbg := &Debugger{
		sync:      sync,
		cpus:      cpus,
		focusline: -1,
		// runch:     make(chan bool, 1),
		running: make([]bool, len(cpus)),
		pcchain: make([][]uint32, len(cpus)),
		breakch: make(chan string),
	}

	for idx, cpu := range cpus {
		dbg.pcchain[idx] = make([]uint32, 1)
		cpu.SetDebugger(dbgForCpu{dbg, idx})
	}

	return dbg
}

func (dbg dbgForCpu) Trace(pc uint32) {
	idx := dbg.cpuidx
	lpc := dbg.pcchain[idx][len(dbg.pcchain[idx])-1]
	if !(pc >= lpc && pc <= lpc+4) {
		dbg.pcchain[idx] = append(dbg.pcchain[idx], pc)
		if len(dbg.pcchain[idx]) > 16 {
			dbg.pcchain[idx] = dbg.pcchain[idx][len(dbg.pcchain[idx])-16:]
		}
	}

	if msg, found := dbg.checkBreapoint(idx, pc); found {
		dbg.curcpu = dbg.cpuidx
		dbg.Break(msg)
	}
}

func (dbg *Debugger) Break(msg string) {
	dbg.breakch <- msg
	<-dbg.breakch
}

func (dbg *Debugger) checkBreapoint(cpuidx int, pc uint32) (string, bool) {
	if !dbg.running[cpuidx] {
		return "", true
	}

	for _, b := range dbg.userBkps {
		if b == pc {
			return fmt.Sprintf("user breakpoint at %08x", pc), true
		}
	}
	for idx, b := range dbg.ourBkps {
		if b == pc {
			dbg.ourBkps = append(dbg.ourBkps[:idx], dbg.ourBkps[idx+1:]...)
			return "", true
		}
	}
	return "", false
}

func (dbg *Debugger) runMonitored() string {
	for {
		select {
		case <-dbg.log.NewLog:
			dbg.refreshLog()
			ui.Render(dbg.uiLog)
		case msg := <-dbg.breakch:
			return msg
		}
	}
}

func (dbg *Debugger) resumeEmulation(running bool, cb func()) {
	if running {
		for i := 0; i < len(dbg.cpus); i++ {
			dbg.running[i] = true
		}
	} else {
		for i := 0; i < len(dbg.cpus); i++ {
			if i == dbg.curcpu {
				dbg.running[i] = false
			} else {
				dbg.running[i] = true
			}
		}
	}
	go func() {
		dbg.breakch <- ""
		dbg.runMonitored()
		dbg.stopMonitored()
		if cb != nil {
			cb()
		}
	}()
}

func (dbg *Debugger) stopMonitored() {
	for i := 0; i < len(dbg.cpus); i++ {
		dbg.running[i] = false
	}
}

func (dbg *Debugger) Run() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}

	dbg.curcpu = 0
	dbg.log = newLogReader()

	dbg.initUi()
	defer ui.Close()

	run := func() {
		par := ui.NewPar("Running....\nPress SPACE to break")
		par.Width = 50
		par.Height = 4
		par.Align()
		par.SetX((ui.TermWidth() - par.Width) / 2)
		par.SetY((ui.TermHeight()-par.Height)/2 - 10)
		ui.Render(par)

		dbg.resumeEmulation(true, func() {
			dbg.refreshUi()
		})
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
		dbg.curcpu = idx
		dbg.focusline = -1
		dbg.linepc = nil
		dbg.uiCode.Items = nil
		dbg.uiRegs.Items = nil
		dbg.refreshUi()
	}

	ui.Handle("/sys/kbd/<space>", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
			run()
		} else {
			stop()
		}
	})

	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		if !dbg.running[dbg.curcpu] && dbg.focusline >= 0 {
			pc := dbg.linepc[dbg.focusline]
			runto(pc)
		}
	})

	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
			dbg.focusline--
			dbg.refreshUi()
		}
	})
	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
			dbg.focusline++
			dbg.refreshUi()
		}
	})

	ui.Handle("/sys/kbd/r", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
			// force refresh of disasm screen
			dbg.linepc = nil
			dbg.lines = nil
			dbg.refreshUi()
		}
	})

	ui.Handle("/sys/kbd/s", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
			dbg.resumeEmulation(false, func() {
				dbg.refreshUi()
			})
		}
	})

	ui.Handle("/sys/kbd/n", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
			pc := dbg.cpus[dbg.curcpu].GetPc()
			if pc != dbg.linepc[dbg.pcline] {
				panic("inconsistent pc")
			}
			nextpc := dbg.linepc[dbg.pcline+1]
			dbg.resumeEmulation(false, func() {
				pc := dbg.cpus[dbg.curcpu].GetPc()
				if pc != nextpc {
					runto(nextpc)
				} else {
					dbg.refreshUi()
				}
			})
		}
	})

	ui.Handle("/sys/kbd/1", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
			switchcpu(0)
		}
	})

	ui.Handle("/sys/kbd/2", func(ui.Event) {
		if !dbg.running[dbg.curcpu] {
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

	dbg.runMonitored()
	dbg.refreshUi()
	ui.Loop()
}

func (dbg *Debugger) AddBreakpoint(pc uint32) {
	dbg.userBkps = append(dbg.userBkps, pc)
}
