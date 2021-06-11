package ice

import (
	"reflect"
	"testing"
	"time"
)

func TestBindStruct(t *testing.T) {
	type args struct {
		rv      reflect.Value
		vvm     map[string][]string
		binders []Binder
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				rv: reflect.ValueOf(&struct {
					Name    string
					Age     int
					Pets    *[]string
					Created *time.Time
					hidden  string
				}{}),
				vvm: map[string][]string{
					"Name":    {"Medivh", "Mike"},
					"Age":     {"28"},
					"Pets":    {"Foo", "Bar", "Baz"},
					"Created": {"2020-02-22T22:22:22Z"},
					"hidden":  {"secret"},
				},
				binders: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bindStruct(tt.args.rv, tt.args.vvm, tt.args.binders...); (err != nil) != tt.wantErr {
				t.Errorf("BindStructDefault() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("%+v", tt.args.rv.Interface())
		})
	}
}

func TestBindList(t *testing.T) {
	type args struct {
		rv      reflect.Value
		vv      []string
		binders []Binder
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				rv: reflect.ValueOf(&[]int{}),
				vv: []string{"1", "2", "3"}, binders: nil,
			},
			wantErr: false,
		},
		{
			name: "parse fail",
			args: args{
				rv: reflect.ValueOf(&[]int{}),
				vv: []string{"1", "2", "3", "a"}, binders: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := bindList(tt.args.rv, tt.args.vv, tt.args.binders...); (err != nil) != tt.wantErr {
				t.Errorf("BindListDefault() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("%+v", tt.args.rv.Interface())
		})
	}
}
