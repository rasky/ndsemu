package main

import (
	"encoding/binary"
	"io"
	"ndsemu/emu"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"ndsemu/emu/spi"
	"os"

	"golang.org/x/exp/mmap"
)

var modGamecard = log.NewModule("gamecard")

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

	AuxSpiCnt  hwio.Reg16 `hwio:"bank=0,offset=0x0,rwmask=0xF07F,wcb"`
	AuxSpiData hwio.Reg16 `hwio:"bank=0,offset=0x2,wcb"`
	RomCtrl    hwio.Reg32 `hwio:"bank=0,offset=0x4,rwmask=0xFF7FFFFF,wcb"`
	GcCommand  hwio.Reg64 `hwio:"bank=0,offset=0x8,wcb"`
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
	secAreaOff int

	spi spi.Bus
	bkp *HwBackupRam
}

type noCartridgeReader struct{}

func (nc noCartridgeReader) ReadAt(buf []byte, off int64) (n int, err error) {
	for i := range buf {
		buf[i] = 0xFF
	}
	return len(buf), nil
}

func NewGamecard(biosfn string, bkp *HwBackupRam) *Gamecard {
	gc := &Gamecard{
		key2: NewKey2(),
	}
	hwio.MustInitRegs(gc)
	gc.RomCtrl.WriteCb = gc.WriteROMCTRL
	gc.CardData.ReadCb = gc.ReadCARDDATA

	// Configure spi bus
	gc.spi.SpiBusName = "SpiAux"
	gc.spi.AddDevice(0, bkp)
	gc.bkp = bkp

	f, err := os.Open(biosfn)
	if err != nil {
		panic(err)
	}
	f.ReadAt(gc.key1Tables[:], 0x30)
	f.Close()

	gc.chipid[0] = 0xFF
	gc.chipid[1] = 0xFF
	gc.chipid[2] = 0xFF
	gc.chipid[3] = 0xFF
	gc.ReaderAt = noCartridgeReader{}

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
	f, err := mmap.Open(fn)
	if err != nil {
		return err
	}

	gc.Size = uint64(f.Len())

	gc.MapCart(f)
	gc.closecb = func() { f.Close() }

	// Inititalize chip id
	gc.chipid[0] = 0xC2 // manufacturer (?)
	gc.chipid[1] = 0x7F // ROM size (Mbytes - 1)
	gc.chipid[2] = 0x00 // flags
	gc.chipid[3] = 0x80 // flags

	return nil
}

func (gc *Gamecard) WriteAUXSPICNT(old, value uint16) {
	modGamecard.InfoZ("Write AUXSPICNT").Hex16("value", value).End()
	if (old^value)&(1<<13) != 0 {
		if value&(1<<13) != 0 {
			modGamecard.InfoZ("change AUXSPI: SPI-backup").End()
			gc.spi.BeginTransfer(0)
		} else {
			modGamecard.InfoZ("change AUXSPI: ROM").End()
		}
	}
	gc.bkp.AuxSpiCntWritten(value)
}

func (gc *Gamecard) WriteAUXSPIDATA(_, value uint16) {
	modGamecard.InfoZ("Write AUXSPIDATA").Hex16("value", value).End()

	if gc.AuxSpiCnt.Value&(1<<13) == 0 {
		modGamecard.WarnZ("AUXSPIDATA written, but SPI not selected").End()
		return
	}

	if gc.AuxSpiCnt.Value&3 != 0 {
		modGamecard.WarnZ("AUXSPIDATA written, but wrong SPI frequency").End()
		return
	}

	// Do the SPI transfer. Send one byte, get one byte.
	log.ModSpi.DebugZ("xfer send").Hex8("val", uint8(value)).End()
	read := gc.spi.Transfer(uint8(value))
	gc.AuxSpiData.Value = uint16(read)
	log.ModSpi.DebugZ("xfer recv").Hex8("val", read).End()

	// If chispselect is off, this is the last trasnfer byte,
	// so reset the write buffer to discard current command and restart
	// new one
	if gc.AuxSpiCnt.Value&(1<<6) == 0 {
		gc.spi.EndTransfer()
	}
}

