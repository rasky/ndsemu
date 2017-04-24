package arm

import (
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
	oCpsr         = a.Indirect{jitRegCpu, cpuCpsrOff, 64}
	oCycles       = a.Indirect{jitRegCpu, cpuClockOff, 64}
	oTargetCycles = a.Indirect{jitRegCpu, cpuTargetOff, 64}
	oTightExit    = a.Indirect{jitRegCpu, cpuTightOff, 64}
)

type jitArm struct {
	*a.Assembler
}

const (
	jitFlagCAdd = 1 << iota // Carry flag (result of ADD)
	jitFlagCSub             // Carry flag (result of SUB -> ARM reverse meaning)
	jitFlagCR10             // Carry flag already in R10 (ignore x86 flag)
	jitFlagV
	jitFlagN
	jitFlagZ
)

func (j *jitArm) oArmReg(rn uint32) a.Operand {
	return a.Indirect{jitRegCpu, cpuRegsOff + int32(rn*4), 32}
}

func (j *jitArm) AddCycles(ncycles int32) {
	if ncycles == 1 {
		j.Inc(oCycles)
	} else {
		j.Add(a.Imm{int32(ncycles)}, oCycles)
	}
}

func (j *jitArm) SetCarry(carry bool) {
	if carry {
		j.Or(a.Imm{1 << 29}, jitRegCpsr)
	} else {
		j.And(a.Imm{^(1 << 29)}, jitRegCpsr)
	}
}

// Generate sequence to read the ARM carry into the specified
// register (value will be 0 or 1)
func (j *jitArm) LoadCarry(dst a.Register) {
	j.Bt(a.Imm{29}, jitRegCpsr)
	j.Setcc(a.CC_C, dst)
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
func (j *jitArm) doAluOp2Reg(op uint32, setcarry bool) {
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

		if shtype == 3 { // rot
			j.RorCl(a.Ebx)
			if setcarry {
				// set carry from x86 sign. We can't rely on the x86 carry
				// flag because it is different when CL=32 (for x86 it means
				// 0, so x86 carry is not affected).
				j.Test(a.Ebx, a.Ebx)
				j.Setcc(a.CC_S, a.R10)
			}
		} else {
			switch shtype {
			case 0: // lsl
				j.ShlCl(a.Ebx)
			case 1: // lsr
				j.ShrCl(a.Ebx)
			case 2: // asr
				j.SarCl(a.Ebx)
			}
			if setcarry {
				j.Setcc(a.CC_C, a.R10)
			}
			// Adjust shifts for amounts >= 32; in ARM, shift amounts
			// are well-defined for amounts >= 32, like in Go.
			j.Cmp(a.Imm{32}, a.Ecx)
			j.Sbb(a.Eax, a.Eax)
			j.And(a.Eax, a.Ebx)
			if setcarry {
				// clear carry if shift>=32
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

func (j *jitArm) doOpAlu(op uint32) {
	imm := (op>>25)&1 != 0
	code := (op >> 21) & 0xF
	setflags := (op>>20)&1 != 0
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF

	// ADC, RSC, SBC needs the carry; load it into
	// R9 before it is potentially changed by the op2 decoding
	if code == 5 || code == 7 || code == 9 {
		j.LoadCarry(a.R9)
	}

	if imm {
		rot := uint((op >> 7) & 0x1E)
		op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))

		j.Mov(a.Imm{int32(op2)}, a.Rbx)

		if setflags {
			if rot != 0 {
				j.SetCarry(op2>>31 != 0)
			}
		}
	} else {
		j.doAluOp2Reg(op, setflags)
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
	case 4: // ADD
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Add(a.Ebx, a.Eax)
		flags |= jitFlagCAdd | jitFlagV
	case 5: // ADC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Bt(a.Imm{0}, a.R9) // load into carry flag
		j.Adc(a.Ebx, a.Eax)
		flags |= jitFlagCAdd | jitFlagV
	case 7: // RSC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Bt(a.Imm{0}, a.R9) // load into carry flag
		j.Sbb(a.Eax, a.Ebx)
		flags |= jitFlagCSub | jitFlagV
		destreg = a.Ebx
	case 6: // SBC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Bt(a.Imm{0}, a.R9) // load into carry flag
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
	case 14: // BIC
		j.Movl(j.oArmReg(rnx), a.Eax)
		j.Not(a.Ebx)
		j.And(a.Ebx, a.Eax)
	case 15: // MVN
		j.Not(a.Ebx)
		destreg = a.Ebx
	default:
		panic("unimplemented")
	}

	if setflags {
		j.CopyFlags(flags | jitFlagN | jitFlagZ)
	}

	if !test {
		j.Movl(destreg, j.oArmReg(rdx))

		if rdx == 15 {
			if setflags {
				// EMIT: cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				panic("unimplemented")
			}
			// EMIT: g.writeBranch("reg(res)&^1", "BranchJump")
			panic("unimplemented")
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

func (j *jitArm) DoOp(op uint32) {

	cond := op >> 28
	var jcctarget func()

	switch cond {
	case 0xE, 0xF:
	case 0x0: // Z
		j.Bt(a.Imm{30}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_NC)
	case 0x1: // !Z
		j.Bt(a.Imm{30}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_C)
	case 0x2: // C
		j.Bt(a.Imm{29}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_NC)
	case 0x3: // !C
		j.Bt(a.Imm{29}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_C)
	case 0x4: // N
		j.Bt(a.Imm{31}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_NC)
	case 0x5: // !N
		j.Bt(a.Imm{31}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_C)
	case 0x6: // V
		j.Bt(a.Imm{28}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_NC)
	case 0x7: // !V
		j.Bt(a.Imm{28}, jitRegCpsr)
		jcctarget = j.JccShortForward(a.CC_C)
	default:
		println(cond)
		panic("unimplemented")
	}

	high := (op >> 20) & 0xFF
	low := (op >> 4) & 0xF
	switch {
	case (high>>5) == 0 && low&0x1 == 0:
		j.doOpAlu(op)
	case (high>>5) == 0 && low&0x9 == 1:
		j.doOpAlu(op)
	case (high >> 5) == 1:
		j.doOpAlu(op)
	default:
		// log.ModCpu.FatalZ("unsupported op").Hex32("op", op).End()
	}

	// Complete JCC instruction used for cond (if any)
	if jcctarget != nil {
		jcctarget()
	}
}

func (j *jitArm) doBeginBlock() {
	j.Mov(a.Indirect{a.Rsp, 8, 64}, a.R15)
	j.Movl(a.Indirect{jitRegCpu, cpuCpsrOff, 32}, a.R14d)
}

func (j *jitArm) doEndBlock() {
	j.Movl(a.R14d, a.Indirect{jitRegCpu, cpuCpsrOff, 32})
	j.Ret()
}

func (j *jitArm) DoBlock(ops []uint32) (out func(*Cpu)) {
	j.doBeginBlock()

	closes := make([]func(), 0, len(ops)*2)
	for i, op := range ops {
		j.DoOp(op)

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

	// We reached the end: closes all pending jumps
	for _, cf := range closes {
		cf()
	}

	j.doEndBlock()

	// Build function wrapper
	j.BuildTo(&out)
	return
}
