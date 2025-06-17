# Copilot Instructions for `span/sets`

## Codebase Structure

- `set.go` and `orderedset.go`: Implement generic set data structures.
  - `Set[T comparable]`: Unordered set using a map for storage.
  - `OrderdSet[T cmp.Ordered]`: Set for ordered types, provides sorted values.
- `set_test.go` and `orderedset_test.go`: Unit tests for the above types.

## Conventions

- Use Go generics for type safety and reusability.
- Sets are backed by `map[T]struct{}` for efficient membership checks.
- Constructors accept variadic arguments for initial values.
- Method names use PascalCase (e.g., `Contains`, `Values`).
- Tests use table-driven patterns and Go's standard `testing` package.
- For ordered sets, values are returned in sorted order using `slices.Sort`.

## Patterns

- Prefer immutability: sets are not mutated after creation in current usage.
- Use `cmp.Ordered` for types requiring sorting.
- Avoid exposing internal map directly; provide access via methods.
- Test coverage includes empty, single, and multiple element cases.
- Follow idiomatic Go formatting and naming.

## Additional Notes

- All code is in the `sets` package.
- No external dependencies beyond the Go standard library.
- Keep code concise and idiomatic.