func (gc *Gamecard) WriteROMCTRL(_, value uint32) {
	modGamecard.InfoZ("Write ROMCTL").
		Hex32("val", value).
		Hex32("lr", uint32(nds7.Cpu.Regs[14])).
		Hex64("cmd", emu.Swap64(gc.GcCommand.Value)).
		Bool("irq", gc.AuxSpiCnt.Value&(1<<14) != 0).
		End()

	if gc.RomCtrl.Value&(1<<15) != 0 {
		s0 := uint64(gc.KeySeed0L.Value) | uint64(gc.KeySeed0H.Value)<<32
		s1 := uint64(gc.KeySeed1L.Value) | uint64(gc.KeySeed1H.Value)<<32
		gc.key2 = NewKey2WithSeed(s0, s1)
		modGamecard.InfoZ("Apply KEY2 encryption seeds").
			Hex64("s0", s0).
			Hex64("s1", s1).
			End()

		if false { // DEBUG
			var gamecode [4]byte
			gc.ReadAt(gamecode[:], 0x0C)

			var enccmd, cmd [8]byte
			binary.LittleEndian.PutUint64(enccmd[:], gc.GcCommand.Value)
			key1 := NewKey1(gc.key1Tables[:], gamecode[:], false)
			key1.DecryptBE(cmd[:], enccmd[:])
			modGamecard.InfoZ("key1 cmd decription TEST TEST").
				Blob("enc", enccmd[:]).
				Blob("dex", cmd[:]).
				End()
		}
	}
	if gc.RomCtrl.Value&(1<<13) != 0 {
		modGamecard.InfoZ("Turn on KEY2 encryption for Data").End()
	}
	if gc.RomCtrl.Value&(1<<22) != 0 {
		modGamecard.InfoZ("Turn on KEY2 encryption for Cmd").End()
	}

	if gc.RomCtrl.Value&(1<<31) != 0 {
		size := (gc.RomCtrl.Value >> 24) & 7
		if size == 7 {
			size = 4
		} else if size > 0 {
			size = 0x100 << size
		}
		modGamecard.InfoZ("ROM block transfer").
			Uint32("size", size).
			Hex8("command", uint8(gc.GcCommand.Value&0xFF)).
			End()

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
			modGamecard.FatalZ("status not implemented").Int("stat", int(gc.stat)).End()
		}

		gc.buf = buf
		gc.xferByte(true)
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
		modGamecard.FatalZ("unknown raw command").Hex8("cmd", cmd[0]).End()
	}
	return buf
}

