package hwio

import "testing"

func TestTableRead(t *testing.T) {
	r1 := Reg16{Value: 0x1122}
	r2 := Reg16{Value: 0x3344}
	r3 := Reg32{Value: 0xAABBCCDD}
	r4 := Reg64{Value: 0x005566778899EEFF}

	table := Table{Name: "t1"}
	table.Reset()
	table.MapReg16(0x400014, &r1)
	table.MapReg16(0x400016, &r2)
	table.MapReg64(0x400018, &r4)
	table.MapReg32(0x400020, &r3)

	exp8 := []uint8{
		0x00, 0x00, 0x00, 0x00,
		0x22, 0x11,
		0x44, 0x33,
		0xFF, 0xEE, 0x99, 0x88, 0x77, 0x66, 0x55, 0x00,
		0xDD, 0xCC, 0xBB, 0xAA,
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
	}

	for off, want := range exp32 {
		got := table.Read32(uint32(0x500010 + off*4))
		if want != got {
			t.Errorf("invalid read32 at %x: got:%x,want:%x", off*4, got, want)
		}
	}
}

func TestTableWrite(t *testing.T) {
	r1 := Reg16{Value: 0x1122, RoMask: 0xF0F0}
	r2 := Reg16{Value: 0x3344, RoMask: 0xF0F0}
	r3 := Reg32{Value: 0xAABBCCDD, RoMask: 0xF0F0F0F0}
	r4 := Reg64{Value: 0x115566778899EEFF, RoMask: 0xF0F0F0F0F0F0F0F0}

	table := Table{Name: "t1"}
	table.Reset()
	table.MapReg16(0x400014, &r1)
	table.MapReg16(0x400016, &r2)
	table.MapReg64(0x400018, &r4)
	table.MapReg32(0x400020, &r3)

	for i := 0; i < 0x14; i++ {
		table.Write8(uint32(0x500010+i), 0)
	}

	if r1.Value != 0x1020 || r2.Value != 0x3040 ||
		r3.Value != 0xA0B0C0D0 ||
		r4.Value != 0x105060708090E0F0 {
		t.Error("invalid regs after write8", r1, r2, r3, r4)
	}

	for i := 0; i < 0x14/2; i++ {
		table.Write16(uint32(0x500010+i*2), 0xFFFF)
	}

	if r1.Value != 0x1F2F || r2.Value != 0x3F4F ||
		r3.Value != 0xAFBFCFDF ||
		r4.Value != 0x1F5F6F7F8F9FEFFF {
		t.Error("invalid regs after write16", r1, r2, r3, r4)
	}

	for i := 0; i < 0x14/4; i++ {
		table.Write32(uint32(0x500010+i*4), 0xCCCCCCCC)
	}

	if r1.Value != 0x1C2C || r2.Value != 0x3C4C ||
		r3.Value != 0xACBCCCDC ||
		r4.Value != 0x1C5C6C7C8C9CECFC {
		t.Error("invalid regs after write32", r1, r2, r3, r4)
	}
}
