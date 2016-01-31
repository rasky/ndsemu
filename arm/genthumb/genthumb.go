package main

import (
	"fmt"
	"ndsemu/emu/cpugen"
	"strings"
)

type Generator struct {
	*cpugen.Generator
}

func (g *Generator) writeCycles(cycles int) {
	fmt.Fprintf(g, "cpu.Clock += %d\n", cycles)
}

func (g *Generator) writeBranch(target string, reason string) {
	fmt.Fprintf(g, "cpu.branch(%s, %s)\n", target, reason)
}

func (g *Generator) WriteHeader() {
	g.Generator.WriteHeader()

	fmt.Fprintf(g, "var opThumbAluTable = [16]func(*Cpu, uint16) {\n")
	for i := 0; i < 16; i++ {
		fmt.Fprintf(g, "(*Cpu).opThumbAlu%02X,\n", i)
	}
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(g, "var disasmThumbAluTable = [16]func(*Cpu, uint16, uint32) string {\n")
	for i := 0; i < 16; i++ {
		fmt.Fprintf(g, "(*Cpu).disasmThumbAlu%02X,\n", i)
	}
	fmt.Fprintf(g, "}\n")
}

func (g *Generator) writeOpAluHeader(op uint16) {
	fmt.Fprintf(g, "func (cpu *Cpu) opThumbAlu%02X(op uint16) {\n", (op>>6)&0xF)
	g.Disasm.Reset()
}
func (g *Generator) writeOpAluFooter(op uint16) {
	fmt.Fprintf(g, "}\n\n")
	if g.Disasm.Len() == 0 {
		// panic(fmt.Sprintf("disasm not implemented for op %04x", op))
		return
	}

	fmt.Fprintf(g, "func (cpu *Cpu) disasmThumbAlu%02X(op uint16, pc uint32) string {\n", (op>>6)&0xF)
	fmt.Fprintf(g, g.Disasm.String())
	fmt.Fprintf(g, "}\n\n")
}

func (g *Generator) writeBeginArchSwitch() {
	fmt.Fprintf(g, "switch cpu.arch {\n")
}

func (g *Generator) writeCaseArchSwitch(arch string) {
	fmt.Fprintf(g, "case %s:\n", arch)
}

func (g *Generator) writeEndArchSwitch() {
	fmt.Fprintf(g, "default: panic(\"unimplemented arch-dependent behavior\")\n")
	fmt.Fprintf(g, "}\n")
}

var regnames = []string{
	"r0", "r1", "r2", "r3", "r4", "r5", "r6", "r7",
	"r8", "r9", "r10", "r11", "r12", "sp", "lr", "pc",
}

var f1name = [3]string{"lsl", "lsr", "asr"}

func (g *Generator) writeOpF1Shift(op uint16) {
	opcode := (op >> 11) & 3

	fmt.Fprintf(g, "// %s\n", f1name[opcode])

	fmt.Fprintf(g, "rsx := (op>>3)&7\n")
	fmt.Fprintf(g, "rdx := op&7\n")
	fmt.Fprintf(g, "offset := (op>>6)&0x1F\n")
	fmt.Fprintf(g, "rs := uint32(cpu.Regs[rsx])\n")

	g.WriteDisasm(f1name[opcode], "r:op&7", "r:(op>>3)&7", "d:(op>>6)&0x1F")

	switch opcode {
	case 0: // LSL
		fmt.Fprintf(g, "if offset != 0 { cpu.Cpsr.SetC(rs & (1<<(32-offset)) != 0) }\n")
		fmt.Fprintf(g, "res := rs << offset\n")
	case 1: // LSR
		fmt.Fprintf(g, "if offset == 0 { offset = 32 }\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(rs & (1<<(offset-1)) != 0)\n")
		fmt.Fprintf(g, "res := rs >> offset\n")
	case 2: // ASR
		fmt.Fprintf(g, "if offset == 0 { offset = 32 }\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(rs & (1<<(offset-1)) != 0)\n")
		fmt.Fprintf(g, "res := uint32(int32(rs) >> offset)\n")
	default:
		panic("unreachable")
	}

	fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
	fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
}

var f2name = [4]string{"add", "sub", "add #nn", "sub #nn"}

