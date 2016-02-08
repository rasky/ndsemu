package emu

import "fmt"

type Fixed12 struct {
	V int32
}

func NewFixed12(val int32) Fixed12 {
	return Fixed12{val << 12}
}

func (f Fixed12) ToFloat64() float64 {
	return float64(f.V) / 4096.0
}

func (f Fixed12) ToInt32() int32 {
	return (f.V + 0x80) >> 8
}

func (f Fixed12) AddFixed(v Fixed12) Fixed12 {
	return Fixed12{f.V + v.V}
}

func (f Fixed12) Mul(mul int32) Fixed12 {
	return Fixed12{f.V * mul}
}

func (f Fixed12) Div(den int32) Fixed12 {
	return Fixed12{f.V / den}
}

func (f Fixed12) DivFixed(den Fixed12) Fixed12 {
	return Fixed12{(f.V << 12) / den.V}
}

func (f Fixed12) MulFixed(mul Fixed12) Fixed12 {
	return Fixed12{(f.V * mul.V) >> 12}
}

func (f Fixed12) String() string {
	return fmt.Sprintf("%.3f", f.ToFloat64())
}
