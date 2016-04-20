package homebrew

import (
	"encoding/binary"
	"io"

	"github.com/howeyc/crc16"
)

// PassMe implements the same hack that the real PassMe hardware
// does, by overlaying to a Slot1 (NDS) ROM and patching it so
// that the BIOS is tricked into jumping into the Slot2 area
// (where a homebrew ROM is placed).
type PassMe struct {
	head [0x160]byte
	rom  io.ReaderAt
}

func NewPassMe(rom io.ReaderAt) *PassMe {
	ovl := &PassMe{
		rom: rom,
	}

	rom.ReadAt(ovl.head[:], 0)

	// opcode LDR PC,[027FFE24h] at 27FFE04h
	binary.LittleEndian.PutUint32(ovl.head[0x04:], 0xE59FF018)

	// Set autostart bit
	ovl.head[0x1F] = 0x04

	// Set ARM9 rom offset to nn01nnnnh (above secure area)
	ovl.head[0x22] = 0x01

	// Patch ARM9 entry address to endless loop
	binary.LittleEndian.PutUint32(ovl.head[0x24:], 0x027FFE04)

	// Patch ARM7 entry address in GBA slot
	binary.LittleEndian.PutUint32(ovl.head[0x34:], 0x080000C0)

	// Adjust CRC16
	sum := ^crc16.ChecksumIBM(ovl.head[:0x15E])
	binary.LittleEndian.PutUint16(ovl.head[0x15E:], sum)

	return ovl
}

func (ovl *PassMe) ReadAt(p []byte, off int64) (n int, err error) {
	n, err = ovl.rom.ReadAt(p, off)
	if err != nil {
		return
	}

	if off < int64(len(ovl.head)) {
		copy(p, ovl.head[off:])
	}
	return
}
