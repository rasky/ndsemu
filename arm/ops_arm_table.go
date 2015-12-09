// Generated on 2015-12-09 01:04:52.661328387 +0100 CET
package arm

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
			// MUL
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
		// AND
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
			// MULS
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
		// ANDS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := rn & op2
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
			// MLA
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
		// EOR
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
			// MLAS
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
		// EORS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := rn ^ op2
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
		// SUB
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// SUBS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := rn - op2
		cpu.Cpsr.SetC(rn < res)
		cpu.Cpsr.SetVSub(rn, op2, res)
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
		// RSB
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// RSBS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := op2 - rn
		cpu.Cpsr.SetC(op2 < res)
		cpu.Cpsr.SetVSub(op2, rn, res)
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
			// UMULL
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
		// ADD
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
			// UMULLS
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
		// ADDS
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
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
			// UMLAL
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
		// ADC
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
			// UMLALS
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
		// ADCS
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
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
			// SMULL
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
		// SBC
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
			// SMULLS
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
		// SBCS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		cf := cpu.Cpsr.CB()
		res := rn - op2
		cpu.Cpsr.SetC(rn < res)
		cpu.Cpsr.SetVSub(rn, op2, res)
		res += cf - 1
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
			// SMLAL
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
		// RSC
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
			// SMLALS
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
		// RSCS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		cf := cpu.Cpsr.CB()
		res := op2 - rn
		cpu.Cpsr.SetC(op2 < res)
		cpu.Cpsr.SetVSub(op2, rn, res)
		res += cf - 1
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// TSTS
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// TEQS
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
			// ?
			cpu.InvalidOpArm(op, "unhandled mul-type")
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// CMPS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := rn - op2
		cpu.Cpsr.SetC(rn < res)
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// CMNS
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
		// ORR
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// ORRS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := rn | op2
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
		// MOV
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// MOVS
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
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
		// BIC
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// BICS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := rn & ^op2
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
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
		// MVN
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
			// ?S
			cpu.InvalidOpArm(op, "unhandled mul-type")
		}
	} else {
		// MVNS
		if !cpu.opArmCond(op) {
			return
		}
		rnx := (op >> 16) & 0xF
		rdx := (op >> 12) & 0xF
		op2 := cpu.opDecodeAluOp2Reg(op, true)
		rn := uint32(cpu.Regs[rnx])
		res := ^op2
		cpu.Cpsr.SetNZ(res)
		if rdx == 15 {
			cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
			return
		}
		cpu.Regs[rdx] = reg(res)
		if rdx == 15 {
			cpu.pc = reg(res)
		}
		_ = res
		_ = rn
	}
}

