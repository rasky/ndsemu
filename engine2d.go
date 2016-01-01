package main

import (
	"encoding/binary"
	"ndsemu/emu"
	"ndsemu/emu/hwio"

	log "gopkg.in/Sirupsen/logrus.v0"
)

const (
	cScreenWidth  = 256
	cScreenHeight = 192
)

func le16(data []byte) uint16 {
	return binary.LittleEndian.Uint16(data)
}

type HwEngine2d struct {
	Idx      int
	DispCnt  hwio.Reg32 `hwio:"offset=0x00,wcb"`
	Bg0Cnt   hwio.Reg16 `hwio:"offset=0x08,wcb"`
	Bg1Cnt   hwio.Reg16 `hwio:"offset=0x0A"`
	Bg2Cnt   hwio.Reg16 `hwio:"offset=0x0C"`
	Bg3Cnt   hwio.Reg16 `hwio:"offset=0x0E"`
	Bg0XOfs  hwio.Reg16 `hwio:"offset=0x10,writeonly"`
	Bg0YOfs  hwio.Reg16 `hwio:"offset=0x12,writeonly"`
	Bg1XOfs  hwio.Reg16 `hwio:"offset=0x14,writeonly"`
	Bg1YOfs  hwio.Reg16 `hwio:"offset=0x16,writeonly"`
	Bg2XOfs  hwio.Reg16 `hwio:"offset=0x18,writeonly"`
	Bg2YOfs  hwio.Reg16 `hwio:"offset=0x1A,writeonly"`
	Bg3XOfs  hwio.Reg16 `hwio:"offset=0x1C,writeonly"`
	Bg3YOfs  hwio.Reg16 `hwio:"offset=0x1E,writeonly"`
	Bg2PA    hwio.Reg16 `hwio:"offset=0x20,writeonly"`
	Bg2PB    hwio.Reg16 `hwio:"offset=0x22,writeonly"`
	Bg2PC    hwio.Reg16 `hwio:"offset=0x24,writeonly"`
	Bg2PD    hwio.Reg16 `hwio:"offset=0x26,writeonly"`
	Bg2PX    hwio.Reg32 `hwio:"offset=0x28,writeonly"`
	Bg2PY    hwio.Reg32 `hwio:"offset=0x2C,writeonly"`
	Bg3PA    hwio.Reg16 `hwio:"offset=0x30,writeonly"`
	Bg3PB    hwio.Reg16 `hwio:"offset=0x32,writeonly"`
	Bg3PC    hwio.Reg16 `hwio:"offset=0x34,writeonly"`
	Bg3PD    hwio.Reg16 `hwio:"offset=0x36,writeonly"`
	Bg3PX    hwio.Reg32 `hwio:"offset=0x38,writeonly"`
	Bg3PY    hwio.Reg32 `hwio:"offset=0x3C,writeonly"`
	Win0X    hwio.Reg16 `hwio:"offset=0x40,writeonly"`
	Win1X    hwio.Reg16 `hwio:"offset=0x42,writeonly"`
	Win0Y    hwio.Reg16 `hwio:"offset=0x44,writeonly"`
	Win1Y    hwio.Reg16 `hwio:"offset=0x46,writeonly"`
	WinIn    hwio.Reg16 `hwio:"offset=0x48"`
	WinOut   hwio.Reg16 `hwio:"offset=0x4A"`
	Mosaic   hwio.Reg16 `hwio:"offset=0x4C,writeonly"`
	BldCnt   hwio.Reg16 `hwio:"offset=0x50"`
	BldAlpha hwio.Reg16 `hwio:"offset=0x52"`
	BldY     hwio.Reg32 `hwio:"offset=0x54"`

	vram    []byte
	lineBuf [4 * (cScreenWidth + 16)]byte
}

func NewHwEngine2d(idx int, vram []byte) *HwEngine2d {
	e2d := new(HwEngine2d)
	hwio.MustInitRegs(e2d)
	e2d.Idx = idx
	e2d.vram = vram
	return e2d
}

func (e2d *HwEngine2d) A() bool { return e2d.Idx == 0 }
func (e2d *HwEngine2d) B() bool { return e2d.Idx != 0 }

func (e2d *HwEngine2d) WriteDISPCNT(old, val uint32) {
	log.WithFields(log.Fields{
		"name": string('A' + e2d.Idx),
		"val":  emu.Hex32(val),
	}).Info("[lcd] write dispcnt")
}

func (e2d *HwEngine2d) drawChar16(y int, src []byte, dst []byte, hflip bool) {
	src = src[y*4:]

	if !hflip {
		for x := 0; x < 4; x++ {
			p0, p1 := src[x]&0xF, src[x]>>4
			if p0 != 0 {
				p0 = (p0 << 4) | p0
				dst[0] = p0
				dst[1] = p0
				dst[2] = p0
				dst[3] = p0
			}

			if p1 != 0 {
				p1 = (p1 << 4) | p1
				dst[4] = p1
				dst[5] = p1
				dst[6] = p1
				dst[7] = p1
			}

			dst = dst[8:]
		}
	} else {
		for x := 3; x >= 0; x-- {
			p1, p0 := src[x]&0xF, src[x]>>4
			if p0 != 0 {
				p0 = (p0 << 4) | p0
				dst[0] = p0
				dst[1] = p0
				dst[2] = p0
				dst[3] = p0
			}

			if p1 != 0 {
				p1 = (p1 << 4) | p1
				dst[4] = p1
				dst[5] = p1
				dst[6] = p1
				dst[7] = p1
			}

			dst = dst[8:]
		}
	}
}

