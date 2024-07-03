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

import "reflect"

type Map0[K, V comparable] map[K]V

func NewMap[K, V comparable]() Map0[K, V]                         { return Map0[K, V]{} }
func NewMapValue[K, V comparable](k K, v V) Map0[K, V]            { return Map0[K, V]{k: v} }
func NewMapValue2[K, V comparable](k1, k2 K, v1, v2 V) Map0[K, V] { return Map0[K, V]{k1: v1, k2: v2} }
func NewMapValue3[K, V comparable](k1, k2, k3 K, v1, v2, v3 V) Map0[K, V] {
	return Map0[K, V]{k1: v1, k2: v2, k3: v3}
}

func (m Map0[K, V]) Len() int                { return len(m) }
func (m Map0[K, V]) Put(k K, v V) Map0[K, V] { m[k] = v; return m }

func (m Map0[K, V]) PutAll(entries []*Entry[K, V]) Map0[K, V] {
	for _, entry := range entries {
		m.Put(entry.K, entry.V)
	}
	return m
}

func (m Map0[K, V]) Delete(k K)           { delete(m, k) }
func (m Map0[K, V]) Clear()               { m.Iterate(func(k K, _ V) { m.Delete(k) }) }
func (m Map0[K, V]) ContainsKey(k K) bool { _, have := m[k]; return have }

func (m Map0[K, V]) ContainsValue(v V) bool {
	for _, v0 := range m {
		if reflect.DeepEqual(v0, v) {
			return true
		}
	}
	return false
}

func (m Map0[K, V]) Iterate(f func(k K, v V)) {
	for k, v := range m {
		f(k, v)
	}
}

func (m Map0[K, V]) Keys() []K {
	ks := make([]K, 0)
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func (m Map0[K, V]) Values() []V {
	vs := make([]V, 0)
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func (m Map0[K, V]) KeysReduce(sum K, reduceFunc func(k K, sum *K)) K {
	m.Iterate(func(k K, _ V) { reduceFunc(k, &sum) })
	return sum
}

func (m Map0[K, V]) ValuesReduce(sum V, reduceFunc func(v V, sum *V)) V {
	m.Iterate(func(_ K, v V) { reduceFunc(v, &sum) })
	return sum
}
