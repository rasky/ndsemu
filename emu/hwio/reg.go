package hwio

import (
	"fmt"
	"ndsemu/emu"
	log "ndsemu/emu/logger"
)

type RegFlags uint8

const (
	RegFlagReadOnly RegFlags = (1 << iota)
	RegFlagWriteOnly
)

type Reg64 struct {
	Name   string
	Value  uint64
	RoMask uint64

	Flags   RegFlags
	ReadCb  func(val uint64) uint64
	WriteCb func(old uint64, val uint64)
}

func (reg Reg64) String() string {
	s := fmt.Sprintf("%s{%016x", reg.Name, reg.Value)
	if reg.ReadCb != nil {
		s += ",r!"
	}
	if reg.WriteCb != nil {
		s += ",w!"
	}
	return s + "}"
}

func (reg *Reg64) write(val uint64, romask uint64) {
	romask = romask | reg.RoMask
	old := reg.Value
	reg.Value = (reg.Value & romask) | (val &^ romask)
	if reg.WriteCb != nil {
		reg.WriteCb(old, reg.Value)
	}
}

func (reg *Reg64) Write64(addr uint32, val uint64) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write64 to readonly reg")
		return
	}
	reg.write(val, 0)
}

func (reg *Reg64) Read64(addr uint32) uint64 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read64 from writeonly reg")
		return 0
	}
	if reg.ReadCb != nil {
		return reg.ReadCb(reg.Value)
	}
	return reg.Value
}

func (reg *Reg64) Write32(addr uint32, val uint32) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write32 to readonly reg")
		return
	}
	shift := ((addr & 4) * 8)
	mask := uint64(0xFFFFFFFF) << shift
	reg.write(uint64(val)<<shift, ^mask)
}

func (reg *Reg64) Read32(addr uint32) uint32 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read32 from writeonly reg")
		return 0
	}
	shift := ((addr & 4) * 8)
	return uint32(reg.Read64(addr) >> shift)
}

func (reg *Reg64) Write16(addr uint32, val uint16) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write16 to readonly reg")
		return
	}
	shift := ((addr & 6) * 8)
	mask := uint64(0xFFFF) << shift
	reg.write(uint64(val)<<shift, ^mask)
}

func (reg *Reg64) Read16(addr uint32) uint16 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read16 from writeonly reg")
		return 0
	}
	shift := ((addr & 6) * 8)
	return uint16(reg.Read64(addr) >> shift)
}

func (reg *Reg64) Write8(addr uint32, val uint8) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write8 to readonly reg")
		return
	}
	shift := ((addr & 7) * 8)
	mask := uint64(0xFF) << shift
	reg.write(uint64(val)<<shift, ^mask)
}

func (reg *Reg64) Read8(addr uint32) uint8 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read8 from writeonly reg")
		return 0
	}
	shift := ((addr & 7) * 8)
	return uint8(reg.Read64(addr) >> shift)
}

type Reg32 struct {
	Name   string
	Value  uint32
	RoMask uint32

	Flags   RegFlags
	ReadCb  func(val uint32) uint32
	WriteCb func(old uint32, val uint32)
}

func (reg Reg32) String() string {
	s := fmt.Sprintf("%s{%08x", reg.Name, reg.Value)
	if reg.ReadCb != nil {
		s += ",r!"
	}
	if reg.WriteCb != nil {
		s += ",w!"
	}
	return s + "}"
}

func (reg *Reg32) write(val uint32, romask uint32) {
	romask = romask | reg.RoMask
	old := reg.Value
	reg.Value = (reg.Value & romask) | (val &^ romask)
	if reg.WriteCb != nil {
		reg.WriteCb(old, reg.Value)
	}
}

func (reg *Reg32) Write32(addr uint32, val uint32) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write32 to readonly reg")
		return
	}
	reg.write(val, 0)
}

func (reg *Reg32) Read32(addr uint32) uint32 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read32 from writeonly reg")
		return 0
	}
	if reg.ReadCb != nil {
		return reg.ReadCb(reg.Value)
	}
	return reg.Value
}

