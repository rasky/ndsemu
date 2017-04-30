package arm

import "ndsemu/emu"

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

func (cpu *Cpu) Read32(addr uint32) uint32 {
	if cpu.dbg != nil {
		cpu.dbg.WatchRead(addr)
	}

	// Unaligned memory reads are forcibly aligned.
	// Opcodes behaving different from this default (LDR, SWP, LDRH, LDRSH) are
	// handled within the opcode itself
	addr &^= 3

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

func (cpu *Cpu) Write32(addr uint32, val uint32) {
	if cpu.dbg != nil {
		cpu.dbg.WatchWrite(addr, uint32(val))
	}

	// Unaligned memory writes are forcibly aligned
	addr &^= 3

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
	if cpu.jit != nil {
		cpu.jit.Invalidate(addr)
	}
}

func (cpu *Cpu) Read16(addr uint32) uint16 {
	if cpu.dbg != nil {
		cpu.dbg.WatchRead(addr)
	}

	// Unaligned memory reads are forcibly aligned.
	// Opcodes behaving different from this default (LDR, SWP, LDRH, LDRSH) are
	// handled within the opcode itself
	addr &^= 1

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

func (cpu *Cpu) Write16(addr uint32, val uint16) {
	if cpu.dbg != nil {
		cpu.dbg.WatchWrite(addr, uint32(val))
	}
	cpu.Clock += 1
	// Unaligned memory writes are forcibly aligned
	addr &^= 1

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
	if cpu.jit != nil {
		cpu.jit.Invalidate(addr)
	}
}

func (cpu *Cpu) Read8(addr uint32) uint8 {
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

func (cpu *Cpu) Write8(addr uint32, val uint8) {
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
	if cpu.jit != nil {
		cpu.jit.Invalidate(addr)
	}
}
