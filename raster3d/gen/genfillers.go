package main

import (
	"flag"
	"fmt"
	"io"
	"ndsemu/raster3d/fillerconfig"
	"os"
	"os/exec"
	"time"
)

var filename = flag.String("filename", "-", "output filename")
var maxlayers = flag.Int("layers", 8, "max number of layers to unroll")

const FillerKeyBits = 3 + 1

const (
	TexNone uint = iota
	TexA3I5
	Tex4
	Tex16
	Tex256
	Tex4x4
	TexA5I3
	TexDirect
)

type Generator struct {
	io.Writer
}

func (g *Generator) genFiller(cfg *fillerconfig.FillerConfig) {
	fmt.Fprintf(g, "func (e3d *HwEngine3d) filler_%02x(poly *Polygon, out gfx.Line, zbuf gfx.Line) {\n", cfg.Key())
	fmt.Fprintf(g, "// %+v\n", *cfg)

	fmt.Fprintf(g, "x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()\n")
	fmt.Fprintf(g, "nx := x1-x0; if nx==0 {return}\n")
	fmt.Fprintf(g, "z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()\n")
	fmt.Fprintf(g, "dz := z1.SubFixed(z0).Div(nx)\n")
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "texoff := poly.tex.VramTexOffset\n")
		fmt.Fprintf(g, "tshift := poly.tex.PitchShift\n")
		if cfg.Palettized() {
			fmt.Fprintf(g, "palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))\n")
		}
		fmt.Fprintf(g, "s0, s1 := poly.left[LerpS].Cur(), poly.right[LerpS].Cur()\n")
		fmt.Fprintf(g, "t0, t1 := poly.left[LerpT].Cur(), poly.right[LerpT].Cur()\n")
		fmt.Fprintf(g, "ds := s1.SubFixed(s0).Div(nx)\n")
		fmt.Fprintf(g, "dt := t1.SubFixed(t0).Div(nx)\n")
		fmt.Fprintf(g, "smask := poly.tex.SMask\n")
		fmt.Fprintf(g, "tmask := poly.tex.TMask\n")
	}

	// Pre pixel loop
	switch cfg.TexFormat {
	case Tex4:
		// 4 pixels per byte, decrease texture pitch
		fmt.Fprintf(g, "tshift -= 2\n")
	case Tex16:
		// 2 pixels per byte, decrease texture pitch
		fmt.Fprintf(g, "tshift -= 1\n")
	case TexDirect:
		// 2 bytes per pixel, increase texture pitch
		fmt.Fprintf(g, "tshift += 1\n")
	}

	// Pixel loop var declarations
	if cfg.TexFormat == TexDirect {
		fmt.Fprintf(g, "var px uint16\n")
	} else {
		fmt.Fprintf(g, "var px uint8\n")
	}
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "var s,t uint32\n")
	}

	// **************************
	// PIXEL LOOP
	// **************************
	fmt.Fprintf(g, "out.Add16(int(x0))\n")
	fmt.Fprintf(g, "for x:=x0; x<x1; x++ {\n")

	// z-buffer check
	// fmt.Fprintf(g, "if z0.V >= int32(zbuf.Get32(0)) { goto next }\n")
	fmt.Fprintf(g, "if false { goto next }\n")

	// texture fetch
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask\n")
		switch cfg.TexFormat {
		case Tex4:
			fmt.Fprintf(g, "px = e3d.texVram.Get8(texoff + t<<tshift + s/4)\n")
			fmt.Fprintf(g, "px = px >> (2 * uint(s&3))\n")
			fmt.Fprintf(g, "px &= 0x3\n")
		case Tex16:
			fmt.Fprintf(g, "px = e3d.texVram.Get8(texoff + t<<tshift + s/2)\n")
			fmt.Fprintf(g, "px = px >> (4 * uint(s&1))\n")
			fmt.Fprintf(g, "px &= 0xF\n")
		case Tex256:
			fmt.Fprintf(g, "px = e3d.texVram.Get8(texoff + t<<tshift + s)\n")
		case TexA3I5:
			// FIXME: add alpha blending
			fmt.Fprintf(g, "px = e3d.texVram.Get8(texoff + t<<tshift + s)\n")
			fmt.Fprintf(g, "px &= 0x1F\n")
		case TexA5I3:
			// FIXME: add alpha blending
			fmt.Fprintf(g, "px = e3d.texVram.Get8(texoff + t<<tshift + s)\n")
			fmt.Fprintf(g, "px &= 0x7\n")
		case TexDirect:
			fmt.Fprintf(g, "px = e3d.texVram.Get16(texoff + t<<tshift + s)\n")
		default:
			fmt.Fprintf(g, "px = e3d.texVram.Get8(texoff + t<<tshift + s)\n")
		}
	}

	// color-key check
	if cfg.ColorKey && cfg.Palettized() {
		fmt.Fprintf(g, "if px == 0 { goto next }\n")
	}

	// draw pixel
	if cfg.Palettized() {
		fmt.Fprintf(g, "out.Set16(0, palette.Lookup(px)|0x8000)\n")
	} else {
		fmt.Fprintf(g, "out.Set16(0, uint16(px)|0x8000)\n")
	}
	fmt.Fprintf(g, "zbuf.Set32(0, uint32(z0.V))\n")

	// Pixel loop footer
	fmt.Fprintf(g, "next:\n")
	fmt.Fprintf(g, "out.Add16(1)\n")
	fmt.Fprintf(g, "zbuf.Add32(1)\n")
	fmt.Fprintf(g, "z0 = z0.AddFixed(dz)\n")
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "s0 = s0.AddFixed(ds)\n")
		fmt.Fprintf(g, "t0 = t0.AddFixed(dt)\n")
	}

	fmt.Fprintf(g, "}\n")
	fmt.Fprintf(g, "}\n")
}

func (g *Generator) Run() {
	fmt.Fprintf(g, "// Generated on %v\n", time.Now())
	fmt.Fprintf(g, "package raster3d\n")
	fmt.Fprintf(g, "import \"ndsemu/emu/gfx\"\n")

	for i := uint(0); i < 1<<FillerKeyBits; i++ {
		cfg := fillerconfig.FillerConfigFromKey(i)
		g.genFiller(&cfg)
	}

	fmt.Fprintf(g, "var polygonFillerTable = [%d]func(*HwEngine3d,*Polygon,gfx.Line,gfx.Line) {\n",
		1<<FillerKeyBits)

	for i := uint(0); i < 1<<FillerKeyBits; i++ {
		fmt.Fprintf(g, "(*HwEngine3d).filler_%02x,\n", i)
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
