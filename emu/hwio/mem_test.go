package hwio

import (
	"reflect"
	"testing"
)

type testMemSizes struct {
	Buf1 Mem `hwio:"size=0x200,rw8,rw16=off,rw32=off"`
	Buf2 Mem `hwio:"size=0x200,rw16,rw8=off,rw32=off"`
	Buf3 Mem `hwio:"size=0x200,rw32,rw8=off,rw16=off"`
}

func TestMemSizes(t *testing.T) {
	mem := testMemSizes{}
	MustInitRegs(&mem)

	mem.Buf1.BankIO8().Write8(0xFF000, 0xAB)
	val := mem.Buf1.BankIO8().Read8(0xEE000)
	if val != 0xAB {
		t.Errorf("invalid wr8, got:%x, want:%x", val, 0xAB)
	}
	if mem.Buf1.BankIO16() != nil || mem.Buf1.BankIO32() != nil {
		t.Error("buf1 shouldn't have 16/32 bit access")
	}

	mem.Buf2.BankIO16().Write16(0xFF000, 0xABCD)
	val2 := mem.Buf2.BankIO16().Read16(0xEE000)
	if val2 != 0xABCD {
		t.Errorf("invalid wr8, got:%x, want:%x", val, 0xABCD)
	}
	if mem.Buf2.BankIO8() != nil || mem.Buf2.BankIO32() != nil {
		t.Error("buf2 shouldn't have 8/32 bit access")
	}

	mem.Buf3.BankIO32().Write32(0xFF000, 0xABCD0123)
	val3 := mem.Buf3.BankIO32().Read32(0xEE000)
	if val3 != 0xABCD0123 {
		t.Errorf("invalid wr8, got:%x, want:%x", val, 0xABCD0123)
	}
	if mem.Buf3.BankIO8() != nil || mem.Buf3.BankIO16() != nil {
		t.Error("buf3 shouldn't have 8/16 bit access")
	}
}

func TestMemReadonly(t *testing.T) {
	var buf Mem
	buf.Data = make([]byte, 0x200)

	for _, f16 := range []MemFlags{MemFlag16Unaligned, MemFlag16ForceAlign, MemFlag16Byteswapped} {
		for _, f32 := range []MemFlags{MemFlag32Unaligned, MemFlag32ForceAlign, MemFlag32Byteswapped} {
			buf.Flags = f16 | f32 | MemFlag8 | MemFlagReadOnly
			buf.BankIO8().Write8(0x0, 0xFF)
			buf.BankIO16().Write16(0x2, 0xFFFF)
			buf.BankIO32().Write32(0x4, 0xFFFFFFFF)
		}
	}

	for i := 0; i < 8; i++ {
		if buf.Data[i] != 0 {
			t.Errorf("data written at offset %d", i)
		}
	}
}

func TestMemAlign16(t *testing.T) {
	var buf Mem
	buf.Data = make([]byte, 0x200)

	buf.Flags = MemFlag16Unaligned
	buf.BankIO16().Write16(0x1, 0xABCD)

	buf.Flags = MemFlag16ForceAlign
	buf.BankIO16().Write16(0x5, 0xABCD)

	buf.Flags = MemFlag16Byteswapped
	buf.BankIO16().Write16(0x9, 0xABCD)

	exp := []byte{
		0x00, 0xCD, 0xAB, 0x00, // unaligned
		0xCD, 0xAB, 0x00, 0x00, // force-align
		0xAB, 0xCD, 0x00, 0x00, // byte-swapped
	}

	if !reflect.DeepEqual(exp, buf.Data[:12]) {
		t.Errorf("invalid data, got:%x, exp:%x", buf.Data[:12], exp)
	}
}

func TestMemAlign32(t *testing.T) {
	var buf Mem
	buf.Data = make([]byte, 0x200)

	buf.Flags = MemFlag32Unaligned
	buf.BankIO32().Write32(0x0+1, 0xABCD1234)
	buf.BankIO32().Write32(0x4+2, 0xABCD1234)
	buf.BankIO32().Write32(0x8+3, 0xABCD1234)

	buf.Flags = MemFlag32ForceAlign
	buf.BankIO32().Write32(0x10+0, 0xABCD1234)
	buf.BankIO32().Write32(0x14+1, 0xABCD1234)
	buf.BankIO32().Write32(0x18+2, 0xABCD1234)
	buf.BankIO32().Write32(0x1C+3, 0xABCD1234)

	buf.Flags = MemFlag32Byteswapped
	buf.BankIO32().Write32(0x20+0, 0xABCD1234)
	buf.BankIO32().Write32(0x24+1, 0xABCD1234)
	buf.BankIO32().Write32(0x28+2, 0xABCD1234)
	buf.BankIO32().Write32(0x2C+3, 0xABCD1234)

	exp := []byte{
		// unaligned
		0x00, 0x34, 0x12, 0xCD,
		0xAB, 0x00, 0x34, 0x12,
		0xCD, 0xAB, 0x00, 0x34,
		0x12, 0xCD, 0xAB, 0x00,

		// force-align
		0x34, 0x12, 0xCD, 0xAB,
		0x34, 0x12, 0xCD, 0xAB,
		0x34, 0x12, 0xCD, 0xAB,
		0x34, 0x12, 0xCD, 0xAB,

		// byte-swapped
		0x34, 0x12, 0xCD, 0xAB,
		0xAB, 0x34, 0x12, 0xCD,
		0xCD, 0xAB, 0x34, 0x12,
		0x12, 0xCD, 0xAB, 0x34,
	}

	if !reflect.DeepEqual(exp[:0x10], buf.Data[:0x10]) {
		t.Errorf("invalid data for unaligned, got:%x, exp:%x", buf.Data[:0x10], exp[:0x10])
	}
	if !reflect.DeepEqual(exp[0x10:0x20], buf.Data[0x10:0x20]) {
		t.Errorf("invalid data for force-align, got:%x, exp:%x", buf.Data[0x10:0x20], exp[0x10:0x20])
	}
	if !reflect.DeepEqual(exp[0x20:0x30], buf.Data[0x20:0x30]) {
		t.Errorf("invalid data for byteswapped, got:%x, exp:%x", buf.Data[0x20:0x30], exp[0x20:0x30])
	}
}
