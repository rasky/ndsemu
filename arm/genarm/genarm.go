package main

import (
	"fmt"
	"ndsemu/emu/cpugen"
)

type Generator struct {
	*cpugen.Generator
}

func (g *Generator) writeOpCond(op uint32) {
	fmt.Fprintf(g, "if !cpu.opArmCond(op) { return }\n")
}

func (g *Generator) WriteDisasm(opname string, args ...string) {
	g.Generator.WriteDisasm(fmt.Sprintf("!cpu.disasmAddCond(%q, op)", opname), args...)
}

func (g *Generator) writeOpSwp(op uint32) {
	byt := (op>>22)&1 != 0

	if (op>>24)&0xF != 1 || ((op>>20)&0xF != 0 && (op>>20)&0xF != 4) {
		panic("invalid call to writeOpSwp")
	}

	g.WriteExitIfOpInvalid("op&0x0FB00FF0 != 0x01000090", "invalid opcode decoded as SWP")
	name := "swp"
	if byt {
		name = "swpb"
	}
	fmt.Fprintf(g, "// %s\n", name)

	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := (op >> 16) & 0xF\n")
	fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")

	fmt.Fprintf(g, "rmx := (op >> 0) & 0xF\n")
	fmt.Fprintf(g, "rm := uint32(cpu.Regs[rmx])\n")

	fmt.Fprintf(g, "rdx := (op >> 12) & 0xF\n")

	if byt {
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead8(rn))\n")
		fmt.Fprintf(g, "cpu.opWrite8(rn, uint8(rm))\n")
	} else {
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead32(rn))\n")
		fmt.Fprintf(g, "cpu.opWrite32(rn, rm)\n")
	}

	g.WriteDisasm(name, "r:(op >> 12) & 0xF", "r:(op >> 0) & 0xF", "l:(op >> 16) & 0xF")
}

var mulNames = [16]string{
	"mul", "mla", "?", "?", "umull", "umlal", "smull", "smlal",
	"?", "?", "?", "?", "?", "?", "?", "?",
}

func (g *Generator) writeOpMul(op uint32) {
	setflags := (op>>20)&1 != 0
	code := (op >> 21) & 0xF
	acc := code&1 != 0

	name := mulNames[code]
	if setflags {
		name += "s"
	}
	fmt.Fprintf(g, "// %s\n", name)

	if mulNames[code] == "?" {
		g.WriteOpInvalid("unhandled mul-type")
		g.WriteDisasmInvalid()
		return
	}

	g.writeOpCond(op)
	fmt.Fprintf(g, "rsx := (op >> 8) & 0xF\n")
	fmt.Fprintf(g, "rs := uint32(cpu.Regs[rsx])\n")

	fmt.Fprintf(g, "rmx := (op >> 0) & 0xF\n")
	fmt.Fprintf(g, "rm := uint32(cpu.Regs[rmx])\n")

	fmt.Fprintf(g, "data := (op >> 4) & 0xF;\n")
	g.WriteExitIfOpInvalid("data!=0x9", "unimplemented half-word multiplies")

	fmt.Fprintf(g, "rdx := (op >> 16) & 0xF\n")

	switch code {
	case 0: // MUL
		fmt.Fprintf(g, "res := rm*rs\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
		}
		g.WriteDisasm(name, "r:(op >> 16) & 0xF", "r:(op >> 0) & 0xF", "r:(op >> 8) & 0xF")
	case 1: // MLA
		fmt.Fprintf(g, "rnx := (op >> 12) & 0xF\n")
		fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")
		fmt.Fprintf(g, "res := rm*rs+rn\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
		}
		g.WriteDisasm(name, "r:(op >> 16) & 0xF", "r:(op >> 0) & 0xF", "r:(op >> 8) & 0xF", "r:(op >> 12) & 0xF")
	case 4, 5: // UMULL, UMLAL
		fmt.Fprintf(g, "res64 := uint64(rm)*uint64(rs)\n")
		fmt.Fprintf(g, "rnx := (op >> 12) & 0xF\n")
		if acc {
			fmt.Fprintf(g, "app64 := uint64(cpu.Regs[rnx]) + uint64(cpu.Regs[rdx]) << 32\n")
			fmt.Fprintf(g, "res64 += app64\n")
		}
		fmt.Fprintf(g, "cpu.Regs[rnx] = reg(res64)\n")
		fmt.Fprintf(g, "res := uint32(res64 >> 32)\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetNZ64(res64)\n")
		}
		g.WriteDisasm(name, "r:(op >> 16) & 0xF", "r:(op >> 0) & 0xF", "r:(op >> 8) & 0xF", "r:(op >> 12) & 0xF")
	case 6, 7: // SMULL, SMLAL
		fmt.Fprintf(g, "res64 := int64(int32(rm))*int64(int32(rs))\n")
		fmt.Fprintf(g, "rnx := (op >> 12) & 0xF\n")
		if acc {
			fmt.Fprintf(g, "app64 := uint64(cpu.Regs[rnx]) + uint64(cpu.Regs[rdx]) << 32\n")
			fmt.Fprintf(g, "res64 += int64(app64)\n")
		}
		fmt.Fprintf(g, "cpu.Regs[rnx] = reg(res64)\n")
		fmt.Fprintf(g, "res := uint32(res64 >> 32)\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetNZ64(uint64(res64))\n")
		}
		g.WriteDisasm(name, "r:(op >> 16) & 0xF", "r:(op >> 0) & 0xF", "r:(op >> 8) & 0xF", "r:(op >> 12) & 0xF")
	default:
		panic("unreachable")
	}

	fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
}

