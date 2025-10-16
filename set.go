package sets

import "iter"

// Set is a generic unordered set data structure for any comparable type.
// It provides O(1) membership checks and efficient set operations.
// Duplicate values are automatically handled - adding the same value multiple times
// has no effect beyond the first addition.
type Set[T comparable] struct {
	innerMap map[T]struct{}
}

// NewSet creates a new Set with the given initial items.
// Duplicate items in the input are automatically deduplicated.
func NewSet[T comparable](items ...T) *Set[T] {
	innerMap := make(map[T]struct{})
	for _, item := range items {
		innerMap[item] = struct{}{}
	}
	return &Set[T]{
		innerMap: innerMap,
	}
}

// Contains checks if the given item exists in the set.
// Returns true if the item is present, false otherwise.
func (s *Set[T]) Contains(t T) (found bool) {
	_, found = s.innerMap[t]
	return
}

// Values returns all elements in the set as a slice.
// The order of elements is not guaranteed and may vary between calls.
func (s *Set[T]) Values() []T {
	ret := make([]T, 0, len(s.innerMap))
	for item := range s.innerMap {
		ret = append(ret, item)
	}
	return ret
}

// Add inserts one or more items into the set.
// Adding items that already exist has no effect.
func (s *Set[T]) Add(items ...T) {
	for _, item := range items {
		s.innerMap[item] = struct{}{}
	}
}

// Remove deletes one or more items from the set.
// Removing items that don't exist has no effect.
func (s *Set[T]) Remove(items ...T) {
	for _, item := range items {
		delete(s.innerMap, item)
	}
}

// Len returns the number of elements in the set.
func (s *Set[T]) Len() int {
	return len(s.innerMap)
}

// IsEmpty returns true if the set contains no elements.
func (s *Set[T]) IsEmpty() bool {
	return len(s.innerMap) == 0
}

// Clear removes all elements from the set, making it empty.
func (s *Set[T]) Clear() {
	s.innerMap = make(map[T]struct{})
}

// Union returns a new set containing all elements from both this set and the other set.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.innerMap {
		result.innerMap[item] = struct{}{}
	}
	for item := range other.innerMap {
		result.innerMap[item] = struct{}{}
	}
	return result
}

// Intersection returns a new set containing only elements that exist in both sets.
func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.innerMap {
		if other.Contains(item) {
			result.innerMap[item] = struct{}{}
		}
	}
	return result
}

// Difference returns a new set containing elements that are in this set but not in the other set.
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.innerMap {
		if !other.Contains(item) {
			result.innerMap[item] = struct{}{}
		}
	}
	return result
}

// Equal returns true if both sets contain exactly the same elements.
func (s *Set[T]) Equal(other *Set[T]) bool {
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

// All returns an iterator over all elements in the set.
// The iteration order is not guaranteed and may vary between calls.
// The iterator can be used with range loops: for v := range set.All() { ... }
func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for item := range s.innerMap {
			if !yield(item) {
				return
			}
		}
	}
}
