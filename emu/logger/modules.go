package logger

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

var modMask ModuleMask = 0 // ModuleMaskAll

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

func EnableModules(mask ModuleMask) {
	modMask |= mask
}

func DisableModules(mask ModuleMask) {
	modMask &^= mask
}

// Implement the whole logging interface directly on modules

func (mod Module) WithFields(fields Fields) Entry {
	return Log(mod).WithFields(fields)
}

func (mod Module) WithDelayedFields(getfields func() Fields) Entry {
	return Log(mod).WithDelayedFields(getfields)
}

func (mod Module) WithField(key string, value interface{}) Entry {
	return Log(mod).WithField(key, value)
}

func (mod Module) Debug(args ...interface{}) {
	Log(mod).Debug(args...)
}

func (mod Module) Print(args ...interface{}) {
	Log(mod).Print(args...)
}

func (mod Module) Info(args ...interface{}) {
	Log(mod).Info(args...)
}

func (mod Module) Warn(args ...interface{}) {
	Log(mod).Warn(args...)
}

func (mod Module) Warning(args ...interface{}) {
	Log(mod).Warning(args...)
}

func (mod Module) Error(args ...interface{}) {
	Log(mod).Error(args...)
}

func (mod Module) Fatal(args ...interface{}) {
	Log(mod).Fatal(args...)
}

func (mod Module) Panic(args ...interface{}) {
	Log(mod).Panic(args...)
}

// printf-like family

func (mod Module) Debugf(format string, args ...interface{}) {
	Log(mod).Debugf(format, args...)
}

func (mod Module) Printf(format string, args ...interface{}) {
	Log(mod).Printf(format, args...)
}

func (mod Module) Infof(format string, args ...interface{}) {
	Log(mod).Infof(format, args...)
}

func (mod Module) Warnf(format string, args ...interface{}) {
	Log(mod).Warnf(format, args...)
}

func (mod Module) Warningf(format string, args ...interface{}) {
	Log(mod).Warningf(format, args...)
}

func (mod Module) Errorf(format string, args ...interface{}) {
	Log(mod).Errorf(format, args...)
}

func (mod Module) Fatalf(format string, args ...interface{}) {
	Log(mod).Fatalf(format, args...)
}

func (mod Module) Panicf(format string, args ...interface{}) {
	Log(mod).Panicf(format, args...)
}

// New-line style family

func (mod Module) Debugln(args ...interface{}) {
	Log(mod).Debugln(args...)
}

func (mod Module) Println(args ...interface{}) {
	Log(mod).Println(args...)
}

func (mod Module) Infoln(args ...interface{}) {
	Log(mod).Infoln(args...)
}

func (mod Module) Warnln(args ...interface{}) {
	Log(mod).Warnln(args...)
}

func (mod Module) Warningln(args ...interface{}) {
	Log(mod).Warningln(args...)
}

func (mod Module) Errorln(args ...interface{}) {
	Log(mod).Errorln(args...)
}

func (mod Module) Fatalln(args ...interface{}) {
	Log(mod).Fatalln(args...)
}

func (mod Module) Panicln(args ...interface{}) {
	Log(mod).Panicln(args...)
}
