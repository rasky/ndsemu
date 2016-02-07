package main

import (
	"encoding/binary"
	"ndsemu/emu"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
)

var modGx = log.NewModule("gx")

type GxCmdCode uint8

type GxCmd struct {
	when int64     // instant at which this cmd was pushed into the fifo
	code GxCmdCode // 8-bit cmd code
	parm uint32    // 32-bit cmd arg
}

type GxCmdDesc struct {
	parms   int
	ncycles int64
	exec    func(*GeometryEngine, []GxCmd)
}

type HwGeometry struct {
	GxFifo hwio.Mem   `hwio:"bank=0,offset=0x0,size=4,vsize=0x40,rw8=off,rw16=off,wcb"`
	GxCmd  hwio.Mem   `hwio:"bank=0,offset=0x40,size=4,vsize=0x190,rw8=off,rw16=off,wcb"`
	GxStat hwio.Reg32 `hwio:"bank=1,offset=0,rwmask=0xC0008000,rcb,wcb"`

	fifoRegCmd uint32
	fifoRegCnt int

	gx     GeometryEngine
	busy   bool
	cycles int64
	fifo   []GxCmd
}

func NewHwGeometry() *HwGeometry {
	g := new(HwGeometry)
	hwio.MustInitRegs(g)
	return g
}

func (g *HwGeometry) ReadGXSTAT(val uint32) uint32 {
	// Sync to the current CPU cycle, so that we return an accurate
	// value
	g.Run(Emu.Sync.Cycles())

	if len(g.fifo) < 128 {
		val |= (1 << 25) // less-than-half-full
	}
	if len(g.fifo) == 0 {
		val |= (1 << 26) // empty
	}
	val |= (uint32(len(g.fifo)) & 0x1ff) << 16
	if g.busy {
		val |= (1 << 27) // busy bit
	}
	modGx.Infof("read GXSTAT: %08x", val)
	return val
}

func (g *HwGeometry) WriteGXSTAT(old, val uint32) {
	g.GxStat.Value &^= 0x8000
	if val&0x8000 != 0 {
		old &^= 0x8000
		// TODO: also reset matrix stat pointer
	}
	g.GxStat.Value |= old & 0x8000
	modGx.Infof("write GXSTAT: %08x", val)
	if val&0xc0000000 != 0 {
		modGx.Fatal("IRQ GX FIFO not emulated")
	}
}

func (g *HwGeometry) WriteGXFIFO(addr uint32, bytes int) {
	if bytes != 4 {
		modGx.Error("non 32-bit write to GXFIFO")
	}

	val := binary.LittleEndian.Uint32(g.GxFifo.Data[0:4])
	modGx.WithField("val", emu.Hex32(val)).Error("write to GXFIFO")

	// If there is a command that's waiting for arguments, then
	// this is one of the arguments; send it to the FIFO right away
	if g.fifoRegCnt != 0 {
		g.fifoPush(uint8(g.fifoRegCmd&0xFF), val)
		g.fifoRegCnt -= 1
		if g.fifoRegCnt == 0 {
			g.fifoRegCmd >>= 8
		}
	} else {
		// Otherwise this is a new packed command
		g.fifoRegCmd = val
	}

	// Unpack next command. Notice that we don't treat "unpacked"
	// commands differently: after all, they're just like a
	// packed command contained just one command.
	for g.fifoRegCmd != 0 {
		nextcmd := uint8(g.fifoRegCmd & 0xFF)
		g.fifoRegCnt = gxCmdDescs[nextcmd].parms

		// If it requires argument, exit; we need to wait for them
		if g.fifoRegCnt != 0 {
			break
		}

		// No arguments: send it straight away to the fifo, and
		// restart loop unpacking next command
		g.fifoPush(nextcmd, 0)
		g.fifoRegCmd >>= 8
	}
}

func (g *HwGeometry) fifoPush(code uint8, parm uint32) {
	if len(g.fifo) < 256 {
		cmd := GxCmd{
			when: Emu.Sync.Cycles(),
			code: GxCmdCode(code),
			parm: parm,
		}
		g.fifo = append(g.fifo, cmd)
	} else {
		modGx.Errorf("attempt to push to full GX FIFO")
	}
}

func (g *HwGeometry) WriteGXCMD(addr uint32, bytes int) {
	if bytes != 4 {
		modGx.Error("non 32-bit write to GXCMD")
	}

	val := binary.LittleEndian.Uint32(g.GxCmd.Data[0:4])
	cmd := uint8((addr-0x4000440)/4 + 0x10)
	modGx.Infof("write gxcmd: push %02x %08x", cmd, val)
	g.fifoPush(cmd, val)
}

func (g *HwGeometry) Reset() {
	g.fifo = nil
	g.cycles = 0
	g.busy = false
}

func (g *HwGeometry) Frequency() emu.Fixed8 {
	return emu.NewFixed8(cBusClock)
}

func (g *HwGeometry) Cycles() int64 {
	return g.cycles
}

func (g *HwGeometry) Run(target int64) {
	for len(g.fifo) > 0 {
		// Peek first command in the FIFO
		cmd := g.fifo[0]
		desc := &gxCmdDescs[cmd.code]
		cycles := g.gx.CalcCmdCycles(cmd.code)

		// Compute the number of fifo entry for arguments
		// in addition to the first one.
		nparms := desc.parms
		if nparms > 0 {
			nparms -= 1
		}

		// Check if all parameters are available, otherwise we
		// can't execute it.
		if len(g.fifo) < 1+nparms {
			g.busy = false
			g.cycles = target
			return
		}

		// Check if we need to simulate waiting for the last command
		// to arrive. This might be needed if the CPU was slower at
		// sending commands compared to the geometry engine to execute.
		tlastarg := g.fifo[nparms].when
		if g.cycles < tlastarg {
			g.cycles = tlastarg
		}

		// Check if we can execute the command in this timeslice
		if g.cycles+cycles > target {
			// Not enough cycles in this timeslice; mark the geometry
			// engine as busy because the timeslice ends in the middle
			// of a command computation.
			g.busy = true
			return
		}

		if desc.exec == nil {
			modGx.WithField("cmd", g.fifo[0].code).Error("unimplemented command")
		} else {
			modGx.WithField("cmd", g.fifo[0].code).Info("exec command")
			desc.exec(&g.gx, g.fifo[:nparms+1])
		}

		g.fifo = g.fifo[nparms+1:]
		g.cycles += cycles
	}
}
