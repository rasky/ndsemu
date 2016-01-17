package emu

import (
	"unsafe"
)

func Read32LE(mem []byte) uint32 {
	_ = mem[3] // trigger panic if out of bounds
	return *(*uint32)(unsafe.Pointer(&mem[0]))
}

func Write32LE(mem []byte, val uint32) {
	_ = mem[3] // trigger panic if out of bounds
	*(*uint32)(unsafe.Pointer(&mem[0])) = val
}

func Read16LE(mem []byte) uint16 {
	_ = mem[1] // trigger panic if out of bounds
	return *(*uint16)(unsafe.Pointer(&mem[0]))
}

func Write16LE(mem []byte, val uint16) {
	_ = mem[1] // trigger panic if out of bounds
	*(*uint16)(unsafe.Pointer(&mem[0])) = val
}
