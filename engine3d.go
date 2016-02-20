package main

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"os"
	"sync"
)

var mod3d = log.NewModule("e3d")

// Swap buffers (marker of end-of-frame, with double-buffering)
type E3DCmd_SwapBuffers struct {
}

// New viewport, in pixel coordinates (0-255 / 0-191)
type E3DCmd_SetViewport struct {
	vx0, vy0, vx1, vy1 int
}

// New vertex to be pushed in Vertex RAM, with coordinates in
// clip space (after model-view-proj)
type E3DCmd_Vertex struct {
	x, y, z, w emu.Fixed12
	s, t       emu.Fixed12
}

// New polygon to be pushed in Polygon RAM
type E3DCmd_Polygon struct {
	vtx  [4]int        // indices of vertices in Vertex RAM
	attr uint32        // misc flags
	tex  RenderTexture // texture for this polygon
}

type RenderVertexFlags uint32

const (
	RVFClipLeft RenderVertexFlags = 1 << iota
	RVFClipRight
	RVFClipTop
	RVFClipBottom
	RVFClipNear
	RVFClipFar
	RVFTransformed // vertex has been already transformed to screen space

	RVFClipAnything = (RVFClipLeft | RVFClipRight | RVFClipTop | RVFClipBottom | RVFClipNear | RVFClipFar)
)

type RenderVertex struct {
	// Coordinates in clip-space
	cx, cy, cz, cw emu.Fixed12

	// Screen coordinates (fractional part is always zero)
	x, y, z emu.Fixed12

	// Texture coordinates
	s, t emu.Fixed12

	// Misc flags
	flags RenderVertexFlags
}

type lerp struct {
	cur   emu.Fixed12
	delta [2]emu.Fixed12
	start emu.Fixed12
}

func newLerp(start emu.Fixed12, d0 emu.Fixed12, d1 emu.Fixed12) lerp {
	return lerp{start: start, delta: [2]emu.Fixed12{d0, d1}}
}

func (l *lerp) Reset() {
	l.cur = l.start
}

func (l *lerp) Cur() emu.Fixed12 {
	return l.cur
}

func (l *lerp) Next(didx int) {
	l.cur = l.cur.AddFixed(l.delta[didx])
}

func (l lerp) String() string {
	return fmt.Sprintf("lerp(%v (%v,%v) [%v])", l.cur, l.delta[0], l.delta[1], l.start)
}

// NOTE: these flags match the polygon attribute word defined in the
// geometry coprocessor.
type RenderPolygonFlags uint32

const (
	RPFQuad RenderPolygonFlags = 1 << 31
)

const (
	LerpX = iota // coordinate on screen (X)
	LerpZ        // depth on screen (Z or W)
	LerpT        // texture X coordinate (T)
	LerpS        // texture Y coordinate (S)
	NumLerps
)

type RenderPolygon struct {
	vtx   [4]int
	flags RenderPolygonFlags
	tex   RenderTexture

	// y coordinate of middle vertex
	hy int32

	// linear interpolators for left and right edge of the polygon
	left  [NumLerps]lerp
	right [NumLerps]lerp

	// texture pointer
	texptr []byte
}

type RTexFormat int
type RTexFlags int

const (
	RTexNone RTexFormat = iota
	RTexA3I5
	RTex4
	RTex16
	RTex256
	RTex4x4
	RTexA5I3
	RTexDirect
)

const (
	RTexSFlip RTexFlags = 1 << iota
	RTexTFlip
	RTexSRepeat
	RTexTRepeat
)

type RenderTexture struct {
	VramTexOffset uint32
	VramPalOffset uint32
	SMask, TMask  uint32
	PitchShift    uint
	Transparency  bool
	Format        RTexFormat
	Flags         RTexFlags
}

type HwEngine3d struct {
	Disp3dCnt hwio.Reg32 `hwio:"offset=0,rwmask=0x7FFF"`

	// Channel to receive new primitives (sent by GxFifo)
	CmdCh chan interface{}

	// Current viewport (last received viewport command)
	viewport E3DCmd_SetViewport

	vertexRams [2][4096]RenderVertex
	polyRams   [2][4096]RenderPolygon

	// Current vram/pram (being drawn)
	curVram []RenderVertex
	curPram []RenderPolygon

	// Next vram/pram (being accumulated for next frame)
	nextVram []RenderVertex
	nextPram []RenderPolygon

	framecnt  int
	frameLock sync.Mutex
}

