package main

import (
	"ndsemu/emu"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

type DmaEvent int

const (
	DmaEventInvalid   DmaEvent = iota // invalid event (out-of-band value)
	DmaEventImmediate                 // immediate event (for immediate channels)
	DmaEventGamecard
	DmaEventHBlank
	DmaEventGxFifo
	DmaEventGbaSoundFifo
	DmaEventGbaVideoCapture
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

	debugRepeat  bool
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

	switch {
	case dma.Cpu == CpuNds9:
		start := (dma.DmaCntrl.Value >> 11) & 7
		switch start {
		case 0:
			return DmaEventImmediate
		case 2:
			return DmaEventHBlank
		case 5:
			return DmaEventGamecard
		case 7:
			return DmaEventGxFifo
		default:
			log.ModDma.FatalZ("DMA start not implemented").Uint16("event", start).End()
			return DmaEventInvalid
		}
	case dma.Cpu == CpuNds7 && Emu.Mode == ModeNds:
		start := (dma.DmaCntrl.Value >> 12) & 3
		switch start {
		case 0:
			return DmaEventImmediate
		case 2:
			return DmaEventGamecard
		default:
			log.ModDma.FatalZ("DMA start not implemented").Uint16("event", start).End()
			return DmaEventInvalid
		}
	case dma.Cpu == CpuNds7 && Emu.Mode == ModeGba:
		start := (dma.DmaCntrl.Value >> 12) & 3
		switch start {
		case 0:
			return DmaEventImmediate
		case 2:
			return DmaEventHBlank
		case 3:
			switch dma.Channel {
			case 0:
				log.ModDma.ErrorZ("DMA start 3 prohibited on channel 0").End()
				return DmaEventInvalid
			case 1, 2:
				log.ModDma.WarnZ("DMA sound FIFO not implemented").End()
				return DmaEventGbaSoundFifo
			case 3:
				log.ModDma.WarnZ("DMA video capture not implemented").End()
				return DmaEventGbaVideoCapture
			}
		default:
			log.ModDma.FatalZ("DMA start not implemented").Uint16("event", start).End()
			return DmaEventInvalid
		}
	}

	panic("unreachable")
}

func (dma *HwDmaChannel) WriteDMACNTRL(old, val uint16) {
	dma.debugRepeat = false

	// Check if this write activated a DMA channel. If it did,
	// we might to do something right away, depending on the start
	// event type.
	if ((old^val)>>15)&1 != 0 {
		evt := dma.startEvent()
		switch evt {
		case DmaEventImmediate:
			// DMA in immediate mode must be triggered immediately
			dma.TriggerEvent(DmaEventImmediate)

		case DmaEventGxFifo:
			// Sync up the geometry engine up to the current point
			// before making it see that the DMA is active; this
			// way, we get cycle-accurate emulation. Whenever the
			// geometry engine is then scheduled, it will trigger this
			// DMA (assuming the FIFO is empty enough).
			dma.DmaCntrl.Value = old
			Emu.Hw.Geom.Run(Emu.Sync.Cycles())
			dma.DmaCntrl.Value = val
		}
	}
}

func (dma *HwDmaChannel) xfer() {
	ctrl := dma.DmaCntrl.Value
	sad := dma.DmaSad.Value
	dad := dma.DmaDad.Value

	irq := (ctrl>>14)&1 != 0
	start := (ctrl >> 11) & 7
	w32 := (ctrl>>10)&1 != 0
	repeat := (ctrl>>9)&1 != 0
	sinc := (ctrl >> 7) & 3
	dinc := (ctrl >> 5) & 3

	if sinc == 3 {
		log.ModDma.FatalZ("sinc=3 should not happen").End()
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

	if !repeat || !dma.debugRepeat {
		if repeat {
			dma.debugRepeat = true
		}
		log.ModDma.InfoZ("transfer").
			Hex32("sad", sad).
			Hex32("dad", dad).
			Hex32("cnt", cnt).
			Uint16("sinc", sinc).
			Uint16("dinc", dinc).
			Bool("irq", irq).
			Uint32("wsize", wordsize).
			End()
	}
	if dad == 0 {
		// nds9.Cpu.Exception(arm.ExceptionDataAbort)
		// Emu.DebugBreak("DMA to zero")
		dma.disable()
		return
	}

	if dma.Cpu == CpuNds9 && start == 7 {
		// GFXFIFO dma is different from others because it is technically
		// a single-transfer, while actually data is flushed in batches
		// of 112 words. So we need to trick this function into repeat mode
		// and avoid triggering irq, unless the transfer is really finished.
		if cnt > 112 {
			irq = false
			repeat = true
			dma.DmaCount.Value = uint16(cnt - 112)
			cnt = 112
		}
	}

	dma.inProgress = true
	for ; cnt != 0; cnt-- {
		if w32 {
			dma.Bus.Write32(dad, dma.Bus.Read32(sad))
		} else {
			dma.Bus.Write16(dad, dma.Bus.Read16(sad))
		}

		// Notify jit engine that we wrote to that address
		if jit := nds9.Cpu.Jit(); jit != nil {
			jit.Invalidate(dad)
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
		log.ModDma.FatalZ("invalid DMA event triggered (?)").End()
	}

	if dma.inProgress {
		// Event GxFifo is scheduled by the Geometry engine any time the
		// FIFO is less than half full. Since the engine can also run in
		// the middle of a DMA trasnfer, it might happen that there
		// are multiple calls pending (ex: super mario 64). Ignore it.
		if event == DmaEventGxFifo && dma.pendingEvent == DmaEventGxFifo {
			return
		}
		if dma.pendingEvent != DmaEventInvalid {
			log.ModDma.FatalZ("too many pending DMA events").End()
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
