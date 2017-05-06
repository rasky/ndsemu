package main

import (
	log "ndsemu/emu/logger"
	"ndsemu/emu/spi"
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

func (ff *HwFirmwareFlash) SpiTransfer(data []byte) ([]byte, spi.ReqStatus) {
	cmd := data[0]
	switch cmd {
	case 0:
		// Dummy command that is sent as part of the last byte transfer
		// FIXME: we could fix this at the spibus level
		return nil, spi.ReqFinish

	case FFCodeRead:
		if len(data) < 4 {
			return nil, spi.ReqContinue
		}
		if len(data) == 4 {
			ff.addr = uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
			modFw.InfoZ("HEAD").Hex32("addr", ff.addr).End()
		}

		buf := make([]byte, 1024)
		ff.f.ReadAt(buf, int64(ff.addr))
		ff.addr += 1024
		return buf, spi.ReqContinue
	case FFCodeRdsr:
		status := uint8(0)
		if ff.wen {
			status |= 2
		}
		modFw.InfoZ("read status").Hex8("val", status).End()
		return []byte{status}, spi.ReqFinish
	case FFCodeWren:
		modFw.InfoZ("write enabled").End()
		ff.wen = true
		return nil, spi.ReqFinish
	case FFCodeWrdi:
		modFw.InfoZ("write disabled").End()
		ff.wen = false
		return nil, spi.ReqFinish
	case FFCodePw:
		if len(data) < 4 {
			return nil, spi.ReqContinue
		}
		if len(data) == 4 {
			ff.addr = uint32(data[1])<<16 | uint32(data[2])<<8 | uint32(data[3])
			modFw.InfoZ("WRITE").Hex32("addr", ff.addr).End()
		}
		// Put away buffer data; will be written just once, in SpiEnd
		ff.wbuf = data[4:]
		return nil, spi.ReqContinue
	default:
		modFw.ErrorZ("unsupported command").Hex8("cmd", cmd).End()
		return nil, spi.ReqFinish
	}
}

func (ff *HwFirmwareFlash) SpiEnd() {
	if ff.wbuf != nil {
		ff.f.WriteAt(ff.wbuf, int64(ff.addr))
		ff.wbuf = nil
	}
}
