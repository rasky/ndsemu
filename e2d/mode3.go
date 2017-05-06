package e2d

import "ndsemu/emu/gfx"

/************************************************
 * Display Mode 3: Main memory display
 ************************************************/

func (e2d *HwEngine2d) Mode3_BeginFrame() {
	modLcd.WarnZ("mode 3 not implemented").End()
}
func (e2d *HwEngine2d) Mode3_EndFrame() {}

func (e2d *HwEngine2d) Mode3_BeginLine(y int, screen gfx.Line) {
	for x := 0; x < cScreenWidth; x++ {
		screen.Set32(x, 0x0000FF)
	}
}
func (e2d *HwEngine2d) Mode3_EndLine(y int) {}
