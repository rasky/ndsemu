package fixed

import "fmt"

type F8 struct {
	v int64
}

func NewF8(val int64) F8 {
	return F8{val << 8}
}

func (f F8) ToFloat64() float64 {
	return float64(f.v) / 256.0
}

func (f F8) ToInt64() int64 {
	return (f.v + 0x80) >> 8
}

func (f F8) Mul(mul int64) F8 {
	return F8{f.v * mul}
}

func (f F8) Div(den int64) F8 {
	return F8{f.v / den}
}

func (f F8) DivFixed(den F8) F8 {
	return F8{(f.v << 8) / den.v}
}

func (f F8) MulFixed(mul F8) F8 {
	return F8{(f.v * mul.v) >> 8}
}

func (f F8) String() string {
	return fmt.Sprintf("%.3f", f.ToFloat64())
}
