package reflectutils

import (
	"reflect"
	"strings"
)

func ParseStructTag(i interface{}, key string) map[string]string {
	r := map[string]string{}
	m := ParseStructTags(i, key)
	for k, vm := range m {
		if v, ok := vm[key]; ok {
			r[k] = v
		}
	}
	return r
}

func ParseStructTags(i interface{}, keys ...string) map[string]map[string]string {
	if len(keys) == 0 {
		return map[string]map[string]string{}
	}
	rv := DeepUnrefValue(reflect.ValueOf(i))
	if rv.Kind() != reflect.Struct {
		panic("ice: parse struct tags to map: require struct type")
	}
	r := map[string]map[string]string{}
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		rf := rt.Field(i)
		r[rf.Name] = map[string]string{}
		for _, k := range keys {
			if v, ok := rf.Tag.Lookup(k); ok {
				r[rf.Name][k] = v
			}
		}
	}
	return r
}

func ParseStructChildTags(i interface{}, parentKey string, itemSep string, kvSep string, keys ...string) map[string]map[string]string {
	rv := DeepUnrefValue(reflect.ValueOf(i))
	if rv.Kind() != reflect.Struct {
		panic("ice: parse struct child tags to map: require struct type")
	}
	r := map[string]map[string]string{}
	rt := reflect.TypeOf(i)
	for i := 0; i < rt.NumField(); i++ {
		var (
			f     = rt.Field(i)
			s     = f.Tag.Get(parentKey)
			items []string
			kv    []string
			k, v  string
		)
		r[f.Name] = map[string]string{}
		items = strings.Split(s, itemSep)
		for _, item := range items {
			kv = strings.Split(item, kvSep)
			if len(kv) >= 1 {
				k = kv[0]
			}
			if len(kv) >= 2 {
				v = kv[1]
			}
			if len(keys) == 0 || containString(keys, k) {
				r[f.Name][k] = v
			}
		}
	}
	return r
}

func containString(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
