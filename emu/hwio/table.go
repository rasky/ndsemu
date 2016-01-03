package hwio

import (
	"fmt"
	"ndsemu/emu"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type BankIO8 interface {
	Read8(addr uint32) uint8
	Write8(addr uint32, val uint8)
}

type BankIO16 interface {
	Read16(addr uint32) uint16
	Write16(addr uint32, val uint16)
}

type BankIO32 interface {
	Read32(addr uint32) uint32
	Write32(addr uint32, val uint32)
}

type BankIO interface {
	BankIO8
	BankIO16
	BankIO32
}

type Table struct {
	Name string

	dense8 [0x10000]uint8
	table8 []BankIO8

	dense16 [0x10000 / 2]uint8
	table16 []BankIO16

	dense32 [0x10000 / 4]uint8
	table32 []BankIO32
}

type io32to16 Table

func (t *io32to16) Read32(addr uint32) uint32 {
	val1 := (*Table)(t).Read16(addr + 0)
	val2 := (*Table)(t).Read16(addr + 2)
	return uint32(val1) | uint32(val2)<<16
}

func (t *io32to16) Write32(addr uint32, val uint32) {
	(*Table)(t).Write16(addr+0, uint16(val&0xFFFF))
	(*Table)(t).Write16(addr+2, uint16(val>>16))
}

type io16to8 Table

func (t *io16to8) Read16(addr uint32) uint16 {
	val1 := (*Table)(t).Read8(addr + 0)
	val2 := (*Table)(t).Read8(addr + 1)
	return uint16(val1) | uint16(val2)<<8
}

func (t *io16to8) Write16(addr uint32, val uint16) {
	(*Table)(t).Write8(addr+0, uint8(val&0xFF))
	(*Table)(t).Write8(addr+1, uint8(val>>8))
}

func (t *Table) Reset() {
	for i := range t.dense8 {
		t.dense8[i] = 0
	}
	for i := range t.dense16 {
		t.dense16[i] = 0
	}
	for i := range t.dense32 {
		t.dense32[i] = 0
	}
	t.table8 = []BankIO8{nil}
	t.table16 = []BankIO16{nil}
	t.table32 = []BankIO32{nil}
}

// Map a register bank (that is, a structure containing mulitple IoReg* fields).
// For this function to work, registers must have a struct tag "hwio", containing
// the following fields:
//
//      offset=0x12     Byte-offset within the register bank at which this
//                      register is mapped. There is no default value: if this
//                      option is missing, the register is assumed not to be
//                      part of the bank, and is ignored by this call.
//
//      bank=NN         Ordinal bank number (if not specified, default to zero).
//                      This option allows for a structure to expose multiple
//                      banks, as regs can be grouped by bank by specified the
//                      bank number.
//
func (t *Table) MapBank(addr uint32, bank interface{}, bankNum int) {
	regs, err := bankGetRegs(bank, bankNum)
	if err != nil {
		panic(err)
	}

	for _, reg := range regs {
		switch r := reg.regPtr.(type) {
		case *Mem:
			t.MapMem(addr+reg.offset, r)
		case *Reg64:
			t.MapReg64(addr+reg.offset, r)
		case *Reg32:
			t.MapReg32(addr+reg.offset, r)
		case *Reg16:
			t.MapReg16(addr+reg.offset, r)
		case *Reg8:
			t.MapReg8(addr+reg.offset, r)
		default:
			panic(fmt.Errorf("invalid reg type: %T", r))
		}
	}
}

func (t *Table) mapBus32(addr uint32, size uint32, io BankIO32, allowremap bool) {
	t.table32 = append(t.table32, io)
	idx := len(t.table32) - 1
	if idx == 0 {
		panic("table was not reset")
	}
	if idx >= 256 {
		panic("too many regs")
	}
	addr >>= 2
	for i := uint32(0); i < size>>2; i++ {
		if !allowremap && t.dense32[addr+i] != 0 {
			panic("address already mapped")
		}
		t.dense32[addr+i] = uint8(idx)
	}
}

func (t *Table) mapBus16(addr uint32, size uint32, io BankIO16, allowremap bool) {
	t.table16 = append(t.table16, io)
	idx := len(t.table16) - 1
	if idx == 0 {
		panic("table was not reset")
	}
	if idx >= 256 {
		panic("too many regs")
	}
	addr >>= 1
	for i := uint32(0); i < size>>1; i++ {
		if !allowremap && t.dense16[addr+i] != 0 {
			panic("address already mapped")
		}
		t.dense16[addr+i] = uint8(idx)
	}
}

func (t *Table) mapBus8(addr uint32, size uint32, io BankIO8, allowremap bool) {
	t.table8 = append(t.table8, io)
	idx := len(t.table8) - 1
	if idx == 0 {
		panic("table was not reset")
	}
	if idx >= 256 {
		panic("too many regs")
	}
	for i := uint32(0); i < size; i++ {
		if !allowremap && t.dense8[addr+i] != 0 {
			panic(fmt.Sprintf("address already mapped: %08x", addr+i))
		}
		t.dense8[addr+i] = uint8(idx)
	}
}

func (t *Table) MapReg64(addr uint32, io *Reg64) {
	addr &= 0xFFFF
	if addr&7 != 0 {
		panic("unaligned mapping")
	}
	t.mapBus8(addr, 8, io, false)
	t.mapBus16(addr, 8, io, false)
	t.mapBus32(addr, 8, io, false)
}

func (t *Table) MapReg32(addr uint32, io *Reg32) {
	addr &= 0xFFFF
	if addr&3 != 0 {
		panic("unaligned mapping")
	}
	t.mapBus8(addr, 4, io, false)
	t.mapBus16(addr, 4, io, false)
	t.mapBus32(addr, 4, io, false)
}

func (t *Table) MapReg16(addr uint32, io *Reg16) {
	addr &= 0xFFFF
	if addr&1 != 0 {
		panic("unaligned mapping")
	}
	t.mapBus8(addr, 2, io, false)
	t.mapBus16(addr, 2, io, false)
	t.mapBus32(addr, 4, (*io32to16)(t), true)
}

func (t *Table) MapReg8(addr uint32, io *Reg8) {
	addr &= 0xFFFF
	t.mapBus8(addr, 1, io, false)
	t.mapBus16(addr, 2, (*io16to8)(t), true)
	t.mapBus32(addr, 4, (*io32to16)(t), true)
}

func (t *Table) MapMem(addr uint32, mem *Mem) {
	if len(mem.Data)&(len(mem.Data)-1) != 0 {
		panic("memory buffer size is not pow2")
	}
	addr &= 0xFFFF
	if mem.Flags&MemFlag8 != 0 {
		t.mapBus8(addr, uint32(mem.Size()), mem8(mem.Data), false)
	}
	if mem.Flags&MemFlag16ForceAlign != 0 {
		t.mapBus16(addr, uint32(mem.Size()), mem16LittleEndianForceAlign(mem.Data), false)
	} else if mem.Flags&MemFlag16Unaligned != 0 {
		t.mapBus16(addr, uint32(mem.Size()), mem16LittleEndianUnaligned(mem.Data), false)
	} else if mem.Flags&MemFlag16Byteswapped != 0 {
		t.mapBus16(addr, uint32(mem.Size()), mem16LittleEndianByteSwap(mem.Data), false)
	}
	if mem.Flags&MemFlag32ForceAlign != 0 {
		t.mapBus32(addr, uint32(mem.Size()), mem32LittleEndianForceAlign(mem.Data), false)
	} else if mem.Flags&MemFlag32Unaligned != 0 {
		t.mapBus32(addr, uint32(mem.Size()), mem32LittleEndianUnaligned(mem.Data), false)
	} else if mem.Flags&MemFlag32Byteswapped != 0 {
		t.mapBus32(addr, uint32(mem.Size()), mem32LittleEndianByteSwap(mem.Data), false)
	}
}

func (t *Table) Read8(addr uint32) uint8 {
	io := t.table8[t.dense8[addr&0xFFFF]]
	if io == nil {
		if addr >= 0x4000400 && addr < 0x4000510 {
			return 0
		}
		log.WithFields(log.Fields{
			"name": t.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] unmapped Read8")
		return 0
	}
	return io.Read8(addr)
}

func (t *Table) Write8(addr uint32, val uint8) {
	io := t.table8[t.dense8[addr&0xFFFF]]
	if io == nil {
		log.WithFields(log.Fields{
			"name": t.Name,
			"val":  emu.Hex8(val),
			"addr": emu.Hex32(addr),
		}).Error("[IO] unmapped Write8")
		return
	}
	io.Write8(addr, val)
}

func (t *Table) Read16(addr uint32) uint16 {
	io := t.table16[t.dense16[(addr&0xFFFF)>>1]]
	if io == nil {
		log.WithFields(log.Fields{
			"name": t.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] unmapped Read16")
		return 0
	}
	return io.Read16(addr)
}

func (t *Table) Write16(addr uint32, val uint16) {
	io := t.table16[t.dense16[(addr&0xFFFF)>>1]]
	if io == nil {
		log.WithFields(log.Fields{
			"name": t.Name,
			"val":  emu.Hex16(val),
			"addr": emu.Hex32(addr),
		}).Error("[IO] unmapped Write16")
		return
	}
	io.Write16(addr, val)
}

func (t *Table) Read32(addr uint32) uint32 {
	io := t.table32[t.dense32[(addr&0xFFFF)>>2]]
	if io == nil {
		log.WithFields(log.Fields{
			"name": t.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] unmapped Read32")
		return 0
	}
	return io.Read32(addr)
}

func (t *Table) Write32(addr uint32, val uint32) {
	io := t.table32[t.dense32[(addr&0xFFFF)>>2]]
	if io == nil {
		log.WithFields(log.Fields{
			"name": t.Name,
			"val":  emu.Hex32(val),
			"addr": emu.Hex32(addr),
		}).Error("[IO] unmapped Write32")
		return
	}
	io.Write32(addr, val)
}
