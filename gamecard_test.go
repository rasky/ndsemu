package main

import (
	"encoding/binary"
	"os"
	"reflect"
	"testing"
)

func TestEncryption(t *testing.T) {
	f, err := os.Open("bios/biosnds7.rom")
	if err != nil {
		t.Fatal(err)
	}

	data := make([]byte, 18*4+256*4*4)
	f.ReadAt(data, 0x30)
	f.Close()

	c := NewKey1(data, []byte("AZEP"))

	var test [8]byte
	binary.BigEndian.PutUint64(test[:], 0x2229b690c67c17ff)
	c.DecryptBE(test[:], test[:])

	var exp [8]byte
	binary.BigEndian.PutUint64(exp[:], 0x46e79a8337bc195d)

	if !reflect.DeepEqual(test, exp) {
		t.Errorf("decryption error, got:%x, want:%x", test, exp)
	}

	c.EncryptBE(exp[:], exp[:])

	var exp2 [8]byte
	binary.BigEndian.PutUint64(exp2[:], 0x2229b690c67c17ff)

	if !reflect.DeepEqual(exp, exp2) {
		t.Errorf("decryption error, got:%x, want:%x", exp, exp2)
	}
}
