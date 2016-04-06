// Generated on 2016-04-06 02:57:57.034909889 +0200 CEST
package arm

import "bytes"
import "strconv"

var opThumbAluTable = [16]func(*Cpu, uint16){
	(*Cpu).opThumbAlu00,
	(*Cpu).opThumbAlu01,
	(*Cpu).opThumbAlu02,
	(*Cpu).opThumbAlu03,
	(*Cpu).opThumbAlu04,
	(*Cpu).opThumbAlu05,
	(*Cpu).opThumbAlu06,
	(*Cpu).opThumbAlu07,
	(*Cpu).opThumbAlu08,
	(*Cpu).opThumbAlu09,
	(*Cpu).opThumbAlu0A,
	(*Cpu).opThumbAlu0B,
	(*Cpu).opThumbAlu0C,
	(*Cpu).opThumbAlu0D,
	(*Cpu).opThumbAlu0E,
	(*Cpu).opThumbAlu0F,
}
var disasmThumbAluTable = [16]func(*Cpu, uint16, uint32) string{
	(*Cpu).disasmThumbAlu00,
	(*Cpu).disasmThumbAlu01,
	(*Cpu).disasmThumbAlu02,
	(*Cpu).disasmThumbAlu03,
	(*Cpu).disasmThumbAlu04,
	(*Cpu).disasmThumbAlu05,
	(*Cpu).disasmThumbAlu06,
	(*Cpu).disasmThumbAlu07,
	(*Cpu).disasmThumbAlu08,
	(*Cpu).disasmThumbAlu09,
	(*Cpu).disasmThumbAlu0A,
	(*Cpu).disasmThumbAlu0B,
	(*Cpu).disasmThumbAlu0C,
	(*Cpu).disasmThumbAlu0D,
	(*Cpu).disasmThumbAlu0E,
	(*Cpu).disasmThumbAlu0F,
}

