package sets

type Set[T comparable] struct {
	innerMap map[T]struct{}
}

func NewSet[T comparable](items ...T) *Set[T] {
	innerMap := make(map[T]struct{})
	for _, item := range items {
		innerMap[item] = struct{}{}
	}
	return &Set[T]{
		innerMap: innerMap,
	}
}

func (s *Set[T]) Contains(t T) (found bool) {
	_, found = s.innerMap[t]
	return
}

func (s *Set[T]) Values() []T {
	ret := make([]T, 0, len(s.innerMap))
	for item := range s.innerMap {
		ret = append(ret, item)
	}
	return ret
}
