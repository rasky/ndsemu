package gamecard

import (
	"fmt"
	"io"
	"os"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type Gamecard struct {
	card      io.ReaderAt
	closecb   func()
	regSpiCnt uint16
	regRomCtl uint32
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
	gc.card = data
}

func (gc *Gamecard) MapCartFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

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
	gc.regRomCtl = value
	log.WithField("val", fmt.Sprintf("%08x", value)).Info("[cartidge] Write ROMCTL")
}

func (gc *Gamecard) ReadROMCTL() uint32 {
	log.WithField("val", fmt.Sprintf("%08x", gc.regRomCtl)).Info("[cartidge] Read ROMCTL")
	return gc.regRomCtl
}
