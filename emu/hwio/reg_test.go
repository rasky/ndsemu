package hwio

import "testing"

func TestReg8(t *testing.T) {
	r := Reg8{Value: 0x11, RoMask: 0xF0}

	if r.Read8(0) != 0x11 {
		t.Errorf("invalid read: %x", r.Read8(0))
	}
	if r.Read8(9999) != 0x11 {
		t.Errorf("invalid read with offset: %x", r.Read8(9999))
	}

	r.Write8(0, 0x77)
	if r.Value != 0x17 {
		t.Errorf("writemask not respected: %x", r.Value)
	}
	r.Write8(99999, 0x88)
	if r.Value != 0x18 {
		t.Errorf("writemask with offset not respected: %x", r.Value)
	}
}

func TestReg16(t *testing.T) {
	r := Reg16{Value: 0x1122, RoMask: 0xF0F0}

	if r.Read16(0) != 0x1122 {
		t.Errorf("invalid read: %x", r.Read16(0))
	}
	if r.Read16(9999) != 0x1122 {
		t.Errorf("invalid read with offset: %x", r.Read16(9999))
	}

	r.Write16(0, 0x7777)
	if r.Value != 0x1727 {
		t.Errorf("writemask not respected: %x", r.Value)
	}
	r.Write16(99999, 0x8888)
	if r.Value != 0x1828 {
		t.Errorf("writemask with offset not respected: %x", r.Value)
	}
}

func TestReg16RW8(t *testing.T) {
	r := Reg16{Value: 0x1122, RoMask: 0xF0F0}

	if r.Read8(0) != 0x22 {
		t.Errorf("invalid read8 0: %x", r.Read8(0))
	}
	if r.Read8(1) != 0x11 {
		t.Errorf("invalid read8 1: %x", r.Read8(1))
	}
	if r.Read8(0x122) != 0x22 {
		t.Errorf("invalid read8 0x122: %x", r.Read8(0x122))
	}
	if r.Read8(0x123) != 0x11 {
		t.Errorf("invalid read8 0x123: %x", r.Read8(0x123))
	}

	r.Write8(0, 0x77)
	if r.Value != 0x1127 {
		t.Errorf("invalid write8 0: %x", r.Value)
	}
	r.Write8(1, 0x77)
	if r.Value != 0x1727 {
		t.Errorf("invalid write8 1: %x", r.Value)
	}

	r.Write8(0x998, 0x88)
	if r.Value != 0x1728 {
		t.Errorf("invalid write8 0x998: %x", r.Value)
	}
	r.Write8(0x999, 0x88)
	if r.Value != 0x1828 {
		t.Errorf("invalid write8 0x999: %x", r.Value)
	}
}

func TestReg32(t *testing.T) {
	r := Reg32{Value: 0x11223344, RoMask: 0xF0F0F0F0}

	if r.Read32(0) != 0x11223344 {
		t.Errorf("invalid read: %x", r.Read16(0))
	}
	if r.Read32(9999) != 0x11223344 {
		t.Errorf("invalid read with offset: %x", r.Read16(9999))
	}

	r.Write32(0, 0x77777777)
	if r.Value != 0x17273747 {
		t.Errorf("writemask not respected: %x", r.Value)
	}
	r.Write32(99999, 0x88888888)
	if r.Value != 0x18283848 {
		t.Errorf("writemask with offset not respected: %x", r.Value)
	}
}

func TestReg32RW16(t *testing.T) {
	r := Reg32{Value: 0x11223344, RoMask: 0xF0F0F0F0}

	if r.Read16(0) != 0x3344 {
		t.Errorf("invalid read16 0: %x", r.Read16(0))
	}
	if r.Read16(2) != 0x1122 {
		t.Errorf("invalid read16 2: %x", r.Read16(2))
	}
	if r.Read16(0x124) != 0x3344 {
		t.Errorf("invalid read16 0x124: %x", r.Read16(0x124))
	}
	if r.Read16(0x126) != 0x1122 {
		t.Errorf("invalid read16 0x126: %x", r.Read16(0x124))
	}

	r.Write16(0, 0x7777)
	if r.Value != 0x11223747 {
		t.Errorf("invalid write16 0: %x", r.Value)
	}
	r.Write16(2, 0x7777)
	if r.Value != 0x17273747 {
		t.Errorf("invalid write16 1: %x", r.Value)
	}

	r.Write16(0x998, 0x8888)
	if r.Value != 0x17273848 {
		t.Errorf("invalid write16 0x998: %x", r.Value)
	}
	r.Write16(0x99A, 0x8888)
	if r.Value != 0x18283848 {
		t.Errorf("invalid write16 0x99A: %x", r.Value)
	}
}

func TestReg32RW8(t *testing.T) {
	r := Reg32{Value: 0x11223344, RoMask: 0xF0F0F0F0}

	if r.Read8(0) != 0x44 {
		t.Errorf("invalid read16 0: %x", r.Read8(0))
	}
	if r.Read8(1) != 0x33 {
		t.Errorf("invalid read16 1: %x", r.Read8(1))
	}
	if r.Read8(2) != 0x22 {
		t.Errorf("invalid read16 2: %x", r.Read8(2))
	}
	if r.Read8(3) != 0x11 {
		t.Errorf("invalid read16 3: %x", r.Read8(3))
	}
	if r.Read8(0x124) != 0x44 {
		t.Errorf("invalid read16 0x124: %x", r.Read8(0x124))
	}
	if r.Read8(0x125) != 0x33 {
		t.Errorf("invalid read16 0x125: %x", r.Read8(0x125))
	}
	if r.Read8(0x126) != 0x22 {
		t.Errorf("invalid read16 0x126: %x", r.Read8(0x126))
	}
	if r.Read8(0x127) != 0x11 {
		t.Errorf("invalid read16 0x127: %x", r.Read8(0x127))
	}

	r.Write8(0, 0x77)
	if r.Value != 0x11223347 {
		t.Errorf("invalid write8 0: %x", r.Value)
	}
	r.Write8(1, 0x77)
	if r.Value != 0x11223747 {
		t.Errorf("invalid write8 1: %x", r.Value)
	}
	r.Write8(2, 0x77)
	if r.Value != 0x11273747 {
		t.Errorf("invalid write8 2: %x", r.Value)
	}
	r.Write8(3, 0x77)
	if r.Value != 0x17273747 {
		t.Errorf("invalid write8 3: %x", r.Value)
	}

	r.Write8(0x998, 0x88)
	if r.Value != 0x17273748 {
		t.Errorf("invalid write16 0x998: %x", r.Value)
	}
	r.Write8(0x999, 0x88)
	if r.Value != 0x17273848 {
		t.Errorf("invalid write16 0x999: %x", r.Value)
	}
	r.Write8(0x99A, 0x88)
	if r.Value != 0x17283848 {
		t.Errorf("invalid write16 0x99A: %x", r.Value)
	}
	r.Write8(0x99B, 0x88)
	if r.Value != 0x18283848 {
		t.Errorf("invalid write16 0x99B: %x", r.Value)
	}
}
