package arm

import (
	"fmt"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type reg uint32

func boolToReg(f bool) reg {
	// Use a form that the compiler can optimize into a SETxx (https://github.com/golang/go/issues/6011)
	var i reg
	if f {
		i = 1
	}
	return i
}

func (r reg) Bit(n uint) bool { return ((uint32(r) >> n) & 1) != 0 }

func (r *reg) BitSet(n uint)   { *r |= 1 << n }
func (r *reg) BitClear(n uint) { *r &= ^(1 << n) }
func (r *reg) BitChange(n uint, f bool) {
	i := boolToReg(f)
	*r = ((*r) &^ (1 << n)) | i<<n
}

func (r *reg) SetWithMask(val uint32, mask uint32) {
	*r = reg(((uint32)(*r) &^ mask) | (val & mask))
}

func (r reg) String() string {
	return fmt.Sprintf("%08X", uint32(r))
}

type regCpsr struct {
	N, Z, C, V, Q, _i, _f, _t bool
	_mode                     uint8
}

func (r regCpsr) CB() uint32 { return uint32(boolToReg(r.C)) }
func (r regCpsr) I() bool    { return r._i }
func (r regCpsr) F() bool    { return r._f }
func (r regCpsr) T() bool    { return r._t }

func (r *regCpsr) SetC(val bool) {
	r.C = val
}

func (r *regCpsr) SetNZ(val uint32) {
	r.N = int32(val) < 0
	r.Z = val == 0
}

func (r *regCpsr) SetNZ64(val uint64) {
	r.N = int64(val) < 0
	r.Z = val == 0
}

func (r *regCpsr) SetVAdd(s1, s2, res uint32) {
	r.V = ^(s1^s2)&(s1^res)&0x80000000 != 0
}

func (r *regCpsr) SetVSub(s1, s2, res uint32) {
	r.V = (s1^s2)&(s1^res)&0x80000000 != 0
}

func (r *regCpsr) GetMode() CpuMode {
	return CpuMode(r._mode)
}

func (r *regCpsr) Uint32() uint32 {
	val := uint32(r._mode) & 0x1F
	if r.N {
		val |= 1 << 31
	}
	if r.Z {
		val |= 1 << 30
	}
	if r.C {
		val |= 1 << 29
	}
	if r.V {
		val |= 1 << 28
	}
	if r.Q {
		val |= 1 << 27
	}
	if r._i {
		val |= 1 << 7
	}
	if r._f {
		val |= 1 << 6
	}
	if r._t {
		val |= 1 << 5
	}
	return val
}

func (r *regCpsr) Set(val uint32, cpu *Cpu) {
	r.SetMode(CpuMode(val&0x1F), cpu)

	r._t = val&(1<<5) != 0
	r._f = val&(1<<6) != 0
	r._i = val&(1<<7) != 0
	r.Q = val&(1<<27) != 0
	r.V = val&(1<<28) != 0
	r.C = val&(1<<29) != 0
	r.Z = val&(1<<30) != 0
	r.N = val&(1<<31) != 0

	// The T/I/F bits are potentially changed, we must force
	// exit the tight loop, to check if the new bits will cause
	// an interrupt right away.
	cpu.tightExit = true
}

func (r *regCpsr) SetWithMask(val uint32, mask uint32, cpu *Cpu) {
	old := r.Uint32()
	val = (val & mask) | (old &^ mask)
	r.Set(val, cpu)
}

func (r *regCpsr) SetT(val bool, cpu *Cpu) {
	r._t = val
	cpu.tightExit = true
}

func (r *regCpsr) SetI(val bool, cpu *Cpu) {
	r._i = val
	cpu.tightExit = true
}

func (r *regCpsr) SetF(val bool, cpu *Cpu) {
	r._f = val
	cpu.tightExit = true
}

func (r *regCpsr) SetMode(mode CpuMode, cpu *Cpu) {
	oldmode := CpuMode(r._mode)
	r._mode = uint8(mode) & 0x1F

	if oldmode == mode {
		return
	}

	switch oldmode {
	case CpuModeUser, CpuModeSystem:
		copy(cpu.UsrBank[:], cpu.Regs[13:15])
	case CpuModeFiq:
		copy(cpu.FiqBank2[:], cpu.Regs[8:13])
		copy(cpu.Regs[8:13], cpu.UsrBank2[:])
		copy(cpu.FiqBank[:], cpu.Regs[13:15])
	case CpuModeIrq:
		copy(cpu.IrqBank[:], cpu.Regs[13:15])
	case CpuModeSupervisor:
		copy(cpu.SvcBank[:], cpu.Regs[13:15])
	case CpuModeAbort:
		copy(cpu.AbtBank[:], cpu.Regs[13:15])
	case CpuModeUndefined:
		copy(cpu.UndBank[:], cpu.Regs[13:15])
	default:
		log.Fatalf("unknown CPU oldmode: %v", oldmode)
	}

	switch mode {
	case CpuModeUser, CpuModeSystem:
		copy(cpu.Regs[13:15], cpu.UsrBank[:])
	case CpuModeFiq:
		copy(cpu.UsrBank2[:], cpu.Regs[8:13])
		copy(cpu.Regs[8:13], cpu.FiqBank2[:])
		copy(cpu.Regs[13:15], cpu.FiqBank[:])
	case CpuModeIrq:
		copy(cpu.Regs[13:15], cpu.IrqBank[:])
	case CpuModeSupervisor:
		copy(cpu.Regs[13:15], cpu.SvcBank[:])
	case CpuModeAbort:
		copy(cpu.Regs[13:15], cpu.AbtBank[:])
	case CpuModeUndefined:
		copy(cpu.Regs[13:15], cpu.UndBank[:])
	default:
		log.Fatalf("unknown CPU newmode: %v", mode)
	}
}

type CpuMode uint8

const (
	CpuModeUser       CpuMode = 0x10
	CpuModeFiq        CpuMode = 0x11
	CpuModeIrq        CpuMode = 0x12
	CpuModeSupervisor CpuMode = 0x13
	CpuModeAbort      CpuMode = 0x17
	CpuModeUndefined  CpuMode = 0x18
	CpuModeSystem     CpuMode = 0x1F
)

func (m CpuMode) String() string {
	switch m {
	case CpuModeUser:
		return "CpuModeUser"
	case CpuModeFiq:
		return "CpuModeFiq"
	case CpuModeIrq:
		return "CpuModeIrq"
	case CpuModeSupervisor:
		return "CpuModeSupervisor"
	case CpuModeAbort:
		return "CpuModeAbort"
	case CpuModeUndefined:
		return "CpuModeUndefined"
	case CpuModeSystem:
		return "CpuModeSystem"
	default:
		return fmt.Sprintf("CpuMode(%02x)", int(m))
	}
}
