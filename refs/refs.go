package refs

import (
	"reflect"
	"time"
)

func String(value string) *string {
	return &value
}

func UnwrapString(value *string, defaultValue ...string) string {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(string)
}

func Int(value int) *int {
	return &value
}

func UnwrapInt(value *int, defaultValue ...int) int {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int)
}

func Int8(value int8) *int8 {
	return &value
}

func UnwrapInt8(value *int8, defaultValue ...int8) int8 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int8)
}

func Int16(value int16) *int16 {
	return &value
}

func UnwrapInt16(value *int16, defaultValue ...int16) int16 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int16)
}

func Int32(value int32) *int32 {
	return &value
}

func UnwrapInt32(value *int32, defaultValue ...int32) int32 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int32)
}

func Int64(value int64) *int64 {
	return &value
}

func UnwrapInt64(value *int64, defaultValue ...int64) int64 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int64)
}

func Uint8(value uint8) *uint8 {
	return &value
}

func UnwrapUint8(value *uint8, defaultValue ...uint8) uint8 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint8)
}

func Uint16(value uint16) *uint16 {
	return &value
}

func UnwrapUint16(value *uint16, defaultValue ...uint16) uint16 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint16)
}

func Uint32(value uint32) *uint32 {
	return &value
}

func UnwrapUint32(value *uint32, defaultValue ...uint32) uint32 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint32)
}

func Uint64(value uint64) *uint64 {
	return &value
}

func UnwrapUint64(value *uint64, defaultValue ...uint64) uint64 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint64)
}

func Float32(value float32) *float32 {
	return &value
}

func UnwrapFloat32(value *float32, defaultValue ...float32) float32 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(float32)
}

func Float64(value float64) *float64 {
	return &value
}

func UnwrapFloat64(value *float64, defaultValue ...float64) float64 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(float64)
}

func Byte(value byte) *byte {
	return &value
}

func UnwrapByte(value *byte, defaultValue ...byte) byte {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(byte)
}

func Rune(value rune) *rune {
	return &value
}

func UnwrapRune(value *rune, defaultValue ...rune) rune {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(rune)
}

func Bool(value bool) *bool {
	return &value
}

func UnwrapBool(value *bool, defaultValue ...bool) bool {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(bool)
}

func Time(value time.Time) *time.Time {
	return &value
}

func UnwrapTime(value *time.Time, defaultValue ...time.Time) time.Time {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(time.Time)
}
