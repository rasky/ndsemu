package main

import (
	"ndsemu/emu"
	"ndsemu/emu/hwio"

	"gopkg.in/Sirupsen/logrus.v0"
)

type HwDmaChannel struct {
	Cpu     CpuNum
	Channel int
	Bus     emu.Bus
	Irq     *HwIrq

	DmaSad   hwio.Reg32 `hwio:"offset=0x00"`
	DmaDad   hwio.Reg32 `hwio:"offset=0x04"`
	DmaCount hwio.Reg16 `hwio:"offset=0x08"`
	DmaCntrl hwio.Reg16 `hwio:"offset=0x0A,wcb"`
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

func (dma *HwDmaChannel) WriteDMACNTRL(old, val uint16) {
	if (val>>15)&1 == 0 {
		return
	}

	irq := (val>>14)&1 != 0
	start := (val >> 13) & 3
	w32 := (val>>10)&1 != 0
	repeat := (val>>9)&1 != 0
	sinc := (val >> 7) & 3
	dinc := (val >> 5) & 3

	if repeat {
		Emu.Log().Fatal("DMA repeat not implemented")
	}
	if sinc == 3 {
		Emu.Log().Fatal("sinc=3 should not happen")
	}
	if start != 0 {
		Emu.Log().Fatalf("DMA start=%d not implemented", start)
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

	sad := dma.DmaSad.Value
	dad := dma.DmaDad.Value

	Emu.Log().WithFields(logrus.Fields{
		"sad":  emu.Hex32(sad),
		"dad":  emu.Hex32(dad),
		"cnt":  emu.Hex32(cnt),
		"sinc": sinc,
		"dinc": dinc,
		"irq":  irq,
	}).Infof("[dma] transfer")

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

	if irq {
		dma.Irq.Raise(IrqDma0 << uint(dma.Channel))
	}

	// Signal that transfer is finished
	dma.DmaCntrl.Value &^= (1 << 15)
}
