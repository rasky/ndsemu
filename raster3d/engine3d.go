package raster3d

import (
	"fmt"
	"image"
	icolor "image/color"
	"image/png"
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"ndsemu/raster3d/fillerconfig"
	"os"
	"sort"
	"sync"
)

var mod3d = log.NewModule("e3d")

type buffer3d struct {
	Vram []Vertex
	Pram []Polygon
}

func (b *buffer3d) Reset() {
	b.Vram = b.Vram[:0]
	b.Pram = b.Pram[:0]
}

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
type decompTexCache struct {
	sync.Mutex
	data map[uint32][]uint8
}

func (d *decompTexCache) Reset() {
	d.Lock()
	d.data = make(map[uint32][]uint8)
	d.Unlock()
}

func (d *decompTexCache) Get(pos uint32) []uint8 {
	d.Lock()
	tex := d.data[pos]
	d.Unlock()
	return tex
}

func (d *decompTexCache) Put(pos uint32, tex []uint8) {
	d.Lock()
	d.data[pos] = tex
	d.Unlock()
}

type HwEngine3d struct {
	Disp3dCnt hwio.Reg32 `hwio:"offset=0,rwmask=0x7FFF"`
	ToonTable hwio.Mem   `hwio:"bank=1,offset=0x80,size=0x40,writeonly"`

	// Channel to receive new primitives (sent by GxFifo)
	CmdCh chan interface{}

	// Current viewport (last received viewport command)
	viewport Primitive_SetViewport

	pool sync.Pool

	// Current vram/pram (being drawn)
	cur buffer3d

	// Next vram/pram (being accumulated for next frame)
	next buffer3d

	nextCh chan buffer3d

	// Texture/palette VRAM
	texVram VramTextureBank
	palVram VramTexturePaletteBank

	// Cache for decompressed textures. This currently handles
	// Tex4x4 format as it's too hard to polyfill directly from
	// the compressed format.
	decompTex decompTexCache

	framecnt int
}

func NewHwEngine3d() *HwEngine3d {
	e3d := new(HwEngine3d)
	hwio.MustInitRegs(e3d)

	e3d.CmdCh = make(chan interface{}, 4096)

	e3d.pool.New = func() interface{} {
		return buffer3d{
			Vram: make([]Vertex, 0, 8192),
			Pram: make([]Polygon, 0, 8192),
		}
	}
	e3d.next = e3d.pool.Get().(buffer3d)
	e3d.nextCh = make(chan buffer3d) // must be non buffered!

	go e3d.recvCmd()
	return e3d
}

func (e3d *HwEngine3d) recvCmd() {
	for {
		cmdi := <-e3d.CmdCh
		switch cmd := cmdi.(type) {
		case Primitive_SwapBuffers:
			e3d.cmdSwapBuffers()
		case Primitive_SetViewport:
			e3d.viewport = cmd
		case Primitive_Polygon:
			e3d.cmdPolygon(cmd)
		case Primitive_Vertex:
			e3d.cmdVertex(cmd)
		default:
			panic("invalid command received in HwEnginge3D")
		}
	}
}

func (vtx *Vertex) calcClippingFlags() {

	// Compute clipping flags (once per vertex)
	if vtx.cx.V < -vtx.cw.V {
		vtx.flags |= RVFClipLeft
	}
	if vtx.cx.V > vtx.cw.V {
		vtx.flags |= RVFClipRight
	}
	if vtx.cy.V < -vtx.cw.V {
		vtx.flags |= RVFClipBottom
	}
	if vtx.cy.V > vtx.cw.V {
		vtx.flags |= RVFClipTop
	}
	if vtx.cw.V < 0 {
		vtx.flags |= RVFClipNear
	}
	// if vtx.cz.V > vtx.cw.V {
	// 	vtx.flags |= RVFClipFar
	// }

	// If w==0, we just flag the vertex as fully outside of the screen
	// FIXME: properly handle invalid inputs
	// if vtx.cw.V == 0 {
	// 	vtx.flags |= RVFClipAnything
	// }
}

