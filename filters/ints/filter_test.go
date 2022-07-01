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

package ints

import (
	"testing"
)

func Test(t *testing.T) {
	for _, tc := range []struct {
		name        string
		val, expect any
	}{
		{"Even", Even(10), true},
		{"Odd", Odd(10), false},
		{"Eq", Eq(10)(10), true},
		{"Gt", Gt(10)(10), false},
		{"GtEq", GtEq(10)(10), true},
		{"Lt", Lt(10)(10), false},
		{"LtEq", LtEq(10)(10), true},
		{"Eq0", Eq0(10), false},
		{"NotEq0", NotEq0(10), true},
		{"Gt0", Gt0(10), true},
		{"GtEq0", GtEq0(10), true},
		{"Lt0", Lt0(10), false},
		{"LtEq0", LtEq0(10), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.val != tc.expect {
				t.Error("test failed!")
			}
		})
	}
}
