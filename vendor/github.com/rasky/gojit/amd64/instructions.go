package amd64

type ImmRm struct {
	op  []byte
	sub byte
}

type Instruction struct {
	Mnemonic string
	imm_r    []byte
	imm_rm   ImmRm
	imm8_rm  ImmRm
	r_rm     []byte
	rm_r     []byte
	bits     byte
}

type j []byte

func as32Bits(i64 *Instruction) *Instruction {
	i32 := *i64
	i32.Mnemonic += "l"
	i32.bits = 32
	return &i32
}

var (
	InstAdd  = &Instruction{"add", nil, ImmRm{j{0x81}, 0}, ImmRm{j{0x83}, 0}, j{0x01}, j{0x03}, 64}
	InstAddb = &Instruction{"addb", nil, ImmRm{j{0x80}, 0}, ImmRm{}, j{0x00}, j{0x02}, 8}

	InstAdc  = &Instruction{"adc", nil, ImmRm{j{0x81}, 2}, ImmRm{j{0x83}, 2}, j{0x11}, j{0x13}, 64}
	InstAdcb = &Instruction{"adcb", nil, ImmRm{j{0x80}, 2}, ImmRm{}, j{0x10}, j{0x12}, 8}

	InstAnd  = &Instruction{"and", nil, ImmRm{j{0x81}, 4}, ImmRm{j{0x83}, 4}, j{0x21}, j{0x23}, 64}
	InstAndb = &Instruction{"andb", nil, ImmRm{j{0x80}, 4}, ImmRm{}, j{0x20}, j{0x22}, 8}

	InstCmp  = &Instruction{"cmp", nil, ImmRm{j{0x81}, 7}, ImmRm{j{0x83}, 7}, j{0x39}, j{0x3B}, 64}
	InstCmpb = &Instruction{"cmpb", nil, ImmRm{j{0x80}, 7}, ImmRm{}, j{0x38}, j{0x3A}, 8}

	InstOr  = &Instruction{"or", nil, ImmRm{j{0x81}, 1}, ImmRm{j{0x83}, 1}, j{0x09}, j{0x0B}, 64}
	InstOrb = &Instruction{"orb", nil, ImmRm{j{0x80}, 1}, ImmRm{}, j{0x08}, j{0x0A}, 8}

	InstSbb  = &Instruction{"sbb", nil, ImmRm{j{0x81}, 3}, ImmRm{j{0x83}, 3}, j{0x19}, j{0x1B}, 64}
	InstSbbb = &Instruction{"sbbb", nil, ImmRm{j{0x80}, 3}, ImmRm{}, j{0x18}, j{0x1A}, 8}

	InstSub  = &Instruction{"sub", nil, ImmRm{j{0x81}, 5}, ImmRm{j{0x83}, 5}, j{0x29}, j{0x2B}, 64}
	InstSubb = &Instruction{"subb", nil, ImmRm{j{0x80}, 5}, ImmRm{}, j{0x28}, j{0x2A}, 8}

	InstTest  = &Instruction{"test", nil, ImmRm{j{0xF7}, 0}, ImmRm{}, j{0x85}, nil, 64}
	InstTestb = &Instruction{"testb", nil, ImmRm{j{0xF6}, 0}, ImmRm{}, j{0x84}, nil, 8}

	InstXor  = &Instruction{"xor", nil, ImmRm{j{0x81}, 6}, ImmRm{}, j{0x31}, j{0x33}, 64}
	InstXorb = &Instruction{"xorb", nil, ImmRm{j{0x80}, 6}, ImmRm{}, j{0x30}, j{0x32}, 8}

	InstRol  = &Instruction{"rol", nil, ImmRm{}, ImmRm{j{0xC1}, 0}, nil, nil, 64}
	InstRolb = &Instruction{"rolb", nil, ImmRm{}, ImmRm{j{0xC0}, 0}, nil, nil, 8}
	InstRor  = &Instruction{"rol", nil, ImmRm{}, ImmRm{j{0xC1}, 1}, nil, nil, 64}
	InstRorb = &Instruction{"rolb", nil, ImmRm{}, ImmRm{j{0xC0}, 1}, nil, nil, 8}
	InstRcl  = &Instruction{"rcl", nil, ImmRm{}, ImmRm{j{0xC1}, 2}, nil, nil, 64}
	InstRclb = &Instruction{"rclb", nil, ImmRm{}, ImmRm{j{0xC0}, 2}, nil, nil, 8}
	InstRcr  = &Instruction{"rcl", nil, ImmRm{}, ImmRm{j{0xC1}, 3}, nil, nil, 64}
	InstRcrb = &Instruction{"rclb", nil, ImmRm{}, ImmRm{j{0xC0}, 3}, nil, nil, 8}

	InstShl  = &Instruction{"shl", nil, ImmRm{}, ImmRm{j{0xC1}, 4}, nil, nil, 64}
	InstShlb = &Instruction{"shlb", nil, ImmRm{}, ImmRm{j{0xC0}, 4}, nil, nil, 8}
	InstShr  = &Instruction{"shr", nil, ImmRm{}, ImmRm{j{0xC1}, 5}, nil, nil, 64}
	InstShrb = &Instruction{"shrb", nil, ImmRm{}, ImmRm{j{0xC0}, 5}, nil, nil, 8}
	InstSar  = &Instruction{"sar", nil, ImmRm{}, ImmRm{j{0xC1}, 7}, nil, nil, 64}
	InstSarb = &Instruction{"sarb", nil, ImmRm{}, ImmRm{j{0xC0}, 7}, nil, nil, 8}

	InstBt  = &Instruction{"bt", nil, ImmRm{j{0x0F, 0xBA}, 4}, ImmRm{}, j{0x0F, 0xA3}, nil, 8}
	InstBtc = &Instruction{"btc", nil, ImmRm{j{0x0F, 0xBA}, 7}, ImmRm{}, j{0x0F, 0xBB}, nil, 8}
	InstBtr = &Instruction{"btr", nil, ImmRm{j{0x0F, 0xBA}, 6}, ImmRm{}, j{0x0F, 0xB3}, nil, 8}
	InstBts = &Instruction{"bts", nil, ImmRm{j{0x0F, 0xBA}, 5}, ImmRm{}, j{0x0F, 0xAB}, nil, 8}

	InstLea = &Instruction{"lea", nil, ImmRm{nil, 0}, ImmRm{}, nil, j{0x8D}, 64}

	InstMov  = &Instruction{"mov", j{0xB8}, ImmRm{j{0xc7}, 0}, ImmRm{}, j{0x89}, j{0x8b}, 64}
	InstMovl = as32Bits(InstMov)
	InstMovb = &Instruction{"movb", j{0xB0}, ImmRm{j{0xc6}, 0}, ImmRm{}, j{0x88}, j{0x8a}, 8}

	InstBsr = &Instruction{"bsr", nil, ImmRm{}, ImmRm{}, nil, j{0x0F, 0xBD}, 64}
	InstBsf = &Instruction{"bsf", nil, ImmRm{}, ImmRm{}, nil, j{0x0F, 0xBC}, 64}
)
