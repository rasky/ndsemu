package gfx

import (
	"reflect"
	"unsafe"
)

type Line struct {
	ptr uintptr
}

func NewLine(mem []byte) Line {
	return Line{uintptr(unsafe.Pointer(&mem[0]))}
}

func (l Line) IsNil() bool { return l.ptr == 0 }

func (l *Line) Add8(x int) {
	l.ptr += uintptr(x)
}

func (l *Line) Add16(x int) {
	l.ptr += uintptr(x * 2)
}

func (l *Line) Add32(x int) {
	l.ptr += uintptr(x * 4)
}

func (l Line) Get8(x int) uint8 {
	xx := uintptr(x)
	return *(*uint8)(unsafe.Pointer(l.ptr + xx))
}

func (l Line) Get16(x int) uint16 {
	xx := uintptr(x * 2)
	return *(*uint16)(unsafe.Pointer(l.ptr + xx))
}

func (l Line) Get32(x int) uint32 {
	xx := uintptr(x * 4)
	return *(*uint32)(unsafe.Pointer(l.ptr + xx))
}

func (l Line) Set8(x int, val uint8) {
	xx := uintptr(x)
	*(*uint8)(unsafe.Pointer(l.ptr + xx)) = val
}

func (l Line) Set16(x int, val uint16) {
	xx := uintptr(x * 2)
	*(*uint16)(unsafe.Pointer(l.ptr + xx)) = val
}

func (l Line) Set32(x int, val uint32) {
	xx := uintptr(x * 4)
	*(*uint32)(unsafe.Pointer(l.ptr + xx)) = val
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

func NewBufferMem(w, h int) Buffer {
	mem := make([]byte, w*h*4)
	return NewBuffer(unsafe.Pointer(&mem[0]), w, h, w*4)
}

func (buf *Buffer) Pointer() unsafe.Pointer {
	return buf.ptr
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
