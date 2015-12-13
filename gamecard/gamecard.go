package gamecard

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type Gamecard struct {
	io.ReaderAt
	closecb   func()
	regSpiCnt uint16
	regRomCtl uint32
	Size      uint64

	cmd [8]byte
	buf []byte
}

func NewGamecard() *Gamecard {
	gc := &Gamecard{
		regSpiCnt: 0x0,
	}
	return gc
}

func (gc *Gamecard) MapCart(data io.ReaderAt) {
	if gc.closecb != nil {
		gc.closecb()
		gc.closecb = nil
	}
	gc.ReaderAt = data
}

func (gc *Gamecard) MapCartFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	size, err := f.Seek(0, 2)
	if err != nil {
		return err
	}
	f.Seek(0, 0)
	gc.Size = uint64(size)

	gc.MapCart(f)
	gc.closecb = func() { f.Close() }
	return nil
}

func (gc *Gamecard) WriteAUXSPICNT(value uint16) {
	gc.regSpiCnt = value
	log.WithField("val", fmt.Sprintf("%04x", value)).Info("[cartidge] Write AUXSPICNT")
}

func (gc *Gamecard) ReadAUXSPICNT() uint16 {
	log.WithField("val", fmt.Sprintf("%04x", gc.regSpiCnt)).Info("[cartidge] Read AUXSPICNT")
	return gc.regSpiCnt
}

func (gc *Gamecard) WriteAUXSPIDATA(value uint16) {
	log.WithField("val", fmt.Sprintf("%04x", value)).Info("[cartidge] Write AUXSPIDATA")
}

func (gc *Gamecard) ReadAUXSPIDATA() uint16 {
	log.WithField("val", fmt.Sprintf("%04x", 0)).Info("[cartidge] Read AUXSPIDATA")
	return 0
}

func (gc *Gamecard) WriteROMCTL(value uint32) {
	gc.regRomCtl = value &^ (1 << 23)
	log.WithField("val", fmt.Sprintf("%08x", value)).Info("[cartidge] Write ROMCTL")
	if gc.regRomCtl&(1<<31) != 0 {
		size := (gc.regRomCtl >> 24) & 7
		if size == 7 {
			size = 4
		} else if size > 0 {
			size = 0x100 << size
		}
		log.Infof("[gamecard] ROM block transfer: size: %d, command: %x", size, gc.cmd)

		gc.buf = make([]byte, size)

		switch gc.cmd[0] {
		case 0x9F:
			// Dummy command: read 0xFF
			for i := range gc.buf {
				gc.buf[i] = 0xFF
			}

		case 0x00:
			// Read header
			gc.ReadAt(gc.buf, 0)

		case 0x90:
			// Get ROM chip ID
			gc.buf[0] = 0xC2 // manufacturer (?)
			gc.buf[1] = 0x7F // ROM size (Mbytes - 1)
			gc.buf[2] = 0x00 // flags
			gc.buf[3] = 0x00 // flags

		case 0x3C:
			// Activate KEY1
			for i := range gc.buf {
				gc.buf[i] = 0xFF
			}

		default:
			log.Fatalf("[gamecard] unknown command: %x", gc.cmd[0])
		}

		if size > 0 {
			// Signal data ready
			gc.regRomCtl |= (1 << 23)
		} else {
			// end of transfer
			gc.regRomCtl &^= (1 << 31)
		}
	}
}

func (gc *Gamecard) ReadROMCTL() uint32 {
	// log.WithField("val", fmt.Sprintf("%08x", gc.regRomCtl)).Info("[cartidge] Read ROMCTL")
	return gc.regRomCtl
}

func (gc *Gamecard) WriteCommand(addr uint32, value uint8) {
	gc.cmd[addr&7] = value
}

func (gc *Gamecard) ReadData() uint32 {
	if len(gc.buf) == 0 {
		return 0
	}
	data := binary.LittleEndian.Uint32(gc.buf[0:4])
	gc.buf = gc.buf[4:]
	if len(gc.buf) == 0 {
		// End of data
		log.Info("[gamecard] end of transfer")
		gc.regRomCtl &^= (1 << 31)
	}
	return data
}
