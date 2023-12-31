package sdl

// #include "sdl_wrapper.h"
//
// #if !(SDL_VERSION_ATLEAST(2,0,5))
//
// enum
// {
// #if SDL_BYTEORDER == SDL_BIG_ENDIAN
//     SDL_PIXELFORMAT_RGBA32 = SDL_PIXELFORMAT_RGBA8888,
//     SDL_PIXELFORMAT_ARGB32 = SDL_PIXELFORMAT_ARGB8888,
//     SDL_PIXELFORMAT_BGRA32 = SDL_PIXELFORMAT_BGRA8888,
//     SDL_PIXELFORMAT_ABGR32 = SDL_PIXELFORMAT_ABGR8888
// #else
//     SDL_PIXELFORMAT_RGBA32 = SDL_PIXELFORMAT_ABGR8888,
//     SDL_PIXELFORMAT_ARGB32 = SDL_PIXELFORMAT_BGRA8888,
//     SDL_PIXELFORMAT_BGRA32 = SDL_PIXELFORMAT_ARGB8888,
//     SDL_PIXELFORMAT_ABGR32 = SDL_PIXELFORMAT_RGBA8888
// #endif
// };
//
// #endif
//
// int bytesPerPixel(Uint32 format) {
//   return SDL_BYTESPERPIXEL(format);
// }
//
// int bitsPerPixel(Uint32 format) {
//   return SDL_BITSPERPIXEL(format);
// }
import "C"
import (
	"image/color"
	"unsafe"
)

// PixelFormat contains pixel format information.
// (https://wiki.libsdl.org/SDL_PixelFormat)
type PixelFormat struct {
	Format        uint32       // one of the PIXELFORMAT values (https://wiki.libsdl.org/SDL_PixelFormatEnum)
	Palette       *Palette     // palette structure associated with this pixel format, or nil if the format doesn't have a palette (https://wiki.libsdl.org/SDL_Palette)
	BitsPerPixel  uint8        // the number of significant bits in a pixel value, eg: 8, 15, 16, 24, 32
	BytesPerPixel uint8        // the number of bytes required to hold a pixel value, eg: 1, 2, 3, 4
	_             [2]uint8     // padding
	Rmask         uint32       // a mask representing the location of the red component of the pixel
	Gmask         uint32       // a mask representing the location of the green component of the pixel
	Bmask         uint32       // a mask representing the location of the blue component of the pixel
	Amask         uint32       // a mask representing the location of the alpha component of the pixel or 0 if the pixel format doesn't have any alpha information
	rLoss         uint8        // (internal use)
	gLoss         uint8        // (internal use)
	bLoss         uint8        // (internal use)
	aLoss         uint8        // (internal use)
	rShift        uint8        // (internal use)
	gShift        uint8        // (internal use)
	bShift        uint8        // (internal use)
	aShift        uint8        // (internal use)
	refCount      int32        // (internal use)
	next          *PixelFormat // (internal use)
}
type cPixelFormat C.SDL_PixelFormat

// Palette contains palette information.
// (https://wiki.libsdl.org/SDL_Palette)
type Palette struct {
	Ncolors  int32  // the number of colors in the palette
	Colors   *Color // an array of Color structures representing the palette (https://wiki.libsdl.org/SDL_Color)
	version  uint32 // incrementally tracks changes to the palette (internal use)
	refCount int32  // reference count (internal use)
}
type cPalette C.SDL_Palette

// Color represents a color. This implements image/color.Color interface.
// (https://wiki.libsdl.org/SDL_Color)
type Color color.RGBA

// Uint32 return uint32 representation of RGBA color.
func (c Color) Uint32() uint32 {
	var v uint32
	v |= uint32(c.R) << 24
	v |= uint32(c.G) << 16
	v |= uint32(c.B) << 8
	v |= uint32(c.A)
	return v
}

