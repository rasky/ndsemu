package main

import (
	"encoding/binary"
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

type HwSoundChannel struct {
	SndCnt hwio.Reg32 `hwio:"offset=0x00,wcb"`
	SndSad hwio.Reg32 `hwio:"offset=0x04,rwmask=0x07FFFFFF"`
	SndTmr hwio.Reg16 `hwio:"offset=0x08"`
	SndPnt hwio.Reg16 `hwio:"offset=0x0A"`
	SndLen hwio.Reg32 `hwio:"offset=0x0C,rwmask=0x001FFFFF"`

	snd *HwSound
	idx int
}

type HwSound struct {
	Bus emu.Bus

	Ch [16]HwSoundChannel

	voice [16]struct {
		mem  []byte
		pos  uint
		step uint
		on   bool
		mode int
	}

	SndGCnt hwio.Reg32 `hwio:"bank=1,offset=0x0"`
	// The NDS7 BIOS brings this register to 0x200 at boot, with a slow loop
	// with delay that takes ~1 second. If we reset it at 0x200, it will just
	// skip everything and the emulator will boot faster.
	SndBias    hwio.Reg32 `hwio:"bank=1,offset=0x4,reset=0x200,rwmask=0x3FF"`
	SndCap0Cnt hwio.Reg8  `hwio:"bank=1,offset=0x8,rwmask=0x8F"`
	SndCap1Cnt hwio.Reg8  `hwio:"bank=1,offset=0x9,rwmask=0x8F"`
}

func NewHwSound(bus emu.Bus) *HwSound {
	snd := new(HwSound)
	snd.Bus = bus
	for i := 0; i < 16; i++ {
		hwio.MustInitRegs(&snd.Ch[i])
		snd.Ch[i].snd = snd
		snd.Ch[i].idx = i
	}
	hwio.MustInitRegs(snd)
	return snd
}

func (ch *HwSoundChannel) WriteSNDCNT(old, new uint32) {
	if (old^new)&(1<<31) != 0 {
		if new&(1<<31) != 0 {
			ch.snd.startChannel(ch.idx)
		} else {
			ch.snd.stopChannel(ch.idx)
		}
	}
}

func (snd *HwSound) startChannel(idx int) {
	ch := &snd.Ch[idx]
	v := &snd.voice[idx]

	ptr := snd.Bus.FetchPointer(ch.SndSad.Value)
	mode := int((ch.SndCnt.Value >> 29) & 3)
	length := uint32(ch.SndPnt.Value)*4 + ch.SndLen.Value*4

	freq := cBusClock / 2 / int64((-int16(ch.SndTmr.Value)))
	if freq < 0 {
		panic("negative frequency?")
	}
	v.on = true
	v.mem = ptr[:length]
	v.pos = 0
	v.step = uint((freq << 16) / cAudioFreq)
	v.mode = mode
	if v.mode == 2 {
		v.mem = snd.adpcmDecompress(v.mem)
		v.mode = 1
	}

	if idx == 4 {
		numsteps := 0
		if v.step != 0 {
			numsteps = (len(v.mem) << 16) / 2 / int(v.step)
		}
		fmt.Printf("prelen=%d, len=%d, freq=%d, step=%.2f, numsteps=%d\n", length, len(v.mem), freq, float64(v.step)/65536.0, numsteps)
	}

	log.ModSound.WithField("ch", idx).WithField("freq", freq).Info("start channel")
}

func (snd *HwSound) stopChannel(idx int) {
	v := &snd.voice[idx]
	v.on = false
	snd.Ch[idx].SndCnt.Value &^= 1 << 31
	log.ModSound.WithField("ch", idx).Info("stop channel")
}

var (
	voldiv          = [4]uint32{0, 1, 2, 4}
	adpcmIndexTable = [8]int16{-1, -1, -1, -1, 2, 4, 6, 8}
	adpcmTable      = [89]uint16{
		0x0007, 0x0008, 0x0009, 0x000A, 0x000B, 0x000C, 0x000D, 0x000E, 0x0010, 0x0011, 0x0013, 0x0015,
		0x0017, 0x0019, 0x001C, 0x001F, 0x0022, 0x0025, 0x0029, 0x002D, 0x0032, 0x0037, 0x003C, 0x0042,
		0x0049, 0x0050, 0x0058, 0x0061, 0x006B, 0x0076, 0x0082, 0x008F, 0x009D, 0x00AD, 0x00BE, 0x00D1,
		0x00E6, 0x00FD, 0x0117, 0x0133, 0x0151, 0x0173, 0x0198, 0x01C1, 0x01EE, 0x0220, 0x0256, 0x0292,
		0x02D4, 0x031C, 0x036C, 0x03C3, 0x0424, 0x048E, 0x0502, 0x0583, 0x0610, 0x06AB, 0x0756, 0x0812,
		0x08E0, 0x09C3, 0x0ABD, 0x0BD0, 0x0CFF, 0x0E4C, 0x0FBA, 0x114C, 0x1307, 0x14EE, 0x1706, 0x1954,
		0x1BDC, 0x1EA5, 0x21B6, 0x2515, 0x28CA, 0x2CDF, 0x315B, 0x364B, 0x3BB9, 0x41B2, 0x4844, 0x4F7E,
		0x5771, 0x602F, 0x69CE, 0x7462, 0x7FFF,
	}
)

func (snd *HwSound) adpcmDecompress(buf []byte) []byte {
	// ioutil.WriteFile("sound.adpcm", buf, 0666)
	head := binary.LittleEndian.Uint32(buf[:4])
	buf = buf[4:]
	pcm := int32(int16(head & 0xFFFF))
	index := int16(head>>16) & 0x7F

	res := make([]byte, 0, len(buf)*4)

	dec := func(sample uint8) {
		diff := adpcmTable[index] / 8
		if sample&1 != 0 {
			diff += adpcmTable[index] / 4
		}
		if sample&2 != 0 {
			diff += adpcmTable[index] / 2
		}
		if sample&4 != 0 {
			diff += adpcmTable[index] / 1
		}
		if sample&8 == 0 {
			pcm += int32(diff)
			if pcm > 0x7FFF {
				pcm = 0x7FFF
			}
		} else {
			pcm -= int32(diff)
			if pcm < -0x7FFF {
				pcm = -0x7FFF
			}
		}

		index += adpcmIndexTable[sample&7]
		if index < 0 {
			index = 0
		} else if index > 88 {
			index = 88
		}
	}

	for i := range buf {
		dec(buf[i] & 0xF)
		res = append(res, uint8(pcm&0xFF))
		res = append(res, uint8((pcm>>8)&0xFF))
		dec(buf[i] >> 4)
		res = append(res, uint8(pcm&0xFF))
		res = append(res, uint8((pcm>>8)&0xFF))
	}

	return res
}

func (snd *HwSound) RunOneFrame(buf []int16) {
	for i := 0; i < len(buf); i += 2 {
		l, r := snd.step()

		// Extend to 16-bit range
		l <<= 6
		r <<= 6

		// Convert to signed
		buf[i] = int16(l - 0x8000)
		buf[i+1] = int16(r - 0x8000)
	}
}

func mulvol64(s int64, vol int64) int64 {
	if vol == 127 {
		return s
	}
	return (s * vol) >> 7
}

// Emulate one tick of audio, producing a couple of (unsigned) 16-bit audio samples
func (snd *HwSound) step() (uint16, uint16) {
	var lmix, rmix int64

	for i := 0; i < 16; i++ {
		var sample int64

		cntrl := snd.Ch[i].SndCnt.Value
		voice := &snd.voice[i]

		if !voice.on {
			continue
		}

		pos := voice.pos >> 16
		switch voice.mode {
		case 0:
			if int(pos) >= len(voice.mem) {
				snd.stopChannel(i)
				continue
			}
			sample = int64(int8(voice.mem[pos])) << 8
		case 1:
			if int(pos*2+1) >= len(voice.mem) {
				snd.stopChannel(i)
				continue
			}
			sample = int64(int16(binary.LittleEndian.Uint16(voice.mem[pos*2:])))
		case 2:
			log.ModSound.WithField("ch", i).Info("unsupported sound format ADPCM")
			snd.stopChannel(i)
			continue
		case 3:
			if i < 8 {
				panic("unsupported sound format #3 in channel 0-7")
			} else if i < 14 {
				panic("unsupported sound format PSG")
			} else {
				panic("unsupported sound format noise")
			}
		}
		voice.pos += voice.step

		sample <<= 16

		// Apply volume divider
		sample >>= voldiv[(cntrl>>8)&3]

		// Apply channel volume
		sample = mulvol64(sample, int64(cntrl&127))

		// Apply panning
		pan := int64((cntrl >> 16) & 127)
		lsample := mulvol64(sample, 127-pan)
		rsample := mulvol64(sample, pan)

		// Mix
		lmix += int64(lsample &^ 0xFF)
		rmix += int64(rsample &^ 0xFF)
	}

	// Apply master volume
	gvol := int64(snd.SndGCnt.Value & 127)
	lmix = mulvol64(lmix, gvol)
	rmix = mulvol64(rmix, gvol)

	// Adjust volume after mixing
	lmix >>= 6
	rmix >>= 6

	// Round
	lmix >>= 16
	rmix >>= 16

	// Bias
	lmix += int64(snd.SndBias.Value)
	rmix += int64(snd.SndBias.Value)

	// Clamp
	if lmix < 0 {
		lmix = 0
	} else if lmix > 0x3FF {
		lmix = 0x3FF
	}
	if rmix < 0 {
		rmix = 0
	} else if rmix > 0x3FF {
		rmix = 0x3FF
	}

	return uint16(lmix), uint16(rmix)
}
