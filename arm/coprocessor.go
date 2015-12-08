package arm

type Coprocessor interface {
	Read(op uint32, cn, cm, cp uint32) uint32
	Write(op uint32, cn, cm, cp uint32, value uint32)
	Exec(op uint32, cn, cm, cp, cd uint32)
}