func (g *Generator) writeOpF2Add(op uint16) {
	opcode := (op >> 9) & 3
	imm := opcode&2 != 0

	fmt.Fprintf(g, "// %s\n", f2name[opcode])

	fmt.Fprintf(g, "rsx := (op>>3)&7\n")
	fmt.Fprintf(g, "rdx := op&7\n")
	fmt.Fprintf(g, "rs := uint32(cpu.Regs[rsx])\n")

	if imm {
		g.WriteDisasm(f2name[opcode&1], "r:op&7", "r:(op>>3)&7", "d:(op>>6)&7")
		fmt.Fprintf(g, "val := uint32((op>>6)&7)\n")
	} else {
		g.WriteDisasm(f2name[opcode&1], "r:op&7", "r:(op>>3)&7", "r:(op>>6)&7")
		fmt.Fprintf(g, "rnx := (op>>6)&7\n")
		fmt.Fprintf(g, "val := uint32(cpu.Regs[rnx])\n")
	}

	switch opcode {
	case 0, 2: // ADD
		fmt.Fprintf(g, "res := rs + val\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(res<rs)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVAdd(rs, val, res)\n")
	case 1, 3: // SUB
		fmt.Fprintf(g, "res := rs - val\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(rs>=val)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVSub(rs, val, res)\n")
	}

	fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
	fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
}

var f3name = [4]string{"mov", "cmp", "add", "sub"}

func (g *Generator) writeOpF3AluImm(op uint16) {
	opcode := (op >> 11) & 3
	rdx := (op >> 8) & 7

	fmt.Fprintf(g, "// %s\n", f3name[opcode])
	g.WriteDisasm(f3name[opcode], "r:(op>>8)&7", "x:op&0xFF")

	test := false
	fmt.Fprintf(g, "imm := uint32(op&0xFF)\n")
	switch opcode {
	case 0: // MOV
		fmt.Fprintf(g, "res := imm\n")
	case 2: // ADD
		fmt.Fprintf(g, "rd := uint32(cpu.Regs[%d])\n", rdx)
		fmt.Fprintf(g, "res := rd + imm\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(res<rd)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVAdd(rd, imm, res)\n")
	case 1: // CMP
		test = true
		fallthrough
	case 3: // SUB
		fmt.Fprintf(g, "rd := uint32(cpu.Regs[%d])\n", rdx)
		fmt.Fprintf(g, "res := rd - imm\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(rd>=imm)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVSub(rd, imm, res)\n")
	default:
		panic("unreachable")
	}
	fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
	if !test {
		fmt.Fprintf(g, "cpu.Regs[%d] = reg(res)\n", rdx)
	}
}

func (g *Generator) writeOpF4Alu(op uint16) {
	// F4 is the only format of opcodes where the real opcode is encoded in
	// bits below the 8th, so our dispatch table can't properly differentiate
	// between all instructions. Instead of doing all the decoding at runtime,
	// we do a second-level dispatching:
	fmt.Fprintf(g, "opThumbAluTable[(op>>6)&0xF](cpu, op)\n")
	fmt.Fprintf(&g.Disasm, "return disasmThumbAluTable[(op>>6)&0xF](cpu, op, pc)\n")
}

var f5name = [4]string{"add(h)", "cmp(h)", "mov(h)", "bx/blx"}

func (g *Generator) writeOpF5HiReg(op uint16) {
	opcode := (op >> 8) & 3

	fmt.Fprintf(g, "// %s\n", f5name[opcode])
	fmt.Fprintf(g, "rdx := (op&7) | (op&0x80)>>4\n")
	fmt.Fprintf(g, "rsx := ((op>>3)&0xF)\n")
	fmt.Fprintf(g, "rs := uint32(cpu.Regs[rsx])\n")

	switch opcode {
	case 0: // ADD
		// NOTE: this does not affect flags (!)
		fmt.Fprintf(g, "rd := uint32(cpu.Regs[rdx])\n")
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(rd+rs)\n")
		fmt.Fprintf(g, "if rdx==15 { \n")
		g.writeBranch("    cpu.Regs[15]&^1", "BranchJump")
		fmt.Fprintf(g, "}\n")
		g.WriteDisasm("add", "r:(op&7) | (op&0x80)>>4", "r:((op>>3)&0xF)")
	case 1: // CMP
		fmt.Fprintf(g, "rd := uint32(cpu.Regs[rdx])\n")
		fmt.Fprintf(g, "res := rd-rs\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(rd>=res)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVSub(rd, rs, res)\n")
		g.WriteDisasm("cmp", "r:(op&7) | (op&0x80)>>4", "r:((op>>3)&0xF)")
	case 2: // MOV
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(rs)\n")
		fmt.Fprintf(g, "if rdx==15 {\n")
		g.writeBranch("   reg(rs)&^1", "BranchJump")
		fmt.Fprintf(g, "}\n")
		g.WriteDisasm("mov", "r:(op&7) | (op&0x80)>>4", "r:((op>>3)&0xF)")
	case 3: // BX/BLX
		fmt.Fprintf(g, "if op&0x80 != 0 { cpu.Regs[14] = (cpu.Regs[15]-2)|1 }\n")
		fmt.Fprintf(g, "newpc := reg(rs)&^1\n")
		fmt.Fprintf(g, "if rs&1==0 { cpu.Cpsr.SetT(false); newpc &^= 3 }\n")
		g.writeBranch("newpc", "BranchCall")
		fmt.Fprintf(g, "_=rdx\n")

		fmt.Fprintf(&g.Disasm, "if op&0x80 != 0 {\n")
		g.WriteDisasm("blx", "r:((op>>3)&0xF)")
		fmt.Fprintf(&g.Disasm, "} else {\n")
		g.WriteDisasm("bx", "r:((op>>3)&0xF)")
		fmt.Fprintf(&g.Disasm, "}\n")
	default:
		panic("unreachable")
	}
}

