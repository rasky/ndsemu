package main

import (
	"encoding/binary"
	"ndsemu/emu"
	log "ndsemu/emu/logger"
)

var modTsc = log.NewModule("tsc")

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

func (ff *HwTouchScreen) SpiTransfer(data []byte) ([]byte, SpiStatus) {
	cmd := data[0]
	if cmd&0x80 == 0 {
		return nil, SpiFinish
	}

	powdown := cmd & 3
	ref := (cmd >> 2) & 1
	bits8 := (cmd>>3)&1 != 0
	adchan := (cmd >> 4) & 7

	modTsc.WithFields(log.Fields{
		"8bits":   bits8,
		"ref":     ref,
		"powdown": powdown,
		"value":   emu.Hex8(cmd),
	}).Infof("reading channel %s", tscChanNames[adchan])

	// Output value is always generated in the 12-bit range, and it is then
	// optionally truncated to 8 bit
	var output uint16
	switch adchan {
	case 0:
		output = 0x800
	case 1: // Y coord
		if ff.penDown {
			adcY1 := binary.LittleEndian.Uint16(Emu.Mem.Ram[0x3FFC80+0x5A:])
			scrY1 := Emu.Mem.Ram[0x3FFC80+0x5D]
			adcY2 := binary.LittleEndian.Uint16(Emu.Mem.Ram[0x3FFC80+0x60:])
			scrY2 := Emu.Mem.Ram[0x3FFC80+0x63]
			output = uint16((ff.penY-int(scrY1)+1)*int(adcY2-adcY1)/int(scrY2-scrY1) + int(adcY1))

			// log.Infof("coord Y:%d -> Out:%x", ff.penY, output)
		} else {
			output = 0xFFF
		}
	case 5: // X coord
		if ff.penDown {
			adcX1 := binary.LittleEndian.Uint16(Emu.Mem.Ram[0x3FFC80+0x58:])
			scrX1 := Emu.Mem.Ram[0x3FFC80+0x5C]
			adcX2 := binary.LittleEndian.Uint16(Emu.Mem.Ram[0x3FFC80+0x5E:])
			scrX2 := Emu.Mem.Ram[0x3FFC80+0x62]
			output = uint16((ff.penX-int(scrX1)+1)*int(adcX2-adcX1)/int(scrX2-scrX1) + int(adcX1))

			// log.Infof("coord :%d -> Out:%x", ff.penC, output)
		} else {
			output = 0x0
		}
	case 6: // microphone
		output = 0x0
	default:
		modTsc.Warnf("channel %s unimplemented", tscChanNames[adchan])
	}

	// While sending, there is always one initial 0 bit, so we always need
	// two bytes
	if !bits8 {
		return []byte{
			uint8(output >> 5), // 7 bit + leading 0
			uint8(output << 3),
		}, SpiFinish
	} else {
		output >>= 4
		return []byte{
			uint8(output >> 1),
			uint8(output << 7),
		}, SpiFinish
	}
}

func (ff *HwTouchScreen) SpiBegin() {}
func (ff *HwTouchScreen) SpiEnd()   {}
