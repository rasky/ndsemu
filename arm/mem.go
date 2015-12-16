package arm

import (
	"fmt"

	log "gopkg.in/Sirupsen/logrus.v0"
)

func (cpu *Cpu) opFetch16(addr uint32) uint16 {
	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckITcm(addr); ptr != nil {
			return *(*uint16)(ptr)
		}
	}
	return cpu.bus.Read16(addr)
}

func (cpu *Cpu) opFetch32(addr uint32) uint32 {
	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckITcm(addr); ptr != nil {
			return *(*uint32)(ptr)
		}
	}
	return cpu.bus.Read32(addr)
}

func (cpu *Cpu) opRead32(addr uint32) uint32 {
	if addr&3 != 0 {
		log.WithFields(log.Fields{
			"pc":   cpu.GetPC(),
			"addr": fmt.Sprintf("%08x", addr),
		}).Error("unaligned read32 memory access")

		rot := uint(addr&3) * 8
		res := cpu.bus.Read32(addr &^ 3)
		return res>>rot | res<<(32-rot)
	}

	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckTcm(addr); ptr != nil {
			return *(*uint32)(ptr)
		}
	}

	return cpu.bus.Read32(addr)
}

func (cpu *Cpu) opWrite32(addr uint32, val uint32) {
	if addr&3 != 0 {
		log.WithFields(log.Fields{
			"pc":   cpu.pc - 4,
			"addr": fmt.Sprintf("%08x", addr),
		}).Error("unaligned write32 memory access")

		rot := uint(addr&3) * 8
		val = (val << rot) | (val >> (32 - rot))
		addr &^= 3
	}

	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckTcm(addr); ptr != nil {
			*(*uint32)(ptr) = val
			return
		}
	}

	cpu.bus.Write32(addr, val)
}

func (cpu *Cpu) opRead16(addr uint32) uint16 {
	if addr&1 != 0 {
		log.WithFields(log.Fields{
			"pc":   cpu.pc - 4,
			"addr": fmt.Sprintf("%08x", addr),
		}).Error("unaligned read16 memory access")

		res := cpu.bus.Read16(addr &^ 1)
		return res>>8 | res<<8
	}

	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckTcm(addr); ptr != nil {
			return *(*uint16)(ptr)
		}
	}

	return cpu.bus.Read16(addr)
}

func (cpu *Cpu) opWrite16(addr uint32, val uint16) {
	if addr&1 != 0 {
		log.WithFields(log.Fields{
			"pc":   cpu.pc - 4,
			"addr": fmt.Sprintf("%08x", addr),
		}).Error("unaligned write16 memory access")

		val = val<<8 | val>>8
		addr &^= 1
	}

	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckTcm(addr); ptr != nil {
			*(*uint16)(ptr) = val
			return
		}
	}
	cpu.bus.Write16(addr, val)
}

func (cpu *Cpu) opRead8(addr uint32) uint8 {
	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckTcm(addr); ptr != nil {
			return *(*uint8)(ptr)
		}
	}
	return cpu.bus.Read8(addr)
}

func (cpu *Cpu) opWrite8(addr uint32, val uint8) {
	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckTcm(addr); ptr != nil {
			*(*uint8)(ptr) = val
			return
		}
	}

	cpu.bus.Write8(addr, val)
}
