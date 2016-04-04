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

	wbuf []byte
	addr uint32
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

func (ff *HwFirmwareFlash) SpiBegin() {
	ff.addr = 0
	ff.wbuf = nil
}

func (ff *HwFirmwareFlash) SpiTransfer(data []byte) ([]byte, SpiStatus) {
	cmd := data[0]
	switch cmd {
	case FFCodeRead:
		if len(data) < 4 {
			return nil, SpiContinue
		}
		if len(data) == 4 {
			ff.addr = uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
			modFw.WithFields(log.Fields{
				"addr": fmt.Sprintf("%06x", ff.addr),
			}).Info("READ")
			return nil, SpiContinue
		}

		buf := make([]byte, 1024)
		ff.f.ReadAt(buf, int64(ff.addr))
		ff.addr += 1024
		return buf, SpiContinue
	case FFCodeRdsr:
		status := uint8(0)
		if ff.wen {
			status |= 2
		}
		modFw.WithField("val", fmt.Sprintf("%02x", status)).Info("read status")
		return []byte{status}, SpiFinish
	case FFCodeWren:
		modFw.Info("write enabled")
		ff.wen = true
		return nil, SpiFinish
	case FFCodeWrdi:
		modFw.Info("write disabled")
		ff.wen = false
		return nil, SpiFinish
	case FFCodePw:
		if len(data) < 4 {
			return nil, SpiContinue
		}
		if len(data) == 4 {
			ff.addr = uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
			modFw.WithFields(log.Fields{
				"addr": fmt.Sprintf("%06x", ff.addr),
			}).Info("WRITE")
			return nil, SpiContinue
		}
		// Put away buffer data; will be written just once, in SpiEnd
		ff.wbuf = data[4:]
		return nil, SpiContinue
	default:
		modFw.Errorf("unsupported command %02x", cmd)
		return nil, SpiFinish
	}
}

func (ff *HwFirmwareFlash) SpiEnd() {
	if ff.wbuf != nil {
		ff.f.WriteAt(ff.wbuf, int64(ff.addr))
		ff.wbuf = nil
	}
}
