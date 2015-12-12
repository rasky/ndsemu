package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

var filename = flag.String("filename", "-", "output filename")

type Generator struct {
	io.Writer
}

func (g *Generator) WriteHeader() {
	fmt.Fprintf(g, "// Generated on %v\n", time.Now())
	fmt.Fprintf(g, "package arm\n")
	fmt.Fprintf(g, "var opArmTable = [256]func(*Cpu, uint32) {\n")
	for i := 0; i < 256; i++ {
		fmt.Fprintf(g, "(*Cpu).opArm%02X,\n", i)
	}
	fmt.Fprintf(g, "}\n")
}

func (g *Generator) writeOpHeader(op uint32) {
	fmt.Fprintf(g, "func (cpu *Cpu) opArm%02X(op uint32) {\n", (op>>20)&0xFF)
}
func (g *Generator) writeOpFooter(op uint32) {
	fmt.Fprintf(g, "}\n\n")
}

func (g *Generator) writeOpCond(op uint32) {
	fmt.Fprintf(g, "if !cpu.opArmCond(op) { return }\n")
}

func (g *Generator) writeOpInvalid(op uint32, msg string) {
	fmt.Fprintf(g, "cpu.InvalidOpArm(op, %q)\n", msg)
}

func (g *Generator) writeExitIfOpInvalid(cond string, op uint32, msg string) {
	fmt.Fprintf(g, "if %s {\n", cond)
	g.writeOpInvalid(op, msg)
	fmt.Fprintf(g, "return\n}\n")
}

var mulNames = [16]string{
	"MUL", "MLA", "?", "?", "UMULL", "UMLAL", "SMULL", "SMLAL",
	"?", "?", "?", "?", "?", "?", "?", "?",
}

func (g *Generator) writeOpMul(op uint32) {
	setflags := (op>>20)&1 != 0
	code := (op >> 21) & 0xF
	acc := code&1 != 0

	fmt.Fprintf(g, "// %s", mulNames[code])
	if setflags {
		fmt.Fprintf(g, "S")
	}
	fmt.Fprintf(g, "\n")

	if mulNames[code] == "?" {
		g.writeOpInvalid(op, "unhandled mul-type")
		return
	}

	g.writeOpCond(op)
	fmt.Fprintf(g, "rsx := (op >> 8) & 0xF\n")
	fmt.Fprintf(g, "rs := uint32(cpu.Regs[rsx])\n")

	fmt.Fprintf(g, "rmx := (op >> 0) & 0xF\n")
	fmt.Fprintf(g, "rm := uint32(cpu.Regs[rmx])\n")

	fmt.Fprintf(g, "data := (op >> 4) & 0xF;\n")
	g.writeExitIfOpInvalid("data!=0x9", op, "unimplemented half-word multiplies")

	fmt.Fprintf(g, "rdx := (op >> 16) & 0xF\n")

	switch code {
	case 0: // MUL
		fmt.Fprintf(g, "res := rm*rs\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
		}
	case 1: // MLA
		fmt.Fprintf(g, "rnx := (op >> 12) & 0xF\n")
		fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")
		fmt.Fprintf(g, "res := rm*rs+rn\n")
		if setflags {
			fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
		}
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
	default:
		panic("unreachable")
	}

	fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
}

func (g *Generator) writeOpBx(op uint32) {
	if op>>20 != 0x12 {
		fmt.Fprintf(os.Stderr, "OP BX ERR: %08x\n", op)
	}
	fmt.Fprintf(g, "// BX / BLX_reg\n")
	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := op&0xF\n")
	fmt.Fprintf(g, "rn := cpu.Regs[rnx]\n")
	fmt.Fprintf(g, "if op&0x20 != 0 { cpu.Regs[14] = cpu.Regs[15]-4 }\n")
	fmt.Fprintf(g, "if rn&1 != 0 { cpu.Cpsr.SetT(true); rn &^= 1 }\n")
	fmt.Fprintf(g, "cpu.pc = rn\n")
}

