package logger

import (
	"fmt"
	"ndsemu/emu"
	"strconv"
	"time"
)

type FieldType int

const (
	FieldTypeUnknown FieldType = iota
	FieldTypeBool
	FieldTypeString
	FieldTypeHex8
	FieldTypeHex16
	FieldTypeHex32
	FieldTypeHex64
	FieldTypeInt
	FieldTypeUint
	FieldTypeError
	FieldTypeDuration
	FieldTypeStringer
	FieldTypeFixed12
	FieldTypeVector12
)

type ZField struct {
	Type FieldType
	Key  string

	// Possible values. Only one of these is populated, depedning on Type
	String    string
	Integer   uint64
	Duration  time.Duration
	Error     error
	Fixed12   emu.Fixed12
	Vector12  [4]emu.Fixed12
	Interface interface{}
	Boolean   bool
}

func (f *ZField) Value() string {
	switch f.Type {
	case FieldTypeBool:
		if f.Boolean {
			return "true"
		}
		return "false"
	case FieldTypeString:
		return f.String
	case FieldTypeUint:
		return strconv.FormatUint(f.Integer, 10)
	case FieldTypeInt:
		return strconv.FormatInt(int64(f.Integer), 10)
	case FieldTypeHex8:
		return fmt.Sprintf("%02x", uint(f.Integer))
	case FieldTypeHex16:
		return fmt.Sprintf("%04x", uint(f.Integer))
	case FieldTypeHex32:
		return fmt.Sprintf("%08x", uint(f.Integer))
	case FieldTypeHex64:
		return fmt.Sprintf("%016x", uint(f.Integer))
	case FieldTypeError:
		return f.Error.Error()
	case FieldTypeDuration:
		return f.Duration.String()
	case FieldTypeStringer:
		return f.Interface.(fmt.Stringer).String()
	case FieldTypeFixed12:
		return f.Fixed12.String()
	case FieldTypeVector12:
		return fmt.Sprint(f.Vector12)
	}
	return ""
}
