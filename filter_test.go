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
	"github.com/go-the-way/streams/reduces"
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

func TestFilterThenCount(t *testing.T) {
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
			if fArr := FilterThenCount(arr, tc.f); !reflect.DeepEqual(fArr, len(tc.expect)) {
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

func TestFilterThenGroupBy(t *testing.T) {
	mF := func(e int) (int, int) { return e, e }
	arr := []int{0, 1, 2, 3, 4}
	for _, tc := range []struct {
		name   string
		f      func(e int) bool
		gpf    func(e int) (int, int)
		expect map[int][]int
	}{
		{"Gt0", func(e int) bool { return e > 0 }, mF, map[int][]int{1: {1}, 2: {2}, 3: {3}, 4: {4}}},
		{"GtEq0", func(e int) bool { return e >= 0 }, mF, map[int][]int{0: {0}, 1: {1}, 2: {2}, 3: {3}, 4: {4}}},
		{"Lt0", func(e int) bool { return e < 0 }, mF, map[int][]int{}},
		{"LtEq0", func(e int) bool { return e <= 0 }, mF, map[int][]int{0: {0}}},
		{"Even", func(e int) bool { return e%2 == 0 }, mF, map[int][]int{0: {0}, 2: {2}, 4: {4}}},
		{"Odd", func(e int) bool { return e%2 != 0 }, mF, map[int][]int{1: {1}, 3: {3}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := FilterThenGroupBy(arr, tc.f, tc.gpf); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}

func TestFilterThenGroupBy2(t *testing.T) {
	type st struct {
		id1, id2 int
	}
	mF := func(s st) (int, int) { return s.id1, s.id2 }
	arr := []st{{1, 1}, {1, 2}, {2, 2}, {3, 3}}
	expect := map[int][]int{1: {1, 2}, 2: {2}, 3: {3}}
	if fArr := FilterThenGroupBy(arr, func(t st) bool { return true }, mF); !reflect.DeepEqual(fArr, expect) {
		t.Error("test failed!")
	}
}

func TestFilterThenReduce(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4}
	for _, tc := range []struct {
		name   string
		f      func(e int) bool
		expect int
	}{
		{"Gt0", func(e int) bool { return e > 0 }, 10},
		{"GtEq0", func(e int) bool { return e >= 0 }, 10},
		{"Lt0", func(e int) bool { return e < 0 }, 0},
		{"LtEq0", func(e int) bool { return e <= 0 }, 0},
		{"Even", func(e int) bool { return e%2 == 0 }, 6},
		{"Odd", func(e int) bool { return e%2 != 0 }, 4},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := FilterThenReduce(arr, tc.f, 0, reduces.Int); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}
