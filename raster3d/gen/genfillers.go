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
	fmt.Fprintf(g, "func (e3d *HwEngine3d) filler_%03x(poly *Polygon, out gfx.Line, zbuf gfx.Line, abuf gfx.Line) {\n", cfg.Key())
	fmt.Fprintf(g, "// %+v\n", *cfg)

	if cfg.FillMode == fillerconfig.FillModeSolid && cfg.TexWithAlpha() {
		cfg.FillMode = fillerconfig.FillModeAlpha
	}

	fmt.Fprintf(g, "x0, x1 := poly.left[LerpX].Cur().NearInt32(), poly.right[LerpX].Cur().NearInt32()\n")
	fmt.Fprintf(g, "nx := x1-x0; if nx==0 {return}\n")
	fmt.Fprintf(g, "z0, z1 := poly.left[LerpZ].Cur(), poly.right[LerpZ].Cur()\n")
	fmt.Fprintf(g, "dz := z1.SubFixed(z0).Div(nx)\n")
	fmt.Fprintf(g, "c0, c1 := color(poly.left[LerpRGB].CurAsInt()), color(poly.right[LerpRGB].CurAsInt())\n")
	fmt.Fprintf(g, "dc := c1.SubColor(c0).Div(nx)\n")
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "texoff := poly.tex.VramTexOffset\n")
		fmt.Fprintf(g, "tshift := poly.tex.PitchShift\n")
		if cfg.Palettized() {
			fmt.Fprintf(g, "palette := e3d.palVram.Palette(int(poly.tex.VramPalOffset))\n")
		}
		fmt.Fprintf(g, "s0, s1 := poly.left[LerpS].Cur12(), poly.right[LerpS].Cur12()\n")
		fmt.Fprintf(g, "t0, t1 := poly.left[LerpT].Cur12(), poly.right[LerpT].Cur12()\n")
		fmt.Fprintf(g, "ds := s1.SubFixed(s0).Div(nx)\n")
		fmt.Fprintf(g, "dt := t1.SubFixed(t0).Div(nx)\n")
		fmt.Fprintf(g, "smask := poly.tex.SMask\n")
		fmt.Fprintf(g, "tmask := poly.tex.TMask\n")
	}
	if cfg.FillMode == fillerconfig.FillModeAlpha {
		fmt.Fprintf(g, "polyalpha := uint8(poly.flags.Alpha())<<1\n")
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
	case Tex4x4:
		fmt.Fprintf(g, "decompTexBuf := e3d.decompTex.Get(texoff)\n")
		fmt.Fprintf(g, "decompTex := gfx.NewLine(decompTexBuf)\n")
	}

	// Pixel loop var declarations
	fmt.Fprintf(g, "var px uint16\n")
	fmt.Fprintf(g, "var pxa uint8\n")
	fmt.Fprintf(g, "pxa = 63\n")
	fmt.Fprintf(g, "var px0 uint8\n")
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "var s,t uint32\n")
	}

	// **************************
	// PIXEL LOOP
	// **************************
	fmt.Fprintf(g, "out.Add32(int(x0))\n")
	fmt.Fprintf(g, "zbuf.Add32(int(x0))\n")
	if cfg.FillMode == fillerconfig.FillModeAlpha {
		fmt.Fprintf(g, "abuf.Add8(int(x0))\n")
	}
	fmt.Fprintf(g, "for x:=x0; x<x1; x++ {\n")

	// z-buffer check
	fmt.Fprintf(g, "// zbuffer check\n")
	fmt.Fprintf(g, "if z0.V >= int32(zbuf.Get32(0)) { goto next }\n")

	// texture fetch
	fmt.Fprintf(g, "// texel fetch\n")
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "s, t = uint32(s0.TruncInt32())&smask, uint32(t0.TruncInt32())&tmask\n")
		switch cfg.TexFormat {
		case Tex4:
			fmt.Fprintf(g, "px0 = e3d.texVram.Get8(texoff + t<<tshift + s/4)\n")
			fmt.Fprintf(g, "px0 = px0 >> (2 * uint(s&3))\n")
			fmt.Fprintf(g, "px0 &= 0x3\n")
			if cfg.ColorKey != 0 {
				fmt.Fprintf(g, "// color key check\n")
				fmt.Fprintf(g, "if px0 == 0 { goto next }\n")
			}
			fmt.Fprintf(g, "px = palette.Lookup(px0)\n")
		case Tex16:
			fmt.Fprintf(g, "px0 = e3d.texVram.Get8(texoff + t<<tshift + s/2)\n")
			fmt.Fprintf(g, "px0 = px0 >> (4 * uint(s&1))\n")
			fmt.Fprintf(g, "px0 &= 0xF\n")
			if cfg.ColorKey != 0 {
				fmt.Fprintf(g, "// color key check\n")
				fmt.Fprintf(g, "if px0 == 0 { goto next }\n")
			}
			fmt.Fprintf(g, "px = palette.Lookup(px0)\n")
		case Tex256:
			fmt.Fprintf(g, "px0 = e3d.texVram.Get8(texoff + t<<tshift + s)\n")
			if cfg.ColorKey != 0 {
				fmt.Fprintf(g, "if px0 == 0 { goto next }\n")
			}
			fmt.Fprintf(g, "px = palette.Lookup(px0)\n")
		case TexA3I5:
			fmt.Fprintf(g, "px0 = e3d.texVram.Get8(texoff + t<<tshift + s)\n")
			fmt.Fprintf(g, "pxa = (px0 >> 5)\n")
			fmt.Fprintf(g, "pxa = pxa | (pxa<<3)\n")
			fmt.Fprintf(g, "px0 &= 0x1F\n")
			if cfg.ColorKey != 0 {
				fmt.Fprintf(g, "// color key check\n")
				fmt.Fprintf(g, "if px0 == 0 { goto next }\n")
			}
			fmt.Fprintf(g, "px = uint16(px0)|uint16(px0)<<5|uint16(px0)<<10\n")
		case TexA5I3:
			fmt.Fprintf(g, "px0 = e3d.texVram.Get8(texoff + t<<tshift + s)\n")
			fmt.Fprintf(g, "pxa = px0 >> 3\n")
			fmt.Fprintf(g, "pxa = (pxa>>5) | (pxa<<1)\n")
			fmt.Fprintf(g, "px0 &= 0x7\n")
			fmt.Fprintf(g, "px0 <<= 2\n")
			if cfg.ColorKey != 0 {
				fmt.Fprintf(g, "// color key check\n")
				fmt.Fprintf(g, "if px0 == 0 { goto next }\n")
			}
			fmt.Fprintf(g, "px = uint16(px0)|uint16(px0)<<5|uint16(px0)<<10\n")
		case TexDirect:
			fmt.Fprintf(g, "px = e3d.texVram.Get16(texoff + t<<tshift + s*2)\n")
			fmt.Fprintf(g, "if px & 0x8000 != 0 { pxa = 63 }\n")
			fmt.Fprintf(g, "px &= 0x7FFF\n")
		case Tex4x4:
			fmt.Fprintf(g, "px = decompTex.Get16(int(t<<tshift + s))\n")
			// fmt.Fprintf(g, "px = emu.Read16LE(decompTexBuf[int(t<<tshift + s)*2:])\n")
			// Tex4x4 is always color-keyed
			if cfg.ColorKey != 0 {
				fmt.Fprintf(g, "// color key check\n")
				fmt.Fprintf(g, "if px == 0 { goto next }\n")
			}
		default:
			panic("unsupported")
		}
	}

	// color mode: combine texture pixel and vertex color
	fmt.Fprintf(g, "// apply vertex color to texel\n")
	switch cfg.ColorMode {
	case fillerconfig.ColorModeModulation:
		fmt.Fprintf(g, "if true {\n")
		fmt.Fprintf(g, "pxc := newColorFrom555U(px)\n")
		fmt.Fprintf(g, "pxc = pxc.Modulate(c0)\n")
		fmt.Fprintf(g, "px = pxc.To555U()\n")
		if cfg.FillMode == fillerconfig.FillModeAlpha {
			fmt.Fprintf(g, "pxa = uint8((int32(pxa+1)*int32(polyalpha+1)-1)>>6)\n")
		}
		fmt.Fprintf(g, "}\n")
	case fillerconfig.ColorModeDecal:
		fmt.Fprintf(g, "if true {\n")
		fmt.Fprintf(g, "pxc := newColorFrom555U(px)\n")
		fmt.Fprintf(g, "pxc = pxc.Decal(c0, pxa)\n")
		fmt.Fprintf(g, "px = pxc.To555U()\n")
		if cfg.FillMode == fillerconfig.FillModeAlpha {
			fmt.Fprintf(g, "pxa = polyalpha\n")
		}
		fmt.Fprintf(g, "}\n")
	case fillerconfig.ColorModeToon, fillerconfig.ColorModeHighlight:
		fmt.Fprintf(g, "if true {\n")
		fmt.Fprintf(g, "tc0 := emu.Read16LE(e3d.ToonTable.Data[((c0.R()>>1)&0x1F)*2:])\n")
		fmt.Fprintf(g, "tc :=  newColorFrom555U(tc0)\n")
		fmt.Fprintf(g, "pxc := newColorFrom555U(px)\n")
		fmt.Fprintf(g, "pxc = pxc.Modulate(tc)\n")
		if cfg.ColorMode == fillerconfig.ColorModeHighlight {
			fmt.Fprintf(g, "pxc = pxc.AddSat(tc)\n")
		}
		fmt.Fprintf(g, "px = pxc.To555U()\n")
		if cfg.FillMode == fillerconfig.FillModeAlpha {
			fmt.Fprintf(g, "pxa = uint8((int32(pxa+1)*int32(polyalpha+1)-1)>>6)\n")
		}
		fmt.Fprintf(g, "}\n")
	case fillerconfig.ColorModeShadow:
		fmt.Fprintf(g, "px = 0\n")
		if cfg.FillMode == fillerconfig.FillModeAlpha {
			fmt.Fprintf(g, "pxa = polyalpha\n")
		}
	}

	fmt.Fprintf(g, "// alpha blending with background\n")
	if cfg.FillMode == fillerconfig.FillModeAlpha {
		fmt.Fprintf(g, "if pxa == 0 { goto next }\n")
		fmt.Fprintf(g, "if true {\n")
		fmt.Fprintf(g, "bkg := uint16(out.Get32(0))\n")
		fmt.Fprintf(g, "bkga := abuf.Get8(0)\n")
		fmt.Fprintf(g, "if bkga != 0 { px = rgbAlphaMix(px, bkg, pxa>>1) }\n")
		fmt.Fprintf(g, "if pxa > bkga { abuf.Set8(0, pxa) }\n")
		fmt.Fprintf(g, "}\n")
	}

	// draw pixel
	fmt.Fprintf(g, "// draw color and z\n")
	fmt.Fprintf(g, "out.Set32(0, uint32(px)|0x80000000)\n")
	fmt.Fprintf(g, "zbuf.Set32(0, uint32(z0.V))\n")

	// Pixel loop footer
	fmt.Fprintf(g, "next:\n")
	fmt.Fprintf(g, "out.Add32(1)\n")
	fmt.Fprintf(g, "zbuf.Add32(1)\n")
	if cfg.FillMode == fillerconfig.FillModeAlpha {
		fmt.Fprintf(g, "abuf.Add8(1)\n")
	}
	fmt.Fprintf(g, "z0 = z0.AddFixed(dz)\n")
	fmt.Fprintf(g, "c0 = c0.AddDelta(dc)\n")
	if cfg.TexFormat > 0 {
		fmt.Fprintf(g, "s0 = s0.AddFixed(ds)\n")
		fmt.Fprintf(g, "t0 = t0.AddFixed(dt)\n")
	}

	fmt.Fprintf(g, "}\n")
	fmt.Fprintf(g, "_=px0\n")
	fmt.Fprintf(g, "_=pxa\n")
	fmt.Fprintf(g, "}\n")
}

func (g *Generator) Run() {
	fmt.Fprintf(g, "// Generated on %v\n", time.Now())
	fmt.Fprintf(g, "package raster3d\n")
	fmt.Fprintf(g, "import \"ndsemu/emu/gfx\"\n")
	fmt.Fprintf(g, "import \"ndsemu/emu\"\n")

	for i := uint(0); i < fillerconfig.FillerKeyMax; i++ {
		cfg := fillerconfig.FillerConfigFromKey(i)
		g.genFiller(&cfg)
	}

	fmt.Fprintf(g, "var polygonFillerTable = [%d]func(*HwEngine3d,*Polygon,gfx.Line,gfx.Line,gfx.Line) {\n",
		fillerconfig.FillerKeyMax)

	for i := uint(0); i < fillerconfig.FillerKeyMax; i++ {
		fmt.Fprintf(g, "(*HwEngine3d).filler_%03x,\n", i)
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