func (g *Generator) writeOpPsrTransfer(op uint32) {
	imm := (op>>25)&1 != 0
	if (op>>26)&3 != 0 || (op>>23)&0x3 != 2 || (op>>20)&1 != 0 {
		panic("invalid psr decoding")
	}
	spsr := (op>>22)&1 != 0
	tostat := (op>>21)&1 != 0

	if !tostat {
		fmt.Fprintf(g, "// MRS\n")
		g.writeOpCond(op)
		fmt.Fprintf(g, "mask := (op>>16)&0xF\n")
		g.writeExitIfOpInvalid("mask != 0xF", op, "unimplemented SWP")
		fmt.Fprintf(g, "rdx := (op>>12)&0xF\n")
		g.writeExitIfOpInvalid("rdx == 15", op, "write to PC in MRS")
		if spsr {
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(*cpu.RegSpsr())")
		} else {
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.Cpsr.Uint32())")
		}
	} else {
		if !imm && (op>>20)&0xFF == 0x12 {
			fmt.Fprintf(g, "if (op>>4)&0xFF != 0 {")
			g.writeOpBx(op)
			fmt.Fprintf(g, "return\n}\n")
		}

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
			g.writeExitIfOpInvalid("(op>>4)&0xFF != 0", op, "unimplemented BX")
			fmt.Fprintf(g, "rmx := op & 0xF\n")
			fmt.Fprintf(g, "val := uint32(cpu.Regs[rmx])\n")
		}

		if spsr {
			fmt.Fprintf(g, "cpu.RegSpsr().SetWithMask(val, mask)\n")
		} else {
			fmt.Fprintf(g, "cpu.Cpsr.SetWithMask(val, mask, cpu)\n")
		}
	}
}

var aluNames = [16]string{
	"AND", "EOR", "SUB", "RSB", "ADD", "ADC", "SBC", "RSC",
	"TST", "TEQ", "CMP", "CMN", "ORR", "MOV", "BIC", "MVN",
}

func (g *Generator) writeOpAlu(op uint32) {

	imm := (op>>25)&1 != 0
	code := (op >> 21) & 0xF
	setflags := (op>>20)&1 != 0

	if code >= 8 && code < 12 && !setflags {
		g.writeOpPsrTransfer(op)
		return
	}

	fmt.Fprintf(g, "// %s", aluNames[code])
	if setflags {
		fmt.Fprintf(g, "S")
	}
	fmt.Fprintf(g, "\n")

	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := (op >> 16) & 0xF\n")
	fmt.Fprintf(g, "rdx := (op >> 12) & 0xF\n")

	if imm {
		fmt.Fprintf(g, "rot := uint((op>>7)&0x1E)\n")
		fmt.Fprintf(g, "op2 := ((op&0xFF)>>rot) | ((op&0xFF)<<(32-rot))\n")
	} else {
		fmt.Fprintf(g, "op2 := cpu.opDecodeAluOp2Reg(op, true)\n")
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
		g.writeExitIfOpInvalid("rnx!=0", op, "rn!=0 on NOV")
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
			panic("this should be a psr transfer")
		}
		g.writeExitIfOpInvalid("rdx != 0 && rdx != 15", op, "invalid rdx on test")
	}

	fmt.Fprintf(g, "_ = res; _ = rn\n")
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
		fmt.Fprintf(g, "cpu.Regs[15] += 2\n")
	}
	fmt.Fprintf(g, "cpu.pc = cpu.Regs[15]\n")
	fmt.Fprintf(g, "cpu.Cpsr.SetT(true)\n")
	fmt.Fprintf(g, "return\n}\n")

	if link {
		fmt.Fprintf(g, "// BL\n")
	} else {
		fmt.Fprintf(g, "// B\n")
	}
	g.writeOpCond(op)
	g.writeOpBranchInner(link)
	fmt.Fprintf(g, "cpu.pc = cpu.Regs[15]\n")
}

