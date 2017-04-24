package amd64

import (
	"fmt"
)

func (a *Assembler) Inc(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xff)
	o.ModRM(a, Register{})
}

func (a *Assembler) Dec(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xff)
	o.ModRM(a, Register{1, 0})
}

func (a *Assembler) Incb(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xfe)
	o.ModRM(a, Register{})
}

func (a *Assembler) Decb(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xfe)
	o.ModRM(a, Register{1, 0})
}

func (a *Assembler) Not(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xf7)
	o.ModRM(a, Register{2, 0})
}

func (a *Assembler) Notb(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xf6)
	o.ModRM(a, Register{2, 0})
}

func (a *Assembler) ShlCl(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xd3)
	o.ModRM(a, Register{4, 0})
}

func (a *Assembler) ShrCl(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xd3)
	o.ModRM(a, Register{5, 0})
}

func (a *Assembler) SarCl(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xd3)
	o.ModRM(a, Register{7, 0})
}

func (a *Assembler) RolCl(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xd3)
	o.ModRM(a, Register{0, 0})
}

func (a *Assembler) RorCl(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xd3)
	o.ModRM(a, Register{1, 0})
}

func (a *Assembler) RclCl(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xd3)
	o.ModRM(a, Register{2, 0})
}

func (a *Assembler) RcrCl(o Operand) {
	o.Rex(a, Register{})
	a.byte(0xd3)
	o.ModRM(a, Register{3, 0})
}

func (asm *Assembler) arithmeticImmReg(insn *Instruction, src Imm, dst Register) {
	if insn.imm_r != nil {
		asm.rex(false, false, false, dst.Val > 7)
		if len(insn.imm_r) > 1 {
			panic("unsupported")
		}
		asm.byte(insn.imm_r[0] | (dst.Val & 7))
	} else {
		asm.rex(dst.Bits == 64, false, false, dst.Val > 7)
		asm.bytes(insn.imm_rm.op)
		asm.modrm(MOD_REG, insn.imm_rm.sub, dst.Val&7)
	}
}

func (asm *Assembler) arithmeticRegReg(insn *Instruction, src Register, dst Register) {
	if insn.r_rm != nil {
		dst.Rex(asm, src)
		asm.bytes(insn.r_rm)
		dst.ModRM(asm, src)
	} else {
		src.Rex(asm, dst)
		asm.bytes(insn.rm_r)
		src.ModRM(asm, dst)
	}
}

func (asm *Assembler) Arithmetic(insn *Instruction, src, dst Operand) {
	switch s := src.(type) {
	case Imm:
		if dr, ok := dst.(Register); ok {
			asm.arithmeticImmReg(insn, s, dr)
		} else {
			dst.Rex(asm, Register{insn.imm_rm.sub, 0})
			asm.bytes(insn.imm_rm.op)
			dst.ModRM(asm, Register{insn.imm_rm.sub, 0})
		}
		immbits := insn.imm_rm.immbits
		if immbits == 0 {
			immbits = insn.bits
		}
		if immbits == 8 {
			asm.byte(byte(s.Val))
		} else {
			asm.int32(uint32(s.Val))
		}
		return
	case Register:
		if dr, ok := dst.(Register); ok {
			asm.arithmeticRegReg(insn, s, dr)
		} else {
			dst.Rex(asm, s)
			asm.bytes(insn.r_rm)
			dst.ModRM(asm, s)
		}
		return
	}
	// if the LHS is neither an immediate nor a register, the rhs
	// must be a register
	dr, ok := dst.(Register)
	if !ok {
		panic(fmt.Sprintf("arithmetic: %#v/%#v not supported!", src, dst))
	}

	src.Rex(asm, dr)
	asm.bytes(insn.rm_r)
	src.ModRM(asm, dr)
}

