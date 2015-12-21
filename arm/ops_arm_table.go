// Generated on 2015-12-21 02:48:57.541764887 +0100 CET
package arm

import "bytes"
import "strconv"

var opArmTable = [256]func(*Cpu, uint32){
	(*Cpu).opArm00,
	(*Cpu).opArm01,
	(*Cpu).opArm02,
	(*Cpu).opArm03,
	(*Cpu).opArm04,
	(*Cpu).opArm05,
	(*Cpu).opArm06,
	(*Cpu).opArm07,
	(*Cpu).opArm08,
	(*Cpu).opArm09,
	(*Cpu).opArm0A,
	(*Cpu).opArm0B,
	(*Cpu).opArm0C,
	(*Cpu).opArm0D,
	(*Cpu).opArm0E,
	(*Cpu).opArm0F,
	(*Cpu).opArm10,
	(*Cpu).opArm11,
	(*Cpu).opArm12,
	(*Cpu).opArm13,
	(*Cpu).opArm14,
	(*Cpu).opArm15,
	(*Cpu).opArm16,
	(*Cpu).opArm17,
	(*Cpu).opArm18,
	(*Cpu).opArm19,
	(*Cpu).opArm1A,
	(*Cpu).opArm1B,
	(*Cpu).opArm1C,
	(*Cpu).opArm1D,
	(*Cpu).opArm1E,
	(*Cpu).opArm1F,
	(*Cpu).opArm20,
	(*Cpu).opArm21,
	(*Cpu).opArm22,
	(*Cpu).opArm23,
	(*Cpu).opArm24,
	(*Cpu).opArm25,
	(*Cpu).opArm26,
	(*Cpu).opArm27,
	(*Cpu).opArm28,
	(*Cpu).opArm29,
	(*Cpu).opArm2A,
	(*Cpu).opArm2B,
	(*Cpu).opArm2C,
	(*Cpu).opArm2D,
	(*Cpu).opArm2E,
	(*Cpu).opArm2F,
	(*Cpu).opArm30,
	(*Cpu).opArm31,
	(*Cpu).opArm32,
	(*Cpu).opArm33,
	(*Cpu).opArm34,
	(*Cpu).opArm35,
	(*Cpu).opArm36,
	(*Cpu).opArm37,
	(*Cpu).opArm38,
	(*Cpu).opArm39,
	(*Cpu).opArm3A,
	(*Cpu).opArm3B,
	(*Cpu).opArm3C,
	(*Cpu).opArm3D,
	(*Cpu).opArm3E,
	(*Cpu).opArm3F,
	(*Cpu).opArm40,
	(*Cpu).opArm41,
	(*Cpu).opArm42,
	(*Cpu).opArm43,
	(*Cpu).opArm44,
	(*Cpu).opArm45,
	(*Cpu).opArm46,
	(*Cpu).opArm47,
	(*Cpu).opArm48,
	(*Cpu).opArm49,
	(*Cpu).opArm4A,
	(*Cpu).opArm4B,
	(*Cpu).opArm4C,
	(*Cpu).opArm4D,
	(*Cpu).opArm4E,
	(*Cpu).opArm4F,
	(*Cpu).opArm50,
	(*Cpu).opArm51,
	(*Cpu).opArm52,
	(*Cpu).opArm53,
	(*Cpu).opArm54,
	(*Cpu).opArm55,
	(*Cpu).opArm56,
	(*Cpu).opArm57,
	(*Cpu).opArm58,
	(*Cpu).opArm59,
	(*Cpu).opArm5A,
	(*Cpu).opArm5B,
	(*Cpu).opArm5C,
	(*Cpu).opArm5D,
	(*Cpu).opArm5E,
	(*Cpu).opArm5F,
	(*Cpu).opArm60,
	(*Cpu).opArm61,
	(*Cpu).opArm62,
	(*Cpu).opArm63,
	(*Cpu).opArm64,
	(*Cpu).opArm65,
	(*Cpu).opArm66,
	(*Cpu).opArm67,
	(*Cpu).opArm68,
	(*Cpu).opArm69,
	(*Cpu).opArm6A,
	(*Cpu).opArm6B,
	(*Cpu).opArm6C,
	(*Cpu).opArm6D,
	(*Cpu).opArm6E,
	(*Cpu).opArm6F,
	(*Cpu).opArm70,
	(*Cpu).opArm71,
	(*Cpu).opArm72,
	(*Cpu).opArm73,
	(*Cpu).opArm74,
	(*Cpu).opArm75,
	(*Cpu).opArm76,
	(*Cpu).opArm77,
	(*Cpu).opArm78,
	(*Cpu).opArm79,
	(*Cpu).opArm7A,
	(*Cpu).opArm7B,
	(*Cpu).opArm7C,
	(*Cpu).opArm7D,
	(*Cpu).opArm7E,
	(*Cpu).opArm7F,
	(*Cpu).opArm80,
	(*Cpu).opArm81,
	(*Cpu).opArm82,
	(*Cpu).opArm83,
	(*Cpu).opArm84,
	(*Cpu).opArm85,
	(*Cpu).opArm86,
	(*Cpu).opArm87,
	(*Cpu).opArm88,
	(*Cpu).opArm89,
	(*Cpu).opArm8A,
	(*Cpu).opArm8B,
	(*Cpu).opArm8C,
	(*Cpu).opArm8D,
	(*Cpu).opArm8E,
	(*Cpu).opArm8F,
	(*Cpu).opArm90,
	(*Cpu).opArm91,
	(*Cpu).opArm92,
	(*Cpu).opArm93,
	(*Cpu).opArm94,
	(*Cpu).opArm95,
	(*Cpu).opArm96,
	(*Cpu).opArm97,
	(*Cpu).opArm98,
	(*Cpu).opArm99,
	(*Cpu).opArm9A,
	(*Cpu).opArm9B,
	(*Cpu).opArm9C,
	(*Cpu).opArm9D,
	(*Cpu).opArm9E,
	(*Cpu).opArm9F,
	(*Cpu).opArmA0,
	(*Cpu).opArmA1,
	(*Cpu).opArmA2,
	(*Cpu).opArmA3,
	(*Cpu).opArmA4,
	(*Cpu).opArmA5,
	(*Cpu).opArmA6,
	(*Cpu).opArmA7,
	(*Cpu).opArmA8,
	(*Cpu).opArmA9,
	(*Cpu).opArmAA,
	(*Cpu).opArmAB,
	(*Cpu).opArmAC,
	(*Cpu).opArmAD,
	(*Cpu).opArmAE,
	(*Cpu).opArmAF,
	(*Cpu).opArmB0,
	(*Cpu).opArmB1,
	(*Cpu).opArmB2,
	(*Cpu).opArmB3,
	(*Cpu).opArmB4,
	(*Cpu).opArmB5,
	(*Cpu).opArmB6,
	(*Cpu).opArmB7,
	(*Cpu).opArmB8,
	(*Cpu).opArmB9,
	(*Cpu).opArmBA,
	(*Cpu).opArmBB,
	(*Cpu).opArmBC,
	(*Cpu).opArmBD,
	(*Cpu).opArmBE,
	(*Cpu).opArmBF,
	(*Cpu).opArmC0,
	(*Cpu).opArmC1,
	(*Cpu).opArmC2,
	(*Cpu).opArmC3,
	(*Cpu).opArmC4,
	(*Cpu).opArmC5,
	(*Cpu).opArmC6,
	(*Cpu).opArmC7,
	(*Cpu).opArmC8,
	(*Cpu).opArmC9,
	(*Cpu).opArmCA,
	(*Cpu).opArmCB,
	(*Cpu).opArmCC,
	(*Cpu).opArmCD,
	(*Cpu).opArmCE,
	(*Cpu).opArmCF,
	(*Cpu).opArmD0,
	(*Cpu).opArmD1,
	(*Cpu).opArmD2,
	(*Cpu).opArmD3,
	(*Cpu).opArmD4,
	(*Cpu).opArmD5,
	(*Cpu).opArmD6,
	(*Cpu).opArmD7,
	(*Cpu).opArmD8,
	(*Cpu).opArmD9,
	(*Cpu).opArmDA,
	(*Cpu).opArmDB,
	(*Cpu).opArmDC,
	(*Cpu).opArmDD,
	(*Cpu).opArmDE,
	(*Cpu).opArmDF,
	(*Cpu).opArmE0,
	(*Cpu).opArmE1,
	(*Cpu).opArmE2,
	(*Cpu).opArmE3,
	(*Cpu).opArmE4,
	(*Cpu).opArmE5,
	(*Cpu).opArmE6,
	(*Cpu).opArmE7,
	(*Cpu).opArmE8,
	(*Cpu).opArmE9,
	(*Cpu).opArmEA,
	(*Cpu).opArmEB,
	(*Cpu).opArmEC,
	(*Cpu).opArmED,
	(*Cpu).opArmEE,
	(*Cpu).opArmEF,
	(*Cpu).opArmF0,
	(*Cpu).opArmF1,
	(*Cpu).opArmF2,
	(*Cpu).opArmF3,
	(*Cpu).opArmF4,
	(*Cpu).opArmF5,
	(*Cpu).opArmF6,
	(*Cpu).opArmF7,
	(*Cpu).opArmF8,
	(*Cpu).opArmF9,
	(*Cpu).opArmFA,
	(*Cpu).opArmFB,
	(*Cpu).opArmFC,
	(*Cpu).opArmFD,
	(*Cpu).opArmFE,
	(*Cpu).opArmFF,
}
var disasmArmTable = [256]func(*Cpu, uint32, uint32) string{
	(*Cpu).disasmArm00,
	(*Cpu).disasmArm01,
	(*Cpu).disasmArm02,
	(*Cpu).disasmArm03,
	(*Cpu).disasmArm04,
	(*Cpu).disasmArm05,
	(*Cpu).disasmArm06,
	(*Cpu).disasmArm07,
	(*Cpu).disasmArm08,
	(*Cpu).disasmArm09,
	(*Cpu).disasmArm0A,
	(*Cpu).disasmArm0B,
	(*Cpu).disasmArm0C,
	(*Cpu).disasmArm0D,
	(*Cpu).disasmArm0E,
	(*Cpu).disasmArm0F,
	(*Cpu).disasmArm10,
	(*Cpu).disasmArm11,
	(*Cpu).disasmArm12,
	(*Cpu).disasmArm13,
	(*Cpu).disasmArm14,
	(*Cpu).disasmArm15,
	(*Cpu).disasmArm16,
	(*Cpu).disasmArm17,
	(*Cpu).disasmArm18,
	(*Cpu).disasmArm19,
	(*Cpu).disasmArm1A,
	(*Cpu).disasmArm1B,
	(*Cpu).disasmArm1C,
	(*Cpu).disasmArm1D,
	(*Cpu).disasmArm1E,
	(*Cpu).disasmArm1F,
	(*Cpu).disasmArm20,
	(*Cpu).disasmArm21,
	(*Cpu).disasmArm22,
	(*Cpu).disasmArm23,
	(*Cpu).disasmArm24,
	(*Cpu).disasmArm25,
	(*Cpu).disasmArm26,
	(*Cpu).disasmArm27,
	(*Cpu).disasmArm28,
	(*Cpu).disasmArm29,
	(*Cpu).disasmArm2A,
	(*Cpu).disasmArm2B,
	(*Cpu).disasmArm2C,
	(*Cpu).disasmArm2D,
	(*Cpu).disasmArm2E,
	(*Cpu).disasmArm2F,
	(*Cpu).disasmArm30,
	(*Cpu).disasmArm31,
	(*Cpu).disasmArm32,
	(*Cpu).disasmArm33,
	(*Cpu).disasmArm34,
	(*Cpu).disasmArm35,
	(*Cpu).disasmArm36,
	(*Cpu).disasmArm37,
	(*Cpu).disasmArm38,
	(*Cpu).disasmArm39,
	(*Cpu).disasmArm3A,
	(*Cpu).disasmArm3B,
	(*Cpu).disasmArm3C,
	(*Cpu).disasmArm3D,
	(*Cpu).disasmArm3E,
	(*Cpu).disasmArm3F,
	(*Cpu).disasmArm40,
	(*Cpu).disasmArm41,
	(*Cpu).disasmArm42,
	(*Cpu).disasmArm43,
	(*Cpu).disasmArm44,
	(*Cpu).disasmArm45,
	(*Cpu).disasmArm46,
	(*Cpu).disasmArm47,
	(*Cpu).disasmArm48,
	(*Cpu).disasmArm49,
	(*Cpu).disasmArm4A,
	(*Cpu).disasmArm4B,
	(*Cpu).disasmArm4C,
	(*Cpu).disasmArm4D,
	(*Cpu).disasmArm4E,
	(*Cpu).disasmArm4F,
	(*Cpu).disasmArm50,
	(*Cpu).disasmArm51,
	(*Cpu).disasmArm52,
	(*Cpu).disasmArm53,
	(*Cpu).disasmArm54,
	(*Cpu).disasmArm55,
	(*Cpu).disasmArm56,
	(*Cpu).disasmArm57,
	(*Cpu).disasmArm58,
	(*Cpu).disasmArm59,
	(*Cpu).disasmArm5A,
	(*Cpu).disasmArm5B,
	(*Cpu).disasmArm5C,
	(*Cpu).disasmArm5D,
	(*Cpu).disasmArm5E,
	(*Cpu).disasmArm5F,
	(*Cpu).disasmArm60,
	(*Cpu).disasmArm61,
	(*Cpu).disasmArm62,
	(*Cpu).disasmArm63,
	(*Cpu).disasmArm64,
	(*Cpu).disasmArm65,
	(*Cpu).disasmArm66,
	(*Cpu).disasmArm67,
	(*Cpu).disasmArm68,
	(*Cpu).disasmArm69,
	(*Cpu).disasmArm6A,
	(*Cpu).disasmArm6B,
	(*Cpu).disasmArm6C,
	(*Cpu).disasmArm6D,
	(*Cpu).disasmArm6E,
	(*Cpu).disasmArm6F,
	(*Cpu).disasmArm70,
	(*Cpu).disasmArm71,
	(*Cpu).disasmArm72,
	(*Cpu).disasmArm73,
	(*Cpu).disasmArm74,
	(*Cpu).disasmArm75,
	(*Cpu).disasmArm76,
	(*Cpu).disasmArm77,
	(*Cpu).disasmArm78,
	(*Cpu).disasmArm79,
	(*Cpu).disasmArm7A,
	(*Cpu).disasmArm7B,
	(*Cpu).disasmArm7C,
	(*Cpu).disasmArm7D,
	(*Cpu).disasmArm7E,
	(*Cpu).disasmArm7F,
	(*Cpu).disasmArm80,
	(*Cpu).disasmArm81,
	(*Cpu).disasmArm82,
	(*Cpu).disasmArm83,
	(*Cpu).disasmArm84,
	(*Cpu).disasmArm85,
	(*Cpu).disasmArm86,
	(*Cpu).disasmArm87,
	(*Cpu).disasmArm88,
	(*Cpu).disasmArm89,
	(*Cpu).disasmArm8A,
	(*Cpu).disasmArm8B,
	(*Cpu).disasmArm8C,
	(*Cpu).disasmArm8D,
	(*Cpu).disasmArm8E,
	(*Cpu).disasmArm8F,
	(*Cpu).disasmArm90,
	(*Cpu).disasmArm91,
	(*Cpu).disasmArm92,
	(*Cpu).disasmArm93,
	(*Cpu).disasmArm94,
	(*Cpu).disasmArm95,
	(*Cpu).disasmArm96,
	(*Cpu).disasmArm97,
	(*Cpu).disasmArm98,
	(*Cpu).disasmArm99,
	(*Cpu).disasmArm9A,
	(*Cpu).disasmArm9B,
	(*Cpu).disasmArm9C,
	(*Cpu).disasmArm9D,
	(*Cpu).disasmArm9E,
	(*Cpu).disasmArm9F,
	(*Cpu).disasmArmA0,
	(*Cpu).disasmArmA1,
	(*Cpu).disasmArmA2,
	(*Cpu).disasmArmA3,
	(*Cpu).disasmArmA4,
	(*Cpu).disasmArmA5,
	(*Cpu).disasmArmA6,
	(*Cpu).disasmArmA7,
	(*Cpu).disasmArmA8,
	(*Cpu).disasmArmA9,
	(*Cpu).disasmArmAA,
	(*Cpu).disasmArmAB,
	(*Cpu).disasmArmAC,
	(*Cpu).disasmArmAD,
	(*Cpu).disasmArmAE,
	(*Cpu).disasmArmAF,
	(*Cpu).disasmArmB0,
	(*Cpu).disasmArmB1,
	(*Cpu).disasmArmB2,
	(*Cpu).disasmArmB3,
	(*Cpu).disasmArmB4,
	(*Cpu).disasmArmB5,
	(*Cpu).disasmArmB6,
	(*Cpu).disasmArmB7,
	(*Cpu).disasmArmB8,
	(*Cpu).disasmArmB9,
	(*Cpu).disasmArmBA,
	(*Cpu).disasmArmBB,
	(*Cpu).disasmArmBC,
	(*Cpu).disasmArmBD,
	(*Cpu).disasmArmBE,
	(*Cpu).disasmArmBF,
	(*Cpu).disasmArmC0,
	(*Cpu).disasmArmC1,
	(*Cpu).disasmArmC2,
	(*Cpu).disasmArmC3,
	(*Cpu).disasmArmC4,
	(*Cpu).disasmArmC5,
	(*Cpu).disasmArmC6,
	(*Cpu).disasmArmC7,
	(*Cpu).disasmArmC8,
	(*Cpu).disasmArmC9,
	(*Cpu).disasmArmCA,
	(*Cpu).disasmArmCB,
	(*Cpu).disasmArmCC,
	(*Cpu).disasmArmCD,
	(*Cpu).disasmArmCE,
	(*Cpu).disasmArmCF,
	(*Cpu).disasmArmD0,
	(*Cpu).disasmArmD1,
	(*Cpu).disasmArmD2,
	(*Cpu).disasmArmD3,
	(*Cpu).disasmArmD4,
	(*Cpu).disasmArmD5,
	(*Cpu).disasmArmD6,
	(*Cpu).disasmArmD7,
	(*Cpu).disasmArmD8,
	(*Cpu).disasmArmD9,
	(*Cpu).disasmArmDA,
	(*Cpu).disasmArmDB,
	(*Cpu).disasmArmDC,
	(*Cpu).disasmArmDD,
	(*Cpu).disasmArmDE,
	(*Cpu).disasmArmDF,
	(*Cpu).disasmArmE0,
	(*Cpu).disasmArmE1,
	(*Cpu).disasmArmE2,
	(*Cpu).disasmArmE3,
	(*Cpu).disasmArmE4,
	(*Cpu).disasmArmE5,
	(*Cpu).disasmArmE6,
	(*Cpu).disasmArmE7,
	(*Cpu).disasmArmE8,
	(*Cpu).disasmArmE9,
	(*Cpu).disasmArmEA,
	(*Cpu).disasmArmEB,
	(*Cpu).disasmArmEC,
	(*Cpu).disasmArmED,
	(*Cpu).disasmArmEE,
	(*Cpu).disasmArmEF,
	(*Cpu).disasmArmF0,
	(*Cpu).disasmArmF1,
	(*Cpu).disasmArmF2,
	(*Cpu).disasmArmF3,
	(*Cpu).disasmArmF4,
	(*Cpu).disasmArmF5,
	(*Cpu).disasmArmF6,
	(*Cpu).disasmArmF7,
	(*Cpu).disasmArmF8,
	(*Cpu).disasmArmF9,
	(*Cpu).disasmArmFA,
	(*Cpu).disasmArmFB,
	(*Cpu).disasmArmFC,
	(*Cpu).disasmArmFD,
	(*Cpu).disasmArmFE,
	(*Cpu).disasmArmFF,
}

