package intervaltree

import (
	"github.com/rdleal/intervalst/interval"
	"golang.org/x/exp/constraints"
)

type CmpFn[I constraints.Ordered] func(x, y I) int

type IntervalTree[I constraints.Ordered, V any] struct {
	tree *interval.MultiValueSearchTree[V, I]
}

func NewIntervalTree[I constraints.Ordered, V any](cmp CmpFn[I]) *IntervalTree[I, V] {
	return &IntervalTree[I, V]{
		tree: interval.NewMultiValueSearchTreeWithOptions[V, I](
			interval.CmpFunc[I](cmp), interval.TreeWithIntervalPoint(),
		),
	}
}

func (it *IntervalTree[I, V]) Add(i, j I, value V) *IntervalTree[I, V] {
	if i > j {
		i, j = j, i
	}
	_ = it.tree.Insert(i, j, value)
	return it
}

func (it *IntervalTree[I, V]) At(i I) []V {
	return it.Intersect(i, i)
}

func (it *IntervalTree[I, V]) Intersect(i, j I) []V {
	values, _ := it.tree.AllIntersections(i, j)
	return values
}
