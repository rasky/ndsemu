package main

import (
	"fmt"
	"ndsemu/emu/fixed"
	"ndsemu/raster3d"
)

type vector [4]fixed.F12
type matrix [4]vector
type color [3]uint8
type fcolor [3]fixed.F12

func newFcolorFrom555(r, g, b uint32) (f fcolor) {
	f[0] = fixed.NewF12(int32(r)).Div(31)
	f[1] = fixed.NewF12(int32(g)).Div(31)
	f[2] = fixed.NewF12(int32(b)).Div(31)
	return
}

func (f fcolor) ToColor() (c color) {
	c[0] = uint8(f[0].Mul(31).TruncInt32())
	c[1] = uint8(f[1].Mul(31).TruncInt32())
	c[2] = uint8(f[2].Mul(31).TruncInt32())
	return
}

func (f fcolor) ToClampedColor() (c color) {
	r := f[0].Mul(31).TruncInt32()
	g := f[1].Mul(31).TruncInt32()
	b := f[2].Mul(31).TruncInt32()
	if r < 0 {
		r = 0
	}
	if g < 0 {
		g = 0
	}
	if b < 0 {
		b = 0
	}
	if r > 31 {
		r = 31
	}
	if g > 31 {
		g = 31
	}
	if b > 31 {
		b = 31
	}
	c[0] = uint8(r)
	c[1] = uint8(g)
	c[2] = uint8(b)
	return
}

func (c color) To32bit() uint32 {
	return uint32(c[0]) | uint32(c[1])<<8 | uint32(c[2])<<16
}

func (v vector) String() string {
	return fmt.Sprintf("v3d(%v,%v,%v,%v)", v[0], v[1], v[2], v[3])
}

func newMatrixIdentity() (m matrix) {
	m[0][0] = fixed.NewF12(1)
	m[1][1] = fixed.NewF12(1)
	m[2][2] = fixed.NewF12(1)
	m[3][3] = fixed.NewF12(1)
	return
}

func newMatrixTrans(x, y, z fixed.F12) matrix {
	m := newMatrixIdentity()
	m[3][0] = x
	m[3][1] = y
	m[3][2] = z
	return m
}

func newMatrixScale(x, y, z fixed.F12) (m matrix) {
	m[0][0] = x
	m[1][1] = y
	m[2][2] = z
	m[3][3] = fixed.NewF12(1)
	return
}

func (mtx *matrix) Row(j int) vector {
	return mtx[j]
}

func (mtx *matrix) Col(i int) (res vector) {
	res[0] = mtx[0][i]
	res[1] = mtx[1][i]
	res[2] = mtx[2][i]
	res[3] = mtx[3][i]
	return res
}

func matMul(a, b matrix) (c matrix) {
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			c[j][i] = a.Row(j).Dot(b.Col(i))
		}
	}
	return
}

func (v vector) Dot(vb vector) (res fixed.F12) {
	val := int64(v[0].V) * int64(vb[0].V)
	val += int64(v[1].V) * int64(vb[1].V)
	val += int64(v[2].V) * int64(vb[2].V)
	val += int64(v[3].V) * int64(vb[3].V)
	val >>= 12
	if val != int64(int32(val)) {
		// Overflow! We use saturation here as it is
		// the most probable solution used by the
		// hardware as well.
		if val > 0 {
			val = 0x7FFFFFFF
		} else {
			val = 0x80000000
		}
	}
	res.V = int32(val)
	return
}

func (v vector) Dot3(vb vector) (res fixed.F12) {
	val := int64(v[0].V) * int64(vb[0].V)
	val += int64(v[1].V) * int64(vb[1].V)
	val += int64(v[2].V) * int64(vb[2].V)
	val >>= 12
	if val != int64(int32(val)) {
		// Overflow! We use saturation here as it is
		// the most probable solution used by the
		// hardware as well.
		if val > 0 {
			val = 0x7FFFFFFF
		} else {
			val = 0x80000000
		}
	}
	res.V = int32(val)
	return
}

func (mtx *matrix) VecMul(vec vector) (res vector) {
	res[0] = vec.Dot(mtx.Col(0))
	res[1] = vec.Dot(mtx.Col(1))
	res[2] = vec.Dot(mtx.Col(2))
	res[3] = vec.Dot(mtx.Col(3))
	return
}

func (mtx *matrix) VecMul3x3(vec vector) (res vector) {
	res[0] = vec.Dot3(mtx.Col(0))
	res[1] = vec.Dot3(mtx.Col(1))
	res[2] = vec.Dot3(mtx.Col(2))
	return
}

const (
	MatDiffuse  = 0
	MatAmbient  = 1
	MatSpecular = 2
	MatEmission = 3

	MtxProjection = 0
	MtxPosition   = 1
	MtxDirection  = 2 // aka "vector", used for light
	MtxTexture    = 3
)

