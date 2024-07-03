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

func TestToSet(t *testing.T) {
	type toSet struct{ id int }
	arr := []*toSet{{10}}
	expect := NewSetValue(10)
	if m := ToSet(arr, func(e *toSet) int { return e.id }); !reflect.DeepEqual(m, expect) {
		t.Error("test failed!")
	}
}

func TestToSetThenMap(t *testing.T) {
	type toSet struct{ id int }
	arr := []*toSet{{10}}
	expect := []int{10}
	if m := ToSetThenMap(arr, func(e *toSet) int { return e.id }, func(r int) int { return r }); !reflect.DeepEqual(m, expect) {
		t.Error("test failed!")
	}
}

func TestToSet2(t *testing.T) {
	type toSet struct {
		id   int
		name string
	}
	arr := []*toSet{{10, "Apple"}}
	expect1, expect2 := NewSetValue(10), NewSetValue("Apple")
	if s1, s2 := ToSet2(arr, func(e *toSet) int { return e.id }, func(e *toSet) string { return e.name }); !reflect.DeepEqual(s1, expect1) || !reflect.DeepEqual(s2, expect2) {
		t.Error("test failed!")
	}
}

func TestToSet3(t *testing.T) {
	type toSet struct {
		id   int
		name string
		age  int
	}
	arr := []*toSet{{10, "Apple", 11}}
	expect1, expect2, expect3 := NewSetValue(10), NewSetValue("Apple"), NewSetValue(11)
	if s1, s2, s3 := ToSet3(arr, func(e *toSet) int { return e.id }, func(e *toSet) string { return e.name }, func(e *toSet) int { return e.age }); !reflect.DeepEqual(s1, expect1) || !reflect.DeepEqual(s2, expect2) || !reflect.DeepEqual(s3, expect3) {
		t.Error("test failed!")
	}
}
