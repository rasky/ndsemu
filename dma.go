package main

import (
	"ndsemu/emu"
	"ndsemu/emu/hwio"

	"gopkg.in/Sirupsen/logrus.v0"
)

type DmaEvent int

const (
	DmaEventInvalid   DmaEvent = iota // invalid event (out-of-band value)
	DmaEventImmediate                 // immediate event (for immediate channels)
	DmaEventGamecard
)

type HwDmaFill struct {
	Dma0Fill hwio.Reg32 `hwio:"offset=0x00"`
	Dma1Fill hwio.Reg32 `hwio:"offset=0x04"`
	Dma2Fill hwio.Reg32 `hwio:"offset=0x08"`
	Dma3Fill hwio.Reg32 `hwio:"offset=0x0C"`
}

func NewHwDmaFill() *HwDmaFill {
	dma := new(HwDmaFill)
	hwio.MustInitRegs(dma)
	return dma
}

type HwDmaChannel struct {
	Cpu     CpuNum
	Channel int
	Bus     emu.Bus
	Irq     *HwIrq

	DmaSad   hwio.Reg32 `hwio:"offset=0x00"`
	DmaDad   hwio.Reg32 `hwio:"offset=0x04"`
	DmaCount hwio.Reg16 `hwio:"offset=0x08"`
	DmaCntrl hwio.Reg16 `hwio:"offset=0x0A,wcb"`

	inProgress   bool
	pendingEvent DmaEvent
}

func NewHwDmaChannel(cpu CpuNum, ch int, bus emu.Bus, irq *HwIrq) *HwDmaChannel {
	dma := &HwDmaChannel{
		Cpu:     cpu,
		Channel: ch,
		Bus:     bus,
		Irq:     irq,
	}
	hwio.MustInitRegs(dma)
	return dma
}

func (dma *HwDmaChannel) disable() {
	dma.DmaCntrl.Value &^= (1 << 15)
}

func (dma *HwDmaChannel) enabled() bool {
	return (dma.DmaCntrl.Value>>15)&1 != 0
}

// Return the event that will trigger the start of this DMA channel. If the
// channel is disabled, DmaEventInvalid is returned.
func (dma *HwDmaChannel) startEvent() DmaEvent {
	if !dma.enabled() {
		return DmaEventInvalid
	}

	if dma.Cpu == CpuNds9 {
		start := (dma.DmaCntrl.Value >> 11) & 7
		switch start {
		case 0:
			return DmaEventImmediate
		case 5:
			return DmaEventGamecard
		default:
			Emu.Log().Fatalf("DMA start=%d not implemented", start)
			return DmaEventInvalid
		}
	} else {
		start := (dma.DmaCntrl.Value >> 12) & 3
		switch start {
		case 0:
			return DmaEventImmediate
		case 2:
			return DmaEventGamecard
		default:
			Emu.Log().Fatalf("DMA start=%d not implemented", start)
			return DmaEventInvalid
		}
	}
}

func (dma *HwDmaChannel) WriteDMACNTRL(old, val uint16) {
	if (val>>15)&1 != 0 {
		dma.TriggerEvent(DmaEventImmediate)
		return
	}
}

func (dma *HwDmaChannel) xfer() {
	ctrl := dma.DmaCntrl.Value
	sad := dma.DmaSad.Value
	dad := dma.DmaDad.Value

	irq := (ctrl>>14)&1 != 0
	w32 := (ctrl>>10)&1 != 0
	repeat := (ctrl>>9)&1 != 0
	sinc := (ctrl >> 7) & 3
	dinc := (ctrl >> 5) & 3

	if sinc == 3 {
		Emu.Log().Fatal("sinc=3 should not happen")
	}

	cnt := uint32(dma.DmaCount.Value)
	if dma.Cpu == CpuNds9 {
		cnt |= (uint32(dma.DmaCntrl.Value) & 0x1F) << 16
		if cnt == 0 {
			cnt = 0x200000
		}
	} else {
		if cnt == 0 {
			if dma.Channel == 3 {
				cnt = 0x10000
			} else {
				cnt = 0x4000
			}
		}
	}

	wordsize := uint32(2)
	if w32 {
		wordsize = 4
	}

	if !repeat || cnt != 1 {
		Emu.Log().WithFields(logrus.Fields{
			"sad":  emu.Hex32(sad),
			"dad":  emu.Hex32(dad),
			"cnt":  emu.Hex32(cnt),
			"sinc": sinc,
			"dinc": dinc,
			"irq":  irq,
		}).Infof("[dma] transfer")
	}
	if dad == 0 {
		// nds9.Cpu.Exception(arm.ExceptionDataAbort)
		// Emu.DebugBreak("DMA to zero")
		return
	}

	dma.inProgress = true
	for ; cnt != 0; cnt-- {
		if w32 {
			dma.Bus.Write32(dad, dma.Bus.Read32(sad))
		} else {
			dma.Bus.Write16(dad, dma.Bus.Read16(sad))
		}

		if sinc == 0 || sinc == 3 {
			sad += wordsize
		} else if sinc == 1 {
			sad -= wordsize
		}
		if dinc == 0 || dinc == 3 {
			dad += wordsize
		} else if dinc == 1 {
			dad -= wordsize
		}
	}
	dma.inProgress = false

	if irq {
		dma.Irq.Raise(IrqDma0 << uint(dma.Channel))
	}

	if !repeat {
		dma.disable()
	} else {
		// Update registers for next repeat. Notice that these should be
		// internal copies of registers, but the external visible registers
		// are writeonly anyway, so we can reuse those for our own goals.
		dma.DmaSad.Value = sad

		// dest-increment 3 is "reload each repetition"
		if dinc != 3 {
			dma.DmaDad.Value = dad
		}
	}
}

func (dma *HwDmaChannel) TriggerEvent(event DmaEvent) {
	if event == DmaEventInvalid {
		Emu.Log().Fatalf("invalid DMA event triggered (?)")
	}

	if dma.inProgress {
		if dma.pendingEvent != DmaEventInvalid {
			Emu.Log().Fatalf("too many pending DMA events")
		}
		dma.pendingEvent = event
	} else {
		for event != DmaEventInvalid {
			if dma.startEvent() == event {
				dma.xfer()
			}
			// A new event might have been triggered while the DMA was in
			// progress (for instance, reading from gamecard triggers new
			// data to be ready and thus a new event to be scheduled). We
			// check here with a loop (instead of using recursion that would
			// grow the stack a lot, since tail recursion is not implemented
			// in the Go compiler)
			event = dma.pendingEvent
			dma.pendingEvent = DmaEventInvalid
		}
	}
}