type GeometryEngine struct {
	// Current matrices
	mtxmode int
	mtx     [4]matrix // 0=proj, 1=pos, 2=dir, 3=tex
	clipmtx matrix    // current clip matrix (pos * proj)

	// Matrix stacks
	mtxStackProj     [1]matrix
	mtxStackPos      [32]matrix
	mtxStackDir      [32]matrix
	mtxStackTex      [1]matrix
	mtxStackProjPtr  int
	mtxStackPosPtr   int
	mtxStackTexPtr   int
	mtxStackOverflow bool

	// Viewport
	vx0, vy0 int
	vx1, vy1 int

	// Material and lights
	material [4]fcolor
	lights   [4]struct {
		dir   vector
		half  vector
		color fcolor
	}
	specTable   [128]fixed.F12
	specTableOn bool

	// Textures
	texinfo  raster3d.Texture
	textrans int

	// Polygons and display lists
	polyattr uint32
	displist struct {
		primtype int
		color    color
		polyattr uint32
		s, t     fixed.F12
		s0, t0   fixed.F12
		cnt      int
		lastvtx  vector
	}
	vcnt int

	// Pos/Vec test results
	posTestResult vector
	vecTestResult vector

	// Channel to send commands to 3D rasterizer engine
	// E3dCmdCh chan []interface{}
	// primBuf  []interface{}
	e3d *raster3d.HwEngine3d
}

func (gx *GeometryEngine) CalcCmdCycles(code GxCmdCode) int64 {
	cycles := gxCmdDescs[code].ncycles
	if gx.mtxmode == 2 && (code == 0x18 || code == 0x19 || code == 0x1A || code == 0x1C) {
		cycles += 30
	}
	return cycles
}

func (gx *GeometryEngine) recalcClipMtx() {
	gx.clipmtx = matMul(gx.mtx[1], gx.mtx[0])
	// modGx.Infof("proj mtx: %v", gx.mtx[0])
	// modGx.Infof("pos mtx: %v", gx.mtx[1])
	// modGx.Infof("clip mtx: %v", gx.clipmtx)
}

func (gx *GeometryEngine) cmdNop(parms []GxCmd) {}

/******************************************************************
 * Matrix commands
 ******************************************************************/

func (gx *GeometryEngine) cmdMtxMode(parms []GxCmd) {
	gx.mtxmode = int(parms[0].parm & 3)
	// modGx.InfoZ("mtx mode").Int("mode", gx.mtxmode).End()
}

func (gx *GeometryEngine) cmdMtxLoad4x4(parms []GxCmd) {
	for j := 0; j < 4; j++ {
		for i := 0; i < 4; i++ {
			gx.mtx[gx.mtxmode][j][i].V = int32(parms[j*4+i].parm)
			// matrix mode 2 -> applies also to position matrix
			if gx.mtxmode == 2 {
				gx.mtx[1][j][i].V = int32(parms[j*4+i].parm)
			}
		}
	}
	if gx.mtxmode != 3 {
		gx.recalcClipMtx()
	}
	modGx.InfoZ("mtx load 4x4").
		Int("mode", gx.mtxmode).
		Vector12("r0", gx.mtx[gx.mtxmode][0]).
		Vector12("r1", gx.mtx[gx.mtxmode][1]).
		Vector12("r2", gx.mtx[gx.mtxmode][2]).
		Vector12("r3", gx.mtx[gx.mtxmode][3]).
		End()
}

func (gx *GeometryEngine) cmdMtxLoad4x3(parms []GxCmd) {
	for j := 0; j < 4; j++ {
		for i := 0; i < 3; i++ {
			gx.mtx[gx.mtxmode][j][i].V = int32(parms[j*3+i].parm)
			// matrix mode 2 -> applies also to position matrix
			if gx.mtxmode == 2 {
				gx.mtx[1][j][i].V = int32(parms[j*3+i].parm)
			}
		}
	}

	gx.mtx[gx.mtxmode][0][3] = fixed.NewF12(0)
	gx.mtx[gx.mtxmode][1][3] = fixed.NewF12(0)
	gx.mtx[gx.mtxmode][2][3] = fixed.NewF12(0)
	gx.mtx[gx.mtxmode][3][3] = fixed.NewF12(1)
	if gx.mtxmode == 2 {
		gx.mtx[1][0][3] = fixed.NewF12(0)
		gx.mtx[1][1][3] = fixed.NewF12(0)
		gx.mtx[1][2][3] = fixed.NewF12(0)
		gx.mtx[1][3][3] = fixed.NewF12(1)
	}

	if gx.mtxmode != 3 {
		gx.recalcClipMtx()
	}
	modGx.InfoZ("mtx load 4x3").
		Int("mode", gx.mtxmode).
		Vector12("r0", gx.mtx[gx.mtxmode][0]).
		Vector12("r1", gx.mtx[gx.mtxmode][1]).
		Vector12("r2", gx.mtx[gx.mtxmode][2]).
		End()
}

func (gx *GeometryEngine) cmdMtxIdentity(parms []GxCmd) {
	gx.mtx[gx.mtxmode] = newMatrixIdentity()
	// matrix mode 2 -> applies also to position matrix
	if gx.mtxmode == 2 {
		gx.mtx[1] = newMatrixIdentity()
	}
	if gx.mtxmode != 3 {
		gx.recalcClipMtx()
	}
	modGx.InfoZ("mtx identity").Int("mode", gx.mtxmode).End()
}

