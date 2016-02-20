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

type FillerConfig struct {
	TexFormat uint // 3 bits (0..7)
	ColorKey  bool // 1 bit
}

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
	return
}

func FillerConfigFromKey(k uint) (cfg FillerConfig) {
	cfg.TexFormat = (k & 7)
	if k&(1<<3) != 0 {
		cfg.ColorKey = true
	}
	return
}
