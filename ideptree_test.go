package deptree

import (
	"reflect"
	"testing"
)

func TestIDepTreeBuilder_AddNode(t *testing.T) {
	type args struct {
		node Node
	}
	tests := []struct {
		name string
		dtb  *IDepTreeBuilder
		args args
	}{
		{
			name: "AddNode",
			dtb:  NewIDepTreeBuilder(),
			args: args{
				node: &testNode{
					nodeId: "test",
					deps:   []string{"test1", "test2"},
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

func TestIDepTreeBuilder_Build(t *testing.T) {
	tests := []struct {
		name    string
		dtb     IDepTreeBuilder
		want    *IDepTree
		wantErr bool
	}{
		{
			name: "Build",
			dtb: IDepTreeBuilder{
				builder: &DepTreeBuilder{deps: map[string][]string{
					"test":  {"test1", "test2"},
					"test1": {},
					"test2": {},
				}},
				nodes: map[string]Node{
					"test": &testNode{
						nodeId: "test",
						deps:   []string{"test1", "test2"},
					},
				},
			},
			want: &IDepTree{
				tree: &DepTree{deps: map[string][]string{"test": {"test1", "test2"}}},
				nodes: map[string]Node{
					"test": &testNode{
						nodeId: "test",
						deps:   []string{"test1", "test2"},
					},
				},
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
			if !reflect.DeepEqual(got.tree.deps["test"], tt.want.tree.deps["test"]) {
				t.Errorf("Build() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDepTree_List(t *testing.T) {
	type args struct {
		top Node
	}
	tests := []struct {
		name string
		dt   IDepTree
		args args
		want []Node
	}{
		{
			name: "ListDesc",
			dt: IDepTree{
				tree: &DepTree{deps: map[string][]string{"test": {"test1", "test2"}}},
				nodes: map[string]Node{
					"test": &testNode{
						nodeId: "test",
						deps:   []string{"test1", "test2"},
					},
				},
			},
			args: args{
				top: &testNode{
					nodeId: "test",
					deps:   []string{"test1", "test2"},
				},
			},
			want: []Node{
				&testNode{
					nodeId: "test",
					deps:   []string{"test1", "test2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.ListDesc(tt.args.top); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListDesc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDepTree_ListReversed(t *testing.T) {
	type args struct {
		top Node
	}
	tests := []struct {
		name string
		dt   IDepTree
		args args
		want []Node
	}{
		{
			name: "ListAsc",
			dt: IDepTree{
				tree: &DepTree{deps: map[string][]string{"test": {"test1", "test2"}}},
				nodes: map[string]Node{
					"test": &testNode{
						nodeId: "test",
						deps:   []string{"test1", "test2"},
					},
				},
			},
			args: args{
				top: &testNode{
					nodeId: "test",
					deps:   []string{"test1", "test2"},
				},
			},
			want: []Node{
				&testNode{
					nodeId: "test",
					deps:   []string{"test1", "test2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.ListAsc(tt.args.top); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAsc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDepTree_ListReversedStr(t *testing.T) {
	type args struct {
		top string
	}
	tests := []struct {
		name string
		dt   IDepTree
		args args
		want []Node
	}{
		{
			name: "ListAscStr",
			dt: IDepTree{
				tree: &DepTree{deps: map[string][]string{"test": {"test1", "test2"}}},
				nodes: map[string]Node{
					"test": &testNode{
						nodeId: "test",
						deps:   []string{"test1", "test2"},
					},
				},
			},
			args: args{
				top: "test",
			},
			want: []Node{
				&testNode{
					nodeId: "test",
					deps:   []string{"test1", "test2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.ListAscStr(tt.args.top); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListAscStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIDepTree_ListStr(t *testing.T) {
	type args struct {
		top string
	}
	tests := []struct {
		name string
		dt   IDepTree
		args args
		want []Node
	}{
		{
			name: "ListDescStr",
			dt: IDepTree{
				tree: &DepTree{deps: map[string][]string{"test": {"test1", "test2"}}},
				nodes: map[string]Node{
					"test": &testNode{
						nodeId: "test",
						deps:   []string{"test1", "test2"},
					},
				},
			},
			args: args{
				top: "test",
			},
			want: []Node{
				&testNode{
					nodeId: "test",
					deps:   []string{"test1", "test2"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dt.ListDescStr(tt.args.top); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListDescStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
