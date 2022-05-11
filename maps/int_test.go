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

package maps

import "testing"

func TestInt(t *testing.T) {
	for _, tc := range []struct {
		name   string
		str    string
		except string
	}{
		{"Int2BinaryString", Int2BinaryString(10), "1010"},
		{"Int2OctalString", Int2OctalString(10), "12"},
		{"Int2String", Int2String(15), "15"},
		{"Int2HexString", Int2HexString(15), "f"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.str != tc.except {
				t.Error("test failed")
			}
		})
	}
}
