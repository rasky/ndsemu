package main

import (
	"encoding/binary"
	"hash/crc64"
	"ndsemu/emu"
	"ndsemu/emu/hw"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"

	"github.com/hashicorp/golang-lru/simplelru"
)

const (
	kLoopManual   = 0
	kLoopInfinite = 1
	kLoopOneShot  = 2

	kMode8bit     = 0
	kMode16bit    = 1
	kModeAdpcm    = 2
	kModePsgNoise = 3

	kPosNoLoop uint = ^uint(0)
)

// Checksum table used to quickly hash a voice to get a key to cache it
var ctable = crc64.MakeTable(crc64.ECMA)

type HwSoundChannel struct {
	SndCnt hwio.Reg32 `hwio:"offset=0x00,wcb"`
	SndSad hwio.Reg32 `hwio:"offset=0x04,rwmask=0x07FFFFFF"`
	SndTmr hwio.Reg16 `hwio:"offset=0x08,wcb"`
	SndPnt hwio.Reg16 `hwio:"offset=0x0A"`
	SndLen hwio.Reg32 `hwio:"offset=0x0C,rwmask=0x001FFFFF"`

	snd *HwSound
	idx int
}

type HwSound struct {
	Bus emu.Bus

	Ch [16]HwSoundChannel

	voice [16]struct {
		mem   []byte
		tmr   uint32
		pos   uint
		on    bool
		mode  int
		loop  int
		delay int
	}

	capture [2]struct {
		on     bool
		tmr    uint32
		reset  uint32
		wpos   uint32
		loop   bool
		bit8   bool
		add    bool
		single bool
		regdad *uint32
		reglen *uint32
	}

	cache *simplelru.LRU

	SndGCnt hwio.Reg32 `hwio:"bank=1,offset=0x0"`
	// The NDS7 BIOS brings this register to 0x200 at boot, with a slow loop
	// with delay that takes ~1 second. If we reset it at 0x200, it will just
	// skip everything and the emulator will boot faster.
	SndBias    hwio.Reg32 `hwio:"bank=1,offset=0x4,reset=0x200,rwmask=0x3FF"`
	SndCap0Cnt hwio.Reg8  `hwio:"bank=1,offset=0x8,rwmask=0x8F,wcb"`
	SndCap1Cnt hwio.Reg8  `hwio:"bank=1,offset=0x9,rwmask=0x8F,wcb"`
	SndCap0Dad hwio.Reg32 `hwio:"bank=1,offset=0x10,writeonly"`
	SndCap1Dad hwio.Reg32 `hwio:"bank=1,offset=0x18,writeonly"`
	SndCap0Len hwio.Reg32 `hwio:"bank=1,offset=0x14"`
	SndCap1Len hwio.Reg32 `hwio:"bank=1,offset=0x1C"`
}

