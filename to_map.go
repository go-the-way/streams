package streams

// ToMap function
//
// E: element type
//
// K: the map key type
//
// V: the map value type
//
// toMapFunc: the to map function
func ToMap[E any, K comparable, V any](toMapFunc func(e E) (K, V), es ...E) map[K]V {
	m := make(map[K]V, 0)
	for _, e := range es {
		k, v := toMapFunc(e)
		m[k] = v
	}
	return m
}

// ToMap2 function
//
// E: element type
//
// K1: the map1 key type
//
// V1: the map1 value type
//
// K2: the map2 key type
//
// V2: the map2 value type
//
// toMapFunc1: the to map1 function
//
// toMapFunc2: the to map2 function
func ToMap2[E any, K1, K2 comparable, V1, V2 any](toMapFunc1 func(e E) (K1, V1), toMapFunc2 func(e E) (K2, V2), es ...E) (map[K1]V1, map[K2]V2) {
	m1 := make(map[K1]V1, 0)
	m2 := make(map[K2]V2, 0)
	for _, e := range es {
		k1, v1 := toMapFunc1(e)
		k2, v2 := toMapFunc2(e)
		m1[k1] = v1
		m2[k2] = v2
	}
	return m1, m2
}