func (g *Generator) writeOpBx(op uint32) {
	link := op&0x20 != 0

	name := "bx"
	if link {
		name = "blx"
	}
	fmt.Fprintf(g, "// %s reg\n", name)

	g.WriteExitIfOpInvalid("op&0x0FFFFFD0 != 0x012FFF10", "invalid opcode decoded as BX/BLX")
	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := op&0xF\n")
	fmt.Fprintf(g, "rn := cpu.Regs[rnx]\n")
	if link {
		fmt.Fprintf(g, "cpu.Regs[14] = cpu.Regs[15]-4\n")
	}
	fmt.Fprintf(g, "if rn&1 != 0 { cpu.Cpsr.SetT(true); rn &^= 1 } else { rn &^= 3 }\n")
	fmt.Fprintf(g, "cpu.pc = rn\n")

	// disasm
	g.WriteDisasm(name, "r:op&0xF")
}

func (g *Generator) writeOpPsrTransfer(op uint32) {
	imm := (op>>25)&1 != 0
	if (op>>26)&3 != 0 || (op>>23)&0x3 != 2 || (op>>20)&1 != 0 {
		panic("invalid psr decoding")
	}
	spsr := (op>>22)&1 != 0
	tostat := (op>>21)&1 != 0

	if imm {
		g.WriteExitIfOpInvalid("op&0x0FB00000 != 0x03200000", "invalid opcode decoded as PSR_imm")
	} else {
		g.WriteExitIfOpInvalid("op&0x0F900FF0 != 0x01000000", "invalid opcode decoded as PSR_reg")
	}

	if !tostat {
		fmt.Fprintf(g, "// MRS\n")
		g.writeOpCond(op)
		fmt.Fprintf(g, "mask := (op>>16)&0xF\n")
		g.WriteExitIfOpInvalid("mask != 0xF", "mask should be 0xF in MRS (is it SWP?)")
		fmt.Fprintf(g, "rdx := (op>>12)&0xF\n")
		g.WriteExitIfOpInvalid("rdx == 15", "write to PC in MRS")
		if spsr {
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(*cpu.RegSpsr())")
		} else {
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.Cpsr.Uint32())")
		}
	} else {
		fmt.Fprintf(g, "// MSR\n")
		g.writeOpCond(op)
		fmt.Fprintf(g, "mask := uint32(0)\n")
		fmt.Fprintf(g, "if (op>>19)&1!=0 { mask |= 0xFF000000 }\n")
		fmt.Fprintf(g, "if (op>>18)&1!=0 { mask |= 0x00FF0000 }\n")
		fmt.Fprintf(g, "if (op>>17)&1!=0 { mask |= 0x0000FF00 }\n")
		fmt.Fprintf(g, "if (op>>16)&1!=0 { mask |= 0x000000FF }\n")

		if imm {
			fmt.Fprintf(g, "val := op & 0xFF\n")
			fmt.Fprintf(g, "shcnt := uint(((op >> 8) & 0xF)*2)\n")
			fmt.Fprintf(g, "val = (val >> shcnt) | (val << (32-shcnt))\n")
		} else {
			fmt.Fprintf(g, "rmx := op & 0xF\n")
			fmt.Fprintf(g, "val := uint32(cpu.Regs[rmx])\n")
		}

		if spsr {
			fmt.Fprintf(g, "cpu.RegSpsr().SetWithMask(val, mask)\n")
		} else {
			fmt.Fprintf(g, "cpu.Cpsr.SetWithMask(val, mask, cpu)\n")
		}
	}

	g.WriteDisasmInvalid()
}