func (gx *GeometryEngine) matMulToCurrent(mt matrix) {
	gx.mtx[gx.mtxmode] = matMul(mt, gx.mtx[gx.mtxmode])
	// matrix mode 2 -> applies also to position matrix
	if gx.mtxmode == 2 {
		gx.mtx[1] = matMul(mt, gx.mtx[1])
	}
	if gx.mtxmode != 3 {
		gx.recalcClipMtx()
	}
}

func (gx *GeometryEngine) cmdMtxTrans(parms []GxCmd) {
	var x, y, z fixed.F12
	x.V = int32(parms[0].parm)
	y.V = int32(parms[1].parm)
	z.V = int32(parms[2].parm)
	mt := newMatrixTrans(x, y, z)
	gx.matMulToCurrent(mt)
}

func (gx *GeometryEngine) cmdMtxScale(parms []GxCmd) {
	var x, y, z fixed.F12
	x.V = int32(parms[0].parm)
	y.V = int32(parms[1].parm)
	z.V = int32(parms[2].parm)
	mt := newMatrixScale(x, y, z)

	// cmdMtxScale doesn't scale light direction matrix (for obvious reasons, as that
	// would make any light non-normalized). Basically, mode=2 is the same as mode=1,
	// and just the position matrix is affected.
	oldm := gx.mtxmode
	if gx.mtxmode == 2 {
		gx.mtxmode = 1
	}
	gx.matMulToCurrent(mt)
	gx.mtxmode = oldm
}

func (gx *GeometryEngine) cmdMtxMult4x4(parms []GxCmd) {
	var mtx matrix
	for i := 0; i < 16; i++ {
		mtx[i/4][i%4].V = int32(parms[i].parm)
	}
	gx.matMulToCurrent(mtx)
}

func (gx *GeometryEngine) cmdMtxMult4x3(parms []GxCmd) {
	var mtx matrix
	for i := 0; i < 12; i++ {
		mtx[i/3][i%3].V = int32(parms[i].parm)
	}
	mtx[3][3] = fixed.NewF12(1)
	gx.matMulToCurrent(mtx)
}

func (gx *GeometryEngine) cmdMtxMult3x3(parms []GxCmd) {
	var mtx matrix
	for i := 0; i < 9; i++ {
		mtx[i/3][i%3].V = int32(parms[i].parm)
	}
	mtx[3][3] = fixed.NewF12(1)
	gx.matMulToCurrent(mtx)
}

func (gx *GeometryEngine) cmdViewport(parms []GxCmd) {
	gx.vx0 = int((parms[0].parm >> 0) & 0xFF)
	gx.vy0 = int((parms[0].parm >> 8) & 0xFF)
	gx.vx1 = int((parms[0].parm >> 16) & 0xFF)
	gx.vy1 = int((parms[0].parm >> 24) & 0xFF)
	gx.e3d.CmdViewport(raster3d.Primitive_SetViewport{
		VX0: gx.vx0, VX1: gx.vx1, VY0: gx.vy0, VY1: gx.vy1,
	})
}

/******************************************************************
 * Matrix stack commands
 ******************************************************************/

func (gx *GeometryEngine) cmdMtxPush(parms []GxCmd) {
	switch gx.mtxmode {
	case MtxProjection:
		if gx.mtxStackProjPtr > 0 {
			gx.mtxStackOverflow = true
		}

		// The "1" entry is a mirror of "0", so always access 0
		gx.mtxStackProj[0] = gx.mtx[MtxProjection]
		gx.mtxStackProjPtr++
		gx.mtxStackProjPtr &= 1
	case MtxTexture:
		if gx.mtxStackTexPtr > 0 {
			gx.mtxStackOverflow = true
		}
		// The "1" entry is a mirror of "0", so always access 0
		gx.mtxStackTex[0] = gx.mtx[MtxTexture]
		gx.mtxStackTexPtr++
		gx.mtxStackTexPtr &= 1
	case MtxPosition, MtxDirection:
		if gx.mtxStackPosPtr > 30 {
			gx.mtxStackOverflow = true
		}

		gx.mtxStackPos[gx.mtxStackPosPtr&31] = gx.mtx[MtxPosition]
		gx.mtxStackDir[gx.mtxStackPosPtr&31] = gx.mtx[MtxDirection]
		gx.mtxStackPosPtr++
		gx.mtxStackPosPtr &= 63
	default:
		modGx.FatalZ("unknown matrix mode").Int("mode", int(gx.mtxmode)).End()
	}
}

