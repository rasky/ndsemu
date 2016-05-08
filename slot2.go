package main

import (
	"io"
	"io/ioutil"
	"ndsemu/homebrew"
	"os"
)

var highz [16]byte = [...]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}

type HwSlot2 struct {
	Rom []byte
	Ram [64 * 1024]byte
}

func NewHwSlot2() *HwSlot2 {
	return &HwSlot2{
		Rom: highz[:],
	}
}

func roundup2(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

func (slot *HwSlot2) mapCart(data []byte, concat bool) error {
	if concat {
		slot.Rom = append(slot.Rom, data...)
	} else {
		slot.Rom = data
	}

	sz := roundup2(len(slot.Rom))
	if sz != len(slot.Rom) {
		data2 := make([]byte, sz)
		for i := 0; i < len(data2); i++ {
			data2[i] = 0xFF
		}
		copy(data2, slot.Rom)
		slot.Rom = data2
	}

	return nil
}

func (slot *HwSlot2) MapCart(f io.Reader) error {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	return slot.mapCart(data, false)
}

func (slot *HwSlot2) MapCartFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer f.Close()
	return slot.MapCart(f)
}

func (slot *HwSlot2) HomebrewMapFatFile(fn string) error {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	// Patch the FAT partition to be usable by the FCSR backend of libfat
	homebrew.FcsrPatchFatPartition(data)

	return slot.mapCart(data, true)
}

func (slot *HwSlot2) UnmapCart() {
	slot.Rom = highz[:]
}
