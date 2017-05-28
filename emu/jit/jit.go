// JIT manager for helping implementing a working translator in an emulator.
//
// This packages implements a generic JIT high-level engine, which can be used
// by a CPU emulator to supplement the emulation with JIT compilation of hotter
// loops/functions. It is not meant to be used as the only way to execute code.
//
// JIT compilation happens in background, using a separate goroutine; this allows
// not to cause stalls in emulation as compilation is in progress.
//
// The main API entry point is Jit.Lookup(): this function can be called with a
// PC address to check if an already-compiled function is available; if not, some
// metrics are internally recorded and, when a function gets hot, it is automatically
// translated in background, so that a subsequen call to Jit.Lookup() eventually
// returns the compiled code.
//
// It is also necessary to call Jit.Invalidate() when the memory is written, so that
// the engine can discard already compiled blocks affected by the write. The function
// is very fast and is meant to be called on every memory write. There's also
// Jit.InvalidateRange() and Jit.InvalidateAll().
//
// Notice that Jit does not handle the actual translation, which is CPU-specific
// and target-specific (though we can assume amd64 for now); in fact, it requires
// an object implementing the Compiler interface to delegate actual compilation to
// an external package. With this decoupling, each CPU interpreter package can
// implement its own JIT (using go-jit as helper, for instance), leavin the high-level
// management to Jit.
package jit

import (
	log "ndsemu/emu/logger"
	"runtime"
	"sync/atomic"
	"unsafe"

	"github.com/edsrzf/mmap-go"
)

const (
	pageSize         = 1024 * 1024 // size of mmap page that contains JIT code
	jitCallThreshold = 255         // after how many calls the code block will be JIT'd
)

var modJit = log.NewModule("jit")

// Canary value used to mark blocks that have been scheduled for compilation
var pendingCanary = unsafe.Pointer(new(block))

const notCompiling = 0xFFFFFFFF

// Compiler is an interface to an object that is able to do an actual
// JIT compilation of a code-block to the target architecture.
// It is used by Jit to perform compilation.
type Compiler interface {
	// Compile the code block beginning with pc into out. Returns a callable
	// pointer to the compiled JIT code, the size of the input block in bytes,
	// and the size of JIT code in the output buffer.
	//
	// If the output buffer is too small for the block, it must return nil,-1,-1;
	// the JIT manager will allocate a larger buffer and call again the function.
	//
	// If the JIT is unable to compile the current code (for any reason, like
	// unsupported opcodes, etc.), it must return nil,0,0. Jit will avoid calling
	// JitCompileBlock() again for the same block.
	//
	// The definition of "block" is not specified; implementers are able to
	// decide what fits best the specific CPU architecture.
	//
	// This JIT engine currently does not handle jumping in the middle of a
	// compiled block, so implementers should probably consider inner loops with
	// conditional jumps (eg: IFs, FORs, WHILEs, etc.) as part of the block, while
	// a function call should probably terminate the block (as the JIT would be
	// unable to use the following part after the function returns).
	JitCompileBlock(pc uint32, out []byte) (jit func(), blockSize int, outSize int)
}

// Config is the configuration for Jit.
type Config struct {
	// Required alignment of the program counter (eg: 4 bytes = shift by 2).
	// This is used to save memory and speed up operations.
	// Notice that the JIT will panic if an unaligned address is ever addressed.
	PcAlignmentShift uint

	// Maximum size of a block that the compiler will generate. This is
	// currently required to allow the JIT to perform fast invalidation;
	// basically, if the background process is compiling a function starting
	// at X, all writes between X and X+MaxBlockSize will invalidate the
	// function being compiled.
	MaxBlockSize int
}

type block struct {
	// Pointer to recompiled function emulating this block
	jitcode func()

	// Beginning of this jit block
	pcstart uint32
	size    uint32
}

