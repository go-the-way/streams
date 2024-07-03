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

package reducefunc

import "testing"

func TestNumber(t *testing.T) {
	var num int
	Number(1, &num)
	if num != 1 {
		t.Error("test failed!")
	}
}

func TestNumberFunc(t *testing.T) {
	var num int
	NumberFunc[int]()(1, &num)
	if num != 1 {
		t.Error("test failed!")
	}
}

func TestSlice(t *testing.T) {
	var ns []int
	Slice([]int{1}, &ns)
	if ns[0] != 1 {
		t.Error("test failed!")
	}
}

func TestSliceFunc(t *testing.T) {
	var ns []int
	SliceFunc[int]()([]int{1}, &ns)
	if ns[0] != 1 {
		t.Error("test failed!")
	}
}
