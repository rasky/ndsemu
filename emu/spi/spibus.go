package spi

import log "ndsemu/emu/logger"

var modSpi = log.NewModule("spi")

type Bus struct {
	devs  map[int]Device // registered devices
	tdev  Device         // current device that is transferring data
	req   []byte         // current request
	reply []byte         // current reply
}

func (spi *Bus) AddDevice(addr int, dev Device) {
	if spi.devs == nil {
		spi.devs = make(map[int]Device)
	}

	if _, found := spi.devs[addr]; found {
		panic("spi: address already assigned")
	}
	spi.devs[addr] = dev
}

func (spi *Bus) BeginTransfer(addr int) {
	if spi.tdev != nil {
		if spi.tdev != spi.devs[addr] {
			modSpi.Warnf("wrong new device=%d", addr)
			// panic("SPI changed device during transfer")
		} else {
			return
		}
	}
	spi.tdev = spi.devs[addr]
	if spi.tdev == nil {
		modSpi.Fatalf("SPI device %d not implemented", addr)
	}
	modSpi.Infof("begin transfer device=%d (%T)", addr, spi.tdev)
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
		modSpi.Warn("SPIDATA written but no transfer")
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
		modSpi.Warn("EndTransfer called but no transfer")
		return
	}
	spi.tdev.SpiEnd()
	spi.tdev = nil
	modSpi.Info("end of transfer")
}