var aluNames = [16]string{
	"and", "eor", "sub", "rsb", "add", "adc", "sbc", "rsc",
	"tst", "teq", "cmp", "cmn", "orr", "mov", "bic", "mvn",
}

func (g *Generator) writeOpAlu(op uint32) {
	imm := (op>>25)&1 != 0
	code := (op >> 21) & 0xF
	setflags := (op>>20)&1 != 0

	if code >= 8 && code < 12 && !setflags {
		g.WriteOpInvalid("invalid ALU test function without flags")
		g.WriteDisasmInvalid()
		return
	}

	name := aluNames[code]
	if setflags {
		name += "s"
	}
	fmt.Fprintf(g, "// %s\n", name)

	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := (op >> 16) & 0xF\n")
	fmt.Fprintf(g, "rdx := (op >> 12) & 0xF\n")

	var disop2 string
	if imm {
		fmt.Fprintf(g, "rot := uint((op>>7)&0x1E)\n")
		fmt.Fprintf(g, "op2 := ((op&0xFF)>>rot) | ((op&0xFF)<<(32-rot))\n")
		disop2 = "x:((op&0xFF)>>((op>>7)&0x1E)) | ((op&0xFF)<<(32-((op>>7)&0x1E)))"
	} else {
		fmt.Fprintf(g, "op2 := cpu.opDecodeAluOp2Reg(op, true)\n")
		disop2 = "s:cpu.disasmOp2(op)"
	}

	// NOTE: we lookup register after opDecodeAluOp2Reg, because it can modify
	// Regs[15] (PC) because of pipelining.
	fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")

	test := false
	switch code {
	case 8: // TST
		test = true
		fallthrough
	case 0: // AND
		fmt.Fprintf(g, "res := rn & op2\n")
	case 9: // TEQ
		test = true
		fallthrough
	case 1: // XOR
		fmt.Fprintf(g, "res := rn ^ op2\n")
	case 10: // CMP
		test = true
		fallthrough
	case 2: // SUB
		fmt.Fprintf(g, "res := rn - op2\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetC(res<=rn)\n")
			fmt.Fprintf(g, "cpu.Cpsr.SetVSub(rn,op2,res)\n")
		}
	case 3: // RSB
		fmt.Fprintf(g, "res := op2 - rn\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetC(res<=op2)\n")
			fmt.Fprintf(g, "cpu.Cpsr.SetVSub(op2,rn,res)\n")
		}
	case 11: // CMN
		test = true
		fallthrough
	case 4: // ADD
		fmt.Fprintf(g, "res := rn + op2\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetC(rn>res)\n")
			fmt.Fprintf(g, "cpu.Cpsr.SetVAdd(rn,op2,res)\n")
		}
	case 5: // ADC
		fmt.Fprintf(g, "cf := cpu.Cpsr.CB()\n")
		fmt.Fprintf(g, "res := rn + op2\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetC(rn>res)\n")
			fmt.Fprintf(g, "cpu.Cpsr.SetVAdd(rn,op2,res)\n")
		}
		fmt.Fprintf(g, "res += cf\n")
	case 6: // SBC
		fmt.Fprintf(g, "cf := cpu.Cpsr.CB()\n")
		fmt.Fprintf(g, "res := rn - op2\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetC(res<=rn)\n")
			fmt.Fprintf(g, "cpu.Cpsr.SetVSub(rn,op2,res)\n")
		}
		fmt.Fprintf(g, "res += cf - 1\n")
	case 7: // RSC
		fmt.Fprintf(g, "cf := cpu.Cpsr.CB()\n")
		fmt.Fprintf(g, "res := op2 - rn\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetC(res<=op2)\n")
			fmt.Fprintf(g, "cpu.Cpsr.SetVSub(op2,rn,res)\n")
		}
		fmt.Fprintf(g, "res += cf - 1\n")
	case 12: // ORR
		fmt.Fprintf(g, "res := rn | op2\n")
	case 13: // MOV
		g.WriteExitIfOpInvalid("rnx!=0", "rn!=0 on NOV")
		fmt.Fprintf(g, "res := op2\n")
	case 14: // BIC
		fmt.Fprintf(g, "res := rn & ^op2\n")
	case 15: // MVN
		fmt.Fprintf(g, "res := ^op2\n")
	}

	if setflags {
		fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
	}

	if !test {
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
		fmt.Fprintf(g, "if rdx == 15 {\n")
		fmt.Fprintf(g, "cpu.pc = reg(res)\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)\n")
		}
		fmt.Fprintf(g, "}\n")
	} else {
		if !setflags {
			panic("unreachable")
		}
		g.WriteExitIfOpInvalid("rdx != 0 && rdx != 15", "invalid rdx on test")
	}

	fmt.Fprintf(g, "_ = res; _ = rn\n")

	if test {
		if setflags {
			name = name[:len(name)-1] // remove superflous 's' for test opcodes
		}
		g.WriteDisasm(name, "r:(op>>16)&0xF", disop2)
	} else if code == 13 || code == 15 {
		g.WriteDisasm(name, "r:(op>>12)&0xF", disop2)
	} else {
		g.WriteDisasm(name, "r:(op>>12)&0xF", "r:(op>>16)&0xF", disop2)
	}
}

