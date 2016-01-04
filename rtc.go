package main

import (
	"ndsemu/emu/hwio"
	"time"

	log "gopkg.in/Sirupsen/logrus.v0"
)

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

	// log.Infof("[s3w] cs=%v clk=%v datadir=%v data=%d", cs, clk, datadir, val&1)

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
				s.dev.WriteData(s.data)
				s.cnt = 0
			}
		} else {
			// READ
			if s.cnt == 8 || s.datadir {
				s.data = s.dev.ReadData()
				s.cnt = 0
			} else {
				s.data <<= 1
				s.cnt++
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
	val |= (s.data >> 7)
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
	rtc.regStatus1 = 0x2 // 24h-mode
	rtc.regStatus2 = 0x0
	rtc.HwSerial3W.dev = rtc
	hwio.MustInitRegs(&rtc.HwSerial3W)
	return rtc
}

func (rtc *HwRtc) ReadData() uint8 {
	if rtc.writing {
		log.Warnf("[rtc] read during register writing")
		return 0
	}

	if rtc.idx >= len(rtc.buf) {
		log.Warnf("[rtc] read but not data setup")
		return 0
	}

	data := rtc.buf[rtc.idx]
	rtc.idx++
	return data
}

func (rtc *HwRtc) bcd(value uint) uint8 {
	if value > 99 {
		log.Warnf("[rtc] cannot convert value %d to BCD", value)
		return 0xFF
	}

	value = (value/10)*16 + (value % 10)
	return uint8(value)
}

func (rtc *HwRtc) writeReg(val uint8) {
	reglen := [8]int{1, 1, 7, 3, 1, 3, 1, 1}

	rtc.buf = append(rtc.buf, val)
	if len(rtc.buf) != reglen[rtc.idx] {
		log.Warnf("[rtc] partial writing reg %d: %02x", rtc.idx, val)
		return
	}
	rtc.writing = false
	log.Warnf("[rtc] final writing reg %d: %02x", rtc.idx, val)

	switch rtc.idx {
	case 4:
		rtc.regStatus2 = val
		log.Infof("[rtc] write status register #2: %02x", val)
	default:
		log.Warnf("[rtc] unimplemented register write: %d=%x", rtc.idx, rtc.buf)
	}
}

func (rtc *HwRtc) WriteData(val uint8) {
	if rtc.writing {
		rtc.writeReg(val)
		return
	}

	if val&0xF != 6 {
		log.Warnf("[rtc] invalid command %02x", val)
		return
	}

	read := val&0x80 != 0
	reg := (val >> 4) & 7

	if !read {
		log.Warnf("[rtc] begin writing reg %d", reg)
		rtc.writing = true
		rtc.buf = nil
		rtc.idx = int(reg)
		return
	}

	rtc.buf = nil
	rtc.idx = 0
	switch reg {
	case 0:
		log.Info("[rtc] read status register #1", rtc.regStatus1)
		rtc.buf = append(rtc.buf, rtc.regStatus1)
	case 2:
		now := time.Now()
		now = time.Date(2016, 1, 2, 15, 34, 28, 0, time.UTC)
		rtc.buf = append(rtc.buf,
			rtc.bcd(uint(now.Year()-2000)),
			rtc.bcd(uint(now.Month())),
			rtc.bcd(uint(now.Day())),
			rtc.bcd(uint(now.Weekday())),
			rtc.bcd(uint(now.Hour())),
			rtc.bcd(uint(now.Minute())),
			rtc.bcd(uint(now.Second())),
		)
		log.Infof("[rtc] read datetime %x", rtc.buf)
	case 4:
		rtc.buf = append(rtc.buf, rtc.regStatus2)
		log.Infof("[rtc] read status register #2: %x", rtc.regStatus2)
	default:
		log.Warnf("[rtc] unimplemented register read: %d", reg)
	}
}
