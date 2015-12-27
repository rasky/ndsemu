package debugger

import (
	"encoding/hex"
	"fmt"
	"strings"

	ui "github.com/gizak/termui"
)

func (dbg *Debugger) initUi() {
	dbg.uiCode = ui.NewList()
	dbg.uiCode.BorderLabel = "Code"
	dbg.uiCode.BorderFg = ui.ColorGreen

	dbg.uiRegs = ui.NewList()
	dbg.uiRegs.BorderLabel = "Regs"
	dbg.uiRegs.BorderFg = ui.ColorGreen

	dbg.uiCalls = ui.NewList()
	dbg.uiCalls.BorderLabel = "Calls"
	dbg.uiCalls.BorderFg = ui.ColorGreen

	dbg.uiLog = ui.NewList()
	dbg.uiLog.BorderLabel = "Logging"
	dbg.uiLog.BorderFg = ui.ColorGreen
	dbg.uiLog.Height = 20
	dbg.log.SetNumLines(20)

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 0, dbg.uiCode),
			ui.NewCol(1, 0, dbg.uiCalls),
			ui.NewCol(6, 0,
				dbg.uiRegs,
				dbg.uiLog,
			),
		),
	)

	ui.Body.Align()
}

func (dbg *Debugger) disasmBlock(pc uint32, sz int, area uint32) {
	nlines := ui.TermHeight() - 4
	dbg.lines = make([]string, nlines)
	dbg.linepc = make([]uint32, nlines)

	for i := 0; i < nlines; i++ {
		var text string
		var buf []byte
		if pc>>24 != area>>24 {
			// avoid disassembling cross-block
			text = "unknown"
			buf = make([]byte, sz)
		} else {
			text, buf = dbg.cpus[dbg.curcpu].Disasm(pc)
		}
		datahex := hex.EncodeToString(buf)
		dbg.linepc[i] = pc
		dbg.lines[i] = fmt.Sprintf("   %08x  %-16s%s", pc, datahex, text)
		pc += uint32(len(buf))
	}
}

func (dbg *Debugger) refreshCode() {
	curpc := dbg.cpus[dbg.curcpu].GetPc()

	const margin = 4
	if dbg.focusline == dbg.pcline {
		dbg.focusline = -1
	}
	dbg.pcline = -1
	for i := margin; i < len(dbg.linepc)-margin; i++ {
		if dbg.linepc[i] == curpc {
			dbg.pcline = i
			break
		}
	}

	nlines := ui.TermHeight() - 4
	if dbg.pcline < 0 {
		dbg.focusline = -1
		_, data := dbg.cpus[dbg.curcpu].Disasm(curpc)
		dbg.disasmBlock(curpc-uint32((nlines/2)*len(data)), len(data), curpc)
	}

	final := make([]string, 0, nlines)
	for i := 0; i < len(dbg.linepc); i++ {
		line := dbg.lines[i]
		if dbg.linepc[i] == curpc {
			line = fmt.Sprintf("[%s%s](bg-green)", line,
				strings.Repeat(" ", dbg.uiCode.Width-5-len(line)))
			dbg.pcline = i
			if dbg.focusline == -1 {
				dbg.focusline = i
			}
		} else if i == dbg.focusline {
			line = fmt.Sprintf("[%s%s](bg-red)", line,
				strings.Repeat(" ", dbg.uiCode.Width-5-len(line)))
		}
		final = append(final, line)
	}

	dbg.uiCode.Items = final
	dbg.uiCode.Height = len(final) + 2
}

func (dbg *Debugger) refreshRegs() {
	cpu := dbg.cpus[dbg.curcpu]
	names := cpu.GetRegNames()
	values := cpu.GetRegs()

	snames := cpu.GetSpecialRegNames()
	svalues := cpu.GetSpecialRegs()

	lines := make([]string, len(names))
	for idx := range names {
		text := fmt.Sprintf("%4s: %08x", names[idx], values[idx])
		if len(dbg.uiRegs.Items) > idx && !strings.Contains(dbg.uiRegs.Items[idx], text) {
			text = fmt.Sprintf("[%s](fg-bold)", text)
		}

		if idx < len(snames) {
			c2 := fmt.Sprintf("%8s: %s", snames[idx], svalues[idx])
			if len(dbg.uiRegs.Items) > idx && !strings.Contains(dbg.uiRegs.Items[idx], c2) {
				c2 = fmt.Sprintf("[%s](fg-bold)", c2)
			}
			text += "     " + c2
		}

		lines[idx] = text
	}
	dbg.uiRegs.Items = lines
	dbg.uiRegs.Height = len(lines) + 2
}

func (dbg *Debugger) refreshCall() {
	chain := dbg.pcchain[dbg.curcpu]
	calls := make([]string, 0, len(chain))
	for _, pc := range chain {
		calls = append(calls, fmt.Sprintf("%08x", pc))
	}

	dbg.uiCalls.Items = calls
	dbg.uiCalls.Height = dbg.uiRegs.Height
}

func (dbg *Debugger) refreshLog() {
	dbg.uiLog.Items = dbg.log.Lines()
}

func (dbg *Debugger) refreshUi() {
	dbg.refreshCode()
	dbg.refreshRegs()
	dbg.refreshLog()
	dbg.refreshCall()

	ui.Body.Align()
	ui.Render(ui.Body)
}
