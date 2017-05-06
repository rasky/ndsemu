package logger

import (
	"gopkg.in/Sirupsen/logrus.v0"
)

type ModuleMask uint64
type Module uint

const (
	ModuleMaskAll ModuleMask = 0xFFFFFFFFFFFFFFFF
)

// Predefine a few "common" module constants. The idea is to have a few
// "standard" modules that can be used for easy logging, but it's always
// possible for an emulator to define additional modules through NewModule()
const (
	ModEmu Module = iota + 1
	ModCpu
	ModIrq
	ModMem
	ModSync
	ModHw
	ModHwIo
	ModGfx
	ModSerial
	ModCrypt
	ModDma
	ModTimer
	Mod3d
	ModInput
	ModSound
	endStandardMods
)

var modCount = endStandardMods

var modDebugMask ModuleMask = 0

var modNames = []string{
	"<error>", "emu", "cpu", "irq", "mem", "sync", "hw", "hwio", "gfx",
	"serial", "crypt", "dma", "timer", "3d", "input", "sound",
}

func NewModule(name string) Module {
	mod := modCount
	modCount++
	modNames = append(modNames, name)
	return mod
}

func ModuleByName(name string) (Module, bool) {
	for idx, s := range modNames {
		if s == name {
			return Module(idx), true
		}
	}
	return Module(0xFFFFFFFF), false
}

func EnableDebugModules(mask ModuleMask) {
	modDebugMask |= mask
}

func DisableDebugModules(mask ModuleMask) {
	modDebugMask &^= mask
}

func (mod Module) Mask() ModuleMask {
	return 1 << ModuleMask(mod)
}

func (mod Module) Enabled(level logrus.Level) bool {
	return level <= logrus.WarnLevel || modDebugMask&mod.Mask() != 0
}

// Implement the whole logging interface directly on modules

func (mod Module) WithFields(fields Fields) Entry {
	return Entry{mod: mod}.WithFields(fields)
}

func (mod Module) WithDelayedFields(getfields func() Fields) Entry {
	return Entry{mod: mod}.WithDelayedFields(getfields)
}

func (mod Module) WithField(key string, value interface{}) Entry {
	return Entry{mod: mod}.WithField(key, value)
}

// printf-like family

func (mod Module) Debugf(format string, args ...interface{}) {
	Entry{mod: mod}.Debugf(format, args...)
}

func (mod Module) Printf(format string, args ...interface{}) {
	Entry{mod: mod}.Printf(format, args...)
}

func (mod Module) Infof(format string, args ...interface{}) {
	Entry{mod: mod}.Infof(format, args...)
}

func (mod Module) Warnf(format string, args ...interface{}) {
	Entry{mod: mod}.Warnf(format, args...)
}

func (mod Module) Warningf(format string, args ...interface{}) {
	Entry{mod: mod}.Warningf(format, args...)
}

func (mod Module) Errorf(format string, args ...interface{}) {
	Entry{mod: mod}.Errorf(format, args...)
}

func (mod Module) Fatalf(format string, args ...interface{}) {
	Entry{mod: mod}.Fatalf(format, args...)
}

func (mod Module) Panicf(format string, args ...interface{}) {
	Entry{mod: mod}.Panicf(format, args...)
}

// New-style fast functions

func (mod Module) logz(lvl logrus.Level, msg string) *EntryZ {
	if mod.Enabled(lvl) {
		e := NewEntryZ()
		e.lvl = lvl
		e.msg = msg
		e.mod = mod
		return e
	}
	return nil
}

func (mod Module) DebugZ(msg string) *EntryZ { return mod.logz(logrus.DebugLevel, msg) }
func (mod Module) InfoZ(msg string) *EntryZ  { return mod.logz(logrus.InfoLevel, msg) }
func (mod Module) WarnZ(msg string) *EntryZ  { return mod.logz(logrus.WarnLevel, msg) }
func (mod Module) ErrorZ(msg string) *EntryZ { return mod.logz(logrus.ErrorLevel, msg) }
func (mod Module) FatalZ(msg string) *EntryZ { return mod.logz(logrus.FatalLevel, msg) }
func (mod Module) PanicZ(msg string) *EntryZ { return mod.logz(logrus.PanicLevel, msg) }
