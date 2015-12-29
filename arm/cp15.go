package arm

import log "gopkg.in/Sirupsen/logrus.v0"

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
}

func (c *Cp15) CheckITcm(addr uint32) []uint8 {
	if c.regControl.Bit(18) { // ITCM enable
		base := uint32(c.regItcmVsize) & 0xFFFFF000
		size := uint32(512 << uint((c.regItcmVsize>>1)&0x1F))
		if addr >= base && addr < base+size {
			addr = (addr - base) & c.itcmSizeMask
			return c.itcm[addr:]
		}
	}
	return nil
}

func (c *Cp15) CheckDTcm(addr uint32) []uint8 {
	if c.regControl.Bit(16) { // DTCM enable
		base := uint32(c.regDtcmVsize) & 0xFFFFF000
		size := uint32(512 << uint((c.regDtcmVsize>>1)&0x1F))
		if addr >= base && addr < base+size {
			addr = (addr - base) & c.dtcmSizeMask
			return c.dtcm[addr:]
		}
	}
	return nil
}

// Check whether the specified address is mapped within the TCM area. If it does,
// return a pointer to it, otherwise return nil
func (c *Cp15) CheckTcm(addr uint32) []uint8 {
	// Between ITCM and DTCM, ITCM has priority, so check that first.
	if ptr := c.CheckITcm(addr); ptr != nil {
		return ptr
	}
	if ptr := c.CheckDTcm(addr); ptr != nil {
		return ptr
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
		log.WithField("op", op).Error("[CP15] invalid op in read")
		return 0
	}

	switch {
	case cn == 1 && cm == 0 && cp == 0:
		log.WithField("val", c.regControl).WithField("pc", c.cpu.GetPC()).Info("[CP15] read control reg")
		return uint32(c.regControl)
	case cn == 9 && cm == 1 && cp == 0:
		log.WithField("val", c.regDtcmVsize).WithField("pc", c.cpu.GetPC()).Info("[CP15] read DTCM size")
		return uint32(c.regDtcmVsize)
	case cn == 9 && cm == 1 && cp == 1:
		log.WithField("val", c.regItcmVsize).WithField("pc", c.cpu.GetPC()).Info("[CP15] read ITCM size")
		return uint32(c.regItcmVsize)
	default:
		log.WithField("pc", c.cpu.GetPC()).Warnf("[CP15] unhandled read C%d,C%d,%d", cn, cm, cp)
		return 0
	}

}

func (c *Cp15) Write(op uint32, cn, cm, cp uint32, value uint32) {
	if op != 0 {
		log.WithField("op", op).Error("[CP15] invalid op in write")
		return
	}

	switch {
	case cn == 1 && cm == 0 && cp == 0:
		c.regControl.SetWithMask(value, c.regControlRwMask)
		if c.regControl.Bit(17) || c.regControl.Bit(19) {
			log.Fatal("DTCM/ITCM load mode")
		}
		log.WithField("val", c.regControl).WithField("pc", c.cpu.GetPC()).Info("[CP15] write control reg")
		if c.regControl.Bit(18) {
			base := uint32(c.regItcmVsize) & 0xFFFFF000
			size := uint32(512 << uint((c.regItcmVsize>>1)&0x1F))
			log.WithFields(log.Fields{
				"base": reg(base),
				"size": size,
			}).Info("[CP15] Activated ITCM")
		}
		if c.regControl.Bit(16) {
			base := uint32(c.regDtcmVsize) & 0xFFFFF000
			size := uint32(512 << uint((c.regDtcmVsize>>1)&0x1F))
			log.WithFields(log.Fields{
				"base": reg(base),
				"size": size,
			}).Info("[CP15] Activated DTCM")
		}
	case cn == 9 && cm == 1 && cp == 0:
		c.regDtcmVsize = reg(value)
		log.WithField("val", c.regDtcmVsize).WithField("pc", c.cpu.GetPC()).Info("[CP15] write DTCM size")
	case cn == 9 && cm == 1 && cp == 1:
		c.regItcmVsize = reg(value)
		log.WithField("val", c.regItcmVsize).WithField("pc", c.cpu.GetPC()).Info("[CP15] write ITCM size")

	case cn == 7:
		if (cm == 0 && cp == 4) || (cm == 8 && cp == 2) {
			// Halt processor (wait for interrupt
			log.WithField("pc", c.cpu.GetPC()).Info("[CP15} halt cpu")
			c.cpu.SetLine(LineHalt, true)
		}
		// anything else is a cache command, ignore

	default:
		log.WithField("pc", c.cpu.GetPC()).Warnf("[CP15] unhandled write C%d,C%d,%d", cn, cm, cp)
		return
	}
}

func (c *Cp15) Exec(op uint32, cn, cm, cp, cd uint32) {
	log.WithField("op", op).WithField("pc", c.cpu.GetPC()).Error("[CP15] invalid op in exec")
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
