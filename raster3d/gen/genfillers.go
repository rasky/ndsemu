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

const FillerKeyBits = 3 + 1

type FillerConfig struct {
	TexFormat uint // 3 bits (0..7)
	ColorKey  bool // 1 bit
}

func (cfg *FillerConfig) Key() (k uint) {
	k |= (cfg.TexFormat & 7) << 0
	if cfg.ColorKey {
		k |= 1 << 3
	}
	return
}

func FillerConfigFromKey(k uint) (cfg FillerConfig) {
	cfg.TexFormat = k & 7
	if k&(1<<3) != 0 {
		cfg.ColorKey = true
	}
	return
}

type Generator struct {
	io.Writer
}

func (g *Generator) genFiller(cfg *FillerConfig) {
	fmt.Fprintf(g, "func (poly *RenderPolygon) filler_%02x(out Line) {\n", cfg.Key())
	fmt.Fprintf(g, "// %+v\n", *cfg)

	fmt.Fprintf(g, "}\n")
}

func (g *Generator) Run() {
	fmt.Fprintf(g, "// Generated on %v\n", time.Now())
	fmt.Fprintf(g, "package main\n")

	for i := uint(0); i < 1<<FillerKeyBits; i++ {
		cfg := FillerConfigFromKey(i)
		g.genFiller(&cfg)
	}

	fmt.Fprintf(g, "var polygonFillerTable = [%d]func(*RenderPolygon,Line) {\n",
		1<<FillerKeyBits)

	for i := uint(0); i < 1<<FillerKeyBits; i++ {
		fmt.Fprintf(g, "(*RenderPolygon).filler_%x,\n", i)
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
