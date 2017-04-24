package logger

import (
	"gopkg.in/Sirupsen/logrus.v0"
)

type Fields logrus.Fields

// Like a logrus.Entry, but is nullable. This allows us to selectively disable
// logging while also removing all code overhead associated with it
type Entry struct {
	mod        Module
	lazyfields [8]func() Fields
}

func (entry Entry) log() *logrus.Entry {
	final := logrus.StandardLogger().WithField("_mod", modNames[entry.mod])
	for _, lf := range entry.lazyfields {
		if lf != nil {
			final = final.WithFields(logrus.Fields(lf()))
		}
	}

	fields := make(logrus.Fields, 8)
	for _, c := range contexts {
		c.AddLogContext(fields)
	}
	return final.WithFields(fields)
}

func (entry Entry) WithFields(fields Fields) Entry {
	return entry.WithDelayedFields(func() Fields { return fields })
}

func (entry Entry) WithField(key string, value interface{}) Entry {
	return entry.WithDelayedFields(func() Fields {
		return Fields{
			key: value,
		}
	})
}

func (entry Entry) WithDelayedFields(getfields func() Fields) Entry {
	for idx := range entry.lazyfields {
		if entry.lazyfields[idx] == nil {
			entry.lazyfields[idx] = getfields
			return entry
		}
	}
	return entry
}

func (entry Entry) Debug(args ...interface{}) {
	if entry.mod.Enabled(logrus.DebugLevel) {
		entry.log().Debug(args...)
	}
}

func (entry Entry) Print(args ...interface{}) {
	if entry.mod.Enabled(logrus.InfoLevel) {
		entry.log().Print(args...)
	}
}

func (entry Entry) Info(args ...interface{}) {
	if entry.mod.Enabled(logrus.InfoLevel) {
		entry.log().Info(args...)
	}
}

func (entry Entry) Warn(args ...interface{}) {
	if entry.mod.Enabled(logrus.WarnLevel) {
		entry.log().Warn(args...)
	}
}

func (entry Entry) Warning(args ...interface{}) {
	if entry.mod.Enabled(logrus.WarnLevel) {
		entry.log().Warning(args...)
	}
}

func (entry Entry) Error(args ...interface{}) {
	if entry.mod.Enabled(logrus.ErrorLevel) {
		entry.log().Error(args...)
	}
}

func (entry Entry) Fatal(args ...interface{}) {
	if entry.mod.Enabled(logrus.FatalLevel) {
		entry.log().Fatal(args...)
	}
}

func (entry Entry) Panic(args ...interface{}) {
	if entry.mod.Enabled(logrus.PanicLevel) {
		entry.log().Panic(args...)
	}
}

// printf-like family

func (entry Entry) Debugf(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.DebugLevel) {
		entry.log().Debugf(format, args...)
	}
}

func (entry Entry) Printf(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.InfoLevel) {
		entry.log().Printf(format, args...)
	}
}

func (entry Entry) Infof(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.InfoLevel) {
		entry.log().Infof(format, args...)
	}
}

func (entry Entry) Warnf(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.WarnLevel) {
		entry.log().Warnf(format, args...)
	}
}

func (entry Entry) Warningf(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.WarnLevel) {
		entry.log().Warningf(format, args...)
	}
}

func (entry Entry) Errorf(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.ErrorLevel) {
		entry.log().Errorf(format, args...)
	}
}

func (entry Entry) Fatalf(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.FatalLevel) {
		entry.log().Fatalf(format, args...)
	}
}

func (entry Entry) Panicf(format string, args ...interface{}) {
	if entry.mod.Enabled(logrus.PanicLevel) {
		entry.log().Panicf(format, args...)
	}
}

// New-line style family

func (entry Entry) Debugln(args ...interface{}) {
	if entry.mod.Enabled(logrus.DebugLevel) {
		entry.log().Debugln(args...)
	}
}

func (entry Entry) Println(args ...interface{}) {
	if entry.mod.Enabled(logrus.InfoLevel) {
		entry.log().Println(args...)
	}
}

func (entry Entry) Infoln(args ...interface{}) {
	if entry.mod.Enabled(logrus.InfoLevel) {
		entry.log().Infoln(args...)
	}
}

func (entry Entry) Warnln(args ...interface{}) {
	if entry.mod.Enabled(logrus.WarnLevel) {
		entry.log().Warnln(args...)
	}
}

func (entry Entry) Warningln(args ...interface{}) {
	if entry.mod.Enabled(logrus.WarnLevel) {
		entry.log().Warningln(args...)
	}
}

func (entry Entry) Errorln(args ...interface{}) {
	if entry.mod.Enabled(logrus.ErrorLevel) {
		entry.log().Errorln(args...)
	}
}

func (entry Entry) Fatalln(args ...interface{}) {
	if entry.mod.Enabled(logrus.FatalLevel) {
		entry.log().Fatalln(args...)
	}
}

func (entry Entry) Panicln(args ...interface{}) {
	if entry.mod.Enabled(logrus.PanicLevel) {
		entry.log().Panicln(args...)
	}
}
