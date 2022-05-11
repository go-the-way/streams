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

func TestEach(t *testing.T) {
	arr := []int{0, 1, 2}
	except := []int{1, 2, 3}
	newArr := make([]int, 0)
	Each(func(e int) { newArr = append(newArr, e+1) }, arr...)
	if !reflect.DeepEqual(newArr, except) {
		t.Error("test failed")
	}
}

func TestEachPtr(t *testing.T) {
	arr := []int{0, 1, 2}
	except := []int{1, 2, 3}
	EachPtr(func(e *int) { *e += 1 }, arr...)
	if !reflect.DeepEqual(arr, except) {
		t.Error("test failed")
	}
}
