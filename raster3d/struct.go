package raster3d

import (
	"ndsemu/emu"
	"ndsemu/emu/gfx"
)

type RenderVertexFlags uint32

const (
	RVFClipLeft RenderVertexFlags = 1 << iota
	RVFClipRight
	RVFClipTop
	RVFClipBottom
	RVFClipNear
	RVFClipFar
	RVFTransformed // vertex has been already transformed to screen space

	RVFClipMask = (RVFClipLeft | RVFClipRight | RVFClipTop | RVFClipBottom | RVFClipNear | RVFClipFar)
)

type Vertex struct {
	// Coordinates in clip-space
	cx, cy, cz, cw emu.Fixed12

	// Screen coordinates (fractional part is always zero)
	x, y, z emu.Fixed12

	// Texture coordinates
	s, t emu.Fixed12

	// Vertex color
	rgb color

	// Misc flags
	flags RenderVertexFlags
}

// NOTE: these flags match the polygon attribute word defined in the
// geometry coprocessor.
type PolygonFlags uint32

const (
	PFRenderBack  PolygonFlags = 1 << 6
	PFRenderFront              = 1 << 7
	PFQuad                     = 1 << 31
)

func (f PolygonFlags) Alpha() int      { return int(f>>16) & 0x1F }
func (f PolygonFlags) ColorMode() uint { return uint(f>>4) & 3 }

const (
	LerpX   = iota // coordinate on screen (X)
	LerpZ          // depth on screen (Z or W)
	LerpT          // texture X coordinate (T)
	LerpS          // texture Y coordinate (S)
	LerpRGB        // vertex color (RGB)
	NumLerps
)

//go:generate go run gen/genfillers.go -filename polyfillers.go

type Polygon struct {
	vtx   [3]*Vertex
	flags PolygonFlags
	tex   Texture

	// y coordinate of middle vertex
	hy int32

	// linear interpolators for left and right edge of the polygon
	left  [NumLerps]lerp
	right [NumLerps]lerp

	// polyfiller for this polygon
	filler func(*HwEngine3d, *Polygon, gfx.Line, gfx.Line, gfx.Line)

	// texture pointer
	texptr []byte
}

//go:generate stringer -type TexFormat
type TexFormat int
type TexFlags int

const (
	TexNone TexFormat = iota
	TexA3I5
	Tex4
	Tex16
	Tex256
	Tex4x4
	TexA5I3
	TexDirect
)

const (
	TexSFlip TexFlags = 1 << iota
	TexTFlip
	TexSRepeat
	TexTRepeat
)

type Texture struct {
	VramTexOffset uint32
	VramPalOffset uint32
	Width, Height uint32
	PitchShift    uint

	// Masks to implement fast clamping in polyfillers. They're
	// set to ^(texturesize-1) if clamping is active, or 0 if not
	// so that clamp does not get triggered.
	SClampMask uint32
	TClampMask uint32

	// Masks to implement fast texture flipping in polyfillers.
	// They're set to the texture size (eg: 0x100) so that they
	// become a mask to check whether the coordinate is being
	// repeated an odd number of times.
	SFlipMask uint32
	TFlipMask uint32

	Transparency bool
	Format       TexFormat
	Flags        TexFlags
}
