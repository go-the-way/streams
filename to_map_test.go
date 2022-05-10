package streams

import "testing"

func TestToMap(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	t.Log(ToMap(func(e *toMap) (int, string) {
		return e.id, e.name
	}, &toMap{10, "Apple"}, &toMap{20, "Pear"}))
}

func TestToMap2(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	t.Log(ToMap2(func(e *toMap) (int, string) {
		return e.id, e.name
	}, func(e *toMap) (string, int) {
		return e.name, e.id
	}, &toMap{10, "Apple"}, &toMap{20, "Pear"}))
}
