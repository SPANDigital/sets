# Go Sets 🎯

A fast, type-safe, and easy-to-use set library for Go with three flavors to match your needs.

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.23-blue.svg)](https://go.dev/)

## Features

✨ **Three set types** - Choose the behavior you need
🚀 **Fast lookups** - O(1) membership checks
🔒 **Type-safe** - Powered by Go generics
🔄 **Iterator support** - Works with Go 1.23+ range loops
📦 **Zero dependencies** - Just the Go standard library

## Installation

```bash
go get github.com/spandigital/sets
```

**Requires Go 1.23 or later** for iterator support.

## Quick Start

```go
import "github.com/spandigital/sets"

// Create a set
fruits := sets.NewSet("apple", "banana", "cherry")

// Check membership
fruits.Contains("apple")  // true

// Add and remove
fruits.Add("date")
fruits.Remove("banana")

// Iterate
for fruit := range fruits.All() {
    fmt.Println(fruit)
}
```

## Three Types of Sets

### 1. Set - Fast and Unordered

Perfect when you just need fast lookups and don't care about order.

```go
tags := sets.NewSet("golang", "databases", "web")

tags.Add("api")
tags.Contains("golang")  // true
tags.Len()              // 4

// Order is random
for tag := range tags.All() {
    fmt.Println(tag)  // unpredictable order
}
```

**Use when:** Speed is priority and order doesn't matter.

### 2. OrderedSet - Sorted Values

Returns values in sorted order. Great for consistent output.

```go
scores := sets.NewOrderedSet(95, 87, 92, 88)

for score := range scores.All() {
    fmt.Println(score)  // 87, 88, 92, 95
}

scores.Values()  // [87, 88, 92, 95]
```

**Use when:** You need sorted output (leaderboards, alphabetical lists, etc.)

### 3. InsertionOrderedSet - Preserves Order

Remembers the order you added items. Perfect for maintaining sequences.

```go
steps := sets.NewInsertionOrderedSet("login", "verify", "dashboard")
steps.Add("logout")

for step := range steps.All() {
    fmt.Println(step)  // login, verify, dashboard, logout
}
```

**Use when:** Order of insertion matters (workflows, history, etc.)

## Common Operations

All set types support the same operations:

```go
s := sets.NewSet(1, 2, 3)

// Basic operations
s.Add(4, 5, 6)           // Add multiple items
s.Remove(2)              // Remove item
s.Contains(3)            // true
s.Len()                  // 5
s.IsEmpty()              // false
s.Clear()                // Remove all

// Set operations
s1 := sets.NewSet(1, 2, 3)
s2 := sets.NewSet(2, 3, 4)

s1.Union(s2)             // {1, 2, 3, 4}
s1.Intersection(s2)      // {2, 3}
s1.Difference(s2)        // {1}
s1.Equal(s2)             // false
```

## Iteration

All sets support Go 1.23+ range-over-function iteration:

```go
colors := sets.NewOrderedSet("red", "green", "blue")

// Standard range loop
for color := range colors.All() {
    fmt.Println(color)
}

// Early termination works
for color := range colors.All() {
    if color == "green" {
        break  // stops iteration
    }
}

// Or get a slice
allColors := colors.Values()  // []string{"blue", "green", "red"}
```

## Performance Characteristics

| Operation | Set | OrderedSet | InsertionOrderedSet |
|-----------|-----|------------|---------------------|
| Contains  | O(1) | O(1) | O(1) |
| Add       | O(1) | O(1) | O(1) amortized |
| Remove    | O(1) | O(1) | O(n) |
| Values    | O(n) | O(n log n) | O(n) |
| Iteration | O(n) | O(n log n) | O(n) |

## Examples

### Deduplicating a Slice

```go
items := []string{"a", "b", "a", "c", "b"}
unique := sets.NewSet(items...)
fmt.Println(unique.Values())  // [a, b, c] (random order)
```

### Finding Common Elements

```go
alice := sets.NewSet("golang", "python", "rust")
bob := sets.NewSet("python", "javascript", "rust")

common := alice.Intersection(bob)
// common contains: python, rust
```

### Maintaining Order

```go
history := sets.NewInsertionOrderedSet[string]()
history.Add("home")
history.Add("search")
history.Add("profile")
history.Add("search")  // duplicate, ignored

// history.All() yields: home, search, profile
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) file for details.

Copyright (c) 2024 Span Digital

