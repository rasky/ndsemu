package main

import "ndsemu/arm"

type HwIrq struct {
	Cpu    *arm.Cpu
	master uint16
	enable uint32
	flags  uint32
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

	IrqIpcSync     IrqType = (1 << 16)
	IrqIpcSendFifo IrqType = (1 << 17)
	IrqIpcRecvFifo IrqType = (1 << 18)

	IrqGameCardData  IrqType = (1 << 19)
	IrqGameCardEject IrqType = (1 << 20)

	IrqTimers IrqType = (IrqTimer0 | IrqTimer1 | IrqTimer2 | IrqTimer3)
)

func (irq *HwIrq) ReadIME() uint16 {
	return irq.master
}

func (irq *HwIrq) ReadIE() uint32 {
	return irq.enable
}

func (irq *HwIrq) ReadIF() uint32 {
	return irq.flags
}

func (irq *HwIrq) updateLineStatus() {
	irqstat := irq.master != 0 && (irq.enable&irq.flags) != 0
	if irqstat {
		if (irq.enable&irq.flags)&^uint32(IrqTimers) != 0 {
			Emu.Log().Infof("[irq] trigger %08x", irq.flags&irq.enable)
		}
	}
	irq.Cpu.SetLine(arm.LineIrq, irqstat)
}

func (irq *HwIrq) WriteIME(ime uint16) {
	irq.master = ime & 1
	irq.updateLineStatus()
}

func (irq *HwIrq) WriteIE(ie uint32) {
	irq.enable = ie
	if ie&^uint32(IrqVBlank|IrqTimers|IrqIpcRecvFifo) != 0 {
		Emu.Log().Infof("[irq] IE: %08x", ie&^uint32(IrqVBlank|IrqTimers|IrqIpcRecvFifo))
	}
	irq.updateLineStatus()
}

func (irq *HwIrq) WriteIF(ifx uint32) {
	irq.flags &^= ifx
	if ifx&^uint32(IrqTimers) != 0 {
		Emu.Log().Infof("[irq] IF: %08x", ifx)
	}
	irq.updateLineStatus()
}

func (irq *HwIrq) Raise(irqtype IrqType) {
	irq.flags |= uint32(irqtype)
	irq.updateLineStatus()
}
