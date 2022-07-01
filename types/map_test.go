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

package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	mm := MakeMap[string, int]()
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
	m := map[int]string{1: "1", 2: "2"}
	expect := []int{1, 2}
	if ks := MapKeys(m); !reflect.DeepEqual(ks, expect) {
		t.Fatal("test failed!")
	}
}

func TestMapValues(t *testing.T) {
	m := map[int]string{1: "1", 2: "2"}
	expect := []string{"1", "2"}
	if vs := MapValues(m); !reflect.DeepEqual(vs, expect) {
		t.Fatal("test failed!")
	}
}