// Pixel types.
const (
	PIXELTYPE_UNKNOWN  = C.SDL_PIXELTYPE_UNKNOWN
	PIXELTYPE_INDEX1   = C.SDL_PIXELTYPE_INDEX1
	PIXELTYPE_INDEX4   = C.SDL_PIXELTYPE_INDEX4
	PIXELTYPE_INDEX8   = C.SDL_PIXELTYPE_INDEX8
	PIXELTYPE_PACKED8  = C.SDL_PIXELTYPE_PACKED8
	PIXELTYPE_PACKED16 = C.SDL_PIXELTYPE_PACKED16
	PIXELTYPE_PACKED32 = C.SDL_PIXELTYPE_PACKED32
	PIXELTYPE_ARRAYU8  = C.SDL_PIXELTYPE_ARRAYU8
	PIXELTYPE_ARRAYU16 = C.SDL_PIXELTYPE_ARRAYU16
	PIXELTYPE_ARRAYU32 = C.SDL_PIXELTYPE_ARRAYU32
	PIXELTYPE_ARRAYF16 = C.SDL_PIXELTYPE_ARRAYF16
	PIXELTYPE_ARRAYF32 = C.SDL_PIXELTYPE_ARRAYF32
)

// Bitmap pixel order high bit -> low bit.
const (
	BITMAPORDER_NONE = C.SDL_BITMAPORDER_NONE
	BITMAPORDER_4321 = C.SDL_BITMAPORDER_4321
	BITMAPORDER_1234 = C.SDL_BITMAPORDER_1234
)

// Packed component order high bit -> low bit.
const (
	PACKEDORDER_NONE = C.SDL_PACKEDORDER_NONE
	PACKEDORDER_XRGB = C.SDL_PACKEDORDER_XRGB
	PACKEDORDER_RGBX = C.SDL_PACKEDORDER_RGBX
	PACKEDORDER_ARGB = C.SDL_PACKEDORDER_ARGB
	PACKEDORDER_RGBA = C.SDL_PACKEDORDER_RGBA
	PACKEDORDER_XBGR = C.SDL_PACKEDORDER_XBGR
	PACKEDORDER_BGRX = C.SDL_PACKEDORDER_BGRX
	PACKEDORDER_ABGR = C.SDL_PACKEDORDER_ABGR
	PACKEDORDER_BGRA = C.SDL_PACKEDORDER_BGRA
)

// Array component order low byte -> high byte.
const (
	ARRAYORDER_NONE = C.SDL_ARRAYORDER_NONE
	ARRAYORDER_RGB  = C.SDL_ARRAYORDER_RGB
	ARRAYORDER_RGBA = C.SDL_ARRAYORDER_RGBA
	ARRAYORDER_ARGB = C.SDL_ARRAYORDER_ARGB
	ARRAYORDER_BGR  = C.SDL_ARRAYORDER_BGR
	ARRAYORDER_BGRA = C.SDL_ARRAYORDER_BGRA
	ARRAYORDER_ABGR = C.SDL_ARRAYORDER_ABGR
)

// Packed component layout.
const (
	PACKEDLAYOUT_NONE    = C.SDL_PACKEDLAYOUT_NONE
	PACKEDLAYOUT_332     = C.SDL_PACKEDLAYOUT_332
	PACKEDLAYOUT_4444    = C.SDL_PACKEDLAYOUT_4444
	PACKEDLAYOUT_1555    = C.SDL_PACKEDLAYOUT_1555
	PACKEDLAYOUT_5551    = C.SDL_PACKEDLAYOUT_5551
	PACKEDLAYOUT_565     = C.SDL_PACKEDLAYOUT_565
	PACKEDLAYOUT_8888    = C.SDL_PACKEDLAYOUT_8888
	PACKEDLAYOUT_2101010 = C.SDL_PACKEDLAYOUT_2101010
	PACKEDLAYOUT_1010102 = C.SDL_PACKEDLAYOUT_1010102
)

