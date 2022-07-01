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

package types

import "reflect"

// Map define
type Map[K, V comparable] map[K]V

// MakeMap make Map instance
func MakeMap[K, V comparable]() Map[K, V] {
	return Map[K, V]{}
}

// Put a K-V
func (m Map[K, V]) Put(k K, v V) { m[k] = v }

// Delete key
func (m Map[K, V]) Delete(k K) { delete(m, k) }

// ContainsKey key
func (m Map[K, V]) ContainsKey(k K) bool { _, have := m[k]; return have }

// ContainsValue value
func (m Map[K, V]) ContainsValue(v V) bool {
	for _, _v := range m {
		if reflect.DeepEqual(_v, v) {
			return true
		}
	}
	return false
}

// Iterate K-V
func (m Map[K, V]) Iterate(f func(k K, v V)) {
	for k, v := range m {
		f(k, v)
	}
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	ks := make([]K, 0)
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func MapValues[K comparable, V any](m map[K]V) []V {
	vs := make([]V, 0)
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}
