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

type Entry[K comparable, V any] struct {
	K K
	V V
}

func NewEntry[K comparable, V any](k K, v V) *Entry[K, V] { return &Entry[K, V]{K: k, V: v} }

func MapToEntry[K comparable, V any](m map[K]V) []*Entry[K, V] {
	ts := make([]*Entry[K, V], 0)
	for k, v := range m {
		ts = append(ts, NewEntry[K, V](k, v))
	}
	return ts
}

func EntryToMap[K comparable, V any](es []*Entry[K, V]) map[K]V {
	m := make(map[K]V, 0)
	for _, e := range es {
		m[e.K] = e.V
	}
	return m
}
