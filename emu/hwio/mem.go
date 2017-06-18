package hwio

import (
	"ndsemu/emu"
	log "ndsemu/emu/logger"
	"unsafe"
)

// 16-bit / 32-bit access to memory with correct unalignment
// This is the main structure used for linear memory access, and
// should be used by default for most memory areas.
// We use this structure by pointer rather than by value because
// it is stored as BankIO interface within Table, and checking if
// a concrete pointer type is behind the interface is faster than
// checking a non-pointer type (see for instance Table.Write32
// that checks for *memUnaligned32 as fast-path).
type memUnalignedLE struct {
	ptr  unsafe.Pointer
	mask uint32
	wcb  func(uint32, int)
	ro   uint8 // 0: read/write, 1: readonly, 2: silent readonly (no log)
}

func newMemUnalignedLE(mem []byte, wcb func(uint32, int), roflag uint8) *memUnalignedLE {
	if len(mem)&(len(mem)-1) != 0 {
		panic("memory buffer size is not pow2")
	}
	return &memUnalignedLE{
		ptr:  unsafe.Pointer(&mem[0]),
		mask: uint32(len(mem) - 1),
		wcb:  wcb,
		ro:   roflag,
	}
}

func (m *memUnalignedLE) FetchPointer(addr uint32) []uint8 {
	off := uintptr(addr & m.mask)
	buf := (*[1 << 30]uint8)(unsafe.Pointer(uintptr(m.ptr) + off))
	len := m.mask + 1 - uint32(off)
	return buf[:len:len]
}

func (m *memUnalignedLE) Read8(addr uint32) uint8 {
	off := uintptr(addr & m.mask)
	return *(*uint8)(unsafe.Pointer(uintptr(m.ptr) + off))
}

func (m *memUnalignedLE) Write8CheckRO(addr uint32, val uint8) bool {
	off := uintptr(addr & m.mask)
	if m.ro == 0 {
		*(*uint8)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
		if m.wcb != nil {
			m.wcb(addr, 1)
		}
		return true
	}
	return m.ro == 2 // fake success if we're in silent mode
}

func (m *memUnalignedLE) Write8(addr uint32, val uint8) {
	if !m.Write8CheckRO(addr, val) {
		log.ModHwIo.WithFields(log.Fields{
			"val":  emu.Hex8(val),
			"addr": emu.Hex32(addr),
		}).Error("Write8 to readonly memory")
	}
}

func (m *memUnalignedLE) Read16(addr uint32) uint16 {
	off := uintptr(addr & m.mask)
	return *(*uint16)(unsafe.Pointer(uintptr(m.ptr) + off))
}

func (m *memUnalignedLE) Write16CheckRO(addr uint32, val uint16) bool {
	off := uintptr(addr & m.mask)
	if m.ro == 0 {
		*(*uint16)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
		if m.wcb != nil {
			m.wcb(addr, 2)
		}
		return true
	}
	return m.ro == 2 // fake success if we're in silent mode
}

func (m *memUnalignedLE) Write16(addr uint32, val uint16) {
	if !m.Write16CheckRO(addr, val) {
		log.ModHwIo.WithFields(log.Fields{
			"val":  emu.Hex16(val),
			"addr": emu.Hex32(addr),
		}).Error("Write16 to readonly memory")
	}
}

func (m *memUnalignedLE) Read32(addr uint32) uint32 {
	off := uintptr(addr & m.mask)
	return *(*uint32)(unsafe.Pointer(uintptr(m.ptr) + off))
}

func (m *memUnalignedLE) Write32CheckRO(addr uint32, val uint32) bool {
	off := uintptr(addr & m.mask)
	if m.ro == 0 {
		*(*uint32)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
		if m.wcb != nil {
			m.wcb(addr, 4)
		}
		return true
	}
	return m.ro == 2 // fake success if we're in silent mode
}

func (m *memUnalignedLE) Write32(addr uint32, val uint32) {
	if !m.Write32CheckRO(addr, val) {
		log.ModHwIo.WithFields(log.Fields{
			"val":  emu.Hex32(val),
			"addr": emu.Hex32(addr),
		}).Error("Write32 to readonly memory")
	}
}

// Access to memory with forced to 16/32-bit boundary.
// Eg: Read16(1) == Read16(0)
type memForceAlignLE memUnalignedLE

func (m *memForceAlignLE) Read16(addr uint32) uint16 {
	return (*memUnalignedLE)(m).Read16(addr &^ 1)
}

func (m *memForceAlignLE) Write16(addr uint32, val uint16) {
	(*memUnalignedLE)(m).Write16(addr&^1, val)
}

func (m *memForceAlignLE) Read32(addr uint32) uint32 {
	return (*memUnalignedLE)(m).Read32(addr &^ 3)
}

func (m *memForceAlignLE) Write32(addr uint32, val uint32) {
	(*memUnalignedLE)(m).Write32(addr&^3, val)
}

