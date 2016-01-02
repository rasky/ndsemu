package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"ndsemu/emu"
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
	GcCommand  hwio.Reg64 `hwio:"bank=0,offset=0x8,writeonly,wcb"`
	KeySeed0L  hwio.Reg32 `hwio:"bank=0,offset=0x10,writeonly"`
	KeySeed1L  hwio.Reg32 `hwio:"bank=0,offset=0x14,writeonly"`
	KeySeed0H  hwio.Reg16 `hwio:"bank=0,offset=0x18,rwmask=0x7f,writeonly"`
	KeySeed1H  hwio.Reg16 `hwio:"bank=0,offset=0x1A,rwmask=0x7f,writeonly"`

	CardData hwio.Reg32 `hwio:"bank=1,offset=0x0,readonly,rcb"`

	chipid     [4]byte
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

	// Inititalize chip id
	gc.chipid[0] = 0xC2 // manufacturer (?)
	gc.chipid[1] = 0x7F // ROM size (Mbytes - 1)
	gc.chipid[2] = 0x00 // flags
	gc.chipid[3] = 0x80 // flags

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
		"cmd": emu.Hex64(gc.GcCommand.Value),
		"irq": gc.AuxSpiCnt.Value&(1<<14) != 0,
	}).Info("[cartidge] Write ROMCTL")

	if gc.RomCtrl.Value&(1<<15) != 0 {
		s0 := uint64(gc.KeySeed0L.Value) | uint64(gc.KeySeed0H.Value)<<32
		s1 := uint64(gc.KeySeed1L.Value) | uint64(gc.KeySeed1H.Value)<<32
		gc.key2 = NewKey2WithSeed(s0, s1)
		log.WithFields(log.Fields{
			"s0": emu.Hex64(s0),
			"s1": emu.Hex64(s1),
		}).Infof("[gamecard] Apply KEY2 encryption seeds")

		if true {
			var gamecode [4]byte
			gc.ReadAt(gamecode[:], 0x0C)

			var enccmd, cmd [8]byte
			binary.LittleEndian.PutUint64(enccmd[:], gc.GcCommand.Value)
			key1 := NewKey1(gc.key1Tables[:], gamecode[:])
			key1.DecryptBE(cmd[:], enccmd[:])
			log.WithFields(log.Fields{
				"enc": fmt.Sprintf("%x", enccmd),
				"dec": fmt.Sprintf("%x", cmd),
			}).Infof("[gamecard] key1 cmd decription TEST TEST")
		}
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
			gc.stat = gcStatusKey1A
			buf = gc.cmdKey1(size)
		case gcStatusKey2:
			buf = gc.cmdKey2(size)
		default:
			log.Fatalf("[gamecard] status not implemented: %d", gc.stat)
		}

		gc.buf = buf
		gc.updateStatus()
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

	case 0x90:
		// Get ROM chip ID
		copy(buf[:], gc.chipid[:])

	case 0x3C:
		// Activate KEY1
		gc.stat = gcStatusKey1A
		for i := range buf {
			buf[i] = 0xFF
		}

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

	case 0x1:
		log.Infof("[gamecard] cmd: read ROM ID 2")
		buf := make([]byte, 4)
		gc.key2.Encrypt(buf, gc.chipid[:])
		return buf

	default:
		log.Fatalf("[gamecard] unknown key1 decrypted command: %x", cmd[0])
		return nil
	}
}

func (gc *Gamecard) cmdKey2(size uint32) []byte {
	var cmd [8]byte
	binary.LittleEndian.PutUint64(cmd[:], gc.GcCommand.Value)

	buf := make([]byte, size)
	switch cmd[0] {
	case 0xB7:
		// Encrypted load
		off := int64(binary.BigEndian.Uint32(cmd[1:5])) & int64(gc.Size-1)
		gc.ReadAt(buf, off)

		// Apply key2 encryption
		// gc.key2.Encrypt(buf, buf)

		log.Infof("[gamecard] encrypted load from offset %x (enc:%x)", off, cmd)
		return buf

	case 0xB8:
		copy(buf, gc.chipid[:])

		// Apply key2 encryption
		// gc.key2.Encrypt(buf, buf)
		return buf

	default:
		log.Fatalf("[gamecard] unknown key2 command: %x", cmd)
		return nil
	}
}

func (gc *Gamecard) updateStatus() {
	if len(gc.buf) == 0 {
		log.Info("[gamecard] end of transfer")
		gc.RomCtrl.Value &^= (1 << 31)
		gc.RomCtrl.Value &^= (1 << 23)
		if gc.AuxSpiCnt.Value&(1<<14) != 0 {
			gc.Irq.Raise(IrqGameCardData)
		}
	} else {
		// Signal data ready
		gc.RomCtrl.Value |= (1 << 23)
		nds9.TriggerDmaEvent(DmaEventGamecard)
		nds7.TriggerDmaEvent(DmaEventGamecard)
	}
}

func (gc *Gamecard) WriteGCCOMMAND(_, val uint64) {
	// Emu.DebugBreak("write gccommand")
	Emu.Log().Infof("[gamecard] Write COMMAND: %08x", val)
}

func (gc *Gamecard) ReadCARDDATA(_ uint32) uint32 {
	if len(gc.buf) == 0 {
		log.WithField("pc7", nds7.Cpu.GetPC()).Warn("[gamecard] read DATA but not pending data")
		return 0
	}
	data := binary.LittleEndian.Uint32(gc.buf[0:4])
	gc.buf = gc.buf[4:]
	// log.Infof("[gamecard] read DATA: %08x", data)
	gc.updateStatus()
	return data
}
