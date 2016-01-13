package main

import (
	"io"
	"ndsemu/emu"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modMemCnt = log.NewModule("memcnt")

type HwMemoryController struct {
	Nds9 *NDS9
	Nds7 *NDS7

	// Registers accessible by NDS9
	VramCntA hwio.Reg8 `hwio:"bank=0,offset=0x0,rwmask=0x9f,writeonly,wcb"`
	VramCntB hwio.Reg8 `hwio:"bank=0,offset=0x1,rwmask=0x9f,writeonly,wcb"`
	VramCntC hwio.Reg8 `hwio:"bank=0,offset=0x2,rwmask=0x9f,writeonly,wcb"`
	VramCntD hwio.Reg8 `hwio:"bank=0,offset=0x3,rwmask=0x9f,writeonly,wcb"`
	VramCntE hwio.Reg8 `hwio:"bank=0,offset=0x4,rwmask=0x9f,writeonly,wcb"`
	VramCntF hwio.Reg8 `hwio:"bank=0,offset=0x5,rwmask=0x9f,writeonly,wcb"`
	VramCntG hwio.Reg8 `hwio:"bank=0,offset=0x6,rwmask=0x9f,writeonly,wcb"`
	WramCnt  hwio.Reg8 `hwio:"bank=0,offset=0x7,rwmask=0x3,wcb"`
	VramCntH hwio.Reg8 `hwio:"bank=0,offset=0x8,rwmask=0x9f,writeonly,wcb"`
	VramCntI hwio.Reg8 `hwio:"bank=0,offset=0x9,rwmask=0x9f,writeonly,wcb"`

	// Read-only access by NDS7
	WramStat hwio.Reg8 `hwio:"bank=1,offset=0x1,readonly,rcb"`

	ExMemCnt  hwio.Reg16 `hwio:"wcb"`
	ExMemStat hwio.Reg16 `hwio:"rwmask=0x007F,wcb"`

	wram [32 * 1024]byte

	// Banks of VRAM that can be mapped to different addresses
	vram      [9][]byte
	unmapVram [9]func()

	// Current mapping of BG Extended Palette. We keep the mapping for each
	// engine (A & B), and each slot (4 of them, 8KB each one).
	BgExtPalette [2][4][]byte

	// Current mapping of OBJ Extended Palette. We keep the mapping for each
	// engine (A & B), with a single palette (8 KB)
	ObjExtPalette [2][]byte

	TexturePalette [6][]byte
}

func NewMemoryController(nds9 *NDS9, nds7 *NDS7, vram []byte) *HwMemoryController {
	mc := &HwMemoryController{
		Nds9: nds9,
		Nds7: nds7,
	}
	hwio.MustInitRegs(mc)

	mc.vram[0] = vram[0 : 128*1024]
	vram = vram[128*1024:]

	mc.vram[1] = vram[0 : 128*1024]
	vram = vram[128*1024:]

	mc.vram[2] = vram[0 : 128*1024]
	vram = vram[128*1024:]

	mc.vram[3] = vram[0 : 128*1024]
	vram = vram[128*1024:]

	mc.vram[4] = vram[0 : 64*1024]
	vram = vram[64*1024:]

	mc.vram[5] = vram[0 : 16*1024]
	vram = vram[16*1024:]

	mc.vram[6] = vram[0 : 16*1024]
	vram = vram[16*1024:]

	mc.vram[7] = vram[0 : 32*1024]
	vram = vram[32*1024:]

	mc.vram[8] = vram[0 : 16*1024]
	vram = vram[16*1024:]

	if len(vram) != 0 {
		panic("invalid vram size")
	}

	return mc
}

func (mc *HwMemoryController) WriteWRAMCNT(_, val uint8) {
	mc.Nds9.Bus.Unmap(0x03000000, 0x03FFFFFF)
	mc.Nds7.Bus.Unmap(0x03000000, 0x037FFFFF)

	switch val {
	case 0: // NDS9 32K - NDS7 its own wram
		mc.Nds9.Bus.MapMemorySlice(0x03000000, 0x03FFFFFF, mc.wram[:], false)
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, Emu.Mem.Wram[:], false)

	case 1: // NDS9 16K (2nd) - NDS7 16K (1st)
		mc.Nds9.Bus.MapMemorySlice(0x03000000, 0x03FFFFFF, mc.wram[16*1024:], false)
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, mc.wram[:16*1024], false)

	case 2: // NDS9 16K (1st) - NDS7 16K (2nd)
		mc.Nds9.Bus.MapMemorySlice(0x03000000, 0x03FFFFFF, mc.wram[:16*1024], false)
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, mc.wram[16*1024:], false)

	case 3: // NDS9 unmapped - NDS7 32K
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, mc.wram[:], false)

	default:
		panic("unreachable")
	}
}

func (mc *HwMemoryController) ReadWRAMSTAT(_ uint8) uint8 {
	return mc.WramCnt.Value
}

