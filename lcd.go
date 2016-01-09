package main

import (
	"ndsemu/emu/hwio"

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
	Irq *HwIrq

	DispStat hwio.Reg16 `hwio:"offset=4,rwmask=0xFFF8"`
	VCount   hwio.Reg16 `hwio:"offset=6,readonly"`
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
	return uint16(y)
}

func (lcd *HwLcd) SyncEvent(x, y int) {
	switch x {
	case 0:
		if y == cVBlankFirstLine {
			if lcd.DispStat.Value&cVBlankIrq != 0 {
				log.Info("[LCD] VBlank IRQ")
				lcd.Irq.Raise(IrqVBlank)
			}
		}

		vmatch := int(lcd.DispStat.Value>>8 | (lcd.DispStat.Value&0x80)<<1)
		if y == vmatch && lcd.DispStat.Value&cVMatchIrq != 0 {
			log.Info("[LCD] VMatch IRQ on NDS9")
			lcd.Irq.Raise(IrqVMatch)
		}

	case cHBlankFirstDot:
		if !(y >= cVBlankFirstLine && y <= cVBlankLastLine) {
			if lcd.DispStat.Value&cHBlankIrq != 0 {
				log.Info("[LCD] HBlank IRQ on NDS9")
				lcd.Irq.Raise(IrqHBlank)
			}
		}
	default:
		// panic("unreachable")
	}
}
