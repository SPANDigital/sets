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
			wantAfterSort: []string{"one", "two", "three"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Values(); !reflect.DeepEqual(got, tt.wantAfterSort) {
				t.Errorf("Values() = %v, wantAfterSort %v", sort(got), sort(tt.wantAfterSort))
			}
		})
	}
}
