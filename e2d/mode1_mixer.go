package e2d

import "ndsemu/emu"

// A pixel in a layer of the layer manager. It is composed as follows:
//   Bits 0-11: color index in the palette
//   Bit 12: set if the pixel uses the extended palette for its layer (either obj or bg)
//   Bit 26-28: layer index (0-3 bkg, 4 obj)
//   Bit 29-30: priority
//   Bit 31: direct color
type LayerPixel uint32

func (p LayerPixel) ColorIndex() uint16  { return uint16(p & 0xFFF) }
func (p LayerPixel) ExtPal() uint32      { return uint32(p>>12) & 1 }
func (p LayerPixel) Priority() uint32    { return uint32(p>>29) & 3 }
func (p LayerPixel) Layer() uint32       { return uint32(p>>26) & 7 }
func (p LayerPixel) Sprite() bool        { return (p>>28)&1 != 0 }
func (p LayerPixel) Direct() bool        { return int32(p) < 0 }
func (p LayerPixel) DirectColor() uint16 { return uint16(p & 0x7FFF) }
func (p LayerPixel) Transparent() bool   { return p == 0 }

type WindowPixel uint32

func (w WindowPixel) BgEnabled(lidx uint32) bool { return (w>>lidx)&1 != 0 }
func (w WindowPixel) SpritesEnabled() bool       { return (w>>4)&1 != 0 }
func (w WindowPixel) FxEnabled() bool            { return (w>>5)&1 != 0 }

func e2dMixer_Normal(layers []uint32, ctx interface{}) uint32 {
	var objpix, bgpix LayerPixel
	var pix uint16
	var cram []uint8
	var c16 uint16
	var fx bool
	e2d := ctx.(*HwEngine2d)

	// Incoming layers:
	// 0-3: BG layers in priority order
	// 4: Sprite layer
	// 5: 3D layer (if disabled -- this is used for capture)
	// 6: Window layer

	// Get the window pixel. Note that this also removes all bound checks
	wnd := WindowPixel(layers[6])

	// Extract the layers. They've been already sorted in priority order,
	// so the first layer with a non-transparent pixel is the one that gets
	// drawn.
	bgpix = LayerPixel(layers[0])
	if !bgpix.Transparent() && wnd.BgEnabled(bgpix.Layer()) {
		cram = e2d.bgAllPals[0+bgpix.ExtPal()]
		goto checkobj
	}
	bgpix = LayerPixel(layers[1])
	if !bgpix.Transparent() && wnd.BgEnabled(bgpix.Layer()) {
		cram = e2d.bgAllPals[2+bgpix.ExtPal()]
		goto checkobj
	}
	bgpix = LayerPixel(layers[2])
	if !bgpix.Transparent() && wnd.BgEnabled(bgpix.Layer()) {
		cram = e2d.bgAllPals[4+bgpix.ExtPal()]
		goto checkobj
	}
	bgpix = LayerPixel(layers[3])
	if !bgpix.Transparent() && wnd.BgEnabled(bgpix.Layer()) {
		cram = e2d.bgAllPals[6+bgpix.ExtPal()]
		goto checkobj
	}

	// No bglayer was drawn here, so see if there is at least an obj, in which
	// case we draw it directly
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() && wnd.SpritesEnabled() {
		goto drawobj
	}

	// No objlayer, and no bglayer. Draw the backdrop.
	// pix = 0
	cram = e2d.bgPal
	fx = (e2d.BldCnt.Value>>5)&1 != 0
	goto lookup

checkobj:
	// We found a bg pixel; now check if there is an object pixel here: if so,
	// we need to check the priority to choose between bg and obj which pixel
	// to draw (if the priorities are equal, objects win)
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() && wnd.SpritesEnabled() && objpix.Priority() <= bgpix.Priority() {
		goto drawobj
	}

	// Draw background pixel
	fx = (e2d.BldCnt.Value>>bgpix.Layer())&1 != 0
	if bgpix.Direct() {
		c16 = bgpix.DirectColor()
		goto draw
	}
	pix = bgpix.ColorIndex()
	goto lookup

drawobj:
	fx = (e2d.BldCnt.Value>>4)&1 != 0
	if objpix.Direct() {
		c16 = objpix.DirectColor()
		goto draw
	}
	pix = objpix.ColorIndex()
	cram = e2d.objAllPals[objpix.ExtPal()]

lookup:
	c16 = emu.Read16LE(cram[pix*2:])

draw:
	if fx && wnd.FxEnabled() {
		// Apply special brightness effects (if any)
		r, g, b := c16&0x1f, (c16>>5)&0x1F, (c16>>10)&0x1F
		c16 = e2d.effectBrightR[r] | e2d.effectBrightG[g] | e2d.effectBrightB[b]
	}

	// Return the output value
	return uint32(c16)
}
