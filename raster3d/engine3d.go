package raster3d

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"ndsemu/raster3d/fillerconfig"
	"os"
	"sync"
)

var mod3d = log.NewModule("e3d")

type HwEngine3d struct {
	Disp3dCnt hwio.Reg32 `hwio:"offset=0,rwmask=0x7FFF"`

	// Channel to receive new primitives (sent by GxFifo)
	CmdCh chan interface{}

	// Current viewport (last received viewport command)
	viewport Primitive_SetViewport

	vertexRams [2][4096]Vertex
	polyRams   [2][4096]Polygon

	// Current vram/pram (being drawn)
	curVram []Vertex
	curPram []Polygon

	// Next vram/pram (being accumulated for next frame)
	nextVram []Vertex
	nextPram []Polygon

	// Texture/palette VRAM
	texVram VramTextureBank
	palVram VramTexturePaletteBank

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

func (e3d *HwEngine3d) cmdVertex(cmd Primitive_Vertex) {
	vtx := Vertex{
		cx: cmd.X,
		cy: cmd.Y,
		cz: cmd.Z,
		cw: cmd.W,
		s:  cmd.S,
		t:  cmd.T,
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

func (e3d *HwEngine3d) cmdPolygon(cmd Primitive_Polygon) {
	poly := Polygon{
		vtx:   cmd.Vtx,
		flags: PolygonFlags(cmd.Attr),
		tex:   cmd.Tex,
	}

	// FIXME: for now, skip all polygons outside the screen
	count := 3
	if poly.flags&PFQuad != 0 {
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
		// Facing the back: see if we must render the back
		if poly.flags&PFRenderBack == 0 {
			return
		}
	} else {
		// Facing the front: see if we must render the front
		if poly.flags&PFRenderFront == 0 {
			return
		}
	}

	if count == 4 {
		// Since we're done with clipping, split quad in two
		// triangles, to make the renderer only care about
		// triangles.
		p1, p2 := poly, poly

		p1.flags &^= PFQuad
		p2.flags &^= PFQuad
		p2.vtx[1], p2.vtx[2] = p2.vtx[2], p2.vtx[3]

		e3d.nextPram = append(e3d.nextPram, p1, p2)
	} else {
		e3d.nextPram = append(e3d.nextPram, poly)
	}
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
	vtx.z = vtx.cz.AddFixed(vtx.cw).Div(2).DivFixed(vtx.cw).Round()

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

	texMappingEnabled := e3d.Disp3dCnt.Value&1 != 0

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

		// If the polygon degrades to a segment, skip it for now (we don't support segments)
		v0, v1, v2 := &e3d.curVram[poly.vtx[0]], &e3d.curVram[poly.vtx[1]], &e3d.curVram[poly.vtx[2]]
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
			ColorKey:  poly.tex.Transparency,
		}
		if !texMappingEnabled {
			fcfg.TexFormat = 0
		}
		poly.filler = polygonFillerTable[fcfg.Key()]
	}

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

			poly.filler(e3d, poly, line, zbuffer)

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

func (e3d *HwEngine3d) SetVram(tex VramTextureBank, pal VramTexturePaletteBank) {
	e3d.texVram = tex
	e3d.palVram = pal
}

func (e3d *HwEngine3d) BeginFrame() {
	// Acquire the frame lock, we will begin drawing now
	e3d.frameLock.Lock()
}

func (e3d *HwEngine3d) EndFrame() {
	// Release the frame lock, drawing is finished
	e3d.frameLock.Unlock()
}