// Pixel format values.
const (
	PIXELFORMAT_UNKNOWN     = C.SDL_PIXELFORMAT_UNKNOWN
	PIXELFORMAT_INDEX1LSB   = C.SDL_PIXELFORMAT_INDEX1LSB
	PIXELFORMAT_INDEX1MSB   = C.SDL_PIXELFORMAT_INDEX1MSB
	PIXELFORMAT_INDEX4LSB   = C.SDL_PIXELFORMAT_INDEX4LSB
	PIXELFORMAT_INDEX4MSB   = C.SDL_PIXELFORMAT_INDEX4MSB
	PIXELFORMAT_INDEX8      = C.SDL_PIXELFORMAT_INDEX8
	PIXELFORMAT_RGB332      = C.SDL_PIXELFORMAT_RGB332
	PIXELFORMAT_RGB444      = C.SDL_PIXELFORMAT_RGB444
	PIXELFORMAT_RGB555      = C.SDL_PIXELFORMAT_RGB555
	PIXELFORMAT_BGR555      = C.SDL_PIXELFORMAT_BGR555
	PIXELFORMAT_ARGB4444    = C.SDL_PIXELFORMAT_ARGB4444
	PIXELFORMAT_RGBA4444    = C.SDL_PIXELFORMAT_RGBA4444
	PIXELFORMAT_ABGR4444    = C.SDL_PIXELFORMAT_ABGR4444
	PIXELFORMAT_BGRA4444    = C.SDL_PIXELFORMAT_BGRA4444
	PIXELFORMAT_ARGB1555    = C.SDL_PIXELFORMAT_ARGB1555
	PIXELFORMAT_RGBA5551    = C.SDL_PIXELFORMAT_RGBA5551
	PIXELFORMAT_ABGR1555    = C.SDL_PIXELFORMAT_ABGR1555
	PIXELFORMAT_BGRA5551    = C.SDL_PIXELFORMAT_BGRA5551
	PIXELFORMAT_RGB565      = C.SDL_PIXELFORMAT_RGB565
	PIXELFORMAT_BGR565      = C.SDL_PIXELFORMAT_BGR565
	PIXELFORMAT_RGB24       = C.SDL_PIXELFORMAT_RGB24
	PIXELFORMAT_BGR24       = C.SDL_PIXELFORMAT_BGR24
	PIXELFORMAT_RGB888      = C.SDL_PIXELFORMAT_RGB888
	PIXELFORMAT_RGBX8888    = C.SDL_PIXELFORMAT_RGBX8888
	PIXELFORMAT_BGR888      = C.SDL_PIXELFORMAT_BGR888
	PIXELFORMAT_BGRX8888    = C.SDL_PIXELFORMAT_BGRX8888
	PIXELFORMAT_ARGB8888    = C.SDL_PIXELFORMAT_ARGB8888
	PIXELFORMAT_RGBA8888    = C.SDL_PIXELFORMAT_RGBA8888
	PIXELFORMAT_ABGR8888    = C.SDL_PIXELFORMAT_ABGR8888
	PIXELFORMAT_BGRA8888    = C.SDL_PIXELFORMAT_BGRA8888
	PIXELFORMAT_ARGB2101010 = C.SDL_PIXELFORMAT_ARGB2101010
	PIXELFORMAT_YV12        = C.SDL_PIXELFORMAT_YV12
	PIXELFORMAT_IYUV        = C.SDL_PIXELFORMAT_IYUV
	PIXELFORMAT_YUY2        = C.SDL_PIXELFORMAT_YUY2
	PIXELFORMAT_UYVY        = C.SDL_PIXELFORMAT_UYVY
	PIXELFORMAT_YVYU        = C.SDL_PIXELFORMAT_YVYU
)

// Pixel format variables.
var (
	PIXELFORMAT_RGBA32 = C.SDL_PIXELFORMAT_RGBA32
	PIXELFORMAT_ARGB32 = C.SDL_PIXELFORMAT_ARGB32
	PIXELFORMAT_BGRA32 = C.SDL_PIXELFORMAT_BGRA32
	PIXELFORMAT_ABGR32 = C.SDL_PIXELFORMAT_ABGR32
)

