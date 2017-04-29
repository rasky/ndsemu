package arm

import (
	"fmt"
	log "ndsemu/emu/logger"
	"unsafe"

	a "github.com/rasky/gojit/amd64"
)

// ABI for JIT functions:
//    R15 = *Cpu
//    R15 = Cpsr (temporary)
var (
	jitRegCpu  = a.R15
	jitRegCpsr = a.R14d
)

var (
	cpuRegsOff    = int32(unsafe.Offsetof(Cpu{}.Regs))
	cpuCpsrOff    = int32(unsafe.Offsetof(Cpu{}.Cpsr))
	cpuClockOff   = int32(unsafe.Offsetof(Cpu{}.Clock))
	cpuTargetOff  = int32(unsafe.Offsetof(Cpu{}.targetCycles))
	cpuTightOff   = int32(unsafe.Offsetof(Cpu{}.tightExit))
	oCpsr         = a.Indirect{jitRegCpu, cpuCpsrOff, 32}
	oCycles       = a.Indirect{jitRegCpu, cpuClockOff, 64}
	oTargetCycles = a.Indirect{jitRegCpu, cpuTargetOff, 64}
	oTightExit    = a.Indirect{jitRegCpu, cpuTightOff, 1}
)

type jitArm struct {
	*a.Assembler
	Cpu *Cpu

	inCallBlock   bool
	afterCall     bool
	frameSize     int32
	callFrameSize int32
}

const (
	jitFlagCAdd = 1 << iota // Carry flag (result of ADD)
	jitFlagCSub             // Carry flag (result of SUB -> ARM reverse meaning)
	jitFlagCR10             // Carry flag already in R10 (ignore x86 flag)
	jitFlagV
	jitFlagN
	jitFlagZ
)

const (
	branchFlagExchange    = 1 << iota // Check if bit 0 is set, and switch to thumb
	branchFlagCpsrRestore             // Load SPSR into CPSR
)

func (j *jitArm) oArmReg(rn uint32) a.Operand {
	if j.afterCall { // self-check
		panic("cannot access ARM registers during call block")
	}
	return a.Indirect{jitRegCpu, cpuRegsOff + int32(rn*4), 32}
}

// Return an operand addressing a slot within the stack section holding the
// arguments to the jit-compiled function.
// Currently, this contains only a pointer to Cpu.
func (j *jitArm) ArgSlot(off int32, bits byte) a.Operand {
	// Go through the local frames, plus 8 bytes for BP, and 8 bytes for return address
	return a.Indirect{a.Rsp, 16 + j.frameSize + j.callFrameSize, bits}
}

// Return an operand addressing a slot within the stack section holding
// the frame (local variables for the JIT-compiled function)
func (j *jitArm) FrameSlot(off int32, bits byte) a.Operand {
	return a.Indirect{a.Rsp, off + j.callFrameSize, bits}
}

// Return an operand addressing a slog within the stack section holding
// the arguments to a Go function in the process of being called. This can
// only be called within the closure passed to CallBlock()
func (j *jitArm) CallSlot(off int32, bits byte) a.Operand {
	if !j.inCallBlock {
		panic("cannot call CallSlot outside of a call block")
	}
	return a.Indirect{a.Rsp, off, bits}
}

func (j *jitArm) CallBlock(framesize int32, cont func()) {
	if j.inCallBlock {
		panic("reentrant callblock not supported")
	}
	if framesize&7 != 0 {
		panic("unaligned frame size")
	}
	j.inCallBlock = true
	j.doEndBlock()
	j.Sub(a.Imm{framesize}, a.Rsp)
	j.callFrameSize = framesize
	cont()
	j.Add(a.Imm{framesize}, a.Rsp)
	j.callFrameSize = 0
	j.doBeginBlock()
	j.inCallBlock = false
	j.afterCall = false
}

func (j *jitArm) CallFuncGo(f interface{}) {
	if !j.inCallBlock {
		panic("CallFuncGo without CallBock")
	}
	j.Assembler.CallFuncGo(f)
	j.afterCall = true
}

func (j *jitArm) AddCycles(ncycles int32) {
	if j.afterCall { // self-check
		panic("cannot access cycle counter during call block")
	}
	if ncycles == 1 {
		j.Inc(oCycles)
	} else {
		j.Add(a.Imm{int32(ncycles)}, oCycles)
	}
}

// Copy x86 flags into ARM flags
func (j *jitArm) CopyFlags(flags int) {
	if flags&(jitFlagCAdd|jitFlagCSub) == (jitFlagCAdd | jitFlagCSub) {
		panic("both carry flag types selected")
	}

	mask := uint32(0)

	// Carry flag on ARM is reversed after subtractions.
	// Basically, after a SUBS, the carry flag is set if there
	// is NO overflow. This means that we need to reverse the
	// x86 carry after subtraction
	if flags&jitFlagCAdd != 0 {
		j.Setcc(a.CC_C, a.R10d)
		mask |= 1 << 29
	}
	if flags&jitFlagCSub != 0 {
		j.Setcc(a.CC_NC, a.R10d)
		mask |= 1 << 29
	}
	if flags&jitFlagCR10 != 0 {
		mask |= 1 << 29
	}
	if flags&jitFlagV != 0 {
		j.Setcc(a.CC_O, a.R11d)
		mask |= 1 << 28
	}
	if flags&jitFlagN != 0 {
		j.Setcc(a.CC_S, a.R12d)
		mask |= 1 << 31
	}
	if flags&jitFlagZ != 0 {
		j.Setcc(a.CC_Z, a.R13d)
		mask |= 1 << 30
	}

	j.And(a.Imm{int32(^mask)}, jitRegCpsr)

	if flags&(jitFlagCAdd|jitFlagCSub|jitFlagCR10) != 0 {
		j.Shl(a.Imm{29}, a.R10d)
		j.Or(a.R10d, jitRegCpsr)
	}
	if flags&jitFlagV != 0 {
		j.Shl(a.Imm{28}, a.R11d)
		j.Or(a.R11d, jitRegCpsr)
	}
	if flags&jitFlagN != 0 {
		j.Shl(a.Imm{31}, a.R12d)
		j.Or(a.R12d, jitRegCpsr)
	}
	if flags&jitFlagZ != 0 {
		j.Shl(a.Imm{30}, a.R13d)
		j.Or(a.R13d, jitRegCpsr)
	}
}

