package emu

import (
	"fmt"
	"reflect"
	"unsafe"
)

type BankPointer struct {
	ptr unsafe.Pointer
}

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

func NewBankPointerIO(bio BankIO) BankPointer {
	ptr := unsafe.Pointer(&bio)
	ptr = unsafe.Pointer(uintptr(ptr) | 2)
	return BankPointer{ptr}
}

func NewBankPointerMem(ptr unsafe.Pointer, ro bool) BankPointer {
	if ro {
		ptr = unsafe.Pointer(uintptr(ptr) | 1)
	}
	return BankPointer{ptr}
}

func (bp BankPointer) Empty() bool {
	return bp.ptr == nil
}

func (bp BankPointer) IsMemory() bool {
	return uintptr(bp.ptr)&2 == 0
}
func (bp BankPointer) IsIO() bool {
	return !bp.IsMemory()
}

func (bp BankPointer) ReadOnly() bool {
	return uintptr(bp.ptr)&1 != 0
}

func (bp BankPointer) MemBase() unsafe.Pointer {
	return unsafe.Pointer(uintptr(bp.ptr) &^ 3)
}

func (bp BankPointer) Mem(off uint32) unsafe.Pointer {
	return unsafe.Pointer(uintptr(bp.MemBase()) + uintptr(off))
}

func (bp BankPointer) IO() *BankIO {
	return (*BankIO)(bp.MemBase())
}

const (
	cAddressSpaceBits = 28
	cAddressSpaceMask = (1 << cAddressSpaceBits) - 1
	cBankSizeBits     = 14 // 16*1024
	cBankSize         = (1 << cBankSizeBits)
	cBankMask         = cBankSize - 1
	cBankShift        = cAddressSpaceBits - cBankSizeBits
	cNumBanks         = 1 << cBankShift
)

type BankedBus struct {
	Banks         [cNumBanks]BankPointer
	NumWaitStates int
}

func (bus *BankedBus) WaitStates() int {
	return bus.NumWaitStates
}

func bankNumFromAddress(address uint32) uint32 {
	return (address & cAddressSpaceMask) >> cBankSizeBits
}

func (bus *BankedBus) Read8(address uint32) uint8 {
	bnk := bus.Banks[bankNumFromAddress(address)]
	if bnk.Empty() {
		Log().WithField("addr", Hex32(address)).Error("unmapped Read8")
		return 0xFF
	}
	if bnk.IsIO() {
		return (*bnk.IO()).Read8(address)
	}
	val := *(*uint8)(bnk.Mem(address & cBankMask))
	return val
}

func (bus *BankedBus) Read16(address uint32) uint16 {
	bnk := bus.Banks[bankNumFromAddress(address)]
	if bnk.Empty() {
		Log().WithField("addr", Hex32(address)).Error("unmapped Read16")
		DebugBreak("unmapped Read16")
		return 0xFF
	}
	if bnk.IsIO() {
		return (*bnk.IO()).Read16(address)
	}
	val := *(*uint16)(bnk.Mem(address & cBankMask))
	return val
}

func (bus *BankedBus) Read32(address uint32) uint32 {
	bnk := bus.Banks[bankNumFromAddress(address)]
	if bnk.Empty() {
		Log().WithField("addr", Hex32(address)).Error("unmapped Read32")
		DebugBreak("unmapped Read32")
		return 0xFFFFFFFF
	}
	if bnk.IsIO() {
		return (*bnk.IO()).Read32(address)
	}
	val := *(*uint32)(bnk.Mem(address & cBankMask))
	return val
}

func (bus *BankedBus) Write8(address uint32, value uint8) {
	bnk := bus.Banks[bankNumFromAddress(address)]
	if bnk.Empty() {
		Log().WithField("ptr", fmt.Sprintf("[%08x]=%02x", address, value)).Error("unmapped Write8")
		return
	}
	if bnk.IsIO() {
		(*bnk.IO()).Write8(address, value)
		return
	}
	// Memory
	if bnk.ReadOnly() {
		Log().WithField("ptr", fmt.Sprintf("[%08x]=%02x", address, value)).Error("Write8 to ROM")
		return
	}
	*(*uint8)(bnk.Mem(address & cBankMask)) = value
}

func (bus *BankedBus) Write16(address uint32, value uint16) {
	bnk := bus.Banks[bankNumFromAddress(address)]
	if bnk.Empty() {
		Log().WithField("ptr", fmt.Sprintf("[%08x]=%04x", address, value)).Error("unmapped Write16")
		DebugBreak("unmapped Write16")
		return
	}
	if bnk.IsIO() {
		(*bnk.IO()).Write16(address, value)
		return
	}
	// Memory
	if bnk.ReadOnly() {
		Log().WithField("ptr", fmt.Sprintf("[%08x]=%02x", address, value)).Error("Write16 to ROM")
		DebugBreak("Write16 to ROM")
		return
	}
	*(*uint16)(bnk.Mem(address & cBankMask)) = value
}

