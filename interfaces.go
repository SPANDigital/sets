package sets

import "iter"

// Containser checks if an element exists in a collection.
type Containser[T any] interface {
	Contains(T) bool
}

// Valuer returns all elements from a collection.
type Valuer[T any] interface {
	Values() []T
}

// Adder inserts elements into a collection.
type Adder[T any] interface {
	Add(...T)
}

// Remover deletes elements from a collection.
type Remover[T any] interface {
	Remove(...T)
}

// Lener returns the number of elements in a collection.
type Lener interface {
	Len() int
}

// EmptyChecker determines if a collection is empty.
type EmptyChecker interface {
	IsEmpty() bool
}

// Clearer removes all elements from a collection.
type Clearer interface {
	Clear()
}

// Unioner combines two collections into a new one.
type Unioner[T any, S any] interface {
	Union(S) S
}

// Intersectioner returns common elements between two collections.
type Intersectioner[T any, S any] interface {
	Intersection(S) S
}

// Differencer returns elements in one collection but not another.
type Differencer[T any, S any] interface {
	Difference(S) S
}

// Equaler compares two collections for equality.
type Equaler[S any] interface {
	Equal(S) bool
}

// Iterable provides an iterator over elements in a collection.
type Iterable[T any] interface {
	All() iter.Seq[T]
}
