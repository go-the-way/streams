package streams

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	is := Map[int, string](func(in int) string {
		return fmt.Sprintf("%x", in)
	}, 1, 2, 3, 4, 5)
	t.Log(is)
}
