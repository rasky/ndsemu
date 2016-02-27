package e2d

import "ndsemu/emu"

// A pixel in a layer of the layer manager. It is composed as follows:
//   Bits 0-11: color index in the palette
//   Bit 12: set if the pixel uses the extended palette for its layer (either obj or bg)
//   Bit 13-14: priority
//   Bit 15: direct color
type LayerPixel uint16

func (p LayerPixel) Color() uint16       { return uint16(p & 0xFFF) }
func (p LayerPixel) ExtPal() bool        { return uint16(p&(1<<12)) != 0 }
func (p LayerPixel) Priority() uint16    { return uint16(p>>13) & 3 }
func (p LayerPixel) Direct() bool        { return int16(p) < 0 }
func (p LayerPixel) DirectColor() uint16 { return uint16(p & 0x7FFF) }
func (p LayerPixel) Transparent() bool   { return p == 0 }

func e2dMixer_Normal(layers []uint32, ctx interface{}) (res uint32) {
	var objpix, bgpix LayerPixel
	var pix uint16
	var cram []uint8
	var c16 uint16
	e2d := ctx.(*HwEngine2d)

	// Extract the layers. They've been already sorted in priority order,
	// so the first layer with a non-transparent pixel is the one that gets
	// drawn.
	bgpix = LayerPixel(layers[0])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[0]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}
	bgpix = LayerPixel(layers[1])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[1]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}
	bgpix = LayerPixel(layers[2])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[2]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}
	bgpix = LayerPixel(layers[3])
	if !bgpix.Transparent() {
		if bgpix.ExtPal() {
			cram = e2d.bgExtPals[3]
		} else {
			cram = e2d.bgPal
		}
		goto checkobj
	}

	// No bglayer was drawn here, so see if there is at least an obj, in which
	// case we draw it directly
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() {
		pix = objpix.Color()
		if objpix.ExtPal() {
			cram = e2d.objExtPal
		} else {
			cram = e2d.objPal
		}
		goto lookup
	}

	// No objlayer, and no bglayer. Draw the backdrop.
	// pix = 0
	cram = e2d.bgPal
	goto lookup

checkobj:
	// We found a bg pixel; now check if there is an object pixel here: if so,
	// we need to check the priority to choose between bg and obj which pixel
	// to draw (if the priorities are equal, objects win)
	objpix = LayerPixel(layers[4])
	if !objpix.Transparent() && objpix.Priority() <= bgpix.Priority() {
		pix = objpix.Color()
		if objpix.ExtPal() {
			cram = e2d.objExtPal
		} else {
			cram = e2d.objPal
		}
	} else {
		if bgpix.Direct() {
			c16 = bgpix.DirectColor()
			goto draw
		}
		pix = bgpix.Color()
	}

lookup:
	c16 = emu.Read16LE(cram[pix*2:])
draw:
	// Just return the 16-bit value for now, the post-processing
	// function will take care of the last step
	return uint32(c16)
}
