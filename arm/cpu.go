package arm

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

type Arch int

const (
	ARMv4 Arch = 4
	ARMv5 Arch = 5
)

type Cpu struct {
	Regs  [16]reg
	Cpsr  regCpsr
	Clock int64

	UsrBank  [2]reg
	FiqBank  [2]reg
	SvcBank  [2]reg
	AbtBank  [2]reg
	IrqBank  [2]reg
	UndBank  [2]reg
	SpsrBank [5]reg

	UsrBank2 [5]reg
	FiqBank2 [5]reg

	arch Arch
	bus  Bus
	pc   reg
	cp15 *Cp15
	cops [16]Coprocessor
}

func NewCpu(arch Arch, bus Bus) *Cpu {
	cpu := &Cpu{bus: bus, arch: arch}
	cpu.Cpsr.r = 0x13 // mode supervisor
	return cpu
}

func (cpu *Cpu) SetPC(addr uint32) {
	cpu.Regs[15] = reg(addr)
}

func (cpu *Cpu) RegSpsrForMode(mode CpuMode) *reg {
	switch mode {
	case CpuModeUser, CpuModeSystem:
		return &cpu.SpsrBank[0]
	case CpuModeFiq:
		return &cpu.SpsrBank[1]
	case CpuModeSupervisor:
		return &cpu.SpsrBank[2]
	case CpuModeAbort:
		return &cpu.SpsrBank[2]
	case CpuModeIrq:
		return &cpu.SpsrBank[3]
	case CpuModeUndefined:
		return &cpu.SpsrBank[4]
	default:
		log.Fatalf("unsupported mode in RegSpsr(): %v", mode)
		panic("unreachable")
	}
}

func (cpu *Cpu) RegF14ForMode(mode CpuMode) *reg {
	switch mode {
	case CpuModeUser, CpuModeSystem:
		return &cpu.UsrBank[1]
	case CpuModeFiq:
		return &cpu.FiqBank[1]
	case CpuModeSupervisor:
		return &cpu.SvcBank[1]
	case CpuModeAbort:
		return &cpu.AbtBank[1]
	case CpuModeIrq:
		return &cpu.IrqBank[1]
	case CpuModeUndefined:
		return &cpu.UndBank[1]
	default:
		log.Fatalf("unsupported mode in RegSpsr(): %v", mode)
		panic("unreachable")
	}
}

func (cpu *Cpu) RegSpsr() *reg {
	return cpu.RegSpsrForMode(cpu.Cpsr.GetMode())
}

func (cpu *Cpu) MapCoprocessor(copnum int, cop Coprocessor) {
	cpu.cops[copnum] = cop
}

func (cpu *Cpu) EnableCp15() *Cp15 {
	cpu.cp15 = newCp15(cpu)
	cpu.cops[15] = cpu.cp15
	return cpu.cp15
}

type Exception int

const (
	ExceptionReset           Exception = 0
	ExceptionUndefined       Exception = 1
	ExceptionSwi             Exception = 2
	ExceptionPrefetchAbort   Exception = 3
	ExceptionDataAbort       Exception = 4
	ExceptionAddressOverflow Exception = 5
	ExceptionIrq             Exception = 6
	ExceptionFiq             Exception = 7
)

// CPU mode to enter when the exception is raised
var excMode = [8]CpuMode{
	CpuModeSupervisor,
	CpuModeUndefined,
	CpuModeSupervisor,
	CpuModeAbort,
	CpuModeAbort,
	CpuModeSupervisor,
	CpuModeIrq,
	CpuModeFiq,
}

func (cpu *Cpu) Exception(exc Exception) {
	newmode := excMode[exc]

	// Check if FIQ/IRQ are disabled
	if exc == ExceptionFiq && cpu.Cpsr.F() {
		return
	}
	if exc == ExceptionIrq && cpu.Cpsr.I() {
		return
	}

	*cpu.RegSpsrForMode(newmode) = cpu.Cpsr.r
	*cpu.RegF14ForMode(newmode) = cpu.pc
	cpu.Cpsr.SetT(false)
	cpu.Cpsr.SetWithMask(uint32(newmode), 0x1F, cpu)
	cpu.Cpsr.SetI(true)
	if exc == ExceptionReset || exc == ExceptionFiq {
		cpu.Cpsr.SetF(true)
	}

	if cpu.cp15 != nil {
		cpu.Regs[15] = reg(cpu.cp15.ExceptionVector())
	} else {
		cpu.Regs[15] = 0x00000000
	}

	cpu.Regs[15] += reg(exc * 4)
	cpu.pc = cpu.Regs[15]
}

func (cpu *Cpu) Reset() {
	cpu.Exception(ExceptionReset)
}
