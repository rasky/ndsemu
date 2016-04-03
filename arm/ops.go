package arm

import (
	"unsafe"

	log "gopkg.in/Sirupsen/logrus.v0"
)

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

var flagCondTable = [8]reg{
	1 << 30, 1 << 30, // Z
	1 << 29, 1 << 29, // C
	1 << 31, 1 << 31, // N
	1 << 28, 1 << 28, // V
}

func (cpu *Cpu) opArmCond(cond uint) bool {
	// Cond 0..7 is a simple flag test (Z, C, N, V). Odd values
	// of cond means a reverse-test (that is, test if the flag
	// is not set).
	if cond < 8 {
		cpsr := cpu.Cpsr.r
		// Avoid a non-predictable branch using a bittrick to
		// reverse cpsr when cont bit 0 is set.
		cpsr ^= reg(int32(cond<<31) >> 31)
		return cpsr&flagCondTable[cond] != 0
	}
	// Cond 8..13 are more complex checks
	switch cond {
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
	}
	panic("unreachable")
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
	cpu.Clock += 2
	cpu.tightExit = true
	cpu.prevpc = cpu.pc
	cpu.pc = newpc
	if cpu.Cpsr.T() {
		if cpu.pc&1 != 0 {
			cpu.breakpoint("disaligned PC in thumb (%v->%v)", cpu.prevpc, cpu.pc)
		}
	} else {
		if cpu.pc&3 != 0 {
			cpu.breakpoint("disaligned PC in arm (%v->%v)", cpu.prevpc, cpu.pc)
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
		lines := cpu.lines
		if lines&LineHalt != 0 {
			cpu.Clock = cpu.targetCycles
			return
		}
		// Check for interrupts outside of the tight loop. This theoretically
		// should be done after each opcode, but we optimize it. We're sure that
		// we don't delay this too much because:
		//    * Every line activation immediately triggers an exit from the
		//      tight loop.
		//    * Every changes in Cpsr that affect F or I also triggers an exit
		//      from the tight loop (as interrupts might be asserted on the
		//      line but on hold because Cpsr flags)
		if lines&LineFiq != 0 && !cpu.Cpsr.F() {
			cpu.Exception(ExceptionFiq)
			continue
		}
		if lines&LineIrq != 0 && !cpu.Cpsr.I() {
			cpu.Exception(ExceptionIrq)
			continue
		}

		mem := cpu.opFetchPointer(uint32(cpu.pc))
		if mem == nil {
			cpu.breakpoint("ARMv%d jump to non-linear memory at %v", cpu.arch, cpu.pc)
		}
		if mem[0] == 0 && mem[1] == 0 && mem[2] == 0 && mem[3] == 0 {
			cpu.breakpoint("ARMv%d jump to 0 area at %v from %v", cpu.arch, cpu.pc, cpu.prevpc)
		}

		// Welcome to the tight loop. This is the innest execution loop that is
		// very very hot performance-wise. This loops is kept mostly for linear
		// execution of opcodes. The fetch pointer is linearly incremented. We
		// exit from the tight loop when:
		//
		//  * We reached the target cycles requested by the caller.
		//  * The linear bank of memory that we're fetching from is finished.
		//    The memory bus implementation is expected to return a chunk as
		//    large as possible, but we could still reach the end of a bank
		//    (even if the memory is *seen* as linear by the ARM).
		//  * The processor jumped. For simplicity, we conside all branches
		//    the same, and we just bail out of the tight loop (even if the
		//    branch target lays within the same memory bank).
		//  * An interrupt could be triggered. This can happen whenever the I/F
		//    flags are modified in CPSR, or when a physical line is asserted.
		//
		//  The last two conditions are triggered externally (see SetLine(),
		//  cpsr.SetWithMask(), and branch()), and we get notified through
		//  the tightExit variable.
		//
		//  NOTE: the code has been hand-optimized to overcome Go compiler
		//  limitations. We use unsafe.Pointer even though we actually check
		//  for memory bounds (though in a more optimized way).
		//
		memptr := unsafe.Pointer(&mem[0])
		memlen := len(mem)
		cpu.tightExit = false

		if !cpu.Cpsr.T() {
			for memlen > 3 {
				cpu.Regs[15] = cpu.pc + 8 // simulate pipeline with prefetch
				cpu.pc += 4

				if trace != nil {
					trace(uint32(cpu.pc - 4))
				}

				op := *(*uint32)(memptr)
				memlen -= 4
				memptr = unsafe.Pointer(uintptr(memptr) + 4)
				cpu.Clock++

				// Check the condition flags on each instruction (bits 28-31).
				// * 0xE means always, and is by far the most common occurrence.
				// * 0xF is used for special instructions that do not support
				// condition flags (so, they are always executed)
				// * Anything else goes through a call opArmCond()
				if op >= 0xE0000000 || cpu.opArmCond(uint(op>>28)) {
					opArmTable[(((op>>16)&0xFF0)|((op>>4)&0xF))&0xFFF](cpu, op)
				}
				if cpu.Clock >= cpu.targetCycles || cpu.tightExit {
					break
				}
			}
		} else {
			for memlen > 1 {
				cpu.Regs[15] = cpu.pc + 4 // simulate pipeline with prefetch
				cpu.pc += 2

				if trace != nil {
					trace(uint32(cpu.pc - 2))
				}

				op := *(*uint16)(memptr)
				memlen -= 2
				memptr = unsafe.Pointer(uintptr(memptr) + 2)

				cpu.Clock++
				opThumbTable[(op>>8)&0xFF](cpu, op)
				if cpu.Clock >= cpu.targetCycles || cpu.tightExit {
					break
				}
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
