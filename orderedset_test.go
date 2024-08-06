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
			if got := NewOrderedSet(tt.args.items...); !reflect.DeepEqual(got.innerMap, tt.want) {
				t.Errorf("NewOrderedSet() = %v, wantAfterSort %v", got, tt.want)
			}
		})
	}
}

func TestOrderdSet_Contains(t *testing.T) {
	type args[T cmp.Ordered] struct {
		t T
	}
	type testCase[T cmp.Ordered] struct {
		name      string
		s         *OrderdSet[T]
		args      args[T]
		wantFound bool
	}
	tests := []testCase[string]{
		{
			name: "empty-contains",
			s:    NewOrderedSet[string](),
			args: args[string]{
				t: "x",
			},
			wantFound: false,
		},
		{
			name: "three-does-contain",
			s:    NewOrderedSet("one", "two", "three"),
			args: args[string]{
				t: "one",
			},
			wantFound: true,
		},
		{
			name: "three-does-not-contain",
			s:    NewOrderedSet("one", "two", "three"),
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

func TestOrderdSet_Values(t *testing.T) {
	type testCase[T cmp.Ordered] struct {
		name string
		s    *OrderdSet[T]
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
