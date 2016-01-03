package gfx

import "sync"

//go:generate go run genmixer/genmixer.go -filename fastmixer.go

type LayerCtx struct {
	endLineCh  chan bool
	nextLineCh chan Line
	restart    bool
}

func (ctx *LayerCtx) NextLine() Line {
	ctx.endLineCh <- true
	return <-ctx.nextLineCh
}

func (ctx *LayerCtx) waitReady() {
	v := <-ctx.endLineCh
	if !v {
		panic("layer process exited while waiting for it to be ready")
	}
}

func (ctx *LayerCtx) waitDead() {
	v := <-ctx.endLineCh
	if v {
		panic("layer process was ready while waiting for it to die")
	}
}

type Layer interface {
	// DrawFrame is the entry point of the drawing code for each layer. It receives
	// as input a LayerCtx object, the index of the layer itself in the layer manager,
	// and the y coordinate of the first line to be drawn.
	//
	// The function must do the initial setup, and then call ctx.NextLine() when
	// ready; this function will return a gfx.Line buffer where the line must be
	// drawn, or a null line if the function must exit immediately. This is an
	// example of the skeleton of a correct DrawFrame implementation:
	//
	//  func (l *MyLayer) DrawFrame(ctx *gfx.LayerCtx, idx int, y int) {
	//
	//      [initial setup...]
	//
	//      for {
	//          out := ctx.NextLine()
	//          if out.IsNil() {
	//              return
	//          }
	//
	//          [ draw line y into out ]
	//
	//          y++
	//      }
	//  }
	//
	// layerIdx is passed as a way to reuse the same code for multiple layers. If
	// you have a different codepath for each layer, then you can safely ignore this
	// argument.
	//
	DrawLayer(ctx *LayerCtx, layerIdx int, y int)
}

type layerData struct {
	Layer
	ctx LayerCtx
	buf []byte
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
	// layers in the layer manager
	Mixer func([]uint32) uint32
}

type LayerManager struct {
	Cfg     LayerManagerConfig
	layers  []*layerData
	y       int
	setupWg sync.WaitGroup
	lineWg  sync.WaitGroup
}

func (lm *LayerManager) AddLayer(l Layer) {
	lm.layers = append(lm.layers, &layerData{
		Layer: l,
		buf:   make([]byte, (lm.Cfg.Width+lm.Cfg.OverflowPixels*2)*lm.Cfg.LayerBpp),
		ctx: LayerCtx{
			endLineCh:  make(chan bool, 1),
			nextLineCh: make(chan Line, 1),
		},
	})
}

func (lm *LayerManager) BeginFrame() {
	lm.y = -1

	lm.setupWg.Add(1)
	go func() {
		lm.setupFrame()
		lm.setupWg.Done()
	}()
}

func (lm *LayerManager) setupFrame() {
	for idx, l := range lm.layers {
		go func(l *layerData, idx int) {
			l.DrawLayer(&l.ctx, idx, 0)
			l.ctx.endLineCh <- false
		}(l, idx)
	}

	// Wait for all layers to have finished their initial setup
	for _, l := range lm.layers {
		<-l.ctx.endLineCh
	}
}

func (lm *LayerManager) drawLine(line Line) {
	off0 := lm.Cfg.OverflowPixels * lm.Cfg.LayerBpp

	// Send new line to each layer
	for idx, l := range lm.layers {
		if l.ctx.restart {
			// Restart drawing on this layer (from the current line)
			l.ctx.restart = false
			l.ctx.nextLineCh <- Line{0}
			go l.DrawLayer(&l.ctx, idx, lm.y)
			<-l.ctx.endLineCh
		}

		l.ctx.nextLineCh <- NewLine(l.buf[off0:])
	}

	// Wait for each layer to finish its current line
	for _, l := range lm.layers {
		l.ctx.waitReady()
	}

	// Now run the mixer
	idx := (len(lm.layers) << 4) | ((lm.Cfg.LayerBpp - 1) << 2) | (lm.Cfg.ScreenBpp - 1)
	fastMixerTable[idx](lm, line)
}

// Begin drawing next line in background, onto the specified screen buffer
func (lm *LayerManager) BeginLine(line Line) {
	if lm.y < 0 {
		lm.setupWg.Wait()
	}
	lm.y++
	lm.lineWg.Add(1)
	go func() {
		lm.drawLine(line)
		lm.lineWg.Done()
	}()
}

// Wait for the current line to be fully drawn.
func (lm *LayerManager) EndLine() {
	lm.lineWg.Wait()
}

// Force a restart of the draw routine of a layer. After calling this function,
// on the next line, the specified layer will receive a nil line object as
// return of its ctx.NextLine() call, and then the Layer.DrawFrame() function
// will be restarted (from the correct line)
func (lm *LayerManager) RestartDraw(layerIdx int) {
	lm.layers[layerIdx].ctx.restart = true
}

func (lm *LayerManager) EndFrame() {
	if lm.y != lm.Cfg.Height-1 {
		panic("end frame called before all lines")
	}

	for _, l := range lm.layers {
		l.ctx.nextLineCh <- Line{0}
		l.ctx.waitDead()
	}
}
