package streams

import "testing"

type (
	intReduce struct{ int }
	strReduce struct{ string }
)

func TestIntReduce(t *testing.T) {
	t.Log(Reduce[int, *intReduce](
		func(e *intReduce, sum *int) { *sum += e.int },
		*new(int),
		&intReduce{10}, &intReduce{10}))
}

func TestStrReduce(t *testing.T) {
	t.Log(Reduce[string, *strReduce](
		func(e *strReduce, sum *string) { *sum += e.string },
		*new(string),
		&strReduce{"hello"}, &strReduce{" "}, &strReduce{"world"}))
}
