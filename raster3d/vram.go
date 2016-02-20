package raster3d

import (
	"ndsemu/emu"
	"unsafe"
)

type VramTextureBank struct {
	Slots [4][]byte
}

func (vt *VramTextureBank) Get8(off uint32) uint8 {
	return vt.Slots[off>>17][off&0x1FFFF]
}

func (vt *VramTextureBank) Get16(off uint32) uint16 {
	return emu.Read16LE(vt.Slots[off>>17][off&0x1FFFF:])
}

/************************************************
 * VramTexturePalette & VramTexturePaletteBank
 ************************************************/

type VramTexturePalette struct {
	ptr unsafe.Pointer
}

func (vt VramTexturePalette) Lookup(c uint8) uint16 {
	return *(*uint16)(unsafe.Pointer(uintptr(vt.ptr) + uintptr(int(c)*2)))
}

type VramTexturePaletteBank struct {
	Slots [6][]byte
}

func (vt *VramTexturePaletteBank) Palette(off int) VramTexturePalette {
	return VramTexturePalette{ptr: unsafe.Pointer(&vt.Slots[off>>14][off&0x3FFF])}
}
