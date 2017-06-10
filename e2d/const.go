package e2d

import (
	log "ndsemu/emu/logger"
)

type HwType int

const (
	HwGba HwType = 0
	HwNds        = 1
)

var modLcd = log.ModGfx

var gKeyState []byte
