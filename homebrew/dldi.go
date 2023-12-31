package homebrew

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	log "ndsemu/emu/logger"
)

var modHbrew = log.NewModule("hbrew")

var (
	dldiMagicString = []byte("\xED\xA5\x8D\xBF Chishm\x00")
	dldiVersion     = 1

	ErrNoDldiFound  = errors.New("cannot find DLDI header in ROM")
	ErrInvalidPatch = errors.New("patch is not a valid DLDI")
	ErrPatchTooBig  = errors.New("patch is larger than available space in ROM")
)

type fixMask byte

const (
	fixAll fixMask = 1 << iota
	fixGlue
	fixGot
	fixBss
)

type dldiHeader struct {
	Magic       [12]byte
	Version     byte
	DriverSize  byte
	FixSections fixMask
	AvailSpace  byte
	Name        [48]byte
	TextStart   uint32
	TextEnd     uint32
	GlueStart   uint32
	GlueEnd     uint32
	GotStart    uint32
	GotEnd      uint32
	BssStart    uint32
	BssEnd      uint32
	IoInterface struct {
		IoType           uint32
		Features         uint32
		FuncStartup      uint32
		FuncIsInserted   uint32
		FuncReadSectors  uint32
		FuncWriteSectors uint32
		FuncClearStatus  uint32
		FuncShutdown     uint32
	}
}

func dldiFindHeader(rom io.ReaderAt) int64 {
	var buf [4096]byte
	off := int64(0)
	for {
		n, err := rom.ReadAt(buf[:], off)
		if idx := bytes.Index(buf[:n], dldiMagicString); idx >= 0 {
			return off + int64(idx)
		}
		if err == io.EOF {
			return -1
		}
	}
}

func relocPointers(data []byte, hstart, hend uint32, reloc uint32) {
	for i := 0; i < len(data); i += 4 {
		ptr := binary.LittleEndian.Uint32(data[i:])
		if ptr >= hstart && ptr < hend {
			binary.LittleEndian.PutUint32(data[i:], ptr+reloc)
		}
	}
}

func DldiPatch(rom []byte, patch []byte) error {
	off := bytes.Index(rom, dldiMagicString)
	if off < 0 {
		return ErrNoDldiFound
	}

	var dh, ah dldiHeader

	binary.Read(bytes.NewReader(rom[off:]), binary.LittleEndian, &ah)
	if bytes.Compare(ah.Magic[:], dldiMagicString) != 0 || int(ah.Version) != dldiVersion {
		return ErrNoDldiFound
	}

	binary.Read(bytes.NewReader(patch), binary.LittleEndian, &dh)
	if bytes.Compare(dh.Magic[:], dldiMagicString) != 0 || int(dh.Version) != dldiVersion {
		return ErrInvalidPatch
	}

	if dh.DriverSize > ah.AvailSpace {
		return ErrPatchTooBig
	}

	// Resize patch file to make sure it's as big as the header declares. Enlarged area is
	// probably just BSS
	if len(patch) < 1<<dh.DriverSize {
		big := make([]byte, 1<<dh.DriverSize)
		copy(big, patch)
		patch = big
	}

	astart := ah.TextStart
	if astart == 0 {
		astart = ah.IoInterface.FuncStartup - 0x80
	}
	dstart := dh.TextStart
	dend := dh.TextStart + (1 << dh.DriverSize)
	reloc := astart - dstart

	modHbrew.Infof("install DLDI driver: %s", string(dh.Name[:bytes.IndexByte(dh.Name[:], 0)]))

	dh.AvailSpace = ah.AvailSpace
	if dh.FixSections&fixAll != 0 {
		relocPointers(patch[dh.TextStart-dstart:dh.TextEnd-dstart], dstart, dend, reloc)
	}
	if dh.FixSections&fixGlue != 0 {
		relocPointers(patch[dh.GlueStart-dstart:dh.GlueEnd-dstart], dstart, dend, reloc)
	}
	if dh.FixSections&fixGot != 0 {
		relocPointers(patch[dh.GotStart-dstart:dh.GotEnd-dstart], dstart, dend, reloc)
	}
	if dh.FixSections&fixBss != 0 {
		for i := dh.BssStart - dstart; i < dh.BssEnd-dstart; i++ {
			patch[i] = 0
		}
	}

	dh.TextStart += reloc
	dh.TextEnd += reloc
	dh.GlueStart += reloc
	dh.GlueEnd += reloc
	dh.GotStart += reloc
	dh.GotEnd += reloc
	dh.BssStart += reloc
	dh.BssEnd += reloc
	dh.IoInterface.FuncStartup += reloc
	dh.IoInterface.FuncIsInserted += reloc
	dh.IoInterface.FuncReadSectors += reloc
	dh.IoInterface.FuncWriteSectors += reloc
	dh.IoInterface.FuncClearStatus += reloc
	dh.IoInterface.FuncShutdown += reloc

	// Write back patched header into patch buffer
	var headbuf bytes.Buffer
	binary.Write(&headbuf, binary.LittleEndian, &dh)
	copy(patch, headbuf.Bytes())

	// Apply patch
	copy(rom[off:], patch)
	return nil
}
