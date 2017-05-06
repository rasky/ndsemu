package logger

import (
	"bytes"
	"fmt"
	"ndsemu/emu/fixed"
	"os"
	"strings"
	"sync"
	"time"

	logrus "gopkg.in/Sirupsen/logrus.v0"
)

type EntryZ struct {
	lvl   logrus.Level
	msg   string
	mod   Module
	zfbuf [16]ZField
	zfidx int
	buf   bytes.Buffer
}

var ezpool = sync.Pool{
	New: func() interface{} {
		return new(EntryZ)
	},
}

func NewEntryZ() *EntryZ {
	return ezpool.Get().(*EntryZ)
}

func (z *EntryZ) Bool(key string, value bool) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeBool
		f.Key = key
		f.Boolean = value
		z.zfidx++
	}
	return z
}

func (z *EntryZ) String(key string, value string) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeString
		f.Key = key
		f.String = value
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Stringer(key string, value fmt.Stringer) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeStringer
		f.Key = key
		f.Interface = value
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Int(key string, value int) *EntryZ     { return z.Int64(key, int64(value)) }
func (z *EntryZ) Int8(key string, value int8) *EntryZ   { return z.Int64(key, int64(value)) }
func (z *EntryZ) Int16(key string, value int16) *EntryZ { return z.Int64(key, int64(value)) }
func (z *EntryZ) Int32(key string, value int32) *EntryZ { return z.Int64(key, int64(value)) }
func (z *EntryZ) Int64(key string, value int64) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeInt
		f.Key = key
		f.Integer = uint64(value)
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Uint(key string, value uint) *EntryZ     { return z.Uint64(key, uint64(value)) }
func (z *EntryZ) Uint8(key string, value uint8) *EntryZ   { return z.Uint64(key, uint64(value)) }
func (z *EntryZ) Uint16(key string, value uint16) *EntryZ { return z.Uint64(key, uint64(value)) }
func (z *EntryZ) Uint32(key string, value uint32) *EntryZ { return z.Uint64(key, uint64(value)) }
func (z *EntryZ) Uint64(key string, value uint64) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeUint
		f.Key = key
		f.Integer = value
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Hex8(key string, value uint8) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeHex8
		f.Key = key
		f.Integer = uint64(value)
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Hex16(key string, value uint16) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeHex16
		f.Key = key
		f.Integer = uint64(value)
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Hex32(key string, value uint32) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeHex32
		f.Key = key
		f.Integer = uint64(value)
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Hex64(key string, value uint64) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeHex64
		f.Key = key
		f.Integer = value
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Error(key string, err error) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeError
		f.Key = key
		f.Error = err
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Duration(key string, d time.Duration) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeDuration
		f.Key = key
		f.Duration = d
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Fixed12(key string, value fixed.F12) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeFixed12
		f.Key = key
		f.Fixed12 = value
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Vector12(key string, value [4]fixed.F12) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeVector12
		f.Key = key
		f.Vector12 = value
		z.zfidx++
	}
	return z
}

func (z *EntryZ) Blob(key string, value []byte) *EntryZ {
	if z != nil {
		f := &z.zfbuf[z.zfidx]
		f.Type = FieldTypeBlob
		f.Key = key
		f.Blob = value
		z.zfidx++
	}
	return z
}

var logfuncs = []func(*logrus.Entry, ...interface{}){
	(*logrus.Entry).Panic,
	(*logrus.Entry).Fatal,
	(*logrus.Entry).Error,
	(*logrus.Entry).Warn,
	(*logrus.Entry).Info,
	(*logrus.Entry).Debug,
}

const (
	nocolor = 0
	red     = 31
	green   = 32
	yellow  = 33
	blue    = 34
	gray    = 37
)

func (z *EntryZ) end() {
	for _, c := range contexts {
		c.AddLogContext(z)
	}

	modname := modNames[z.mod]
	frame := "xxx"
	levelText := strings.ToUpper(z.lvl.String())[0:4]

	// Extract special fields
	for i := 0; i < z.zfidx; i++ {
		switch z.zfbuf[i].Key {
		case "_frame":
			frame = z.zfbuf[i].Value()
		}
	}

	var levelColor int
	switch z.lvl {
	case logrus.DebugLevel:
		levelColor = gray
	case logrus.WarnLevel:
		levelColor = yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = red
	default:
		levelColor = blue
	}

	fmt.Fprintf(&z.buf, "\x1b[%dm%s\x1b[0m[%05s] [%s] %-38s ",
		levelColor, levelText, frame, modname, z.msg)
	for i := 0; i < z.zfidx; i++ {
		e := &z.zfbuf[i]
		if e.Key[0] == '_' {
			continue
		}
		fmt.Fprintf(&z.buf, " \x1b[%dm%s\x1b[0m=%s", levelColor, e.Key, e.Value())
	}
	z.buf.WriteByte('\n')
	output.Write(z.buf.Bytes())
	z.buf.Reset()

	if z.lvl == logrus.FatalLevel {
		os.Exit(1)
	} else if z.lvl == logrus.PanicLevel {
		panic("raising panic in logger")
	}

	// Recycle entry
	buf := z.buf
	*z = EntryZ{}
	z.buf = buf
	ezpool.Put(z)
}

func (z *EntryZ) End() {
	if z != nil {
		z.end()
	}
}
