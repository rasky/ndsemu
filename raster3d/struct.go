package raster3d

import "ndsemu/emu"

type RenderVertexFlags uint32

const (
	RVFClipLeft RenderVertexFlags = 1 << iota
	RVFClipRight
	RVFClipTop
	RVFClipBottom
	RVFClipNear
	RVFClipFar
	RVFTransformed // vertex has been already transformed to screen space

	RVFClipAnything = (RVFClipLeft | RVFClipRight | RVFClipTop | RVFClipBottom | RVFClipNear | RVFClipFar)
)

type Vertex struct {
	// Coordinates in clip-space
	cx, cy, cz, cw emu.Fixed12

	// Screen coordinates (fractional part is always zero)
	x, y, z emu.Fixed12

	// Texture coordinates
	s, t emu.Fixed12

	// Misc flags
	flags RenderVertexFlags
}

// NOTE: these flags match the polygon attribute word defined in the
// geometry coprocessor.
type PolygonFlags uint32

const (
	PFQuad PolygonFlags = 1 << 31
)

const (
	LerpX = iota // coordinate on screen (X)
	LerpZ        // depth on screen (Z or W)
	LerpT        // texture X coordinate (T)
	LerpS        // texture Y coordinate (S)
	NumLerps
)

type Polygon struct {
	vtx   [4]int
	flags PolygonFlags
	tex   Texture

	// y coordinate of middle vertex
	hy int32

	// linear interpolators for left and right edge of the polygon
	left  [NumLerps]lerp
	right [NumLerps]lerp

	// texture pointer
	texptr []byte
}

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
	SMask, TMask  uint32
	PitchShift    uint
	Transparency  bool
	Format        TexFormat
	Flags         TexFlags
}
