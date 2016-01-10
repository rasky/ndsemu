package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

var filename = flag.String("filename", "-", "output filename")
var maxlayers = flag.Int("layers", 8, "max number of layers to unroll")

type Generator struct {
	io.Writer
}

func (g *Generator) genMixer(nl int, obpp, ibpp int) {
	fmt.Fprintf(g, "func (lm *LayerManager) fastmixer_%d_%d_%d(screen Line) {\n", nl, ibpp*8, obpp*8)
	fmt.Fprintf(g, "var inbuf [%d]uint32\n", nl)
	fmt.Fprintf(g, "in := inbuf[:]\n")

	fmt.Fprintf(g, "width := lm.Cfg.Width\n")
	fmt.Fprintf(g, "off0 := lm.Cfg.OverflowPixels * %d\n", ibpp)
	fmt.Fprintf(g, "mix := lm.Cfg.Mixer\n")
	fmt.Fprintf(g, "mixctx := lm.Cfg.MixerCtx\n")

	for i := 0; i < nl; i++ {
		fmt.Fprintf(g, "l%d := NewLine(lm.layers[lm.PriorityOrder[%d]].linebuf[off0:])\n", i, i)
	}

	fmt.Fprintf(g, "for x := 0; x < width; x++ {\n")
	for i := 0; i < nl; i++ {
		switch ibpp {
		case 1:
			fmt.Fprintf(g, "inbuf[%d] = uint32(l%d.Get8(x))\n", i, i)
		case 2:
			fmt.Fprintf(g, "inbuf[%d] = uint32(l%d.Get16(x))\n", i, i)
		case 4:
			fmt.Fprintf(g, "inbuf[%d] = l%d.Get32(x)\n", i, i)
		default:
			panic("unreachable")
		}
	}

	fmt.Fprintf(g, "out := mix(in, mixctx)\n")

	switch obpp {
	case 1:
		fmt.Fprintf(g, "screen.Set8(x, uint8(out))\n")
	case 2:
		fmt.Fprintf(g, "screen.Set16(x, uint16(out))\n")
	case 4:
		fmt.Fprintf(g, "screen.Set32(x, uint32(out))\n")
	default:
		panic("unimplemented")
	}

	fmt.Fprintf(g, "}\n")
	fmt.Fprintf(g, "}\n\n")
}

func (g *Generator) Run() {
	fmt.Fprintf(g, "// Generated on %v\n", time.Now())
	fmt.Fprintf(g, "package gfx\n")

	for nl := 1; nl < *maxlayers; nl++ {
		for _, obpp := range []int{1, 2, 4} {
			for _, ibpp := range []int{1, 2, 4} {
				g.genMixer(nl, obpp, ibpp)
			}
		}
	}

	fmt.Fprintf(g, "var fastMixerTable = [%d]func(*LayerManager,Line) {\n",
		*maxlayers*16)
	for i := 0; i < *maxlayers*16; i++ {
		nl := i / 16
		ibpp := ((i >> 2) & 0x3) + 1
		obpp := (i & 0x3) + 1

		if nl == 0 || ibpp == 3 || obpp == 3 {
			fmt.Fprintf(g, "nil,\n")
			continue
		}

		fmt.Fprintf(g, "(*LayerManager).fastmixer_%d_%d_%d,\n", nl, ibpp*8, obpp*8)
	}

	fmt.Fprintf(g, "}\n")
}

func main() {
	flag.Parse()

	var f io.Writer
	if *filename == "-" {
		f = os.Stdout
	} else {
		ff, err := os.Create(*filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer func() {
			if r := recover(); r != nil {
				panic(r)
			}
			cmd := exec.Command("go", "fmt", *filename)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				os.Exit(1)
			}
		}()
		defer ff.Close()
		f = ff
	}

	g := Generator{f}
	g.Run()
}