func (g *Generator) writeOpF6LdrPc(op uint16) {
	rdx := (op >> 8) & 7
	fmt.Fprintf(g, "// ldr pc\n")
	fmt.Fprintf(g, "pc := uint32(cpu.Regs[15]) &^ 2\n")
	fmt.Fprintf(g, "pc += uint32((op & 0xFF)*4)\n")
	fmt.Fprintf(g, "cpu.Regs[%d] = reg(cpu.opRead32(pc))\n", rdx)
	g.writeCycles(1)
	g.WriteDisasm("ldr", "r:(op>>8)&7", "P:(op & 0xFF)*4")
}

var f7name = [4]string{"str", "strb", "ldr", "ldrb"}
var f8name = [4]string{"strh", "ldsb", "ldrh", "ldsh"}

func (g *Generator) writeOpF7F8LdrStr(op uint16) {
	opcode := (op >> 10) & 3
	f8 := op&(1<<9) != 0

	name := f7name[opcode]
	if f8 {
		name = f8name[opcode]
	}
	fmt.Fprintf(g, "// %s\n", name)
	g.WriteDisasm(name, "r:op&7", "m:(op>>3)&7:(op>>6)&7")

	fmt.Fprintf(g, "rox := (op>>6)&7\n")
	fmt.Fprintf(g, "rbx := (op>>3)&7\n")
	fmt.Fprintf(g, "rdx := op&7\n")
	fmt.Fprintf(g, "addr := uint32(cpu.Regs[rbx] + cpu.Regs[rox])\n")

	if !f8 {
		switch opcode {
		case 0: // STR
			fmt.Fprintf(g, "cpu.opWrite32(addr, uint32(cpu.Regs[rdx]))\n")
		case 1: // STRB
			fmt.Fprintf(g, "cpu.opWrite8(addr, uint8(cpu.Regs[rdx]))\n")
		case 2: // LDR
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead32(addr))\n")
		case 3: // LDRB
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead8(addr))\n")
		default:
			panic("unreachable")
		}
	} else {
		switch opcode {
		case 0: // STRH
			fmt.Fprintf(g, "cpu.opWrite16(addr, uint16(cpu.Regs[rdx]))\n")
		case 1: // LDSB
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(int8(cpu.opRead8(addr)))\n")
		case 2: // LDRH
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead16(addr))\n")
		case 3: // LDSH
			fmt.Fprintf(g, "cpu.Regs[rdx] = reg(int16(cpu.opRead16(addr)))\n")
		default:
			panic("unreachable")
		}
	}
	g.writeCycles(1)
}

var f9name = [4]string{"str #nn", "ldr #nn", "strb #nn", "ldrb #nn"}

