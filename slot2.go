package main

import (
	"io"
	"io/ioutil"
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

func (slot *HwSlot2) MapCart(f io.Reader) error {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	data2 := make([]byte, roundup2(len(data)))
	for i := 0; i < len(data2); i++ {
		data2[i] = 0xFF
	}
	copy(data2, data)

	slot.Rom = data2
	return nil
}

func (slot *HwSlot2) MapCartFile(fn string) error {
	f, err := os.Open(fn)
	if err != nil {
		return err
	}

	defer f.Close()
	return slot.MapCart(f)
}

func (slot *HwSlot2) UnmapCart() {
	slot.Rom = highz[:]
}
