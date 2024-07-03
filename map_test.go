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
	"github.com/go-the-way/streams/reducefunc"
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := []string{"1", "2", "3"}
	if a := Map(arr, func(e int) string { return fmt.Sprintf("%d", e) }); !reflect.DeepEqual(a, expect) {
		t.Error("test failed!")
	}
}

func TestMapThenFilter(t *testing.T) {
	{
		arr := []int{1, 2, 3}
		expect := []string{"1", "2", "3"}
		if a := MapThenFilter(arr, func(e int) string { return fmt.Sprintf("%d", e) }, func(e string) bool { return true }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed!")
		}
	}
	{
		arr := []int{1, 2, 3}
		expect := make([]string, 0)
		if a := MapThenFilter(arr, func(e int) string { return fmt.Sprintf("%d", e) }, func(e string) bool { return false }); !reflect.DeepEqual(a, expect) {
			t.Error("test failed!")
		}
	}
}

func TestMapThenReduce(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := "123"
	if a := MapThenReduce(arr, func(e int) string { return fmt.Sprintf("%d", e) }, "", func(r string, sum *string) { *sum += r }); !reflect.DeepEqual(a, expect) {
		t.Error("test failed!")
	}
}

func TestMapThenGroupBy(t *testing.T) {
	arr := []int{1}
	expect := map[string][]string{"1": {"1"}}
	if a := MapThenGroupBy(arr, func(e int) string { return fmt.Sprintf("%d", e) }, func(str string) (string, string) { return str, str }); !reflect.DeepEqual(a, expect) {
		t.Error("test failed!")
	}
}

func TestMapMap(t *testing.T) {
	m := map[int]string{1: "1"}
	expect := []int{1}
	if val := MapMap(m, func(k int, v string) int { return k }); !reflect.DeepEqual(val, expect) {
		t.Error("test failed!")
	}
}

func TestMapMapThenFilter(t *testing.T) {
	m := map[int]string{1: "1", 2: "2"}
	expect := []int{1}
	if val := MapMapThenFilter(m, func(k int, v string) int { return k }, func(r int) bool { return r == 1 }); !reflect.DeepEqual(val, expect) {
		t.Error("test failed!")
	}
}

func TestMapMapThenReduce(t *testing.T) {
	m := map[int]string{1: "1", 2: "2"}
	expect := 3
	if val := MapMapThenReduce(m, func(k int, v string) int { return k }, 0, reducefunc.Number[int]); !reflect.DeepEqual(val, expect) {
		t.Error("test failed!")
	}
}
