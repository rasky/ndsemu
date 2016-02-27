package e2d

import "io"

const VramSmallestBankSize = 8 * 1024

// VramLinearBank is an abstraction that linearizes the vram banks mapped by
// the NDS9 for the graphic engines.
//
// VRAM is made by different separate banks, that can be mapped at different
// addresses and with different orders by the NDS9 (see the HwMemoryContorller).
// So for instance, the NDS9 might map at 0x62000000 the banks C, B, A, in that
// order, consecutively.
//
// The graphic engine accesses VRAM through the same memory mapping; for the
// purpose of writing our own code in a sane way, VramLinearBank can be used
// to index the VRAM over the different banks.
type VramLinearBank struct {
	Ptr [32][]uint8
}

// VramLinearBankId is an enum that is used in calls to VramLinearBank to
// declare which kind of VRAM bank we want to access.
type VramLinearBankId int

const (
	// Request access to the BG VRAM
	VramLinearBG VramLinearBankId = iota

	// Request access to the OAM RAM
	VramLinearOAM

	// Request access to the BG Extended Palettes
	VramLinearBGExtPal

	// Request access to the OBJ Extended Palette
	VramLinearOBJExtPal
)

type MemoryController interface {

	// Get access to the palette RAM for the specified engine
	VramPalette(engine int) []byte

	// Get access to the OAM RAM
	VramOAM(engine int) []byte

	// Return the VRAM linear bank that will be accessed by the specified engine.
	// The linear bank is 256k big, and can be accessed as 8-bit or 16-bit.
	// byteOffset is the offset within the VRAM from which the 256k bank starts.
	//
	// If the requested bank is unmapped, a zero-filled area is returned. If the
	// requested bank is mapped for less than 256K, the missing areas will be
	// zero-filled as well.
	VramLinearBank(engine int, which VramLinearBankId, baseOffset int) VramLinearBank

	// Get access to the raw VRAM bank (beyond any mapping)
	VramRawBank(bank int) []byte
}

func (vb *VramLinearBank) Dump(w io.Writer) {
	for i := range vb.Ptr {
		w.Write(vb.Ptr[i][:VramSmallestBankSize])
	}
}

func (vb *VramLinearBank) FetchPointer(off int) []uint8 {
	bank := vb.Ptr[off/VramSmallestBankSize]
	off &= (VramSmallestBankSize - 1)
	return bank[off:]
}

func (vb *VramLinearBank) Get8(off int) uint8 {
	ptr := vb.FetchPointer(off)
	return ptr[0]
}

func (vb *VramLinearBank) Get16(off int) uint16 {
	ptr := vb.FetchPointer(off * 2)
	return uint16(ptr[0]) | uint16(ptr[1])<<8
}
