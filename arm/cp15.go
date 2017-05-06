package arm

import (
	log "ndsemu/emu/logger"
)

var modCp15 = log.NewModule("cp15")

type Cp15 struct {
	cpu              *Cpu
	regControl       reg
	regControlRwMask uint32
	regDtcmVsize     reg
	regItcmVsize     reg

	itcm         []byte
	dtcm         []byte
	itcmSizeMask uint32
	dtcmSizeMask uint32

	itcmBegin uint32
	itcmEnd   uint32
	dtcmBegin uint32
	dtcmEnd   uint32
}

// updateTcmConfig() recalculates the variables xtcmBegin/xtcmEnd, used by
// CheckXTcm() to check whether an address lies withn the xTCM area.
func (c *Cp15) updateTcmConfig() {
	if c.regControl.Bit(18) { // ITCM enable
		c.itcmBegin = uint32(c.regItcmVsize) & 0xFFFFF000
		c.itcmEnd = c.itcmBegin + uint32(512<<uint((c.regItcmVsize>>1)&0x1F))
	} else {
		// If the area is disabled, set these vars to a zero-sized area that
		// will never match the checks in CheckITcm. We use 0xFFFFFFFF instead
		// of 0x0 so that the first check in CheckITcm returns false (rather
		// than the second)
		c.itcmBegin = 0xFFFFFFFF
		c.itcmEnd = 0xFFFFFFFF
	}

	if c.regControl.Bit(16) { // DTCM enable
		c.dtcmBegin = uint32(c.regDtcmVsize) & 0xFFFFF000
		c.dtcmEnd = c.dtcmBegin + uint32(512<<uint((c.regDtcmVsize>>1)&0x1F))
	} else {
		c.dtcmBegin = 0xFFFFFFFF
		c.dtcmEnd = 0xFFFFFFFF
	}
}

// Check whether the specified address falls within the ITCM area, and returns
// a slice to the referenced point. Returns nil if the address is outside, or
// ITCM is disabled.
func (c *Cp15) CheckITcm(addr uint32) []uint8 {
	if addr >= c.itcmBegin && addr < c.itcmEnd {
		addr = (addr - c.itcmBegin) & c.itcmSizeMask
		return c.itcm[addr:]
	}
	return nil
}

// Check whether the specified address falls within the DTCM area, and returns
// a slice to the referenced point. Returns nil if the address is outside, or
// DTCM is disabled.
func (c *Cp15) CheckDTcm(addr uint32) []uint8 {
	if addr >= c.dtcmBegin && addr < c.dtcmEnd {
		addr = (addr - c.dtcmBegin) & c.dtcmSizeMask
		return c.dtcm[addr:]
	}
	return nil
}

func (c *Cp15) ExceptionVector() uint32 {
	if c.regControl.Bit(13) {
		return 0xFFFF0000
	} else {
		return 0x00000000
	}
}

func (c *Cp15) Read(op uint32, cn, cm, cp uint32) uint32 {
	if op != 0 {
		modCp15.ErrorZ("invalid op in read").Uint32("op", op).End()
		return 0
	}

	switch {
	case cn == 1 && cm == 0 && cp == 0:
		// modCp15.WithField("val", c.regControl).WithField("pc", c.cpu.GetPC()).Info("read control reg")
		return uint32(c.regControl)
	case cn == 9 && cm == 1 && cp == 0:
		// modCp15.WithField("val", c.regDtcmVsize).WithField("pc", c.cpu.GetPC()).Info("read DTCM size")
		return uint32(c.regDtcmVsize)
	case cn == 9 && cm == 1 && cp == 1:
		// modCp15.WithField("val", c.regItcmVsize).WithField("pc", c.cpu.GetPC()).Info("read ITCM size")
		return uint32(c.regItcmVsize)
	default:
		modCp15.WarnZ("unhandled read").Uint32("cn", cn).Uint32("cm", cm).Uint32("cp", cp).End()
		return 0
	}

}