func (g *Generator) writeOpF9Strb(op uint16) {
	opcode := (op >> 11) & 3
	fmt.Fprintf(g, "// %s\n", f9name[opcode])
	fmt.Fprintf(g, "offset := uint32((op>>6)&0x1F)\n")
	fmt.Fprintf(g, "rbx := (op>>3)&0x7\n")
	fmt.Fprintf(g, "rdx := op&0x7\n")
	fmt.Fprintf(g, "rb := uint32(cpu.Regs[rbx])\n")
	switch opcode {
	case 0: // STR
		fmt.Fprintf(g, "offset *= 4\n")
		fmt.Fprintf(g, "rd := uint32(cpu.Regs[rdx])\n")
		fmt.Fprintf(g, "cpu.opWrite32(rb+offset, rd)\n")
	case 1: // LDR
		fmt.Fprintf(g, "offset *= 4\n")
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead32(rb+offset))\n")
	case 2: // STRB
		fmt.Fprintf(g, "rd := uint8(cpu.Regs[rdx])\n")
		fmt.Fprintf(g, "cpu.opWrite8(rb+offset, rd)\n")
	case 3: // LDRB
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead8(rb+offset))\n")
	default:
		panic("unreachable")
	}
	g.writeCycles(1)
	idx := strings.IndexByte(f9name[opcode], ' ')
	if opcode == 0 || opcode == 1 {
		g.WriteDisasm(f9name[opcode][:idx], "r:op&7", "n:(op>>3)&0x7:((op>>6)&0x1F)*4")
	} else {
		g.WriteDisasm(f9name[opcode][:idx], "r:op&7", "n:(op>>3)&0x7:(op>>6)&0x1F")
	}
}

var f10name = [2]string{"strh #nn", "ldrh #nn"}

func (g *Generator) writeOpF10Strh(op uint16) {
	opcode := (op >> 11) & 1
	fmt.Fprintf(g, "// %s\n", f10name[opcode])
	fmt.Fprintf(g, "offset := uint32((op>>6)&0x1F)\n")
	fmt.Fprintf(g, "rbx := (op>>3)&0x7\n")
	fmt.Fprintf(g, "rdx := op&0x7\n")
	fmt.Fprintf(g, "rb := uint32(cpu.Regs[rbx])\n")

	switch opcode {
	case 0: // STRH
		fmt.Fprintf(g, "offset *= 2\n")
		fmt.Fprintf(g, "rd := uint16(cpu.Regs[rdx])\n")
		fmt.Fprintf(g, "cpu.opWrite16(rb+offset, rd)\n")
	case 1: // LDRH
		fmt.Fprintf(g, "offset *= 2\n")
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(cpu.opRead16(rb+offset))\n")
	default:
		panic("unreachable")
	}
	g.writeCycles(1)
	idx := strings.IndexByte(f10name[opcode], ' ')
	g.WriteDisasm(f10name[opcode][:idx], "r:op&7", "n:(op>>3)&0x7:((op>>6)&0x1F)*2")
}

var f11name = [2]string{"str [sp+nn]", "ldr [sp+nn]"}

func (g *Generator) writeOpF11Strsp(op uint16) {
	opcode := (op >> 11) & 1
	rdx := (op >> 8) & 0x7
	fmt.Fprintf(g, "// %s\n", f11name[opcode])
	fmt.Fprintf(g, "offset := (op&0xFF)*4\n")
	fmt.Fprintf(g, "sp := uint32(cpu.Regs[13])\n")
	switch opcode {
	case 0: // STR
		fmt.Fprintf(g, "cpu.opWrite32(sp+uint32(offset), uint32(cpu.Regs[%d]))\n", rdx)
	case 1: // LDR
		fmt.Fprintf(g, "cpu.Regs[%d] = reg(cpu.opRead32(sp+uint32(offset)))\n", rdx)
	default:
		panic("unreachable")
	}
	g.writeCycles(1)
	idx := strings.IndexByte(f11name[opcode], ' ')
	g.WriteDisasm(f11name[opcode][:idx], "r:(op>>8)&7", "n:13:(op&0xFF)*4")
}

var f12name = [2]string{"add pc", "add sp"}

func (g *Generator) writeOpF12AddPc(op uint16) {
	opcode := (op >> 11) & 1
	rdx := (op >> 8) & 0x7
	fmt.Fprintf(g, "// %s\n", f12name[opcode])
	fmt.Fprintf(g, "offset := (op&0xFF)*4\n")
	switch opcode {
	case 0: // ADD PC
		fmt.Fprintf(g, "cpu.Regs[%d] = (cpu.Regs[15] &^ 2) + reg(offset)\n", rdx)
		g.WriteDisasm("add", "r:(op>>8)&7", "r:15", "x:(op&0xFF)*4")
	case 1: // ADD SP
		fmt.Fprintf(g, "cpu.Regs[%d] = cpu.Regs[13] + reg(offset)\n", rdx)
		g.WriteDisasm("add", "r:(op>>8)&7", "r:13", "x:(op&0xFF)*4")
	default:
		panic("unreachable")
	}
}

