package main

import (
	"ndsemu/emu"
	"ndsemu/emu/gfx"
	log "ndsemu/emu/logger"
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
}

// New polygon to be pushed in Polygon RAM
type E3DCmd_Polygon struct {
	vtx  [4]int // indices of vertices in Vertex RAM
	attr uint32 // misc flags
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

	flags RenderVertexFlags

	// Screen coordinates
	sx, sy, sz emu.Fixed12
}

// NOTE: these flags match the polygon attribute word defined in the
// geometry coprocessor.
type RenderPolygonFlags uint32

const (
	RPFQuad RenderPolygonFlags = 1 << 31
)

type RenderPolygon struct {
	vtx   [4]int
	flags RenderPolygonFlags
}

type HwEngine3d struct {
	// Current viewport (last received viewport command)
	viewport E3DCmd_SetViewport
	// plinecnt [192]cnt
	// plines   [1024][192]int

	vertexRams [4096][2]RenderVertex
	polyRams   [4096][2]RenderPolygon

	// Current vram/pram (being drawn)
	curVram []RenderVertex
	curPram []RenderPolygon

	// Next vram/pram (being accumulated for next frame)
	nextVram []RenderVertex
	nextPram []RenderPolygon

	framecnt  int
	frameLock sync.Mutex

	// Channel to receive new commands
	CmdCh chan interface{}
}

func NewHwEngine3d() *HwEngine3d {
	e3d := new(HwEngine3d)
	e3d.nextVram = e3d.vertexRams[e3d.framecnt&1][:0]
	e3d.nextPram = e3d.polyRams[e3d.framecnt&1][:0]
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

	if count == 4 {
		// Since we're done with clipping, split quad in two
		// triangles, to make the renderer only care about
		// triangles.
		p1, p2 := poly, poly

		p1.flags &^= RPFQuad
		p2.flags &^= RPFQuad
		copy(p2.vtx[0:3], p2.vtx[1:4])

		e3d.nextPram = append(e3d.nextPram, p1, p2)
	} else {
		e3d.nextPram = append(e3d.nextPram, poly)
	}
}

func (e3d *HwEngine3d) vtxTransform(vtx *RenderVertex) {
	if vtx.flags&RVFTransformed != 0 {
		return
	}

	viewwidth := emu.NewFixed12(int32(e3d.viewport.vx1 - e3d.viewport.vx0 + 1))
	viewheight := emu.NewFixed12(int32(e3d.viewport.vy1 - e3d.viewport.vy0 + 1))
	// Compute viewsize / (2*v.w) in two steps, to avoid overflows
	// (viewwidth could be 256<<12, which would overflow when further
	// shifted in preparation for division)
	dx := viewwidth.Div(2).DivFixed(vtx.cw)
	dy := viewheight.Div(2).DivFixed(vtx.cw)

	// sx = (v.x + v.w) * viewwidth / (2*v.w) + viewx0
	// sy = (v.y + v.w) * viewheight / (2*v.w) + viewy0
	vtx.sx = vtx.cx.AddFixed(vtx.cw).MulFixed(dx).Add(int32(e3d.viewport.vx0))
	vtx.sy = vtx.cy.AddFixed(vtx.cw).MulFixed(dy).Add(int32(e3d.viewport.vy0))
	vtx.sz = vtx.cz.AddFixed(vtx.cw).Div(2).DivFixed(vtx.cw)

	vtx.flags |= RVFTransformed
}

func (e3d *HwEngine3d) dumpNextScene() {
	mod3d.Infof("begin scene")
	for _, poly := range e3d.nextPram {
		v0 := &e3d.nextVram[poly.vtx[0]]
		v1 := &e3d.nextVram[poly.vtx[1]]
		v2 := &e3d.nextVram[poly.vtx[2]]
		mod3d.Infof("tri SV: (%.2f,%.2f)-(%.2f,%.2f)-(%.2f,%.2f)",
			v0.sx.ToFloat64(), v0.sy.ToFloat64(),
			v1.sx.ToFloat64(), v1.sy.ToFloat64(),
			v2.sx.ToFloat64(), v2.sy.ToFloat64())
	}
	mod3d.Infof("end scene")
}

func (e3d *HwEngine3d) cmdSwapBuffers() {
	// The next frame primitives are complete; we can now do full-frame processing
	// in preparation for drawing next frame
	e3d.dumpNextScene()

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

}

func (e3d *HwEngine3d) BeginFrame() {
	// Acquire the frame lock, we will begin drawing now
	e3d.frameLock.Lock()
}

func (e3d *HwEngine3d) EndFrame() {
	// Release the frame lock, drawing is finished
	e3d.frameLock.Unlock()
}
