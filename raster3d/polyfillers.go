// Generated on 2016-03-03 02:56:17.894222929 +0100 CET
package raster3d

import "ndsemu/emu/gfx"

func (e3d *HwEngine3d) filler_00(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
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
	}
}
func (e3d *HwEngine3d) filler_01(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_02(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_03(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_04(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_05(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:0}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_06(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_07(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_08(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:0}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
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
	}
}
func (e3d *HwEngine3d) filler_09(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_0a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_0b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_0c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_0d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:0}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_0e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_0f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:0}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_10(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
	}
}
func (e3d *HwEngine3d) filler_11(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_12(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_13(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_14(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get8(texoff + t<<tshift + s)
		out.Set16(0, palette.Lookup(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_15(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:1}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_16(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_17(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_18(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:1}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
	out.Add16(int(x0))
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
	}
}
func (e3d *HwEngine3d) filler_19(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_1a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_1b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_1c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_1d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:1}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_1e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_1f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:1}
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
	zbuf.Add32(int(x0))
	abuf.Add8(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
			goto next
		}
		s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
		px = e3d.texVram.Get16(texoff + t<<tshift + s)
		out.Set16(0, uint16(px)|0x8000)
		zbuf.Set32(0, uint32(z0.V))
	next:
		out.Add16(1)
		zbuf.Add32(1)
		abuf.Add32(1)
		z0 = z0.AddFixed(dz)
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_20(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
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
	}
}
func (e3d *HwEngine3d) filler_21(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_22(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_23(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_24(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_25(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:2}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_26(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_27(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_28(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:2}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
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
	}
}
func (e3d *HwEngine3d) filler_29(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_2a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_2b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_2c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_2d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:2}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_2e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_2f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:2}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_30(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:false FillMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
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
	}
}
func (e3d *HwEngine3d) filler_31(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:false FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_32(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:false FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_33(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:false FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_34(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:false FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_35(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:false FillMode:3}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_36(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:false FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_37(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:false FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_38(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:0 ColorKey:true FillMode:3}
	x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
	nx := x1 - x0
	if nx == 0 {
		return
	}
	z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()
	dz := z1.SubFixed(z0).Div(nx)
	var px uint8
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
	}
}
func (e3d *HwEngine3d) filler_39(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:1 ColorKey:true FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_3a(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:2 ColorKey:true FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_3b(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:3 ColorKey:true FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_3c(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:4 ColorKey:true FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_3d(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:5 ColorKey:true FillMode:3}
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
	decompTexBuf := e3d.decompTex.Get(texoff)
	decompTex := gfx.NewLine(decompTexBuf)
	var px uint16
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
		s0 = s0.AddFixed(ds)
		t0 = t0.AddFixed(dt)
	}
}
func (e3d *HwEngine3d) filler_3e(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:6 ColorKey:true FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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
func (e3d *HwEngine3d) filler_3f(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {
	// {TexFormat:7 ColorKey:true FillMode:3}
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
	zbuf.Add32(int(x0))
	for x := x0; x < x1; x++ {
		if z0.V >= int32(zbuf.Get32(0)) {
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

var polygonFillerTable = [64]func(*HwEngine3d, *Polygon, gfx.Line, gfx.Line, gfx.Line){
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
}
