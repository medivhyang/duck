package reflectutil

import (
	"reflect"
)

func MapStruct(src interface{}, dst interface{}, nameMapFunc func(string) string) {
	srvRv := reflect.ValueOf(src)
	if DeepUnrefType(srvRv.Type()).Kind() != reflect.Struct {
		panic("ice: map struct: require struct type")
	}
	dstRv := reflect.ValueOf(dst)
	if dstRv.Type().Kind() != reflect.Ptr && DeepUnrefType(dstRv.Type()).Kind() != reflect.Struct {
		panic("ice: map struct: require struct pointer type")
	}
	m := ParseStructToMap(src)
	copyMap := make(map[string]interface{}, len(m))
	for k, v := range m {
		copyMap[k] = v
	}
	for name, value := range copyMap {
		newName := nameMapFunc(name)
		if newName != name {
			m[newName] = value
			delete(m, name)
		}
	}
	ParseMapToStruct(m, dst, nil)
}

func ParseMapToStruct(m map[string]interface{}, i interface{}, nameMapFunc func(string) string) {
	dstRv := reflect.ValueOf(i)
	dstRt := dstRv.Type()
	if dstRt.Kind() != reflect.Ptr && DeepUnrefType(dstRt).Kind() != reflect.Struct {
		panic("ice: parse map to struct: require struct pointer type")
	}
	drv := DeepUnrefAndNewValue(dstRv)
	fieldMap := map[string]reflect.Value{}
	for i := 0; i < drv.NumField(); i++ {
		fieldMap[drv.Type().Field(i).Name] = drv.Field(i)
	}
	for k, v := range m {
		if nameMapFunc != nil {
			k = nameMapFunc(k)
		}
		fv, ok := fieldMap[k]
		if !ok {
			continue
		}
		ft := fv.Type()
		srvRv := reflect.ValueOf(v)
		srvRt := srvRv.Type()
		if DeepUnrefType(srvRt) != DeepUnrefType(fv.Type()) {
			continue
		}
		switch DeepUnrefType(srvRt).Kind() {
		case reflect.Slice:
			l := DeepUnrefValue(srvRv).Len()
			DeepUnrefAndNewValue(fv).Set(reflect.MakeSlice(DeepUnrefType(ft), l, l))
			fallthrough
		case reflect.Array:
			reflect.Copy(DeepUnrefAndNewValue(fv), DeepUnrefAndNewValue(srvRv))
		default:
			fv.Set(reflect.ValueOf(v))
		}
	}
}

func ParseStructToMap(i interface{}) map[string]interface{} {
	rv := DeepUnrefValue(reflect.ValueOf(i))
	if rv.Kind() != reflect.Struct {
		panic("ice: parse struct to map: require struct type")
	}
	r := map[string]interface{}{}
	for i := 0; i < rv.NumField(); i++ {
		r[rv.Type().Field(i).Name] = rv.Field(i).Interface()
	}
	return r
}
