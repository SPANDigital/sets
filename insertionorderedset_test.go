package sets

import (
	"reflect"
	"testing"
)

func TestNewInsertionOrderedSet(t *testing.T) {
	type args[T comparable] struct {
		items []T
	}
	type testCase[T comparable] struct {
		name      string
		args      args[T]
		wantOrder []T
		wantLen   int
	}
	tests := []testCase[string]{
		{
			name:      "empty",
			args:      struct{ items []string }{items: nil},
			wantOrder: []string{},
			wantLen:   0,
		},
		{
			name:      "one",
			args:      struct{ items []string }{items: []string{"one"}},
			wantOrder: []string{"one"},
			wantLen:   1,
		},
		{
			name:      "three",
			args:      struct{ items []string }{items: []string{"one", "two", "three"}},
			wantOrder: []string{"one", "two", "three"},
			wantLen:   3,
		},
		{
			name:      "with duplicates",
			args:      struct{ items []string }{items: []string{"one", "two", "one", "three"}},
			wantOrder: []string{"one", "two", "three"},
			wantLen:   3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewInsertionOrderedSet(tt.args.items...)
			if got.Len() != tt.wantLen {
				t.Errorf("NewInsertionOrderedSet() len = %v, want %v", got.Len(), tt.wantLen)
			}
			if !reflect.DeepEqual(got.Values(), tt.wantOrder) {
				t.Errorf("NewInsertionOrderedSet() order = %v, want %v", got.Values(), tt.wantOrder)
			}
		})
	}
}

func TestInsertionOrderedSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		t T
	}
	type testCase[T comparable] struct {
		name      string
		s         *InsertionOrderedSet[T]
		args      args[T]
		wantFound bool
	}
	tests := []testCase[string]{
		{
			name: "empty-contains",
			s:    NewInsertionOrderedSet[string](),
			args: args[string]{
				t: "x",
			},
			wantFound: false,
		},
		{
			name: "three-does-contain",
			s:    NewInsertionOrderedSet("one", "two", "three"),
			args: args[string]{
				t: "two",
			},
			wantFound: true,
		},
		{
			name: "three-does-not-contain",
			s:    NewInsertionOrderedSet("one", "two", "three"),
			args: args[string]{
				t: "four",
			},
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFound := tt.s.Contains(tt.args.t); gotFound != tt.wantFound {
				t.Errorf("Contains() = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestInsertionOrderedSet_Values(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		s    *InsertionOrderedSet[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "empty",
			s:    NewInsertionOrderedSet[string](),
			want: []string{},
		},
		{
			name: "one",
			s:    NewInsertionOrderedSet("alpha"),
			want: []string{"alpha"},
		},
		{
			name: "preserves insertion order",
			s:    NewInsertionOrderedSet("charlie", "alpha", "beta"),
			want: []string{"charlie", "alpha", "beta"},
		},
		{
			name: "reverse order",
			s:    NewInsertionOrderedSet("three", "two", "one"),
			want: []string{"three", "two", "one"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionOrderedSet_Add(t *testing.T) {
	s := NewInsertionOrderedSet(1, 2)
	s.Add(3)
	if !s.Contains(3) {
		t.Error("Add failed to add element")
	}
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Verify order is preserved
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(s.Values(), expected) {
		t.Errorf("Order not preserved after Add, got %v, want %v", s.Values(), expected)
	}

	// Test adding duplicate
	s.Add(2)
	if s.Len() != 3 {
		t.Errorf("Adding duplicate changed length, got %d", s.Len())
	}
	if !reflect.DeepEqual(s.Values(), expected) {
		t.Errorf("Order changed after adding duplicate, got %v, want %v", s.Values(), expected)
	}
}

func TestInsertionOrderedSet_Remove(t *testing.T) {
	s := NewInsertionOrderedSet(1, 2, 3, 4)
	s.Remove(2)
	if s.Contains(2) {
		t.Error("Remove failed to remove element")
	}
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Verify order is preserved after removal
	expected := []int{1, 3, 4}
	if !reflect.DeepEqual(s.Values(), expected) {
		t.Errorf("Order not preserved after Remove, got %v, want %v", s.Values(), expected)
	}

	// Test removing non-existent
	s.Remove(10)
	if s.Len() != 3 {
		t.Errorf("Removing non-existent changed length, got %d", s.Len())
	}
}

func TestInsertionOrderedSet_Len(t *testing.T) {
	tests := []struct {
		name string
		s    *InsertionOrderedSet[int]
		want int
	}{
		{"empty", NewInsertionOrderedSet[int](), 0},
		{"one", NewInsertionOrderedSet(1), 1},
		{"three", NewInsertionOrderedSet(1, 2, 3), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertionOrderedSet_IsEmpty(t *testing.T) {
	s := NewInsertionOrderedSet[int]()
	if !s.IsEmpty() {
		t.Error("New set should be empty")
	}

	s.Add(1)
	if s.IsEmpty() {
		t.Error("Set with elements should not be empty")
	}
}

func TestInsertionOrderedSet_Clear(t *testing.T) {
	s := NewInsertionOrderedSet(1, 2, 3)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Clear failed to empty set")
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", s.Len())
	}
}

func TestInsertionOrderedSet_Union(t *testing.T) {
	s1 := NewInsertionOrderedSet(1, 2, 3)
	s2 := NewInsertionOrderedSet(3, 4, 5)
	result := s1.Union(s2)

	// Should have all elements
	if result.Len() != 5 {
		t.Errorf("Union length = %d, want 5", result.Len())
	}

	// Should preserve s1's order first, then add new elements from s2 in order
	expected := []int{1, 2, 3, 4, 5}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Union values = %v, want %v", got, expected)
	}
}

func TestInsertionOrderedSet_Union_OrderPreservation(t *testing.T) {
	s1 := NewInsertionOrderedSet("c", "a", "b")
	s2 := NewInsertionOrderedSet("b", "d", "e")
	result := s1.Union(s2)

	// s1's order first, then new elements from s2
	expected := []string{"c", "a", "b", "d", "e"}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Union order = %v, want %v", got, expected)
	}
}

func TestInsertionOrderedSet_Intersection(t *testing.T) {
	s1 := NewInsertionOrderedSet(1, 2, 3, 4)
	s2 := NewInsertionOrderedSet(3, 4, 5, 6)
	result := s1.Intersection(s2)

	if result.Len() != 2 {
		t.Errorf("Intersection length = %d, want 2", result.Len())
	}

	// Should preserve s1's order for common elements
	expected := []int{3, 4}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Intersection values = %v, want %v", got, expected)
	}
}

func TestInsertionOrderedSet_Intersection_OrderPreservation(t *testing.T) {
	s1 := NewInsertionOrderedSet("d", "c", "b", "a")
	s2 := NewInsertionOrderedSet("a", "b", "e", "f")
	result := s1.Intersection(s2)

	// Should preserve s1's order
	expected := []string{"b", "a"}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Intersection order = %v, want %v", got, expected)
	}
}

func TestInsertionOrderedSet_Difference(t *testing.T) {
	s1 := NewInsertionOrderedSet(1, 2, 3, 4)
	s2 := NewInsertionOrderedSet(3, 4, 5, 6)
	result := s1.Difference(s2)

	if result.Len() != 2 {
		t.Errorf("Difference length = %d, want 2", result.Len())
	}

	// Should preserve s1's order
	expected := []int{1, 2}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Difference values = %v, want %v", got, expected)
	}
}

func TestInsertionOrderedSet_Equal(t *testing.T) {
	s1 := NewInsertionOrderedSet(1, 2, 3)
	s2 := NewInsertionOrderedSet(3, 2, 1) // Different order
	s3 := NewInsertionOrderedSet(1, 2, 4) // Different elements

	// Equal should ignore order
	if !s1.Equal(s2) {
		t.Error("Equal sets not detected as equal (order should not matter)")
	}
	if s1.Equal(s3) {
		t.Error("Unequal sets detected as equal")
	}
}

func TestInsertionOrderedSet_MultipleOperations(t *testing.T) {
	s := NewInsertionOrderedSet[string]()
	s.Add("first")
	s.Add("second")
	s.Add("third")
	s.Remove("second")
	s.Add("fourth")

	expected := []string{"first", "third", "fourth"}
	got := s.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("After multiple operations, order = %v, want %v", got, expected)
	}
}

func TestInsertionOrderedSet_All(t *testing.T) {
	s := NewInsertionOrderedSet("charlie", "alpha", "delta", "beta")

	// Test iteration preserves insertion order
	collected := make([]string, 0)
	for v := range s.All() {
		collected = append(collected, v)
	}

	expected := []string{"charlie", "alpha", "delta", "beta"}
	if !reflect.DeepEqual(collected, expected) {
		t.Errorf("All() = %v, want %v (insertion order)", collected, expected)
	}
}

func TestInsertionOrderedSet_All_EarlyTermination(t *testing.T) {
	s := NewInsertionOrderedSet("first", "second", "third", "fourth", "fifth")

	// Test early termination with break
	collected := make([]string, 0)
	for v := range s.All() {
		collected = append(collected, v)
		if len(collected) == 3 {
			break
		}
	}

	// Should get first 3 in insertion order
	expected := []string{"first", "second", "third"}
	if !reflect.DeepEqual(collected, expected) {
		t.Errorf("Early termination collected %v, want %v", collected, expected)
	}
}

func TestInsertionOrderedSet_All_Empty(t *testing.T) {
	s := NewInsertionOrderedSet[string]()

	count := 0
	for range s.All() {
		count++
	}
	if count != 0 {
		t.Errorf("Empty set All() yielded %d elements, want 0", count)
	}
}

func TestInsertionOrderedSet_All_AfterRemoval(t *testing.T) {
	s := NewInsertionOrderedSet(1, 2, 3, 4, 5)
	s.Remove(3)

	// Test that iteration order is maintained after removal
	collected := make([]int, 0)
	for v := range s.All() {
		collected = append(collected, v)
	}

	expected := []int{1, 2, 4, 5}
	if !reflect.DeepEqual(collected, expected) {
		t.Errorf("All() after removal = %v, want %v", collected, expected)
	}
}
