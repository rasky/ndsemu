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
	ColorModeHighlight
)

// Fill mode (how the polygon is filled)
const (
	FillModeSolid uint = iota
	FillModeAlpha
	FillModeWireframe
)

type FillerConfig struct {
	TexFormat uint // 3 bits (0..7)
	ColorKey  bool // 1 bit
	FillMode  uint // 2 bits (0=solid, 1=alpha, 2=wireframe)
	ColorMode uint // 2 bits (0=modul, 1=decal, 2=toon, 3=highlight)
}

const FillerKeyBits = 3 + 1 + 2 + 2

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

func (cfg *FillerConfig) Key() (k uint) {
	k |= (cfg.TexFormat & 7) << 0
	if cfg.ColorKey {
		k |= 1 << 3
	}
	k |= (cfg.FillMode & 3) << 4
	k |= (cfg.ColorMode & 3) << 6
	return
}

func FillerConfigFromKey(k uint) (cfg FillerConfig) {
	cfg.TexFormat = (k & 7)
	if k&(1<<3) != 0 {
		cfg.ColorKey = true
	}
	cfg.FillMode = (k >> 4) & 3
	cfg.ColorMode = (k >> 6) & 3
	return
}
