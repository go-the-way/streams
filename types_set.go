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

type Set[E comparable] map[E]struct{}

func NewSet[E comparable]() Set[E]            { return Set[E]{} }
func NewSetValue[E comparable](e ...E) Set[E] { s := Set[E]{}; s.AddAll(e...); return s }
func (s Set[E]) Len() int                     { return len(s) }
func (s Set[E]) Add(element E) Set[E]         { s[element] = struct{}{}; return s }

func (s Set[E]) AddAll(elements ...E) Set[E] {
	for _, e := range elements {
		s.Add(e)
	}
	return s
}

func (s Set[E]) Delete(element E) Set[E] { delete(s, element); return s }
func (s Set[E]) Clear() Set[E]           { s.Iterate(func(e E) { s.Delete(e) }); return s }
func (s Set[E]) Contains(element E) bool { _, have := s[element]; return have }

func (s Set[E]) Iterate(f func(element E)) {
	for k := range s {
		f(k)
	}
}

func (s Set[E]) Filter(filterFunc func(e E) bool) Set[E] {
	newSet := NewSet[E]()
	s.Iterate(func(e E) {
		if filterFunc(e) {
			newSet.Add(e)
		}
	})
	return newSet
}

func (s Set[E]) Reduce(id E, reduceFunc func(e E, sum *E)) E {
	s.Iterate(func(e E) { reduceFunc(e, &id) })
	return id
}

func (s Set[E]) List() []E {
	l := make([]E, 0, len(s))
	s.Iterate(func(e E) { l = append(l, e) })
	return l
}
