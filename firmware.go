package main

import (
	"fmt"
	"io"
	"os"

	log "gopkg.in/Sirupsen/logrus.v0"
)

const (
	FFCodeRead uint8 = 0x03
	FFCodeRdsr uint8 = 0x05
	FFCodeWren uint8 = 0x06
	FFCodeWrdi uint8 = 0x04
	FFCodePw   uint8 = 0x0A
)

type HwFirmwareFlash struct {
	f   io.ReaderAt
	wen bool
}

func NewHwFirmwareFlash(fn string) *HwFirmwareFlash {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &HwFirmwareFlash{
		f: f,
	}
}

func (ff *HwFirmwareFlash) transfer(ch chan uint8) {
	recv := func(val uint8) uint8 {
		data := <-ch
		ch <- val
		return data
	}

	cmd := recv(0)
	switch cmd {
	case FFCodeRead:
		a1, a2, a3 := recv(0), recv(0), recv(0)
		addr := uint32(a1)<<16 | uint32(a2)<<8 | uint32(a3)
		Emu.Log().WithFields(log.Fields{
			"addr": fmt.Sprintf("%06x", addr),
		}).Info("[firmware] READ")
		var buf []byte
		for _ = range ch {
			if len(buf) == 0 {
				buf = make([]byte, 1024)
				ff.f.ReadAt(buf, int64(addr))
				addr += uint32(len(buf))
			}
			data := buf[0]
			buf = buf[1:]
			ch <- data
		}
	case FFCodeRdsr:
		status := uint8(0)
		if ff.wen {
			status |= 2
		}
		log.WithField("val", fmt.Sprintf("%02x", status)).Info("[firmware] read status")
		recv(status)
	case FFCodeWren:
		log.Info("[firmware] write enabled")
		ff.wen = true
	case FFCodeWrdi:
		log.Info("[firmware] write disabled")
		ff.wen = false
	case FFCodePw:
		a1, a2, a3 := recv(0), recv(0), recv(0)
		addr := uint32(a1)<<16 | uint32(a2)<<8 | uint32(a3)
		Emu.Log().WithFields(log.Fields{
			"addr": fmt.Sprintf("%06x", addr),
		}).Info("[firmware] WRITE")
		var buf []byte
		for c := range ch {
			log.Infof("[firmware] byte %02x", c)
			buf = append(buf, c)
			ch <- 0
		}
		Emu.Log().Infof("[firmware] write finished (%d bytes)", len(buf))
	default:
		log.Errorf("[firmware] unsupported command %02x", cmd)
	}
}

func (ff *HwFirmwareFlash) BeginTransfer() chan uint8 {
	ch := make(chan uint8)
	go ff.transfer(ch)
	return ch
}