func (c *Cp15) Write(op uint32, cn, cm, cp uint32, value uint32) {
	if op != 0 {
		modCp15.ErrorZ("invalid op in write").Uint32("op", op).End()
		return
	}

	switch {
	case cn == 1 && cm == 0 && cp == 0:
		c.regControl.SetWithMask(value, c.regControlRwMask)
		if c.regControl.Bit(17) || c.regControl.Bit(19) {
			modCp15.FatalZ("DTCM/ITCM load mode").End()
		}
		modCp15.InfoZ("write control reg").
			Hex32("val", uint32(c.regControl)).
			End()
		if c.regControl.Bit(18) {
			base := uint32(c.regItcmVsize) & 0xFFFFF000
			size := uint32(512 << uint((c.regItcmVsize>>1)&0x1F))
			modCp15.InfoZ("Activated ITCM").
				Hex32("base", base).
				Hex32("size", size).
				End()
		} else {
			modCp15.InfoZ("Disabled ITCM").End()
		}
		if c.regControl.Bit(16) {
			base := uint32(c.regDtcmVsize) & 0xFFFFF000
			size := uint32(512 << uint((c.regDtcmVsize>>1)&0x1F))
			modCp15.InfoZ("Activated DTCM").
				Hex32("base", base).
				Hex32("size", size).
				End()
		} else {
			modCp15.InfoZ("Disabled DTCM").End()
		}
		c.updateTcmConfig()
	case cn == 9 && cm == 1 && cp == 0:
		c.regDtcmVsize = reg(value)
		c.updateTcmConfig()
		modCp15.InfoZ("write DTCM size").Hex32("val", uint32(c.regDtcmVsize)).End()
	case cn == 9 && cm == 1 && cp == 1:
		c.regItcmVsize = reg(value)
		c.updateTcmConfig()
		modCp15.InfoZ("write ITCM size").Hex32("val", uint32(c.regItcmVsize)).End()

	case cn == 6:
		modCp15.InfoZ("PU region configuration").
			Uint32("region", cm).
			Bool("enable", value&1 != 0).
			Hex32("base", (value>>12)*4096).
			Hex32("size", 2<<((value>>1)&0x1F)).
			End()

	case cn == 7:
		if (cm == 0 && cp == 4) || (cm == 8 && cp == 2) {
			// Halt processor (wait for interrupt
			modCp15.InfoZ("halt cpu").Hex32("pc", uint32(c.cpu.GetPC())).End()
			c.cpu.SetLine(LineHalt, true)
		}
		// anything else is a cache command, ignore

	default:
		modCp15.WarnZ("unhandled write").Uint32("cn", cn).Uint32("cm", cm).Uint32("cp", cp).Hex32("value", value).End()
		return
	}
}

func (c *Cp15) Exec(op uint32, cn, cm, cp, cd uint32) {
	modCp15.ErrorZ("invalid op in exec").Hex32("op", op).End()
	return
}

// ConfigureTcm activates emulation of TCM (tightly-coupled memory), which
// is a level-1 memory with zero waitstates. ARM supports two different areas
// (ITCM and DTCM), with possibly different sizes, and CP15 has specific
// registers to map TCM in the virtual address space.
func (c *Cp15) ConfigureTcm(itcmSize int, dtcmSize int) {
	if itcmSize > 0 {
		c.itcm = make([]byte, itcmSize)
	} else {
		c.itcm = nil
	}

	if dtcmSize > 0 {
		c.dtcm = make([]byte, dtcmSize)
	} else {
		c.dtcm = nil
	}

	// (assuming pow2)
	c.itcmSizeMask = uint32(itcmSize - 1)
	c.dtcmSizeMask = uint32(dtcmSize - 1)
}

// Configure the CP15 Control Register. Value is the initial value of the register,
// while rwmask specifies which bits can be modified at runtime, and which bits
// are fixed.
func (c *Cp15) ConfigureControlReg(value uint32, rwmask uint32) {
	c.regControl = reg(value)
	c.regControlRwMask = rwmask
}

func newCp15(cpu *Cpu) *Cp15 {
	return &Cp15{
		cpu:              cpu,
		regControlRwMask: 0xFFFFFFFF,
	}
}
