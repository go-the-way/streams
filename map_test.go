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
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	expect := []string{"1", "2", "3"}
	if a := Map(func(e int) string { return fmt.Sprintf("%d", e) }, arr...); !reflect.DeepEqual(a, expect) {
		t.Error("test failed")
	}
}

func TestMapMap(t *testing.T) {
	m := map[string]int{"1": 1, "2": 2}
	expect := map[string]int{"1": 2, "2": 3}
	if a := MapMap(func(k string, v int) (string, int) { return k, v + 1 }, m); !reflect.DeepEqual(a, expect) {
		t.Error("test failed")
	}
}
