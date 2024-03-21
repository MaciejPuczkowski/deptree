package deptree

import (
	"fmt"
	"strings"
)

// DepTreeBuilder builds a dependency tree based on the strings representing node ids.
type DepTreeBuilder struct {
	isIntegral bool
	deps       map[string][]string
}

// NewDepTreeBuilder returns a new dependency tree builder.
func NewDepTreeBuilder() *DepTreeBuilder {
	return &DepTreeBuilder{
		deps: make(map[string][]string),
	}
}

// AddDeps adds dependencies to the dependency tree.
func (dtb *DepTreeBuilder) AddDeps(node string, deps ...string) {
	if _, ok := dtb.deps[node]; !ok {
		dtb.deps[node] = make([]string, 0)
	}
	dtb.deps[node] = append(dtb.deps[node], deps...)
	dtb.isIntegral = false
}

// Build builds a dependency tree from the dependency tree builder. Error is returned if the dependency tree
// contains a cycle or violates an integrity. The integrity is violated if a node for a dependency is not added to the
// builder. It means that if you provide "B" as dependency for "A", then you need to provide "B" with no dependencies.
// You can also call function ForceIntegrity() that automatically adds missing nodes to the builder.
func (dtb *DepTreeBuilder) Build() (*DepTree, error) {
	if err := dtb.integrityCheck(); err != nil {
		return nil, err
	}
	if err := dtb.cyclesCheck(); err != nil {
		return nil, err
	}
	newMap := make(map[string][]string)
	for k, v := range dtb.deps {
		newMap[k] = make([]string, len(v))
		for i, d := range v {
			newMap[k][i] = d
		}
	}
	return &DepTree{deps: newMap}, nil
}

// ForceIntegrity adds missing nodes to the dependency tree builder. You can call this function if you don't want to
// provide nodes with empty dependencies, and you know it is not an issue for a client code.
func (dtb *DepTreeBuilder) ForceIntegrity() {
	dtb.isIntegral = true
	for _, deps := range dtb.deps {
		for _, dep := range deps {
			if _, ok := dtb.deps[dep]; !ok {
				dtb.deps[dep] = make([]string, 0)
			}
		}
	}
}

func (dtb *DepTreeBuilder) integrityCheck() error {
	for _, children := range dtb.deps {
		for _, child := range children {
			if _, ok := dtb.deps[child]; !ok {
				return fmt.Errorf("%w: missing dependency \"%s\"", ErrIntegrity, child)
			}
		}
	}
	return nil
}

func (dtb *DepTreeBuilder) cyclesCheck() error {
	for node, _ := range dtb.deps {
		if ch, err := dtb.cycleCheckFor(node, node); err != nil {
			return fmt.Errorf("%w: %s", err, strings.Join(ch, "->"))
		}
	}
	return nil
}

func (dtb *DepTreeBuilder) cycleCheckFor(top, current string) ([]string, error) {
	chain := []string{current}
	if deps, ok := dtb.deps[current]; ok {
		for _, dep := range deps {
			if top == dep {
				return chain, fmt.Errorf("%w: cycle detected", ErrIntegrity)
			}
			ch, err := dtb.cycleCheckFor(top, dep)
			if ch != nil {
				chain = append(chain, ch...)
			}
			if err != nil {
				return chain, err
			}
		}
	}
	return chain, nil
}
