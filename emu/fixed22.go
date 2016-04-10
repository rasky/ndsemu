package emu

import "fmt"

type Fixed22 struct {
	V int32
}

func NewFixed22(val int32) Fixed22 {
	return Fixed22{val << 22}
}

func newFixed22FromInt64(val int64) Fixed22 {
	val32 := int32(val)
	if int64(val32) != val {
		fmt.Printf("%v %x\n", val, val)
		panic("fixed point overflow")
	}
	return Fixed22{val32}
}

func (f Fixed22) ToFloat64() float64 {
	return float64(f.V) / 4194304.0
}

func (f Fixed22) NearInt32() int32 {
	return (f.V + (1 << 21)) >> 22
}

func (f Fixed22) TruncInt32() int32 {
	return f.V >> 22
}

func (f Fixed22) Add(add int32) Fixed22 {
	return Fixed22{f.V + add<<22}
}

func (f Fixed22) AddFixed(v Fixed22) Fixed22 {
	return Fixed22{f.V + v.V}
}

func (f Fixed22) SubFixed(v Fixed22) Fixed22 {
	return Fixed22{f.V - v.V}
}

func (f Fixed22) Mul(mul int32) Fixed22 {
	return Fixed22{f.V * mul}
}

func (f Fixed22) Div(den int32) Fixed22 {
	return Fixed22{f.V / den}
}

func (f Fixed22) Neg() Fixed22 {
	return Fixed22{-f.V}
}

func (f Fixed22) DivFixed(den Fixed22) Fixed22 {
	return newFixed22FromInt64((int64(f.V) << 22) / int64(den.V))
}

func (f Fixed22) MulFixed(mul Fixed22) Fixed22 {
	return newFixed22FromInt64((int64(f.V) * int64(mul.V)) >> 22)
}

func (f Fixed22) Round() Fixed22 {
	return NewFixed22(f.NearInt32())
}

func (f Fixed22) Clamp(min, max Fixed22) Fixed22 {
	if f.V < min.V {
		f.V = min.V
	}
	if f.V > max.V {
		f.V = max.V
	}
	return f
}

// func (f Fixed22) mulFixedNearest(mul Fixed22) Fixed22 {
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
func (f Fixed22) Lerp(f2 Fixed22, ratio Fixed22) Fixed22 {
	return f2.SubFixed(f).MulFixed(ratio).AddFixed(f)
}

func (f Fixed22) String() string {
	return fmt.Sprintf("%.4f", f.ToFloat64())
}
