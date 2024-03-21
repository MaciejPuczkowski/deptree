# DepTree
is a package that allows you to sort the dependencies tree to a flat list. It may be
used t sort imports, packages, migrations, executions and many other things that needs to be listed or 
executed in a proper order. If "A" depends on "B", you need to execute B before
A. It's what the package does. Takes your dependency tree and gives your flat list to go through.


# Usage:
### Tree based on strings:
```go
builder := NewDepTreeBuilder()
// add node test with dependencies to test1 and test2
builder.AddDeps("test", "test1", "test2")
// add node test1 with dependencies to test2 and test3
builder.AddDeps("test1", "test2", "test3")
// create nodes for the depenencies
builder.ForceIntegrity()
tree, _ := builder.Build()
list = tree.ListAsc("test")
```
The very basic usage is by sorting strings representing the node ids. The ids must be unique.
Method "ForceIntegrity" is called to create the nodes for the depenencies we did not provide as the nodes.
Otherwise the Build() would return integrity error, because we did not provide nodes: "test2" and "test3"
with empty dependencies, by default the builder will assume that the dependency is missing and return an error.

### Using generic builder:
```go
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
tree, _ := builder.Build()
nodes := tree.ListAsc(node[0])

```
Generic builder **NDepTreeBuilder** allows you to use object of the specific types as the nodes
as long their type implements the **Node** interface. The interface requires methods NodeId() and Deps() which are
self-explanatory. For the generic builder we provide the nodes by **AddNode** function. This
time we can't force integrity because the builder doesn't know how to create the missing object.
So all dependencies must be provided as the nodes, otherwise **Build** will return the integrity error.
**ListAsc** returns a list of the specific objects provided to the generic template.

## Using interface:
```go
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
nodes, _ := tree.ListAsc(nodes[0])
```
Using interfaces is very similar to using generics. The difference is that we
provide the nodes as the list of the **Node** interface. It allows us to sort objects
of different types but it leaves to the client code checking the proper types. 
Especially the **ListAsc** function returns list of the **Node** interface which
is not much useful for the client code most likely. It requires further type's checking.