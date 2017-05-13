package logger

import (
	"io"
	"os"

	"gopkg.in/Sirupsen/logrus.v0"
)

type textFormatter struct {
	logrus.TextFormatter
}

var output io.Writer
var outIsTerminal bool

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

func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
	output = out
}

func init() {
	SetOutput(os.Stdout)
	outIsTerminal = IsTerminal(1)
	logrus.SetFormatter(&textFormatter{logrus.TextFormatter{}})
}
