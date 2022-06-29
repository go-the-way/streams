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

// MapMap function
//
// KI: input key type
//
// KO: output key type
//
// VI: input value type
//
// VO: output value type
//
// mapFunc: the map function
func MapMap[KI, KO comparable, VI, VO any](m map[KI]VI, mapFunc func(ki KI, vi VI) (KO, VO)) map[KO]VO {
	om := make(map[KO]VO, 0)
	for k, v := range m {
		ko, vo := mapFunc(k, v)
		om[ko] = vo
	}
	return om
}

// MapMapThenFilter function
//
// KI: input key type
//
// KO: output key type
//
// VI: input value type
//
// VO: output value type
//
// mapFunc: the map function
//
// filterFunc: the filter function
func MapMapThenFilter[KI, KO comparable, VI, VO any](m map[KI]VI, mapFunc func(ki KI, vi VI) (KO, VO), filterFunc func(KO, VO) bool) map[KO]VO {
	om := make(map[KO]VO, 0)
	for k, v := range m {
		if ko, vo := mapFunc(k, v); filterFunc(ko, vo) {
			om[ko] = vo
		}
	}
	return om
}

// MapMapThenReduce function
//
// KI: input key type
//
// KO: output key type
//
// VI: input value type
//
// VO: output value type
//
// R: result type
//
// mapFunc: the map function
//
// reduceFunc: the reduce function
func MapMapThenReduce[KI, KO comparable, VI, VO, R any](m map[KI]VI, mapFunc func(ki KI, vi VI) (KO, VO), r R, reduceFunc func(k KO, v VO, sum *R)) R {
	om := make(map[KO]VO, 0)
	for k, v := range m {
		ko, vo := mapFunc(k, v)
		om[ko] = vo
	}
	return ReduceMap(om, r, reduceFunc)
}
