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

// GroupBy function
//
// T: element type
//
// K: the map key type
//
// V: the map value type
//
// groupByFunc: the group by function
func GroupBy[T any, K comparable, V any](ts []T, groupByFunc func(t T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	ForEach(ts, func(_ int, t T) {
		k, v := groupByFunc(t)
		if vs, have := m[k]; have {
			vs = append(vs, v)
			m[k] = vs
		} else {
			m[k] = []V{v}
		}
	})
	return m
}

// GroupByThenMap function
//
// T: element type
//
// K: the map key type
//
// V: the map value type
//
// groupByFunc: the group by function
func GroupByThenMap[T any, K comparable, V, VV any](ts []T, groupByFunc func(t T) (K, V), mapFunc func(vs []V) VV) map[K]VV {
	m := make(map[K][]V)
	mm := make(map[K]VV)
	ForEach(ts, func(_ int, t T) {
		k, v := groupByFunc(t)
		if vs, have := m[k]; have {
			vs = append(vs, v)
			m[k] = vs
		} else {
			m[k] = []V{v}
		}
	})
	MapEach(m, func(k K, v []V) {
		mm[k] = mapFunc(v)
	})
	return mm
}
