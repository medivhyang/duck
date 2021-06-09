package reflectutils

import (
	"reflect"
	"testing"
)

func TestParseStructTag(t *testing.T) {
	type args struct {
		i   interface{}
		key string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "normal",
			args: args{
				i: struct {
					Name string `name:"NewName"`
					Age  int    `name:"NewAge"`
					Pets []string
				}{},
				key: "name",
			},
			want: map[string]string{
				"Name": "NewName",
				"Age":  "NewAge",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStructTag(tt.args.i, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStructTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseStructTags(t *testing.T) {
	type args struct {
		i    interface{}
		keys []string
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]string
	}{
		{
			name: "normal",
			args: args{
				i: struct {
					Name string   `name:"NewName" require:"true"`
					Age  int      `name:"NewAge" require:"false"`
					Pets []string `name:"NewPets"`
				}{},
				keys: []string{"name", "require"},
			},
			want: map[string]map[string]string{
				"Name": {
					"name":    "NewName",
					"require": "true",
				},
				"Age": {
					"name":    "NewAge",
					"require": "false",
				},
				"Pets": {
					"name": "NewPets",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStructTags(tt.args.i, tt.args.keys...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStructTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseStructChildTags(t *testing.T) {
	type args struct {
		i         interface{}
		parentKey string
		itemSep   string
		kvSep     string
		keys      []string
	}
	tests := []struct {
		name string
		args args
		want map[string]map[string]string
	}{
		{
			name: "normal",
			args: args{
				i: struct {
					Name string `binding:"name=NewName,require=true,default=xxx"`
					Age  int    `binding:"name=NewAge,require=false"`
				}{},
				parentKey: "binding",
				itemSep:   ",",
				kvSep:     "=",
				keys:      []string{"name", "require"},
			},
			want: map[string]map[string]string{
				"Name": {
					"name":    "NewName",
					"require": "true",
				},
				"Age": {
					"name":    "NewAge",
					"require": "false",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseStructChildTags(tt.args.i, tt.args.parentKey, tt.args.itemSep, tt.args.kvSep, tt.args.keys...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseStructChildTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