func (gx *GeometryEngine) cmdMtxPop(parms []GxCmd) {
	switch gx.mtxmode {
	case MtxProjection:
		// NOTE: the offset parameter is ignored
		gx.mtxStackProjPtr -= 1
		gx.mtxStackProjPtr &= 1

		if gx.mtxStackProjPtr > 0 {
			gx.mtxStackOverflow = true
		}

		// The "1" entry is a mirror of "0", so always access 0
		gx.mtx[MtxProjection] = gx.mtxStackProj[0]
		gx.recalcClipMtx()
	case MtxTexture:
		// NOTE: the offset parameter is ignored
		gx.mtxStackTexPtr -= 1
		gx.mtxStackTexPtr &= 1

		if gx.mtxStackTexPtr > 0 {
			gx.mtxStackOverflow = true
		}

		// The "1" entry is a mirror of "0", so always access 0
		gx.mtx[MtxTexture] = gx.mtxStackTex[0]
	case MtxPosition, MtxDirection:
		// 6-bit signed offset, -30 / +31
		offset := int32((parms[0].parm&0x3F)<<26) >> 26
		gx.mtxStackPosPtr -= int(offset)
		gx.mtxStackPosPtr &= 63

		if gx.mtxStackPosPtr > 30 {
			gx.mtxStackOverflow = true
		}

		gx.mtx[MtxPosition] = gx.mtxStackPos[gx.mtxStackPosPtr&31]
		gx.mtx[MtxDirection] = gx.mtxStackDir[gx.mtxStackPosPtr&31]
		gx.recalcClipMtx()
	default:
		modGx.FatalZ("unknown matrix mode").Int("mode", int(gx.mtxmode)).End()
	}
}

func (gx *GeometryEngine) cmdMtxStore(parms []GxCmd) {
	switch gx.mtxmode {
	case 0:
		gx.mtxStackProj[0] = gx.mtx[0]
	case 1, 2:
		idx := int(parms[0].parm & 0x1F)
		if idx > 30 {
			// OVERFLOW FLAG
			modGx.FatalZ("MTX_STORE caused overflow in pos stack").End()
		}
		gx.mtxStackPos[idx] = gx.mtx[1]
		gx.mtxStackDir[idx] = gx.mtx[2]
	case 3:
		gx.mtxStackTex[0] = gx.mtx[3]
	}
}

func (gx *GeometryEngine) cmdMtxRestore(parms []GxCmd) {
	switch gx.mtxmode {
	case 0:
		gx.mtx[0] = gx.mtxStackProj[0]
		gx.recalcClipMtx()
	case 1, 2:
		idx := int(parms[0].parm & 0x1F)
		if idx > 30 {
			// OVERFLOW FLAG
			modGx.FatalZ("MTX_RESTORE caused overflow in pos stack").End()
		}
		gx.mtx[1] = gx.mtxStackPos[idx]
		gx.mtx[2] = gx.mtxStackDir[idx]
		gx.recalcClipMtx()
	case 3:
		gx.mtx[3] = gx.mtxStackTex[0]
	}
}

func (gx *GeometryEngine) cmdPolyAttr(parms []GxCmd) {
	gx.polyattr = parms[0].parm
}

func (gx *GeometryEngine) cmdBeginVtxs(parms []GxCmd) {
	gx.displist.polyattr = gx.polyattr
	gx.displist.primtype = int(parms[0].parm & 3)
	gx.displist.cnt = 0
	// modGx.Infof("begin vtx prim=%d, attr=%08x", gx.displist.primtype, gx.displist.polyattr)
}

func (gx *GeometryEngine) cmdEndVtxs(parms []GxCmd) {
	// dummy command, it is actually ignored by hardware
}

func (gx *GeometryEngine) cmdColor(parms []GxCmd) {
	r, g, b := uint8(parms[0].parm>>0)&0x1F, uint8(parms[0].parm>>5)&0x1F, uint8(parms[0].parm>>10)&0x1F

	gx.displist.color[0] = r
	gx.displist.color[1] = g
	gx.displist.color[2] = b
	modGx.InfoZ("color").Hex32("rgb", uint32(r)|uint32(g)<<8|uint32(b)<<16).End()
}

func (gx *GeometryEngine) cmdTexCoord(parms []GxCmd) {
	sx, tx := int16(parms[0].parm&0xFFFF), int16(parms[0].parm>>16)
	s, t := fixed.F12{V: int32(sx) << 8}, fixed.F12{V: int32(tx) << 8}

	switch gx.textrans {
	case 0:
		gx.displist.s = s
		gx.displist.t = t

	case 1:
		texv := vector{s, t, fixed.NewF12(1).Div(16), fixed.NewF12(1).Div(16)}
		s = texv.Dot(gx.mtx[MtxTexture].Col(0))
		t = texv.Dot(gx.mtx[MtxTexture].Col(1))

		// Internally, S/T are calculated as 1.11.4 (16bit); we truncate them
		// to the same precision.
		gx.displist.s = fixed.F12{V: int32(int16(s.V>>8)) << 8}
		gx.displist.t = fixed.F12{V: int32(int16(t.V>>8)) << 8}

	case 2, 3:
		// set basic coordinates, but will be modified by normal/vertex command
		gx.displist.s0 = s
		gx.displist.t0 = t

	default:
		panic("unreachable")
	}
}