func (bus *BankedBus) Write32(address, value uint32) {
	bnk := bus.Banks[bankNumFromAddress(address)]
	if bnk.Empty() {
		Log().WithField("ptr", fmt.Sprintf("[%08x]=%08x", address, value)).Error("unmapped Write32")
		return
	}
	if bnk.IsIO() {
		(*bnk.IO()).Write32(address, value)
		return
	}
	if bnk.ReadOnly() {
		Log().WithField("ptr", fmt.Sprintf("[%08x]=%02x", address, value)).Error("Write32 to ROM")
		DebugBreak("Write32 to ROM")
		return
	}
	*(*uint32)(bnk.Mem(address & cBankMask)) = value
}

func (bus *BankedBus) FetchPointer(address uint32) []uint8 {
	bnk := bus.Banks[bankNumFromAddress(address)]
	if bnk.Empty() {
		// log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("unmapped FetchPointer")
		return nil
	}
	if bnk.IsIO() {
		// log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("FetchPointer to I/O area")
		return nil
	}

	off := address & cBankMask
	size := int(cBankSize - off)
	sh := reflect.SliceHeader{Data: uintptr(bnk.Mem(off)), Len: size, Cap: size}
	return *(*[]uint8)(unsafe.Pointer(&sh))
}

// MapMemory maps a physical memory bank within a virtual address space.
// start/end respectively specify the beginning and the end of the virtual area (the covered
// area must be a 64k multiple at this point). ptr is the unsafe.Pointer to the memory,
// and physsize is the physical size of the memory.
// If the physical size is bigger than the virtual size, a panic is generated.
// If the physical size is smaller than the virtual size, the memory is mirrored within the
// virtual area; if it's not a multiple, a panic is generated
func (bus *BankedBus) MapMemory(start uint32, end uint32, ptr unsafe.Pointer, physsize int, ro bool) {
	if start&^cAddressSpaceMask != 0 || end&^cAddressSpaceMask != 0 {
		panic("invalid bus mapping")
	}
	if start > end {
		panic("invalid start < end")
	}
	if start&cBankMask != 0 || (end+1)&cBankMask != 0 {
		panic("start/end not at bank boundaries")
	}

	vsize := end - start + 1
	if vsize%uint32(physsize) != 0 {
		panic("vsize is not a multiple of physsize")
	}

	if physsize%cBankSize != 0 {
		if !ro {
			panic("physsize is not a multiple of banksize")
		}

		// Create a rounded-up buffer containing a copy of the
		// read-only memory, padded with 0xFF.
		newsize := (physsize + cBankSize - 1) / cBankSize * cBankSize
		buf := make([]uint8, newsize)
		for i := 0; i < physsize; i++ {
			buf[i] = *(*uint8)(ptr)
			ptr = unsafe.Pointer(uintptr(ptr) + 1)
		}
		for i := physsize; i < newsize; i++ {
			buf[i] = 0xFF
		}
		ptr = unsafe.Pointer(&buf[0])
		physsize = newsize
	}

	bnk := bankNumFromAddress(start)
	for j := 0; j < int(vsize)/physsize; j++ {
		p := ptr
		for i := 0; i < int(physsize/cBankSize); i++ {
			if !bus.Banks[bnk].Empty() {
				panic("bank is already mapped")
			}
			bus.Banks[bnk] = NewBankPointerMem(unsafe.Pointer(p), ro)
			p = unsafe.Pointer(uintptr(p) + uintptr(cBankSize))
			bnk++
		}
	}
}

func (bus *BankedBus) MapMemorySlice(start uint32, end uint32, buf []byte, ro bool) {
	bus.MapMemory(start, end, unsafe.Pointer(&buf[0]), len(buf), ro)
}

func (bus *BankedBus) MapIORegs(start uint32, end uint32, io BankIO) {
	if start&^cAddressSpaceMask != 0 || end&^cAddressSpaceMask != 0 {
		panic("invalid bus mapping")
	}
	if start > end {
		panic("invalid start < end")
	}
	if start&cBankMask != 0 || (end+1)&cBankMask != 0 {
		panic("start/end not at bank boundaries")
	}

	vsize := end - start + 1
	bnk := bankNumFromAddress(start)
	for j := 0; j < int(vsize/cBankSize); j++ {
		if !bus.Banks[bnk].Empty() {
			panic("bank is already mapped")
		}
		bus.Banks[bnk] = NewBankPointerIO(io)
		bnk++
	}
}

func (bus *BankedBus) Unmap(start uint32, end uint32) {
	if start&^cAddressSpaceMask != 0 || end&^cAddressSpaceMask != 0 {
		panic("invalid bus mapping")
	}
	if start > end {
		panic("invalid start < end")
	}
	if start&cBankMask != 0 || (end+1)&cBankMask != 0 {
		panic("start/end not at bank boundaries")
	}

	vsize := end - start + 1
	bnk := bankNumFromAddress(start)
	for j := 0; j < int(vsize/cBankSize); j++ {
		bus.Banks[bnk] = BankPointer{}
		bnk++
	}

}
