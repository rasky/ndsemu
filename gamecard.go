package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"ndsemu/emu/hwio"
	"os"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type gcStatus int

const (
	gcStatusRaw gcStatus = iota
	gcStatusKey1A
	gcStatusKey1B
	gcStatusKey2
)

type Gamecard struct {
	io.ReaderAt
	Irq     *HwIrq
	closecb func()
	Size    uint64

	AuxSpiCnt  hwio.Reg16 `hwio:"bank=0,offset=0x0,rwmask=0xF07F"`
	AuxSpiData hwio.Reg16 `hwio:"bank=0,offset=0x2,rcb,wcb"`
	RomCtrl    hwio.Reg32 `hwio:"bank=0,offset=0x4,rwmask=0xFF7FFFFF,wcb"`
	GcCommand  hwio.Reg64 `hwio:"bank=0,offset=0x8,writeonly"`

	CardData hwio.Reg32 `hwio:"bank=1,offset=0x0,readonly,rcb"`

	stat       gcStatus
	buf        []byte
	key1Tables [(18 + 1024) * 4]byte
	key2       Key2
}

func NewGamecard(irq *HwIrq, biosfn string) *Gamecard {
	gc := &Gamecard{
		Irq:  irq,
		key2: NewKey2(),
	}
	hwio.MustInitRegs(gc)

	f, err := os.Open(biosfn)
	if err != nil {
		panic(err)
	}
	f.ReadAt(gc.key1Tables[:], 0x30)
	f.Close()

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

func (gc *Gamecard) WriteAUXSPIDATA(_, value uint16) {
	// Emu.DebugBreak("[cartidge] Write AUXSPIDATA")
}

func (gc *Gamecard) ReadAUXSPIDATA(_ uint16) uint16 {
	// Emu.DebugBreak("[cartidge] Read AUXSPIDATA")
	return 0
}

func (gc *Gamecard) WriteROMCTRL(_, value uint32) {
	log.WithFields(log.Fields{
		"val": fmt.Sprintf("%08x", value),
		"pc7": nds7.Cpu.GetPC(),
		"lr":  nds7.Cpu.Regs[14],
	}).Info("[cartidge] Write ROMCTL")

	if gc.RomCtrl.Value&(1<<15) != 0 {
		log.Infof("[gamecard] Apply KEY2 encryption seeds")
	}
	if gc.RomCtrl.Value&(1<<13) != 0 {
		log.Infof("[gamecard] Turn on KEY2 encryption for Data")
	}
	if gc.RomCtrl.Value&(1<<22) != 0 {
		log.Infof("[gamecard] Turn on KEY2 encryption for Cmd")
	}

	if gc.RomCtrl.Value&(1<<31) != 0 {
		size := (gc.RomCtrl.Value >> 24) & 7
		if size == 7 {
			size = 4
		} else if size > 0 {
			size = 0x100 << size
		}
		log.Infof("[gamecard] ROM block transfer: size: %d, command: %x", size, (gc.GcCommand.Value & 0xFF))

		var buf []byte
		switch gc.stat {
		case gcStatusRaw:
			buf = gc.cmdRaw(size)
		case gcStatusKey1A:
			// we do nothing here and wait for the command to be reissued
			gc.stat = gcStatusKey1B
		case gcStatusKey1B:
			buf = gc.cmdKey1(size)
			gc.stat = gcStatusKey1A
		default:
			log.Fatal("[gamecard] status key2 not implemented")
		}

		gc.buf = buf
		if len(gc.buf) > 0 {
			// Signal data ready
			gc.RomCtrl.Value |= (1 << 23)
		} else {
			gc.endOfTransfer()
		}
	}
}

func (gc *Gamecard) cmdRaw(size uint32) []byte {
	buf := make([]byte, size)

	var cmd [8]byte
	binary.LittleEndian.PutUint64(cmd[:], gc.GcCommand.Value)

	switch cmd[0] {
	case 0x9F:
		// Dummy command: read 0xFF
		for i := range buf {
			buf[i] = 0xFF
		}

	case 0x00:
		// Read header
		gc.ReadAt(buf, 0)

	case 0x90, 0xB8:
		// Get ROM chip ID
		buf[0] = 0xC2 // manufacturer (?)
		buf[1] = 0x7F // ROM size (Mbytes - 1)
		buf[2] = 0x00 // flags
		buf[3] = 0x00 // flags

		if cmd[0] == 0xB7 { // encrypted
			gc.key2.Encrypt(buf, buf)
		}

	case 0x3C:
		// Activate KEY1
		gc.stat = gcStatusKey1A
		for i := range buf {
			buf[i] = 0xFF
		}

	case 0xB7:
		// Encrypted load
		off := int64(binary.BigEndian.Uint32(cmd[1:5]))
		gc.ReadAt(buf, off)

		// Apply key2 encryption
		gc.key2.Encrypt(buf, buf)

	default:
		log.Fatalf("[gamecard] unknown raw command: %x", cmd[0])
	}
	return buf
}

func (gc *Gamecard) cmdKey1(size uint32) []byte {
	var gamecode [4]byte
	gc.ReadAt(gamecode[:], 0x0C)

	var enccmd, cmd [8]byte
	binary.LittleEndian.PutUint64(enccmd[:], gc.GcCommand.Value)
	key1 := NewKey1(gc.key1Tables[:], gamecode[:])
	key1.DecryptBE(cmd[:], enccmd[:])
	log.WithFields(log.Fields{
		"enc": fmt.Sprintf("%x", enccmd),
		"dec": fmt.Sprintf("%x", cmd),
	}).Infof("[gamecard] key1 cmd decription")

	switch cmd[0] >> 4 {
	case 0x4:
		log.Infof("[gamecard] cmd: turn on KEY2")
		buf := make([]byte, 0x910+4)
		for i := 0; i < 0x910; i++ {
			buf[i] = 0xFF
		}
		return nil

	case 0xA:
		log.Infof("[gamecard] cmd: switch to KEY2 status")
		gc.stat = gcStatusKey2
		buf := make([]byte, 0x910)
		for i := 0; i < 0x910; i++ {
			buf[i] = 0xFF
		}
		return nil

	default:
		log.Fatalf("[gamecard] unknown key1 decrypted command: %x", cmd[0])
		return nil
	}
}

func (gc *Gamecard) endOfTransfer() {
	log.Info("[gamecard] end of transfer")
	gc.RomCtrl.Value &^= (1 << 31)
	gc.RomCtrl.Value &^= (1 << 23)
	if gc.AuxSpiCnt.Value&(1<<14) != 0 {
		gc.Irq.Raise(IrqGameCardData)
	}
}

func (gc *Gamecard) ReadCARDDATA(_ uint32) uint32 {
	if len(gc.buf) == 0 {
		log.WithField("pc7", nds7.Cpu.GetPC()).Warn("[gamecard] read DATA but not pending data")
		return 0
	}
	data := binary.BigEndian.Uint32(gc.buf[0:4])
	gc.buf = gc.buf[4:]
	if len(gc.buf) == 0 {
		// End of data
		gc.endOfTransfer()
	}
	return data
}
