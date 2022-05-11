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
func Map[I, O any](mapFunc func(in I) O, ins ...I) []O {
	os := make([]O, len(ins), len(ins))
	for i, in := range ins {
		os[i] = mapFunc(in)
	}
	return os
}

// MapMap function
//
// KI: input key type
//
// VI: input value type
//
// KO: output key type
//
// VO: output value type
//
// mapFunc: the map function
func MapMap[KI comparable, VI any, KO comparable, VO any](mapFunc func(ki KI, vi VI) (KO, VO), m map[KI]VI) map[KO]VO {
	om := make(map[KO]VO, 0)
	for k, v := range m {
		ko, vo := mapFunc(k, v)
		om[ko] = vo
	}
	return om
}
