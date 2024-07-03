// Copyright 2022 streams Author. All Rights Reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//      http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package streams

// ToMap function
//
// T: element type
//
// K: the map key type
//
// V: the map value type
//
// toMapFunc: the to map function
func ToMap[T any, K comparable, V any](ts []T, toMapFunc func(t T) (K, V)) map[K]V {
	m := make(map[K]V)
	for _, t := range ts {
		k, v := toMapFunc(t)
		m[k] = v
	}
	return m
}

// ToMapThenMap function
//
// T: element type
//
// K: the map key type
//
// KK: the map map key type
//
// V: the map value type
//
// VV: the map map value type
//
// toMapFunc: the to map function
func ToMapThenMap[T any, K, KK comparable, V, VV any](ts []T, toMapFunc func(t T) (K, V), mapFunc func(k K, v V) (KK, VV)) map[KK]VV {
	m := make(map[KK]VV)
	for _, t := range ts {
		k, v := toMapFunc(t)
		kk, vv := mapFunc(k, v)
		m[kk] = vv
	}
	return m
}

// ToMap2 function
//
// T: element type
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
func ToMap2[T any, K1, K2 comparable, V1, V2 any](ts []T, toMapFunc1 func(t T) (K1, V1), toMapFunc2 func(t T) (K2, V2)) (map[K1]V1, map[K2]V2) {
	m1 := make(map[K1]V1)
	m2 := make(map[K2]V2)
	for _, t := range ts {
		k1, v1 := toMapFunc1(t)
		k2, v2 := toMapFunc2(t)
		m1[k1] = v1
		m2[k2] = v2
	}
	return m1, m2
}

// ToMap3 function
//
// T: element type
//
// K1: the map1 key type
//
// K2: the map2 key type
//
// K3: the map3 key type
//
// V1: the map1 value type
//
// V2: the map2 value type
//
// V3: the map3 value type
//
// toMapFunc1: the to map1 function
//
// toMapFunc2: the to map2 function
//
// toMapFunc3: the to map3 function
func ToMap3[T any, K1, K2, K3 comparable, V1, V2, V3 any](ts []T, toMapFunc1 func(t T) (K1, V1), toMapFunc2 func(t T) (K2, V2), toMapFunc3 func(t T) (K3, V3)) (map[K1]V1, map[K2]V2, map[K3]V3) {
	m1 := make(map[K1]V1)
	m2 := make(map[K2]V2)
	m3 := make(map[K3]V3)
	for _, t := range ts {
		k1, v1 := toMapFunc1(t)
		k2, v2 := toMapFunc2(t)
		k3, v3 := toMapFunc3(t)
		m1[k1] = v1
		m2[k2] = v2
		m3[k3] = v3
	}
	return m1, m2, m3
}

// MapToMap function
//
// KI: input key type
//
// KO: output key type
//
// VI: input value type
//
// VO: output value type
//
// mapFunc: the map function
func MapToMap[KI, KO comparable, VI, VO any](m map[KI]VI, toMapFunc func(ki KI, vi VI) (KO, VO)) map[KO]VO {
	om := make(map[KO]VO)
	for k, v := range m {
		ko, vo := toMapFunc(k, v)
		om[ko] = vo
	}
	return om
}

// MapToMapThenFilter function
//
// KI: input key type
//
// KO: output key type
//
// VI: input value type
//
// VO: output value type
//
// mapFunc: the map function
//
// filterFunc: the filter function
func MapToMapThenFilter[KI, KO comparable, VI, VO any](m map[KI]VI, toMapFunc func(ki KI, vi VI) (KO, VO), filterFunc func(KO, VO) bool) map[KO]VO {
	om := make(map[KO]VO)
	for k, v := range m {
		if ko, vo := toMapFunc(k, v); filterFunc(ko, vo) {
			om[ko] = vo
		}
	}
	return om
}

// MapToMapThenReduce function
//
// KI: input key type
//
// KO: output key type
//
// VI: input value type
//
// VO: output value type
//
// R: result type
//
// mapFunc: the map function
//
// reduceFunc: the reduce function
func MapToMapThenReduce[KI, KO comparable, VI, VO, R any](m map[KI]VI, toMapFunc func(ki KI, vi VI) (KO, VO), r R, reduceFunc func(k KO, v VO, sum *R)) R {
	om := make(map[KO]VO)
	for k, v := range m {
		ko, vo := toMapFunc(k, v)
		om[ko] = vo
	}
	return ReduceMap(om, r, reduceFunc)
}
