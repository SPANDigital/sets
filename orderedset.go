package sets

import (
	"cmp"
	"iter"
	"slices"
)

// OrderedSet is a generic set data structure for ordered types (integers, floats, strings).
// Unlike Set, OrderedSet returns values in sorted order when calling Values().
// It provides O(1) membership checks and efficient set operations.
// Duplicate values are automatically handled - adding the same value multiple times
// has no effect beyond the first addition.
type OrderedSet[T cmp.Ordered] struct {
	innerMap map[T]struct{}
}

// NewOrderedSet creates a new OrderedSet with the given initial items.
// Duplicate items in the input are automatically deduplicated.
func NewOrderedSet[T cmp.Ordered](items ...T) *OrderedSet[T] {
	innerMap := make(map[T]struct{})
	for _, item := range items {
		innerMap[item] = struct{}{}
	}
	return &OrderedSet[T]{
		innerMap: innerMap,
	}
}

// Contains checks if the given item exists in the set.
// Returns true if the item is present, false otherwise.
func (s *OrderedSet[T]) Contains(t T) (found bool) {
	_, found = s.innerMap[t]
	return
}

// Values returns all elements in the set as a sorted slice.
// Elements are sorted in ascending order.
func (s *OrderedSet[T]) Values() []T {
	ret := make([]T, 0, len(s.innerMap))
	for item := range s.innerMap {
		ret = append(ret, item)
	}
	slices.Sort(ret)
	return ret
}

// Add inserts one or more items into the set.
// Adding items that already exist has no effect.
func (s *OrderedSet[T]) Add(items ...T) {
	for _, item := range items {
		s.innerMap[item] = struct{}{}
	}
}

// Remove deletes one or more items from the set.
// Removing items that don't exist has no effect.
func (s *OrderedSet[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s.innerMap, item)
	}
}

// Len returns the number of elements in the set.
func (s *OrderedSet[T]) Len() int {
	return len(s.innerMap)
}

// IsEmpty returns true if the set contains no elements.
func (s *OrderedSet[T]) IsEmpty() bool {
	return len(s.innerMap) == 0
}

// Clear removes all elements from the set, making it empty.
func (s *OrderedSet[T]) Clear() {
	s.innerMap = make(map[T]struct{})
}

// Union returns a new set containing all elements from both this set and the other set.
func (s *OrderedSet[T]) Union(other *OrderedSet[T]) *OrderedSet[T] {
	result := NewOrderedSet[T]()
	for item := range s.innerMap {
		result.innerMap[item] = struct{}{}
	}
	for item := range other.innerMap {
		result.innerMap[item] = struct{}{}
	}
	return result
}

// Intersection returns a new set containing only elements that exist in both sets.
func (s *OrderedSet[T]) Intersection(other *OrderedSet[T]) *OrderedSet[T] {
	result := NewOrderedSet[T]()
	for item := range s.innerMap {
		if other.Contains(item) {
			result.innerMap[item] = struct{}{}
		}
	}
	return result
}

// Difference returns a new set containing elements that are in this set but not in the other set.
func (s *OrderedSet[T]) Difference(other *OrderedSet[T]) *OrderedSet[T] {
	result := NewOrderedSet[T]()
	for item := range s.innerMap {
		if !other.Contains(item) {
			result.innerMap[item] = struct{}{}
		}
	}
	return result
}

// Equal returns true if both sets contain exactly the same elements.
func (s *OrderedSet[T]) Equal(other *OrderedSet[T]) bool {
	if len(s.innerMap) != len(other.innerMap) {
		return false
	}
	for item := range s.innerMap {
		if !other.Contains(item) {
			return false
		}
	}
	return true
}

// All returns an iterator over all elements in the set in sorted order.
// Elements are yielded in ascending order.
// The iterator can be used with range loops: for v := range set.All() { ... }
func (s *OrderedSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		values := s.Values()
		for _, item := range values {
			if !yield(item) {
				return
			}
		}
	}
}
