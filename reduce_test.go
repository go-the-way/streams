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

import "testing"

func TestIntReduce(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expect := 10
	if sum := Reduce(arr, 0, func(e int, sum *int) { *sum = *sum + e }); sum != expect {
		t.Error("test failed")
	}
}

func TestStringReduce(t *testing.T) {
	arr := []string{"1", "2", "3", "4"}
	expect := "1234"
	if sum := Reduce(arr, "", func(e string, sum *string) { *sum = *sum + e }); sum != expect {
		t.Error("test failed")
	}
}

func TestReduceMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 2}
	expect := 6
	if sum := ReduceMap(m, 0, func(k, v int, sum *int) { *sum += k + v }); sum != expect {
		t.Error("test failed")
	}
}