func (g *Generator) writeOpF13AddSp(op uint16) {
	fmt.Fprintf(g, "// add sp\n")

	fmt.Fprintf(g, "offset := (op&0x7F)*4\n")
	fmt.Fprintf(g, "if op&0x80 == 0 {\n")
	fmt.Fprintf(g, "  cpu.Regs[13] += reg(offset)\n")
	fmt.Fprintf(g, "} else {\n")
	fmt.Fprintf(g, "  cpu.Regs[13] -= reg(offset)\n")
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(&g.Disasm, "if op&0x80 == 0 {\n")
	g.WriteDisasm("add", "r:13", "x:(op&0x7f)*4")
	fmt.Fprintf(&g.Disasm, "} else {\n")
	g.WriteDisasm("sub", "r:13", "x:(op&0x7f)*4")
	fmt.Fprintf(&g.Disasm, "}\n")
}

func (g *Generator) writeOpF14PushPop(op uint16) {
	pop := (op>>11)&1 != 0

	if pop {
		fmt.Fprintf(g, "// pop\n")
	} else {
		fmt.Fprintf(g, "// push\n")
		fmt.Fprintf(g, "count := popcount16(op&0x1FF)\n")
	}

	fmt.Fprintf(g, "sp := uint32(cpu.Regs[13])\n")
	if !pop {
		fmt.Fprintf(g, "sp -= uint32(count*4)\n")
		fmt.Fprintf(g, "cpu.Regs[13] = reg(sp)\n")
	}

	for i := 0; i < 9; i++ {
		fmt.Fprintf(g, "if (op>>%d)&1 != 0 {\n", i)
		regnum := i
		if i == 8 {
			if pop {
				regnum = 15
			} else {
				regnum = 14
			}
		}
		if pop {
			if regnum == 15 {
				g.writeBeginArchSwitch()

				g.writeCaseArchSwitch("ARMv4")
				fmt.Fprintf(g, "  pc := reg(cpu.opRead32(sp) &^ 1)\n")
				g.writeBranch("pc", "BranchReturn")

				g.writeCaseArchSwitch("ARMv5")
				fmt.Fprintf(g, "  pc := reg(cpu.opRead32(sp))\n")
				fmt.Fprintf(g, "  if pc&1 == 0 { cpu.Cpsr.SetT(false); pc = pc&^3 } else { pc = pc&^1 }\n")
				g.writeBranch("pc", "BranchReturn")

				g.writeEndArchSwitch()

			} else {
				fmt.Fprintf(g, "  cpu.Regs[%d] = reg(cpu.opRead32(sp))\n", regnum)
			}

		} else {
			fmt.Fprintf(g, "  cpu.opWrite32(sp, uint32(cpu.Regs[%d]))\n", regnum)
		}
		fmt.Fprintf(g, "  sp += 4\n")
		fmt.Fprintf(g, "}\n")
	}

	if pop {
		fmt.Fprintf(g, "cpu.Regs[13] = reg(sp)\n")
	}
	g.writeCycles(1)

	if pop {
		if (op>>8)&1 != 0 {
			g.WriteDisasm("pop", "k:op&0xff|0x8000")
		} else {
			g.WriteDisasm("pop", "k:op&0xff")
		}
	} else {
		if (op>>8)&1 != 0 {
			g.WriteDisasm("push", "k:op&0xff|0x4000")
		} else {
			g.WriteDisasm("push", "k:op&0xff")
		}
	}
}

func (g *Generator) writeOpF15LdmStm(op uint16) {
	load := (op>>11)&1 != 0
	rbx := (op >> 8) & 0x7

	if load {
		fmt.Fprintf(g, "// ldm\n")
	} else {
		fmt.Fprintf(g, "// stm\n")
		g.WriteExitIfOpInvalid(fmt.Sprintf("op&(1<<%d) != 0", rbx), "unimplemented: base reg in register list in STM")
	}

	fmt.Fprintf(g, "ptr := uint32(cpu.Regs[%d])\n", rbx)

	fmt.Fprintf(g, "if op&0xFF==0 {\n")
	g.writeBeginArchSwitch()

	g.writeCaseArchSwitch("ARMv4")
	if load {
		fmt.Fprintf(g, "  cpu.Regs[15] = reg(cpu.opRead32(ptr))\n")
	} else {
		fmt.Fprintf(g, "  cpu.opWrite32(ptr, uint32(cpu.Regs[15]))\n")
	}
	fmt.Fprintf(g, "ptr+=0x40\n")

	g.writeCaseArchSwitch("ARMv5")
	fmt.Fprintf(g, "ptr+=0x40\n")

	g.writeEndArchSwitch()
	fmt.Fprintf(g, "cpu.Regs[%d] = reg(ptr)\n", rbx)
	fmt.Fprintf(g, "return\n")
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(g, "wb := true\n")
	for i := 0; i < 8; i++ {
		fmt.Fprintf(g, "if (op>>%d)&1 != 0 {\n", i)
		regnum := i
		if load {
			fmt.Fprintf(g, "  cpu.Regs[%d] = reg(cpu.opRead32(ptr))\n", regnum)
			if regnum == int(rbx) {
				fmt.Fprintf(g, "wb = false\n")
			}
		} else {
			fmt.Fprintf(g, "  cpu.opWrite32(ptr, uint32(cpu.Regs[%d]))\n", regnum)
		}
		fmt.Fprintf(g, "  ptr += 4\n")
		fmt.Fprintf(g, "}\n")
	}

	fmt.Fprintf(g, "if wb { cpu.Regs[%d] = reg(ptr) }\n", rbx)
	g.writeCycles(1)
	if load {
		g.WriteDisasm("ldm", "R:(op>>8)&7:(op>>((op>>8)&7))&1", "k:op&0xFF")
	} else {
		g.WriteDisasm("stm", "R:(op>>8)&7:(op>>((op>>8)&7))&1", "k:op&0xFF")
	}
}