func NewHwEngine3d() *HwEngine3d {
	e3d := new(HwEngine3d)
	hwio.MustInitRegs(e3d)
	e3d.nextVram = e3d.vertexRams[0][:0]
	e3d.nextPram = e3d.polyRams[0][:0]
	e3d.CmdCh = make(chan interface{}, 1024)
	go e3d.recvCmd()
	return e3d
}

func (e3d *HwEngine3d) recvCmd() {
	for {
		cmdi := <-e3d.CmdCh
		switch cmd := cmdi.(type) {
		case E3DCmd_SwapBuffers:
			e3d.cmdSwapBuffers()
		case E3DCmd_SetViewport:
			e3d.viewport = cmd
		case E3DCmd_Polygon:
			e3d.cmdPolygon(cmd)
		case E3DCmd_Vertex:
			e3d.cmdVertex(cmd)
		default:
			panic("invalid command received in HwEnginge3D")
		}
	}
}

func (e3d *HwEngine3d) cmdVertex(cmd E3DCmd_Vertex) {
	vtx := RenderVertex{
		cx: cmd.x,
		cy: cmd.y,
		cz: cmd.z,
		cw: cmd.w,
		s:  cmd.s,
		t:  cmd.t,
	}

	// Compute clipping flags (once per vertex)
	if vtx.cx.V < -vtx.cw.V {
		vtx.flags |= RVFClipLeft
	}
	if vtx.cx.V > vtx.cw.V {
		vtx.flags |= RVFClipRight
	}
	if vtx.cy.V < -vtx.cw.V {
		vtx.flags |= RVFClipTop
	}
	if vtx.cy.V > vtx.cw.V {
		vtx.flags |= RVFClipBottom
	}
	// if vtx.cz.V < 0 {
	// 	vtx.flags |= RVFClipNear
	// }
	// if vtx.cz.V > vtx.cw.V {
	// 	vtx.flags |= RVFClipFar
	// }

	// If w==0, we just flag the vertex as fully outside of the screen
	// FIXME: properly handle invalid inputs
	if vtx.cw.V == 0 {
		vtx.flags |= RVFClipAnything
	}

	e3d.nextVram = append(e3d.nextVram, vtx)
}

func (e3d *HwEngine3d) cmdPolygon(cmd E3DCmd_Polygon) {
	poly := RenderPolygon{
		vtx:   cmd.vtx,
		flags: RenderPolygonFlags(cmd.attr),
		tex:   cmd.tex,
	}

	// FIXME: for now, skip all polygons outside the screen
	count := 3
	if poly.flags&RPFQuad != 0 {
		count = 4
	}
	clipping := false
	for i := 0; i < count; i++ {
		if poly.vtx[i] >= len(e3d.nextVram) || poly.vtx[i] < 0 {
			mod3d.Fatalf("wrong polygon index: %d (num vtx: %d)", poly.vtx[i], len(e3d.nextVram))
		}
		vtx := e3d.nextVram[poly.vtx[i]]
		if vtx.flags&RVFClipAnything != 0 {
			clipping = true
			break
		}
	}

	if clipping {
		// FIXME: implement clipping
		return
	}

	// Transform all vertices (that weren't transformed already)
	for i := 0; i < count; i++ {
		e3d.vtxTransform(&e3d.nextVram[poly.vtx[i]])
	}

	// Do backface culling
	v0, v1, v2 := &e3d.nextVram[poly.vtx[0]], &e3d.nextVram[poly.vtx[1]], &e3d.nextVram[poly.vtx[2]]
	d0x := v0.x.SubFixed(v1.x)
	d0y := v0.y.SubFixed(v1.y)
	d1x := v2.x.SubFixed(v1.x)
	d1y := v2.y.SubFixed(v1.y)
	if int64(d0x.V)*int64(d1y.V) <= int64(d1x.V)*int64(d0y.V) {
		return
	}

	if count == 4 {
		// Since we're done with clipping, split quad in two
		// triangles, to make the renderer only care about
		// triangles.
		p1, p2 := poly, poly

		p1.flags &^= RPFQuad
		p2.flags &^= RPFQuad
		p2.vtx[1], p2.vtx[2] = p2.vtx[2], p2.vtx[3]

		e3d.nextPram = append(e3d.nextPram, p1, p2)
	} else {
		e3d.nextPram = append(e3d.nextPram, poly)
	}
}