func (e2d *HwEngine2d) drawChar256(y int, src []byte, dst []byte, hflip bool) {
	src = src[y*8:]
	if hflip {
		for x := 0; x < 8; x++ {
			p0 := src[x]
			if p0 != 0 {
				dst[0] = p0
				dst[1] = p0
				dst[2] = p0
				dst[3] = p0
			}
			dst = dst[4:]
		}
	} else {
		for x := 7; x >= 0; x-- {
			p0 := src[x]
			if p0 != 0 {
				dst[0] = p0
				dst[1] = p0
				dst[2] = p0
				dst[3] = p0
			}
			dst = dst[4:]
		}
	}
}

func (e2d *HwEngine2d) drawLayer0(y int, numLayer int, line []byte) {
	mapBase := 0
	charBase := 0
	if e2d.A() {
		mapBase = int((e2d.DispCnt.Value>>27)&7) * 64 * 1024
		charBase = int((e2d.DispCnt.Value>>24)&7) * 64 * 1024
	}

	var mapx, mapy int
	if numLayer == 0 {
		mapBase += int((e2d.Bg0Cnt.Value>>8)&0xF) * 2 * 1024
		charBase += int((e2d.Bg0Cnt.Value>>2)&0xF) * 16 * 1024
		mapx = int(e2d.Bg0XOfs.Value)
		mapy = y + int(e2d.Bg0YOfs.Value)
	} else if numLayer == 3 {
		mapBase += int((e2d.Bg3Cnt.Value>>8)&0xF) * 2 * 1024
		charBase += int((e2d.Bg3Cnt.Value>>2)&0xF) * 16 * 1024
		mapx = int(e2d.Bg3XOfs.Value)
		mapy = y + int(e2d.Bg3YOfs.Value)
	}

	depth256 := (e2d.DispCnt.Value>>7)&1 != 0
	// tmap := e2d.vram[mapBase : mapBase+0x800]
	var tmap, chars []byte
	if e2d.A() {
		tmap = nds9.Bus.FetchPointer(uint32(0x6000000 + mapBase))
		chars = nds9.Bus.FetchPointer(uint32(0x6000000 + charBase))
	} else {
		tmap = nds9.Bus.FetchPointer(uint32(0x6200000 + mapBase))
		chars = nds9.Bus.FetchPointer(uint32(0x6200000 + charBase))
	}

	// numAreas := (e2d.DispCnt.Value >> 14) & 3

	mapy &= 255
	mapLine := tmap[32*2*(mapy/8):]

	// if y == 15 {
	// 	for x := 0; x < cScreenWidth*4; x++ {
	// 		line[x] = 0xff
	// 	}
	// 	return
	// }

	for x := 0; x < cScreenWidth/8; x++ {

		mapx &= 255
		tile := le16(mapLine[2*(mapx/8):])
		mapx += 8

		tnum := int(tile & 1023)
		hflip := (tile>>10)&1 != 0
		vflip := (tile>>11)&1 != 0
		ty := mapy & 7
		if vflip {
			ty = 7 - ty
		}

		if depth256 {
			e2d.drawChar256(ty, chars[tnum*64:], line, hflip)
		} else {
			e2d.drawChar16(ty, chars[tnum*32:], line, hflip)
		}
		line = line[8*4:]
	}

}

func (e2d *HwEngine2d) WriteBG0CNT(_, val uint16) {
	log.Infof("[lcd] bg0cnt=%x", val)
}

func (e2d *HwEngine2d) DrawLine(y int, line []byte) {
	bg0on := (e2d.DispCnt.Value >> 8) & 1
	bg1on := (e2d.DispCnt.Value >> 9) & 1
	bg2on := (e2d.DispCnt.Value >> 10) & 1
	bg3on := (e2d.DispCnt.Value >> 11) & 1
	objon := (e2d.DispCnt.Value >> 12) & 1
	win0on := (e2d.DispCnt.Value >> 13) & 1
	win1on := (e2d.DispCnt.Value >> 14) & 1
	objwinon := (e2d.DispCnt.Value >> 15) & 1

	if y == 0 {
		log.Infof("[lcd %s] bg=[%d,%d,%d,%d] obj=%d win=[%d,%d,%d]",
			string('A'+e2d.Idx), bg0on, bg1on, bg2on, bg3on, objon, win0on, win1on, objwinon)
		log.Infof("[lcd %s] scroll0=[%d,%d] scroll3=[%d,%d] size1=%d size3=%d",
			string('A'+e2d.Idx),
			e2d.Bg0XOfs.Value, e2d.Bg0YOfs.Value,
			e2d.Bg3XOfs.Value, e2d.Bg3YOfs.Value,
			e2d.Bg0Cnt.Value>>14, e2d.Bg3Cnt.Value>>13)
	}

	if bg0on != 0 {
		e2d.drawLayer0(y, 0, e2d.lineBuf[4*8:])
	}
	if bg3on != 0 {
		e2d.drawLayer0(y, 3, e2d.lineBuf[4*8:])
	}

	copy(line, e2d.lineBuf[4*8:])
}