// Jit is a generic JIT manager, that handles common tasks that are necessary
// for all JITs, with the exclusion of the actual code generation. Code
// generation is handled by an instance implementing the Compiler interface,
// that Jit uses.
type Jit struct {
	comp Compiler
	cfg  *Config

	taskCh         chan uint32 // Channel were functions to be compiled get queued
	bkgCompilingPc uint64      // If != notCompiling, this is the PC of the block being compiled

	pageLastIdx  int                     // free index into the last page
	pages        [][]byte                // memory mapped pages
	blocks       [65536][]unsafe.Pointer // *block, needs unsafe.Pointer for atomic
	blockMetrics [65536][]byte

	HACK_OtherJit *Jit
}

// NewJit creates a new Jit object. It requires a Compiler which will be
// called to do the actual translation, and a Config instance that configures
// some advanced parameters of Jit.
func NewJit(comp Compiler, cfg *Config) *Jit {
	j := &Jit{
		cfg:    cfg,
		comp:   comp,
		taskCh: make(chan uint32, 256),
	}
	go j.bkgProc()
	return j
}

// Lookup checks if there is a JIT function available for this address, and
// returns it if so.
// If it's not available, it will keep some metrics of the most called
// addresses, and eventually schedule a background compilation that will
// make the JIT code available.
func (j *Jit) Lookup(pc uint32) func() {
	align := j.cfg.PcAlignmentShift
	if bg := j.blocks[pc>>16]; bg != nil {
		b := (*block)(atomic.LoadPointer(&bg[(pc&0xFFFF)>>align]))
		if b != nil {
			// b.jitcode could be null if the block was compiled but the
			// compilation failed for any reason. In this case, we correctly
			// return nil, but don't update metrics or trigger a new compilation
			// as it would fail as well.
			return b.jitcode
		}
	}

	// No code found. Bump metrics for this target
	j.updateMetrics(pc)
	return nil
}

func (j *Jit) updateMetrics(pc uint32) {
	align := j.cfg.PcAlignmentShift
	mg := j.blockMetrics[pc>>16]
	if mg == nil {
		mg = make([]byte, (1<<16)>>align)
		j.blockMetrics[pc>>16] = mg
	}
	idx := (pc & 0xFFFF) >> align
	mg[idx]++
	if mg[idx] == jitCallThreshold {
		// This call target is used a lot. We want to trigger background compilation.
		// Before spawning the background process, allocate the block where the
		// code will be stored. This allows Invalidate() to be called concurrently
		// with background compilation, since the result will just be discarded.
		bg := j.blocks[pc>>16]
		if bg == nil {
			bg = make([]unsafe.Pointer, (1<<16)>>align)
			j.blocks[pc>>16] = bg
		}
		bptr := &bg[(pc&0xFFFF)>>align]

		modJit.InfoZ("requesting compilation").Hex32("pc", pc).End()

		if !atomic.CompareAndSwapPointer(bptr, nil, pendingCanary) {
			modJit.WarnZ("requested JIT for function already compiled").Hex32("pc", pc).End()
			return
		}

		select {
		case j.taskCh <- pc:
		default:
			// The channel is full; don't block but try again next time
			// we call this target
			mg[idx]--
			modJit.InfoZ("compilation queue full").Hex32("pc", pc).End()
		}
	}
}

// Invalidate invalidates a specific program counter (eg: after the CPU wrote
// to this address).
// This function is very fast and is meant to be called for every memory write.
func (j *Jit) invalidate(pc uint32) {
	// Fast-path: in the normal case the CPU is not writing onto the code
	if bg := j.blocks[pc>>16]; bg != nil {
		align := j.cfg.PcAlignmentShift
		atomic.StorePointer(&bg[(pc&0xFFFF)>>align], nil)
		// If we're recompiling a function right now, check if we must invalidate it.
		// There's no race condition here: even if bkgCompilingPc gets set just after
		// we load it, it means that the JIT compiler hasn't accessed the memory yet
		// so we don't need to invalidate the block pointer.
		if pcfunc := atomic.LoadUint64(&j.bkgCompilingPc); pcfunc != notCompiling {
			if uint64(pc)-pcfunc < uint64(j.cfg.MaxBlockSize) {
				atomic.StorePointer(&j.blocks[pc>>16][(pc&0xFFFF)>>align], nil)
			}
		}
	}
}

func (j *Jit) Invalidate(pc uint32) {
	j.invalidate(pc)
	if j.HACK_OtherJit != nil {
		j.HACK_OtherJit.invalidate(pc)
	}
}

