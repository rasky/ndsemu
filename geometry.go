package main

import (
	"encoding/binary"
	"fmt"
	"ndsemu/emu"
	"ndsemu/emu/hwio"
	log "ndsemu/emu/logger"
	"ndsemu/raster3d"
)

var modGxFifo = log.NewModule("gxfifo")
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
	// Bank 0 (0x4000400). Main geometry FIFO
	GxFifo  hwio.Mem `hwio:"bank=0,offset=0x0,size=4,vsize=0x40,rw8=off,rw16=off,wcb"`
	GxCmd   hwio.Mem `hwio:"bank=0,offset=0x40,size=4,vsize=0x190,rw8=off,rw16=off,wcb"`
	ClipMtx hwio.Mem `hwio:"bank=0,offset=0x240,size=0x40,readonly"`
	DirMtx  hwio.Mem `hwio:"bank=0,offset=0x280,size=0x40,readonly"`

	// Bank 1 (0x4000600). Status and results
	GxStat   hwio.Reg32 `hwio:"bank=1,offset=0,rwmask=0xC0008000,rcb,wcb"`
	RamCount hwio.Reg32 `hwio:"bank=1,offset=4,readonly,rcb"`

	// Bank 2 (0x4000300). Various tables and parameters
	Edge0 hwio.Reg16 `hwio:"bank=2,offset=0x30,writeonly"`
	Edge1 hwio.Reg16 `hwio:"bank=2,offset=0x32,writeonly"`
	Edge2 hwio.Reg16 `hwio:"bank=2,offset=0x34,writeonly"`
	Edge3 hwio.Reg16 `hwio:"bank=2,offset=0x36,writeonly"`
	Edge4 hwio.Reg16 `hwio:"bank=2,offset=0x38,writeonly"`
	Edge5 hwio.Reg16 `hwio:"bank=2,offset=0x3A,writeonly"`
	Edge6 hwio.Reg16 `hwio:"bank=2,offset=0x3C,writeonly"`
	Edge7 hwio.Reg16 `hwio:"bank=2,offset=0x3E,writeonly"`

	fifoRegCmd uint32
	fifoRegCnt int

	irq    *HwIrq
	gx     GeometryEngine
	busy   bool
	cycles int64
	fifo   []GxCmd
}

func NewHwGeometry(irq *HwIrq, e3d *raster3d.HwEngine3d) *HwGeometry {
	g := new(HwGeometry)
	g.irq = irq
	g.gx.E3dCmdCh = e3d.CmdCh
	hwio.MustInitRegs(g)
	return g
}

func (g *HwGeometry) ReadGXSTAT(val uint32) uint32 {
	// Sync to the current CPU cycle, so that we return an accurate
	// value
	g.Run(Emu.Sync.Cycles())

	// FIXME: for now, always return OK to "box test" command (not implemented)
	val |= 1 << 1

	if g.fifoLessThanHalfFull() {
		val |= (1 << 25)
	}
	if g.fifoEmpty() {
		val |= (1 << 26) // empty
	}
	if g.busy {
		val |= (1 << 27) // busy bit
	}

	// Bits 16-24: Entries in the FIFO
	val |= (uint32(len(g.fifo)) & 0x1ff) << 16

	// Bits 8-12: Position matrix stack (only 5 bits)
	val |= (uint32(g.gx.mtxStackPosPtr) & 0x1F) << 8

	// Bit 13: Projection matrix stack (1 bit)
	val |= (uint32(g.gx.mtxStackProjPtr) & 0x1) << 13

	// Bit 14: true if there is a PUSH/POP command pending
	for _, cmd := range g.fifo {
		if cmd.code == GX_MTX_POP || cmd.code == GX_MTX_PUSH {
			val |= (1 << 14)
			break
		}
	}

	modGxFifo.Infof("read GXSTAT: %08x", val)
	return val
}

func (g *HwGeometry) ReadRAMCOUNT(_ uint32) uint32 {
	vtx := Emu.Hw.E3d.NumVertices()
	poly := Emu.Hw.E3d.NumPolygons()

	if vtx > 2048 {
		vtx = 2048
	}
	if poly > 6144 {
		poly = 6144
	}

	// FIXME: used by zelda-hourglass, but we can't
	// implement it properly until clipping is implemented,
	// otherwise the returned value makes the game hang
	//
	// modGxFifo.Warnf("RAMCOUNT: %d,%d", vtx, poly)
	return 0 //uint32(vtx)<<16 | uint32(poly)
}

func (g *HwGeometry) WriteGXSTAT(old, val uint32) {
	g.GxStat.Value &^= 0x8000
	if val&0x8000 != 0 {
		old &^= 0x8000
		// Also reset matrix stat pointer
		g.gx.mtxStackProjPtr = 0
	}
	g.GxStat.Value |= old & 0x8000
	modGxFifo.Infof("write GXSTAT: %08x", val)
}

