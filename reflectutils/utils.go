package reflectutils

import "reflect"

func DeepUnrefValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}

func DeepUnrefAndNewValue(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}
	return v
}

func DeepUnrefType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func IsUnexportedStructField(sf reflect.StructField) bool {
	return sf.PkgPath != "" && !sf.Anonymous
}

func ToSlice(i interface{}) []interface{} {
	drv := DeepUnrefValue(reflect.ValueOf(i))
	if drv.Kind() != reflect.Slice {
		panic("ice: to slice: require slice type")
	}
	l := drv.Len()
	r := make([]interface{}, l)
	for i := 0; i < l; i++ {
		r[i] = drv.Index(i).Interface()
	}
	return r
}

func GetStructFieldNames(i interface{}) []string {
	rv := DeepUnrefValue(reflect.ValueOf(i))
	if rv.Kind() != reflect.Struct {
		panic("ice: get struct field names: require struct type")
	}
	var r []string
	for i := 0; i < rv.NumField(); i++ {
		r = append(r, rv.Type().Field(i).Name)
	}
	return r
}
