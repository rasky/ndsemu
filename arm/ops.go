package arm

import log "gopkg.in/Sirupsen/logrus.v0"

//go:generate go run genarm/genarm.go -filename ops_arm_table.go
//go:generate go run genthumb/genthumb.go -filename ops_thumb_table.go

var popcount = [256]uint8{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

func popcount8(val uint8) uint {
	return uint(popcount[val])
}
func popcount16(val uint16) uint {
	return popcount8(uint8(val&0xFF)) + popcount8(uint8(val>>8))
}

// func init() {
// 	for i := 0; i < 256; i++ {
// 		cnt := 0
// 		for j := 0; j < 8; j++ {
// 			if i&(1<<uint(j)) != 0 {
// 				cnt++
// 			}
// 		}
// 		popcount[i] = cnt
// 	}
// 	fmt.Printf("%#v\n", popcount)
// }

func (cpu *Cpu) InvalidOpArm(op uint32, msg string) {
	cpu.breakpoint("invalid ARM opcode at %v (%04X): %s", cpu.GetPC(), op, msg)
}

func (cpu *Cpu) InvalidOpThumb(op uint16, msg string) {
	cpu.breakpoint("invalid thumb opcode at %v (%04X): %s", cpu.pc-2, op, msg)
}

func (cpu *Cpu) opArmCond(op uint32) bool {
	switch op >> 28 {
	case 0:
		return cpu.Cpsr.Z()
	case 1:
		return !cpu.Cpsr.Z()
	case 2:
		return cpu.Cpsr.C()
	case 3:
		return !cpu.Cpsr.C()
	case 4:
		return cpu.Cpsr.N()
	case 5:
		return !cpu.Cpsr.N()
	case 6:
		return cpu.Cpsr.V()
	case 7:
		return !cpu.Cpsr.V()
	case 8:
		return cpu.Cpsr.C() && !cpu.Cpsr.Z()
	case 9:
		return !cpu.Cpsr.C() || cpu.Cpsr.Z()
	case 10:
		return cpu.Cpsr.N() == cpu.Cpsr.V()
	case 11:
		return cpu.Cpsr.N() != cpu.Cpsr.V()
	case 12:
		return !cpu.Cpsr.Z() && cpu.Cpsr.N() == cpu.Cpsr.V()
	case 13:
		return cpu.Cpsr.Z() || cpu.Cpsr.N() != cpu.Cpsr.V()
	case 14:
		return true
	case 15:
		return false
	}
	return false
}

func (cpu *Cpu) opCopRead(copnum uint32, op uint32, cn, cm, cp uint32) uint32 {
	cop := cpu.cops[copnum]
	if cop == nil {
		log.WithFields(log.Fields{
			"pc":  cpu.pc,
			"cop": copnum,
		}).Error("invalid coprocessor access")
		return 0xFFFFFFFF
	}

	return cop.Read(op, cn, cm, cp)
}

func (cpu *Cpu) opCopWrite(copnum uint32, op uint32, cn, cm, cp uint32, value uint32) {
	cop := cpu.cops[copnum]
	if cop == nil {
		log.WithFields(log.Fields{
			"pc":  cpu.pc,
			"cop": copnum,
		}).Error("invalid coprocessor access")
		return
	}

	cop.Write(op, cn, cm, cp, value)
}

func (cpu *Cpu) opCopExec(copnum uint32, op uint32, cn, cm, cp, cd uint32) {
	cop := cpu.cops[copnum]
	if cop == nil {
		log.WithFields(log.Fields{
			"pc":  cpu.pc,
			"cop": copnum,
		}).Error("invalid coprocessor access")
		return
	}

	cop.Exec(op, cn, cm, cp, cd)
}

type BranchType int

const (
	BranchJump BranchType = iota
	BranchCall
	BranchReturn
	BranchInterrupt
)

func (cpu *Cpu) branch(newpc reg, reason BranchType) {
	prevpc := cpu.pc
	cpu.Clock += 2
	cpu.tightExit = true
	cpu.pc = newpc
	if cpu.Cpsr.T() {
		if cpu.pc&1 != 0 {
			cpu.breakpoint("disaligned PC in thumb (%v->%v)", prevpc, cpu.pc)
		}
	} else {
		if cpu.pc&3 != 0 {
			cpu.breakpoint("disaligned PC in arm (%v->%v)", prevpc, cpu.pc)
		}
	}
}

func (cpu *Cpu) Retarget(until int64) {
	if cpu.targetCycles > until {
		cpu.targetCycles = until
	}
}

func (cpu *Cpu) Run(until int64) {
	cpu.targetCycles = until

	var trace func(uint32)
	if cpu.dbg != nil {
		trace = cpu.dbg.Trace
	}

	for cpu.Clock < cpu.targetCycles {
		if cpu.lines&LineHalt != 0 {
			cpu.Clock = cpu.targetCycles
			return
		}

		mem := cpu.opFetchPointer(uint32(cpu.pc))
		if mem == nil {
			cpu.breakpoint("[CPU] ARMv%d jump to non-linear memory at %v", cpu.arch, cpu.pc)
		}
		if mem[0] == 0 && mem[1] == 0 && mem[2] == 0 && mem[3] == 0 {
			cpu.breakpoint("[CPU] ARMv%d jump to 0 area at %v from %v", cpu.arch, cpu.pc, cpu.prevpc)
		}

		if !cpu.Cpsr.T() {
			cpu.tightExit = false
			for cpu.Clock < cpu.targetCycles && !cpu.tightExit && len(mem) > 0 {
				cpu.Regs[15] = cpu.pc + 8 // simulate pipeline with prefetch
				cpu.pc += 4

				lines := cpu.lines
				if lines&LineFiq != 0 && !cpu.Cpsr.F() {
					cpu.Exception(ExceptionFiq)
					cpu.Clock += 3 // FIXME
					break
				}
				if lines&LineIrq != 0 && !cpu.Cpsr.I() {
					cpu.Exception(ExceptionIrq)
					cpu.Clock += 3 // FIXME
					break
				}
				if lines&LineHalt != 0 {
					cpu.pc -= 4
					cpu.Clock = cpu.targetCycles
					return
				}
				if trace != nil {
					trace(uint32(cpu.pc - 4))
				}

				op := uint32(mem[0]) | uint32(mem[1])<<8 | uint32(mem[2])<<16 | uint32(mem[3])<<24
				mem = mem[4:]
				cpu.Clock++
				opArmTable[((op>>16)&0xFF0)|((op>>4)&0xF)](cpu, op)
			}
		} else {
			cpu.tightExit = false
			for cpu.Clock < cpu.targetCycles && !cpu.tightExit && len(mem) > 0 {
				cpu.Regs[15] = cpu.pc + 4 // simulate pipeline with prefetch
				cpu.pc += 2

				lines := cpu.lines
				if lines&LineFiq != 0 && !cpu.Cpsr.F() {
					cpu.Exception(ExceptionFiq)
					cpu.Clock += 3 // FIXME
					break
				}
				if lines&LineIrq != 0 && !cpu.Cpsr.I() {
					cpu.Exception(ExceptionIrq)
					cpu.Clock += 3 // FIXME
					break
				}
				if lines&LineHalt != 0 {
					cpu.pc -= 2
					cpu.Clock = cpu.targetCycles
					return
				}

				if trace != nil {
					trace(uint32(cpu.pc - 2))
				}
				op := uint16(mem[0]) | uint16(mem[1])<<8
				mem = mem[2:]
				cpu.Clock++
				opThumbTable[(op>>8)&0xFF](cpu, op)
			}
		}
	}

	cpu.Regs[15] = cpu.pc
}

func (cpu *Cpu) GetPC() reg {
	thumb := cpu.Cpsr.T()
	if !thumb {
		return cpu.pc - 4
	} else {
		return cpu.pc - 2
	}
}
