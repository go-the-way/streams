package streams

import "testing"

func TestFilter(t *testing.T) {
	t.Log(Filter[int](func(e int) bool {
		return e%2 == 0
	}, 1, 2, 3, 4, 6, 7, 8, 9, 10))
}
