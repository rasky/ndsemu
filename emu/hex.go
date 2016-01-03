package emu

import (
	"encoding/binary"

	"fmt"
)

func Swap16(val uint16) uint16 {
	return (val >> 8) | (val << 8)
}

func Swap32(val uint32) uint32 {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], val)
	return binary.BigEndian.Uint32(buf[:])
}

func Swap64(val uint64) uint64 {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], val)
	return binary.BigEndian.Uint64(buf[:])
}

func Hex64(val uint64) string {
	return fmt.Sprintf("%016x", val)
}

func Hex32(val uint32) string {
	return fmt.Sprintf("%08x", val)
}

func Hex16(val uint16) string {
	return fmt.Sprintf("%04x", val)
}

func Hex8(val uint8) string {
	return fmt.Sprintf("%02x", val)
}
