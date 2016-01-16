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

// 16-bit access to memory with correct unalignment
type mem16LittleEndianUnaligned []uint8

func (m mem16LittleEndianUnaligned) Read16(addr uint32) uint16 {
	off := (addr & uint32(len(m)-1))
	_ = m[off+1] // trigger panic for out of bounds
	return *(*uint16)(unsafe.Pointer(&m[off]))
}

func (m mem16LittleEndianUnaligned) Write16(addr uint32, val uint16) {
	off := (addr & uint32(len(m)-1))
	_ = m[off+1] // trigger panic for out of bounds
	*(*uint16)(unsafe.Pointer(&m[off])) = val
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

// 16-bit access to memory with correct unalignment
type mem32LittleEndianUnaligned []uint8

func (m mem32LittleEndianUnaligned) Read32(addr uint32) uint32 {
	off := (addr & uint32(len(m)-1))
	_ = m[off+3] // trigger panic for out of bounds
	return *(*uint32)(unsafe.Pointer(&m[off]))
}

func (m mem32LittleEndianUnaligned) Write32(addr uint32, val uint32) {
	off := (addr & uint32(len(m)-1))
	_ = m[off+3] // trigger panic for out of bounds
	*(*uint32)(unsafe.Pointer(&m[off])) = val
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
