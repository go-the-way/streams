# streams

Basic Stream API written in Go.(Only supports Go1.18+)

[![CircleCI](https://circleci.com/gh/go-the-way/streams/tree/main.svg?style=shield)](https://circleci.com/gh/go-the-way/streams/tree/main)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/go-the-way/streams)
[![codecov](https://codecov.io/gh/go-the-way/streams/branch/main/graph/badge.svg?token=8MAR3J959H)](https://codecov.io/gh/go-the-way/streams)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-the-way/streams)](https://goreportcard.com/report/github.com/go-the-way/streams)
[![GoDoc](https://pkg.go.dev/badge/github.com/go-the-way/streams?status.svg)](https://pkg.go.dev/github.com/go-the-way/streams?tab=doc)

## Functions
- [Filter](#filter)
- [Map](#map)
- [MapMap](#mapmap)
- [Reduce](#reduce)
- [Skip](#skip)
- [GroupBy](#groupby)
- [ToMap](#tomap)
- [ToMap2](#tomap2)
- [Each](#each)
- [EachPtr](#eachptr)

### Filter

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	fmt.Println(streams.Filter[int](func(e int) bool {
		return e%2 == 0
	}, 1, 2, 3, 4, 6, 7, 8, 9, 10))
}
```

### Map

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	fmt.Println(streams.Map[int, string](func(in int) string {
		return fmt.Sprintf("%x", in)
	}, 1, 2, 3, 4, 5))
}
```

### MapMap

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	m := map[string]int{"1": 1, "2": 2}
	fmt.Println(streams.MapMap(func(k string, v int) (string, int) { return k, v + 1 }, m))
}
```

### Reduce

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

type (
	intReduce struct{ int }
)

func main() {
	fmt.Println(streams.Reduce[int, *intReduce](
		func(e *intReduce, sum *int) { *sum += e.int },
		*new(int),
		&intReduce{10}, &intReduce{10}))
}
```

### Skip

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	fmt.Println(streams.Skip(1, 1, 2, 3, 4))
}
```

### GroupBy

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	type groupBy struct {
		id   int
		name string
	}
	fmt.Println(streams.GroupBy(func(e *groupBy) (string, *groupBy) {
		return e.name, e
	},
		&groupBy{10, "Apple"},
		&groupBy{10, "Banana"},
		&groupBy{30, "Banana"},
		&groupBy{10, "Pear"},
		&groupBy{20, "Pear"},
	))
}
```

### ToMap

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	type toMap struct {
		id   int
		name string
	}
	fmt.Println(streams.ToMap(func(e *toMap) (int, string) {
		return e.id, e.name
	}, &toMap{10, "Apple"}, &toMap{20, "Pear"}))
}
```

### ToMap2

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	type toMap struct {
		id   int
		name string
	}
	fmt.Println(streams.ToMap2(func(e *toMap) (int, string) {
		return e.id, e.name
	}, func(e *toMap) (string, int) {
		return e.name, e.id
	}, &toMap{10, "Apple"}, &toMap{20, "Pear"}))
}
```

### Each

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	arr := []int{0, 1, 2}
	newArr := make([]int, 0)
	streams.Each(func(e int) { newArr = append(newArr, e+1) }, arr...)
	fmt.Println(newArr)
}
```

### EachPtr

```go
package main

import (
	"fmt"

	"github.com/go-the-way/streams"
)

func main() {
	arr := []int{0, 1, 2} 
	streams.EachPtr(func(e *int) { *e += 1 }, arr...)
	fmt.Println(arr)
}
```
