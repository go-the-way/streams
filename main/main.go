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
