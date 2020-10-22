package options

import (
	"reflect"
	"time"
)

type String struct {
	Valid bool
	Value string
}

func (o String) Unwrap(defaultValue ...string) string {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(string)
}

type Int struct {
	Valid bool
	Value int
}

func (o Int) Unwrap(defaultValue ...int) int {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int)
}

type Int8 struct {
	Valid bool
	Value int8
}

func (o Int8) Unwrap(defaultValue ...int8) int8 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int8)
}

type Int16 struct {
	Valid bool
	Value int16
}

func (o Int16) Unwrap(defaultValue ...int16) int16 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int16)
}

type Int32 struct {
	Valid bool
	Value int32
}

func (o Int32) Unwrap(defaultValue ...int32) int32 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int32)
}

type Int64 struct {
	Valid bool
	Value int64
}

func (o Int64) Unwrap(defaultValue ...int64) int64 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(int64)
}

type Uint8 struct {
	Valid bool
	Value uint8
}

func (o Uint8) Unwrap(defaultValue ...uint8) uint8 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint8)
}

type Uint16 struct {
	Valid bool
	Value uint16
}

func (o Uint16) Unwrap(defaultValue ...uint16) uint16 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint16)
}

type Uint32 struct {
	Valid bool
	Value uint32
}

func (o Uint32) Unwrap(defaultValue ...uint32) uint32 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint32)
}

type Uint64 struct {
	Valid bool
	Value uint64
}

func (o Uint64) Unwrap(defaultValue ...uint64) uint64 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(uint64)
}

type Float32 struct {
	Valid bool
	Value float32
}

func (o Float32) Unwrap(defaultValue ...float32) float32 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(float32)
}

type Float64 struct {
	Valid bool
	Value float64
}

func (o Float64) Unwrap(defaultValue ...float64) float64 {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(float64)
}

type Byte struct {
	Valid bool
	Value byte
}

func (o Byte) Unwrap(defaultValue ...byte) byte {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(byte)
}

type Rune struct {
	Valid bool
	Value rune
}

func (o Rune) Unwrap(defaultValue ...rune) rune {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(rune)
}

type Bool struct {
	Valid bool
	Value bool
}

func (o Bool) Unwrap(defaultValue ...bool) bool {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(bool)
}

type Time struct {
	Valid bool
	Value time.Time
}

func (o Time) Unwrap(defaultValue ...time.Time) time.Time {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(time.Time)
}