// InvalidateRange invalidates a range of addresses, starting from pc, for a total
// of size bytes. It discards any JIT code related to this area
func (j *Jit) InvalidateRange(pc uint32, size int) {
	align := j.cfg.PcAlignmentShift

	// Extend the range backwards to cover the maximum block size. This
	// helps invalidating a range which falls within a block.
	if pc < uint32(j.cfg.MaxBlockSize) {
		pc = 0
	} else {
		pc -= uint32(j.cfg.MaxBlockSize)
	}
	size += j.cfg.MaxBlockSize

	// If the whole range is within a block group that's empty,
	// we can just exit right away.
	if j.blocks[pc>>16] == nil && pc>>16 == (pc+uint32(size)-1)>>16 {
		return
	}

	j.lockBkgProc(0)
	for size > 0 {
		if bg := j.blocks[pc>>16]; bg != nil {
			atomic.StorePointer(&bg[(pc&0xFFFF)>>align], nil)
		}
		pc += 1 << align
		size -= 1 << align
	}
	j.unlockBkgProc()
}

// Invalidate the whole memory; this discards any compiled JIT code
func (j *Jit) InvalidateAll() {
	j.lockBkgProc(0)
	// Clean all blocks
	for i := range j.blocks {
		j.blocks[i] = nil
	}
	// Also release all pages
	for i := range j.pages {
		m := mmap.MMap(j.pages[i])
		m.Unmap()
	}
	j.pages = nil
	j.unlockBkgProc()
}

func (j *Jit) newPage() {
	b, err := mmap.MapRegion(nil, 1024*1024, mmap.EXEC|mmap.RDWR, mmap.ANON, int64(0))
	if err != nil {
		panic(err)
	}
	j.pages = append(j.pages, b)
	j.pageLastIdx = 0
}

func (j *Jit) getFreePage() []byte {
	if len(j.pages) == 0 {
		j.newPage()
	}
	return j.pages[len(j.pages)-1][j.pageLastIdx:]
}

// Background process that handles compilation
func (j *Jit) bkgProc() {
	for pc := range j.taskCh {
		j.lockBkgProc(pc)

		// Compile the new call target
		code, insize, outsize := j.comp.JitCompileBlock(pc, j.getFreePage())
		if outsize < 0 {
			// If the free part of the page wasn't big enough,
			// try again after allocating a larger page.
			j.newPage()
			code, insize, outsize = j.comp.JitCompileBlock(pc, j.getFreePage())
		}
		j.pageLastIdx += outsize

		// Sanity check: the compiler should not create a block bigger than
		// MaxBlockSize (otherwise InvalidateRange might fail)
		if insize > j.cfg.MaxBlockSize {
			modJit.FatalZ("compiled block is too big").Hex32("pc", pc).Int("sz", insize).Int("max", j.cfg.MaxBlockSize).End()
		}

		modJit.InfoZ("compiled block").Hex32("pc", pc).Int("insn", insize/4).End()

		// Prepare the new block. Notice that code could be nil (if
		// compilation failed), but we don't care and still save the block
		// to avoid trying compiling this block again.
		b := new(block)
		b.jitcode = code
		b.pcstart = pc
		b.size = uint32(insize)

		// Store it atomically. We should find the pending canary in the slot;
		// if we don't, it means that the block was invalidated while we were
		// recompiling it, so we can just ignore the error.
		bg := j.blocks[pc>>16]
		if bg != nil {
			bptr := &bg[(pc&0xFFFF)>>j.cfg.PcAlignmentShift]
			atomic.CompareAndSwapPointer(bptr, pendingCanary, unsafe.Pointer(b))
		}

		j.unlockBkgProc()
	}
}

func (j *Jit) lockBkgProc(pc uint32) {
	for atomic.CompareAndSwapUint64(&j.bkgCompilingPc, notCompiling, uint64(pc)) {
		// TODO: check if time.Sleep(1*time.Microsecond is better)
		runtime.Gosched()
	}
}

func (j *Jit) unlockBkgProc() {
	atomic.StoreUint64(&j.bkgCompilingPc, notCompiling)
}