// These define alpha as the opacity of a surface.
const (
	ALPHA_OPAQUE      = C.SDL_ALPHA_OPAQUE
	ALPHA_TRANSPARENT = C.SDL_ALPHA_TRANSPARENT
)

func (format *PixelFormat) cptr() *C.SDL_PixelFormat {
	return (*C.SDL_PixelFormat)(unsafe.Pointer(format))
}

func (palette *Palette) cptr() *C.SDL_Palette {
	return (*C.SDL_Palette)(unsafe.Pointer(palette))
}

/*
 * the following code is modified version of the code from bitbucket.org/dooots/go-sdl2
 */

// GetPixelFormatName returns the human readable name of a pixel format.
// (https://wiki.libsdl.org/SDL_GetPixelFormatName)
func GetPixelFormatName(format uint) string {
	return C.GoString(C.SDL_GetPixelFormatName(C.Uint32(format)))
}

// PixelFormatEnumToMasks converts one of the enumerated pixel formats to a bpp value and RGBA masks.
// (https://wiki.libsdl.org/SDL_PixelFormatEnumToMasks)
func PixelFormatEnumToMasks(format uint) (bpp int, rmask, gmask, bmask, amask uint32, err error) {
	result := C.SDL_PixelFormatEnumToMasks(C.Uint32(format), (*C.int)(unsafe.Pointer(&bpp)),
		(*C.Uint32)(&rmask), (*C.Uint32)(&gmask), (*C.Uint32)(&bmask),
		(*C.Uint32)(&amask))
	if result == C.SDL_FALSE {
		err = GetError()
	}
	return
}

// MasksToPixelFormatEnum converts a bpp value and RGBA masks to an enumerated pixel format.
// (https://wiki.libsdl.org/SDL_MasksToPixelFormatEnum)
func MasksToPixelFormatEnum(bpp int, rmask, gmask, bmask, amask uint32) uint {
	return uint(C.SDL_MasksToPixelFormatEnum(C.int(bpp), C.Uint32(rmask), C.Uint32(gmask),
		C.Uint32(bmask), C.Uint32(amask)))
}

// AllocFormat creates a PixelFormat structure corresponding to a pixel format.
// (https://wiki.libsdl.org/SDL_AllocFormat)
func AllocFormat(format uint) (*PixelFormat, error) {
	r := (*PixelFormat)(unsafe.Pointer(C.SDL_AllocFormat(C.Uint32(format))))
	if r == nil {
		return nil, GetError()
	}
	return r, nil
}

// Free frees the PixelFormat structure allocated by AllocFormat().
// (https://wiki.libsdl.org/SDL_FreeFormat)
func (format *PixelFormat) Free() {
	C.SDL_FreeFormat((*C.SDL_PixelFormat)(unsafe.Pointer(format)))
}

// AllocPalette creates a palette structure with the specified number of color entries.
// (https://wiki.libsdl.org/SDL_AllocPalette)
func AllocPalette(ncolors int) (*Palette, error) {
	r := (*Palette)(unsafe.Pointer(C.SDL_AllocPalette(C.int(ncolors))))
	if r == nil {
		return nil, GetError()
	}
	return r, nil
}

// SetPalette sets the palette for the pixel format structure.
// (https://wiki.libsdl.org/SDL_SetPixelFormatPalette)
func (format *PixelFormat) SetPalette(palette *Palette) error {
	r := C.SDL_SetPixelFormatPalette((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.SDL_Palette)(unsafe.Pointer(palette)))
	if r != 0 {
		return GetError()
	}
	return nil
}

// SetColors sets a range of colors in the palette.
// (https://wiki.libsdl.org/SDL_SetPaletteColors)
func (palette *Palette) SetColors(colors []Color) error {
	if colors == nil {
		return nil
	}
	var ptr *C.SDL_Color
	if len(colors) > 0 {
		ptr = (*C.SDL_Color)(unsafe.Pointer(&colors[0]))
	}

	r := C.SDL_SetPaletteColors((*C.SDL_Palette)(unsafe.Pointer(palette)),
		ptr, 0, C.int(len(colors)))
	if r != 0 {
		return GetError()
	}
	return nil
}

