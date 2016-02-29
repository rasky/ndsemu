// Generated on 2016-02-29 00:14:31.281061131 +0100 CET
package raster3d

import "ndsemu/emu/gfx"

func (e3d *HwEngine3d) filler_00(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:0 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
	}
}
func (e3d *HwEngine3d) filler_01(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:1 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		px &= 0x1F
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_02(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:2 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
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
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px = px >> (2 * uint(s&3))
		px &= 0x3
		out.Set16(0, palette.Lookup(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_03(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:3 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
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
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px = px >> (4 * uint(s&1))
		px &= 0xF
		out.Set16(0, palette.Lookup(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_04(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:4 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		out.Set16(0, palette.Lookup(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_05(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:5 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_06(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:6 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		px &= 0x7
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_07(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:7 ColorKey:false}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
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
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_08(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:0 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
	}
}
func (e3d *HwEngine3d) filler_09(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:1 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		px &= 0x1F
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_0a(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:2 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
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
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s/4)
		px = px >> (2 * uint(s&3))
		px &= 0x3
		if px == 0 {
			goto next
		}
		out.Set16(0, palette.Lookup(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_0b(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:3 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
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
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s/2)
		px = px >> (4 * uint(s&1))
		px &= 0xF
		if px == 0 {
			goto next
		}
		out.Set16(0, palette.Lookup(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_0c(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:4 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		if px == 0 {
			goto next
		}
		out.Set16(0, palette.Lookup(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_0d(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:5 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_0e(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:6 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	texoff := poly.tex.VramTexOffset
	tshift := poly.tex.PitchShift
	s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()
	t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()
	ds := s1.SubFixed(s0).Div(nx)
	dt := t1.SubFixed(t0).Div(nx)
	smask := poly.tex.SMask
	tmask := poly.tex.TMask
	var px uint8
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		px &= 0x7
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_0f(poly *Polygon, out gfx.Line, zbuf gfx.Line) {
	// {TexFormat:7 ColorKey:true}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
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
	var s, t uint32
	out.Add16(int(x0))
	for x := x0; x < x1; x++ {
		if false {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}

var polygonFillerTable = [16]func(*HwEngine3d, *Polygon, gfx.Line, gfx.Line){
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
}
