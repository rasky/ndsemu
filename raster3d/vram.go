package raster3d

import (
	"ndsemu/emu"
	"unsafe"
)

// VramTextureBank abstracts access to the VRAM banks
// that were mapped for texture pixels (texels).
//
// The NDS allows up to 4 different 128K banks to be mapped
// for textures (through the memory controller). Moreover,
// textures can be very large, and a single texture might span
// across two different banks (this means that we can't find out
// which bank a texture lies within, and then get a unsafe.Pointer
// within that bank).
type VramTextureBank struct {
	Slots [4][]byte
}

func (vt *VramTextureBank) Get8(off uint32) uint8 {
	return vt.Slots[off>>17][off&0x1FFFF]
}

func (vt *VramTextureBank) Get16(off uint32) uint16 {
	return emu.Read16LE(vt.Slots[off>>17][off&0x1FFFF:])
}

// VramTexturePaletteBank abstracts access to the VRAM banks
// that were mapped for texture palettes.
//
// The NDS allows up to 6 different 16K banks to be mapped for
// texture palettes (notice that these can be physical 16K banks,
// or parts of larger banks; this whole mess is handled by the
// memory controller anyway). Fortunately, a palette never spans
// across two banks, so for each palette we can create an
// optimized VramTexturePalette object (that is a wrapper of
// an unsafe.Pointer), so that the rasterizer can access it
// without much overhead.
type VramTexturePaletteBank struct {
	Slots [6][]byte
}

func (vt *VramTexturePaletteBank) Palette(off int) VramTexturePalette {
	return VramTexturePalette{ptr: unsafe.Pointer(&vt.Slots[off>>14][off&0x3FFF])}
}

type VramTexturePalette struct {
	ptr unsafe.Pointer
}

func (vt VramTexturePalette) Lookup(c uint8) uint16 {
	return *(*uint16)(unsafe.Pointer(uintptr(vt.ptr) + uintptr(int(c)*2)))
}
