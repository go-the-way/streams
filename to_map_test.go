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

func TestToMap(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	arr := []*toMap{{10, "Apple"}, {20, "Pear"}}
	except := map[int]string{10: "Apple", 20: "Pear"}
	if m := ToMap(func(e *toMap) (int, string) {
		return e.id, e.name
	}, arr...); !reflect.DeepEqual(m, except) {
		t.Error("test failed")
	}
}

func TestToMap2(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	arr := []*toMap{{10, "Apple"}, {20, "Pear"}}
	except1 := map[int]string{10: "Apple", 20: "Pear"}
	except2 := map[string]int{"Apple": 10, "Pear": 20}
	if m1, m2 := ToMap2(func(e *toMap) (int, string) {
		return e.id, e.name
	}, func(e *toMap) (string, int) {
		return e.name, e.id
	}, arr...); !reflect.DeepEqual(m1, except1) || !reflect.DeepEqual(m2, except2) {
		t.Error("test failed")
	}
}