func (g *Generator) writeOpBranchInner(link bool) {
	fmt.Fprintf(g, "off := int32(op<<9)>>7\n")
	if link {
		fmt.Fprintf(g, "cpu.Regs[14] = cpu.Regs[15]-4\n")
	}
	fmt.Fprintf(g, "cpu.Regs[15] += reg(off)\n")
}

func (g *Generator) writeOpBranch(op uint32) {
	link := op&(1<<24) != 0

	fmt.Fprintf(g, "if op>>28 == 0xF {\n")
	fmt.Fprintf(g, "// BLX_imm\n")
	// BLX is always a link-branch, and linkbit is used as halfword offset
	g.writeOpBranchInner(true)
	if link {
		fmt.Fprintf(g, "  cpu.Regs[15] += 2\n")
	}
	fmt.Fprintf(g, "  cpu.pc = cpu.Regs[15]\n")
	fmt.Fprintf(g, "  cpu.Cpsr.SetT(true)\n")
	fmt.Fprintf(g, "  return\n")
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(&g.Disasm, "if op>>28 == 0xF {\n")
	g.WriteDisasm("blx", "o:int32(op<<9)>>7")
	fmt.Fprintf(&g.Disasm, "}\n")

	if link {
		fmt.Fprintf(g, "// BL\n")
		g.WriteDisasm("bl", "o:int32(op<<9)>>7")
	} else {
		fmt.Fprintf(g, "// B\n")
		g.WriteDisasm("b", "o:int32(op<<9)>>7")
	}
	g.writeOpCond(op)
	g.writeOpBranchInner(link)
	fmt.Fprintf(g, "cpu.pc = cpu.Regs[15]\n")

}

