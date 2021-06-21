package reflectutil

import (
	"reflect"
	"testing"
	"time"
)

func TestMapStruct(t *testing.T) {
	type args struct {
		src     interface{}
		dst     interface{}
		nameMap map[string]string
	}
	tests := []struct {
		name string
		args args
	}{
		{

			name: "normal",
			args: args{
				src: struct {
					Name   string
					Age    int
					Pets   *[]string
					Hidden string
				}{
					Name: "Medivh",
					Age:  3,
					Pets: &[]string{"Foo", "Bar", "Baz"},
				},
				dst: &struct {
					Name    string
					Age2    int
					Pets    []string
					Created *time.Time
				}{},
				nameMap: map[string]string{
					"Age": "Age2",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MapStruct(tt.args.src, tt.args.dst, tt.args.nameMap)
			t.Logf("%+v", tt.args.dst)
		})
	}
}

func TestParseMapToStruct(t *testing.T) {
	type args struct {
		m map[string]interface{}
		i interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "normal",
			args: args{
				m: map[string]interface{}{
					"Name": "Medivh",
					"Age":  32,
					"Pets": []string{"Foo", "Bar", "Baz"},
				},
				i: &struct {
					Name string
					Age  int
					Pets []string
				}{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseMapToStruct(tt.args.m, tt.args.i)
			t.Logf("%+v\n", tt.args.i)
		})
	}
}

func TestParseStructToMap(t *testing.T) {
	type args struct {
		i interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "normal",
			args: args{
				i: struct {
					Name string
					Age  int
					Pets *[]string
				}{
					Name: "Medivh",
					Age:  3,
					Pets: &[]string{"Foo", "Bar", "Baz"},
				},
			},
			want: map[string]interface{}{
				"Name": "Medivh",
				"Age":  3,
				"Pets": &[]string{"Foo", "Bar", "Baz"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStructToMap(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStructToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
