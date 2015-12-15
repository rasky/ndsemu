package arm

import (
	"fmt"

	log "gopkg.in/Sirupsen/logrus.v0"
)

//go:generate go run genarm/genarm.go -filename ops_arm_table.go
//go:generate go run genthumb/genthumb.go -filename ops_thumb_table.go

var popcount = [256]uint8{
	0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
	4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8,
}

func popcount8(val uint8) uint {
	return uint(popcount[val])
}
func popcount16(val uint16) uint {
	return popcount8(uint8(val&0xFF)) + popcount8(uint8(val>>8))
}

// func init() {
// 	for i := 0; i < 256; i++ {
// 		cnt := 0
// 		for j := 0; j < 8; j++ {
// 			if i&(1<<uint(j)) != 0 {
// 				cnt++
// 			}
// 		}
// 		popcount[i] = cnt
// 	}
// 	fmt.Printf("%#v\n", popcount)
// }

// Called by opXX generated functions that use a register as second operand;
// the decoding is complex with many subcases, so we write the code here and
// just call it
func (cpu *Cpu) opDecodeAluOp2Reg(op uint32, setcarry bool) uint32 {
	var shift uint32
	if op&0x10 != 0 {
		if (op>>7)&0x1 != 0 {
			cpu.InvalidOpArm(op, "opDecodeAluOp2Reg: bit 7 is not zero")
		}
		// Increment PC if shifting by register (in case it's accessed
		// by this instruction as an operand)
		cpu.Regs[15] += 4
		shift = uint32(cpu.Regs[(op>>8)&0xF] & 0xFF)
		if shift >= 32 {
			cpu.InvalidOpArm(op, "opDecodeAluOp2Reg: big shift factor not implemented")
		}
	} else {
		shift = uint32((op >> 7) & 0x1F)
	}

	rm := uint32(cpu.Regs[op&0xF])
	shtype := (op >> 5) & 3
	if shift == 0 && shtype != 0 {
		fmt.Println(shtype)
		cpu.InvalidOpArm(op, "opDecodeAluOp2Reg: zero shift factor not implemented")
	}

	if shtype == 0 {
		if shift == 0 {
			return rm
		}
		if setcarry {
			cpu.Cpsr.SetC((rm>>(32-shift))&1 != 0)
		}
		res := rm << shift
		return res
	} else if shtype == 1 {
		if shift == 0 {
			shift = 32
		}
		if setcarry {
			cpu.Cpsr.SetC((rm>>(shift-1))&1 != 0)
		}
		res := rm >> shift
		return res
	} else if shtype == 2 {
		if shift == 0 {
			shift = 32
		}
		if setcarry {
			cpu.Cpsr.SetC((rm>>(shift-1))&1 != 0)
		}
		res := uint32(int32(rm) >> shift)
		return res
	} else if shtype == 3 {
		if shift == 0 {
			cb := cpu.Cpsr.CB()
			cpu.Cpsr.SetC(rm&1 != 0)
			return (rm >> 1) | (cb << 31)
		}
		shift &= 31
		res := (rm >> shift) | (rm << (32 - shift))
		if setcarry {
			cpu.Cpsr.SetC(rm>>31 != 0)
		}
		return res
	}
	panic("unreachable")
}

func (cpu *Cpu) InvalidOpArm(op uint32, msg string) {
	panic(fmt.Sprintf("FATAL: invalid ARM opcode (%08X): %s", op, msg))
}

func (cpu *Cpu) InvalidOpThumb(op uint16, msg string) {
	log.WithFields(log.Fields{
		"pc": cpu.pc - 2,
		"op": fmt.Sprintf("%04x", op),
	}).Fatalf("invalid thumb opcode: %s", msg)
}

