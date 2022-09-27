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

package ts

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	mm := NewMap[string, int]()
	mm.Put("hello", 10)
	mm.Put("world", 20)
	mm.Delete("world")

	if !mm.ContainsKey("hello") {
		t.Fatal("test failed!")
	}

	if mm.ContainsKey("world") {
		t.Fatal("test failed!")
	}

	if !mm.ContainsValue(10) {
		t.Fatal("test failed!")
	}

	if mm.ContainsValue(20) {
		t.Fatal("test failed!")
	}

	str := ""
	mm.Iterate(func(k string, v int) {
		str += fmt.Sprintf("%s%d", k, v)
	})

	if str != "hello10" {
		t.Fatal("test failed!")
	}
}

func TestMapKeys(t *testing.T) {
	m := NewMapValue2[int, string](1, 2, "1", "2")
	expect := []int{1, 2}
	expect2 := []int{2, 1}
	if ks := m.Keys(); !(reflect.DeepEqual(ks, expect) || reflect.DeepEqual(ks, expect2)) {
		t.Fatal("test failed!")
	}
}

func TestMapValues(t *testing.T) {
	m := NewMapValue2(1, 2, "1", "2")
	expect := []string{"1", "2"}
	if vs := m.Values(); !reflect.DeepEqual(vs, expect) {
		t.Fatal("test failed!")
	}
}

func TestNewMapValue(t *testing.T) {
	m := NewMapValue(1, 1)
	if !(m.ContainsKey(1) && m.ContainsValue(1)) {
		t.Fatal("test failed!")
	}
}

func TestNewMapValue3(t *testing.T) {
	m := NewMapValue3(1, 2, 3, 1, 2, 3)
	if !(m.ContainsKey(1) && m.ContainsValue(1)) {
		t.Fatal("test failed!")
	}
	if !(m.ContainsKey(2) && m.ContainsValue(2)) {
		t.Fatal("test failed!")
	}
	if !(m.ContainsKey(2) && m.ContainsValue(2)) {
		t.Fatal("test failed!")
	}
}

func TestMapClear(t *testing.T) {
	m := NewMapValue(1, 1)
	m.Clear()
	if m.ContainsKey(1) {
		t.Fatal("test failed!")
	}
}

func TestMapPuAll(t *testing.T) {
	m := NewMap[int, int]()
	m.PutAll([]*Entry[int, int]{NewEntry(2, 2), NewEntry(3, 3)})
	if !(m.ContainsKey(2) && m.ContainsValue(2)) {
		t.Fatal("test failed!")
	}
	if !(m.ContainsKey(3) && m.ContainsValue(3)) {
		t.Fatal("test failed!")
	}
}

func TestMapLen(t *testing.T) {
	m := NewMapValue(1, 2)
	if m.Len() != len(m) {
		t.Error("test failed!")
	}
}

func TestMapKeyReduce(t *testing.T) {
	s := NewMapValue2(1, 2, 1, 2)
	if sum := s.KeyReduce(0, func(e int, sum *int) { *sum += e }); sum != 3 {
		t.Error("test failed!")
	}
}

func TestMapValueReduce(t *testing.T) {
	s := NewMapValue2(1, 2, 1, 2)
	if sum := s.ValueReduce(0, func(e int, sum *int) { *sum += e }); sum != 3 {
		t.Error("test failed!")
	}
}