func (gx *GeometryEngine) cmdTexImageParam(parms []GxCmd) {
	gx.texinfo.VramTexOffset = (parms[0].parm & 0xFFFF) * 8
	gx.texinfo.Width = 8 << ((parms[0].parm >> 20) & 7)
	gx.texinfo.Height = 8 << ((parms[0].parm >> 23) & 7)
	gx.texinfo.SFlipMask = 0
	gx.texinfo.TFlipMask = 0
	gx.texinfo.SClampMask = 0
	gx.texinfo.TClampMask = 0
	gx.texinfo.PitchShift = uint(3 + (parms[0].parm>>20)&7)
	gx.texinfo.Format = raster3d.TexFormat((parms[0].parm >> 26) & 7)
	gx.texinfo.ColorKey = (parms[0].parm>>29)&1 != 0
	gx.texinfo.Flags = 0
	if (parms[0].parm>>16)&1 != 0 {
		gx.texinfo.Flags |= raster3d.TexSRepeat
	} else {
		gx.texinfo.SClampMask = ^(gx.texinfo.Width - 1)
	}
	if (parms[0].parm>>17)&1 != 0 {
		gx.texinfo.Flags |= raster3d.TexTRepeat
	} else {
		gx.texinfo.TClampMask = ^(gx.texinfo.Height - 1)
	}
	if (parms[0].parm>>18)&1 != 0 {
		if gx.texinfo.Flags&raster3d.TexSRepeat == 0 {
			modGx.Warnf("texture with S Flip but not Repeat")
		} else {
			gx.texinfo.Flags |= raster3d.TexSFlip
			gx.texinfo.SFlipMask = gx.texinfo.Width
		}
	}
	if (parms[0].parm>>19)&1 != 0 {
		if gx.texinfo.Flags&raster3d.TexTRepeat == 0 {
			modGx.Warnf("texture with T Flip but not Repeat")
		} else {
			gx.texinfo.Flags |= raster3d.TexTFlip
			gx.texinfo.TFlipMask = gx.texinfo.Height
		}
	}

	gx.textrans = int((parms[0].parm >> 30) & 3)
	modGx.InfoZ("teximage").
		Uint32("w", gx.texinfo.Width).
		Uint32("h", gx.texinfo.Height).
		Int("fmt", int(gx.texinfo.Format)).
		Int("trans", gx.textrans).
		End()
}

func (gx *GeometryEngine) cmdTexPaletteBase(parms []GxCmd) {
	gx.texinfo.VramPalOffset = (parms[0].parm & 0x1FFF) * 16
}

func (gx *GeometryEngine) pushVertex(v vector) {

	gx.displist.lastvtx = v
	vw := gx.clipmtx.VecMul(v)
	s, t := gx.displist.s, gx.displist.t

	if gx.textrans == 3 {
		// Vertex source: texture coordinates are calculated in any VTX command
		s = gx.displist.s0.AddFixed(v.Dot3(gx.mtx[MtxTexture].Col(0)))
		t = gx.displist.t0.AddFixed(v.Dot3(gx.mtx[MtxTexture].Col(1)))
	}

	modGx.InfoZ("vertex").
		Vector12("obj", v).
		Vector12("wrd", vw).
		Hex32("rgb", gx.displist.color.To32bit()).
		End()

	gx.e3d.CmdVertex(raster3d.Primitive_Vertex{
		X: vw[0], Y: vw[1], Z: vw[2], W: vw[3],
		S: s, T: t,
		C: [3]uint8(gx.displist.color),
	})
	gx.vcnt++

	gx.displist.cnt++
	poly := raster3d.Primitive_Polygon{
		Attr: gx.displist.polyattr,
		Tex:  gx.texinfo,
	}

	// Adjust for palette offset difference for texformat 2
	// We do it here because it's the single point where we are sure
	// of both the texture format and the palette being used (since
	// they're set through different commands that can arrive in any
	// order)
	if poly.Tex.Format == raster3d.Tex4 {
		poly.Tex.VramPalOffset /= 2
	}

	switch gx.displist.primtype {
	case 0: // tri list
		if gx.displist.cnt%3 != 0 {
			break
		}
		poly.Vtx[0] = gx.vcnt - 3
		poly.Vtx[1] = gx.vcnt - 2
		poly.Vtx[2] = gx.vcnt - 1
		gx.e3d.CmdPolygon(poly)
	case 2: // tri strip
		if gx.displist.cnt >= 3 {
			poly.Vtx[0] = gx.vcnt - 3
			poly.Vtx[1] = gx.vcnt - 2
			poly.Vtx[2] = gx.vcnt - 1
			if gx.displist.cnt&1 == 0 {
				poly.Vtx[1], poly.Vtx[2] = poly.Vtx[2], poly.Vtx[1]
			}
			gx.e3d.CmdPolygon(poly)
		}

	case 1: // quad list
		if gx.displist.cnt%4 != 0 {
			break
		}
		poly.Vtx[0] = gx.vcnt - 4
		poly.Vtx[1] = gx.vcnt - 3
		poly.Vtx[2] = gx.vcnt - 2
		poly.Vtx[3] = gx.vcnt - 1
		poly.Attr |= (1 << 31) // overload bit 31 to specify quad
		gx.e3d.CmdPolygon(poly)
	case 3: // quad strip
		if gx.displist.cnt >= 4 && gx.displist.cnt&1 == 0 {
			poly.Vtx[0] = gx.vcnt - 4
			poly.Vtx[1] = gx.vcnt - 3
			poly.Vtx[2] = gx.vcnt - 1
			poly.Vtx[3] = gx.vcnt - 2
			poly.Attr |= (1 << 31) // overload bit 31 to specify quad
			gx.e3d.CmdPolygon(poly)
		}
	}
}

