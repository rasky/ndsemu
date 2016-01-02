package emu

import "fmt"

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
