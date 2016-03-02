package e2d

import (
	"fmt"
	"ndsemu/emu/gfx"
)

/************************************************
 * Display Mode 1: BG & OBJ layers
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

func (e2d *HwEngine2d) Mode1_BeginFrame() {
	// Set the 4 BG layer priorities
	for i := 0; i < 4; i++ {
		pri := uint(e2d.bgregs[i].priority())
		e2d.lm.SetLayerPriority(i, pri)
	}
	e2d.lm.SetLayerPriority(4, 100) // put sprites always last in the mixer

	bgmode := e2d.DispCnt.Value & 7
	bg3d := (e2d.DispCnt.Value>>3)&1 != 0

	// Switch bg0 between text and 3d layer, depending on setting in DISPCNT
	// Notice also that in bgmode 6, layer 0 is always 3d
	if e2d.A() {
		if bg3d || bgmode == 6 {
			e2d.Mode1_setBgMode(0, BgMode3D)
		} else {
			e2d.Mode1_setBgMode(0, BgModeText)
		}
	}

	// Bg1 is always text
	e2d.Mode1_setBgMode(1, BgModeText)

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
		e2d.Mode1_setBgMode(2, BgModeText)
		e2d.Mode1_setBgMode(3, BgModeText)
	case 1:
		e2d.Mode1_setBgMode(2, BgModeText)
		e2d.Mode1_setBgMode(3, BgModeAffine)
	case 2:
		e2d.Mode1_setBgMode(2, BgModeAffine)
		e2d.Mode1_setBgMode(3, BgModeAffine)
	case 3:
		e2d.Mode1_setBgMode(2, BgModeText)
		e2d.Mode1_setBgMode(3, ext[3])
	case 4:
		e2d.Mode1_setBgMode(2, BgModeAffine)
		e2d.Mode1_setBgMode(3, ext[3])
	case 5:
		e2d.Mode1_setBgMode(2, ext[2])
		e2d.Mode1_setBgMode(3, ext[3])
	case 6:
		if e2d.A() {
			// Large bitmap mode only supported on engine A
			e2d.Mode1_setBgMode(2, BgModeLargeBitmap)
		}
	}
	e2d.lm.BeginFrame()

	bg0on := (e2d.DispCnt.Value >> 8) & 1
	bg1on := (e2d.DispCnt.Value >> 9) & 1
	bg2on := (e2d.DispCnt.Value >> 10) & 1
	bg3on := (e2d.DispCnt.Value >> 11) & 1

	objon := (e2d.DispCnt.Value >> 12) & 1
	win0on := (e2d.DispCnt.Value >> 13) & 1
	win1on := (e2d.DispCnt.Value >> 14) & 1
	objwinon := (e2d.DispCnt.Value >> 15) & 1

	modLcd.Infof("%s: modes=%v bg=[%d,%d,%d,%d] obj=%d win=[%d,%d,%d]",
		string('A'+e2d.Idx), e2d.bgmodes, bg0on, bg1on, bg2on, bg3on, objon, win0on, win1on, objwinon)
	// modLcd.Infof("%s: scroll0=[%d,%d] scroll1=[%d,%d] scroll2=[%d,%d] scroll3=[%d,%d] size0=%d size3=%d",
	// 	string('A'+e2d.Idx),
	// 	e2d.Bg0XOfs.Value, e2d.Bg0YOfs.Value,
	// 	e2d.Bg1XOfs.Value, e2d.Bg1YOfs.Value,
	// 	e2d.Bg2XOfs.Value, e2d.Bg2YOfs.Value,
	// 	e2d.Bg3XOfs.Value, e2d.Bg3YOfs.Value,
	// 	e2d.Bg0Cnt.Value>>14, e2d.Bg3Cnt.Value>>13)
}

func (e2d *HwEngine2d) Mode1_EndFrame() {
	e2d.lm.EndFrame()
}

func (e2d *HwEngine2d) Mode1_BeginLine(y int, screen gfx.Line) {
	pram := e2d.mc.VramPalette(e2d.Idx)
	e2d.bgPal = pram[:512]
	e2d.objPal = pram[512:]

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
	for i := range e2d.bgExtPals {
		// When storing the pointer for each layer, we want to respect the
		// priority order; in fact, this array will be accessed by the mixer
		// function, so in that context bgExtPals[2] means the third layer in
		// priority order.
		lidx := e2d.lm.PriorityOrder[i]

		// Compute the BG Ext Palette slot used by each bg layer. Normally,
		// BG0 uses Slot 0, BG3 uses Slot 3, etc. but BG0 and BG1 can optionally
		// use a different slot (depending on bit 13 of BGxCNT register)
		slotnum := lidx
		if lidx == 0 && e2d.Bg0Cnt.Value&(1<<13) != 0 {
			slotnum = 2
		}
		if lidx == 1 && e2d.Bg1Cnt.Value&(1<<13) != 0 {
			slotnum = 3
		}

		e2d.bgExtPals[i] = bgextpal.FetchPointer(8 * 1024 * slotnum)
	}

	objextpal := e2d.mc.VramLinearBank(e2d.Idx, VramLinearOBJExtPal, 0)
	e2d.objExtPal = objextpal.FetchPointer(0)

	e2d.lm.BeginLine(screen)
}

func (e2d *HwEngine2d) Mode1_EndLine(y int) {
	e2d.lm.EndLine()
}

func (e2d *HwEngine2d) Mode1_setBgMode(lidx int, mode BgMode) {
	if e2d.bgmodes[lidx] == mode {
		return
	}

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