func (g *Generator) writeOpCoprocessor(op uint32) {
	if op&(1<<24) == 0 {
		copread := (op>>20)&1 != 0
		if copread {
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

		fmt.Fprintf(g, "if (op&0x10) == 0 { // CDP\n")
		fmt.Fprintf(g, "cpu.opCopExec(copnum, opc, cn, cm, cp, rdx)\n")
		fmt.Fprintf(g, "return\n}\n")

		if copread {
			fmt.Fprintf(g, "res := cpu.opCopRead(copnum, opc, cn, cm, cp)\n")
			fmt.Fprintf(g, "if rdx==15 { cpu.Cpsr.SetWithMask(res, 0xF0000000, cpu) }")
			fmt.Fprintf(g, "else { cpu.Regs[rdx] = reg(res) }")
		} else {
			fmt.Fprintf(g, "cpu.Regs[15]+=4\n")
			fmt.Fprintf(g, "rd := cpu.Regs[rdx]\n")
			fmt.Fprintf(g, "cpu.opCopWrite(copnum, opc, cn, cm, cp, uint32(rd))\n")
		}

	} else {
		g.writeOpInvalid(op, "unhandled cop")
	}
}

func (g *Generator) writeOpMemory(op uint32) {
	shreg := (op>>25)&1 != 0
	pre := (op>>24)&1 != 0
	up := (op>>23)&1 != 0
	byt := (op>>22)&1 != 0
	wb := (op>>21)&1 != 0
	load := (op>>20)&1 != 0

	g.writeExitIfOpInvalid("(op>>28)==0xF", op, "PLD not supported")
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

	if load {
		if byt {
			fmt.Fprintf(g, "res := uint32(cpu.opRead8(rn))\n")
		} else {
			fmt.Fprintf(g, "res := cpu.opRead32(rn)\n")
		}
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
		fmt.Fprintf(g, "if rdx == 15 { cpu.Cpsr.SetT((res&1)!=0); cpu.pc = reg(res&^1) }\n")
	} else {
		fmt.Fprintf(g, "rd := cpu.Regs[rdx]\n")
		if byt {
			fmt.Fprintf(g, "cpu.opWrite8(rn, uint8(rd))\n")
		} else {
			fmt.Fprintf(g, "cpu.opWrite32(rn, uint32(rd))\n")
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
			g.writeOpInvalid(op, "forced-unprivileged memory access")
		} else {
			wb = true
		}
	}

	if wb {
		fmt.Fprintf(g, "cpu.Regs[rnx] = reg(rn)\n")
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

	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := (op>>16)&0xF\n")
	fmt.Fprintf(g, "rdx := (op>>12)&0xF\n")
	fmt.Fprintf(g, "rn := uint32(cpu.Regs[rnx])\n")
	fmt.Fprintf(g, "cpu.Regs[15]+=4\n")

	if imm {
		fmt.Fprintf(g, "off := (op&0xF) | ((op&0xF00)>>4)\n")
	} else {
		fmt.Fprintf(g, "rmx := op&0xF\n")
		g.writeExitIfOpInvalid("rmx==15", op, "halfword: invalid rm==15")
		fmt.Fprintf(g, "off := uint32(cpu.Regs[rmx])\n")
	}

	if pre {
		if up {
			fmt.Fprintf(g, "rn += off\n")
		} else {
			fmt.Fprintf(g, "rn -= off\n")
		}
	}

	fmt.Fprintf(g, "code := (op>>5)&3\n")
	fmt.Fprintf(g, "switch code {\n")

	fmt.Fprintf(g, "case 1:\n")
	if load {
		fmt.Fprintf(g, "// LDRH\n")
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead16(rn))\n")
	} else {
		fmt.Fprintf(g, "// STRH\n")
		fmt.Fprintf(g, "cpu.opWrite16(rn, uint16(cpu.Regs[rdx]))\n")
	}

	fmt.Fprintf(g, "case 2:\n")
	if load {
		fmt.Fprintf(g, "// LDRSB\n")
		fmt.Fprintf(g, "data := int32(int8(cpu.opRead8(rn)))\n")
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(data)\n")
	} else {
		fmt.Fprintf(g, "// LDRD\n")
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead32(rn))\n")
		fmt.Fprintf(g, "cpu.Regs[rdx+1] = reg(cpu.opRead32(rn+4))\n")
	}

	fmt.Fprintf(g, "case 3:\n")
	if load {
		fmt.Fprintf(g, "// LDRSH\n")
		fmt.Fprintf(g, "data := int32(int16(cpu.opRead16(rn)))\n")
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(data)\n")
	} else {
		fmt.Fprintf(g, "// STRD\n")
		fmt.Fprintf(g, "cpu.opWrite32(rn, uint32(cpu.Regs[rdx]))\n")
		fmt.Fprintf(g, "cpu.opWrite32(rn+4, uint32(cpu.Regs[rdx+1]))\n")
	}

	fmt.Fprintf(g, "default:\n")
	g.writeOpInvalid(op, "halfword invalid op")

	fmt.Fprintf(g, "}\n")

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

	if load {
		if up && !pre {
			fmt.Fprintf(g, "// POP\n")
		} else {
			fmt.Fprintf(g, "// LDM")
			if up {
				fmt.Fprintf(g, "I")
			} else {
				fmt.Fprintf(g, "D")
			}
			if pre {
				fmt.Fprintf(g, "B\n")
			} else {
				fmt.Fprintf(g, "A\n")
			}
		}
	} else {
		if !up && pre {
			fmt.Fprintf(g, "// PUSH\n")
		} else {
			fmt.Fprintf(g, "// STM")
			if up {
				fmt.Fprintf(g, "I")
			} else {
				fmt.Fprintf(g, "D")
			}
			if pre {
				fmt.Fprintf(g, "B\n")
			} else {
				fmt.Fprintf(g, "A\n")
			}
		}
	}

	g.writeOpCond(op)
	fmt.Fprintf(g, "rnx := (op>>16)&0xF\n")
	g.writeExitIfOpInvalid("rnx==15", op, "invalid use of PC in LDM/STM")
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
		fmt.Fprintf(g, "  if cpu.Regs[15]&1 != 0 { cpu.Cpsr.SetT(true); cpu.Regs[15] = cpu.Regs[15] &^ 1 }\n")
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
}

