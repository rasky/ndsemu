package arm

import (
	"fmt"
	"ndsemu/emu"

	log "gopkg.in/Sirupsen/logrus.v0"
)

// WARNING: This whole file is *very* hot, as it stands between the CPU and
// the memory bus. For every memory access, before accessing the bus, we need:
//
// 	1) Check if there is a debugger installed, and call the wathcpoint
// 	2) Check if the address is misaligned, and handle it the way the CPU does
// 	3) Check if the address falls within DTCM or ITCM (if there is a CP15 and
// 	they are active).
//
// 	The code isn't pretty because it is manually optimized.
// 	DO NOT REFACTOR WITHOUT RUNNING MICRO-BENCHMARKS
//
func (cpu *Cpu) opFetchPointer(addr uint32) []uint8 {
	if cpu.cp15 != nil {
		if ptr := cpu.cp15.CheckITcm(addr); ptr != nil {
			return ptr
		}
	}
	return cpu.bus.FetchPointer(addr)
}

func (cpu *Cpu) opRead32(addr uint32) uint32 {
	if cpu.dbg != nil {
		cpu.dbg.WatchRead(addr)
	}

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
		ptr := cpu.cp15.CheckITcm(addr)
		if ptr == nil {
			ptr = cpu.cp15.CheckDTcm(addr)
			if ptr == nil {
				goto nodtcm
			}
		}

		cpu.Clock += 1
		return emu.Read32LE(ptr)
	}

nodtcm:
	cpu.Clock += cpu.memCycles
	return cpu.bus.Read32(addr)
}

func (cpu *Cpu) opWrite32(addr uint32, val uint32) {
	if cpu.dbg != nil {
		cpu.dbg.WatchWrite(addr, uint32(val))
	}
	if addr&3 != 0 {
		log.WithFields(log.Fields{
			"pc":   cpu.pc - 4,
			"addr": fmt.Sprintf("%08x", addr),
		}).Warn("unaligned write32 memory access")

		// Unaligned memory writes are forcibly aligned
		addr &^= 3
	}

	if cpu.cp15 != nil {
		ptr := cpu.cp15.CheckITcm(addr)
		if ptr == nil {
			ptr = cpu.cp15.CheckDTcm(addr)
			if ptr == nil {
				goto nodtcm
			}
		}

		cpu.Clock += 1
		emu.Write32LE(ptr, val)
		return
	}

nodtcm:
	cpu.Clock += cpu.memCycles
	cpu.bus.Write32(addr, val)
}

func (cpu *Cpu) opRead16(addr uint32) uint16 {
	if cpu.dbg != nil {
		cpu.dbg.WatchRead(addr)
	}
	if addr&1 != 0 {
		log.WithFields(log.Fields{
			"pc":   cpu.pc - 4,
			"addr": fmt.Sprintf("%08x", addr),
		}).Error("unaligned read16 memory access")

		res := cpu.bus.Read16(addr &^ 1)
		return res>>8 | res<<8
	}

	if cpu.cp15 != nil {
		ptr := cpu.cp15.CheckITcm(addr)
		if ptr == nil {
			ptr = cpu.cp15.CheckDTcm(addr)
			if ptr == nil {
				goto nodtcm
			}
		}
		cpu.Clock += 1
		return emu.Read16LE(ptr)
	}

nodtcm:
	cpu.Clock += cpu.memCycles
	return cpu.bus.Read16(addr)
}

func (cpu *Cpu) opWrite16(addr uint32, val uint16) {
	if cpu.dbg != nil {
		cpu.dbg.WatchWrite(addr, uint32(val))
	}
	cpu.Clock += 1
	if addr&1 != 0 {
		log.WithFields(log.Fields{
			"pc":   cpu.pc - 4,
			"addr": fmt.Sprintf("%08x", addr),
		}).Warn("unaligned write16 memory access")

		// Unaligned memory writes are forcibly aligned
		addr &^= 1
	}

	if cpu.cp15 != nil {
		ptr := cpu.cp15.CheckITcm(addr)
		if ptr == nil {
			ptr = cpu.cp15.CheckDTcm(addr)
			if ptr == nil {
				goto nodtcm
			}
		}
		cpu.Clock += 1
		emu.Write16LE(ptr, val)
		return
	}
nodtcm:
	cpu.Clock += cpu.memCycles
	cpu.bus.Write16(addr, val)
}

func (cpu *Cpu) opRead8(addr uint32) uint8 {
	if cpu.dbg != nil {
		cpu.dbg.WatchRead(addr)
	}
	cpu.Clock += 1
	if cpu.cp15 != nil {
		ptr := cpu.cp15.CheckITcm(addr)
		if ptr == nil {
			ptr = cpu.cp15.CheckDTcm(addr)
			if ptr == nil {
				goto nodtcm
			}
		}
		cpu.Clock += 1
		return ptr[0]
	}
nodtcm:
	cpu.Clock += cpu.memCycles
	return cpu.bus.Read8(addr)
}

func (cpu *Cpu) opWrite8(addr uint32, val uint8) {
	if cpu.dbg != nil {
		cpu.dbg.WatchWrite(addr, uint32(val))
	}
	cpu.Clock += 1
	if cpu.cp15 != nil {
		ptr := cpu.cp15.CheckITcm(addr)
		if ptr == nil {
			ptr = cpu.cp15.CheckDTcm(addr)
			if ptr == nil {
				goto nodtcm
			}
		}
		cpu.Clock += 1
		ptr[0] = uint8(val & 0xFF)
		return
	}
nodtcm:
	cpu.Clock += cpu.memCycles
	cpu.bus.Write8(addr, val)
}
