package gfx

import (
	"sync"
)

//go:generate go run genmixer/genmixer.go -filename fastmixer.go

type Layer interface {
	// DrawFrame is the entry point of the drawing code for each layer. It receives
	// the index of the layer itself in the layer manager.
	//
	// The function must do the initial setup, and then return a closure (function)
	// that will be invoked to draw each line (gfx.Line), in sequence.
	// This is an example of the skeleton of a correct DrawFrame implementation:
	//
	//  func (l *MyLayer) DrawFrame(layerIdx int) func(Line) {
	//
	//      [initial setup...]
	//      y = 0
	//
	//      return func(out gfx.Line) {
	//          [ draw line y into out ]
	//          y++
	//      }
	//  }
	//
	// layerIdx is passed as a way to reuse the same code for multiple layers. If
	// you have a different codepath for each layer, then you can safely ignore this
	// argument.
	//
	DrawLayer(layerIdx int) func(Line)
}

type layerData struct {
	Layer
	pri     uint       // priority value for this layer
	linebuf []byte     // pixel buffer for this layer
	next    func(Line) // function to draw next line
}

type LayerManagerConfig struct {
	// Size of the screen
	Width, Height int

	// Bytes per pixel of the screen (usually 4)
	ScreenBpp int

	// Bytes per pixel of a layer. This can be lower than the depth of the
	// screen, so that each layer draw less bytes, befor being mixed in
	LayerBpp int

	// Overflow pixels for each line. This is a number of horizontal
	// pixels each layer will have available before and after the line
	// buffer, so that they can avoid clipping at the pixel level.
	OverflowPixels int

	// Per-pixel mixer function. This function is used to mix all the layers
	// into the final output, and is called one time for each pixel. It receives
	// as input the pixel value as fetched for each layer, and must return the
	// output value to be copied onto the screen.
	// The input slice is guaranteed to have as many elements as the number of
	// layers in the layer manager.
	// In the input slice, pixels are sorted by priority order, not pixel order;
	// for instance, pixels[0] is not coming from the first layer, but from the
	// layer with the lower priority.
	Mixer func(pixels []uint32, ctx interface{}) uint32

	// Context to be passed to the mixer function
	MixerCtx interface{}

	// Per-line post-processing function. This function can be used
	// to apply post-processing computation on the whole line, after it
	// went through the mixer. For performance reasons, you may want to
	// keep all processing within the mixer, but sometimes that's not possible
	// or can be slower than doing everything in one go later.
	PostProc func(line Line, ctx interface{})

	PostProcCtx interface{}
}

type LayerManager struct {
	// Configuration of the layer manager. The configuration can safely be
	// changed between frames, but only some fields can be changed between
	// lines. Please refer to the documentation of each field.
	Cfg LayerManagerConfig

	// Indices of layers, sorted by priority. This array is sorted by the
	// LayerManager while operating, and it is guaranteed to be correctly
	// populated if accessed from within the mixer function.
	PriorityOrder []int

	layers  []*layerData
	mixbuf  [8]uint32
	y       int
	setupWg sync.WaitGroup
	lineWg  sync.WaitGroup
}

func (lm *LayerManager) AddLayer(l Layer) int {
	idx := len(lm.layers)
	lm.layers = append(lm.layers, nil)
	lm.ChangeLayer(idx, l)
	return idx
}

func (lm *LayerManager) ChangeLayer(lidx int, l Layer) {
	lm.layers[lidx] = &layerData{
		Layer: l,
	}
}

// Set the priority value for a layer. The priorty is an unsigned value that is
// used to sort layers before passing them to the mixer. In case of identical
// priority, the layer index is used to disambiguate the order (basically,
// the sort is stable with respect to the natural order).
// For instance, if there are 4 layers that have priorities [3,0,4,0], the order
// of pixels passed to the mixer will be: layer #1, layer #3, layer #0, layer #2
func (lm *LayerManager) SetLayerPriority(lidx int, pri uint) {

	// Since we don't sort layers from scratch everytime, we can't guarantee
	// a stable sort with respect to the natural order, in face of different
	// changes. Thus, we must account for our sort to be unstable, and tweak
	// the priority order to disambiguate identical priorities.
	if pri > 1<<20 {
		panic("priority out of range")
	}
	lm.layers[lidx].pri = pri<<8 | uint(lidx)
}

func (lm *LayerManager) BeginFrame() {
	lm.y = -1

	// Initialize the order array (this only triggers the first time,
	// or if there was a change in layer allocation between frames)
	if len(lm.PriorityOrder) != len(lm.layers) {
		lm.PriorityOrder = make([]int, len(lm.layers))
		for i := range lm.PriorityOrder {
			lm.PriorityOrder[i] = i
		}
	}

	buflen := (lm.Cfg.Width + lm.Cfg.OverflowPixels*2) * lm.Cfg.LayerBpp

	for idx, l := range lm.layers {
		l.next = l.DrawLayer(idx)

		// Allocate the line buffer for this layer, if we haven't already
		if len(l.linebuf) != buflen {
			l.linebuf = make([]byte, buflen)
		}
	}
}

func (lm *LayerManager) drawLine(line Line) {
	// Sort layers by priority (that is, sort the order array so that it
	// contains the layers indices in priority order).
	// We use an inline insertion sort. In most cases, order will already be
	// sorted, so insertion sort degrades to simply checking that the array
	// is still sorted; even if we do some sorting, we're still much
	// faster than calling sort.Sort with all the overhead it brings.
	layers := lm.layers
	order := lm.PriorityOrder
	for i := 1; i < len(order); i++ {
		cur := order[i]
		j := i - 1
		for layers[order[j]].pri > layers[cur].pri {
			order[j+1] = order[j]
			j--
			if j < 0 {
				break
			}
		}
		order[j+1] = cur
	}
	lm.PriorityOrder = order

	// Send new line to each layer
	off0 := lm.Cfg.OverflowPixels * lm.Cfg.LayerBpp
	for _, l := range lm.layers {
		for i := range l.linebuf {
			l.linebuf[i] = 0x0
		}
		l.next(NewLine(l.linebuf[off0:]))
	}

	// Now run the mixer
	idx := (len(lm.layers) << 4) | ((lm.Cfg.LayerBpp - 1) << 2) | (lm.Cfg.ScreenBpp - 1)
	fastMixerTable[idx](lm, line)
}

// Begin drawing next line (possibly in background), onto the specified screen buffer
func (lm *LayerManager) BeginLine(line Line) {
	if lm.y < 0 {
		lm.setupWg.Wait()
	}
	lm.y++
	lm.drawLine(line)
}

// Wait for the current line to be fully drawn.
func (lm *LayerManager) EndLine() {

}

func (lm *LayerManager) EndFrame() {
	if lm.y != lm.Cfg.Height-1 {
		panic("end frame called before all lines")
	}
}

// Wrapper for rendering functions
type LayerFunc struct {
	Func func(lidx int) func(Line)
}

func (lf LayerFunc) DrawLayer(lidx int) func(Line) {
	return lf.Func(lidx)
}
