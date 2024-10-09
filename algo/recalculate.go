package algo

import "container/list"

type Recalculate[V any, AGG any] struct {
	vals *list.List

	aggFunc func(AGG, V) AGG
	initAgg AGG
}

var _ AggregationAlgo[any, any] = (*Recalculate[any, any])(nil)

func NewRecalculate[V any, AGG any](aggFunc func(AGG, V) AGG, initAgg AGG) AggregationAlgo[V, AGG] {
	r := Recalculate[V, AGG]{
		vals:    list.New(),
		aggFunc: aggFunc,
		initAgg: initAgg,
	}
	return &r
}

func (r *Recalculate[V, AGG]) Insert(val V) {
	r.vals.PushBack(val)
}

func (r *Recalculate[V, AGG]) Evict() {
	r.vals.Remove(r.vals.Front())
}

func (r *Recalculate[V, AGG]) Query() AGG {
	agg := r.initAgg
	p := r.vals.Front()
	for p != nil {
		v, ok := p.Value.(V)
		if !ok {
			panic("invalid value")
		}
		agg = r.aggFunc(agg, v)
		p = p.Next()
	}
	return agg
}
