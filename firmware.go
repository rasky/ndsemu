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
		log.WithFields(log.Fields{
			"addr": fmt.Sprintf("%06x", addr),
			"pc7":  fmt.Sprintf("%v", nds7.Cpu.GetPC()),
			"pc9":  fmt.Sprintf("%v", nds9.Cpu.GetPC()),
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
			// log.Infof("[firmware] reading: %02x", data)
			ch <- data
		}
	case FFCodeRdsr:
		status := uint8(0)
		if ff.wen {
			status |= 2
		}
		recv(status)
	default:
		log.Errorf("[firmware] unsupported command %02x", cmd)
		close(ch)
	}
}

func (ff *HwFirmwareFlash) BeginTransfer() chan uint8 {
	ch := make(chan uint8)
	go ff.transfer(ch)
	return ch
}