func (mc *HwMemoryController) WriteEXMEMCNT(_, val uint16) {
	// Writable by NDS9. EXMEMSTAT reflects EXMEMCNT in higher bits
	mc.ExMemStat.Value |= val & 0xFF80
}

func (mc *HwMemoryController) WriteEXMEMSTAT(_, val uint16) {
	// Writable by NDS7. Low bits are also carried over to EXMEMCNT, and since
	// there is a rwmask here (preserving the higher bits), we can just copy it
	mc.ExMemCnt.Value = mc.ExMemStat.Value
}

func (mc *HwMemoryController) mapVram9(idx byte, start uint32, end uint32) {
	modMemCnt.WithFields(log.Fields{
		"bank": string(idx),
		"addr": emu.Hex32(start),
	}).Infof("mapping VRAM on NDS9")
	idx -= 'A'
	mc.Nds9.Bus.Unmap(start, end)
	mc.Nds9.Bus.MapMemorySlice(start, end, mc.vram[idx], false)
	mc.unmapVram[idx] = func() {
		modMemCnt.WithFields(log.Fields{
			"bank":  string(idx + 'A'),
			"start": emu.Hex32(start),
			"end":   emu.Hex32(end),
		}).Warn("unmap")
		mc.Nds9.Bus.Unmap(start, end)
		mc.Nds9.Bus.MapMemorySlice(start, end, make([]byte, 4096), true)
	}
}

func (mc *HwMemoryController) mapBgExtPalette(idx byte, engIdx int, firstslot int) {
	modMemCnt.WithFields(log.Fields{
		"bank": string(idx),
		"slot": "bg-ext-palette",
	}).Infof("mapping VRAM on NDS9")
	idx -= 'A'

	var i int
	ptr := mc.vram[idx]
	for i := firstslot; i < 4 && len(ptr) > 0; i++ {
		mc.BgExtPalette[engIdx][i] = ptr[:8*1024]
		ptr = ptr[8*1024:]
	}
	lastslot := i
	mc.unmapVram[idx] = func() {
		for i := firstslot; i < lastslot; i++ {
			mc.BgExtPalette[engIdx][i] = nil
		}
	}
}

func (mc *HwMemoryController) mapObjExtPalette(idx byte, engIdx int) {
	modMemCnt.WithFields(log.Fields{
		"bank": string(idx),
		"slot": "obj-ext-palette",
	}).Infof("mapping VRAM on NDS9")
	idx -= 'A'
	mc.ObjExtPalette[engIdx] = mc.vram[idx][:8*1024]
	mc.unmapVram[idx] = func() {
		mc.ObjExtPalette[engIdx] = nil
	}
}

func (mc *HwMemoryController) mapTexturePalette(idx byte, slotnum int) {
	modMemCnt.WithFields(log.Fields{
		"bank": string(idx),
		"slot": "texture-palette",
	}).Infof("mapping VRAM on NDS9")
	idx -= 'A'
	mc.TexturePalette[slotnum] = mc.vram[idx][:16*1024]
	mc.unmapVram[idx] = func() {
		mc.TexturePalette[slotnum] = nil
	}
}

func (mc *HwMemoryController) writeVramCnt(idx byte, val uint8) (int, int) {
	idx -= 'A'
	if mc.unmapVram[idx] != nil {
		mc.unmapVram[idx]()
		mc.unmapVram[idx] = nil
	}
	if val&0x80 == 0 {
		return -1, -1
	}
	return int(val & 7), int((val >> 3) & 3)
}

