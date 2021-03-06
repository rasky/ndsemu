package raster3d

import (
	"fmt"
	"ndsemu/emu/fixed"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hw"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"ndsemu/raster3d/fillerconfig"
	"os"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

var mod3d = log.NewModule("e3d")

type buffer3d struct {
	Pram     []Polygon
	Vram     []Vertex
	ClipVram []Vertex
}

func (b *buffer3d) Reset() {
	b.Pram = b.Pram[:0]
	b.Vram = b.Vram[:0]
	b.ClipVram = b.ClipVram[:0]
}

type HwEngine3d struct {
	Disp3dCnt  hwio.Reg32 `hwio:"offset=0,rwmask=0x7FFF"`
	ToonTable  hwio.Mem   `hwio:"bank=1,offset=0x80,size=0x40,writeonly"`
	ClearColor hwio.Reg32 `hwio:"bank=1,offset=0x50,writeonly"`
	ClearDepth hwio.Reg32 `hwio:"bank=1,offset=0x54,writeonly"`
	FogColor   hwio.Reg32 `hwio:"bank=1,offset=0x58,writeonly"`
	FogOffset  hwio.Reg32 `hwio:"bank=1,offset=0x5C,rwmask=0x7FFF,writeonly"`
	FogTable   hwio.Mem   `hwio:"bank=1,offset=0x60,size=0x20,writeonly"`

	// Registeres shared with e2d
	dispcnt *uint32
	bg0cnt  *uint16
	bg0xofs *uint16

	// Current viewport (last received viewport command)
	viewport Primitive_SetViewport

	pool sync.Pool

	// Current vram/pram (being drawn)
	cur buffer3d

	// Next vram/pram (being accumulated for next frame)
	next   buffer3d
	nextCh chan buffer3d

	// Texture/palette VRAM
	texVram VramTextureBank
	palVram VramTexturePaletteBank

	// Cache for decompressed textures. This currently handles
	// Tex4x4 format as it's too hard to polyfill directly from
	// the compressed format.
	texCache texCache

	backbuf [256 * 192 * 4]uint8
	backY   int32

	framecnt int
}

func NewHwEngine3d() *HwEngine3d {
	e3d := new(HwEngine3d)
	hwio.MustInitRegs(e3d)

	e3d.pool.New = func() interface{} {
		return buffer3d{
			Vram: make([]Vertex, 0, 8192),
			Pram: make([]Polygon, 0, 8192),
		}
	}
	e3d.next = e3d.pool.Get().(buffer3d)
	e3d.nextCh = make(chan buffer3d, 1)

	return e3d
}

func (e3d *HwEngine3d) SetBgRegs(dispcnt *uint32, bg0cnt, bg0xofs *uint16) {
	e3d.dispcnt = dispcnt
	e3d.bg0cnt = bg0cnt
	e3d.bg0xofs = bg0xofs
}

func (vtx *Vertex) calcClippingFlags() {
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
	if vtx.cz.V < -vtx.cw.V {
		vtx.flags |= RVFClipNear
	}
	if vtx.cz.V > vtx.cw.V {
		vtx.flags |= RVFClipFar
	}
}

func (e3d *HwEngine3d) CmdViewport(cmd Primitive_SetViewport) {
	e3d.viewport = cmd
}

func (e3d *HwEngine3d) CmdVertex(cmd Primitive_Vertex) {
	r, g, b := int32(cmd.C[0]), int32(cmd.C[1]), int32(cmd.C[2])

	// Expand to 0-63 range, which is that used internally by
	// rasterizer
	r = r<<1 + (r+31)>>5
	g = g<<1 + (g+31)>>5
	b = b<<1 + (b+31)>>5

	vtx := Vertex{
		cx: cmd.X,
		cy: cmd.Y,
		cz: cmd.Z,
		cw: cmd.W,
		s:  cmd.S.ToF32(),
		t:  cmd.T.ToF32(),
		r:  fixed.NewF12(r),
		g:  fixed.NewF12(g),
		b:  fixed.NewF12(b),
	}
	vtx.calcClippingFlags()
	e3d.next.Vram = append(e3d.next.Vram, vtx)
}

func (e3d *HwEngine3d) CmdPolygon(cmd Primitive_Polygon) {

	flags := PolygonFlags(cmd.Attr)

	var vinbuf [16]*Vertex
	vtxs := vinbuf[:0]

	count := 3
	if flags&PFQuad != 0 {
		count = 4
		flags &^= PFQuad
	}

	clipany := RenderVertexFlags(0)
	clipall := RVFClipMask
	for i := 0; i < count; i++ {
		if cmd.Vtx[i] >= len(e3d.next.Vram) || cmd.Vtx[i] < 0 {
			mod3d.FatalZ("wrong polygon index").Int("index", cmd.Vtx[i]).Int("numvtx", len(e3d.next.Vram)).End()
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

	// Do clipping
	if clipany != 0 {
		vtxs = e3d.polyClip(vtxs)
		if vtxs == nil {
			return
		}
	}

	// Transform all vertices (that weren't transformed already)
	for _, vtx := range vtxs {
		e3d.vtxTransform(vtx)
	}

	// Split the clipped polygon into triangles
	for i := 1; i < len(vtxs)-1; i++ {
		trivtxs := [3]*Vertex{vtxs[0], vtxs[i], vtxs[i+1]}

		// Do backface culling and ignore polygon if it should
		// not be drawn.
		// Notice that we can'd to culling on quads as they could be concave
		d0x := trivtxs[0].x.SubFixed(trivtxs[1].x)
		d0y := trivtxs[0].y.SubFixed(trivtxs[1].y)
		d1x := trivtxs[2].x.SubFixed(trivtxs[1].x)
		d1y := trivtxs[2].y.SubFixed(trivtxs[1].y)
		if int64(d0x.V)*int64(d1y.V) <= int64(d1x.V)*int64(d0y.V) {
			// Facing the back: see if we must render the back
			if flags&PFRenderBack == 0 {
				continue
			}
		} else {
			// Facing the front: see if we must render the front
			if flags&PFRenderFront == 0 {
				continue
			}
		}

		poly := Polygon{
			flags: flags,
			tex:   cmd.Tex,
			vtx:   trivtxs,
		}
		e3d.next.Pram = append(e3d.next.Pram, poly)
	}
}

func (v0 *Vertex) Lerp(v1 *Vertex, ratio fixed.F12, vout *Vertex) {
	vout.cx = v0.cx.Lerp(v1.cx, ratio)
	vout.cy = v0.cy.Lerp(v1.cy, ratio)
	vout.cz = v0.cz.Lerp(v1.cz, ratio)
	vout.cw = v0.cw.Lerp(v1.cw, ratio)
	vout.s = v0.s.Lerp(v1.s, ratio.ToF32())
	vout.t = v0.t.Lerp(v1.t, ratio.ToF32())
	vout.r = v0.r.Lerp(v1.r, ratio)
	vout.g = v0.g.Lerp(v1.g, ratio)
	vout.b = v0.b.Lerp(v1.b, ratio)
}

var clipFormulas = [...]struct {
	Plane         RenderVertexFlags
	PlaneCoord    func(*Vertex) fixed.F12
	PlaneSetCoord func(*Vertex, fixed.F12)
}{
	{
		RVFClipFar,
		func(v *Vertex) fixed.F12 { return v.cz },
		func(v *Vertex, f fixed.F12) { v.cz = f },
	},
	{
		RVFClipNear,
		func(v *Vertex) fixed.F12 { return v.cz.Neg() },
		func(v *Vertex, f fixed.F12) { v.cz = f.Neg() },
	},
	{
		RVFClipTop,
		func(v *Vertex) fixed.F12 { return v.cy },
		func(v *Vertex, f fixed.F12) { v.cy = f },
	},
	{
		RVFClipBottom,
		func(v *Vertex) fixed.F12 { return v.cy.Neg() },
		func(v *Vertex, f fixed.F12) { v.cy = f.Neg() },
	},
	{
		RVFClipLeft,
		func(v *Vertex) fixed.F12 { return v.cx.Neg() },
		func(v *Vertex, f fixed.F12) { v.cx = f.Neg() },
	},
	{
		RVFClipRight,
		func(v *Vertex) fixed.F12 { return v.cx },
		func(v *Vertex, f fixed.F12) { v.cx = f },
	},
}

func (e3d *HwEngine3d) polyClip(poly []*Vertex) []*Vertex {
	var clipbuf [16]*Vertex
	clipped := clipbuf[:0]

	for _, clipInfo := range clipFormulas {
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

				e3d.next.ClipVram = append(e3d.next.ClipVram, Vertex{})
				vout := &e3d.next.ClipVram[len(e3d.next.ClipVram)-1]

				v0.Lerp(v, ratio, vout)
				clipInfo.PlaneSetCoord(vout, vout.cw)
				vout.calcClippingFlags()

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

		// FIXME: check if it's really required
		if len(clipped) == 0 {
			return nil
		}

		poly = append(poly[:0], clipped...)
		clipped = clipbuf[:0]
	}

	// Fix rounding errors; all coordinates should now be
	// in the range [-w, w]
	for _, v := range poly {
		v.cx = v.cx.Clamp(v.cw.Neg(), v.cw)
		v.cy = v.cy.Clamp(v.cw.Neg(), v.cw)
		v.cz = v.cz.Clamp(v.cw.Neg(), v.cw)
	}

	return poly
}

func (e3d *HwEngine3d) vtxTransform(vtx *Vertex) {
	if vtx.flags&RVFTransformed != 0 {
		return
	}

	viewwidth := fixed.NewF12(int32(e3d.viewport.VX1 - e3d.viewport.VX0))
	viewheight := fixed.NewF12(int32(e3d.viewport.VY1 - e3d.viewport.VY0))

	if vtx.cw.V == 0 {
		vtx.cw.V = 1
	}

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

	// Clamp screen coord. This is only required because clipping in clip-space
	// cannot be accurate with fixed point coordinates (at least not with 12 bit),
	// and thus it can generate coordinates that are slightly out
	vtx.x = vtx.x.Clamp(fixed.NewF12(int32(e3d.viewport.VX0)), fixed.NewF12(int32(e3d.viewport.VX1)))
	vtx.y = vtx.y.Clamp(fixed.NewF12(int32(e3d.viewport.VY0)), fixed.NewF12(int32(e3d.viewport.VY1)))

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
		var dxl0, dxl1, dxr0, dxr1 fixed.F32
		var ddl0, ddl1, ddr0, ddr1 fixed.F32
		var dsl0, dsl1, dsr0, dsr1 fixed.F32
		var dtl0, dtl1, dtr0, dtr1 fixed.F32
		var drl0, drl1, drr0, drr1 fixed.F32
		var dgl0, dgl1, dgr0, dgr1 fixed.F32
		var dbl0, dbl1, dbr0, dbr1 fixed.F32

		dxl0 = v1.x.SubFixed(v0.x).ToF32()
		ddl0 = v1.d.SubFixed(v0.d)
		dsl0 = v1.s.SubFixed(v0.s)
		dtl0 = v1.t.SubFixed(v0.t)
		drl0 = v1.r.SubFixed(v0.r).ToF32()
		dgl0 = v1.g.SubFixed(v0.g).ToF32()
		dbl0 = v1.b.SubFixed(v0.b).ToF32()

		dxl1 = v2.x.SubFixed(v1.x).ToF32()
		ddl1 = v2.d.SubFixed(v1.d)
		dsl1 = v2.s.SubFixed(v1.s)
		dtl1 = v2.t.SubFixed(v1.t)
		drl1 = v2.r.SubFixed(v1.r).ToF32()
		dgl1 = v2.g.SubFixed(v1.g).ToF32()
		dbl1 = v2.b.SubFixed(v1.b).ToF32()

		if hy1 > 0 {
			dxl0 = dxl0.Div(hy1)
			ddl0 = ddl0.Div(hy1)
			dsl0 = dsl0.Div(hy1)
			dtl0 = dtl0.Div(hy1)
			drl0 = drl0.Div(hy1)
			dgl0 = dgl0.Div(hy1)
			dbl0 = dbl0.Div(hy1)
		}
		if hy2 > 0 {
			dxl1 = dxl1.Div(hy2)
			ddl1 = ddl1.Div(hy2)
			dsl1 = dsl1.Div(hy2)
			dtl1 = dtl1.Div(hy2)
			drl1 = drl1.Div(hy2)
			dgl1 = dgl1.Div(hy2)
			dbl1 = dbl1.Div(hy2)
		}
		if hy1+hy2 > 0 {
			dxr0 = v2.x.SubFixed(v0.x).ToF32().Div(hy1 + hy2)
			ddr0 = v2.d.SubFixed(v0.d).Div(hy1 + hy2)
			dsr0 = v2.s.SubFixed(v0.s).Div(hy1 + hy2)
			dtr0 = v2.t.SubFixed(v0.t).Div(hy1 + hy2)
			drr0 = v2.r.SubFixed(v0.r).ToF32().Div(hy1 + hy2)
			dgr0 = v2.g.SubFixed(v0.g).ToF32().Div(hy1 + hy2)
			dbr0 = v2.b.SubFixed(v0.b).ToF32().Div(hy1 + hy2)

			dxr1 = dxr0
			ddr1 = ddr0
			dsr1 = dsr0
			dtr1 = dtr0
			drr1 = drr0
			dgr1 = dgr0
			dbr1 = dbr0
		}

		// Now create interpolator instances
		poly.left[LerpX] = newLerp(v0.x.ToF32(), dxl0, dxl1)
		poly.right[LerpX] = newLerp(v0.x.ToF32(), dxr0, dxr1)

		poly.left[LerpD] = newLerp(v0.d, ddl0, ddl1)
		poly.right[LerpD] = newLerp(v0.d, ddr0, ddr1)

		poly.left[LerpS] = newLerp(v0.s, dsl0, dsl1)
		poly.right[LerpS] = newLerp(v0.s, dsr0, dsr1)

		poly.left[LerpT] = newLerp(v0.t, dtl0, dtl1)
		poly.right[LerpT] = newLerp(v0.t, dtr0, dtr1)

		poly.left[LerpR] = newLerp(v0.r.ToF32(), drl0, drl1)
		poly.right[LerpR] = newLerp(v0.r.ToF32(), drr0, drr1)

		poly.left[LerpG] = newLerp(v0.g.ToF32(), dgl0, dgl1)
		poly.right[LerpG] = newLerp(v0.g.ToF32(), dgr0, dgr1)

		poly.left[LerpB] = newLerp(v0.b.ToF32(), dbl0, dbl1)
		poly.right[LerpB] = newLerp(v0.b.ToF32(), dbr0, dbr1)

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
	}
}

func (e3d *HwEngine3d) sortPolys(alphaYSort bool) {
	// Do a stable sort, so that we keep the existing order for all
	// solid polygons (alpha=0x1F). This should be consistent with the order
	// the NDS renderes the display list.
	sort.SliceStable(e3d.next.Pram, func(i, j int) bool {
		pi := &e3d.next.Pram[i]
		pj := &e3d.next.Pram[j]

		// First: all solid polygons before alpha polygons
		isolid := !pi.UseAlpha()
		jsolid := !pj.UseAlpha()
		if isolid && !jsolid {
			return true
		} else if jsolid && !isolid {
			return false
		}
		// Polygons have the same translucent / solid status

		// Second: sort by "y" solid polygons (or all polygons if
		// alphaYSort is true)
		if isolid || alphaYSort {
			// vtx[0] is always the one with lowest y (after preparePolys())
			iy := pi.vtx[0].y.V
			jy := pj.vtx[0].y.V

			if iy < jy {
				return true
			} else if jy < iy {
				return false
			}
		}
		// Now polygons have the same y
		// FIXME: sort left-to-right? For now, ignore.

		// Polygons have the same sorting properties.
		// Return false so that stable sorting keeps them in the right order
		return false
	})
}

func (e3d *HwEngine3d) polysSetDepth(wbuffering bool) {
	for idx := range e3d.next.Pram {
		poly := &e3d.next.Pram[idx]

		for i := 0; i < 3; i++ {
			v := poly.vtx[i]
			if v.flags&RVFDepth != 0 {
				continue
			}

			var d fixed.F12
			if wbuffering {
				// Use W as depth value
				d = v.cw

			} else {
				// Use Z as depth value: first, do the perspective divide and
				// scale it into the range -0x800 / +0x800
				d = v.cz.MulDivFixed(fixed.NewF12(0x800), v.cw)

				// Now move it "almost" to the 0x0 - 0x1000 range.
				// We actually want to avoid a pure 0x0 because it would generate
				// divisions by zero, so we get close to it but not exactly on it
				d.V += 0x7FFEFF
			}

			// Invert depth. We switch from 20.12 to 10.22,
			// attempting to preserve as much precision as possible
			v.d = d.ToF32().Inv()

			// Divide texture coordinates by depth, to prepare
			// for perspective correction
			v.s = v.s.MulFixed(v.d)
			v.t = v.t.MulFixed(v.d)

			// Remember that we've already processed this vertex
			v.flags |= RVFDepth
		}
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
	mod3d.InfoZ("end scene").End()
}

func (e3d *HwEngine3d) CmdSwapBuffers(cmd Primitive_SwapBuffers) {
	// The next frame primitives are complete; we can now do full-frame processing
	// in preparation for drawing next frame

	// Turn on wbuffering instead of zbuffering, if requested
	e3d.polysSetDepth(cmd.WBuffering)

	// Computer interpolators/slopes for all polygons
	e3d.preparePolys()

	// Sort polygons from back to front.
	// NOTE: we need to do this AFTER preparePolys(), becase we
	// need to sort by Y coordinate, and preparePolys() rotate the
	// vertices so that vtx[0] is always the topmost vertex
	e3d.sortPolys(cmd.AlphaYSort)

	// Debug dump of scene
	// e3d.dumpNextScene()

	e3d.framecnt++

	// Send the next buffer to the main rendering thread.
	select {
	case e3d.nextCh <- e3d.next:
	default:
		panic("two scenes queued")
	}

	// Get a new buffer from the pool, ready for next frame
	e3d.next = e3d.pool.Get().(buffer3d)
}

func (e3d *HwEngine3d) drawScene() {
	texMappingEnabled := e3d.Disp3dCnt.Value&(1<<0) != 0
	highlightEnabled := e3d.Disp3dCnt.Value&(1<<1) != 0
	alphaBlendingEnabled := e3d.Disp3dCnt.Value&(1<<3) != 0

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
		y0 := v0.y.TruncInt32()
		y1 := v2.y.TruncInt32()
		if poly.UseAlpha() {
			y1--
		}
		for j := y0; j <= y1; j++ {
			if polyPerLine[j] == nil {
				polyPerLine[j] = make([]uint16, 0, 128)
			}
			polyPerLine[j] = append(polyPerLine[j], uint16(idx))
		}

		// Setup the correct polyfiller for this
		fcfg := fillerconfig.FillerConfig{
			TexFormat: uint(poly.tex.Format),
			ColorMode: uint(poly.flags.ColorMode()),
		}
		if poly.tex.Flags&TexSRepeat == 0 || poly.tex.Flags&TexTRepeat == 0 {
			// If either S or T uses clamping, we need to use the full
			// calculation of texture coordinates.
			fcfg.TexCoords = fillerconfig.TexCoordsFull
		} else if poly.tex.Flags&TexSFlip != 0 || poly.tex.Flags&TexTFlip != 0 {
			// If either S or T uses flipping, we need to use the repeatorflip
			// calculation.
			fcfg.TexCoords = fillerconfig.TexCoordsRepeatOrFlip
		} else {
			// If both S and T uses repeat, we can use the repeatonly fast-path
			if poly.tex.Flags != TexSRepeat|TexTRepeat {
				panic("invalid tex flags")
			}
			fcfg.TexCoords = fillerconfig.TexCoordsRepeatOnly
		}
		// TexCoordsRepeatOnly is the fastest path. It can be selected
		// only if both S/T are configured in repeat mode, but we can also
		// force it if all vertices uses S/T within the texture size;
		// in that case, the repeat mode is basically ignored anyway, so
		// we can simply select the fastest path.
		if fcfg.TexCoords != fillerconfig.TexCoordsRepeatOnly {
			wrap := false
			for _, v := range poly.vtx {
				if v.s.V < 0 || v.t.V < 0 ||
					v.s.DivFixed(v.d).TruncInt32() >= int32(poly.tex.Width) ||
					v.t.DivFixed(v.d).TruncInt32() >= int32(poly.tex.Height) {
					wrap = true
					break
				}
			}
			if !wrap {
				fcfg.TexCoords = fillerconfig.TexCoordsRepeatOnly
			}
		}
		if poly.tex.ColorKey {
			fcfg.ColorKey = 1
		}
		if !texMappingEnabled {
			fcfg.TexFormat = 0
		}

		if alphaBlendingEnabled {
			fcfg.FillMode = fillerconfig.FillModeAlpha
		} else {
			fcfg.FillMode = fillerconfig.FillModeSolid
		}

		if fcfg.ColorMode == fillerconfig.ColorModeToon && highlightEnabled {
			fcfg.ColorMode = fillerconfig.ColorModeHighlight
		}

		poly.filler = polygonFillerTable[fcfg.Key()]
	}

	// FIXME: move elsewhere. Theoretically, 3d rendering begins
	// 19 lines before the first screen line, so it would be possible
	// to begin processing textures somewhere in the middle of vblank.
	// To be 100% sure, we can't do that when we receieve SwapBuffers
	// (that is, in the middle of previous frame) as the texture data
	// could not be ready.
	e3d.texCache.Update(e3d.cur.Pram, e3d)

	for y := 0; y < 192; y++ {
		if e3d.Disp3dCnt.Value&(1<<14) != 0 {
			panic("bitmap")
		}

		line := gfx.NewLine(e3d.backbuf[4*256*y:])

		clearColor := (e3d.ClearColor.Value & 0x7FFF) | 0x80000000
		clearAlpha := uint8(e3d.ClearColor.Value>>16) & 0x1F
		clearDepth := uint32(e3d.ClearDepth.Value)
		clearDepth = (clearDepth * 0x200) + 0x1FF // gbatek is wrong

		var abuf [256]byte
		var zbuf [256 * 4]byte
		zbuffer := gfx.NewLine(zbuf[:])
		abuffer := gfx.NewLine(abuf[:])
		for i := 0; i < 256; i++ {
			line.Set32(i, clearColor)
			abuffer.Set8(i, clearAlpha)
			zbuffer.Set32(i, clearDepth)
		}

		// Draw polygons that are visibile in this line
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

		// Now mark pixels with alpha 0 as fully transparent,
		// and embed 5-bit alpha in pixel in other cases.
		// This will be used for 3d/2d transparency
		for i := 0; i < 256; i++ {
			alpha := abuffer.Get8(i)
			if alpha == 0 {
				line.Set32(i, 0)
			} else {
				line.Set32(i, line.Get32(i)|uint32(alpha)<<16|1<<24)
			}
		}

		atomic.StoreInt32(&e3d.backY, int32(y))
	}
}

func (e3d *HwEngine3d) Draw3D(lidx int) func(gfx.Line) {
	y := int32(0)

	keystate := hw.GetKeyboardState()

	return func(out gfx.Line) {
		xofs := int(*e3d.bg0xofs & 511)
		pri := uint32(*e3d.bg0cnt&3) << 29

		// Wait until the 3D drawing goroutine has completed this line
		for atomic.LoadInt32(&e3d.backY) < y {
			time.Sleep(10 * time.Microsecond)
		}

		// Check if layer 0 is enabled, otherwise ignore
		if *e3d.dispcnt&(1<<8) == 0 {
			y++
			return
		}

		if keystate[hw.SCANCODE_5] != 0 {
			y++
			return
		}

		// Copy the line into the output buffer, applying horizontal offset
		line := gfx.NewLine(e3d.backbuf[y*4*256:])
		for i := 0; i < 256; i++ {
			x := (i + xofs) & 511
			if x < 256 {
				pix := line.Get32(x)
				// Ignore pixel that are fully transparent (= not drawn)
				// Notice that pixels that ended up with alpha=0 have been cleared
				// at the end of drawScene().
				if pix != 0 {
					out.Set32(i, pix|pri)
				}
			}

		}
		y++
	}
}

func (e3d *HwEngine3d) SetVram(tex VramTextureBank, pal VramTexturePaletteBank) {
	e3d.texVram = tex
	e3d.palVram = pal
}

func (e3d *HwEngine3d) BeginFrame() {
	e3d.backY = -1
	go e3d.drawScene()
}

func (e3d *HwEngine3d) EndFrame() {
	// CAVEAT: on the first frame, EndFrame() is called
	// without BeginFrame()! See emulator.go:hsync()

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