func (gx *GeometryEngine) cmdVtx16(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm))
	v[1].V = int32(int16(parms[0].parm >> 16))
	v[2].V = int32(int16(parms[1].parm))
	v[3] = fixed.NewF12(1)
	modGx.DebugZ("vtx16").Hex32("p0", parms[0].parm).Hex32("p1", parms[1].parm).Vector12("v", v).End()
	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtx10(parms []GxCmd) {
	var v vector
	v[0].V = int32(((parms[0].parm>>0)&0x3FF)<<22) >> 16
	v[1].V = int32(((parms[0].parm>>10)&0x3FF)<<22) >> 16
	v[2].V = int32(((parms[0].parm>>20)&0x3FF)<<22) >> 16
	v[3] = fixed.NewF12(1)
	modGx.DebugZ("vtx10").Hex32("p0", parms[0].parm).Vector12("v", v).End()
	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxXY(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm))
	v[1].V = int32(int16(parms[0].parm >> 16))
	v[2].V = gx.displist.lastvtx[2].V
	v[3] = fixed.NewF12(1)
	modGx.DebugZ("vxy").Hex32("p0", parms[0].parm).Vector12("v", v).End()
	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxXZ(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm))
	v[1].V = gx.displist.lastvtx[1].V
	v[2].V = int32(int16(parms[0].parm >> 16))
	v[3] = fixed.NewF12(1)
	modGx.DebugZ("vxz").Hex32("p0", parms[0].parm).Vector12("v", v).End()
	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxYZ(parms []GxCmd) {
	var v vector
	v[0].V = gx.displist.lastvtx[0].V
	v[1].V = int32(int16(parms[0].parm))
	v[2].V = int32(int16(parms[0].parm >> 16))
	v[3] = fixed.NewF12(1)
	modGx.DebugZ("vyz").Hex32("p0", parms[0].parm).Vector12("v", v).End()
	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxDiff(parms []GxCmd) {
	xd := int32(((parms[0].parm>>0)&0x3FF)<<22) >> 22
	yd := int32(((parms[0].parm>>10)&0x3FF)<<22) >> 22
	zd := int32(((parms[0].parm>>20)&0x3FF)<<22) >> 22

	var v vector
	v[0].V = gx.displist.lastvtx[0].V + xd
	v[1].V = gx.displist.lastvtx[1].V + yd
	v[2].V = gx.displist.lastvtx[2].V + zd
	v[3] = fixed.NewF12(1)
	modGx.DebugZ("vdiff").Hex32("p0", parms[0].parm).Vector12("v", v).End()
	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdDifAmb(parms []GxCmd) {
	gx.material[MatDiffuse] = newFcolorFrom555(
		(parms[0].parm>>0)&0x1F, (parms[0].parm>>5)&0x1F, (parms[0].parm>>10)&0x1F)

	if (parms[0].parm>>15)&1 != 0 {
		gx.displist.color = gx.material[MatDiffuse].ToColor()
	}

	gx.material[MatAmbient] = newFcolorFrom555(
		(parms[0].parm>>16)&0x1F, (parms[0].parm>>21)&0x1F, (parms[0].parm>>26)&0x1F)
}

func (gx *GeometryEngine) cmdSpeEmi(parms []GxCmd) {
	gx.material[MatSpecular] = newFcolorFrom555(
		(parms[0].parm>>0)&0x1F, (parms[0].parm>>5)&0x1F, (parms[0].parm>>10)&0x1F)

	gx.specTableOn = (parms[0].parm>>15)&1 != 0

	gx.material[MatEmission] = newFcolorFrom555(
		(parms[0].parm>>16)&0x1F, (parms[0].parm>>21)&0x1F, (parms[0].parm>>26)&0x1F)
}

func (gx *GeometryEngine) cmdLightColor(parms []GxCmd) {
	idx := parms[0].parm >> 30
	gx.lights[idx].color = newFcolorFrom555(
		(parms[0].parm>>0)&0x1F, (parms[0].parm>>5)&0x1F, (parms[0].parm>>10)&0x1F)
}

func (gx *GeometryEngine) cmdLightVector(parms []GxCmd) {
	idx := parms[0].parm >> 30

	x := fixed.F12{V: int32(((parms[0].parm>>0)&0x3FF)<<22) >> 19}
	y := fixed.F12{V: int32(((parms[0].parm>>10)&0x3FF)<<22) >> 19}
	z := fixed.F12{V: int32(((parms[0].parm>>20)&0x3FF)<<22) >> 19}

	v := vector{x, y, z}
	gx.lights[idx].dir = gx.mtx[MtxDirection].VecMul3x3(v)

	gx.lights[idx].half = gx.lights[idx].dir
	gx.lights[idx].half[2] = gx.lights[idx].half[2].Add(-1)
	gx.lights[idx].half[0].V /= 2
	gx.lights[idx].half[1].V /= 2
	gx.lights[idx].half[2].V /= 2
}

func (gx *GeometryEngine) cmdNormal(parms []GxCmd) {
	var n vector
	n[0].V = int32(((parms[0].parm>>0)&0x3FF)<<22) >> 19
	n[1].V = int32(((parms[0].parm>>10)&0x3FF)<<22) >> 19
	n[2].V = int32(((parms[0].parm>>20)&0x3FF)<<22) >> 19
	n[3].V = 0.0

	if gx.textrans == 2 {
		gx.displist.s = gx.displist.s0.AddFixed(n.Dot3(gx.mtx[MtxDirection].Col(0)))
		gx.displist.t = gx.displist.t0.AddFixed(n.Dot3(gx.mtx[MtxDirection].Col(1)))
	}

	n = gx.mtx[MtxDirection].VecMul3x3(n)

	color := gx.material[MatEmission]
	for i := uint(0); i < 4; i++ {
		// Check if light is activated
		if gx.displist.polyattr&(1<<i) == 0 {
			continue
		}
		// modGx.WithDelayedFields(func() log.Fields {
		// 	return log.Fields{
		// 		"light": gx.lights[i],
		// 		"diff":  gx.lights[i].dir.Dot(n),
		// 		"shine": gx.lights[i].half.Dot(n),
		// 	}
		// }).Infof("light")
		difflvl := gx.lights[i].dir.Dot(n)
		difflvl.V = -difflvl.V
		if difflvl.V < 0 {
			difflvl.V = 0
		}

		shinelvl := gx.lights[i].half.Dot(n)
		shinelvl.V = -shinelvl.V
		shinelvl.V *= shinelvl.V
		if shinelvl.V < 0 {
			shinelvl.V = 0
		} else if shinelvl.V >= 1<<12 {
			shinelvl.V = (1 << 12) - 1
		}

		if gx.specTableOn {
			shinelvl = gx.specTable[(shinelvl.V >> 5)]
		}

		for x := 0; x < 3; x++ {
			color[x].V += gx.material[MatSpecular][x].MulFixed(gx.lights[i].color[x]).MulFixed(shinelvl).V
			color[x].V += gx.material[MatDiffuse][x].MulFixed(gx.lights[i].color[x]).MulFixed(difflvl).V
			color[x].V += gx.material[MatAmbient][x].MulFixed(gx.lights[i].color[x]).V
		}
	}

	gx.displist.color = color.ToClampedColor()
	modGx.InfoZ("normal").
		Vector12("n", n).
		Fixed12("r", color[0]).
		Fixed12("g", color[1]).
		Fixed12("b", color[2]).
		Hex8("lights", uint8(gx.displist.polyattr&0xF)).
		End()
}

func (gx *GeometryEngine) cmdShininess(parms []GxCmd) {
	for i := 0; i < 128; i++ {
		val := parms[i>>2].parm
		val >>= uint(i&3) * 8
		val &= 0xFF
		gx.specTable[i].V = int32(val << 4)
	}
}

func (gx *GeometryEngine) cmdPosTest(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm & 0xFFFF))
	v[1].V = int32(int16(parms[0].parm >> 16))
	v[2].V = int32(int16(parms[1].parm & 0xFFFF))
	v[3].V = 1 << 12
	gx.displist.lastvtx = v // overwrite last vertex
	gx.posTestResult = gx.clipmtx.VecMul(v)
}

