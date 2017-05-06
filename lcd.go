package main

import (
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modLcd = log.NewModule("lcd")

const (
	cVBlankFlag = (1 << 0)
	cHBlankFlag = (1 << 1)
	cVMatchFlag = (1 << 2)
	cVBlankIrq  = (1 << 3)
	cHBlankIrq  = (1 << 4)
	cVMatchIrq  = (1 << 5)

	cVBlankFirstLine = 192
	cVBlankLastLine  = 261

	cHBlankFirstDot = 268
)

type HwLcd struct {
	Irq *HwIrq

	DispStat hwio.Reg16 `hwio:"offset=4,rwmask=0xFFF8,rcb"`
	VCount   hwio.Reg16 `hwio:"offset=6,readonly,rcb"`
}

func NewHwLcd(irq *HwIrq) *HwLcd {
	lcd := &HwLcd{Irq: irq}
	hwio.MustInitRegs(lcd)
	return lcd
}

func (lcd *HwLcd) ReadDISPSTAT(stat uint16) uint16 {
	x, y := Emu.Sync.DotPos()

	// VBlank: not set on line 227
	if y >= cVBlankFirstLine && y <= cVBlankLastLine {
		stat |= cVBlankFlag
	}

	// HBlank: the flag is kept to 0 for 268 dots / ~1600 cycles (even
	// if the screen is 256 dots)
	if x > cHBlankFirstDot {
		stat |= cHBlankFlag
	}

	// VMatch
	vmatch := int(stat>>8 | (stat&0x80)<<1)
	if y == vmatch {
		stat |= cVMatchFlag
	}

	return stat
}

func (lcd *HwLcd) ReadVCOUNT(_ uint16) uint16 {
	_, y := Emu.Sync.DotPos()
	return uint16(y) & 0x1FF
}

func (lcd *HwLcd) SyncEvent(x, y int) {
	switch x {
	case 0:
		if y == cVBlankFirstLine {
			if lcd.DispStat.Value&cVBlankIrq != 0 {
				modLcd.InfoZ("VBlank IRQ").String("irq", lcd.Irq.Name).End()
				lcd.Irq.Raise(IrqVBlank)
			}
		}

		vmatch := int(lcd.DispStat.Value>>8 | (lcd.DispStat.Value&0x80)<<1)
		if y == vmatch && lcd.DispStat.Value&cVMatchIrq != 0 {
			modLcd.InfoZ("VMatch IRQ").String("irq", lcd.Irq.Name).Int("line", y).End()
			lcd.Irq.Raise(IrqVMatch)
		}

	case cHBlankFirstDot:
		if !(y >= cVBlankFirstLine && y <= cVBlankLastLine) {
			if lcd.DispStat.Value&cHBlankIrq != 0 {
				// modLcd.WithField("irq", lcd.Irq.Name).Info("HBlank IRQ")
				lcd.Irq.Raise(IrqHBlank)
			}
		}
	default:
		// panic("unreachable")
	}
}
