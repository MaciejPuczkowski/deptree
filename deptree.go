package deptree

import "fmt"

var ErrIntegrity = fmt.Errorf("integrity error")

// DepTree is the main dependency manager.
type DepTree struct {
	deps map[string][]string
}

func (dt *DepTree) ListAsc(top ...string) []string {
	result := make([]string, 0)
	for _, t := range top {
		result = dt.sanitize(append(dt.sanitize(result), dt.listFor(t)...))
	}
	return result
}

func (dt *DepTree) ListDesc(top ...string) []string {
	rs := dt.ListAsc(top...)
	result := make([]string, len(rs))
	for i, r := range rs {
		result[len(rs)-i-1] = r
	}
	return result
}

func (dt *DepTree) sanitize(deps []string) []string {
	result := make([]string, 0)
	test := make(map[string]bool)
	for _, d := range deps {
		test[d] = false
	}
	l := len(deps) - 1
	for i := range deps {
		j := l - i
		d := deps[j]
		if test[d] {
			continue
		}
		test[d] = true
		result = append(result, d)
	}
	return result
}

func (dt *DepTree) listFor(top string) []string {
	if _, ok := dt.deps[top]; !ok {
		return []string{}
	}
	result := []string{top}
	for _, dep := range dt.deps[top] {
		result = append(result, dt.listFor(dep)...)
	}
	return result
}
