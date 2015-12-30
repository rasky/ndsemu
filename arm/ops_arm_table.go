// Generated on 2015-12-30 12:33:50.2007402 +0100 CET
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm001(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm002(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm003(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm004(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm005(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm006(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm007(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 1
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 2
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 3
	} else {
		cpu.Clock += 4
	}
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

func (cpu *Cpu) opArm00B(op uint32) {
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
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmArm00B(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strh", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm00D(op uint32) {
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
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmArm00D(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrd", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm00F(op uint32) {
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
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmArm00F(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("strd", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm010(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm011(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm012(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm013(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm014(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm015(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm016(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm017(op uint32) {
	// ands
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 1
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 2
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 3
	} else {
		cpu.Clock += 4
	}
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

func (cpu *Cpu) opArm01B(op uint32) {
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
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmArm01B(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrh", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm01D(op uint32) {
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
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmArm01D(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrsb", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm01F(op uint32) {
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
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) disasmArm01F(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("ldrsh", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	return out.String()
}

func (cpu *Cpu) opArm020(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm021(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm022(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm023(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm024(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm025(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm026(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm027(op uint32) {
	// eor
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 2
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 3
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 4
	} else {
		cpu.Clock += 5
	}
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm031(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm032(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm033(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm034(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm035(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm036(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm037(op uint32) {
	// eors
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 2
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 3
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 4
	} else {
		cpu.Clock += 5
	}
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm041(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm042(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm043(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm044(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm045(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm046(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm047(op uint32) {
	// sub
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cpu.Clock += 1
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm050(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm051(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm052(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm053(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm054(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm055(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm056(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm057(op uint32) {
	// subs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm060(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm061(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm062(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm063(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm064(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm065(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm066(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm067(op uint32) {
	// rsb
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm070(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm071(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm072(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm073(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm074(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm075(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm076(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm077(op uint32) {
	// rsbs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm080(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm081(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm082(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm083(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm084(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm085(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm086(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm087(op uint32) {
	// add
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 1
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 2
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 3
	} else {
		cpu.Clock += 4
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm08B(op uint32) {
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
	// STRH
	cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm08D(op uint32) {
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
	// LDRD
	cpu.Regs[rdx] = reg(cpu.opRead32(rn))
	cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm08F(op uint32) {
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
	// STRD
	cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
	cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm090(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm091(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm092(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm093(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm094(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm095(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm096(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm097(op uint32) {
	// adds
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 1
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 2
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 3
	} else {
		cpu.Clock += 4
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm09B(op uint32) {
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
	// LDRH
	cpu.Regs[rdx] = reg(cpu.opRead16(rn))
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm09D(op uint32) {
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
	// LDRSB
	data := int32(int8(cpu.opRead8(rn)))
	cpu.Regs[rdx] = reg(data)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm09F(op uint32) {
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
	// LDRSH
	data := int32(int16(cpu.opRead16(rn)))
	cpu.Regs[rdx] = reg(data)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm0A0(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm0A1(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0A2(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0A3(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0A4(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0A5(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0A6(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0A7(op uint32) {
	// adc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 2
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 3
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 4
	} else {
		cpu.Clock += 5
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm0B1(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0B2(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0B3(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0B4(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0B5(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0B6(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0B7(op uint32) {
	// adcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 2
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 3
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 4
	} else {
		cpu.Clock += 5
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm0C1(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0C2(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0C3(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0C4(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0C5(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0C6(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0C7(op uint32) {
	// sbc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 1
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 2
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 3
	} else {
		cpu.Clock += 4
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0CB(op uint32) {
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
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm0CD(op uint32) {
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm0CF(op uint32) {
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
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm0D0(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm0D1(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0D2(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0D3(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0D4(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0D5(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0D6(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0D7(op uint32) {
	// sbcs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 1
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 2
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 3
	} else {
		cpu.Clock += 4
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
	out.WriteString(RegNames[arg3])
	return out.String()
}

func (cpu *Cpu) opArm0DB(op uint32) {
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm0DD(op uint32) {
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm0DF(op uint32) {
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm0E0(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm0E1(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0E2(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0E3(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0E4(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0E5(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0E6(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0E7(op uint32) {
	// rsc
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 2
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 3
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 4
	} else {
		cpu.Clock += 5
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm0F1(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0F2(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0F3(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0F4(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0F5(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0F6(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm0F7(op uint32) {
	// rscs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rs&0xFFFFFF00 == 0 || ^rs&0xFFFFFF00 == 0 {
		cpu.Clock += 2
	} else if rs&0xFFFF0000 == 0 || ^rs&0xFFFF0000 == 0 {
		cpu.Clock += 3
	} else if rs&0xFF000000 == 0 || ^rs&0xFF000000 == 0 {
		cpu.Clock += 4
	} else {
		cpu.Clock += 5
	}
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
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := (op >> 16) & 0xF
	out.WriteString(RegNames[arg1])
	out.WriteString(", ")
	arg2 := (op >> 0) & 0xF
	out.WriteString(RegNames[arg2])
	out.WriteString(", ")
	arg3 := (op >> 8) & 0xF
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

func (cpu *Cpu) disasmArm100(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mrs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	out.WriteString("cpsr")
	return out.String()
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
	cpu.Clock += 1
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

func (cpu *Cpu) opArm10B(op uint32) {
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm10D(op uint32) {
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm10F(op uint32) {
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm110(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm111(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm112(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm113(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm114(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm115(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm116(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm117(op uint32) {
	// tsts
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Clock += 1
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

func (cpu *Cpu) disasmArm120(op uint32, pc uint32) string {
	dst := "cpsr_"
	if (op>>19)&1 != 0 {
		dst += "f"
	}
	if (op>>18)&1 != 0 {
		dst += "s"
	}
	if (op>>17)&1 != 0 {
		dst += "x"
	}
	if (op>>16)&1 != 0 {
		dst += "c"
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("msr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := dst
	out.WriteString(arg0)
	out.WriteString(", ")
	arg1 := op & 0xF
	out.WriteString(RegNames[arg1])
	return out.String()
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
	cpu.branch(rn, BranchJump)
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
	cpu.branch(rn, BranchCall)
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
	cpu.Clock += 1
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm130(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm131(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm132(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm133(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm134(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm135(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm136(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm137(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) disasmArm140(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("mrs", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := (op >> 12) & 0xF
	out.WriteString(RegNames[arg0])
	out.WriteString(", ")
	arg1 := cpu.disasmSpsrName()
	out.WriteString(arg1)
	return out.String()
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
	cpu.Clock += 1
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

func (cpu *Cpu) opArm14B(op uint32) {
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm14D(op uint32) {
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm14F(op uint32) {
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm150(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
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
	_ = cf
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

func (cpu *Cpu) opArm151(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm152(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
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
	_ = cf
}

func (cpu *Cpu) opArm153(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm154(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
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
	_ = cf
}

func (cpu *Cpu) opArm155(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm156(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm157(op uint32) {
	// cmps
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
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
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Clock += 1
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

func (cpu *Cpu) disasmArm160(op uint32, pc uint32) string {
	dst := cpu.disasmSpsrName() + "_"
	if (op>>19)&1 != 0 {
		dst += "f"
	}
	if (op>>18)&1 != 0 {
		dst += "s"
	}
	if (op>>17)&1 != 0 {
		dst += "x"
	}
	if (op>>16)&1 != 0 {
		dst += "c"
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("msr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := dst
	out.WriteString(arg0)
	out.WriteString(", ")
	arg1 := op & 0xF
	out.WriteString(RegNames[arg1])
	return out.String()
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
	cpu.Clock += 1
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm170(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
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
	_ = cf
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

func (cpu *Cpu) opArm171(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm172(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
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
	_ = cf
}

func (cpu *Cpu) opArm173(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm174(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
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
	_ = cf
}

func (cpu *Cpu) opArm175(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm176(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
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
	_ = cf
}

func (cpu *Cpu) opArm177(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
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
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm180(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm181(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm182(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm183(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm184(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm185(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm186(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm187(op uint32) {
	// orr
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cpu.Clock += 1
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm190(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm191(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm192(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm193(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm194(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm195(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm196(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm197(op uint32) {
	// orrs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm1A0(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm1A1(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1A2(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1A3(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1A4(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1A5(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1A6(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1A7(op uint32) {
	// mov
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cpu.Clock += 1
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm1B0(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm1B1(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1B2(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1B3(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1B4(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1B5(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1B6(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1B7(op uint32) {
	// movs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	if rnx != 0 {
		cpu.InvalidOpArm(op, "rn!=0 on NOV")
		return
	}
	res := op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm1C0(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm1C1(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1C2(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1C3(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1C4(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1C5(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1C6(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1C7(op uint32) {
	// bic
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cpu.Clock += 1
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm1D0(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm1D1(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1D2(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1D3(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1D4(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1D5(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1D6(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1D7(op uint32) {
	// bics
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm1E0(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm1E1(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1E2(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1E3(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1E4(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1E5(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1E6(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1E7(op uint32) {
	// mvn
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cpu.Clock += 1
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
	if rdx == 14 {
		cpu.InvalidOpArm(op, "LDRD PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArm1F0(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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

func (cpu *Cpu) opArm1F1(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(32-shift))&1 != 0)
	op2 <<= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1F2(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1F3(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 >>= shift
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1F4(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1F5(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	cpu.Cpsr.SetC((op2>>(shift-1))&1 != 0)
	op2 = uint32(int32(op2) >> shift)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1F6(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		cpu.Cpsr.SetC(op2&1 != 0)
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
}

func (cpu *Cpu) opArm1F7(op uint32) {
	// mvns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=true
	op2 := uint32(cpu.Regs[op&0xF])
	cpu.Regs[15] += 4
	shift := uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
	if shift == 0 {
		goto op2end
	}
	if shift >= 32 {
		cpu.InvalidOpArm(op, "big shift register not implemented")
		return
	}
	cpu.Clock += 1
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
	cpu.Cpsr.SetC(op2>>31 != 0)
op2end:
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSB PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "LDRSH PC not implemented")
		return
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm200(op uint32) {
	// and
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn ^ op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn + op2
	res += cf
	cpu.Cpsr.SetC(rn > res)
	cpu.Cpsr.SetVAdd(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	res += cf - 1
	cpu.Cpsr.SetC(res <= rn)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	res += cf - 1
	cpu.Cpsr.SetC(res <= op2)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
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
	_ = cf
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

func (cpu *Cpu) disasmArm320(op uint32, pc uint32) string {
	dst := "cpsr_"
	if (op>>19)&1 != 0 {
		dst += "f"
	}
	if (op>>18)&1 != 0 {
		dst += "s"
	}
	if (op>>17)&1 != 0 {
		dst += "x"
	}
	if (op>>16)&1 != 0 {
		dst += "c"
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("msr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := dst
	out.WriteString(arg0)
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> uint(((op>>8)&0xF)*2)) | ((op & 0xFF) << (32 - uint(((op>>8)&0xF)*2))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm330(op uint32) {
	// teqs
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
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
	_ = cf
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
	cf := cpu.Cpsr.CB()
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
	_ = cf
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

func (cpu *Cpu) disasmArm360(op uint32, pc uint32) string {
	dst := cpu.disasmSpsrName() + "_"
	if (op>>19)&1 != 0 {
		dst += "f"
	}
	if (op>>18)&1 != 0 {
		dst += "s"
	}
	if (op>>17)&1 != 0 {
		dst += "x"
	}
	if (op>>16)&1 != 0 {
		dst += "c"
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("msr", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := dst
	out.WriteString(arg0)
	out.WriteString(", ")
	arg1 := int64(((op & 0xFF) >> uint(((op>>8)&0xF)*2)) | ((op & 0xFF) << (32 - uint(((op>>8)&0xF)*2))))
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg1, 16))
	return out.String()
}

func (cpu *Cpu) opArm370(op uint32) {
	// cmns
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	cf := cpu.Cpsr.CB()
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
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn | op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
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
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
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
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn & ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cf := cpu.Cpsr.CB()
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := ^op2
	cpu.Cpsr.SetNZ(res)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
		cpu.branch(reg(res)&^1, BranchJump)
	}
	_ = res
	_ = rn
	_ = cf
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm601(op uint32) {
	// undefined
	cpu.Exception(ExceptionUndefined)
}

func (cpu *Cpu) opArm602(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm604(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm606(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm612(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm614(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm616(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm622(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm624(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm626(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm632(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm634(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm636(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm642(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm644(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm646(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm652(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm654(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm656(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm662(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm664(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm666(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm672(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm674(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm676(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn -= off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm682(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm684(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm686(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm692(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm694(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm696(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6A2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6A4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6A6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6B2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6B4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6B6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm6C2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6C4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6C6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm6D2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6D4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6D6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6E2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6E4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6E6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6F2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6F4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm6F6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	rn += off
	cpu.InvalidOpArm(op, "forced-unprivileged memory access")
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
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

func (cpu *Cpu) opArm702(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm704(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm706(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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

func (cpu *Cpu) opArm712(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm714(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm716(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm722(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm724(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm726(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm732(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm734(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm736(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
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

func (cpu *Cpu) opArm742(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm744(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm746(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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

func (cpu *Cpu) opArm752(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm754(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm756(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm762(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm764(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm766(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm772(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm774(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm776(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn -= off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
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

func (cpu *Cpu) opArm782(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm784(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm786(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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

func (cpu *Cpu) opArm792(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm794(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm796(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm7A2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7A4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7A6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite32(rn, uint32(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm7B2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7B4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7B6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	res := cpu.opRead32(rn)
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
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

func (cpu *Cpu) opArm7C2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7C4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7C6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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

func (cpu *Cpu) opArm7D2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7D4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7D6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm7E2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7E4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7E6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	rd := cpu.Regs[rdx]
	cpu.opWrite8(rn, uint8(rd))
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsl, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		goto op2end
	}
	op2 <<= shift
op2end:
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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

func (cpu *Cpu) opArm7F2(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=lsr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 >>= shift
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7F4(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=asr, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 {
		shift = 32
	}
	op2 = uint32(int32(op2) >> shift)
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
}

func (cpu *Cpu) opArm7F6(op uint32) {
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
	cf := cpu.Cpsr.CB()
	// op2: shtype=ror, byreg=false
	op2 := uint32(cpu.Regs[op&0xF])
	shift := uint32((op >> 7) & 0x1F)
	if shift == 0 { // becomes RRX #1
		op2 = (op2 >> 1) | (cf << 31)
		goto op2end
	}
	shift &= 31
	op2 = (op2 >> shift) | (op2 << (32 - shift))
op2end:
	_ = cf
	off := op2
	rn += off
	res := uint32(cpu.opRead8(rn))
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.Cpsr.SetT((res & 1) != 0)
		cpu.branch(reg(res&^1), BranchJump)
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
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
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
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
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
			rn += 4
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
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
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	cpu.Clock += 1
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
	cpu.Clock += 1
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
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			var val uint32
			val = uint32(cpu.Regs[i])
			cpu.opWrite32(rn, val)
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	usrbnk := true
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
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
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
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
	oldmode := cpu.Cpsr.GetMode()
	if usrbnk {
		cpu.Cpsr.SetMode(CpuModeUser, cpu)
	}
	for i := 0; mask != 0; i++ {
		if mask&1 != 0 {
			rn += 4
			val := reg(cpu.opRead32(rn))
			cpu.Regs[i] = val
			if i == 15 {
				cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)
				if cpu.Regs[15]&1 != 0 {
					cpu.Cpsr.SetT(true)
					cpu.Regs[15] &^= 1
				} else {
					cpu.Regs[15] &^= 3
				}
				cpu.branch(cpu.Regs[15], BranchJump)
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
	if usrbnk {
		cpu.Cpsr.SetMode(oldmode, cpu)
	}
	cpu.Clock += 1
}

func (cpu *Cpu) opArmA00(op uint32) {
	if op>>28 == 0xF {
		// BLX_imm
		off := int32(op<<8) >> 6
		cpu.Regs[14] = cpu.Regs[15] - 4
		cpu.Regs[15] += reg(off)
		cpu.Cpsr.SetT(true)
		cpu.branch(cpu.Regs[15], BranchCall)
		return
	}
	// B
	if !cpu.opArmCond(op) {
		return
	}
	off := int32(op<<8) >> 6
	cpu.Regs[15] += reg(off)
	cpu.branch(cpu.Regs[15], BranchCall)
}

func (cpu *Cpu) disasmArmA00(op uint32, pc uint32) string {
	if op>>28 == 0xF {
		var out bytes.Buffer
		out.WriteString("blx       ")
		arg0 := int32(int32(op<<8) >> 6)
		arg0x := pc + 8 + uint32(arg0)
		out.WriteString(strconv.FormatInt(int64(arg0x), 16))
		return out.String()
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("b", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := int32(int32(op<<8) >> 6)
	arg0x := pc + 8 + uint32(arg0)
	out.WriteString(strconv.FormatInt(int64(arg0x), 16))
	return out.String()
}

func (cpu *Cpu) opArmB00(op uint32) {
	if op>>28 == 0xF {
		// BLX_imm
		off := int32(op<<8) >> 6
		cpu.Regs[14] = cpu.Regs[15] - 4
		cpu.Regs[15] += reg(off)
		cpu.Regs[15] += 2
		cpu.Cpsr.SetT(true)
		cpu.branch(cpu.Regs[15], BranchCall)
		return
	}
	// BL
	if !cpu.opArmCond(op) {
		return
	}
	off := int32(op<<8) >> 6
	cpu.Regs[14] = cpu.Regs[15] - 4
	cpu.Regs[15] += reg(off)
	cpu.branch(cpu.Regs[15], BranchCall)
}

func (cpu *Cpu) disasmArmB00(op uint32, pc uint32) string {
	if op>>28 == 0xF {
		var out bytes.Buffer
		out.WriteString("blx       ")
		arg0 := int32(int32(op<<8) >> 6)
		arg0x := pc + 8 + uint32(arg0)
		out.WriteString(strconv.FormatInt(int64(arg0x), 16))
		return out.String()
	}
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("bl", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := int32(int32(op<<8) >> 6)
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
	cpu.Clock += 1
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
	cpu.Clock += 1
}

func (cpu *Cpu) opArmF00(op uint32) {
	cpu.Exception(ExceptionSwi)
	cpu.Clock += 2
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
	(*Cpu).opArm000, (*Cpu).opArm001, (*Cpu).opArm002, (*Cpu).opArm003,
	(*Cpu).opArm004, (*Cpu).opArm005, (*Cpu).opArm006, (*Cpu).opArm007,
	(*Cpu).opArm000, (*Cpu).opArm009, (*Cpu).opArm002, (*Cpu).opArm00B,
	(*Cpu).opArm004, (*Cpu).opArm00D, (*Cpu).opArm006, (*Cpu).opArm00F,
	(*Cpu).opArm010, (*Cpu).opArm011, (*Cpu).opArm012, (*Cpu).opArm013,
	(*Cpu).opArm014, (*Cpu).opArm015, (*Cpu).opArm016, (*Cpu).opArm017,
	(*Cpu).opArm010, (*Cpu).opArm019, (*Cpu).opArm012, (*Cpu).opArm01B,
	(*Cpu).opArm014, (*Cpu).opArm01D, (*Cpu).opArm016, (*Cpu).opArm01F,
	(*Cpu).opArm020, (*Cpu).opArm021, (*Cpu).opArm022, (*Cpu).opArm023,
	(*Cpu).opArm024, (*Cpu).opArm025, (*Cpu).opArm026, (*Cpu).opArm027,
	(*Cpu).opArm020, (*Cpu).opArm029, (*Cpu).opArm022, (*Cpu).opArm00B,
	(*Cpu).opArm024, (*Cpu).opArm00D, (*Cpu).opArm026, (*Cpu).opArm00F,
	(*Cpu).opArm030, (*Cpu).opArm031, (*Cpu).opArm032, (*Cpu).opArm033,
	(*Cpu).opArm034, (*Cpu).opArm035, (*Cpu).opArm036, (*Cpu).opArm037,
	(*Cpu).opArm030, (*Cpu).opArm039, (*Cpu).opArm032, (*Cpu).opArm01B,
	(*Cpu).opArm034, (*Cpu).opArm01D, (*Cpu).opArm036, (*Cpu).opArm01F,
	(*Cpu).opArm040, (*Cpu).opArm041, (*Cpu).opArm042, (*Cpu).opArm043,
	(*Cpu).opArm044, (*Cpu).opArm045, (*Cpu).opArm046, (*Cpu).opArm047,
	(*Cpu).opArm040, (*Cpu).opArm049, (*Cpu).opArm042, (*Cpu).opArm04B,
	(*Cpu).opArm044, (*Cpu).opArm04D, (*Cpu).opArm046, (*Cpu).opArm04F,
	(*Cpu).opArm050, (*Cpu).opArm051, (*Cpu).opArm052, (*Cpu).opArm053,
	(*Cpu).opArm054, (*Cpu).opArm055, (*Cpu).opArm056, (*Cpu).opArm057,
	(*Cpu).opArm050, (*Cpu).opArm049, (*Cpu).opArm052, (*Cpu).opArm05B,
	(*Cpu).opArm054, (*Cpu).opArm05D, (*Cpu).opArm056, (*Cpu).opArm05F,
	(*Cpu).opArm060, (*Cpu).opArm061, (*Cpu).opArm062, (*Cpu).opArm063,
	(*Cpu).opArm064, (*Cpu).opArm065, (*Cpu).opArm066, (*Cpu).opArm067,
	(*Cpu).opArm060, (*Cpu).opArm049, (*Cpu).opArm062, (*Cpu).opArm04B,
	(*Cpu).opArm064, (*Cpu).opArm04D, (*Cpu).opArm066, (*Cpu).opArm04F,
	(*Cpu).opArm070, (*Cpu).opArm071, (*Cpu).opArm072, (*Cpu).opArm073,
	(*Cpu).opArm074, (*Cpu).opArm075, (*Cpu).opArm076, (*Cpu).opArm077,
	(*Cpu).opArm070, (*Cpu).opArm049, (*Cpu).opArm072, (*Cpu).opArm05B,
	(*Cpu).opArm074, (*Cpu).opArm05D, (*Cpu).opArm076, (*Cpu).opArm05F,
	(*Cpu).opArm080, (*Cpu).opArm081, (*Cpu).opArm082, (*Cpu).opArm083,
	(*Cpu).opArm084, (*Cpu).opArm085, (*Cpu).opArm086, (*Cpu).opArm087,
	(*Cpu).opArm080, (*Cpu).opArm089, (*Cpu).opArm082, (*Cpu).opArm08B,
	(*Cpu).opArm084, (*Cpu).opArm08D, (*Cpu).opArm086, (*Cpu).opArm08F,
	(*Cpu).opArm090, (*Cpu).opArm091, (*Cpu).opArm092, (*Cpu).opArm093,
	(*Cpu).opArm094, (*Cpu).opArm095, (*Cpu).opArm096, (*Cpu).opArm097,
	(*Cpu).opArm090, (*Cpu).opArm099, (*Cpu).opArm092, (*Cpu).opArm09B,
	(*Cpu).opArm094, (*Cpu).opArm09D, (*Cpu).opArm096, (*Cpu).opArm09F,
	(*Cpu).opArm0A0, (*Cpu).opArm0A1, (*Cpu).opArm0A2, (*Cpu).opArm0A3,
	(*Cpu).opArm0A4, (*Cpu).opArm0A5, (*Cpu).opArm0A6, (*Cpu).opArm0A7,
	(*Cpu).opArm0A0, (*Cpu).opArm0A9, (*Cpu).opArm0A2, (*Cpu).opArm08B,
	(*Cpu).opArm0A4, (*Cpu).opArm08D, (*Cpu).opArm0A6, (*Cpu).opArm08F,
	(*Cpu).opArm0B0, (*Cpu).opArm0B1, (*Cpu).opArm0B2, (*Cpu).opArm0B3,
	(*Cpu).opArm0B4, (*Cpu).opArm0B5, (*Cpu).opArm0B6, (*Cpu).opArm0B7,
	(*Cpu).opArm0B0, (*Cpu).opArm0B9, (*Cpu).opArm0B2, (*Cpu).opArm09B,
	(*Cpu).opArm0B4, (*Cpu).opArm09D, (*Cpu).opArm0B6, (*Cpu).opArm09F,
	(*Cpu).opArm0C0, (*Cpu).opArm0C1, (*Cpu).opArm0C2, (*Cpu).opArm0C3,
	(*Cpu).opArm0C4, (*Cpu).opArm0C5, (*Cpu).opArm0C6, (*Cpu).opArm0C7,
	(*Cpu).opArm0C0, (*Cpu).opArm0C9, (*Cpu).opArm0C2, (*Cpu).opArm0CB,
	(*Cpu).opArm0C4, (*Cpu).opArm0CD, (*Cpu).opArm0C6, (*Cpu).opArm0CF,
	(*Cpu).opArm0D0, (*Cpu).opArm0D1, (*Cpu).opArm0D2, (*Cpu).opArm0D3,
	(*Cpu).opArm0D4, (*Cpu).opArm0D5, (*Cpu).opArm0D6, (*Cpu).opArm0D7,
	(*Cpu).opArm0D0, (*Cpu).opArm0D9, (*Cpu).opArm0D2, (*Cpu).opArm0DB,
	(*Cpu).opArm0D4, (*Cpu).opArm0DD, (*Cpu).opArm0D6, (*Cpu).opArm0DF,
	(*Cpu).opArm0E0, (*Cpu).opArm0E1, (*Cpu).opArm0E2, (*Cpu).opArm0E3,
	(*Cpu).opArm0E4, (*Cpu).opArm0E5, (*Cpu).opArm0E6, (*Cpu).opArm0E7,
	(*Cpu).opArm0E0, (*Cpu).opArm0E9, (*Cpu).opArm0E2, (*Cpu).opArm0CB,
	(*Cpu).opArm0E4, (*Cpu).opArm0CD, (*Cpu).opArm0E6, (*Cpu).opArm0CF,
	(*Cpu).opArm0F0, (*Cpu).opArm0F1, (*Cpu).opArm0F2, (*Cpu).opArm0F3,
	(*Cpu).opArm0F4, (*Cpu).opArm0F5, (*Cpu).opArm0F6, (*Cpu).opArm0F7,
	(*Cpu).opArm0F0, (*Cpu).opArm0F9, (*Cpu).opArm0F2, (*Cpu).opArm0DB,
	(*Cpu).opArm0F4, (*Cpu).opArm0DD, (*Cpu).opArm0F6, (*Cpu).opArm0DF,
	(*Cpu).opArm100, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm109, (*Cpu).opArm101, (*Cpu).opArm10B,
	(*Cpu).opArm101, (*Cpu).opArm10D, (*Cpu).opArm101, (*Cpu).opArm10F,
	(*Cpu).opArm110, (*Cpu).opArm111, (*Cpu).opArm112, (*Cpu).opArm113,
	(*Cpu).opArm114, (*Cpu).opArm115, (*Cpu).opArm116, (*Cpu).opArm117,
	(*Cpu).opArm110, (*Cpu).opArm049, (*Cpu).opArm112, (*Cpu).opArm11B,
	(*Cpu).opArm114, (*Cpu).opArm11D, (*Cpu).opArm116, (*Cpu).opArm11F,
	(*Cpu).opArm120, (*Cpu).opArm121, (*Cpu).opArm101, (*Cpu).opArm123,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm049, (*Cpu).opArm101, (*Cpu).opArm12B,
	(*Cpu).opArm101, (*Cpu).opArm12D, (*Cpu).opArm101, (*Cpu).opArm12F,
	(*Cpu).opArm130, (*Cpu).opArm131, (*Cpu).opArm132, (*Cpu).opArm133,
	(*Cpu).opArm134, (*Cpu).opArm135, (*Cpu).opArm136, (*Cpu).opArm137,
	(*Cpu).opArm130, (*Cpu).opArm049, (*Cpu).opArm132, (*Cpu).opArm13B,
	(*Cpu).opArm134, (*Cpu).opArm13D, (*Cpu).opArm136, (*Cpu).opArm13F,
	(*Cpu).opArm140, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm149, (*Cpu).opArm101, (*Cpu).opArm14B,
	(*Cpu).opArm101, (*Cpu).opArm14D, (*Cpu).opArm101, (*Cpu).opArm14F,
	(*Cpu).opArm150, (*Cpu).opArm151, (*Cpu).opArm152, (*Cpu).opArm153,
	(*Cpu).opArm154, (*Cpu).opArm155, (*Cpu).opArm156, (*Cpu).opArm157,
	(*Cpu).opArm150, (*Cpu).opArm049, (*Cpu).opArm152, (*Cpu).opArm15B,
	(*Cpu).opArm154, (*Cpu).opArm15D, (*Cpu).opArm156, (*Cpu).opArm15F,
	(*Cpu).opArm160, (*Cpu).opArm161, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101, (*Cpu).opArm101,
	(*Cpu).opArm101, (*Cpu).opArm049, (*Cpu).opArm101, (*Cpu).opArm16B,
	(*Cpu).opArm101, (*Cpu).opArm16D, (*Cpu).opArm101, (*Cpu).opArm16F,
	(*Cpu).opArm170, (*Cpu).opArm171, (*Cpu).opArm172, (*Cpu).opArm173,
	(*Cpu).opArm174, (*Cpu).opArm175, (*Cpu).opArm176, (*Cpu).opArm177,
	(*Cpu).opArm170, (*Cpu).opArm049, (*Cpu).opArm172, (*Cpu).opArm17B,
	(*Cpu).opArm174, (*Cpu).opArm17D, (*Cpu).opArm176, (*Cpu).opArm17F,
	(*Cpu).opArm180, (*Cpu).opArm181, (*Cpu).opArm182, (*Cpu).opArm183,
	(*Cpu).opArm184, (*Cpu).opArm185, (*Cpu).opArm186, (*Cpu).opArm187,
	(*Cpu).opArm180, (*Cpu).opArm049, (*Cpu).opArm182, (*Cpu).opArm18B,
	(*Cpu).opArm184, (*Cpu).opArm18D, (*Cpu).opArm186, (*Cpu).opArm18F,
	(*Cpu).opArm190, (*Cpu).opArm191, (*Cpu).opArm192, (*Cpu).opArm193,
	(*Cpu).opArm194, (*Cpu).opArm195, (*Cpu).opArm196, (*Cpu).opArm197,
	(*Cpu).opArm190, (*Cpu).opArm049, (*Cpu).opArm192, (*Cpu).opArm19B,
	(*Cpu).opArm194, (*Cpu).opArm19D, (*Cpu).opArm196, (*Cpu).opArm19F,
	(*Cpu).opArm1A0, (*Cpu).opArm1A1, (*Cpu).opArm1A2, (*Cpu).opArm1A3,
	(*Cpu).opArm1A4, (*Cpu).opArm1A5, (*Cpu).opArm1A6, (*Cpu).opArm1A7,
	(*Cpu).opArm1A0, (*Cpu).opArm049, (*Cpu).opArm1A2, (*Cpu).opArm1AB,
	(*Cpu).opArm1A4, (*Cpu).opArm1AD, (*Cpu).opArm1A6, (*Cpu).opArm1AF,
	(*Cpu).opArm1B0, (*Cpu).opArm1B1, (*Cpu).opArm1B2, (*Cpu).opArm1B3,
	(*Cpu).opArm1B4, (*Cpu).opArm1B5, (*Cpu).opArm1B6, (*Cpu).opArm1B7,
	(*Cpu).opArm1B0, (*Cpu).opArm049, (*Cpu).opArm1B2, (*Cpu).opArm1BB,
	(*Cpu).opArm1B4, (*Cpu).opArm1BD, (*Cpu).opArm1B6, (*Cpu).opArm1BF,
	(*Cpu).opArm1C0, (*Cpu).opArm1C1, (*Cpu).opArm1C2, (*Cpu).opArm1C3,
	(*Cpu).opArm1C4, (*Cpu).opArm1C5, (*Cpu).opArm1C6, (*Cpu).opArm1C7,
	(*Cpu).opArm1C0, (*Cpu).opArm049, (*Cpu).opArm1C2, (*Cpu).opArm1CB,
	(*Cpu).opArm1C4, (*Cpu).opArm1CD, (*Cpu).opArm1C6, (*Cpu).opArm1CF,
	(*Cpu).opArm1D0, (*Cpu).opArm1D1, (*Cpu).opArm1D2, (*Cpu).opArm1D3,
	(*Cpu).opArm1D4, (*Cpu).opArm1D5, (*Cpu).opArm1D6, (*Cpu).opArm1D7,
	(*Cpu).opArm1D0, (*Cpu).opArm049, (*Cpu).opArm1D2, (*Cpu).opArm1DB,
	(*Cpu).opArm1D4, (*Cpu).opArm1DD, (*Cpu).opArm1D6, (*Cpu).opArm1DF,
	(*Cpu).opArm1E0, (*Cpu).opArm1E1, (*Cpu).opArm1E2, (*Cpu).opArm1E3,
	(*Cpu).opArm1E4, (*Cpu).opArm1E5, (*Cpu).opArm1E6, (*Cpu).opArm1E7,
	(*Cpu).opArm1E0, (*Cpu).opArm049, (*Cpu).opArm1E2, (*Cpu).opArm1EB,
	(*Cpu).opArm1E4, (*Cpu).opArm1ED, (*Cpu).opArm1E6, (*Cpu).opArm1EF,
	(*Cpu).opArm1F0, (*Cpu).opArm1F1, (*Cpu).opArm1F2, (*Cpu).opArm1F3,
	(*Cpu).opArm1F4, (*Cpu).opArm1F5, (*Cpu).opArm1F6, (*Cpu).opArm1F7,
	(*Cpu).opArm1F0, (*Cpu).opArm049, (*Cpu).opArm1F2, (*Cpu).opArm1FB,
	(*Cpu).opArm1F4, (*Cpu).opArm1FD, (*Cpu).opArm1F6, (*Cpu).opArm1FF,
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
	(*Cpu).opArm600, (*Cpu).opArm601, (*Cpu).opArm602, (*Cpu).opArm601,
	(*Cpu).opArm604, (*Cpu).opArm601, (*Cpu).opArm606, (*Cpu).opArm601,
	(*Cpu).opArm600, (*Cpu).opArm601, (*Cpu).opArm602, (*Cpu).opArm601,
	(*Cpu).opArm604, (*Cpu).opArm601, (*Cpu).opArm606, (*Cpu).opArm601,
	(*Cpu).opArm610, (*Cpu).opArm601, (*Cpu).opArm612, (*Cpu).opArm601,
	(*Cpu).opArm614, (*Cpu).opArm601, (*Cpu).opArm616, (*Cpu).opArm601,
	(*Cpu).opArm610, (*Cpu).opArm601, (*Cpu).opArm612, (*Cpu).opArm601,
	(*Cpu).opArm614, (*Cpu).opArm601, (*Cpu).opArm616, (*Cpu).opArm601,
	(*Cpu).opArm620, (*Cpu).opArm601, (*Cpu).opArm622, (*Cpu).opArm601,
	(*Cpu).opArm624, (*Cpu).opArm601, (*Cpu).opArm626, (*Cpu).opArm601,
	(*Cpu).opArm620, (*Cpu).opArm601, (*Cpu).opArm622, (*Cpu).opArm601,
	(*Cpu).opArm624, (*Cpu).opArm601, (*Cpu).opArm626, (*Cpu).opArm601,
	(*Cpu).opArm630, (*Cpu).opArm601, (*Cpu).opArm632, (*Cpu).opArm601,
	(*Cpu).opArm634, (*Cpu).opArm601, (*Cpu).opArm636, (*Cpu).opArm601,
	(*Cpu).opArm630, (*Cpu).opArm601, (*Cpu).opArm632, (*Cpu).opArm601,
	(*Cpu).opArm634, (*Cpu).opArm601, (*Cpu).opArm636, (*Cpu).opArm601,
	(*Cpu).opArm640, (*Cpu).opArm601, (*Cpu).opArm642, (*Cpu).opArm601,
	(*Cpu).opArm644, (*Cpu).opArm601, (*Cpu).opArm646, (*Cpu).opArm601,
	(*Cpu).opArm640, (*Cpu).opArm601, (*Cpu).opArm642, (*Cpu).opArm601,
	(*Cpu).opArm644, (*Cpu).opArm601, (*Cpu).opArm646, (*Cpu).opArm601,
	(*Cpu).opArm650, (*Cpu).opArm601, (*Cpu).opArm652, (*Cpu).opArm601,
	(*Cpu).opArm654, (*Cpu).opArm601, (*Cpu).opArm656, (*Cpu).opArm601,
	(*Cpu).opArm650, (*Cpu).opArm601, (*Cpu).opArm652, (*Cpu).opArm601,
	(*Cpu).opArm654, (*Cpu).opArm601, (*Cpu).opArm656, (*Cpu).opArm601,
	(*Cpu).opArm660, (*Cpu).opArm601, (*Cpu).opArm662, (*Cpu).opArm601,
	(*Cpu).opArm664, (*Cpu).opArm601, (*Cpu).opArm666, (*Cpu).opArm601,
	(*Cpu).opArm660, (*Cpu).opArm601, (*Cpu).opArm662, (*Cpu).opArm601,
	(*Cpu).opArm664, (*Cpu).opArm601, (*Cpu).opArm666, (*Cpu).opArm601,
	(*Cpu).opArm670, (*Cpu).opArm601, (*Cpu).opArm672, (*Cpu).opArm601,
	(*Cpu).opArm674, (*Cpu).opArm601, (*Cpu).opArm676, (*Cpu).opArm601,
	(*Cpu).opArm670, (*Cpu).opArm601, (*Cpu).opArm672, (*Cpu).opArm601,
	(*Cpu).opArm674, (*Cpu).opArm601, (*Cpu).opArm676, (*Cpu).opArm601,
	(*Cpu).opArm680, (*Cpu).opArm601, (*Cpu).opArm682, (*Cpu).opArm601,
	(*Cpu).opArm684, (*Cpu).opArm601, (*Cpu).opArm686, (*Cpu).opArm601,
	(*Cpu).opArm680, (*Cpu).opArm601, (*Cpu).opArm682, (*Cpu).opArm601,
	(*Cpu).opArm684, (*Cpu).opArm601, (*Cpu).opArm686, (*Cpu).opArm601,
	(*Cpu).opArm690, (*Cpu).opArm601, (*Cpu).opArm692, (*Cpu).opArm601,
	(*Cpu).opArm694, (*Cpu).opArm601, (*Cpu).opArm696, (*Cpu).opArm601,
	(*Cpu).opArm690, (*Cpu).opArm601, (*Cpu).opArm692, (*Cpu).opArm601,
	(*Cpu).opArm694, (*Cpu).opArm601, (*Cpu).opArm696, (*Cpu).opArm601,
	(*Cpu).opArm6A0, (*Cpu).opArm601, (*Cpu).opArm6A2, (*Cpu).opArm601,
	(*Cpu).opArm6A4, (*Cpu).opArm601, (*Cpu).opArm6A6, (*Cpu).opArm601,
	(*Cpu).opArm6A0, (*Cpu).opArm601, (*Cpu).opArm6A2, (*Cpu).opArm601,
	(*Cpu).opArm6A4, (*Cpu).opArm601, (*Cpu).opArm6A6, (*Cpu).opArm601,
	(*Cpu).opArm6B0, (*Cpu).opArm601, (*Cpu).opArm6B2, (*Cpu).opArm601,
	(*Cpu).opArm6B4, (*Cpu).opArm601, (*Cpu).opArm6B6, (*Cpu).opArm601,
	(*Cpu).opArm6B0, (*Cpu).opArm601, (*Cpu).opArm6B2, (*Cpu).opArm601,
	(*Cpu).opArm6B4, (*Cpu).opArm601, (*Cpu).opArm6B6, (*Cpu).opArm601,
	(*Cpu).opArm6C0, (*Cpu).opArm601, (*Cpu).opArm6C2, (*Cpu).opArm601,
	(*Cpu).opArm6C4, (*Cpu).opArm601, (*Cpu).opArm6C6, (*Cpu).opArm601,
	(*Cpu).opArm6C0, (*Cpu).opArm601, (*Cpu).opArm6C2, (*Cpu).opArm601,
	(*Cpu).opArm6C4, (*Cpu).opArm601, (*Cpu).opArm6C6, (*Cpu).opArm601,
	(*Cpu).opArm6D0, (*Cpu).opArm601, (*Cpu).opArm6D2, (*Cpu).opArm601,
	(*Cpu).opArm6D4, (*Cpu).opArm601, (*Cpu).opArm6D6, (*Cpu).opArm601,
	(*Cpu).opArm6D0, (*Cpu).opArm601, (*Cpu).opArm6D2, (*Cpu).opArm601,
	(*Cpu).opArm6D4, (*Cpu).opArm601, (*Cpu).opArm6D6, (*Cpu).opArm601,
	(*Cpu).opArm6E0, (*Cpu).opArm601, (*Cpu).opArm6E2, (*Cpu).opArm601,
	(*Cpu).opArm6E4, (*Cpu).opArm601, (*Cpu).opArm6E6, (*Cpu).opArm601,
	(*Cpu).opArm6E0, (*Cpu).opArm601, (*Cpu).opArm6E2, (*Cpu).opArm601,
	(*Cpu).opArm6E4, (*Cpu).opArm601, (*Cpu).opArm6E6, (*Cpu).opArm601,
	(*Cpu).opArm6F0, (*Cpu).opArm601, (*Cpu).opArm6F2, (*Cpu).opArm601,
	(*Cpu).opArm6F4, (*Cpu).opArm601, (*Cpu).opArm6F6, (*Cpu).opArm601,
	(*Cpu).opArm6F0, (*Cpu).opArm601, (*Cpu).opArm6F2, (*Cpu).opArm601,
	(*Cpu).opArm6F4, (*Cpu).opArm601, (*Cpu).opArm6F6, (*Cpu).opArm601,
	(*Cpu).opArm700, (*Cpu).opArm601, (*Cpu).opArm702, (*Cpu).opArm601,
	(*Cpu).opArm704, (*Cpu).opArm601, (*Cpu).opArm706, (*Cpu).opArm601,
	(*Cpu).opArm700, (*Cpu).opArm601, (*Cpu).opArm702, (*Cpu).opArm601,
	(*Cpu).opArm704, (*Cpu).opArm601, (*Cpu).opArm706, (*Cpu).opArm601,
	(*Cpu).opArm710, (*Cpu).opArm601, (*Cpu).opArm712, (*Cpu).opArm601,
	(*Cpu).opArm714, (*Cpu).opArm601, (*Cpu).opArm716, (*Cpu).opArm601,
	(*Cpu).opArm710, (*Cpu).opArm601, (*Cpu).opArm712, (*Cpu).opArm601,
	(*Cpu).opArm714, (*Cpu).opArm601, (*Cpu).opArm716, (*Cpu).opArm601,
	(*Cpu).opArm720, (*Cpu).opArm601, (*Cpu).opArm722, (*Cpu).opArm601,
	(*Cpu).opArm724, (*Cpu).opArm601, (*Cpu).opArm726, (*Cpu).opArm601,
	(*Cpu).opArm720, (*Cpu).opArm601, (*Cpu).opArm722, (*Cpu).opArm601,
	(*Cpu).opArm724, (*Cpu).opArm601, (*Cpu).opArm726, (*Cpu).opArm601,
	(*Cpu).opArm730, (*Cpu).opArm601, (*Cpu).opArm732, (*Cpu).opArm601,
	(*Cpu).opArm734, (*Cpu).opArm601, (*Cpu).opArm736, (*Cpu).opArm601,
	(*Cpu).opArm730, (*Cpu).opArm601, (*Cpu).opArm732, (*Cpu).opArm601,
	(*Cpu).opArm734, (*Cpu).opArm601, (*Cpu).opArm736, (*Cpu).opArm601,
	(*Cpu).opArm740, (*Cpu).opArm601, (*Cpu).opArm742, (*Cpu).opArm601,
	(*Cpu).opArm744, (*Cpu).opArm601, (*Cpu).opArm746, (*Cpu).opArm601,
	(*Cpu).opArm740, (*Cpu).opArm601, (*Cpu).opArm742, (*Cpu).opArm601,
	(*Cpu).opArm744, (*Cpu).opArm601, (*Cpu).opArm746, (*Cpu).opArm601,
	(*Cpu).opArm750, (*Cpu).opArm601, (*Cpu).opArm752, (*Cpu).opArm601,
	(*Cpu).opArm754, (*Cpu).opArm601, (*Cpu).opArm756, (*Cpu).opArm601,
	(*Cpu).opArm750, (*Cpu).opArm601, (*Cpu).opArm752, (*Cpu).opArm601,
	(*Cpu).opArm754, (*Cpu).opArm601, (*Cpu).opArm756, (*Cpu).opArm601,
	(*Cpu).opArm760, (*Cpu).opArm601, (*Cpu).opArm762, (*Cpu).opArm601,
	(*Cpu).opArm764, (*Cpu).opArm601, (*Cpu).opArm766, (*Cpu).opArm601,
	(*Cpu).opArm760, (*Cpu).opArm601, (*Cpu).opArm762, (*Cpu).opArm601,
	(*Cpu).opArm764, (*Cpu).opArm601, (*Cpu).opArm766, (*Cpu).opArm601,
	(*Cpu).opArm770, (*Cpu).opArm601, (*Cpu).opArm772, (*Cpu).opArm601,
	(*Cpu).opArm774, (*Cpu).opArm601, (*Cpu).opArm776, (*Cpu).opArm601,
	(*Cpu).opArm770, (*Cpu).opArm601, (*Cpu).opArm772, (*Cpu).opArm601,
	(*Cpu).opArm774, (*Cpu).opArm601, (*Cpu).opArm776, (*Cpu).opArm601,
	(*Cpu).opArm780, (*Cpu).opArm601, (*Cpu).opArm782, (*Cpu).opArm601,
	(*Cpu).opArm784, (*Cpu).opArm601, (*Cpu).opArm786, (*Cpu).opArm601,
	(*Cpu).opArm780, (*Cpu).opArm601, (*Cpu).opArm782, (*Cpu).opArm601,
	(*Cpu).opArm784, (*Cpu).opArm601, (*Cpu).opArm786, (*Cpu).opArm601,
	(*Cpu).opArm790, (*Cpu).opArm601, (*Cpu).opArm792, (*Cpu).opArm601,
	(*Cpu).opArm794, (*Cpu).opArm601, (*Cpu).opArm796, (*Cpu).opArm601,
	(*Cpu).opArm790, (*Cpu).opArm601, (*Cpu).opArm792, (*Cpu).opArm601,
	(*Cpu).opArm794, (*Cpu).opArm601, (*Cpu).opArm796, (*Cpu).opArm601,
	(*Cpu).opArm7A0, (*Cpu).opArm601, (*Cpu).opArm7A2, (*Cpu).opArm601,
	(*Cpu).opArm7A4, (*Cpu).opArm601, (*Cpu).opArm7A6, (*Cpu).opArm601,
	(*Cpu).opArm7A0, (*Cpu).opArm601, (*Cpu).opArm7A2, (*Cpu).opArm601,
	(*Cpu).opArm7A4, (*Cpu).opArm601, (*Cpu).opArm7A6, (*Cpu).opArm601,
	(*Cpu).opArm7B0, (*Cpu).opArm601, (*Cpu).opArm7B2, (*Cpu).opArm601,
	(*Cpu).opArm7B4, (*Cpu).opArm601, (*Cpu).opArm7B6, (*Cpu).opArm601,
	(*Cpu).opArm7B0, (*Cpu).opArm601, (*Cpu).opArm7B2, (*Cpu).opArm601,
	(*Cpu).opArm7B4, (*Cpu).opArm601, (*Cpu).opArm7B6, (*Cpu).opArm601,
	(*Cpu).opArm7C0, (*Cpu).opArm601, (*Cpu).opArm7C2, (*Cpu).opArm601,
	(*Cpu).opArm7C4, (*Cpu).opArm601, (*Cpu).opArm7C6, (*Cpu).opArm601,
	(*Cpu).opArm7C0, (*Cpu).opArm601, (*Cpu).opArm7C2, (*Cpu).opArm601,
	(*Cpu).opArm7C4, (*Cpu).opArm601, (*Cpu).opArm7C6, (*Cpu).opArm601,
	(*Cpu).opArm7D0, (*Cpu).opArm601, (*Cpu).opArm7D2, (*Cpu).opArm601,
	(*Cpu).opArm7D4, (*Cpu).opArm601, (*Cpu).opArm7D6, (*Cpu).opArm601,
	(*Cpu).opArm7D0, (*Cpu).opArm601, (*Cpu).opArm7D2, (*Cpu).opArm601,
	(*Cpu).opArm7D4, (*Cpu).opArm601, (*Cpu).opArm7D6, (*Cpu).opArm601,
	(*Cpu).opArm7E0, (*Cpu).opArm601, (*Cpu).opArm7E2, (*Cpu).opArm601,
	(*Cpu).opArm7E4, (*Cpu).opArm601, (*Cpu).opArm7E6, (*Cpu).opArm601,
	(*Cpu).opArm7E0, (*Cpu).opArm601, (*Cpu).opArm7E2, (*Cpu).opArm601,
	(*Cpu).opArm7E4, (*Cpu).opArm601, (*Cpu).opArm7E6, (*Cpu).opArm601,
	(*Cpu).opArm7F0, (*Cpu).opArm601, (*Cpu).opArm7F2, (*Cpu).opArm601,
	(*Cpu).opArm7F4, (*Cpu).opArm601, (*Cpu).opArm7F6, (*Cpu).opArm601,
	(*Cpu).opArm7F0, (*Cpu).opArm601, (*Cpu).opArm7F2, (*Cpu).opArm601,
	(*Cpu).opArm7F4, (*Cpu).opArm601, (*Cpu).opArm7F6, (*Cpu).opArm601,
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
	(*Cpu).disasmArm000, (*Cpu).disasmArm009, (*Cpu).disasmArm000, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm000, (*Cpu).disasmArm00D, (*Cpu).disasmArm000, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010,
	(*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010, (*Cpu).disasmArm010,
	(*Cpu).disasmArm010, (*Cpu).disasmArm019, (*Cpu).disasmArm010, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm010, (*Cpu).disasmArm01D, (*Cpu).disasmArm010, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020,
	(*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020, (*Cpu).disasmArm020,
	(*Cpu).disasmArm020, (*Cpu).disasmArm029, (*Cpu).disasmArm020, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm020, (*Cpu).disasmArm00D, (*Cpu).disasmArm020, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030,
	(*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030, (*Cpu).disasmArm030,
	(*Cpu).disasmArm030, (*Cpu).disasmArm039, (*Cpu).disasmArm030, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm030, (*Cpu).disasmArm01D, (*Cpu).disasmArm030, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040,
	(*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040, (*Cpu).disasmArm040,
	(*Cpu).disasmArm040, (*Cpu).disasmArm049, (*Cpu).disasmArm040, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm040, (*Cpu).disasmArm00D, (*Cpu).disasmArm040, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050,
	(*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050, (*Cpu).disasmArm050,
	(*Cpu).disasmArm050, (*Cpu).disasmArm049, (*Cpu).disasmArm050, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm050, (*Cpu).disasmArm01D, (*Cpu).disasmArm050, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060,
	(*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060, (*Cpu).disasmArm060,
	(*Cpu).disasmArm060, (*Cpu).disasmArm049, (*Cpu).disasmArm060, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm060, (*Cpu).disasmArm00D, (*Cpu).disasmArm060, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070,
	(*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070, (*Cpu).disasmArm070,
	(*Cpu).disasmArm070, (*Cpu).disasmArm049, (*Cpu).disasmArm070, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm070, (*Cpu).disasmArm01D, (*Cpu).disasmArm070, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080,
	(*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080, (*Cpu).disasmArm080,
	(*Cpu).disasmArm080, (*Cpu).disasmArm089, (*Cpu).disasmArm080, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm080, (*Cpu).disasmArm00D, (*Cpu).disasmArm080, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090,
	(*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090, (*Cpu).disasmArm090,
	(*Cpu).disasmArm090, (*Cpu).disasmArm099, (*Cpu).disasmArm090, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm090, (*Cpu).disasmArm01D, (*Cpu).disasmArm090, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0, (*Cpu).disasmArm0A0,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm0A9, (*Cpu).disasmArm0A0, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm0A0, (*Cpu).disasmArm00D, (*Cpu).disasmArm0A0, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0, (*Cpu).disasmArm0B0,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm0B9, (*Cpu).disasmArm0B0, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm0B0, (*Cpu).disasmArm01D, (*Cpu).disasmArm0B0, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0, (*Cpu).disasmArm0C0,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm0C9, (*Cpu).disasmArm0C0, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm0C0, (*Cpu).disasmArm00D, (*Cpu).disasmArm0C0, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0, (*Cpu).disasmArm0D0,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm0D9, (*Cpu).disasmArm0D0, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm0D0, (*Cpu).disasmArm01D, (*Cpu).disasmArm0D0, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0, (*Cpu).disasmArm0E0,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm0E9, (*Cpu).disasmArm0E0, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm0E0, (*Cpu).disasmArm00D, (*Cpu).disasmArm0E0, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0, (*Cpu).disasmArm0F0,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm0F9, (*Cpu).disasmArm0F0, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm0F0, (*Cpu).disasmArm01D, (*Cpu).disasmArm0F0, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm100, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm109, (*Cpu).disasmArm049, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm049, (*Cpu).disasmArm00D, (*Cpu).disasmArm049, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110,
	(*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110, (*Cpu).disasmArm110,
	(*Cpu).disasmArm110, (*Cpu).disasmArm049, (*Cpu).disasmArm110, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm110, (*Cpu).disasmArm01D, (*Cpu).disasmArm110, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm120, (*Cpu).disasmArm121, (*Cpu).disasmArm049, (*Cpu).disasmArm123,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm049, (*Cpu).disasmArm00D, (*Cpu).disasmArm049, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130,
	(*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130, (*Cpu).disasmArm130,
	(*Cpu).disasmArm130, (*Cpu).disasmArm049, (*Cpu).disasmArm130, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm130, (*Cpu).disasmArm01D, (*Cpu).disasmArm130, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm140, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm149, (*Cpu).disasmArm049, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm049, (*Cpu).disasmArm00D, (*Cpu).disasmArm049, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150,
	(*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150, (*Cpu).disasmArm150,
	(*Cpu).disasmArm150, (*Cpu).disasmArm049, (*Cpu).disasmArm150, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm150, (*Cpu).disasmArm01D, (*Cpu).disasmArm150, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm160, (*Cpu).disasmArm161, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049,
	(*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm049, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm049, (*Cpu).disasmArm00D, (*Cpu).disasmArm049, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170,
	(*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170, (*Cpu).disasmArm170,
	(*Cpu).disasmArm170, (*Cpu).disasmArm049, (*Cpu).disasmArm170, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm170, (*Cpu).disasmArm01D, (*Cpu).disasmArm170, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180,
	(*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180, (*Cpu).disasmArm180,
	(*Cpu).disasmArm180, (*Cpu).disasmArm049, (*Cpu).disasmArm180, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm180, (*Cpu).disasmArm00D, (*Cpu).disasmArm180, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190,
	(*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190, (*Cpu).disasmArm190,
	(*Cpu).disasmArm190, (*Cpu).disasmArm049, (*Cpu).disasmArm190, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm190, (*Cpu).disasmArm01D, (*Cpu).disasmArm190, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0, (*Cpu).disasmArm1A0,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm049, (*Cpu).disasmArm1A0, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm1A0, (*Cpu).disasmArm00D, (*Cpu).disasmArm1A0, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0, (*Cpu).disasmArm1B0,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm049, (*Cpu).disasmArm1B0, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm1B0, (*Cpu).disasmArm01D, (*Cpu).disasmArm1B0, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0, (*Cpu).disasmArm1C0,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm049, (*Cpu).disasmArm1C0, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm1C0, (*Cpu).disasmArm00D, (*Cpu).disasmArm1C0, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0, (*Cpu).disasmArm1D0,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm049, (*Cpu).disasmArm1D0, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm1D0, (*Cpu).disasmArm01D, (*Cpu).disasmArm1D0, (*Cpu).disasmArm01F,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0, (*Cpu).disasmArm1E0,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm049, (*Cpu).disasmArm1E0, (*Cpu).disasmArm00B,
	(*Cpu).disasmArm1E0, (*Cpu).disasmArm00D, (*Cpu).disasmArm1E0, (*Cpu).disasmArm00F,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0, (*Cpu).disasmArm1F0,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm049, (*Cpu).disasmArm1F0, (*Cpu).disasmArm01B,
	(*Cpu).disasmArm1F0, (*Cpu).disasmArm01D, (*Cpu).disasmArm1F0, (*Cpu).disasmArm01F,
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
	(*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320,
	(*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320,
	(*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320,
	(*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320, (*Cpu).disasmArm320,
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
	(*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360,
	(*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360,
	(*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360,
	(*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360, (*Cpu).disasmArm360,
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
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm600, (*Cpu).disasmArm049, (*Cpu).disasmArm600, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm610, (*Cpu).disasmArm049, (*Cpu).disasmArm610, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm640, (*Cpu).disasmArm049, (*Cpu).disasmArm640, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm650, (*Cpu).disasmArm049, (*Cpu).disasmArm650, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm680, (*Cpu).disasmArm049, (*Cpu).disasmArm680, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm690, (*Cpu).disasmArm049, (*Cpu).disasmArm690, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6C0, (*Cpu).disasmArm049, (*Cpu).disasmArm6C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm6D0, (*Cpu).disasmArm049, (*Cpu).disasmArm6D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm700, (*Cpu).disasmArm049, (*Cpu).disasmArm700, (*Cpu).disasmArm049,
	(*Cpu).disasmArm700, (*Cpu).disasmArm049, (*Cpu).disasmArm700, (*Cpu).disasmArm049,
	(*Cpu).disasmArm700, (*Cpu).disasmArm049, (*Cpu).disasmArm700, (*Cpu).disasmArm049,
	(*Cpu).disasmArm700, (*Cpu).disasmArm049, (*Cpu).disasmArm700, (*Cpu).disasmArm049,
	(*Cpu).disasmArm710, (*Cpu).disasmArm049, (*Cpu).disasmArm710, (*Cpu).disasmArm049,
	(*Cpu).disasmArm710, (*Cpu).disasmArm049, (*Cpu).disasmArm710, (*Cpu).disasmArm049,
	(*Cpu).disasmArm710, (*Cpu).disasmArm049, (*Cpu).disasmArm710, (*Cpu).disasmArm049,
	(*Cpu).disasmArm710, (*Cpu).disasmArm049, (*Cpu).disasmArm710, (*Cpu).disasmArm049,
	(*Cpu).disasmArm720, (*Cpu).disasmArm049, (*Cpu).disasmArm720, (*Cpu).disasmArm049,
	(*Cpu).disasmArm720, (*Cpu).disasmArm049, (*Cpu).disasmArm720, (*Cpu).disasmArm049,
	(*Cpu).disasmArm720, (*Cpu).disasmArm049, (*Cpu).disasmArm720, (*Cpu).disasmArm049,
	(*Cpu).disasmArm720, (*Cpu).disasmArm049, (*Cpu).disasmArm720, (*Cpu).disasmArm049,
	(*Cpu).disasmArm730, (*Cpu).disasmArm049, (*Cpu).disasmArm730, (*Cpu).disasmArm049,
	(*Cpu).disasmArm730, (*Cpu).disasmArm049, (*Cpu).disasmArm730, (*Cpu).disasmArm049,
	(*Cpu).disasmArm730, (*Cpu).disasmArm049, (*Cpu).disasmArm730, (*Cpu).disasmArm049,
	(*Cpu).disasmArm730, (*Cpu).disasmArm049, (*Cpu).disasmArm730, (*Cpu).disasmArm049,
	(*Cpu).disasmArm740, (*Cpu).disasmArm049, (*Cpu).disasmArm740, (*Cpu).disasmArm049,
	(*Cpu).disasmArm740, (*Cpu).disasmArm049, (*Cpu).disasmArm740, (*Cpu).disasmArm049,
	(*Cpu).disasmArm740, (*Cpu).disasmArm049, (*Cpu).disasmArm740, (*Cpu).disasmArm049,
	(*Cpu).disasmArm740, (*Cpu).disasmArm049, (*Cpu).disasmArm740, (*Cpu).disasmArm049,
	(*Cpu).disasmArm750, (*Cpu).disasmArm049, (*Cpu).disasmArm750, (*Cpu).disasmArm049,
	(*Cpu).disasmArm750, (*Cpu).disasmArm049, (*Cpu).disasmArm750, (*Cpu).disasmArm049,
	(*Cpu).disasmArm750, (*Cpu).disasmArm049, (*Cpu).disasmArm750, (*Cpu).disasmArm049,
	(*Cpu).disasmArm750, (*Cpu).disasmArm049, (*Cpu).disasmArm750, (*Cpu).disasmArm049,
	(*Cpu).disasmArm760, (*Cpu).disasmArm049, (*Cpu).disasmArm760, (*Cpu).disasmArm049,
	(*Cpu).disasmArm760, (*Cpu).disasmArm049, (*Cpu).disasmArm760, (*Cpu).disasmArm049,
	(*Cpu).disasmArm760, (*Cpu).disasmArm049, (*Cpu).disasmArm760, (*Cpu).disasmArm049,
	(*Cpu).disasmArm760, (*Cpu).disasmArm049, (*Cpu).disasmArm760, (*Cpu).disasmArm049,
	(*Cpu).disasmArm770, (*Cpu).disasmArm049, (*Cpu).disasmArm770, (*Cpu).disasmArm049,
	(*Cpu).disasmArm770, (*Cpu).disasmArm049, (*Cpu).disasmArm770, (*Cpu).disasmArm049,
	(*Cpu).disasmArm770, (*Cpu).disasmArm049, (*Cpu).disasmArm770, (*Cpu).disasmArm049,
	(*Cpu).disasmArm770, (*Cpu).disasmArm049, (*Cpu).disasmArm770, (*Cpu).disasmArm049,
	(*Cpu).disasmArm780, (*Cpu).disasmArm049, (*Cpu).disasmArm780, (*Cpu).disasmArm049,
	(*Cpu).disasmArm780, (*Cpu).disasmArm049, (*Cpu).disasmArm780, (*Cpu).disasmArm049,
	(*Cpu).disasmArm780, (*Cpu).disasmArm049, (*Cpu).disasmArm780, (*Cpu).disasmArm049,
	(*Cpu).disasmArm780, (*Cpu).disasmArm049, (*Cpu).disasmArm780, (*Cpu).disasmArm049,
	(*Cpu).disasmArm790, (*Cpu).disasmArm049, (*Cpu).disasmArm790, (*Cpu).disasmArm049,
	(*Cpu).disasmArm790, (*Cpu).disasmArm049, (*Cpu).disasmArm790, (*Cpu).disasmArm049,
	(*Cpu).disasmArm790, (*Cpu).disasmArm049, (*Cpu).disasmArm790, (*Cpu).disasmArm049,
	(*Cpu).disasmArm790, (*Cpu).disasmArm049, (*Cpu).disasmArm790, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm049, (*Cpu).disasmArm7A0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm049, (*Cpu).disasmArm7A0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm049, (*Cpu).disasmArm7A0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7A0, (*Cpu).disasmArm049, (*Cpu).disasmArm7A0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm049, (*Cpu).disasmArm7B0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm049, (*Cpu).disasmArm7B0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm049, (*Cpu).disasmArm7B0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7B0, (*Cpu).disasmArm049, (*Cpu).disasmArm7B0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm049, (*Cpu).disasmArm7C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm049, (*Cpu).disasmArm7C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm049, (*Cpu).disasmArm7C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7C0, (*Cpu).disasmArm049, (*Cpu).disasmArm7C0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm049, (*Cpu).disasmArm7D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm049, (*Cpu).disasmArm7D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm049, (*Cpu).disasmArm7D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7D0, (*Cpu).disasmArm049, (*Cpu).disasmArm7D0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm049, (*Cpu).disasmArm7E0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm049, (*Cpu).disasmArm7E0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm049, (*Cpu).disasmArm7E0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7E0, (*Cpu).disasmArm049, (*Cpu).disasmArm7E0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm049, (*Cpu).disasmArm7F0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm049, (*Cpu).disasmArm7F0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm049, (*Cpu).disasmArm7F0, (*Cpu).disasmArm049,
	(*Cpu).disasmArm7F0, (*Cpu).disasmArm049, (*Cpu).disasmArm7F0, (*Cpu).disasmArm049,
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