// Emit JIT code for op2 decoding in ALU opcodes.
// op2 goes into EBX
func (j *jitArm) emitAluOp2Reg(op uint32, setcarry bool) {
	shtype := (op >> 5) & 3
	byreg := op&0x10 != 0

	// load value to be shifted/rotated into ebx
	j.Movl(j.oArmReg(op&0xF), a.Ebx)

	// now load into ECX that shift amount, that can be
	// either foudn in a register or as immediate
	if byreg {
		if (op>>7)&1 != 0 {
			panic("bit7 in op2 with reg, should not be here")
		}
		// cpu.Regs[15] += 4
		j.Add(a.Imm{4}, j.oArmReg(15))
		// move shift amount from armreg into ECX
		j.Movl(j.oArmReg((op>>8)&0xF), a.Ecx)
		// and ecx & 0xFF
		j.And(a.Imm{0xFF}, a.Ecx)
		// if ecx == 0 -> jump forward (ebx is ok as-is)
		op2end := j.JccShortForward(a.CC_Z)
		j.AddCycles(1)

		switch shtype {
		case 3: // rot
			j.RorCl(a.Ebx)
			if setcarry {
				// set carry from x86 sign. We can't rely on the x86 carry
				// flag because it is different when CL=32 (for x86 it means
				// 0, so x86 carry is not affected).
				j.Test(a.Ebx, a.Ebx)
				j.Setcc(a.CC_S, a.R10)
			}
		case 2: // asr
			// Calculate shift = max(shift, 31). We actually put 0xFFFFFFFF
			// in ECX, but that is parsed as 31 by x86
			j.Cmp(a.Imm{32}, a.Ecx)
			j.Sbb(a.Eax, a.Eax)
			j.Not(a.Eax)
			j.Or(a.Eax, a.Ecx)

			// Shift right. This is now always performed correctly as we
			// maxed out the value before.
			j.SarCl(a.Ebx)

			if setcarry {
				j.Setcc(a.CC_C, a.R10) // x86 carry in R10

				// If the shift value was >= 32, EBX is either 0 or FFFFFFFF,
				// and the carry must be 0 or 1 (respectively).
				j.Test(a.Eax, a.Eax)
				j.Cmovcc(a.CC_NZ, a.Ebx, a.R10d)
				j.And(a.Imm{1}, a.R10d)
			}

		case 0, 1: // lsl / lsr
			if shtype == 0 {
				j.ShlCl(a.Ebx)
			} else {
				j.ShrCl(a.Ebx)
			}
			if !setcarry {
				// Adjust shifts for amounts >= 32; in ARM, shift amounts
				// are well-defined for amounts >= 32, like in Go.
				j.Cmp(a.Imm{32}, a.Ecx)
				j.Sbb(a.Eax, a.Eax)
				j.And(a.Eax, a.Ebx)
			} else {
				// We need to both adjust the result for shift >= 32 and
				// compute carry flag. The ARM carry flag can be computed like this:
				//   shift < 32: use x86 carry
				//   shift == 32: nothing was shifted (it's shift=0 in x86 semantic);
				//                use bit 0 or 31 of EBX (depending on shift direction)
				//   shift > 32: carry must be zero
				j.Setcc(a.CC_C, a.R10) // x86 carry in R10
				if shtype == 0 {
					j.Bt(a.Imm{0}, a.Ebx)
				} else {
					j.Bt(a.Imm{31}, a.Ebx)
				}
				j.Setcc(a.CC_C, a.R11) // EBX bit 0 or 31 in R11 (this will only be used if shift==32)

				j.Cmp(a.Imm{32}, a.Ecx)
				j.Cmovcc(a.CC_Z, a.R11, a.R10) // shift == 32 -> EBX 0/31 bit in R10
				j.Sbb(a.Eax, a.Eax)
				j.And(a.Eax, a.Ebx)

				j.Cmp(a.Imm{33}, a.Ecx) // shift >= 33 -> clear R10
				j.Sbb(a.Eax, a.Eax)
				j.And(a.Eax, a.Ebx)
				j.And(a.Eax, a.R10d)
			}
		}

		if setcarry {
			j.CopyFlags(jitFlagCR10)
		}

		op2end()
	} else {
		shift := (op >> 7) & 0x1F

		switch shtype {
		case 0: // lsl
			if shift == 0 {
				return
			}
			j.Shl(a.Imm{int32(shift)}, a.Ebx)
			if setcarry {
				j.CopyFlags(jitFlagCAdd)
			}
		case 1, 2: // lsr/asr
			if shift == 0 {
				// Equal to >>32 in Go, so bit31 is carry
				// and then clear the output
				if setcarry {
					j.Bt(a.Imm{31}, a.Ebx)
					j.CopyFlags(jitFlagCAdd)
				}
				j.Xor(a.Ebx, a.Ebx)
			} else {
				if shtype == 1 {
					j.Shr(a.Imm{int32(shift)}, a.Ebx)
				} else {
					j.Sar(a.Imm{int32(shift)}, a.Ebx)
				}
				if setcarry {
					j.CopyFlags(jitFlagCAdd)
				}
			}
		case 3: // ror
			if shift == 0 {
				// shift == 0 -> rcr #1
				j.Bt(a.Imm{29}, jitRegCpsr)
				j.Rcr(a.Imm{1}, a.Ebx)
			} else {
				j.Ror(a.Imm{int32(shift)}, a.Ebx)
			}
			if setcarry {
				j.CopyFlags(jitFlagCAdd)
			}
		}
	}
}

func (j *jitArm) emitOpAlu(op uint32) {
	imm := (op>>25)&1 != 0
	code := (op >> 21) & 0xF
	setflags := (op>>20)&1 != 0
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF

	// ADC, RSC, SBC needs the initial carry; load it into
	// R9 before it is potentially changed by the op2 decoding
	if code == 5 || code == 7 || code == 6 {
		j.Bt(a.Imm{29}, jitRegCpsr)
		j.Setcc(a.CC_C, a.R9)
	}

	if imm {
		rot := uint((op >> 7) & 0x1E)
		op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))

		j.Mov(a.Imm{int32(op2)}, a.Rbx)

		if setflags {
			if rot != 0 {
				if op2>>31 != 0 {
					j.Or(a.Imm{1 << 29}, jitRegCpsr)
				} else {
					j.And(a.Imm{^(1 << 29)}, jitRegCpsr)
				}
			}
		}
	} else {
		j.emitAluOp2Reg(op, setflags)
	}

	destreg := a.Eax
	flags := 0
	test := false
	switch code {
	case 8: // TST
		test = true
		fallthrough
	case 0: // AND
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.And(a.Ebx, a.Eax)
	case 9: // TEQ
		test = true
		fallthrough
	case 1: // XOR
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Xor(a.Ebx, a.Eax)
	case 10: // CMP
		test = true
		fallthrough
	case 2: // SUB
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Sub(a.Ebx, a.Eax)
		flags |= jitFlagCSub | jitFlagV
	case 3: // RSB
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Sub(a.Eax, a.Ebx)
		destreg = a.Ebx
		flags |= jitFlagCSub | jitFlagV
	case 11: // CMN
		test = true
		fallthrough
	case 4: // ADD
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Add(a.Ebx, a.Eax)
		flags |= jitFlagCAdd | jitFlagV
	case 5: // ADC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Bt(a.Imm{0}, a.R9) // load ARM carry into carry flag
		j.Adc(a.Ebx, a.Eax)
		flags |= jitFlagCAdd | jitFlagV
	case 7: // RSC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Bt(a.Imm{0}, a.R9) // load ARM carry into carry flag
		j.Cmc()              // complement carry: for subtraction, it's reversed
		j.Sbb(a.Eax, a.Ebx)
		flags |= jitFlagCSub | jitFlagV
		destreg = a.Ebx
	case 6: // SBC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Bt(a.Imm{0}, a.R9) // load ARM carry into carry flag
		j.Cmc()              // complement carry: for subtraction, it's reversed
		j.Sbb(a.Ebx, a.Eax)
		flags |= jitFlagCSub | jitFlagV
	case 12: // ORR
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Or(a.Ebx, a.Eax)
	case 13: // MOV
		if rnx != 0 {
			panic("rnx!=0 on MOV")
		}
		destreg = a.Ebx
		if setflags {
			// load the N/Z carry flags
			j.Test(a.Ebx, a.Ebx)
		}
	case 14: // BIC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Not(a.Ebx)
		j.And(a.Ebx, a.Eax)
	case 15: // MVN
		j.Not(a.Ebx)
		// NOT does not touch flags, so add a TEST in case we need N/Z
		if setflags {
			j.Test(a.Ebx, a.Ebx)
		}
		destreg = a.Ebx
	default:
		panic("unreachable")
	}

	// Update flags only if dest not PC
	// See ARM7TDMI Manual, §4.5.4
	// TODO: is this valid for ARM9 as well?
	if setflags && rdx != 15 {
		j.CopyFlags(flags | jitFlagN | jitFlagZ)
	}

	if !test {
		j.Movl(destreg, j.oArmReg(rdx))

		if rdx == 15 {
			bflags := 0
			if setflags {
				bflags = branchFlagCpsrRestore
			}
			// Clean only BIT 0, as we might have just switched back to thumb
			j.Btr(a.Imm{0}, destreg)
			j.emitBranch(destreg, BranchJump, bflags)
		}
	} else {
		if !setflags {
			panic("unreachable")
		}
		if rdx != 0 && rdx != 15 {
			panic("invalid rdx on test")
		}
	}
}

