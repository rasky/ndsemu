package emu

import "fmt"

type Fixed8 struct {
	v int64
}

func NewFixed8(val int64) Fixed8 {
	return Fixed8{val << 8}
}

func (f Fixed8) ToFloat64() float64 {
	return float64(f.v) / 256.0
}

func (f Fixed8) ToInt64() int64 {
	return (f.v + 0x80) >> 8
}

func (f Fixed8) Mul(mul int64) Fixed8 {
	return Fixed8{f.v * mul}
}

func (f Fixed8) Div(den int64) Fixed8 {
	return Fixed8{f.v / den}
}

func (f Fixed8) DivFixed(den Fixed8) Fixed8 {
	return Fixed8{(f.v << 8) / den.v}
}

func (f Fixed8) MulFixed(mul Fixed8) Fixed8 {
	return Fixed8{(f.v * mul.v) >> 8}
}

func (f Fixed8) String() string {
	return fmt.Sprintf("%.3f", f.ToFloat64())
}
