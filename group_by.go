package streams

// GroupBy function
//
// E: element type
//
// K: the map key type
//
// V: the map value type
//
// groupByFunc: the group by function
func GroupBy[E any, K comparable, V any](groupByFunc func(e E) (K, V), es ...E) map[K][]V {
	m := make(map[K][]V)
	for _, e := range es {
		k, v := groupByFunc(e)
		if vs, have := m[k]; have {
			vs = append(vs, v)
			m[k] = vs
		} else {
			m[k] = []V{v}
		}
	}
	return m
}