func (j *jitArm) emitOpMemory(op uint32) {
	shreg := (op>>25)&1 != 0
	pre := (op>>24)&1 != 0
	up := (op>>23)&1 != 0
	byt := (op>>22)&1 != 0
	wb := (op>>21)&1 != 0
	load := (op>>20)&1 != 0

	if op>>8 == 0xF {
		panic("PLD not supported")
	}

	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF

	j.Movl(j.oArmReg(rnx), a.Eax)
	j.Add(a.Imm{4}, j.oArmReg(15)) // cpu.Regs[15]+=4

	var off a.Operand
	if shreg {
		j.emitAluOp2Reg(op, false)
		off = a.Ebx
	} else {
		off = a.Imm{int32(op & 0xFFF)}
	}

	if pre {
		if up {
			j.Add(off, a.Eax)
		} else {
			j.Sub(off, a.Eax)
		}
	} else {
		// save computed offset for later, will be used during post
		j.Movl(off, j.FrameSlot(0x4, 32))
	}

	j.Movl(a.Eax, j.FrameSlot(0x0, 32)) // save address for later

	// Allocate stack frame to prepare for calls
	j.CallBlock(0x18, func() {
		if load {
			if byt {
				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.CallFuncGo((*Cpu).Read8)
				j.Xor(a.Edx, a.Edx) // FIXME: use MOVZX
				j.Movb(j.CallSlot(0x10, 8), a.Dl)
			} else {
				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.CallFuncGo((*Cpu).Read32)
				j.Movl(j.CallSlot(0x10, 32), a.Edx)
				j.Movl(j.FrameSlot(0x0, 32), a.Ecx) // restore address

				// rotate value read from memory in case address was misaligned
				// it's faster to always do it rather than checking
				j.And(a.Imm{3}, a.Ecx)
				j.Shl(a.Imm{3}, a.Ecx)
				j.RorCl(a.Edx)
			}
		} else {
			j.Mov(j.oArmReg(rdx), a.Edx)

			if byt {
				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.And(a.Imm{0xFF}, a.Edx) // FIXME: use MOVZX
				j.Movl(a.Edx, j.CallSlot(0xC, 32))
				j.CallFuncGo((*Cpu).Write8)
			} else {
				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.Movl(a.Edx, j.CallSlot(0xC, 32))
				j.CallFuncGo((*Cpu).Write32)
			}
		}
	})

	// Restore address if we need it
	if !pre || wb {
		j.Movl(j.FrameSlot(0x0, 32), a.Eax)
	}

	if !pre {
		// Restore offset. It wasn't added yet to the address since
		// we're in post-mode, so do it now
		j.Movl(j.FrameSlot(0x4, 32), a.Ebx)

		if up {
			j.Add(a.Ebx, a.Eax)
		} else {
			j.Sub(a.Ebx, a.Eax)
		}
		if wb {
			// writeback always enabled for post. wb bit is "force unprivileged"
			panic("unimplemented forced-unprivileged memory access")
		}
		wb = true
	}

	if wb {
		j.Movl(a.Eax, j.oArmReg(rnx))
	}

	if load {
		// Store the value read into the ARM CPU register
		// We must do this here, after having restored the block state
		j.Movl(a.Edx, j.oArmReg(rdx))
		if rdx == 15 {
			// Emit branch to target.
			j.emitBranch(a.Edx, BranchJump, branchFlagExchange)
		}
	}

	j.AddCycles(1)
}