func (cpu *Cpu) opThumb00(op uint16) {
	// lsl
	rsx := (op >> 3) & 7
	rdx := op & 7
	offset := (op >> 6) & 0x1F
	rs := uint32(cpu.Regs[rsx])
	if offset != 0 {
		cpu.Cpsr.SetC(rs&(1<<(32-offset)) != 0)
	}
	res := rs << offset
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumb00(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("lsl       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64((op >> 6) & 0x1F)
	out.WriteString("#")
	out.WriteString(strconv.FormatInt(arg2, 10))
	return out.String()
}

func (cpu *Cpu) opThumb08(op uint16) {
	// lsr
	rsx := (op >> 3) & 7
	rdx := op & 7
	offset := (op >> 6) & 0x1F
	rs := uint32(cpu.Regs[rsx])
	if offset == 0 {
		offset = 32
	}
	cpu.Cpsr.SetC(rs&(1<<(offset-1)) != 0)
	res := rs >> offset
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumb08(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("lsr       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64((op >> 6) & 0x1F)
	out.WriteString("#")
	out.WriteString(strconv.FormatInt(arg2, 10))
	return out.String()
}

func (cpu *Cpu) opThumb10(op uint16) {
	// asr
	rsx := (op >> 3) & 7
	rdx := op & 7
	offset := (op >> 6) & 0x1F
	rs := uint32(cpu.Regs[rsx])
	if offset == 0 {
		offset = 32
	}
	cpu.Cpsr.SetC(rs&(1<<(offset-1)) != 0)
	res := uint32(int32(rs) >> offset)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumb10(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("asr       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64((op >> 6) & 0x1F)
	out.WriteString("#")
	out.WriteString(strconv.FormatInt(arg2, 10))
	return out.String()
}

func (cpu *Cpu) opThumb18(op uint16) {
	// add
	rsx := (op >> 3) & 7
	rdx := op & 7
	rs := uint32(cpu.Regs[rsx])
	rnx := (op >> 6) & 7
	val := uint32(cpu.Regs[rnx])
	res := rs + val
	cpu.Cpsr.SetC(res < rs)
	cpu.Cpsr.SetVAdd(rs, val, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumb18(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("add       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 6) & 7
	out.WriteString(RegNames[arg2])
	return out.String()
}

func (cpu *Cpu) opThumb1A(op uint16) {
	// sub
	rsx := (op >> 3) & 7
	rdx := op & 7
	rs := uint32(cpu.Regs[rsx])
	rnx := (op >> 6) & 7
	val := uint32(cpu.Regs[rnx])
	res := rs - val
	cpu.Cpsr.SetC(rs >= val)
	cpu.Cpsr.SetVSub(rs, val, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumb1A(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("sub       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 6) & 7
	out.WriteString(RegNames[arg2])
	return out.String()
}

func (cpu *Cpu) opThumb1C(op uint16) {
	// add #nn
	rsx := (op >> 3) & 7
	rdx := op & 7
	rs := uint32(cpu.Regs[rsx])
	val := uint32((op >> 6) & 7)
	res := rs + val
	cpu.Cpsr.SetC(res < rs)
	cpu.Cpsr.SetVAdd(rs, val, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumb1C(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("add       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64((op >> 6) & 7)
	out.WriteString("#")
	out.WriteString(strconv.FormatInt(arg2, 10))
	return out.String()
}

func (cpu *Cpu) opThumb1E(op uint16) {
	// sub #nn
	rsx := (op >> 3) & 7
	rdx := op & 7
	rs := uint32(cpu.Regs[rsx])
	val := uint32((op >> 6) & 7)
	res := rs - val
	cpu.Cpsr.SetC(rs >= val)
	cpu.Cpsr.SetVSub(rs, val, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumb1E(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("sub       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64((op >> 6) & 7)
	out.WriteString("#")
	out.WriteString(strconv.FormatInt(arg2, 10))
	return out.String()
}

func (cpu *Cpu) opThumb20(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[0] = reg(res)
}

func (cpu *Cpu) disasmThumb20(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("mov       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(op & 0xFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opThumb21(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[1] = reg(res)
}

func (cpu *Cpu) opThumb22(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[2] = reg(res)
}

func (cpu *Cpu) opThumb23(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[3] = reg(res)
}

func (cpu *Cpu) opThumb24(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[4] = reg(res)
}

func (cpu *Cpu) opThumb25(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[5] = reg(res)
}

func (cpu *Cpu) opThumb26(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[6] = reg(res)
}

func (cpu *Cpu) opThumb27(op uint16) {
	// mov
	imm := uint32(op & 0xFF)
	res := imm
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[7] = reg(res)
}

func (cpu *Cpu) opThumb28(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[0])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) disasmThumb28(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("cmp       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(op & 0xFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opThumb29(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[1])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) opThumb2A(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[2])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) opThumb2B(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[3])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) opThumb2C(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[4])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) opThumb2D(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[5])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) opThumb2E(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[6])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) opThumb2F(op uint16) {
	// cmp
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[7])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) opThumb30(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[0])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[0] = reg(res)
}

func (cpu *Cpu) disasmThumb30(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("add       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(op & 0xFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opThumb31(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[1])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[1] = reg(res)
}

func (cpu *Cpu) opThumb32(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[2])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[2] = reg(res)
}

func (cpu *Cpu) opThumb33(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[3])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[3] = reg(res)
}

func (cpu *Cpu) opThumb34(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[4])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[4] = reg(res)
}

func (cpu *Cpu) opThumb35(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[5])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[5] = reg(res)
}

func (cpu *Cpu) opThumb36(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[6])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[6] = reg(res)
}

func (cpu *Cpu) opThumb37(op uint16) {
	// add
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[7])
	res := rd + imm
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[7] = reg(res)
}

func (cpu *Cpu) opThumb38(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[0])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[0] = reg(res)
}

func (cpu *Cpu) disasmThumb38(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("sub       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(op & 0xFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opThumb39(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[1])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[1] = reg(res)
}

func (cpu *Cpu) opThumb3A(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[2])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[2] = reg(res)
}

func (cpu *Cpu) opThumb3B(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[3])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[3] = reg(res)
}

func (cpu *Cpu) opThumb3C(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[4])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[4] = reg(res)
}

func (cpu *Cpu) opThumb3D(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[5])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[5] = reg(res)
}

func (cpu *Cpu) opThumb3E(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[6])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[6] = reg(res)
}

func (cpu *Cpu) opThumb3F(op uint16) {
	// sub
	imm := uint32(op & 0xFF)
	rd := uint32(cpu.Regs[7])
	res := rd - imm
	cpu.Cpsr.SetC(rd >= imm)
	cpu.Cpsr.SetVSub(rd, imm, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[7] = reg(res)
}

func (cpu *Cpu) opThumb40(op uint16) {
	opThumbAluTable[(op>>6)&0xF](cpu, op)
}

func (cpu *Cpu) disasmThumb40(op uint16, pc uint32) string {
	return disasmThumbAluTable[(op>>6)&0xF](cpu, op, pc)
}

func (cpu *Cpu) opThumb44(op uint16) {
	// add(h)
	rdx := (op & 7) | (op&0x80)>>4
	rsx := ((op >> 3) & 0xF)
	rs := uint32(cpu.Regs[rsx])
	rd := uint32(cpu.Regs[rdx])
	cpu.Regs[rdx] = reg(rd + rs)
	if rdx == 15 {
		cpu.branch(cpu.Regs[15]&^1, BranchJump)
	}
}

func (cpu *Cpu) disasmThumb44(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("add       ")
	arg0 := (op & 7) | (op&0x80)>>4
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := ((op >> 3) & 0xF)
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumb45(op uint16) {
	// cmp(h)
	rdx := (op & 7) | (op&0x80)>>4
	rsx := ((op >> 3) & 0xF)
	rs := uint32(cpu.Regs[rsx])
	rd := uint32(cpu.Regs[rdx])
	res := rd - rs
	cpu.Cpsr.SetNZ(res)
	cpu.Cpsr.SetC(rd >= res)
	cpu.Cpsr.SetVSub(rd, rs, res)
}

func (cpu *Cpu) disasmThumb45(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("cmp       ")
	arg0 := (op & 7) | (op&0x80)>>4
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := ((op >> 3) & 0xF)
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumb46(op uint16) {
	// mov(h)
	rdx := (op & 7) | (op&0x80)>>4
	rsx := ((op >> 3) & 0xF)
	rs := uint32(cpu.Regs[rsx])
	cpu.Regs[rdx] = reg(rs)
	if rdx == 15 {
		cpu.branch(reg(rs)&^1, BranchJump)
	}
}

func (cpu *Cpu) disasmThumb46(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("mov       ")
	arg0 := (op & 7) | (op&0x80)>>4
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := ((op >> 3) & 0xF)
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumb47(op uint16) {
	// bx/blx
	rdx := (op & 7) | (op&0x80)>>4
	rsx := ((op >> 3) & 0xF)
	rs := uint32(cpu.Regs[rsx])
	if op&0x80 != 0 {
		cpu.Regs[14] = (cpu.Regs[15] - 2) | 1
	}
	newpc := reg(rs) &^ 1
	if rs&1 == 0 {
		cpu.Cpsr.SetT(false)
		newpc &^= 3
	}
	cpu.branch(newpc, BranchCall)
	_ = rdx
}

func (cpu *Cpu) disasmThumb47(op uint16, pc uint32) string {
	if op&0x80 != 0 {
		var out bytes.Buffer
		out.WriteString("blx       ")
		arg0 := ((op >> 3) & 0xF)
		out.WriteString(RegNames[arg0])
		return out.String()
	} else {
		var out bytes.Buffer
		out.WriteString("bx        ")
		arg0 := ((op >> 3) & 0xF)
		out.WriteString(RegNames[arg0])
		return out.String()
	}
}

func (cpu *Cpu) opThumb48(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[0] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb48(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldr       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint32((op & 0xFF) * 4)
	arg1 += uint32((pc + 4) &^ 2)
	arg1v := cpu.opRead32(arg1)
	out.WriteString("= 0x")
	out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	return out.String()
}

func (cpu *Cpu) opThumb49(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[1] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb4A(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[2] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb4B(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[3] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb4C(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[4] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb4D(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[5] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb4E(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[6] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb4F(op uint16) {
	// ldr pc
	pc := uint32(cpu.Regs[15]) &^ 2
	pc += uint32((op & 0xFF) * 4)
	cpu.Regs[7] = reg(cpu.opRead32(pc))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb50(op uint16) {
	// str
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.opWrite32(addr, uint32(cpu.Regs[rdx]))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb50(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("str       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb52(op uint16) {
	// strh
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.opWrite16(addr, uint16(cpu.Regs[rdx]))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb52(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("strh      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb54(op uint16) {
	// strb
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.opWrite8(addr, uint8(cpu.Regs[rdx]))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb54(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("strb      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb56(op uint16) {
	// ldsb
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.Regs[rdx] = reg(int8(cpu.opRead8(addr)))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb56(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldsb      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb58(op uint16) {
	// ldr
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.Regs[rdx] = reg(cpu.opRead32(addr))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb58(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldr       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb5A(op uint16) {
	// ldrh
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.Regs[rdx] = reg(cpu.opRead16(addr))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb5A(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldrh      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb5C(op uint16) {
	// ldrb
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.Regs[rdx] = reg(cpu.opRead8(addr))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb5C(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldrb      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb5E(op uint16) {
	// ldsh
	rox := (op >> 6) & 7
	rbx := (op >> 3) & 7
	rdx := op & 7
	addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])
	cpu.Regs[rdx] = reg(int16(cpu.opRead16(addr)))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb5E(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldsh      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 7
	arg1b := (op >> 6) & 7
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(RegNames[arg1b])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opThumb60(op uint16) {
	// str #nn
	offset := uint32((op >> 6) & 0x1F)
	rbx := (op >> 3) & 0x7
	rdx := op & 0x7
	rb := uint32(cpu.Regs[rbx])
	offset *= 4
	rd := uint32(cpu.Regs[rdx])
	cpu.opWrite32(rb+offset, rd)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb60(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("str       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 0x7
	arg1b := ((op >> 6) & 0x1F) * 4
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb68(op uint16) {
	// ldr #nn
	offset := uint32((op >> 6) & 0x1F)
	rbx := (op >> 3) & 0x7
	rdx := op & 0x7
	rb := uint32(cpu.Regs[rbx])
	offset *= 4
	cpu.Regs[rdx] = reg(cpu.opRead32(rb + offset))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb68(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldr       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 0x7
	arg1b := ((op >> 6) & 0x1F) * 4
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb70(op uint16) {
	// strb #nn
	offset := uint32((op >> 6) & 0x1F)
	rbx := (op >> 3) & 0x7
	rdx := op & 0x7
	rb := uint32(cpu.Regs[rbx])
	rd := uint8(cpu.Regs[rdx])
	cpu.opWrite8(rb+offset, rd)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb70(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("strb      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 0x7
	arg1b := (op >> 6) & 0x1F
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb78(op uint16) {
	// ldrb #nn
	offset := uint32((op >> 6) & 0x1F)
	rbx := (op >> 3) & 0x7
	rdx := op & 0x7
	rb := uint32(cpu.Regs[rbx])
	cpu.Regs[rdx] = reg(cpu.opRead8(rb + offset))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb78(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldrb      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 0x7
	arg1b := (op >> 6) & 0x1F
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb80(op uint16) {
	// strh #nn
	offset := uint32((op >> 6) & 0x1F)
	rbx := (op >> 3) & 0x7
	rdx := op & 0x7
	rb := uint32(cpu.Regs[rbx])
	offset *= 2
	rd := uint16(cpu.Regs[rdx])
	cpu.opWrite16(rb+offset, rd)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb80(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("strh      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 0x7
	arg1b := ((op >> 6) & 0x1F) * 2
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb88(op uint16) {
	// ldrh #nn
	offset := uint32((op >> 6) & 0x1F)
	rbx := (op >> 3) & 0x7
	rdx := op & 0x7
	rb := uint32(cpu.Regs[rbx])
	offset *= 2
	cpu.Regs[rdx] = reg(cpu.opRead16(rb + offset))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb88(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldrh      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 3) & 0x7
	arg1b := ((op >> 6) & 0x1F) * 2
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb90(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[0]))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb90(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("str       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := 13
	arg1b := (op & 0xFF) * 4
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb91(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[1]))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb92(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[2]))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb93(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[3]))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb94(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[4]))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb95(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[5]))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb96(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[6]))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb97(op uint16) {
	// str [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[7]))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb98(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[0] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumb98(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldr       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := 13
	arg1b := (op & 0xFF) * 4
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+4)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
	}
	return out.String()
}

func (cpu *Cpu) opThumb99(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[1] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb9A(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[2] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb9B(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[3] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb9C(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[4] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb9D(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[5] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb9E(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[6] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumb9F(op uint16) {
	// ldr [sp+nn]
	offset := (op & 0xFF) * 4
	sp := uint32(cpu.Regs[13])
	cpu.Regs[7] = reg(cpu.opRead32(sp + uint32(offset)))
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbA0(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[0] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) disasmThumbA0(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("add       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := 15
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64((op & 0xFF) * 4)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opThumbA1(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[1] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) opThumbA2(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[2] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) opThumbA3(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[3] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) opThumbA4(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[4] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) opThumbA5(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[5] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) opThumbA6(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[6] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) opThumbA7(op uint16) {
	// add pc
	offset := (op & 0xFF) * 4
	cpu.Regs[7] = (cpu.Regs[15] &^ 2) + reg(offset)
}

func (cpu *Cpu) opThumbA8(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[0] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) disasmThumbA8(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("add       ")
	arg0 := (op >> 8) & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := 13
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64((op & 0xFF) * 4)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opThumbA9(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[1] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) opThumbAA(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[2] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) opThumbAB(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[3] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) opThumbAC(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[4] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) opThumbAD(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[5] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) opThumbAE(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[6] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) opThumbAF(op uint16) {
	// add sp
	offset := (op & 0xFF) * 4
	cpu.Regs[7] = cpu.Regs[13] + reg(offset)
}

func (cpu *Cpu) opThumbB0(op uint16) {
	// add sp
	offset := (op & 0x7F) * 4
	if op&0x80 == 0 {
		cpu.Regs[13] += reg(offset)
	} else {
		cpu.Regs[13] -= reg(offset)
	}
}

func (cpu *Cpu) disasmThumbB0(op uint16, pc uint32) string {
	if op&0x80 == 0 {
		var out bytes.Buffer
		out.WriteString("add       ")
		arg0 := 13
		out.WriteString(RegNames[arg0])
		out.WriteString(", ")
		arg1 := int64((op & 0x7f) * 4)
		out.WriteString("#0x")
		out.WriteString(strconv.FormatInt(arg1, 16))
		return out.String()
	} else {
		var out bytes.Buffer
		out.WriteString("sub       ")
		arg0 := 13
		out.WriteString(RegNames[arg0])
		out.WriteString(", ")
		arg1 := int64((op & 0x7f) * 4)
		out.WriteString("#0x")
		out.WriteString(strconv.FormatInt(arg1, 16))
		return out.String()
	}
}

func (cpu *Cpu) opThumbB1(op uint16) {
	cpu.InvalidOpThumb(op, "not implemented")
}

func (cpu *Cpu) disasmThumbB1(op uint16, pc uint32) string {
	return "dw " + strconv.FormatInt(int64(op), 16)
}

func (cpu *Cpu) opThumbB4(op uint16) {
	// push
	count := popcount16(op & 0x1FF)
	sp := uint32(cpu.Regs[13])
	sp -= uint32(count * 4)
	cpu.Regs[13] = reg(sp)
	if (op>>0)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[0]))
		sp += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[1]))
		sp += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[2]))
		sp += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[3]))
		sp += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[4]))
		sp += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[5]))
		sp += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[6]))
		sp += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[7]))
		sp += 4
	}
	if (op>>8)&1 != 0 {
		cpu.opWrite32(sp, uint32(cpu.Regs[14]))
		sp += 4
	}
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumbB4(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("push      ")
	arg0 := op & 0xff
	out.WriteString("{")
	for i := 0; arg0 != 0; i++ {
		if arg0&1 != 0 {
			out.WriteString(RegNames[i])
			arg0 >>= 1
			if arg0 != 0 {
				out.WriteString(", ")
			}
		} else {
			arg0 >>= 1
		}
	}
	out.WriteString("}")
	return out.String()
}

func (cpu *Cpu) disasmThumbB5(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("push      ")
	arg0 := op&0xff | 0x4000
	out.WriteString("{")
	for i := 0; arg0 != 0; i++ {
		if arg0&1 != 0 {
			out.WriteString(RegNames[i])
			arg0 >>= 1
			if arg0 != 0 {
				out.WriteString(", ")
			}
		} else {
			arg0 >>= 1
		}
	}
	out.WriteString("}")
	return out.String()
}

func (cpu *Cpu) opThumbBC(op uint16) {
	// pop
	sp := uint32(cpu.Regs[13])
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(sp))
		sp += 4
	}
	if (op>>8)&1 != 0 {
		switch cpu.arch {
		case ARMv4:
			pc := reg(cpu.opRead32(sp) &^ 1)
			cpu.branch(pc, BranchReturn)
		case ARMv5:
			pc := reg(cpu.opRead32(sp))
			if pc&1 == 0 {
				cpu.Cpsr.SetT(false)
				pc = pc &^ 3
			} else {
				pc = pc &^ 1
			}
			cpu.branch(pc, BranchReturn)
		default:
			panic("unimplemented arch-dependent behavior")
		}
		sp += 4
	}
	cpu.Regs[13] = reg(sp)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumbBC(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("pop       ")
	arg0 := op & 0xff
	out.WriteString("{")
	for i := 0; arg0 != 0; i++ {
		if arg0&1 != 0 {
			out.WriteString(RegNames[i])
			arg0 >>= 1
			if arg0 != 0 {
				out.WriteString(", ")
			}
		} else {
			arg0 >>= 1
		}
	}
	out.WriteString("}")
	return out.String()
}

func (cpu *Cpu) disasmThumbBD(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("pop       ")
	arg0 := op&0xff | 0x8000
	out.WriteString("{")
	for i := 0; arg0 != 0; i++ {
		if arg0&1 != 0 {
			out.WriteString(RegNames[i])
			arg0 >>= 1
			if arg0 != 0 {
				out.WriteString(", ")
			}
		} else {
			arg0 >>= 1
		}
	}
	out.WriteString("}")
	return out.String()
}

func (cpu *Cpu) opThumbC0(op uint16) {
	// stm
	if op&(1<<0) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[0])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[0] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[0] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumbC0(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("stm       ")
	arg0r := (op >> 8) & 7
	arg0w := (op >> ((op >> 8) & 7)) & 1
	out.WriteString(RegNames[arg0r])
	if arg0w != 0 {
		out.WriteString("!")
	}
	out.WriteString(", ")
	arg1 := op & 0xFF
	out.WriteString("{")
	for i := 0; arg1 != 0; i++ {
		if arg1&1 != 0 {
			out.WriteString(RegNames[i])
			arg1 >>= 1
			if arg1 != 0 {
				out.WriteString(", ")
			}
		} else {
			arg1 >>= 1
		}
	}
	out.WriteString("}")
	return out.String()
}

func (cpu *Cpu) opThumbC1(op uint16) {
	// stm
	if op&(1<<1) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[1])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[1] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[1] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbC2(op uint16) {
	// stm
	if op&(1<<2) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[2])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[2] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[2] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbC3(op uint16) {
	// stm
	if op&(1<<3) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[3])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[3] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[3] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbC4(op uint16) {
	// stm
	if op&(1<<4) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[4])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[4] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[4] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbC5(op uint16) {
	// stm
	if op&(1<<5) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[5])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[5] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[5] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbC6(op uint16) {
	// stm
	if op&(1<<6) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[6])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[6] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[6] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbC7(op uint16) {
	// stm
	if op&(1<<7) != 0 {
		cpu.InvalidOpThumb(op, "unimplemented: base reg in register list in STM")
		return
	}
	ptr := uint32(cpu.Regs[7])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.opWrite32(ptr, uint32(cpu.Regs[15]))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[7] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[0]))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[1]))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[2]))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[3]))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[4]))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[5]))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[6]))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.opWrite32(ptr, uint32(cpu.Regs[7]))
		ptr += 4
	}
	if wb {
		cpu.Regs[7] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbC8(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[0])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[0] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if wb {
		cpu.Regs[0] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) disasmThumbC8(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ldm       ")
	arg0r := (op >> 8) & 7
	arg0w := (op >> ((op >> 8) & 7)) & 1
	out.WriteString(RegNames[arg0r])
	if arg0w != 0 {
		out.WriteString("!")
	}
	out.WriteString(", ")
	arg1 := op & 0xFF
	out.WriteString("{")
	for i := 0; arg1 != 0; i++ {
		if arg1&1 != 0 {
			out.WriteString(RegNames[i])
			arg1 >>= 1
			if arg1 != 0 {
				out.WriteString(", ")
			}
		} else {
			arg1 >>= 1
		}
	}
	out.WriteString("}")
	return out.String()
}

func (cpu *Cpu) opThumbC9(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[1])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[1] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if wb {
		cpu.Regs[1] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbCA(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[2])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[2] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if wb {
		cpu.Regs[2] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbCB(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[3])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[3] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if wb {
		cpu.Regs[3] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbCC(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[4])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[4] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if wb {
		cpu.Regs[4] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbCD(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[5])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[5] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if wb {
		cpu.Regs[5] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbCE(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[6])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[6] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if wb {
		cpu.Regs[6] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbCF(op uint16) {
	// ldm
	ptr := uint32(cpu.Regs[7])
	if op&0xFF == 0 {
		switch cpu.arch {
		case ARMv4:
			cpu.Regs[15] = reg(cpu.opRead32(ptr))
			ptr += 0x40
		case ARMv5:
			ptr += 0x40
		default:
			panic("unimplemented arch-dependent behavior")
		}
		cpu.Regs[7] = reg(ptr)
		return
	}
	wb := true
	if (op>>0)&1 != 0 {
		cpu.Regs[0] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>1)&1 != 0 {
		cpu.Regs[1] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>2)&1 != 0 {
		cpu.Regs[2] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>3)&1 != 0 {
		cpu.Regs[3] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>4)&1 != 0 {
		cpu.Regs[4] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>5)&1 != 0 {
		cpu.Regs[5] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>6)&1 != 0 {
		cpu.Regs[6] = reg(cpu.opRead32(ptr))
		ptr += 4
	}
	if (op>>7)&1 != 0 {
		cpu.Regs[7] = reg(cpu.opRead32(ptr))
		wb = false
		ptr += 4
	}
	if wb {
		cpu.Regs[7] = reg(ptr)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opThumbD0(op uint16) {
	// beq
	if cpu.Cpsr.Z() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD0(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("beq       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD1(op uint16) {
	// bne
	if !cpu.Cpsr.Z() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD1(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bne       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD2(op uint16) {
	// bhs
	if cpu.Cpsr.C() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD2(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bhs       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD3(op uint16) {
	// blo
	if !cpu.Cpsr.C() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD3(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("blo       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD4(op uint16) {
	// bmi
	if cpu.Cpsr.N() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD4(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bmi       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD5(op uint16) {
	// bpl
	if !cpu.Cpsr.N() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD5(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bpl       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD6(op uint16) {
	// bvs
	if cpu.Cpsr.V() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD6(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bvs       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD7(op uint16) {
	// bvc
	if !cpu.Cpsr.V() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD7(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bvc       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD8(op uint16) {
	// bhi
	if cpu.Cpsr.C() && !cpu.Cpsr.Z() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD8(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bhi       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbD9(op uint16) {
	// bls
	if !cpu.Cpsr.C() || cpu.Cpsr.Z() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbD9(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bls       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbDA(op uint16) {
	// bge
	if cpu.Cpsr.N() == cpu.Cpsr.V() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbDA(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bge       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbDB(op uint16) {
	// blt
	if cpu.Cpsr.N() != cpu.Cpsr.V() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbDB(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("blt       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbDC(op uint16) {
	// bgt
	if !cpu.Cpsr.Z() && cpu.Cpsr.N() == cpu.Cpsr.V() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbDC(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bgt       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbDD(op uint16) {
	// ble
	if cpu.Cpsr.Z() || cpu.Cpsr.N() != cpu.Cpsr.V() {
		offset := int8(uint8(op & 0xFF))
		offset32 := int32(offset) * 2
		cpu.branch(cpu.Regs[15]+reg(offset32), BranchJump)
	}
}

func (cpu *Cpu) disasmThumbDD(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ble       ")
	arg0 := int32(int32(int8(uint8(op&0xFF))) * 2)
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbDE(op uint16) {
	// b undef
	cpu.InvalidOpThumb(op, "invalid F16 with opcode==14")
}

func (cpu *Cpu) opThumbDF(op uint16) {
	// swi
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmThumbDF(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("swi       ")
	arg0 := int64(op & 0xFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg0, 16))
	return out.String()
}

func (cpu *Cpu) opThumbE0(op uint16) {
	// b
	pc := cpu.Regs[15] + reg(int32(int16(op<<5)>>4))
	cpu.branch(pc, BranchJump)
}

func (cpu *Cpu) disasmThumbE0(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("b         ")
	arg0 := int32(int32(int16(op<<5) >> 4))
	arg0x := pc + 4 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opThumbE8(op uint16) {
	// blx step 2
	newpc := cpu.Regs[14] + reg((op&0x7FF)<<1)
	cpu.Regs[14] = (cpu.Regs[15] - 2) | 1
	newpc &^= 2
	cpu.Cpsr.SetT(false)
	cpu.branch(newpc, BranchCall)
}

func (cpu *Cpu) disasmThumbE8(op uint16, pc uint32) string {
	return "[continued]"
}

func (cpu *Cpu) opThumbF0(op uint16) {
	// bl/blx step 1
	cpu.Regs[14] = cpu.Regs[15] + reg(int32(uint32(op&0x7FF)<<21)>>9)
}

func (cpu *Cpu) disasmThumbF0(op uint16, pc uint32) string {
	mem := cpu.opFetchPointer(pc + 2)
	op2 := uint16(mem[0]) | uint16(mem[1])<<8
	nextpc := (int32(uint32(op&0x7FF)<<21) >> 9) + int32((op2&0x7FF)<<1)
	if (op2>>12)&1 == 0 {
		var out bytes.Buffer
		out.WriteString("blx       ")
		arg0 := int32(nextpc)
		arg0x := pc + 4 + uint32(arg0)
		out.WriteString(strconv.FormatInt(int64(arg0x), 16))
		return out.String()
	} else {
		var out bytes.Buffer
		out.WriteString("bl        ")
		arg0 := int32(nextpc)
		arg0x := pc + 4 + uint32(arg0)
		out.WriteString(strconv.FormatInt(int64(arg0x), 16))
		return out.String()
	}
}

func (cpu *Cpu) opThumbF8(op uint16) {
	// bl step 2
	newpc := cpu.Regs[14] + reg((op&0x7FF)<<1)
	cpu.Regs[14] = (cpu.Regs[15] - 2) | 1
	cpu.branch(newpc, BranchCall)
}

func (cpu *Cpu) opThumbAlu00(op uint16) {
	// ands
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd & rs
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu00(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("ands      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu01(op uint16) {
	// eors
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd ^ rs
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu01(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("eors      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu02(op uint16) {
	// lsls
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	shift := (rs & 0xFF)
	if shift != 0 {
		cpu.Cpsr.SetC((rd<<(shift-1))&0x80000000 != 0)
	}
	res := rd << shift
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu02(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("lsls      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu03(op uint16) {
	// lsrs
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	shift := (rs & 0xFF)
	if shift != 0 {
		cpu.Cpsr.SetC((rd>>(shift-1))&1 != 0)
	}
	res := rd >> shift
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu03(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("lsrs      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu04(op uint16) {
	// asrs
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	shift := (rs & 0xFF)
	if shift != 0 {
		cpu.Cpsr.SetC((int32(rd)>>(shift-1))&1 != 0)
	}
	res := uint32(int32(rd) >> shift)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu04(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("asrs      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu05(op uint16) {
	// adcs
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	cf := cpu.Cpsr.CB()
	res := rd + rs
	res += cf
	if cf == 0 {
		cpu.Cpsr.SetC(res < rd)
	} else {
		cpu.Cpsr.SetC(res <= rd)
	}
	cpu.Cpsr.SetVAdd(rd, rs, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu05(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("adcs      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu06(op uint16) {
	// sbcs
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	cf := cpu.Cpsr.CB()
	res := rd - rs
	res += cf - 1
	if cf == 0 {
		cpu.Cpsr.SetC(rd > rs)
	} else {
		cpu.Cpsr.SetC(rd >= rs)
	}
	cpu.Cpsr.SetVSub(rd, rs, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu06(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("sbcs      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu07(op uint16) {
	// rors
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	rot := (rs & 0xFF)
	if rot != 0 {
		cpu.Cpsr.SetC((rd<<(rot-1))&0x80000000 != 0)
	}
	rot = (rs & 0x1F)
	res := (rd >> rot) | (rd << (32 - rot))
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu07(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("rors      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu08(op uint16) {
	// tst
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd & rs
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) disasmThumbAlu08(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("tst       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu09(op uint16) {
	// negs
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	res := 0 - rs
	cpu.Cpsr.SetC(false)
	cpu.Cpsr.SetVSub(0, rs, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu09(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("negs      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu0A(op uint16) {
	// cmp
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd - rs
	cpu.Cpsr.SetC(rd >= rs)
	cpu.Cpsr.SetVSub(rd, rs, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) disasmThumbAlu0A(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("cmp       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu0B(op uint16) {
	// cmn
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd + rs
	cpu.Cpsr.SetC(res < rd)
	cpu.Cpsr.SetVAdd(rd, rs, res)
	cpu.Cpsr.SetNZ(res)
}

func (cpu *Cpu) disasmThumbAlu0B(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("cmn       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu0C(op uint16) {
	// orrs
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd | rs
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu0C(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("orrs      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu0D(op uint16) {
	// muls
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd * rs
	if cpu.arch <= ARMv4 {
		cpu.Cpsr.SetC(false)
	}
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu0D(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("muls      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu0E(op uint16) {
	// bics
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	rd := uint32(cpu.Regs[rdx])
	res := rd &^ rs
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu0E(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("bics      ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opThumbAlu0F(op uint16) {
	// mvn
	rsx := (op >> 3) & 0x7
	rs := uint32(cpu.Regs[rsx])
	rdx := op & 0x7
	res := ^rs
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmThumbAlu0F(op uint16, pc uint32) string {
	var out bytes.Buffer
	out.WriteString("mvn       ")
	arg0 := op & 7
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 3) & 7
	out.WriteString(RegNames[arg1])
	return out.String()
}

var opThumbTable = [256]func(*Cpu, uint16){
	(*Cpu).opThumb00, (*Cpu).opThumb00, (*Cpu).opThumb00, (*Cpu).opThumb00,
	(*Cpu).opThumb00, (*Cpu).opThumb00, (*Cpu).opThumb00, (*Cpu).opThumb00,
	(*Cpu).opThumb08, (*Cpu).opThumb08, (*Cpu).opThumb08, (*Cpu).opThumb08,
	(*Cpu).opThumb08, (*Cpu).opThumb08, (*Cpu).opThumb08, (*Cpu).opThumb08,
	(*Cpu).opThumb10, (*Cpu).opThumb10, (*Cpu).opThumb10, (*Cpu).opThumb10,
	(*Cpu).opThumb10, (*Cpu).opThumb10, (*Cpu).opThumb10, (*Cpu).opThumb10,
	(*Cpu).opThumb18, (*Cpu).opThumb18, (*Cpu).opThumb1A, (*Cpu).opThumb1A,
	(*Cpu).opThumb1C, (*Cpu).opThumb1C, (*Cpu).opThumb1E, (*Cpu).opThumb1E,
	(*Cpu).opThumb20, (*Cpu).opThumb21, (*Cpu).opThumb22, (*Cpu).opThumb23,
	(*Cpu).opThumb24, (*Cpu).opThumb25, (*Cpu).opThumb26, (*Cpu).opThumb27,
	(*Cpu).opThumb28, (*Cpu).opThumb29, (*Cpu).opThumb2A, (*Cpu).opThumb2B,
	(*Cpu).opThumb2C, (*Cpu).opThumb2D, (*Cpu).opThumb2E, (*Cpu).opThumb2F,
	(*Cpu).opThumb30, (*Cpu).opThumb31, (*Cpu).opThumb32, (*Cpu).opThumb33,
	(*Cpu).opThumb34, (*Cpu).opThumb35, (*Cpu).opThumb36, (*Cpu).opThumb37,
	(*Cpu).opThumb38, (*Cpu).opThumb39, (*Cpu).opThumb3A, (*Cpu).opThumb3B,
	(*Cpu).opThumb3C, (*Cpu).opThumb3D, (*Cpu).opThumb3E, (*Cpu).opThumb3F,
	(*Cpu).opThumb40, (*Cpu).opThumb40, (*Cpu).opThumb40, (*Cpu).opThumb40,
	(*Cpu).opThumb44, (*Cpu).opThumb45, (*Cpu).opThumb46, (*Cpu).opThumb47,
	(*Cpu).opThumb48, (*Cpu).opThumb49, (*Cpu).opThumb4A, (*Cpu).opThumb4B,
	(*Cpu).opThumb4C, (*Cpu).opThumb4D, (*Cpu).opThumb4E, (*Cpu).opThumb4F,
	(*Cpu).opThumb50, (*Cpu).opThumb50, (*Cpu).opThumb52, (*Cpu).opThumb52,
	(*Cpu).opThumb54, (*Cpu).opThumb54, (*Cpu).opThumb56, (*Cpu).opThumb56,
	(*Cpu).opThumb58, (*Cpu).opThumb58, (*Cpu).opThumb5A, (*Cpu).opThumb5A,
	(*Cpu).opThumb5C, (*Cpu).opThumb5C, (*Cpu).opThumb5E, (*Cpu).opThumb5E,
	(*Cpu).opThumb60, (*Cpu).opThumb60, (*Cpu).opThumb60, (*Cpu).opThumb60,
	(*Cpu).opThumb60, (*Cpu).opThumb60, (*Cpu).opThumb60, (*Cpu).opThumb60,
	(*Cpu).opThumb68, (*Cpu).opThumb68, (*Cpu).opThumb68, (*Cpu).opThumb68,
	(*Cpu).opThumb68, (*Cpu).opThumb68, (*Cpu).opThumb68, (*Cpu).opThumb68,
	(*Cpu).opThumb70, (*Cpu).opThumb70, (*Cpu).opThumb70, (*Cpu).opThumb70,
	(*Cpu).opThumb70, (*Cpu).opThumb70, (*Cpu).opThumb70, (*Cpu).opThumb70,
	(*Cpu).opThumb78, (*Cpu).opThumb78, (*Cpu).opThumb78, (*Cpu).opThumb78,
	(*Cpu).opThumb78, (*Cpu).opThumb78, (*Cpu).opThumb78, (*Cpu).opThumb78,
	(*Cpu).opThumb80, (*Cpu).opThumb80, (*Cpu).opThumb80, (*Cpu).opThumb80,
	(*Cpu).opThumb80, (*Cpu).opThumb80, (*Cpu).opThumb80, (*Cpu).opThumb80,
	(*Cpu).opThumb88, (*Cpu).opThumb88, (*Cpu).opThumb88, (*Cpu).opThumb88,
	(*Cpu).opThumb88, (*Cpu).opThumb88, (*Cpu).opThumb88, (*Cpu).opThumb88,
	(*Cpu).opThumb90, (*Cpu).opThumb91, (*Cpu).opThumb92, (*Cpu).opThumb93,
	(*Cpu).opThumb94, (*Cpu).opThumb95, (*Cpu).opThumb96, (*Cpu).opThumb97,
	(*Cpu).opThumb98, (*Cpu).opThumb99, (*Cpu).opThumb9A, (*Cpu).opThumb9B,
	(*Cpu).opThumb9C, (*Cpu).opThumb9D, (*Cpu).opThumb9E, (*Cpu).opThumb9F,
	(*Cpu).opThumbA0, (*Cpu).opThumbA1, (*Cpu).opThumbA2, (*Cpu).opThumbA3,
	(*Cpu).opThumbA4, (*Cpu).opThumbA5, (*Cpu).opThumbA6, (*Cpu).opThumbA7,
	(*Cpu).opThumbA8, (*Cpu).opThumbA9, (*Cpu).opThumbAA, (*Cpu).opThumbAB,
	(*Cpu).opThumbAC, (*Cpu).opThumbAD, (*Cpu).opThumbAE, (*Cpu).opThumbAF,
	(*Cpu).opThumbB0, (*Cpu).opThumbB1, (*Cpu).opThumbB1, (*Cpu).opThumbB1,
	(*Cpu).opThumbB4, (*Cpu).opThumbB4, (*Cpu).opThumbB1, (*Cpu).opThumbB1,
	(*Cpu).opThumbB1, (*Cpu).opThumbB1, (*Cpu).opThumbB1, (*Cpu).opThumbB1,
	(*Cpu).opThumbBC, (*Cpu).opThumbBC, (*Cpu).opThumbB1, (*Cpu).opThumbB1,
	(*Cpu).opThumbC0, (*Cpu).opThumbC1, (*Cpu).opThumbC2, (*Cpu).opThumbC3,
	(*Cpu).opThumbC4, (*Cpu).opThumbC5, (*Cpu).opThumbC6, (*Cpu).opThumbC7,
	(*Cpu).opThumbC8, (*Cpu).opThumbC9, (*Cpu).opThumbCA, (*Cpu).opThumbCB,
	(*Cpu).opThumbCC, (*Cpu).opThumbCD, (*Cpu).opThumbCE, (*Cpu).opThumbCF,
	(*Cpu).opThumbD0, (*Cpu).opThumbD1, (*Cpu).opThumbD2, (*Cpu).opThumbD3,
	(*Cpu).opThumbD4, (*Cpu).opThumbD5, (*Cpu).opThumbD6, (*Cpu).opThumbD7,
	(*Cpu).opThumbD8, (*Cpu).opThumbD9, (*Cpu).opThumbDA, (*Cpu).opThumbDB,
	(*Cpu).opThumbDC, (*Cpu).opThumbDD, (*Cpu).opThumbDE, (*Cpu).opThumbDF,
	(*Cpu).opThumbE0, (*Cpu).opThumbE0, (*Cpu).opThumbE0, (*Cpu).opThumbE0,
	(*Cpu).opThumbE0, (*Cpu).opThumbE0, (*Cpu).opThumbE0, (*Cpu).opThumbE0,
	(*Cpu).opThumbE8, (*Cpu).opThumbE8, (*Cpu).opThumbE8, (*Cpu).opThumbE8,
	(*Cpu).opThumbE8, (*Cpu).opThumbE8, (*Cpu).opThumbE8, (*Cpu).opThumbE8,
	(*Cpu).opThumbF0, (*Cpu).opThumbF0, (*Cpu).opThumbF0, (*Cpu).opThumbF0,
	(*Cpu).opThumbF0, (*Cpu).opThumbF0, (*Cpu).opThumbF0, (*Cpu).opThumbF0,
	(*Cpu).opThumbF8, (*Cpu).opThumbF8, (*Cpu).opThumbF8, (*Cpu).opThumbF8,
	(*Cpu).opThumbF8, (*Cpu).opThumbF8, (*Cpu).opThumbF8, (*Cpu).opThumbF8,
}
var disasmThumbTable = [256]func(*Cpu, uint16, uint32) string{
	(*Cpu).disasmThumb00, (*Cpu).disasmThumb00, (*Cpu).disasmThumb00, (*Cpu).disasmThumb00,
	(*Cpu).disasmThumb00, (*Cpu).disasmThumb00, (*Cpu).disasmThumb00, (*Cpu).disasmThumb00,
	(*Cpu).disasmThumb08, (*Cpu).disasmThumb08, (*Cpu).disasmThumb08, (*Cpu).disasmThumb08,
	(*Cpu).disasmThumb08, (*Cpu).disasmThumb08, (*Cpu).disasmThumb08, (*Cpu).disasmThumb08,
	(*Cpu).disasmThumb10, (*Cpu).disasmThumb10, (*Cpu).disasmThumb10, (*Cpu).disasmThumb10,
	(*Cpu).disasmThumb10, (*Cpu).disasmThumb10, (*Cpu).disasmThumb10, (*Cpu).disasmThumb10,
	(*Cpu).disasmThumb18, (*Cpu).disasmThumb18, (*Cpu).disasmThumb1A, (*Cpu).disasmThumb1A,
	(*Cpu).disasmThumb1C, (*Cpu).disasmThumb1C, (*Cpu).disasmThumb1E, (*Cpu).disasmThumb1E,
	(*Cpu).disasmThumb20, (*Cpu).disasmThumb20, (*Cpu).disasmThumb20, (*Cpu).disasmThumb20,
	(*Cpu).disasmThumb20, (*Cpu).disasmThumb20, (*Cpu).disasmThumb20, (*Cpu).disasmThumb20,
	(*Cpu).disasmThumb28, (*Cpu).disasmThumb28, (*Cpu).disasmThumb28, (*Cpu).disasmThumb28,
	(*Cpu).disasmThumb28, (*Cpu).disasmThumb28, (*Cpu).disasmThumb28, (*Cpu).disasmThumb28,
	(*Cpu).disasmThumb30, (*Cpu).disasmThumb30, (*Cpu).disasmThumb30, (*Cpu).disasmThumb30,
	(*Cpu).disasmThumb30, (*Cpu).disasmThumb30, (*Cpu).disasmThumb30, (*Cpu).disasmThumb30,
	(*Cpu).disasmThumb38, (*Cpu).disasmThumb38, (*Cpu).disasmThumb38, (*Cpu).disasmThumb38,
	(*Cpu).disasmThumb38, (*Cpu).disasmThumb38, (*Cpu).disasmThumb38, (*Cpu).disasmThumb38,
	(*Cpu).disasmThumb40, (*Cpu).disasmThumb40, (*Cpu).disasmThumb40, (*Cpu).disasmThumb40,
	(*Cpu).disasmThumb44, (*Cpu).disasmThumb45, (*Cpu).disasmThumb46, (*Cpu).disasmThumb47,
	(*Cpu).disasmThumb48, (*Cpu).disasmThumb48, (*Cpu).disasmThumb48, (*Cpu).disasmThumb48,
	(*Cpu).disasmThumb48, (*Cpu).disasmThumb48, (*Cpu).disasmThumb48, (*Cpu).disasmThumb48,
	(*Cpu).disasmThumb50, (*Cpu).disasmThumb50, (*Cpu).disasmThumb52, (*Cpu).disasmThumb52,
	(*Cpu).disasmThumb54, (*Cpu).disasmThumb54, (*Cpu).disasmThumb56, (*Cpu).disasmThumb56,
	(*Cpu).disasmThumb58, (*Cpu).disasmThumb58, (*Cpu).disasmThumb5A, (*Cpu).disasmThumb5A,
	(*Cpu).disasmThumb5C, (*Cpu).disasmThumb5C, (*Cpu).disasmThumb5E, (*Cpu).disasmThumb5E,
	(*Cpu).disasmThumb60, (*Cpu).disasmThumb60, (*Cpu).disasmThumb60, (*Cpu).disasmThumb60,
	(*Cpu).disasmThumb60, (*Cpu).disasmThumb60, (*Cpu).disasmThumb60, (*Cpu).disasmThumb60,
	(*Cpu).disasmThumb68, (*Cpu).disasmThumb68, (*Cpu).disasmThumb68, (*Cpu).disasmThumb68,
	(*Cpu).disasmThumb68, (*Cpu).disasmThumb68, (*Cpu).disasmThumb68, (*Cpu).disasmThumb68,
	(*Cpu).disasmThumb70, (*Cpu).disasmThumb70, (*Cpu).disasmThumb70, (*Cpu).disasmThumb70,
	(*Cpu).disasmThumb70, (*Cpu).disasmThumb70, (*Cpu).disasmThumb70, (*Cpu).disasmThumb70,
	(*Cpu).disasmThumb78, (*Cpu).disasmThumb78, (*Cpu).disasmThumb78, (*Cpu).disasmThumb78,
	(*Cpu).disasmThumb78, (*Cpu).disasmThumb78, (*Cpu).disasmThumb78, (*Cpu).disasmThumb78,
	(*Cpu).disasmThumb80, (*Cpu).disasmThumb80, (*Cpu).disasmThumb80, (*Cpu).disasmThumb80,
	(*Cpu).disasmThumb80, (*Cpu).disasmThumb80, (*Cpu).disasmThumb80, (*Cpu).disasmThumb80,
	(*Cpu).disasmThumb88, (*Cpu).disasmThumb88, (*Cpu).disasmThumb88, (*Cpu).disasmThumb88,
	(*Cpu).disasmThumb88, (*Cpu).disasmThumb88, (*Cpu).disasmThumb88, (*Cpu).disasmThumb88,
	(*Cpu).disasmThumb90, (*Cpu).disasmThumb90, (*Cpu).disasmThumb90, (*Cpu).disasmThumb90,
	(*Cpu).disasmThumb90, (*Cpu).disasmThumb90, (*Cpu).disasmThumb90, (*Cpu).disasmThumb90,
	(*Cpu).disasmThumb98, (*Cpu).disasmThumb98, (*Cpu).disasmThumb98, (*Cpu).disasmThumb98,
	(*Cpu).disasmThumb98, (*Cpu).disasmThumb98, (*Cpu).disasmThumb98, (*Cpu).disasmThumb98,
	(*Cpu).disasmThumbA0, (*Cpu).disasmThumbA0, (*Cpu).disasmThumbA0, (*Cpu).disasmThumbA0,
	(*Cpu).disasmThumbA0, (*Cpu).disasmThumbA0, (*Cpu).disasmThumbA0, (*Cpu).disasmThumbA0,
	(*Cpu).disasmThumbA8, (*Cpu).disasmThumbA8, (*Cpu).disasmThumbA8, (*Cpu).disasmThumbA8,
	(*Cpu).disasmThumbA8, (*Cpu).disasmThumbA8, (*Cpu).disasmThumbA8, (*Cpu).disasmThumbA8,
	(*Cpu).disasmThumbB0, (*Cpu).disasmThumbB1, (*Cpu).disasmThumbB1, (*Cpu).disasmThumbB1,
	(*Cpu).disasmThumbB4, (*Cpu).disasmThumbB5, (*Cpu).disasmThumbB1, (*Cpu).disasmThumbB1,
	(*Cpu).disasmThumbB1, (*Cpu).disasmThumbB1, (*Cpu).disasmThumbB1, (*Cpu).disasmThumbB1,
	(*Cpu).disasmThumbBC, (*Cpu).disasmThumbBD, (*Cpu).disasmThumbB1, (*Cpu).disasmThumbB1,
	(*Cpu).disasmThumbC0, (*Cpu).disasmThumbC0, (*Cpu).disasmThumbC0, (*Cpu).disasmThumbC0,
	(*Cpu).disasmThumbC0, (*Cpu).disasmThumbC0, (*Cpu).disasmThumbC0, (*Cpu).disasmThumbC0,
	(*Cpu).disasmThumbC8, (*Cpu).disasmThumbC8, (*Cpu).disasmThumbC8, (*Cpu).disasmThumbC8,
	(*Cpu).disasmThumbC8, (*Cpu).disasmThumbC8, (*Cpu).disasmThumbC8, (*Cpu).disasmThumbC8,
	(*Cpu).disasmThumbD0, (*Cpu).disasmThumbD1, (*Cpu).disasmThumbD2, (*Cpu).disasmThumbD3,
	(*Cpu).disasmThumbD4, (*Cpu).disasmThumbD5, (*Cpu).disasmThumbD6, (*Cpu).disasmThumbD7,
	(*Cpu).disasmThumbD8, (*Cpu).disasmThumbD9, (*Cpu).disasmThumbDA, (*Cpu).disasmThumbDB,
	(*Cpu).disasmThumbDC, (*Cpu).disasmThumbDD, (*Cpu).disasmThumbB1, (*Cpu).disasmThumbDF,
	(*Cpu).disasmThumbE0, (*Cpu).disasmThumbE0, (*Cpu).disasmThumbE0, (*Cpu).disasmThumbE0,
	(*Cpu).disasmThumbE0, (*Cpu).disasmThumbE0, (*Cpu).disasmThumbE0, (*Cpu).disasmThumbE0,
	(*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8,
	(*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8,
	(*Cpu).disasmThumbF0, (*Cpu).disasmThumbF0, (*Cpu).disasmThumbF0, (*Cpu).disasmThumbF0,
	(*Cpu).disasmThumbF0, (*Cpu).disasmThumbF0, (*Cpu).disasmThumbF0, (*Cpu).disasmThumbF0,
	(*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8,
	(*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8, (*Cpu).disasmThumbE8,
}
