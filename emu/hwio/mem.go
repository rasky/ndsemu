package hwio

import (
	"unsafe"

	"encoding/binary"
)

type mem8 []uint8

func (m mem8) Read8(addr uint32) uint8 {
	off := addr & uint32(len(m)-1)
	return m[off]
}

func (m mem8) Write8(addr uint32, val uint8) {
	off := addr & uint32(len(m)-1)
	m[off] = val
}

func (m mem8) FetchPointer(addr uint32) []uint8 {
	off := addr & uint32(len(m)-1)
	return m[off:]
}

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
}

func newMemUnalignedLE(mem []byte) *memUnalignedLE {
	if len(mem)&(len(mem)-1) != 0 {
		panic("memory buffer size is not pow2")
	}
	return &memUnalignedLE{
		ptr:  unsafe.Pointer(&mem[0]),
		mask: uint32(len(mem) - 1),
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

func (m *memUnalignedLE) Write8(addr uint32, val uint8) {
	off := uintptr(addr & m.mask)
	*(*uint8)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
}

func (m *memUnalignedLE) Read16(addr uint32) uint16 {
	off := uintptr(addr & m.mask)
	return *(*uint16)(unsafe.Pointer(uintptr(m.ptr) + off))
}

func (m *memUnalignedLE) Write16(addr uint32, val uint16) {
	off := uintptr(addr & m.mask)
	*(*uint16)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
}

func (m *memUnalignedLE) Read32(addr uint32) uint32 {
	off := uintptr(addr & m.mask)
	return *(*uint32)(unsafe.Pointer(uintptr(m.ptr) + off))
}

func (m *memUnalignedLE) Write32(addr uint32, val uint32) {
	off := uintptr(addr & m.mask)
	*(*uint32)(unsafe.Pointer(uintptr(m.ptr) + off)) = val
}

// 16-bit access to memory with forced to 16-bit boundary.
// Eg: Read16(1) == Read16(0)
type mem16LittleEndianForceAlign []uint8

func (m mem16LittleEndianForceAlign) Read16(addr uint32) uint16 {
	off := (addr & uint32(len(m)-1)) &^ 1
	_ = m[off+1] // trigger panic for out of bounds
	return *(*uint16)(unsafe.Pointer(&m[off]))
}

func (m mem16LittleEndianForceAlign) Write16(addr uint32, val uint16) {
	off := (addr & uint32(len(m)-1)) &^ 1
	_ = m[off+1] // trigger panic for out of bounds
	*(*uint16)(unsafe.Pointer(&m[off])) = val
}

// 16-bit access to memory with byteswapping on unaligned access
// Eg: Read16(1) == byteswap(Read16(0))
type mem16LittleEndianByteSwap []uint8

func (m mem16LittleEndianByteSwap) Read16(addr uint32) uint16 {
	off := (addr & uint32(len(m)-1))
	return uint16(m[off]) | uint16(m[off^1])
}

func (m mem16LittleEndianByteSwap) Write16(addr uint32, val uint16) {
	off := (addr & uint32(len(m)-1))
	m[off] = uint8(val)
	m[off^1] = uint8(val >> 8)
}

// 16-bit access to memory with forced to 16-bit boundary.
// Eg: Read16(1) == Read16(0)
type mem32LittleEndianForceAlign []uint8

func (m mem32LittleEndianForceAlign) Read32(addr uint32) uint32 {
	off := (addr & uint32(len(m)-1)) &^ 3
	_ = m[off+3] // trigger panic for out of bounds
	return *(*uint32)(unsafe.Pointer(&m[off]))
}

func (m mem32LittleEndianForceAlign) Write32(addr uint32, val uint32) {
	off := (addr & uint32(len(m)-1)) &^ 3
	_ = m[off+3] // trigger panic for out of bounds
	*(*uint32)(unsafe.Pointer(&m[off])) = val
}

// 16-bit access to memory with byteswapping on unaligned access
// Eg: Read16(1) == byteswap(Read16(0))
type mem32LittleEndianByteSwap []uint8

func (m mem32LittleEndianByteSwap) Read32(addr uint32) uint32 {
	off := (addr & uint32(len(m)-1))
	rot := (off & 3) * 8
	off = off &^ 3
	val := binary.LittleEndian.Uint32(m[off : off+4])
	return (val >> rot) | (val << (32 - rot))
}

func (m mem32LittleEndianByteSwap) Write32(addr uint32, val uint32) {
	off := (addr & uint32(len(m)-1))
	rot := (off & 3) * 8
	off = off &^ 3
	val = (val << rot) | (val >> (32 - rot))
	binary.LittleEndian.PutUint32(m[off:off+4], val)
}

type MemFlags int

const (
	MemFlag8 MemFlags = (1 << iota)
	MemFlag16ForceAlign
	MemFlag16Unaligned
	MemFlag16Byteswapped
	MemFlag32ForceAlign
	MemFlag32Unaligned
	MemFlag32Byteswapped
)

type Mem struct {
	Name  string
	Data  []byte
	VSize int
	Flags MemFlags
}