func (e3d *HwEngine3d) cmdVertex(cmd Primitive_Vertex) {
	vtx := Vertex{
		cx:  cmd.X,
		cy:  cmd.Y,
		cz:  cmd.Z,
		cw:  cmd.W,
		s:   cmd.S,
		t:   cmd.T,
		rgb: newColorFrom555(cmd.C[0], cmd.C[1], cmd.C[2]),
	}
	vtx.calcClippingFlags()
	e3d.next.Vram = append(e3d.next.Vram, vtx)
}

func (e3d *HwEngine3d) cmdPolygon(cmd Primitive_Polygon) {

	flags := PolygonFlags(cmd.Attr)

	var vinbuf [4]*Vertex
	vtxs := vinbuf[:0]

	// FIXME: for now, skip all polygons outside the screen
	count := 3
	if flags&PFQuad != 0 {
		count = 4
		flags &^= PFQuad
	}

	clipany := RenderVertexFlags(0)
	clipall := RVFClipMask
	for i := 0; i < count; i++ {
		if cmd.Vtx[i] >= len(e3d.next.Vram) || cmd.Vtx[i] < 0 {
			mod3d.Fatalf("wrong polygon index: %d (num vtx: %d)", cmd.Vtx[i], len(e3d.next.Vram))
		}
		vtx := &e3d.next.Vram[cmd.Vtx[i]]
		clipany |= (vtx.flags & RVFClipMask)
		clipall &= (vtx.flags & RVFClipMask)
		vtxs = append(vtxs, vtx)
	}

	// If all vertices are out of the same plane (any of them),
	// the polygon is fully out, so clip it.
	if clipall != 0 {
		return
	}

	// Transform all vertices (that weren't transformed already)
	for _, vtx := range vtxs {
		e3d.vtxTransform(vtx)
	}

	// Do backface culling
	d0x := vtxs[0].x.SubFixed(vtxs[1].x)
	d0y := vtxs[0].y.SubFixed(vtxs[1].y)
	d1x := vtxs[2].x.SubFixed(vtxs[1].x)
	d1y := vtxs[2].y.SubFixed(vtxs[1].y)
	if int64(d0x.V)*int64(d1y.V) <= int64(d1x.V)*int64(d0y.V) {
		// Facing the back: see if we must render the back
		if flags&PFRenderBack == 0 {
			return
		}
	} else {
		// Facing the front: see if we must render the front
		if flags&PFRenderFront == 0 {
			return
		}
	}

	// Do clipping
	if clipany != 0 {
		vtxs = e3d.polyClip(vtxs)
		if vtxs == nil {
			return
		}
	}

	// Transform all vertices (that weren't transformed already)
	for _, vtx := range vtxs {
		// vtx.calcClippingFlags()
		// if vtx.flags != 0 {
		// 	fmt.Printf("vtx:(%v,%v,%v,%v) clip=%x clipany=%x\n", vtx.cx, vtx.cy, vtx.cz, vtx.cw, vtx.flags, clipany)
		// 	panic("clipping failed")
		// }
		e3d.vtxTransform(vtx)
	}

	// Split the clipped polygon into triangles and add them to pram
	for i := 1; i < len(vtxs)-1; i++ {
		poly := Polygon{
			flags: flags,
			tex:   cmd.Tex,
			vtx: [3]*Vertex{
				vtxs[0], vtxs[i], vtxs[i+1],
			},
		}
		e3d.next.Pram = append(e3d.next.Pram, poly)
	}
}

func (v0 *Vertex) Lerp(v1 *Vertex, ratio emu.Fixed12) *Vertex {
	vout := new(Vertex)
	vout.cx = v0.cx.Lerp(v1.cx, ratio)
	vout.cy = v0.cy.Lerp(v1.cy, ratio)
	vout.cz = v0.cz.Lerp(v1.cz, ratio)
	vout.cw = v0.cw.Lerp(v1.cw, ratio)
	vout.s = v0.s.Lerp(v1.s, ratio)
	vout.t = v0.t.Lerp(v1.t, ratio)
	vout.rgb = v0.rgb.Lerp(v1.rgb, ratio)
	return vout
}