func (j *jitArm) emitOpHalfWord(op uint32) {
	pre := (op>>24)&1 != 0
	up := (op>>23)&1 != 0
	imm := (op>>22)&1 != 0
	wb := (op>>21)&1 != 0
	if !pre {
		wb = true
	}
	load := (op>>20)&1 != 0
	code := (op >> 5) & 3

	if code == 0 {
		panic("invalid opcode decoded as LD/STR")
	}

	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF

	j.Movl(j.oArmReg(rnx), a.Eax)
	j.Add(a.Imm{4}, j.oArmReg(15)) // cpu.Regs[15]+=4

	var off a.Operand
	if imm {
		off = a.Imm{int32((op & 0xF) | ((op & 0xF00) >> 4))}
	} else {
		rmx := op & 0xF
		if rmx == 15 {
			panic("halfword: unimplemented rm==15")
		}
		off = j.oArmReg(rmx)
	}

	if pre {
		if up {
			j.Add(off, a.Eax)
		} else {
			j.Sub(off, a.Eax)
		}
	} else {
		// save computed offset for later, will be used during post
		j.Movl(off, j.FrameSlot(0x4, 32))
	}

	j.Movl(a.Eax, j.FrameSlot(0x0, 32)) // save address for later

	// Allocate stack frame to prepare for calls
	j.CallBlock(0x18, func() {
		switch code {
		case 1: // LDRH/STRH
			if load {
				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.CallFuncGo((*Cpu).Read16)
				// FIXME: use 16-bit instructions when added to go-jit
				j.Movl(j.CallSlot(0x10, 32), a.Edx) // only lower 16-bits are used
				j.And(a.Imm{0xFFFF}, a.Edx)
				if j.Cpu.arch < ARMv5 {
					// Convert to branchless RORW once we have 16-bit opcodes
					j.Bt(a.Imm{0}, j.FrameSlot(0x0, 32))
					close := j.JccShortForward(a.CC_NC)
					j.Movl(a.Edx, a.Ecx)
					j.Shr(a.Imm{8}, a.Edx)
					j.Shl(a.Imm{8}, a.Ecx)
					j.Or(a.Ecx, a.Edx)
					j.And(a.Imm{0xFFFF}, a.Edx)
					close()
				}
			} else {
				j.Movl(j.oArmReg(rdx), a.Edx)
				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.Movl(a.Edx, j.CallSlot(0xC, 32))
				j.CallFuncGo((*Cpu).Write16)
			}
		case 2: // LDRSB / LDRD
			if load {
				// LDRSB
				if rdx == 15 {
					panic("LDRSB PC not implemented")
				}
				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.CallFuncGo((*Cpu).Read8)
				j.Xor(a.Edx, a.Edx)
				j.Movb(j.CallSlot(0x10, 8), a.Dl) // FIXME: use MOVSX
				j.Shl(a.Imm{24}, a.Edx)
				j.Sar(a.Imm{24}, a.Edx)
			} else {
				// LDRD
				load = true // this is a load as well!
				panic("LDRD not implemented")
			}

		case 3: // LDRSH / STRD
			if load {
				// LDRSH
				if rdx == 15 {
					panic("LDRSH PC not implemented")
				}

				j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
				j.Movl(a.Eax, j.CallSlot(0x8, 32))
				j.CallFuncGo((*Cpu).Read16)
				// FIXME: use 16-bit instructions when added to go-jit
				j.Movl(j.CallSlot(0x10, 32), a.Edx) // only lower 16-bits are used
				j.Shl(a.Imm{16}, a.Edx)
				j.Sar(a.Imm{16}, a.Edx)
				if j.Cpu.arch < ARMv5 {
					// On ARMv4, LDRSH on unaligned address basically ignores
					// the lower byte and sign extends the higher
					j.Bt(a.Imm{0}, j.FrameSlot(0x0, 32))
					close := j.JccShortForward(a.CC_NC)
					j.Sar(a.Imm{8}, a.Edx)
					close()
				}
			} else {
				panic("STRD not implemented")
			}
		}
	})

	if load {
		j.Movl(a.Edx, j.oArmReg(rdx))
	}
	if wb {
		j.Movl(j.FrameSlot(0x0, 32), a.Eax)

		if !pre {
			if up {
				j.Add(j.FrameSlot(0x4, 32), a.Eax)
			} else {
				j.Sub(j.FrameSlot(0x4, 32), a.Eax)
			}
		}

		j.Movl(a.Eax, j.oArmReg(rnx))
	}

	j.AddCycles(1)
}

func (j *jitArm) emitOpSwp(op uint32) {
	byt := (op>>22)&1 != 0
	if (op>>24)&0xF != 1 || ((op>>20)&0xF != 0 && (op>>20)&0xF != 4) {
		panic("invalid call to emitOpSwp")
	}
	if op&0x0FB00FF0 != 0x01000090 {
		panic("invalid opcode decoded as SWP")
	}

	rnx := (op >> 16) & 0xF
	rmx := (op >> 0) & 0xF
	rdx := (op >> 12) & 0xF
	j.Movl(j.oArmReg(rnx), a.Eax)
	j.Movl(j.oArmReg(rmx), a.Ebx)

	j.Movl(a.Eax, j.FrameSlot(0x10, 32)) // save address for later
	j.Movl(a.Ebx, j.FrameSlot(0x14, 32)) // save value to write to memory for later

	j.CallBlock(0x18, func() {
		// edx := cpu.Read8/32(rn)
		if byt {
			j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
			j.Movl(a.Eax, j.CallSlot(0x8, 32))
			j.CallFuncGo((*Cpu).Read8)
			j.Xor(a.Edx, a.Edx) // FIXME: use MOVZX
			j.Movb(j.CallSlot(0x10, 8), a.Dl)
		} else {
			j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
			j.Movl(a.Eax, j.CallSlot(0x8, 32))
			j.CallFuncGo((*Cpu).Read32)
			j.Movl(j.CallSlot(0x10, 32), a.Edx)
			j.Movl(j.FrameSlot(0x10, 32), a.Ecx) // restore address

			// rotate value read from memory in case address was misaligned
			// it's faster to always do it rather than checking
			j.And(a.Imm{3}, a.Ecx)
			j.Shl(a.Imm{3}, a.Ecx)
			j.RorCl(a.Edx)
		}
	})

	j.Movl(a.Edx, j.FrameSlot(0x18, 32)) // save value to write to register for later
	j.Movl(j.FrameSlot(0x10, 32), a.Eax) // address
	j.Movl(j.FrameSlot(0x14, 32), a.Edx) // value to write

	j.CallBlock(0x18, func() {
		// cpu.Write8/32(rn, rm)
		if byt {
			j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
			j.Movl(a.Eax, j.CallSlot(0x8, 32))
			j.And(a.Imm{0xFF}, a.Edx) // FIXME: use MOVZX
			j.Movl(a.Edx, j.CallSlot(0xC, 32))
			j.CallFuncGo((*Cpu).Write8)
		} else {
			j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
			j.Movl(a.Eax, j.CallSlot(0x8, 32))
			j.Movl(a.Edx, j.CallSlot(0xC, 32))
			j.CallFuncGo((*Cpu).Write32)
		}
	})

	j.Movl(j.FrameSlot(0x18, 32), a.Edx) // value to write to register
	j.Movl(a.Edx, j.oArmReg(rdx))
	j.AddCycles(1)
}

func (j *jitArm) emitOpSwi(op uint32) {
	j.CallBlock(0x10, func() {
		j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
		j.Mov(a.Imm{int32(ExceptionSwi)}, j.CallSlot(0x8, 64))
		j.CallFuncGo((*Cpu).Exception)
	})
	j.AddCycles(2)
}

func (j *jitArm) emitOpClz(op uint32) {
	if op&0x0FFF0FF0 != 0x016F0F10 {
		panic("invalid opcode decoded as clz")
	}
	if j.Cpu.arch < ARMv5 {
		panic("invalid CLZ opcode on pre-ARMv5 CPU")
	}

	rmx := op & 0xF
	rdx := (op >> 12) & 0xF

	j.Xor(a.Rax, a.Rax)
	j.Movl(j.oArmReg(rmx), a.Eax)
	j.Shl(a.Imm{1}, a.Rax)
	j.Bts(a.Imm{0}, a.Rax)
	j.Bsr(a.Rax, a.Rbx)
	j.Sub(a.Imm{32}, a.Ebx)
	j.Neg(a.Ebx)
	j.Movl(a.Ebx, j.oArmReg(rdx))
}