func (gx *GeometryEngine) cmdVecTest(parms []GxCmd) {
	var n vector
	n[0].V = int32(((parms[0].parm>>0)&0x3FF)<<22) >> 19
	n[1].V = int32(((parms[0].parm>>10)&0x3FF)<<22) >> 19
	n[2].V = int32(((parms[0].parm>>20)&0x3FF)<<22) >> 19
	n[3].V = 0
	gx.vecTestResult = gx.mtx[MtxDirection].VecMul3x3(n)
	// modGx.Warnf("n:%v res:%v mtx:%v", n, gx.vecTestResult, gx.mtx[MtxDirection])
}

func (gx *GeometryEngine) cmdSwapBuffers(parms []GxCmd) {
	gx.e3d.CmdSwapBuffers(raster3d.Primitive_SwapBuffers{
		AlphaYSort: parms[0].parm&1 != 0,
		WBuffering: parms[0].parm&2 != 0,
	})
	gx.vcnt = 0
}

//go:generate stringer -type GxCmdCode
const (
	GX_NOP GxCmdCode = 0x0

	GX_MTX_MODE     GxCmdCode = 0x10
	GX_MTX_PUSH     GxCmdCode = 0x11
	GX_MTX_POP      GxCmdCode = 0x12
	GX_MTX_STORE    GxCmdCode = 0x13
	GX_MTX_RESTORE  GxCmdCode = 0x14
	GX_MTX_IDENTITY GxCmdCode = 0x15
	GX_MTX_LOAD_4x4 GxCmdCode = 0x16
	GX_MTX_LOAD_4x3 GxCmdCode = 0x17
	GX_MTX_MULT_4x4 GxCmdCode = 0x18
	GX_MTX_MULT_4x3 GxCmdCode = 0x19
	GX_MTX_MULT_3x3 GxCmdCode = 0x1A
	GX_MTX_SCALE    GxCmdCode = 0x1B
	GX_MTX_TRANS    GxCmdCode = 0x1C

	GX_COLOR          GxCmdCode = 0x20
	GX_NORMAL         GxCmdCode = 0x21
	GX_TEXCOORD       GxCmdCode = 0x22
	GX_VTX_16         GxCmdCode = 0x23
	GX_VTX_10         GxCmdCode = 0x24
	GX_VTX_XY         GxCmdCode = 0x25
	GX_VTX_XZ         GxCmdCode = 0x26
	GX_VTX_YZ         GxCmdCode = 0x27
	GX_VTX_DIFF       GxCmdCode = 0x28
	GX_POLYGON_ATTR   GxCmdCode = 0x29
	GX_TEXIMAGE_PARAM GxCmdCode = 0x2A
	GX_PLTT_BASE      GxCmdCode = 0x2B

	GX_DIF_AMB      GxCmdCode = 0x30
	GX_SPE_EMI      GxCmdCode = 0x31
	GX_LIGHT_VECTOR GxCmdCode = 0x32
	GX_LIGHT_COLOR  GxCmdCode = 0x33
	GX_SHININESS    GxCmdCode = 0x34

	GX_BEGIN_VTXS GxCmdCode = 0x40
	GX_END_VTXS   GxCmdCode = 0x41

	GX_SWAP_BUFFERS GxCmdCode = 0x50

	GX_VIEWPORT GxCmdCode = 0x60

	GX_BOX_TEST GxCmdCode = 0x70
	GX_POS_TEST GxCmdCode = 0x71
	GX_VEC_TEST GxCmdCode = 0x72
)

