package algo

import (
	"cmp"
	"container/list"
)

// Sliding window using ordered statistical tree
type OrderedStatsTreeSW[V cmp.Ordered] struct {
	vals *list.List
	tree *AVLNode[V]
}

func NewOrderedStatsTreeSW[V cmp.Ordered]() AggregationAlgo[V, V] {
	return &OrderedStatsTreeSW[V]{
		vals: list.New(),
		tree: nil,
	}
}

func (o *OrderedStatsTreeSW[V]) Insert(val V) {
	o.vals.PushBack(val)
	o.tree = o.tree.Insert(val)
}

func (o *OrderedStatsTreeSW[V]) Evict() {
	e := o.vals.Remove(o.vals.Front()).(V)
	o.tree = o.tree.Delete(e)
}

// Query the median value
func (o *OrderedStatsTreeSW[V]) Query() V {
	index := o.vals.Len() / 2
	return o.tree.Select(index)
}
