package main

import (
	"ndsemu/arm"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

type HwIrq struct {
	Name string
	Cpu  *arm.Cpu

	Ime hwio.Reg32 `hwio:"offset=0x08,rwmask=0x1,wcb"`
	Ie  hwio.Reg32 `hwio:"offset=0x10,wcb"`
	If  hwio.Reg32 `hwio:"offset=0x14,wcb"`

	// Mask of level-triggerd IRQs (can't be asserted by CPU)
	lvlirq uint32
}

type IrqType uint32

const (
	IrqVBlank IrqType = (1 << 0)
	IrqHBlank IrqType = (1 << 1)
	IrqVMatch IrqType = (1 << 2)
	IrqTimer0 IrqType = (1 << 3)
	IrqTimer1 IrqType = (1 << 4)
	IrqTimer2 IrqType = (1 << 5)
	IrqTimer3 IrqType = (1 << 6)
	IrqRtc    IrqType = (1 << 7) // nds7 only

	IrqDma0 IrqType = (1 << 8)
	IrqDma1 IrqType = (1 << 9)
	IrqDma2 IrqType = (1 << 10)
	IrqDma3 IrqType = (1 << 11)

	IrqIpcSync     IrqType = (1 << 16)
	IrqIpcSendFifo IrqType = (1 << 17)
	IrqIpcRecvFifo IrqType = (1 << 18)

	IrqGameCardData  IrqType = (1 << 19)
	IrqGameCardEject IrqType = (1 << 20)

	IrqGxFifo IrqType = (1 << 21)

	IrqTimers IrqType = (IrqTimer0 | IrqTimer1 | IrqTimer2 | IrqTimer3)
)

func NewHwIrq(name string, cpu *arm.Cpu) *HwIrq {
	irq := &HwIrq{Name: name, Cpu: cpu}
	hwio.MustInitRegs(irq)
	return irq
}

func (irq *HwIrq) Log() log.Entry {
	return log.ModIrq.WithField("name", irq.Name)
}

func (irq *HwIrq) WriteIME(_, _ uint32) {
	// irq.Log().Info("", irq.Ime)
	irq.updateLineStatus()
}

func (irq *HwIrq) updateLineStatus() {
	// if irq.Cpu == nds9.Cpu {
	// 	irq.Log().Info("", irq.Ime, irq.Ie, irq.If)
	// }
	irqstat := irq.Ime.Value != 0 && (irq.Ie.Value&irq.If.Value) != 0
	if irqstat {
		if (irq.Ie.Value&irq.If.Value)&^uint32(IrqTimers|IrqVBlank) != 0 {
			irq.Log().Infof("trigger %08x", irq.If.Value&irq.Ie.Value)
		}
	}
	irq.Cpu.SetLine(arm.LineIrq, irqstat)
}

func (irq *HwIrq) WriteIE(_, ie uint32) {
	if ie&^uint32(IrqVBlank|IrqTimers|IrqIpcRecvFifo) != 0 {
		irq.Log().Infof("IE: %08x", ie&^uint32(IrqVBlank|IrqTimers|IrqIpcRecvFifo))
	}
	irq.updateLineStatus()
}

func (irq *HwIrq) WriteIF(old, ifx uint32) {
	// ignore acknowledge of level-triggered interrupts
	ifx &^= irq.lvlirq
	// acknowledge the irqs in the write mask
	irq.If.Value = old &^ ifx
	if ifx&^uint32(IrqTimers) != 0 {
		irq.Log().Infof("Irq ACK: %08x", ifx)
	}
	irq.updateLineStatus()
}

// Raise an edge-triggered interrupt. The interrupt is shown in the IF
// register, and can be acknowledged by the CPU at any time by writing
// to the same reg.
func (irq *HwIrq) Raise(irqtype IrqType) {
	irq.If.Value |= uint32(irqtype)
	// irq.Log().Info("raise", irq.If)
	irq.updateLineStatus()
}

// Assert a level-triggered interrupt. The interrupt is shown in the IF
// register, but the CPU can't acknowledge it. It will be acknowledged
// only by a subsequent call to Assert().
func (irq *HwIrq) Assert(irqtype IrqType, set bool) {
	if set {
		irq.lvlirq |= uint32(irqtype)
		irq.If.Value |= uint32(irqtype)
	} else {
		irq.lvlirq &^= uint32(irqtype)
	}
	irq.updateLineStatus()
}
