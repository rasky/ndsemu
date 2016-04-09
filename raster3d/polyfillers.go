// Generated on 2016-04-09 11:40:13.041583741 +0200 CEST
package raster3d

import "ndsemu/emu/gfx"
import "ndsemu/emu"

func (e3d *HwEngine3d) filler_000(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_001(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_002(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_003(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_004(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_005(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_006(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_007(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_008(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_009(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_00a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_00b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_00c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_00d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_00e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_00f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_010(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_011(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_012(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_013(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_014(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_015(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_017(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_018(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_019(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_01a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_01b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_01c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_01d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_01e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_01f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_020(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_021(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_022(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_023(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_024(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_025(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_027(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_028(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_029(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_02a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_02b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_02c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_02d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_02e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_02f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_030(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_031(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_032(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_033(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_034(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_035(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_036(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_037(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_038(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:0}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_039(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_03a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_03b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_03c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_03d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_03e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_03f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_040(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_041(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_042(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_043(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_044(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_045(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_046(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_047(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_048(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_049(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_04a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_04b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_04c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_04d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_04e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_04f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_050(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_051(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_052(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_053(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_054(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_055(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_056(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_057(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_058(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_059(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_05a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_05b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_05c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_05d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_05e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_05f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_060(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_061(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_062(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_063(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_064(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_065(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_066(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_067(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_068(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_069(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_06a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_06b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_06c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_06d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_06e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_06f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_070(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_071(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_072(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_073(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_074(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_075(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_076(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_077(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_078(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:1}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_079(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_07a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_07b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_07c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_07d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_07e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_07f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_080(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_081(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_082(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_083(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_084(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_085(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_086(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_087(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_088(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_089(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_08a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_08b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_08c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_08d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_08e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_08f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_090(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_091(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_092(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_093(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_094(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_095(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_096(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_097(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_098(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_099(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_09a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_09b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_09c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_09d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_09e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_09f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0a0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0a1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0a2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0a3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0a4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0a5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_0a6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0a7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0a8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0a9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0aa(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ab(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ac(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ad(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ae(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0af(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0b0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0b1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0b2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0b3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0b4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0b5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_0b6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0b7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0b8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:2}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0b9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ba(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0bb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0bc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0bd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0be(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0bf(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0c0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0c1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0c2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0c3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0c4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0c5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_0c6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0c7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0c8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0c9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ca(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0cb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0cc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0cd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ce(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0cf(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0d0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0d1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0d2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0d3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0d4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0d5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_0d6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0d7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0d8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0d9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0da(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0db(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0dc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0dd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0de(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0df(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0e0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0e1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0e2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0e3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0e4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0e5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_0e6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0e7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0e8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0e9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ea(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0eb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ec(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ed(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ee(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ef(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0f0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0f1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0f2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0f3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0f4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0f5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_0f6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0f7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_0f8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:3}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_0f9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0fa(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0fb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0fc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0fd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0fe(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_0ff(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_100(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:0 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_101(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_102(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_103(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_104(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_105(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_106(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_107(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_108(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:0 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_109(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_10a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_10b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_10c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_10d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_10e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_10f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:0 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_110(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:1 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_111(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_112(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_113(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_114(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_115(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_116(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_117(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_118(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:1 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_119(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_11a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_11b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_11c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_11d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	polyalpha := uint8(poly.flags.Alpha()) << 1
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_11e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_11f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:1 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_120(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:2 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_121(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_122(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_123(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_124(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_125(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_126(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_127(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_128(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:2 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_129(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_12a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_12b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_12c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_12d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_12e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_12f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:2 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_130(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:0 FillMode:3 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_131(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:0 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_132(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:0 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_133(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:0 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_134(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:0 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_135(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:0 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
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
func (e3d *HwEngine3d) filler_136(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:0 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_137(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:0 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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
func (e3d *HwEngine3d) filler_138(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:1 FillMode:3 ColorMode:4}
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
		// texel fetch
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
func (e3d *HwEngine3d) filler_139(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:1 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_13a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:1 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_13b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:1 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_13c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:1 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_13d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:1 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	decompTexBuf := e3d.decompTex.Get(texoff)
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_13e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:1 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
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
func (e3d *HwEngine3d) filler_13f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:1 FillMode:3 ColorMode:4}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())
	dc := c1.SubColor(c0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
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
		// texel fetch
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
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

var polygonFillerTable = [320]func(*HwEngine3d, *Polygon, gfx.Line, gfx.Line, gfx.Line){
	(*HwEngine3d).filler_000,
	(*HwEngine3d).filler_001,
	(*HwEngine3d).filler_002,
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
	(*HwEngine3d).filler_018,
	(*HwEngine3d).filler_019,
	(*HwEngine3d).filler_01a,
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
	(*HwEngine3d).filler_027,
	(*HwEngine3d).filler_028,
	(*HwEngine3d).filler_029,
	(*HwEngine3d).filler_02a,
	(*HwEngine3d).filler_02b,
	(*HwEngine3d).filler_02c,
	(*HwEngine3d).filler_02d,
	(*HwEngine3d).filler_02e,
	(*HwEngine3d).filler_02f,
	(*HwEngine3d).filler_030,
	(*HwEngine3d).filler_031,
	(*HwEngine3d).filler_032,
	(*HwEngine3d).filler_033,
	(*HwEngine3d).filler_034,
	(*HwEngine3d).filler_035,
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
	(*HwEngine3d).filler_042,
	(*HwEngine3d).filler_043,
	(*HwEngine3d).filler_044,
	(*HwEngine3d).filler_045,
	(*HwEngine3d).filler_046,
	(*HwEngine3d).filler_047,
	(*HwEngine3d).filler_048,
	(*HwEngine3d).filler_049,
	(*HwEngine3d).filler_04a,
	(*HwEngine3d).filler_04b,
	(*HwEngine3d).filler_04c,
	(*HwEngine3d).filler_04d,
	(*HwEngine3d).filler_04e,
	(*HwEngine3d).filler_04f,
	(*HwEngine3d).filler_050,
	(*HwEngine3d).filler_051,
	(*HwEngine3d).filler_052,
	(*HwEngine3d).filler_053,
	(*HwEngine3d).filler_054,
	(*HwEngine3d).filler_055,
	(*HwEngine3d).filler_056,
	(*HwEngine3d).filler_057,
	(*HwEngine3d).filler_058,
	(*HwEngine3d).filler_059,
	(*HwEngine3d).filler_05a,
	(*HwEngine3d).filler_05b,
	(*HwEngine3d).filler_05c,
	(*HwEngine3d).filler_05d,
	(*HwEngine3d).filler_05e,
	(*HwEngine3d).filler_05f,
	(*HwEngine3d).filler_060,
	(*HwEngine3d).filler_061,
	(*HwEngine3d).filler_062,
	(*HwEngine3d).filler_063,
	(*HwEngine3d).filler_064,
	(*HwEngine3d).filler_065,
	(*HwEngine3d).filler_066,
	(*HwEngine3d).filler_067,
	(*HwEngine3d).filler_068,
	(*HwEngine3d).filler_069,
	(*HwEngine3d).filler_06a,
	(*HwEngine3d).filler_06b,
	(*HwEngine3d).filler_06c,
	(*HwEngine3d).filler_06d,
	(*HwEngine3d).filler_06e,
	(*HwEngine3d).filler_06f,
	(*HwEngine3d).filler_070,
	(*HwEngine3d).filler_071,
	(*HwEngine3d).filler_072,
	(*HwEngine3d).filler_073,
	(*HwEngine3d).filler_074,
	(*HwEngine3d).filler_075,
	(*HwEngine3d).filler_076,
	(*HwEngine3d).filler_077,
	(*HwEngine3d).filler_078,
	(*HwEngine3d).filler_079,
	(*HwEngine3d).filler_07a,
	(*HwEngine3d).filler_07b,
	(*HwEngine3d).filler_07c,
	(*HwEngine3d).filler_07d,
	(*HwEngine3d).filler_07e,
	(*HwEngine3d).filler_07f,
	(*HwEngine3d).filler_080,
	(*HwEngine3d).filler_081,
	(*HwEngine3d).filler_082,
	(*HwEngine3d).filler_083,
	(*HwEngine3d).filler_084,
	(*HwEngine3d).filler_085,
	(*HwEngine3d).filler_086,
	(*HwEngine3d).filler_087,
	(*HwEngine3d).filler_088,
	(*HwEngine3d).filler_089,
	(*HwEngine3d).filler_08a,
	(*HwEngine3d).filler_08b,
	(*HwEngine3d).filler_08c,
	(*HwEngine3d).filler_08d,
	(*HwEngine3d).filler_08e,
	(*HwEngine3d).filler_08f,
	(*HwEngine3d).filler_090,
	(*HwEngine3d).filler_091,
	(*HwEngine3d).filler_092,
	(*HwEngine3d).filler_093,
	(*HwEngine3d).filler_094,
	(*HwEngine3d).filler_095,
	(*HwEngine3d).filler_096,
	(*HwEngine3d).filler_097,
	(*HwEngine3d).filler_098,
	(*HwEngine3d).filler_099,
	(*HwEngine3d).filler_09a,
	(*HwEngine3d).filler_09b,
	(*HwEngine3d).filler_09c,
	(*HwEngine3d).filler_09d,
	(*HwEngine3d).filler_09e,
	(*HwEngine3d).filler_09f,
	(*HwEngine3d).filler_0a0,
	(*HwEngine3d).filler_0a1,
	(*HwEngine3d).filler_0a2,
	(*HwEngine3d).filler_0a3,
	(*HwEngine3d).filler_0a4,
	(*HwEngine3d).filler_0a5,
	(*HwEngine3d).filler_0a6,
	(*HwEngine3d).filler_0a7,
	(*HwEngine3d).filler_0a8,
	(*HwEngine3d).filler_0a9,
	(*HwEngine3d).filler_0aa,
	(*HwEngine3d).filler_0ab,
	(*HwEngine3d).filler_0ac,
	(*HwEngine3d).filler_0ad,
	(*HwEngine3d).filler_0ae,
	(*HwEngine3d).filler_0af,
	(*HwEngine3d).filler_0b0,
	(*HwEngine3d).filler_0b1,
	(*HwEngine3d).filler_0b2,
	(*HwEngine3d).filler_0b3,
	(*HwEngine3d).filler_0b4,
	(*HwEngine3d).filler_0b5,
	(*HwEngine3d).filler_0b6,
	(*HwEngine3d).filler_0b7,
	(*HwEngine3d).filler_0b8,
	(*HwEngine3d).filler_0b9,
	(*HwEngine3d).filler_0ba,
	(*HwEngine3d).filler_0bb,
	(*HwEngine3d).filler_0bc,
	(*HwEngine3d).filler_0bd,
	(*HwEngine3d).filler_0be,
	(*HwEngine3d).filler_0bf,
	(*HwEngine3d).filler_0c0,
	(*HwEngine3d).filler_0c1,
	(*HwEngine3d).filler_0c2,
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
	(*HwEngine3d).filler_0d8,
	(*HwEngine3d).filler_0d9,
	(*HwEngine3d).filler_0da,
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
	(*HwEngine3d).filler_0e7,
	(*HwEngine3d).filler_0e8,
	(*HwEngine3d).filler_0e9,
	(*HwEngine3d).filler_0ea,
	(*HwEngine3d).filler_0eb,
	(*HwEngine3d).filler_0ec,
	(*HwEngine3d).filler_0ed,
	(*HwEngine3d).filler_0ee,
	(*HwEngine3d).filler_0ef,
	(*HwEngine3d).filler_0f0,
	(*HwEngine3d).filler_0f1,
	(*HwEngine3d).filler_0f2,
	(*HwEngine3d).filler_0f3,
	(*HwEngine3d).filler_0f4,
	(*HwEngine3d).filler_0f5,
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
	(*HwEngine3d).filler_102,
	(*HwEngine3d).filler_103,
	(*HwEngine3d).filler_104,
	(*HwEngine3d).filler_105,
	(*HwEngine3d).filler_106,
	(*HwEngine3d).filler_107,
	(*HwEngine3d).filler_108,
	(*HwEngine3d).filler_109,
	(*HwEngine3d).filler_10a,
	(*HwEngine3d).filler_10b,
	(*HwEngine3d).filler_10c,
	(*HwEngine3d).filler_10d,
	(*HwEngine3d).filler_10e,
	(*HwEngine3d).filler_10f,
	(*HwEngine3d).filler_110,
	(*HwEngine3d).filler_111,
	(*HwEngine3d).filler_112,
	(*HwEngine3d).filler_113,
	(*HwEngine3d).filler_114,
	(*HwEngine3d).filler_115,
	(*HwEngine3d).filler_116,
	(*HwEngine3d).filler_117,
	(*HwEngine3d).filler_118,
	(*HwEngine3d).filler_119,
	(*HwEngine3d).filler_11a,
	(*HwEngine3d).filler_11b,
	(*HwEngine3d).filler_11c,
	(*HwEngine3d).filler_11d,
	(*HwEngine3d).filler_11e,
	(*HwEngine3d).filler_11f,
	(*HwEngine3d).filler_120,
	(*HwEngine3d).filler_121,
	(*HwEngine3d).filler_122,
	(*HwEngine3d).filler_123,
	(*HwEngine3d).filler_124,
	(*HwEngine3d).filler_125,
	(*HwEngine3d).filler_126,
	(*HwEngine3d).filler_127,
	(*HwEngine3d).filler_128,
	(*HwEngine3d).filler_129,
	(*HwEngine3d).filler_12a,
	(*HwEngine3d).filler_12b,
	(*HwEngine3d).filler_12c,
	(*HwEngine3d).filler_12d,
	(*HwEngine3d).filler_12e,
	(*HwEngine3d).filler_12f,
	(*HwEngine3d).filler_130,
	(*HwEngine3d).filler_131,
	(*HwEngine3d).filler_132,
	(*HwEngine3d).filler_133,
	(*HwEngine3d).filler_134,
	(*HwEngine3d).filler_135,
	(*HwEngine3d).filler_136,
	(*HwEngine3d).filler_137,
	(*HwEngine3d).filler_138,
	(*HwEngine3d).filler_139,
	(*HwEngine3d).filler_13a,
	(*HwEngine3d).filler_13b,
	(*HwEngine3d).filler_13c,
	(*HwEngine3d).filler_13d,
	(*HwEngine3d).filler_13e,
	(*HwEngine3d).filler_13f,
}
