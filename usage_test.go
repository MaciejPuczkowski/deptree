package deptree

import (
	"reflect"
	"testing"
)

type testNode struct {
	nodeId string
	deps   []string
}

func (tn *testNode) NodeId() string {
	return tn.nodeId
}

func (tn *testNode) Deps() []string {
	return tn.deps
}

func TestUseCase_GettingStartedString(t *testing.T) {
	builder := NewDepTreeBuilder()
	builder.AddDeps("test", "test1", "test2")
	builder.AddDeps("test1", "test2", "test3")
	builder.ForceIntegrity()
	tree, err := builder.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []string{"test", "test1", "test3", "test2"}
	actual := tree.ListDesc("test")
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
	expected = []string{"test2", "test3", "test1", "test"}
	actual = tree.ListAsc("test")
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}

}
func TestUseCase_GettingStartedMultiTop(t *testing.T) {
	builder := NewDepTreeBuilder()
	builder.AddDeps("test", "test1", "test2")
	builder.AddDeps("test1", "test2", "test3")
	builder.AddDeps("test5", "test2", "test3")
	builder.ForceIntegrity()
	tree, err := builder.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []string{"test3", "test2", "test5", "test1", "test"}
	actual := tree.ListAsc("test", "test5")
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}

}
func TestUseCase_GettingStartedSortALL(t *testing.T) {
	builder := NewDepTreeBuilder()
	builder.AddDeps("test", "test1", "test2")
	builder.AddDeps("test1", "test2", "test3")
	builder.AddDeps("test5", "test2", "test3")
	builder.ForceIntegrity()
	tree, err := builder.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []string{"test3", "test2", "test1", "test", "test5"}
	actual := tree.ListAsc("test", "test5", "test1", "test3", "test2", "test", "test", "test3")
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}

}
func TestUseCase_GettingStartedSortALLNodes(t *testing.T) {
	nodes := []*testNode{
		{nodeId: "test", deps: []string{"test1", "test2"}},
		{nodeId: "test1", deps: []string{"test2", "test3"}},
		{nodeId: "test5", deps: []string{"test2", "test3"}},
		{nodeId: "test2", deps: []string{}},
		{nodeId: "test3", deps: []string{}},
	}
	builder := NewNDepTreeBuilder[*testNode]()
	for _, node := range nodes {
		builder.AddNode(node)
	}
	tree, err := builder.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := []string{"test3", "test2", "test5", "test1", "test"}
	actualNodes := tree.ListAsc(nodes...)
	actual := make([]string, len(actualNodes))
	for i, node := range actualNodes {
		actual[i] = node.NodeId()
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}

}

func TestUseCase_GettingStartedNode(t *testing.T) {
	nodes := []*testNode{
		{nodeId: "test", deps: []string{"test1", "test2"}},
		{nodeId: "test1", deps: []string{"test2", "test3"}},
		{nodeId: "test2", deps: []string{"test3", "test4"}},
		{nodeId: "test3", deps: []string{"test4", "test5"}},
		{nodeId: "test4", deps: []string{}},
		{nodeId: "test5", deps: []string{}},
	}
	builder := NewIDepTreeBuilder()
	for _, node := range nodes {
		builder.AddNode(node)
	}
	tree, _ := builder.Build()
	expected := []string{"test", "test1", "test2", "test3", "test5", "test4"}
	actualNodes := tree.ListDesc(nodes[0])
	actual := make([]string, len(actualNodes))
	for i, node := range actualNodes {
		actual[i] = node.NodeId()
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
	actualNodes = tree.ListDescStr("test")
	actual = make([]string, len(actualNodes))
	for i, node := range actualNodes {
		actual[i] = node.NodeId()
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}

}

func TestUseCase_GettingStartedTemplate(t *testing.T) {
	nodes := []*testNode{
		{nodeId: "test", deps: []string{"test1", "test2"}},
		{nodeId: "test1", deps: []string{"test2", "test3"}},
		{nodeId: "test2", deps: []string{"test3", "test4"}},
		{nodeId: "test3", deps: []string{"test4", "test5"}},
		{nodeId: "test4", deps: []string{}},
		{nodeId: "test5", deps: []string{}},
	}
	builder := NewNDepTreeBuilder[*testNode]()
	for _, node := range nodes {
		builder.AddNode(node)
	}
	tree, _ := builder.Build()
	expected := []string{"test", "test1", "test2", "test3", "test5", "test4"}
	actualNodes := tree.ListDesc(nodes[0])
	actual := make([]string, len(actualNodes))
	for i, node := range actualNodes {
		actual[i] = node.nodeId
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}
	actualNodes = tree.ListDescStr("test")
	actual = make([]string, len(actualNodes))
	for i, node := range actualNodes {
		actual[i] = node.nodeId
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected: %v, actual: %v", expected, actual)
	}

}