// Free frees the palette created with AllocPalette().
// (https://wiki.libsdl.org/SDL_FreePalette)
func (palette *Palette) Free() {
	C.SDL_FreePalette((*C.SDL_Palette)(unsafe.Pointer(palette)))
}

// MapRGB maps an RGB triple to an opaque pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL_MapRGB)
func MapRGB(format *PixelFormat, r, g, b uint8) uint32 {
	return uint32(C.SDL_MapRGB((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// MapRGBA maps an RGBA quadruple to a pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL_MapRGBA)
func MapRGBA(format *PixelFormat, r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapRGBA((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// GetRGB returns RGB values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL_GetRGB)
func GetRGB(pixel uint32, format *PixelFormat) (r, g, b uint8) {
	C.SDL_GetRGB(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	return
}

// GetRGBA returns RGBA values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL_GetRGBA)
func GetRGBA(pixel uint32, format *PixelFormat) (r, g, b, a uint8) {
	C.SDL_GetRGBA(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	return
}

// CalculateGammaRamp calculates a 256 entry gamma ramp for a gamma value.
// (https://wiki.libsdl.org/SDL_CalculateGammaRamp)
func CalculateGammaRamp(gamma float32, ramp *[256]uint16) {
	C.SDL_CalculateGammaRamp(C.float(gamma), (*C.Uint16)(unsafe.Pointer(&ramp[0])))
}

// BytesPerPixel returns the number of bytes per pixel for the given format
func BytesPerPixel(format uint32) int {
	return int(C.bytesPerPixel(C.Uint32(format)))
}

// BitsPerPixel returns the number of bits per pixel for the given format
func BitsPerPixel(format uint32) int {
	return int(C.bitsPerPixel(C.Uint32(format)))
}

type RGB444 struct {
	R, G, B byte
}

func (c RGB444) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 12
	g = uint32(c.G) << 12
	b = uint32(c.B) << 12
	return
}

var (
	RGB444Model   color.Model = color.ModelFunc(rgb444Model)
	RGB332Model   color.Model = color.ModelFunc(rgb332Model)
	RGB565Model   color.Model = color.ModelFunc(rgb565Model)
	RGB555Model   color.Model = color.ModelFunc(rgb555Model)
	BGR565Model   color.Model = color.ModelFunc(bgr565Model)
	BGR555Model   color.Model = color.ModelFunc(bgr555Model)
	ARGB4444Model color.Model = color.ModelFunc(argb4444Model)
	ABGR4444Model color.Model = color.ModelFunc(abgr4444Model)
	RGBA4444Model color.Model = color.ModelFunc(rgba4444Model)
	BGRA4444Model color.Model = color.ModelFunc(bgra4444Model)
	ARGB1555Model color.Model = color.ModelFunc(argb1555Model)
	RGBA5551Model color.Model = color.ModelFunc(rgba5551Model)
	ABGR1555Model color.Model = color.ModelFunc(abgr1555Model)
	BGRA5551Model color.Model = color.ModelFunc(bgra5551Model)
	RGBA8888Model color.Model = color.ModelFunc(rgba8888Model)
	BGRA8888Model color.Model = color.ModelFunc(bgra8888Model)
)

func rgb444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB444{uint8(r >> 12), uint8(g >> 12), uint8(b >> 12)}
}

type RGB332 struct {
	R, G, B byte
}

func (c RGB332) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 13
	g = uint32(c.G) << 13
	b = uint32(c.B) << 14
	return
}

func rgb332Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB332{uint8(r >> 13), uint8(g >> 13), uint8(b >> 14)}
}

type RGB565 struct {
	R, G, B byte
}

func (c RGB565) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 10
	b = uint32(c.B) << 11
	return
}

func rgb565Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB565{uint8(r >> 11), uint8(g >> 10), uint8(b >> 11)}
}

type RGB555 struct {
	R, G, B byte
}

func (c RGB555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 11
	b = uint32(c.B) << 11
	return
}

func rgb555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGB555{uint8(r >> 11), uint8(g >> 11), uint8(b >> 11)}
}

type BGR565 struct {
	B, G, R byte
}

func (c BGR565) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 10
	b = uint32(c.B) << 11
	return
}

func bgr565Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return BGR565{uint8(b >> 11), uint8(g >> 10), uint8(r >> 11)}
}

