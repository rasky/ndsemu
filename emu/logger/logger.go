package logger

import (
	"os"

	"gopkg.in/Sirupsen/logrus.v0"
)

type textFormatter struct {
	logrus.TextFormatter
}

func (f *textFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if mod, found := entry.Data["_mod"]; found {
		entry.Message = "[" + mod.(string) + "] " + entry.Message
		delete(entry.Data, "_mod")
	}
	return f.TextFormatter.Format(entry)
}

func Disable() {
	logrus.SetLevel(logrus.PanicLevel)
}

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&textFormatter{logrus.TextFormatter{}})
}
