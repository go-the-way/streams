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

// Set define
type Set[E comparable] map[E]struct{}

// MakeSet make set instance
func MakeSet[E comparable]() Set[E] {
	return Set[E]{}
}

// Add an element
func (s Set[E]) Add(element E) { s[element] = struct{}{} }

// Delete element
func (s Set[E]) Delete(element E) { delete(s, element) }

// Contains element
func (s Set[E]) Contains(element E) bool { _, have := s[element]; return have }

// Iterate element
func (s Set[E]) Iterate(f func(element E)) {
	for k := range s {
		f(k)
	}
}
