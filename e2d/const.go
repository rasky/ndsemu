package e2d

import (
	log "ndsemu/emu/logger"
)

const (
	cScreenWidth  = 256
	cScreenHeight = 192
)

type HwType int

const (
	HwGba HwType = 0
	HwNds        = 1
)

var modLcd = log.ModGfx

var gKeyState []byte
