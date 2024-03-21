package deptree

import (
	"reflect"
	"testing"
)

func TestDepTree_List(t *testing.T) {
	type fields struct {
		deps map[string][]string
	}
	type args struct {
		top string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "linear dependency tree",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
					"b": {"c", "e"},
					"c": {"d"},
					"e": {},
					"d": {},
				},
			},
			args: args{
				top: "a",
			},
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "linear dependency tree with top in the middle",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
					"b": {"c", "e"},
					"c": {"d"},
					"e": {},
					"d": {},
				},
			},
			args: args{
				top: "b",
			},
			want: []string{"b", "c", "d", "e"},
		},
		{
			name: "parallel trees 1",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
					"b": {"c", "e"},
					"c": {"d"},
					"e": {},
					"d": {},
					"1": {"2", "3"},
					"2": {},
					"3": {},
				},
			},
			args: args{
				top: "b",
			},
			want: []string{"b", "c", "d", "e"},
		},
		{
			name: "parallel trees 2",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
					"b": {"c", "e"},
					"c": {"d"},
					"e": {},
					"d": {},
					"1": {"2", "3"},
					"2": {},
					"3": {},
				},
			},
			args: args{
				top: "1",
			},
			want: []string{"1", "2", "3"},
		},
		{
			name: "branched tree",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
					"b": {"c", "e"},
					"c": {"d", "e"},
					"e": {"1", "2", "3"},
					"d": {},
					"1": {"2", "3"},
					"2": {},
					"3": {},
				},
			},
			args: args{
				top: "a",
			},
			want: []string{"a", "b", "c", "d", "e", "1", "2", "3"},
		},
		{
			name: "mixed trees 1",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
					"b": {"c", "e"},
					"c": {"d"},
					"e": {},
					"d": {"3"},
					"1": {"2", "3"},
					"2": {},
					"3": {},
				},
			},
			args: args{
				top: "a",
			},
			want: []string{"a", "b", "c", "d", "3", "e"},
		},
		{
			name: "empty tree",
			fields: fields{
				deps: map[string][]string{},
			},
			args: args{
				top: "a",
			},
			want: []string{},
		},
		{
			name: "single item tree ",
			fields: fields{
				deps: map[string][]string{
					"a": {},
				},
			},
			args: args{
				top: "a",
			},
			want: []string{"a"},
		},
		{
			name: "ignore not existent",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
				},
			},
			args: args{
				top: "a",
			},
			want: []string{"a"},
		},
		{
			name: "ignore not existent top",
			fields: fields{
				deps: map[string][]string{
					"a": {"b"},
				},
			},
			args: args{
				top: "b",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dt := &DepTree{
				deps: tt.fields.deps,
			}
			if got := dt.ListDesc(tt.args.top); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListDesc() = %v, want %v", got, tt.want)
			}
			reversedWant := make([]string, len(tt.want))
			for i, r := range tt.want {
				reversedWant[len(tt.want)-i-1] = r
			}
			if got := dt.ListAsc(tt.args.top); !reflect.DeepEqual(got, reversedWant) {
				t.Errorf("ListAsc() = %v, want %v", got, reversedWant)
			}
		})
	}
}