var f16name = [16]string{
	"beq", "bne", "bhs", "blo", "bmi", "bpl", "bvs", "bvc",
	"bhi", "bls", "bge", "blt", "bgt", "ble", "b undef", "swi",
}

var f16cond = [14]string{
	"cpu.Cpsr.Z()",  // BEQ
	"!cpu.Cpsr.Z()", // BNE
	"cpu.Cpsr.C()",  // BHS
	"!cpu.Cpsr.C()", // BLO
	"cpu.Cpsr.N()",  // BMI
	"!cpu.Cpsr.N()", // BPL
	"cpu.Cpsr.V()",  // BVS
	"!cpu.Cpsr.V()", // BVC

	"cpu.Cpsr.C() && !cpu.Cpsr.Z()", // BHI
	"!cpu.Cpsr.C() || cpu.Cpsr.Z()", // BLS

	"cpu.Cpsr.N() == cpu.Cpsr.V()", // BGE
	"cpu.Cpsr.N() != cpu.Cpsr.V()", // BLT

	"!cpu.Cpsr.Z() && cpu.Cpsr.N() == cpu.Cpsr.V()", // BGT
	"cpu.Cpsr.Z() || cpu.Cpsr.N() != cpu.Cpsr.V()",  // BLE
}

func (g *Generator) writeOpF16BranchCond(op uint16) {
	opcode := (op >> 8) & 0xF

	fmt.Fprintf(g, "// %s\n", f16name[opcode])
	if opcode == 14 {
		g.WriteOpInvalid("invalid F16 with opcode==14")
		g.WriteDisasmInvalid()
		return
	}
	if opcode == 15 {
		fmt.Fprintf(g, "cpu.Exception(ExceptionSwi)\n")
		g.WriteDisasm("swi", "x:op&0xFF")
		return
	}

	fmt.Fprintf(g, "if %s {\n", f16cond[opcode])
	fmt.Fprintf(g, "offset := int8(uint8(op&0xFF))\n")
	fmt.Fprintf(g, "offset32 := int32(offset)*2\n")
	g.writeBranch("cpu.Regs[15]+reg(offset32)", "BranchJump")
	fmt.Fprintf(g, "}\n")
	g.WriteDisasm(f16name[opcode], "o:int32(int8(uint8(op&0xFF)))*2")
}

func (g *Generator) writeOpF18Branch(op uint16) {
	fmt.Fprintf(g, "// b\n")
	fmt.Fprintf(g, "pc := cpu.Regs[15] + reg(int32(int16(op<<5)>>4))\n")
	g.writeBranch("pc", "BranchJump")
	g.WriteDisasm("b", "o:int32(int16(op<<5)>>4)")
}

func (g *Generator) writeOpF19LongBranch1(op uint16) {
	fmt.Fprintf(g, "// bl/blx step 1\n")
	fmt.Fprintf(g, "cpu.Regs[14] = cpu.Regs[15] + reg(int32(uint32(op&0x7FF)<<21)>>9)\n")

	fmt.Fprintf(&g.Disasm, "mem := cpu.opFetchPointer(pc+2)\n")
	fmt.Fprintf(&g.Disasm, "op2 := uint16(mem[0]) | uint16(mem[1])<<8\n")
	fmt.Fprintf(&g.Disasm, "nextpc := (int32(uint32(op&0x7FF)<<21)>>9) + int32((op2&0x7FF)<<1)\n")
	fmt.Fprintf(&g.Disasm, "if (op2>>12)&1 == 0{\n")
	g.WriteDisasm("blx", "o:nextpc")
	fmt.Fprintf(&g.Disasm, "} else {\n")
	g.WriteDisasm("bl", "o:nextpc")
	fmt.Fprintf(&g.Disasm, "}\n")
}

