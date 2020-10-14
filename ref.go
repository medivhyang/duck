package duck

import (
	"reflect"
	"time"
)

func RefString(value string) *string {
	return &value
}

func DerefString(value *string, defaultValue ...string) string {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(string)
}

func RefInt(value int) *int {
	return &value
}

func DerefInt(value *int, defaultValue ...int) int {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int)
}

func RefInt8(value int8) *int8 {
	return &value
}

func DerefInt8(value *int8, defaultValue ...int8) int8 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int8)
}

func RefInt16(value int16) *int16 {
	return &value
}

func DerefInt16(value *int16, defaultValue ...int16) int16 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int16)
}

func RefInt32(value int32) *int32 {
	return &value
}

func DerefInt32(value *int32, defaultValue ...int32) int32 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int32)
}

func RefInt64(value int64) *int64 {
	return &value
}

func DerefInt64(value *int64, defaultValue ...int64) int64 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(int64)
}

func RefUint8(value uint8) *uint8 {
	return &value
}

func DerefUint8(value *uint8, defaultValue ...uint8) uint8 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint8)
}

func RefUint16(value uint16) *uint16 {
	return &value
}

func DerefUint16(value *uint16, defaultValue ...uint16) uint16 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint16)
}

func RefUint32(value uint32) *uint32 {
	return &value
}

func DerefUint32(value *uint32, defaultValue ...uint32) uint32 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint32)
}

func RefUint64(value uint64) *uint64 {
	return &value
}

func DerefUint64(value *uint64, defaultValue ...uint64) uint64 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(uint64)
}

func RefFloat32(value float32) *float32 {
	return &value
}

func DerefFloat32(value *float32, defaultValue ...float32) float32 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(float32)
}

func RefFloat64(value float64) *float64 {
	return &value
}

func DerefFloat64(value *float64, defaultValue ...float64) float64 {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(float64)
}

func RefByte(value byte) *byte {
	return &value
}

func DerefByte(value *byte, defaultValue ...byte) byte {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(byte)
}

func RefRune(value rune) *rune {
	return &value
}

func DerefRune(value *rune, defaultValue ...rune) rune {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(rune)
}

func RefBool(value bool) *bool {
	return &value
}

func DerefBool(value *bool, defaultValue ...bool) bool {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(bool)
}

func RefTime(value time.Time) *time.Time {
	return &value
}

func DerefTime(value *time.Time, defaultValue ...time.Time) time.Time {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(time.Time)
}