func (j *jitArm) emitOpMul(op uint32) {
	setflags := (op>>20)&1 != 0
	code := (op >> 21) & 0xF
	acc := code&1 != 0
	htopy := (op>>6)&1 != 0
	htopx := (op>>5)&1 != 0
	halfwidth := code >= 8

	if halfwidth {
		if setflags {
			panic("half-width multiply with setflags")
		}
		if j.Cpu.arch < ARMv5 {
			panic("half-width multiply not available before ARMv5")
		}
	}

	rsx := (op >> 8) & 0xF
	rmx := (op >> 0) & 0xF
	rdx := (op >> 16) & 0xF
	rnx := (op >> 12) & 0xF // used for MLA

	// Need this 64-bit safe for 64-bit ops and for computing the cycles
	j.Xor(a.Rax, a.Rax)
	j.Movl(j.oArmReg(rsx), a.Eax)

	if !halfwidth {
		j.Xor(a.Rdx, a.Rdx)

		// First MSB set in RS (if RS==0, then 0)
		j.Mov(a.Rax, a.Rbx)
		j.Bsr(a.Rbx, a.Rcx)
		j.Cmovcc(a.CC_Z, a.Rdx, a.Rcx)

		// First MSB set in ^RS (if ^RS==0, then 0)
		j.Not(a.Ebx)
		j.Bsr(a.Rbx, a.Rdi)
		j.Cmovcc(a.CC_Z, a.Rdx, a.Rdi)

		// Take the minimum between the two indices
		j.Cmp(a.Rdi, a.Rcx)
		j.Cmovcc(a.CC_A, a.Rdi, a.Rcx)

		// Convert into nibble count
		j.Shr(a.Imm{3}, a.Rcx)

		// Compute cycles: number of nibbles + 1 (then, if mla, +1)
		if acc {
			j.Add(a.Imm{2}, a.Rcx)
		} else {
			j.Inc(a.Rcx)
		}
		j.Add(a.Rcx, oCycles)
	}

	switch code {
	case 0, 1: // MUL, MLA
		j.Mul(j.oArmReg(rmx))
		if acc {
			j.Add(j.oArmReg(rnx), a.Eax)
		}
		if setflags {
			if !acc { // no need for TEST after ADD
				j.Test(a.Eax, a.Eax)
			}
			j.CopyFlags(jitFlagN | jitFlagZ)
		}
		j.Movl(a.Eax, j.oArmReg(rdx))
	case 4, 5, 6, 7: // UMULL, UMLAL, SMULL, SMLAL
		j.Movl(j.oArmReg(rmx), a.Ebx) // FIXME: use MOVSX
		if code < 6 {
			j.Mul(a.Rbx)
		} else {
			j.Shl(a.Imm{32}, a.Rbx)
			j.Shl(a.Imm{32}, a.Rax)
			j.Sar(a.Imm{32}, a.Rbx)
			j.Sar(a.Imm{32}, a.Rax)
			j.Imul(a.Rbx)
		}
		if acc {
			j.Xor(a.Rdx, a.Rdx)
			j.Xor(a.Rcx, a.Rcx)
			j.Movl(j.oArmReg(rdx), a.Edx)
			j.Movl(j.oArmReg(rnx), a.Ecx)
			j.Shl(a.Imm{32}, a.Rdx)
			j.Add(a.Rcx, a.Rdx)
			j.Add(a.Rdx, a.Rax)
		}
		if setflags {
			if !acc { // no need for TEST after ADD
				j.Test(a.Rax, a.Rax)
			}
			j.CopyFlags(jitFlagN | jitFlagZ)
		}
		j.Movl(a.Eax, j.oArmReg(rnx))
		j.Shr(a.Imm{32}, a.Rax)
		j.Movl(a.Eax, j.oArmReg(rdx))
	case 8, 0xb: // SMLAxy/SMULxy
		j.Movl(j.oArmReg(rmx), a.Ebx)
		if !htopx {
			j.Shl(a.Imm{16}, a.Ebx)
		}
		if !htopy {
			j.Shl(a.Imm{16}, a.Eax)
		}
		j.Sar(a.Imm{16}, a.Ebx)
		j.Sar(a.Imm{16}, a.Eax)
		j.Imul(a.Ebx)
		if code == 8 { // SMLAxy
			j.Add(j.oArmReg(rnx), a.Eax)
		}
		j.Movl(a.Eax, j.oArmReg(rdx))
	case 9: // SMULWy/SMLAWy
		acc = !htopx // MLA flag is reversed in htopx
		j.Movl(j.oArmReg(rmx), a.Ebx)
		if htopy {
			j.Shr(a.Imm{16}, a.Rax)
		}
		j.Shl(a.Imm{32}, a.Rbx)
		j.Shl(a.Imm{48}, a.Rax)
		j.Sar(a.Imm{32}, a.Rbx)
		j.Sar(a.Imm{48}, a.Rax)
		j.Imul(a.Rbx)
		j.Sar(a.Imm{16}, a.Rax)
		if !htopx { // SMLAWy
			j.Add(j.oArmReg(rnx), a.Eax)
		}
		j.Movl(a.Eax, j.oArmReg(rdx))

	default:
		panic("unimplemented")
	}
}

