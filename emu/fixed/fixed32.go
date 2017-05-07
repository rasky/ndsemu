package fixed

import "fmt"

type F32 struct {
	V int64
}

func NewF32(val int32) F32 {
	return F32{int64(val) << 32}
}

func (f F32) ToFloat64() float64 {
	return float64(f.V) / float64(1<<32)
}

func (f F32) NearInt32() int32 {
	return int32((f.V + (1 << 31)) >> 32)
}

func (f F32) TruncInt32() int32 {
	return int32(f.V >> 32)
}

func (f F32) Add(add int32) F32 {
	return F32{f.V + int64(add)<<32}
}

func (f F32) AddFixed(v F32) F32 {
	return F32{f.V + v.V}
}

func (f F32) SubFixed(v F32) F32 {
	return F32{f.V - v.V}
}

func (f F32) Mul(mul int32) F32 {
	return F32{f.V * int64(mul)}
}

func (f F32) Div(den int32) F32 {
	return F32{f.V / int64(den)}
}

func (f F32) Neg() F32 {
	return F32{-f.V}
}

// TODO: implement
// func (f F32) DivFixed(den F32) F32 {
// 	return F32{((int64(f.V) << 32) / int64(den.V))}
// }

// implemented in assembly
func mul128(x, y int64) (z1, z0 int64)

func (f F32) MulFixed(mul F32) F32 {
	hi, lo := mul128(f.V, mul.V)
	return F32{hi<<32 | int64(lo>>32)}
}

func (f F32) Round() F32 {
	return NewF32(f.NearInt32())
}

func (f F32) Clamp(min, max F32) F32 {
	if f.V < min.V {
		f.V = min.V
	}
	if f.V > max.V {
		f.V = max.V
	}
	return f
}

// func (f F32) mulFixedNearest(mul F32) F32 {
// 	res := int64(f.V) * int64(mul.V)
// 	if res >= 0 {
// 		res += 1 << 21
// 	} else {
// 		res -= 1 << 21
// 	}
// 	return newFromInt64(res >> 22)
// }

// Lerp computes a linear interpolation between f and f2.
// Returns f + (f2-f)*ratio
func (f F32) Lerp(f2 F32, ratio F32) F32 {
	return f2.SubFixed(f).MulFixed(ratio).AddFixed(f)
}

func (f F32) String() string {
	return fmt.Sprintf("%.4f", f.ToFloat64())
}
