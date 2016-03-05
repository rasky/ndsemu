package raster3d

// RGB color (6 bit precisions, spaced 8-bits each one).
// This matches the internal precision used by the NDS hardware
// when handling colors in rasterizer
type color int32

type colorDelta int32

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

// Compute a subtraction component-wise.
// The result of this computation is not a real color, but can be used to computer
// an interpolater between two colors.
func (c1 color) SubColor(c2 color) colorDelta {
	r1, g1, b1 := int32(c1.R()), int32(c1.G()), int32(c1.B())
	r2, g2, b2 := int32(c2.R()), int32(c2.G()), int32(c2.B())

	rdiff := r1 - r2
	gdiff := g1 - g2
	bdiff := b1 - b2

	return colorDelta(rdiff + (gdiff << 8) | (bdiff << 16))
}

func (c1 color) AddDelta(d colorDelta) color {
	return c1 + color(d)
}

func (c colorDelta) Div(n int32) colorDelta {
	return c / colorDelta(n)
}

func (c1 color) Modulate(c2 color) color {
	r1, g1, b1 := int32(c1.R()), int32(c1.G()), int32(c1.B())
	r2, g2, b2 := int32(c2.R()), int32(c2.G()), int32(c2.B())

	r := ((r1+1)*(r2+1) - 1) >> 6
	g := ((g1+1)*(g2+1) - 1) >> 6
	b := ((b1+1)*(b2+1) - 1) >> 6

	return newColorFrom666(uint8(r), uint8(g), uint8(b))
}

func (c1 color) Decal(c2 color, alpha uint8) color {
	r1, g1, b1 := int32(c1.R()), int32(c1.G()), int32(c1.B())
	r2, g2, b2 := int32(c2.R()), int32(c2.G()), int32(c2.B())

	r := (r1*int32(alpha) + r2*(63-int32(alpha))) >> 6
	g := (g1*int32(alpha) + g2*(63-int32(alpha))) >> 6
	b := (b1*int32(alpha) + b2*(63-int32(alpha))) >> 6

	return newColorFrom666(uint8(r), uint8(g), uint8(b))
}

func (c1 color) AddSat(c2 color) color {
	r1, g1, b1 := int32(c1.R()), int32(c1.G()), int32(c1.B())
	r2, g2, b2 := int32(c2.R()), int32(c2.G()), int32(c2.B())

	r := r1 + r2
	g := g1 + g2
	b := b1 + b2
	if r > 63 {
		r = 63
	}
	if g > 63 {
		g = 63
	}
	if b > 63 {
		b = 63
	}

	return newColorFrom666(uint8(r), uint8(g), uint8(b))
}