func (gc *Gamecard) cmdKey1(size uint32) []byte {
	var gamecode [4]byte
	gc.ReadAt(gamecode[:], 0x0C)

	var enccmd, cmd [8]byte
	binary.LittleEndian.PutUint64(enccmd[:], gc.GcCommand.Value)
	key1 := NewKey1(gc.key1Tables[:], gamecode[:], false)
	key1.DecryptBE(cmd[:], enccmd[:])
	modGamecard.InfoZ("key1 cmd decription").
		Blob("enc", enccmd[:]).
		Blob("dec", cmd[:]).
		End()

	switch cmd[0] >> 4 {
	case 0x4:
		modGamecard.InfoZ("cmd: turn on KEY2").End()
		buf := make([]byte, 0x910)
		for i := 0; i < 0x910; i++ {
			buf[i] = 0xFF
		}
		return nil

	case 0x1:
		modGamecard.InfoZ("cmd: read ROM ID 2").End()
		buf := make([]byte, 4)
		copy(buf, gc.chipid[:])
		return buf

	case 0x2:
		off := int(cmd[0]&0xF)<<12 | int(cmd[1])<<4 | int(cmd[2])>>4
		off *= 0x1000

		// This command is issued 8 times with the same offset, to load 8
		// different secure area blocks.
		if gc.secAreaOff == 0 {
			gc.secAreaOff = off
		} else if !(gc.secAreaOff >= off && gc.secAreaOff < off+0x1000) {
			modGamecard.ErrorZ("invalid secure area loading: we didn't get 8 repetitions").End()
			emu.DebugBreak("invalid secure area loading")
		}

		buf := make([]byte, 512)
		gc.ReadAt(buf, int64(gc.secAreaOff))
		modGamecard.InfoZ("cmd: get secure area block").Hex32("offset", uint32(gc.secAreaOff)).End()

		// Set encryption area ID, that is not present in unencrypted ROMs
		if gc.secAreaOff == 0x4000 {
			copy(buf[0:8], []byte("encryObj"))
		}
		// Apply encryption of secure area
		if gc.secAreaOff < 0x4800 {
			keyl3 := NewKey1(gc.key1Tables[:], gamecode[:], true)
			for i := 0; i < len(buf); i += 8 {
				keyl3.EncryptLE(buf[i:i+8], buf[i:i+8])
			}
		}
		// Secure area ID (first 8 bytes) has two layers of encryption
		if gc.secAreaOff == 0x4000 {
			key1.EncryptLE(buf[0:8], buf[0:8])
		}

		gc.secAreaOff += 0x200
		if gc.secAreaOff == off+0x1000 {
			// This was the last repetition, switch back to normal key1 mode
			gc.stat = gcStatusKey1A
			gc.secAreaOff = 0
		} else {
			// we still need to wait for reptitions, stay in key1b mode
			gc.stat = gcStatusKey1B
		}

		return buf

	case 0xA:
		modGamecard.InfoZ("cmd: switch to KEY2 status").End()
		gc.stat = gcStatusKey2
		buf := make([]byte, 0x910)
		for i := 0; i < 0x910; i++ {
			buf[i] = 0xFF
		}
		return nil

	default:
		modGamecard.FatalZ("unknown key1 decrypted command").Hex8("cmd", cmd[0]).End()
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

		// Access at secure area and lower is forbidden in key2 mode
		if off < 0x8000 {
			off = 0x8000 + off&0x1FF
		}

		gc.ReadAt(buf, off)

		// Apply key2 encryption
		// gc.key2.Encrypt(buf, buf)

		modGamecard.InfoZ("encrypted load").
			Hex32("offset", uint32(off)).
			Blob("enc", cmd[:]).
			End()
		return buf

	case 0x94, 0xD6:
		// FIXME: NAND support
		copy(buf[:], gc.chipid[:])
		return buf

	case 0xB8:
		copy(buf, gc.chipid[:])

		// Apply key2 encryption
		// gc.key2.Encrypt(buf, buf)
		return buf

	default:
		modGamecard.ErrorZ("unknown key2 command").Blob("cmd", cmd[:]).End()
		return nil
	}
}

func (gc *Gamecard) xferByte(first bool) {
	if len(gc.buf) == 0 {
		modGamecard.InfoZ("end of transfer").End()
		gc.RomCtrl.Value &^= (1 << 31)
		gc.RomCtrl.Value &^= (1 << 23)
		if gc.AuxSpiCnt.Value&(1<<14) != 0 {
			gc.Irq.Raise(IrqGameCardData)
		}
		return
	}

	// Compute correct delay for data transfer.
	// This is required by at least Tongari Boushi, which
	// is very picky about the correct delay.
	clkrate := int64(5)
	if gc.RomCtrl.Value&(1<<27) != 0 {
		clkrate = int64(8)
	}
	if gc.stat == gcStatusKey2 {
		clkrate *= int64((gc.RomCtrl.Value>>16)&0x3F) + 4
	} else {
		clkrate *= int64(gc.RomCtrl.Value&0x1FFF) + 4 + 4
	}

	cycles := Emu.Sync.Cycles()
	Emu.Sync.ScheduleEvent(cycles+clkrate, func() {
		data := binary.LittleEndian.Uint32(gc.buf[0:4])
		gc.buf = gc.buf[4:]
		gc.CardData.Value = data

		gc.RomCtrl.Value |= (1 << 23) // signal data available
		nds9.TriggerDmaEvent(DmaEventGamecard)
		nds7.TriggerDmaEvent(DmaEventGamecard)
	})
}

func (gc *Gamecard) WriteGCCOMMAND(_, val uint64) {
	modGamecard.InfoZ("Write COMMAND").Hex64("val", val).End()
}

func (gc *Gamecard) ReadCARDDATA(_ uint32) uint32 {
	if gc.RomCtrl.Value&(1<<23) == 0 {
		modGamecard.WarnZ("read without pending data").End()
	} else {
		gc.RomCtrl.Value &^= (1 << 23)
		gc.xferByte(false)
	}

	return gc.CardData.Value
}
