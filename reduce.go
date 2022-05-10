package streams

// Reduce function
//
// ID: identity value type
//
// E: element type
//
// accumulatorFunc: the accumulator function
func Reduce[ID, E any](accumulatorFunc func(e E, sum *ID), id ID, es ...E) ID {
	for _, e := range es {
		accumulatorFunc(e, &id)
	}
	return id
}
