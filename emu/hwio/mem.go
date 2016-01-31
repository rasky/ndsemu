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
	ro   bool
}

func newMemUnalignedLE(mem []byte, wcb func(uint32, int), readonly bool) *memUnalignedLE {
	if len(mem)&(len(mem)-1) != 0 {
		panic("memory buffer size is not pow2")
	}
	return &memUnalignedLE{
		ptr:  unsafe.Pointer(&mem[0]),
		mask: uint32(len(mem) - 1),
		wcb:  wcb,
		ro:   readonly,
	}
}

func (m *memUnalignedLE) FetchPointer(addr uint32) []uint8 {
	off := uintptr(addr & m.mask)
	buf := (*[1 << 30]uint8)(unsafe.Pointer(uintptr(m.ptr) + off))
	return buf[:m.mask+1-uint32(off)]
}

func (m *memUnalignedLE) Read8(addr uint32) uint8 {
	off := uintptr(addr & m.mask)
	return *(*uint8)(unsafe.Pointer(uintptr(m.ptr) + off))
}

func (m *memUnalignedLE) Write8CheckRO(addr uint32, val uint8) bool {
	off := uintptr(addr & m.mask)
	if !m.ro {
		*(*uint8)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
		if m.wcb != nil {
			m.wcb(addr, 1)
		}
		return true
	}
	return false
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
	if !m.ro {
		*(*uint16)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
		if m.wcb != nil {
			m.wcb(addr, 2)
		}
		return true
	}
	return false
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
	if !m.ro {
		*(*uint32)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
		if m.wcb != nil {
			m.wcb(addr, 4)
		}
		return true
	}
	return false
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
	MemFlagReadOnly                             // all writes are forbidden
)

type Mem struct {
	Name    string            // name of the memory area (for debugging)
	Data    []byte            // actual memory buffer
	VSize   int               // virtual size of the memory (can be bigger than physical size)
	Flags   MemFlags          // flags determining how the memory can be accessed
	WriteCb func(uint32, int) // optional write callback (receives full address and number of bytes written)
}
