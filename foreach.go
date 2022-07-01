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

// ForEach function
//
// T: element type
//
// eachFunc: the each function
func ForEach[T any](ts []T, eachFunc func(i int, t T)) {
	for i, t := range ts {
		eachFunc(i, t)
	}
}

// ForEachPtr function
//
// T: element type
//
// eachFunc: the each function
func ForEachPtr[T any](ts []T, eachFunc func(i int, t *T)) {
	for i := range ts {
		eachFunc(i, &ts[i])
	}
}

// MapEach function
//
// K: the map key type
//
// V: the map value type
//
// eachFunc: the each function
func MapEach[K comparable, V any](m map[K]V, eachFunc func(k K, v V)) {
	for k, v := range m {
		eachFunc(k, v)
	}
}
