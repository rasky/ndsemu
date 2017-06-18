package arm

func (cpu *Cpu) JitCompileBlock(pc uint32, out []byte) (func(), int, int) {
	panic("JIT not implemented on 32-bit platforms")
}