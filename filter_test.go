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

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4}
	for _, tc := range []struct {
		name   string
		f      func(e int) bool
		expect []int
	}{
		{"Gt0", func(e int) bool { return e > 0 }, []int{1, 2, 3, 4}},
		{"GtEq0", func(e int) bool { return e >= 0 }, []int{0, 1, 2, 3, 4}},
		{"Lt0", func(e int) bool { return e < 0 }, []int{}},
		{"LtEq0", func(e int) bool { return e <= 0 }, []int{0}},
		{"Even", func(e int) bool { return e%2 == 0 }, []int{0, 2, 4}},
		{"Odd", func(e int) bool { return e%2 != 0 }, []int{1, 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := Filter(arr, tc.f); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}

func TestFilterThenMap(t *testing.T) {
	mF := func(e int) int { return e + 1 }
	arr := []int{0, 1, 2, 3, 4}
	for _, tc := range []struct {
		name   string
		f      func(e int) bool
		mf     func(e int) int
		expect []int
	}{
		{"Gt0", func(e int) bool { return e > 0 }, mF, []int{2, 3, 4, 5}},
		{"GtEq0", func(e int) bool { return e >= 0 }, mF, []int{1, 2, 3, 4, 5}},
		{"Lt0", func(e int) bool { return e < 0 }, mF, []int{}},
		{"LtEq0", func(e int) bool { return e <= 0 }, mF, []int{1}},
		{"Even", func(e int) bool { return e%2 == 0 }, mF, []int{1, 3, 5}},
		{"Odd", func(e int) bool { return e%2 != 0 }, mF, []int{2, 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := FilterThenMap(arr, tc.f, tc.mf); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}

func TestFilterThenToMap(t *testing.T) {
	mF := func(e int) (int, struct{}) { return e, struct{}{} }
	arr := []int{0, 1, 2, 3, 4}
	for _, tc := range []struct {
		name   string
		f      func(e int) bool
		mf     func(e int) (int, struct{})
		expect map[int]struct{}
	}{
		{"Gt0", func(e int) bool { return e > 0 }, mF, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}},
		{"GtEq0", func(e int) bool { return e >= 0 }, mF, map[int]struct{}{0: {}, 1: {}, 2: {}, 3: {}, 4: {}}},
		{"Lt0", func(e int) bool { return e < 0 }, mF, map[int]struct{}{}},
		{"LtEq0", func(e int) bool { return e <= 0 }, mF, map[int]struct{}{0: {}}},
		{"Even", func(e int) bool { return e%2 == 0 }, mF, map[int]struct{}{0: {}, 2: {}, 4: {}}},
		{"Odd", func(e int) bool { return e%2 != 0 }, mF, map[int]struct{}{1: {}, 3: {}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := FilterThenToMap(arr, tc.f, tc.mf); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}