func (g *Generator) writeOpF19LongBranch2(op uint16) {
	blx := (op>>12)&1 == 0
	if blx {
		fmt.Fprintf(g, "// blx step 2\n")
	} else {
		fmt.Fprintf(g, "// bl step 2\n")
	}
	fmt.Fprintf(g, "newpc := cpu.Regs[14] + reg((op&0x7FF)<<1)\n")
	fmt.Fprintf(g, "cpu.Regs[14] = (cpu.Regs[15]-2) | 1\n")
	if blx {
		fmt.Fprintf(g, "newpc &^= 2\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetT(false)\n")
	}
	g.writeBranch("newpc", "BranchCall")
	fmt.Fprintf(&g.Disasm, "return \"[continued]\"\n")
}

var opaluname = [16]string{
	"ands", "eors", "lsls", "lsrs", "asrs", "adcs", "sbcs", "rors",
	"tst", "negs", "cmp", "cmn", "orrs", "muls", "bics", "mvn",
}

func (g *Generator) WriteAluOp(op uint16) {
	if op>>10 != 0x10 {
		panic("invalid ALU opcode")
	}
	opcode := (op >> 6) & 0xF

	g.writeOpAluHeader(op)
	fmt.Fprintf(g, "// %s\n", opaluname[opcode])
	fmt.Fprintf(g, "rsx := (op>>3)&0x7\n")
	fmt.Fprintf(g, "rs := uint32(cpu.Regs[rsx])\n")
	fmt.Fprintf(g, "rdx := op&0x7\n")
	if opcode != 9 && opcode != 0xF {
		fmt.Fprintf(g, "rd := uint32(cpu.Regs[rdx])\n")
	}
	g.WriteDisasm(opaluname[opcode], "r:op&7", "r:(op>>3)&7")

	test := false
	switch opcode {
	case 8: // TST
		test = true
		fallthrough
	case 0: // AND
		fmt.Fprintf(g, "res := rd & rs\n")
	case 1: // EOR
		fmt.Fprintf(g, "res := rd ^ rs\n")
	case 2: // LSL
		fmt.Fprintf(g, "rot := (rs&0xFF)\n")
		fmt.Fprintf(g, "if rot != 0 { cpu.Cpsr.SetC(rd & (1<<(32-rot)) != 0) }\n")
		fmt.Fprintf(g, "res := rd << rot\n")
	case 3: // LSR
		fmt.Fprintf(g, "rot := (rs&0xFF)\n")
		fmt.Fprintf(g, "if rot != 0 { cpu.Cpsr.SetC(rd & (1<<(rot-1)) != 0) }\n")
		fmt.Fprintf(g, "res := rd >> rot\n")
	case 4: // ASR
		fmt.Fprintf(g, "rot := (rs&0xFF)\n")
		fmt.Fprintf(g, "if rot != 0 { cpu.Cpsr.SetC(rd & (1<<(rot-1)) != 0) }\n")
		fmt.Fprintf(g, "res := uint32(int32(rd) >> rot)\n")
	case 5: // ADC
		fmt.Fprintf(g, "cf := cpu.Cpsr.CB()\n")
		fmt.Fprintf(g, "res := rd + rs\n")
		fmt.Fprintf(g, "res += cf\n")
		fmt.Fprintf(g, "if cf == 0 {\n")
		fmt.Fprintf(g, "  cpu.Cpsr.SetC(res<rd)\n")
		fmt.Fprintf(g, "} else {\n")
		fmt.Fprintf(g, "  cpu.Cpsr.SetC(res<=rd)\n")
		fmt.Fprintf(g, "}\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVAdd(rd, rs, res)\n")
	case 6: // SBC
		fmt.Fprintf(g, "cf := cpu.Cpsr.CB()\n")
		fmt.Fprintf(g, "res := rd - rs\n")
		fmt.Fprintf(g, "res += cf-1\n")
		fmt.Fprintf(g, "if cf == 0 {\n")
		fmt.Fprintf(g, "  cpu.Cpsr.SetC(rd>rs)\n")
		fmt.Fprintf(g, "} else {\n")
		fmt.Fprintf(g, "  cpu.Cpsr.SetC(rd>=rs)\n")
		fmt.Fprintf(g, "}\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVSub(rd, rs, res)\n")
	case 7: // ROR
		fmt.Fprintf(g, "rot := (rs&0xFF)\n")
		fmt.Fprintf(g, "if rot != 0 { cpu.Cpsr.SetC(rd & (1<<(rot-1)) != 0) }\n")
		fmt.Fprintf(g, "rot = (rs&0x1F)\n")
		fmt.Fprintf(g, "res := (rd >> rot) | (rd << (32-rot))\n")
	case 9: // NEG
		fmt.Fprintf(g, "res := 0 - rs\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(false)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVSub(0, rs, res)\n")
	case 10: // CMP
		test = true
		fmt.Fprintf(g, "res := rd - rs\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(rd>=rs)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVSub(rd, rs, res)\n")
	case 11: // CMN
		test = true
		fmt.Fprintf(g, "res := rd + rs\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetC(res<rd)\n")
		fmt.Fprintf(g, "cpu.Cpsr.SetVAdd(rd, rs, res)\n")
	case 12: // ORR
		fmt.Fprintf(g, "res := rd | rs\n")
	case 13: // MUL
		fmt.Fprintf(g, "res := rd * rs\n")
		fmt.Fprintf(g, "if cpu.arch <= ARMv4 { cpu.Cpsr.SetC(false) }\n")
	case 14: // BIC
		fmt.Fprintf(g, "res := rd &^ rs\n")
	case 15: // MVN
		fmt.Fprintf(g, "res := ^rs\n")
	default:
		panic("unreachable")
	}

	fmt.Fprintf(g, "cpu.Cpsr.SetNZ(res)\n")
	if !test {
		fmt.Fprintf(g, "cpu.Regs[rdx] = reg(res)\n")
	}
	g.writeOpAluFooter(op)
}

