// Generated on 2016-04-12 23:59:19.530178744 +0200 CEST
package raster3d

import "ndsemu/emu/gfx"
import "ndsemu/emu"

func (e3d *HwEngine3d) filler_000(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_001 skipped, because of identical polyfiller:
//     001 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_002 skipped, because of identical polyfiller:
//     002 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

func (e3d *HwEngine3d) filler_003(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_004(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_005(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_006(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_007(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_008(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_009(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_00a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_00b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_00c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_00d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_00e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_00f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_010(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_011(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_012(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_013(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_014(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_015(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_016(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_017(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_018 skipped, because of identical polyfiller:
//     018 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_019 skipped, because of identical polyfiller:
//     019 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_01a skipped, because of identical polyfiller:
//     01a -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

func (e3d *HwEngine3d) filler_01b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_01c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_01d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_01e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_01f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_020(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_021(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_022(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_023(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_024(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_025(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_026(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_027 skipped, because of identical polyfiller:
//     027 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}
//     00f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_028 skipped, because of identical polyfiller:
//     028 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}
//     010 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_029 skipped, because of identical polyfiller:
//     029 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}
//     011 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

func (e3d *HwEngine3d) filler_02a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_02b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_02c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_02d skipped, because of identical polyfiller:
//     02d -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
//     015 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_02e skipped, because of identical polyfiller:
//     02e -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
//     016 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_02f skipped, because of identical polyfiller:
//     02f -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
//     017 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

func (e3d *HwEngine3d) filler_030(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_031 skipped, because of identical polyfiller:
//     031 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
//     030 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}

// filler_032 skipped, because of identical polyfiller:
//     032 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
//     030 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}

// filler_033 skipped, because of identical polyfiller:
//     033 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
//     003 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_034 skipped, because of identical polyfiller:
//     034 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
//     004 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_035 skipped, because of identical polyfiller:
//     035 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
//     005 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

func (e3d *HwEngine3d) filler_036(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_037(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_038(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_039(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_03a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_03b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_03c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_03d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_03e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_03f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_040(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_041(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_042 skipped, because of identical polyfiller:
//     042 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
//     012 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_043 skipped, because of identical polyfiller:
//     043 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
//     013 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_044 skipped, because of identical polyfiller:
//     044 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
//     014 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_045 skipped, because of identical polyfiller:
//     045 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}
//     015 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_046 skipped, because of identical polyfiller:
//     046 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}
//     016 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_047 skipped, because of identical polyfiller:
//     047 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}
//     017 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_048 skipped, because of identical polyfiller:
//     048 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
//     030 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}

// filler_049 skipped, because of identical polyfiller:
//     049 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
//     030 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}

// filler_04a skipped, because of identical polyfiller:
//     04a -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
//     030 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}

// filler_04b skipped, because of identical polyfiller:
//     04b -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
//     01b -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_04c skipped, because of identical polyfiller:
//     04c -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
//     01c -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_04d skipped, because of identical polyfiller:
//     04d -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
//     01d -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

func (e3d *HwEngine3d) filler_04e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_04f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_050(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_051(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_052(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_053(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_054(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_055(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_056(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_057 skipped, because of identical polyfiller:
//     057 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
//     03f -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:0}

// filler_058 skipped, because of identical polyfiller:
//     058 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
//     040 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:1}

// filler_059 skipped, because of identical polyfiller:
//     059 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
//     041 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:0 TexCoords:2}

// filler_05a skipped, because of identical polyfiller:
//     05a -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
//     02a -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_05b skipped, because of identical polyfiller:
//     05b -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
//     02b -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_05c skipped, because of identical polyfiller:
//     05c -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
//     02c -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

// filler_05d skipped, because of identical polyfiller:
//     05d -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:0}
//     015 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_05e skipped, because of identical polyfiller:
//     05e -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:1}
//     016 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_05f skipped, because of identical polyfiller:
//     05f -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:0 TexCoords:2}
//     017 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_060 skipped, because of identical polyfiller:
//     060 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_061 skipped, because of identical polyfiller:
//     061 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_062 skipped, because of identical polyfiller:
//     062 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

func (e3d *HwEngine3d) filler_063(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_064(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_065(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_066 skipped, because of identical polyfiller:
//     066 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
//     006 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_067 skipped, because of identical polyfiller:
//     067 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
//     007 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_068 skipped, because of identical polyfiller:
//     068 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
//     008 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_069 skipped, because of identical polyfiller:
//     069 -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
//     009 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_06a skipped, because of identical polyfiller:
//     06a -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
//     00a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_06b skipped, because of identical polyfiller:
//     06b -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
//     00b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_06c skipped, because of identical polyfiller:
//     06c -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
//     00c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_06d skipped, because of identical polyfiller:
//     06d -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
//     00d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_06e skipped, because of identical polyfiller:
//     06e -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
//     00e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_06f skipped, because of identical polyfiller:
//     06f -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
//     00f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_070 skipped, because of identical polyfiller:
//     070 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
//     010 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_071 skipped, because of identical polyfiller:
//     071 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
//     011 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

func (e3d *HwEngine3d) filler_072(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_073(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_074(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_075(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_076(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_077(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_078 skipped, because of identical polyfiller:
//     078 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_079 skipped, because of identical polyfiller:
//     079 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_07a skipped, because of identical polyfiller:
//     07a -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

func (e3d *HwEngine3d) filler_07b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_07c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_07d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_07e skipped, because of identical polyfiller:
//     07e -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
//     01e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_07f skipped, because of identical polyfiller:
//     07f -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
//     01f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_080 skipped, because of identical polyfiller:
//     080 -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
//     020 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

// filler_081 skipped, because of identical polyfiller:
//     081 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
//     021 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_082 skipped, because of identical polyfiller:
//     082 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
//     022 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_083 skipped, because of identical polyfiller:
//     083 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
//     023 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

// filler_084 skipped, because of identical polyfiller:
//     084 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
//     024 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_085 skipped, because of identical polyfiller:
//     085 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
//     025 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_086 skipped, because of identical polyfiller:
//     086 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
//     026 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

// filler_087 skipped, because of identical polyfiller:
//     087 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
//     00f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_088 skipped, because of identical polyfiller:
//     088 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
//     010 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_089 skipped, because of identical polyfiller:
//     089 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
//     011 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

func (e3d *HwEngine3d) filler_08a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_08b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_08c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_08d skipped, because of identical polyfiller:
//     08d -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}
//     075 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}

// filler_08e skipped, because of identical polyfiller:
//     08e -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}
//     076 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}

// filler_08f skipped, because of identical polyfiller:
//     08f -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}
//     077 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}

// filler_090 skipped, because of identical polyfiller:
//     090 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_091 skipped, because of identical polyfiller:
//     091 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_092 skipped, because of identical polyfiller:
//     092 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_093 skipped, because of identical polyfiller:
//     093 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     063 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}

// filler_094 skipped, because of identical polyfiller:
//     094 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     064 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}

// filler_095 skipped, because of identical polyfiller:
//     095 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     065 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}

// filler_096 skipped, because of identical polyfiller:
//     096 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     006 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_097 skipped, because of identical polyfiller:
//     097 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     007 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_098 skipped, because of identical polyfiller:
//     098 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     008 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_099 skipped, because of identical polyfiller:
//     099 -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     009 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_09a skipped, because of identical polyfiller:
//     09a -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     00a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_09b skipped, because of identical polyfiller:
//     09b -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     00b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_09c skipped, because of identical polyfiller:
//     09c -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     00c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_09d skipped, because of identical polyfiller:
//     09d -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     00d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_09e skipped, because of identical polyfiller:
//     09e -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     00e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_09f skipped, because of identical polyfiller:
//     09f -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     00f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0a0 skipped, because of identical polyfiller:
//     0a0 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     010 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_0a1 skipped, because of identical polyfiller:
//     0a1 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     011 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_0a2 skipped, because of identical polyfiller:
//     0a2 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     072 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}

// filler_0a3 skipped, because of identical polyfiller:
//     0a3 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     073 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}

// filler_0a4 skipped, because of identical polyfiller:
//     0a4 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     074 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}

// filler_0a5 skipped, because of identical polyfiller:
//     0a5 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:0}
//     075 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}

// filler_0a6 skipped, because of identical polyfiller:
//     0a6 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:1}
//     076 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}

// filler_0a7 skipped, because of identical polyfiller:
//     0a7 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:0 TexCoords:2}
//     077 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}

// filler_0a8 skipped, because of identical polyfiller:
//     0a8 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0a9 skipped, because of identical polyfiller:
//     0a9 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0aa skipped, because of identical polyfiller:
//     0aa -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     000 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0ab skipped, because of identical polyfiller:
//     0ab -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     07b -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}

// filler_0ac skipped, because of identical polyfiller:
//     0ac -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     07c -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}

// filler_0ad skipped, because of identical polyfiller:
//     0ad -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     07d -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}

// filler_0ae skipped, because of identical polyfiller:
//     0ae -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     01e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0af skipped, because of identical polyfiller:
//     0af -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     01f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_0b0 skipped, because of identical polyfiller:
//     0b0 -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     020 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

// filler_0b1 skipped, because of identical polyfiller:
//     0b1 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     021 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0b2 skipped, because of identical polyfiller:
//     0b2 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     022 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_0b3 skipped, because of identical polyfiller:
//     0b3 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     023 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

// filler_0b4 skipped, because of identical polyfiller:
//     0b4 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     024 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0b5 skipped, because of identical polyfiller:
//     0b5 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     025 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:1}

// filler_0b6 skipped, because of identical polyfiller:
//     0b6 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     026 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0 TexCoords:2}

// filler_0b7 skipped, because of identical polyfiller:
//     0b7 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     00f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:0}

// filler_0b8 skipped, because of identical polyfiller:
//     0b8 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     010 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:1}

// filler_0b9 skipped, because of identical polyfiller:
//     0b9 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     011 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0 TexCoords:2}

// filler_0ba skipped, because of identical polyfiller:
//     0ba -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     08a -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:0}

// filler_0bb skipped, because of identical polyfiller:
//     0bb -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     08b -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:1}

// filler_0bc skipped, because of identical polyfiller:
//     0bc -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     08c -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:0 TexCoords:2}

// filler_0bd skipped, because of identical polyfiller:
//     0bd -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:0}
//     075 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:0}

// filler_0be skipped, because of identical polyfiller:
//     0be -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:1}
//     076 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:1}

// filler_0bf skipped, because of identical polyfiller:
//     0bf -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:0 TexCoords:2}
//     077 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0 TexCoords:2}

func (e3d *HwEngine3d) filler_0c0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_0c1 skipped, because of identical polyfiller:
//     0c1 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_0c2 skipped, because of identical polyfiller:
//     0c2 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

func (e3d *HwEngine3d) filler_0c3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0c4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0c5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0c6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0c7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0c8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0c9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0ca(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0cb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0cc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0cd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0ce(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0cf(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0d7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_0d8 skipped, because of identical polyfiller:
//     0d8 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_0d9 skipped, because of identical polyfiller:
//     0d9 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_0da skipped, because of identical polyfiller:
//     0da -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

func (e3d *HwEngine3d) filler_0db(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0dc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0dd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0de(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0df(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0e0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0e1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0e2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0e3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0e4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0e5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0e6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_0e7 skipped, because of identical polyfiller:
//     0e7 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}
//     0cf -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_0e8 skipped, because of identical polyfiller:
//     0e8 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}
//     0d0 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_0e9 skipped, because of identical polyfiller:
//     0e9 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}
//     0d1 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

func (e3d *HwEngine3d) filler_0ea(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0eb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0ec(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_0ed skipped, because of identical polyfiller:
//     0ed -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
//     0d5 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_0ee skipped, because of identical polyfiller:
//     0ee -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
//     0d6 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_0ef skipped, because of identical polyfiller:
//     0ef -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
//     0d7 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

func (e3d *HwEngine3d) filler_0f0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_0f1 skipped, because of identical polyfiller:
//     0f1 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
//     0f0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}

// filler_0f2 skipped, because of identical polyfiller:
//     0f2 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
//     0f0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}

// filler_0f3 skipped, because of identical polyfiller:
//     0f3 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
//     0c3 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_0f4 skipped, because of identical polyfiller:
//     0f4 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
//     0c4 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_0f5 skipped, because of identical polyfiller:
//     0f5 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
//     0c5 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

func (e3d *HwEngine3d) filler_0f6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0f7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0f8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0f9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0fa(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0fb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0fc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0fd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0fe(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_0ff(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_100(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_101(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_102 skipped, because of identical polyfiller:
//     102 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
//     0d2 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_103 skipped, because of identical polyfiller:
//     103 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
//     0d3 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_104 skipped, because of identical polyfiller:
//     104 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
//     0d4 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_105 skipped, because of identical polyfiller:
//     105 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}
//     0d5 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_106 skipped, because of identical polyfiller:
//     106 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}
//     0d6 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_107 skipped, because of identical polyfiller:
//     107 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}
//     0d7 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_108 skipped, because of identical polyfiller:
//     108 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
//     0f0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}

// filler_109 skipped, because of identical polyfiller:
//     109 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
//     0f0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}

// filler_10a skipped, because of identical polyfiller:
//     10a -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
//     0f0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}

// filler_10b skipped, because of identical polyfiller:
//     10b -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
//     0db -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_10c skipped, because of identical polyfiller:
//     10c -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
//     0dc -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_10d skipped, because of identical polyfiller:
//     10d -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
//     0dd -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

func (e3d *HwEngine3d) filler_10e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_10f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_110(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_111(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_112(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_113(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_114(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_115(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_116(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_117 skipped, because of identical polyfiller:
//     117 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
//     0ff -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:0}

// filler_118 skipped, because of identical polyfiller:
//     118 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
//     100 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:1}

// filler_119 skipped, because of identical polyfiller:
//     119 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
//     101 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:1 TexCoords:2}

// filler_11a skipped, because of identical polyfiller:
//     11a -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
//     0ea -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_11b skipped, because of identical polyfiller:
//     11b -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
//     0eb -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_11c skipped, because of identical polyfiller:
//     11c -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
//     0ec -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

// filler_11d skipped, because of identical polyfiller:
//     11d -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:0}
//     0d5 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_11e skipped, because of identical polyfiller:
//     11e -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:1}
//     0d6 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_11f skipped, because of identical polyfiller:
//     11f -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:1 TexCoords:2}
//     0d7 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_120 skipped, because of identical polyfiller:
//     120 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_121 skipped, because of identical polyfiller:
//     121 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_122 skipped, because of identical polyfiller:
//     122 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

func (e3d *HwEngine3d) filler_123(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_124(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_125(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_126 skipped, because of identical polyfiller:
//     126 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
//     0c6 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_127 skipped, because of identical polyfiller:
//     127 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
//     0c7 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_128 skipped, because of identical polyfiller:
//     128 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
//     0c8 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_129 skipped, because of identical polyfiller:
//     129 -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
//     0c9 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_12a skipped, because of identical polyfiller:
//     12a -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
//     0ca -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_12b skipped, because of identical polyfiller:
//     12b -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
//     0cb -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_12c skipped, because of identical polyfiller:
//     12c -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
//     0cc -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_12d skipped, because of identical polyfiller:
//     12d -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
//     0cd -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_12e skipped, because of identical polyfiller:
//     12e -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
//     0ce -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_12f skipped, because of identical polyfiller:
//     12f -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
//     0cf -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_130 skipped, because of identical polyfiller:
//     130 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
//     0d0 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_131 skipped, because of identical polyfiller:
//     131 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
//     0d1 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

func (e3d *HwEngine3d) filler_132(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_133(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_134(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_135(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_136(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_137(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_138 skipped, because of identical polyfiller:
//     138 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_139 skipped, because of identical polyfiller:
//     139 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_13a skipped, because of identical polyfiller:
//     13a -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

func (e3d *HwEngine3d) filler_13b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_13c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_13d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_13e skipped, because of identical polyfiller:
//     13e -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
//     0de -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_13f skipped, because of identical polyfiller:
//     13f -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
//     0df -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_140 skipped, because of identical polyfiller:
//     140 -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
//     0e0 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

// filler_141 skipped, because of identical polyfiller:
//     141 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
//     0e1 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_142 skipped, because of identical polyfiller:
//     142 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
//     0e2 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_143 skipped, because of identical polyfiller:
//     143 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
//     0e3 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

// filler_144 skipped, because of identical polyfiller:
//     144 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
//     0e4 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_145 skipped, because of identical polyfiller:
//     145 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
//     0e5 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_146 skipped, because of identical polyfiller:
//     146 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
//     0e6 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

// filler_147 skipped, because of identical polyfiller:
//     147 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
//     0cf -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_148 skipped, because of identical polyfiller:
//     148 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
//     0d0 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_149 skipped, because of identical polyfiller:
//     149 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
//     0d1 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

func (e3d *HwEngine3d) filler_14a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_14b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_14c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_14d skipped, because of identical polyfiller:
//     14d -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}
//     135 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}

// filler_14e skipped, because of identical polyfiller:
//     14e -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}
//     136 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}

// filler_14f skipped, because of identical polyfiller:
//     14f -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}
//     137 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}

// filler_150 skipped, because of identical polyfiller:
//     150 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_151 skipped, because of identical polyfiller:
//     151 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_152 skipped, because of identical polyfiller:
//     152 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_153 skipped, because of identical polyfiller:
//     153 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     123 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}

// filler_154 skipped, because of identical polyfiller:
//     154 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     124 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}

// filler_155 skipped, because of identical polyfiller:
//     155 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     125 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}

// filler_156 skipped, because of identical polyfiller:
//     156 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     0c6 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_157 skipped, because of identical polyfiller:
//     157 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     0c7 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_158 skipped, because of identical polyfiller:
//     158 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     0c8 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_159 skipped, because of identical polyfiller:
//     159 -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     0c9 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_15a skipped, because of identical polyfiller:
//     15a -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     0ca -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_15b skipped, because of identical polyfiller:
//     15b -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     0cb -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_15c skipped, because of identical polyfiller:
//     15c -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     0cc -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_15d skipped, because of identical polyfiller:
//     15d -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     0cd -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_15e skipped, because of identical polyfiller:
//     15e -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     0ce -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_15f skipped, because of identical polyfiller:
//     15f -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     0cf -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_160 skipped, because of identical polyfiller:
//     160 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     0d0 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_161 skipped, because of identical polyfiller:
//     161 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     0d1 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_162 skipped, because of identical polyfiller:
//     162 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     132 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}

// filler_163 skipped, because of identical polyfiller:
//     163 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     133 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}

// filler_164 skipped, because of identical polyfiller:
//     164 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     134 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}

// filler_165 skipped, because of identical polyfiller:
//     165 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:0}
//     135 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}

// filler_166 skipped, because of identical polyfiller:
//     166 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:1}
//     136 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}

// filler_167 skipped, because of identical polyfiller:
//     167 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:1 TexCoords:2}
//     137 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}

// filler_168 skipped, because of identical polyfiller:
//     168 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_169 skipped, because of identical polyfiller:
//     169 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_16a skipped, because of identical polyfiller:
//     16a -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     0c0 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_16b skipped, because of identical polyfiller:
//     16b -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     13b -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}

// filler_16c skipped, because of identical polyfiller:
//     16c -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     13c -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}

// filler_16d skipped, because of identical polyfiller:
//     16d -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     13d -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}

// filler_16e skipped, because of identical polyfiller:
//     16e -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     0de -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_16f skipped, because of identical polyfiller:
//     16f -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     0df -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_170 skipped, because of identical polyfiller:
//     170 -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     0e0 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

// filler_171 skipped, because of identical polyfiller:
//     171 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     0e1 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_172 skipped, because of identical polyfiller:
//     172 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     0e2 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_173 skipped, because of identical polyfiller:
//     173 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     0e3 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

// filler_174 skipped, because of identical polyfiller:
//     174 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     0e4 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:0}

// filler_175 skipped, because of identical polyfiller:
//     175 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     0e5 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:1}

// filler_176 skipped, because of identical polyfiller:
//     176 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     0e6 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1 TexCoords:2}

// filler_177 skipped, because of identical polyfiller:
//     177 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     0cf -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:0}

// filler_178 skipped, because of identical polyfiller:
//     178 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     0d0 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:1}

// filler_179 skipped, because of identical polyfiller:
//     179 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     0d1 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1 TexCoords:2}

// filler_17a skipped, because of identical polyfiller:
//     17a -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     14a -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:0}

// filler_17b skipped, because of identical polyfiller:
//     17b -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     14b -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:1}

// filler_17c skipped, because of identical polyfiller:
//     17c -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     14c -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:1 TexCoords:2}

// filler_17d skipped, because of identical polyfiller:
//     17d -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:0}
//     135 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:0}

// filler_17e skipped, because of identical polyfiller:
//     17e -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:1}
//     136 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:1}

// filler_17f skipped, because of identical polyfiller:
//     17f -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:1 TexCoords:2}
//     137 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1 TexCoords:2}

func (e3d *HwEngine3d) filler_180(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_181 skipped, because of identical polyfiller:
//     181 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_182 skipped, because of identical polyfiller:
//     182 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

func (e3d *HwEngine3d) filler_183(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_184(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_185(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_186(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_187(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_188(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_189(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_18a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_18b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_18c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_18d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_18e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_18f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_190(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_191(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_192(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_193(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_194(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_195(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_196(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_197(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_198 skipped, because of identical polyfiller:
//     198 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_199 skipped, because of identical polyfiller:
//     199 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_19a skipped, because of identical polyfiller:
//     19a -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

func (e3d *HwEngine3d) filler_19b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_19c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_19d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_19e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_19f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1a0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1a1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1a2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1a3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1a4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1a5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1a6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_1a7 skipped, because of identical polyfiller:
//     1a7 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}
//     18f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1a8 skipped, because of identical polyfiller:
//     1a8 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}
//     190 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1a9 skipped, because of identical polyfiller:
//     1a9 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}
//     191 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

func (e3d *HwEngine3d) filler_1aa(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1ab(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1ac(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_1ad skipped, because of identical polyfiller:
//     1ad -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
//     195 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1ae skipped, because of identical polyfiller:
//     1ae -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
//     196 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1af skipped, because of identical polyfiller:
//     1af -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
//     197 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

func (e3d *HwEngine3d) filler_1b0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_1b1 skipped, because of identical polyfiller:
//     1b1 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
//     1b0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}

// filler_1b2 skipped, because of identical polyfiller:
//     1b2 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
//     1b0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}

// filler_1b3 skipped, because of identical polyfiller:
//     1b3 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
//     183 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1b4 skipped, because of identical polyfiller:
//     1b4 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
//     184 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1b5 skipped, because of identical polyfiller:
//     1b5 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
//     185 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

func (e3d *HwEngine3d) filler_1b6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1b7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1b8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1b9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1ba(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1bb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1bc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1bd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1be(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1bf(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1c0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1c1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_1c2 skipped, because of identical polyfiller:
//     1c2 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
//     192 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1c3 skipped, because of identical polyfiller:
//     1c3 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
//     193 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1c4 skipped, because of identical polyfiller:
//     1c4 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
//     194 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_1c5 skipped, because of identical polyfiller:
//     1c5 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}
//     195 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1c6 skipped, because of identical polyfiller:
//     1c6 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}
//     196 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1c7 skipped, because of identical polyfiller:
//     1c7 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}
//     197 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_1c8 skipped, because of identical polyfiller:
//     1c8 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
//     1b0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}

// filler_1c9 skipped, because of identical polyfiller:
//     1c9 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
//     1b0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}

// filler_1ca skipped, because of identical polyfiller:
//     1ca -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
//     1b0 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}

// filler_1cb skipped, because of identical polyfiller:
//     1cb -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
//     19b -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1cc skipped, because of identical polyfiller:
//     1cc -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
//     19c -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1cd skipped, because of identical polyfiller:
//     1cd -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
//     19d -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

func (e3d *HwEngine3d) filler_1ce(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1cf(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1d0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1d1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1d2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1d3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1d4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1d5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1d6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_1d7 skipped, because of identical polyfiller:
//     1d7 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
//     1bf -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:0}

// filler_1d8 skipped, because of identical polyfiller:
//     1d8 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
//     1c0 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:1}

// filler_1d9 skipped, because of identical polyfiller:
//     1d9 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
//     1c1 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:2 TexCoords:2}

// filler_1da skipped, because of identical polyfiller:
//     1da -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
//     1aa -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1db skipped, because of identical polyfiller:
//     1db -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
//     1ab -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1dc skipped, because of identical polyfiller:
//     1dc -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
//     1ac -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

// filler_1dd skipped, because of identical polyfiller:
//     1dd -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:0}
//     195 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1de skipped, because of identical polyfiller:
//     1de -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:1}
//     196 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1df skipped, because of identical polyfiller:
//     1df -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:2 TexCoords:2}
//     197 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_1e0 skipped, because of identical polyfiller:
//     1e0 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1e1 skipped, because of identical polyfiller:
//     1e1 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1e2 skipped, because of identical polyfiller:
//     1e2 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

func (e3d *HwEngine3d) filler_1e3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1e4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1e5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_1e6 skipped, because of identical polyfiller:
//     1e6 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
//     186 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1e7 skipped, because of identical polyfiller:
//     1e7 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
//     187 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1e8 skipped, because of identical polyfiller:
//     1e8 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
//     188 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_1e9 skipped, because of identical polyfiller:
//     1e9 -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
//     189 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1ea skipped, because of identical polyfiller:
//     1ea -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
//     18a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1eb skipped, because of identical polyfiller:
//     1eb -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
//     18b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_1ec skipped, because of identical polyfiller:
//     1ec -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
//     18c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1ed skipped, because of identical polyfiller:
//     1ed -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
//     18d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1ee skipped, because of identical polyfiller:
//     1ee -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
//     18e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_1ef skipped, because of identical polyfiller:
//     1ef -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
//     18f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1f0 skipped, because of identical polyfiller:
//     1f0 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
//     190 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_1f1 skipped, because of identical polyfiller:
//     1f1 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
//     191 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

func (e3d *HwEngine3d) filler_1f2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1f3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1f4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1f5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1f6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1f7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_1f8 skipped, because of identical polyfiller:
//     1f8 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1f9 skipped, because of identical polyfiller:
//     1f9 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1fa skipped, because of identical polyfiller:
//     1fa -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

func (e3d *HwEngine3d) filler_1fb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1fc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_1fd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_1fe skipped, because of identical polyfiller:
//     1fe -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
//     19e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_1ff skipped, because of identical polyfiller:
//     1ff -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
//     19f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_200 skipped, because of identical polyfiller:
//     200 -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
//     1a0 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

// filler_201 skipped, because of identical polyfiller:
//     201 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
//     1a1 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_202 skipped, because of identical polyfiller:
//     202 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
//     1a2 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_203 skipped, because of identical polyfiller:
//     203 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
//     1a3 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

// filler_204 skipped, because of identical polyfiller:
//     204 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
//     1a4 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_205 skipped, because of identical polyfiller:
//     205 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
//     1a5 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_206 skipped, because of identical polyfiller:
//     206 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
//     1a6 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

// filler_207 skipped, because of identical polyfiller:
//     207 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
//     18f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_208 skipped, because of identical polyfiller:
//     208 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
//     190 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_209 skipped, because of identical polyfiller:
//     209 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
//     191 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

func (e3d *HwEngine3d) filler_20a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_20b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_20c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_20d skipped, because of identical polyfiller:
//     20d -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}
//     1f5 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}

// filler_20e skipped, because of identical polyfiller:
//     20e -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}
//     1f6 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}

// filler_20f skipped, because of identical polyfiller:
//     20f -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}
//     1f7 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}

// filler_210 skipped, because of identical polyfiller:
//     210 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_211 skipped, because of identical polyfiller:
//     211 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_212 skipped, because of identical polyfiller:
//     212 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_213 skipped, because of identical polyfiller:
//     213 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     1e3 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}

// filler_214 skipped, because of identical polyfiller:
//     214 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     1e4 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}

// filler_215 skipped, because of identical polyfiller:
//     215 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     1e5 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}

// filler_216 skipped, because of identical polyfiller:
//     216 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     186 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_217 skipped, because of identical polyfiller:
//     217 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     187 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_218 skipped, because of identical polyfiller:
//     218 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     188 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_219 skipped, because of identical polyfiller:
//     219 -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     189 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_21a skipped, because of identical polyfiller:
//     21a -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     18a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_21b skipped, because of identical polyfiller:
//     21b -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     18b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_21c skipped, because of identical polyfiller:
//     21c -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     18c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_21d skipped, because of identical polyfiller:
//     21d -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     18d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_21e skipped, because of identical polyfiller:
//     21e -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     18e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_21f skipped, because of identical polyfiller:
//     21f -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     18f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_220 skipped, because of identical polyfiller:
//     220 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     190 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_221 skipped, because of identical polyfiller:
//     221 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     191 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_222 skipped, because of identical polyfiller:
//     222 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     1f2 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}

// filler_223 skipped, because of identical polyfiller:
//     223 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     1f3 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}

// filler_224 skipped, because of identical polyfiller:
//     224 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     1f4 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}

// filler_225 skipped, because of identical polyfiller:
//     225 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:0}
//     1f5 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}

// filler_226 skipped, because of identical polyfiller:
//     226 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:1}
//     1f6 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}

// filler_227 skipped, because of identical polyfiller:
//     227 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:2 TexCoords:2}
//     1f7 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}

// filler_228 skipped, because of identical polyfiller:
//     228 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_229 skipped, because of identical polyfiller:
//     229 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_22a skipped, because of identical polyfiller:
//     22a -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     180 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_22b skipped, because of identical polyfiller:
//     22b -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     1fb -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}

// filler_22c skipped, because of identical polyfiller:
//     22c -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     1fc -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}

// filler_22d skipped, because of identical polyfiller:
//     22d -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     1fd -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}

// filler_22e skipped, because of identical polyfiller:
//     22e -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     19e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_22f skipped, because of identical polyfiller:
//     22f -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     19f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_230 skipped, because of identical polyfiller:
//     230 -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     1a0 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

// filler_231 skipped, because of identical polyfiller:
//     231 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     1a1 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_232 skipped, because of identical polyfiller:
//     232 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     1a2 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_233 skipped, because of identical polyfiller:
//     233 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     1a3 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

// filler_234 skipped, because of identical polyfiller:
//     234 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     1a4 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:0}

// filler_235 skipped, because of identical polyfiller:
//     235 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     1a5 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:1}

// filler_236 skipped, because of identical polyfiller:
//     236 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     1a6 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2 TexCoords:2}

// filler_237 skipped, because of identical polyfiller:
//     237 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     18f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:0}

// filler_238 skipped, because of identical polyfiller:
//     238 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     190 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:1}

// filler_239 skipped, because of identical polyfiller:
//     239 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     191 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2 TexCoords:2}

// filler_23a skipped, because of identical polyfiller:
//     23a -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     20a -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:0}

// filler_23b skipped, because of identical polyfiller:
//     23b -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     20b -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:1}

// filler_23c skipped, because of identical polyfiller:
//     23c -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     20c -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:2 TexCoords:2}

// filler_23d skipped, because of identical polyfiller:
//     23d -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:0}
//     1f5 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:0}

// filler_23e skipped, because of identical polyfiller:
//     23e -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:1}
//     1f6 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:1}

// filler_23f skipped, because of identical polyfiller:
//     23f -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:2 TexCoords:2}
//     1f7 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2 TexCoords:2}

func (e3d *HwEngine3d) filler_240(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_241 skipped, because of identical polyfiller:
//     241 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_242 skipped, because of identical polyfiller:
//     242 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

func (e3d *HwEngine3d) filler_243(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_244(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_245(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_246(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_247(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_248(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_249(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_24a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_24b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_24c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_24d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_24e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_24f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_250(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_251(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_252(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_253(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_254(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_255(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_256(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_257(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_258 skipped, because of identical polyfiller:
//     258 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_259 skipped, because of identical polyfiller:
//     259 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_25a skipped, because of identical polyfiller:
//     25a -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

func (e3d *HwEngine3d) filler_25b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_25c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_25d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_25e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_25f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_260(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_261(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_262(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_263(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_264(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_265(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_266(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_267 skipped, because of identical polyfiller:
//     267 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}
//     24f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_268 skipped, because of identical polyfiller:
//     268 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}
//     250 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_269 skipped, because of identical polyfiller:
//     269 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}
//     251 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

func (e3d *HwEngine3d) filler_26a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_26b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_26c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_26d skipped, because of identical polyfiller:
//     26d -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
//     255 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_26e skipped, because of identical polyfiller:
//     26e -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
//     256 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_26f skipped, because of identical polyfiller:
//     26f -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
//     257 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

func (e3d *HwEngine3d) filler_270(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_271 skipped, because of identical polyfiller:
//     271 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
//     270 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}

// filler_272 skipped, because of identical polyfiller:
//     272 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
//     270 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}

// filler_273 skipped, because of identical polyfiller:
//     273 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
//     243 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_274 skipped, because of identical polyfiller:
//     274 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
//     244 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_275 skipped, because of identical polyfiller:
//     275 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
//     245 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

func (e3d *HwEngine3d) filler_276(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_277(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_278(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_279(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_27a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_27b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_27c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_27d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_27e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_27f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_280(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_281(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_282 skipped, because of identical polyfiller:
//     282 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
//     252 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_283 skipped, because of identical polyfiller:
//     283 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
//     253 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_284 skipped, because of identical polyfiller:
//     284 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
//     254 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_285 skipped, because of identical polyfiller:
//     285 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}
//     255 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_286 skipped, because of identical polyfiller:
//     286 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}
//     256 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_287 skipped, because of identical polyfiller:
//     287 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}
//     257 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_288 skipped, because of identical polyfiller:
//     288 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
//     270 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}

// filler_289 skipped, because of identical polyfiller:
//     289 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
//     270 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}

// filler_28a skipped, because of identical polyfiller:
//     28a -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
//     270 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}

// filler_28b skipped, because of identical polyfiller:
//     28b -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
//     25b -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_28c skipped, because of identical polyfiller:
//     28c -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
//     25c -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_28d skipped, because of identical polyfiller:
//     28d -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
//     25d -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

func (e3d *HwEngine3d) filler_28e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_28f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_290(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_291(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_292(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_293(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_294(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_295(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_296(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		px = 0
		pxa = polyalpha
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_297 skipped, because of identical polyfiller:
//     297 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
//     27f -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:0}

// filler_298 skipped, because of identical polyfiller:
//     298 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
//     280 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:1}

// filler_299 skipped, because of identical polyfiller:
//     299 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
//     281 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:3 TexCoords:2}

// filler_29a skipped, because of identical polyfiller:
//     29a -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
//     26a -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_29b skipped, because of identical polyfiller:
//     29b -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
//     26b -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_29c skipped, because of identical polyfiller:
//     29c -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
//     26c -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

// filler_29d skipped, because of identical polyfiller:
//     29d -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:0}
//     255 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_29e skipped, because of identical polyfiller:
//     29e -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:1}
//     256 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_29f skipped, because of identical polyfiller:
//     29f -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:3 TexCoords:2}
//     257 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2a0 skipped, because of identical polyfiller:
//     2a0 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2a1 skipped, because of identical polyfiller:
//     2a1 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2a2 skipped, because of identical polyfiller:
//     2a2 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

func (e3d *HwEngine3d) filler_2a3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2a4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2a5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_2a6 skipped, because of identical polyfiller:
//     2a6 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
//     246 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2a7 skipped, because of identical polyfiller:
//     2a7 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
//     247 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2a8 skipped, because of identical polyfiller:
//     2a8 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
//     248 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2a9 skipped, because of identical polyfiller:
//     2a9 -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
//     249 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2aa skipped, because of identical polyfiller:
//     2aa -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
//     24a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2ab skipped, because of identical polyfiller:
//     2ab -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
//     24b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2ac skipped, because of identical polyfiller:
//     2ac -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
//     24c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2ad skipped, because of identical polyfiller:
//     2ad -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
//     24d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2ae skipped, because of identical polyfiller:
//     2ae -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
//     24e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2af skipped, because of identical polyfiller:
//     2af -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
//     24f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2b0 skipped, because of identical polyfiller:
//     2b0 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
//     250 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2b1 skipped, because of identical polyfiller:
//     2b1 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
//     251 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

func (e3d *HwEngine3d) filler_2b2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2b3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2b4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2b5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2b6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2b7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_2b8 skipped, because of identical polyfiller:
//     2b8 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2b9 skipped, because of identical polyfiller:
//     2b9 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2ba skipped, because of identical polyfiller:
//     2ba -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

func (e3d *HwEngine3d) filler_2bb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2bc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2bd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_2be skipped, because of identical polyfiller:
//     2be -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
//     25e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2bf skipped, because of identical polyfiller:
//     2bf -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
//     25f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2c0 skipped, because of identical polyfiller:
//     2c0 -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
//     260 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2c1 skipped, because of identical polyfiller:
//     2c1 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
//     261 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2c2 skipped, because of identical polyfiller:
//     2c2 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
//     262 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2c3 skipped, because of identical polyfiller:
//     2c3 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
//     263 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2c4 skipped, because of identical polyfiller:
//     2c4 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
//     264 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2c5 skipped, because of identical polyfiller:
//     2c5 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
//     265 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2c6 skipped, because of identical polyfiller:
//     2c6 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
//     266 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2c7 skipped, because of identical polyfiller:
//     2c7 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
//     24f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2c8 skipped, because of identical polyfiller:
//     2c8 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
//     250 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2c9 skipped, because of identical polyfiller:
//     2c9 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
//     251 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

func (e3d *HwEngine3d) filler_2ca(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2cb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_2cc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		px = 0
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_2cd skipped, because of identical polyfiller:
//     2cd -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}
//     2b5 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}

// filler_2ce skipped, because of identical polyfiller:
//     2ce -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}
//     2b6 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}

// filler_2cf skipped, because of identical polyfiller:
//     2cf -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}
//     2b7 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}

// filler_2d0 skipped, because of identical polyfiller:
//     2d0 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2d1 skipped, because of identical polyfiller:
//     2d1 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2d2 skipped, because of identical polyfiller:
//     2d2 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2d3 skipped, because of identical polyfiller:
//     2d3 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     2a3 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}

// filler_2d4 skipped, because of identical polyfiller:
//     2d4 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     2a4 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}

// filler_2d5 skipped, because of identical polyfiller:
//     2d5 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     2a5 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}

// filler_2d6 skipped, because of identical polyfiller:
//     2d6 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     246 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2d7 skipped, because of identical polyfiller:
//     2d7 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     247 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2d8 skipped, because of identical polyfiller:
//     2d8 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     248 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2d9 skipped, because of identical polyfiller:
//     2d9 -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     249 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2da skipped, because of identical polyfiller:
//     2da -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     24a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2db skipped, because of identical polyfiller:
//     2db -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     24b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2dc skipped, because of identical polyfiller:
//     2dc -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     24c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2dd skipped, because of identical polyfiller:
//     2dd -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     24d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2de skipped, because of identical polyfiller:
//     2de -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     24e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2df skipped, because of identical polyfiller:
//     2df -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     24f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2e0 skipped, because of identical polyfiller:
//     2e0 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     250 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2e1 skipped, because of identical polyfiller:
//     2e1 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     251 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2e2 skipped, because of identical polyfiller:
//     2e2 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     2b2 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}

// filler_2e3 skipped, because of identical polyfiller:
//     2e3 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     2b3 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}

// filler_2e4 skipped, because of identical polyfiller:
//     2e4 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     2b4 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}

// filler_2e5 skipped, because of identical polyfiller:
//     2e5 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:0}
//     2b5 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}

// filler_2e6 skipped, because of identical polyfiller:
//     2e6 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:1}
//     2b6 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}

// filler_2e7 skipped, because of identical polyfiller:
//     2e7 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:3 TexCoords:2}
//     2b7 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}

// filler_2e8 skipped, because of identical polyfiller:
//     2e8 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2e9 skipped, because of identical polyfiller:
//     2e9 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2ea skipped, because of identical polyfiller:
//     2ea -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     240 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2eb skipped, because of identical polyfiller:
//     2eb -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     2bb -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}

// filler_2ec skipped, because of identical polyfiller:
//     2ec -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     2bc -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}

// filler_2ed skipped, because of identical polyfiller:
//     2ed -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     2bd -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}

// filler_2ee skipped, because of identical polyfiller:
//     2ee -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     25e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2ef skipped, because of identical polyfiller:
//     2ef -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     25f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2f0 skipped, because of identical polyfiller:
//     2f0 -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     260 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2f1 skipped, because of identical polyfiller:
//     2f1 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     261 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2f2 skipped, because of identical polyfiller:
//     2f2 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     262 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2f3 skipped, because of identical polyfiller:
//     2f3 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     263 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2f4 skipped, because of identical polyfiller:
//     2f4 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     264 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2f5 skipped, because of identical polyfiller:
//     2f5 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     265 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2f6 skipped, because of identical polyfiller:
//     2f6 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     266 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2f7 skipped, because of identical polyfiller:
//     2f7 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     24f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:0}

// filler_2f8 skipped, because of identical polyfiller:
//     2f8 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     250 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:1}

// filler_2f9 skipped, because of identical polyfiller:
//     2f9 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     251 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3 TexCoords:2}

// filler_2fa skipped, because of identical polyfiller:
//     2fa -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     2ca -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:0}

// filler_2fb skipped, because of identical polyfiller:
//     2fb -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     2cb -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:1}

// filler_2fc skipped, because of identical polyfiller:
//     2fc -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     2cc -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:3 TexCoords:2}

// filler_2fd skipped, because of identical polyfiller:
//     2fd -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:0}
//     2b5 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:0}

// filler_2fe skipped, because of identical polyfiller:
//     2fe -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:1}
//     2b6 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:1}

// filler_2ff skipped, because of identical polyfiller:
//     2ff -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:3 TexCoords:2}
//     2b7 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3 TexCoords:2}

func (e3d *HwEngine3d) filler_300(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_301 skipped, because of identical polyfiller:
//     301 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_302 skipped, because of identical polyfiller:
//     302 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

func (e3d *HwEngine3d) filler_303(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_304(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_305(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_306(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_307(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_308(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_309(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_30a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_30b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_30c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_30d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_30e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_30f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_310(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_311(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_312(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_313(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_314(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_315(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_316(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_317(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_318 skipped, because of identical polyfiller:
//     318 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_319 skipped, because of identical polyfiller:
//     319 -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_31a skipped, because of identical polyfiller:
//     31a -> {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

func (e3d *HwEngine3d) filler_31b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_31c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_31d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_31e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_31f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_320(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_321(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_322(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_323(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_324(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_325(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_326(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_327 skipped, because of identical polyfiller:
//     327 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}
//     30f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_328 skipped, because of identical polyfiller:
//     328 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}
//     310 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_329 skipped, because of identical polyfiller:
//     329 -> {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}
//     311 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

func (e3d *HwEngine3d) filler_32a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_32b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_32c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_32d skipped, because of identical polyfiller:
//     32d -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
//     315 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_32e skipped, because of identical polyfiller:
//     32e -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
//     316 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_32f skipped, because of identical polyfiller:
//     32f -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
//     317 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

func (e3d *HwEngine3d) filler_330(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}

// filler_331 skipped, because of identical polyfiller:
//     331 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
//     330 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}

// filler_332 skipped, because of identical polyfiller:
//     332 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
//     330 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}

// filler_333 skipped, because of identical polyfiller:
//     333 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
//     303 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_334 skipped, because of identical polyfiller:
//     334 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
//     304 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_335 skipped, because of identical polyfiller:
//     335 -> {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
//     305 -> {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

func (e3d *HwEngine3d) filler_336(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_337(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_338(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_339(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_33a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_33b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_33c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_33d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_33e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_33f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_340(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_341(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.texCache.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = decompTex.Get16(int(t<<tshift + s))
		// color key check
		if px == 0 {
			goto next
		}
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_342 skipped, because of identical polyfiller:
//     342 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
//     312 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_343 skipped, because of identical polyfiller:
//     343 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
//     313 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_344 skipped, because of identical polyfiller:
//     344 -> {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
//     314 -> {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_345 skipped, because of identical polyfiller:
//     345 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}
//     315 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_346 skipped, because of identical polyfiller:
//     346 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}
//     316 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_347 skipped, because of identical polyfiller:
//     347 -> {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}
//     317 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_348 skipped, because of identical polyfiller:
//     348 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
//     330 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}

// filler_349 skipped, because of identical polyfiller:
//     349 -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
//     330 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}

// filler_34a skipped, because of identical polyfiller:
//     34a -> {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
//     330 -> {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}

// filler_34b skipped, because of identical polyfiller:
//     34b -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
//     31b -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_34c skipped, because of identical polyfiller:
//     34c -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
//     31c -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_34d skipped, because of identical polyfiller:
//     34d -> {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
//     31d -> {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

func (e3d *HwEngine3d) filler_34e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_34f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_350(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 2
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_351(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_352(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_353(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	tshift -= 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		// color key check
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_354(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_355(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_356(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	polyalpha := uint8(poly.flags.Alpha()) << 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		// alpha blending with background
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := uint16(out.Get32(0))
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_357 skipped, because of identical polyfiller:
//     357 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
//     33f -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:0}

// filler_358 skipped, because of identical polyfiller:
//     358 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
//     340 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:1}

// filler_359 skipped, because of identical polyfiller:
//     359 -> {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
//     341 -> {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:4 TexCoords:2}

// filler_35a skipped, because of identical polyfiller:
//     35a -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
//     32a -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_35b skipped, because of identical polyfiller:
//     35b -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
//     32b -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_35c skipped, because of identical polyfiller:
//     35c -> {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
//     32c -> {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

// filler_35d skipped, because of identical polyfiller:
//     35d -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:0}
//     315 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_35e skipped, because of identical polyfiller:
//     35e -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:1}
//     316 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_35f skipped, because of identical polyfiller:
//     35f -> {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:4 TexCoords:2}
//     317 -> {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_360 skipped, because of identical polyfiller:
//     360 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_361 skipped, because of identical polyfiller:
//     361 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_362 skipped, because of identical polyfiller:
//     362 -> {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

func (e3d *HwEngine3d) filler_363(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_364(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_365(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_366 skipped, because of identical polyfiller:
//     366 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
//     306 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_367 skipped, because of identical polyfiller:
//     367 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
//     307 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_368 skipped, because of identical polyfiller:
//     368 -> {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
//     308 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_369 skipped, because of identical polyfiller:
//     369 -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
//     309 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_36a skipped, because of identical polyfiller:
//     36a -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
//     30a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_36b skipped, because of identical polyfiller:
//     36b -> {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
//     30b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_36c skipped, because of identical polyfiller:
//     36c -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
//     30c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_36d skipped, because of identical polyfiller:
//     36d -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
//     30d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_36e skipped, because of identical polyfiller:
//     36e -> {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
//     30e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_36f skipped, because of identical polyfiller:
//     36f -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
//     30f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_370 skipped, because of identical polyfiller:
//     370 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
//     310 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_371 skipped, because of identical polyfiller:
//     371 -> {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
//     311 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

func (e3d *HwEngine3d) filler_372(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_373(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_374(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_375(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_376(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_377(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	tshift += 1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px = e3d.texVram.Get16(texoff + t<<tshift + s*2)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_378 skipped, because of identical polyfiller:
//     378 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_379 skipped, because of identical polyfiller:
//     379 -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_37a skipped, because of identical polyfiller:
//     37a -> {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

func (e3d *HwEngine3d) filler_37b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_37c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_37d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_37e skipped, because of identical polyfiller:
//     37e -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
//     31e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_37f skipped, because of identical polyfiller:
//     37f -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
//     31f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_380 skipped, because of identical polyfiller:
//     380 -> {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
//     320 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

// filler_381 skipped, because of identical polyfiller:
//     381 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
//     321 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_382 skipped, because of identical polyfiller:
//     382 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
//     322 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_383 skipped, because of identical polyfiller:
//     383 -> {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
//     323 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

// filler_384 skipped, because of identical polyfiller:
//     384 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
//     324 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_385 skipped, because of identical polyfiller:
//     385 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
//     325 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_386 skipped, because of identical polyfiller:
//     386 -> {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
//     326 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

// filler_387 skipped, because of identical polyfiller:
//     387 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
//     30f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_388 skipped, because of identical polyfiller:
//     388 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
//     310 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_389 skipped, because of identical polyfiller:
//     389 -> {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
//     311 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

func (e3d *HwEngine3d) filler_38a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sclamp, tclamp := poly.tex.SClampMask, poly.tex.TClampMask
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		if s&sclamp != 0 {
			s = ^uint32(int32(s) >> 31)
		}
		if t&tclamp != 0 {
			t = ^uint32(int32(t) >> 31)
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_38b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	sflip, tflip := poly.tex.SFlipMask, poly.tex.TFlipMask
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		if s&sflip != 0 {
			s = ^s
		}
		if t&tflip != 0 {
			t = ^t
		}
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

func (e3d *HwEngine3d) filler_38c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()
	t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()
	ds, dt := s1.SubFixed(s0).Div(nx), t1.SubFixed(t0).Div(nx)
	smask, tmask := poly.tex.Width-1, poly.tex.Height-1
	var px uint16
	var pxa uint8
	pxa = 63
	var px0 uint8
	var s, t uint32
	out.Add32(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		// zbuffer check
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		// texel coords
		s, t = uint32(s0.TruncInt32()), uint32(t0.TruncInt32())
		s, t = s&smask, t&tmask
		// texel fetch
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		// color key check
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		// apply vertex color to texel
		if true {
			tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])
			tc := newColorFrom555U(tc0)
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(tc)
			pxc = pxc.AddSat(tc)
			px = pxc.To555U()
		}
		// alpha blending with background
		// draw color and z
		out.Set32(0, uint32(px)|0x80000000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add32(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

// filler_38d skipped, because of identical polyfiller:
//     38d -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}
//     375 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}

// filler_38e skipped, because of identical polyfiller:
//     38e -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}
//     376 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}

// filler_38f skipped, because of identical polyfiller:
//     38f -> {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}
//     377 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}

// filler_390 skipped, because of identical polyfiller:
//     390 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_391 skipped, because of identical polyfiller:
//     391 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_392 skipped, because of identical polyfiller:
//     392 -> {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_393 skipped, because of identical polyfiller:
//     393 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     363 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}

// filler_394 skipped, because of identical polyfiller:
//     394 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     364 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}

// filler_395 skipped, because of identical polyfiller:
//     395 -> {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     365 -> {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}

// filler_396 skipped, because of identical polyfiller:
//     396 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     306 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_397 skipped, because of identical polyfiller:
//     397 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     307 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_398 skipped, because of identical polyfiller:
//     398 -> {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     308 -> {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_399 skipped, because of identical polyfiller:
//     399 -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     309 -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_39a skipped, because of identical polyfiller:
//     39a -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     30a -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_39b skipped, because of identical polyfiller:
//     39b -> {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     30b -> {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_39c skipped, because of identical polyfiller:
//     39c -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     30c -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_39d skipped, because of identical polyfiller:
//     39d -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     30d -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_39e skipped, because of identical polyfiller:
//     39e -> {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     30e -> {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_39f skipped, because of identical polyfiller:
//     39f -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     30f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3a0 skipped, because of identical polyfiller:
//     3a0 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     310 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_3a1 skipped, because of identical polyfiller:
//     3a1 -> {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     311 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_3a2 skipped, because of identical polyfiller:
//     3a2 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     372 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}

// filler_3a3 skipped, because of identical polyfiller:
//     3a3 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     373 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}

// filler_3a4 skipped, because of identical polyfiller:
//     3a4 -> {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     374 -> {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}

// filler_3a5 skipped, because of identical polyfiller:
//     3a5 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:0}
//     375 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}

// filler_3a6 skipped, because of identical polyfiller:
//     3a6 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:1}
//     376 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}

// filler_3a7 skipped, because of identical polyfiller:
//     3a7 -> {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:4 TexCoords:2}
//     377 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}

// filler_3a8 skipped, because of identical polyfiller:
//     3a8 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3a9 skipped, because of identical polyfiller:
//     3a9 -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3aa skipped, because of identical polyfiller:
//     3aa -> {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     300 -> {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3ab skipped, because of identical polyfiller:
//     3ab -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     37b -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}

// filler_3ac skipped, because of identical polyfiller:
//     3ac -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     37c -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}

// filler_3ad skipped, because of identical polyfiller:
//     3ad -> {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     37d -> {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}

// filler_3ae skipped, because of identical polyfiller:
//     3ae -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     31e -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3af skipped, because of identical polyfiller:
//     3af -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     31f -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_3b0 skipped, because of identical polyfiller:
//     3b0 -> {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     320 -> {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

// filler_3b1 skipped, because of identical polyfiller:
//     3b1 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     321 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3b2 skipped, because of identical polyfiller:
//     3b2 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     322 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_3b3 skipped, because of identical polyfiller:
//     3b3 -> {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     323 -> {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

// filler_3b4 skipped, because of identical polyfiller:
//     3b4 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     324 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3b5 skipped, because of identical polyfiller:
//     3b5 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     325 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:1}

// filler_3b6 skipped, because of identical polyfiller:
//     3b6 -> {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     326 -> {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4 TexCoords:2}

// filler_3b7 skipped, because of identical polyfiller:
//     3b7 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     30f -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:0}

// filler_3b8 skipped, because of identical polyfiller:
//     3b8 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     310 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:1}

// filler_3b9 skipped, because of identical polyfiller:
//     3b9 -> {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     311 -> {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4 TexCoords:2}

// filler_3ba skipped, because of identical polyfiller:
//     3ba -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     38a -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:0}

// filler_3bb skipped, because of identical polyfiller:
//     3bb -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     38b -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:1}

// filler_3bc skipped, because of identical polyfiller:
//     3bc -> {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     38c -> {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:4 TexCoords:2}

// filler_3bd skipped, because of identical polyfiller:
//     3bd -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:0}
//     375 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:0}

// filler_3be skipped, because of identical polyfiller:
//     3be -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:1}
//     376 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:1}

// filler_3bf skipped, because of identical polyfiller:
//     3bf -> {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:4 TexCoords:2}
//     377 -> {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4 TexCoords:2}

var polygonFillerTable = [960]func(*HwEngine3d, *Polygon, gfx.Line, gfx.Line, gfx.Line){
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_003,
	(*HwEngine3d).filler_004,
	(*HwEngine3d).filler_005,
	(*HwEngine3d).filler_006,
	(*HwEngine3d).filler_007,
	(*HwEngine3d).filler_008,
	(*HwEngine3d).filler_009,
	(*HwEngine3d).filler_00a,
	(*HwEngine3d).filler_00b,
	(*HwEngine3d).filler_00c,
	(*HwEngine3d).filler_00d,
	(*HwEngine3d).filler_00e,
	(*HwEngine3d).filler_00f,
	(*HwEngine3d).filler_010,
	(*HwEngine3d).filler_011,
	(*HwEngine3d).filler_012,
	(*HwEngine3d).filler_013,
	(*HwEngine3d).filler_014,
	(*HwEngine3d).filler_015,
	(*HwEngine3d).filler_016,
	(*HwEngine3d).filler_017,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_01b,
	(*HwEngine3d).filler_01c,
	(*HwEngine3d).filler_01d,
	(*HwEngine3d).filler_01e,
	(*HwEngine3d).filler_01f,
	(*HwEngine3d).filler_020,
	(*HwEngine3d).filler_021,
	(*HwEngine3d).filler_022,
	(*HwEngine3d).filler_023,
	(*HwEngine3d).filler_024,
	(*HwEngine3d).filler_025,
	(*HwEngine3d).filler_026,
	(*HwEngine3d).filler_00f,
	(*HwEngine3d).filler_010,
	(*HwEngine3d).filler_011,
	(*HwEngine3d).filler_02a,
	(*HwEngine3d).filler_02b,
	(*HwEngine3d).filler_02c,
	(*HwEngine3d).filler_015,
	(*HwEngine3d).filler_016,
	(*HwEngine3d).filler_017,
	(*HwEngine3d).filler_030,
	(*HwEngine3d).filler_030,
	(*HwEngine3d).filler_030,
	(*HwEngine3d).filler_003,
	(*HwEngine3d).filler_004,
	(*HwEngine3d).filler_005,
	(*HwEngine3d).filler_036,
	(*HwEngine3d).filler_037,
	(*HwEngine3d).filler_038,
	(*HwEngine3d).filler_039,
	(*HwEngine3d).filler_03a,
	(*HwEngine3d).filler_03b,
	(*HwEngine3d).filler_03c,
	(*HwEngine3d).filler_03d,
	(*HwEngine3d).filler_03e,
	(*HwEngine3d).filler_03f,
	(*HwEngine3d).filler_040,
	(*HwEngine3d).filler_041,
	(*HwEngine3d).filler_012,
	(*HwEngine3d).filler_013,
	(*HwEngine3d).filler_014,
	(*HwEngine3d).filler_015,
	(*HwEngine3d).filler_016,
	(*HwEngine3d).filler_017,
	(*HwEngine3d).filler_030,
	(*HwEngine3d).filler_030,
	(*HwEngine3d).filler_030,
	(*HwEngine3d).filler_01b,
	(*HwEngine3d).filler_01c,
	(*HwEngine3d).filler_01d,
	(*HwEngine3d).filler_04e,
	(*HwEngine3d).filler_04f,
	(*HwEngine3d).filler_050,
	(*HwEngine3d).filler_051,
	(*HwEngine3d).filler_052,
	(*HwEngine3d).filler_053,
	(*HwEngine3d).filler_054,
	(*HwEngine3d).filler_055,
	(*HwEngine3d).filler_056,
	(*HwEngine3d).filler_03f,
	(*HwEngine3d).filler_040,
	(*HwEngine3d).filler_041,
	(*HwEngine3d).filler_02a,
	(*HwEngine3d).filler_02b,
	(*HwEngine3d).filler_02c,
	(*HwEngine3d).filler_015,
	(*HwEngine3d).filler_016,
	(*HwEngine3d).filler_017,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_063,
	(*HwEngine3d).filler_064,
	(*HwEngine3d).filler_065,
	(*HwEngine3d).filler_006,
	(*HwEngine3d).filler_007,
	(*HwEngine3d).filler_008,
	(*HwEngine3d).filler_009,
	(*HwEngine3d).filler_00a,
	(*HwEngine3d).filler_00b,
	(*HwEngine3d).filler_00c,
	(*HwEngine3d).filler_00d,
	(*HwEngine3d).filler_00e,
	(*HwEngine3d).filler_00f,
	(*HwEngine3d).filler_010,
	(*HwEngine3d).filler_011,
	(*HwEngine3d).filler_072,
	(*HwEngine3d).filler_073,
	(*HwEngine3d).filler_074,
	(*HwEngine3d).filler_075,
	(*HwEngine3d).filler_076,
	(*HwEngine3d).filler_077,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_07b,
	(*HwEngine3d).filler_07c,
	(*HwEngine3d).filler_07d,
	(*HwEngine3d).filler_01e,
	(*HwEngine3d).filler_01f,
	(*HwEngine3d).filler_020,
	(*HwEngine3d).filler_021,
	(*HwEngine3d).filler_022,
	(*HwEngine3d).filler_023,
	(*HwEngine3d).filler_024,
	(*HwEngine3d).filler_025,
	(*HwEngine3d).filler_026,
	(*HwEngine3d).filler_00f,
	(*HwEngine3d).filler_010,
	(*HwEngine3d).filler_011,
	(*HwEngine3d).filler_08a,
	(*HwEngine3d).filler_08b,
	(*HwEngine3d).filler_08c,
	(*HwEngine3d).filler_075,
	(*HwEngine3d).filler_076,
	(*HwEngine3d).filler_077,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_063,
	(*HwEngine3d).filler_064,
	(*HwEngine3d).filler_065,
	(*HwEngine3d).filler_006,
	(*HwEngine3d).filler_007,
	(*HwEngine3d).filler_008,
	(*HwEngine3d).filler_009,
	(*HwEngine3d).filler_00a,
	(*HwEngine3d).filler_00b,
	(*HwEngine3d).filler_00c,
	(*HwEngine3d).filler_00d,
	(*HwEngine3d).filler_00e,
	(*HwEngine3d).filler_00f,
	(*HwEngine3d).filler_010,
	(*HwEngine3d).filler_011,
	(*HwEngine3d).filler_072,
	(*HwEngine3d).filler_073,
	(*HwEngine3d).filler_074,
	(*HwEngine3d).filler_075,
	(*HwEngine3d).filler_076,
	(*HwEngine3d).filler_077,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_07b,
	(*HwEngine3d).filler_07c,
	(*HwEngine3d).filler_07d,
	(*HwEngine3d).filler_01e,
	(*HwEngine3d).filler_01f,
	(*HwEngine3d).filler_020,
	(*HwEngine3d).filler_021,
	(*HwEngine3d).filler_022,
	(*HwEngine3d).filler_023,
	(*HwEngine3d).filler_024,
	(*HwEngine3d).filler_025,
	(*HwEngine3d).filler_026,
	(*HwEngine3d).filler_00f,
	(*HwEngine3d).filler_010,
	(*HwEngine3d).filler_011,
	(*HwEngine3d).filler_08a,
	(*HwEngine3d).filler_08b,
	(*HwEngine3d).filler_08c,
	(*HwEngine3d).filler_075,
	(*HwEngine3d).filler_076,
	(*HwEngine3d).filler_077,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c3,
	(*HwEngine3d).filler_0c4,
	(*HwEngine3d).filler_0c5,
	(*HwEngine3d).filler_0c6,
	(*HwEngine3d).filler_0c7,
	(*HwEngine3d).filler_0c8,
	(*HwEngine3d).filler_0c9,
	(*HwEngine3d).filler_0ca,
	(*HwEngine3d).filler_0cb,
	(*HwEngine3d).filler_0cc,
	(*HwEngine3d).filler_0cd,
	(*HwEngine3d).filler_0ce,
	(*HwEngine3d).filler_0cf,
	(*HwEngine3d).filler_0d0,
	(*HwEngine3d).filler_0d1,
	(*HwEngine3d).filler_0d2,
	(*HwEngine3d).filler_0d3,
	(*HwEngine3d).filler_0d4,
	(*HwEngine3d).filler_0d5,
	(*HwEngine3d).filler_0d6,
	(*HwEngine3d).filler_0d7,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0db,
	(*HwEngine3d).filler_0dc,
	(*HwEngine3d).filler_0dd,
	(*HwEngine3d).filler_0de,
	(*HwEngine3d).filler_0df,
	(*HwEngine3d).filler_0e0,
	(*HwEngine3d).filler_0e1,
	(*HwEngine3d).filler_0e2,
	(*HwEngine3d).filler_0e3,
	(*HwEngine3d).filler_0e4,
	(*HwEngine3d).filler_0e5,
	(*HwEngine3d).filler_0e6,
	(*HwEngine3d).filler_0cf,
	(*HwEngine3d).filler_0d0,
	(*HwEngine3d).filler_0d1,
	(*HwEngine3d).filler_0ea,
	(*HwEngine3d).filler_0eb,
	(*HwEngine3d).filler_0ec,
	(*HwEngine3d).filler_0d5,
	(*HwEngine3d).filler_0d6,
	(*HwEngine3d).filler_0d7,
	(*HwEngine3d).filler_0f0,
	(*HwEngine3d).filler_0f0,
	(*HwEngine3d).filler_0f0,
	(*HwEngine3d).filler_0c3,
	(*HwEngine3d).filler_0c4,
	(*HwEngine3d).filler_0c5,
	(*HwEngine3d).filler_0f6,
	(*HwEngine3d).filler_0f7,
	(*HwEngine3d).filler_0f8,
	(*HwEngine3d).filler_0f9,
	(*HwEngine3d).filler_0fa,
	(*HwEngine3d).filler_0fb,
	(*HwEngine3d).filler_0fc,
	(*HwEngine3d).filler_0fd,
	(*HwEngine3d).filler_0fe,
	(*HwEngine3d).filler_0ff,
	(*HwEngine3d).filler_100,
	(*HwEngine3d).filler_101,
	(*HwEngine3d).filler_0d2,
	(*HwEngine3d).filler_0d3,
	(*HwEngine3d).filler_0d4,
	(*HwEngine3d).filler_0d5,
	(*HwEngine3d).filler_0d6,
	(*HwEngine3d).filler_0d7,
	(*HwEngine3d).filler_0f0,
	(*HwEngine3d).filler_0f0,
	(*HwEngine3d).filler_0f0,
	(*HwEngine3d).filler_0db,
	(*HwEngine3d).filler_0dc,
	(*HwEngine3d).filler_0dd,
	(*HwEngine3d).filler_10e,
	(*HwEngine3d).filler_10f,
	(*HwEngine3d).filler_110,
	(*HwEngine3d).filler_111,
	(*HwEngine3d).filler_112,
	(*HwEngine3d).filler_113,
	(*HwEngine3d).filler_114,
	(*HwEngine3d).filler_115,
	(*HwEngine3d).filler_116,
	(*HwEngine3d).filler_0ff,
	(*HwEngine3d).filler_100,
	(*HwEngine3d).filler_101,
	(*HwEngine3d).filler_0ea,
	(*HwEngine3d).filler_0eb,
	(*HwEngine3d).filler_0ec,
	(*HwEngine3d).filler_0d5,
	(*HwEngine3d).filler_0d6,
	(*HwEngine3d).filler_0d7,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_123,
	(*HwEngine3d).filler_124,
	(*HwEngine3d).filler_125,
	(*HwEngine3d).filler_0c6,
	(*HwEngine3d).filler_0c7,
	(*HwEngine3d).filler_0c8,
	(*HwEngine3d).filler_0c9,
	(*HwEngine3d).filler_0ca,
	(*HwEngine3d).filler_0cb,
	(*HwEngine3d).filler_0cc,
	(*HwEngine3d).filler_0cd,
	(*HwEngine3d).filler_0ce,
	(*HwEngine3d).filler_0cf,
	(*HwEngine3d).filler_0d0,
	(*HwEngine3d).filler_0d1,
	(*HwEngine3d).filler_132,
	(*HwEngine3d).filler_133,
	(*HwEngine3d).filler_134,
	(*HwEngine3d).filler_135,
	(*HwEngine3d).filler_136,
	(*HwEngine3d).filler_137,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_13b,
	(*HwEngine3d).filler_13c,
	(*HwEngine3d).filler_13d,
	(*HwEngine3d).filler_0de,
	(*HwEngine3d).filler_0df,
	(*HwEngine3d).filler_0e0,
	(*HwEngine3d).filler_0e1,
	(*HwEngine3d).filler_0e2,
	(*HwEngine3d).filler_0e3,
	(*HwEngine3d).filler_0e4,
	(*HwEngine3d).filler_0e5,
	(*HwEngine3d).filler_0e6,
	(*HwEngine3d).filler_0cf,
	(*HwEngine3d).filler_0d0,
	(*HwEngine3d).filler_0d1,
	(*HwEngine3d).filler_14a,
	(*HwEngine3d).filler_14b,
	(*HwEngine3d).filler_14c,
	(*HwEngine3d).filler_135,
	(*HwEngine3d).filler_136,
	(*HwEngine3d).filler_137,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_123,
	(*HwEngine3d).filler_124,
	(*HwEngine3d).filler_125,
	(*HwEngine3d).filler_0c6,
	(*HwEngine3d).filler_0c7,
	(*HwEngine3d).filler_0c8,
	(*HwEngine3d).filler_0c9,
	(*HwEngine3d).filler_0ca,
	(*HwEngine3d).filler_0cb,
	(*HwEngine3d).filler_0cc,
	(*HwEngine3d).filler_0cd,
	(*HwEngine3d).filler_0ce,
	(*HwEngine3d).filler_0cf,
	(*HwEngine3d).filler_0d0,
	(*HwEngine3d).filler_0d1,
	(*HwEngine3d).filler_132,
	(*HwEngine3d).filler_133,
	(*HwEngine3d).filler_134,
	(*HwEngine3d).filler_135,
	(*HwEngine3d).filler_136,
	(*HwEngine3d).filler_137,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_13b,
	(*HwEngine3d).filler_13c,
	(*HwEngine3d).filler_13d,
	(*HwEngine3d).filler_0de,
	(*HwEngine3d).filler_0df,
	(*HwEngine3d).filler_0e0,
	(*HwEngine3d).filler_0e1,
	(*HwEngine3d).filler_0e2,
	(*HwEngine3d).filler_0e3,
	(*HwEngine3d).filler_0e4,
	(*HwEngine3d).filler_0e5,
	(*HwEngine3d).filler_0e6,
	(*HwEngine3d).filler_0cf,
	(*HwEngine3d).filler_0d0,
	(*HwEngine3d).filler_0d1,
	(*HwEngine3d).filler_14a,
	(*HwEngine3d).filler_14b,
	(*HwEngine3d).filler_14c,
	(*HwEngine3d).filler_135,
	(*HwEngine3d).filler_136,
	(*HwEngine3d).filler_137,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_183,
	(*HwEngine3d).filler_184,
	(*HwEngine3d).filler_185,
	(*HwEngine3d).filler_186,
	(*HwEngine3d).filler_187,
	(*HwEngine3d).filler_188,
	(*HwEngine3d).filler_189,
	(*HwEngine3d).filler_18a,
	(*HwEngine3d).filler_18b,
	(*HwEngine3d).filler_18c,
	(*HwEngine3d).filler_18d,
	(*HwEngine3d).filler_18e,
	(*HwEngine3d).filler_18f,
	(*HwEngine3d).filler_190,
	(*HwEngine3d).filler_191,
	(*HwEngine3d).filler_192,
	(*HwEngine3d).filler_193,
	(*HwEngine3d).filler_194,
	(*HwEngine3d).filler_195,
	(*HwEngine3d).filler_196,
	(*HwEngine3d).filler_197,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_19b,
	(*HwEngine3d).filler_19c,
	(*HwEngine3d).filler_19d,
	(*HwEngine3d).filler_19e,
	(*HwEngine3d).filler_19f,
	(*HwEngine3d).filler_1a0,
	(*HwEngine3d).filler_1a1,
	(*HwEngine3d).filler_1a2,
	(*HwEngine3d).filler_1a3,
	(*HwEngine3d).filler_1a4,
	(*HwEngine3d).filler_1a5,
	(*HwEngine3d).filler_1a6,
	(*HwEngine3d).filler_18f,
	(*HwEngine3d).filler_190,
	(*HwEngine3d).filler_191,
	(*HwEngine3d).filler_1aa,
	(*HwEngine3d).filler_1ab,
	(*HwEngine3d).filler_1ac,
	(*HwEngine3d).filler_195,
	(*HwEngine3d).filler_196,
	(*HwEngine3d).filler_197,
	(*HwEngine3d).filler_1b0,
	(*HwEngine3d).filler_1b0,
	(*HwEngine3d).filler_1b0,
	(*HwEngine3d).filler_183,
	(*HwEngine3d).filler_184,
	(*HwEngine3d).filler_185,
	(*HwEngine3d).filler_1b6,
	(*HwEngine3d).filler_1b7,
	(*HwEngine3d).filler_1b8,
	(*HwEngine3d).filler_1b9,
	(*HwEngine3d).filler_1ba,
	(*HwEngine3d).filler_1bb,
	(*HwEngine3d).filler_1bc,
	(*HwEngine3d).filler_1bd,
	(*HwEngine3d).filler_1be,
	(*HwEngine3d).filler_1bf,
	(*HwEngine3d).filler_1c0,
	(*HwEngine3d).filler_1c1,
	(*HwEngine3d).filler_192,
	(*HwEngine3d).filler_193,
	(*HwEngine3d).filler_194,
	(*HwEngine3d).filler_195,
	(*HwEngine3d).filler_196,
	(*HwEngine3d).filler_197,
	(*HwEngine3d).filler_1b0,
	(*HwEngine3d).filler_1b0,
	(*HwEngine3d).filler_1b0,
	(*HwEngine3d).filler_19b,
	(*HwEngine3d).filler_19c,
	(*HwEngine3d).filler_19d,
	(*HwEngine3d).filler_1ce,
	(*HwEngine3d).filler_1cf,
	(*HwEngine3d).filler_1d0,
	(*HwEngine3d).filler_1d1,
	(*HwEngine3d).filler_1d2,
	(*HwEngine3d).filler_1d3,
	(*HwEngine3d).filler_1d4,
	(*HwEngine3d).filler_1d5,
	(*HwEngine3d).filler_1d6,
	(*HwEngine3d).filler_1bf,
	(*HwEngine3d).filler_1c0,
	(*HwEngine3d).filler_1c1,
	(*HwEngine3d).filler_1aa,
	(*HwEngine3d).filler_1ab,
	(*HwEngine3d).filler_1ac,
	(*HwEngine3d).filler_195,
	(*HwEngine3d).filler_196,
	(*HwEngine3d).filler_197,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_1e3,
	(*HwEngine3d).filler_1e4,
	(*HwEngine3d).filler_1e5,
	(*HwEngine3d).filler_186,
	(*HwEngine3d).filler_187,
	(*HwEngine3d).filler_188,
	(*HwEngine3d).filler_189,
	(*HwEngine3d).filler_18a,
	(*HwEngine3d).filler_18b,
	(*HwEngine3d).filler_18c,
	(*HwEngine3d).filler_18d,
	(*HwEngine3d).filler_18e,
	(*HwEngine3d).filler_18f,
	(*HwEngine3d).filler_190,
	(*HwEngine3d).filler_191,
	(*HwEngine3d).filler_1f2,
	(*HwEngine3d).filler_1f3,
	(*HwEngine3d).filler_1f4,
	(*HwEngine3d).filler_1f5,
	(*HwEngine3d).filler_1f6,
	(*HwEngine3d).filler_1f7,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_1fb,
	(*HwEngine3d).filler_1fc,
	(*HwEngine3d).filler_1fd,
	(*HwEngine3d).filler_19e,
	(*HwEngine3d).filler_19f,
	(*HwEngine3d).filler_1a0,
	(*HwEngine3d).filler_1a1,
	(*HwEngine3d).filler_1a2,
	(*HwEngine3d).filler_1a3,
	(*HwEngine3d).filler_1a4,
	(*HwEngine3d).filler_1a5,
	(*HwEngine3d).filler_1a6,
	(*HwEngine3d).filler_18f,
	(*HwEngine3d).filler_190,
	(*HwEngine3d).filler_191,
	(*HwEngine3d).filler_20a,
	(*HwEngine3d).filler_20b,
	(*HwEngine3d).filler_20c,
	(*HwEngine3d).filler_1f5,
	(*HwEngine3d).filler_1f6,
	(*HwEngine3d).filler_1f7,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_1e3,
	(*HwEngine3d).filler_1e4,
	(*HwEngine3d).filler_1e5,
	(*HwEngine3d).filler_186,
	(*HwEngine3d).filler_187,
	(*HwEngine3d).filler_188,
	(*HwEngine3d).filler_189,
	(*HwEngine3d).filler_18a,
	(*HwEngine3d).filler_18b,
	(*HwEngine3d).filler_18c,
	(*HwEngine3d).filler_18d,
	(*HwEngine3d).filler_18e,
	(*HwEngine3d).filler_18f,
	(*HwEngine3d).filler_190,
	(*HwEngine3d).filler_191,
	(*HwEngine3d).filler_1f2,
	(*HwEngine3d).filler_1f3,
	(*HwEngine3d).filler_1f4,
	(*HwEngine3d).filler_1f5,
	(*HwEngine3d).filler_1f6,
	(*HwEngine3d).filler_1f7,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_180,
	(*HwEngine3d).filler_1fb,
	(*HwEngine3d).filler_1fc,
	(*HwEngine3d).filler_1fd,
	(*HwEngine3d).filler_19e,
	(*HwEngine3d).filler_19f,
	(*HwEngine3d).filler_1a0,
	(*HwEngine3d).filler_1a1,
	(*HwEngine3d).filler_1a2,
	(*HwEngine3d).filler_1a3,
	(*HwEngine3d).filler_1a4,
	(*HwEngine3d).filler_1a5,
	(*HwEngine3d).filler_1a6,
	(*HwEngine3d).filler_18f,
	(*HwEngine3d).filler_190,
	(*HwEngine3d).filler_191,
	(*HwEngine3d).filler_20a,
	(*HwEngine3d).filler_20b,
	(*HwEngine3d).filler_20c,
	(*HwEngine3d).filler_1f5,
	(*HwEngine3d).filler_1f6,
	(*HwEngine3d).filler_1f7,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_243,
	(*HwEngine3d).filler_244,
	(*HwEngine3d).filler_245,
	(*HwEngine3d).filler_246,
	(*HwEngine3d).filler_247,
	(*HwEngine3d).filler_248,
	(*HwEngine3d).filler_249,
	(*HwEngine3d).filler_24a,
	(*HwEngine3d).filler_24b,
	(*HwEngine3d).filler_24c,
	(*HwEngine3d).filler_24d,
	(*HwEngine3d).filler_24e,
	(*HwEngine3d).filler_24f,
	(*HwEngine3d).filler_250,
	(*HwEngine3d).filler_251,
	(*HwEngine3d).filler_252,
	(*HwEngine3d).filler_253,
	(*HwEngine3d).filler_254,
	(*HwEngine3d).filler_255,
	(*HwEngine3d).filler_256,
	(*HwEngine3d).filler_257,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_25b,
	(*HwEngine3d).filler_25c,
	(*HwEngine3d).filler_25d,
	(*HwEngine3d).filler_25e,
	(*HwEngine3d).filler_25f,
	(*HwEngine3d).filler_260,
	(*HwEngine3d).filler_261,
	(*HwEngine3d).filler_262,
	(*HwEngine3d).filler_263,
	(*HwEngine3d).filler_264,
	(*HwEngine3d).filler_265,
	(*HwEngine3d).filler_266,
	(*HwEngine3d).filler_24f,
	(*HwEngine3d).filler_250,
	(*HwEngine3d).filler_251,
	(*HwEngine3d).filler_26a,
	(*HwEngine3d).filler_26b,
	(*HwEngine3d).filler_26c,
	(*HwEngine3d).filler_255,
	(*HwEngine3d).filler_256,
	(*HwEngine3d).filler_257,
	(*HwEngine3d).filler_270,
	(*HwEngine3d).filler_270,
	(*HwEngine3d).filler_270,
	(*HwEngine3d).filler_243,
	(*HwEngine3d).filler_244,
	(*HwEngine3d).filler_245,
	(*HwEngine3d).filler_276,
	(*HwEngine3d).filler_277,
	(*HwEngine3d).filler_278,
	(*HwEngine3d).filler_279,
	(*HwEngine3d).filler_27a,
	(*HwEngine3d).filler_27b,
	(*HwEngine3d).filler_27c,
	(*HwEngine3d).filler_27d,
	(*HwEngine3d).filler_27e,
	(*HwEngine3d).filler_27f,
	(*HwEngine3d).filler_280,
	(*HwEngine3d).filler_281,
	(*HwEngine3d).filler_252,
	(*HwEngine3d).filler_253,
	(*HwEngine3d).filler_254,
	(*HwEngine3d).filler_255,
	(*HwEngine3d).filler_256,
	(*HwEngine3d).filler_257,
	(*HwEngine3d).filler_270,
	(*HwEngine3d).filler_270,
	(*HwEngine3d).filler_270,
	(*HwEngine3d).filler_25b,
	(*HwEngine3d).filler_25c,
	(*HwEngine3d).filler_25d,
	(*HwEngine3d).filler_28e,
	(*HwEngine3d).filler_28f,
	(*HwEngine3d).filler_290,
	(*HwEngine3d).filler_291,
	(*HwEngine3d).filler_292,
	(*HwEngine3d).filler_293,
	(*HwEngine3d).filler_294,
	(*HwEngine3d).filler_295,
	(*HwEngine3d).filler_296,
	(*HwEngine3d).filler_27f,
	(*HwEngine3d).filler_280,
	(*HwEngine3d).filler_281,
	(*HwEngine3d).filler_26a,
	(*HwEngine3d).filler_26b,
	(*HwEngine3d).filler_26c,
	(*HwEngine3d).filler_255,
	(*HwEngine3d).filler_256,
	(*HwEngine3d).filler_257,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_2a3,
	(*HwEngine3d).filler_2a4,
	(*HwEngine3d).filler_2a5,
	(*HwEngine3d).filler_246,
	(*HwEngine3d).filler_247,
	(*HwEngine3d).filler_248,
	(*HwEngine3d).filler_249,
	(*HwEngine3d).filler_24a,
	(*HwEngine3d).filler_24b,
	(*HwEngine3d).filler_24c,
	(*HwEngine3d).filler_24d,
	(*HwEngine3d).filler_24e,
	(*HwEngine3d).filler_24f,
	(*HwEngine3d).filler_250,
	(*HwEngine3d).filler_251,
	(*HwEngine3d).filler_2b2,
	(*HwEngine3d).filler_2b3,
	(*HwEngine3d).filler_2b4,
	(*HwEngine3d).filler_2b5,
	(*HwEngine3d).filler_2b6,
	(*HwEngine3d).filler_2b7,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_2bb,
	(*HwEngine3d).filler_2bc,
	(*HwEngine3d).filler_2bd,
	(*HwEngine3d).filler_25e,
	(*HwEngine3d).filler_25f,
	(*HwEngine3d).filler_260,
	(*HwEngine3d).filler_261,
	(*HwEngine3d).filler_262,
	(*HwEngine3d).filler_263,
	(*HwEngine3d).filler_264,
	(*HwEngine3d).filler_265,
	(*HwEngine3d).filler_266,
	(*HwEngine3d).filler_24f,
	(*HwEngine3d).filler_250,
	(*HwEngine3d).filler_251,
	(*HwEngine3d).filler_2ca,
	(*HwEngine3d).filler_2cb,
	(*HwEngine3d).filler_2cc,
	(*HwEngine3d).filler_2b5,
	(*HwEngine3d).filler_2b6,
	(*HwEngine3d).filler_2b7,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_2a3,
	(*HwEngine3d).filler_2a4,
	(*HwEngine3d).filler_2a5,
	(*HwEngine3d).filler_246,
	(*HwEngine3d).filler_247,
	(*HwEngine3d).filler_248,
	(*HwEngine3d).filler_249,
	(*HwEngine3d).filler_24a,
	(*HwEngine3d).filler_24b,
	(*HwEngine3d).filler_24c,
	(*HwEngine3d).filler_24d,
	(*HwEngine3d).filler_24e,
	(*HwEngine3d).filler_24f,
	(*HwEngine3d).filler_250,
	(*HwEngine3d).filler_251,
	(*HwEngine3d).filler_2b2,
	(*HwEngine3d).filler_2b3,
	(*HwEngine3d).filler_2b4,
	(*HwEngine3d).filler_2b5,
	(*HwEngine3d).filler_2b6,
	(*HwEngine3d).filler_2b7,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_240,
	(*HwEngine3d).filler_2bb,
	(*HwEngine3d).filler_2bc,
	(*HwEngine3d).filler_2bd,
	(*HwEngine3d).filler_25e,
	(*HwEngine3d).filler_25f,
	(*HwEngine3d).filler_260,
	(*HwEngine3d).filler_261,
	(*HwEngine3d).filler_262,
	(*HwEngine3d).filler_263,
	(*HwEngine3d).filler_264,
	(*HwEngine3d).filler_265,
	(*HwEngine3d).filler_266,
	(*HwEngine3d).filler_24f,
	(*HwEngine3d).filler_250,
	(*HwEngine3d).filler_251,
	(*HwEngine3d).filler_2ca,
	(*HwEngine3d).filler_2cb,
	(*HwEngine3d).filler_2cc,
	(*HwEngine3d).filler_2b5,
	(*HwEngine3d).filler_2b6,
	(*HwEngine3d).filler_2b7,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_303,
	(*HwEngine3d).filler_304,
	(*HwEngine3d).filler_305,
	(*HwEngine3d).filler_306,
	(*HwEngine3d).filler_307,
	(*HwEngine3d).filler_308,
	(*HwEngine3d).filler_309,
	(*HwEngine3d).filler_30a,
	(*HwEngine3d).filler_30b,
	(*HwEngine3d).filler_30c,
	(*HwEngine3d).filler_30d,
	(*HwEngine3d).filler_30e,
	(*HwEngine3d).filler_30f,
	(*HwEngine3d).filler_310,
	(*HwEngine3d).filler_311,
	(*HwEngine3d).filler_312,
	(*HwEngine3d).filler_313,
	(*HwEngine3d).filler_314,
	(*HwEngine3d).filler_315,
	(*HwEngine3d).filler_316,
	(*HwEngine3d).filler_317,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_31b,
	(*HwEngine3d).filler_31c,
	(*HwEngine3d).filler_31d,
	(*HwEngine3d).filler_31e,
	(*HwEngine3d).filler_31f,
	(*HwEngine3d).filler_320,
	(*HwEngine3d).filler_321,
	(*HwEngine3d).filler_322,
	(*HwEngine3d).filler_323,
	(*HwEngine3d).filler_324,
	(*HwEngine3d).filler_325,
	(*HwEngine3d).filler_326,
	(*HwEngine3d).filler_30f,
	(*HwEngine3d).filler_310,
	(*HwEngine3d).filler_311,
	(*HwEngine3d).filler_32a,
	(*HwEngine3d).filler_32b,
	(*HwEngine3d).filler_32c,
	(*HwEngine3d).filler_315,
	(*HwEngine3d).filler_316,
	(*HwEngine3d).filler_317,
	(*HwEngine3d).filler_330,
	(*HwEngine3d).filler_330,
	(*HwEngine3d).filler_330,
	(*HwEngine3d).filler_303,
	(*HwEngine3d).filler_304,
	(*HwEngine3d).filler_305,
	(*HwEngine3d).filler_336,
	(*HwEngine3d).filler_337,
	(*HwEngine3d).filler_338,
	(*HwEngine3d).filler_339,
	(*HwEngine3d).filler_33a,
	(*HwEngine3d).filler_33b,
	(*HwEngine3d).filler_33c,
	(*HwEngine3d).filler_33d,
	(*HwEngine3d).filler_33e,
	(*HwEngine3d).filler_33f,
	(*HwEngine3d).filler_340,
	(*HwEngine3d).filler_341,
	(*HwEngine3d).filler_312,
	(*HwEngine3d).filler_313,
	(*HwEngine3d).filler_314,
	(*HwEngine3d).filler_315,
	(*HwEngine3d).filler_316,
	(*HwEngine3d).filler_317,
	(*HwEngine3d).filler_330,
	(*HwEngine3d).filler_330,
	(*HwEngine3d).filler_330,
	(*HwEngine3d).filler_31b,
	(*HwEngine3d).filler_31c,
	(*HwEngine3d).filler_31d,
	(*HwEngine3d).filler_34e,
	(*HwEngine3d).filler_34f,
	(*HwEngine3d).filler_350,
	(*HwEngine3d).filler_351,
	(*HwEngine3d).filler_352,
	(*HwEngine3d).filler_353,
	(*HwEngine3d).filler_354,
	(*HwEngine3d).filler_355,
	(*HwEngine3d).filler_356,
	(*HwEngine3d).filler_33f,
	(*HwEngine3d).filler_340,
	(*HwEngine3d).filler_341,
	(*HwEngine3d).filler_32a,
	(*HwEngine3d).filler_32b,
	(*HwEngine3d).filler_32c,
	(*HwEngine3d).filler_315,
	(*HwEngine3d).filler_316,
	(*HwEngine3d).filler_317,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_363,
	(*HwEngine3d).filler_364,
	(*HwEngine3d).filler_365,
	(*HwEngine3d).filler_306,
	(*HwEngine3d).filler_307,
	(*HwEngine3d).filler_308,
	(*HwEngine3d).filler_309,
	(*HwEngine3d).filler_30a,
	(*HwEngine3d).filler_30b,
	(*HwEngine3d).filler_30c,
	(*HwEngine3d).filler_30d,
	(*HwEngine3d).filler_30e,
	(*HwEngine3d).filler_30f,
	(*HwEngine3d).filler_310,
	(*HwEngine3d).filler_311,
	(*HwEngine3d).filler_372,
	(*HwEngine3d).filler_373,
	(*HwEngine3d).filler_374,
	(*HwEngine3d).filler_375,
	(*HwEngine3d).filler_376,
	(*HwEngine3d).filler_377,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_37b,
	(*HwEngine3d).filler_37c,
	(*HwEngine3d).filler_37d,
	(*HwEngine3d).filler_31e,
	(*HwEngine3d).filler_31f,
	(*HwEngine3d).filler_320,
	(*HwEngine3d).filler_321,
	(*HwEngine3d).filler_322,
	(*HwEngine3d).filler_323,
	(*HwEngine3d).filler_324,
	(*HwEngine3d).filler_325,
	(*HwEngine3d).filler_326,
	(*HwEngine3d).filler_30f,
	(*HwEngine3d).filler_310,
	(*HwEngine3d).filler_311,
	(*HwEngine3d).filler_38a,
	(*HwEngine3d).filler_38b,
	(*HwEngine3d).filler_38c,
	(*HwEngine3d).filler_375,
	(*HwEngine3d).filler_376,
	(*HwEngine3d).filler_377,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_363,
	(*HwEngine3d).filler_364,
	(*HwEngine3d).filler_365,
	(*HwEngine3d).filler_306,
	(*HwEngine3d).filler_307,
	(*HwEngine3d).filler_308,
	(*HwEngine3d).filler_309,
	(*HwEngine3d).filler_30a,
	(*HwEngine3d).filler_30b,
	(*HwEngine3d).filler_30c,
	(*HwEngine3d).filler_30d,
	(*HwEngine3d).filler_30e,
	(*HwEngine3d).filler_30f,
	(*HwEngine3d).filler_310,
	(*HwEngine3d).filler_311,
	(*HwEngine3d).filler_372,
	(*HwEngine3d).filler_373,
	(*HwEngine3d).filler_374,
	(*HwEngine3d).filler_375,
	(*HwEngine3d).filler_376,
	(*HwEngine3d).filler_377,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_300,
	(*HwEngine3d).filler_37b,
	(*HwEngine3d).filler_37c,
	(*HwEngine3d).filler_37d,
	(*HwEngine3d).filler_31e,
	(*HwEngine3d).filler_31f,
	(*HwEngine3d).filler_320,
	(*HwEngine3d).filler_321,
	(*HwEngine3d).filler_322,
	(*HwEngine3d).filler_323,
	(*HwEngine3d).filler_324,
	(*HwEngine3d).filler_325,
	(*HwEngine3d).filler_326,
	(*HwEngine3d).filler_30f,
	(*HwEngine3d).filler_310,
	(*HwEngine3d).filler_311,
	(*HwEngine3d).filler_38a,
	(*HwEngine3d).filler_38b,
	(*HwEngine3d).filler_38c,
	(*HwEngine3d).filler_375,
	(*HwEngine3d).filler_376,
	(*HwEngine3d).filler_377,
}
