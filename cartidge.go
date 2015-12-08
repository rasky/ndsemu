package main

type CartHeader struct {
	Title      [12]byte
	Gamecode   [4]byte
	Maker      [2]byte
	Unit       byte
	EncSeed    byte
	Capacity   byte
	Reserved1  [9]byte
	RomVersion byte
	Autostart  byte

	Arm9Offset uint32
	Arm9Entry  uint32
	Arm9Ram    uint32
	Arm9Size   uint32

	Arm7Offset uint32
	Arm7Entry  uint32
	Arm7Ram    uint32
	Arm7Size   uint32
}
