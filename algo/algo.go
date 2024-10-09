package algo

type AggregationAlgo[V any, AGG any] interface {
	Insert(V)
	Evict()
	Query() AGG
}
