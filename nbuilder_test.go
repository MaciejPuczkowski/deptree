package deptree

import (
	"reflect"
	"testing"
)

func TestNDepTreeBuilder_AddNode(t *testing.T) {
	type args[N Node] struct {
		node N
	}
	type testCase[N Node] struct {
		name string
		dtb  NDepTreeBuilder[N]
		args args[N]
	}
	tests := []testCase[*testNode]{
		{
			name: "test add node",
			dtb: NDepTreeBuilder[*testNode]{
				builder: NewDepTreeBuilder(),
				nodes:   make(map[string]*testNode),
			},
			args: args[*testNode]{
				node: &testNode{
					nodeId: "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.dtb.AddNode(tt.args.node)
		})
	}
}

func TestNDepTreeBuilder_Build(t *testing.T) {
	type testCase[N Node] struct {
		name    string
		dtb     NDepTreeBuilder[N]
		want    *NDepTree[N]
		wantErr bool
	}
	tests := []testCase[*testNode]{
		{
			name: "test build",
			dtb: NDepTreeBuilder[*testNode]{
				builder: &DepTreeBuilder{
					deps: make(map[string][]string),
				},
				nodes: make(map[string]*testNode),
			},
			want: &NDepTree[*testNode]{
				tree:  &DepTree{deps: make(map[string][]string)},
				nodes: make(map[string]*testNode),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dtb.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() got = %v, want %v", got.tree.deps, tt.want.tree.deps)
			}
		})
	}
}

func TestNewNDepTreeBuilder(t *testing.T) {
	type testCase[N Node] struct {
		name string
		want *NDepTreeBuilder[N]
	}
	tests := []testCase[*testNode]{
		{
			name: "test new node builder",
			want: &NDepTreeBuilder[*testNode]{
				builder: NewDepTreeBuilder(),
				nodes:   make(map[string]*testNode),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNDepTreeBuilder[*testNode](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNDepTreeBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