func (g *Generator) WriteOp(op uint32) {
	g.writeOpHeader(op)

	switch (op >> 25) & 0x7 {
	case 0:
		fmt.Fprintf(g, "if op&0x90 == 0x90 {\n")
		fmt.Fprintf(g, "	if op&0x60 != 0 {\n")
		g.writeOpHalfWord(op)
		fmt.Fprintf(g, "	} else {\n")
		g.writeOpMul(op)
		fmt.Fprintf(g, "	}\n")
		fmt.Fprintf(g, "} else {\n")
		g.writeOpAlu(op)
		fmt.Fprintf(g, "}\n")
	case 1:
		g.writeOpAlu(op)
	case 2, 3:
		g.writeOpMemory(op)
	case 4:
		g.writeOpBlock(op)
	case 5:
		g.writeOpBranch(op)
	case 7:
		g.writeOpCoprocessor(op)
	default:
		g.writeOpInvalid(op, "unimplemented")
	}

	g.writeOpFooter(op)
}

func main() {
	flag.Parse()

	var f io.Writer
	if *filename == "-" {
		f = os.Stdout
	} else {
		ff, err := os.Create(*filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func() {
			cmd := exec.Command("go", "fmt", *filename)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				os.Exit(1)
			}
		}()
		defer ff.Close()
		f = ff
	}

	out := Generator{f}
	out.WriteHeader()
	for op := 0; op < 0x100; op++ {
		out.WriteOp(uint32(op << 20))
	}
}
