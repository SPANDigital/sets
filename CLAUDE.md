# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Go library providing generic set data structures with three implementations:
- `Set[T comparable]`: Unordered set for any comparable type
- `OrderedSet[T cmp.Ordered]`: Ordered set that returns values in sorted order
- `InsertionOrderedSet[T comparable]`: Set that preserves insertion order

Module path: `github.com/spandigital/sets`

## Common Commands

### Testing
```bash
# Run all tests
go test -v ./...

# Run tests for a specific file
go test -v -run TestFunctionName

# Run with coverage
go test -v -cover ./...
```

### Building
```bash
# Build the package
go build -v ./...

# Format code
go fmt ./...
```

## Architecture

### Set Implementation (`set.go`)
- Generic type: `Set[T comparable]` accepts any comparable type
- Backed by `map[T]struct{}` for O(1) membership checks
- `NewSet[T](items ...T)`: Constructor with variadic initialization
- `Contains(t T) bool`: Membership test - O(1)
- `Values() []T`: Returns all elements (order not guaranteed)
- `Add(items ...T)`: Add elements - O(1) per element
- `Remove(items ...T)`: Remove elements - O(1) per element
- Set operations: `Union()`, `Intersection()`, `Difference()`, `Equal()`

### OrderedSet Implementation (`orderedset.go`)
- Generic type: `OrderedSet[T cmp.Ordered]` for sortable types (int, string, etc.)
- Same internal `map[T]struct{}` backing as Set
- `Values()` returns elements in sorted order using `slices.Sort` - O(n log n)
- All methods same as Set, but maintains sorted order in output
- Performance: Same as Set except Values() which sorts

### InsertionOrderedSet Implementation (`insertionorderedset.go`)
- Generic type: `InsertionOrderedSet[T comparable]` preserves insertion order
- Dual data structure: `map[T]struct{}` + `[]T` slice
- `Contains(t T) bool`: O(1) via map lookup
- `Values() []T`: O(n) returns copy in insertion order
- `Add(items ...T)`: O(1) amortized - appends to slice
- `Remove(items ...T)`: O(n) - must find and remove from slice
- Set operations preserve this set's order for Union/Intersection/Difference
- `Equal()` ignores order - only checks membership

### Interfaces (`interfaces.go`)
Single-method interfaces for composability:
- `Containser[T]`, `Valuer[T]`, `Adder[T]`, `Remover[T]`
- `Lener`, `EmptyChecker`, `Clearer`
- `Unioner[T, S]`, `Intersectioner[T, S]`, `Differencer[T, S]`, `Equaler[S]`
- `Iterable[T]` - provides `All() iter.Seq[T]` for standard Go iteration
- All three set types implement all these interfaces

### Iterator Support (`iter` package)
All set types support Go 1.23+ range-over-function iteration:
```go
s := sets.NewOrderedSet(3, 1, 2)
for v := range s.All() {
    fmt.Println(v)  // prints: 1, 2, 3
}
```
- `Set.All()` - iterates in random order
- `OrderedSet.All()` - iterates in sorted order
- `InsertionOrderedSet.All()` - iterates in insertion order
- All support early termination via `break`

## Code Conventions

- Use Go generics for all set implementations
- Method names use PascalCase
- Sets use `map[T]struct{}` pattern (zero memory per element)
- Constructors accept variadic arguments for convenience
- Tests follow table-driven patterns
- No external dependencies beyond Go standard library
- Go version: 1.23 (required for `iter` package)
- All set types implement `iter.Seq[T]` via `All()` method for standard Go iteration
