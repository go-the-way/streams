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
