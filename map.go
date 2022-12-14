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

// Map function
//
// I: input element type
//
// O: output element type
//
// mapFunc: the map function
func Map[I, O any](ins []I, mapFunc func(in I) O) []O {
	os := make([]O, len(ins), len(ins))
	for i, in := range ins {
		os[i] = mapFunc(in)
	}
	return os
}

// MapThenFilter function
//
// I: input element type
//
// O: output element type
//
// mapFunc: the map function
//
// filterFunc: the filter function
func MapThenFilter[I, O any](ins []I, mapFunc func(in I) O, filterFunc func(out O) bool) []O {
	os := make([]O, 0)
	for _, in := range ins {
		if o := mapFunc(in); filterFunc(o) {
			os = append(os, o)
		}
	}
	return os
}

// MapThenReduce function
//
// I: input element type
//
// O: output element type
//
// R: result type
//
// mapFunc: the map function
//
// reduceFunc: the reduce function
func MapThenReduce[I, O, R any](ins []I, mapFunc func(in I) O, r R, reduceFunc func(r O, sum *R)) R {
	os := make([]O, 0)
	for _, in := range ins {
		os = append(os, mapFunc(in))
	}
	return Reduce(os, r, reduceFunc)
}

// MapThenGroupBy function
//
// I: input element type
//
// O: output element type
//
// K: the map key type
//
// V: the map value type
//
// mapFunc: the map function
//
// groupBy: the group by function
func MapThenGroupBy[I, O, K comparable, V any](ins []I, mapFunc func(in I) O, groupByFunc func(o O) (K, V)) map[K][]V {
	os := make([]O, 0)
	for _, in := range ins {
		os = append(os, mapFunc(in))
	}
	return GroupBy(os, groupByFunc)
}

// MapMap function
//
// K: the map key type
//
// V: the map value type
//
// R: result type
//
// mapFunc: the map function
func MapMap[K comparable, V, R any](m map[K]V, mapFunc func(k K, v V) R) []R {
	rs := make([]R, 0)
	for k, v := range m {
		rs = append(rs, mapFunc(k, v))
	}
	return rs
}

// MapMapThenFilter function
//
// I: input element type
//
// O: output element type
//
// mapFunc: the map function
//
// filterFunc: the filter function
func MapMapThenFilter[K comparable, V, R any](m map[K]V, mapFunc func(k K, v V) R, filterFunc func(r R) bool) []R {
	rs := make([]R, 0)
	for k, v := range m {
		if r := mapFunc(k, v); filterFunc(r) {
			rs = append(rs, r)
		}
	}
	return rs
}

// MapMapThenReduce function
//
// K: the map key type
//
// V: the map value type
//
// R: result type
//
// RR: reduce type
//
// mapFunc: the map function
//
// reduceFunc: the reduce function
func MapMapThenReduce[K comparable, V, R, RR any](m map[K]V, mapFunc func(k K, v V) R, rr RR, reduceFunc func(r R, sum *RR)) RR {
	rs := make([]R, 0)
	for k, v := range m {
		rs = append(rs, mapFunc(k, v))
	}
	return Reduce(rs, rr, reduceFunc)
}
