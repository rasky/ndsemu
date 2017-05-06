package logger

import (
	"fmt"
	"ndsemu/emu/fixed"
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

var logfuncs = []func(*logrus.Entry, ...interface{}){
	(*logrus.Entry).Panic,
	(*logrus.Entry).Fatal,
	(*logrus.Entry).Error,
	(*logrus.Entry).Warn,
	(*logrus.Entry).Info,
	(*logrus.Entry).Debug,
}

func (z *EntryZ) end() {
	fields := make(logrus.Fields, z.zfidx+1)
	fields["_mod"] = modNames[z.mod]
	for _, f := range z.zfbuf[:z.zfidx] {
		fields[f.Key] = f.Value()
	}
	for _, c := range contexts {
		c.AddLogContext(fields)
	}
	entry := logrus.StandardLogger().WithFields(fields)
	logfuncs[z.lvl](entry, z.msg)

	// Recycle entry
	*z = EntryZ{}
	ezpool.Put(z)
}

func (z *EntryZ) End() {
	if z != nil {
		z.end()
	}
}
