package e2d

import (
	"ndsemu/emu"
	"ndsemu/emu/gfx"
)

/************************************************
 * Display Mode 2: VRAM display
 ************************************************/

func (e2d *HwEngine2d) Mode2_BeginFrame() {
	block := (e2d.DispCnt.Value >> 18) & 3
	modLcd.Infof("%s: mode=VRAM-Display bank:%s",
		string('A'+e2d.Idx), string('A'+block))

}
func (e2d *HwEngine2d) Mode2_EndFrame() {}

func (e2d *HwEngine2d) Mode2_BeginLine(y int, screen gfx.Line) {
	block := (e2d.DispCnt.Value >> 18) & 3
	vram := e2d.mc.VramRawBank(int(block))[y*cScreenWidth*2:]
	for x := 0; x < cScreenWidth; x++ {
		pix := emu.Read16LE(vram[x*2:])
		screen.Set32(x, uint32(pix))
	}
}
func (e2d *HwEngine2d) Mode2_EndLine(y int) {}
