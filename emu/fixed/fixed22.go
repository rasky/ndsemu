package fixed

import "fmt"

type F22 struct {
	V int32
}

func NewF22(val int32) F22 {
	return F22{val << 22}
}

func newF22FromInt64(val int64) F22 {
	val32 := int32(val)
	if int64(val32) != val {
		fmt.Printf("%v %x\n", val, val)
		// panic("fixed point overflow")
	}
	return F22{val32}
}

func (f F22) ToFloat64() float64 {
	return float64(f.V) / 4194304.0
}

func (f F22) NearInt32() int32 {
	return (f.V + (1 << 21)) >> 22
}

func (f F22) TruncInt32() int32 {
	return f.V >> 22
}

func (f F22) Add(add int32) F22 {
	return F22{f.V + add<<22}
}

func (f F22) AddFixed(v F22) F22 {
	return F22{f.V + v.V}
}

func (f F22) SubFixed(v F22) F22 {
	return F22{f.V - v.V}
}

func (f F22) Mul(mul int32) F22 {
	return F22{f.V * mul}
}

func (f F22) Div(den int32) F22 {
	return F22{f.V / den}
}

func (f F22) Neg() F22 {
	return F22{-f.V}
}

func (f F22) DivFixed(den F22) F22 {
	return newF22FromInt64((int64(f.V) << 22) / int64(den.V))
}

func (f F22) MulFixed(mul F22) F22 {
	return newF22FromInt64((int64(f.V) * int64(mul.V)) >> 22)
}

func (f F22) Round() F22 {
	return NewF22(f.NearInt32())
}

func (f F22) Clamp(min, max F22) F22 {
	if f.V < min.V {
		f.V = min.V
	}
	if f.V > max.V {
		f.V = max.V
	}
	return f
}

// func (f F22) mulFixedNearest(mul F22) F22 {
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
func (f F22) Lerp(f2 F22, ratio F22) F22 {
	return f2.SubFixed(f).MulFixed(ratio).AddFixed(f)
}

func (f F22) String() string {
	return fmt.Sprintf("%.4f", f.ToFloat64())
}
