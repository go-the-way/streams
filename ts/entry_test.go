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
	"reflect"
	"testing"
)

func TestNewEntry(t *testing.T) {
	e1 := NewEntry(10, 20)
	expect := &Entry[int, int]{10, 20}
	if !reflect.DeepEqual(e1, expect) {
		t.Fatal("test failed!")
	}
}

func TestMapToEntry(t *testing.T) {
	m := map[int]string{1: "1", 2: "2"}
	expect := []*Entry[int, string]{{1, "1"}, {2, "2"}}
	if !reflect.DeepEqual(MapToEntry[int, string](m), expect) {
		t.Fatal("test failed!")
	}
}

func TestEntryToMap(t *testing.T) {
	es := []*Entry[int, string]{{1, "1"}, {2, "2"}}
	expect := map[int]string{1: "1", 2: "2"}
	if !reflect.DeepEqual(EntryToMap[int, string](es), expect) {
		t.Fatal("test failed!")
	}
}