var clipFormulas = [...]struct {
	Plane         RenderVertexFlags
	PlaneCoord    func(*Vertex) emu.Fixed12
	PlaneSetCoord func(*Vertex, emu.Fixed12)
}{
	{
		RVFClipNear,
		func(v *Vertex) emu.Fixed12 { return emu.NewFixed12(0) },
		func(v *Vertex, f emu.Fixed12) {},
	},
	{
		RVFClipTop,
		func(v *Vertex) emu.Fixed12 { return v.cy },
		func(v *Vertex, f emu.Fixed12) { v.cy = f },
	},
	{
		RVFClipBottom,
		func(v *Vertex) emu.Fixed12 { return v.cy.Neg() },
		func(v *Vertex, f emu.Fixed12) { v.cy = f.Neg() },
	},
	{
		RVFClipLeft,
		func(v *Vertex) emu.Fixed12 { return v.cx.Neg() },
		func(v *Vertex, f emu.Fixed12) { v.cx = f.Neg() },
	},
	{
		RVFClipRight,
		func(v *Vertex) emu.Fixed12 { return v.cx },
		func(v *Vertex, f emu.Fixed12) { v.cx = f },
	},
}

func (e3d *HwEngine3d) polyClip(poly []*Vertex) (clipped []*Vertex) {
	// fmt.Printf("begin clipping\n")
	// defer fmt.Printf("end clipping\n")
	for _, clipInfo := range clipFormulas {

		// fmt.Printf("begin clipping: plane %x\n", clipInfo.Plane)
		// for _, v := range poly {
		// 	fmt.Printf("vtx: %v,%v,%v,%v flags=%x\n",
		// 		v.cx, v.cy, v.cz, v.cw,
		// 		v.flags)
		// }

		last := poly[len(poly)-1]
		lastOut := last.flags&clipInfo.Plane != 0

		for _, v := range poly {
			if (v.flags&clipInfo.Plane != 0 && !lastOut) ||
				(v.flags&clipInfo.Plane == 0 && lastOut) {

				v0 := last
				dist1 := v0.cw.SubFixed(clipInfo.PlaneCoord(v0))
				dist2 := v.cw.SubFixed(clipInfo.PlaneCoord(v))
				if dist1.SubFixed(dist2).V == 0 {
					return nil
				}
				ratio := dist1.DivFixed(dist1.SubFixed(dist2))

				vout := v0.Lerp(v, ratio)
				clipInfo.PlaneSetCoord(vout, vout.cw)
				// vout.cw = clipInfo.PlaneCoord(vout) // fix rounding errors
				vout.calcClippingFlags()
				// fmt.Printf("clip %d: %v,%v,%v,%v ratio:%v flags:%x\n",
				// 	idx,
				// 	vout.cx, vout.cy, vout.cz, vout.cw,
				// 	ratio, vout.flags)

				clipped = append(clipped, vout)
			}

			last = v
			if v.flags&clipInfo.Plane == 0 {
				lastOut = false
				clipped = append(clipped, v)
			} else {
				lastOut = true
			}
		}

		// for _, v := range clipped {
		// 	if v.flags&clipInfo.Plane != 0 {
		// 		panic("ahahah")
		// 	}
		// }

		// FIXME: check if it's really required
		if len(clipped) == 0 {
			return nil
		}

		poly = clipped
		clipped = nil
	}

	for _, v := range poly {
		v.cx = v.cx.Clamp(v.cw.Neg(), v.cw)
		v.cy = v.cy.Clamp(v.cw.Neg(), v.cw)
		v.cz = v.cz.Clamp(v.cw.Neg(), v.cw)

		// 	if v.flags&RVFClipMask != 0 {
		// 		fmt.Printf("vtx: %v,%v,%v,%v flags=%x\n", v.cx, v.cy, v.cz, v.cw, v.flags)
		// 		panic("ahahah2")
		// 	}
	}

	return poly
}

