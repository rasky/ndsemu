package emu

type Bus interface {
	WaitStates() int

	Read32(address uint32) uint32
	Write32(address uint32, val uint32)

	Read16(address uint32) uint16
	Write16(address uint32, val uint16)

	Read8(address uint32) uint8
	Write8(address uint32, val uint8)
}
