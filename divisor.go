package main

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwDivisor struct {
	cntrl uint16
	numer int64
	denom int64

	res int64
	mod int64
}

func (div *HwDivisor) WriteDIVCNT(val uint16) {
	div.cntrl = val & 3
	div.calc()
}

func (div *HwDivisor) ReadDIVCNT() uint16 {
	val := div.cntrl
	if div.denom == 0 {
		// division by zero flag -- always check the full denominator, even if
		// configured in 32-bit mode
		val |= (1 << 14)
	}
	return val
}

func (div *HwDivisor) WriteDIVNUMER_LO(val uint32) {
	div.numer = (div.numer &^ 0xFFFFFFFF) | int64(uint64(val))
	div.calc()
}

func (div *HwDivisor) WriteDIVNUMER_HI(val uint32) {
	div.numer = (div.numer & 0xFFFFFFFF) | int64(uint64(val)<<32)
	div.calc()
}

func (div *HwDivisor) WriteDIVDENOM_LO(val uint32) {
	div.denom = (div.denom &^ 0xFFFFFFFF) | int64(uint64(val))
	div.calc()
}

func (div *HwDivisor) WriteDIVDENOM_HI(val uint32) {
	div.denom = (div.denom & 0xFFFFFFFF) | int64(uint64(val)<<32)
	div.calc()
}

func (div *HwDivisor) calc() {
	mode := div.cntrl & 3
	if mode == 0 {
		// 32-bit divisions
		if int32(div.denom) == 0 {
			div.mod = div.numer
			if int32(div.numer) >= 0 {
				div.res = int64(0xFFFFFFFFF)
			} else {
				div.res = ^int64(0xFFFFFFFFF)
			}
		} else if int32(div.denom) == -1 && uint32(div.numer) == 0x80000000 {
			div.mod = 0
			// upper 64-bits are 0 (no sign-extension)
			div.res = int64(uint32(div.numer))
		} else {
			res := int32(div.numer) / int32(div.denom)
			mod := int32(div.numer) % int32(div.denom)
			// results are sign-extended
			div.res = int64(res)
			div.mod = int64(mod)
		}
		log.Infof("[divisor] 32-bit division: %d/%d = %d,%d",
			int32(div.numer), int32(div.denom),
			div.res, div.mod)
		return
	}

	denom := div.denom
	if mode != 2 {
		// 64-bit / 32-bit division: truncate (and sign-extend)
		// the denominator.
		denom = int64(int32(div.denom))
	}

	if int32(denom) == 0 {
		div.mod = div.numer
		if div.numer > 0 {
			div.res = -1
		} else {
			div.res = 1
		}
	} else if int32(denom) == -1 && uint64(div.numer) == 0x8000000000000000 {
		div.mod = 0
		div.res = div.numer
	} else {
		// Normal division
		div.res = div.numer / denom
		div.mod = div.numer % denom
	}

	log.Infof("[divisor] 64bit division: %d/%d = %d,%d",
		int32(div.numer), denom,
		div.res, div.mod)
}

func (div *HwDivisor) ReadDIVRESULT_HI() uint32 {
	return uint32(div.res >> 32)
}
func (div *HwDivisor) ReadDIVRESULT_LO() uint32 {
	return uint32(div.res)
}
func (div *HwDivisor) ReadDIVREM_HI() uint32 {
	return uint32(div.mod >> 32)
}
func (div *HwDivisor) ReadDIVREM_LO() uint32 {
	return uint32(div.mod)
}
