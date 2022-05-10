package streams

// Map function
//
// I: input element type
//
// O: output element type
//
// mapFunc: the map function
func Map[I, O any](mapFunc func(in I) O, ins ...I) []O {
	os := make([]O, len(ins), len(ins))
	for i, in := range ins {
		os[i] = mapFunc(in)
	}
	return os
}
