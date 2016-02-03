package main

import (
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"time"
)

var modRtc = log.NewModule("rtc")

type SerialDevice interface {
	ReadData() uint8
	WriteData(val uint8)
}

type HwSerial3W struct {
	cs      bool
	clk     bool
	datadir bool
	data    uint8
	cnt     int

	dev SerialDevice

	Serial hwio.Reg8 `hwio:"rcb,wcb"`
}

func (s *HwSerial3W) WriteSERIAL(_, val uint8) {
	cs := val&(1<<2) != 0
	clk := val&(1<<1) != 0
	datadir := val&(1<<4) != 0

	// log.Infof("cs=%v clk=%v datadir=%v data=%d", cs, clk, datadir, val&1)

	// select rising edge: begin new transfer
	// select and change direction: begin new transfer
	if cs && !s.cs {
		s.data = 0
		if datadir {
			s.cnt = 0
		} else {
			s.cnt = 8
		}
	}

	// select=1 -> active
	// falling edge -> transfer
	if cs && !clk && s.clk {
		if datadir {
			// WRITE
			s.data >>= 1
			s.data |= (val & 1) << 7
			s.cnt++
			if s.cnt == 8 {
				// log.Infof("writing byte: %x", s.data)
				s.dev.WriteData(s.data)
				s.cnt = 0
			}
		} else {
			// READ
			s.data >>= 1
			s.cnt++
			if s.cnt == 8 || s.datadir {
				s.data = s.dev.ReadData()
				s.cnt = 0
				// log.Infof("begin next byte(2): %x (first bit: %d)", s.data, s.data&1)
			} else {
				// log.Infof("prepare next bit: idx=%d, val=%d", s.cnt, s.data&1)
			}
		}
	}

	s.cs = cs
	s.clk = clk
	s.datadir = datadir
}

func (s *HwSerial3W) ReadSERIAL(_ uint8) uint8 {
	val := uint8(0x60)
	if s.clk {
		val |= (1 << 1)
	}
	if s.cs {
		val |= (1 << 2)
	}
	if s.datadir {
		val |= (1 << 4)
	}
	val |= (s.data & 1)
	// log.Infof("reading bit: %d", s.data&1)
	return val
}

// Seiko S-35180
type HwRtc struct {
	HwSerial3W

	regStatus1 uint8
	regStatus2 uint8

	writing bool
	buf     []byte
	idx     int
}

func NewHwRtc() *HwRtc {
	rtc := new(HwRtc)
	rtc.regStatus1 = 0x00 // 0x80: reset to defaults
	rtc.regStatus2 = 0x00
	rtc.HwSerial3W.dev = rtc
	hwio.MustInitRegs(&rtc.HwSerial3W)
	return rtc
}

func (rtc *HwRtc) ResetDefaults() {
	rtc.regStatus1 = 0x80
	rtc.regStatus2 = 0x00
}

func (rtc *HwRtc) ReadData() uint8 {
	if rtc.writing {
		modRtc.Warnf("read during register writing")
		return 0
	}

	if rtc.idx >= len(rtc.buf) {
		modRtc.Warnf("read but not data setup")
		return 0
	}

	data := rtc.buf[rtc.idx]
	rtc.idx++
	return data
}

func (rtc *HwRtc) bcd(value uint) uint8 {
	if value > 99 {
		modRtc.Warnf("cannot convert value %d to BCD", value)
		return 0xFF
	}

	value = (value/10)*16 + (value % 10)
	return uint8(value)
}

var rtcRegnames = [8]string{"sr1", "alarm1", "datetime", "clockadjust", "sr2", "alarm2", "time", "unused"}

func (rtc *HwRtc) writeReg(val uint8) {
	reglen := [8]int{1, 3, 7, 1, 1, 3, 3, 1}

	rtc.buf = append(rtc.buf, val)
	if len(rtc.buf) != reglen[rtc.idx] {
		modRtc.Warnf("partial writing reg %q: %02x", rtcRegnames[rtc.idx], val)
		return
	}
	rtc.writing = false
	modRtc.Warnf("final writing reg %q: %02x", rtcRegnames[rtc.idx], val)

	switch rtc.idx {
	case 0: // sr1
		rtc.regStatus1 = (rtc.regStatus1 & 0xF0) | (val & 0xE)
		modRtc.Infof("write sr1: %02x", val)
	case 4: // sr2
		rtc.regStatus2 = val
		modRtc.Infof("write sr2: %02x", val)
	default:
		modRtc.Warnf("unimplemented register write: %q=%x", rtcRegnames[rtc.idx], rtc.buf)
	}
}

func (rtc *HwRtc) WriteData(val uint8) {
	if rtc.writing {
		rtc.writeReg(val)
		return
	}

	if val&0xF != 6 {
		modRtc.Warnf("invalid command %02x", val)
		return
	}

	read := val&0x80 != 0
	reg := (val >> 4) & 7

	if !read {
		modRtc.Warnf("begin writing reg %q", rtcRegnames[reg])
		rtc.writing = true
		rtc.buf = nil
		rtc.idx = int(reg)
		return
	}

	rtc.buf = nil
	rtc.idx = 0
	switch reg {
	case 0: // sr1
		rtc.buf = append(rtc.buf, rtc.regStatus1)
		// Bit 4-7 are auto-cleared after read
		// (though currently we don't set them in our emulation...)
		rtc.regStatus1 &= 0x0F
	case 2: // datetime
		now := time.Now()

		var hour uint8
		if rtc.regStatus1&2 != 0 {
			// 24H mode
			hour = rtc.bcd(uint(now.Hour()))
		} else {
			// 12H mode, with 12:00 that becomes 0pm instead of 12pm (as per
			// normale human convention)
			hour = rtc.bcd(uint(now.Hour() % 12))
			if now.Hour() >= 12 {
				hour |= 0x40
			}
		}

		rtc.buf = append(rtc.buf,
			rtc.bcd(uint(now.Year()-2000)),
			rtc.bcd(uint(now.Month())),
			rtc.bcd(uint(now.Day())),
			rtc.bcd(uint(now.Weekday())),
			hour,
			rtc.bcd(uint(now.Minute())),
			rtc.bcd(uint(now.Second())),
		)
	case 4:
		rtc.buf = append(rtc.buf, rtc.regStatus2)
	default:
		modRtc.Warnf("unimplemented register read %q", rtcRegnames[reg])
		return
	}

	modRtc.Infof("read %q: %x", rtcRegnames[reg], rtc.buf)
}
