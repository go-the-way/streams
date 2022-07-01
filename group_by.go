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
	m := make(map[K][]V, 0)
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

// GroupByThenToMap function
//
// T: element type
//
// K: the map key type
//
// V: the map value type
//
// KR: the result key type
//
// VR: the result value type
//
// groupByFunc: the group by function
//
// toMapFunc: the toMap function
func GroupByThenToMap[T any, K comparable, V, KR comparable, VR any](ts []T, groupByFunc func(t T) (K, V), toMapFunc func(k K, vs []V) (KR, VR)) map[KR]VR {
	m := make(map[K][]V, 0)
	rm := make(map[KR]VR, 0)
	ForEach(ts, func(_ int, t T) {
		k, v := groupByFunc(t)
		if vs, have := m[k]; have {
			vs = append(vs, v)
			m[k] = vs
		} else {
			m[k] = []V{v}
		}
	})
	MapEach(m, func(k K, vs []V) {
		kr, vr := toMapFunc(k, vs)
		rm[kr] = vr
	})
	return rm
}
