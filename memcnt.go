package main

import (
	"fmt"
	"ndsemu/e2d"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"ndsemu/raster3d"
)

var modMemCnt = log.NewModule("memcnt")

// Memory controller VRAM mapping
// ******************************
//
// VRAM is made of several banks (of different sizes), that can be mapped at different
// addresses (called "areas"), in different address spaces. Depending on the area where
// each bank is mapped, it can be accessed by either ARM9, ARM7 or the GPU. Some areas,
// in fact, are only accessible by the CPU and when a bank is mapped there, it's not
// accessible anymore by the ARMs.
//
// Each bank can be mapped to a single area at a time, but each area can potentially
// be mapped by multiple banks. In the normal case, the banks are mapped at different
// offsets within each area (called "slots"), so nothing weird happen, but overlapping
// banks is indeed possible.
//
// Overlapping banks
// *****************
//
// In case of overlapping banks (that is, multiple banks mapped to the same slot in
// the same area), writes go to all the mapped banks, and reads contain
// the OR of the value of each bank (this is basically what happens when the parallel
// bus is connected to different DRAM chips at the same time). This is currently
// not emulated.
//
// To emulate this special mapping logic, we can't use the normal Map/Unmap functions
// as exported by hwio.Table, because Table assumes that there can be a single object
// mapped to each memory address. Even if overlapping banks are not supported, they
// must be somehow accounted for handling sequences like this:
//
//    Bank A mapped to 6200000
//    Bank B mapped to 6200000 (overlapping, but no R/W is performed)
//    Bank A mapped to 6400000
//
// or even weirder:
//
//    Bank A mapped to 6200000
//    Bank B mapped to 6200000 (overlapping, but no R/W is performed)
//    Bank B mapped to 6000000 (now A is still available at 6200000)
//
// On the last line, if we run a blanket Unmap() of whatever is at 6200000, we
// would unmap the B bank. hwio.Table currently doesn't expose an API like
// "unmap this specific memory slice at this address, if present", so we
// currently punt on unmapping.
//
// GPU mapping addresses
// *********************
//
// Some areas are visible only to the GPU. We implement this by creating a new
// bus (hwio.Table) called GpuBus, and we map VRAM banks into it, at fixed addresses
// that we invented. I assume that this is the way it works on the real hardware too:
// the GPU would do a memory access on a special bus (through the memory controller)
// and would select which area through a part of the address bits.
//

type vramAreaIdx int

const (
	vramAreaLcdc vramAreaIdx = iota
	vramAreaBgA
	vramAreaObjA
	vramAreaBgExtPalA
	vramAreaObjExtPalA
	vramAreaTexture
	vramAreaTexturePal
	vramAreaBgB
	vramAreaObjB
	vramAreaBgExtPalB
	vramAreaObjExtPalB
	vramAreaArm7
	vramAreaCount
	vramAreaInvalid
)

var vramAreaNames = [...]string{
	"lcdc",
	"bg-a", "obj-a", "bgxpal-a", "objxpal-a",
	"tex", "texpal",
	"bg-b", "obj-b", "bgxpal-b", "objxpal-b",
	"arm7",
	"invalid", "invalid",
}

// vramSlot is a single slot within an area. We implement it as a hwio.Mem
// instance, and we dynamically change the hwio.Mem.Data field anytime a new
// bank is mapped.
type vramSlot struct {
	hwio.Mem
	maps [9][]byte // memory buffers mapped to this slot (potentially, one per bank)
	cnt  uint8     // number of banks (buffers) currently mapped to this slot
}

// vramArea represents a single VRAM area
type vramArea struct {
	bus      *hwio.Table // Pointer to the bus that accesses this area (ARM9, ARM7 or GPU)
	addr     uint32      // Base address of this area (within the bus)
	slots    []vramSlot  // Slots this area is composed of
	slotSize uint32      // Size of each slot (16K for now)
}

type HwMemoryController struct {
	Nds9 *NDS9
	Nds7 *NDS7

	// Special GPU bus in which VRAM can be mapped.
	// This is a bus that is used only by GPU (not by CPUs) to access some special
	// memory areas. Memcnt can map memory banks here for GPU usage.
	GpuBus *hwio.Table

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

	// VRAM banks
	vram [9][]byte

	// Areas of VRAM
	vramAreas [vramAreaCount]vramArea

	// Area where each bank is currently mapped to (or vramAreaInvalid otherwise).
	// This is a redundant cache that is used to quickly unmap a bank from its
	// previous area when a new mapping is performed.
	curBankArea [9]vramAreaIdx
}

var zero [16 * 1024]byte