func (g *Generator) writeOpSwi(op uint32) {
	g.WriteDisasm("swi", "x:op&0xFFFFFF")
	fmt.Fprintf(g, "cpu.Exception(ExceptionSwi)\n")
}

func (g *Generator) writeOpCoprocessor(op uint32) {
	copread := (op>>20)&1 != 0
	cdp := (op & 0x10) == 0

	if cdp {
		fmt.Fprintf(g, "// CDP\n")
	} else if copread {
		fmt.Fprintf(g, "// MRC\n")
	} else {
		fmt.Fprintf(g, "// MCR\n")
	}
	fmt.Fprintf(g, "if (op>>28)!=0xF { // MRC2/MCR2\n")
	g.writeOpCond(op)
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(g, "opc   := (op>>21)&0x7\n")
	fmt.Fprintf(g, "cn    := (op>>16)&0xF\n")
	fmt.Fprintf(g, "rdx   := (op>>12)&0xF\n")
	fmt.Fprintf(g, "copnum:= (op>>8)&0xF\n")
	fmt.Fprintf(g, "cp    := (op>>5)&0x7\n")
	fmt.Fprintf(g, "cm    := (op>>0)&0xF\n")

	if cdp {
		fmt.Fprintf(g, "cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)\n")
		return
	} else if copread {
		fmt.Fprintf(g, "res := cpu.opCopRead(copnum, opc, cn, cm, cp)\n")
		fmt.Fprintf(g, "if rdx==15 { cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu) }")
		fmt.Fprintf(g, "else { cpu.Regs[rdx] = reg(res) }")
	} else {
		fmt.Fprintf(g, "cpu.Regs[15]+=4\n")
		fmt.Fprintf(g, "rd := cpu.Regs[rdx]\n")
		fmt.Fprintf(g, "cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))\n")
	}
}

func (g *Generator) writeOpMemory(op uint32) {
	shreg := (op>>25)&1 != 0
	pre := (op>>24)&1 != 0
	up := (op>>23)&1 != 0
	byt := (op>>22)&1 != 0
	wb := (op>>21)&1 != 0
	load := (op>>20)&1 != 0

	g.WriteExitIfOpInvalid("(op>>28)==0xF", "PLD not supported")
	g.writeOpCond(op)

	fmt.Fprintf(g, "rnx := (op>>16)&0xF\n")
	fmt.Fprintf(g, "rdx := (op>>12)&0xF\n")
	fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")
	fmt.Fprintf(g, "cpu.Regs[15]+=4\n")

	if shreg {
		fmt.Fprintf(g, "off := cpu.opDecodeAluOp2Reg(op, false)\n")
	} else {
		fmt.Fprintf(g, "off := op & 0xFFF\n")
	}

	if pre {
		if up {
			fmt.Fprintf(g, "rn += off\n")
		} else {
			fmt.Fprintf(g, "rn -= off\n")
		}
	}

	var name string
	if load {
		if byt {
			fmt.Fprintf(g, "res := uint32(cpu.opRead8(rn))\n")
			name = "ldrb"
		} else {
			fmt.Fprintf(g, "res := cpu.opRead32(rn)\n")
			name = "ldr"
		}
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
		fmt.Fprintf(g, "if rdx == 15 { cpu.Cpsr.SetT((res&1)!=0); cpu.pc = reg(res&^1) }\n")
	} else {
		fmt.Fprintf(g, "rd := cpu.Regs[rdx]\n")
		if byt {
			fmt.Fprintf(g, "cpu.opWrite8(rn, uint8(rd))\n")
			name = "strb"
		} else {
			fmt.Fprintf(g, "cpu.opWrite32(rn, uint32(rd))\n")
			name = "str"
		}
	}

	if !pre {
		if up {
			fmt.Fprintf(g, "rn += off\n")
		} else {
			fmt.Fprintf(g, "rn -= off\n")
		}
		// writeback always enabled for post. wb bit is "force unprivileged"
		if wb {
			g.WriteOpInvalid("forced-unprivileged memory access")
		} else {
			wb = true
		}
	}

	if wb {
		fmt.Fprintf(g, "cpu.Regs[rnx] = reg(rn)\n")
	}

	// disasm
	if pre {
		var off string
		if shreg {
			if up {
				off = "N:(op>>16)&0xF:cpu.disasmOp2(op)"
			} else {
				off = "N:(op>>16)&0xF:\"-\"+cpu.disasmOp2(op)"
			}
		} else {
			if up {
				off = "n:(op>>16)&0xF:int32(op&0xFFF)"
			} else {
				off = "n:(op>>16)&0xF:-int32(op&0xFFF)"
			}
		}
		if wb {
			off += ":!"
		}
		g.WriteDisasm(name, "r:(op>>12)&0xF", off)
	} else {
		var off string
		if shreg {
			if up {
				off = "s:cpu.disasmOp2(op)"
			} else {
				off = "s:\"-\"+cpu.disasmOp2(op)"
			}
		} else {
			if up {
				off = "x:op&0xFFF"
			} else {
				off = "x:-int64(op&0xFFF)"
			}
		}
		g.WriteDisasm(name, "r:(op>>12)&0xF", "l:(op>>16)&0xF", off)
	}
}

