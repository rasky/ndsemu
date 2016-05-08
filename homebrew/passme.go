package homebrew

import (
	"encoding/binary"
	"io"

	"github.com/howeyc/crc16"
)

type ovlReaderAt struct {
	f    io.ReaderAt
	data []byte
	off  int64
}

func (ovl *ovlReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	n, err = ovl.f.ReadAt(p, off)
	if err != nil {
		return
	}

	x0 := off
	if ovl.off > x0 {
		x0 = ovl.off
	}

	x1 := off + int64(n)
	if ovl.off+int64(len(ovl.data)) < x1 {
		x1 = ovl.off + int64(len(ovl.data))
	}

	if x1 > x0 {
		copy(p[x0-off:x1-off], ovl.data[x0-ovl.off:x1-ovl.off])
	}
	return
}

// NewPassMe implements the same hack that the real PassMe hardware
// does, by overlaying to a Slot1 (NDS) ROM and patching it so
// that the BIOS is tricked into jumping into the Slot2 area
// (where a homebrew ROM is placed).
func NewPassMe(rom io.ReaderAt) io.ReaderAt {
	var head [0x160]byte
	rom.ReadAt(head[:], 0)

	// opcode LDR PC,[027FFE24h] at 27FFE04h
	binary.LittleEndian.PutUint32(head[0x04:], 0xE59FF018)

	// Set autostart bit
	head[0x1F] = 0x04

	// Set ARM9 rom offset to nn01nnnnh (above secure area)
	head[0x22] = 0x01

	// Patch ARM9 entry address to endless loop
	binary.LittleEndian.PutUint32(head[0x24:], 0x027FFE04)

	// Patch ARM7 entry address in GBA slot
	binary.LittleEndian.PutUint32(head[0x34:], 0x080000C0)

	// Adjust CRC16
	sum := ^crc16.ChecksumIBM(head[:0x15E])
	binary.LittleEndian.PutUint16(head[0x15E:], sum)

	return &ovlReaderAt{
		f:    rom,
		data: head[:],
		off:  0,
	}
}
