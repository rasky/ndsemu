package e2d

import (
	log "ndsemu/emu/logger"
)

const (
	cScreenWidth  = 256
	cScreenHeight = 192
)

var modLcd = log.NewModule("2d")

var gKeyState []byte
