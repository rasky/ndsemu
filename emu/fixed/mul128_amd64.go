package fixed

// implemented in assembly
func mul128(x, y int64) (hi int64, lo uint64)
func div128(hinum, lonum, den int64) (quo, rem int64)
