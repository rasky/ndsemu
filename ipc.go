package main

import (
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modIpc = log.NewModule("ipc")

type ipcFifo struct {
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

	Ipc9Sync     hwio.Reg16 `hwio:"bank=0,offset=0x0,rwmask=0xFF00,wcb"`
	Ipc7Sync     hwio.Reg16 `hwio:"bank=2,offset=0x0,rwmask=0xFF00,wcb"`
	Ipc9FifoCnt  hwio.Reg16 `hwio:"bank=0,offset=0x4,rcb,wcb"`
	Ipc7FifoCnt  hwio.Reg16 `hwio:"bank=2,offset=0x4,rcb,wcb"`
	Ipc9FifoSend hwio.Reg32 `hwio:"bank=0,offset=0x8,writeonly,wcb"`
	Ipc7FifoSend hwio.Reg32 `hwio:"bank=2,offset=0x8,writeonly,wcb"`
	Ipc9FifoRecv hwio.Reg32 `hwio:"bank=1,offset=0x0,readonly,rcb"`
	Ipc7FifoRecv hwio.Reg32 `hwio:"bank=3,offset=0x0,readonly,rcb"`

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

	hwio.MustInitRegs(ipc)
	return ipc
}

func (ipc *HwIpc) updateIrqFlagsCpu(cpunum CpuNum) {
	send := &ipc.data[cpunum]
	recv := &ipc.data[1-cpunum]

	newEmptyFlag := send.Empty() && send.emptyIrq
	newDataFlag := !recv.Empty() && recv.dataIrq

	// 0->1 transitions: raise irq
	if !ipc.irqEmptyFlag[cpunum] && newEmptyFlag {
		modIpc.InfoZ("trigger IRQ send-empty").Int("cpu", int(cpunum)).End()
		ipc.HwIrq[cpunum].Raise(IrqIpcSendFifo)
	}
	if !ipc.irqDataFlag[cpunum] && newDataFlag {
		modIpc.InfoZ("trigger IRQ recv-data-available").Int("cpu", int(cpunum)).End()
		ipc.HwIrq[cpunum].Raise(IrqIpcRecvFifo)
	}

	ipc.irqEmptyFlag[cpunum] = newEmptyFlag
	ipc.irqDataFlag[cpunum] = newDataFlag
}

func (ipc *HwIpc) updateIrqFlags() {
	ipc.updateIrqFlagsCpu(CpuNds9)
	ipc.updateIrqFlagsCpu(CpuNds7)
}

func (ipc *HwIpc) WriteIPC7SYNC(_, value uint16) {
	// See WriteIPC9SYNC comment for why this is required
	nds9.Run(Emu.Sync.Cycles())

	ipc.Ipc9Sync.Value &^= 0xF
	ipc.Ipc9Sync.Value |= (value >> 8) & 0xF
	if value&(1<<13) != 0 && ipc.Ipc9Sync.Value&(1<<14) != 0 {
		ipc.HwIrq[CpuNds9].Raise(IrqIpcSync)
	}
}

func (ipc *HwIpc) WriteIPC9SYNC(_, value uint16) {
	// Force a sync between the CPUs when the SYNC is being written. This can be necessary
	// in case there is some strict-timing communication going on between the two CPUs.
	// Example: when booting a homebrew ROM, ARM7 (at around 0x80000c0) patches the code
	// that is being run by ARM9 (which is sitting in a tight loop writing to IPC9SYNC),
	// making it jump elsewhere; immediatley after, the ARM7 memcpy's
	// over the ARM9 tight loop, assuming that it has already jumped away.
	// This breaks emulation if we don't sync between the CPUs quick enough.
	nds7.Run(Emu.Sync.Cycles())

	ipc.Ipc7Sync.Value &^= 0xF
	ipc.Ipc7Sync.Value |= (value >> 8) & 0xF
	if value&(1<<13) != 0 && ipc.Ipc7Sync.Value&(1<<14) != 0 {
		ipc.HwIrq[CpuNds7].Raise(IrqIpcSync)
	}
}

func (ipc *HwIpc) ReadIPC9FIFOCNT(_ uint16) uint16 { return ipc.readIPCFIFOCNT(CpuNds9) }
func (ipc *HwIpc) ReadIPC7FIFOCNT(_ uint16) uint16 { return ipc.readIPCFIFOCNT(CpuNds7) }
func (ipc *HwIpc) readIPCFIFOCNT(cpunum CpuNum) uint16 {
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

func (ipc *HwIpc) WriteIPC9FIFOCNT(_, val uint16) { ipc.writeIPCFIFOCNT(CpuNds9, val) }
func (ipc *HwIpc) WriteIPC7FIFOCNT(_, val uint16) { ipc.writeIPCFIFOCNT(CpuNds7, val) }
func (ipc *HwIpc) writeIPCFIFOCNT(cpunum CpuNum, val uint16) {
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
	modIpc.InfoZ("FIFO control").Hex16("val", val).End()
	ipc.updateIrqFlags()
}

func (ipc *HwIpc) WriteIPC9FIFOSEND(_, val uint32) { ipc.writeIPCFIFOSEND(CpuNds9, val) }
func (ipc *HwIpc) WriteIPC7FIFOSEND(_, val uint32) { ipc.writeIPCFIFOSEND(CpuNds7, val) }
func (ipc *HwIpc) writeIPCFIFOSEND(cpunum CpuNum, val uint32) {
	if ipc.enable[cpunum] {
		send := &ipc.data[cpunum]
		if send.Full() {
			ipc.err[cpunum] = true
		}
		send.Push(val)
	}
	modIpc.InfoZ("FIFO push").Hex32("val", val).End()
	ipc.updateIrqFlags()
}

func (ipc *HwIpc) ReadIPC9FIFORECV(_ uint32) uint32 { return ipc.readIPCFIFORECV(CpuNds9) }
func (ipc *HwIpc) ReadIPC7FIFORECV(_ uint32) uint32 { return ipc.readIPCFIFORECV(CpuNds7) }
func (ipc *HwIpc) readIPCFIFORECV(cpunum CpuNum) uint32 {
	recv := &ipc.data[1-cpunum]
	if !ipc.enable[cpunum] {
		return recv.Top()
	}
	if recv.Empty() {
		ipc.err[cpunum] = true
		return recv.Top()
	}

	value := recv.Pop()
	modIpc.InfoZ("FIFO pop").Hex32("val", value).End()
	ipc.updateIrqFlags()
	return value
}
