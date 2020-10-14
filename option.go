package duck

import (
	"reflect"
	"time"
)

type OptionString struct {
	Valid bool
	Value string
}

func (o OptionString) ValueOrDefault(defaultValue ...string) string {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(string)
}

type OptionInt struct {
	Valid bool
	Value int
}

func (o OptionInt) ValueOrDefault(defaultValue ...int) int {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int)
}

type OptionInt8 struct {
	Valid bool
	Value int8
}

func (o OptionInt8) ValueOrDefault(defaultValue ...int8) int8 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int8)
}

type OptionInt16 struct {
	Valid bool
	Value int16
}

func (o OptionInt16) ValueOrDefault(defaultValue ...int16) int16 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int16)
}

type OptionInt32 struct {
	Valid bool
	Value int32
}

func (o OptionInt32) ValueOrDefault(defaultValue ...int32) int32 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int32)
}

type OptionInt64 struct {
	Valid bool
	Value int64
}

func (o OptionInt64) ValueOrDefault(defaultValue ...int64) int64 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int64)
}

type OptionUint8 struct {
	Valid bool
	Value uint8
}

func (o OptionUint8) ValueOrDefault(defaultValue ...uint8) uint8 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint8)
}

type OptionUint16 struct {
	Valid bool
	Value uint16
}

func (o OptionUint16) ValueOrDefault(defaultValue ...uint16) uint16 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint16)
}

type OptionUint32 struct {
	Valid bool
	Value uint32
}

func (o OptionUint32) ValueOrDefault(defaultValue ...uint32) uint32 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint32)
}

type OptionUint64 struct {
	Valid bool
	Value uint64
}

func (o OptionUint64) ValueOrDefault(defaultValue ...uint64) uint64 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint64)
}

type OptionFloat32 struct {
	Valid bool
	Value float32
}

func (o OptionFloat32) ValueOrDefault(defaultValue ...float32) float32 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(float32)
}

type OptionFloat64 struct {
	Valid bool
	Value float64
}

func (o OptionFloat64) ValueOrDefault(defaultValue ...float64) float64 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(float64)
}

type OptionByte struct {
	Valid bool
	Value byte
}

func (o OptionByte) ValueOrDefault(defaultValue ...byte) byte {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(byte)
}

type OptionRune struct {
	Valid bool
	Value rune
}

func (o OptionRune) ValueOrDefault(defaultValue ...rune) rune {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(rune)
}

type OptionBool struct {
	Valid bool
	Value bool
}

func (o OptionBool) ValueOrDefault(defaultValue ...bool) bool {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(bool)
}

type OptionTime struct {
	Valid bool
	Value time.Time
}

func (o OptionTime) ValueOrDefault(defaultValue ...time.Time) time.Time {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(time.Time)
}
