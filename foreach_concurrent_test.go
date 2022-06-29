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
	"sync/atomic"
	"testing"
)

func TestForEachConcurrent(t *testing.T) {
	arr := []int{0, 1, 2}
	val := int32(0)
	expect := int32(6)
	ForEachConcurrent(arr, func(_, e int) { atomic.AddInt32(&val, int32(e+1)) })
	if !reflect.DeepEqual(val, expect) {
		t.Error("test failed")
	}
}

func TestForEachPtrConcurrent(t *testing.T) {
	arr := []int{0, 1, 2}
	expect := []int{1, 2, 3}
	ForEachPtrConcurrent(arr, func(_ int, e *int) { *e += 1 })
	if !reflect.DeepEqual(arr, expect) {
		t.Error("test failed")
	}
}

func TestMapEachConcurrent(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	expect := map[string]int{"a": 2, "b": 3}
	MapEachConcurrent(m, func(k string, v int) { m[k] = v + 1 })
	if !reflect.DeepEqual(m, expect) {
		t.Error("test failed")
	}
}
