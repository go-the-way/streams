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

func TestSlice(t *testing.T) {
	for _, tc := range []struct {
		name        string
		val, expect any
	}{
		{"ReduceString", ReduceString([]string{"1", "2", "3"}), "123"},
		{"ReduceStringWithComma", ReduceStringWithComma([]string{"1", "2", "3"}), "1,2,3"},
		{"ReduceStringWithDot", ReduceStringWithDot([]string{"1", "2", "3"}), "1.2.3"},
		{"ReduceStringWithSpace", ReduceStringWithSpace([]string{"1", "2", "3"}), "1 2 3"},
		{"ReduceStringWithLine", ReduceStringWithLine([]string{"1", "2", "3"}), "1\n2\n3"},
		{"ReduceInt8", ReduceInt8([]int8{1, 2, 3}), int8(6)},
		{"ReduceInt16", ReduceInt16([]int16{1, 2, 3}), int16(6)},
		{"ReduceInt32", ReduceInt32([]int32{1, 2, 3}), int32(6)},
		{"ReduceInt", ReduceInt([]int{1, 2, 3}), 6},
		{"ReduceInt64", ReduceInt64([]int64{1, 2, 3}), int64(6)},
		{"ReduceUint8", ReduceUint8([]uint8{1, 2, 3}), uint8(6)},
		{"ReduceUint16", ReduceUint16([]uint16{1, 2, 3}), uint16(6)},
		{"ReduceUint32", ReduceUint32([]uint32{1, 2, 3}), uint32(6)},
		{"ReduceUint", ReduceUint([]uint{1, 2, 3}), uint(6)},
		{"ReduceUint64", ReduceUint64([]uint64{1, 2, 3}), uint64(6)},
		{"ReduceFloat32", ReduceFloat32([]float32{1, 2, 3}), float32(6)},
		{"ReduceFloat64", ReduceFloat64([]float64{1, 2, 3}), float64(6)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.val != tc.expect {
				t.Error("test failed!")
			}
		})
	}
}