var vramBankMappingDesc = [9][8]struct {
	Area vramAreaIdx
	Base uint32
	Off0 uint32
	Off1 uint32
}{
	'A' - 'A': {
		0: {vramAreaLcdc, 0x6800000, 0, 0},
		1: {vramAreaBgA, 0x6000000, 0x20000, 0x20000 * 2},
		2: {vramAreaObjA, 0x6400000, 0x20000, 0x20000 * 2},
		3: {vramAreaTexture, 0x5000000, 0x20000, 0x20000 * 2},
	},
	'B' - 'A': {
		0: {vramAreaLcdc, 0x6820000, 0, 0},
		1: {vramAreaBgA, 0x6000000, 0x20000, 0x20000 * 2},
		2: {vramAreaObjA, 0x6400000, 0x20000, 0x20000 * 2},
		3: {vramAreaTexture, 0x5000000, 0x20000, 0x20000 * 2},
	},
	'C' - 'A': {
		0: {vramAreaLcdc, 0x6840000, 0, 0},
		1: {vramAreaBgA, 0x6000000, 0x20000, 0x20000 * 2},
		2: {vramAreaObjA, 0x6400000, 0x20000, 0x20000 * 2},
		3: {vramAreaTexture, 0x5000000, 0x20000, 0x20000 * 2},
		4: {vramAreaBgB, 0x6200000, 0, 0},
	},
	'D' - 'A': {
		0: {vramAreaLcdc, 0x6860000, 0, 0},
		1: {vramAreaBgA, 0x6000000, 0x20000, 0x20000 * 2},
		2: {vramAreaArm7, 0x6000000, 0x20000, 0},
		3: {vramAreaTexture, 0x5000000, 0x20000, 0x20000 * 2},
		4: {vramAreaObjB, 0x6600000, 0, 0},
	},
	'E' - 'A': {
		0: {vramAreaLcdc, 0x6880000, 0, 0},
		1: {vramAreaBgA, 0x6000000, 0, 0},
		2: {vramAreaObjA, 0x6400000, 0, 0},
		3: {vramAreaTexturePal, 0x6000000, 0, 0},
		4: {vramAreaBgExtPalA, 0x1000000, 0, 0},
	},
	'F' - 'A': {
		0: {vramAreaLcdc, 0x6890000, 0, 0},
		1: {vramAreaBgA, 0x6000000, 0x4000, 0x10000},
		2: {vramAreaObjA, 0x6400000, 0x4000, 0x10000},
		3: {vramAreaTexturePal, 0x6000000, 0x4000, 0x10000},
		4: {vramAreaBgExtPalA, 0x1000000, 0x4000, 0},
		5: {vramAreaObjExtPalA, 0x2000000, 0x0, 0x0},
	},
	'G' - 'A': {
		0: {vramAreaLcdc, 0x6894000, 0, 0},
		1: {vramAreaBgA, 0x6000000, 0x4000, 0x10000},
		2: {vramAreaObjA, 0x6400000, 0x4000, 0x10000},
		3: {vramAreaTexturePal, 0x6000000, 0x4000, 0x10000},
		4: {vramAreaBgExtPalA, 0x1000000, 0x4000, 0},
		5: {vramAreaObjExtPalA, 0x2000000, 0x0, 0x0},
	},
	'H' - 'A': {
		0: {vramAreaLcdc, 0x6898000, 0, 0},
		1: {vramAreaBgB, 0x6200000, 0, 0},
		2: {vramAreaBgExtPalB, 0x3000000, 0, 0},
	},
	'I' - 'A': {
		0: {vramAreaLcdc, 0x68A0000, 0, 0},
		1: {vramAreaBgB, 0x6208000, 0, 0},
		2: {vramAreaObjB, 0x6600000, 0, 0},
		3: {vramAreaObjExtPalB, 0x4000000, 0, 0},
	},
}

