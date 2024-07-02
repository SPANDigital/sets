# sets
Sets implentations in go

## Examplee

### Set
```go
package main

import (
	"github.com/spandigital/sets"
)

func main() {
	stuff := sets.NewSet("a", "b", "c", "d")
	println(stuff.Contains("a")) //true
	println(stuff.Values())      // all values will be returned but order is not guaranteed
}
```

### Ordered Set
```go
package main

import (
	"github.com/spandigital/sets"
)

func main() {
	stuff := sets.NewOrderedSet("a", "b", "c", "d")
	println(stuff.Contains("a")) //true
	println(stuff.Values())      //"a", "b", "c", "d"
}
```

