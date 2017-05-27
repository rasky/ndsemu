package fixed

import "fmt"

type F12 struct {
	V int32
}

func NewF12(val int32) F12 {
	return F12{val << 12}
}

func newF12FromInt64(val int64) F12 {
	val32 := int32(val)
	if int64(val32) != val {
		fmt.Printf("%v %x\n", val, val)
		// panic("fixed point overflow")
	}
	return F12{val32}
}

func (f F12) ToFloat64() float64 {
	return float64(f.V) / 4096.0
}

func (f F12) NearInt32() int32 {
	return (f.V + (1 << 11)) >> 12
}

func (f F12) TruncInt32() int32 {
	return f.V >> 12
}

func (f F12) Add(add int32) F12 {
	return F12{f.V + add<<12}
}

func (f F12) AddFixed(v F12) F12 {
	return F12{f.V + v.V}
}

func (f F12) SubFixed(v F12) F12 {
	return F12{f.V - v.V}
}

func (f F12) Mul(mul int32) F12 {
	return F12{f.V * mul}
}

func (f F12) Div(den int32) F12 {
	return F12{f.V / den}
}

func (f F12) Neg() F12 {
	return F12{-f.V}
}

func (f F12) DivFixed(den F12) F12 {
	return newF12FromInt64((int64(f.V) << 12) / int64(den.V))
}

func (f F12) Inv() F12 {
	return newF12FromInt64((int64(1) << 24) / int64(f.V))
}

func (f F12) Inv22() F22 {
	return newF22FromInt64((int64(1) << (24 + (22 - 12))) / int64(f.V))
}

func (f F12) MulFixed(mul F12) F12 {
	return newF12FromInt64((int64(f.V) * int64(mul.V)) >> 12)
}

func (f F12) MulDivFixed(mul, div F12) F12 {
	return newF12FromInt64(int64(f.V) * int64(mul.V) / int64(div.V))
}

func (f F12) Round() F12 {
	return NewF12(f.NearInt32())
}

func (f F12) Clamp(min, max F12) F12 {
	if f.V < min.V {
		f.V = min.V
	}
	if f.V > max.V {
		f.V = max.V
	}
	return f
}

// func (f F12) mulFixedNearest(mul F12) F12 {
// 	res := int64(f.V) * int64(mul.V)
// 	if res >= 0 {
// 		res += 1 << 11
// 	} else {
// 		res -= 1 << 11
// 	}
// 	return newFromInt64(res >> 12)
// }

// Lerp computes a linear interpolation between f and f2.
// Returns f + (f2-f)*ratio
func (f F12) Lerp(f2 F12, ratio F12) F12 {
	return f2.SubFixed(f).MulFixed(ratio).AddFixed(f)
}

func (f F12) String() string {
	return fmt.Sprintf("%.4f", f.ToFloat64())
}

func (f F12) ToF32() (r F32) {
	r.V = int64(f.V) << (32 - 12)
	return
}

func (f F12) ToF22Sat() (r F22) {
	r.V = f.V << (22 - 12)
	if r.TruncInt32() != f.TruncInt32() {
		r.V = 0x7FFFFFFF
		if f.V < 0 {
			r.V = -r.V
		}
	}
	return
}
