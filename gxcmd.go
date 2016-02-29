package main

import (
	"fmt"
	"ndsemu/emu"
	"ndsemu/raster3d"
)

type vector [4]emu.Fixed12
type matrix [4]vector
type color [3]uint8

func (v vector) String() string {
	return fmt.Sprintf("v3d(%v,%v,%v,%v)", v[0], v[1], v[2], v[3])
}

func newMatrixIdentity() (m matrix) {
	m[0][0] = emu.NewFixed12(1)
	m[1][1] = emu.NewFixed12(1)
	m[2][2] = emu.NewFixed12(1)
	m[3][3] = emu.NewFixed12(1)
	return
}

func newMatrixTrans(x, y, z emu.Fixed12) matrix {
	m := newMatrixIdentity()
	m[3][0] = x
	m[3][1] = y
	m[3][2] = z
	return m
}

func newMatrixScale(x, y, z emu.Fixed12) (m matrix) {
	m[0][0] = x
	m[1][1] = y
	m[2][2] = z
	m[3][3] = emu.NewFixed12(1)
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

func (v vector) Dot(vb vector) (res emu.Fixed12) {
	val := int64(v[0].V) * int64(vb[0].V)
	val += int64(v[1].V) * int64(vb[1].V)
	val += int64(v[2].V) * int64(vb[2].V)
	val += int64(v[3].V) * int64(vb[3].V)
	res.V = int32(val >> 12)
	return
}

func (mtx *matrix) VecMul(vec vector) (res vector) {
	for i := 0; i < 4; i++ {
		res[i] = vec.Dot(mtx.Col(i))
	}
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
	mtxStackProj    [1]matrix
	mtxStackPos     [32]matrix
	mtxStackDir     [32]matrix
	mtxStackProjPtr int
	mtxStackPosPtr  int

	// Viewport
	vx0, vy0 int
	vx1, vy1 int

	// Material and lights
	material  [4]color
	spectable bool
	lights    [4]struct {
		dir   vector
		color color
	}

	// Textures
	texinfo  raster3d.Texture
	textrans int

	// Polygons and display lists
	polyattr uint32
	displist struct {
		primtype int
		polyattr uint32
		t, s     emu.Fixed12
		cnt      int
		lastvtx  vector
	}
	vcnt int

	// Channel to send commands to 3D rasterizer engine
	E3dCmdCh chan interface{}
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
	modGx.Infof("proj mtx: %v", gx.mtx[0])
	modGx.Infof("pos mtx: %v", gx.mtx[1])
	modGx.Infof("clip mtx: %v", gx.clipmtx)
}

func (gx *GeometryEngine) cmdNop(parms []GxCmd) {

}

/******************************************************************
 * Matrix commands
 ******************************************************************/

func (gx *GeometryEngine) cmdMtxMode(parms []GxCmd) {
	gx.mtxmode = int(parms[0].parm & 3)
	modGx.Infof("mtx mode: %d", gx.mtxmode)
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

	gx.mtx[gx.mtxmode][0][3] = emu.NewFixed12(0)
	gx.mtx[gx.mtxmode][1][3] = emu.NewFixed12(0)
	gx.mtx[gx.mtxmode][2][3] = emu.NewFixed12(0)
	gx.mtx[gx.mtxmode][3][3] = emu.NewFixed12(1)
	if gx.mtxmode == 2 {
		gx.mtx[1][0][3] = emu.NewFixed12(0)
		gx.mtx[1][1][3] = emu.NewFixed12(0)
		gx.mtx[1][2][3] = emu.NewFixed12(0)
		gx.mtx[1][3][3] = emu.NewFixed12(1)
	}

	if gx.mtxmode != 3 {
		gx.recalcClipMtx()
	}
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
	var x, y, z emu.Fixed12
	x.V = int32(parms[0].parm)
	y.V = int32(parms[1].parm)
	z.V = int32(parms[2].parm)
	mt := newMatrixTrans(x, y, z)
	gx.matMulToCurrent(mt)
}

func (gx *GeometryEngine) cmdMtxScale(parms []GxCmd) {
	var x, y, z emu.Fixed12
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
	mtx[3][3] = emu.NewFixed12(1)
	gx.matMulToCurrent(mtx)
}

func (gx *GeometryEngine) cmdMtxMult3x3(parms []GxCmd) {
	var mtx matrix
	for i := 0; i < 9; i++ {
		mtx[i/3][i%3].V = int32(parms[i].parm)
	}
	mtx[3][3] = emu.NewFixed12(1)
	gx.matMulToCurrent(mtx)
}

func (gx *GeometryEngine) cmdViewport(parms []GxCmd) {
	gx.vx0 = int((parms[0].parm >> 0) & 0xFF)
	gx.vy0 = int((parms[0].parm >> 8) & 0xFF)
	gx.vx1 = int((parms[0].parm >> 16) & 0xFF)
	gx.vy1 = int((parms[0].parm >> 24) & 0xFF)
	gx.E3dCmdCh <- raster3d.Primitive_SetViewport{
		VX0: gx.vx0, VX1: gx.vx1, VY0: gx.vy0, VY1: gx.vy1,
	}
}

/******************************************************************
 * Matrix stack commands
 ******************************************************************/

func (gx *GeometryEngine) cmdMtxPush(parms []GxCmd) {
	switch gx.mtxmode {
	case 0:
		if gx.mtxStackProjPtr > 0 {
			// OVERFLOW FLAG
			modGx.Fatal("MTX_PUSH caused overflow in proj stack")
		}

		// The "1" entry is a mirror of "0", so always access 0
		gx.mtxStackProj[0] = gx.mtx[0]
		gx.mtxStackProjPtr++
		gx.mtxStackProjPtr &= 1
	case 1, 2:
		if gx.mtxStackPosPtr > 30 {
			// OVERFLOW FLAG -- even if there are actually 32 entries, so 31 would be OK
			modGx.Fatal("MTX_PUSH caused overflow in pos stack")
		}

		gx.mtxStackPos[gx.mtxStackPosPtr&31] = gx.mtx[1]
		gx.mtxStackDir[gx.mtxStackPosPtr&31] = gx.mtx[2]
		gx.mtxStackPosPtr++
		gx.mtxStackPosPtr &= 63
	default:
		modGx.Fatalf("unsupported MTX_PUSH for mode=%d", gx.mtxmode)
	}
}

func (gx *GeometryEngine) cmdMtxPop(parms []GxCmd) {
	switch gx.mtxmode {
	case 0:
		// NOTE: the offset parameter is ignored
		gx.mtxStackProjPtr -= 1
		gx.mtxStackProjPtr &= 1

		if gx.mtxStackProjPtr > 0 {
			// OVERFLOW FLAG
			modGx.Fatal("MTX_POP caused overflow in proj stack")
		}

		// The "1" entry is a mirror of "0", so always access 0
		gx.mtx[0] = gx.mtxStackProj[0]
		gx.recalcClipMtx()
	case 1, 2:
		// 6-bit signed offset, -30 / +31
		offset := int32((parms[0].parm&0x3F)<<26) >> 26
		gx.mtxStackPosPtr -= int(offset)
		gx.mtxStackPosPtr &= 63

		if gx.mtxStackPosPtr > 30 {
			// OVERFLOW FLAG
			modGx.Fatal("MTX_POP caused overflow in pos stack")
		}

		gx.mtx[1] = gx.mtxStackPos[gx.mtxStackPosPtr&31]
		gx.mtx[2] = gx.mtxStackDir[gx.mtxStackPosPtr&31]
		gx.recalcClipMtx()
	default:
		modGx.Fatalf("unsupported MTX_POP for mode=%d", gx.mtxmode)
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
			modGx.Fatal("MTX_STORE caused overflow in pos stack")
		}
		gx.mtxStackPos[idx] = gx.mtx[1]
		gx.mtxStackDir[idx] = gx.mtx[2]
	default:
		modGx.Fatalf("unsupported MTX_STORE for mode=%d", gx.mtxmode)
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
			modGx.Fatal("MTX_RESTORE caused overflow in pos stack")
		}
		gx.mtx[1] = gx.mtxStackPos[idx]
		gx.mtx[2] = gx.mtxStackDir[idx]
		gx.recalcClipMtx()
	default:
		modGx.Fatalf("unsupported MTX_RESTORE for mode=%d", gx.mtxmode)
	}
}