func (e3d *HwEngine3d) vtxTransform(vtx *Vertex) {
	if vtx.flags&RVFTransformed != 0 {
		return
	}

	viewwidth := emu.NewFixed12(int32(e3d.viewport.VX1 - e3d.viewport.VX0))
	viewheight := emu.NewFixed12(int32(e3d.viewport.VY1 - e3d.viewport.VY0))
	// Compute viewsize / (2*v.w) in two steps, to avoid overflows
	// (viewwidth could be 256<<12, which would overflow when further
	// shifted in preparation for division)
	dx := viewwidth.Div(2).DivFixed(vtx.cw)
	dy := viewheight.Div(2).DivFixed(vtx.cw)

	mirror := vtx.cw.Mul(2)

	// sx = (v.x + v.w) * viewwidth / (2*v.w) + viewx0
	// sy = (v.y + v.w) * viewheight / (2*v.w) + viewy0
	vtx.x = vtx.cx.AddFixed(vtx.cw).MulFixed(dx).Add(int32(e3d.viewport.VX0)).Round()
	vtx.y = mirror.SubFixed(vtx.cy.AddFixed(vtx.cw)).MulFixed(dy).Add(int32(e3d.viewport.VY0)).Round()
	vtx.z = vtx.cw // vtx.cz.AddFixed(vtx.cw).Div(2).DivFixed(vtx.cw)

	// Clamp screen coord. This is only required because clipping in clip-space
	// cannot be accurate with fixed point coordinates (at least not with 12 bit),
	// and thus it can generate coordinates that are slightly out
	vtx.x = vtx.x.Clamp(emu.NewFixed12(int32(e3d.viewport.VX0)), emu.NewFixed12(int32(e3d.viewport.VX1)))
	vtx.y = vtx.y.Clamp(emu.NewFixed12(int32(e3d.viewport.VY0)), emu.NewFixed12(int32(e3d.viewport.VY1)))

	vtx.flags |= RVFTransformed
}

