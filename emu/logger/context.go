package logger

import logrus "gopkg.in/Sirupsen/logrus.v0"

type LogContextAdder interface {
	// Given a log entry being formed, add some context it
	// (by called WithFields)
	AddLogContext(fields logrus.Fields)
}

var contexts []LogContextAdder

func AddContext(c LogContextAdder) {
	contexts = append(contexts, c)
}
