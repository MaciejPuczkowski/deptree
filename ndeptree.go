package deptree

// Node is an interface for any object that can be used as a node in a dependency tree.
type Node interface {
	NodeId() string
	Deps() []string
}

// NDepTree is an object sorting dependencies to the lists
type NDepTree[N Node] struct {
	nodes map[string]N
	tree  *DepTree
}

// ListAsc sorts the dependencies bases on provided top nodes. Ascending order means that dependency comes before
// the node. Many tops may be sorted at once. The function returns the nodes in order merged properly for all tops.
func (dt *NDepTree[N]) ListAsc(top ...N) []N {
	return dt.ListAscStr(dt.stringify(top)...)
}

// ListAscStr takes strings representing node ids. See ListAsc for more details.
func (dt *NDepTree[N]) ListAscStr(top ...string) []N {
	ls := dt.tree.ListAsc(top...)
	result := make([]N, len(ls))
	for i, l := range ls {
		result[i] = dt.nodes[l]
	}
	return result
}

// ListDesc sorts the dependencies bases on provided top nodes. Descending order means that dependency comes after
// the node. Many tops may be sorted at once. The function returns the nodes in order merged properly for all tops.
func (dt *NDepTree[N]) ListDesc(top ...N) []N {
	return dt.ListDescStr(dt.stringify(top)...)
}

func (dt *NDepTree[N]) stringify(nodes []N) []string {
	ids := make([]string, len(nodes))
	for i, n := range nodes {
		ids[i] = n.NodeId()
	}
	return ids
}

// ListDescStr takes strings representing node ids. See ListDesc for more details.
func (dt *NDepTree[N]) ListDescStr(top ...string) []N {
	ls := dt.tree.ListDesc(top...)
	result := make([]N, len(ls))
	for i, l := range ls {
		result[i] = dt.nodes[l]
	}
	return result
}
