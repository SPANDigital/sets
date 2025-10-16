package sets

import "iter"

// InsertionOrderedSet is a generic set data structure that preserves insertion order.
// Unlike Set (unordered) and OrderedSet (sorted), InsertionOrderedSet returns values
// in the order they were first added.
// It provides O(1) membership checks but O(n) removal operations.
// Duplicate values are automatically handled - adding the same value multiple times
// has no effect beyond the first addition.
type InsertionOrderedSet[T comparable] struct {
	innerMap map[T]struct{}
	order    []T
}

// NewInsertionOrderedSet creates a new InsertionOrderedSet with the given initial items.
// Items are stored in the order provided. Duplicate items are deduplicated while
// preserving the first occurrence's position.
func NewInsertionOrderedSet[T comparable](items ...T) *InsertionOrderedSet[T] {
	innerMap := make(map[T]struct{})
	order := make([]T, 0, len(items))
	for _, item := range items {
		if _, exists := innerMap[item]; !exists {
			innerMap[item] = struct{}{}
			order = append(order, item)
		}
	}
	return &InsertionOrderedSet[T]{
		innerMap: innerMap,
		order:    order,
	}
}

// Contains checks if the given item exists in the set.
// Returns true if the item is present, false otherwise.
// Time complexity: O(1)
func (s *InsertionOrderedSet[T]) Contains(t T) bool {
	_, found := s.innerMap[t]
	return found
}

// Values returns all elements in the set as a slice in insertion order.
// The order reflects the sequence in which elements were first added.
// Time complexity: O(n)
func (s *InsertionOrderedSet[T]) Values() []T {
	result := make([]T, len(s.order))
	copy(result, s.order)
	return result
}

// Add inserts one or more items into the set.
// Items are added to the end of the order. Adding items that already exist has no effect.
// Time complexity: O(m) where m is the number of items to add
func (s *InsertionOrderedSet[T]) Add(items ...T) {
	for _, item := range items {
		if _, exists := s.innerMap[item]; !exists {
			s.innerMap[item] = struct{}{}
			s.order = append(s.order, item)
		}
	}
}

// Remove deletes one or more items from the set.
// Removing items that don't exist has no effect.
// Time complexity: O(n*m) where n is set size and m is number of items to remove
func (s *InsertionOrderedSet[T]) Remove(items ...T) {
	for _, item := range items {
		if _, exists := s.innerMap[item]; exists {
			delete(s.innerMap, item)
			// Find and remove from order slice
			for i, v := range s.order {
				if v == item {
					s.order = append(s.order[:i], s.order[i+1:]...)
					break
				}
			}
		}
	}
}

// Len returns the number of elements in the set.
// Time complexity: O(1)
func (s *InsertionOrderedSet[T]) Len() int {
	return len(s.order)
}

// IsEmpty returns true if the set contains no elements.
// Time complexity: O(1)
func (s *InsertionOrderedSet[T]) IsEmpty() bool {
	return len(s.order) == 0
}

// Clear removes all elements from the set, making it empty.
// Time complexity: O(1)
func (s *InsertionOrderedSet[T]) Clear() {
	s.innerMap = make(map[T]struct{})
	s.order = make([]T, 0)
}

// Union returns a new set containing all elements from both this set and the other set.
// Elements from this set appear first in insertion order, followed by new elements
// from the other set in their insertion order.
// Time complexity: O(n + m) where n and m are the sizes of the two sets
func (s *InsertionOrderedSet[T]) Union(other *InsertionOrderedSet[T]) *InsertionOrderedSet[T] {
	result := NewInsertionOrderedSet[T]()
	// Add all from this set first (preserves order)
	for _, item := range s.order {
		result.innerMap[item] = struct{}{}
		result.order = append(result.order, item)
	}
	// Add new items from other set
	for _, item := range other.order {
		if _, exists := result.innerMap[item]; !exists {
			result.innerMap[item] = struct{}{}
			result.order = append(result.order, item)
		}
	}
	return result
}

// Intersection returns a new set containing only elements that exist in both sets.
// The insertion order from this set is preserved for common elements.
// Time complexity: O(n) where n is the size of this set
func (s *InsertionOrderedSet[T]) Intersection(other *InsertionOrderedSet[T]) *InsertionOrderedSet[T] {
	result := NewInsertionOrderedSet[T]()
	for _, item := range s.order {
		if other.Contains(item) {
			result.innerMap[item] = struct{}{}
			result.order = append(result.order, item)
		}
	}
	return result
}

// Difference returns a new set containing elements that are in this set but not in the other set.
// The insertion order from this set is preserved.
// Time complexity: O(n) where n is the size of this set
func (s *InsertionOrderedSet[T]) Difference(other *InsertionOrderedSet[T]) *InsertionOrderedSet[T] {
	result := NewInsertionOrderedSet[T]()
	for _, item := range s.order {
		if !other.Contains(item) {
			result.innerMap[item] = struct{}{}
			result.order = append(result.order, item)
		}
	}
	return result
}

// Equal returns true if both sets contain exactly the same elements.
// Order is not considered for equality - only membership matters.
// Time complexity: O(n) where n is the size of this set
func (s *InsertionOrderedSet[T]) Equal(other *InsertionOrderedSet[T]) bool {
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

// All returns an iterator over all elements in the set in insertion order.
// Elements are yielded in the order they were first added to the set.
// The iterator can be used with range loops: for v := range set.All() { ... }
// Time complexity: O(n)
func (s *InsertionOrderedSet[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, item := range s.order {
			if !yield(item) {
				return
			}
		}
	}
}
