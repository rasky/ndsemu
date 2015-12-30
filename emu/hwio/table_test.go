package hwio

import "testing"

func TestTableRead(t *testing.T) {
	r1 := Reg16{Value: 0x1122}
	r2 := Reg16{Value: 0x3344}
	r3 := Reg32{Value: 0xAABBCCDD}
	r4 := Reg64{Value: 0x005566778899EEFF}
	r5 := Reg8{Value: 0x12}
	r6 := Reg8{Value: 0x34}
	r7 := Reg16{Value: 0xFEEF}
	r8 := Reg8{Value: 0x56}
	r9 := Reg8{Value: 0x78}

	table := Table{Name: "t1"}
	table.Reset()
	table.MapReg16(0x400014, &r1)
	table.MapReg16(0x400016, &r2)
	table.MapReg64(0x400018, &r4)
	table.MapReg32(0x400020, &r3)
	table.MapReg8(0x400024, &r5)
	table.MapReg8(0x400025, &r6)
	table.MapReg16(0x400026, &r7)
	table.MapReg8(0x400028, &r8)
	table.MapReg8(0x40002B, &r9)

	exp8 := []uint8{
		0x00, 0x00, 0x00, 0x00,
		0x22, 0x11,
		0x44, 0x33,
		0xFF, 0xEE, 0x99, 0x88, 0x77, 0x66, 0x55, 0x00,
		0xDD, 0xCC, 0xBB, 0xAA,
		0x12,
		0x34,
		0xEF, 0xFE,
		0x56,
		0x00,
		0x00,
		0x78,
	}

	for off, want := range exp8 {
		got := table.Read8(uint32(0x500010 + off))
		if want != got {
			t.Errorf("invalid read8 at %x: got:%x,want:%x", off, got, want)
		}
	}

	exp16 := []uint16{
		0x0000, 0x0000,
		0x1122,
		0x3344,
		0xEEFF, 0x8899, 0x6677, 0x0055,
		0xCCDD, 0xAABB,
		0x3412,
		0xFEEF,
		0x0056,
		0x7800,
	}

	for off, want := range exp16 {
		got := table.Read16(uint32(0x500010 + off*2))
		if want != got {
			t.Errorf("invalid read16 at %x: got:%x,want:%x", off*2, got, want)
		}
	}

	exp32 := []uint32{
		0x00000000,
		0x33441122,
		0x8899EEFF, 0x00556677,
		0xAABBCCDD,
		0xFEEF3412,
		0x78000056,
	}

	for off, want := range exp32 {
		got := table.Read32(uint32(0x500010 + off*4))
		if want != got {
			t.Errorf("invalid read32 at %x: got:%x,want:%x", off*4, got, want)
		}
	}
}

func TestCallbackNumCalls(t *testing.T) {
	calls := ""

	r1 := Reg16{Value: 0x1234, ReadCb: func(v uint16) uint16 {
		calls += "1"
		return v
	}}
	r2 := Reg8{Value: 0x56, ReadCb: func(v uint8) uint8 {
		calls += "2"
		return v
	}}
	r3 := Reg8{Value: 0x78, ReadCb: func(v uint8) uint8 {
		calls += "3"
		return v
	}}

	table := Table{Name: "t1"}
	table.Reset()
	table.MapReg16(0x400014, &r1)
	table.MapReg8(0x400016, &r2)
	table.MapReg8(0x400017, &r3)

	val := table.Read32(0x400014)
	if val != 0x78561234 {
		t.Errorf("invalid read32, got:%x, want:%x", val, 0x78561234)
	}

	if calls != "123" {
		t.Errorf("invalid calls, got:%s want:%s", calls, "123")
	}
}

func TestTableWrite(t *testing.T) {
	r1 := Reg16{Value: 0x1122, RoMask: 0xF0F0}
	r2 := Reg16{Value: 0x3344, RoMask: 0xF0F0}
	r3 := Reg32{Value: 0xAABBCCDD, RoMask: 0xF0F0F0F0}
	r4 := Reg64{Value: 0x115566778899EEFF, RoMask: 0xF0F0F0F0F0F0F0F0}
	r5 := Reg8{Value: 0x12, RoMask: 0x0F}
	r6 := Reg8{Value: 0x34, RoMask: 0x0F}
	r7 := Reg16{Value: 0xFEEF, RoMask: 0x0F0F}
	r8 := Reg8{Value: 0x56, RoMask: 0x0F}
	r9 := Reg8{Value: 0x78, RoMask: 0x0F}

	table := Table{Name: "t1"}
	table.Reset()
	table.MapReg16(0x400014, &r1)
	table.MapReg16(0x400016, &r2)
	table.MapReg64(0x400018, &r4)
	table.MapReg32(0x400020, &r3)
	table.MapReg8(0x400024, &r5)
	table.MapReg8(0x400025, &r6)
	table.MapReg16(0x400026, &r7)
	table.MapReg8(0x400028, &r8)
	table.MapReg8(0x40002B, &r9)

	for i := 0; i < 0x1C; i++ {
		table.Write8(uint32(0x500010+i), 0)
	}

	if r1.Value != 0x1020 || r2.Value != 0x3040 ||
		r3.Value != 0xA0B0C0D0 ||
		r4.Value != 0x105060708090E0F0 ||
		r5.Value != 0x02 || r6.Value != 0x04 ||
		r7.Value != 0x0E0F || r8.Value != 0x06 || r9.Value != 0x08 {
		t.Error("invalid regs after write8", r1, r2, r3, r4, r5, r6, r7, r8, r9)
	}

	for i := 0; i < 0x1C/2; i++ {
		table.Write16(uint32(0x500010+i*2), 0xFFFF)
	}

	if r1.Value != 0x1F2F || r2.Value != 0x3F4F ||
		r3.Value != 0xAFBFCFDF ||
		r4.Value != 0x1F5F6F7F8F9FEFFF ||
		r5.Value != 0xF2 || r6.Value != 0xF4 ||
		r7.Value != 0xFEFF || r8.Value != 0xF6 || r9.Value != 0xF8 {
		t.Error("invalid regs after write16", r1, r2, r3, r4, r5, r6, r7, r8, r9)
	}

	for i := 0; i < 0x1C/4; i++ {
		table.Write32(uint32(0x500010+i*4), 0xCCCCCCCC)
	}

	if r1.Value != 0x1C2C || r2.Value != 0x3C4C ||
		r3.Value != 0xACBCCCDC ||
		r4.Value != 0x1C5C6C7C8C9CECFC ||
		r5.Value != 0xC2 || r6.Value != 0xC4 ||
		r7.Value != 0xCECF || r8.Value != 0xC6 || r9.Value != 0xC8 {
		t.Error("invalid regs after write32", r1, r2, r3, r4, r5, r6, r7, r8, r9)
	}
}
