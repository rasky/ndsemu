package debugger

import (
	"bufio"
	"io"
	log "ndsemu/emu/logger"
	"regexp"
	"sync"
)

type logReader struct {
	NewLog chan string

	f      *bufio.Reader
	nlines int
	lines  []string
	lock   sync.Mutex
}

func newLogReader() *logReader {
	r, w := io.Pipe()
	log.SetOutput(w)

	lr := &logReader{
		f:      bufio.NewReader(r),
		nlines: 80,
		NewLog: make(chan string, 1),
	}
	go lr.loop()
	return lr
}

func (lr *logReader) SetNumLines(nl int) {
	lr.lock.Lock()
	lr.nlines = nl
	if len(lr.lines) > lr.nlines {
		lr.lines = lr.lines[len(lr.lines)-lr.nlines:]
	}
	lr.lock.Unlock()
}

func (lr *logReader) Lines() (res []string) {
	lr.lock.Lock()
	res = make([]string, len(lr.lines))
	copy(res[:], lr.lines[:])
	lr.lock.Unlock()
	return
}

var rxModName = regexp.MustCompile(`(\[[^0-9].*?\])`)
var rxAnsiRed = regexp.MustCompile(`\033\[31m(.*?)\033\[0m`)
var rxAnsiYellow = regexp.MustCompile(`\033\[33m(.*?)\033\[0m`)
var rxAnsiBlue = regexp.MustCompile(`\033\[34m(.*?)\033\[0m`)

var mkModName = []byte("[$1](fg-green)")
var mkRed = []byte("[$1](fg-red)")
var mkYellow = []byte("[$1](fg-yellow)")
var mkBlue = []byte("[$1](fg-blue)")

func (lr *logReader) markify(text []byte) []byte {
	text = rxModName.ReplaceAll(text, mkModName)
	text = rxAnsiBlue.ReplaceAll(text, mkBlue)
	text = rxAnsiYellow.ReplaceAll(text, mkYellow)
	text = rxAnsiRed.ReplaceAll(text, mkRed)
	return text
}

func (lr *logReader) loop() {
	for {
		line, err := lr.f.ReadSlice('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}

		line = lr.markify(line)
		sline := string(line[:len(line)-1])

		lr.lock.Lock()
		lr.lines = append(lr.lines, sline)
		if len(lr.lines) > lr.nlines {
			lr.lines = lr.lines[len(lr.lines)-lr.nlines:]
		}
		lr.lock.Unlock()

		select {
		case lr.NewLog <- sline:
		default:
		}
	}
}
