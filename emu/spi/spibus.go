package spi

import (
	"fmt"
	log "ndsemu/emu/logger"
)

type Bus struct {
	SpiBusName string // Name of the bus (for logging)

	devs  map[int]Device // registered devices
	names map[int]string // name of the device
	tdev  Device         // current device that is transferring data
	req   []byte         // current request
	reply []byte         // current reply
}

func (spi *Bus) AddDevice(addr int, dev Device) {
	if spi.devs == nil {
		spi.devs = make(map[int]Device)
		spi.names = make(map[int]string)
	}

	if _, found := spi.devs[addr]; found {
		panic("spi: address already assigned")
	}
	spi.devs[addr] = dev
	spi.names[addr] = fmt.Sprintf("%T", dev)
}

func (spi *Bus) BeginTransfer(addr int) {
	if spi.tdev != nil {
		if spi.tdev != spi.devs[addr] {
			log.ModSpi.WarnZ("change device during transfer").
				String("bus", spi.SpiBusName).
				Int("device", addr).
				End()
		} else {
			return
		}
	}
	spi.tdev = spi.devs[addr]
	if spi.tdev == nil {
		log.ModSpi.FatalZ("device not implemented").
			String("bus", spi.SpiBusName).
			Int("device", addr).
			End()
	}
	if addr != 2 {
		log.ModSpi.InfoZ("begin transfer").
			String("bus", spi.SpiBusName).
			Int("device", addr).
			String("name", spi.names[addr]).
			End()
	}
	spi.tdev.SpiBegin()

	if spi.req == nil {
		// first time, allocate a buffer, that we will reuse
		spi.req = make([]byte, 16)
	}
	spi.req = spi.req[:0]
	spi.reply = nil
}

func (spi *Bus) Transfer(val byte) (ret byte) {
	if spi.tdev == nil {
		log.ModSpi.WarnZ("SPIDATA written but no transfer").
			String("bus", spi.SpiBusName).
			End()
		return 0
	}

	// First, compute the return value with any pending
	// reply. As described in the documentation, the return
	// value must already be computed at this time, because
	// the byte being written cannot affect the byte that
	// will be simultaneously read
	if len(spi.reply) == 0 {
		ret = 0
	} else {
		ret = spi.reply[0]
		spi.reply = spi.reply[1:]
	}

	// If the current reply was fully sent-out,
	// send the request to the device, as it's
	// possibly a new request to be processed.
	if len(spi.reply) == 0 {
		var stat ReqStatus
		spi.req = append(spi.req, val)
		spi.reply, stat = spi.tdev.SpiTransfer(spi.req)
		if stat == ReqFinish {
			spi.req = spi.req[:0]
		}
	}

	return ret
}

func (spi *Bus) EndTransfer() {
	if spi.tdev == nil {
		log.ModSpi.WarnZ("EndTransfer called but no transfer").
			String("bus", spi.SpiBusName).
			End()
		return
	}
	spi.Reset()
}

func (spi *Bus) Reset() {
	if spi.tdev != nil {
		spi.tdev.SpiEnd()
		log.ModSpi.InfoZ("end transfer").
			String("bus", spi.SpiBusName).
			End()
		spi.tdev = nil
	}
}
