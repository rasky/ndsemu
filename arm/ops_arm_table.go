// Generated on 2015-12-21 13:26:56.531478137 +0100 CET
package arm

import "bytes"
import "strconv"

func (cpu *Cpu) opArm000(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm000(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("and", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm009(op uint32) {
	// mul
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res := rm * rs
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm009(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mul", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	return out.String()
}

func (cpu *Cpu) opArm010(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm010(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ands", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm019(op uint32) {
	// muls
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res := rm * rs
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm019(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("muls", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	return out.String()
}

func (cpu *Cpu) opArm020(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm020(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("eor", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm029(op uint32) {
	// mla
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	rnx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	res := rm*rs + rn
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm029(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mla", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm030(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm030(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("eors", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm039(op uint32) {
	// mlas
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	rnx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	res := rm*rs + rn
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm039(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mlas", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm040(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm040(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("sub", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm049(op uint32) {
	cpu.InvalidOpArm(op, "invalid opcode decoded as LD/STR half-word")
}

func (cpu *Cpu) disasmArm049(op uint32, pc uint32) string {
	return "dw " + strconv.FormatInt(int64(op), 16)
}

func (cpu *Cpu) opArm04B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm04B(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strh", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm04D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm04D(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrd", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm04F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm04F(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strd", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm050(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm050(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("subs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm05B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm05B(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrh", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm05D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm05D(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrsb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm05F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm05F(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrsh", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm060(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm060(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rsb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm070(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm070(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rsbs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm080(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm080(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("add", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm089(op uint32) {
	// umull
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := uint64(rm) * uint64(rs)
	rnx := (op >> 12) & 0xF
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm089(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("umull", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm090(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm090(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("adds", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm099(op uint32) {
	// umulls
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := uint64(rm) * uint64(rs)
	rnx := (op >> 12) & 0xF
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Cpsr.SetNZ64(res64)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm099(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("umulls", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0A0(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm0A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("adc", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm0A9(op uint32) {
	// umlal
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := uint64(rm) * uint64(rs)
	rnx := (op >> 12) & 0xF
	app64 := uint64(cpu.Regs[rnx]) + uint64(cpu.Regs[rdx])<<32
	res64 += app64
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm0A9(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("umlal", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0B0(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	res += cf
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm0B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("adcs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm0B9(op uint32) {
	// umlals
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := uint64(rm) * uint64(rs)
	rnx := (op >> 12) & 0xF
	app64 := uint64(cpu.Regs[rnx]) + uint64(cpu.Regs[rdx])<<32
	res64 += app64
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Cpsr.SetNZ64(res64)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm0B9(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("umlals", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0C0(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm0C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("sbc", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm0C9(op uint32) {
	// smull
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := int64(int32(rm)) * int64(int32(rs))
	rnx := (op >> 12) & 0xF
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm0C9(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("smull", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0D0(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	res += cf - 1
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm0D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("sbcs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm0D9(op uint32) {
	// smulls
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := int64(int32(rm)) * int64(int32(rs))
	rnx := (op >> 12) & 0xF
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Cpsr.SetNZ64(uint64(res64))
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm0D9(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("smulls", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0E0(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm0E0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rsc", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm0E9(op uint32) {
	// smlal
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := int64(int32(rm)) * int64(int32(rs))
	rnx := (op >> 12) & 0xF
	app64 := uint64(cpu.Regs[rnx]) + uint64(cpu.Regs[rdx])<<32
	res64 += int64(app64)
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm0E9(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("smlal", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0F0(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	res += cf - 1
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm0F0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rscs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm0F9(op uint32) {
	// smlals
	if !cpu.opArmCond(op) {
		return
	}
	rsx := (op >> 8) & 0xF
	rs := uint32(cpu.Regs[rsx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	data := (op >> 4) & 0xF
	if data != 0x9 {
		cpu.InvalidOpArm(op, "unimplemented half-word multiplies")
		return
	}
	rdx := (op >> 16) & 0xF
	res64 := int64(int32(rm)) * int64(int32(rs))
	rnx := (op >> 12) & 0xF
	app64 := uint64(cpu.Regs[rnx]) + uint64(cpu.Regs[rdx])<<32
	res64 += int64(app64)
	cpu.Regs[rnx] = reg(res64)
	res := uint32(res64 >> 32)
	cpu.Cpsr.SetNZ64(uint64(res64))
	cpu.Regs[rdx] = reg(res)
}

func (cpu *Cpu) disasmArm0F9(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("smlals", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm100(op uint32) {
	if op&0x0F900FF0 != 0x01000000 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as PSR_reg")
		return
	}
	// MRS
	if !cpu.opArmCond(op) {
		return
	}
	mask := (op >> 16) & 0xF
	if mask != 0xF {
		cpu.InvalidOpArm(op, "mask should be 0xF in MRS (is it SWP?)")
		return
	}
	rdx := (op >> 12) & 0xF
	if rdx == 15 {
		cpu.InvalidOpArm(op, "write to PC in MRS")
		return
	}
	cpu.Regs[rdx] = reg(cpu.Cpsr.Uint32())
}

func (cpu *Cpu) opArm101(op uint32) {
	cpu.InvalidOpArm(op, "invalid ALU test function without flags")
}

func (cpu *Cpu) opArm109(op uint32) {
	if op&0x0FB00FF0 != 0x01000090 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as SWP")
		return
	}
	// swp
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rn := uint32(cpu.Regs[rnx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	rdx := (op >> 12) & 0xF
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.opWrite32(rn, rm)
}

func (cpu *Cpu) disasmArm109(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("swp", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg2])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm110(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm110(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("tst", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm11B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
}

func (cpu *Cpu) opArm11D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm11F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm120(op uint32) {
	if op&0x0F900FF0 != 0x01000000 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as PSR_reg")
		return
	}
	// MSR
	if !cpu.opArmCond(op) {
		return
	}
	mask := uint32(0)
	if (op>>19)&1 != 0 {
		mask |= 0xFF000000
	}
	if (op>>18)&1 != 0 {
		mask |= 0x00FF0000
	}
	if (op>>17)&1 != 0 {
		mask |= 0x0000FF00
	}
	if (op>>16)&1 != 0 {
		mask |= 0x000000FF
	}
	rmx := op & 0xF
	val := uint32(cpu.Regs[rmx])
	cpu.Cpsr.SetWithMask(val, mask, cpu)
}

func (cpu *Cpu) opArm121(op uint32) {
	// bx reg
	if op&0x0FFFFFD0 != 0x012FFF10 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as BX/BLX")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := op & 0xF
	rn := cpu.Regs[rnx]
	if rn&1 != 0 {
		cpu.Cpsr.SetT(true)
		rn &^= 1
	} else {
		rn &^= 3
	}
	cpu.pc = rn
}

func (cpu *Cpu) disasmArm121(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("bx", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := op & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm123(op uint32) {
	// blx reg
	if op&0x0FFFFFD0 != 0x012FFF10 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as BX/BLX")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := op & 0xF
	rn := cpu.Regs[rnx]
	cpu.Regs[14] = cpu.Regs[15] - 4
	if rn&1 != 0 {
		cpu.Cpsr.SetT(true)
		rn &^= 1
	} else {
		rn &^= 3
	}
	cpu.pc = rn
}

func (cpu *Cpu) disasmArm123(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("blx", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := op & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm12B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm12D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm12F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm130(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm130(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("teq", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm13B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm13D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm13F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn -= off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm140(op uint32) {
	if op&0x0F900FF0 != 0x01000000 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as PSR_reg")
		return
	}
	// MRS
	if !cpu.opArmCond(op) {
		return
	}
	mask := (op >> 16) & 0xF
	if mask != 0xF {
		cpu.InvalidOpArm(op, "mask should be 0xF in MRS (is it SWP?)")
		return
	}
	rdx := (op >> 12) & 0xF
	if rdx == 15 {
		cpu.InvalidOpArm(op, "write to PC in MRS")
		return
	}
	cpu.Regs[rdx] = reg(*cpu.RegSpsr())
}

func (cpu *Cpu) opArm149(op uint32) {
	if op&0x0FB00FF0 != 0x01000090 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as SWP")
		return
	}
	// swpb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rn := uint32(cpu.Regs[rnx])
	rmx := (op >> 0) & 0xF
	rm := uint32(cpu.Regs[rmx])
	rdx := (op >> 12) & 0xF
	cpu.Regs[rdx] = reg(cpu.opRead8(rn))
	cpu.opWrite8(rn, uint8(rm))
}

func (cpu *Cpu) disasmArm149(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("swpb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg2])
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm150(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm150(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("cmp", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm15B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
}

func (cpu *Cpu) opArm15D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm15F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm160(op uint32) {
	if op&0x0F900FF0 != 0x01000000 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as PSR_reg")
		return
	}
	// MSR
	if !cpu.opArmCond(op) {
		return
	}
	mask := uint32(0)
	if (op>>19)&1 != 0 {
		mask |= 0xFF000000
	}
	if (op>>18)&1 != 0 {
		mask |= 0x00FF0000
	}
	if (op>>17)&1 != 0 {
		mask |= 0x0000FF00
	}
	if (op>>16)&1 != 0 {
		mask |= 0x000000FF
	}
	rmx := op & 0xF
	val := uint32(cpu.Regs[rmx])
	cpu.RegSpsr().SetWithMask(val, mask)
}

func (cpu *Cpu) opArm161(op uint32) {
	// clz
	if op&0x0FFF0FF0 != 0x016F0F10 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as CLZ")
		return
	}
	if cpu.arch < ARMv5 {
		cpu.InvalidOpArm(op, "invalid CLZ opcode on pre-ARMv5 CPU")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rdx := (op >> 12) & 0xF
	rm := cpu.Regs[op&0xF]
	var lz int
	for lz = 0; lz < 32; lz++ {
		if int32(rm) < 0 {
			break
		}
		rm <<= 1
	}
	cpu.Regs[rdx] = reg(lz)
}

func (cpu *Cpu) disasmArm161(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("clz", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := op & 0xF
	out.WriteString(RegNames[arg1])
	return out.String()
}

func (cpu *Cpu) opArm16B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm16D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm16F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm170(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm170(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("cmn", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm17B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm17D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm17F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn -= off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm180(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm180(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("orr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm18B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
}

func (cpu *Cpu) opArm18D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
}

func (cpu *Cpu) opArm18F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
}

func (cpu *Cpu) opArm190(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm190(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("orrs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm19B(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
}

func (cpu *Cpu) opArm19D(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm19F(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm1A0(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm1A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mov", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm1AB(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1AD(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1AF(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1B0(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm1B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("movs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm1BB(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1BD(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1BF(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	rmx := op & 0xF
	if rmx == 15 {
		cpu.InvalidOpArm(op, "halfword: invalid rm==15")
		return
	}
	off := uint32(cpu.Regs[rmx])
	rn += off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1C0(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm1C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("bic", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm1CB(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
}

func (cpu *Cpu) opArm1CD(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
}

func (cpu *Cpu) opArm1CF(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
}

func (cpu *Cpu) opArm1D0(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm1D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("bics", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm1DB(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
}

func (cpu *Cpu) opArm1DD(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm1DF(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
}

func (cpu *Cpu) opArm1E0(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm1E0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mvn", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm1EB(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1ED(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1EF(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1F0(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	op2 := cpu.opDecodeAluOp2Reg(op, true)
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm1F0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mvns", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmOp2(op)
	out.WriteString(arg1)
	return out.String()
}

func (cpu *Cpu) opArm1FB(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1FD(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm1FF(op uint32) {
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := (op & 0xF) | ((op & 0xF00) >> 4)
	rn += off
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm200(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm200(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("and", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm210(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm210(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ands", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm220(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm220(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("eor", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm230(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm230(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("eors", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm240(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm240(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("sub", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm250(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm250(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("subs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm260(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm260(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rsb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm270(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm270(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rsbs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm280(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm280(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("add", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm290(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm290(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("adds", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm2A0(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm2A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("adc", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm2B0(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	res += cf
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm2B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("adcs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm2C0(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm2C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("sbc", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm2D0(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	res += cf - 1
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm2D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("sbcs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm2E0(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm2E0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rsc", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm2F0(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	cf := cpu.Cpsr.CB()
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	res += cf - 1
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm2F0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("rscs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm310(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm310(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("tst", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm320(op uint32) {
	if op&0x0FB00000 != 0x03200000 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as PSR_imm")
		return
	}
	// MSR
	if !cpu.opArmCond(op) {
		return
	}
	mask := uint32(0)
	if (op>>19)&1 != 0 {
		mask |= 0xFF000000
	}
	if (op>>18)&1 != 0 {
		mask |= 0x00FF0000
	}
	if (op>>17)&1 != 0 {
		mask |= 0x0000FF00
	}
	if (op>>16)&1 != 0 {
		mask |= 0x000000FF
	}
	val := op & 0xFF
	shcnt := uint(((op >> 8) & 0xF) * 2)
	val = (val >> shcnt) | (val << (32 - shcnt))
	cpu.Cpsr.SetWithMask(val, mask, cpu)
}

func (cpu *Cpu) opArm330(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm330(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("teq", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm350(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm350(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("cmp", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm360(op uint32) {
	if op&0x0FB00000 != 0x03200000 {
		cpu.InvalidOpArm(op, "invalid opcode decoded as PSR_imm")
		return
	}
	// MSR
	if !cpu.opArmCond(op) {
		return
	}
	mask := uint32(0)
	if (op>>19)&1 != 0 {
		mask |= 0xFF000000
	}
	if (op>>18)&1 != 0 {
		mask |= 0x00FF0000
	}
	if (op>>17)&1 != 0 {
		mask |= 0x0000FF00
	}
	if (op>>16)&1 != 0 {
		mask |= 0x000000FF
	}
	val := op & 0xFF
	shcnt := uint(((op >> 8) & 0xF) * 2)
	val = (val >> shcnt) | (val << (32 - shcnt))
	cpu.RegSpsr().SetWithMask(val, mask)
}

func (cpu *Cpu) opArm370(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm370(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("cmn", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm380(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm380(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("orr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm390(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm390(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("orrs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm3A0(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm3A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mov", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm3B0(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm3B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("movs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm3C0(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm3C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("bic", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm3D0(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm3D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("bics", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm3E0(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm3E0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mvn", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm3F0(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) disasmArm3F0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mvns", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> ((op >> 7) & 0x1E)) | ((op & 0xFF) << (32 - ((op >> 7) & 0x1E))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm400(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm400(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(-int64(op & 0xFFF))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm410(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm410(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(-int64(op & 0xFFF))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm420(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm430(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm440(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm440(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(-int64(op & 0xFFF))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm450(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm450(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(-int64(op & 0xFFF))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm460(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm470(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm480(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm480(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(op & 0xFFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm490(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm490(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(op & 0xFFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm4A0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm4B0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm4C0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm4C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(op & 0xFFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm4D0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm4D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := int64(op & 0xFFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg2, 16))
	return out.String()
}

func (cpu *Cpu) opArm4E0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm4F0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm500(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
}

func (cpu *Cpu) disasmArm500(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm510(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm510(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm520(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm520(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm530(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm530(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm540(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
}

func (cpu *Cpu) disasmArm540(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm550(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm550(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm560(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm560(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm570(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm570(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := -int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm580(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
}

func (cpu *Cpu) disasmArm580(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm590(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm590(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm5A0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm5A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm5B0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm5B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm5C0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
}

func (cpu *Cpu) disasmArm5C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm5D0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm5D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !false {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
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

func (cpu *Cpu) opArm5E0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm5E0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm5F0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := op & 0xFFF
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm5F0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := int32(op & 0xFFF)
	if RegNames[arg1a] == "pc" && !true {
		arg1c := uint32(arg1b) + uint32((pc+8)&^2)
		arg1v := cpu.opRead32(arg1c)
		out.WriteString("= 0x")
		out.WriteString(strconv.FormatInt(int64(arg1v), 16))
	} else {
		out.WriteString("[")
		out.WriteString(RegNames[arg1a])
		out.WriteString(", #0x")
		out.WriteString(strconv.FormatInt(int64(arg1b), 16))
		out.WriteString("]")
		out.WriteString("!")
	}
	return out.String()
}

func (cpu *Cpu) opArm600(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm600(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := "-" + cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm610(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm610(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := "-" + cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm620(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm630(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm640(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm640(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := "-" + cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm650(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm650(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := "-" + cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm660(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm670(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm680(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm680(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm690(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm690(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm6A0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm6B0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm6C0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm6C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm6D0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm6D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString("[")
	out.WriteString(RegNames[arg1])
	out.WriteString("]")
	out.WriteString(", ")
	arg2 := cpu.disasmOp2(op)
	out.WriteString(arg2)
	return out.String()
}

func (cpu *Cpu) opArm6E0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm6F0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm700(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
}

func (cpu *Cpu) disasmArm700(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm710(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm710(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm720(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm720(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm730(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm730(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm740(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
}

func (cpu *Cpu) disasmArm740(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm750(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm750(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm760(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm760(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm770(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm770(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := "-" + cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm780(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
}

func (cpu *Cpu) disasmArm780(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm790(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm790(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm7A0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm7A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("str", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm7B0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm7B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm7C0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
}

func (cpu *Cpu) disasmArm7C0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm7D0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
}

func (cpu *Cpu) disasmArm7D0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	return out.String()
}

func (cpu *Cpu) opArm7E0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm7E0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm7F0(op uint32) {
	if (op >> 28) == 0xF {
		cpu.InvalidOpArm(op, "PLD not supported")
		return
	}
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rn := uint32(cpu.Regs[rnx])
	cpu.Regs[15] += 4
	off := cpu.opDecodeAluOp2Reg(op, false)
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.pc = reg(res &^ 1)
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm7F0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1a := (op >> 16) & 0xF
	arg1b := cpu.disasmOp2(op)
	out.WriteString("[")
	out.WriteString(RegNames[arg1a])
	out.WriteString(", ")
	out.WriteString(arg1b)
	out.WriteString("]")
	out.WriteString("!")
	return out.String()
}

func (cpu *Cpu) opArm800(op uint32) {
	// stmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm800(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stmda", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm810(op uint32) {
	// ldmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm810(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldmda", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm820(op uint32) {
	// stmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) disasmArm820(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stmda", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm830(op uint32) {
	// ldmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) disasmArm830(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldmda", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm840(op uint32) {
	// stmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm850(op uint32) {
	// ldmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm860(op uint32) {
	// stmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm870(op uint32) {
	// ldmda
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm880(op uint32) {
	// stm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm880(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stm", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm890(op uint32) {
	// ldm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm890(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldm", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm8A0(op uint32) {
	// stm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm8A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stm", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm8B0(op uint32) {
	// ldm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm8B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldm", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm8C0(op uint32) {
	// stm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm8D0(op uint32) {
	// ldm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm8E0(op uint32) {
	// stm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm8F0(op uint32) {
	// ldm
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm900(op uint32) {
	// stmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm900(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stmdb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm910(op uint32) {
	// ldmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm910(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldmdb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm920(op uint32) {
	// stmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) disasmArm920(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stmdb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm930(op uint32) {
	// ldmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) disasmArm930(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldmdb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm940(op uint32) {
	// stmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm950(op uint32) {
	// ldmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm960(op uint32) {
	// stmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm970(op uint32) {
	// ldmdb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	rn -= uint32(4 * popcount16(mask))
	orn := rn
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm980(op uint32) {
	// stmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm980(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stmib", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm990(op uint32) {
	// ldmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) disasmArm990(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldmib", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm9A0(op uint32) {
	// stmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm9A0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("stmib", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm9B0(op uint32) {
	// ldmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) disasmArm9B0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldmib", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString("!")
	out.WriteString(", ")
	arg1 := uint16(op & 0xFFFF)
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

func (cpu *Cpu) opArm9C0(op uint32) {
	// stmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm9D0(op uint32) {
	// ldmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm9E0(op uint32) {
	// stmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	cpu.Regs[15] += 4
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			if i >= 8 && i < 15 {
				val = uint32(cpu.UsrBank[i-8])
			} else {
				val = uint32(cpu.Regs[i])
			}
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm9F0(op uint32) {
	// ldmib
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	if rnx == 15 {
		cpu.InvalidOpArm(op, "invalid use of PC in LDM/STM")
		return
	}
	rn := uint32(cpu.Regs[rnx])
	mask := uint16(op & 0xFFFF)
	usrbnk := (mask & 0x8000) == 0
	if usrbnk {
		cpu.InvalidOpArm(op, "usrbnk not supported")
		return
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			if usrbnk && i >= 8 && i < 15 {
				cpu.UsrBank[i-8] = val
			} else {
				cpu.Regs[i] = val
			}
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArmA00(op uint32) {
	if op>>28 == 0xF {
		// BLX_imm
		off := int32(op<<9) >> 7
		cpu.Regs[14] = cpu.Regs[15] - 4
		cpu.Regs[15] += reg(off)
		cpu.pc = cpu.Regs[15]
		cpu.Cpsr.SetT(true)
		return
	}
	// B
	if !cpu.opArmCond(op) {
		return
	}
	off := int32(op<<9) >> 7
	cpu.Regs[15] += reg(off)
	cpu.pc = cpu.Regs[15]
}

func (cpu *Cpu) disasmArmA00(op uint32, pc uint32) string {
	if op>>28 == 0xF {
		var out bytes.Buffer
		opcode := cpu.disasmAddCond("blx", op)
		out.WriteString((opcode + "                ")[:10])
		arg0 := int32(int32(op<<9) >> 7)
		arg0x := pc + 8 + uint32(arg0)
		out.WriteString(strconv.FormatInt(int64(arg0x), 16))
		return out.String()
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("b", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := int32(int32(op<<9) >> 7)
	arg0x := pc + 8 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opArmB00(op uint32) {
	if op>>28 == 0xF {
		// BLX_imm
		off := int32(op<<9) >> 7
		cpu.Regs[14] = cpu.Regs[15] - 4
		cpu.Regs[15] += reg(off)
		cpu.Regs[15] += 2
		cpu.pc = cpu.Regs[15]
		cpu.Cpsr.SetT(true)
		return
	}
	// BL
	if !cpu.opArmCond(op) {
		return
	}
	off := int32(op<<9) >> 7
	cpu.Regs[14] = cpu.Regs[15] - 4
	cpu.Regs[15] += reg(off)
	cpu.pc = cpu.Regs[15]
}

func (cpu *Cpu) disasmArmB00(op uint32, pc uint32) string {
	if op>>28 == 0xF {
		var out bytes.Buffer
		opcode := cpu.disasmAddCond("blx", op)
		out.WriteString((opcode + "                ")[:10])
		arg0 := int32(int32(op<<9) >> 7)
		arg0x := pc + 8 + uint32(arg0)
		out.WriteString(strconv.FormatInt(int64(arg0x), 16))
		return out.String()
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("bl", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := int32(int32(op<<9) >> 7)
	arg0x := pc + 8 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opArmC00(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmE00(op uint32) {
	// CDP
	if (op >> 28) != 0xF { // MRC2/MCR2
		if !cpu.opArmCond(op) {
			return
		}
	}
	opc := (op >> 21) & 0x7
	cn := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	copnum := (op >> 8) & 0xF
	cp := (op >> 5) & 0x7
	cm := (op >> 0) & 0xF
	cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
}

func (cpu *Cpu) opArmE01(op uint32) {
	// MCR
	if (op >> 28) != 0xF { // MRC2/MCR2
		if !cpu.opArmCond(op) {
			return
		}
	}
	opc := (op >> 21) & 0x7
	cn := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	copnum := (op >> 8) & 0xF
	cp := (op >> 5) & 0x7
	cm := (op >> 0) & 0xF
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) opArmE11(op uint32) {
	// MRC
	if (op >> 28) != 0xF { // MRC2/MCR2
		if !cpu.opArmCond(op) {
			return
		}
	}
	opc := (op >> 21) & 0x7
	cn := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	copnum := (op >> 8) & 0xF
	cp := (op >> 5) & 0x7
	cm := (op >> 0) & 0xF
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) opArmF00(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF00(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("swi", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := int64(op & 0xFFFFFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg0, 16))
	return out.String()
}

var opArmTable = [4096]func(*Cpu, uint32){
	(*Cpu).opArm000, (*Cpu).opArm000, (*Cpu).opArm000, (*Cpu).opArm000,
	(*Cpu).opArm000, (*Cpu).opArm000, (*Cpu).opArm000, (*Cpu).opArm000,
	(*Cpu).opArm000, (*Cpu).opArm009, (*Cpu).opArm000, (*Cpu).opArm009,
	(*Cpu).opArm000, (*Cpu).opArm009, (*Cpu).opArm000, (*Cpu).opArm009,
	(*Cpu).opArm010, (*Cpu).opArm010, (*Cpu).opArm010, (*Cpu).opArm010,
	(*Cpu).opArm010, (*Cpu).opArm010, (*Cpu).opArm010, (*Cpu).opArm010,
	(*Cpu).opArm010, (*Cpu).opArm019, (*Cpu).opArm010, (*Cpu).opArm019,
	(*Cpu).opArm010, (*Cpu).opArm019, (*Cpu).opArm010, (*Cpu).opArm019,
	(*Cpu).opArm020, (*Cpu).opArm020, (*Cpu).opArm020, (*Cpu).opArm020,
	(*Cpu).opArm020, (*Cpu).opArm020, (*Cpu).opArm020, (*Cpu).opArm020,
	(*Cpu).opArm020, (*Cpu).opArm029, (*Cpu).opArm020, (*Cpu).opArm029,
	(*Cpu).opArm020, (*Cpu).opArm029, (*Cpu).opArm020, (*Cpu).opArm029,
	(*Cpu).opArm030, (*Cpu).opArm030, (*Cpu).opArm030, (*Cpu).opArm030,
	(*Cpu).opArm030, (*Cpu).opArm030, (*Cpu).opArm030, (*Cpu).opArm030,
	(*Cpu).opArm030, (*Cpu).opArm039, (*Cpu).opArm030, (*Cpu).opArm039,
	(*Cpu).opArm030, (*Cpu).opArm039, (*Cpu).opArm030, (*Cpu).opArm039,
	(*Cpu).opArm040, (*Cpu).opArm040, (*Cpu).opArm040, (*Cpu).opArm040,
	(*Cpu).opArm040, (*Cpu).opArm040, (*Cpu).opArm040, (*Cpu).opArm040,
	(*Cpu).opArm040, (*Cpu).opArm049, (*Cpu).opArm040, (*Cpu).opArm04B,
	(*Cpu).opArm040, (*Cpu).opArm04D, (*Cpu).opArm040, (*Cpu).opArm04F,
	(*Cpu).opArm050, (*Cpu).opArm050, (*Cpu).opArm050, (*Cpu).opArm050,
	(*Cpu).opArm050, (*Cpu).opArm050, (*Cpu).opArm050, (*Cpu).opArm050,
	(*Cpu).opArm050, (*Cpu).opArm049, (*Cpu).opArm050, (*Cpu).opArm05B,
	(*Cpu).opArm050, (*Cpu).opArm05D, (*Cpu).opArm050, (*Cpu).opArm05F,
	(*Cpu).opArm060, (*Cpu).opArm060, (*Cpu).opArm060, (*Cpu).opArm060,
	(*Cpu).opArm060, (*Cpu).opArm060, (*Cpu).opArm060, (*Cpu).opArm060,
	(*Cpu).opArm060, (*Cpu).opArm049, (*Cpu).opArm060, (*Cpu).opArm04B,
	(*Cpu).opArm060, (*Cpu).opArm04D, (*Cpu).opArm060, (*Cpu).opArm04F,
	(*Cpu).opArm070, (*Cpu).opArm070, (*Cpu).opArm070, (*Cpu).opArm070,
	(*Cpu).opArm070, (*Cpu).opArm070, (*Cpu).opArm070, (*Cpu).opArm070,
	(*Cpu).opArm070, (*Cpu).opArm049, (*Cpu).opArm070, (*Cpu).opArm05B,
	(*Cpu).opArm070, (*Cpu).opArm05D, (*Cpu).opArm070, (*Cpu).opArm05F,
	(*Cpu).opArm080, (*Cpu).opArm080, (*Cpu).opArm080, (*Cpu).opArm080,
	(*Cpu).opArm080, (*Cpu).opArm080, (*Cpu).opArm080, (*Cpu).opArm080,
	(*Cpu).opArm080, (*Cpu).opArm089, (*Cpu).opArm080, (*Cpu).opArm089,
	(*Cpu).opArm080, (*Cpu).opArm089, (*Cpu).opArm080, (*Cpu).opArm089,
	(*Cpu).opArm090, (*Cpu).opArm090, (*Cpu).opArm090, (*Cpu).opArm090,
	(*Cpu).opArm090, (*Cpu).opArm090, (*Cpu).opArm090, (*Cpu).opArm090,
	(*Cpu).opArm090, (*Cpu).opArm099, (*Cpu).opArm090, (*Cpu).opArm099,
	(*Cpu).opArm090, (*Cpu).opArm099, (*Cpu).opArm090, (*Cpu).opArm099,
	(*Cpu).opArm0A0, (*Cpu).opArm0A0, (*Cpu).opArm0A0, (*Cpu).opArm0A0,
	(*Cpu).opArm0A0, (*Cpu).opArm0A0, (*Cpu).opArm0A0, (*Cpu).opArm0A0,
	(*Cpu).opArm0A0, (*Cpu).opArm0A9, (*Cpu).opArm0A0, (*Cpu).opArm0A9,
	(*Cpu).opArm0A0, (*Cpu).opArm0A9, (*Cpu).opArm0A0, (*Cpu).opArm0A9,
	(*Cpu).opArm0B0, (*Cpu).opArm0B0, (*Cpu).opArm0B0, (*Cpu).opArm0B0,
	(*Cpu).opArm0B0, (*Cpu).opArm0B0, (*Cpu).opArm0B0, (*Cpu).opArm0B0,
	(*Cpu).opArm0B0, (*Cpu).opArm0B9, (*Cpu).opArm0B0, (*Cpu).opArm0B9,
	(*Cpu).opArm0B0, (*Cpu).opArm0B9, (*Cpu).opArm0B0, (*Cpu).opArm0B9,
	(*Cpu).opArm0C0, (*Cpu).opArm0C0, (*Cpu).opArm0C0, (*Cpu).opArm0C0,
	(*Cpu).opArm0C0, (*Cpu).opArm0C0, (*Cpu).opArm0C0, (*Cpu).opArm0C0,
	(*Cpu).opArm0C0, (*Cpu).opArm0C9, (*Cpu).opArm0C0, (*Cpu).opArm0C9,
	(*Cpu).opArm0C0, (*Cpu).opArm0C9, (*Cpu).opArm0C0, (*Cpu).opArm0C9,
	(*Cpu).opArm0D0, (*Cpu).opArm0D0, (*Cpu).opArm0D0, (*Cpu).opArm0D0,
	(*Cpu).opArm0D0, (*Cpu).opArm0D0, (*Cpu).opArm0D0, (*Cpu).opArm0D0,
	(*Cpu).opArm0D0, (*Cpu).opArm0D9, (*Cpu).opArm0D0, (*Cpu).opArm0D9,
	(*Cpu).opArm0D0, (*Cpu).opArm0D9, (*Cpu).opArm0D0, (*Cpu).opArm0D9,
	(*Cpu).opArm0E0, (*Cpu).opArm0E0, (*Cpu).opArm0E0, (*Cpu).opArm0E0,
	(*Cpu).opArm0E0, (*Cpu).opArm0E0, (*Cpu).opArm0E0, (*Cpu).opArm0E0,
	(*Cpu).opArm0E0, (*Cpu).opArm0E9, (*Cpu).opArm0E0, (*Cpu).opArm0E9,
	(*Cpu).opArm0E0, (*Cpu).opArm0E9, (*Cpu).opArm0E0, (*Cpu).opArm0E9,
	(*Cpu).opArm0F0, (*Cpu).opArm0F0, (*Cpu).opArm0F0, (*Cpu).opArm0F0,
	(*Cpu).opArm0F0, (*Cpu).opArm0F0, (*Cpu).opArm0F0, (*Cpu).opArm0F0,
	(*Cpu).opArm0F0, (*Cpu).opArm0F9, (*Cpu).opArm0F0, (*Cpu).opArm0F9,
	(*Cpu).opArm0F0, (*Cpu).opArm0F9, (*Cpu).opArm0F0, (*Cpu).opArm0F9,
	(*Cpu).opArm100, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm109, (*Cpu).opArm101, (*Cpu).opArm109,
	(*Cpu).opArm101, (*Cpu).opArm109, (*Cpu).opArm101, (*Cpu).opArm109,
	(*Cpu).opArm110, (*Cpu).opArm110, (*Cpu).opArm110, (*Cpu).opArm110,
	(*Cpu).opArm110, (*Cpu).opArm110, (*Cpu).opArm110, (*Cpu).opArm110,
	(*Cpu).opArm110, (*Cpu).opArm049, (*Cpu).opArm110, (*Cpu).opArm11B,
	(*Cpu).opArm110, (*Cpu).opArm11D, (*Cpu).opArm110, (*Cpu).opArm11F,
	(*Cpu).opArm120, (*Cpu).opArm121, (*Cpu).opArm101, (*Cpu).opArm123,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm049, (*Cpu).opArm101, (*Cpu).opArm12B,
	(*Cpu).opArm101, (*Cpu).opArm12D, (*Cpu).opArm101, (*Cpu).opArm12F,
	(*Cpu).opArm130, (*Cpu).opArm130, (*Cpu).opArm130, (*Cpu).opArm130,
	(*Cpu).opArm130, (*Cpu).opArm130, (*Cpu).opArm130, (*Cpu).opArm130,
	(*Cpu).opArm130, (*Cpu).opArm049, (*Cpu).opArm130, (*Cpu).opArm13B,
	(*Cpu).opArm130, (*Cpu).opArm13D, (*Cpu).opArm130, (*Cpu).opArm13F,
	(*Cpu).opArm140, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm149, (*Cpu).opArm101, (*Cpu).opArm149,
	(*Cpu).opArm101, (*Cpu).opArm149, (*Cpu).opArm101, (*Cpu).opArm149,
	(*Cpu).opArm150, (*Cpu).opArm150, (*Cpu).opArm150, (*Cpu).opArm150,
	(*Cpu).opArm150, (*Cpu).opArm150, (*Cpu).opArm150, (*Cpu).opArm150,
	(*Cpu).opArm150, (*Cpu).opArm049, (*Cpu).opArm150, (*Cpu).opArm15B,
	(*Cpu).opArm150, (*Cpu).opArm15D, (*Cpu).opArm150, (*Cpu).opArm15F,
	(*Cpu).opArm160, (*Cpu).opArm161, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm049, (*Cpu).opArm101, (*Cpu).opArm16B,
	(*Cpu).opArm101, (*Cpu).opArm16D, (*Cpu).opArm101, (*Cpu).opArm16F,
	(*Cpu).opArm170, (*Cpu).opArm170, (*Cpu).opArm170, (*Cpu).opArm170,
	(*Cpu).opArm170, (*Cpu).opArm170, (*Cpu).opArm170, (*Cpu).opArm170,
	(*Cpu).opArm170, (*Cpu).opArm049, (*Cpu).opArm170, (*Cpu).opArm17B,
	(*Cpu).opArm170, (*Cpu).opArm17D, (*Cpu).opArm170, (*Cpu).opArm17F,
	(*Cpu).opArm180, (*Cpu).opArm180, (*Cpu).opArm180, (*Cpu).opArm180,
	(*Cpu).opArm180, (*Cpu).opArm180, (*Cpu).opArm180, (*Cpu).opArm180,
	(*Cpu).opArm180, (*Cpu).opArm049, (*Cpu).opArm180, (*Cpu).opArm18B,
	(*Cpu).opArm180, (*Cpu).opArm18D, (*Cpu).opArm180, (*Cpu).opArm18F,
	(*Cpu).opArm190, (*Cpu).opArm190, (*Cpu).opArm190, (*Cpu).opArm190,
	(*Cpu).opArm190, (*Cpu).opArm190, (*Cpu).opArm190, (*Cpu).opArm190,
	(*Cpu).opArm190, (*Cpu).opArm049, (*Cpu).opArm190, (*Cpu).opArm19B,
	(*Cpu).opArm190, (*Cpu).opArm19D, (*Cpu).opArm190, (*Cpu).opArm19F,
	(*Cpu).opArm1A0, (*Cpu).opArm1A0, (*Cpu).opArm1A0, (*Cpu).opArm1A0,
	(*Cpu).opArm1A0, (*Cpu).opArm1A0, (*Cpu).opArm1A0, (*Cpu).opArm1A0,
	(*Cpu).opArm1A0, (*Cpu).opArm049, (*Cpu).opArm1A0, (*Cpu).opArm1AB,
	(*Cpu).opArm1A0, (*Cpu).opArm1AD, (*Cpu).opArm1A0, (*Cpu).opArm1AF,
	(*Cpu).opArm1B0, (*Cpu).opArm1B0, (*Cpu).opArm1B0, (*Cpu).opArm1B0,
	(*Cpu).opArm1B0, (*Cpu).opArm1B0, (*Cpu).opArm1B0, (*Cpu).opArm1B0,
	(*Cpu).opArm1B0, (*Cpu).opArm049, (*Cpu).opArm1B0, (*Cpu).opArm1BB,
	(*Cpu).opArm1B0, (*Cpu).opArm1BD, (*Cpu).opArm1B0, (*Cpu).opArm1BF,
	(*Cpu).opArm1C0, (*Cpu).opArm1C0, (*Cpu).opArm1C0, (*Cpu).opArm1C0,
	(*Cpu).opArm1C0, (*Cpu).opArm1C0, (*Cpu).opArm1C0, (*Cpu).opArm1C0,
	(*Cpu).opArm1C0, (*Cpu).opArm049, (*Cpu).opArm1C0, (*Cpu).opArm1CB,
	(*Cpu).opArm1C0, (*Cpu).opArm1CD, (*Cpu).opArm1C0, (*Cpu).opArm1CF,
	(*Cpu).opArm1D0, (*Cpu).opArm1D0, (*Cpu).opArm1D0, (*Cpu).opArm1D0,
	(*Cpu).opArm1D0, (*Cpu).opArm1D0, (*Cpu).opArm1D0, (*Cpu).opArm1D0,
	(*Cpu).opArm1D0, (*Cpu).opArm049, (*Cpu).opArm1D0, (*Cpu).opArm1DB,
	(*Cpu).opArm1D0, (*Cpu).opArm1DD, (*Cpu).opArm1D0, (*Cpu).opArm1DF,
	(*Cpu).opArm1E0, (*Cpu).opArm1E0, (*Cpu).opArm1E0, (*Cpu).opArm1E0,
	(*Cpu).opArm1E0, (*Cpu).opArm1E0, (*Cpu).opArm1E0, (*Cpu).opArm1E0,
	(*Cpu).opArm1E0, (*Cpu).opArm049, (*Cpu).opArm1E0, (*Cpu).opArm1EB,
	(*Cpu).opArm1E0, (*Cpu).opArm1ED, (*Cpu).opArm1E0, (*Cpu).opArm1EF,
	(*Cpu).opArm1F0, (*Cpu).opArm1F0, (*Cpu).opArm1F0, (*Cpu).opArm1F0,
	(*Cpu).opArm1F0, (*Cpu).opArm1F0, (*Cpu).opArm1F0, (*Cpu).opArm1F0,
	(*Cpu).opArm1F0, (*Cpu).opArm049, (*Cpu).opArm1F0, (*Cpu).opArm1FB,
	(*Cpu).opArm1F0, (*Cpu).opArm1FD, (*Cpu).opArm1F0, (*Cpu).opArm1FF,
	(*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200,
	(*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200,
	(*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200,
	(*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200, (*Cpu).opArm200,
	(*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210,
	(*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210,
	(*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210,
	(*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210, (*Cpu).opArm210,
	(*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220,
	(*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220,
	(*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220,
	(*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220, (*Cpu).opArm220,
	(*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230,
	(*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230,
	(*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230,
	(*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230, (*Cpu).opArm230,
	(*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240,
	(*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240,
	(*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240,
	(*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240, (*Cpu).opArm240,
	(*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250,
	(*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250,
	(*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250,
	(*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250, (*Cpu).opArm250,
	(*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260,
	(*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260,
	(*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260,
	(*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260, (*Cpu).opArm260,
	(*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270,
	(*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270,
	(*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270,
	(*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270, (*Cpu).opArm270,
	(*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280,
	(*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280,
	(*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280,
	(*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280, (*Cpu).opArm280,
	(*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290,
	(*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290,
	(*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290,
	(*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290, (*Cpu).opArm290,
	(*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0,
	(*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0,
	(*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0,
	(*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0, (*Cpu).opArm2A0,
	(*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0,
	(*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0,
	(*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0,
	(*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0, (*Cpu).opArm2B0,
	(*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0,
	(*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0,
	(*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0,
	(*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0, (*Cpu).opArm2C0,
	(*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0,
	(*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0,
	(*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0,
	(*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0, (*Cpu).opArm2D0,
	(*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0,
	(*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0,
	(*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0,
	(*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0, (*Cpu).opArm2E0,
	(*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0,
	(*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0,
	(*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0,
	(*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0, (*Cpu).opArm2F0,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310,
	(*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310,
	(*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310,
	(*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310, (*Cpu).opArm310,
	(*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320,
	(*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320,
	(*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320,
	(*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320, (*Cpu).opArm320,
	(*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330,
	(*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330,
	(*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330,
	(*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330, (*Cpu).opArm330,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350,
	(*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350,
	(*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350,
	(*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350, (*Cpu).opArm350,
	(*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360,
	(*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360,
	(*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360,
	(*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360, (*Cpu).opArm360,
	(*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370,
	(*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370,
	(*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370,
	(*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370, (*Cpu).opArm370,
	(*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380,
	(*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380,
	(*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380,
	(*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380, (*Cpu).opArm380,
	(*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390,
	(*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390,
	(*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390,
	(*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390, (*Cpu).opArm390,
	(*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0,
	(*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0,
	(*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0,
	(*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0, (*Cpu).opArm3A0,
	(*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0,
	(*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0,
	(*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0,
	(*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0, (*Cpu).opArm3B0,
	(*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0,
	(*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0,
	(*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0,
	(*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0, (*Cpu).opArm3C0,
	(*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0,
	(*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0,
	(*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0,
	(*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0, (*Cpu).opArm3D0,
	(*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0,
	(*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0,
	(*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0,
	(*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0, (*Cpu).opArm3E0,
	(*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0,
	(*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0,
	(*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0,
	(*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0, (*Cpu).opArm3F0,
	(*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400,
	(*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400,
	(*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400,
	(*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400, (*Cpu).opArm400,
	(*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410,
	(*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410,
	(*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410,
	(*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410, (*Cpu).opArm410,
	(*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420,
	(*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420,
	(*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420,
	(*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420, (*Cpu).opArm420,
	(*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430,
	(*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430,
	(*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430,
	(*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430, (*Cpu).opArm430,
	(*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440,
	(*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440,
	(*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440,
	(*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440, (*Cpu).opArm440,
	(*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450,
	(*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450,
	(*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450,
	(*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450, (*Cpu).opArm450,
	(*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460,
	(*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460,
	(*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460,
	(*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460, (*Cpu).opArm460,
	(*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470,
	(*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470,
	(*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470,
	(*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470, (*Cpu).opArm470,
	(*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480,
	(*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480,
	(*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480,
	(*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480, (*Cpu).opArm480,
	(*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490,
	(*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490,
	(*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490,
	(*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490, (*Cpu).opArm490,
	(*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0,
	(*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0,
	(*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0,
	(*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0, (*Cpu).opArm4A0,
	(*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0,
	(*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0,
	(*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0,
	(*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0, (*Cpu).opArm4B0,
	(*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0,
	(*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0,
	(*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0,
	(*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0, (*Cpu).opArm4C0,
	(*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0,
	(*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0,
	(*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0,
	(*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0, (*Cpu).opArm4D0,
	(*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0,
	(*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0,
	(*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0,
	(*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0, (*Cpu).opArm4E0,
	(*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0,
	(*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0,
	(*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0,
	(*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0, (*Cpu).opArm4F0,
	(*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500,
	(*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500,
	(*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500,
	(*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500, (*Cpu).opArm500,
	(*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510,
	(*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510,
	(*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510,
	(*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510, (*Cpu).opArm510,
	(*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520,
	(*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520,
	(*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520,
	(*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520, (*Cpu).opArm520,
	(*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530,
	(*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530,
	(*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530,
	(*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530, (*Cpu).opArm530,
	(*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540,
	(*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540,
	(*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540,
	(*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540, (*Cpu).opArm540,
	(*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550,
	(*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550,
	(*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550,
	(*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550, (*Cpu).opArm550,
	(*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560,
	(*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560,
	(*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560,
	(*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560, (*Cpu).opArm560,
	(*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570,
	(*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570,
	(*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570,
	(*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570, (*Cpu).opArm570,
	(*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580,
	(*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580,
	(*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580,
	(*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580, (*Cpu).opArm580,
	(*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590,
	(*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590,
	(*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590,
	(*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590, (*Cpu).opArm590,
	(*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0,
	(*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0,
	(*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0,
	(*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0, (*Cpu).opArm5A0,
	(*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0,
	(*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0,
	(*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0,
	(*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0, (*Cpu).opArm5B0,
	(*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0,
	(*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0,
	(*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0,
	(*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0, (*Cpu).opArm5C0,
	(*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0,
	(*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0,
	(*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0,
	(*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0, (*Cpu).opArm5D0,
	(*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0,
	(*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0,
	(*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0,
	(*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0, (*Cpu).opArm5E0,
	(*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0,
	(*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0,
	(*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0,
	(*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0, (*Cpu).opArm5F0,
	(*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600,
	(*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600,
	(*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600,
	(*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600, (*Cpu).opArm600,
	(*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610,
	(*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610,
	(*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610,
	(*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610, (*Cpu).opArm610,
	(*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620,
	(*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620,
	(*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620,
	(*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620, (*Cpu).opArm620,
	(*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630,
	(*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630,
	(*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630,
	(*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630, (*Cpu).opArm630,
	(*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640,
	(*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640,
	(*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640,
	(*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640, (*Cpu).opArm640,
	(*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650,
	(*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650,
	(*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650,
	(*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650, (*Cpu).opArm650,
	(*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660,
	(*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660,
	(*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660,
	(*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660, (*Cpu).opArm660,
	(*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670,
	(*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670,
	(*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670,
	(*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670, (*Cpu).opArm670,
	(*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680,
	(*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680,
	(*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680,
	(*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680, (*Cpu).opArm680,
	(*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690,
	(*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690,
	(*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690,
	(*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690, (*Cpu).opArm690,
	(*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0,
	(*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0,
	(*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0,
	(*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0, (*Cpu).opArm6A0,
	(*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0,
	(*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0,
	(*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0,
	(*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0, (*Cpu).opArm6B0,
	(*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0,
	(*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0,
	(*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0,
	(*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0, (*Cpu).opArm6C0,
	(*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0,
	(*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0,
	(*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0,
	(*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0, (*Cpu).opArm6D0,
	(*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0,
	(*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0,
	(*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0,
	(*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0, (*Cpu).opArm6E0,
	(*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0,
	(*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0,
	(*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0,
	(*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0, (*Cpu).opArm6F0,
	(*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700,
	(*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700,
	(*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700,
	(*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700, (*Cpu).opArm700,
	(*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710,
	(*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710,
	(*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710,
	(*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710, (*Cpu).opArm710,
	(*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720,
	(*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720,
	(*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720,
	(*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720, (*Cpu).opArm720,
	(*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730,
	(*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730,
	(*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730,
	(*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730, (*Cpu).opArm730,
	(*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740,
	(*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740,
	(*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740,
	(*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740, (*Cpu).opArm740,
	(*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750,
	(*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750,
	(*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750,
	(*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750, (*Cpu).opArm750,
	(*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760,
	(*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760,
	(*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760,
	(*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760, (*Cpu).opArm760,
	(*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770,
	(*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770,
	(*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770,
	(*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770, (*Cpu).opArm770,
	(*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780,
	(*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780,
	(*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780,
	(*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780, (*Cpu).opArm780,
	(*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790,
	(*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790,
	(*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790,
	(*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790, (*Cpu).opArm790,
	(*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0,
	(*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0,
	(*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0,
	(*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0, (*Cpu).opArm7A0,
	(*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0,
	(*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0,
	(*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0,
	(*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0, (*Cpu).opArm7B0,
	(*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0,
	(*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0,
	(*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0,
	(*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0, (*Cpu).opArm7C0,
	(*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0,
	(*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0,
	(*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0,
	(*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0, (*Cpu).opArm7D0,
	(*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0,
	(*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0,
	(*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0,
	(*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0, (*Cpu).opArm7E0,
	(*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0,
	(*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0,
	(*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0,
	(*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0, (*Cpu).opArm7F0,
	(*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800,
	(*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800,
	(*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800,
	(*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800, (*Cpu).opArm800,
	(*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810,
	(*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810,
	(*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810,
	(*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810, (*Cpu).opArm810,
	(*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820,
	(*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820,
	(*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820,
	(*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820, (*Cpu).opArm820,
	(*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830,
	(*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830,
	(*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830,
	(*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830, (*Cpu).opArm830,
	(*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840,
	(*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840,
	(*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840,
	(*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840, (*Cpu).opArm840,
	(*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850,
	(*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850,
	(*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850,
	(*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850, (*Cpu).opArm850,
	(*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860,
	(*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860,
	(*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860,
	(*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860, (*Cpu).opArm860,
	(*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870,
	(*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870,
	(*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870,
	(*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870, (*Cpu).opArm870,
	(*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880,
	(*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880,
	(*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880,
	(*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880, (*Cpu).opArm880,
	(*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890,
	(*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890,
	(*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890,
	(*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890, (*Cpu).opArm890,
	(*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0,
	(*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0,
	(*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0,
	(*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0, (*Cpu).opArm8A0,
	(*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0,
	(*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0,
	(*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0,
	(*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0, (*Cpu).opArm8B0,
	(*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0,
	(*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0,
	(*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0,
	(*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0, (*Cpu).opArm8C0,
	(*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0,
	(*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0,
	(*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0,
	(*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0, (*Cpu).opArm8D0,
	(*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0,
	(*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0,
	(*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0,
	(*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0, (*Cpu).opArm8E0,
	(*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0,
	(*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0,
	(*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0,
	(*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0, (*Cpu).opArm8F0,
	(*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900,
	(*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900,
	(*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900,
	(*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900, (*Cpu).opArm900,
	(*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910,
	(*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910,
	(*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910,
	(*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910, (*Cpu).opArm910,
	(*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920,
	(*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920,
	(*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920,
	(*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920, (*Cpu).opArm920,
	(*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930,
	(*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930,
	(*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930,
	(*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930, (*Cpu).opArm930,
	(*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940,
	(*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940,
	(*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940,
	(*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940, (*Cpu).opArm940,
	(*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950,
	(*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950,
	(*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950,
	(*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950, (*Cpu).opArm950,
	(*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960,
	(*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960,
	(*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960,
	(*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960, (*Cpu).opArm960,
	(*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970,
	(*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970,
	(*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970,
	(*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970, (*Cpu).opArm970,
	(*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980,
	(*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980,
	(*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980,
	(*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980, (*Cpu).opArm980,
	(*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990,
	(*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990,
	(*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990,
	(*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990, (*Cpu).opArm990,
	(*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0,
	(*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0,
	(*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0,
	(*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0, (*Cpu).opArm9A0,
	(*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0,
	(*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0,
	(*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0,
	(*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0, (*Cpu).opArm9B0,
	(*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0,
	(*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0,
	(*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0,
	(*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0, (*Cpu).opArm9C0,
	(*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0,
	(*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0,
	(*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0,
	(*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0, (*Cpu).opArm9D0,
	(*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0,
	(*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0,
	(*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0,
	(*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0, (*Cpu).opArm9E0,
	(*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0,
	(*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0,
	(*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0,
	(*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0, (*Cpu).opArm9F0,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00, (*Cpu).opArmA00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00, (*Cpu).opArmB00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00, (*Cpu).opArmC00,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE01, (*Cpu).opArmE00, (*Cpu).opArmE01,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmE00, (*Cpu).opArmE11, (*Cpu).opArmE00, (*Cpu).opArmE11,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
	(*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00, (*Cpu).opArmF00,
}
var disasmArmTable = [4096]func(*Cpu, uint32, uint32) string{
	(*Cpu).disasmArm000, (*Cpu).disasmArm000, (*Cpu).disasmArm000, (*Cpu).disasmArm000,
	(*Cpu).disasmArm000, (*Cpu).disasmArm000, (*Cpu).disasmArm000, (*Cpu).disasmArm000,
	(*Cpu).disasmArm000, (*Cpu).disasmArm009, (*Cpu).disasmArm000, (*Cpu).disasmArm009,
	(*Cpu).disasmArm000, (*Cpu).disasmArm009, (*Cpu).disasmArm000, (*Cpu).disasmArm009,
	(*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010,
	(*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010,
	(*Cpu).disasmArm010, (*Cpu).disasmArm019, (*Cpu).disasmArm010, (*Cpu).disasmArm019,
	(*Cpu).disasmArm010, (*Cpu).disasmArm019, (*Cpu).disasmArm010, (*Cpu).disasmArm019,
	(*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020,
	(*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020,
	(*Cpu).disasmArm020, (*Cpu).disasmArm029, (*Cpu).disasmArm020, (*Cpu).disasmArm029,
	(*Cpu).disasmArm020, (*Cpu).disasmArm029, (*Cpu).disasmArm020, (*Cpu).disasmArm029,
	(*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030,
	(*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030,
	(*Cpu).disasmArm030, (*Cpu).disasmArm039, (*Cpu).disasmArm030, (*Cpu).disasmArm039,
	(*Cpu).disasmArm030, (*Cpu).disasmArm039, (*Cpu).disasmArm030, (*Cpu).disasmArm039,
	(*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040,
	(*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040,
	(*Cpu).disasmArm040, (*Cpu).disasmArm049, (*Cpu).disasmArm040, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm040, (*Cpu).disasmArm04D, (*Cpu).disasmArm040, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050,
	(*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050,
	(*Cpu).disasmArm050, (*Cpu).disasmArm049, (*Cpu).disasmArm050, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm050, (*Cpu).disasmArm05D, (*Cpu).disasmArm050, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060,
	(*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060,
	(*Cpu).disasmArm060, (*Cpu).disasmArm049, (*Cpu).disasmArm060, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm060, (*Cpu).disasmArm04D, (*Cpu).disasmArm060, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070,
	(*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070,
	(*Cpu).disasmArm070, (*Cpu).disasmArm049, (*Cpu).disasmArm070, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm070, (*Cpu).disasmArm05D, (*Cpu).disasmArm070, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080,
	(*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080,
	(*Cpu).disasmArm080, (*Cpu).disasmArm089, (*Cpu).disasmArm080, (*Cpu).disasmArm089,
	(*Cpu).disasmArm080, (*Cpu).disasmArm089, (*Cpu).disasmArm080, (*Cpu).disasmArm089,
	(*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090,
	(*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090,
	(*Cpu).disasmArm090, (*Cpu).disasmArm099, (*Cpu).disasmArm090, (*Cpu).disasmArm099,
	(*Cpu).disasmArm090, (*Cpu).disasmArm099, (*Cpu).disasmArm090, (*Cpu).disasmArm099,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm0A9, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A9,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm0A9, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A9,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm0B9, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B9,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm0B9, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B9,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm0C9, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C9,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm0C9, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C9,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm0D9, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D9,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm0D9, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D9,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm0E9, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E9,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm0E9, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E9,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm0F9, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F9,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm0F9, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F9,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm109, (*Cpu).disasmArm049, (*Cpu).disasmArm109,
	(*Cpu).disasmArm049, (*Cpu).disasmArm109, (*Cpu).disasmArm049, (*Cpu).disasmArm109,
	(*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110,
	(*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110,
	(*Cpu).disasmArm110, (*Cpu).disasmArm049, (*Cpu).disasmArm110, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm110, (*Cpu).disasmArm05D, (*Cpu).disasmArm110, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm049, (*Cpu).disasmArm121, (*Cpu).disasmArm049, (*Cpu).disasmArm123,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm049, (*Cpu).disasmArm04D, (*Cpu).disasmArm049, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130,
	(*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130,
	(*Cpu).disasmArm130, (*Cpu).disasmArm049, (*Cpu).disasmArm130, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm130, (*Cpu).disasmArm05D, (*Cpu).disasmArm130, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm149, (*Cpu).disasmArm049, (*Cpu).disasmArm149,
	(*Cpu).disasmArm049, (*Cpu).disasmArm149, (*Cpu).disasmArm049, (*Cpu).disasmArm149,
	(*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150,
	(*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150,
	(*Cpu).disasmArm150, (*Cpu).disasmArm049, (*Cpu).disasmArm150, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm150, (*Cpu).disasmArm05D, (*Cpu).disasmArm150, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm049, (*Cpu).disasmArm161, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm049, (*Cpu).disasmArm04D, (*Cpu).disasmArm049, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170,
	(*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170,
	(*Cpu).disasmArm170, (*Cpu).disasmArm049, (*Cpu).disasmArm170, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm170, (*Cpu).disasmArm05D, (*Cpu).disasmArm170, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180,
	(*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180,
	(*Cpu).disasmArm180, (*Cpu).disasmArm049, (*Cpu).disasmArm180, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm180, (*Cpu).disasmArm04D, (*Cpu).disasmArm180, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190,
	(*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190,
	(*Cpu).disasmArm190, (*Cpu).disasmArm049, (*Cpu).disasmArm190, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm190, (*Cpu).disasmArm05D, (*Cpu).disasmArm190, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm049, (*Cpu).disasmArm1A0, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm04D, (*Cpu).disasmArm1A0, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm049, (*Cpu).disasmArm1B0, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm05D, (*Cpu).disasmArm1B0, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm049, (*Cpu).disasmArm1C0, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm04D, (*Cpu).disasmArm1C0, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm049, (*Cpu).disasmArm1D0, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm05D, (*Cpu).disasmArm1D0, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm049, (*Cpu).disasmArm1E0, (*Cpu).disasmArm04B,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm04D, (*Cpu).disasmArm1E0, (*Cpu).disasmArm04F,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm049, (*Cpu).disasmArm1F0, (*Cpu).disasmArm05B,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm05D, (*Cpu).disasmArm1F0, (*Cpu).disasmArm05F,
	(*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200,
	(*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200,
	(*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200,
	(*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200, (*Cpu).disasmArm200,
	(*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210,
	(*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210,
	(*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210,
	(*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210, (*Cpu).disasmArm210,
	(*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220,
	(*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220,
	(*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220,
	(*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220, (*Cpu).disasmArm220,
	(*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230,
	(*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230,
	(*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230,
	(*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230, (*Cpu).disasmArm230,
	(*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240,
	(*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240,
	(*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240,
	(*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240, (*Cpu).disasmArm240,
	(*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250,
	(*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250,
	(*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250,
	(*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250, (*Cpu).disasmArm250,
	(*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260,
	(*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260,
	(*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260,
	(*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260, (*Cpu).disasmArm260,
	(*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270,
	(*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270,
	(*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270,
	(*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270, (*Cpu).disasmArm270,
	(*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280,
	(*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280,
	(*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280,
	(*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280, (*Cpu).disasmArm280,
	(*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290,
	(*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290,
	(*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290,
	(*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290, (*Cpu).disasmArm290,
	(*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0,
	(*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0,
	(*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0,
	(*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0, (*Cpu).disasmArm2A0,
	(*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0,
	(*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0,
	(*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0,
	(*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0, (*Cpu).disasmArm2B0,
	(*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0,
	(*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0,
	(*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0,
	(*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0, (*Cpu).disasmArm2C0,
	(*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0,
	(*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0,
	(*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0,
	(*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0, (*Cpu).disasmArm2D0,
	(*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0,
	(*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0,
	(*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0,
	(*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0, (*Cpu).disasmArm2E0,
	(*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0,
	(*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0,
	(*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0,
	(*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0, (*Cpu).disasmArm2F0,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310,
	(*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310,
	(*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310,
	(*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310, (*Cpu).disasmArm310,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330,
	(*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330,
	(*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330,
	(*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330, (*Cpu).disasmArm330,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350,
	(*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350,
	(*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350,
	(*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350, (*Cpu).disasmArm350,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370,
	(*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370,
	(*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370,
	(*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370, (*Cpu).disasmArm370,
	(*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380,
	(*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380,
	(*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380,
	(*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380, (*Cpu).disasmArm380,
	(*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390,
	(*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390,
	(*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390,
	(*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390, (*Cpu).disasmArm390,
	(*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0,
	(*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0,
	(*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0,
	(*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0, (*Cpu).disasmArm3A0,
	(*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0,
	(*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0,
	(*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0,
	(*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0, (*Cpu).disasmArm3B0,
	(*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0,
	(*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0,
	(*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0,
	(*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0, (*Cpu).disasmArm3C0,
	(*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0,
	(*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0,
	(*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0,
	(*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0, (*Cpu).disasmArm3D0,
	(*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0,
	(*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0,
	(*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0,
	(*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0, (*Cpu).disasmArm3E0,
	(*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0,
	(*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0,
	(*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0,
	(*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0, (*Cpu).disasmArm3F0,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400, (*Cpu).disasmArm400,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410, (*Cpu).disasmArm410,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440, (*Cpu).disasmArm440,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450, (*Cpu).disasmArm450,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480, (*Cpu).disasmArm480,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490, (*Cpu).disasmArm490,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0, (*Cpu).disasmArm4C0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0, (*Cpu).disasmArm4D0,
	(*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500,
	(*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500,
	(*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500,
	(*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500, (*Cpu).disasmArm500,
	(*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510,
	(*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510,
	(*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510,
	(*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510, (*Cpu).disasmArm510,
	(*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520,
	(*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520,
	(*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520,
	(*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520, (*Cpu).disasmArm520,
	(*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530,
	(*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530,
	(*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530,
	(*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530, (*Cpu).disasmArm530,
	(*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540,
	(*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540,
	(*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540,
	(*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540, (*Cpu).disasmArm540,
	(*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550,
	(*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550,
	(*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550,
	(*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550, (*Cpu).disasmArm550,
	(*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560,
	(*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560,
	(*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560,
	(*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560, (*Cpu).disasmArm560,
	(*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570,
	(*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570,
	(*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570,
	(*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570, (*Cpu).disasmArm570,
	(*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580,
	(*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580,
	(*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580,
	(*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580, (*Cpu).disasmArm580,
	(*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590,
	(*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590,
	(*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590,
	(*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590, (*Cpu).disasmArm590,
	(*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0,
	(*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0,
	(*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0,
	(*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0, (*Cpu).disasmArm5A0,
	(*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0,
	(*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0,
	(*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0,
	(*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0, (*Cpu).disasmArm5B0,
	(*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0,
	(*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0,
	(*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0,
	(*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0, (*Cpu).disasmArm5C0,
	(*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0,
	(*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0,
	(*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0,
	(*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0, (*Cpu).disasmArm5D0,
	(*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0,
	(*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0,
	(*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0,
	(*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0, (*Cpu).disasmArm5E0,
	(*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0,
	(*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0,
	(*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0,
	(*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0, (*Cpu).disasmArm5F0,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600, (*Cpu).disasmArm600,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610, (*Cpu).disasmArm610,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640, (*Cpu).disasmArm640,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650, (*Cpu).disasmArm650,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680, (*Cpu).disasmArm680,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690, (*Cpu).disasmArm690,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0, (*Cpu).disasmArm6C0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0, (*Cpu).disasmArm6D0,
	(*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700,
	(*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700,
	(*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700,
	(*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700, (*Cpu).disasmArm700,
	(*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710,
	(*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710,
	(*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710,
	(*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710, (*Cpu).disasmArm710,
	(*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720,
	(*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720,
	(*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720,
	(*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720, (*Cpu).disasmArm720,
	(*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730,
	(*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730,
	(*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730,
	(*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730, (*Cpu).disasmArm730,
	(*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740,
	(*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740,
	(*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740,
	(*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740, (*Cpu).disasmArm740,
	(*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750,
	(*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750,
	(*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750,
	(*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750, (*Cpu).disasmArm750,
	(*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760,
	(*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760,
	(*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760,
	(*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760, (*Cpu).disasmArm760,
	(*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770,
	(*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770,
	(*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770,
	(*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770, (*Cpu).disasmArm770,
	(*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780,
	(*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780,
	(*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780,
	(*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780, (*Cpu).disasmArm780,
	(*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790,
	(*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790,
	(*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790,
	(*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790, (*Cpu).disasmArm790,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0, (*Cpu).disasmArm7A0,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0, (*Cpu).disasmArm7B0,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0, (*Cpu).disasmArm7C0,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0, (*Cpu).disasmArm7D0,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0, (*Cpu).disasmArm7E0,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0, (*Cpu).disasmArm7F0,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800, (*Cpu).disasmArm800,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810, (*Cpu).disasmArm810,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820, (*Cpu).disasmArm820,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830, (*Cpu).disasmArm830,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880, (*Cpu).disasmArm880,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890, (*Cpu).disasmArm890,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0, (*Cpu).disasmArm8A0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0, (*Cpu).disasmArm8B0,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900, (*Cpu).disasmArm900,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910, (*Cpu).disasmArm910,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920, (*Cpu).disasmArm920,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930, (*Cpu).disasmArm930,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980, (*Cpu).disasmArm980,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990, (*Cpu).disasmArm990,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0, (*Cpu).disasmArm9A0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0, (*Cpu).disasmArm9B0,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00, (*Cpu).disasmArmA00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00, (*Cpu).disasmArmB00,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
	(*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00, (*Cpu).disasmArmF00,
}
