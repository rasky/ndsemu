package main

import (
	"ndsemu/emu/hw"
	log "ndsemu/emu/logger"
)

var modPower = log.NewModule("powerman")

type HwPowerMan struct {
	cntrl uint8
}

func NewHwPowerMan() *HwPowerMan {
	return &HwPowerMan{}
}

func (ff *HwPowerMan) transfer(ch chan uint8) {
	recv := func(val uint8) uint8 {
		data := <-ch
		ch <- val
		return data
	}

	index := recv(0)
	if index&0x80 == 0 {
		data := recv(0)
		switch index & 0x7F {
		case 0:
			ff.cntrl = data
			modPower.Infof("write control: %02x", data)
		default:
			modPower.Infof("write reg %d: %02x", index&0x7F, data)
		}
	} else {
		var data uint8
		switch index & 0x7F {
		case 0:
			data = ff.cntrl
		case 1:
			// Bit 0: if set, battery is finishing
			if hw.ReadBatteryStatus != nil && hw.ReadBatteryStatus() < 10 {
				data = 0x1
			}
		default:
			modPower.Infof("read reg %d", index&0x7F)
		}
		recv(data)
	}
}

func (ff *HwPowerMan) BeginTransfer() chan uint8 {
	ch := make(chan uint8)
	go ff.transfer(ch)
	return ch
}