func (cpu *Cpu) opArmCond(op uint32) bool {
	switch op >> 28 {
	case 0:
		return cpu.Cpsr.Z()
	case 1:
		return !cpu.Cpsr.Z()
	case 2:
		return cpu.Cpsr.C()
	case 3:
		return !cpu.Cpsr.C()
	case 4:
		return cpu.Cpsr.N()
	case 5:
		return !cpu.Cpsr.N()
	case 6:
		return cpu.Cpsr.V()
	case 7:
		return !cpu.Cpsr.V()
	case 8:
		return cpu.Cpsr.C() && !cpu.Cpsr.Z()
	case 9:
		return !cpu.Cpsr.C() || cpu.Cpsr.Z()
	case 10:
		return cpu.Cpsr.N() == cpu.Cpsr.V()
	case 11:
		return cpu.Cpsr.N() != cpu.Cpsr.V()
	case 12:
		return !cpu.Cpsr.Z() && cpu.Cpsr.N() == cpu.Cpsr.V()
	case 13:
		return cpu.Cpsr.Z() || cpu.Cpsr.N() != cpu.Cpsr.V()
	case 14:
		return true
	case 15:
		return false
	}
	return false
}

func (cpu *Cpu) opCopRead(copnum uint32, op uint32, cn, cm, cp uint32) uint32 {
	cop := cpu.cops[copnum]
	if cop == nil {
		log.WithFields(log.Fields{
			"pc":  cpu.pc,
			"cop": copnum,
		}).Error("invalid coprocessor access")
		return 0xFFFFFFFF
	}

	return cop.Read(op, cn, cm, cp)
}

func (cpu *Cpu) opCopWrite(copnum uint32, op uint32, cn, cm, cp uint32, value uint32) {
	cop := cpu.cops[copnum]
	if cop == nil {
		log.WithFields(log.Fields{
			"pc":  cpu.pc,
			"cop": copnum,
		}).Error("invalid coprocessor access")
		return
	}

	cop.Write(op, cn, cm, cp, value)
}

func (cpu *Cpu) opCopExec(copnum uint32, op uint32, cn, cm, cp, cd uint32) {
	cop := cpu.cops[copnum]
	if cop == nil {
		log.WithFields(log.Fields{
			"pc":  cpu.pc,
			"cop": copnum,
		}).Error("invalid coprocessor access")
		return
	}

	cop.Exec(op, cn, cm, cp, cd)
}

func (cpu *Cpu) Run(until int64) {
	cpu.pc = cpu.Regs[15]
	for cpu.Clock < until {
		lines := cpu.lines
		if lines&LineFiq != 0 && cpu.Cpsr.I() {
			cpu.Exception(ExceptionFiq)
		}
		if lines&LineIrq != 0 && cpu.Cpsr.F() {
			cpu.Exception(ExceptionIrq)
		}
		if lines&LineHalt != 0 {
			cpu.Clock = until
			return
		}
		thumb := cpu.Cpsr.T()
		if !thumb {
			log.WithFields(log.Fields{
				"mode": "arm",
				"pc":   cpu.pc,
			}).Info("[CPU] step")
			op := cpu.bus.Read32(uint32(cpu.pc))
			cpu.Regs[15] = cpu.pc + 8 // simulate pipeline with prefetch
			cpu.pc += 4
			opArmTable[(op>>20)&0xFF](cpu, op)
		} else {
			log.WithFields(log.Fields{
				"mode": "thumb",
				"pc":   cpu.pc,
			}).Info("[CPU] step")
			op := cpu.bus.Read16(uint32(cpu.pc))
			cpu.Regs[15] = cpu.pc + 4 // simulate pipeline with prefetch
			cpu.pc += 2
			opThumbTable[(op>>8)&0xFF](cpu, op)
		}
		cpu.Clock += 1
	}
	cpu.Regs[15] = cpu.pc
}

func (cpu *Cpu) GetPC() reg {
	thumb := cpu.Cpsr.T()
	if !thumb {
		return cpu.pc - 4
	} else {
		return cpu.pc - 2
	}
}
