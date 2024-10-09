package algo

import "container/list"

type SubOnEvict[V any, AGG any] struct {
	vals *list.List

	aggFunc func(AGG, V) AGG
	subFunc func(AGG, V) AGG
	agg     AGG
}

var _ AggregationAlgo[any, any] = (*SubOnEvict[any, any])(nil)

func NewSubOnEvict[V any, AGG any](
	aggFunc func(AGG, V) AGG,
	subFunc func(AGG, V) AGG,
	initAgg AGG,
) AggregationAlgo[V, AGG] {
	return &SubOnEvict[V, AGG]{
		vals:    list.New(),
		aggFunc: aggFunc,
		subFunc: subFunc,
		agg:     initAgg,
	}
}

func (s *SubOnEvict[V, AGG]) Insert(val V) {
	s.vals.PushBack(val)
	s.agg = s.aggFunc(s.agg, val)
}

func (s *SubOnEvict[V, AGG]) Evict() {
	val := s.vals.Remove(s.vals.Front()).(V)
	s.agg = s.subFunc(s.agg, val)
}

func (s *SubOnEvict[V, AGG]) Query() AGG {
	return s.agg
}
