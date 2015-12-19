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
	Disasm      bytes.Buffer
	disasmDedup map[string]uint16
}

var filename = flag.String("filename", "-", "output filename")

func (g *Generator) WriteHeader() {
	fmt.Fprintf(g, "// Generated on %v\n", time.Now())
	fmt.Fprintf(g, "package arm\n")
	fmt.Fprintf(g, "import \"bytes\"\n")
	fmt.Fprintf(g, "import \"strconv\"\n")

	fmt.Fprintf(g, "var opThumbTable = [256]func(*Cpu, uint16) {\n")
	for i := 0; i < 256; i++ {
		fmt.Fprintf(g, "(*Cpu).opThumb%02X,\n", i)
	}
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(g, "var disasmThumbTable = [256]func(*Cpu, uint16, uint32) string {\n")
	for i := 0; i < 256; i++ {
		fmt.Fprintf(g, "(*Cpu).disasmThumb%02X,\n", i)
	}
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(g, "var opThumbAluTable = [16]func(*Cpu, uint16) {\n")
	for i := 0; i < 16; i++ {
		fmt.Fprintf(g, "(*Cpu).opThumbAlu%02X,\n", i)
	}
	fmt.Fprintf(g, "}\n")

	fmt.Fprintf(g, "var disasmThumbAluTable = [16]func(*Cpu, uint16, uint32) string {\n")
	for i := 0; i < 16; i++ {
		fmt.Fprintf(g, "(*Cpu).disasmThumbAlu%02X,\n", i)
	}
	fmt.Fprintf(g, "}\n")
}

func (g *Generator) WriteFooter() {

}

func (g *Generator) WriteOpHeader(op uint16) {
	fmt.Fprintf(g, "func (cpu *Cpu) opThumb%02X(op uint16) {\n", (op>>8)&0xFF)
	g.Disasm.Reset()
}
func (g *Generator) WriteOpFooter(op uint16) {
	fmt.Fprintf(g, "}\n\n")
	if g.Disasm.Len() == 0 {
		// panic(fmt.Sprintf("disasm not implemented for op %04x", op))
		return
	}
	if g.disasmDedup == nil {
		g.disasmDedup = make(map[string]uint16)
	}
	h := md5.Sum(g.Disasm.Bytes())
	hs := hex.EncodeToString(h[:])
	fmt.Fprintf(g, "func (cpu *Cpu) disasmThumb%02X(op uint16, pc uint32) string {\n", (op>>8)&0xFF)
	if op2, ok := g.disasmDedup[hs]; ok {
		fmt.Fprintf(g, "return cpu.disasmThumb%02X(op,pc)\n", (op2>>8)&0xFF)
	} else {
		fmt.Fprintf(g, g.Disasm.String())
		g.disasmDedup[hs] = op
	}
	fmt.Fprintf(g, "}\n\n")

}

func (g *Generator) WriteOpInvalid(op uint16, msg string) {
	fmt.Fprintf(g, "cpu.InvalidOpThumb(op, %q)\n", msg)
}

func (g *Generator) WriteDisasmInvalid() {
	fmt.Fprint(&g.Disasm, "return \"dw \" + strconv.FormatInt(int64(op),16)\n")
}

func (g *Generator) WriteExitIfOpInvalid(cond string, op uint16, msg string) {
	fmt.Fprintf(g, "if %s {\n", cond)
	g.WriteOpInvalid(op, msg)
	fmt.Fprintf(g, "return\n}\n")
}

func (g *Generator) WriteDisasm(opcode string, args ...string) {
	fmt.Fprintf(&g.Disasm, "var out bytes.Buffer\n")
	fmt.Fprintf(&g.Disasm, "out.WriteString(%q)\n", (opcode + "                ")[:10])
	for i, a := range args {
		tmpname := "arg" + strconv.Itoa(i)

		switch a[0] {
		case 'r':
			// register
			fmt.Fprintf(&g.Disasm, "%s:=%s\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%s])\n", tmpname)
		case 'R':
			// register with possible writeback
			idx := strings.LastIndexByte(a, ':')
			fmt.Fprintf(&g.Disasm, "%sr:=%s\n", tmpname, a[2:idx])
			fmt.Fprintf(&g.Disasm, "%sw:=%s\n", tmpname, a[idx+1:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sr])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "if %sw!=0 { out.WriteString(\"!\") }\n", tmpname)
		case 'd':
			fmt.Fprintf(&g.Disasm, "%s:=int64(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"#\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(%s, 10))\n", tmpname)
		case 'x':
			fmt.Fprintf(&g.Disasm, "%s:=int64(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"#0x\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(%s, 16))\n", tmpname)
		case 'm':
			// two-register memory reference
			idx := strings.LastIndexByte(a, ':')
			fmt.Fprintf(&g.Disasm, "%sa:=%s\n", tmpname, a[2:idx])
			fmt.Fprintf(&g.Disasm, "%sb:=%s\n", tmpname, a[idx+1:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"[\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sa])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\", \")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sb])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"]\")\n")
		case 'n':
			// register-imm memory reference
			idx := strings.LastIndexByte(a, ':')
			fmt.Fprintf(&g.Disasm, "%sa:=%s\n", tmpname, a[2:idx])
			fmt.Fprintf(&g.Disasm, "%sb:=%s\n", tmpname, a[idx+1:])
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"[\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(RegNames[%sa])\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\", #0x\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(int64(%sb), 16))\n", tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"]\")\n")
		case 'P':
			// PC-relative memory reference. This is treated different as we
			// can lookup the value from memory at runtime and show it instead
			// of the memory reference itself
			fmt.Fprintf(&g.Disasm, "%s:=uint32(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "%s+=uint32((pc+4)&^2)\n", tmpname)
			fmt.Fprintf(&g.Disasm, "%sv:=cpu.opRead32(%s)\n", tmpname, tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(\"= 0x\")\n")
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(int64(%sv), 16))\n", tmpname)
		case 'k':
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
		case 'o':
			// PC offset (signed)
			fmt.Fprintf(&g.Disasm, "%s:=int32(%s)\n", tmpname, a[2:])
			fmt.Fprintf(&g.Disasm, "%sx:=pc+4+uint32(%s)\n", tmpname, tmpname)
			fmt.Fprintf(&g.Disasm, "out.WriteString(strconv.FormatInt(int64(%sx), 16))\n", tmpname)
		default:
			panic("invalid argument")
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

	do(&Generator{Writer: f})
}
