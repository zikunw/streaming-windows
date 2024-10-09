package algo

import "container/list"

type StackEntry[V any, AGG any] struct {
	val V
	agg AGG
}

type TwoStack[V any, AGG any] struct {
	front *list.List
	back  *list.List

	aggFunc     func(AGG, V) AGG
	combineFunc func(AGG, AGG) AGG
	initAgg     AGG
}

var _ AggregationAlgo[any, any] = (*TwoStack[any, any])(nil)

func NewTwoStack[V any, AGG any](
	aggFunc func(AGG, V) AGG,
	combineFunc func(AGG, AGG) AGG,
	initAgg AGG,
) AggregationAlgo[V, AGG] {
	return &TwoStack[V, AGG]{
		front:   list.New(),
		back:    list.New(),
		aggFunc: aggFunc,
		initAgg: initAgg,
	}
}

func (t *TwoStack[V, AGG]) getFrontAgg() AGG {
	if t.front.Len() == 0 {
		return t.initAgg
	}
	return t.front.Front().Value.(StackEntry[V, AGG]).agg
}

func (t *TwoStack[V, AGG]) getBackAgg() AGG {
	if t.back.Len() == 0 {
		return t.initAgg
	}
	return t.back.Front().Value.(StackEntry[V, AGG]).agg
}

func (t *TwoStack[V, AGG]) Insert(val V) {
	t.back.PushFront(StackEntry[V, AGG]{
		val: val,
		agg: t.aggFunc(t.getBackAgg(), val),
	})
}

func (t *TwoStack[V, AGG]) Evict() {
	// flip if front stack is empty
	if t.front.Len() == 0 {
		for t.back.Len() > 0 {
			val := t.back.Front().Value.(StackEntry[V, AGG]).val
			t.front.PushFront(StackEntry[V, AGG]{
				val: val,
				agg: t.aggFunc(t.getFrontAgg(), val),
			})
			t.back.Remove(t.back.Front())
		}
	}
	t.front.Remove(t.front.Front())
}

func (t *TwoStack[V, AGG]) Query() AGG {
	return t.combineFunc(t.getFrontAgg(), t.getBackAgg())
}
