package emu

import (
	"gopkg.in/Sirupsen/logrus.v0"
)

var Log func() *logrus.Entry

func init() {
	std := logrus.StandardLogger()
	Log = func() *logrus.Entry {
		return logrus.NewEntry(std)
	}
}
