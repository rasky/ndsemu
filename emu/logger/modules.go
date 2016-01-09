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
	ModHwIo
	ModGfx
	ModSerial
	ModCrypt
	ModDma
	ModTimer
	Mod3d
	ModInput
	endStandardMods
)

var modCount = endStandardMods

var modDebugMask ModuleMask = 0

var modNames = []string{
	"<error>", "emu", "cpu", "irq", "mem", "sync", "hwio", "gfx",
	"serial", "crypt", "dma", "timer", "3d", "input",
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

func (mod Module) Enabled(level logrus.Level) bool {
	return level <= logrus.WarnLevel || modDebugMask&(1<<mod) != 0
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

func (mod Module) Debug(args ...interface{}) {
	Entry{mod: mod}.Debug(args...)
}

func (mod Module) Print(args ...interface{}) {
	Entry{mod: mod}.Print(args...)
}

func (mod Module) Info(args ...interface{}) {
	Entry{mod: mod}.Info(args...)
}

func (mod Module) Warn(args ...interface{}) {
	Entry{mod: mod}.Warn(args...)
}

func (mod Module) Warning(args ...interface{}) {
	Entry{mod: mod}.Warning(args...)
}

func (mod Module) Error(args ...interface{}) {
	Entry{mod: mod}.Error(args...)
}

func (mod Module) Fatal(args ...interface{}) {
	Entry{mod: mod}.Fatal(args...)
}

func (mod Module) Panic(args ...interface{}) {
	Entry{mod: mod}.Panic(args...)
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

// New-line style family

func (mod Module) Debugln(args ...interface{}) {
	Entry{mod: mod}.Debugln(args...)
}

func (mod Module) Println(args ...interface{}) {
	Entry{mod: mod}.Println(args...)
}

func (mod Module) Infoln(args ...interface{}) {
	Entry{mod: mod}.Infoln(args...)
}

func (mod Module) Warnln(args ...interface{}) {
	Entry{mod: mod}.Warnln(args...)
}

func (mod Module) Warningln(args ...interface{}) {
	Entry{mod: mod}.Warningln(args...)
}

func (mod Module) Errorln(args ...interface{}) {
	Entry{mod: mod}.Errorln(args...)
}

func (mod Module) Fatalln(args ...interface{}) {
	Entry{mod: mod}.Fatalln(args...)
}

func (mod Module) Panicln(args ...interface{}) {
	Entry{mod: mod}.Panicln(args...)
}