func (gx *GeometryEngine) cmdPolyAttr(parms []GxCmd) {
	gx.polyattr = parms[0].parm
}

func (gx *GeometryEngine) cmdBeginVtxs(parms []GxCmd) {
	gx.displist.polyattr = gx.polyattr
	gx.displist.primtype = int(parms[0].parm & 3)
	gx.displist.cnt = 0
	modGx.Infof("begin vtx prim=%d, attr=%08x", gx.displist.primtype, gx.displist.polyattr)
}

func (gx *GeometryEngine) cmdEndVtxs(parms []GxCmd) {
	// dummy command, it is actually ignored by hardware
}

func (gx *GeometryEngine) cmdNormal(parms []GxCmd) {
	// TODO: implement
}

func (gx *GeometryEngine) cmdColor(parms []GxCmd) {
	// TODO: implement
}

func (gx *GeometryEngine) cmdTexCoord(parms []GxCmd) {
	sx, tx := int16(parms[0].parm&0xFFFF), int16(parms[0].parm>>16)
	s, t := emu.Fixed12{V: int32(sx) << 8}, emu.Fixed12{V: int32(tx) << 8}

	switch gx.textrans {
	case 0:
		gx.displist.s = s
		gx.displist.t = t

	case 1:
		texv := vector{s, t, emu.NewFixed12(1).Div(16), emu.NewFixed12(1).Div(16)}
		s = texv.Dot(gx.mtx[3].Col(0))
		t = texv.Dot(gx.mtx[3].Col(1))

		// Internally, S/T are calculated as 1.11.4 (16bit); we truncate them
		// to the same precision.
		gx.displist.s = emu.Fixed12{V: int32(int16(s.V>>8)) << 8}
		gx.displist.t = emu.Fixed12{V: int32(int16(t.V>>8)) << 8}

	default:
		panic("not implemented")
	}
}