func (j *jitArm) emitOpBlock(op uint32) {
	pre := (op>>24)&1 != 0
	up := (op>>23)&1 != 0
	psr := (op>>22)&1 != 0
	wb := (op>>21)&1 != 0
	load := (op>>20)&1 != 0

	rnx := (op >> 16) & 0xF
	mask := uint16(op & 0xFFFF)
	if rnx == 15 {
		panic("invalid use of PC in LDM/STM")
	}
	if mask == 0 {
		panic("unimplemented empty mask")
	}

	const (
		WbDisabled  = 0
		WbNormal    = 1
		WbUnchanged = 2
	)
	wbmode := WbDisabled
	if wb {
		wbmode = WbNormal

		// Handle special cases when rnx is included in the mask and
		// writeback is enabled
		if mask&(1<<rnx) != 0 {
			// check if it's first register in list
			if mask&((1<<rnx)-1) == 0 {
				wbmode = WbUnchanged
			} else {
				if load {
					onlyreg := mask & ^(1<<rnx) == 0
					lastreg := mask & ^((1<<rnx)-1) == (1 << rnx)
					if j.Cpu.arch >= ARMv5 && (onlyreg || !lastreg) {
						wbmode = WbNormal
					} else {
						wbmode = WbDisabled
					}
				} else {
					if j.Cpu.arch >= ARMv5 {
						wbmode = WbUnchanged
					} else {
						wbmode = WbNormal
					}
				}
			}
		}
	}

	// PSR bit is normally used to specify user bank access.
	// The only exception is when using LDM and PC is part of the regs;
	// in that case, PSR works as S bit in LDR, that is it also loads CPSR with SPSR.
	// So we keep a different
	var loadRestoreCpsr bool
	if psr && load && mask&0x8000 != 0 {
		loadRestoreCpsr = true
		psr = false
	}

	if psr {
		// Get current mode and save it into the frame for later
		j.Movl(jitRegCpsr, a.Edx)
		j.And(a.Imm{0x1F}, a.Edx)
		j.Movl(a.Edx, j.FrameSlot(0x10, 32))

		// Switch to CpuModeUser; this will allow LDM/STM to access the user bank
		j.emitCallCpsrSetMode(a.Imm{int32(CpuModeUser)})
	}

	nregs := popcount16(mask)
	j.Movl(j.oArmReg(rnx), a.Eax)
	if !up {
		j.Sub(a.Imm{int32(4 * nregs)}, a.Eax)
		pre = !pre
	}
	if !load {
		j.Add(a.Imm{4}, j.oArmReg(15)) // simulate prefetching
	}

	for i := uint32(0); mask != 0; i++ {
		if mask&1 != 0 {
			if pre {
				j.Add(a.Imm{4}, a.Eax)
			}
			j.Movl(a.Eax, j.FrameSlot(0x0, 32)) // save address for later
			if load {
				j.CallBlock(0x18, func() {
					j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
					j.Movl(a.Eax, j.CallSlot(0x8, 32)) // argument
					j.CallFuncGo((*Cpu).Read32)
					j.Movl(j.CallSlot(0x10, 32), a.Ebx) // return value
				})
				// Store into the register. We only avoid storing if
				// this is the base register, and we're in WbUnchanged mode
				if i != rnx || wbmode != WbUnchanged {
					j.Movl(a.Ebx, j.oArmReg(i))
				}
				if i == 15 {
					bflags := 0
					if loadRestoreCpsr {
						bflags |= branchFlagCpsrRestore
					}
					if j.Cpu.arch >= ARMv5 {
						bflags |= branchFlagExchange
					}
					j.emitBranch(a.Ebx, BranchJump, bflags)
				}
			} else {
				j.Movl(j.oArmReg(i), a.Ebx)
				j.CallBlock(0x10, func() {
					j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
					j.Movl(a.Eax, j.CallSlot(0x8, 32)) // argument 1: address
					j.Movl(a.Ebx, j.CallSlot(0xC, 32)) // argument 2: value
					j.CallFuncGo((*Cpu).Write32)
				})
			}
			j.Movl(j.FrameSlot(0x0, 32), a.Eax) // restore address
			if !pre {
				j.Add(a.Imm{4}, a.Eax)
			}
		}
		mask >>= 1
	}
	if wbmode == WbNormal {
		if !up {
			j.Sub(a.Imm{int32(4 * nregs)}, a.Eax)
		}
		j.Movl(a.Eax, j.oArmReg(rnx))
	}
	if psr {
		// Restore original mode
		j.Movl(j.FrameSlot(0x10, 32), a.Edx)
		j.emitCallCpsrSetMode(a.Edx)
	}

	j.AddCycles(1)
}

func (j *jitArm) emitBranch(tgt a.Register, reason BranchType, flags int) {
	if tgt == a.Edi {
		panic("cannot call emitBranch with target==Edi (used as temp register)")
	}

	if flags&branchFlagExchange != 0 {
		// branchless approach to emit:
		//    if rn&1 != 0 { cpu.Cpsr.SetT(true); rn &^= 1 } else { rn &^= 3 }
		// we do:
		//    tbit := rn&1
		//    rn ^= tbit
		//    cpu.Cpsr.r |= (tbit << 4)
		//    rn &^= (tbit << 1)
		j.Movl(tgt, a.Edi)     // copy address into EDX
		j.And(a.Imm{1}, a.Edi) // isolate bit 0; if set -> thumb
		j.Xor(a.Edi, tgt)      // turn off bit 0 in address (if it was set)
		j.Shl(a.Imm{5}, a.Edi)
		j.Or(a.Edi, jitRegCpsr) // copy bit 0 (thumb) into CPSR bit 5 (setT())
		j.Shr(a.Imm{4}, a.Edi)
		j.Xor(a.Imm{2}, a.Edi)
		j.Not(a.Edi)
		j.And(a.Edi, tgt)
	}

	if flags&branchFlagCpsrRestore != 0 {
		j.Movl(tgt, j.FrameSlot(0x20, 32)) // save for later
		// EMIT: cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		j.emitCallSpsr()
		j.Movl(a.Indirect{a.Rax, 0, 32}, a.Edi)
		j.emitCallCpsrSetWithMask(a.Edi, 0xFFFFFFFF)
		j.Movl(j.FrameSlot(0x20, 32), tgt)
	}

	j.CallBlock(0x10, func() {
		// Get the closure to the cpu.branch(). Declare type so that we
		// get compiler error if it changes (we might need to update code below)
		var funcBranch func(*Cpu, reg, BranchType) = (*Cpu).branch

		j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
		j.Movl(tgt, j.CallSlot(0x8, 32))
		j.Movl(a.Imm{int32(reason)}, j.CallSlot(0xC, 32))
		j.CallFuncGo(funcBranch)
	})
}

func (j *jitArm) emitLink() {
	// emit: cpu.Regs[14] = cpu.Regs[15]-4
	j.Movl(j.oArmReg(15), a.Ebx)
	j.Sub(a.Imm{4}, a.Ebx)
	j.Movl(a.Ebx, j.oArmReg(14))
}

func (j *jitArm) emitOpBx(op uint32) {
	link := op&0x20 != 0
	if op&0x0FFFFFD0 != 0x012FFF10 {
		panic("invalid opcode decoded as BX/BLX")
	}

	rnx := op & 0xF
	j.Movl(j.oArmReg(rnx), a.Eax)
	if link {
		j.emitLink()
	}

	if link {
		j.emitBranch(a.Eax, BranchCall, branchFlagExchange)
	} else {
		j.emitBranch(a.Eax, BranchJump, branchFlagExchange)
	}
}

func (j *jitArm) emitOpBranch(op uint32) {
	link := op&(1<<24) != 0

	off := int32(op<<8) >> 6

	if op>>28 == 0xF {
		// BLX_imm
		// BLX is always a link-branch, and linkbit is used as halfword offset
		if link {
			off += 2
		}
		j.emitLink()
		j.Bts(a.Imm{5}, jitRegCpsr) // set T flag
	} else {
		// B/BL
		if link {
			j.emitLink()
		}
	}

	j.Add(a.Imm{off}, j.oArmReg(15))
	j.Movl(j.oArmReg(15), a.Eax)
	j.emitBranch(a.Eax, BranchCall, 0)
}

func (j *jitArm) emitCallSpsr() {
	j.CallBlock(0x10, func() {
		var cpuRegSpsr func(*Cpu) *reg = (*Cpu).RegSpsr
		j.Mov(jitRegCpu, j.CallSlot(0x0, 64))
		j.CallFuncGo(cpuRegSpsr)
		j.Mov(j.CallSlot(0x8, 64), a.Rax)
	})
}

