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
	"strings"
	"testing"
)

func TestGroupBy(t *testing.T) {
	type groupBy struct {
		id   string
		name string
	}
	gbs := []*groupBy{
		{"10", "Apple"},
		{"10", "Banana"},
		{"10", "Pear"},
		{"20", "Pear"},
		{"30", "Banana"},
	}
	for _, tc := range []struct {
		name   string
		f      func(e *groupBy) (string, string)
		expect any
	}{
		{"ByID", func(e *groupBy) (string, string) { return e.id, e.name }, map[string][]string{"10": {"Apple", "Banana", "Pear"}, "20": {"Pear"}, "30": {"Banana"}}},
		{"ByName", func(e *groupBy) (string, string) { return e.name, e.id }, map[string][]string{"Apple": {"10"}, "Banana": {"10", "30"}, "Pear": {"10", "20"}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := GroupBy(gbs, tc.f); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed")
			}
		})
	}
}

func TestGroupByThenMap(t *testing.T) {
	type groupByThenMap struct {
		id   string
		name string
	}
	gbs := []*groupByThenMap{
		{"10", "Apple"},
		{"10", "Banana"},
		{"10", "Pear"},
		{"20", "Pear"},
		{"30", "Banana"},
	}
	for _, tc := range []struct {
		name   string
		f      func(e *groupByThenMap) (string, string)
		mf     func(str []string) string
		expect any
	}{
		{"ByID", func(e *groupByThenMap) (string, string) { return e.id, e.name }, func(str []string) string { return strings.Join(str, "") }, map[string]string{"10": "AppleBananaPear", "20": "Pear", "30": "Banana"}},
		{"ByName", func(e *groupByThenMap) (string, string) { return e.name, e.id }, func(str []string) string { return strings.Join(str, "") }, map[string]string{"Apple": "10", "Banana": "1030", "Pear": "1020"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := GroupByThenMap(gbs, tc.f, tc.mf); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed")
			}
		})
	}
}
