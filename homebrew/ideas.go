package homebrew

import (
	"ndsemu/arm"
	log "ndsemu/emu/logger"
)

var modIdeas = log.NewModule("ideas")

// Activate support for debugging with IDEAS protocol.
// The IDEAS protocol uses the SWI call number 0xFC as a debug "puts()" call.
func ActivateIdeasDebug(cpu *arm.Cpu) {
	cpu.SetSwiHle(0xFC, ideasLog)
	log.EnableDebugModules(modIdeas.Mask())
}

func ideasLog(cpu *arm.Cpu) int64 {
	var s []byte
	for i := uint32(0); i < 1024; i++ {
		ch := cpu.Read8(uint32(cpu.Regs[0]) + i)
		if ch == 0 {
			break
		}
		s = append(s, ch)
	}
	if len(s) > 0 {
		if s[len(s)-1] == '\n' {
			s = s[:len(s)-1]
		}
		modIdeas.Info(string(s))
	}
	return 0
}
