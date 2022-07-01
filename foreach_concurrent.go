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

import "sync"

// ForEachConcurrent function
//
// T: element type
//
// eachFunc: the each function
func ForEachConcurrent[T any](ts []T, eachFunc func(i int, t T)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(ts))
	for i, t := range ts {
		go func(i int, t T) { eachFunc(i, t); wg.Done() }(i, t)
	}
	wg.Wait()
}

// ForEachPtrConcurrent function
//
// T: element type
//
// eachFunc: the each function
func ForEachPtrConcurrent[T any](ts []T, eachFunc func(i int, t *T)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(ts))
	for i := range ts {
		go func(i int, t *T) { eachFunc(i, t); wg.Done() }(i, &ts[i])
	}
	wg.Wait()
}

// MapEachConcurrent function
//
// K: the map key type
//
// V: the map value type
//
// eachFunc: the each function
func MapEachConcurrent[K comparable, V any](m map[K]V, eachFunc func(k K, v V)) {
	wg := &sync.WaitGroup{}
	wg.Add(len(m))
	for k, v := range m {
		go func(k K, v V) { eachFunc(k, v); wg.Done() }(k, v)
	}
	wg.Wait()
}