func NewMemoryController(nds9 *NDS9, nds7 *NDS7, vram []byte) *HwMemoryController {
	mc := &HwMemoryController{
		Nds9:   nds9,
		Nds7:   nds7,
		GpuBus: hwio.NewTable("gpubus"),
	}
	hwio.MustInitRegs(mc)

	// Setup VRAM areas
	mc.vramAreas = [...]vramArea{
		vramAreaBgA:        newVramArea("VRAM-A-BG", nds9.Bus, 0x6000000, 0x607FFFF, 16*1024),
		vramAreaBgB:        newVramArea("VRAM-B-BG", nds9.Bus, 0x6200000, 0x621FFFF, 16*1024),
		vramAreaObjA:       newVramArea("VRAM-A-OBJ", nds9.Bus, 0x6400000, 0x645FFFF, 16*1024),
		vramAreaObjB:       newVramArea("VRAM-B-OBJ", nds9.Bus, 0x6600000, 0x661FFFF, 16*1024),
		vramAreaLcdc:       newVramArea("VRAM-LCDC", nds9.Bus, 0x6800000, 0x68A3FFF, 16*1024),
		vramAreaArm7:       newVramArea("VRAM-ARM7", nds7.Bus, 0x6000000, 0x603FFFF, 16*1024),
		vramAreaBgExtPalA:  newVramArea("VRAM-A-BGXPAL", mc.GpuBus, 0x1000000, 0x100FFFF, 16*1024),
		vramAreaObjExtPalA: newVramArea("VRAM-A-OBJXPAL", mc.GpuBus, 0x2000000, 0x2003FFF, 16*1024),
		vramAreaBgExtPalB:  newVramArea("VRAM-B-BGXPAL", mc.GpuBus, 0x3000000, 0x3007FFF, 16*1024),
		vramAreaObjExtPalB: newVramArea("VRAM-B-OBJXPAL", mc.GpuBus, 0x4000000, 0x4003FFF, 16*1024),
		vramAreaTexture:    newVramArea("VRAM-TEX", mc.GpuBus, 0x5000000, 0x507FFFF, 16*1024),
		vramAreaTexturePal: newVramArea("VRAM-TEXPAL", mc.GpuBus, 0x6000000, 0x6017FFF, 16*1024),
	}

	for i := range mc.curBankArea {
		mc.curBankArea[i] = vramAreaInvalid
	}

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

func (mc *HwMemoryController) WriteEXMEMCNT(old, val uint16) {
	// Writable by NDS9. EXMEMSTAT reflects EXMEMCNT in higher bits
	mc.ExMemStat.Value |= val & 0xFF80

	// Bit 11 changed: gamecard nds9/nds7 mapping
	if (old^val)&(1<<11) != 0 {
		if val&(1<<11) != 0 {
			nds9.Bus.UnmapBank(0x40001A0, Emu.Hw.Gc, 0)
			nds9.Bus.UnmapBank(0x4100010, Emu.Hw.Gc, 1)
			nds7.Bus.MapBank(0x40001A0, Emu.Hw.Gc, 0)
			nds7.Bus.MapBank(0x4100010, Emu.Hw.Gc, 1)
			Emu.Hw.Gc.Irq = nds7.Irq
			modMemCnt.InfoZ("mapped gamecard to NDS7").End()
		} else {
			nds7.Bus.UnmapBank(0x40001A0, Emu.Hw.Gc, 0)
			nds7.Bus.UnmapBank(0x4100010, Emu.Hw.Gc, 1)
			nds9.Bus.MapBank(0x40001A0, Emu.Hw.Gc, 0)
			nds9.Bus.MapBank(0x4100010, Emu.Hw.Gc, 1)
			Emu.Hw.Gc.Irq = nds9.Irq
			modMemCnt.InfoZ("mapped gamecard to NDS9").End()
		}
	}

	// Bit 7 changed: GBA slot nds9/nds7 mapping
	if (old^val)&(1<<7) != 0 {
		if val&(1<<7) != 0 {
			// GBA slot mapped to NDS7. Since we don't emulate it yet, when
			// there is no card in the slot, 0xFF is returned
			nds7.Bus.Unmap(0x8000000, 0xAFFFFFF)
			nds7.Bus.MapMemorySlice(0x8000000, 0x9FFFFFF, Emu.Hw.Sl2.Rom[:], true)
			nds7.Bus.MapMemorySlice(0xA000000, 0xAFFFFFF, Emu.Hw.Sl2.Ram[:], false)

			// NDS9 sees a zero-filled region
			nds9.Bus.Unmap(0x8000000, 0xAFFFFFF)
			nds9.Bus.MapMemorySlice(0x8000000, 0xAFFFFFF, zero[:], true)
		} else {
			// GBA slot mapped to NDS9. Same as above, reversing roles
			nds9.Bus.Unmap(0x8000000, 0xAFFFFFF)
			nds9.Bus.MapMemorySlice(0x8000000, 0x9FFFFFF, Emu.Hw.Sl2.Rom[:], true)
			nds9.Bus.MapMemorySlice(0xA000000, 0xAFFFFFF, Emu.Hw.Sl2.Ram[:], false)

			nds7.Bus.Unmap(0x8000000, 0xAFFFFFF)
			nds7.Bus.MapMemorySlice(0x8000000, 0xAFFFFFF, zero[:], true)
		}
	}
}

func (mc *HwMemoryController) WriteEXMEMSTAT(_, val uint16) {
	// Writable by NDS7. Low bits are also carried over to EXMEMCNT, and since
	// there is a rwmask here (preserving the higher bits), we can just copy it
	mc.ExMemCnt.Value = mc.ExMemStat.Value
}

func newVramArea(name string, bus *hwio.Table, begin, end uint32, slotSize uint32) vramArea {
	end += 1
	if (end-begin)%slotSize != 0 {
		panic("slot size is not a multiple of mapped range")
	}

	numSlots := (end - begin) / slotSize

	a := vramArea{
		bus:      bus,
		slots:    make([]vramSlot, numSlots),
		slotSize: slotSize,
		addr:     begin,
	}
	for i := range a.slots {
		mem := &a.slots[i]
		mem.Name = fmt.Sprintf("%s%02d", name, i)
		mem.Data = zero[:]
		mem.VSize = int(slotSize)
		mem.Flags = hwio.MemFlagReadOnly | hwio.MemFlag16Unaligned | hwio.MemFlag32Unaligned
		bus.MapMem(begin+uint32(i)*slotSize, &mem.Mem)
	}
	return a
}

func (a *vramArea) Map(addr uint32, bank byte, mem []byte) {
	if (addr-a.addr)%a.slotSize != 0 {
		panic("mapping not aligned with slots")
	}
	if (addr-a.addr)/a.slotSize >= uint32(len(a.slots)) {
		panic("invalid mapping address")
	}

	// Unmap all the slots affected by this new mapping.
	// This is necessary because we need to call MapMem() everytime
	// we modify the hwio.Mem instance within each slot; this is the
	// way the API of hwio.Mem works.
	a.bus.Unmap(addr, addr+uint32(len(mem))-1)

	sidx := (addr - a.addr) / a.slotSize
	for len(mem) > 0 {
		s := &a.slots[sidx]
		if s.maps[bank] != nil {
			panic("double mapping")
		}
		s.maps[bank] = mem[:a.slotSize:a.slotSize]
		s.cnt++
		if s.cnt == 1 {
			s.Mem.Data = s.maps[bank]
			s.Mem.Flags &^= hwio.MemFlagReadOnly
		} else {
			// FIXME: implement handling of more than one mapping
			s.Mem.Data = zero[:]
			s.Mem.Flags |= hwio.MemFlagReadOnly
			modMemCnt.WarnZ("VRAM double mapping").String("bank", string(bank+'A')).Hex32("addr", addr).End()
		}

		a.bus.MapMem(addr, &s.Mem)
		mem = mem[a.slotSize:]
		sidx++
		addr += a.slotSize
	}
}

func (a *vramArea) Unmap(bank byte) {
	for sidx := range a.slots {
		s := &a.slots[sidx]

		if s.maps[bank] != nil {
			s.maps[bank] = nil
			s.cnt--
			switch s.cnt {
			case 0:
				s.Mem.Data = zero[:]
				s.Mem.Flags |= hwio.MemFlagReadOnly
			case 1:
				for b := range s.maps {
					if s.maps[b] != nil {
						s.Mem.Data = s.maps[b]
						break
					}
				}
			default:
				// FIXME: implement handling of more than one mapping
				s.Mem.Data = zero[:]
				s.Mem.Flags |= hwio.MemFlagReadOnly
			}

			// Apply the new mapping. We need to call MapMem() any time
			// we modify the hwio.Mem instance.
			addr := a.addr + uint32(sidx)*a.slotSize
			a.bus.Unmap(addr, addr+a.slotSize-1)
			a.bus.MapMem(addr, &s.Mem)
		}
	}
}

func (mc *HwMemoryController) writeVRAMCNT(bank byte, val uint8) {
	bank -= 'A'

	// First unmap the bank from its current area (if any)
	if mc.curBankArea[bank] != vramAreaInvalid {
		mc.vramAreas[mc.curBankArea[bank]].Unmap(bank)
		mc.curBankArea[bank] = vramAreaInvalid
	}

	// If we're asked to disable the bank, we're done
	if val&0x80 == 0 {
		return
	}

	// Extract requested mapping description
	mst, off := int(val&7), int((val>>3)&3)
	desc := &vramBankMappingDesc[bank][mst]

	if desc.Base == 0 {
		modMemCnt.ErrorZ("Unimplemented VRAM mapping").String("bank", string(bank+'A')).Int("mst", mst).Int("off", off).End()
		return
	}

	// Compute final mapping address
	addr := desc.Base
	if off&1 != 0 {
		addr += desc.Off0
	}
	if off&2 != 0 {
		addr += desc.Off1
	}

	modMemCnt.InfoZ("mapping VRAM").
		String("bank", string(bank+'A')).Int("mst", mst).Int("off", off).
		String("area", vramAreaNames[desc.Area]).
		Hex32("addr", addr).
		End()

	// Do the mapping (and remember it)
	mc.vramAreas[desc.Area].Map(addr, bank, mc.vram[bank])
	mc.curBankArea[bank] = desc.Area
}

func (mc *HwMemoryController) WriteVRAMCNTA(_, val uint8) { mc.writeVRAMCNT('A', val) }
func (mc *HwMemoryController) WriteVRAMCNTB(_, val uint8) { mc.writeVRAMCNT('B', val) }
func (mc *HwMemoryController) WriteVRAMCNTC(_, val uint8) { mc.writeVRAMCNT('C', val) }
func (mc *HwMemoryController) WriteVRAMCNTD(_, val uint8) { mc.writeVRAMCNT('D', val) }
func (mc *HwMemoryController) WriteVRAMCNTE(_, val uint8) { mc.writeVRAMCNT('E', val) }
func (mc *HwMemoryController) WriteVRAMCNTF(_, val uint8) { mc.writeVRAMCNT('F', val) }
func (mc *HwMemoryController) WriteVRAMCNTG(_, val uint8) { mc.writeVRAMCNT('G', val) }
func (mc *HwMemoryController) WriteVRAMCNTH(_, val uint8) { mc.writeVRAMCNT('H', val) }
func (mc *HwMemoryController) WriteVRAMCNTI(_, val uint8) { mc.writeVRAMCNT('I', val) }

/********************************************
 * Engine2D VRAM
 ********************************************/

var empty [e2d.VramSmallestBankSize]byte

var linearBankAddr = [4]struct {
	bus   string
	addrs [2]uint32
}{
	e2d.VramLinearBGExtPal:  {"GPU", [2]uint32{0x1000000, 0x3000000}},
	e2d.VramLinearOBJExtPal: {"GPU", [2]uint32{0x2000000, 0x4000000}},
	e2d.VramLinearBG:        {"ARM9", [2]uint32{0x6000000, 0x6200000}},
	e2d.VramLinearOAM:       {"ARM9", [2]uint32{0x6400000, 0x6600000}},
}

// Return the VRAM linear bank that will be accessed by the specified engine.
// The linear bank is 256k big, and can be accessed as 8-bit or 16-bit.
// byteOffset is the offset within the VRAM from which the 256k bank starts.
//
// If the requested bank is unmapped, a zero-filled area is returned. If the
// requested bank is mapped for less than 256K, the missing areas will be
// zero-filled as well.
func (mc *HwMemoryController) VramLinearBank(engine int, which e2d.VramLinearBankId, baseOffset int) (vb e2d.VramLinearBank) {
	bus := mc.Nds9.Bus
	if linearBankAddr[which].bus == "GPU" {
		bus = mc.GpuBus
	}
	addr := linearBankAddr[which].addrs[engine] + uint32(baseOffset)

	for i := 0; i < 32; i++ {
		vb.Ptr[i] = bus.FetchPointer(addr)
		// vb.Ptr[i] = vb.Ptr[i][:e2d.VramSmallestBankSize:e2d.VramSmallestBankSize]
		if vb.Ptr[i] == nil {
			vb.Ptr[i] = empty[:]
		}
		addr += e2d.VramSmallestBankSize
	}
	return
}

func (mc *HwMemoryController) VramPalette(engine int) []byte {
	return Emu.Mem.PaletteRam[engine*1024 : engine*1024+1024]
}

func (mc *HwMemoryController) VramOAM(engine int) []byte {
	return Emu.Mem.OamRam[0x400*engine : 0x400+0x400*engine]
}

func (mc *HwMemoryController) VramRawBank(bank int) []byte {
	return mc.vram[bank]
}

/********************************************
 * Raster3D VRAM
 ********************************************/

func (mc *HwMemoryController) VramTextureBank() (tb raster3d.VramTextureBank) {
	for i := range tb.Slots {
		tb.Slots[i] = mc.GpuBus.FetchPointer(0x5000000 + uint32(i*16*1024))
	}
	return
}

func (mc *HwMemoryController) VramTexturePaletteBank() (tpb raster3d.VramTexturePaletteBank) {
	for i := range tpb.Slots {
		tpb.Slots[i] = mc.GpuBus.FetchPointer(0x6000000 + uint32(i*16*1024))
	}
	return
}
