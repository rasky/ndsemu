package main

import (
	"fmt"
	"ndsemu/emu"
)

type matrix struct {
	v [16]emu.Fixed22
}

func newMatrixIdentity() (m matrix) {
	m.v[0] = emu.NewFixed22(1)
	m.v[5] = emu.NewFixed22(1)
	m.v[10] = emu.NewFixed22(1)
	m.v[15] = emu.NewFixed22(1)
	return
}

type GeometryEngine struct {
	mtxmode  int
	mtx      [4]matrix // 0=proj, 1=pos, 2=vector, 3=tex
	vx0, vy0 int
	vx1, vy1 int
}

func (gx *GeometryEngine) cmdNop(parms []GxCmd) {

}

func (gx *GeometryEngine) cmdMtxMode(parms []GxCmd) {
	gx.mtxmode = int(parms[0].code & 3)
}

func (gx *GeometryEngine) cmdMtxLoad4x4(parms []GxCmd) {
	for i := 0; i < 16; i++ {
		gx.mtx[gx.mtxmode].v[i].V = int32(parms[0].parm)
		// matrix mode 2 -> applies also to position matrix
		if gx.mtxmode == 2 {
			gx.mtx[1].v[i].V = int32(parms[0].parm)
		}
	}
}

func (gx *GeometryEngine) cmdMtxIdentity(parms []GxCmd) {
	for i := 0; i < 16; i++ {
		gx.mtx[gx.mtxmode] = newMatrixIdentity()
		// matrix mode 2 -> applies also to position matrix
		if gx.mtxmode == 2 {
			gx.mtx[1] = newMatrixIdentity()
		}
	}
}

func (gx *GeometryEngine) cmdViewport(parms []GxCmd) {
	gx.vx0 = int((parms[0].parm >> 0) & 0xFF)
	gx.vy0 = int((parms[0].parm >> 8) & 0xFF)
	gx.vx1 = int((parms[0].parm >> 16) & 0xFF)
	gx.vx1 = int((parms[0].parm >> 24) & 0xFF)
}

func (gx *GeometryEngine) cmdSwapBuffers(parms []GxCmd) {

}

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

func (c GxCmdCode) String() string {
	switch c {
	case GX_NOP:
		return "GX_NOP"
	case GX_MTX_MODE:
		return "GX_MTX_MODE"
	case GX_MTX_PUSH:
		return "GX_MTX_PUSH"
	case GX_MTX_POP:
		return "GX_MTX_POP"
	case GX_MTX_STORE:
		return "GX_MTX_STORE"
	case GX_MTX_RESTORE:
		return "GX_MTX_RESTORE"
	case GX_MTX_IDENTITY:
		return "GX_MTX_IDENTITY"
	case GX_MTX_LOAD_4x4:
		return "GX_MTX_LOAD_4x4"
	case GX_MTX_LOAD_4x3:
		return "GX_MTX_LOAD_4x3"
	case GX_MTX_MULT_4x4:
		return "GX_MTX_MULT_4x4"
	case GX_MTX_MULT_4x3:
		return "GX_MTX_MULT_4x3"
	case GX_MTX_MULT_3x3:
		return "GX_MTX_MULT_3x3"
	case GX_MTX_SCALE:
		return "GX_MTX_SCALE"
	case GX_MTX_TRANS:
		return "GX_MTX_TRANS"
	case GX_COLOR:
		return "GX_COLOR"
	case GX_NORMAL:
		return "GX_NORMAL"
	case GX_TEXCOORD:
		return "GX_TEXCOORD"
	case GX_VTX_16:
		return "GX_VTX_16"
	case GX_VTX_10:
		return "GX_VTX_10"
	case GX_VTX_XY:
		return "GX_VTX_XY"
	case GX_VTX_XZ:
		return "GX_VTX_XZ"
	case GX_VTX_YZ:
		return "GX_VTX_YZ"
	case GX_VTX_DIFF:
		return "GX_VTX_DIFF"
	case GX_POLYGON_ATTR:
		return "GX_POLYGON_ATTR"
	case GX_TEXIMAGE_PARAM:
		return "GX_TEXIMAGE_PARAM"
	case GX_PLTT_BASE:
		return "GX_PLTT_BASE"
	case GX_DIF_AMB:
		return "GX_DIF_AMB"
	case GX_SPE_EMI:
		return "GX_SPE_EMI"
	case GX_LIGHT_VECTOR:
		return "GX_LIGHT_VECTOR"
	case GX_LIGHT_COLOR:
		return "GX_LIGHT_COLOR"
	case GX_SHININESS:
		return "GX_SHININESS"
	case GX_BEGIN_VTXS:
		return "GX_BEGIN_VTXS"
	case GX_END_VTXS:
		return "GX_END_VTXS"
	case GX_SWAP_BUFFERS:
		return "GX_SWAP_BUFFERS"
	case GX_VIEWPORT:
		return "GX_VIEWPORT"
	case GX_BOX_TEST:
		return "GX_BOX_TEST"
	case GX_POS_TEST:
		return "GX_POS_TEST"
	case GX_VEC_TEST:
		return "GX_VEC_TEST"
	default:
		return fmt.Sprintf("GX_UNKNOWN(%02x)", uint8(c))
	}
}

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
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x20
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x24
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
	// 0x28
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
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
	{0, 0, nil}, {0, 0, nil}, {0, 0, nil}, {0, 0, nil},
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
