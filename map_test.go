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
	"github.com/go-the-way/streams/reduces"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := []string{"1", "2", "3"}
	if a := Map(arr, func(e int) string { return fmt.Sprintf("%d", e) }); !reflect.DeepEqual(a, expect) {
		t.Error("test failed")
	}
}

func TestMapThenFilter(t *testing.T) {
	{
		arr := []int{1, 2, 3}
		expect := []string{"1", "2", "3"}
		if a := MapThenFilter(arr, func(e int) string { return fmt.Sprintf("%d", e) }, func(e string) bool { return true }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed")
		}
	}
	{
		arr := []int{1, 2, 3}
		expect := make([]string, 0)
		if a := MapThenFilter(arr, func(e int) string { return fmt.Sprintf("%d", e) }, func(e string) bool { return false }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed")
		}
	}
}

func TestMapThenReduce(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := "123"
	if a := MapThenReduce(arr, func(e int) string { return fmt.Sprintf("%d", e) }, "", reduces.String); !reflect.DeepEqual(a, expect) {
		t.Error("test failed")
	}
}

func TestMapMap(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2}
	expect := map[string]int{"1": 2, "2": 3}
	if a := MapMap(m, func(k string, v int) (string, int) { return k, v + 1 }); !reflect.DeepEqual(a, expect) {
		t.Error("test failed")
	}
}

func TestMapMapThenFilter(t *testing.T) {
	{
		m := map[string]int{"1": 1, "2": 2}
		expect := map[string]int{"1": 2, "2": 3}
		if a := MapMapThenFilter(m, func(k string, v int) (string, int) { return k, v + 1 }, func(k string, v int) bool { return true }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed")
		}
	}
	{
		m := map[string]int{"1": 1, "2": 2}
		expect := map[string]int{}
		if a := MapMapThenFilter(m, func(k string, v int) (string, int) { return k, v + 1 }, func(k string, v int) bool { return false }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed")
		}
	}
}

func TestMapMapThenReduce(t *testing.T) {
	rF := func(k string, v int, sum *string) { *sum += fmt.Sprintf("%s%d", k, v) }
	m := map[string]int{"1": 1, "2": 2}
	expect := "1223"
	if a := MapMapThenReduce(m, func(k string, v int) (string, int) { return k, v + 1 }, "", rF); !reflect.DeepEqual(a, expect) {
		t.Error("test failed")
	}
}
