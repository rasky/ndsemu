package cpugen

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Generator struct {
	io.Writer
	Prefix      string
	OpSize      string
	GenDisasm   bool
	PcRelOff    int
	Disasm      bytes.Buffer
	disasmDedup map[string]int
}

var filename = flag.String("filename", "-", "output filename")

func (g *Generator) WriteHeader() {
	fmt.Fprintf(g, "// Generated on %v\n", time.Now())
	fmt.Fprintf(g, "package arm\n")
	if g.GenDisasm {
		fmt.Fprintf(g, "import \"bytes\"\n")
		fmt.Fprintf(g, "import \"strconv\"\n")
	}

	fmt.Fprintf(g, "var op%sTable = [256]func(*Cpu, %s) {\n", g.Prefix, g.OpSize)
	for i := 0; i < 256; i++ {
		fmt.Fprintf(g, "(*Cpu).op%s%02X,\n", g.Prefix, i)
	}
	fmt.Fprintf(g, "}\n")

	if g.GenDisasm {
		fmt.Fprintf(g, "var disasm%sTable = [256]func(*Cpu, %s, uint32) string {\n", g.Prefix, g.OpSize)
		for i := 0; i < 256; i++ {
			fmt.Fprintf(g, "(*Cpu).disasm%s%02X,\n", g.Prefix, i)
		}
		fmt.Fprintf(g, "}\n")
	}

}

func (g *Generator) WriteFooter() {

}

func (g *Generator) WriteOpHeader(opnum int) {
	fmt.Fprintf(g, "func (cpu *Cpu) op%s%02X(op %s) {\n", g.Prefix, opnum, g.OpSize)
	g.Disasm.Reset()
}

func (g *Generator) WriteOpFooter(opnum int) {
	fmt.Fprintf(g, "}\n\n")
	if !g.GenDisasm {
		return
	}

	if g.Disasm.Len() == 0 {
		g.WriteDisasmInvalid()
	}
	if g.disasmDedup == nil {
		g.disasmDedup = make(map[string]int)
	}
	h := md5.Sum(g.Disasm.Bytes())
	hs := hex.EncodeToString(h[:])
	fmt.Fprintf(g, "func (cpu *Cpu) disasm%s%02X(op %s, pc uint32) string {\n", g.Prefix, opnum, g.OpSize)
	if opnum2, ok := g.disasmDedup[hs]; ok {
		fmt.Fprintf(g, "return cpu.disasm%s%02X(op,pc)\n", g.Prefix, opnum2)
	} else {
		fmt.Fprintf(g, g.Disasm.String())
		g.disasmDedup[hs] = opnum
	}
	fmt.Fprintf(g, "}\n\n")

}

func (g *Generator) WriteOpInvalid(msg string) {
	fmt.Fprintf(g, "cpu.InvalidOp%s(op, %q)\n", g.Prefix, msg)
}

func (g *Generator) WriteDisasmInvalid() {
	fmt.Fprint(&g.Disasm, "return \"dw \" + strconv.FormatInt(int64(op),16)\n")
}

func (g *Generator) WriteExitIfOpInvalid(cond string, msg string) {
	fmt.Fprintf(g, "if %s {\n", cond)
	g.WriteOpInvalid(msg)
	fmt.Fprintf(g, "return\n}\n")
}

