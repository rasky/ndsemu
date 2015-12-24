package arm

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

var RegNames = [16]string{
	"r0", "r1", "r2", "r3", "r4", "r5", "r6", "r7",
	"r8", "r9", "r10", "r11", "r12", "sp", "lr", "pc",
}

func (cpu *Cpu) GetRegNames() []string {
	return RegNames[:]
}

func (cpu *Cpu) GetRegs() []uint32 {
	var val [16]uint32
	for i := 0; i < 16; i++ {
		val[i] = uint32(cpu.Regs[i])
	}
	return val[:]
}

func (cpu *Cpu) SetReg(n int, val uint32) {
	cpu.Regs[n] = reg(val)
}

func (cpu *Cpu) GetSpecialRegNames() []string {
	return []string{"Flags", "Mode", "Insn", "Spsr", "Clock", "Lines"}
}

func (cpu *Cpu) GetSpecialRegs() []string {
	flags := ""
	if cpu.Cpsr.N() {
		flags += "N"
	} else {
		flags += "-"
	}
	if cpu.Cpsr.Z() {
		flags += "Z"
	} else {
		flags += "-"
	}
	if cpu.Cpsr.C() {
		flags += "C"
	} else {
		flags += "-"
	}
	if cpu.Cpsr.V() {
		flags += "V"
	} else {
		flags += "-"
	}
	if cpu.Cpsr.Q() {
		flags += "Q"
	} else {
		flags += "-"
	}
	if cpu.Cpsr.I() {
		flags += "I"
	} else {
		flags += "-"
	}
	if cpu.Cpsr.F() {
		flags += "F"
	} else {
		flags += "-"
	}

	insn := "arm"
	if cpu.Cpsr.T() {
		insn = "thumb"
	}

	spsr := "--"
	mode := cpu.Cpsr.GetMode()
	if mode != CpuModeSystem && mode != CpuModeUser {
		spsr = fmt.Sprint(*cpu.RegSpsr())
	}

	lines := ""
	if cpu.lines&LineIrq != 0 {
		lines += "I"
	} else {
		lines += "-"
	}
	if cpu.lines&LineFiq != 0 {
		lines += "F"
	} else {
		lines += "-"
	}
	if cpu.lines&LineHalt != 0 {
		lines += "H"
	} else {
		lines += "-"
	}

	return []string{
		flags,
		fmt.Sprint(cpu.Cpsr.GetMode()),
		insn,
		spsr,
		strconv.FormatInt(cpu.Clock, 10),
		lines,
	}
}

func (cpu *Cpu) GetPc() uint32 {
	return uint32(cpu.Regs[15])
}

var condName = [16]string{
	"eq", "ne", "hs", "lo", "mi", "pl", "vs", "vc",
	"hi", "ls", "ge", "lt", "gt", "le", "", ".ERROR",
}

func (cpu *Cpu) disasmAddCond(opname string, op uint32) string {
	suff := condName[(op>>28)&0xF]
	return opname + suff
}

var op2Shift = [4]string{"lsl", "lsr", "asr", "ror"}

func (cpu *Cpu) disasmOp2(op uint32) string {
	rmx := RegNames[op&0xF]
	shift := op2Shift[(op>>5)&3]
	if op&0x10 != 0 {
		// shift by reg
		return rmx + " " + shift + " " + RegNames[(op>>8)&0xF]
	} else {
		// shift by imm
		shiftimm := int64((op >> 7) & 0x1F)
		if shiftimm == 0 {
			switch shift {
			case "lsl":
				return rmx
			case "lsr":
				return rmx + " lsr #32"
			case "asr":
				return rmx + " asr #32"
			case "ror":
				return rmx + " rrx #1"
			}
		}
		return rmx + " " + shift + " #" + strconv.FormatInt(shiftimm, 10)
	}
}

func (cpu *Cpu) disasmSpsrName() string {
	switch cpu.Cpsr.GetMode() {
	case CpuModeUser, CpuModeSystem:
		return "spsr_ERR"
	case CpuModeFiq:
		return "spsr_fiq"
	case CpuModeSupervisor:
		return "spsr_svc"
	case CpuModeAbort:
		return "spsr_abt"
	case CpuModeIrq:
		return "spsr_irq"
	case CpuModeUndefined:
		return "spsr_und"
	default:
		panic("unreachable")
	}
}