func (j *jitArm) emitCallCpsrSetMode(mode a.Operand) {
	j.CallBlock(0x10, func() {
		var cpuSetMode func(CpuMode, *Cpu) = j.Cpu.Cpsr.SetMode
		j.Movl(mode, j.CallSlot(0x0, 32))
		j.MovAbs(uint64(uintptr(unsafe.Pointer(j.Cpu))), a.Rax)
		j.Mov(a.Rax, j.CallSlot(0x8, 64))
		j.CallFuncGo(cpuSetMode)
	})
}

func (j *jitArm) emitCallCpsrSetWithMask(val a.Operand, mask uint32) {
	j.CallBlock(0x10, func() {
		var cpuRegCpsrSetWithMask func(val uint32, mask uint32, cpu *Cpu) = j.Cpu.Cpsr.SetWithMask
		j.Movl(val, j.CallSlot(0x0, 32))
		j.Movl(a.Imm{int32(mask)}, j.CallSlot(0x4, 32))
		j.MovAbs(uint64(uintptr(unsafe.Pointer(j.Cpu))), a.Rax)
		j.Mov(a.Rax, j.CallSlot(0x8, 64))
		j.CallFuncGo(cpuRegCpsrSetWithMask)
	})
}

func (j *jitArm) emitOpPsrTransfer(op uint32) {
	imm := (op>>25)&1 != 0
	if (op>>26)&3 != 0 || (op>>23)&0x3 != 2 || (op>>20)&1 != 0 {
		panic("invalid psr decoding")
	}
	spsr := (op>>22)&1 != 0
	tostat := (op>>21)&1 != 0

	if imm {
		if op&0x0FB00000 != 0x03200000 {
			panic("invalid opcode decoded as PSR_imm")
		}
	} else {
		if op&0x0F900FF0 != 0x01000000 {
			panic("invalid opcode decoded as PSR_reg")
		}
	}

	if !tostat {
		// MRS: from SPR to REG
		mask := (op >> 16) & 0xF
		if mask != 0xF {
			panic("mask should be 0xF in MRS (is it SWP?)")
		}
		rdx := (op >> 12) & 0xF
		if rdx == 15 {
			panic("write to PC in MRS")
		}

		var valueop a.Operand
		if spsr {
			j.emitCallSpsr()
			j.Mov(a.Indirect{a.Rax, 0, 32}, a.Ebx)
			valueop = a.Ebx
		} else {
			valueop = jitRegCpsr
		}
		j.Movl(valueop, j.oArmReg(rdx))
	} else {
		// MSR: from REG to SPR
		mask := uint32(0)
		if op&(1<<19) != 0 {
			mask |= 0xFF000000
		}
		if op&(1<<18) != 0 {
			mask |= 0x00FF0000
		}
		if op&(1<<17) != 0 {
			mask |= 0x0000FF00
		}
		if op&(1<<16) != 0 {
			mask |= 0x000000FF
		}

		var valueop a.Operand
		if imm {
			val := op & 0xFF
			shcnt := uint(((op >> 8) & 0xF) * 2)
			val = (val >> shcnt) | (val << (32 - shcnt))

			val &= mask
			valueop = a.Imm{int32(val)}
		} else {
			rmx := op & 0xF
			j.Movl(j.oArmReg(rmx), a.Edx)
			j.And(a.Imm{int32(mask)}, a.Edx)
			valueop = a.Edx
		}

		if spsr {
			j.Movl(valueop, j.FrameSlot(0x0, 32)) // save for later
			j.emitCallSpsr()
			j.Movl(a.Indirect{a.Rax, 0, 32}, a.Ebx)
			j.And(a.Imm{int32(^mask)}, a.Ebx)
			j.Or(j.FrameSlot(0x0, 32), a.Ebx)
			j.Movl(a.Ebx, a.Indirect{a.Rax, 0, 32})
		} else {
			j.emitCallCpsrSetWithMask(valueop, mask)
		}
	}
}

func (j *jitArm) emitOpCoprocessor(op uint32) {
	copread := (op>>20)&1 != 0
	cdp := (op & 0x10) == 0

	opc := (op >> 21) & 0x7
	cn := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	copnum := (op >> 8) & 0xF
	cp := (op >> 5) & 0x7
	cm := (op >> 0) & 0xF

	if cdp {
		// CDP
		j.CallBlock(0x18, func() {
			var cpuOpCopExec func(uint32, uint32, uint32, uint32, uint32, uint32) = j.Cpu.opCopExec
			j.Movl(a.Imm{int32(copnum)}, j.CallSlot(0x0, 32))
			j.Movl(a.Imm{int32(opc)}, j.CallSlot(0x4, 32))
			j.Movl(a.Imm{int32(cn)}, j.CallSlot(0x8, 32))
			j.Movl(a.Imm{int32(cm)}, j.CallSlot(0xC, 32))
			j.Movl(a.Imm{int32(cp)}, j.CallSlot(0x10, 32))
			j.Movl(a.Imm{int32(rdx)}, j.CallSlot(0x14, 32))
			j.CallFuncGo(cpuOpCopExec)
		})
	} else if copread {
		// MRC
		j.CallBlock(0x18, func() {
			var cpuOpCopRead func(uint32, uint32, uint32, uint32, uint32) uint32 = j.Cpu.opCopRead
			j.Movl(a.Imm{int32(copnum)}, j.CallSlot(0x0, 32))
			j.Movl(a.Imm{int32(opc)}, j.CallSlot(0x4, 32))
			j.Movl(a.Imm{int32(cn)}, j.CallSlot(0x8, 32))
			j.Movl(a.Imm{int32(cm)}, j.CallSlot(0xC, 32))
			j.Movl(a.Imm{int32(cp)}, j.CallSlot(0x10, 32))
			j.CallFuncGo(cpuOpCopRead)
			j.Movl(j.CallSlot(0x18, 32), a.Eax)
		})
		if rdx == 15 {
			j.Shr(a.Imm{28}, a.Eax)
			j.Shl(a.Imm{4}, jitRegCpsr)
			j.Shl(a.Imm{28}, a.Eax)
			j.Shr(a.Imm{4}, jitRegCpsr)
			j.Or(a.Eax, jitRegCpsr)
		} else {
			j.Movl(a.Eax, j.oArmReg(rdx))
		}
	} else {
		// MCR
		j.Add(a.Imm{4}, j.oArmReg(15))
		j.Movl(j.oArmReg(rdx), a.Eax)
		j.CallBlock(0x18, func() {
			var cpuOpCopWrite func(uint32, uint32, uint32, uint32, uint32, uint32) = j.Cpu.opCopWrite
			j.Movl(a.Imm{int32(copnum)}, j.CallSlot(0x0, 32))
			j.Movl(a.Imm{int32(opc)}, j.CallSlot(0x4, 32))
			j.Movl(a.Imm{int32(cn)}, j.CallSlot(0x8, 32))
			j.Movl(a.Imm{int32(cm)}, j.CallSlot(0xC, 32))
			j.Movl(a.Imm{int32(cp)}, j.CallSlot(0x10, 32))
			j.Movl(a.Eax, j.CallSlot(0x14, 32))
			j.CallFuncGo(cpuOpCopWrite)
		})
	}

	j.AddCycles(1)
}

