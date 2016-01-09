package logger

import (
	"gopkg.in/Sirupsen/logrus.v0"
)

type Fields logrus.Fields

// Like a logrus.Entry, but is nullable. This allows us to selectively disable
// logging while also removing all code overhead associated with it
type Entry struct {
	*logrus.Entry
}

func (entry Entry) WithFields(fields Fields) Entry {
	if entry.Entry != nil {
		return Entry{entry.Entry.WithFields(logrus.Fields(fields))}
	}
	return entry
}

func (entry Entry) WithDelayedFields(getfields func() Fields) Entry {
	if entry.Entry != nil {
		return Entry{entry.Entry.WithFields(logrus.Fields(getfields()))}
	}
	return entry
}

func (entry Entry) WithField(key string, value interface{}) Entry {
	if entry.Entry != nil {
		return Entry{entry.Entry.WithField(key, value)}
	}
	return entry
}

func (entry Entry) Debug(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Debug(args...)
	}
}

func (entry Entry) Print(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Print(args...)
	}
}

func (entry Entry) Info(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Info(args...)
	}
}

func (entry Entry) Warn(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Warn(args...)
	}
}

func (entry Entry) Warning(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Warning(args...)
	}
}

func (entry Entry) Error(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Error(args...)
	}
}

func (entry Entry) Fatal(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Fatal(args...)
	}
}

func (entry Entry) Panic(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Panic(args...)
	}
}

// printf-like family

func (entry Entry) Debugf(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Debugf(format, args...)
	}
}

func (entry Entry) Printf(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Printf(format, args...)
	}
}

func (entry Entry) Infof(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Infof(format, args...)
	}
}

func (entry Entry) Warnf(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Warnf(format, args...)
	}
}

func (entry Entry) Warningf(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Warningf(format, args...)
	}
}

func (entry Entry) Errorf(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Errorf(format, args...)
	}
}

func (entry Entry) Fatalf(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Fatalf(format, args...)
	}
}

func (entry Entry) Panicf(format string, args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Panicf(format, args...)
	}
}

// New-line style family

func (entry Entry) Debugln(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Debugln(args...)
	}
}

func (entry Entry) Println(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Println(args...)
	}
}

func (entry Entry) Infoln(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Infoln(args...)
	}
}

func (entry Entry) Warnln(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Warnln(args...)
	}
}

func (entry Entry) Warningln(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Warningln(args...)
	}
}

func (entry Entry) Errorln(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Errorln(args...)
	}
}

func (entry Entry) Fatalln(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Fatalln(args...)
	}
}

func (entry Entry) Panicln(args ...interface{}) {
	if entry.Entry != nil {
		entry.Entry.Panicln(args...)
	}
}
