package streams

import "testing"

func TestGroupBy(t *testing.T) {
	type groupBy struct {
		id   int
		name string
	}
	t.Log(GroupBy(func(e *groupBy) (string, *groupBy) {
		return e.name, e
	},
		&groupBy{10, "Apple"},
		&groupBy{10, "Banana"},
		&groupBy{30, "Banana"},
		&groupBy{10, "Pear"},
		&groupBy{20, "Pear"},
	))

}
