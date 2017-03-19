package hw

/*
typedef unsigned char Uint8;
typedef unsigned short Uint16;
typedef signed short Int16;

extern void audioCallbackC(void *userdata, Uint8 *stream, int len);
*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type AudioBuffer []int16

var gout *Output

//export audioCallbackGo
func audioCallbackGo(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	n := int(length) / 2
	hdr := reflect.SliceHeader{Data: uintptr(unsafe.Pointer(stream)), Len: n, Cap: n}
	buf := *(*[]int16)(unsafe.Pointer(&hdr))
	gout.audioCallback(buf)
}

func (out *Output) audioSpecSetCallback(spec *sdl.AudioSpec) {
	spec.Callback = sdl.AudioCallback(C.audioCallbackC)
	spec.UserData = nil
	if gout != nil && gout != out {
		panic("FIXME: two outputs not supported")
	}
	gout = out
}
