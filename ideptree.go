package deptree

// IDepTreeBuilder collects nodes needed to build a dependency tree. The node is an object implementing Node interface.
// The difference between IDepTreeBuilder and NDepTreeBuilder is that IDepTreeBuilder may be used for
// object of different types implementing Node interface, but then the client code must be aware of the type of
// the object. NDepTreeBuilder is a generic type and returns a tree specific to the given type. The client code is
// limited to the type but the client code doesn't have to check the returned types.
type IDepTreeBuilder NDepTreeBuilder[Node]

// NewIDepTreeBuilder returns a new IDepTreeBuilder.
func NewIDepTreeBuilder() *IDepTreeBuilder {
	return (*IDepTreeBuilder)(&NDepTreeBuilder[Node]{
		builder: NewDepTreeBuilder(),
		nodes:   make(map[string]Node),
	})
}

// AddNode adds a node to the IDepTreeBuilder.
func (dtb *IDepTreeBuilder) AddNode(node Node) {
	(*NDepTreeBuilder[Node])(dtb).AddNode(node)
}

// Build builds a dependency tree from the IDepTreeBuilder. If node for a dependency is not added
func (dtb *IDepTreeBuilder) Build() (*IDepTree, error) {
	t, err := (*NDepTreeBuilder[Node])(dtb).Build()
	if err != nil {
		return nil, err
	}
	return (*IDepTree)(t), nil
}

// IDepTree is an object sorting dependencies to the lists
type IDepTree NDepTree[Node]

// ListAsc sorts the dependencies bases on provided top nodes. Ascending order means that dependency comes before
// the node. Many tops may be sorted at once. The function returns the nodes in order merged properly for all tops.
func (dt *IDepTree) ListAsc(top ...Node) []Node {
	return (*NDepTree[Node])(dt).ListAsc(top...)
}

// ListAscStr takes strings representing node ids. See ListAsc for more details.
func (dt *IDepTree) ListAscStr(top ...string) []Node {
	return (*NDepTree[Node])(dt).ListAscStr(top...)
}

// ListDesc sorts the dependencies bases on provided top nodes. Descending order means that dependency comes after
// the node. Many tops may be sorted at once. The function returns the nodes in order merged properly for all tops.
func (dt *IDepTree) ListDesc(top ...Node) []Node {
	return (*NDepTree[Node])(dt).ListDesc(top...)
}

// ListDescStr takes strings representing node ids. See ListDesc for more details.
func (dt *IDepTree) ListDescStr(top ...string) []Node {
	return (*NDepTree[Node])(dt).ListDescStr(top...)
}
