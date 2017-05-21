package main

import (
	log "ndsemu/emu/logger"
	"ndsemu/emu/spi"
	"os"

	"github.com/edsrzf/mmap-go"
)

var modBackup = log.NewModule("backup")

// HwBackupRam implements the save ram presents in most cartridge.
// It implements the spi.Device interface
type HwBackupRam struct {
	sram     mmap.MMap
	addrSize int
	addr     int

	wbuf           []byte
	writeEnabled   bool
	autodetect     bool
	auxCntrWritten bool

	fn string
	f  *os.File
}

func NewHwBackupRam() *HwBackupRam {
	b := &HwBackupRam{
		autodetect: true,
	}
	for idx := range b.sram {
		b.sram[idx] = 0xFF
	}
	return b
}

func (b *HwBackupRam) MapSaveFile(fn string) error {
	b.fn = fn
	return nil
}

func (b *HwBackupRam) checkSize(addr int) {
	size := len(b.sram)
	if addr < size {
		return
	}

	if size == 0 {
		size = 512
	}
	for size <= addr {
		size *= 2
	}

	var err error
	if b.sram != nil {
		b.sram.Unmap()
	}
	if b.f == nil {
		b.f, err = os.OpenFile(b.fn, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		if fi, err := b.f.Stat(); err != nil {
			panic(err)
		} else if fi.Size() > int64(size) {
			size = int(fi.Size())
		}
	}
	b.f.Truncate(int64(size))
	b.sram, err = mmap.MapRegion(b.f, -1, mmap.RDWR, 0, 0)
	if err != nil {
		panic(err)
	}
}

func (b *HwBackupRam) tryAutoDetect(data []byte) bool {
	if len(data) == 1 {
		b.auxCntrWritten = false
		return false
	}

	if b.auxCntrWritten {
		b.addrSize = len(data) - 2
		modBackup.WarnZ("autodetect addr size").Int("size", b.addrSize).End()
		b.autodetect = false
		return true
	}
	modBackup.InfoZ("autodetect failed, waiting").End()
	return false
}

func (b *HwBackupRam) SpiTransfer(data []byte) ([]byte, spi.ReqStatus) {

	switch data[0] {
	case 0x0:
		// Dummy command that is sometimes sent. Ignore it
		return nil, spi.ReqFinish

	case 0x1: // WRSR
		if len(data) < 2 {
			return nil, spi.ReqContinue
		}
		modBackup.InfoZ("cmd WRSR").Hex8("val", data[1]).End()
		return nil, spi.ReqFinish

	case 0x5: // RDSR
		modBackup.InfoZ("cmd RDSR").End()
		var sr uint8
		if b.writeEnabled {
			sr |= 2
		}
		return []byte{sr}, spi.ReqFinish

	case 0x4: // WRDI
		modBackup.InfoZ("cmd WRDI").End()
		b.writeEnabled = false
		return nil, spi.ReqFinish

	case 0x6: // WREN
		modBackup.InfoZ("cmd WREN").End()
		b.writeEnabled = true
		return nil, spi.ReqFinish

	case 0x3, 0xB: // RD
		if b.autodetect && !b.tryAutoDetect(data) {
			return nil, spi.ReqContinue
		}

		if len(data) < 1+b.addrSize {
			return nil, spi.ReqContinue
		}

		if len(data) == 1+b.addrSize {
			b.addr = 0
			for _, v := range data[1:] {
				b.addr <<= 8
				b.addr |= int(v)
			}
			if b.addrSize == 1 && data[0] == 0xB {
				// For 0.5k EEPROMS, cmd 0xB means "read high"
				b.addr += 0x100
			}
		}

		b.checkSize(b.addr)
		modBackup.InfoZ("cmd RD").Int("addr", b.addr).End()
		buf := make([]byte, 256)
		sz := len(b.sram) - b.addr
		if sz > 256 {
			sz = 256
		}
		copy(buf[:sz], b.sram[b.addr:b.addr+sz])
		return buf, spi.ReqContinue

	case 0x2, 0xA: // WR
		if !b.writeEnabled {
			modBackup.FatalZ("writing with write disabled").End()
		}
		if b.autodetect {
			modBackup.FatalZ("writing while autodetecting size").End()
		}

		if len(data) < 1+b.addrSize {
			return nil, spi.ReqContinue
		}

		if len(data) == 1+b.addrSize {
			b.addr = 0
			for _, v := range data[1:] {
				b.addr <<= 8
				b.addr |= int(v)
			}
			if b.addrSize == 1 && data[0] == 0xA {
				// For 0.5k EEPROMS, cmd 0xB means "read high"
				b.addr += 0x100
			}
			modBackup.InfoZ("cmd WR").Int("addr", b.addr).End()
		}

		// Copy the whole buffer every time; I know it's inefficient,
		// but never mind...
		b.checkSize(b.addr)
		copy(b.sram[b.addr:], data[1+b.addrSize:])
		return nil, spi.ReqContinue

	default:
		modBackup.ErrorZ("unimplemented command").Blob("data", data).Int("len", len(data)).End()
		if len(data) == 16 {
			modBackup.FatalZ("data too long").End()
		}
		return nil, spi.ReqContinue
	}
}

// This is a callback to inform that the AUXSPI control register has been written.
// It is normally not required on normal execution because it is already processed
// by gamecard to control the spi bus (that eventually calls HwBackupRam as spi.Device);
// but it's very useful when we are doing auto-detection of backup RAM size.
func (b *HwBackupRam) AuxSpiCntWritten(value uint16) {
	b.auxCntrWritten = true
}

func (b *HwBackupRam) SpiBegin() {
	modBackup.InfoZ("begin transfer").End()
}

func (b *HwBackupRam) SpiEnd() {
	modBackup.InfoZ("end transfer").End()
	b.sram.Flush()
}
