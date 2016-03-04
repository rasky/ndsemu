// Generated on 2016-03-04 03:17:10.983083827 +0100 CET
package raster3d

import "ndsemu/emu/gfx"
import "ndsemu/emu"

func (e3d *HwEngine3d) filler_00(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:0 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_01(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_02(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_03(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_04(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_05(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_06(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_07(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_08(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:0 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_09(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_0a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_0b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_0c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_0d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_0e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_0f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:0 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_10(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:1 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_11(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_12(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_13(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_14(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_15(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_16(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_17(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_18(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:1 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_19(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_1a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_1b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_1c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_1d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_1e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_1f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:1 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
			pxa = uint8((int32(pxa+1)*int32(polyalpha+1) - 1) >> 6)
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_20(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:2 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_21(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_22(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_23(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_24(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_25(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_26(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_27(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_28(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:2 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_29(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_2a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_2b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_2c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_2d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_2e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_2f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:2 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_30(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:3 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_31(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_32(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_33(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_34(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_35(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_36(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_37(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_38(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:3 ColorMode:0}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_39(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_3a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_3b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_3c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_3d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_3e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_3f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:3 ColorMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Modulate(c0)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_40(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:0 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_41(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_42(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_43(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_44(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_45(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_46(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_47(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_48(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:0 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_49(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_4a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_4b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_4c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_4d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_4e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_4f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:0 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_50(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:1 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_51(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_52(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_53(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_54(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_55(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_56(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_57(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_58(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:1 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_59(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_5a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_5b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_5c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_5d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_5e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_5f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:1 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_60(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:2 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_61(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_62(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_63(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_64(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_65(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_66(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_67(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_68(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:2 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_69(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_6a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_6b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_6c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_6d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_6e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_6f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:2 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_70(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:3 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_71(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_72(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_73(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_74(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_75(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_76(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_77(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_78(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:3 ColorMode:1}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_79(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_7a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_7b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_7c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_7d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_7e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_7f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:3 ColorMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			pxc := newColorFrom555U(px)
			pxc = pxc.Decal(c0, pxa)
			px = pxc.To555U()
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_80(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:0 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_81(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_82(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_83(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_84(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_85(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_86(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_87(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_88(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:0 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_89(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_8a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_8b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_8c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_8d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_8e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_8f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:0 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_90(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:1 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_91(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_92(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_93(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_94(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_95(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_96(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_97(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_98(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:1 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_99(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_9a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_9b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_9c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_9d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_9e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_9f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:1 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
			pxa = polyalpha
		}
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:2 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:2 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_a9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_aa(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ab(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ac(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ad(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ae(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_af(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:2 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:3 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:3 ColorMode:2}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_b9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ba(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_bb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_bc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_bd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_be(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_bf(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:3 ColorMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		if true {
			px = emu.Read16LE(e3d.ToonTable.Data[(px&0x1F)*2:])
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:0 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:0 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_c9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ca(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_cb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_cc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_cd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ce(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_cf(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:0 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:1 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:1 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_d9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_da(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_db(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_dc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_dd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_de(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_df(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:1 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		pxa = polyalpha
		if pxa == 0 {
			goto next
		}
		if true {
			bkg := out.Get16(0)
			bkga := abuf.Get8(0)
			if bkga != 0 {
				px = rgbAlphaMix(px, bkg, pxa>>1)
			}
			if pxa > bkga {
				abuf.Set8(0, pxa)
			}
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add8(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:2 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:2 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_e9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ea(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_eb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ec(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ed(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ee(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ef(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:2 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f0(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:3 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f1(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f2(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f3(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f4(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f5(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f6(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f7(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f8(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:3 ColorMode:3}
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_f9(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = (px0 >> 5)
		pxa = pxa | (pxa << 3)
		px0 &= 0x1F
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_fa(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px0 = px0 >> (2 * uint(s&3))
		px0 &= 0x3
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_fb(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px0 = px0 >> (4 * uint(s&1))
		px0 &= 0xF
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_fc(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px0 == 0 {
			goto next
		}
		px = palette.Lookup(px0)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_fd(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = decompTex.Get16(int(t<<tshift + s))
		if px == 0 {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_fe(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px0 = e3d.texVram.Get8(texoff + t<<tshift + s)
		pxa = px0 >> 3
		pxa = (pxa >> 5) | (pxa << 1)
		px0 &= 0x7
		px0 <<= 2
		if px0 == 0 {
			goto next
		}
		px = uint16(px0) | uint16(px0)<<5 | uint16(px0)<<10
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}
func (e3d *HwEngine3d) filler_ff(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:3 ColorMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
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
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		if px&0x8000 != 0 {
			pxa = 63
		}
		px &= 0x7FFF
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		c0 = c0.AddDelta(dc)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
	_ = px0
	_ = pxa
}

var polygonFillerTable = [256]func(*HwEngine3d, *Polygon, gfx.Line, gfx.Line, gfx.Line){
	(*HwEngine3d).filler_00,
	(*HwEngine3d).filler_01,
	(*HwEngine3d).filler_02,
	(*HwEngine3d).filler_03,
	(*HwEngine3d).filler_04,
	(*HwEngine3d).filler_05,
	(*HwEngine3d).filler_06,
	(*HwEngine3d).filler_07,
	(*HwEngine3d).filler_08,
	(*HwEngine3d).filler_09,
	(*HwEngine3d).filler_0a,
	(*HwEngine3d).filler_0b,
	(*HwEngine3d).filler_0c,
	(*HwEngine3d).filler_0d,
	(*HwEngine3d).filler_0e,
	(*HwEngine3d).filler_0f,
	(*HwEngine3d).filler_10,
	(*HwEngine3d).filler_11,
	(*HwEngine3d).filler_12,
	(*HwEngine3d).filler_13,
	(*HwEngine3d).filler_14,
	(*HwEngine3d).filler_15,
	(*HwEngine3d).filler_16,
	(*HwEngine3d).filler_17,
	(*HwEngine3d).filler_18,
	(*HwEngine3d).filler_19,
	(*HwEngine3d).filler_1a,
	(*HwEngine3d).filler_1b,
	(*HwEngine3d).filler_1c,
	(*HwEngine3d).filler_1d,
	(*HwEngine3d).filler_1e,
	(*HwEngine3d).filler_1f,
	(*HwEngine3d).filler_20,
	(*HwEngine3d).filler_21,
	(*HwEngine3d).filler_22,
	(*HwEngine3d).filler_23,
	(*HwEngine3d).filler_24,
	(*HwEngine3d).filler_25,
	(*HwEngine3d).filler_26,
	(*HwEngine3d).filler_27,
	(*HwEngine3d).filler_28,
	(*HwEngine3d).filler_29,
	(*HwEngine3d).filler_2a,
	(*HwEngine3d).filler_2b,
	(*HwEngine3d).filler_2c,
	(*HwEngine3d).filler_2d,
	(*HwEngine3d).filler_2e,
	(*HwEngine3d).filler_2f,
	(*HwEngine3d).filler_30,
	(*HwEngine3d).filler_31,
	(*HwEngine3d).filler_32,
	(*HwEngine3d).filler_33,
	(*HwEngine3d).filler_34,
	(*HwEngine3d).filler_35,
	(*HwEngine3d).filler_36,
	(*HwEngine3d).filler_37,
	(*HwEngine3d).filler_38,
	(*HwEngine3d).filler_39,
	(*HwEngine3d).filler_3a,
	(*HwEngine3d).filler_3b,
	(*HwEngine3d).filler_3c,
	(*HwEngine3d).filler_3d,
	(*HwEngine3d).filler_3e,
	(*HwEngine3d).filler_3f,
	(*HwEngine3d).filler_40,
	(*HwEngine3d).filler_41,
	(*HwEngine3d).filler_42,
	(*HwEngine3d).filler_43,
	(*HwEngine3d).filler_44,
	(*HwEngine3d).filler_45,
	(*HwEngine3d).filler_46,
	(*HwEngine3d).filler_47,
	(*HwEngine3d).filler_48,
	(*HwEngine3d).filler_49,
	(*HwEngine3d).filler_4a,
	(*HwEngine3d).filler_4b,
	(*HwEngine3d).filler_4c,
	(*HwEngine3d).filler_4d,
	(*HwEngine3d).filler_4e,
	(*HwEngine3d).filler_4f,
	(*HwEngine3d).filler_50,
	(*HwEngine3d).filler_51,
	(*HwEngine3d).filler_52,
	(*HwEngine3d).filler_53,
	(*HwEngine3d).filler_54,
	(*HwEngine3d).filler_55,
	(*HwEngine3d).filler_56,
	(*HwEngine3d).filler_57,
	(*HwEngine3d).filler_58,
	(*HwEngine3d).filler_59,
	(*HwEngine3d).filler_5a,
	(*HwEngine3d).filler_5b,
	(*HwEngine3d).filler_5c,
	(*HwEngine3d).filler_5d,
	(*HwEngine3d).filler_5e,
	(*HwEngine3d).filler_5f,
	(*HwEngine3d).filler_60,
	(*HwEngine3d).filler_61,
	(*HwEngine3d).filler_62,
	(*HwEngine3d).filler_63,
	(*HwEngine3d).filler_64,
	(*HwEngine3d).filler_65,
	(*HwEngine3d).filler_66,
	(*HwEngine3d).filler_67,
	(*HwEngine3d).filler_68,
	(*HwEngine3d).filler_69,
	(*HwEngine3d).filler_6a,
	(*HwEngine3d).filler_6b,
	(*HwEngine3d).filler_6c,
	(*HwEngine3d).filler_6d,
	(*HwEngine3d).filler_6e,
	(*HwEngine3d).filler_6f,
	(*HwEngine3d).filler_70,
	(*HwEngine3d).filler_71,
	(*HwEngine3d).filler_72,
	(*HwEngine3d).filler_73,
	(*HwEngine3d).filler_74,
	(*HwEngine3d).filler_75,
	(*HwEngine3d).filler_76,
	(*HwEngine3d).filler_77,
	(*HwEngine3d).filler_78,
	(*HwEngine3d).filler_79,
	(*HwEngine3d).filler_7a,
	(*HwEngine3d).filler_7b,
	(*HwEngine3d).filler_7c,
	(*HwEngine3d).filler_7d,
	(*HwEngine3d).filler_7e,
	(*HwEngine3d).filler_7f,
	(*HwEngine3d).filler_80,
	(*HwEngine3d).filler_81,
	(*HwEngine3d).filler_82,
	(*HwEngine3d).filler_83,
	(*HwEngine3d).filler_84,
	(*HwEngine3d).filler_85,
	(*HwEngine3d).filler_86,
	(*HwEngine3d).filler_87,
	(*HwEngine3d).filler_88,
	(*HwEngine3d).filler_89,
	(*HwEngine3d).filler_8a,
	(*HwEngine3d).filler_8b,
	(*HwEngine3d).filler_8c,
	(*HwEngine3d).filler_8d,
	(*HwEngine3d).filler_8e,
	(*HwEngine3d).filler_8f,
	(*HwEngine3d).filler_90,
	(*HwEngine3d).filler_91,
	(*HwEngine3d).filler_92,
	(*HwEngine3d).filler_93,
	(*HwEngine3d).filler_94,
	(*HwEngine3d).filler_95,
	(*HwEngine3d).filler_96,
	(*HwEngine3d).filler_97,
	(*HwEngine3d).filler_98,
	(*HwEngine3d).filler_99,
	(*HwEngine3d).filler_9a,
	(*HwEngine3d).filler_9b,
	(*HwEngine3d).filler_9c,
	(*HwEngine3d).filler_9d,
	(*HwEngine3d).filler_9e,
	(*HwEngine3d).filler_9f,
	(*HwEngine3d).filler_a0,
	(*HwEngine3d).filler_a1,
	(*HwEngine3d).filler_a2,
	(*HwEngine3d).filler_a3,
	(*HwEngine3d).filler_a4,
	(*HwEngine3d).filler_a5,
	(*HwEngine3d).filler_a6,
	(*HwEngine3d).filler_a7,
	(*HwEngine3d).filler_a8,
	(*HwEngine3d).filler_a9,
	(*HwEngine3d).filler_aa,
	(*HwEngine3d).filler_ab,
	(*HwEngine3d).filler_ac,
	(*HwEngine3d).filler_ad,
	(*HwEngine3d).filler_ae,
	(*HwEngine3d).filler_af,
	(*HwEngine3d).filler_b0,
	(*HwEngine3d).filler_b1,
	(*HwEngine3d).filler_b2,
	(*HwEngine3d).filler_b3,
	(*HwEngine3d).filler_b4,
	(*HwEngine3d).filler_b5,
	(*HwEngine3d).filler_b6,
	(*HwEngine3d).filler_b7,
	(*HwEngine3d).filler_b8,
	(*HwEngine3d).filler_b9,
	(*HwEngine3d).filler_ba,
	(*HwEngine3d).filler_bb,
	(*HwEngine3d).filler_bc,
	(*HwEngine3d).filler_bd,
	(*HwEngine3d).filler_be,
	(*HwEngine3d).filler_bf,
	(*HwEngine3d).filler_c0,
	(*HwEngine3d).filler_c1,
	(*HwEngine3d).filler_c2,
	(*HwEngine3d).filler_c3,
	(*HwEngine3d).filler_c4,
	(*HwEngine3d).filler_c5,
	(*HwEngine3d).filler_c6,
	(*HwEngine3d).filler_c7,
	(*HwEngine3d).filler_c8,
	(*HwEngine3d).filler_c9,
	(*HwEngine3d).filler_ca,
	(*HwEngine3d).filler_cb,
	(*HwEngine3d).filler_cc,
	(*HwEngine3d).filler_cd,
	(*HwEngine3d).filler_ce,
	(*HwEngine3d).filler_cf,
	(*HwEngine3d).filler_d0,
	(*HwEngine3d).filler_d1,
	(*HwEngine3d).filler_d2,
	(*HwEngine3d).filler_d3,
	(*HwEngine3d).filler_d4,
	(*HwEngine3d).filler_d5,
	(*HwEngine3d).filler_d6,
	(*HwEngine3d).filler_d7,
	(*HwEngine3d).filler_d8,
	(*HwEngine3d).filler_d9,
	(*HwEngine3d).filler_da,
	(*HwEngine3d).filler_db,
	(*HwEngine3d).filler_dc,
	(*HwEngine3d).filler_dd,
	(*HwEngine3d).filler_de,
	(*HwEngine3d).filler_df,
	(*HwEngine3d).filler_e0,
	(*HwEngine3d).filler_e1,
	(*HwEngine3d).filler_e2,
	(*HwEngine3d).filler_e3,
	(*HwEngine3d).filler_e4,
	(*HwEngine3d).filler_e5,
	(*HwEngine3d).filler_e6,
	(*HwEngine3d).filler_e7,
	(*HwEngine3d).filler_e8,
	(*HwEngine3d).filler_e9,
	(*HwEngine3d).filler_ea,
	(*HwEngine3d).filler_eb,
	(*HwEngine3d).filler_ec,
	(*HwEngine3d).filler_ed,
	(*HwEngine3d).filler_ee,
	(*HwEngine3d).filler_ef,
	(*HwEngine3d).filler_f0,
	(*HwEngine3d).filler_f1,
	(*HwEngine3d).filler_f2,
	(*HwEngine3d).filler_f3,
	(*HwEngine3d).filler_f4,
	(*HwEngine3d).filler_f5,
	(*HwEngine3d).filler_f6,
	(*HwEngine3d).filler_f7,
	(*HwEngine3d).filler_f8,
	(*HwEngine3d).filler_f9,
	(*HwEngine3d).filler_fa,
	(*HwEngine3d).filler_fb,
	(*HwEngine3d).filler_fc,
	(*HwEngine3d).filler_fd,
	(*HwEngine3d).filler_fe,
	(*HwEngine3d).filler_ff,
}
