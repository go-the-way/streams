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

package ts

import "reflect"

type Map[K, V comparable] map[K]V

func NewMap[K, V comparable]() Map[K, V]                         { return Map[K, V]{} }
func NewMapValue[K, V comparable](k K, v V) Map[K, V]            { return Map[K, V]{k: v} }
func NewMapValue2[K, V comparable](k1, k2 K, v1, v2 V) Map[K, V] { return Map[K, V]{k1: v1, k2: v2} }
func NewMapValue3[K, V comparable](k1, k2, k3 K, v1, v2, v3 V) Map[K, V] {
	return Map[K, V]{k1: v1, k2: v2, k3: v3}
}
func (m Map[K, V]) Put(k K, v V) Map[K, V] { m[k] = v; return m }
func (m Map[K, V]) Delete(k K)             { delete(m, k) }
func (m Map[K, V]) Clear()                 { m.Iterate(func(k K, _ V) { m.Delete(k) }) }
func (m Map[K, V]) ContainsKey(k K) bool   { _, have := m[k]; return have }
func (m Map[K, V]) Len() int               { return len(m) }

func (m Map[K, V]) PutAll(entries []*Entry[K, V]) Map[K, V] {
	for _, entry := range entries {
		m.Put(entry.K, entry.V)
	}
	return m
}

func (m Map[K, V]) ContainsValue(v V) bool {
	for _, _v := range m {
		// TODO: reflect equal?
		if reflect.DeepEqual(_v, v) {
			return true
		}
	}
	return false
}

func (m Map[K, V]) Iterate(f func(k K, v V)) {
	for k, v := range m {
		f(k, v)
	}
}

func (m Map[K, V]) Keys() []K {
	ks := make([]K, 0)
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func (m Map[K, V]) Values() []V {
	vs := make([]V, 0)
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func (m Map[K, V]) KeyReduce(id K, reduceFunc func(k K, sum *K)) K {
	m.Iterate(func(k K, v V) { reduceFunc(k, &id) })
	return id
}

func (m Map[K, V]) ValueReduce(id V, reduceFunc func(v V, sum *V)) V {
	m.Iterate(func(k K, v V) { reduceFunc(v, &id) })
	return id
}
