package e2d

import (
	"fmt"
	"ndsemu/emu/gfx"
)

/************************************************
 * BG & OBJ layers
 ************************************************/

type BgMode int

//go:generate stringer -type BgMode
const (
	BgModeText BgMode = iota
	BgModeAffine
	BgModeAffineMap16
	BgModeAffineBitmap
	BgModeAffineBitmapDirect
	BgModeLargeBitmap
	BgMode3D
)

func (e2d *HwEngine2d) layers_BeginFrame() {
	bgmode := e2d.DispCnt.Value & 7
	bg3d := (e2d.DispCnt.Value>>3)&1 != 0

	// Switch bg0 between text and 3d layer, depending on setting in DISPCNT
	// Notice also that in bgmode 6, layer 0 is always 3d
	if e2d.A() {
		if bg3d || bgmode == 6 {
			e2d.setBgMode(0, BgMode3D)
			e2d.lm.ChangeLayer(5, gfx.NullLayer{})
			e2d.l3dIdx = 0
		} else {
			e2d.setBgMode(0, BgModeText)
			e2d.lm.ChangeLayer(5, e2d.l3d)
			e2d.l3dIdx = 5
		}
	} else {
		e2d.l3dIdx = -1
	}

	// Bg1 is always text
	e2d.setBgMode(1, BgModeText)

	// Compute extended mode (only for bg2/bg3)
	var ext [4]BgMode
	for idx := 2; idx < 4; idx++ {
		if *e2d.bgregs[idx].Cnt&(1<<7) == 0 {
			ext[idx] = BgModeAffineMap16
		} else if *e2d.bgregs[idx].Cnt&(1<<2) == 0 {
			ext[idx] = BgModeAffineBitmap
		} else {
			ext[idx] = BgModeAffineBitmapDirect
		}
	}

	// Change bg2/bg3 depending on mode
	switch bgmode {
	case 0:
		e2d.setBgMode(2, BgModeText)
		e2d.setBgMode(3, BgModeText)
	case 1:
		e2d.setBgMode(2, BgModeText)
		e2d.setBgMode(3, BgModeAffine)
	case 2:
		e2d.setBgMode(2, BgModeAffine)
		e2d.setBgMode(3, BgModeAffine)
	case 3:
		e2d.setBgMode(2, BgModeText)
		e2d.setBgMode(3, ext[3])
	case 4:
		e2d.setBgMode(2, BgModeAffine)
		e2d.setBgMode(3, ext[3])
	case 5:
		e2d.setBgMode(2, ext[2])
		e2d.setBgMode(3, ext[3])
	case 6:
		if e2d.A() {
			// Large bitmap mode only supported on engine A
			e2d.setBgMode(2, BgModeLargeBitmap)
		}
	}

	// Set the 4 BG layer priorities
	for i := 0; i < 4; i++ {
		pri := uint(e2d.bgregs[i].priority())
		e2d.lm.SetLayerPriority(i, pri)
	}
	e2d.lm.SetLayerPriority(4, 100) // put sprites always above BG layers
	e2d.lm.SetLayerPriority(5, 101) // 3D layer
	e2d.lm.SetLayerPriority(6, 102) // and window last

	// Begin frame in layer manager
	e2d.lm.BeginFrame()

	// DEBUG
	bg0on := (e2d.DispCnt.Value >> 8) & 1
	bg1on := (e2d.DispCnt.Value >> 9) & 1
	bg2on := (e2d.DispCnt.Value >> 10) & 1
	bg3on := (e2d.DispCnt.Value >> 11) & 1

	objon := (e2d.DispCnt.Value >> 12) & 1
	win0on := (e2d.DispCnt.Value >> 13) & 1
	win1on := (e2d.DispCnt.Value >> 14) & 1
	objwinon := (e2d.DispCnt.Value >> 15) & 1

	modLcd.Infof("%s: modes=%v bg=[%d,%d,%d,%d] pri=[%d %d %d %d] obj=%d win=[%d,%d,%d]",
		string('A'+e2d.Idx), e2d.bgmodes, bg0on, bg1on, bg2on, bg3on,
		e2d.bgregs[0].priority(), e2d.bgregs[1].priority(), e2d.bgregs[2].priority(), e2d.bgregs[3].priority(),
		objon, win0on, win1on, objwinon)
	// modLcd.Infof("%s: scroll0=[%d,%d] scroll1=[%d,%d] scroll2=[%d,%d] scroll3=[%d,%d] size0=%d size3=%d",
	// 	string('A'+e2d.Idx),
	// 	e2d.Bg0XOfs.Value, e2d.Bg0YOfs.Value,
	// 	e2d.Bg1XOfs.Value, e2d.Bg1YOfs.Value,
	// 	e2d.Bg2XOfs.Value, e2d.Bg2YOfs.Value,
	// 	e2d.Bg3XOfs.Value, e2d.Bg3YOfs.Value,
	// 	e2d.Bg0Cnt.Value>>14, e2d.Bg3Cnt.Value>>13)
}

