package emu

import "fmt"

type Fixed22 struct {
	V int32
}

func NewFixed22(val int32) Fixed22 {
	return Fixed22{val << 12}
}

func (f Fixed22) ToFloat64() float64 {
	return float64(f.V) / 256.0
}

func (f Fixed22) ToInt32() int32 {
	return (f.V + 0x80) >> 8
}

func (f Fixed22) Mul(mul int32) Fixed22 {
	return Fixed22{f.V * mul}
}

func (f Fixed22) Div(den int32) Fixed22 {
	return Fixed22{f.V / den}
}

func (f Fixed22) DivFixed(den Fixed22) Fixed22 {
	return Fixed22{(f.V << 8) / den.V}
}

func (f Fixed22) MulFixed(mul Fixed22) Fixed22 {
	return Fixed22{(f.V * mul.V) >> 8}
}

func (f Fixed22) String() string {
	return fmt.Sprintf("%.3f", f.ToFloat64())
}
