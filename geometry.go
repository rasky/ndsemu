package main

import (
	"encoding/binary"
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

type GxFifo struct {
	cmds [256]GxCmd
	r    int64
	w    int64
}

func (f *GxFifo) Reset() {
	f.r = 0
	f.w = 0
}

func (f *GxFifo) Push(cmd GxCmd) {
	if f.Full() {
		panic("gxfifo push full")
	}
	f.cmds[f.w&255] = cmd
	f.w++
}

func (f *GxFifo) Pop() GxCmd {
	if f.r >= f.w {
		panic("gxfifo pop empty")
	}
	cmd := f.cmds[f.r&255]
	f.r++
	return cmd
}

func (f *GxFifo) Top() *GxCmd {
	if f.Empty() {
		return nil
	}
	return &f.cmds[f.r&255]
}

func (f *GxFifo) HasCmdTest() bool {
	for i := f.r; i < f.w; i++ {
		cmd := &f.cmds[i&255]
		if cmd.code == GX_VEC_TEST || cmd.code == GX_POS_TEST || cmd.code == GX_BOX_TEST {
			return true
		}
	}
	return false
}

func (f *GxFifo) HasCmdMatrix() bool {
	for i := f.r; i < f.w; i++ {
		cmd := &f.cmds[i&255]
		if cmd.code == GX_MTX_PUSH || cmd.code == GX_MTX_POP {
			return true
		}
	}
	return false
}

// FIXME: these function don't include the 4-slot pipe into account.
// We should probably increase the fifo by 4 entries
func (f *GxFifo) Len() int               { return int(f.w - f.r) }
func (f *GxFifo) Empty() bool            { return f.r >= f.w }
func (f *GxFifo) Full() bool             { return f.Len() >= 256 }
func (f *GxFifo) LessThanHalfFull() bool { return f.Len() < 128 }

type HwGeometry struct {
	// Bank 0 (0x4000400). Main geometry FIFO
	GxFifo  hwio.Mem `hwio:"bank=0,offset=0x0,size=4,vsize=0x40,rw8=off,rw16=off,wcb"`
	GxCmd   hwio.Mem `hwio:"bank=0,offset=0x40,size=4,vsize=0x190,rw8=off,rw16=off,wcb"`
	ClipMtx hwio.Mem `hwio:"bank=0,offset=0x240,size=0x40,readonly"`
	DirMtx  hwio.Mem `hwio:"bank=0,offset=0x280,size=0x40,readonly"`

	// Bank 1 (0x4000600). Status and results
	GxStat     hwio.Reg32 `hwio:"bank=1,offset=0,rwmask=0xC0008000,rcb,wcb"`
	RamCount   hwio.Reg32 `hwio:"bank=1,offset=4,readonly,rcb"`
	VecResultX hwio.Reg16 `hwio:"bank=1,offset=0x30,readonly,rcb"`
	VecResultY hwio.Reg16 `hwio:"bank=1,offset=0x32,readonly,rcb"`
	VecResultZ hwio.Reg16 `hwio:"bank=1,offset=0x34,readonly,rcb"`

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
	fifo   GxFifo

	// Static buffer for a single gxcmd extacted from the FIFO
	// We use a static buffer to avoid allocating every time (as we
	// process one command at a time)
	curcmd [36]GxCmd

	// Debug statistics on one frame's worth of GX FIFO commands
	framestats struct {
		start  int64 // cycles at which the frame started (last SwapBuffers)
		numcmd int   // number of commands sent
	}
}

func NewHwGeometry(irq *HwIrq, e3d *raster3d.HwEngine3d) *HwGeometry {
	g := new(HwGeometry)
	g.irq = irq
	g.gx.E3dCmdCh = e3d.CmdCh
	hwio.MustInitRegs(g)
	// FIXME: these callbacks were already populated through reflection, but the resulting
	// runtime trampoline is much slower (and allocates!). Since these registers are written
	// A LOT during each frame, manually bind the callback.
	g.GxFifo.WriteCb = g.WriteGXFIFO
	g.GxCmd.WriteCb = g.WriteGXCMD
	return g
}

func (g *HwGeometry) ReadGXSTAT(val uint32) uint32 {
	// Sync to the current CPU cycle, so that we return an accurate
	// value
	g.Run(Emu.Sync.Cycles())

	// Bit 0: true if there is a box/pos/vec test pending
	if g.fifo.HasCmdTest() {
		val |= (1 << 0)
	}

	// FIXME: for now, always return OK to "box test" command (not implemented)
	val |= 1 << 1

	if g.fifo.LessThanHalfFull() {
		val |= (1 << 25)
	}
	if g.fifo.Empty() {
		val |= (1 << 26) // empty
	}
	if g.busy {
		val |= (1 << 27) // busy bit
	}

	// Bits 16-24: Entries in the FIFO
	val |= uint32(g.fifo.Len()&0x1ff) << 16

	// Bits 8-12: Position matrix stack (only 5 bits)
	val |= (uint32(g.gx.mtxStackPosPtr) & 0x1F) << 8

	// Bit 13: Projection matrix stack (1 bit)
	val |= (uint32(g.gx.mtxStackProjPtr) & 0x1) << 13

	// Bit 14: true if there is a PUSH/POP command pending
	if g.fifo.HasCmdMatrix() {
		val |= (1 << 14)
	}

	// Bit 15: true if there was a matrix stack overflow
	if g.gx.mtxStackOverflow {
		val |= (1 << 15)
	}

	// modGxFifo.Infof("read GXSTAT: %08x", val)
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

func (g *HwGeometry) readVecResult(vec emu.Fixed12) uint16 {
	// Convert into a 16-bit value, with 4-bits of sign and 8-bit of fractional part
	// First thing, copy the 12-bits fraction into the result.
	n := uint16(vec.V & 0xFFF)

	// Then check the integer part; we can only represent 0 or -1, any other value
	// overflows. So basically, the top 4 bits are:
	//   integer == 0 -> 0x0
	//   integer > 0 -> 0xF
	//   integer == -1 -> 0xF
	//   integer < -1 -> 0x0
	i := int32(vec.V) >> 12
	if i > 0 || i == -1 {
		n |= 0xF000
	}
	return n
}

func (g *HwGeometry) ReadVECRESULTX(_ uint16) uint16 { return g.readVecResult(g.gx.vecTestResult[0]) }
func (g *HwGeometry) ReadVECRESULTY(_ uint16) uint16 { return g.readVecResult(g.gx.vecTestResult[1]) }
func (g *HwGeometry) ReadVECRESULTZ(_ uint16) uint16 { return g.readVecResult(g.gx.vecTestResult[2]) }

func (g *HwGeometry) WriteGXSTAT(old, val uint32) {
	g.GxStat.Value &^= 0x8000
	if val&0x8000 != 0 {
		// turn off overflow flag
		old &^= 0x8000
		// Also reset matrix stat pointer
		g.gx.mtxStackProjPtr = 0
	}
	g.GxStat.Value |= old & 0x8000
}

func (g *HwGeometry) WriteGXFIFO(addr uint32, bytes int) {

	if bytes != 4 {
		modGxFifo.Error("non 32-bit write to GXFIFO")
	}

	now := Emu.Sync.Cycles()
	val := binary.LittleEndian.Uint32(g.GxFifo.Data[0:4])
	// modGxFifo.WithFields(log.Fields{
	// 	"val":    emu.Hex32(val),
	// 	"curcmd": emu.Hex32(g.fifoRegCmd),
	// 	"curcnt": g.fifoRegCnt,
	// }).Info("write to GXFIFO")

	// If there is a command that's waiting for arguments, then
	// this is one of the arguments; send it to the FIFO right away
	if g.fifoRegCnt != 0 {
		g.fifoPush(now, uint8(g.fifoRegCmd&0xFF), val)
		g.fifoRegCnt -= 1
		if g.fifoRegCnt > 0 {
			g.updateIrq()
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
		g.fifoPush(now, nextcmd, 0)
		g.fifoRegCmd >>= 8
	}
	g.updateIrq()
}

func (g *HwGeometry) updateIrq() {
	// Update the IRQ level. Notice that the GX FIFO IRQ is level-triggered
	// so the line stays set for the whole time the condition is true.
	switch g.GxStat.Value >> 30 {
	case 1:
		g.irq.Assert(IrqGxFifo, g.fifo.LessThanHalfFull())
	case 2:
		g.irq.Assert(IrqGxFifo, g.fifo.Empty())
	default:
		g.irq.Assert(IrqGxFifo, false)
	}
}

func (g *HwGeometry) fifoPush(when int64, code uint8, parm uint32) {
	// If the FIFO is full, try synchronize the geometry processor
	// up to the current timestamp. This might be enough to flush
	// the FIFO a little bit and make room for the new command.
	if g.fifo.Full() {
		g.Run(Emu.Sync.Cycles())
	}

	// If the FIFO is still full, it means that the CPU is
	// really writing to a full FIFO. The CPU will be blocked
	// until the FIFO frees up a space.
	panicCount := 0
	for g.fifo.Full() {
		// Burn CPU cycles that should be enough to execute
		// the next FIFO command.
		cycles := g.nextCmdCycles()
		if cycles == 0 {
			// Technically, there shouldn't be 0-cycles commands,
			// but unknown commands are marked as 0 while it probably
			// takes at least 1 cycle to the geometry engine to pull
			// them from the FIFO and skip them.
			cycles = 1
		}
		nds9.Cpu.Clock += cycles * 2

		// Now synchronize the geometry engine. Since the CPU has
		// burnt some cycles, this should allow us to run a little
		// bit and consume the FIFO.
		g.Run(Emu.Sync.Cycles())

		// Debug counter to avoid infinite loop: if the FIFO is not
		// consumed, it's a bug in our code, just abort.
		panicCount++
		if panicCount == 1024 {
			panic("stalled geometry engine?")
		}
	}

	cmd := GxCmd{
		when: when,
		code: GxCmdCode(code),
		parm: parm,
	}
	g.fifo.Push(cmd)
	// modGxFifo.WithField("val", fmt.Sprintf("%02x-%08x", code, parm)).WithField("len", len(g.fifo)).Infof("gxfifo push")
	// g.updateIrq()
}

func (g *HwGeometry) WriteGXCMD(addr uint32, bytes int) {
	if bytes != 4 {
		modGxFifo.Error("non 32-bit write to GXCMD")
	}

	val := binary.LittleEndian.Uint32(g.GxCmd.Data[0:4])
	// modGxFifo.WithField("val", emu.Hex32(val)).WithField("addr", emu.Hex32(addr)).Infof("Write GXCMD")
	cmd := uint8((addr-0x4000440)/4 + 0x10)
	now := Emu.Sync.Cycles()
	g.fifoPush(now, cmd, val)
	g.updateIrq()
}

func (g *HwGeometry) Reset() {
	g.fifo.Reset()
	g.cycles = 0
	g.busy = false
}

func (g *HwGeometry) Frequency() emu.Fixed8 {
	return emu.NewFixed8(cBusClock)
}

func (g *HwGeometry) Cycles() int64 {
	return g.cycles
}

func (g *HwGeometry) nextCmdCycles() int64 {
	if cmd := g.fifo.Top(); cmd != nil {
		return g.gx.CalcCmdCycles(cmd.code)
	}
	return 0
}

func (g *HwGeometry) Run(target int64) {
	if g.fifo.LessThanHalfFull() {
		// modGxFifo.WithField("fifolen", len(g.fifo)).Info("trigger GXFIFO DMA")
		nds9.TriggerDmaEvent(DmaEventGxFifo)
	}

	for g.cycles < target {
		// Peek first command in the FIFO
		cmd := g.fifo.Top()
		if cmd == nil {
			g.busy = false
			break
		}

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
		if g.fifo.Len() < 1+nparms {
			g.busy = false
			break
		}

		// Extract parameters
		for i := 0; i < nparms+1; i++ {
			g.curcmd[i] = g.fifo.Pop()
		}

		// Check if we need to simulate waiting for the last command
		// to arrive. This might be needed if the CPU was slower at
		// sending commands compared to the geometry engine to execute.
		tlastarg := g.curcmd[nparms].when
		if g.cycles < tlastarg {
			g.cycles = tlastarg
		}

		if desc.exec == nil {
			// modGx.WithField("cmd", g.fifo[0].code).Error("unimplemented command")
		} else {
			// modGx.WithField("cmd", g.fifo[0].code).Info("exec command")
			desc.exec(&g.gx, g.curcmd[:nparms+1])
		}

		// SwapBuffer is special, because it always waits for a VBlank before
		// beginning execution, plus its own 392 cycles. So move the timing
		// forward until we reach the next vblank point, so that we simulate
		// having waited until then.
		if cmd.code == GX_SWAP_BUFFERS {
			x, y := Emu.Sync.DotPos()
			dpd := Emu.Sync.DotPosDistance(0, 192)
			modGx.Infof("SwapBuffers: %d cmd, total:%d; now at (%d,%d) to vsync: %d",
				g.framestats.numcmd, g.cycles-g.framestats.start, x, y, dpd)
			g.cycles += dpd
			g.framestats.numcmd = 0
			g.framestats.start = g.cycles + cycles
		} else {
			g.framestats.numcmd++
		}
		g.cycles += cycles

		if g.fifo.LessThanHalfFull() {
			// modGxFifo.WithField("fifolen", len(g.fifo)).Info("trigger GXFIFO DMA")
			nds9.TriggerDmaEvent(DmaEventGxFifo)
		}
		g.busy = true
	}

	// We get here if there are no more full commands to execute in the FIFO,
	// even though there was some time available in this timeslice.
	if g.cycles < target {
		g.cycles = target
	}

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
