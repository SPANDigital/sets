package sets

import (
	"cmp"
	"reflect"
	"slices"
	"testing"
)

func TestNewSet(t *testing.T) {
	type args[T comparable] struct {
		items []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want map[T]struct{}
	}
	tests := []testCase[string]{
		{
			name: "empty",
			args: struct{ items []string }{items: nil},
			want: map[string]struct{}{},
		},
		{
			name: "one",
			args: struct{ items []string }{items: []string{"two"}},
			want: map[string]struct{}{"two": struct{}{}},
		},
		{
			name: "two",
			args: struct{ items []string }{items: []string{"one", "two"}},
			want: map[string]struct{}{"one": struct{}{}, "two": struct{}{}},
		},
		{
			name: "three",
			args: struct{ items []string }{items: []string{"one", "two", "three"}},
			want: map[string]struct{}{"one": struct{}{}, "two": struct{}{}, "three": struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(tt.args.items...); !reflect.DeepEqual(got.innerMap, tt.want) {
				t.Errorf("NewSet() = %v, wantAfterSort %v", got, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args[T comparable] struct {
		t T
	}
	type testCase[T comparable] struct {
		name      string
		s         *Set[T]
		args      args[T]
		wantFound bool
	}
	tests := []testCase[string]{
		{
			name: "empty-contains",
			s:    NewSet[string](),
			args: args[string]{
				t: "x",
			},
			wantFound: false,
		},
		{
			name: "three-does-contain",
			s:    NewSet("one", "two", "three"),
			args: args[string]{
				t: "one",
			},
			wantFound: true,
		},
		{
			name: "three-does-not-contain",
			s:    NewSet("one", "two", "three"),
			args: args[string]{
				t: "four",
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

func sort[T cmp.Ordered](s []T) []T {
	slices.Sort(s)
	return s
}

func TestSet_Values(t *testing.T) {
	type testCase[T comparable] struct {
		name          string
		s             *Set[T]
		wantAfterSort []T
	}
	tests := []testCase[string]{
		{
			name:          "empty",
			s:             NewSet[string](),
			wantAfterSort: []string{},
		},
		{
			name:          "one",
			s:             NewSet("one"),
			wantAfterSort: []string{"one"},
		},
		{
			name:          "two",
			s:             NewSet("one", "two"),
			wantAfterSort: []string{"one", "two"},
		},
		{
			name:          "three",
			s:             NewSet("one", "two", "three"),
			wantAfterSort: []string{"one", "three", "two"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Values()
			sortedGot := make([]string, len(got))
			copy(sortedGot, got)
			slices.Sort(sortedGot)
			if !reflect.DeepEqual(sortedGot, tt.wantAfterSort) {
				t.Errorf("Values() sorted = %v, want %v", sortedGot, tt.wantAfterSort)
			}
		})
	}
}

func TestSet_Add(t *testing.T) {
	s := NewSet("one", "two")
	s.Add("three")
	if !s.Contains("three") {
		t.Error("Add failed to add element")
	}
	if s.Len() != 3 {
		t.Errorf("Expected length 3, got %d", s.Len())
	}

	// Test adding duplicate
	s.Add("one")
	if s.Len() != 3 {
		t.Errorf("Adding duplicate changed length, got %d", s.Len())
	}
}

func TestSet_Remove(t *testing.T) {
	s := NewSet("one", "two", "three")
	s.Remove("two")
	if s.Contains("two") {
		t.Error("Remove failed to remove element")
	}
	if s.Len() != 2 {
		t.Errorf("Expected length 2, got %d", s.Len())
	}

	// Test removing non-existent
	s.Remove("four")
	if s.Len() != 2 {
		t.Errorf("Removing non-existent changed length, got %d", s.Len())
	}
}

func TestSet_Len(t *testing.T) {
	tests := []struct {
		name string
		s    *Set[string]
		want int
	}{
		{"empty", NewSet[string](), 0},
		{"one", NewSet("a"), 1},
		{"three", NewSet("a", "b", "c"), 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_IsEmpty(t *testing.T) {
	s := NewSet[string]()
	if !s.IsEmpty() {
		t.Error("New set should be empty")
	}

	s.Add("one")
	if s.IsEmpty() {
		t.Error("Set with elements should not be empty")
	}
}

func TestSet_Clear(t *testing.T) {
	s := NewSet("one", "two", "three")
	s.Clear()
	if !s.IsEmpty() {
		t.Error("Clear failed to empty set")
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", s.Len())
	}
}

func TestSet_Union(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(3, 4, 5)
	result := s1.Union(s2)

	expected := []int{1, 2, 3, 4, 5}
	if result.Len() != 5 {
		t.Errorf("Union length = %d, want 5", result.Len())
	}
	for _, v := range expected {
		if !result.Contains(v) {
			t.Errorf("Union missing element %d", v)
		}
	}
}

func TestSet_Intersection(t *testing.T) {
	s1 := NewSet(1, 2, 3, 4)
	s2 := NewSet(3, 4, 5, 6)
	result := s1.Intersection(s2)

	if result.Len() != 2 {
		t.Errorf("Intersection length = %d, want 2", result.Len())
	}
	if !result.Contains(3) || !result.Contains(4) {
		t.Error("Intersection missing expected elements")
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := NewSet(1, 2, 3, 4)
	s2 := NewSet(3, 4, 5, 6)
	result := s1.Difference(s2)

	if result.Len() != 2 {
		t.Errorf("Difference length = %d, want 2", result.Len())
	}
	if !result.Contains(1) || !result.Contains(2) {
		t.Error("Difference missing expected elements")
	}
	if result.Contains(3) || result.Contains(4) {
		t.Error("Difference contains unexpected elements")
	}
}

func TestSet_Equal(t *testing.T) {
	s1 := NewSet(1, 2, 3)
	s2 := NewSet(3, 2, 1)
	s3 := NewSet(1, 2, 4)

	if !s1.Equal(s2) {
		t.Error("Equal sets not detected as equal")
	}
	if s1.Equal(s3) {
		t.Error("Unequal sets detected as equal")
	}
}

func TestSet_All(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5)

	// Test iteration collects all values
	collected := make([]int, 0)
	for v := range s.All() {
		collected = append(collected, v)
	}
	if len(collected) != 5 {
		t.Errorf("All() iterated %d elements, want 5", len(collected))
	}

	// Verify all values were seen
	for _, expected := range []int{1, 2, 3, 4, 5} {
		found := false
		for _, v := range collected {
			if v == expected {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("All() did not yield %d", expected)
		}
	}
}

func TestSet_All_EarlyTermination(t *testing.T) {
	s := NewSet(1, 2, 3, 4, 5)

	// Test early termination with break
	count := 0
	for range s.All() {
		count++
		if count == 3 {
			break
		}
	}
	if count != 3 {
		t.Errorf("Early termination failed, got %d iterations, want 3", count)
	}
}

func TestSet_All_Empty(t *testing.T) {
	s := NewSet[int]()

	count := 0
	for range s.All() {
		count++
	}
	if count != 0 {
		t.Errorf("Empty set All() yielded %d elements, want 0", count)
	}
}