func (gx *GeometryEngine) cmdTexImageParam(parms []GxCmd) {
	gx.texinfo.VramTexOffset = (parms[0].parm & 0xFFFF) * 8
	gx.texinfo.SMask = 8<<((parms[0].parm>>20)&7) - 1
	gx.texinfo.TMask = 8<<((parms[0].parm>>23)&7) - 1
	gx.texinfo.PitchShift = uint(3 + (parms[0].parm>>20)&7)
	gx.texinfo.Format = raster3d.TexFormat((parms[0].parm >> 26) & 7)
	gx.texinfo.Transparency = (parms[0].parm>>29)&1 != 0
	gx.texinfo.Flags = 0
	if (parms[0].parm>>16)&1 != 0 {
		gx.texinfo.Flags |= raster3d.TexSRepeat
	}
	if (parms[0].parm>>17)&1 != 0 {
		gx.texinfo.Flags |= raster3d.TexTRepeat
	}
	if (parms[0].parm>>18)&1 != 0 {
		gx.texinfo.Flags |= raster3d.TexSFlip
	}
	if (parms[0].parm>>19)&1 != 0 {
		gx.texinfo.Flags |= raster3d.TexTFlip
	}

	gx.textrans = int((parms[0].parm >> 30) & 3)
	if gx.textrans != 0 && gx.textrans != 1 {
		modGx.Fatalf("texture trans mode %d not implemented", gx.textrans)
	}
}

