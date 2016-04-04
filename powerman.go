package main

import log "ndsemu/emu/logger"

var modPower = log.NewModule("powerman")

type HwPowerMan struct {
	cntrl uint8
}

func NewHwPowerMan() *HwPowerMan {
	return &HwPowerMan{}
}

func (ff *HwPowerMan) SpiTransfer(data []byte) ([]byte, SpiStatus) {
	index := data[0]
	if index&0x80 == 0 {
		// Write reg
		if len(data) < 2 {
			return nil, SpiContinue
		}
		val := data[1]
		switch index & 0x7F {
		case 0:
			ff.cntrl = val
			modPower.Infof("write control: %02x", data)
		default:
			modPower.Infof("write reg %d: %02x", index&0x7F, val)
		}
		return nil, SpiFinish
	} else {
		// Read reg
		switch index & 0x7F {
		case 0:
			return []byte{ff.cntrl}, SpiFinish
		default:
			modPower.Infof("read reg %d", index&0x7F)
			return nil, SpiFinish
		}
	}
}

func (ff *HwPowerMan) SpiBegin() {}
func (ff *HwPowerMan) SpiEnd()   {}
