package reflectutil

import (
	"errors"
	"reflect"
	"strconv"
	"time"
)

type Binder interface {
	Match(rv reflect.Value) bool
	Bind(rv reflect.Value, s string) error
}

var DefaultBinders = []Binder{&TimeBinder{}, &BaseBinder{}}

func Bind(i interface{}, s string, binders ...Binder) error {
	if len(binders) == 0 {
		binders = DefaultBinders
	}
	return bind(reflect.ValueOf(i), s, binders...)
}

func bind(rv reflect.Value, s string, binders ...Binder) error {
	for _, b := range binders {
		if b.Match(rv) {
			return b.Bind(rv, s)
		}
	}
	return errors.New("ice: bind: unsupported type")
}

func BindList(i interface{}, ss []string, binders ...Binder) error {
	if len(binders) == 0 {
		binders = DefaultBinders
	}
	return bindList(reflect.ValueOf(i), ss, binders...)
}

func bindList(rv reflect.Value, ss []string, binders ...Binder) error {
	drv := DeepUnrefAndNewValue(rv)
	switch DeepUnrefType(rv.Type()).Kind() {
	case reflect.Slice:
		drv.Set(reflect.MakeSlice(drv.Type(), len(ss), len(ss)))
		for i, v := range ss {
			if err := bind(drv.Index(i), v, binders...); err != nil {
				return err
			}
		}
		return nil
	case reflect.Array:
		if len(ss) >= drv.Len() {
			return errors.New("ice: bind list: too much values")
		}
		for i, v := range ss {
			return bind(drv.Field(i), v, binders...)
		}
		return nil
	default:
		return errors.New("ice: bind list: unsupported type")
	}
}

func BindStruct(i interface{}, values map[string][]string, binders ...Binder) error {
	if len(binders) == 0 {
		binders = DefaultBinders
	}
	return bindStruct(reflect.ValueOf(i), values, binders...)
}

func bindStruct(rv reflect.Value, values map[string][]string, binders ...Binder) error {
	return bindStructFunc(rv, func(s string) []string {
		return values[s]
	}, binders...)
}

func BindStructFunc(i interface{}, valuesFn func(string) []string, binders ...Binder) error {
	if len(binders) == 0 {
		binders = DefaultBinders
	}
	return bindStructFunc(reflect.ValueOf(i), valuesFn, binders...)
}

func bindStructFunc(rv reflect.Value, valuesFn func(string) []string, binders ...Binder) error {
	if valuesFn == nil {
		return nil
	}
	if len(binders) == 0 {
		binders = DefaultBinders
	}
	drv := DeepUnrefValue(rv)
	switch drv.Kind() {
	case reflect.Struct:
		for i := 0; i < drv.NumField(); i++ {
			if IsUnexportedStructField(drv.Type().Field(i)) {
				continue
			}
			var (
				fv   = drv.Field(i)
				name = drv.Type().Field(i).Name
				vv   = valuesFn(name)
			)
			if len(vv) == 0 {
				continue
			}
			switch DeepUnrefType(fv.Type()).Kind() {
			case reflect.Array, reflect.Slice:
				if err := bindList(fv, vv, binders...); err != nil {
					return err
				}
			default:
				v := ""
				if len(vv) > 0 {
					v = vv[0]
				}
				if err := bind(fv, v, binders...); err != nil {
					return err
				}
			}
		}
	default:
		return errors.New("ice: bind struct: unsupported type")
	}
	return nil
}

type BaseBinder struct{}

func (b *BaseBinder) Match(rv reflect.Value) bool {
	drv := DeepUnrefValue(rv)
	switch drv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		return true
	default:
		return false
	}
}

func (b *BaseBinder) Bind(rv reflect.Value, v string) error {
	switch rv.Kind() {
	case reflect.String:
		rv.SetString(v)
	case reflect.Int:
		return b.bindInt(rv, v, 10, 0)
	case reflect.Int8:
		return b.bindInt(rv, v, 10, 8)
	case reflect.Int16:
		return b.bindInt(rv, v, 10, 16)
	case reflect.Int32:
		return b.bindInt(rv, v, 10, 32)
	case reflect.Int64:
		return b.bindInt(rv, v, 10, 64)
	case reflect.Uint:
		return b.bindInt(rv, v, 10, 0)
	case reflect.Uint8:
		return b.bindInt(rv, v, 10, 8)
	case reflect.Uint16:
		return b.bindUint(rv, v, 10, 16)
	case reflect.Uint32:
		return b.bindUint(rv, v, 10, 32)
	case reflect.Uint64:
		return b.bindUint(rv, v, 10, 64)
	case reflect.Bool:
		return b.bindBool(rv, v)
	case reflect.Float32:
		return b.bindFloat(rv, v, 32)
	case reflect.Float64:
		return b.bindFloat(rv, v, 64)
	}
	return nil
}

func (b *BaseBinder) bindInt(rv reflect.Value, v string, base int, size int) error {
	if v == "" {
		v = "0"
	}
	i, err := strconv.ParseInt(v, base, size)
	if err != nil {
		return err
	}
	rv.SetInt(i)
	return nil
}

func (b *BaseBinder) bindUint(rv reflect.Value, v string, base int, size int) error {
	if v == "" {
		v = "0"
	}
	ui, err := strconv.ParseUint(v, base, size)
	if err != nil {
		return err
	}
	rv.SetUint(ui)
	return nil
}

func (b *BaseBinder) bindFloat(rv reflect.Value, v string, size int) error {
	if v == "" {
		v = "0"
	}
	f, err := strconv.ParseFloat(v, size)
	if err != nil {
		return err
	}
	rv.SetFloat(f)
	return nil
}

func (b *BaseBinder) bindBool(rv reflect.Value, v string) error {
	if v == "" {
		v = "0"
	}
	bl, err := strconv.ParseBool(v)
	if err != nil {
		return err
	}
	rv.SetBool(bl)
	return nil
}

type TimeBinder struct {
	Layout   string
	Location *time.Location
}

func (b *TimeBinder) Match(rv reflect.Value) bool {
	_, ok := rv.Interface().(time.Time)
	if ok {
		return true
	}
	_, ok2 := rv.Interface().(*time.Time)
	return ok2
}

func (b *TimeBinder) Bind(rv reflect.Value, v string) error {
	layout := time.RFC3339
	if b.Layout != "" {
		layout = time.RFC3339
	}
	location := time.Local
	if b.Location != nil {
		location = b.Location
	}
	t, err := time.ParseInLocation(layout, v, location)
	if err != nil {
		return err
	}
	switch rv.Interface().(type) {
	case time.Time:
		rv.Set(reflect.ValueOf(t))
	case *time.Time:
		rv.Set(reflect.ValueOf(&t))
	}
	return nil
}