func (gx *GeometryEngine) cmdTexPaletteBase(parms []GxCmd) {
	gx.texinfo.VramPalOffset = (parms[0].parm & 0x1FFF) * 16
}

func (gx *GeometryEngine) pushVertex(v vector) {

	gx.displist.lastvtx = v
	vw := gx.clipmtx.VecMul(v)

	modGx.Infof("vertex: (%.2f,%.2f,%.2f) -> (%.2f,%.2f,%.2f,%.2f)",
		v[0].ToFloat64(), v[1].ToFloat64(), v[2].ToFloat64(),
		vw[0].ToFloat64(), vw[1].ToFloat64(), vw[2].ToFloat64(), vw[3].ToFloat64(),
	)

	gx.E3dCmdCh <- raster3d.Primitive_Vertex{
		X: vw[0], Y: vw[1], Z: vw[2], W: vw[3],
		S: gx.displist.s, T: gx.displist.t,
	}
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
		gx.E3dCmdCh <- poly
	case 2: // tri strip
		if gx.displist.cnt >= 3 {
			poly.Vtx[0] = gx.vcnt - 3
			poly.Vtx[1] = gx.vcnt - 2
			poly.Vtx[2] = gx.vcnt - 1
			if gx.displist.cnt&1 == 0 {
				poly.Vtx[1], poly.Vtx[2] = poly.Vtx[2], poly.Vtx[1]
			}
			gx.E3dCmdCh <- poly
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
		gx.E3dCmdCh <- poly
	case 3: // quad strip
		if gx.displist.cnt >= 4 && gx.displist.cnt&1 == 0 {
			poly.Vtx[0] = gx.vcnt - 4
			poly.Vtx[1] = gx.vcnt - 3
			poly.Vtx[2] = gx.vcnt - 1
			poly.Vtx[3] = gx.vcnt - 2
			poly.Attr |= (1 << 31) // overload bit 31 to specify quad
			gx.E3dCmdCh <- poly
		}
	}
}

