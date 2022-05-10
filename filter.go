package streams

// Filter function
//
// E: element type
//
// filterFunc: the filter function
func Filter[E any](filterFunc func(e E) bool, es ...E) []E {
	nEs := make([]E, 0)
	for _, e := range es {
		if filterFunc(e) {
			nEs = append(nEs, e)
		}
	}
	return nEs
}
