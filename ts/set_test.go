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

func TestSet(t *testing.T) {
	set := NewSet[string]()
	set.Add("hello")
	set.Add("world")
	set.Delete("world")

	if !set.Contains("hello") {
		t.Fatal("test failed!")
	}

	if set.Contains("world") {
		t.Fatal("test failed!")
	}

	str := ""
	set.Iterate(func(t string) {
		str += fmt.Sprintf("%s", t)
	})

	if str != "hello" {
		t.Fatal("test failed!")
	}
}

func TestSetClear(t *testing.T) {
	s := NewSet[int]()
	s.AddAll(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	s.Clear()
	if len(s) > 0 {
		t.Error("test failed!")
	}
}

func TestSetAddAll(t *testing.T) {
	s := NewSet[int]()
	s.AddAll(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)

	for _, i := range []int{1, 2, 3, 4, 5} {
		if !s.Contains(i) {
			t.Error("test failed!")
		}
	}
}

func TestNewSet(t *testing.T) {
	s := NewSetValue(1, 2, 3, 4)
	if !s.Contains(1) {
		t.Error("test failed!")
	}
	if !s.Contains(2) {
		t.Error("test failed!")
	}
	if !s.Contains(3) {
		t.Error("test failed!")
	}
	if !s.Contains(4) {
		t.Error("test failed!")
	}
}

func TestSetFilter(t *testing.T) {
	s := NewSetValue(1, 2, 3, 4)
	if c := len(s.Filter(func(e int) bool { return e > 0 })); c != 4 {
		t.Error("test failed!")
	}
}

func TestSetReduce(t *testing.T) {
	s := NewSetValue(1, 2, 3, 4)
	if sum := s.Reduce(0, func(e int, sum *int) { *sum += e }); sum != 10 {
		t.Error("test failed!")
	}
}

func TestSetList(t *testing.T) {
	s := NewSetValue(1, 2)
	if list := s.List(); !reflect.DeepEqual([]int{1, 2}, list) {
		t.Error("test failed!")
	}
}

func TestSetLen(t *testing.T) {
	s := NewSetValue(1, 2)
	if s.Len() != len(s) {
		t.Error("test failed!")
	}
}