var gxCmdDescs = []GxCmdDesc{
	// 0x0
	{0, 1, (*GeometryEngine).cmdNop}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x4
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x8
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0xC
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x10
	{1, 1, (*GeometryEngine).cmdMtxMode}, {0, 17, (*GeometryEngine).cmdMtxPush}, {1, 36, (*GeometryEngine).cmdMtxPop}, {1, 17, (*GeometryEngine).cmdMtxStore},
	// 0x14
	{1, 36, (*GeometryEngine).cmdMtxRestore}, {0, 19, (*GeometryEngine).cmdMtxIdentity}, {16, 34, (*GeometryEngine).cmdMtxLoad4x4}, {12, 30, (*GeometryEngine).cmdMtxLoad4x3},
	// 0x18
	{16, 35, (*GeometryEngine).cmdMtxMult4x4}, {12, 31, (*GeometryEngine).cmdMtxMult4x3}, {9, 28, (*GeometryEngine).cmdMtxMult3x3}, {3, 22, (*GeometryEngine).cmdMtxScale},
	// 0x1C
	{3, 22, (*GeometryEngine).cmdMtxTrans}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x20
	{1, 1, (*GeometryEngine).cmdColor}, {1, 9, (*GeometryEngine).cmdNormal}, {1, 1, (*GeometryEngine).cmdTexCoord}, {2, 9, (*GeometryEngine).cmdVtx16},
	// 0x24
	{1, 8, (*GeometryEngine).cmdVtx10}, {1, 8, (*GeometryEngine).cmdVtxXY}, {1, 8, (*GeometryEngine).cmdVtxXZ}, {1, 8, (*GeometryEngine).cmdVtxYZ},
	// 0x28
	{1, 8, (*GeometryEngine).cmdVtxDiff}, {1, 1, (*GeometryEngine).cmdPolyAttr}, {1, 1, (*GeometryEngine).cmdTexImageParam}, {1, 1, (*GeometryEngine).cmdTexPaletteBase},
	// 0x2C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x30
	{1, 4, (*GeometryEngine).cmdDifAmb}, {1, 4, (*GeometryEngine).cmdSpeEmi}, {1, 6, (*GeometryEngine).cmdLightVector}, {1, 1, (*GeometryEngine).cmdLightColor},
	// 0x34
	{32, 32, (*GeometryEngine).cmdShininess}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x38
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x3C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x40
	{1, 1, (*GeometryEngine).cmdBeginVtxs}, {0, 1, (*GeometryEngine).cmdEndVtxs}, {0, 0, nil}, {0, 0, nil},
	// 0x44
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x48
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x4C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x50
	{1, 392, (*GeometryEngine).cmdSwapBuffers}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x54
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x58
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x5C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x60
	{1, 1, (*GeometryEngine).cmdViewport}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x64
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x68
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x6C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x70
	{0, 0, nil}, {2, 9, (*GeometryEngine).cmdPosTest}, {1, 5, (*GeometryEngine).cmdVecTest}, {0, 0, nil},
	// 0x74
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x78
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x7C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
}
