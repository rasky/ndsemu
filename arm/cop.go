package arm

import (
	"unsafe"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type Coprocessor interface {
	Read(op uint32, cn, cm, cp uint32) uint32
	Write(op uint32, cn, cm, cp uint32, value uint32)
	Exec(op uint32, cn, cm, cp, cd uint32)
}

const cDtcmPhysicalSize = 16 * 1024
const cItcmPhysicalSize = 32 * 1024

type Cp15 struct {
	cpu          *Cpu
	regControl   reg
	regDtcmVsize reg
	regItcmVsize reg

	dtcm [cDtcmPhysicalSize]byte
	itcm [cItcmPhysicalSize]byte
}

// Check whether the specified address is mapped within the TCM area. If it does,
// return a pointer to it, otherwise return nil
func (c *Cp15) CheckTcm(addr uint32) unsafe.Pointer {

	// Between ITCM and DTCM, ITCM has priority, so check that first.
	if c.regControl.Bit(18) { // ITCM enable
		base := uint32(c.regItcmVsize) & 0xFFFFF000
		size := uint32(512 << uint((c.regItcmVsize>>1)&0x1F))
		if addr >= base && addr < base+size {
			addr = (addr - base) % cItcmPhysicalSize
			return unsafe.Pointer(&c.itcm[addr])
		}
	}

	if c.regControl.Bit(16) { // DTCM enable
		base := uint32(c.regDtcmVsize) & 0xFFFFF000
		size := uint32(512 << uint((c.regDtcmVsize>>1)&0x1F))
		if addr >= base && addr < base+size {
			addr = (addr - base) % cDtcmPhysicalSize
			return unsafe.Pointer(&c.dtcm[addr])
		}
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
		const rwmask = (1 << 0) | (1 << 2) | (1 << 7) | (0xFF << 12)
		c.regControl.SetWithMask(value, rwmask)
		log.WithField("val", c.regControl).WithField("pc", c.cpu.GetPC()).Info("[CP15] write control reg")
	case cn == 9 && cm == 1 && cp == 0:
		c.regDtcmVsize = reg(value)
		log.WithField("val", c.regDtcmVsize).WithField("pc", c.cpu.GetPC()).Info("[CP15] write DTCM size")
	case cn == 9 && cm == 1 && cp == 1:
		c.regItcmVsize = reg(value)
		log.WithField("val", c.regItcmVsize).WithField("pc", c.cpu.GetPC()).Info("[CP15] write ITCM size")
	default:
		log.WithField("pc", c.cpu.GetPC()).Warnf("[CP15] unhandled write C%d,C%d,%d", cn, cm, cp)
		return
	}
}

func (c *Cp15) Exec(op uint32, cn, cm, cp, cd uint32) {
	log.WithField("op", op).WithField("pc", c.cpu.GetPC()).Error("[CP15] invalid op in exec")
	return
}

func NewCp15(cpu *Cpu) *Cp15 {
	return &Cp15{
		cpu:        cpu,
		regControl: reg(0x2078),
	}
}