func (gx *GeometryEngine) cmdVtx16(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm))
	v[1].V = int32(int16(parms[0].parm >> 16))
	v[2].V = int32(int16(parms[1].parm))
	v[3] = emu.NewFixed12(1)
	modGx.Infof("v16: %08x %08x -> %v\n", parms[0].parm, parms[1].parm, v)

	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtx10(parms []GxCmd) {
	var v vector
	v[0].V = int32(((parms[0].parm>>0)&0x3FF)<<22) >> 16
	v[1].V = int32(((parms[0].parm>>10)&0x3FF)<<22) >> 16
	v[2].V = int32(((parms[0].parm>>20)&0x3FF)<<22) >> 16
	v[3] = emu.NewFixed12(1)
	modGx.Infof("v10: %08x -> %v\n", parms[0].parm, v)

	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxXY(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm))
	v[1].V = int32(int16(parms[0].parm >> 16))
	v[2].V = gx.displist.lastvtx[2].V
	v[3] = emu.NewFixed12(1)
	modGx.Infof("vxy: %08x -> %v\n", parms[0].parm, v)

	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxXZ(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm))
	v[1].V = gx.displist.lastvtx[1].V
	v[2].V = int32(int16(parms[0].parm >> 16))
	v[3] = emu.NewFixed12(1)
	modGx.Infof("vxz: %08x -> %v\n", parms[0].parm, v)

	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxYZ(parms []GxCmd) {
	var v vector
	v[0].V = gx.displist.lastvtx[0].V
	v[1].V = int32(int16(parms[0].parm))
	v[2].V = int32(int16(parms[0].parm >> 16))
	v[3] = emu.NewFixed12(1)
	modGx.Infof("vyz: %08x -> %v\n", parms[0].parm, v)

	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdVtxDiff(parms []GxCmd) {
	xd := int32(((parms[0].parm>>0)&0x3FF)<<22) >> 19
	yd := int32(((parms[0].parm>>10)&0x3FF)<<22) >> 19
	zd := int32(((parms[0].parm>>20)&0x3FF)<<22) >> 19

	var v vector
	v[0].V = gx.displist.lastvtx[0].V + xd
	v[1].V = gx.displist.lastvtx[1].V + yd
	v[2].V = gx.displist.lastvtx[2].V + zd
	v[3] = emu.NewFixed12(1)
	modGx.Infof("vdiff: %08x -> %v\n", parms[0].parm, v)

	gx.pushVertex(v)
}

func (gx *GeometryEngine) cmdDifAmb(parms []GxCmd) {
	gx.material[MatDiffuse][0] = uint8((parms[0].parm >> 0) & 0x1F)
	gx.material[MatDiffuse][1] = uint8((parms[0].parm >> 5) & 0x1F)
	gx.material[MatDiffuse][2] = uint8((parms[0].parm >> 10) & 0x1F)

	if (parms[0].parm>>15)&1 != 0 {
		// FIXME: handle bit 15
		// apply as vertex color
	}

	gx.material[MatAmbient][0] = uint8((parms[0].parm >> 16) & 0x1F)
	gx.material[MatAmbient][1] = uint8((parms[0].parm >> 21) & 0x1F)
	gx.material[MatAmbient][2] = uint8((parms[0].parm >> 26) & 0x1F)
}

func (gx *GeometryEngine) cmdSpeEmi(parms []GxCmd) {
	gx.material[MatSpecular][0] = uint8((parms[0].parm >> 0) & 0x1F)
	gx.material[MatSpecular][1] = uint8((parms[0].parm >> 5) & 0x1F)
	gx.material[MatSpecular][2] = uint8((parms[0].parm >> 10) & 0x1F)

	gx.spectable = (parms[0].parm>>15)&1 != 0

	gx.material[MatEmission][0] = uint8((parms[0].parm >> 16) & 0x1F)
	gx.material[MatEmission][1] = uint8((parms[0].parm >> 21) & 0x1F)
	gx.material[MatEmission][2] = uint8((parms[0].parm >> 26) & 0x1F)
}

func (gx *GeometryEngine) cmdLightColor(parms []GxCmd) {
	idx := parms[0].parm >> 30
	gx.lights[idx].color[0] = uint8((parms[0].parm >> 0) & 0x1F)
	gx.lights[idx].color[1] = uint8((parms[0].parm >> 5) & 0x1F)
	gx.lights[idx].color[2] = uint8((parms[0].parm >> 10) & 0x1F)
}

func (gx *GeometryEngine) cmdLightVector(parms []GxCmd) {
	idx := parms[0].parm >> 30

	x := emu.Fixed12{V: int32(((parms[0].parm>>0)&0x3FF)<<22) >> 19}
	y := emu.Fixed12{V: int32(((parms[0].parm>>10)&0x3FF)<<22) >> 19}
	z := emu.Fixed12{V: int32(((parms[0].parm>>20)&0x3FF)<<22) >> 19}

	v := vector{x, y, z, emu.NewFixed12(1)}
	gx.lights[idx].dir = gx.mtx[2].VecMul(v)
}

func (gx *GeometryEngine) cmdSwapBuffers(parms []GxCmd) {
	gx.E3dCmdCh <- raster3d.Primitive_SwapBuffers{}
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
	{0, 0, (*GeometryEngine).cmdNop}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
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
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x38
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x3C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x40
	{1, 1, (*GeometryEngine).cmdBeginVtxs}, {0, 0, (*GeometryEngine).cmdEndVtxs}, {0, 0, nil}, {0, 0, nil},
	// 0x44
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x48
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x4C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x50
	{0, 392, (*GeometryEngine).cmdSwapBuffers}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
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
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x74
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x78
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x7C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
}
