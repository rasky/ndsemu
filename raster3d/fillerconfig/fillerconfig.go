// This package is shared between raster3d and generator
// It must not import raster3d for circular dependencies
package fillerconfig

// Texture formats
const (
	TexNone uint = iota
	TexA3I5
	Tex4
	Tex16
	Tex256
	Tex4x4
	TexA5I3
	TexDirect
)

// Color mode (how vertex color and texture pixels are combined)
const (
	ColorModeModulation uint = iota
	ColorModeDecal
	ColorModeToon
	ColorModeShadow
	ColorModeHighlight
)

// Fill mode (how the polygon is filled)
const (
	FillModeSolid uint = iota
	FillModeAlpha
	FillModeWireframe
)

// Tex Coords mode. We don't support different modes for s/t coordinates
// to avoid combinatorial explosion, so we support 3 different modes, in
// increasing order of speed:
//
// * Full: supports all hw modes (clamp, repeat, or repeat+flip)
// * RepeatOrFlip: doesn't support clamp (implements repeat or repeat+flip)
// * RepeatOnly: supports only repeat
//
// So for instance, RepeatOnly can be used only if *both* s/t are in
// repeat mode; if s is configured in repeat, and t is configured in clamp,
// Full must be used.
const (
	TexCoordsFull uint = iota // supports any combination of
	TexCoordsRepeatOrFlip
	TexCoordsRepeatOnly
)

type FillerConfig struct {
	TexFormat uint // 8 values (0..7)
	ColorKey  uint // 2 values (0..1)
	FillMode  uint // 4 values (0=solid, 1=alpha, 2=wireframe)
	ColorMode uint // 5 values (0=modul, 1=decal, 2=toon, 3=shadow, 4=highlight)
	TexCoords uint // 3 values (0=full, 1=noclamp, 2=onlywrap)
}

const FillerKeyMax = 8 * 2 * 4 * 5 * 3

func (cfg *FillerConfig) Palettized() bool {
	switch cfg.TexFormat {
	case Tex4, Tex16, Tex256:
		return true
	default:
		return false
	}
}

func (cfg *FillerConfig) TexWithAlpha() bool {
	switch cfg.TexFormat {
	case TexA3I5, TexA5I3, TexDirect:
		return true
	default:
		return false
	}
}

func (cfg *FillerConfig) Key() uint {
	k := uint(cfg.ColorMode & 7)
	k = (k * 4) + (cfg.FillMode & 3)
	k = (k * 2) + (cfg.ColorKey & 1)
	k = (k * 8) + (cfg.TexFormat & 7)
	k = (k * 3) + (cfg.TexCoords & 3)
	return k
}

func FillerConfigFromKey(k uint) (cfg FillerConfig) {
	cfg.TexCoords = k % 3
	k /= 3
	cfg.TexFormat = k % 8
	k /= 8
	cfg.ColorKey = k % 2
	k /= 2
	cfg.FillMode = k % 4
	k /= 4
	cfg.ColorMode = k % 5
	return
}
