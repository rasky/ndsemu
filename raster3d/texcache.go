package raster3d

import (
	"fmt"
	"image"
	icolor "image/color"
	"image/png"
	"ndsemu/emu"
	"os"
)

// Cache holding decompressed textures. It is a locked dictionary,
// indexed by the texture address in VRAM, and containing the texture
// raw bits.
// FIXME: possibile improvements:
//   * As a key, use a fast hash of texture bits (eg: crc64); this would
//     allow reusing the same texture across different frames. We need to
//     benchmark whether it's a net win
//   * Once we switch to the above, we could use a proper LRU cache to
//     also decrease pressure on the GC; as things stand, texture buffers
//     are allocated / deallocated 60 times per second.
type texCache struct {
	data map[uint32][]uint8
}

func (d *texCache) Reset() {
	d.data = make(map[uint32][]uint8)
}

func (d *texCache) Get(pos uint32) []uint8 {
	tex := d.data[pos]
	return tex
}

func (d *texCache) Put(pos uint32, tex []uint8) {
	d.data[pos] = tex
}

func (cache *texCache) Update(polys []Polygon, e3d *HwEngine3d) {
	cache.Reset()

	for idx := range polys {
		poly := &polys[idx]
		decompFunc := decompTexFuncs[poly.tex.Format]
		if decompFunc == nil {
			continue
		}

		off := poly.tex.VramTexOffset
		if buf := cache.Get(off); buf != nil {
			if len(buf) != int((poly.tex.Width)*(poly.tex.Height)*2) {
				panic("different compressed texture size in same frame")
			}
			continue
		}

		out := decompFunc(cache, poly, e3d)

		if false {
			f, err := os.Create(fmt.Sprintf("tex-%x.png", poly.tex.VramTexOffset))
			if err == nil {
				png.Encode(f, &Image555{
					buf: out,
					w:   int(poly.tex.Width),
					h:   int(poly.tex.Height),
				})
				f.Close()
			}
		}

		cache.Put(poly.tex.VramTexOffset, out)
	}
}

var decompTexFuncs = map[TexFormat]func(*texCache, *Polygon, *HwEngine3d) []byte{
	Tex4x4: (*texCache).decompTex4x4,
}

func (cache *texCache) decompTex4x4(poly *Polygon, e3d *HwEngine3d) []byte {
	off := poly.tex.VramTexOffset
	out := make([]uint8, (poly.tex.Width)*(poly.tex.Height)*2)

	var xtraoff uint32
	switch off / (128 * 1024) {
	case 0:
		xtraoff = 128*1024 + off/2
	case 2:
		xtraoff = 128*1024 + (off-2*128*1024)/2 + 0x10000
	default:
		xtraoff = 128 * 1024
		panic("compressed texture in wrong slot?")
	}

	mod3d.Infof("tex:%d, xtraoff:%d, size:%d,%d",
		off, xtraoff, poly.tex.Width, poly.tex.Height)

	for y := 0; y < int(poly.tex.Height); y += 4 {
		for x := 0; x < int(poly.tex.Width); x += 4 {
			xtra := e3d.texVram.Get16(xtraoff)
			xtraoff += 2

			mode := xtra >> 14
			paloff := uint32(xtra & 0x3FFF)
			pal := e3d.palVram.Palette(int(poly.tex.VramPalOffset + paloff*4))

			var colors [4]uint16
			colors[0] = pal.Lookup(0)
			colors[1] = pal.Lookup(1)
			switch mode {
			case 0:
				colors[2] = pal.Lookup(1)
			case 1:
				colors[2] = rgbMix(colors[0], 1, colors[1], 1)
			case 2:
				colors[2] = pal.Lookup(2)
				colors[3] = pal.Lookup(3)
			case 3:
				colors[2] = rgbMix(colors[0], 5, colors[1], 3)
				colors[3] = rgbMix(colors[0], 3, colors[1], 5)
			}

			for j := 0; j < 4; j++ {
				pack := e3d.texVram.Get8(off)
				off++
				for i := 0; i < 4; i++ {
					tex := (pack >> uint(i*2)) & 3
					emu.Write16LE(out[((y+j)<<poly.tex.PitchShift+(x+i))*2:], colors[tex])
				}
			}
		}
	}

	return out
}

// Implement image.Image interface for a linear rgb555 buffer
type Image555 struct {
	buf  []uint8
	w, h int
}

func (i *Image555) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i *Image555) ColorModel() icolor.Model {
	return icolor.RGBAModel
}

func (i *Image555) At(x, y int) icolor.Color {
	c0 := emu.Read16LE(i.buf[(y*i.w+x)*2:])

	var c1 icolor.RGBA
	c1.R = uint8(c0>>0) & 0x1F
	c1.G = uint8(c0>>5) & 0x1F
	c1.B = uint8(c0>>10) & 0x1F
	c1.R = (c1.R << 3) | (c1.R >> 2)
	c1.G = (c1.G << 3) | (c1.G >> 2)
	c1.B = (c1.B << 3) | (c1.B >> 2)
	c1.A = 0xFF

	return c1
}