func (g *Generator) writeOpHalfWord(op uint32) {
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
		g.WriteOpInvalid("invalid opcode decoded as LD/STR half-word")
		g.WriteDisasmInvalid()
		return
	}

	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := (op>>16)&0xF\n")
	fmt.Fprintf(g, "rdx := (op>>12)&0xF\n")
	fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")
	fmt.Fprintf(g, "cpu.Regs[15]+=4\n")

	var disargs []string
	disargs = append(disargs, "r:(op>>12)&0xF")

	if imm {
		fmt.Fprintf(g, "off := (op&0xF) | ((op&0xF00)>>4)\n")
	} else {
		fmt.Fprintf(g, "rmx := op&0xF\n")
		g.WriteExitIfOpInvalid("rmx==15", "halfword: invalid rm==15")
		fmt.Fprintf(g, "off := uint32(cpu.Regs[rmx])\n")
	}

	if pre {
		if up {
			fmt.Fprintf(g, "rn += off\n")
		} else {
			fmt.Fprintf(g, "rn -= off\n")
		}
	}

	switch code {
	case 1:
		if load {
			fmt.Fprintf(g, "// LDRH\n")
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead16(rn))\n")
			g.WriteDisasm("ldrh", disargs...)
		} else {
			fmt.Fprintf(g, "// STRH\n")
			fmt.Fprintf(g, "cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))\n")
			g.WriteDisasm("strh", disargs...)
		}
	case 2:
		if load {
			fmt.Fprintf(g, "// LDRSB\n")
			fmt.Fprintf(g, "data := int32(int8(cpu.opRead8(rn)))\n")
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(data)\n")
			g.WriteDisasm("ldrsb", disargs...)
		} else {
			fmt.Fprintf(g, "// LDRD\n")
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead32(rn))\n")
			fmt.Fprintf(g, "cpu.Regs[rdx+1] = reg(cpu.opRead32(rn+4))\n")
			g.WriteDisasm("ldrd", disargs...)
		}
	case 3:
		if load {
			fmt.Fprintf(g, "// LDRSH\n")
			fmt.Fprintf(g, "data := int32(int16(cpu.opRead16(rn)))\n")
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(data)\n")
			g.WriteDisasm("ldrsh", disargs...)
		} else {
			fmt.Fprintf(g, "// STRD\n")
			fmt.Fprintf(g, "cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))\n")
			fmt.Fprintf(g, "cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))\n")
			g.WriteDisasm("strd", disargs...)
		}
	}

	if !pre {
		if up {
			fmt.Fprintf(g, "rn += off\n")
		} else {
			fmt.Fprintf(g, "rn -= off\n")
		}
	}

	if wb {
		fmt.Fprintf(g, "cpu.Regs[rnx] = reg(rn)\n")
	}
}