func (cpu *Cpu) opArm00(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm00(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm01(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm01(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm02(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm02(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm03(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm03(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm04(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm04(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm05(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm05(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm06(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm06(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm07(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn -= off
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm07(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm08(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm08(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm09(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm09(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm0A(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm0A(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm0B(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm0B(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm0C(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm0C(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm0D(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm0D(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm0E(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm0E(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm0F(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			rn += off
			cpu.Regs[rnx] = reg(rn)
		} else {
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
	} else {
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
}

func (cpu *Cpu) disasmArm0F(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
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
}

func (cpu *Cpu) opArm10(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
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
	} else {
		// MRS
		if !cpu.opArmCond(op) {
			return
		}
		mask := (op >> 16) & 0xF
		if mask != 0xF {
			cpu.InvalidOpArm(op, "unimplemented SWP")
			return
		}
		rdx := (op >> 12) & 0xF
		if rdx == 15 {
			cpu.InvalidOpArm(op, "write to PC in MRS")
			return
		}
		cpu.Regs[rdx] = reg(cpu.Cpsr.Uint32())
	}
}

func (cpu *Cpu) disasmArm10(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
		return "dw " + strconv.FormatInt(int64(op), 16)
	}
}

func (cpu *Cpu) opArm11(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm11(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm12(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		if (op>>4)&0xFF != 0 { // BX / BLX_reg
			if !cpu.opArmCond(op) {
				return
			}
			rnx := op & 0xF
			rn := cpu.Regs[rnx]
			if op&0x20 != 0 {
				cpu.Regs[14] = cpu.Regs[15] - 4
			}
			if rn&1 != 0 {
				cpu.Cpsr.SetT(true)
				rn &^= 1
			} else {
				rn &^= 3
			}
			cpu.pc = rn
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
		if (op>>4)&0xFF != 0 {
			cpu.InvalidOpArm(op, "unimplemented BX")
			return
		}
		rmx := op & 0xF
		val := uint32(cpu.Regs[rmx])
		cpu.Cpsr.SetWithMask(val, mask, cpu)
	}
}

func (cpu *Cpu) disasmArm12(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
		if op&0x20 != 0 {
			var out bytes.Buffer
			opcode := cpu.disasmAddCond("blx", op)
			out.WriteString((opcode + "                ")[:10])
			arg0 := op & 0xF
			out.WriteString(RegNames[arg0])
			return out.String()
		} else {
			var out bytes.Buffer
			opcode := cpu.disasmAddCond("bx", op)
			out.WriteString((opcode + "                ")[:10])
			arg0 := op & 0xF
			out.WriteString(RegNames[arg0])
			return out.String()
		}
		return "dw " + strconv.FormatInt(int64(op), 16)
	}
}

func (cpu *Cpu) opArm13(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm13(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm14(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn -= off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
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
	} else {
		// MRS
		if !cpu.opArmCond(op) {
			return
		}
		mask := (op >> 16) & 0xF
		if mask != 0xF {
			cpu.InvalidOpArm(op, "unimplemented SWP")
			return
		}
		rdx := (op >> 12) & 0xF
		if rdx == 15 {
			cpu.InvalidOpArm(op, "write to PC in MRS")
			return
		}
		cpu.Regs[rdx] = reg(*cpu.RegSpsr())
	}
}

func (cpu *Cpu) disasmArm14(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
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
	} else {
		return "dw " + strconv.FormatInt(int64(op), 16)
	}
}

func (cpu *Cpu) opArm15(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn -= off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm15(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm16(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn -= off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
		if (op>>4)&0xFF != 0 {
			cpu.InvalidOpArm(op, "unimplemented BX")
			return
		}
		rmx := op & 0xF
		val := uint32(cpu.Regs[rmx])
		cpu.RegSpsr().SetWithMask(val, mask)
	}
}

func (cpu *Cpu) disasmArm16(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
		return "dw " + strconv.FormatInt(int64(op), 16)
	}
}

func (cpu *Cpu) opArm17(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn -= off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm17(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm18(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm18(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm19(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm19(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm1A(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm1A(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm1B(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
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
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm1B(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm1C(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn += off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm1C(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm1D(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn += off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm1D(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm1E(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn += off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// STRH
				cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))
			case 2:
				// LDRD
				cpu.Regs[rdx] = reg(cpu.opRead32(rn))
				cpu.Regs[rdx+1] = reg(cpu.opRead32(rn + 4))
			case 3:
				// STRD
				cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))
				cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm1E(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("strd", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm1F(op uint32) {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			if !cpu.opArmCond(op) {
				return
			}
			rnx := (op >> 16) & 0xF
			rdx := (op >> 12) & 0xF
			rn := uint32(cpu.Regs[rnx])
			cpu.Regs[15] += 4
			off := (op & 0xF) | ((op & 0xF00) >> 4)
			rn += off
			code := (op >> 5) & 3
			switch code {
			case 1:
				// LDRH
				cpu.Regs[rdx] = reg(cpu.opRead16(rn))
			case 2:
				// LDRSB
				data := int32(int8(cpu.opRead8(rn)))
				cpu.Regs[rdx] = reg(data)
			case 3:
				// LDRSH
				data := int32(int16(cpu.opRead16(rn)))
				cpu.Regs[rdx] = reg(data)
			default:
				cpu.InvalidOpArm(op, "halfword invalid op")
			}
			cpu.Regs[rnx] = reg(rn)
		} else {
			// ?s
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
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
}

func (cpu *Cpu) disasmArm1F(op uint32, pc uint32) string {
	if op&0x90 == 0x90 {
		if op&0x60 != 0 {
			code := (op >> 5) & 3
			switch code {
			case 1:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 2:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsb", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			case 3:
				var out bytes.Buffer
				opcode := cpu.disasmAddCond("ldrsh", op)
				out.WriteString((opcode + "                ")[:10])
				arg0 := (op >> 12) & 0xF
				out.WriteString(RegNames[arg0])
				return out.String()
			default:
				return "dw " + strconv.FormatInt(int64(op), 16)
			}
		} else {
			return "dw " + strconv.FormatInt(int64(op), 16)
		}
	} else {
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
}

func (cpu *Cpu) opArm20(op uint32) {
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

func (cpu *Cpu) disasmArm20(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm21(op uint32) {
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

func (cpu *Cpu) disasmArm21(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm22(op uint32) {
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

func (cpu *Cpu) disasmArm22(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm23(op uint32) {
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

func (cpu *Cpu) disasmArm23(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm24(op uint32) {
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

func (cpu *Cpu) disasmArm24(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm25(op uint32) {
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

func (cpu *Cpu) disasmArm25(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm26(op uint32) {
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

func (cpu *Cpu) disasmArm26(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm27(op uint32) {
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

func (cpu *Cpu) disasmArm27(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm28(op uint32) {
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

func (cpu *Cpu) disasmArm28(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm29(op uint32) {
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

func (cpu *Cpu) disasmArm29(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm2A(op uint32) {
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

func (cpu *Cpu) disasmArm2A(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm2B(op uint32) {
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

func (cpu *Cpu) disasmArm2B(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm2C(op uint32) {
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

func (cpu *Cpu) disasmArm2C(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm2D(op uint32) {
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

func (cpu *Cpu) disasmArm2D(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm2E(op uint32) {
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

func (cpu *Cpu) disasmArm2E(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm2F(op uint32) {
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

func (cpu *Cpu) disasmArm2F(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm30(op uint32) {
	// MRS
	if !cpu.opArmCond(op) {
		return
	}
	mask := (op >> 16) & 0xF
	if mask != 0xF {
		cpu.InvalidOpArm(op, "unimplemented SWP")
		return
	}
	rdx := (op >> 12) & 0xF
	if rdx == 15 {
		cpu.InvalidOpArm(op, "write to PC in MRS")
		return
	}
	cpu.Regs[rdx] = reg(cpu.Cpsr.Uint32())
}

func (cpu *Cpu) disasmArm30(op uint32, pc uint32) string {
	return "dw " + strconv.FormatInt(int64(op), 16)
}

func (cpu *Cpu) opArm31(op uint32) {
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

func (cpu *Cpu) disasmArm31(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm32(op uint32) {
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

func (cpu *Cpu) disasmArm32(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArm33(op uint32) {
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

func (cpu *Cpu) disasmArm33(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm34(op uint32) {
	// MRS
	if !cpu.opArmCond(op) {
		return
	}
	mask := (op >> 16) & 0xF
	if mask != 0xF {
		cpu.InvalidOpArm(op, "unimplemented SWP")
		return
	}
	rdx := (op >> 12) & 0xF
	if rdx == 15 {
		cpu.InvalidOpArm(op, "write to PC in MRS")
		return
	}
	cpu.Regs[rdx] = reg(*cpu.RegSpsr())
}

func (cpu *Cpu) disasmArm34(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArm35(op uint32) {
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

func (cpu *Cpu) disasmArm35(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm36(op uint32) {
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

func (cpu *Cpu) disasmArm36(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArm37(op uint32) {
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

func (cpu *Cpu) disasmArm37(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm38(op uint32) {
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

func (cpu *Cpu) disasmArm38(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm39(op uint32) {
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

func (cpu *Cpu) disasmArm39(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm3A(op uint32) {
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

func (cpu *Cpu) disasmArm3A(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm3B(op uint32) {
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

func (cpu *Cpu) disasmArm3B(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm3C(op uint32) {
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

func (cpu *Cpu) disasmArm3C(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm3D(op uint32) {
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

func (cpu *Cpu) disasmArm3D(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm3E(op uint32) {
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

func (cpu *Cpu) disasmArm3E(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm3F(op uint32) {
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

func (cpu *Cpu) disasmArm3F(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm40(op uint32) {
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

func (cpu *Cpu) disasmArm40(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm41(op uint32) {
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

func (cpu *Cpu) disasmArm41(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm42(op uint32) {
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

func (cpu *Cpu) disasmArm42(op uint32, pc uint32) string {
	return cpu.disasmArm40(op, pc)
}

func (cpu *Cpu) opArm43(op uint32) {
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

func (cpu *Cpu) disasmArm43(op uint32, pc uint32) string {
	return cpu.disasmArm41(op, pc)
}

func (cpu *Cpu) opArm44(op uint32) {
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

func (cpu *Cpu) disasmArm44(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm45(op uint32) {
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

func (cpu *Cpu) disasmArm45(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm46(op uint32) {
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

func (cpu *Cpu) disasmArm46(op uint32, pc uint32) string {
	return cpu.disasmArm44(op, pc)
}

func (cpu *Cpu) opArm47(op uint32) {
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

func (cpu *Cpu) disasmArm47(op uint32, pc uint32) string {
	return cpu.disasmArm45(op, pc)
}

func (cpu *Cpu) opArm48(op uint32) {
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

func (cpu *Cpu) disasmArm48(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm49(op uint32) {
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

func (cpu *Cpu) disasmArm49(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm4A(op uint32) {
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

func (cpu *Cpu) disasmArm4A(op uint32, pc uint32) string {
	return cpu.disasmArm48(op, pc)
}

func (cpu *Cpu) opArm4B(op uint32) {
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

func (cpu *Cpu) disasmArm4B(op uint32, pc uint32) string {
	return cpu.disasmArm49(op, pc)
}

func (cpu *Cpu) opArm4C(op uint32) {
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

func (cpu *Cpu) disasmArm4C(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm4D(op uint32) {
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

func (cpu *Cpu) disasmArm4D(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm4E(op uint32) {
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

func (cpu *Cpu) disasmArm4E(op uint32, pc uint32) string {
	return cpu.disasmArm4C(op, pc)
}

func (cpu *Cpu) opArm4F(op uint32) {
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

func (cpu *Cpu) disasmArm4F(op uint32, pc uint32) string {
	return cpu.disasmArm4D(op, pc)
}

func (cpu *Cpu) opArm50(op uint32) {
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

func (cpu *Cpu) disasmArm50(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm51(op uint32) {
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

func (cpu *Cpu) disasmArm51(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm52(op uint32) {
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

func (cpu *Cpu) disasmArm52(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm53(op uint32) {
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

func (cpu *Cpu) disasmArm53(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm54(op uint32) {
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

func (cpu *Cpu) disasmArm54(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm55(op uint32) {
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

func (cpu *Cpu) disasmArm55(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm56(op uint32) {
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

func (cpu *Cpu) disasmArm56(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm57(op uint32) {
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

func (cpu *Cpu) disasmArm57(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm58(op uint32) {
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

func (cpu *Cpu) disasmArm58(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm59(op uint32) {
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

func (cpu *Cpu) disasmArm59(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm5A(op uint32) {
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

func (cpu *Cpu) disasmArm5A(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm5B(op uint32) {
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

func (cpu *Cpu) disasmArm5B(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm5C(op uint32) {
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

func (cpu *Cpu) disasmArm5C(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm5D(op uint32) {
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

func (cpu *Cpu) disasmArm5D(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm5E(op uint32) {
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

func (cpu *Cpu) disasmArm5E(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm5F(op uint32) {
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

func (cpu *Cpu) disasmArm5F(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm60(op uint32) {
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

func (cpu *Cpu) disasmArm60(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm61(op uint32) {
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

func (cpu *Cpu) disasmArm61(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm62(op uint32) {
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

func (cpu *Cpu) disasmArm62(op uint32, pc uint32) string {
	return cpu.disasmArm60(op, pc)
}

func (cpu *Cpu) opArm63(op uint32) {
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

func (cpu *Cpu) disasmArm63(op uint32, pc uint32) string {
	return cpu.disasmArm61(op, pc)
}

func (cpu *Cpu) opArm64(op uint32) {
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

func (cpu *Cpu) disasmArm64(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm65(op uint32) {
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

func (cpu *Cpu) disasmArm65(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm66(op uint32) {
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

func (cpu *Cpu) disasmArm66(op uint32, pc uint32) string {
	return cpu.disasmArm64(op, pc)
}

func (cpu *Cpu) opArm67(op uint32) {
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

func (cpu *Cpu) disasmArm67(op uint32, pc uint32) string {
	return cpu.disasmArm65(op, pc)
}

func (cpu *Cpu) opArm68(op uint32) {
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

func (cpu *Cpu) disasmArm68(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm69(op uint32) {
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

func (cpu *Cpu) disasmArm69(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm6A(op uint32) {
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

func (cpu *Cpu) disasmArm6A(op uint32, pc uint32) string {
	return cpu.disasmArm68(op, pc)
}

func (cpu *Cpu) opArm6B(op uint32) {
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

func (cpu *Cpu) disasmArm6B(op uint32, pc uint32) string {
	return cpu.disasmArm69(op, pc)
}

func (cpu *Cpu) opArm6C(op uint32) {
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

func (cpu *Cpu) disasmArm6C(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm6D(op uint32) {
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

func (cpu *Cpu) disasmArm6D(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm6E(op uint32) {
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

func (cpu *Cpu) disasmArm6E(op uint32, pc uint32) string {
	return cpu.disasmArm6C(op, pc)
}

func (cpu *Cpu) opArm6F(op uint32) {
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

func (cpu *Cpu) disasmArm6F(op uint32, pc uint32) string {
	return cpu.disasmArm6D(op, pc)
}

func (cpu *Cpu) opArm70(op uint32) {
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

func (cpu *Cpu) disasmArm70(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm71(op uint32) {
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

func (cpu *Cpu) disasmArm71(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm72(op uint32) {
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

func (cpu *Cpu) disasmArm72(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm73(op uint32) {
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

func (cpu *Cpu) disasmArm73(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm74(op uint32) {
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

func (cpu *Cpu) disasmArm74(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm75(op uint32) {
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

func (cpu *Cpu) disasmArm75(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm76(op uint32) {
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

func (cpu *Cpu) disasmArm76(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm77(op uint32) {
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

func (cpu *Cpu) disasmArm77(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm78(op uint32) {
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

func (cpu *Cpu) disasmArm78(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm79(op uint32) {
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

func (cpu *Cpu) disasmArm79(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm7A(op uint32) {
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

func (cpu *Cpu) disasmArm7A(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm7B(op uint32) {
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

func (cpu *Cpu) disasmArm7B(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm7C(op uint32) {
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

func (cpu *Cpu) disasmArm7C(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm7D(op uint32) {
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

func (cpu *Cpu) disasmArm7D(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm7E(op uint32) {
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

func (cpu *Cpu) disasmArm7E(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm7F(op uint32) {
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

func (cpu *Cpu) disasmArm7F(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm80(op uint32) {
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

func (cpu *Cpu) disasmArm80(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm81(op uint32) {
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

func (cpu *Cpu) disasmArm81(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm82(op uint32) {
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

func (cpu *Cpu) disasmArm82(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm83(op uint32) {
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

func (cpu *Cpu) disasmArm83(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm84(op uint32) {
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

func (cpu *Cpu) disasmArm84(op uint32, pc uint32) string {
	return cpu.disasmArm80(op, pc)
}

func (cpu *Cpu) opArm85(op uint32) {
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

func (cpu *Cpu) disasmArm85(op uint32, pc uint32) string {
	return cpu.disasmArm81(op, pc)
}

func (cpu *Cpu) opArm86(op uint32) {
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

func (cpu *Cpu) disasmArm86(op uint32, pc uint32) string {
	return cpu.disasmArm82(op, pc)
}

func (cpu *Cpu) opArm87(op uint32) {
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

func (cpu *Cpu) disasmArm87(op uint32, pc uint32) string {
	return cpu.disasmArm83(op, pc)
}

func (cpu *Cpu) opArm88(op uint32) {
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

func (cpu *Cpu) disasmArm88(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm89(op uint32) {
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

func (cpu *Cpu) disasmArm89(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm8A(op uint32) {
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

func (cpu *Cpu) disasmArm8A(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm8B(op uint32) {
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

func (cpu *Cpu) disasmArm8B(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm8C(op uint32) {
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

func (cpu *Cpu) disasmArm8C(op uint32, pc uint32) string {
	return cpu.disasmArm88(op, pc)
}

func (cpu *Cpu) opArm8D(op uint32) {
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

func (cpu *Cpu) disasmArm8D(op uint32, pc uint32) string {
	return cpu.disasmArm89(op, pc)
}

func (cpu *Cpu) opArm8E(op uint32) {
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

func (cpu *Cpu) disasmArm8E(op uint32, pc uint32) string {
	return cpu.disasmArm8A(op, pc)
}

func (cpu *Cpu) opArm8F(op uint32) {
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

func (cpu *Cpu) disasmArm8F(op uint32, pc uint32) string {
	return cpu.disasmArm8B(op, pc)
}

func (cpu *Cpu) opArm90(op uint32) {
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

func (cpu *Cpu) disasmArm90(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm91(op uint32) {
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

func (cpu *Cpu) disasmArm91(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm92(op uint32) {
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

func (cpu *Cpu) disasmArm92(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm93(op uint32) {
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

func (cpu *Cpu) disasmArm93(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm94(op uint32) {
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

func (cpu *Cpu) disasmArm94(op uint32, pc uint32) string {
	return cpu.disasmArm90(op, pc)
}

func (cpu *Cpu) opArm95(op uint32) {
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

func (cpu *Cpu) disasmArm95(op uint32, pc uint32) string {
	return cpu.disasmArm91(op, pc)
}

func (cpu *Cpu) opArm96(op uint32) {
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

func (cpu *Cpu) disasmArm96(op uint32, pc uint32) string {
	return cpu.disasmArm92(op, pc)
}

func (cpu *Cpu) opArm97(op uint32) {
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

func (cpu *Cpu) disasmArm97(op uint32, pc uint32) string {
	return cpu.disasmArm93(op, pc)
}

func (cpu *Cpu) opArm98(op uint32) {
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

func (cpu *Cpu) disasmArm98(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm99(op uint32) {
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

func (cpu *Cpu) disasmArm99(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm9A(op uint32) {
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

func (cpu *Cpu) disasmArm9A(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm9B(op uint32) {
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

func (cpu *Cpu) disasmArm9B(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArm9C(op uint32) {
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

func (cpu *Cpu) disasmArm9C(op uint32, pc uint32) string {
	return cpu.disasmArm98(op, pc)
}

func (cpu *Cpu) opArm9D(op uint32) {
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

func (cpu *Cpu) disasmArm9D(op uint32, pc uint32) string {
	return cpu.disasmArm99(op, pc)
}

func (cpu *Cpu) opArm9E(op uint32) {
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

func (cpu *Cpu) disasmArm9E(op uint32, pc uint32) string {
	return cpu.disasmArm9A(op, pc)
}

func (cpu *Cpu) opArm9F(op uint32) {
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

func (cpu *Cpu) disasmArm9F(op uint32, pc uint32) string {
	return cpu.disasmArm9B(op, pc)
}

func (cpu *Cpu) opArmA0(op uint32) {
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

func (cpu *Cpu) disasmArmA0(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArmA1(op uint32) {
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

func (cpu *Cpu) disasmArmA1(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA2(op uint32) {
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

func (cpu *Cpu) disasmArmA2(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA3(op uint32) {
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

func (cpu *Cpu) disasmArmA3(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA4(op uint32) {
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

func (cpu *Cpu) disasmArmA4(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA5(op uint32) {
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

func (cpu *Cpu) disasmArmA5(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA6(op uint32) {
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

func (cpu *Cpu) disasmArmA6(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA7(op uint32) {
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

func (cpu *Cpu) disasmArmA7(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA8(op uint32) {
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

func (cpu *Cpu) disasmArmA8(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmA9(op uint32) {
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

func (cpu *Cpu) disasmArmA9(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmAA(op uint32) {
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

func (cpu *Cpu) disasmArmAA(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmAB(op uint32) {
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

func (cpu *Cpu) disasmArmAB(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmAC(op uint32) {
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

func (cpu *Cpu) disasmArmAC(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmAD(op uint32) {
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

func (cpu *Cpu) disasmArmAD(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmAE(op uint32) {
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

func (cpu *Cpu) disasmArmAE(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmAF(op uint32) {
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

func (cpu *Cpu) disasmArmAF(op uint32, pc uint32) string {
	return cpu.disasmArmA0(op, pc)
}

func (cpu *Cpu) opArmB0(op uint32) {
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

func (cpu *Cpu) disasmArmB0(op uint32, pc uint32) string {
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

func (cpu *Cpu) opArmB1(op uint32) {
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

func (cpu *Cpu) disasmArmB1(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB2(op uint32) {
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

func (cpu *Cpu) disasmArmB2(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB3(op uint32) {
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

func (cpu *Cpu) disasmArmB3(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB4(op uint32) {
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

func (cpu *Cpu) disasmArmB4(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB5(op uint32) {
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

func (cpu *Cpu) disasmArmB5(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB6(op uint32) {
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

func (cpu *Cpu) disasmArmB6(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB7(op uint32) {
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

func (cpu *Cpu) disasmArmB7(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB8(op uint32) {
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

func (cpu *Cpu) disasmArmB8(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmB9(op uint32) {
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

func (cpu *Cpu) disasmArmB9(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmBA(op uint32) {
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

func (cpu *Cpu) disasmArmBA(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmBB(op uint32) {
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

func (cpu *Cpu) disasmArmBB(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmBC(op uint32) {
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

func (cpu *Cpu) disasmArmBC(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmBD(op uint32) {
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

func (cpu *Cpu) disasmArmBD(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmBE(op uint32) {
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

func (cpu *Cpu) disasmArmBE(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmBF(op uint32) {
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

func (cpu *Cpu) disasmArmBF(op uint32, pc uint32) string {
	return cpu.disasmArmB0(op, pc)
}

func (cpu *Cpu) opArmC0(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC0(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC1(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC1(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC2(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC2(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC3(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC3(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC4(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC4(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC5(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC5(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC6(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC6(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC7(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC7(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC8(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC8(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmC9(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmC9(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmCA(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmCA(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmCB(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmCB(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmCC(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmCC(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmCD(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmCD(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmCE(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmCE(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmCF(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmCF(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD0(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD0(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD1(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD1(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD2(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD2(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD3(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD3(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD4(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD4(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD5(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD5(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD6(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD6(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD7(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD7(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD8(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD8(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmD9(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmD9(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmDA(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmDA(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmDB(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmDB(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmDC(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmDC(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmDD(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmDD(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmDE(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmDE(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmDF(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) disasmArmDF(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE0(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmE0(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE1(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmE1(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE2(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmE2(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE3(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmE3(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE4(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmE4(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE5(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmE5(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE6(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmE6(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE7(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmE7(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE8(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmE8(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmE9(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmE9(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmEA(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmEA(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmEB(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmEB(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmEC(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmEC(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmED(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmED(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmEE(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	cpu.Regs[15] += 4
	rd := cpu.Regs[rdx]
	cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))
}

func (cpu *Cpu) disasmArmEE(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmEF(op uint32) {
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
	if (op & 0x10) == 0 { // CDP
		cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)
		return
	}
	res := cpu.opCopRead(copnum, opc, cn, cm, cp)
	if rdx == 15 {
		cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu)
	} else {
		cpu.Regs[rdx] = reg(res)
	}
}

func (cpu *Cpu) disasmArmEF(op uint32, pc uint32) string {
	return cpu.disasmArm30(op, pc)
}

func (cpu *Cpu) opArmF0(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF0(op uint32, pc uint32) string {
	var out bytes.Buffer
	opcode := cpu.disasmAddCond("swi", op)
	out.WriteString((opcode + "                ")[:10])
	arg0 := int64(op & 0xFF)
	out.WriteString("#0x")
	out.WriteString(strconv.FormatInt(arg0, 16))
	return out.String()
}

func (cpu *Cpu) opArmF1(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF1(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF2(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF2(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF3(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF3(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF4(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF4(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF5(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF5(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF6(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF6(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF7(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF7(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF8(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF8(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmF9(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmF9(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmFA(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmFA(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmFB(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmFB(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmFC(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmFC(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmFD(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmFD(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmFE(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmFE(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}

func (cpu *Cpu) opArmFF(op uint32) {
	cpu.Exception(ExceptionSwi)
}

func (cpu *Cpu) disasmArmFF(op uint32, pc uint32) string {
	return cpu.disasmArmF0(op, pc)
}
