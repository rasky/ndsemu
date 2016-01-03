// Generated on 2016-01-04 00:30:42.744734591 +0100 CET
package gfx

func (lm *LayerManager) fastmixer_1_8_8(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_1_16_8(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_1_32_8(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_1_8_16(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_1_16_16(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_1_32_16(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_1_8_32(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_1_16_32(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_1_32_32(screen Line) {
	var inbuf [1]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_2_8_8(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_2_16_8(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_2_32_8(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_2_8_16(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_2_16_16(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_2_32_16(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_2_8_32(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_2_16_32(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_2_32_32(screen Line) {
	var inbuf [2]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_3_8_8(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_3_16_8(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_3_32_8(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_3_8_16(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_3_16_16(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_3_32_16(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_3_8_32(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_3_16_32(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_3_32_32(screen Line) {
	var inbuf [3]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_4_8_8(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_4_16_8(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_4_32_8(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_4_8_16(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_4_16_16(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_4_32_16(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_4_8_32(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_4_16_32(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_4_32_32(screen Line) {
	var inbuf [4]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_5_8_8(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_5_16_8(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_5_32_8(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_5_8_16(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_5_16_16(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_5_32_16(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_5_8_32(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_5_16_32(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_5_32_32(screen Line) {
	var inbuf [5]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_6_8_8(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		inbuf[5] = uint32(l5.Get8(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_6_16_8(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		inbuf[5] = uint32(l5.Get16(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_6_32_8(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		inbuf[5] = l5.Get32(x)
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_6_8_16(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		inbuf[5] = uint32(l5.Get8(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_6_16_16(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		inbuf[5] = uint32(l5.Get16(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_6_32_16(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		inbuf[5] = l5.Get32(x)
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_6_8_32(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		inbuf[5] = uint32(l5.Get8(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_6_16_32(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		inbuf[5] = uint32(l5.Get16(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_6_32_32(screen Line) {
	var inbuf [6]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		inbuf[5] = l5.Get32(x)
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_7_8_8(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		inbuf[5] = uint32(l5.Get8(x))
		inbuf[6] = uint32(l6.Get8(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_7_16_8(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		inbuf[5] = uint32(l5.Get16(x))
		inbuf[6] = uint32(l6.Get16(x))
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_7_32_8(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		inbuf[5] = l5.Get32(x)
		inbuf[6] = l6.Get32(x)
		out := mix(in)
		screen.Set8(x, uint8(out))
	}
}

func (lm *LayerManager) fastmixer_7_8_16(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		inbuf[5] = uint32(l5.Get8(x))
		inbuf[6] = uint32(l6.Get8(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_7_16_16(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		inbuf[5] = uint32(l5.Get16(x))
		inbuf[6] = uint32(l6.Get16(x))
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_7_32_16(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		inbuf[5] = l5.Get32(x)
		inbuf[6] = l6.Get32(x)
		out := mix(in)
		screen.Set16(x, uint16(out))
	}
}

func (lm *LayerManager) fastmixer_7_8_32(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 4
	off0 := lm.Cfg.OverflowPixels * 1
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get8(x))
		inbuf[1] = uint32(l1.Get8(x))
		inbuf[2] = uint32(l2.Get8(x))
		inbuf[3] = uint32(l3.Get8(x))
		inbuf[4] = uint32(l4.Get8(x))
		inbuf[5] = uint32(l5.Get8(x))
		inbuf[6] = uint32(l6.Get8(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_7_16_32(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 8
	off0 := lm.Cfg.OverflowPixels * 2
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = uint32(l0.Get16(x))
		inbuf[1] = uint32(l1.Get16(x))
		inbuf[2] = uint32(l2.Get16(x))
		inbuf[3] = uint32(l3.Get16(x))
		inbuf[4] = uint32(l4.Get16(x))
		inbuf[5] = uint32(l5.Get16(x))
		inbuf[6] = uint32(l6.Get16(x))
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

func (lm *LayerManager) fastmixer_7_32_32(screen Line) {
	var inbuf [7]uint32
	in := inbuf[:]
	width := lm.Cfg.Width * 16
	off0 := lm.Cfg.OverflowPixels * 4
	mix := lm.Cfg.Mixer
	l0 := NewLine(lm.layers[0].buf[off0:])
	l1 := NewLine(lm.layers[1].buf[off0:])
	l2 := NewLine(lm.layers[2].buf[off0:])
	l3 := NewLine(lm.layers[3].buf[off0:])
	l4 := NewLine(lm.layers[4].buf[off0:])
	l5 := NewLine(lm.layers[5].buf[off0:])
	l6 := NewLine(lm.layers[6].buf[off0:])
	for x := 0; x < width; x++ {
		inbuf[0] = l0.Get32(x)
		inbuf[1] = l1.Get32(x)
		inbuf[2] = l2.Get32(x)
		inbuf[3] = l3.Get32(x)
		inbuf[4] = l4.Get32(x)
		inbuf[5] = l5.Get32(x)
		inbuf[6] = l6.Get32(x)
		out := mix(in)
		screen.Set32(x, uint32(out))
	}
}

var fastMixerTable = [128]func(*LayerManager, Line){
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_1_8_8,
	(*LayerManager).fastmixer_1_8_16,
	nil,
	(*LayerManager).fastmixer_1_8_32,
	(*LayerManager).fastmixer_1_16_8,
	(*LayerManager).fastmixer_1_16_16,
	nil,
	(*LayerManager).fastmixer_1_16_32,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_1_32_8,
	(*LayerManager).fastmixer_1_32_16,
	nil,
	(*LayerManager).fastmixer_1_32_32,
	(*LayerManager).fastmixer_2_8_8,
	(*LayerManager).fastmixer_2_8_16,
	nil,
	(*LayerManager).fastmixer_2_8_32,
	(*LayerManager).fastmixer_2_16_8,
	(*LayerManager).fastmixer_2_16_16,
	nil,
	(*LayerManager).fastmixer_2_16_32,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_2_32_8,
	(*LayerManager).fastmixer_2_32_16,
	nil,
	(*LayerManager).fastmixer_2_32_32,
	(*LayerManager).fastmixer_3_8_8,
	(*LayerManager).fastmixer_3_8_16,
	nil,
	(*LayerManager).fastmixer_3_8_32,
	(*LayerManager).fastmixer_3_16_8,
	(*LayerManager).fastmixer_3_16_16,
	nil,
	(*LayerManager).fastmixer_3_16_32,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_3_32_8,
	(*LayerManager).fastmixer_3_32_16,
	nil,
	(*LayerManager).fastmixer_3_32_32,
	(*LayerManager).fastmixer_4_8_8,
	(*LayerManager).fastmixer_4_8_16,
	nil,
	(*LayerManager).fastmixer_4_8_32,
	(*LayerManager).fastmixer_4_16_8,
	(*LayerManager).fastmixer_4_16_16,
	nil,
	(*LayerManager).fastmixer_4_16_32,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_4_32_8,
	(*LayerManager).fastmixer_4_32_16,
	nil,
	(*LayerManager).fastmixer_4_32_32,
	(*LayerManager).fastmixer_5_8_8,
	(*LayerManager).fastmixer_5_8_16,
	nil,
	(*LayerManager).fastmixer_5_8_32,
	(*LayerManager).fastmixer_5_16_8,
	(*LayerManager).fastmixer_5_16_16,
	nil,
	(*LayerManager).fastmixer_5_16_32,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_5_32_8,
	(*LayerManager).fastmixer_5_32_16,
	nil,
	(*LayerManager).fastmixer_5_32_32,
	(*LayerManager).fastmixer_6_8_8,
	(*LayerManager).fastmixer_6_8_16,
	nil,
	(*LayerManager).fastmixer_6_8_32,
	(*LayerManager).fastmixer_6_16_8,
	(*LayerManager).fastmixer_6_16_16,
	nil,
	(*LayerManager).fastmixer_6_16_32,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_6_32_8,
	(*LayerManager).fastmixer_6_32_16,
	nil,
	(*LayerManager).fastmixer_6_32_32,
	(*LayerManager).fastmixer_7_8_8,
	(*LayerManager).fastmixer_7_8_16,
	nil,
	(*LayerManager).fastmixer_7_8_32,
	(*LayerManager).fastmixer_7_16_8,
	(*LayerManager).fastmixer_7_16_16,
	nil,
	(*LayerManager).fastmixer_7_16_32,
	nil,
	nil,
	nil,
	nil,
	(*LayerManager).fastmixer_7_32_8,
	(*LayerManager).fastmixer_7_32_16,
	nil,
	(*LayerManager).fastmixer_7_32_32,
}
