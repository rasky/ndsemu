package main

import "ndsemu/emu"

type vector [4]emu.Fixed12
type matrix [4]vector

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

type RenderVertex struct {
	v vector
}

type RenderPolygon struct {
	attr uint32 // misc flags
	vtx  [4]int // indices of vertices in Vertex RAM
}

type GeometryEngine struct {
	mtxmode  int
	mtx      [4]matrix // 0=proj, 1=pos, 2=vector, 3=tex
	clipmtx  matrix    // current clip matrix (pos * proj)
	vx0, vy0 int
	vx1, vy1 int
	polyattr uint32
	displist struct {
		primtype int
		polyattr uint32
		cnt      int
	}

	vertexRam  []RenderVertex
	polygonRam []RenderPolygon
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
}

func (gx *GeometryEngine) cmdNop(parms []GxCmd) {

}

func (gx *GeometryEngine) cmdMtxMode(parms []GxCmd) {
	gx.mtxmode = int(parms[0].code & 3)
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
	gx.vx1 = int((parms[0].parm >> 24) & 0xFF)
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

func (gx *GeometryEngine) pushPolygon(poly RenderPolygon) {
	gx.polygonRam = append(gx.polygonRam, poly)
}

func (gx *GeometryEngine) pushVertex(vtx RenderVertex) {
	gx.vertexRam = append(gx.vertexRam, vtx)
	gx.displist.cnt += 1
	vcnt := len(gx.vertexRam)

	poly := RenderPolygon{
		attr: gx.displist.polyattr,
	}

	switch gx.displist.primtype {
	case 0: // tri list
		if gx.displist.cnt%3 != 0 {
			break
		}
		fallthrough
	case 2: // tri strip
		if gx.displist.cnt >= 3 {
			poly.vtx[0] = vcnt - 2
			poly.vtx[1] = vcnt - 1
			poly.vtx[2] = vcnt - 0
			gx.pushPolygon(poly)
		}

	case 1: // quad list
		if gx.displist.cnt%4 != 0 {
			break
		}
		fallthrough
	case 3: // quad strip
		if gx.displist.cnt >= 4 {
			poly.vtx[0] = vcnt - 3
			poly.vtx[1] = vcnt - 2
			poly.vtx[2] = vcnt - 1
			poly.vtx[3] = vcnt - 0
			poly.attr |= (1 << 31) // overload bit 31 to specify quad
			gx.pushPolygon(poly)
		}
	}
}

func (gx *GeometryEngine) cmdVtx16(parms []GxCmd) {
	var v vector
	v[0].V = int32(int16(parms[0].parm))
	v[1].V = int32(int16(parms[0].parm >> 16))
	v[2].V = int32(int16(parms[1].parm))
	v[3].V = 1

	s := gx.clipmtx.VecMul(v)
	vtx := RenderVertex{s}
	gx.pushVertex(vtx)
}

func (gx *GeometryEngine) cmdSwapBuffers(parms []GxCmd) {
	gx.vertexRam = nil
	gx.polygonRam = nil
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
	{0, 0, nil}, {0, 19, (*GeometryEngine).cmdMtxIdentity}, {16, 34, (*GeometryEngine).cmdMtxLoad4x4}, {0, 0, nil},
	// 0x18
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x1C
	{3, 22, (*GeometryEngine).cmdMtxTrans}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x20
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {2, 9, (*GeometryEngine).cmdVtx16},
	// 0x24
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x28
	{0, 0, nil}, {1, 1, (*GeometryEngine).cmdPolyAttr}, {0, 0, nil}, {0, 0, nil},
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
