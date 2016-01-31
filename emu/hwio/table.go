package hwio

import (
	"fmt"
	"ndsemu/emu"
	log "ndsemu/emu/logger"
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
	ws   int

	table8  radixTree
	table16 radixTree
	table32 radixTree
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

func NewTable(name string) *Table {
	t := new(Table)
	t.Name = name
	t.Reset()
	return t
}

func (t *Table) SetWaitStates(ws int) {
	t.ws = ws
}

func (t *Table) Reset() {
	t.table8 = radixTree{}
	t.table16 = radixTree{}
	t.table32 = radixTree{}
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

func (t *Table) UnmapBank(addr uint32, bank interface{}, bankNum int) {
	regs, err := bankGetRegs(bank, bankNum)
	if err != nil {
		panic(err)
	}

	for _, reg := range regs {
		switch r := reg.regPtr.(type) {
		case *Mem:
			t.Unmap(addr+reg.offset, addr+reg.offset+uint32(r.VSize)-1)
		case *Reg64:
			t.Unmap(addr+reg.offset, addr+reg.offset+7)
		case *Reg32:
			t.Unmap(addr+reg.offset, addr+reg.offset+3)
		case *Reg16:
			t.Unmap(addr+reg.offset, addr+reg.offset+1)
		case *Reg8:
			t.Unmap(addr+reg.offset, addr+reg.offset+0)
		default:
			panic(fmt.Errorf("invalid reg type: %T", r))
		}
	}
}

func (t *Table) mapBus32(addr uint32, size uint32, io BankIO32, allowremap bool) {
	// fmt.Printf("mapping: %08x-%08x %T\n", addr, addr+size-1, io)
	err := t.table32.InsertRange(addr, addr+size-1, io)
	if err != nil {
		panic(err)
	}
}

func (t *Table) mapBus16(addr uint32, size uint32, io BankIO16, allowremap bool) {
	// fmt.Printf("mapping: %08x-%08x %T\n", addr, addr+size-1, io)
	err := t.table16.InsertRange(addr, addr+size-1, io)
	if err != nil {
		panic(err)
	}
}

func (t *Table) mapBus8(addr uint32, size uint32, io BankIO8, allowremap bool) {
	err := t.table8.InsertRange(addr, addr+size-1, io)
	if err != nil {
		panic(err)
	}
}

func (t *Table) MapReg64(addr uint32, io *Reg64) {
	if addr&7 != 0 {
		panic("unaligned mapping")
	}
	t.mapBus8(addr, 8, io, false)
	t.mapBus16(addr&^1, 8, io, false)
	t.mapBus32(addr&^3, 8, io, false)
}

func (t *Table) MapReg32(addr uint32, io *Reg32) {
	if addr&3 != 0 {
		panic("unaligned mapping")
	}
	t.mapBus8(addr, 4, io, false)
	t.mapBus16(addr&^1, 4, io, false)
	t.mapBus32(addr&^3, 4, io, false)
}

func (t *Table) MapReg16(addr uint32, io *Reg16) {
	if addr&1 != 0 {
		panic("unaligned mapping")
	}
	t.mapBus8(addr, 2, io, false)
	t.mapBus16(addr&^1, 2, io, false)
	t.mapBus32(addr&^3, 4, (*io32to16)(t), true)
}

func (t *Table) MapReg8(addr uint32, io *Reg8) {
	t.mapBus8(addr, 1, io, false)
	t.mapBus16(addr&^1, 2, (*io16to8)(t), true)
	t.mapBus32(addr&^3, 4, (*io32to16)(t), true)
}

func (t *Table) MapMem(addr uint32, mem *Mem) {
	if len(mem.Data)&(len(mem.Data)-1) != 0 {
		panic("memory buffer size is not pow2")
	}

	readonly := mem.Flags&MemFlagReadOnly != 0
	smem := newMemUnalignedLE(mem.Data, mem.WriteCb, readonly)

	if mem.Flags&MemFlag8 != 0 {
		t.mapBus8(addr, uint32(mem.VSize), smem, false)
	}
	if mem.Flags&MemFlag16ForceAlign != 0 {
		t.mapBus16(addr, uint32(mem.VSize), (*memForceAlignLE)(smem), false)
	} else if mem.Flags&MemFlag16Unaligned != 0 {
		t.mapBus16(addr, uint32(mem.VSize), smem, false)
	} else if mem.Flags&MemFlag16Byteswapped != 0 {
		t.mapBus16(addr, uint32(mem.VSize), (*memByteSwappedLE)(smem), false)
	}
	if mem.Flags&MemFlag32ForceAlign != 0 {
		t.mapBus32(addr, uint32(mem.VSize), (*memForceAlignLE)(smem), false)
	} else if mem.Flags&MemFlag32Unaligned != 0 {
		t.mapBus32(addr, uint32(mem.VSize), smem, false)
	} else if mem.Flags&MemFlag32Byteswapped != 0 {
		t.mapBus32(addr, uint32(mem.VSize), (*memByteSwappedLE)(smem), false)
	}
}

func (t *Table) MapMemorySlice(addr uint32, end uint32, mem []uint8, readonly bool) {
	flags := MemFlag8 | MemFlag16Unaligned | MemFlag32Unaligned
	if readonly {
		flags |= MemFlagReadOnly
	}
	t.MapMem(addr, &Mem{
		Data:  mem,
		Flags: flags,
		VSize: int(end - addr + 1),
	})
}

func (t *Table) Unmap(begin uint32, end uint32) {
	t.table8.RemoveRange(begin, end)
	t.table16.RemoveRange(begin, end)
	t.table32.RemoveRange(begin, end)
}

func (t *Table) Read8(addr uint32) uint8 {
	io := t.table8.Search(addr)
	if io == nil {
		log.ModHwIo.WithFields(log.Fields{
			"name": t.Name,
			"addr": emu.Hex32(addr),
		}).Error("unmapped Read8")
		return 0
	}
	if mem, ok := io.(*memUnalignedLE); ok {
		return mem.Read8(addr)
	}
	return io.(BankIO8).Read8(addr)
}

func (t *Table) Write8(addr uint32, val uint8) {
	io := t.table8.Search(addr)
	if io == nil {
		log.ModHwIo.WithFields(log.Fields{
			"name": t.Name,
			"val":  emu.Hex8(val),
			"addr": emu.Hex32(addr),
		}).Error("unmapped Write8")
		return
	}
	if mem, ok := io.(*memUnalignedLE); ok {
		// NOTE: we use the CheckRO format so that the success codepath
		// (that is, when the memory is read-write) is fully inlined and
		// requires no function call.
		ok := mem.Write8CheckRO(addr, val)
		if !ok {
			log.ModHwIo.WithFields(log.Fields{
				"name": t.Name,
				"val":  emu.Hex8(val),
				"addr": emu.Hex32(addr),
			}).Error("Write8 to ROM")
		}
		return
	}
	io.(BankIO8).Write8(addr, val)
}

func (t *Table) Read16(addr uint32) uint16 {
	io := t.table16.Search(addr)
	if io == nil {
		log.ModHwIo.WithFields(log.Fields{
			"name": t.Name,
			"addr": emu.Hex32(addr),
		}).Error("unmapped Read16")
		return 0
	}
	if mem, ok := io.(*memUnalignedLE); ok {
		return mem.Read16(addr)
	}
	return io.(BankIO16).Read16(addr)
}

func (t *Table) Write16(addr uint32, val uint16) {
	io := t.table16.Search(addr)
	if io == nil {
		log.ModHwIo.WithFields(log.Fields{
			"name": t.Name,
			"val":  emu.Hex16(val),
			"addr": emu.Hex32(addr),
		}).Error("unmapped Write16")
		return
	}
	if mem, ok := io.(*memUnalignedLE); ok {
		// NOTE: we use the CheckRO format so that the success codepath
		// (that is, when the memory is read-write) is fully inlined and
		// requires no function call.
		ok := mem.Write16CheckRO(addr, val)
		if !ok {
			log.ModHwIo.WithFields(log.Fields{
				"name": t.Name,
				"val":  emu.Hex16(val),
				"addr": emu.Hex32(addr),
			}).Error("Write16 to ROM")
		}
		return
	}
	io.(BankIO16).Write16(addr, val)
}

func (t *Table) Read32(addr uint32) uint32 {
	io := t.table32.Search(addr)
	if io == nil {
		log.ModHwIo.WithFields(log.Fields{
			"name": t.Name,
			"addr": emu.Hex32(addr),
		}).Error("unmapped Read32")
		return 0
	}
	if mem, ok := io.(*memUnalignedLE); ok {
		return mem.Read32(addr)
	}
	return io.(BankIO32).Read32(addr)
}

func (t *Table) Write32(addr uint32, val uint32) {
	io := t.table32.Search(addr)
	if io == nil {
		log.ModHwIo.WithFields(log.Fields{
			"name": t.Name,
			"val":  emu.Hex32(val),
			"addr": emu.Hex32(addr),
		}).Error("unmapped Write32")
		return
	}
	if mem, ok := io.(*memUnalignedLE); ok {
		// NOTE: we use the CheckRO format so that the success codepath
		// (that is, when the memory is read-write) is fully inlined and
		// requires no function call.
		ok := mem.Write32CheckRO(addr, val)
		if !ok {
			log.ModHwIo.WithFields(log.Fields{
				"name": t.Name,
				"val":  emu.Hex32(val),
				"addr": emu.Hex32(addr),
			}).Error("Write32 to ROM")
		}
		return
	}
	io.(BankIO32).Write32(addr, val)
}

func (t *Table) FetchPointer(addr uint32) []uint8 {
	io := t.table8.Search(addr)
	if mem, ok := io.(*memUnalignedLE); ok {
		return mem.FetchPointer(addr)
	}
	return nil
}

func (t *Table) WaitStates() int {
	return t.ws
}