type BGR555 struct {
	B, G, R byte
}

func (c BGR555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 11
	g = uint32(c.G) << 11
	b = uint32(c.B) << 11
	return
}

func bgr555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return BGR555{uint8(b >> 11), uint8(g >> 11), uint8(r >> 11)}
}

type ARGB4444 struct {
	A, R, G, B byte
}

func (c ARGB4444) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A) << 4
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	return
}

func argb4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return ARGB4444{uint8(a >> 4), uint8(r >> 4), uint8(g >> 4), uint8(b >> 4)}
}

type ABGR4444 struct {
	A, B, G, R byte
}

func (c ABGR4444) RGBA() (r, g, b, a uint32) {
	a = uint32(c.A) << 4
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	return
}

func abgr4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return ABGR4444{uint8(a >> 4), uint8(b >> 4), uint8(g >> 4), uint8(r >> 4)}
}

type RGBA4444 struct {
	R, G, B, A byte
}

func (c RGBA4444) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	a = uint32(c.A) << 4
	return
}

func rgba4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA4444{uint8(r >> 4), uint8(g >> 4), uint8(b >> 4), uint8(a >> 4)}
}

type BGRA4444 struct {
	B, G, R, A byte
}

func (c BGRA4444) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 4
	g = uint32(c.G) << 4
	b = uint32(c.B) << 4
	a = uint32(c.A) << 4
	return
}

func bgra4444Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return BGRA4444{uint8(b >> 4), uint8(g >> 4), uint8(r >> 4), uint8(a >> 4)}
}

type ARGB1555 struct {
	A, R, G, B byte
}

func (c ARGB1555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func argb1555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return ARGB1555{uint8(a), uint8(r >> 3), uint8(g >> 3), uint8(b >> 3)}
}

type RGBA5551 struct {
	R, G, B, A byte
}

func (c RGBA5551) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func rgba5551Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return RGBA5551{uint8(r >> 3), uint8(g >> 3), uint8(b >> 3), uint8(a)}
}

type ABGR1555 struct {
	A, R, G, B byte
}

func (c ABGR1555) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func abgr1555Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return ABGR1555{uint8(a), uint8(r >> 3), uint8(g >> 3), uint8(b >> 3)}
}

type BGRA5551 struct {
	B, G, R, A byte
}

func (c BGRA5551) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) << 3
	g = uint32(c.G) << 3
	b = uint32(c.B) << 3
	if c.A > 0 {
		tmp := int32(-1)
		a = uint32(tmp)
	}
	return
}

func bgra5551Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	if a > 0 {
		a = 1
	} else {
		a = 0
	}
	return BGRA5551{uint8(b >> 3), uint8(g >> 3), uint8(r >> 3), uint8(a)}
}

type RGBA8888 struct {
	R, G, B, A byte
}

func (c RGBA8888) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	g = uint32(c.G)
	b = uint32(c.B)
	a = uint32(c.A)
	return
}

func rgba8888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return RGBA8888{uint8(r), uint8(g), uint8(b), uint8(a)}
}

type BGRA8888 struct {
	B, G, R, A byte
}

func (c BGRA8888) RGBA() (r, g, b, a uint32) {
	b = uint32(c.B)
	g = uint32(c.G)
	r = uint32(c.R)
	a = uint32(c.A)
	return
}

func bgra8888Model(c color.Color) color.Color {
	if _, ok := c.(color.RGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return BGRA8888{uint8(b), uint8(g), uint8(r), uint8(a)}
}