func (a *Assembler) MovAbs(src uint64, dst Register) {
	a.rex(true, false, false, dst.Val > 7)
	if len(InstMov.imm_r) > 1 {
		panic("unsupported")
	}
	a.byte(InstMov.imm_r[0] | (dst.Val & 7))
	a.int64(src)
}

func (a *Assembler) Add(src, dst Operand)   { a.Arithmetic(InstAdd, src, dst) }
func (a *Assembler) Addb(src, dst Operand)  { a.Arithmetic(InstAddb, src, dst) }
func (a *Assembler) Adc(src, dst Operand)   { a.Arithmetic(InstAdc, src, dst) }
func (a *Assembler) Adcb(src, dst Operand)  { a.Arithmetic(InstAdcb, src, dst) }
func (a *Assembler) And(src, dst Operand)   { a.Arithmetic(InstAnd, src, dst) }
func (a *Assembler) Andb(src, dst Operand)  { a.Arithmetic(InstAndb, src, dst) }
func (a *Assembler) Cmp(src, dst Operand)   { a.Arithmetic(InstCmp, src, dst) }
func (a *Assembler) Cmpb(src, dst Operand)  { a.Arithmetic(InstCmpb, src, dst) }
func (a *Assembler) Mov(src, dst Operand)   { a.Arithmetic(InstMov, src, dst) }
func (a *Assembler) Movl(src, dst Operand)  { a.Arithmetic(InstMovl, src, dst) }
func (a *Assembler) Movb(src, dst Operand)  { a.Arithmetic(InstMovb, src, dst) }
func (a *Assembler) Or(src, dst Operand)    { a.Arithmetic(InstOr, src, dst) }
func (a *Assembler) Orb(src, dst Operand)   { a.Arithmetic(InstOrb, src, dst) }
func (a *Assembler) Lea(src, dst Operand)   { a.Arithmetic(InstLea, src, dst) }
func (a *Assembler) Sbb(src, dst Operand)   { a.Arithmetic(InstSbb, src, dst) }
func (a *Assembler) Sbbb(src, dst Operand)  { a.Arithmetic(InstSbbb, src, dst) }
func (a *Assembler) Sub(src, dst Operand)   { a.Arithmetic(InstSub, src, dst) }
func (a *Assembler) Subb(src, dst Operand)  { a.Arithmetic(InstSubb, src, dst) }
func (a *Assembler) Test(src, dst Operand)  { a.Arithmetic(InstTest, src, dst) }
func (a *Assembler) Testb(src, dst Operand) { a.Arithmetic(InstTestb, src, dst) }
func (a *Assembler) Xor(src, dst Operand)   { a.Arithmetic(InstXor, src, dst) }
func (a *Assembler) Xorb(src, dst Operand)  { a.Arithmetic(InstXorb, src, dst) }

func (a *Assembler) Rol(src, dst Operand)  { a.Arithmetic(InstRol, src, dst) }
func (a *Assembler) Rolb(src, dst Operand) { a.Arithmetic(InstRolb, src, dst) }
func (a *Assembler) Ror(src, dst Operand)  { a.Arithmetic(InstRor, src, dst) }
func (a *Assembler) Rorb(src, dst Operand) { a.Arithmetic(InstRorb, src, dst) }
func (a *Assembler) Rcl(src, dst Operand)  { a.Arithmetic(InstRcl, src, dst) }
func (a *Assembler) Rclb(src, dst Operand) { a.Arithmetic(InstRclb, src, dst) }
func (a *Assembler) Rcr(src, dst Operand)  { a.Arithmetic(InstRcr, src, dst) }
func (a *Assembler) Rcrb(src, dst Operand) { a.Arithmetic(InstRcrb, src, dst) }
func (a *Assembler) Shl(src, dst Operand)  { a.Arithmetic(InstShl, src, dst) }
func (a *Assembler) Shlb(src, dst Operand) { a.Arithmetic(InstShlb, src, dst) }
func (a *Assembler) Shr(src, dst Operand)  { a.Arithmetic(InstShr, src, dst) }
func (a *Assembler) Shrb(src, dst Operand) { a.Arithmetic(InstShrb, src, dst) }
func (a *Assembler) Sar(src, dst Operand)  { a.Arithmetic(InstSar, src, dst) }
func (a *Assembler) Sarb(src, dst Operand) { a.Arithmetic(InstSarb, src, dst) }

