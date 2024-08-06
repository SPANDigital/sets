# sets
Sets implentations in go

## Example

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

    numStuff := sets.NewSet(1, 2, 3, 4)
	println(numStuff.Contains(2)) //true
	println(numStuff.Values())    // all values will be returned but order is not guaranteed
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
	
	numStuff := sets.NewOrderedSet(1, 2, 3, 4, 5, 6)
	println(numStuff.Contains(2)) //true
	println(numStuff.Values())    //1, 2, 3, 4, 5, 6
}
```