func (e3d *HwEngine3d) preparePolys() {

	for idx := range e3d.next.Pram {
		poly := &e3d.next.Pram[idx]
		v0, v1, v2 := poly.vtx[0], poly.vtx[1], poly.vtx[2]

		// Sort the three vertices by the Y coordinate (v0=top, v1=middle, 2=bottom)
		if v0.y.V > v1.y.V {
			v0, v1 = v1, v0
			poly.vtx[0], poly.vtx[1] = poly.vtx[1], poly.vtx[0]
		}
		if v0.y.V > v2.y.V {
			v0, v2 = v2, v0
			poly.vtx[0], poly.vtx[2] = poly.vtx[2], poly.vtx[0]
		}
		if v1.y.V > v2.y.V {
			v1, v2 = v2, v1
			poly.vtx[1], poly.vtx[2] = poly.vtx[2], poly.vtx[1]
		}

		hy1 := v1.y.TruncInt32() - v0.y.TruncInt32()
		hy2 := v2.y.TruncInt32() - v1.y.TruncInt32()
		if hy1 < 0 || hy2 < 0 {
			panic("invalid y order")
		}

		// Calculate the four slopes for each coordinate.  The coordinates
		// we need to interpolate are: position (X), depth (Z), texture (S & T).
		//
		// Assuming a triangle where:
		//    * v0 is at top
		//    * v1 is middle left
		//    * v2 is bottom
		// we need two slopes for the left segments (from v0 to v1, and then from v1 to v2), and
		// one slope for the right segment (from v0 to v2). To make the line-based rasterizer
		// simpler, we consider the triangle virtually split in half at the v1
		// level, so we calculate two slopes for each half-triangle; in our example, the right
		// slopes for the upper and lower part will obviously be the same (as it's just one
		// segment).
		var dxl0, dxl1, dxr0, dxr1 emu.Fixed22
		var dzl0, dzl1, dzr0, dzr1 emu.Fixed22
		var dsl0, dsl1, dsr0, dsr1 emu.Fixed12
		var dtl0, dtl1, dtr0, dtr1 emu.Fixed12
		var dcl0, dcl1, dcr0, dcr1 colorDelta

		dxl0 = v1.x.SubFixed(v0.x).ToFixed22()
		dzl0 = v1.z.SubFixed(v0.z).ToFixed22()
		dsl0 = v1.s.SubFixed(v0.s)
		dtl0 = v1.t.SubFixed(v0.t)
		dcl0 = v1.rgb.SubColor(v0.rgb)

		dxl1 = v2.x.SubFixed(v1.x).ToFixed22()
		dzl1 = v1.z.SubFixed(v1.z).ToFixed22()
		dsl1 = v2.s.SubFixed(v1.s)
		dtl1 = v2.t.SubFixed(v1.t)
		dcl1 = v2.rgb.SubColor(v1.rgb)

		if hy1 > 0 {
			dxl0 = dxl0.Div(hy1)
			dzl0 = dzl0.Div(hy1)
			dsl0 = dsl0.Div(hy1)
			dtl0 = dtl0.Div(hy1)
			dcl0 = dcl0.Div(hy1)
		}
		if hy2 > 0 {
			dxl1 = dxl1.Div(hy2)
			dzl1 = dzl1.Div(hy2)
			dsl1 = dsl1.Div(hy2)
			dtl1 = dtl1.Div(hy2)
			dcl1 = dcl1.Div(hy2)
		}
		if hy1+hy2 > 0 {
			dxr0 = v2.x.SubFixed(v0.x).ToFixed22().Div(hy1 + hy2)
			dzr0 = v2.z.SubFixed(v0.z).ToFixed22().Div(hy1 + hy2)
			dsr0 = v2.s.SubFixed(v0.s).Div(hy1 + hy2)
			dtr0 = v2.t.SubFixed(v0.t).Div(hy1 + hy2)
			dcr0 = v2.rgb.SubColor(v0.rgb).Div(hy1 + hy2)

			dxr1 = dxr0
			dzr1 = dzr0
			dsr1 = dsr0
			dtr1 = dtr0
			dcr1 = dcr0
		}

		// Now create interpolator instances
		poly.left[LerpX] = newLerp(v0.x.ToFixed22(), dxl0, dxl1)
		poly.right[LerpX] = newLerp(v0.x.ToFixed22(), dxr0, dxr1)

		poly.left[LerpZ] = newLerp(v0.z.ToFixed22(), dzl0, dzl1)
		poly.right[LerpZ] = newLerp(v0.z.ToFixed22(), dzr0, dzr1)

		poly.left[LerpS] = newLerp12(v0.s, dsl0, dsl1)
		poly.right[LerpS] = newLerp12(v0.s, dsr0, dsr1)

		poly.left[LerpT] = newLerp12(v0.t, dtl0, dtl1)
		poly.right[LerpT] = newLerp12(v0.t, dtr0, dtr1)

		poly.left[LerpRGB] = newLerpFromInt(int32(v0.rgb), int32(dcl0), int32(dcl1))
		poly.right[LerpRGB] = newLerpFromInt(int32(v0.rgb), int32(dcr0), int32(dcr1))

		// If v0 and v1 lies on the same line (top segment), there is no upper
		// half of the triangle. In this case, we need the initial values of the lerp
		// to reflect this. Given that there was no divsion above, delta[0] is the
		// full different between v1 and v0, so we just need to add it to the start
		// coordinate (v0) to transform it into v1.
		if hy1 == 0 {
			if v0.x.V < v1.x.V {
				poly.left, poly.right = poly.right, poly.left
				for idx := range poly.left {
					rp := &poly.right[idx]
					rp.start += rp.delta[0]
				}
			} else {
				for idx := range poly.left {
					rp := &poly.left[idx]
					rp.start += rp.delta[0]
				}

			}

		} else {
			// We have assumed that the middle vertex is "on the left" (that is,
			// the segment between v0 and v1 is part of the left perimeter).
			// We check if that's the case, by simply comparing the calculated
			// slopes. If it's not true, we just swap all the calculated
			// interpolators.
			if dxl0.V > dxr0.V {
				poly.left, poly.right = poly.right, poly.left
			}
		}

		poly.hy = v1.y.TruncInt32()

		// if poly.flags.ColorMode() == fillerconfig.ColorModeToon {
		// 	left := poly.left[LerpRGB]
		// 	right := poly.right[LerpRGB]
		// 	left.Reset()
		// 	right.Reset()
		// 	log.Mod3d.Infof("toon polygon dump: %x,%x,%x", v0.rgb, v1.rgb, v2.rgb)
		// 	for i := 0; i < int(hy1); i++ {
		// 		log.Mod3d.Infof("step: %x,%x", left.cur, right.cur)
		// 		left.Next(0)
		// 		right.Next(0)
		// 	}
		// }

	}
}

