package main

import (
	log "gopkg.in/Sirupsen/logrus.v0"
)

type HwIrq struct {
	master uint32
	enable uint32
	ack    uint32
}

func (irq *HwIrq) ReadIME() uint32 {
	return irq.master
}

func (irq *HwIrq) ReadIE() uint32 {
	return irq.enable
}

func (irq *HwIrq) ReadIF() uint32 {
	return irq.ack
}

func (irq *HwIrq) WriteIME(ime uint32) {
	irq.master = ime & 1
	log.Infof("[irq] IME:%x", irq.master)
}

func (irq *HwIrq) WriteIE(ie uint32) {
	irq.enable = ie
	log.Infof("[irq] IE:%x", irq.enable)
}

func (irq *HwIrq) WriteIF(ifx uint32) {
	log.Infof("[irq] IF:%x", ifx)
}
