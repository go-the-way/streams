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

// Filter function
//
// T: element type
//
// filterFunc: the filter function
func Filter[T any](ts []T, filterFunc func(t T) bool) []T {
	nEs := make([]T, 0)
	for _, t := range ts {
		if filterFunc(t) {
			nEs = append(nEs, t)
		}
	}
	return nEs
}

// FilterThenMap function
//
// T: element type
//
// R: result type
//
// filterFunc: the filter function
func FilterThenMap[T, R any](ts []T, filterFunc func(t T) bool, mapFunc func(t T) R) []R {
	rs := make([]R, 0)
	for _, t := range ts {
		if filterFunc(t) {
			rs = append(rs, mapFunc(t))
		}
	}
	return rs
}

// FilterThenToMap function
//
// T: element type
//
// K: map key type
//
// V: map value type
//
// filterFunc: the filter function
func FilterThenToMap[T any, K comparable, V any](ts []T, filterFunc func(t T) bool, toMapFunc func(t T) (K, V)) map[K]V {
	m := make(map[K]V, 0)
	for _, t := range ts {
		if filterFunc(t) {
			k, v := toMapFunc(t)
			m[k] = v
		}
	}
	return m
}
