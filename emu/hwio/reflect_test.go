package hwio

import "testing"

type test1 struct {
	Reg1   Reg16 `hwio:"offset=0x111,reset=0x123,rwmask=0x1,wcb"`
	Reg2   Reg32 `hwio:"offset=0x444,bank=1,rcb"`
	called bool
}

func (t *test1) WriteREG1(old uint16, val uint16) {
	t.called = true
}

func (t *test1) ReadREG2(val uint32) uint32 {
	return val | 1
}

func TestReflect(t *testing.T) {
	ts := &test1{}

	err := InitRegs(ts)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ts)

	if ts.Reg1.Name != "Reg1" || ts.Reg2.Name != "Reg2" {
		t.Error("invalid names:", ts.Reg1, ts.Reg2)
	}

	if ts.Reg2.Read32(0) != 1 {
		t.Error("invalid read32:", ts.Reg2.Read32(0))
	}

	val := ts.Reg1.Read16(0)
	if val != 0x123 {
		t.Error("invalid read16", val)
	}

	ts.Reg1.Write16(0, 0)
	if ts.Reg1.Value != 0x122 {
		t.Error("invalid read after rwmask", ts.Reg1.Value)
	}
	if !ts.called {
		t.Error("callback not called")
	}
}

func TestParseBank(t *testing.T) {
	ts := &test1{}
	info, err := bankGetRegs(ts, 0)
	if err != nil {
		t.Fatal(err)
	}
	if len(info) != 1 {
		t.Fatal("wrong number of regs in bank:", len(info))
	}
	if info[0].offset != 0x111 {
		t.Errorf("invalid reg offset: %x", info[0].offset)
	}

	rptr, ok := info[0].regPtr.(*Reg16)
	if !ok {
		t.Errorf("invalid reg ptr type: %T", info[0].regPtr)
	} else if rptr != &ts.Reg1 {
		t.Errorf("invalid reg ptr")
	}

	info, err = bankGetRegs(ts, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(info) != 1 {
		t.Fatal("wrong number of regs in bank:", len(info))
	}
	if info[0].offset != 0x444 {
		t.Errorf("invalid reg offset: %x", info[0].offset)
	}

	rptr2, ok := info[0].regPtr.(*Reg32)
	if !ok {
		t.Errorf("invalid reg ptr type: %T", info[0].regPtr)
	} else if rptr2 != &ts.Reg2 {
		t.Errorf("invalid reg ptr")
	}
}

func TestReadWriteOnly(t *testing.T) {
	type test2 struct {
		Reg1 Reg32 `hwio:"reset=0x123,readonly"`
		Reg2 Reg32 `hwio:"writeonly"`
	}

	ts := &test2{}
	err := InitRegs(ts)
	if err != nil {
		t.Fatal(err)
	}

	ts.Reg1.Write32(0, 0) // this should be ignored
	if ts.Reg1.Read32(0) != 0x123 {
		t.Error("invalid reg1 read:", ts.Reg1.Read32(0))
	}

	ts.Reg2.Write32(0, 0x123)
	if ts.Reg2.Read32(0) != 0 {
		t.Error("invalid reg2 read:", ts.Reg2.Read32(0))
	}
}
