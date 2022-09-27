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
	"fmt"
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	arr := []*toMap{{10, "Apple"}, {20, "Pear"}}
	expect := map[int]string{10: "Apple", 20: "Pear"}
	if m := ToMap(arr, func(e *toMap) (int, string) {
		return e.id, e.name
	}); !reflect.DeepEqual(m, expect) {
		t.Error("test failed!")
	}
}

func TestToMapThenMap(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	arr := []*toMap{{10, "Apple"}, {20, "Pear"}}
	expect := map[string]int{"Apple": 10, "Pear": 20}
	if m := ToMapThenMap(arr, func(e *toMap) (int, string) {
		return e.id, e.name
	}, func(k int, v string) (string, int) {
		return v, k
	}); !reflect.DeepEqual(m, expect) {
		t.Error("test failed!")
	}
}

func TestToMap2(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	arr := []*toMap{{10, "Apple"}, {20, "Pear"}}
	except1 := map[int]string{10: "Apple", 20: "Pear"}
	except2 := map[string]int{"Apple": 10, "Pear": 20}
	if m1, m2 := ToMap2(arr, func(e *toMap) (int, string) {
		return e.id, e.name
	}, func(e *toMap) (string, int) {
		return e.name, e.id
	}); !reflect.DeepEqual(m1, except1) || !reflect.DeepEqual(m2, except2) {
		t.Error("test failed!")
	}
}

func TestToMap3(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	arr := []*toMap{{10, "Apple"}, {20, "Pear"}}
	except1 := map[int]string{10: "Apple", 20: "Pear"}
	except2 := map[string]int{"Apple": 10, "Pear": 20}
	except3 := map[string]int{"Apple1": 10, "Pear1": 20}
	if m1, m2, m3 := ToMap3(arr, func(e *toMap) (int, string) {
		return e.id, e.name
	}, func(e *toMap) (string, int) {
		return e.name, e.id
	}, func(e *toMap) (string, int) {
		return fmt.Sprintf("%s1", e.name), e.id
	}); !reflect.DeepEqual(m1, except1) || !reflect.DeepEqual(m2, except2) || !reflect.DeepEqual(m3, except3) {
		t.Error("test failed!")
	}
}

func TestMapToMap(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2}
	expect := map[string]int{"1": 2, "2": 3}
	if a := MapToMap(m, func(k string, v int) (string, int) { return k, v + 1 }); !reflect.DeepEqual(a, expect) {
		t.Error("test failed!")
	}
}

func TestMapToMapThenFilter(t *testing.T) {
	{
		m := map[string]int{"1": 1, "2": 2}
		expect := map[string]int{"1": 2, "2": 3}
		if a := MapToMapThenFilter(m, func(k string, v int) (string, int) { return k, v + 1 }, func(k string, v int) bool { return true }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed!")
		}
	}
	{
		m := map[string]int{"1": 1, "2": 2}
		expect := map[string]int{}
		if a := MapToMapThenFilter(m, func(k string, v int) (string, int) { return k, v + 1 }, func(k string, v int) bool { return false }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed!")
		}
	}
}

func TestMapToMapThenReduce(t *testing.T) {
	rF := func(k string, v int, sum *string) { *sum += fmt.Sprintf("%s%d", k, v) }
	m := map[string]int{"1": 1}
	expect := "12"
	a := MapToMapThenReduce(m, func(k string, v int) (string, int) { return k, v + 1 }, "", rF)
	if !reflect.DeepEqual(a, expect) {
		t.Error("test failed!")
	}
}
