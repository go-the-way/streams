package streams

// Skip function
//
// E: element type
func Skip[E any](index int, es ...E) []E {
	nEs := make([]E, 0)
	nEs = append(nEs, es[:index]...)
	nEs = append(nEs, es[index+1:]...)
	return nEs
}
