package raster3d

import (
	"ndsemu/emu/fixed"
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
	RVFDepth       // prepared for perspective correction

	RVFClipMask = (RVFClipLeft | RVFClipRight | RVFClipTop | RVFClipBottom | RVFClipNear | RVFClipFar)
)

type Vertex struct {
	// Coordinates in clip-space
	cx, cy, cz, cw fixed.F12

	// Screen coordinates (fractional part is always zero)
	x, y fixed.F12

	// Texture coordinates
	s, t fixed.F32

	// Vertex color
	r, g, b fixed.F12

	// Depth coordinate. This is the coordinate to do perspective
	// correction; it can either be W or 1/Z
	d fixed.F32

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

type PolygonColorMode uint32

const (
	PCMModulate PolygonColorMode = 0
	PCMDecal                     = 1
	PCMToon                      = 2
	PCMShadow                    = 3
)

func (f PolygonFlags) Alpha() int                  { return int(f>>16) & 0x1F }
func (f PolygonFlags) ColorMode() PolygonColorMode { return PolygonColorMode(f>>4) & 3 }

const (
	LerpX = iota // coordinate on screen (X)
	LerpD        // inverse depth for perspective correction (1/Z or W)
	LerpT        // texture X coordinate (T)
	LerpS        // texture Y coordinate (S)
	LerpR        // vertex color (R component)
	LerpG        // vertex color (G component)
	LerpB        // vertex color (B component)
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

	ColorKey bool
	Format   TexFormat
	Flags    TexFlags
}

func (p *Polygon) UseAlpha() bool {
	// Check if the texture contains an alpha value (and we're using it)
	mode := p.flags.ColorMode()
	if (p.tex.Format == TexA5I3 || p.tex.Format == TexA3I5) &&
		(mode == PCMModulate || mode == PCMToon) {
		return true
	}

	// Check if the vertex/polygon alpha is semi-transaparent
	alpha := p.flags.Alpha()
	return alpha > 0 && alpha < 31
}
