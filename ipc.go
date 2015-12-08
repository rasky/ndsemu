package main

import (
	"fmt"

	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwIpc struct {
	data [2]uint16
}

func (ipc *HwIpc) WriteIPCSYNC(cpunum CpuNum, value uint16) {
	ipc.data[cpunum] = value
	if value&(1<<13) != 0 || value&(1<<14) != 0 {
		panic("not implemented: IPCSYNC IRQ emulation")
	}
	log.WithFields(log.Fields{
		"cpu":   cpunum,
		"value": fmt.Sprintf("%04x", ipc.data[cpunum]),
	}).Info("[IPC] Write sync")
}

func (ipc *HwIpc) ReadIPCSYNC(cpunum CpuNum) uint16 {
	value := ipc.data[cpunum]
	value = (value &^ 0xF) | ((ipc.data[1-cpunum] >> 8) & 0xF)
	log.WithFields(log.Fields{
		"cpu":   cpunum,
		"value": fmt.Sprintf("%04x", value),
	}).Info("[IPC] Read sync")
	return value
}