func (cpu *Cpu) Disasm(pc uint32) (string, []byte) {
	thumb := cpu.Cpsr.T()
	if !thumb {
		var buf [4]byte
		op := cpu.opFetch32(pc)
		n := disasmArmTable[((op>>16)&0xFF0)|((op>>4)&0xF)](cpu, op, pc)
		binary.LittleEndian.PutUint32(buf[:], op)
		return n, buf[:]
	} else {
		var buf [2]byte
		op := cpu.opFetch16(pc)
		n := disasmThumbTable[(op>>8)&0xFF](cpu, op, pc)
		binary.LittleEndian.PutUint16(buf[:], op)
		return n, buf[:]
	}
}

func (cpu *Cpu) DumpStatus() {

	fmt.Printf("--------- Status at %v ----------\n", cpu.GetPC())
	text, _ := cpu.Disasm(uint32(cpu.GetPC()))
	fmt.Println(text)
	fmt.Printf("R0:%v  R1:%v  R2:%v  R3:%v  R4:%v  R5:%v  R6:%v  R7:%v\n",
		cpu.Regs[0], cpu.Regs[1], cpu.Regs[2], cpu.Regs[3], cpu.Regs[4], cpu.Regs[5], cpu.Regs[6], cpu.Regs[7])
	fmt.Printf("R8:%v  R9:%v R10:%v R11:%v R12:%v  SP:%v  LR:%v  PC:%v\n",
		cpu.Regs[8], cpu.Regs[9], cpu.Regs[10], cpu.Regs[11], cpu.Regs[12], cpu.Regs[13], cpu.Regs[14], cpu.Regs[15])

	special := cpu.GetSpecialRegs()
	fmt.Printf("Flags: %s | Mode: %s | Insn: %s | Spsr:%v | Clock:%v\n",
		special[0], special[1], special[2], special[3], special[4])
}

/*************************************************************************
 * Manual debugging code - to be removed after proper debugging support
 *************************************************************************/

var EXPAND = 0

func (cpu *Cpu) Trace() {

	// if cpu.DebugTrace > 0 {
	// 	cpu.DumpStatus()
	// 	cpu.DebugTrace--
	// }

	// if cpu.GetPC() >= 0x166A && cpu.GetPC() <= 0x1696 {
	// 	cpu.DumpStatus()
	// }

	if cpu.GetPC() == 0x2F1C || cpu.GetPC() == 0x2F24 {
		fmt.Printf("IntrWait: RAMIF=%08x/%08x WAIT=%v\n",
			cpu.opRead32(0x380FFF8), cpu.opRead32(0x3FFFFF8),
			cpu.Regs[1])
	}

	if cpu.GetPC() == 0x2038 {
		fmt.Println("EXPAND BEGIN")
		EXPAND = 1

	}

	if cpu.GetPC() == 0x20B6 {
		fmt.Println("EXPAND FINISHED")
		EXPAND = 0
	}

	if EXPAND == 0 && cpu.GetPC() == 0x20CA && cpu.Regs[4] == 17 {
		fmt.Printf("DEC IN: %v %v\n", cpu.Regs[0], cpu.Regs[6])
	}
	if EXPAND == 0 && cpu.GetPC() == 0x20EC {
		fmt.Printf("DEC OUT: %v %v\n", cpu.Regs[1], cpu.Regs[0])
	}

	if EXPAND == 0 && cpu.GetPC() == 0x2008 && cpu.Regs[4] == 0 {
		fmt.Printf("ENC IN: %v %v\n", cpu.Regs[0], cpu.Regs[6])
	}
	if EXPAND == 0 && cpu.GetPC() == 0x202A {
		fmt.Printf("ENC OUT: %v %v\n", cpu.Regs[1], cpu.Regs[0])
	}

	// if cpu.GetPC() >= 0xFFFF0940 && cpu.GetPC() <= 0xFFFF0960 {
	// 	cpu.DumpStatus()
	// }
}
