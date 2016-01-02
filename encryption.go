package main

import "encoding/binary"

// Key1 is a slightly modified Blowfish implementation. The main differences are:
//   1) data is accessed as little-endian rather than big-endian
//   2) the standard tables are not used; a custom set of tables (stored in the
//      BIOS) are used instead.
type Key1 struct {
	p              [18]uint32
	s0, s1, s2, s3 [256]uint32
}

func NewKey1(biosTables []byte, gameCode []byte) *Key1 {
	var c Key1

	// Copy tables from bios into our class
	for i := 0; i < 18; i++ {
		c.p[i] = binary.LittleEndian.Uint32(biosTables[:4])
		biosTables = biosTables[4:]
	}
	for i := 0; i < 256; i++ {
		c.s0[i] = binary.LittleEndian.Uint32(biosTables[:4])
		biosTables = biosTables[4:]
	}
	for i := 0; i < 256; i++ {
		c.s1[i] = binary.LittleEndian.Uint32(biosTables[:4])
		biosTables = biosTables[4:]
	}
	for i := 0; i < 256; i++ {
		c.s2[i] = binary.LittleEndian.Uint32(biosTables[:4])
		biosTables = biosTables[4:]
	}
	for i := 0; i < 256; i++ {
		c.s3[i] = binary.LittleEndian.Uint32(biosTables[:4])
		biosTables = biosTables[4:]
	}

	// Apply a custom key expansion algorithm, using the 4-byte gamecode (from
	// the cartridge header) as key. The algorithm is built upon the standard
	// Blowfish key expansion, but there is some additional stretching going
	// on.
	idcode := binary.LittleEndian.Uint32(gameCode)

	var keycode [12]byte
	binary.LittleEndian.PutUint32(keycode[0:4], idcode)
	binary.LittleEndian.PutUint32(keycode[4:8], idcode/2)
	binary.LittleEndian.PutUint32(keycode[8:12], idcode*2)

	c.EncryptLE(keycode[4:12], keycode[4:12])
	c.EncryptLE(keycode[0:8], keycode[0:8])
	c.expandKey(keycode[0:8])

	c.EncryptLE(keycode[4:12], keycode[4:12])
	c.EncryptLE(keycode[0:8], keycode[0:8])
	c.expandKey(keycode[0:8])

	return &c
}

func (c *Key1) EncryptLE(dst, src []byte) {
	r := binary.LittleEndian.Uint32(src[:4])
	l := binary.LittleEndian.Uint32(src[4:])
	l, r = c.encryptBlock(l, r)
	binary.LittleEndian.PutUint32(dst[:4], r)
	binary.LittleEndian.PutUint32(dst[4:], l)
}

func (c *Key1) DecryptLE(dst, src []byte) {
	r := binary.LittleEndian.Uint32(src[:4])
	l := binary.LittleEndian.Uint32(src[4:])
	l, r = c.decryptBlock(l, r)
	binary.LittleEndian.PutUint32(dst[:4], r)
	binary.LittleEndian.PutUint32(dst[4:], l)
}

func (c *Key1) EncryptBE(dst, src []byte) {
	l := binary.BigEndian.Uint32(src[:4])
	r := binary.BigEndian.Uint32(src[4:])
	l, r = c.encryptBlock(l, r)
	binary.BigEndian.PutUint32(dst[:4], l)
	binary.BigEndian.PutUint32(dst[4:], r)
}

func (c *Key1) DecryptBE(dst, src []byte) {
	l := binary.BigEndian.Uint32(src[:4])
	r := binary.BigEndian.Uint32(src[4:])
	l, r = c.decryptBlock(l, r)
	binary.BigEndian.PutUint32(dst[:4], l)
	binary.BigEndian.PutUint32(dst[4:], r)
}

// Standard blowfish encryption round
func (c *Key1) encryptBlock(l, r uint32) (uint32, uint32) {
	xl, xr := l, r
	xl ^= c.p[0]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[1]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[2]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[3]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[4]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[5]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[6]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[7]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[8]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[9]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[10]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[11]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[12]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[13]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[14]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[15]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[16]
	xr ^= c.p[17]
	return xr, xl
}

// Standard blowfish decryption round
func (c *Key1) decryptBlock(l, r uint32) (uint32, uint32) {
	xl, xr := l, r
	xl ^= c.p[17]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[16]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[15]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[14]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[13]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[12]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[11]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[10]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[9]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[8]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[7]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[6]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[5]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[4]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[3]
	xr ^= ((c.s0[byte(xl>>24)] + c.s1[byte(xl>>16)]) ^ c.s2[byte(xl>>8)]) + c.s3[byte(xl)] ^ c.p[2]
	xl ^= ((c.s0[byte(xr>>24)] + c.s1[byte(xr>>16)]) ^ c.s2[byte(xr>>8)]) + c.s3[byte(xr)] ^ c.p[1]
	xr ^= c.p[0]
	return xr, xl
}

// Standard blowfish key expansion
func (c *Key1) expandKey(key []byte) {
	j := 0
	for i := 0; i < 18; i++ {
		var d uint32
		for k := 0; k < 4; k++ {
			d = d<<8 | uint32(key[j])
			j++
			if j >= len(key) {
				j = 0
			}
		}
		c.p[i] ^= d
	}

	var l, r uint32
	for i := 0; i < 18; i += 2 {
		l, r = c.encryptBlock(l, r)
		c.p[i], c.p[i+1] = l, r
	}

	for i := 0; i < 256; i += 2 {
		l, r = c.encryptBlock(l, r)
		c.s0[i], c.s0[i+1] = l, r
	}
	for i := 0; i < 256; i += 2 {
		l, r = c.encryptBlock(l, r)
		c.s1[i], c.s1[i+1] = l, r
	}
	for i := 0; i < 256; i += 2 {
		l, r = c.encryptBlock(l, r)
		c.s2[i], c.s2[i+1] = l, r
	}
	for i := 0; i < 256; i += 2 {
		l, r = c.encryptBlock(l, r)
		c.s3[i], c.s3[i+1] = l, r
	}
}

// Key2 is a simple stream cipher that uses 2 39-bit LSFRs to generate
// the PRNG to encrypt the ciphertext. The LSFRs have the following polynomials:
// 	 L1 = x^5+x^17+x^18+x^31
// 	 L2 = x^5+x^18+x^23+x^31
type Key2 struct {
	x, y uint64
}

func br39(val uint64) uint64 {
	ret := uint64(0)
	for i := 0; i < 39; i++ {
		ret |= (val & 1) << uint(38-i)
		val >>= 1
	}
	return ret
}

// Initialize with default seed
func NewKey2() Key2 {
	return NewKey2WithSeed(0x58C56DE0E8, 0x5C879B9B05)
}

func NewKey2WithSeed(x, y uint64) Key2 {
	return Key2{
		x: br39(x),
		y: br39(y),
	}
}

func (k *Key2) Encrypt(out, in []byte) {
	for idx, v := range in {
		x := uint8((k.x >> 5) ^ (k.x >> 17) ^ (k.x >> 18) ^ (k.x >> 31))
		y := uint8((k.y >> 5) ^ (k.y >> 23) ^ (k.y >> 18) ^ (k.y >> 31))
		out[idx] = v ^ x ^ y
		k.x = k.x<<8 | uint64(x)
		k.y = k.y<<8 | uint64(y)
	}
}
