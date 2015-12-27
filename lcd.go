package main

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

const (
	cVBlankFlag = (1 << 0)
	cHBlankFlag = (1 << 1)
	cVMatchFlag = (1 << 2)
	cVBlankIrq  = (1 << 3)
	cHBlankIrq  = (1 << 4)
	cVMatchIrq  = (1 << 5)

	cVBlankFirstLine = 160
	cVBlankLastLine  = 226

	cHBlankFirstDot = 268
)

type HwLcd struct {
	Irq9 *HwIrq
	Irq7 *HwIrq

	dispstat uint16
}

func NewHwLcd(irq9 *HwIrq, irq7 *HwIrq) *HwLcd {
	return &HwLcd{Irq9: irq9, Irq7: irq7}
}

func (lcd *HwLcd) WriteDISPSTAT(val uint16) {
	lcd.dispstat = val &^ 0x7
}

func (lcd *HwLcd) ReadDISPSTAT() uint16 {
	stat := lcd.dispstat
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

func (lcd *HwLcd) ReadVCOUNT() uint16 {
	_, y := Emu.Sync.DotPos()
	return uint16(y)
}

func (lcd *HwLcd) SyncEvent(x, y int) {
	switch x {
	case 0:
		if y == cVBlankFirstLine {
			if lcd.dispstat&cVBlankIrq != 0 {
				log.Info("[LCD] VBlank IRQ")
				lcd.Irq9.Raise(IrqVBlank)
				lcd.Irq7.Raise(IrqVBlank)
			}
		}
		vmatch := int(lcd.dispstat>>8 | (lcd.dispstat&0x80)<<1)
		if y == vmatch {
			if lcd.dispstat&cVMatchIrq != 0 {
				log.Info("[LCD] VMatch IRQ")
				lcd.Irq9.Raise(IrqVMatch)
				lcd.Irq7.Raise(IrqVMatch)
			}
		}
	case cHBlankFirstDot:
		if !(y >= cVBlankFirstLine && y <= cVBlankLastLine) {
			if lcd.dispstat&cHBlankIrq != 0 {
				log.Info("[LCD] HBlank IRQ")
				lcd.Irq9.Raise(IrqHBlank)
				lcd.Irq7.Raise(IrqHBlank)
			}
		}
	default:
		// panic("unreachable")
	}
}
