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

func TestGenerate(t *testing.T) {
	expect := []int{2, 4, 6, 8}
	es := Generate(1, 4, func(index, step int) int { return step * 2 })
	if !reflect.DeepEqual(expect, es) {
		t.Error("test failed")
	}
}
