package e2d

import "ndsemu/emu/gfx"

/************************************************
 * Display Mode 0: display off
 ************************************************/

func (e2d *HwEngine2d) Mode0_BeginFrame() {}
func (e2d *HwEngine2d) Mode0_EndFrame()   {}

func (e2d *HwEngine2d) Mode0_BeginLine(y int, screen gfx.Line) {
	for x := 0; x < cScreenWidth; x++ {
		// Display off -> draw white
		screen.Set32(x, 0xFFFFFF)
	}
}
func (e2d *HwEngine2d) Mode0_EndLine(y int) {}
