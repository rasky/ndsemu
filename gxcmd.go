package main

import (
	"fmt"
	"ndsemu/emu"
)

type vector [4]emu.Fixed12
type matrix [4]vector

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
			c[j][i] = a.Col(i).Dot(b.Row(j))
		}
	}
	return
}

func (v vector) Dot(vb vector) (res emu.Fixed12) {
	res.V = (v[0].V*vb[0].V + v[1].V*vb[1].V + v[2].V*vb[2].V + v[3].V*vb[3].V) >> 12
	return
}

func (mtx *matrix) VecMul(vec vector) (res vector) {
	for i := 0; i < 4; i++ {
		res[i] = vec.Dot(mtx.Col(i))
	}
	return
}

type GeometryEngine struct {
	mtxmode  int
	mtx      [4]matrix // 0=proj, 1=pos, 2=vector, 3=tex
	clipmtx  matrix    // current clip matrix (pos * proj)
	vx0, vy0 int
	vx1, vy1 int
	polyattr uint32
	texinfo  RenderTexture
	textrans int
	displist struct {
		primtype int
		polyattr uint32
		t, s     emu.Fixed12
		cnt      int
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
	for j := 0; j < 3; j++ {
		for i := 0; i < 4; i++ {
			gx.mtx[gx.mtxmode][j][i].V = int32(parms[j*4+i].parm)
			// matrix mode 2 -> applies also to position matrix
			if gx.mtxmode == 2 {
				gx.mtx[1][j][i].V = int32(parms[j*4+i].parm)
			}
		}
	}

	gx.mtx[gx.mtxmode][3][0] = emu.NewFixed12(0)
	gx.mtx[gx.mtxmode][3][1] = emu.NewFixed12(0)
	gx.mtx[gx.mtxmode][3][2] = emu.NewFixed12(0)
	gx.mtx[gx.mtxmode][3][3] = emu.NewFixed12(1)
	if gx.mtxmode == 2 {
		gx.mtx[1][3][0] = emu.NewFixed12(0)
		gx.mtx[1][3][1] = emu.NewFixed12(0)
		gx.mtx[1][3][2] = emu.NewFixed12(0)
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

func (gx *GeometryEngine) cmdMtxTrans(parms []GxCmd) {
	var x, y, z emu.Fixed12
	x.V = int32(parms[0].parm)
	y.V = int32(parms[1].parm)
	z.V = int32(parms[2].parm)
	mt := newMatrixTrans(x, y, z)

	gx.mtx[gx.mtxmode] = matMul(gx.mtx[gx.mtxmode], mt)
	// matrix mode 2 -> applies also to position matrix
	if gx.mtxmode == 2 {
		gx.mtx[1] = matMul(gx.mtx[1], mt)
	}
	if gx.mtxmode != 3 {
		gx.recalcClipMtx()
	}
}

func (gx *GeometryEngine) cmdViewport(parms []GxCmd) {
	gx.vx0 = int((parms[0].parm >> 0) & 0xFF)
	gx.vy0 = int((parms[0].parm >> 8) & 0xFF)
	gx.vx1 = int((parms[0].parm >> 16) & 0xFF)
	gx.vy1 = int((parms[0].parm >> 24) & 0xFF)
	gx.E3dCmdCh <- E3DCmd_SetViewport{
		vx0: gx.vx0, vx1: gx.vx1, vy0: gx.vy0, vy1: gx.vy1,
	}
}

func (gx *GeometryEngine) cmdPolyAttr(parms []GxCmd) {
	gx.polyattr = parms[0].parm
}

func (gx *GeometryEngine) cmdBeginVtxs(parms []GxCmd) {
	gx.displist.polyattr = gx.polyattr
	gx.displist.primtype = int(parms[0].parm & 3)
	gx.displist.cnt = 0
}

func (gx *GeometryEngine) cmdEndVtxs(parms []GxCmd) {
	// dummy command, it is actually ignored by hardware
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
		gx.displist.s = texv.Dot(gx.mtx[3].Col(0))
		gx.displist.t = texv.Dot(gx.mtx[3].Col(1))

	default:
		panic("not implemented")
	}
}

func (gx *GeometryEngine) cmdTexImageParam(parms []GxCmd) {
	gx.texinfo.VramTexOffset = (parms[0].parm & 0xFFFF) * 8
	gx.texinfo.SMask = 8<<((parms[0].parm>>20)&7) - 1
	gx.texinfo.TMask = 8<<((parms[0].parm>>23)&7) - 1
	gx.texinfo.PitchShift = uint(3 + (parms[0].parm>>20)&7)
	gx.texinfo.Format = RTexFormat((parms[0].parm >> 26) & 7)
	gx.texinfo.Transparency = (parms[0].parm>>29)&1 != 0
	gx.texinfo.Flags = 0
	if (parms[0].parm>>16)&1 != 0 {
		gx.texinfo.Flags |= RTexSRepeat
	}
	if (parms[0].parm>>17)&1 != 0 {
		gx.texinfo.Flags |= RTexTRepeat
	}
	if (parms[0].parm>>18)&1 != 0 {
		gx.texinfo.Flags |= RTexSFlip
	}
	if (parms[0].parm>>19)&1 != 0 {
		gx.texinfo.Flags |= RTexTFlip
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
	vw := gx.clipmtx.VecMul(v)
	modGx.Infof("vertex: (%.2f,%.2f,%.2f) -> (%.2f,%.2f,%.2f,%.2f)",
		v[0].ToFloat64(), v[1].ToFloat64(), v[2].ToFloat64(),
		vw[0].ToFloat64(), vw[1].ToFloat64(), vw[2].ToFloat64(), vw[3].ToFloat64(),
	)

	gx.E3dCmdCh <- E3DCmd_Vertex{
		x: vw[0], y: vw[1], z: vw[2], w: vw[3],
		s: gx.displist.s, t: gx.displist.t,
	}
	gx.vcnt++

	gx.displist.cnt++
	poly := E3DCmd_Polygon{
		attr: gx.displist.polyattr,
		tex:  gx.texinfo,
	}

	// Adjust for palette offset difference for texformat 2
	// We do it here because it's the single point where we are sure
	// of both the texture format and the palette being used (since
	// they're set through different commands that can arrive in any
	// order)
	if poly.tex.Format == RTex4 {
		poly.tex.VramPalOffset /= 2
	}

	switch gx.displist.primtype {
	case 0: // tri list
		if gx.displist.cnt%3 != 0 {
			break
		}
		fallthrough
	case 2: // tri strip
		if gx.displist.cnt >= 3 {
			poly.vtx[0] = gx.vcnt - 3
			poly.vtx[1] = gx.vcnt - 2
			poly.vtx[2] = gx.vcnt - 1
			gx.E3dCmdCh <- poly
		}

	case 1: // quad list
		if gx.displist.cnt%4 != 0 {
			break
		}
		fallthrough
	case 3: // quad strip
		if gx.displist.cnt >= 4 {
			poly.vtx[0] = gx.vcnt - 4
			poly.vtx[1] = gx.vcnt - 3
			poly.vtx[2] = gx.vcnt - 2
			poly.vtx[3] = gx.vcnt - 1
			poly.attr |= (1 << 31) // overload bit 31 to specify quad
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

func (gx *GeometryEngine) cmdSwapBuffers(parms []GxCmd) {
	gx.E3dCmdCh <- E3DCmd_SwapBuffers{}
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
	{1, 1, (*GeometryEngine).cmdMtxMode}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x14
	{0, 0, nil}, {0, 19, (*GeometryEngine).cmdMtxIdentity}, {16, 34, (*GeometryEngine).cmdMtxLoad4x4}, {12, 30, (*GeometryEngine).cmdMtxLoad4x3},
	// 0x18
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x1C
	{3, 22, (*GeometryEngine).cmdMtxTrans}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x20
	{0, 0, nil}, {0, 0, nil}, {1, 1, (*GeometryEngine).cmdTexCoord}, {2, 9, (*GeometryEngine).cmdVtx16},
	// 0x24
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x28
	{0, 0, nil}, {1, 1, (*GeometryEngine).cmdPolyAttr}, {1, 1, (*GeometryEngine).cmdTexImageParam}, {1, 1, (*GeometryEngine).cmdTexPaletteBase},
	// 0x2C
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x30
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
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
