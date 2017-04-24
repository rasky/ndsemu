package arm

import (
	"encoding/binary"
	"math/rand"
	"testing"

	a "github.com/rasky/gojit/amd64"
	"golang.org/x/arch/x86/x86asm"
)

func TestAlu(t *testing.T) {
	jita, err := a.NewGoABI(1024 * 1024)
	if err != nil {
		t.Fatal(err)
	}
	jit := &jitArm{jita}

	testf := func(op uint32, exp string) {
		var buf [4]byte
		binary.LittleEndian.PutUint32(buf[:], op)
		op = binary.BigEndian.Uint32(buf[:])

		jita.Off = 0
		f := jit.DoBlock([]uint32{op})
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

		var cpu1, cpu2 Cpu

		for i := 0; i < 1024; i++ {
			var pre [16]reg
			// Test a few times with different registers/flags
			for j := 0; j < 16; j++ {
				pre[j] = reg(rand.Uint32())
				cpu1.Regs[j] = pre[j]
			}
			cpu1.Cpsr.r = reg(rand.Uint32() & 0xF0000000)
			cpu2 = cpu1

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
					t.Errorf("R%d differs: exp:%v jit:%v", i, cpu1.Regs[i], cpu2.Regs[i])
				}
			}
			if cpu1.Cpsr != cpu2.Cpsr {
				t.Errorf("R0:%v R7:%v", pre[0], pre[7])
				t.Errorf("Cpsr differs: exp:%v jit:%v", cpu1.Cpsr, cpu2.Cpsr)
			}
			if cpu1.Clock != cpu2.Clock {
				t.Errorf("Clock differs: exp:%v jit:%v", cpu1.Clock, cpu2.Clock)
			}
		}
	}

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

}