type polySorter []Polygon

func (p polySorter) Len() int      { return len(p) }
func (p polySorter) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p polySorter) Less(i, j int) bool {
	ai := p[i].flags.Alpha()
	aj := p[j].flags.Alpha()
	return aj > ai
}

func (e3d *HwEngine3d) sortPolys() {
	// Do a stable sort, so that we keep the exisint order for all
	// solid polygon. This should be consistent with the order the NDS
	// renderes the display list.
	sort.Stable(polySorter(e3d.next.Pram))
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

func (e3d *HwEngine3d) decompressTextures() {

	e3d.decompTex.Reset()

	for _, poly := range e3d.cur.Pram {
		if poly.tex.Format != Tex4x4 {
			continue
		}

		off := poly.tex.VramTexOffset
		if buf := e3d.decompTex.Get(off); buf != nil {
			if len(buf) != int((poly.tex.SMask+1)*(poly.tex.TMask+1)*2) {
				panic("different compressed texture size in same frame")
			}
			continue
		}

		out := make([]uint8, (poly.tex.SMask+1)*(poly.tex.TMask+1)*2)

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
			off, xtraoff, poly.tex.SMask+1, poly.tex.TMask+1)

		for y := 0; y < int(poly.tex.TMask+1); y += 4 {
			for x := 0; x < int(poly.tex.SMask+1); x += 4 {
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

		if false {
			f, err := os.Create(fmt.Sprintf("tex-%x.png", poly.tex.VramTexOffset))
			if err == nil {
				png.Encode(f, &Image555{
					buf: out,
					w:   int(poly.tex.SMask + 1),
					h:   int(poly.tex.TMask + 1),
				})
				f.Close()
			}
		}

		e3d.decompTex.Put(poly.tex.VramTexOffset, out)
	}
}

func (e3d *HwEngine3d) dumpNextScene() {
	if e3d.framecnt == 0 {
		os.Remove("dump3d.txt")
	}
	f, err := os.OpenFile("dump3d.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Fprintf(f, "begin scene\n")
	for idx, poly := range e3d.next.Pram {
		v0, v1, v2 := poly.vtx[0], poly.vtx[1], poly.vtx[2]
		fmt.Fprintf(f, "tri %d:\n", idx)
		fmt.Fprintf(f, "    ccoord: (%v,%v,%v,%v)-(%v,%v,%v,%v)-(%v,%v,%v,%v)\n",
			v0.cx, v0.cy, v0.cz, v0.cw,
			v1.cx, v1.cy, v1.cz, v1.cw,
			v2.cx, v2.cy, v2.cz, v2.cw)
		fmt.Fprintf(f, "    scoord: (%v,%v)-(%v,%v)-(%v,%v)\n",
			v0.x.TruncInt32(), v0.y.TruncInt32(),
			v1.x.TruncInt32(), v1.y.TruncInt32(),
			v2.x.TruncInt32(), v2.y.TruncInt32())
		// fmt.Fprintf(f, "    tex: (%v,%v)-(%v,%v)-(%v,%v)\n",
		// 	v0.s, v0.t,
		// 	v1.s, v1.t,
		// 	v2.s, v2.t)
		// fmt.Fprintf(f, "    left lerps: %v\n", poly.left)
		// fmt.Fprintf(f, "    right lerps: %v\n", poly.right)
		// fmt.Fprintf(f, "    hy: %v\n", poly.hy)
		fmt.Fprintf(f, "    flags: %08x\n", poly.flags)
		fmt.Fprintf(f, "    tex: fmt=%d, flips=%v, flipt=%v, reps=%v, rept=%v\n",
			poly.tex.Format, poly.tex.Flags&TexSFlip != 0, poly.tex.Flags&TexTFlip != 0,
			poly.tex.Flags&TexSRepeat != 0, poly.tex.Flags&TexTRepeat != 0)
	}
	mod3d.Infof("end scene")
}

func (e3d *HwEngine3d) cmdSwapBuffers() {
	// The next frame primitives are complete; we can now do full-frame processing
	// in preparation for drawing next frame
	e3d.preparePolys()
	e3d.sortPolys()
	// e3d.dumpNextScene()
	e3d.framecnt++

	// Send the next buffer to the main rendering thread. Since the channel
	// is not buffered, this call will block until the other side reads, which is
	// at VBlank start. This is exactly what we expect from SwapBuffers: it blocks
	// until next VBlank.
	e3d.nextCh <- e3d.next

	// Get a new buffer from the pool, ready for next frame
	e3d.next = e3d.pool.Get().(buffer3d)
}

func (e3d *HwEngine3d) Draw3D(ctx *gfx.LayerCtx, lidx int, y int) {

	texMappingEnabled := e3d.Disp3dCnt.Value&1 != 0

	// Initialize rasterizer.
	var polyPerLine [192][]uint16
	for idx := range e3d.cur.Pram {
		poly := &e3d.cur.Pram[idx]

		// Set current segment to the initial one computed in preparePolys
		// This is required because we might need to redraw the exact
		// same 3D scene multiple times, so each time we want to start
		// from the beginning.
		for idx := range poly.left {
			poly.left[idx].Reset()
			poly.right[idx].Reset()
		}

		// If the polygon degrades to a segment, skip it for now (we don't support segments)
		v0, v1, v2 := poly.vtx[0], poly.vtx[1], poly.vtx[2]
		if (v0.x == v1.x && v0.y == v1.y) || (v1.x == v2.x && v1.y == v2.y) {
			// FIXME: implement segments
			continue
		}

		// Update the per-line polygon list, by adding this polygon's index
		// to the lines in which it is visible.
		for j := v0.y.TruncInt32(); j <= v2.y.TruncInt32(); j++ {
			polyPerLine[j] = append(polyPerLine[j], uint16(idx))
		}

		// Setup the correct polyfiller for this
		fcfg := fillerconfig.FillerConfig{
			TexFormat: uint(poly.tex.Format),
			ColorMode: poly.flags.ColorMode(),
		}
		if poly.tex.Transparency {
			fcfg.ColorKey = 1
		}
		if !texMappingEnabled {
			fcfg.TexFormat = 0
		}
		switch poly.flags.Alpha() {
		case 31:
			fcfg.FillMode = fillerconfig.FillModeSolid
		case 0:
			fcfg.FillMode = fillerconfig.FillModeWireframe
		default:
			fcfg.FillMode = fillerconfig.FillModeAlpha
		}
		if fcfg.ColorMode == fillerconfig.ColorModeToon && e3d.Disp3dCnt.Value&(1<<1) != 0 {
			fcfg.ColorMode = fillerconfig.ColorModeHighlight
		}

		// log.Mod3d.Infof("polygon: %d - %+v (key:%x, alpha: %d)", idx, fcfg, fcfg.Key(), poly.flags.Alpha())
		// if fcfg.ColorMode == fillerconfig.ColorModeToon {
		// 	log.Mod3d.Infof("toon: (%x,%x,%x)-(%x,%x,%x)\n",
		// 		poly.left[LerpRGB].cur, poly.left[LerpRGB].delta[0], poly.left[LerpRGB].delta[1],
		// 		poly.right[LerpRGB].cur, poly.right[LerpRGB].delta[0], poly.right[LerpRGB].delta[1])
		// }

		poly.filler = polygonFillerTable[fcfg.Key()]
	}

	// FIXME: move elsewhere. Theoretically, 3d rendering begins
	// 19 lines before the first screen line, so it would be possible
	// to begin processing textures somewhere in the middle of vblank.
	// To be 100% sure, we can't do that when we receieve SwapBuffers
	// (that is, in the middle of previous frame) as the texture data
	// could not be ready.
	e3d.decompressTextures()

	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		if e3d.Disp3dCnt.Value&(1<<14) != 0 {
			panic("bitmap")
		}

		var abuf [256]byte
		var zbuf [256 * 4]byte
		zbuffer := gfx.NewLine(zbuf[:])
		abuffer := gfx.NewLine(abuf[:])
		for i := 0; i < 256; i++ {
			zbuffer.Set32(i, 0x7FFFFFFF)
			abuffer.Set8(i, 0x1F)
		}

		for _, idx := range polyPerLine[y] {
			poly := &e3d.cur.Pram[idx]

			x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
			if x0 < 0 || x1 >= 256 || x1 < x0 {
				fmt.Printf("%v,%v\n", poly.vtx[0].x.TruncInt32(), poly.vtx[0].y.TruncInt32())
				fmt.Printf("%v,%v\n", poly.vtx[1].x.TruncInt32(), poly.vtx[1].y.TruncInt32())
				fmt.Printf("%v,%v\n", poly.vtx[2].x.TruncInt32(), poly.vtx[2].y.TruncInt32())
				fmt.Printf("left lerps: %v\n", poly.left)
				fmt.Printf("right lerps: %v\n", poly.right)
				fmt.Printf("x0,x1=%v,%v   y=%v, hy=%v\n", x0, x1, y, poly.hy)
				// panic("out of bounds")
			} else {
				poly.filler(e3d, poly, line, zbuffer, abuffer)
			}

			if int32(y) < poly.hy {
				for idx := 0; idx < NumLerps; idx++ {
					poly.left[idx].Next(0)
					poly.right[idx].Next(0)
				}
			} else {
				for idx := 0; idx < NumLerps; idx++ {
					poly.left[idx].Next(1)
					poly.right[idx].Next(1)
				}
			}
		}

		/*
			// TO BE TESTED
			for i := 0; i < 256; i++ {
				if abuffer.Get8(i) == 0 {
					line.Set16(i, 0)
				}
			}
		*/
		y++
	}
}

func (e3d *HwEngine3d) SetVram(tex VramTextureBank, pal VramTexturePaletteBank) {
	e3d.texVram = tex
	e3d.palVram = pal
}

func (e3d *HwEngine3d) BeginFrame() {
}

func (e3d *HwEngine3d) EndFrame() {
	// We're now at vblank start. Read the pending buffer from SwapBuffers (if any).
	select {
	case next := <-e3d.nextCh:
		// OK got a new buffer. Recycle the current one into the pool
		e3d.cur.Reset()
		e3d.pool.Put(e3d.cur)
		e3d.cur = next
	default:
		// If there's no pending buffer, then it means that there was no new geometry
		// commands, or the commands are taking more than 1/60th of second to be elaborated;
		// in any case, it's too late; the same frame will be drawn again.
	}
}

func (e3d *HwEngine3d) NumVertices() int {
	return len(e3d.cur.Vram)
}

func (e3d *HwEngine3d) NumPolygons() int {
	return len(e3d.cur.Pram)
}
