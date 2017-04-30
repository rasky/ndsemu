package arm

import (
	"encoding/binary"
	"fmt"
	"math/rand"
	"runtime/debug"
	"testing"
	"time"

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

func (d *debugBus) Read(op uint32, cn, cm, cp uint32) uint32 {
	d.Accesses = append(d.Accesses, fmt.Sprintf("COPREAD:%08x:%d:%d:%d", op, cn, cm, cp))
	return uint32(d.RandData[len(d.Accesses)])
}

func (d *debugBus) Write(op uint32, cn, cm, cp uint32, value uint32) {
	d.Accesses = append(d.Accesses, fmt.Sprintf("COPWRITE:%08x:%d:%d:%d:%08x", op, cn, cm, cp, value))
}

func (d *debugBus) Exec(op uint32, cn, cm, cp uint32, value uint32) {
	d.Accesses = append(d.Accesses, fmt.Sprintf("COPEXEC:%08x:%d:%d:%d:%08x", op, cn, cm, cp, value))
}

var specials = []uint32{
	0x0, 0x0, 0xFFFFFFFF, 0xFFFFFFFE,
	0x1, 0x2, 0x80000000, 0x80000001,
}

func randSpecials() uint32 {
	// Return "special values" which are more likely to trigger bugs on
	// edge conditions
	return specials[rand.Uint32()&3]
}

func TestAlu(t *testing.T) {
	debug.SetGCPercent(-1) // Disable GC for now

	jita, err := a.NewGoABI(1024 * 1024)
	if err != nil {
		t.Fatal(err)
	}

	bus1 := new(debugBus)
	bus2 := new(debugBus)

	var cpu1, cpu2 Cpu
	jit := &jitArm{Assembler: jita, Cpu: &cpu2}

	var total []uint32

	testf1 := func(op uint32, exp string, mod func(*Cpu)) {
		var buf [4]byte
		binary.LittleEndian.PutUint32(buf[:], op)
		op = binary.BigEndian.Uint32(buf[:])
		total = append(total, op)

		jita.Off = 0
		f := jit.EmitBlock([]uint32{op})
		x86bin := jit.Buf[:jit.Off]

		t.Logf("Testing ARM Opcode (ARMv%d):\t%08x  %s", cpu1.arch, op, exp)
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
			t.Logf("%04x %-28x %s", pc, x86bin[:size], text)

			x86bin = x86bin[size:]
			pc += uint64(size)
		}

		for i := 0; i < 1024; i++ {
			var pre [16]reg

			randf := rand.Uint32
			if i < 128 {
				randf = randSpecials
			}

			// Generate random CPU state
			for j := 0; j < 16; j++ {
				cpu1.Regs[j] = reg(randf())
			}
			cpu1.Regs[15] &^= 3 // PC must be aligned
			cpu1.pc = reg(rand.Uint32()) &^ 3
			cpu1.Cpsr.r = (reg(rand.Uint32()) & 0xF0000000) | reg(CpuModeUser)
			cpu1.Clock = 0

			// Generate new random data
			bus1.RandData = make([]uint32, 0, 16)
			for i := 0; i < 16; i++ {
				bus1.RandData = append(bus1.RandData, rand.Uint32())
			}

			// Reset bus monitor
			cpu1.bus = bus1
			for i := 0; i < 16; i++ {
				cpu1.MapCoprocessor(i, bus1)
			}
			bus1.Accesses = nil

			// Test-specific modifications
			if mod != nil {
				mod(&cpu1)
			}

			// Save for debug
			pre = cpu1.Regs

			// Copy into second CPU for comparison
			cpu2 = cpu1
			cpu2.bus = bus2
			for i := 0; i < 16; i++ {
				cpu2.MapCoprocessor(i, bus2)
			}
			bus2.Accesses = nil
			bus2.RandData = bus1.RandData

			// Run interpreter over this instruction
			cpu1.Clock++
			if op >= 0xE0000000 || cpu1.opArmCond(uint(op>>28)) {
				opArmTable[(((op>>16)&0xFF0)|((op>>4)&0xF))&0xFFF](&cpu1, op)
			}

			// fmt.Println(disasmArmTable[((op>>16)&0xFF0)|((op>>4)&0xF)](&cpu1, op, 0))

			// Run jit over the same instruction
			f(&cpu2)

			// Compare cpu1 and cpu2 regs
			for i := 0; i < 16; i++ {
				if cpu1.Regs[i] != cpu2.Regs[i] {
					t.Fatalf("R%d differs: exp:%v jit:%v", i, cpu1.Regs[i], cpu2.Regs[i])
				}
			}
			if cpu1.Cpsr != cpu2.Cpsr {
				t.Fatalf("Cpsr differs: exp:%v jit:%v", cpu1.Cpsr, cpu2.Cpsr)
			}

			for i := 0; i < 5; i++ {
				if cpu1.SpsrBank[i] != cpu2.SpsrBank[i] {
					t.Errorf("Spsr[%d] differs: exp:%v jit:%v", i, cpu1.SpsrBank[i], cpu2.SpsrBank[i])
				}
			}
			if cpu1.pc != cpu2.pc {
				t.Errorf("pc differs: exp:%v jit:%v", cpu1.pc, cpu2.pc)
			}
			if cpu1.Clock != cpu2.Clock {
				t.Errorf("Clock differs: exp:%v jit:%v", cpu1.Clock, cpu2.Clock)
			}
			if len(bus1.Accesses) != len(bus2.Accesses) {
				t.Errorf("Different IO accesses: exp:%v jit:%v", bus1.Accesses, bus2.Accesses)
			} else {
				for i := range bus1.Accesses {
					if bus1.Accesses[i] != bus2.Accesses[i] {
						t.Errorf("Different IO accesses: exp:%v jit:%v", bus1.Accesses, bus2.Accesses)
						break
					}
				}
			}
			_ = pre
		}
	}

	testf := func(op uint32, exp string) {
		testf1(op, exp, nil)
	}

	for _, a := range []Arch{ARMv4, ARMv5} {
		cpu1.arch = a
		cpu2.arch = a

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
		testf(0x1fb08be3, "orr       r11, r11, #0x1f")
		testf(0x63a823e0, "eor       r10, r3, r3 ror #16")
		testf(0x6334a0e1, "mov       r3, r3 ror #8")
		testf(0x0140c3e0, "sbc       r4, r3, r1")
		testf(0x0000d4e2, "sbcs      r0, r4, #0x0")
		testf(0x47729ae0, "adds      r7, r10, r7 asr #4")
		testf(0x70470000, "andeq     r4, r0, r0 ror r7")
		testf(0x70471000, "andeqs    r4, r0, r0 ror r7")
		testf(0x8330b1e0, "adcs      r3, r1, r3 lsl #1")
		testf(0x01304330, "sublo     r3, r3, r1")
		testf(0x00106112, "rsbne     r1, r1, #0x0")
		testf(0x02311ce2, "ands      r3, r12, #0x80000000")
		testf(0x48b08fe2, "add       r11, pc, #0x48")
		testf(0xe0061c13, "tstne     r12, #0xe000000")
		testf(0x01003ce3, "teq       r12, #0x1")
		testf(0x310213e1, "tst       r3, r1 lsr r2")
		testf(0x18007ce3, "cmn       r12, #0x18")
		testf(0x132011e0, "ands      r2, r1, r3 lsl r0")
		testf(0x5b5ab0e1, "movs      r5, r11 asr r10")
		testf(0x47729ae0, "adds      r7, r10, r7 asr #4")
		testf(0x6880ff01, "mvnseq    r8, r8 rrx #1")
		testf(0x9363f5e2, "rscs      r6, r5, #0x4c000002")
		testf(0x02f18fe0, "add       pc, pc, r2 lsl #2")
		testf1(0x04f05ee2, "subs      pc, lr, #0x4", func(cpu *Cpu) {
			cpu.Cpsr.r = (reg(rand.Uint32()) & 0xF0000000) | reg(CpuModeSupervisor)
			// Random spsr in user mode, with random T bit to check address masking.
			// Also random LR with bit 0 set, so that we can check it's being cleared
			*cpu.RegSpsr() = (reg(rand.Uint32()) & 0xF0000020) | reg(CpuModeUser)
			cpu.Regs[14] &^= 2

		})
		testf1(0x0ef0b0e1, "movs      pc, lr", func(cpu *Cpu) {
			cpu.Cpsr.r = (reg(rand.Uint32()) & 0xF0000000) | reg(CpuModeSupervisor)
			// Random spsr in user mode, with random T bit to check address masking.
			// Also random LR with bit 0 set, so that we can check it's being cleared
			*cpu.RegSpsr() = (reg(rand.Uint32()) & 0xF0000020) | reg(CpuModeUser)
			cpu.Regs[14] &^= 2
		})

		// MEM ------------------------------------------
		testf(0x020081e7, "str       r0, [r1, r2])")
		testf(0x04a099e4, "ldr       r10, [r9], #0x4")
		testf(0x01b0d3e4, "ldrb      r11, [r3], #0x1")
		testf(0x0d10c0e5, "strb      r1, [r0, #0xd]")
		testf(0x01b0c0e4, "strb      r11, [r0], #0x1")
		testf(0x18a09be5, "ldr       r10, [r11, #0x18]")
		testf(0x08101ce5, "ldr       r1, [r12, #-0x8]")
		testf(0x0cc19be7, "ldr       r12, [r11, r12 lsl #2]")
		testf(0x3c309fb5, "ldrlt     r3, = 0xfffe")
		testf(0x010073e5, "ldrb      r0, [r3, #0x-1]!")
		testf(0x010062e5, "strb      r0, [r2, #0x-1]!")
		testf(0x09a053e7, "ldrb      r10, [r3, -r9]")
		testf(0x33ee31d5, "ldrle     lr, [r1, -#0xe33]!")
		testf(0x33ee31c5, "ldrgt     lr, [r1, -#0xe33]!")
		testf1(0x0080bde8, "ldm       sp!, {pc}", func(cpu *Cpu) {
			bus := cpu.bus.(*debugBus)
			// Turn off second bit that will be returned by LDR, because
			// it might cause a jump to a not aligned address. Bit 0 is used
			// to switch to Thumb, so keep it.
			for i := range bus.RandData {
				bus.RandData[i] &^= 2
				if cpu.arch < ARMv5 {
					// ARM4 doesn't support switching to thumb in LDM
					bus.RandData[i] &^= 1
				}
			}
		})
		testf1(0xff7fd0e8, "ldm       r0, {r0, r1, r2, r3, r4, r5, r6, r7, r8…", func(cpu *Cpu) {
			bus := cpu.bus.(*debugBus)
			// Turn off second bit that will be returned by LDR, because
			// it might cause a jump to a not aligned address. Bit 0 is used
			// to switch to Thumb, so keep it.
			for i := range bus.RandData {
				bus.RandData[i] &^= 2
				if cpu.arch < ARMv5 {
					// ARM4 doesn't support switching to thumb in LDM
					bus.RandData[i] &^= 1
				}
			}
		})
		testf1(0x04f010e5, "ldr       pc, [r0, #-0x4]", func(cpu *Cpu) {
			bus := cpu.bus.(*debugBus)
			// Turn off second bit that will be returned by LDR, because
			// it might cause a jump to a not aligned address. Bit 0 is used
			// to switch to Thumb, so keep it.
			for i := range bus.RandData {
				bus.RandData[i] &^= 2
			}
		})

		// SWI ------------------------------------------

		// SWP ------------------------------------------
		testf(0x9aa043e1, "swpb      r10, r10, [r3]")
		testf(0x900001e1, "swp       r0, r0, [r1]")

		// CLZ ------------------------------------------
		if cpu1.arch >= 5 {
			testf(0x112f6fe1, "clz       r2, r1")
		}

		// MUL ------------------------------------------
		testf(0x9a0b00e0, "mul       r0, r10, r11")
		testf(0x910016e0, "muls      r6, r1, r0")
		testf(0x950124e0, "mla       r4, r5, r1, r0")
		testf(0x950134e0, "mlas      r4, r5, r1, r0")
		testf(0x953281e0, "umull     r3, r1, r5, r2")
		testf(0x9584c4e0, "smull     r8, r4, r5, r4")
		testf(0x9584d4e0, "smulls    r8, r4, r5, r4")
		testf(0x9363e5e0, "smlal     r6, r5, r3, r3")
		testf(0x9363f5e0, "smlals    r6, r5, r3, r3")
		if cpu1.arch >= 5 {
			testf(0x843a03e1, "smlabb    r3, r4, r10, r3")
			testf(0x860767e1, "smulbb    r7, r6, r7")
			testf(0xa00520e1, "smulwb    r0, r0, r5")
			testf(0xe10521e1, "smulwt    r1, r1, r5")
		}

		// BLK ------------------------------------------
		testf(0x0f50bd28, "ldm       sp!, {r0, r1, r2, r3, r12, lr}")
		testf(0x0c50bde9, "ldmib     sp!, {r2, r3, r12, lr}")
		testf(0x0f502de9, "stmdb     sp!, {r0, r1, r2, r3, r12, lr}")
		testf(0xff7fb1e9, "ldmib     r1!, {r0, r1, r2, r3, r4, r5, r6, r7, r8, r9, r10, r11, r12, sp, lr}")

		// HLF ------------------------------------------
		testf(0xb010c3e1, "strh      r1, [r3, #0x0]")
		testf(0xb200d1e1, "ldrh      r0, [r1, #0x2]")
		testf(0xb280c2e0, "strh      r8, [r2], #0x2")
		testf(0xb28072e1, "ldrh      r8, [r2, #0x-2]!")
		testf(0xd120d3e0, "ldrsb     r2, [r3], #0x1")
		testf(0xf004d5e1, "ldrsh     r0, [r5, #0x40]")
		testf(0xb30081b1, "strhlt    r0, [r1, r3]")
		testf(0x7f6051e4, "ldrb      r6, [r1], -#0x7f")
		testf(0xb29050e0, "ldrh      r9, [r0], #0xfffffffe")

		// BLX ------------------------------------------
		testf(0x10ff2fe1, "bx        r0")
		testf(0x31ff2fe1, "blx       r1")
		testf(0x1eff2fe1, "bx        lr")
		testf(0x1cff2f01, "bxeq      r12")
		testf(0xceffffea, "b         20d0bf4")
		testf(0x0300000a, "beq       37fb6b8")
		testf(0x300100bb, "bllt      2197fdc")
		testf(0xf0fbffeb, "bl        38001d4")
		testf(0xcc57619a, "bls       1855f88")
		testf(0x300100fb, "bx        2197fdc")

		// PSR ------------------------------------------
		testf1(0x0bf029e1, "msr       cpsr_fc, r11", func(cpu *Cpu) {
			cpu.Regs[11] = (reg(rand.Uint32()) & 0xF0000000) | reg(CpuModeIrq)
		})
		testf1(0x0ef06fe1, "msr       spsr_irq_fsxc, lr", func(cpu *Cpu) {
			cpu.Cpsr.r = (reg(rand.Uint32()) & 0xF0000000) | reg(CpuModeSupervisor)
			cpu.Regs[14] = (reg(rand.Uint32()) & 0xF0000000) | reg(CpuModeUser)
		})
		testf(0x00c00fF1, "mrs       r12, cpsr")

		// COP ------------------------------------------
		testf(0x114f19ee, "mrc       p15, #0, r4, c9, c1, #0")
		testf(0x9a0f07ee, "mcr       p15, #0, r0, c7, c10, #4")

		// FLAGS ----------------------------------------
		testf(0x01007305, "ldrbeq    r0, [r3, #0x-1]!")
		testf(0x01007315, "ldrbne    r0, [r3, #0x-1]!")
		testf(0x01007325, "ldrbhs    r0, [r3, #0x-1]!")
		testf(0x01007335, "ldrblo    r0, [r3, #0x-1]!")
		testf(0x01007345, "ldrbmi    r0, [r3, #0x-1]!")
		testf(0x01007355, "ldrbpl    r0, [r3, #0x-1]!")
		testf(0x01007365, "ldrbvs    r0, [r3, #0x-1]!")
		testf(0x01007375, "ldrbvc    r0, [r3, #0x-1]!")
		testf(0x01007385, "ldrbhi    r0, [r3, #0x-1]!")
		testf(0x01007395, "ldrbls    r0, [r3, #0x-1]!")
		testf(0x010073a5, "ldrbge    r0, [r3, #0x-1]!")
		testf(0x010073b5, "ldrblt    r0, [r3, #0x-1]!")
		testf(0x010073c5, "ldrbgt    r0, [r3, #0x-1]!")
		testf(0x010073d5, "ldrble    r0, [r3, #0x-1]!")
	}

	total = append(total, total...)
	total = append(total, total...)
	total = append(total, total...)
	total = append(total, total...)
	t0 := time.Now()
	jit.EmitBlock(total)
	t.Logf("Compilation of %d ops: %v", len(total), time.Since(t0))
}
