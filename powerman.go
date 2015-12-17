package main

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwPowerMan struct {
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
	data := recv(0)
	if index&0x80 == 0 {
		switch index & 0x7F {
		default:
			log.Infof("[powerman] write reg %d: %02x", index&0x7F, data)
		}
	} else {
		log.Infof("[powerman] read reg %d", index&0x7F)
	}
}

func (ff *HwPowerMan) BeginTransfer() chan uint8 {
	ch := make(chan uint8)
	go ff.transfer(ch)
	return ch
}
