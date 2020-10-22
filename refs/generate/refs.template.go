package generate

import (
	"github.com/cheekybits/genny/generic"
	"reflect"
)

//go:generate genny -in=$GOFILE -out=refs.generate.go gen "TREF=string,int,int8,int16,int32,int64,uint8,uint16,uint32,uint64,float32,float64,byte,rune,bool,time.Time"

type TREF generic.Type

func RefTREF(value TREF) *TREF {
	return &value
}

func UnrapTREF(value *TREF, defaultValue ...TREF) TREF {
	if value != nil {
		return *value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return reflect.New(reflect.TypeOf(value).Elem()).Elem().Interface().(TREF)
}
