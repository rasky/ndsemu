package raster3d

import (
	"ndsemu/emu/fixed"
)

// RGB color (6 bit precisions, spaced 8-bits each one).
// This matches the internal precision used by the NDS hardware
// when handling colors in rasterizer
type color uint32

func newColorFrom666(r, g, b uint8) color {
	return color(r) | color(g)<<8 | color(b)<<16
}
func newColorFrom555(r, g, b uint8) color {
	r = (r * 2) + (r+31)/32
	g = (g * 2) + (g+31)/32
	b = (b * 2) + (b+31)/32
	return color(r) | color(g)<<8 | color(b)<<16
}

func newColorFrom555U(rgb uint16) color {
	return newColorFrom555(uint8(rgb&0x1F), uint8(rgb>>5)&0x1F, uint8(rgb>>10)&0x1F)
}

func (c color) To555U() uint16 {
	r, g, b := c.R()>>1, c.G()>>1, c.B()>>1
	return uint16(r) | uint16(g)<<5 | uint16(b)<<10
}

func (c color) R() uint8 { return uint8(c >> 0) }
func (c color) G() uint8 { return uint8(c >> 8) }
func (c color) B() uint8 { return uint8(c >> 16) }

func (c color) RF32() fixed.F32 { return fixed.NewF32(int32(c.R())) }
func (c color) GF32() fixed.F32 { return fixed.NewF32(int32(c.G())) }
func (c color) BF32() fixed.F32 { return fixed.NewF32(int32(c.B())) }

// Compute a subtraction component-wise.
// The result of this computation is not a real color, but can be used to computer
// an interpolater between two colors.
func (c1 color) SubColor(c2 color) (fixed.F32, fixed.F32, fixed.F32) {
	r1, g1, b1 := int16(c1.R()), int16(c1.G()), int16(c1.B())
	r2, g2, b2 := int16(c2.R()), int16(c2.G()), int16(c2.B())

	rdiff := (r1 - r2)
	gdiff := (g1 - g2)
	bdiff := (b1 - b2)
	return fixed.NewF32(int32(rdiff)), fixed.NewF32(int32(gdiff)), fixed.NewF32(int32(bdiff))
}

func (c1 color) Decal(c2 color, alpha uint8) color {
	r1, g1, b1 := int32(c1.R()), int32(c1.G()), int32(c1.B())
	r2, g2, b2 := int32(c2.R()), int32(c2.G()), int32(c2.B())

	r := (r1*int32(alpha) + r2*(63-int32(alpha))) >> 6
	g := (g1*int32(alpha) + g2*(63-int32(alpha))) >> 6
	b := (b1*int32(alpha) + b2*(63-int32(alpha))) >> 6

	return newColorFrom666(uint8(r), uint8(g), uint8(b))
}

func (c1 color) Lerp(c2 color, ratio fixed.F12) color {
	r1, g1, b1 := int32(c1.R()), int32(c1.G()), int32(c1.B())
	r2, g2, b2 := int32(c2.R()), int32(c2.G()), int32(c2.B())

	r := r1 + fixed.NewF12(r2-r1).MulFixed(ratio).NearInt32()
	g := g1 + fixed.NewF12(g2-g1).MulFixed(ratio).NearInt32()
	b := b1 + fixed.NewF12(b2-b1).MulFixed(ratio).NearInt32()

	return newColorFrom666(uint8(r), uint8(g), uint8(b))
}

func rgbMix(c1 uint16, f1 int, c2 uint16, f2 int) uint16 {
	r1, g1, b1 := (c1 & 0x1F), ((c1 >> 5) & 0x1F), ((c1 >> 10) & 0x1F)
	r2, g2, b2 := (c2 & 0x1F), ((c2 >> 5) & 0x1F), ((c2 >> 10) & 0x1F)

	r := (int(r1)*f1 + int(r2)*f2) / (f1 + f2)
	g := (int(g1)*f1 + int(g2)*f2) / (f1 + f2)
	b := (int(b1)*f1 + int(b2)*f2) / (f1 + f2)

	return uint16(r) | uint16(g<<5) | uint16(b<<10)
}

func rgbAlphaMix(c1 uint16, c2 uint16, alpha uint8) uint16 {
	r1, g1, b1 := (c1 & 0x1F), ((c1 >> 5) & 0x1F), ((c1 >> 10) & 0x1F)
	r2, g2, b2 := (c2 & 0x1F), ((c2 >> 5) & 0x1F), ((c2 >> 10) & 0x1F)

	a1 := uint(alpha + 1)
	a2 := uint(31 - alpha)
	r := (uint(r1)*a1 + uint(r2)*a2) >> 5
	g := (uint(g1)*a1 + uint(g2)*a2) >> 5
	b := (uint(b1)*a1 + uint(b2)*a2) >> 5

	return uint16(r) | uint16(g)<<5 | uint16(b)<<10
}