func (e2d *HwEngine2d) layers_EndFrame() {
	e2d.lm.EndFrame()
}

func (e2d *HwEngine2d) layers_BeginLine(y int, screen gfx.Line) {
	pram := e2d.mc.VramPalette(e2d.Idx)
	bgPal := pram[:512]
	objPal := pram[512:]

	// Fetch direct pointers to extended palettes (both bg and obj). These
	// will be used by the mixer function to access the actual colors.
	//
	// Extended palettes are active if bit 30 (bg) and/or 31 (obj) in
	// DISPCNT are set, and of course the appropriate VRAM banks need to
	// be mapped. In any case, VramLinearBank() returnes an empty memory
	// area if the banks are unmapped, and we can even try to use it
	// (and display all black). This is possibly consistent with what NDS
	// does if the extended palettes are used without being properly
	// configured.
	bgextpal := e2d.mc.VramLinearBank(e2d.Idx, VramLinearBGExtPal, 0)
	objextpal := e2d.mc.VramLinearBank(e2d.Idx, VramLinearOBJExtPal, 0)

	for i := 0; i < 4; i++ {
		// Compute the BG Ext Palette slot used by each bg layer. Normally,
		// BG0 uses Slot 0, BG3 uses Slot 3, etc. but BG0 and BG1 can optionally
		// use a different slot (depending on bit 13 of BGxCNT register)
		slotnum := i
		if i == 0 && e2d.Bg0Cnt.Value&(1<<13) != 0 {
			slotnum = 2
		}
		if i == 1 && e2d.Bg1Cnt.Value&(1<<13) != 0 {
			slotnum = 3
		}

		e2d.allPals[i][0] = bgPal
		e2d.allPals[i][1] = bgextpal.FetchPointer(8 * 1024 * slotnum)
	}

	// OBJ layer
	e2d.allPals[4][0] = objPal
	e2d.allPals[4][1] = objextpal.FetchPointer(0)

	// Backdrop layer
	e2d.allPals[5][0] = bgPal
	e2d.allPals[5][1] = nil // not used

	e2d.lm.BeginLine(screen)
}

func (e2d *HwEngine2d) layers_EndLine(y int) {
	e2d.lm.EndLine()
}

func (e2d *HwEngine2d) setBgMode(lidx int, mode BgMode) {
	e2d.bgmodes[lidx] = mode
	switch mode {
	case BgModeText:
		e2d.lm.ChangeLayer(lidx, gfx.LayerFunc{Func: e2d.DrawBG})
	case BgMode3D:
		e2d.lm.ChangeLayer(lidx, e2d.l3d)
	case BgModeAffineMap16, BgModeAffineBitmapDirect, BgModeAffineBitmap, BgModeAffine:
		e2d.lm.ChangeLayer(lidx, gfx.LayerFunc{Func: e2d.DrawBGAffine})
	default:
		panic(fmt.Errorf("bgmode %v not implemented", mode))
	}
}
