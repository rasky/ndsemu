package e2d

import (
	"fmt"
	"ndsemu/emu/gfx"
)

func (e2d *HwEngine2d) winXCoord(winid int) (int, int) {
	var xreg uint16
	if winid == 0 {
		xreg = e2d.Win0X.Value
	} else {
		xreg = e2d.Win1X.Value
	}

	x2 := xreg & 0xFF
	x1 := xreg >> 8
	screenWidth := uint16(e2d.ScreenWidth())
	if x2 > screenWidth || x1 > x2 {
		x2 = screenWidth
	}
	return int(x1), int(x2)
}

func (e2d *HwEngine2d) winYCoord(winid int) (int, int) {
	var yreg uint16
	if winid == 0 {
		yreg = e2d.Win0Y.Value
	} else {
		yreg = e2d.Win1Y.Value
	}

	y2 := yreg & 0xFF
	y1 := yreg >> 8
	screenHeight := uint16(e2d.ScreenHeight())
	if y2 > screenHeight || y1 > y2 {
		y2 = screenHeight
	}
	return int(y1), int(y2)
}

func (e2d *HwEngine2d) DrawWindow(lidx int) func(gfx.Line) {
	cScreenWidth := e2d.ScreenWidth()

	// Allocate object window buffer
	objWin := make([]byte, (cScreenWidth+e2d.lm.Cfg.OverflowPixels*2)*4)
	objWinLine := gfx.NewLine(objWin)
	objWinLine.Add32(e2d.lm.Cfg.OverflowPixels)

	drawObjWin := e2d.DrawOBJWindow(lidx)

	if false {
		e2d.debugWindow()
	}

	y := 0
	return func(out gfx.Line) {
		w0en := (e2d.DispCnt.Value>>13)&1 != 0
		w1en := (e2d.DispCnt.Value>>14)&1 != 0
		objwen := (e2d.DispCnt.Value>>15)&1 != 0

		// Always draw the object window, so that it's
		// at the correct line.
		for i := 0; i < cScreenWidth; i++ {
			objWinLine.Set32(i, 0)
		}
		drawObjWin(objWinLine)

		// Default: everything enabled
		for i := 0; i < cScreenWidth; i++ {
			out.Set32(i, 0xFFFFFFFF)
		}

		// Draw the "outside area" (lowest pri)
		if w0en || w1en || objwen {
			mask := uint8(e2d.WinOut.Value & 0xFF)
			for i := 0; i < cScreenWidth; i++ {
				out.Set32(i, uint32(mask))
			}
		}

		// Draw the object window.
		if objwen {
			mask := uint8(e2d.WinOut.Value >> 8)
			for i := 0; i < cScreenWidth; i++ {
				if objWinLine.Get32(i) != 0 {
					out.Set32(i, uint32(mask))
				}
			}
		}

		// Draw the window 1
		if w1en {
			mask := uint8(e2d.WinIn.Value >> 8)
			x1, x2 := e2d.winXCoord(1)
			y1, y2 := e2d.winYCoord(1)
			if y >= y1 && y < y2 {
				for x := x1; x < x2; x++ {
					out.Set32(x, uint32(mask))
				}
			}
		}

		// Draw the window 0 (highest pri)
		if w0en {
			mask := uint8(e2d.WinIn.Value & 0xFF)
			x1, x2 := e2d.winXCoord(0)
			y1, y2 := e2d.winYCoord(0)
			if y >= y1 && y < y2 {
				for x := x1; x < x2; x++ {
					out.Set32(x, uint32(mask))
				}
			}
		}

		y++
	}
}

func (e2d *HwEngine2d) debugWindow() {
	w0en := (e2d.DispCnt.Value>>13)&1 != 0
	w1en := (e2d.DispCnt.Value>>14)&1 != 0
	objwen := (e2d.DispCnt.Value>>15)&1 != 0

	w0mask := uint8(e2d.WinIn.Value & 0xFF)
	w1mask := uint8(e2d.WinIn.Value >> 8)
	objmask := uint8(e2d.WinOut.Value >> 8)
	woutmask := uint8(e2d.WinOut.Value & 0xFF)

	w0x1, w0x2 := e2d.winXCoord(0)
	w0y1, w0y2 := e2d.winYCoord(0)
	w1x1, w1x2 := e2d.winXCoord(1)
	w1y1, w1y2 := e2d.winYCoord(1)

	names := [...]string{"bg0", "bg1", "bg2", "bg3", "obj", "fx"}
	var w0names, w1names, objnames, woutnames []string

	for i := range names {
		if w0mask&(1<<uint(i)) != 0 {
			w0names = append(w0names, names[i])
		}
		if w1mask&(1<<uint(i)) != 0 {
			w1names = append(w1names, names[i])
		}
		if objmask&(1<<uint(i)) != 0 {
			objnames = append(objnames, names[i])
		}
		if woutmask&(1<<uint(i)) != 0 {
			woutnames = append(woutnames, names[i])
		}
	}

	if w0en || w1en || objwen {
		fmt.Printf("WIN0:%v:(%d,%d)-(%d,%d)  WIN1:%v:(%d,%d)-(%d,%d)  OBJ:%v\n",
			w0en, w0x1, w0y1, w0x2, w0y2,
			w1en, w1x1, w1y1, w1x2, w1y2,
			objwen)
		fmt.Printf("WIN0:%v   WIN1:%v  OBJ:%v  WINOUT:%v\n", w0names, w1names, objnames, woutnames)
	}
}
