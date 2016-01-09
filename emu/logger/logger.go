package logger

import (
	"gopkg.in/Sirupsen/logrus.v0"
)

func beginLogging(mod Module) Entry {
	entry := Entry{logrus.StandardLogger().WithField("_mod", modNames[mod])}
	for _, c := range contexts {
		entry = c.AddLogContext(entry)
	}
	return entry
}

func Log(mod Module) Entry {
	// NOTE: keep this function light. We want this function to be
	// inlined in all call sites, as much as possible, so that skipping
	// disabled modules become very quick.

	// Check if the module is enabled or not. If not, return a nil
	// Entry that will effectively do nothing
	if modMask&ModuleMask(1<<mod) != 0 {
		return beginLogging(mod)
	}
	return Entry{nil}
}
