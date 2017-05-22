package main

import (
	"ndsemu/emu/hw"
	"ndsemu/emu/spi"

	log "ndsemu/emu/logger"
)

var modPower = log.NewModule("powerman")

type HwPowerMan struct {
	cntrl   uint8
	mic     bool
	micgain int
}

func NewHwPowerMan() *HwPowerMan {
	return &HwPowerMan{}
}

func (pow *HwPowerMan) AudioEnabled() bool {
	return pow.cntrl&(1<<0) != 0 && pow.cntrl&(1<<1) == 0
}

func (ff *HwPowerMan) SpiTransfer(data []byte) ([]byte, spi.ReqStatus) {
	index := data[0]
	if index&0x80 == 0 {
		// Write reg
		if len(data) < 2 {
			return nil, spi.ReqContinue
		}
		val := data[1]
		switch index & 0x7F {
		case 0:
			ff.cntrl = val
			modPower.InfoZ("write control").Hex8("val", val).End()
		case 2:
			ff.mic = val&1 != 0
			modPower.InfoZ("enable microphone").End()
		case 3:
			ff.micgain = 20 * int((val&3)+1)
			modPower.InfoZ("set microphone gain").Int("gain", ff.micgain).End()
		default:
			modPower.WarnZ("write unknown reg").Uint8("reg", index&0x7F).Hex8("val", val).End()
		}
		return nil, spi.ReqFinish
	} else {
		// Read reg
		switch index & 0x7F {
		case 0:
			return []byte{ff.cntrl}, spi.ReqFinish
		case 1:
			// Bit 0: if set, battery is finishing
			val := uint8(0)
			if hw.ReadBatteryStatus != nil && hw.ReadBatteryStatus() < 10 {
				val |= 1
			}
			return []byte{val}, spi.ReqFinish
		case 2:
			val := uint8(0)
			if ff.mic {
				val |= 1
			}
			return []byte{val}, spi.ReqFinish
		default:
			modPower.WarnZ("read unknown reg").Uint8("reg", index&0x7F).End()
			return nil, spi.ReqFinish
		}
	}
}

func (ff *HwPowerMan) SpiBegin() {}
func (ff *HwPowerMan) SpiEnd()   {}
