package gfx

import (
	"reflect"
	"unsafe"
)

type Line struct {
	ptr uintptr
}

func (l Line) SetRGB(x int, r, g, b uint8) {
	xx := uintptr(x * 4)
	*(*uint8)(unsafe.Pointer(l.ptr + xx)) = r
	*(*uint8)(unsafe.Pointer(l.ptr + xx + 1)) = g
	*(*uint8)(unsafe.Pointer(l.ptr + xx + 2)) = b
}

type Buffer struct {
	ptr           unsafe.Pointer
	Width, Height int
	pitch         int
}

func NewBuffer(ptr unsafe.Pointer, w, h, pitch int) Buffer {
	return Buffer{ptr: ptr, Width: w, Height: h, pitch: pitch}
}

func (buf *Buffer) Line(y int) Line {
	if y >= 0 && y < buf.Height {
		ptr := uintptr(buf.ptr) + uintptr(y*buf.pitch)
		return Line{ptr}
	}
	panic("invalid line")
}

func (buf *Buffer) LineAsSlice(y int) []uint8 {
	if y >= 0 && y < buf.Height {
		ptr := uintptr(buf.ptr) + uintptr(y*buf.pitch)
		slice := reflect.SliceHeader{Data: ptr, Len: buf.Width * 4, Cap: buf.Width * 4}
		return *(*[]uint8)(unsafe.Pointer(&slice))
	}
	return nil
}