func (g *Generator) writeOpBlock(op uint32) {
	pre := (op>>24)&1 != 0
	up := (op>>23)&1 != 0
	psr := (op>>22)&1 != 0
	wb := (op>>21)&1 != 0
	load := (op>>20)&1 != 0

	var name string
	if load {
		name = "ldm"
		if !(up && !pre) {
			if up {
				name += "i"
			} else {
				name += "d"
			}
			if pre {
				name += "b"
			} else {
				name += "a"
			}
		}
	} else {
		name = "stm"
		if !(up && !pre) {
			if up {
				name += "i"
			} else {
				name += "d"
			}
			if pre {
				name += "b"
			} else {
				name += "a"
			}
		}
	}
	fmt.Fprintf(g, "// %s\n", name)

	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := (op>>16)&0xF\n")
	g.WriteExitIfOpInvalid("rnx==15", "invalid use of PC in LDM/STM")
	fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")
	fmt.Fprintf(g, "mask := uint16(op&0xFFFF)\n")
	if !up {
		fmt.Fprintf(g, "rn -= uint32(4*popcount16(mask))\n")
		if wb {
			fmt.Fprintf(g, "orn := rn\n")
		}
		pre = !pre
	}
	if !load {
		fmt.Fprintf(g, "cpu.Regs[15] += 4\n") // simulate prefetching
	}
	if load && psr {
		fmt.Fprintf(g, "usrbnk := (mask&0x8000)==0\n")
		g.WriteExitIfOpInvalid("usrbnk", "usrbnk not supported")
	}
	fmt.Fprintf(g, "for i:=0; mask != 0; i++ {\n")
	fmt.Fprintf(g, "  if mask&1 != 0 {\n")
	if pre {
		fmt.Fprintf(g, "rn += 4\n")
	}
	if load {
		fmt.Fprintf(g, "val := reg(cpu.opRead32(rn))\n")
		if psr {
			fmt.Fprintf(g, "if usrbnk && i>=8 && i<15 {cpu.UsrBank[i-8]=val} else {cpu.Regs[i]=val}\n")
		} else {
			fmt.Fprintf(g, "cpu.Regs[i] = val\n")
		}
		fmt.Fprintf(g, "if i==15 {\n")
		if psr {
			fmt.Fprintf(g, "cpu.Cpsr.Set(uint32(*cpu.RegSpsr()), cpu)\n")
		}
		fmt.Fprintf(g, "  if cpu.Regs[15]&1 != 0 {cpu.Cpsr.SetT(true); cpu.Regs[15] &^= 1} else {cpu.Regs[15] &^= 3}\n")
		fmt.Fprintf(g, "  cpu.pc = cpu.Regs[15]\n")
		fmt.Fprintf(g, "}\n")
	} else {
		fmt.Fprintf(g, "var val uint32\n")
		if psr {
			// read user bank
			fmt.Fprintf(g, "if i>=8 && i<15 {val=uint32(cpu.UsrBank[i-8])} else {val=uint32(cpu.Regs[i])}\n")
		} else {
			fmt.Fprintf(g, "val = uint32(cpu.Regs[i])\n")
		}
		fmt.Fprintf(g, "cpu.opWrite32(rn, val)\n")
	}
	if !pre {
		fmt.Fprintf(g, "rn += 4\n")
	}
	fmt.Fprintf(g, "  }\n")
	fmt.Fprintf(g, "  mask >>= 1\n")
	fmt.Fprintf(g, "}\n")
	if wb {
		if up {
			fmt.Fprintf(g, "cpu.Regs[rnx] = reg(rn)\n")
		} else {
			fmt.Fprintf(g, "cpu.Regs[rnx] = reg(orn)\n")
		}
	}

	sreg := "r:(op>>16)&0xF"
	if wb {
		sreg += ":!"
	}
	g.WriteDisasm(name, sreg, "k:uint16(op&0xFFFF)")
}