func (g *Generator) WriteOp(op uint16) {
	g.WriteOpHeader(int((op >> 8) & 0xFF))

	oph := op >> 8
	switch {
	case oph>>5 == 0x0 && (oph>>3)&3 != 3: // F1
		g.writeOpF1Shift(op)

	case oph>>5 == 0x0 && (oph>>3)&3 == 3: // F2
		g.writeOpF2Add(op)

	case oph>>5 == 0x1: // F3
		g.writeOpF3AluImm(op)

	case oph>>2 == 0x10: // F4
		g.writeOpF4Alu(op)

	case oph>>2 == 0x11: // F5
		g.writeOpF5HiReg(op)

	case oph>>3 == 0x9: // F6
		g.writeOpF6LdrPc(op)

	case oph>>4 == 0x5: // F7 & F8
		g.writeOpF7F8LdrStr(op)

	case oph>>5 == 0x3: // F9
		g.writeOpF9Strb(op)

	case oph>>4 == 0x8: // F10
		g.writeOpF10Strh(op)

	case oph>>4 == 0x9: // F11
		g.writeOpF11Strsp(op)

	case oph>>4 == 0xA: // F12
		g.writeOpF12AddPc(op)

	case oph>>4 == 0xB && oph&0xF == 0: // F13
		g.writeOpF13AddSp(op)

	case oph>>4 == 0xB && oph&6 == 4: // F14
		g.writeOpF14PushPop(op)

	case oph>>4 == 0xC: // F15
		g.writeOpF15LdmStm(op)

	case oph>>4 == 0xD: // F16
		g.writeOpF16BranchCond(op)

	case oph>>3 == 0x1C: // F18
		g.writeOpF18Branch(op)

	case oph>>3 == 0x1E: // F19
		g.writeOpF19LongBranch1(op)
	case oph>>3 == 0x1F || oph>>3 == 0x1D: // F19
		g.writeOpF19LongBranch2(op)

	default:
		g.WriteOpInvalid("not implemented")
		g.WriteDisasmInvalid()
	}

	g.WriteOpFooter(int((op >> 8) & 0xFF))
}

func main() {
	cpugen.Main(func(g *cpugen.Generator) {
		out := Generator{g}
		out.Prefix = "Thumb"
		out.OpSize = "uint16"
		out.GenDisasm = true
		out.PcRelOff = 4
		out.TableBits = 8
		out.WriteHeader()
		for op := 0; op < 0x100; op++ {
			out.WriteOp(uint16(op << 8))
		}
		for op := 0; op < 0x10; op++ {
			out.WriteAluOp(uint16(op<<6) | uint16(0x10<<10))
		}

		out.WriteFooter()
	})
}