func (reg *Reg32) Write16(addr uint32, val uint16) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write16 to readonly reg")
		return
	}
	shift := ((addr & 2) * 8)
	mask := uint32(0xFFFF) << shift
	reg.write(uint32(val)<<shift, ^mask)
}

func (reg *Reg32) Read16(addr uint32) uint16 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read16 from writeonly reg")
		return 0
	}
	shift := ((addr & 2) * 8)
	return uint16(reg.Read32(addr) >> shift)
}

func (reg *Reg32) Write8(addr uint32, val uint8) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write8 to readonly reg")
		return
	}
	shift := ((addr & 3) * 8)
	mask := uint32(0xFF) << shift
	reg.write(uint32(val)<<shift, ^mask)
}

func (reg *Reg32) Read8(addr uint32) uint8 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read8 from writeonly reg")
		return 0
	}
	shift := ((addr & 3) * 8)
	return uint8(reg.Read32(addr) >> shift)
}

type Reg16 struct {
	Name   string
	Value  uint16
	RoMask uint16

	Flags   RegFlags
	ReadCb  func(val uint16) uint16
	WriteCb func(old uint16, val uint16)
}

func (reg Reg16) String() string {
	s := fmt.Sprintf("%s{%04x", reg.Name, reg.Value)
	if reg.ReadCb != nil {
		s += ",r!"
	}
	if reg.WriteCb != nil {
		s += ",w!"
	}
	return s + "}"
}

func (reg *Reg16) write(val uint16, romask uint16) {
	romask = romask | reg.RoMask
	old := reg.Value
	reg.Value = (reg.Value & romask) | (val &^ romask)
	if reg.WriteCb != nil {
		reg.WriteCb(old, reg.Value)
	}
}

func (reg *Reg16) Write16(addr uint32, val uint16) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write16 to readonly reg")
		return
	}
	reg.write(val, 0)
}

func (reg *Reg16) Read16(addr uint32) uint16 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read16 from writeonly reg")
		return 0
	}
	if reg.ReadCb != nil {
		return reg.ReadCb(reg.Value)
	}
	return reg.Value
}

func (reg *Reg16) Write8(addr uint32, val uint8) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write8 to readonly reg")
		return
	}
	shift := ((addr & 1) * 8)
	mask := uint16(0xFF) << shift
	reg.write(uint16(val)<<shift, ^mask)
}

func (reg *Reg16) Read8(addr uint32) uint8 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read8 from writeonly reg")
		return 0
	}
	shift := ((addr & 1) * 8)
	return uint8(reg.Read16(addr) >> shift)
}

type Reg8 struct {
	Name   string
	Value  uint8
	RoMask uint8

	Flags   RegFlags
	ReadCb  func(val uint8) uint8
	WriteCb func(old uint8, val uint8)
}

func (reg Reg8) String() string {
	s := fmt.Sprintf("%s{%02x", reg.Name, reg.Value)
	if reg.ReadCb != nil {
		s += ",r!"
	}
	if reg.WriteCb != nil {
		s += ",w!"
	}
	return s + "}"
}

func (reg *Reg8) write(val uint8, romask uint8) {
	romask = romask | reg.RoMask
	old := reg.Value
	reg.Value = (reg.Value & romask) | (val &^ romask)
	if reg.WriteCb != nil {
		reg.WriteCb(old, reg.Value)
	}
}

func (reg *Reg8) Write8(addr uint32, val uint8) {
	if reg.Flags&RegFlagReadOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Write16 to readonly reg")
		return
	}
	reg.write(val, 0)
}

func (reg *Reg8) Read8(addr uint32) uint8 {
	if reg.Flags&RegFlagWriteOnly != 0 {
		log.ModHwIo.WithFields(log.Fields{
			"name": reg.Name,
			"addr": emu.Hex32(addr),
		}).Error("[IO] invalid Read16 from writeonly reg")
		return 0
	}
	if reg.ReadCb != nil {
		return reg.ReadCb(reg.Value)
	}
	return reg.Value
}
