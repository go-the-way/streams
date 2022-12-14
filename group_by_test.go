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
				t.Error("test failed!")
			}
		})
	}
}

func TestGroupByThenToMap(t *testing.T) {
	type groupByThenToMap struct {
		id   string
		name string
	}
	gbs := []*groupByThenToMap{
		{"10", "Apple"},
		{"10", "Banana"},
		{"10", "Pear"},
		{"20", "Pear"},
		{"30", "Banana"},
	}
	for _, tc := range []struct {
		name   string
		f      func(e *groupByThenToMap) (string, string)
		mf     func(s string, str []string) (string, string)
		expect any
	}{
		{"ByID", func(e *groupByThenToMap) (string, string) { return e.id, e.name }, func(s string, str []string) (string, string) { return s, strings.Join(str, "") }, map[string]string{"10": "AppleBananaPear", "20": "Pear", "30": "Banana"}},
		{"ByName", func(e *groupByThenToMap) (string, string) { return e.name, e.id }, func(s string, str []string) (string, string) { return s, strings.Join(str, "") }, map[string]string{"Apple": "10", "Banana": "1030", "Pear": "1020"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := GroupByThenToMap(gbs, tc.f, tc.mf); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
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
		{"10", ""},
	}
	for _, tc := range []struct {
		name            string
		f               func(e *groupByThenMap) (string, string)
		mf              func(s string, str []string) string
		expect, expect2 any
	}{
		{"ByID", func(e *groupByThenMap) (string, string) { return e.id, e.name }, func(s string, str []string) string { return s + strings.Join(str, "") }, []string{"10Apple"}, []string{"10Apple"}},
		{"ByName", func(e *groupByThenMap) (string, string) { return e.name, e.id }, func(s string, str []string) string { return s + strings.Join(str, "") }, []string{"Apple10", "10"}, []string{"10", "Apple10"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := GroupByThenMap(gbs, tc.f, tc.mf); !reflect.DeepEqual(fArr, tc.expect) && !reflect.DeepEqual(fArr, tc.expect2) {
				t.Error("test failed!")
			}
		})
	}
}

func TestGroupByCounting(t *testing.T) {
	type groupByCounting struct {
		id   string
		name string
	}
	gbs := []*groupByCounting{
		{"10", "Apple"},
	}
	for _, tc := range []struct {
		name   string
		f      func(e *groupByCounting) string
		expect any
	}{
		{"ByID", func(e *groupByCounting) string { return e.id }, map[string]int{"10": 1}},
		{"ByName", func(e *groupByCounting) string { return e.name }, map[string]int{"Apple": 1}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := GroupByCounting(gbs, tc.f); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}

func TestGroupByCountingThenMap(t *testing.T) {
	type groupByCountingThenMap struct {
		id   string
		name string
	}
	gbs := []*groupByCountingThenMap{
		{"10", "Apple"},
	}
	for _, tc := range []struct {
		name   string
		f      func(e *groupByCountingThenMap) string
		mf     func(k string, v int) string
		expect any
	}{
		{"ByID", func(e *groupByCountingThenMap) string { return e.id }, func(k string, v int) string { return fmt.Sprintf("%s%d", k, v) }, []string{"101"}},
		{"ByName", func(e *groupByCountingThenMap) string { return e.name }, func(k string, v int) string { return fmt.Sprintf("%s%d", k, v) }, []string{"Apple1"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := GroupByCountingThenMap(gbs, tc.f, tc.mf); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}

func TestGroupByCountingThenToMap(t *testing.T) {
	type groupByCountingThenToMap struct {
		id   string
		name string
	}
	gbs := []*groupByCountingThenToMap{
		{"10", "Apple"},
	}
	for _, tc := range []struct {
		name   string
		f      func(e *groupByCountingThenToMap) string
		mf     func(k string, v int) (int, string)
		expect any
	}{
		{"ByID", func(e *groupByCountingThenToMap) string { return e.id }, func(k string, v int) (int, string) { return v, k }, map[int]string{1: "10"}},
		{"ByName", func(e *groupByCountingThenToMap) string { return e.name }, func(k string, v int) (int, string) { return v, k }, map[int]string{1: "Apple"}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if fArr := GroupByCountingThenToMap(gbs, tc.f, tc.mf); !reflect.DeepEqual(fArr, tc.expect) {
				t.Error("test failed!")
			}
		})
	}
}