func NewHwSound(bus emu.Bus) *HwSound {
	cache, err := simplelru.NewLRU(128, nil)
	if err != nil {
		panic(err)
	}

	snd := new(HwSound)
	snd.Bus = bus
	snd.cache = cache
	for i := 0; i < 16; i++ {
		hwio.MustInitRegs(&snd.Ch[i])
		snd.Ch[i].snd = snd
		snd.Ch[i].idx = i
	}
	snd.capture[0].regdad = &snd.SndCap0Dad.Value
	snd.capture[1].regdad = &snd.SndCap1Dad.Value
	snd.capture[0].reglen = &snd.SndCap0Len.Value
	snd.capture[1].reglen = &snd.SndCap1Len.Value
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

func (ch *HwSoundChannel) WriteSNDTMR(_, new uint16) {
	// Å write to SNDTMR also takes effect while the voice is playing
	// so copy the value into the latched register we increment at every tick.
	ch.snd.voice[ch.idx].tmr = uint32(new)
}

func (snd *HwSound) startChannel(idx int) {
	ch := &snd.Ch[idx]
	v := &snd.voice[idx]

	ptr := snd.Bus.FetchPointer(ch.SndSad.Value)
	mode := int((ch.SndCnt.Value >> 29) & 3)
	length := uint32(ch.SndPnt.Value)*4 + ch.SndLen.Value*4
	loop := int((ch.SndCnt.Value >> 27) & 3)

	if ch.SndCnt.Value&(1<<15) != 0 {
		panic("hold")
	}

	v.on = false // will put true at the end of the function, if no error
	v.mem = ptr[:length]
	v.pos = 0
	v.delay = 3
	v.tmr = uint32(ch.SndTmr.Value)
	v.mode = mode
	v.loop = loop

	var sum uint64
	switch v.mode {
	case kModeAdpcm:
		v.delay = 11
		sum = crc64.Checksum(v.mem, ctable)
		if buf, found := snd.cache.Get(sum); found {
			v.mem = buf.([]byte)
		} else {
			v.mem = snd.adpcmDecompress(v.mem)
			// go ioutil.WriteFile(fmt.Sprintf("%x.raw", sum), v.mem, 0666)
			snd.cache.Add(sum, v.mem)
		}
	case kModePsgNoise:
		v.delay = 1
		if idx >= 8 || idx <= 13 {
			// Mode PSG
			v.mem = psgTable[(ch.SndCnt.Value>>24)&3][:]
		} else {
			log.ModSound.WithField("ch", idx).Error("unsupported PSG/noise mode on this channel")
			return
		}
	}

	if ch.SndCnt.Value&(1<<15) != 0 {
		panic("hold value")
	}

	log.ModSound.InfoZ("start channel").
		Int("ch", idx).
		Int("mode", mode).
		Hex32("rpos", ch.SndSad.Value).
		Uint32("len", length).
		Uint("ptlen", uint(ch.SndPnt.Value)*4).
		Hex64("sum", sum).
		Int("loop", loop).
		Hex16("tmr", ch.SndTmr.Value).
		Int64("clk", nds7.Cycles()).
		End()
	v.on = true
}

func (snd *HwSound) stopChannel(idx int) {
	v := &snd.voice[idx]
	v.on = false
	snd.Ch[idx].SndCnt.Value &^= 1 << 31
	log.ModSound.InfoZ("stop channel").Int("idx", idx).End()
}

func (snd *HwSound) loopChannel(idx int) uint {
	if snd.voice[idx].loop == kLoopInfinite {
		off := snd.Ch[idx].SndPnt.Value * 4
		switch snd.voice[idx].mode {
		case kModeAdpcm:
			off -= 4
			fallthrough
		case kMode16bit:
			off /= 2
		}
		return uint(off)
	}
	return kPosNoLoop
}

func (snd *HwSound) WriteSNDCAP0CNT(old, new uint8) { snd.writeSNDCAPCNT(0, old, new) }
func (snd *HwSound) WriteSNDCAP1CNT(old, new uint8) { snd.writeSNDCAPCNT(1, old, new) }
func (snd *HwSound) writeSNDCAPCNT(idx int, old, new uint8) {
	if (old^new)&(1<<7) != 0 {
		if new&(1<<7) != 0 {
			snd.startCapture(idx, new)
		} else {
			snd.stopCapture(idx, new)
		}
	}
}

func (snd *HwSound) startCapture(idx int, cnt uint8) {
	cap := &snd.capture[idx]
	cap.on = true
	cap.loop = cnt&(1<<2) == 0
	cap.bit8 = cnt&(1<<3) != 0
	cap.single = cnt&(1<<1) != 0
	cap.add = cnt&(1<<1) != 0
	cap.wpos = *cap.regdad
	cap.reset = uint32(snd.Ch[idx*2+1].SndTmr.Value)
	cap.tmr = cap.reset
	log.ModSound.InfoZ("start capture").
		Int("idx", idx).
		Bool("loop", cap.loop).
		Bool("8bit", cap.bit8).
		Bool("single", cap.single).
		Bool("add", cap.add).
		Hex32("wpos", cap.wpos).
		Hex32("wlen", *cap.reglen*4).
		Hex16("tmr", uint16(cap.reset)).
		Int64("clk", nds7.Cycles()).
		End()
}

func (snd *HwSound) stopCapture(idx int, cnt uint8) {
	cap := &snd.capture[idx]
	cap.on = false
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
	psgTable = [8][16]uint8{
		{0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0xff, 0x7f}, // _______-
		{0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0xff, 0x7f, 0xff, 0x7f}, // ______--
		{0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f}, // _____---
		{0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f}, // ____----
		{0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f}, // ___-----
		{0x01, 0x80, 0x01, 0x80, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f}, // __------
		{0x01, 0x80, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f, 0xff, 0x7f}, // _-------
		{0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80, 0x01, 0x80}, // ________
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
		diff += (adpcmTable[index] / 4) * uint16((sample>>0)&1)
		diff += (adpcmTable[index] / 2) * uint16((sample>>1)&1)
		diff += (adpcmTable[index] / 1) * uint16((sample>>2)&1)
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
		l = l<<6 | l>>4
		r = r<<6 | r>>4

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
	var chbuf [4]int64

	// Master enable
	if snd.SndGCnt.Value&(1<<15) == 0 {
		return uint16(snd.SndBias.Value), uint16(snd.SndBias.Value)
	}

	scans := []int{
		hw.SCANCODE_0,
		hw.SCANCODE_1,
		hw.SCANCODE_2,
		hw.SCANCODE_3,
		hw.SCANCODE_4,
		hw.SCANCODE_5,
		hw.SCANCODE_6,
		hw.SCANCODE_7,
		hw.SCANCODE_8,
		hw.SCANCODE_9,
	}

	keys := hw.GetKeyboardState()
	mask := 0xFFFF
	pressed := 0
	for i := 0; i < len(scans); i++ {
		if keys[scans[i]] != 0 {
			pressed |= 1 << uint(i)
		}
	}
	if pressed != 0 {
		mask = pressed
	}

	for i := 0; i < 16; i++ {
		var sample int64

		cntrl := snd.Ch[i].SndCnt.Value
		voice := &snd.voice[i]

		if !voice.on {
			continue
		}
		if mask&(1<<uint(i)) == 0 {
			continue
		}

		voice.tmr += cTimerStepPerSample
		for voice.tmr >= 0x10000 {
			if voice.delay >= 0 {
				voice.delay--
			} else {
				voice.pos++
			}
			voice.tmr = uint32(snd.Ch[i].SndTmr.Value) + (voice.tmr - 0x10000)
		}
		if voice.delay >= 0 {
			continue
		}

		switch voice.mode {
		case kMode8bit:
			if int(voice.pos) >= len(voice.mem) {
				voice.pos = voice.pos + snd.loopChannel(i) - uint(len(voice.mem))
				if voice.pos == kPosNoLoop {
					snd.stopChannel(i)
					continue
				}
			}
			sample = int64(int8(voice.mem[voice.pos])) << 8
		case kMode16bit, kModeAdpcm:
			if int(voice.pos*2+1) >= len(voice.mem) {
				voice.pos = voice.pos + snd.loopChannel(i) - uint(len(voice.mem)/2)
				if voice.pos == kPosNoLoop || int(voice.pos*2+1) >= len(voice.mem) {
					snd.stopChannel(i)
					continue
				}
			}
			sample = int64(int16(binary.LittleEndian.Uint16(voice.mem[voice.pos*2:])))
		case kModePsgNoise:
			for int(voice.pos*2+1) >= len(voice.mem) {
				voice.pos -= uint(len(voice.mem)) / 2
			}
			sample = int64(int16(binary.LittleEndian.Uint16(voice.mem[voice.pos*2:])))
		}

		// Convert into fixed point to keep some precision
		sample <<= 8

		// Apply volume divider
		sample >>= voldiv[(cntrl>>8)&3]

		// Apply channel volume
		sample = mulvol64(sample, int64(cntrl&127))

		if i < 4 {
			// Save copy of channels used in capture
			chbuf[i] = sample

			// Check specific "Channel 1/3 disable" bit
			if i == 1 && snd.SndGCnt.Value&(1<<12) != 0 {
				continue
			}
			if i == 3 && snd.SndGCnt.Value&(1<<13) != 0 {
				continue
			}
		}

		// Apply panning
		pan := int64((cntrl >> 16) & 127)
		lsample := mulvol64(sample, 127-pan)
		rsample := mulvol64(sample, pan)

		// Mix
		lmix += int64(lsample)
		rmix += int64(rsample)
	}

	// Handle capture
	for i := 0; i < 2; i++ {
		cap := &snd.capture[i]
		if cap.on {
			var sample int64
			if !cap.single {
				if i == 0 {
					sample = lmix
				} else {
					sample = rmix
				}
				if sample > 0x7FFF00 {
					sample = 0x7FFF00
				}
				if sample < -0x800000 {
					sample = -0x800000
				}
			} else {
				sample = chbuf[i*2]
			}
			if cap.add {
				panic("capture with addition")
			}

			cap.tmr += cTimerStepPerSample
			for cap.tmr >= 0x10000 {
				if cap.bit8 {
					snd.Bus.Write8(cap.wpos, uint8(sample>>16))
					cap.wpos++
				} else {
					snd.Bus.Write16(cap.wpos, uint16(sample>>8))
					cap.wpos += 2
				}

				cap.tmr = uint32(cap.reset) + (cap.tmr - 0x10000)
				if cap.wpos >= *cap.regdad+*cap.reglen*4 {
					if cap.loop {
						cap.wpos = *cap.regdad
					} else {
						cap.on = false
					}
				}
			}
		}
	}

	switch (snd.SndGCnt.Value >> 8) & 3 {
	case 1:
		lmix = chbuf[1]
	case 2:
		lmix = chbuf[3]
	case 3:
		lmix = chbuf[1] + chbuf[3]
	}
	switch (snd.SndGCnt.Value >> 10) & 3 {
	case 1:
		rmix = chbuf[1]
	case 2:
		rmix = chbuf[3]
	case 3:
		rmix = chbuf[1] + chbuf[3]
	}

	// Apply master volume
	gvol := int64(snd.SndGCnt.Value & 127)
	lmix = mulvol64(lmix, gvol)
	rmix = mulvol64(rmix, gvol)

	// Adjust volume after mixing
	lmix >>= 6
	rmix >>= 6

	// Convert from fixed into integer (strip fraction)
	lmix >>= 8
	rmix >>= 8

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