func (e3d *HwEngine3d) vtxTransform(vtx *RenderVertex) {
	if vtx.flags&RVFTransformed != 0 {
		return
	}

	viewwidth := emu.NewFixed12(int32(e3d.viewport.vx1 - e3d.viewport.vx0))
	viewheight := emu.NewFixed12(int32(e3d.viewport.vy1 - e3d.viewport.vy0))
	// Compute viewsize / (2*v.w) in two steps, to avoid overflows
	// (viewwidth could be 256<<12, which would overflow when further
	// shifted in preparation for division)
	dx := viewwidth.Div(2).DivFixed(vtx.cw)
	dy := viewheight.Div(2).DivFixed(vtx.cw)

	mirror := vtx.cw.Mul(2)

	// sx = (v.x + v.w) * viewwidth / (2*v.w) + viewx0
	// sy = (v.y + v.w) * viewheight / (2*v.w) + viewy0
	vtx.x = vtx.cx.AddFixed(vtx.cw).MulFixed(dx).Add(int32(e3d.viewport.vx0)).Round()
	vtx.y = mirror.SubFixed(vtx.cy.AddFixed(vtx.cw)).MulFixed(dy).Add(int32(e3d.viewport.vy0)).Round()
	vtx.z = vtx.cw // vtx.cz.AddFixed(vtx.cw).Div(2).DivFixed(vtx.cw).Round()

	vtx.flags |= RVFTransformed
}

