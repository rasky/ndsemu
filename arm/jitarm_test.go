package arm

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"testing"

	a "github.com/rasky/gojit/amd64"
	"golang.org/x/arch/x86/x86asm"
)

type debugBus struct {
	Accesses []string
	RandData []uint32
}

func (d *debugBus) Read8(addr uint32) uint8 {
	d.Accesses = append(d.Accesses, fmt.Sprintf("R8:%08x", addr))
	return uint8(d.RandData[len(d.Accesses)])
}

func (d *debugBus) Read16(addr uint32) uint16 {
	d.Accesses = append(d.Accesses, fmt.Sprintf("R16:%08x", addr))
	return uint16(d.RandData[len(d.Accesses)])
}

func (d *debugBus) Read32(addr uint32) uint32 {
	d.Accesses = append(d.Accesses, fmt.Sprintf("R32:%08x", addr))
	return uint32(d.RandData[len(d.Accesses)])
}

func (d *debugBus) Write8(addr uint32, val uint8) {
	d.Accesses = append(d.Accesses, fmt.Sprintf("W8:%08x:%02x", addr, val))
}
func (d *debugBus) Write16(addr uint32, val uint16) {
	d.Accesses = append(d.Accesses, fmt.Sprintf("W16:%08x:%04x", addr, val))
}
func (d *debugBus) Write32(addr uint32, val uint32) {
	d.Accesses = append(d.Accesses, fmt.Sprintf("W32:%08x:%08x", addr, val))
}

func (d *debugBus) WaitStates() int {
	return 0xA
}
func (d *debugBus) FetchPointer(addr uint32) []byte {
	panic("unimplemented")
}

func TestAlu(t *testing.T) {
	jita, err := a.NewGoABI(1024 * 1024)
	if err != nil {
		t.Fatal(err)
	}

	bus1 := new(debugBus)
	bus2 := new(debugBus)
	bus1.RandData = make([]uint32, 0, 4096)
	for i := 0; i < 4096; i++ {
		bus1.RandData = append(bus1.RandData, rand.Uint32())
	}
	bus2.RandData = bus1.RandData

	var cpu1, cpu2 Cpu
	jit := &jitArm{jita, &cpu2}

	testf := func(op uint32, exp string) {
		var buf [4]byte
		binary.LittleEndian.PutUint32(buf[:], op)
		op = binary.BigEndian.Uint32(buf[:])

		jita.Off = 0
		f := jit.EmitBlock([]uint32{op})
		x86bin := jit.Buf[:jit.Off]

		t.Logf("Testing ARM Opcode:\t%08x  %s", op, exp)
		pc := uint64(0)
		for len(x86bin) > 0 {
			inst, err := x86asm.Decode(x86bin, 64)
			var text string
			size := inst.Len
			if err != nil || size == 0 || inst.Op == 0 {
				size = 1
				text = "?"
			} else {
				text = x86asm.GoSyntax(inst, pc, nil)
			}
			t.Logf("%04x  %s", pc, text)

			x86bin = x86bin[size:]
			pc += uint64(size)
		}

		for i := 0; i < 1024; i++ {
			var pre [16]reg

			// Generate random CPU state
			for j := 0; j < 16; j++ {
				pre[j] = reg(rand.Uint32())
				cpu1.Regs[j] = pre[j]
			}
			cpu1.Cpsr.r = reg(rand.Uint32() & 0xF0000000)
			cpu2 = cpu1

			// Reset bus monitor
			cpu1.bus = bus1
			cpu2.bus = bus2
			bus1.Accesses = nil
			bus2.Accesses = nil

			// Run interpreter over this instruction
			cpu1.Clock++
			if op >= 0xE0000000 || cpu1.opArmCond(uint(op>>28)) {
				opArmTable[(((op>>16)&0xFF0)|((op>>4)&0xF))&0xFFF](&cpu1, op)
			}

			// Run jit over the same instruction
			f(&cpu2)

			// Compare cpu1 and cpu2 regs
			for i := 0; i < 16; i++ {
				if cpu1.Regs[i] != cpu2.Regs[i] {
					t.Fatalf("R%d differs: exp:%v jit:%v", i, cpu1.Regs[i], cpu2.Regs[i])
				}
			}
			if cpu1.Cpsr != cpu2.Cpsr {
				t.Errorf("R0:%v R7:%v", pre[0], pre[7])
				t.Errorf("Cpsr differs: exp:%v jit:%v", cpu1.Cpsr, cpu2.Cpsr)
			}
			if cpu1.Clock != cpu2.Clock {
				t.Errorf("Clock differs: exp:%v jit:%v", cpu1.Clock, cpu2.Clock)
			}
			if len(bus1.Accesses) != len(bus2.Accesses) {
				t.Errorf("Different mem accesses: exp:%v jit:%v", bus1.Accesses, bus2.Accesses)
			} else {
				for i := range bus1.Accesses {
					if bus1.Accesses[i] != bus2.Accesses[i] {
						t.Errorf("Different mem accesses: exp:%v jit:%v", bus1.Accesses, bus2.Accesses)
						break
					}
				}
			}
		}
	}

	// ALU ------------------------------------------
	testf(0x01c3a0e3, "mov       r12, #0x4000000")
	testf(0xff0d80e2, "add       r0, r0, #0x3fc0")
	testf(0x04d040e2, "sub       sp, r0, #0x4")
	testf(0x04d04d02, "subeq     sp, sp, #0x4")
	testf(0x1f00a0e3, "mov       r0, #0x1f")
	testf(0xff14c1e3, "bic       r1, r1, #0xff000000")
	testf(0x011050e2, "subs      r1, r0, #0x1")
	testf(0x022080e0, "add       r2, r0, r2")
	testf(0x010053e1, "cmp       r3, r1")
	testf(0x3a77a0e1, "mov       r7, r10 lsr r7")
	testf(0x63a823e0, "eor       r10, r3, r3 ror #16")
	testf(0x6334a0e1, "mov       r3, r3 ror #8")
	testf(0x47729ae0, "adds      r7, r10, r7 asr #4")
	testf(0x70470000, "andeq     r4, r0, r0 ror r7")
	testf(0x70471000, "andeqs    r4, r0, r0 ror r7")

	// MEM ------------------------------------------
	testf(0x020081e7, "str       r0, [r1, r2])")
	testf(0x04a099e4, "ldr       r10, [r9], #0x4")
	testf(0x01b0d3e4, "ldrb      r11, [r3], #0x1")
	testf(0x0d10c0e5, "strb      r1, [r0, #0xd]")
	testf(0x01b0c0e4, "strb      r11, [r0], #0x1")
	testf(0x18a09be5, "ldr       r10, [r11, #0x18]")
	testf(0x08101ce5, "ldr       r1, [r12, #-0x8]")
	// testf(0x04f010e5, "ldr       pc, [r0, #-0x4]")

	// SWI ------------------------------------------
}
