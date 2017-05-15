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

const (
	// Initialize the pixel as layer backdrop (5) with lowest priority (3).
	BackdropPixel LayerPixel = (5<<26 | 3<<29)
)

type WindowPixel uint32

func (w WindowPixel) BgEnabled(lidx uint32) bool { return (w>>lidx)&1 != 0 }
func (w WindowPixel) SpritesEnabled() bool       { return w&(1<<4) != 0 }
func (w WindowPixel) FxEnabled() bool            { return w&(1<<5) != 0 }

func e2dMixer_Normal(layers []uint32, ctx interface{}) uint32 {
	var pix1 LayerPixel // top-level pixel
	var pix2 LayerPixel // second-level pixel (for blending effect)
	var rgb1 uint16

	e2d := ctx.(*HwEngine2d)

	// Incoming layers:
	// 0-3: BG layers in priority order
	// 4: Sprite layer
	// 5: 3D layer (if disabled -- this is used for capture)
	// 6: Window layer

	// Get the window pixel. Note that this also removes all bound checks
	wnd := WindowPixel(layers[6])

	// Check if there's a visible obj pixel; if so, extract its priority
	// otherwise use a out-of-band priority (0xF, real max is 3)
	objpri := uint32(0xF)
	objpix := LayerPixel(layers[4])
	if !objpix.Transparent() && wnd.SpritesEnabled() {
		objpri = objpix.Priority()
	}

	// Extract the layers. They've been already sorted in priority order,
	// so the first layer with a non-transparent pixel is the one that gets
	// drawn.
	pix1 = LayerPixel(layers[0])
	if !pix1.Transparent() && wnd.BgEnabled(pix1.Layer()) {
		if objpri <= pix1.Priority() {
			goto foundobjbg
		}
		goto findsecond1
	}
	pix1 = LayerPixel(layers[1])
	if !pix1.Transparent() && wnd.BgEnabled(pix1.Layer()) {
		if objpri <= pix1.Priority() {
			goto foundobjbg
		}
		goto findsecond2
	}
	pix1 = LayerPixel(layers[2])
	if !pix1.Transparent() && wnd.BgEnabled(pix1.Layer()) {
		if objpri <= pix1.Priority() {
			goto foundobjbg
		}
		goto findsecond3
	}
	pix1 = LayerPixel(layers[3])
	if !pix1.Transparent() && wnd.BgEnabled(pix1.Layer()) {
		if objpri <= pix1.Priority() {
			goto foundobjbg
		}
		goto findsecond4
	}

	// Initialize the pixel as layer backdrop (5) with lowest priority (3).
	// If the obj layer had a pixel, we're done (pix1=obj / pix2=backdrop)
	pix1 = BackdropPixel
	if objpri <= 3 {
		goto foundobjbg
	}

	// There's only backdrop. No pixel2
	pix2 = 0xFFFFFFFF
	goto draw

foundobjbg:
	// The obj pixel is in front of the bg pixel.
	pix2 = pix1
	pix1 = objpix
	goto draw

findsecond1:
	pix2 = LayerPixel(layers[1])
	if !pix2.Transparent() && wnd.BgEnabled(pix2.Layer()) {
		if objpri <= pix2.Priority() {
			pix2 = objpix
		}
		goto draw
	}
findsecond2:
	pix2 = LayerPixel(layers[2])
	if !pix2.Transparent() && wnd.BgEnabled(pix2.Layer()) {
		if objpri <= pix2.Priority() {
			pix2 = objpix
		}
		goto draw
	}
findsecond3:
	pix2 = LayerPixel(layers[3])
	if !pix2.Transparent() && wnd.BgEnabled(pix2.Layer()) {
		if objpri <= pix2.Priority() {
			pix2 = objpix
		}
		goto draw
	}
findsecond4:
	// No other bg layer contains pixels. If there was a obj pixel,
	// it's our second pixel, otherwise it's the backdrop
	if objpri <= 3 {
		pix2 = objpix
	} else {
		pix2 = BackdropPixel
	}
	goto draw

draw:
	lidx := pix1.Layer()
	if pix1.Direct() {
		rgb1 = pix1.DirectColor()
	} else {
		cram := e2d.allPals[lidx][pix1.ExtPal()]
		rgb1 = emu.Read16LE(cram[pix1.ColorIndex()*2:])
	}

	// Check if there's a special effect and is active on our first layer,
	// otherwise return immediately
	bld := e2d.BldCnt.Value
	fxmode := e2d.effectMode
	if !wnd.FxEnabled() || fxmode == 0 || (bld>>lidx)&1 == 0 {
		goto exit
	}

	// Special effect. Check if it's blending or brightness change
	if fxmode == 1 {
		// Alpha blending. We need to check the second pixel: the blending
		// is performed only if the second pixel is marked as target #2
		// in the blending control register (bits 8...13).
		l2idx := pix2.Layer()
		if (bld>>8>>l2idx)&1 != 0 {
			var rgb2 uint16
			if pix2.Direct() {
				rgb2 = pix2.DirectColor()
			} else {
				cram := e2d.allPals[l2idx][pix2.ExtPal()]
				rgb2 = emu.Read16LE(cram[pix2.ColorIndex()*2:])
			}

			r1, g1, b1 := rgb1&0x1f, (rgb1>>5)&0x1F, (rgb1>>10)&0x1F
			r2, g2, b2 := rgb2&0x1f, (rgb2>>5)&0x1F, (rgb2>>10)&0x1F

			// blend
			r1 = (r1*e2d.effectAlpha1 + r2*e2d.effectAlpha2) >> 4
			g1 = (g1*e2d.effectAlpha1 + g2*e2d.effectAlpha2) >> 4
			b1 = (b1*e2d.effectAlpha1 + b2*e2d.effectAlpha2) >> 4

			// clamp
			if r1 > 31 {
				r1 = 31
			}
			if g1 > 31 {
				g1 = 31
			}
			if b1 > 31 {
				b1 = 31
			}

			rgb1 = r1 | g1<<5 | b1<<10
		}
	} else {
		// Apply special brightness effects (if any)
		r, g, b := rgb1&0x1f, (rgb1>>5)&0x1F, (rgb1>>10)&0x1F
		rgb1 = e2d.effectBrightR[r] | e2d.effectBrightG[g] | e2d.effectBrightB[b]
	}

exit:
	// Return the output value
	return uint32(rgb1)
}
