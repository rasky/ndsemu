package emu

import "fmt"

type Fixed12 struct {
	V int32
}

func NewFixed12(val int32) Fixed12 {
	return Fixed12{val << 12}
}

func newFromInt64(val int64) Fixed12 {
	val32 := int32(val)
	if int64(val32) != val {
		fmt.Printf("%v %x\n", val, val)
		panic("fixed point overflow")
	}
	return Fixed12{val32}
}

func (f Fixed12) ToFloat64() float64 {
	return float64(f.V) / 4096.0
}

func (f Fixed12) NearInt32() int32 {
	return (f.V + (1 << 11)) >> 12
}

func (f Fixed12) TruncInt32() int32 {
	return f.V >> 12
}

func (f Fixed12) Add(add int32) Fixed12 {
	return Fixed12{f.V + add<<12}
}

func (f Fixed12) AddFixed(v Fixed12) Fixed12 {
	return Fixed12{f.V + v.V}
}

func (f Fixed12) SubFixed(v Fixed12) Fixed12 {
	return Fixed12{f.V - v.V}
}

func (f Fixed12) Mul(mul int32) Fixed12 {
	return Fixed12{f.V * mul}
}

func (f Fixed12) Div(den int32) Fixed12 {
	return Fixed12{f.V / den}
}

func (f Fixed12) DivFixed(den Fixed12) Fixed12 {
	return newFromInt64((int64(f.V) << 12) / int64(den.V))
}

func (f Fixed12) MulFixed(mul Fixed12) Fixed12 {
	return newFromInt64((int64(f.V) * int64(mul.V)) >> 12)
}

func (f Fixed12) Round() Fixed12 {
	return NewFixed12(f.NearInt32())
}

func (f Fixed12) String() string {
	return fmt.Sprintf("%.4f", f.ToFloat64())
}
