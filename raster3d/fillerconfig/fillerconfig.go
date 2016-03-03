// This package is shared between raster3d and generator
// It must not import raster3d for circular dependencies
package fillerconfig

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

const (
	FillModeSolid uint = iota
	FillModeAlpha
	FillModeWireframe
)

type FillerConfig struct {
	TexFormat uint // 3 bits (0..7)
	ColorKey  bool // 1 bit
	FillMode  uint // 2 bits (0=solid, 1=alpha, 2=wireframe)
}

const FillerKeyBits = 3 + 1 + 2

func (cfg *FillerConfig) Palettized() bool {
	switch cfg.TexFormat {
	case Tex4, Tex16, Tex256:
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
	return
}

func FillerConfigFromKey(k uint) (cfg FillerConfig) {
	cfg.TexFormat = (k & 7)
	if k&(1<<3) != 0 {
		cfg.ColorKey = true
	}
	cfg.FillMode = (k >> 4) & 3
	return
}