func (g *Generator) WriteDisasm(opcode string, args ...string) {
	fmt.Fprintf(&g.Disasm, "var out bytes.Buffer\n")
	if opcode[0] == '!' {
		fmt.Fprintf(&g.Disasm, "opcode := %s\n", opcode[1:])
		fmt.Fprintf(&g.Disasm, "out.WriteString((opcode + \"                \")[:10])\n")
	} else {
		fmt.Fprintf(&g.Disasm, "out.WriteString(%q)\n", (opcode + "                ")[:10])
	}
	for i, a := range args {
		tmpname := "arg" + strconv.Itoa(i)

		switch a[0:2] {
		case "r:":
			// register
			wb := false
			if strings.HasSuffix(a, ":!") {
				wb = true
				a = a[:len(a)-2]
			}
			fmt.Fprintf(&g.Disasm, "%s:=%s\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%s])\n", tmpname)
			if wb {
				fmt.Fprintf(&g.Disasm, "out.WriteString(\"!\")\n")
			}
		case "R:":
			// register with possible writeback
			idx := strings.LastIndexByte(a, ':')
			fmt.Fprintf(&g.Disasm, "%sr:=%s\n", tmpname, a[2:idx])
			fmt.Fprintf(&g.Disasm, "%sw:=%s\n", tmpname, a[idx+1:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sr])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "if %sw!=0 { out.WriteString(\"!\") }\n", tmpname)
		case "d:":
			fmt.Fprintf(&g.Disasm, "%s:=int64(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"#\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(%s, 10))\n", tmpname)
		case "x:":
			fmt.Fprintf(&g.Disasm, "%s:=int64(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"#0x\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(%s, 16))\n", tmpname)
		case "l:":
			// one-register memory reference
			fmt.Fprintf(&g.Disasm, "%s:=%s\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"[\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%s])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"]\")\n")
		case "m:":
			// two-register memory reference
			wb := false
			if strings.HasSuffix(a, ":!") {
				wb = true
				a = a[:len(a)-2]
			}
			idx := strings.LastIndexByte(a, ':')
			fmt.Fprintf(&g.Disasm, "%sa:=%s\n", tmpname, a[2:idx])
			fmt.Fprintf(&g.Disasm, "%sb:=%s\n", tmpname, a[idx+1:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"[\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sa])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\", \")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sb])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"]\")\n")
			if wb {
				fmt.Fprintf(&g.Disasm, "out.WriteString(\"!\")\n")
			}
		case "n:":
			// register-imm memory reference
			wb := false
			if strings.HasSuffix(a, ":!") {
				wb = true
				a = a[:len(a)-2]
			}
			idx := strings.LastIndexByte(a, ':')
			fmt.Fprintf(&g.Disasm, "%sa:=%s\n", tmpname, a[2:idx])
			fmt.Fprintf(&g.Disasm, "%sb:=%s\n", tmpname, a[idx+1:])
			fmt.Fprintf(&g.Disasm, "if RegNames[%sa]==\"pc\" && !%v {\n", tmpname, wb)
			fmt.Fprintf(&g.Disasm, "%sc:=uint32(%sb)+uint32((pc+%d)&^2)\n", tmpname, tmpname, g.PcRelOff)
			fmt.Fprintf(&g.Disasm, "%sv:=cpu.opRead32(%sc)\n", tmpname, tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"= 0x\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(int64(%sv), 16))\n", tmpname)
			fmt.Fprintf(&g.Disasm, "} else {\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"[\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sa])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\", #0x\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(int64(%sb), 16))\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"]\")\n")
			if wb {
				fmt.Fprintf(&g.Disasm, "out.WriteString(\"!\")\n")
			}
			fmt.Fprintf(&g.Disasm, "}\n")
		case "N:":
			// register-string memory reference
			wb := false
			if strings.HasSuffix(a, ":!") {
				wb = true
				a = a[:len(a)-2]
			}
			idx := strings.LastIndexByte(a, ':')
			fmt.Fprintf(&g.Disasm, "%sa:=%s\n", tmpname, a[2:idx])
			fmt.Fprintf(&g.Disasm, "%sb:=%s\n", tmpname, a[idx+1:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"[\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sa])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\", \")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(%sb)\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"]\")\n")
			if wb {
				fmt.Fprintf(&g.Disasm, "out.WriteString(\"!\")\n")
			}
		case "P:":
			// PC-relative memory reference. This is treated different as we
			// can lookup the value from memory at runtime and show it instead
			// of the memory reference itself
			fmt.Fprintf(&g.Disasm, "%s:=uint32(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "%s+=uint32((pc+%d)&^2)\n", tmpname, g.PcRelOff)
			fmt.Fprintf(&g.Disasm, "%sv:=cpu.opRead32(%s)\n", tmpname, tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"= 0x\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(int64(%sv), 16))\n", tmpname)
		case "k:":
			// register bitmask
			fmt.Fprintf(&g.Disasm, "%s:=%s\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"{\")\n")
			fmt.Fprintf(&g.Disasm, "for i:=0;%s!=0;i++ {\n", tmpname)
			fmt.Fprintf(&g.Disasm, "  if %s&1 != 0 {\n", tmpname)
			fmt.Fprintf(&g.Disasm, "    out.WriteString(RegNames[i])\n")
			fmt.Fprintf(&g.Disasm, "    %s>>=1\n", tmpname)
			fmt.Fprintf(&g.Disasm, "    if %s != 0 { out.WriteString(\", \") }\n", tmpname)
			fmt.Fprintf(&g.Disasm, "  } else { \n")
			fmt.Fprintf(&g.Disasm, "    %s>>=1\n", tmpname)
			fmt.Fprintf(&g.Disasm, "  }\n")
			fmt.Fprintf(&g.Disasm, "}\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"}\")\n")
		case "o:":
			// PC offset (signed)
			fmt.Fprintf(&g.Disasm, "%s:=int32(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "%sx:=pc+%d+uint32(%s)\n", tmpname, g.PcRelOff, tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(int64(%sx), 16))\n", tmpname)
		case "s:":
			// the specified code produces the output
			fmt.Fprintf(&g.Disasm, "%s:=%s\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(%s)\n", tmpname)
		default:
			if a[1] == ':' {
				panic("invalid argument")
			}
			fmt.Fprintf(&g.Disasm, "out.WriteString(%q)\n", a)
		}

		if i < len(args)-1 {
			fmt.Fprintf(&g.Disasm, "out.WriteString(\", \")\n")
		}
	}
	fmt.Fprintf(&g.Disasm, "return out.String()\n")
}

func Main(do func(g *Generator)) {
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

	do(&Generator{
		Writer:    f,
		OpSize:    "uint32",
		GenDisasm: false,
	})
}
