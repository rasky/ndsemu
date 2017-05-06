package logger

type LogContextAdder interface {
	// Given a log entry being formed, add some context it
	// (by called WithFields)
	AddLogContext(entry *EntryZ)
}

var contexts []LogContextAdder

func AddContext(c LogContextAdder) {
	contexts = append(contexts, c)
}
