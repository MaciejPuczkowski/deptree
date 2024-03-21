package deptree

import (
	"reflect"
	"testing"
)

func TestDepTreeBuilder_AddDeps(t *testing.T) {
	type fields struct {
		deps map[string][]string
	}
	type args struct {
		node string
		deps []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		result map[string][]string
	}{
		{
			name:   "AddDeps to empty tree",
			fields: fields{deps: make(map[string][]string)},
			args: args{
				node: "a",
				deps: []string{"b", "c"},
			},
			result: map[string][]string{
				"a": {"b", "c"},
			},
		},
		{
			name:   "AddDeps to existing node",
			fields: fields{deps: map[string][]string{"a": {"b", "c"}}},
			args: args{
				node: "a",
				deps: []string{"d", "e"},
			},
			result: map[string][]string{
				"a": {"b", "c", "d", "e"},
			},
		},
		{
			name:   "AddDeps to existing tree",
			fields: fields{deps: map[string][]string{"d": {"b", "a"}}},
			args: args{
				node: "a",
				deps: []string{"b", "c"},
			},
			result: map[string][]string{
				"a": {"b", "c"},
				"d": {"b", "a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dtb := &DepTreeBuilder{
				deps: tt.fields.deps,
			}
			dtb.AddDeps(tt.args.node, tt.args.deps...)
			if !reflect.DeepEqual(dtb.deps, tt.result) {
				t.Errorf("AddDeps() = %v, want %v", dtb.deps, tt.result)
			}
		})
	}
}

func TestDepTreeBuilder_Build(t *testing.T) {
	type fields struct {
		deps map[string][]string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *DepTree
		wantErr bool
	}{
		{
			name:    "Build empty tree",
			fields:  fields{deps: make(map[string][]string)},
			want:    &DepTree{deps: make(map[string][]string)},
			wantErr: false,
		},
		{
			name:    "Build existing tree",
			fields:  fields{deps: map[string][]string{"a": {"b", "c"}, "b": {}, "c": {}}},
			want:    &DepTree{deps: map[string][]string{"a": {"b", "c"}, "b": {}, "c": {}}},
			wantErr: false,
		},
		{
			name:    "Build with dependency missed tree",
			fields:  fields{deps: map[string][]string{"a": {"b", "c"}}},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Build with cycle",
			fields:  fields{deps: map[string][]string{"a": {"b", "c"}, "b": {}, "c": {"d"}, "d": {"a"}}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dtb := &DepTreeBuilder{
				deps: tt.fields.deps,
			}
			got, err := dtb.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}