func (g *Generator) writeOpClz(op uint32) {

	fmt.Fprintf(g, "// clz\n")
	g.WriteExitIfOpInvalid("op&0x0FFF0FF0 != 0x016F0F10", "invalid opcode decoded as CLZ")
	g.WriteExitIfOpInvalid("cpu.arch < ARMv5", "invalid CLZ opcode on pre-ARMv5 CPU")
	g.writeOpCond(op)

	fmt.Fprintf(g, "rdx := (op>>12)&0xF\n")
	fmt.Fprintf(g, "rm := cpu.Regs[op&0xF]\n")
	fmt.Fprintf(g, "var lz int\n")
	fmt.Fprintf(g, "for lz=0;lz<32;lz++ {\n")
	fmt.Fprintf(g, "	if int32(rm)<0 { break }\n")
	fmt.Fprintf(g, "	rm <<= 1\n")
	fmt.Fprintf(g, "}\n")
	fmt.Fprintf(g, "cpu.Regs[rdx] = reg(lz)\n")

	g.WriteDisasm("clz", "r:(op>>12)&0xF", "r:op&0xF")
}

func (g *Generator) WriteOp(op uint32) {
	high := (op >> 20) & 0xFF
	low := (op >> 4) & 0xF

	g.WriteOpHeader(int(high<<4 | low))

	// E1C410B0

	switch {
	case high == 0x12 && low&0xD == 0x1:
		g.writeOpBx(op)
	case high == 0x16 && low == 0x1:
		g.writeOpClz(op)
	case (high & 0xFB) == 0x32:
		g.writeOpPsrTransfer(op)
	case (high&0xF9) == 0x10 && low == 0:
		g.writeOpPsrTransfer(op)
	case (high&0xFC) == 0 && low&0x9 == 0x9:
		g.writeOpMul(op)
	case (high&0xF8) == 8 && low&0x9 == 0x9:
		g.writeOpMul(op)
	case (high&0xFB) == 0x10 && low&0x9 == 0x9:
		g.writeOpSwp(op)
	case (high>>5) == 0 && low&0x9 == 9: // TransReg10 / TransImm10
		g.writeOpHalfWord(op)
	case (high>>5) == 0 && low&0x1 == 0:
		g.writeOpAlu(op)
	case (high>>5) == 0 && low&0x9 == 1:
		g.writeOpAlu(op)
	case (high >> 5) == 1:
		g.writeOpAlu(op)
	case (high>>5) == 2 || (high>>5) == 3: // TransImm9 / TransReg9
		g.writeOpMemory(op)
	case (high >> 5) == 4:
		g.writeOpBlock(op)
	case (high >> 5) == 5:
		g.writeOpBranch(op)
	case (high>>5) == 7 && (high>>4)&1 == 0:
		g.writeOpCoprocessor(op)
	case (high>>5) == 7 && (high>>4)&1 == 1:
		g.writeOpSwi(op)
	default:
		g.WriteOpInvalid("unimplemented")
		// panic("unreachable")
	}

	g.WriteOpFooter(int(high<<4 | low))
}

func main() {
	cpugen.Main(func(g *cpugen.Generator) {
		out := Generator{g}
		out.Prefix = "Arm"
		out.OpSize = "uint32"
		out.GenDisasm = true
		out.PcRelOff = 8
		out.TableBits = 12
		out.WriteHeader()
		for op := 0; op < 0x100; op++ {
			for op2 := 0; op2 < 0x10; op2++ {
				out.WriteOp(uint32(op<<20) | uint32(op2<<4))
			}
		}
		out.WriteFooter()
	})
}
