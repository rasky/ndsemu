package main

import (
	"ndsemu/emu"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwTouchScreen struct {
	penX, penY int
	penDown    bool
}

func NewHwTouchScreen() *HwTouchScreen {
	return &HwTouchScreen{}
}

var tscChanNames = [8]string{
	"temp0", "touch_y", "battery", "touch_z1",
	"touch_z2", "touch_x", "aux", "temp1",
}

func (ff *HwTouchScreen) SetPen(down bool, x, y int) {
	ff.penX = x
	ff.penY = y
	ff.penDown = down
}

func (ff *HwTouchScreen) transfer(ch chan uint8) {
	recv := func(val uint8) uint8 {
		data, ok := <-ch
		if !ok {
			return 0
		}
		ch <- val
		return data
	}

	var cmd uint8
	for cmd = range ch {
		ch <- 0
		if cmd&0x80 != 0 {
			break
		}
	}

	if cmd&0x80 == 0 {
		return
	}

start:
	powdown := cmd & 3
	ref := (cmd >> 2) & 1
	bits8 := (cmd>>3)&1 != 0
	adchan := (cmd >> 4) & 7

	Emu.Log().WithFields(log.Fields{
		"8bits":   bits8,
		"ref":     ref,
		"powdown": powdown,
		"value":   emu.Hex8(cmd),
	}).Infof("[tsc] reading channel %s", tscChanNames[adchan])

	// Output value is always generated in the 12-bit range, and it is then
	// optionally truncated to 8 bit
	var output uint16
	switch adchan {
	case 0:
		output = 0x800
	case 1: // Y coord
		if ff.penDown {
			output = uint16(ff.penY) << 4
			log.Info("[tsc] coord Y:", ff.penY)
		} else {
			output = 0xFFF
		}
	case 5: // X coord
		if ff.penDown {
			output = uint16(ff.penX) << 4
			log.Info("[tsc] coord X:", ff.penX)
		} else {
			output = 0x0
		}
	default:
		log.Warnf("[tsc] channel %s unimplemented", tscChanNames[adchan])
	}

	// While sending, there is always one initial 0 bit, so we always need
	// two bytes
	if !bits8 {
		recv(uint8(output >> 5)) // 7 bit + leading 0
		cmd = recv(uint8(output << 3))
	} else {
		output >>= 4
		recv(uint8(output >> 1))
		cmd = recv(uint8(output << 7))
	}

	if cmd&0x80 != 0 {
		goto start
	}

	for _ = range ch {
		ch <- 0
	}
}

func (ff *HwTouchScreen) BeginTransfer() chan uint8 {
	ch := make(chan uint8)
	go ff.transfer(ch)
	return ch
}