// 16-bit access to memory with byteswapping on unaligned access
// Eg: Read16(1) == byteswap(Read16(0))
type memByteSwappedLE memUnalignedLE

func (m *memByteSwappedLE) Read16(addr uint32) uint16 {
	rot := (addr & 1) * 8
	val := (*memUnalignedLE)(m).Read16(addr &^ 1)
	return (val >> rot) | (val << (16 - rot))
}

func (m *memByteSwappedLE) Write16(addr uint32, val uint16) {
	rot := (addr & 1) * 8
	val = (val << rot) | (val >> (16 - rot))
	(*memUnalignedLE)(m).Write16(addr&^1, val)
}

func (m *memByteSwappedLE) Read32(addr uint32) uint32 {
	rot := (addr & 3) * 8
	val := (*memUnalignedLE)(m).Read32(addr &^ 3)
	return (val >> rot) | (val << (32 - rot))
}

func (m *memByteSwappedLE) Write32(addr uint32, val uint32) {
	rot := (addr & 3) * 8
	val = (val << rot) | (val >> (32 - rot))
	(*memUnalignedLE)(m).Write32(addr&^3, val)
}

type MemFlags int

const (
	MemFlag8             MemFlags = (1 << iota) // 8-bit access is allowed
	MemFlag16Unaligned                          // 16-bit access is allowed, even if unaligned
	MemFlag16ForceAlign                         // 16-bit access is allowed, and it is forcibly aligned to 16-bit boundary
	MemFlag16Byteswapped                        // 16-bit access is allowed, and if not aligned the data is byteswapped
	MemFlag32Unaligned                          // 32-bit access is allowed, even if unaligned
	MemFlag32ForceAlign                         // 32-bit access is allowed, and it is forcibly aligned to 32-bit boundary
	MemFlag32Byteswapped                        // 32-bit access is allowed, and if not aligned the data is byteswapped
	MemFlag8ReadOnly                            // 8-bit accesses are read-only (requires MemFlag8)
	MemFlag16ReadOnly                           // 16-bit accesses are read-only (requires one of MemFlag16*)
	MemFlag32ReadOnly                           // 32-bit accesses are read-only (requires one of MemFlag32*)
	MemFlagNoROLog                              // skip logging attempts to write when configured to readonly

	MemFlagReadOnly = MemFlag8ReadOnly | MemFlag16ReadOnly | MemFlag32ReadOnly // all writes are forbidden
)

// Linear memory area that can be mapped into a Table.
//
// NOTE: this structure does not directly implement the BankIO interface
// for performance reasons. In fact, it would be inefficient to parse all
// the flags at runtime for each memory access to correctly implment it; so,
// clients must call the BankIO8, BankIO16, BankIO32 methods to create
// adaptors that implement memory access depending on the memory bank
// configuration.
type Mem struct {
	Name    string            // name of the memory area (for debugging)
	Data    []byte            // actual memory buffer
	VSize   int               // virtual size of the memory (can be bigger than physical size)
	Flags   MemFlags          // flags determining how the memory can be accessed
	WriteCb func(uint32, int) // optional write callback (receives full address and number of bytes written)
}

func (mem *Mem) roFlag(robit MemFlags) uint8 {
	var roflag uint8
	if mem.Flags&robit != 0 {
		if mem.Flags&MemFlagNoROLog != 0 {
			roflag = 2
		} else {
			roflag = 1
		}
	}
	return roflag
}

func (mem *Mem) BankIO8() BankIO8 {
	if mem.Flags&MemFlag8 == 0 {
		return nil
	}
	roflag := mem.roFlag(MemFlag8ReadOnly)
	return newMemUnalignedLE(mem.Data, mem.WriteCb, roflag)
}

func (mem *Mem) BankIO16() BankIO16 {
	roflag := mem.roFlag(MemFlag16ReadOnly)
	smem := newMemUnalignedLE(mem.Data, mem.WriteCb, roflag)
	if mem.Flags&MemFlag16Unaligned != 0 {
		return smem
	}
	if mem.Flags&MemFlag16ForceAlign != 0 {
		return (*memForceAlignLE)(smem)
	}
	if mem.Flags&MemFlag16Byteswapped != 0 {
		return (*memByteSwappedLE)(smem)
	}
	return nil
}

func (mem *Mem) BankIO32() BankIO32 {
	roflag := mem.roFlag(MemFlag32ReadOnly)
	smem := newMemUnalignedLE(mem.Data, mem.WriteCb, roflag)
	if mem.Flags&MemFlag32Unaligned != 0 {
		return smem
	}
	if mem.Flags&MemFlag32ForceAlign != 0 {
		return (*memForceAlignLE)(smem)
	}
	if mem.Flags&MemFlag32Byteswapped != 0 {
		return (*memByteSwappedLE)(smem)
	}
	return nil
}
