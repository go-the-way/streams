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

import (
	"testing"
)

func TestString(t *testing.T) {
	for _, tc := range []struct {
		name        string
		str, expect any
	}{
		{"String2Bool", String2Bool("T"), true},
		{"String2Int", String2Int("10"), 10},
		{"String2Byte", String2Byte("10"), byte(10)},
		{"String2Int8", String2Int8("10"), int8(10)},
		{"String2Int16", String2Int16("10"), int16(10)},
		{"String2Int32", String2Int32("10"), int32(10)},
		{"String2Int64", String2Int64("10"), int64(10)},
		{"String2UInt", String2UInt("10"), uint(10)},
		{"String2UInt8", String2UInt8("10"), uint8(10)},
		{"String2UInt16", String2UInt16("10"), uint16(10)},
		{"String2UInt32", String2UInt32("10"), uint32(10)},
		{"String2UInt64", String2UInt64("10"), uint64(10)},
		{"String2Float32", String2Float32("10.25"), float32(10.25)},
		{"String2Float64", String2Float64("10.25"), 10.25},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.str != tc.expect {
				t.Error("test failed")
			}
		})
	}
}
