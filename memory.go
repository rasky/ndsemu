package main

import (
	"fmt"
	"unsafe"

	log "gopkg.in/Sirupsen/logrus.v0"
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

type BankedBus struct {
	Banks [4096]BankPointer
}

func (bus *BankedBus) Read8(address uint32) uint8 {
	bnk := bus.Banks[(address>>16)&0xFFF]
	if bnk.Empty() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("unmapped Read8")
		return 0xFF
	}
	if bnk.IsIO() {
		return (*bnk.IO()).Read8(address)
	}
	val := *(*uint8)(bnk.Mem(address & 0xFFFF))
	return val
}

func (bus *BankedBus) Read16(address uint32) uint16 {
	bnk := bus.Banks[(address>>16)&0xFFF]
	if bnk.Empty() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("unmapped Read16")
		return 0xFF
	}
	if bnk.IsIO() {
		return (*bnk.IO()).Read16(address)
	}
	val := *(*uint16)(bnk.Mem(address & 0xFFFF))
	return val
}

func (bus *BankedBus) Read32(address uint32) uint32 {
	bnk := bus.Banks[(address>>16)&0xFFF]
	if bnk.Empty() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("unmapped Read32")
		return 0xFFFFFFFF
	}
	if bnk.IsIO() {
		return (*bnk.IO()).Read32(address)
	}
	val := *(*uint32)(bnk.Mem(address & 0xFFFF))
	return val
}

func (bus *BankedBus) Write8(address uint32, value uint8) {
	bnk := bus.Banks[(address>>16)&0xFFF]
	if bnk.Empty() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("unmapped Write8")
		return
	}
	if bnk.IsIO() {
		(*bnk.IO()).Write8(address, value)
		return
	}
	// Memory
	if bnk.ReadOnly() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("writing to ROM")
		return
	}
	log.WithFields(log.Fields{
		"ptr": fmt.Sprintf("%08x", address),
		"val": fmt.Sprintf("%02x", value),
	}).Info("Write8")
	*(*uint8)(bnk.Mem(address & 0xFFFF)) = value
}

func (bus *BankedBus) Write16(address uint32, value uint16) {
	bnk := bus.Banks[(address>>16)&0xFFF]
	if bnk.Empty() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("unmapped Write16")
		return
	}
	if bnk.IsIO() {
		(*bnk.IO()).Write16(address, value)
		return
	}
	// Memory
	if bnk.ReadOnly() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("writing to ROM")
		return
	}
	log.WithFields(log.Fields{
		"ptr": fmt.Sprintf("%08x", address),
		"val": fmt.Sprintf("%04x", value),
	}).Info("Write16")
	*(*uint16)(bnk.Mem(address & 0xFFFF)) = value
}

func (bus *BankedBus) Write32(address, value uint32) {
	bnk := bus.Banks[(address>>16)&0xFFF]
	if bnk.Empty() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("unmapped Write32")
		return
	}
	if bnk.IsIO() {
		(*bnk.IO()).Write32(address, value)
		return
	}
	if bnk.ReadOnly() {
		log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("writing to ROM")
		return
	}

	log.WithFields(log.Fields{
		"ptr": fmt.Sprintf("%08x", address),
		"val": fmt.Sprintf("%08x", value),
	}).Info("Write32")
	*(*uint32)(bnk.Mem(address & 0xFFFF)) = value
}

func (bus *BankedBus) MapMemory(address uint32, ptr unsafe.Pointer, size int, ro bool) {
	if address&0xF0000000 != 0 {
		panic("invalid bus mapping")
	}
	address &= 0x0FFFFFFF
	bnk := (address >> 16) & 0xFFF
	if bnk<<16 != address {
		if !ro {
			panic("invalid bus mapping")
		}
		buf := make([]byte, 65536)
		for i := 0; i < 65536; i++ {
			if i >= int(address&0xFFFF) && size > 0 {
				buf[i] = *(*byte)(ptr)
				ptr = unsafe.Pointer(uintptr(ptr) + 1)
				size--
			} else {
				buf[i] = 0xFF
			}
		}
		if !bus.Banks[bnk].Empty() {
			panic("reused memory bank")
		}
		bus.Banks[bnk] = NewBankPointerMem(unsafe.Pointer(&buf[0]), ro)
		bnk++
	}

	for size > 0 {
		if !bus.Banks[bnk].Empty() {
			panic("reused memory bank")
		}

		bus.Banks[bnk] = NewBankPointerMem(ptr, ro)
		ptr = unsafe.Pointer(uintptr(ptr) + uintptr(65536))
		size -= 65536
		bnk++
	}
}

/*
type dummyIO32 struct {
	BankIO8
}

func (d32 dummyIO32) Read32(address uint32) uint32 {
	log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("invalid Read32 from IO8")
	return 0xFFFFFFFF
}
func (d32 dummyIO32) Write32(address uint32, val uint32) {
	log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("invalid Write32 to IO8")
}

func (bus *BankedBus) MapIO8(address uint32, io8 BankIO8) {
	if address&0xF0000000 != 0 {
		panic("invalid bus mapping")
	}
	address &= 0x0FFFFFFF
	bnk := address >> 16
	if bnk<<16 != address {
		panic("invalid bus mapping")
	}
	if !bus.Banks[bnk].Empty() {
		panic("reused memory bank")
	}
	bus.Banks[bnk] = NewBankPointerIO(dummyIO32{io8})
}


type dummyIO8 struct {
	BankIO32
}

func (d8 dummyIO8) Read8(address uint32) uint8 {
	log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("invalid Read8 from IO32")
	return 0xFF
}
func (d8 dummyIO8) Write8(address uint32, val uint8) {
	log.WithField("ptr", fmt.Sprintf("%08x", address)).Error("invalid Write8 to IO32")
}

func (bus *BankedBus) MapIO32(address uint32, io32 BankIO32) {
	if address&0xF0000000 != 0 {
		panic("invalid bus mapping")
	}
	address &= 0x0FFFFFFF
	bnk := address >> 16
	if bnk<<16 != address {
		panic("invalid bus mapping")
	}
	if !bus.Banks[bnk].Empty() {
		panic("reused memory bank")
	}
	bus.Banks[bnk] = NewBankPointerIO(dummyIO8{io32})
}
*/

func (bus *BankedBus) MapIORegs(address uint32, io BankIO) {
	if address&0xF0000000 != 0 {
		panic("invalid bus mapping")
	}
	address &= 0x0FFFFFFF
	bnk := address >> 16
	if bnk<<16 != address {
		panic("invalid bus mapping")
	}
	if !bus.Banks[bnk].Empty() {
		panic("reused memory bank")
	}
	bus.Banks[bnk] = NewBankPointerIO(io)
}
