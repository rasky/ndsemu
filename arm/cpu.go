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

	UsrBank  [7]reg
	FiqBank  [7]reg
	SvcBank  [2]reg
	AbtBank  [2]reg
	IrqBank  [2]reg
	UndBank  [2]reg
	SpsrBank [5]reg

	arch Arch
	bus  Bus
	pc   reg
	cp15 *Cp15
	cops [16]Coprocessor
}

func NewCpu(arch Arch, bus Bus) *Cpu {
	cpu := &Cpu{bus: bus, arch: arch}
	return cpu
}

func (cpu *Cpu) RegSpsr() *reg {
	switch mode := cpu.Cpsr.GetMode(); mode {
	case CpuModeUser:
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
		log.Fatalf("unsupport mode in RegSpsr(): %v", mode)
		panic("unreachable")
	}
}

func (cpu *Cpu) MapCoprocessor(copnum int, cop Coprocessor) {
	cpu.cops[copnum] = cop
}

func (cpu *Cpu) EnableCp15() *Cp15 {
	cpu.cp15 = newCp15(cpu)
	cpu.cops[15] = cpu.cp15
	return cpu.cp15
}
