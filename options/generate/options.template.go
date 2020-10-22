package generate

import (
	"github.com/cheekybits/genny/generic"
	"reflect"
)

//go:generate genny -in=$GOFILE -out=options.generate.go gen "TOPTION=string,int,int8,int16,int32,int64,uint8,uint16,uint32,uint64,float32,float64,byte,rune,bool,time.Time"

type TOPTION generic.Type

type OptionTOPTION struct {
	Valid bool
	Value TOPTION
}

func (o OptionTOPTION) Unwrap(defaultValue ...TOPTION) TOPTION {
	if o.Valid {
		return o.Value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(o.Value)).Elem().Interface().(TOPTION)
}
