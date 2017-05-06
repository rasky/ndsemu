package main

import (
	log "ndsemu/emu/logger"
	"ndsemu/emu/spi"
)

var modBackup = log.NewModule("backup")

// HwBackupRam implements the save ram presents in most cartridge.
// It implements the spi.Device interface
type HwBackupRam struct {
	sram     [4 * 1024 * 1024]byte
	addrSize int
	addr     int

	wbuf           []byte
	writeEnabled   bool
	autodetect     bool
	auxCntrWritten bool
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

	case 0x5: // RDSR
		modBackup.InfoZ("cmd RDSR").End()
		return nil, spi.ReqFinish

	case 0x4: // WRDI
		modBackup.InfoZ("cmd WRDI").End()
		b.writeEnabled = false
		return nil, spi.ReqFinish

	case 0x6: // WREN
		modBackup.InfoZ("cmd WREN").End()
		b.writeEnabled = true
		return nil, spi.ReqFinish

	case 0x3: // RD
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
		}

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
			modBackup.InfoZ("cmd WR").Int("addr", b.addr).End()
		}

		// Copy the whole buffer every time; I know it's inefficient,
		// but never mind...
		copy(b.sram[b.addr:], data[1+b.addrSize:])
		return nil, spi.ReqContinue

	// case 0x2: // WR
	// 	if len(gc.spiwbuf) >= 3 {
	// 		addr := int(gc.spiwbuf[1])<<8 + int(gc.spiwbuf[2])
	// 		if len(gc.spiwbuf) == 3 {
	// 			modGamecard.WithField("addr", addr).Info("SPI write backup")
	// 		}
	// 		addr += len(gc.spiwbuf) - 3
	// 		gc.backupSram[addr] = uint8(value)
	// 	}

	// case 0x2: // WRLO
	// 	if len(gc.spiwbuf) >= 2 {
	// 		addr := int(gc.spiwbuf[1])
	// 		if len(gc.spiwbuf) == 2 {
	// 			modGamecard.WithField("addr", addr).Info("SPI write backup 0.5k LO")
	// 		}
	// 		addr += len(gc.spiwbuf) - 2
	// 		gc.backupSram[addr] = uint8(value)
	// 	}

	// case 0xA: // WRHI
	// 	if len(gc.spiwbuf) >= 2 {
	// 		addr := int(gc.spiwbuf[1]) + 0x100
	// 		if len(gc.spiwbuf) == 2 {
	// 			modGamecard.WithField("addr", addr).Info("SPI write backup 0.5k HI")
	// 		}
	// 		addr += len(gc.spiwbuf) - 2
	// 		gc.backupSram[addr] = uint8(value)
	// 	}

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
}