func (mc *HwMemoryController) WriteVRAMCNTA(_, val uint8) {
	mst, ofs := mc.writeVramCnt('A', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('A', 0x6800000, 0x681FFFF)
	case 1:
		base := 0x6000000 + uint32(ofs)*0x20000
		mc.mapVram9('A', base, base+0x1FFFF)
	case 2:
		base := 0x6400000 + uint32(ofs)*0x20000
		mc.mapVram9('A', base, base+0x1FFFF)
	case 3:
		// Texture slot....
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "A",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTB(_, val uint8) {
	mst, ofs := mc.writeVramCnt('B', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('B', 0x6820000, 0x683FFFF)
	case 1:
		base := 0x6000000 + uint32(ofs)*0x20000
		mc.mapVram9('B', base, base+0x1FFFF)
	case 2:
		base := 0x6400000 + uint32(ofs)*0x20000
		mc.mapVram9('B', base, base+0x1FFFF)
	case 3:
		// Texture slot....
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "B",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTC(_, val uint8) {
	mst, ofs := mc.writeVramCnt('C', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('C', 0x6840000, 0x685FFFF)
	case 1:
		base := 0x6000000 + uint32(ofs)*0x20000
		mc.mapVram9('C', base, base+0x1FFFF)
	case 3:
		// texture slot...
	case 4:
		mc.mapVram9('C', 0x6200000, 0x621FFFF)
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "C",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTD(_, val uint8) {
	mst, ofs := mc.writeVramCnt('D', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('D', 0x6860000, 0x687FFFF)
	case 1:
		base := 0x6000000 + uint32(ofs)*0x20000
		mc.mapVram9('D', base, base+0x1FFFF)
	case 3:
		// texture slot...
	case 4:
		mc.mapVram9('D', 0x6600000, 0x661FFFF)
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "D",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTE(_, val uint8) {
	mst, ofs := mc.writeVramCnt('E', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('E', 0x6880000, 0x688FFFF)
	case 1:
		mc.mapVram9('E', 0x6000000, 0x600FFFF)
	case 2:
		mc.mapVram9('E', 0x6400000, 0x640FFFF)
	case 3:
		// texture slot...
	case 4:
		mc.mapBgExtPalette('E', 0, 0)
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "E",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTF(_, val uint8) {
	mst, ofs := mc.writeVramCnt('F', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('F', 0x6890000, 0x6893FFF)
	case 1:
		off := uint32(0x4000*(ofs&1) + 0x8000*(ofs&2))
		mc.mapVram9('F', 0x6000000+off, 0x6003FFF+off)
	case 2:
		off := uint32((ofs&1)*0x4000 + (ofs&2)*0x8000)
		mc.mapVram9('F', 0x6400000+off, 0x6403FFF+off)
	case 3:
		off := int((ofs&1)*1 + (ofs&2)*2)
		mc.mapTexturePalette('F', off)
	case 4:
		mc.mapBgExtPalette('F', 0, ofs*2)
	case 5:
		mc.mapObjExtPalette('F', 0)
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "F",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTG(_, val uint8) {
	mst, ofs := mc.writeVramCnt('G', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('G', 0x6894000, 0x6897FFF)
	case 3:
		// Texture slot....
	case 4:
		mc.mapBgExtPalette('G', 0, ofs*2)
	case 5:
		mc.mapObjExtPalette('G', 0)
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "G",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTH(_, val uint8) {
	mst, ofs := mc.writeVramCnt('H', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('H', 0x6898000, 0x689FFFF)
	case 1:
		mc.mapVram9('H', 0x6200000, 0x6207FFF)
	case 2:
		mc.mapBgExtPalette('H', 1, 0)
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "H",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTI(_, val uint8) {
	mst, ofs := mc.writeVramCnt('I', val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9('I', 0x68A0000, 0x68A3FFF)
	case 1:
		mc.mapVram9('I', 0x6208000, 0x620BFFF)
	case 2:
		mc.mapVram9('I', 0x6600000, 0x6603FFF)
	case 3:
		mc.mapObjExtPalette('I', 1)
	default:
		modMemCnt.WithFields(log.Fields{
			"bank": "I",
			"mst":  mst,
			"ofs":  ofs,
		}).Fatal("invalid vram configuration")
	}
}

const vramSmallestBankSize = 8 * 1024

var empty [vramSmallestBankSize]byte

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
	ptr [32][]uint8
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

// Return the VRAM linear bank that will be accessed by the specified engine.
// The linear banks is 256k big, and can be accessed as 8-bit or 16-bit.
// byteOffset is the offset within the VRAM from which the 256k bank starts.
//
// If the requested bank is unmapped, a zero-filled area is returned. If the
// requested bank is mapped for less than 256K, the missing areas will be
// zero-filled as well.
func (mc *HwMemoryController) VramLinearBank(engine int, which VramLinearBankId, baseOffset int) (vb VramLinearBank) {
	for i := 0; i < 32; i++ {
		var ptr []byte

		switch which {
		case VramLinearBGExtPal:
			if i < len(mc.BgExtPalette[engine]) {
				ptr = mc.BgExtPalette[engine][i]
			}
		case VramLinearOBJExtPal:
			if i == 0 {
				ptr = mc.ObjExtPalette[engine]
			}
		case VramLinearBG:
			ptr = mc.Nds9.Bus.FetchPointer(uint32(0x6000000 + 0x200000*engine + baseOffset + i*vramSmallestBankSize))
		case VramLinearOAM:
			ptr = mc.Nds9.Bus.FetchPointer(uint32(0x6400000 + 0x200000*engine + baseOffset + i*vramSmallestBankSize))
		default:
			panic("unreachable")
		}

		vb.ptr[i] = ptr
		if vb.ptr[i] == nil {
			vb.ptr[i] = empty[:]
		}
	}
	return
}

func (vb *VramLinearBank) Dump(w io.Writer) {
	for i := range vb.ptr {
		w.Write(vb.ptr[i][:vramSmallestBankSize])
	}
}

func (vb *VramLinearBank) FetchPointer(off int) []uint8 {
	bank := vb.ptr[off/vramSmallestBankSize]
	off &= (vramSmallestBankSize - 1)
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
