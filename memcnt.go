package main

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwMemoryController struct {
	Nds9 *NDS9
	Nds7 *NDS7

	vramA   [128 * 1024]byte
	vramB   [128 * 1024]byte
	vramC   [128 * 1024]byte
	vramD   [128 * 1024]byte
	vramE   [64 * 1024]byte
	vramF   [16 * 1024]byte
	vramG   [16 * 1024]byte
	vramH   [32 * 1024]byte
	vramI   [16 * 1024]byte
	wram    [32 * 1024]byte
	wramcnt uint8

	vram      [9][]byte
	vramcnt   [9]uint8
	unmapVram [9]func()
}

func NewMemoryController(nds9 *NDS9, nds7 *NDS7) *HwMemoryController {
	mc := &HwMemoryController{
		Nds9: nds9,
		Nds7: nds7,
	}

	mc.vram[0] = mc.vramA[:]
	mc.vram[1] = mc.vramB[:]
	mc.vram[2] = mc.vramC[:]
	mc.vram[3] = mc.vramD[:]
	mc.vram[4] = mc.vramE[:]
	mc.vram[5] = mc.vramF[:]
	mc.vram[6] = mc.vramG[:]
	mc.vram[7] = mc.vramH[:]
	mc.vram[8] = mc.vramI[:]

	return mc
}

func (mc *HwMemoryController) WriteWRAMCNT(val uint8) {
	mc.wramcnt = val

	mc.Nds9.Bus.Unmap(0x03000000, 0x03FFFFFF)
	mc.Nds7.Bus.Unmap(0x03000000, 0x037FFFFF)

	switch mc.wramcnt & 3 {
	case 0: // NDS9 32K - NDS7 its own wram
		mc.Nds9.Bus.MapMemorySlice(0x03000000, 0x03FFFFFF, mc.wram[:], false)
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, mc.Nds7.WRam[:], false)

	case 1: // NDS9 16K (2nd) - NDS7 16K (1st)
		mc.Nds9.Bus.MapMemorySlice(0x03000000, 0x03FFFFFF, mc.wram[16*1024:], false)
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, mc.wram[:16*1024], false)

	case 2: // NDS9 16K (1st) - NDS7 16K (2nd)
		mc.Nds9.Bus.MapMemorySlice(0x03000000, 0x03FFFFFF, mc.wram[:16*1024], false)
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, mc.wram[16*1024:], false)

	case 3: // NDS9 unmapped - NDS7 32K
		mc.Nds7.Bus.MapMemorySlice(0x03000000, 0x037FFFFF, mc.wram[:], false)
	}
}

func (mc *HwMemoryController) ReadWRAMCNT() uint8 {
	return mc.wramcnt
}

func (mc *HwMemoryController) mapVram9(idx int, start uint32, end uint32) {
	mc.Nds9.Bus.MapMemorySlice(start, end, mc.vram[idx], false)
	mc.unmapVram[idx] = func() {
		mc.Nds9.Bus.Unmap(start, end)
	}
}

func (mc *HwMemoryController) writeVramCnt(idx int, val uint8) (int, int) {
	mc.vramcnt[idx] = val
	if mc.unmapVram[idx] != nil {
		mc.unmapVram[idx]()
		mc.unmapVram[idx] = nil
	}
	if val&0x80 == 0 {
		return -1, -1
	}
	return int(val & 7), int((val >> 3) & 3)
}

func (mc *HwMemoryController) WriteVRAMCNTA(val uint8) {
	mst, ofs := mc.writeVramCnt(0, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(0, 0x6800000, 0x681FFFF)
	case 1:
		base := 0x6000000 + uint32(ofs)*0x20000
		mc.mapVram9(0, base, base+0x1FFFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'A',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTB(val uint8) {
	mst, ofs := mc.writeVramCnt(1, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(1, 0x6820000, 0x683FFFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'B',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTC(val uint8) {
	mst, ofs := mc.writeVramCnt(2, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(2, 0x6840000, 0x685FFFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'C',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTD(val uint8) {
	mst, ofs := mc.writeVramCnt(3, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(3, 0x6860000, 0x687FFFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'D',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTE(val uint8) {
	mst, ofs := mc.writeVramCnt(4, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(4, 0x6880000, 0x688FFFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'E',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTF(val uint8) {
	mst, ofs := mc.writeVramCnt(5, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(5, 0x6890000, 0x6893FFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'F',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTG(val uint8) {
	mst, ofs := mc.writeVramCnt(6, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(6, 0x6894000, 0x6897FFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'G',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTH(val uint8) {
	mst, ofs := mc.writeVramCnt(7, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(7, 0x6898000, 0x689FFFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'H',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}

func (mc *HwMemoryController) WriteVRAMCNTI(val uint8) {
	mst, ofs := mc.writeVramCnt(8, val)
	switch mst {
	case -1:
		return
	case 0:
		mc.mapVram9(8, 0x68A0000, 0x68A3FFF)
	default:
		log.WithFields(log.Fields{
			"bank": 'I',
			"mst":  mst,
			"ofs":  ofs,
		}).Warn("invalid vram configuration")
	}
}
