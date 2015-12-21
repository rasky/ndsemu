package main

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwTouchScreen struct {
}

func NewHwTouchScreen() *HwTouchScreen {
	return &HwTouchScreen{}
}

var tscChanNames = [8]string{
	"temp0", "touch_y", "battery", "touch_z1",
	"touch_z2", "touch_x", "aux", "temp1",
}

func (ff *HwTouchScreen) transfer(ch chan uint8) {
	recv := func(val uint8) uint8 {
		data := <-ch
		ch <- val
		return data
	}

	cmd := recv(0)
	powdown := cmd & 3
	ref := (cmd >> 2) & 1
	bits8 := (cmd>>3)&1 != 0
	adchan := (cmd >> 4) & 7

	log.WithFields(log.Fields{
		"8bits":   bits8,
		"ref":     ref,
		"powdown": powdown,
	}).Infof("[tsc] reading channel %s", tscChanNames[adchan])

	// Output value is always generated in the 12-bit range, and it is then
	// optionally truncated to 8 bit
	var output uint16
	switch adchan {
	default:
		log.Warnf("[tsc] channel %s unimplemented", tscChanNames[adchan])
	}

	// While sending, there is always one initial 0 bit, so we always need
	// two bytes
	if !bits8 {
		recv(uint8(output >> 5)) // 7 bit + leading 0
		recv(uint8(output << 3))
	} else {
		output >>= 4
		recv(uint8(output >> 1))
		recv(uint8(output << 7))
	}
}

func (ff *HwTouchScreen) BeginTransfer() chan uint8 {
	ch := make(chan uint8)
	go ff.transfer(ch)
	return ch
}
