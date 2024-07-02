package sets

import (
	"cmp"
	"slices"
)

type OrderdSet[T cmp.Ordered] struct {
	innerMap map[T]struct{}
}

func NewOrderedSet[T cmp.Ordered](items ...T) *OrderdSet[T] {
	innerMap := make(map[T]struct{})
	for _, item := range items {
		innerMap[item] = struct{}{}
	}
	return &OrderdSet[T]{
		innerMap: innerMap,
	}
}

func (s *OrderdSet[T]) Contains(t T) (found bool) {
	_, found = s.innerMap[t]
	return
}

func (s *OrderdSet[T]) Values() []T {
	ret := make([]T, 0, len(s.innerMap))
	for item := range s.innerMap {
		ret = append(ret, item)
	}
	slices.Sort(ret)
	return ret
}
