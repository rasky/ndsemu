package homebrew

import (
	_ "embed"
)

//go:embed dldi/fcsr.dldi
var fcsrdldi []byte

// FCSR (FlashCart SRAM) is the name of the DLDI module (aka libfat backend)
// that reads the FAT partition from NDS Slot 2, searching for it in the whole
// 0x0Axxxxxxx area.
//
// This basically assumes that the homebrew developer has concatenated the FAT
// partition to the actual game ROM, and then flashed the cart.
// It is also the simplest one to emulate for us, as it doesn't require to
// emulate a real hardware peripheral (eg: a SD card).

// Apply the DLDI patch to the ROM to install the FCSR backend
func FcsrPatchDldi(rom []byte) error {
	return DldiPatch(rom, fcsrdldi)
}

// Patch the FAT partition adding a magic identifier to it.
// This is required because the FCSR code searched for this special
// identifier to find the FAT partition in the Slot 2 area.
// The identified is written at a specific offset in the FAT header
// that is normally unused.
func FcsrPatchFatPartition(fat []byte) {
	copy(fat[0x100:], []byte("FCSR Chishm FAT\x00"))
}