func (cpu *Cpu) opArm20(op uint32) {
	// AND
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

func (cpu *Cpu) opArm21(op uint32) {
	// ANDS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm22(op uint32) {
	// EOR
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

func (cpu *Cpu) opArm23(op uint32) {
	// EORS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm24(op uint32) {
	// SUB
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

func (cpu *Cpu) opArm25(op uint32) {
	// SUBS
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(rn < res)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm26(op uint32) {
	// RSB
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

func (cpu *Cpu) opArm27(op uint32) {
	// RSBS
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := op2 - rn
	cpu.Cpsr.SetC(op2 < res)
	cpu.Cpsr.SetVSub(op2, rn, res)
	cpu.Cpsr.SetNZ(res)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm28(op uint32) {
	// ADD
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

func (cpu *Cpu) opArm29(op uint32) {
	// ADDS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm2A(op uint32) {
	// ADC
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

func (cpu *Cpu) opArm2B(op uint32) {
	// ADCS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm2C(op uint32) {
	// SBC
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

func (cpu *Cpu) opArm2D(op uint32) {
	// SBCS
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
	cpu.Cpsr.SetC(rn < res)
	cpu.Cpsr.SetVSub(rn, op2, res)
	res += cf - 1
	cpu.Cpsr.SetNZ(res)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm2E(op uint32) {
	// RSC
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

func (cpu *Cpu) opArm2F(op uint32) {
	// RSCS
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
	cpu.Cpsr.SetC(op2 < res)
	cpu.Cpsr.SetVSub(op2, rn, res)
	res += cf - 1
	cpu.Cpsr.SetNZ(res)
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
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

func (cpu *Cpu) opArm31(op uint32) {
	// TSTS
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

func (cpu *Cpu) opArm33(op uint32) {
	// TEQS
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

func (cpu *Cpu) opArm35(op uint32) {
	// CMPS
	if !cpu.opArmCond(op) {
		return
	}
	rnx := (op >> 16) & 0xF
	rdx := (op >> 12) & 0xF
	rot := uint((op >> 7) & 0x1E)
	op2 := ((op & 0xFF) >> rot) | ((op & 0xFF) << (32 - rot))
	rn := uint32(cpu.Regs[rnx])
	res := rn - op2
	cpu.Cpsr.SetC(rn < res)
	cpu.Cpsr.SetVSub(rn, op2, res)
	cpu.Cpsr.SetNZ(res)
	if rdx != 0 && rdx != 15 {
		cpu.InvalidOpArm(op, "invalid rdx on test")
		return
	}
	_ = res
	_ = rn
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

func (cpu *Cpu) opArm37(op uint32) {
	// CMNS
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

func (cpu *Cpu) opArm38(op uint32) {
	// ORR
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

func (cpu *Cpu) opArm39(op uint32) {
	// ORRS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm3A(op uint32) {
	// MOV
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

func (cpu *Cpu) opArm3B(op uint32) {
	// MOVS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm3C(op uint32) {
	// BIC
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

func (cpu *Cpu) opArm3D(op uint32) {
	// BICS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
}

func (cpu *Cpu) opArm3E(op uint32) {
	// MVN
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

func (cpu *Cpu) opArm3F(op uint32) {
	// MVNS
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
	if rdx == 15 {
		cpu.InvalidOpArm(op, "unimplemented RD=15 with ALU")
		return
	}
	cpu.Regs[rdx] = reg(res)
	if rdx == 15 {
		cpu.pc = reg(res)
	}
	_ = res
	_ = rn
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

func (cpu *Cpu) opArm80(op uint32) {
	// STMDA
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

func (cpu *Cpu) opArm81(op uint32) {
	// LDMDA
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm82(op uint32) {
	// STMDA
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

func (cpu *Cpu) opArm83(op uint32) {
	// LDMDA
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm84(op uint32) {
	// STMDA
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

func (cpu *Cpu) opArm85(op uint32) {
	// LDMDA
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm86(op uint32) {
	// STMDA
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

func (cpu *Cpu) opArm87(op uint32) {
	// LDMDA
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm88(op uint32) {
	// STMIA
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

func (cpu *Cpu) opArm89(op uint32) {
	// POP
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm8A(op uint32) {
	// STMIA
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

func (cpu *Cpu) opArm8B(op uint32) {
	// POP
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm8C(op uint32) {
	// STMIA
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

func (cpu *Cpu) opArm8D(op uint32) {
	// POP
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm8E(op uint32) {
	// STMIA
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

func (cpu *Cpu) opArm8F(op uint32) {
	// POP
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm90(op uint32) {
	// PUSH
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

func (cpu *Cpu) opArm91(op uint32) {
	// LDMDB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm92(op uint32) {
	// PUSH
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

func (cpu *Cpu) opArm93(op uint32) {
	// LDMDB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm94(op uint32) {
	// PUSH
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

func (cpu *Cpu) opArm95(op uint32) {
	// LDMDB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm96(op uint32) {
	// PUSH
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

func (cpu *Cpu) opArm97(op uint32) {
	// LDMDB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
			rn += 4
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(orn)
}

func (cpu *Cpu) opArm98(op uint32) {
	// STMIB
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

func (cpu *Cpu) opArm99(op uint32) {
	// LDMIB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm9A(op uint32) {
	// STMIB
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

func (cpu *Cpu) opArm9B(op uint32) {
	// LDMIB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
}

func (cpu *Cpu) opArm9C(op uint32) {
	// STMIB
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

func (cpu *Cpu) opArm9D(op uint32) {
	// LDMIB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
}

func (cpu *Cpu) opArm9E(op uint32) {
	// STMIB
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

func (cpu *Cpu) opArm9F(op uint32) {
	// LDMIB
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
					cpu.Regs[15] = cpu.Regs[15] &^ 1
				}
				cpu.pc = cpu.Regs[15]
			}
		}
		mask >>= 1
	}
	cpu.Regs[rnx] = reg(rn)
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

func (cpu *Cpu) opArmC0(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC1(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC2(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC3(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC4(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC5(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC6(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC7(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC8(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmC9(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmCA(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmCB(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmCC(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmCD(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmCE(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmCF(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD0(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD1(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD2(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD3(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD4(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD5(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD6(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD7(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD8(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmD9(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmDA(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmDB(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmDC(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmDD(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmDE(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
}

func (cpu *Cpu) opArmDF(op uint32) {
	cpu.InvalidOpArm(op, "unimplemented")
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

func (cpu *Cpu) opArmF0(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF1(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF2(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF3(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF4(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF5(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF6(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF7(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF8(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmF9(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmFA(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmFB(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmFC(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmFD(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmFE(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}

func (cpu *Cpu) opArmFF(op uint32) {
	cpu.InvalidOpArm(op, "unhandled cop")
}
