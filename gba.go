package main

import (
	"ndsemu/e2d"
	"ndsemu/emu"
)

type GbaMemCnt struct {
	Bus emu.Bus
}

func (mc *GbaMemCnt) VramPalette(engine int) []byte {
	return mc.Bus.FetchPointer(0x05000000)
}

func (mc *GbaMemCnt) VramOAM(engine int) []byte {
	return mc.Bus.FetchPointer(0x07000000)
}

func (mc *GbaMemCnt) VramLinearBank(engine int, which e2d.VramLinearBankId, baseOffset int) e2d.VramLinearBank {
	vram := mc.Bus.FetchPointer(0x06000000)
	vram = vram[baseOffset:] // FIXME

	var bank e2d.VramLinearBank
	for i := range bank.Ptr {
		if len(vram) < 8*1024 {
			bank.Ptr[i] = vram
			break
		}
		bank.Ptr[i] = vram[:8*1024 : 8*1024]
		vram = vram[8*1024:]
	}
	return bank
}

func (mc *GbaMemCnt) VramLcdcBank(bank int) []byte {
	panic("LCDC bank access in GBA mode")
}