func (g *HwGeometry) WriteGXFIFO(addr uint32, bytes int) {
	if bytes != 4 {
		modGxFifo.Error("non 32-bit write to GXFIFO")
	}

	val := binary.LittleEndian.Uint32(g.GxFifo.Data[0:4])
	modGxFifo.WithFields(log.Fields{
		"val":    emu.Hex32(val),
		"curcmd": emu.Hex32(g.fifoRegCmd),
		"curcnt": g.fifoRegCnt,
	}).Info("write to GXFIFO")

	// If there is a command that's waiting for arguments, then
	// this is one of the arguments; send it to the FIFO right away
	if g.fifoRegCnt != 0 {
		g.fifoPush(uint8(g.fifoRegCmd&0xFF), val)
		g.fifoRegCnt -= 1
		if g.fifoRegCnt > 0 {
			return
		}
		// Process next packed command
		g.fifoRegCmd >>= 8
	} else {
		// Otherwise this is a new packed command
		g.fifoRegCmd = val
	}

	// Unpack next command. Notice that we don't treat "unpacked"
	// commands differently: after all, they're just like a
	// packed command contained just one command.
	for g.fifoRegCmd != 0 {
		nextcmd := uint8(g.fifoRegCmd & 0xFF)
		if int(nextcmd) < len(gxCmdDescs) {
			if gxCmdDescs[nextcmd].exec == nil {
				modGxFifo.Fatalf("packed command not implemented: %02x", nextcmd)
			}
			g.fifoRegCnt = gxCmdDescs[nextcmd].parms
		} else {
			g.fifoRegCnt = 0
			modGxFifo.Fatalf("invalid packed command: %02x", nextcmd)
		}

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

func (g *HwGeometry) updateIrq() {
	// Update the IRQ level. Notice that the GX FIFO IRQ is level-triggered
	// so the line stays set for the whole time the condition is true.
	switch g.GxStat.Value >> 30 {
	case 1:
		g.irq.Assert(IrqGxFifo, g.fifoLessThanHalfFull())
	case 2:
		g.irq.Assert(IrqGxFifo, g.fifoEmpty())
	}
}

func (g *HwGeometry) fifoEmpty() bool {
	// FIXME: this doesn't include the 4-slot pipe into account.
	return len(g.fifo) == 0
}

func (g *HwGeometry) fifoLessThanHalfFull() bool {
	// FIXME: this doesn't include the 4-slot pipe into account.
	// We should probably increase the fifo by 4 entries, and then use
	// "< 128+4" here.
	return len(g.fifo) < 128
}

func (g *HwGeometry) fifoPush(code uint8, parm uint32) {
	if len(g.fifo) < 256 {
		cmd := GxCmd{
			when: Emu.Sync.Cycles(),
			code: GxCmdCode(code),
			parm: parm,
		}
		modGxFifo.WithField("val", fmt.Sprintf("%02x-%08x", code, parm)).Info("gxfifo push")
		g.fifo = append(g.fifo, cmd)
		g.updateIrq()
	} else {
		modGxFifo.Errorf("attempt to push to full GX FIFO")
	}
}

func (g *HwGeometry) WriteGXCMD(addr uint32, bytes int) {
	if bytes != 4 {
		modGxFifo.Error("non 32-bit write to GXCMD")
	}

	val := binary.LittleEndian.Uint32(g.GxCmd.Data[0:4])
	modGxFifo.WithField("val", emu.Hex32(val)).WithField("addr", emu.Hex32(addr)).Infof("Write GXCMD")
	cmd := uint8((addr-0x4000440)/4 + 0x10)
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
	if g.fifoLessThanHalfFull() {
		nds9.TriggerDmaEvent(DmaEventGxFifo)
	}

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
			break
		}

		// Check if we need to simulate waiting for the last command
		// to arrive. This might be needed if the CPU was slower at
		// sending commands compared to the geometry engine to execute.
		tlastarg := g.fifo[nparms].when
		if g.cycles < tlastarg {
			g.cycles = tlastarg
		}

		// SwapBuffer is special, because it always waits for a VBlank before
		// beginning execution, plus its own 392 cycles. To simulate this
		// behaviour, if this is the first time we see the command (so it's
		// the moment at which it's beginning execution), calculate the
		// timing of nextvblank, and adjust the "when" attribute of the command;
		// this will make the rest of the code assume that the command arrived
		// at that moment in the FIFO, and thus wait until then to synchronize
		if cmd.code == GX_SWAP_BUFFERS && cmd.parm&(1<<31) == 0 {
			g.fifo[0].parm |= 1 << 31
			g.fifo[0].when = g.cycles + Emu.Sync.DotPosDistance(0, 192)
			continue
		}

		// Check if we can execute the command in this timeslice
		if g.cycles+cycles > target {
			// Not enough cycles in this timeslice; mark the geometry
			// engine as busy because the timeslice ends in the middle
			// of a command computation.
			g.busy = true
			goto end
		}

		if desc.exec == nil {
			modGx.WithField("cmd", g.fifo[0].code).Error("unimplemented command")
		} else {
			modGx.WithField("cmd", g.fifo[0].code).Info("exec command")
			desc.exec(&g.gx, g.fifo[:nparms+1])
		}

		g.fifo = g.fifo[nparms+1:]
		g.cycles += cycles
		if g.fifoLessThanHalfFull() {
			nds9.TriggerDmaEvent(DmaEventGxFifo)
		}
	}

	// We get here if there are no more full commands to execute in the FIFO,
	// even though there was some time available in this timeslice.
	g.cycles = target
	g.busy = false

end:
	g.updateIrq()
	for i := 0; i < 16; i++ {
		v := g.gx.clipmtx[i/4][i%4].V
		binary.LittleEndian.PutUint32(g.ClipMtx.Data[i*4:i*4+4], uint32(v))
	}
	for i := 0; i < 9; i++ {
		v := g.gx.mtx[MtxDirection][i/3][i%3].V
		binary.LittleEndian.PutUint32(g.DirMtx.Data[i*4:i*4+4], uint32(v))
	}
}