func (e3d *HwEngine3d) preparePolys() {

	for idx := range e3d.nextPram {
		poly := &e3d.nextPram[idx]
		v0, v1, v2 := &e3d.nextVram[poly.vtx[0]], &e3d.nextVram[poly.vtx[1]], &e3d.nextVram[poly.vtx[2]]

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
		var dxl0, dxl1, dxr0, dxr1 emu.Fixed12
		var dzl0, dzl1, dzr0, dzr1 emu.Fixed12
		var dsl0, dsl1, dsr0, dsr1 emu.Fixed12
		var dtl0, dtl1, dtr0, dtr1 emu.Fixed12

		dxl0 = v1.x.SubFixed(v0.x)
		dzl0 = v1.z.SubFixed(v0.z)
		dsl0 = v1.s.SubFixed(v0.s)
		dtl0 = v1.t.SubFixed(v0.t)

		dxl1 = v2.x.SubFixed(v1.x)
		dzl1 = v1.z.SubFixed(v1.z)
		dsl1 = v2.s.SubFixed(v1.s)
		dtl1 = v2.t.SubFixed(v1.t)

		if hy1 > 0 {
			dxl0 = dxl0.Div(hy1)
			dzl0 = dzl0.Div(hy1)
			dsl0 = dsl0.Div(hy1)
			dtl0 = dtl0.Div(hy1)
		}
		if hy2 > 0 {
			dxl1 = dxl1.Div(hy2)
			dzl1 = dzl1.Div(hy2)
			dsl1 = dsl1.Div(hy2)
			dtl1 = dtl1.Div(hy2)
		}
		if hy1+hy2 > 0 {
			dxr0 = v2.x.SubFixed(v0.x).Div(hy1 + hy2)
			dzr0 = v2.z.SubFixed(v0.z).Div(hy1 + hy2)
			dsr0 = v2.s.SubFixed(v0.s).Div(hy1 + hy2)
			dtr0 = v2.t.SubFixed(v0.t).Div(hy1 + hy2)

			dxr1 = dxr0
			dzr1 = dzr0
			dsr1 = dsr0
			dtr1 = dtr0
		}

		// Now create interpolator instances
		poly.left[LerpX] = newLerp(v0.x, dxl0, dxl1)
		poly.right[LerpX] = newLerp(v0.x, dxr0, dxr1)

		poly.left[LerpZ] = newLerp(v0.z, dzl0, dzl1)
		poly.right[LerpZ] = newLerp(v0.z, dzr0, dzr1)

		poly.left[LerpS] = newLerp(v0.s, dsl0, dsl1)
		poly.right[LerpS] = newLerp(v0.s, dsr0, dsr1)

		poly.left[LerpT] = newLerp(v0.t, dtl0, dtl1)
		poly.right[LerpT] = newLerp(v0.t, dtr0, dtr1)

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
					rp.start = rp.start.AddFixed(rp.delta[0])
				}
			} else {
				for idx := range poly.left {
					rp := &poly.left[idx]
					rp.start = rp.start.AddFixed(rp.delta[0])
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
	for idx, poly := range e3d.nextPram {
		v0 := &e3d.nextVram[poly.vtx[0]]
		v1 := &e3d.nextVram[poly.vtx[1]]
		v2 := &e3d.nextVram[poly.vtx[2]]
		fmt.Fprintf(f, "tri %d:\n", idx)
		fmt.Fprintf(f, "    ccoord: (%v,%v,%v,%v)-(%v,%v,%v,%v)-(%v,%v,%v,%v)\n",
			v0.cx, v0.cy, v0.cz, v0.cw,
			v1.cx, v1.cy, v1.cz, v1.cw,
			v2.cx, v2.cy, v2.cz, v2.cw)
		fmt.Fprintf(f, "    scoord: (%v,%v)-(%v,%v)-(%v,%v)\n",
			v0.x.TruncInt32(), v0.y.TruncInt32(),
			v1.x.TruncInt32(), v1.y.TruncInt32(),
			v2.x.TruncInt32(), v2.y.TruncInt32())
		fmt.Fprintf(f, "    tex: (%v,%v)-(%v,%v)-(%v,%v)\n",
			v0.s, v0.t,
			v1.s, v1.t,
			v2.s, v2.t)
		// fmt.Fprintf(f, "    left lerps: %v\n", poly.left)
		// fmt.Fprintf(f, "    right lerps: %v\n", poly.right)
		fmt.Fprintf(f, "    hy: %v\n", poly.hy)
		fmt.Fprintf(f, "    tex: fmt=%d, flips=%v, flipt=%v, reps=%v, rept=%v\n",
			poly.tex.Format, poly.tex.Flags&RTexSFlip != 0, poly.tex.Flags&RTexTFlip != 0,
			poly.tex.Flags&RTexSRepeat != 0, poly.tex.Flags&RTexTRepeat != 0)
	}
	mod3d.Infof("end scene")
}

func (e3d *HwEngine3d) cmdSwapBuffers() {
	// The next frame primitives are complete; we can now do full-frame processing
	// in preparation for drawing next frame
	e3d.preparePolys()
	// e3d.dumpNextScene()

	// Now wait for the current frame to be fully drawn,
	// because we don't want to mess with buffers being drawn
	e3d.frameLock.Lock()
	e3d.framecnt++
	e3d.curVram = e3d.nextVram
	e3d.curPram = e3d.nextPram
	e3d.nextVram = e3d.vertexRams[e3d.framecnt&1][:0]
	e3d.nextPram = e3d.polyRams[e3d.framecnt&1][:0]
	e3d.frameLock.Unlock()
}

func (e3d *HwEngine3d) Draw3D(ctx *gfx.LayerCtx, lidx int, y int) {

	// Initialize rasterizer.
	var polyPerLine [192][]uint16
	for idx := range e3d.curPram {
		poly := &e3d.curPram[idx]

		// Set current segment to the initial one computed in preparePolys
		// This is required because we might need to redraw the exact
		// same 3D scene multiple times, so each time we want to start
		// from the beginning.
		for idx := range poly.left {
			poly.left[idx].Reset()
			poly.right[idx].Reset()
		}

		// Update the per-line polygon list, by adding this polygon's index
		// to the lines in which it is visible.
		v0, v1, v2 := &e3d.curVram[poly.vtx[0]], &e3d.curVram[poly.vtx[1]], &e3d.curVram[poly.vtx[2]]
		if (v0.x == v1.x && v0.y == v1.y) || (v1.x == v2.x && v1.y == v2.y) {
			// FIXME: implement segments
			continue
		}

		for j := v0.y.TruncInt32(); j <= v2.y.TruncInt32(); j++ {
			polyPerLine[j] = append(polyPerLine[j], uint16(idx))
		}
	}

	vramTex := Emu.Hw.Mc.VramTextureBank()
	vramPal := Emu.Hw.Mc.VramTexturePaletteBank()
	texMappingEnabled := e3d.Disp3dCnt.Value&1 != 0
	for {
		line := ctx.NextLine()
		if line.IsNil() {
			return
		}

		var zbuf [256 * 4]byte
		zbuffer := gfx.NewLine(zbuf[:])
		for i := 0; i < 256; i++ {
			zbuffer.Set32(i, 0x7FFFFFFF)
		}

		for _, idx := range polyPerLine[y] {
			poly := &e3d.curPram[idx]

			x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()
			if x0 < 0 || x1 >= 256 || x1 < x0 {
				fmt.Printf("%v,%v\n", e3d.curVram[poly.vtx[0]].x.TruncInt32(), e3d.curVram[poly.vtx[0]].y.TruncInt32())
				fmt.Printf("%v,%v\n", e3d.curVram[poly.vtx[1]].x.TruncInt32(), e3d.curVram[poly.vtx[1]].y.TruncInt32())
				fmt.Printf("%v,%v\n", e3d.curVram[poly.vtx[2]].x.TruncInt32(), e3d.curVram[poly.vtx[2]].y.TruncInt32())
				fmt.Printf("left lerps: %v\n", poly.left)
				fmt.Printf("right lerps: %v\n", poly.right)
				panic("out of bounds")
			}

			texoff := poly.tex.VramTexOffset
			tshift := poly.tex.PitchShift
			palette := vramPal.Palette(int(poly.tex.VramPalOffset))
			nx := x1 - x0
			z0 := poly.left[LerpZ].Cur()
			z1 := poly.right[LerpZ].Cur()
			s0 := poly.left[LerpS].Cur()
			s1 := poly.right[LerpS].Cur()
			t0 := poly.left[LerpT].Cur()
			t1 := poly.right[LerpT].Cur()
			dz := z1.SubFixed(z0)
			ds := s1.SubFixed(s0)
			dt := t1.SubFixed(t0)
			smask := poly.tex.SMask
			tmask := poly.tex.TMask
			traspmask := uint8(0xFF)
			if poly.tex.Transparency {
				traspmask = 0
			}
			if nx > 0 {
				dz = dz.Div(nx)
				ds = ds.Div(nx)
				dt = dt.Div(nx)
			}

			fmt := poly.tex.Format
			if !texMappingEnabled {
				fmt = RTexNone
			}
			switch fmt {
			case RTexNone:
				for x := x0; x < x1; x++ {
					line.Set16(int(x), 0xFFFF)
				}
			case RTex16:
				tshift -= 1 // because 2 pixels per bytes
				for x := x0; x < x1; x++ {
					s, t := uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
					px := vramTex.Get8(texoff + t<<tshift + s/2)
					px = px >> (4 * uint(s&1))
					px &= 0xF
					if px|traspmask != 0 && z0.V < int32(zbuffer.Get32(int(x))) {
						line.Set16(int(x), palette.Lookup(px)|0x8000)
						zbuffer.Set32(int(x), uint32(z0.V))
					}

					z0 = z0.AddFixed(dz)
					s0 = s0.AddFixed(ds)
					t0 = t0.AddFixed(dt)
				}
			case RTex256:
				for x := x0; x < x1; x++ {
					s, t := uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
					px := vramTex.Get8(texoff + t<<tshift + s)
					if px|traspmask != 0 && z0.V < int32(zbuffer.Get32(int(x))) {
						line.Set16(int(x), palette.Lookup(px)|0x8000)
						zbuffer.Set32(int(x), uint32(z0.V))
					}

					z0 = z0.AddFixed(dz)
					s0 = s0.AddFixed(ds)
					t0 = t0.AddFixed(dt)
				}

			case RTexA5I3:
				// FIXME: handle alpha blending modes
				for x := x0; x < x1; x++ {
					s, t := uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
					px := vramTex.Get8(texoff + t<<tshift + s)
					if px|traspmask != 0 && z0.V < int32(zbuffer.Get32(int(x))) {
						line.Set16(int(x), palette.Lookup(px&7)|0x8000)
						zbuffer.Set32(int(x), uint32(z0.V))
					}

					z0 = z0.AddFixed(dz)
					s0 = s0.AddFixed(ds)
					t0 = t0.AddFixed(dt)
				}

			case RTexA3I5:
				// FIXME: handle alpha blending modes
				for x := x0; x < x1; x++ {
					s, t := uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
					px := vramTex.Get8(texoff + t<<tshift + s)
					if px|traspmask != 0 && z0.V < int32(zbuffer.Get32(int(x))) {
						line.Set16(int(x), palette.Lookup(px&0x1f)|0x8000)
						zbuffer.Set32(int(x), uint32(z0.V))
					}

					z0 = z0.AddFixed(dz)
					s0 = s0.AddFixed(ds)
					t0 = t0.AddFixed(dt)
				}

			case RTexDirect:
				tshift += 1 // because texel is 2 bytes
				for x := x0; x < x1; x++ {
					s, t := uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask
					px := vramTex.Get16(texoff + t<<tshift + s<<1)
					if z0.V < int32(zbuffer.Get32(int(x))) {
						line.Set16(int(x), px|0x8000)
						zbuffer.Set32(int(x), uint32(z0.V))
					}

					z0 = z0.AddFixed(dz)
					s0 = s0.AddFixed(ds)
					t0 = t0.AddFixed(dt)
				}

			case 5:
			// case 7:

			default:
				mod3d.Fatal("texformat not implemented:", poly.tex.Format)
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

		y++
	}
}

func (e3d *HwEngine3d) BeginFrame() {
	// Acquire the frame lock, we will begin drawing now
	e3d.frameLock.Lock()
}

func (e3d *HwEngine3d) EndFrame() {
	// Release the frame lock, drawing is finished
	e3d.frameLock.Unlock()
}