func (a *Assembler) Bt(src, dst Operand)  { a.Arithmetic(InstBt, src, dst) }
func (a *Assembler) Btc(src, dst Operand) { a.Arithmetic(InstBtc, src, dst) }
func (a *Assembler) Bts(src, dst Operand) { a.Arithmetic(InstBts, src, dst) }
func (a *Assembler) Btr(src, dst Operand) { a.Arithmetic(InstBtr, src, dst) }

func (a *Assembler) Int3()  { a.byte(0xcc) }
func (a *Assembler) Ret()   { a.byte(0xc3) }
func (a *Assembler) Pushf() { a.byte(0x9c) }
func (a *Assembler) Popf()  { a.byte(0x9d) }

func (a *Assembler) Call(dst Operand) {
	if _, ok := dst.(Imm); ok {
		panic("can't call(Imm); use CallRel instead.")
	} else {
		a.byte(0xff)
		dst.ModRM(a, Register{0x2, 64})
	}
}

func (a *Assembler) CallRel(dst uintptr) {
	a.byte(0xe8)
	a.rel32(dst)
}

func (a *Assembler) Push(src Operand) {
	if imm, ok := src.(Imm); ok {
		a.byte(0x68)
		a.int32(uint32(imm.Val))
	} else {
		a.byte(0xff)
		src.ModRM(a, Register{0x6, 64})
	}
}

func (a *Assembler) Pop(dst Operand) {
	switch d := dst.(type) {
	case Imm:
		panic("can't pop imm")
	case Register:
		a.rex(false, false, false, d.Val > 7)
		a.byte(0x58 | (d.Val & 7))
	default:
		dst.Rex(a, Register{0x0, 64})
		a.byte(0x8f)
		dst.ModRM(a, Register{0x0, 64})
	}
}

func (a *Assembler) JmpRel(dst uintptr) {
	a.byte(0xe9)
	a.rel32(dst)
}

func (a *Assembler) JccShort(cc byte, off int8) {
	a.byte(0x70 | cc)
	a.byte(byte(off))
}

// Like JccShort, but the offset is calculated with a closure function
// Use like this:
//     closejcc := a.JccShortDelayed(CC_EQ)
//     // [emit code]
//     closejcc()  // jump here
//
func (a *Assembler) JccShortForward(cc byte) func() {
	a.byte(0x70 | cc)
	off := a.Off
	a.byte(0)
	base := a.Off
	return func() {
		end := a.Off
		if int(int8(end-base)) != end-base {
			panic("jcc: too far!")
		}
		a.Buf[off] = byte(int8(end - base))
	}
}

func (a *Assembler) JccForward(cc byte) func() {
	a.byte(0x0f)
	a.byte(0x80 | cc)
	off := a.Off
	a.int32(0)
	base := a.Off
	return func() {
		end := a.Off
		i := end - base
		a.Buf[off] = byte(i & 0xFF)
		a.Buf[off+1] = byte(i >> 8)
		a.Buf[off+2] = byte(i >> 16)
		a.Buf[off+3] = byte(i >> 24)
	}
}

func (a *Assembler) Setcc(cc byte, dst Register) {
	a.rex(false, false, false, dst.Val > 7)
	a.byte(0x0f)
	a.byte(0x90 | cc)
	dst.ModRM(a, dst)
}

func (a *Assembler) JccRel(cc byte, dst uintptr) {
	a.byte(0x0f)
	a.byte(0x80 | cc)
	a.rel32(dst)
}
