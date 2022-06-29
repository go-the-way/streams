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

func TestTernary(t *testing.T) {
	var val = 0
	for _, tc := range []struct {
		testName string
		cond     bool
		tF, fF   func()
		except   any
	}{
		{"True", true, func() { t.Log("true"); val = 1 }, func() { t.Log("false"); val = 2 }, 1},
		{"False", false, func() { t.Log("true"); val = 1 }, func() { t.Log("false"); val = 2 }, 2},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			Ternary(tc.cond, tc.tF, tc.fF)
			if !reflect.DeepEqual(val, tc.except) {
				t.Error("test failed!")
			}
		})
	}
}

func TestTernaryV(t *testing.T) {
	for _, tc := range []struct {
		testName string
		cond     bool
		tF, fF   func() int
		except   any
	}{
		{"VTrue", true, func() int { t.Log("true"); return 1 }, func() int { t.Log("false"); return 2 }, 1},
		{"VFalse", false, func() int { t.Log("true"); return 1 }, func() int { t.Log("false"); return 2 }, 2},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			val := TernaryV(tc.cond, tc.tF(), tc.fF())
			if !reflect.DeepEqual(val, tc.except) {
				t.Error("test failed!")
			}
		})
	}
}

func TestIfElse(t *testing.T) {
	val := 0
	for _, tc := range []struct {
		testName            string
		ifF, elseIfF, elseF func() (bool, func())
		except              any
	}{
		{"If", func() (bool, func()) { return true, func() { val = 1 } }, func() (bool, func()) { return false, func() { val = 2 } }, func() (bool, func()) { return false, func() { val = 3 } }, 1},
		{"ElseIf", func() (bool, func()) { return false, func() { val = 1 } }, func() (bool, func()) { return true, func() { val = 2 } }, func() (bool, func()) { return false, func() { val = 3 } }, 2},
		{"Else", func() (bool, func()) { return false, func() { val = 1 } }, func() (bool, func()) { return false, func() { val = 2 } }, func() (bool, func()) { return true, func() { val = 3 } }, 3},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			IfElse(tc.ifF, tc.elseIfF, tc.elseF)
			if !reflect.DeepEqual(val, tc.except) {
				t.Error("test failed!")
			}
		})
	}
}

func TestSwitchCase(t *testing.T) {
	val := 0
	for _, tc := range []struct {
		testName               string
		sw                     int
		case1F, case2F, case3F func() (int, func())
		except                 any
	}{
		{"Default", 0, func() (int, func()) { return 1, func() { val = 1 } }, func() (int, func()) { return 2, func() { val = 2 } }, func() (int, func()) { return 3, func() { val = 3 } }, 0},
		{"Case1", 1, func() (int, func()) { return 1, func() { val = 1 } }, func() (int, func()) { return 2, func() { val = 2 } }, func() (int, func()) { return 3, func() { val = 3 } }, 1},
		{"Case2", 2, func() (int, func()) { return 1, func() { val = 1 } }, func() (int, func()) { return 2, func() { val = 2 } }, func() (int, func()) { return 3, func() { val = 3 } }, 2},
		{"Case3", 3, func() (int, func()) { return 1, func() { val = 1 } }, func() (int, func()) { return 2, func() { val = 2 } }, func() (int, func()) { return 3, func() { val = 3 } }, 3},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			SwitchCase(tc.sw, func() {}, tc.case1F, tc.case2F, tc.case3F)
			if !reflect.DeepEqual(val, tc.except) {
				t.Error("test failed!")
			}
		})
	}
}

func TestSwitchCaseV(t *testing.T) {
	for _, tc := range []struct {
		testName               string
		sw                     int
		case1F, case2F, case3F func() (int, string)
		except                 any
	}{
		{"Default", 0, func() (int, string) { return 1, "1" }, func() (int, string) { return 2, "2" }, func() (int, string) { return 3, "3" }, "0"},
		{"CaseV1", 1, func() (int, string) { return 1, "1" }, func() (int, string) { return 2, "2" }, func() (int, string) { return 3, "3" }, "1"},
		{"CaseV2", 2, func() (int, string) { return 1, "1" }, func() (int, string) { return 2, "2" }, func() (int, string) { return 3, "3" }, "2"},
		{"CaseV3", 3, func() (int, string) { return 1, "1" }, func() (int, string) { return 2, "2" }, func() (int, string) { return 3, "3" }, "3"},
	} {
		t.Run(tc.testName, func(t *testing.T) {
			val := SwitchCaseV(tc.sw, func() string { return "0" }, tc.case1F, tc.case2F, tc.case3F)
			if !reflect.DeepEqual(val, tc.except) {
				t.Error("test failed!")
			}
		})
	}
}
