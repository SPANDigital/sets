package sets

import (
	"cmp"
	"reflect"
	"testing"
)

func TestNewOrderedSet(t *testing.T) {
	type args[T cmp.Ordered] struct {
		items []T
	}
	type testCase[T cmp.Ordered] struct {
		name string
		args args[T]
		want map[T]struct{}
	}
	tests := []testCase[int]{
		{
			name: "empty",
			args: struct{ items []int }{items: nil},
			want: map[int]struct{}{},
		},
		{
			name: "one",
			args: struct{ items []int }{items: []int{5}},
			want: map[int]struct{}{5: struct{}{}},
		},
		{
			name: "two",
			args: struct{ items []int }{items: []int{10, 5}},
			want: map[int]struct{}{10: struct{}{}, 5: struct{}{}},
		},
		{
			name: "three",
			args: struct{ items []int }{items: []int{30, 10, 20}},
			want: map[int]struct{}{30: struct{}{}, 10: struct{}{}, 20: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewOrderedSet(tt.args.items...); !reflect.DeepEqual(got.innerMap, tt.want) {
				t.Errorf("NewOrderedSet() = %v, wantAfterSort %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSet_Contains(t *testing.T) {
	type args[T cmp.Ordered] struct {
		t T
	}
	type testCase[T cmp.Ordered] struct {
		name      string
		s         *OrderedSet[T]
		args      args[T]
		wantFound bool
	}
	tests := []testCase[int]{
		{
			name: "empty-contains",
			s:    NewOrderedSet[int](),
			args: args[int]{
				t: 99,
			},
			wantFound: false,
		},
		{
			name: "three-does-contain",
			s:    NewOrderedSet(10, 20, 30),
			args: args[int]{
				t: 20,
			},
			wantFound: true,
		},
		{
			name: "three-does-not-contain",
			s:    NewOrderedSet(10, 20, 30),
			args: args[int]{
				t: 40,
			},
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFound := tt.s.Contains(tt.args.t); gotFound != tt.wantFound {
				t.Errorf("Contains() = %v, wantAfterSort %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestOrderedSet_Values(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		s    *OrderedSet[T]
		want []T
	}
	tests := []testCase[string]{
		{
			name: "empty",
			s:    NewOrderedSet[string](),
			want: []string{},
		},
		{
			name: "one",
			s:    NewOrderedSet("alpha"),
			want: []string{"alpha"},
		},
		{
			name: "two",
			s:    NewOrderedSet("alpha", "beta"),
			want: []string{"alpha", "beta"},
		},
		{
			name: "three",
			s:    NewOrderedSet("alpha", "beta", "gamma"),
			want: []string{"alpha", "beta", "gamma"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, wantAfterSort %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSet_Add(t *testing.T) {
	s := NewOrderedSet(1, 2)
	s.Add(3)
	if !s.Contains(3) {
		t.Error("Add failed to add element")
	}
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Test adding duplicate
	s.Add(1)
	if s.Len() != 3 {
		t.Errorf("Adding duplicate changed length, got %d", s.Len())
	}
}

func TestOrderedSet_Remove(t *testing.T) {
	s := NewOrderedSet(1, 2, 3)
	s.Remove(2)
	if s.Contains(2) {
		t.Error("Remove failed to remove element")
	}
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	// Test removing non-existent
	s.Remove(4)
	if s.Len() != 2 {
		t.Errorf("Removing non-existent changed length, got %d", s.Len())
	}
}

func TestOrderedSet_Len(t *testing.T) {
	tests := []struct {
		name string
		s    *OrderedSet[int]
		want int
	}{
		{"empty", NewOrderedSet[int](), 0},
		{"one", NewOrderedSet(1), 1},
		{"three", NewOrderedSet(1, 2, 3), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrderedSet_IsEmpty(t *testing.T) {
	s := NewOrderedSet[int]()
	if !s.IsEmpty() {
		t.Error("New set should be empty")
	}

	s.Add(1)
	if s.IsEmpty() {
		t.Error("Set with elements should not be empty")
	}
}

func TestOrderedSet_Clear(t *testing.T) {
	s := NewOrderedSet(1, 2, 3)
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Clear failed to empty set")
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", s.Len())
	}
}

func TestOrderedSet_Union(t *testing.T) {
	s1 := NewOrderedSet(1, 2, 3)
	s2 := NewOrderedSet(3, 4, 5)
	result := s1.Union(s2)

	expected := []int{1, 2, 3, 4, 5}
	if result.Len() != 5 {
		t.Errorf("Union length = %d, want 5", result.Len())
	}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Union values = %v, want %v", got, expected)
	}
}

func TestOrderedSet_Intersection(t *testing.T) {
	s1 := NewOrderedSet(1, 2, 3, 4)
	s2 := NewOrderedSet(3, 4, 5, 6)
	result := s1.Intersection(s2)

	expected := []int{3, 4}
	if result.Len() != 2 {
		t.Errorf("Intersection length = %d, want 2", result.Len())
	}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Intersection values = %v, want %v", got, expected)
	}
}

func TestOrderedSet_Difference(t *testing.T) {
	s1 := NewOrderedSet(1, 2, 3, 4)
	s2 := NewOrderedSet(3, 4, 5, 6)
	result := s1.Difference(s2)

	expected := []int{1, 2}
	if result.Len() != 2 {
		t.Errorf("Difference length = %d, want 2", result.Len())
	}
	got := result.Values()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Difference values = %v, want %v", got, expected)
	}
}

func TestOrderedSet_Equal(t *testing.T) {
	s1 := NewOrderedSet(1, 2, 3)
	s2 := NewOrderedSet(3, 2, 1)
	s3 := NewOrderedSet(1, 2, 4)

	if !s1.Equal(s2) {
		t.Error("Equal sets not detected as equal")
	}
	if s1.Equal(s3) {
		t.Error("Unequal sets detected as equal")
	}
}

func TestOrderedSet_All(t *testing.T) {
	s := NewOrderedSet(30, 10, 50, 20, 40)

	// Test iteration collects all values in sorted order
	collected := make([]int, 0)
	for v := range s.All() {
		collected = append(collected, v)
	}

	expected := []int{10, 20, 30, 40, 50}
	if !reflect.DeepEqual(collected, expected) {
		t.Errorf("All() = %v, want %v (sorted order)", collected, expected)
	}
}

func TestOrderedSet_All_EarlyTermination(t *testing.T) {
	s := NewOrderedSet(5, 4, 3, 2, 1)

	// Test early termination with break
	collected := make([]int, 0)
	for v := range s.All() {
		collected = append(collected, v)
		if len(collected) == 3 {
			break
		}
	}

	// Should get first 3 in sorted order: 1, 2, 3
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(collected, expected) {
		t.Errorf("Early termination collected %v, want %v", collected, expected)
	}
}

func TestOrderedSet_All_Empty(t *testing.T) {
	s := NewOrderedSet[int]()

	count := 0
	for range s.All() {
		count++
	}
	if count != 0 {
		t.Errorf("Empty set All() yielded %d elements, want 0", count)
	}
}
