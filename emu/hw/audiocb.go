package hw

/*
#include <stdio.h>
typedef unsigned char Uint8;
typedef unsigned short Uint16;
typedef signed short Int16;

extern void audioCallbackGo(void *userdata, Uint8 *stream, int len);

// NOTE: this cannot be defined in audio.go because that files contains an export directive.
// This is a documentated limitation of cgo:
//
//    Using //export in a file places a restriction on the preamble: since it is
//    copied into two different C output files, it must not contain any definitions,
//    only declarations. If a file contains both definitions and declarations, then the
//    two output files will produce duplicate symbols and the linker will fail. To
//    avoid this, definitions must be placed in preambles in other files, or in C source files.
//
void audioCallbackC(void *userdata, Uint8 *stream, int len) {
	audioCallbackGo(userdata, stream, len);
}
*/
import "C"
