package main

import (
	"fmt"
	log "ndsemu/emu/logger"
	"os"
)

var modFw = log.NewModule("firmware")

const (
	FFCodeRead uint8 = 0x03
	FFCodeRdsr uint8 = 0x05
	FFCodeWren uint8 = 0x06
	FFCodeWrdi uint8 = 0x04
	FFCodePw   uint8 = 0x0A
)

type HwFirmwareFlash struct {
	f   *os.File
	wen bool
}

func NewHwFirmwareFlash() *HwFirmwareFlash {
	return &HwFirmwareFlash{}
}

func (ff *HwFirmwareFlash) MapFirmwareFile(fn string) error {
	f, err := os.OpenFile(fn, os.O_RDWR, 0777)
	if err != nil {
		return err
	}
	ff.f = f
	return nil
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
		modFw.WithFields(log.Fields{
			"addr": fmt.Sprintf("%06x", addr),
		}).Info("READ")
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
		modFw.WithField("val", fmt.Sprintf("%02x", status)).Info("read status")
		recv(status)
	case FFCodeWren:
		modFw.Info("write enabled")
		ff.wen = true
	case FFCodeWrdi:
		modFw.Info("write disabled")
		ff.wen = false
	case FFCodePw:
		a1, a2, a3 := recv(0), recv(0), recv(0)
		addr := uint32(a1)<<16 | uint32(a2)<<8 | uint32(a3)
		modFw.WithFields(log.Fields{
			"addr": fmt.Sprintf("%06x", addr),
		}).Info("WRITE")
		var buf []byte
		for c := range ch {
			buf = append(buf, c)
			ch <- 0
		}
		ff.f.WriteAt(buf, int64(addr))
		modFw.Infof("write finished (%d bytes)", len(buf))
	default:
		modFw.Errorf("unsupported command %02x", cmd)
	}
}

func (ff *HwFirmwareFlash) BeginTransfer() chan uint8 {
	ch := make(chan uint8)
	go ff.transfer(ch)
	return ch
}
