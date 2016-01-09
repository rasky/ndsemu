package main

import (
	"ndsemu/arm"
	"ndsemu/emu/hwio"

	"gopkg.in/Sirupsen/logrus.v0"
)

type HwIrq struct {
	Name string
	Cpu  *arm.Cpu

	Ime hwio.Reg32 `hwio:"offset=0x08,rwmask=0x1,wcb"`
	Ie  hwio.Reg32 `hwio:"offset=0x10,wcb"`
	If  hwio.Reg32 `hwio:"offset=0x14,wcb"`
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

	IrqTimers IrqType = (IrqTimer0 | IrqTimer1 | IrqTimer2 | IrqTimer3)
)

func NewHwIrq(name string, cpu *arm.Cpu) *HwIrq {
	irq := &HwIrq{Name: name, Cpu: cpu}
	hwio.MustInitRegs(irq)
	return irq
}

func (irq *HwIrq) Log() *logrus.Entry {
	return Emu.Log().WithField("name", irq.Name)
}

func (irq *HwIrq) WriteIME(_, _ uint32) {
	// irq.Log().Info("[irq] ", irq.Ime)
	irq.updateLineStatus()
}

func (irq *HwIrq) updateLineStatus() {
	// if irq.Cpu == nds9.Cpu {
	// 	irq.Log().Info("[irq] ", irq.Ime, irq.Ie, irq.If)
	// }
	irqstat := irq.Ime.Value != 0 && (irq.Ie.Value&irq.If.Value) != 0
	if irqstat {
		if (irq.Ie.Value&irq.If.Value)&^uint32(IrqTimers|IrqVBlank) != 0 {
			irq.Log().Infof("[irq] trigger %08x", irq.If.Value&irq.Ie.Value)
		}
	}
	irq.Cpu.SetLine(arm.LineIrq, irqstat)
}

func (irq *HwIrq) WriteIE(_, ie uint32) {
	if ie&^uint32(IrqVBlank|IrqTimers|IrqIpcRecvFifo) != 0 {
		irq.Log().Infof("[irq] IE: %08x", ie&^uint32(IrqVBlank|IrqTimers|IrqIpcRecvFifo))
	}
	irq.updateLineStatus()
}

func (irq *HwIrq) WriteIF(old, ifx uint32) {
	irq.If.Value = old &^ ifx
	if ifx&^uint32(IrqTimers) != 0 {
		irq.Log().Infof("[irq] Irq ACK: %08x", ifx)
	}
	irq.updateLineStatus()
}

func (irq *HwIrq) Raise(irqtype IrqType) {
	irq.If.Value |= uint32(irqtype)
	// irq.Log().Info("[irq] raise", irq.If)
	irq.updateLineStatus()
}