func (j *jitArm) emitOp(op uint32) {

	cond := op >> 28
	var jcctarget func()

	switch cond {
	case 0xE, 0xF:
	case 0x0: // Z
		j.Bt(a.Imm{30}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_NC)
	case 0x1: // !Z
		j.Bt(a.Imm{30}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_C)
	case 0x2: // C
		j.Bt(a.Imm{29}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_NC)
	case 0x3: // !C
		j.Bt(a.Imm{29}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_C)
	case 0x4: // N
		j.Bt(a.Imm{31}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_NC)
	case 0x5: // !N
		j.Bt(a.Imm{31}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_C)
	case 0x6: // V
		j.Bt(a.Imm{28}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_NC)
	case 0x7: // !V
		j.Bt(a.Imm{28}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_C)
	case 0x8: // C && !Z
		j.Bt(a.Imm{29}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_NC)
		j.Bt(a.Imm{30}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_C)
	case 0x9: // !C || Z
		j.Movl(jitRegCpsr, a.Eax)
		j.Shl(a.Imm{1}, a.Eax)
		j.Not(a.Eax)
		j.Or(jitRegCpsr, a.Eax)
		j.Bt(a.Imm{30}, a.Eax)
		jcctarget = j.JccForward(a.CC_NC)
	case 0xC: // !Z && N==V
		j.Bt(a.Imm{30}, jitRegCpsr)
		jcctarget = j.JccForward(a.CC_C)
		fallthrough
	case 0xA, 0xB: // N==V / N!=V
		j.Movl(jitRegCpsr, a.Eax)
		j.Shr(a.Imm{3}, a.Eax)
		j.Xor(jitRegCpsr, a.Eax)
		j.Bt(a.Imm{28}, a.Eax)
		if cond == 0xA || cond == 0xC {
			jcctarget = j.JccForward(a.CC_C)
		} else {
			jcctarget = j.JccForward(a.CC_NC)
		}
	case 0xD: // Z || N==V / N!=V
		j.Movl(jitRegCpsr, a.Eax)
		j.Movl(jitRegCpsr, a.Ebx)
		j.Shr(a.Imm{3}, a.Eax)
		j.Shr(a.Imm{2}, a.Ebx)
		j.Xor(jitRegCpsr, a.Eax)
		j.Or(a.Ebx, a.Eax)
		j.Bt(a.Imm{28}, a.Eax)
		jcctarget = j.JccForward(a.CC_NC)

	default:
		panic(fmt.Errorf("unimplemented cond %x", cond))
	}

	high := (op >> 20) & 0xFF
	low := (op >> 4) & 0xF
	switch {
	case high == 0x12 && low&0xD == 0x1:
		j.emitOpBx(op)
	case high == 0x16 && low == 0x1:
		j.emitOpClz(op)
	case (high & 0xFB) == 0x32:
		j.emitOpPsrTransfer(op)
	case (high&0xF9) == 0x10 && low == 0:
		j.emitOpPsrTransfer(op)
	case (high&0xF9) == 0x10 && low&0x9 == 0x8:
		j.emitOpMul(op) // half-word mul
	case (high&0xFC) == 0 && low&0xF == 0x9:
		j.emitOpMul(op)
	case (high&0xF8) == 8 && low&0xF == 0x9:
		j.emitOpMul(op)
	case (high&0xFB) == 0x10 && low&0xF == 0x9:
		j.emitOpSwp(op)
	case (high>>5) == 0 && low&0x9 == 9: // TransReg10 / TransImm10
		j.emitOpHalfWord(op)
	case (high>>5) == 0 && low&0x1 == 0:
		j.emitOpAlu(op)
	case (high>>5) == 0 && low&0x9 == 1:
		j.emitOpAlu(op)
	case (high >> 5) == 1:
		j.emitOpAlu(op)
	case (high>>5) == 2 || (high>>5) == 3: // TransImm9 / TransReg9
		j.emitOpMemory(op)
	case (high >> 5) == 4:
		j.emitOpBlock(op)
	case (high >> 5) == 5:
		j.emitOpBranch(op)
	case (high>>5) == 7 && (high>>4)&1 == 0:
		j.emitOpCoprocessor(op)
	case (high>>5) == 7 && (high>>4)&1 == 1:
		j.emitOpSwi(op)
	default:
		log.ModCpu.FatalZ("unsupported op").Hex32("op", op).End()
	}

	// Complete JCC instruction used for cond (if any)
	if jcctarget != nil {
		jcctarget()
	}
}

func (j *jitArm) doBeginBlock() {
	j.Mov(j.ArgSlot(0, 64), jitRegCpu)
	j.Movl(a.Indirect{jitRegCpu, cpuCpsrOff, 32}, a.R14d)
}

func (j *jitArm) doEndBlock() {
	j.Movl(a.R14d, a.Indirect{jitRegCpu, cpuCpsrOff, 32})
}

func (j *jitArm) EmitBlock(ops []uint32) (out func(*Cpu)) {

	// Setup frame pointer with a local frame
	j.frameSize = 20 * 4
	j.Sub(a.Imm{int32(j.frameSize + 8)}, a.Rsp)
	j.Mov(a.Rbp, a.Indirect{a.Rsp, int32(j.frameSize), 64})
	j.Lea(a.Indirect{a.Rsp, int32(j.frameSize), 64}, a.Rbp)

	j.doBeginBlock()

	closes := make([]func(), 0, len(ops)*2)
	for i, op := range ops {
		j.emitOp(op)

		// Emit: cpu.Cycles += 1
		// Notice that we can't cache cpu.Cycles into a x86 register
		// because it can be changed externally at any point
		// (eg: the memory bus can add wait states to it).
		j.Mov(oCycles, a.Rax)
		j.Inc(a.Rax)
		j.Mov(a.Rax, oCycles)

		if i != len(ops)-1 {
			// Emit: if cpu.Cycles >= targetCycles -> exit
			// Forward jump (predicted as unlikely)
			j.Cmp(oTargetCycles, a.Rax)
			cf := j.JccForward(a.CC_AE)
			closes = append(closes, cf)

			// if tightExit -> exit
			// Forward jump (predicted as unlikely)
			j.Testb(a.Imm{0}, oTightExit)
			cf = j.JccForward(a.CC_NZ)
			closes = append(closes, cf)
		}
	}

	// We reached the block epilogue, which is the target of all
	// jumps geenrated after each instruction; close the jumps
	// so that they point here.
	for _, cf := range closes {
		cf()
	}

	j.doEndBlock()

	j.Mov(a.Indirect{a.Rsp, int32(j.frameSize), 64}, a.Rbp)
	j.Add(a.Imm{int32(j.frameSize + 8)}, a.Rsp)
	j.Ret()

	// Padding to align to 16-byte boundary. This is a speed
	// trick to better use cachelines, compilers do it.
	for j.Off&15 != 0 {
		j.Int3()
	}

	// Build function wrapper
	j.BuildTo(&out)
	return
}
