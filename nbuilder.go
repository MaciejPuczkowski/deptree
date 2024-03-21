package deptree

// NDepTreeBuilder collects nodes needed to build a dependency tree. The node is an object implementing Node interface.
// The difference between IDepTreeBuilder and NDepTreeBuilder is that IDepTreeBuilder may be used for
// object of different types implementing Node interface, but then the client code must be aware of the type of
// the object. NDepTreeBuilder is a generic type and returns a tree specific to the given type. The client code is
// limited to the type but the client code doesn't have to check the returned types.
type NDepTreeBuilder[N Node] struct {
	nodes   map[string]N
	builder *DepTreeBuilder
}

// NewNDepTreeBuilder returns a new NDepTreeBuilder.
func NewNDepTreeBuilder[N Node]() *NDepTreeBuilder[N] {
	return &NDepTreeBuilder[N]{
		builder: NewDepTreeBuilder(),
		nodes:   make(map[string]N),
	}
}

// AddNode adds a node to the NDepTreeBuilder.
func (dtb *NDepTreeBuilder[N]) AddNode(node N) {
	if _, ok := dtb.nodes[node.NodeId()]; !ok {
		dtb.nodes[node.NodeId()] = node
	}
	dtb.builder.AddDeps(node.NodeId(), node.Deps()...)
}

// Build builds a dependency tree from the NDepTreeBuilder. If node for a dependency is not added to the NDepTreeBuilder
// an integrity error will be returned.
func (dtb *NDepTreeBuilder[N]) Build() (*NDepTree[N], error) {
	tree, err := dtb.builder.Build()
	if err != nil {
		return nil, err
	}
	return &NDepTree[N]{nodes: dtb.nodes, tree: tree}, nil
}
