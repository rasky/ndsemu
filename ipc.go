package main

import (
	"ndsemu/emu/hwio"

	"fmt"
)

type ipcFifo struct {
	fifocnt  uint16
	fifo     []uint32
	emptyIrq bool
	dataIrq  bool
	last     uint32
}

func (ipc *ipcFifo) Empty() bool { return len(ipc.fifo) == 0 }
func (ipc *ipcFifo) Full() bool  { return len(ipc.fifo) == 16 }
func (ipc *ipcFifo) Flush()      { ipc.fifo = nil; ipc.last = 0 }
func (ipc *ipcFifo) Push(val uint32) {
	ipc.fifo = append(ipc.fifo, val)
	if len(ipc.fifo) > 16 {
		panic("fifo overflow")
	}
}
func (ipc *ipcFifo) Top() uint32 {
	if ipc.Empty() {
		return ipc.last
	}
	return ipc.fifo[0]
}
func (ipc *ipcFifo) Pop() uint32 {
	if len(ipc.fifo) == 0 {
		panic("fifo underflow")
	}
	ipc.last = ipc.fifo[0]
	ipc.fifo = ipc.fifo[1:]
	return ipc.last
}

type HwIpc struct {
	HwIrq [2]*HwIrq

	Ipc9Sync hwio.Reg16 `hwio:"bank=0,offset=0x0,rcb,wcb=WriteIPCSYNC"`
	Ipc7Sync hwio.Reg16 `hwio:"bank=1,offset=0x0,rcb,wcb=WriteIPCSYNC"`

	data         [2]ipcFifo
	enable       [2]bool
	err          [2]bool
	irqEmptyFlag [2]bool
	irqDataFlag  [2]bool
}

func NewHwIpc(irq9 *HwIrq, irq7 *HwIrq) *HwIpc {
	ipc := new(HwIpc)
	ipc.HwIrq[CpuNds9] = irq9
	ipc.HwIrq[CpuNds7] = irq7

	hwio.InitRegs(ipc)
	return ipc
}

func (ipc *HwIpc) updateIrqFlagsCpu(cpunum CpuNum) {
	send := &ipc.data[cpunum]
	recv := &ipc.data[1-cpunum]

	newEmptyFlag := send.Empty() && send.emptyIrq
	newDataFlag := !recv.Empty() && recv.dataIrq

	// 0->1 transitions: raise irq
	if !ipc.irqEmptyFlag[cpunum] && newEmptyFlag {
		Emu.Log().Infof("[ipc] trigger IRQ send-empty on CPU %d", cpunum)
		ipc.HwIrq[cpunum].Raise(IrqIpcSendFifo)
	}
	if !ipc.irqDataFlag[cpunum] && newDataFlag {
		Emu.Log().Infof("[ipc] trigger IRQ recv-data-available on CPU %d", cpunum)
		ipc.HwIrq[cpunum].Raise(IrqIpcRecvFifo)
	}

	ipc.irqEmptyFlag[cpunum] = newEmptyFlag
	ipc.irqDataFlag[cpunum] = newDataFlag
}

func (ipc *HwIpc) updateIrqFlags() {
	ipc.updateIrqFlagsCpu(CpuNds9)
	ipc.updateIrqFlagsCpu(CpuNds7)
}

func (ipc *HwIpc) WriteIPCSYNC(_, value uint16) {
	if value&(1<<13) != 0 || value&(1<<14) != 0 {
		Emu.DebugBreak("[ipc] sync IRQ not implemented")
	}
}

func (ipc *HwIpc) ReadIPC7SYNC(value uint16) uint16 {
	value &^= 0xF
	value |= (ipc.Ipc9Sync.Value >> 8) & 0xF
	return value
}

func (ipc *HwIpc) ReadIPC9SYNC(value uint16) uint16 {
	value &^= 0xF
	value |= (ipc.Ipc7Sync.Value >> 8) & 0xF
	return value
}

func (ipc *HwIpc) ReadIPCFIFOCNT(cpunum CpuNum) uint16 {
	send := &ipc.data[cpunum]
	recv := &ipc.data[1-cpunum]
	cnt := uint16(0)
	if send.Empty() {
		cnt |= (1 << 0)
	}
	if send.Full() {
		cnt |= (1 << 1)
	}
	if send.emptyIrq {
		cnt |= (1 << 2)
	}
	if recv.Empty() {
		cnt |= (1 << 8)
	}
	if recv.Full() {
		cnt |= (1 << 9)
	}
	if recv.dataIrq {
		cnt |= (1 << 10)
	}
	if ipc.err[cpunum] {
		cnt |= (1 << 14)
	}
	if ipc.enable[cpunum] {
		cnt |= (1 << 15)
	}
	return cnt
}

func (ipc *HwIpc) WriteIPCFIFOCNT(cpunum CpuNum, val uint16) {
	send := &ipc.data[cpunum]
	recv := &ipc.data[1-cpunum]

	if val&(1<<3) != 0 {
		send.Flush()
	}
	send.emptyIrq = val&(1<<2) != 0
	recv.dataIrq = val&(1<<10) != 0
	ipc.enable[cpunum] = val&(1<<15) != 0
	if val&(1<<14) != 0 {
		ipc.err[cpunum] = false
	}
	Emu.Log().WithField("val", fmt.Sprintf("%04x", val)).Infof("[ipc] FIFO control")
	ipc.updateIrqFlags()
}

func (ipc *HwIpc) WriteIPCFIFOSEND(cpunum CpuNum, val uint32) {
	if ipc.enable[cpunum] {
		send := &ipc.data[cpunum]
		if send.Full() {
			ipc.err[cpunum] = true
		}
		send.Push(val)
	}
	Emu.Log().WithField("val", fmt.Sprintf("%08x", val)).Infof("[ipc] FIFO push")
	ipc.updateIrqFlags()
}

func (ipc *HwIpc) ReadIPCFIFORECV(cpunum CpuNum) uint32 {
	recv := &ipc.data[1-cpunum]
	if !ipc.enable[cpunum] {
		return recv.Top()
	}
	if recv.Empty() {
		ipc.err[cpunum] = true
		return recv.Top()
	}

	value := recv.Pop()
	Emu.Log().WithField("val", fmt.Sprintf("%08x", value)).Infof("[ipc] FIFO pop")
	ipc.updateIrqFlags()
	return value
}
