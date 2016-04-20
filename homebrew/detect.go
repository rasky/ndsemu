package homebrew

import "os"

// Check if the specified ROM file is a Hombrew NDS ROM
func Detect(fn string) (bool, error) {
	f, err := os.Open(fn)
	if err != nil {
		return false, err
	}
	defer f.Close()

	var gameid [4]byte
	f.ReadAt(gameid[:], 0xAC)
	return gameid == [4]byte{'P', 'A', 'S', 'S'}, nil
}
