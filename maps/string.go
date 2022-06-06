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

import "strconv"

var (
	// String2Bool
	// T -> true, F -> false
	String2Bool = func(str string) (b bool) { b, _ = strconv.ParseBool(str); return }

	// String2Int
	// 10 -> 10
	String2Int = func(str string) int { i, _ := strconv.ParseInt(str, 10, 32); return int(i) }

	// String2Byte
	// 10 -> byte(10)
	String2Byte = func(str string) byte { i, _ := strconv.ParseInt(str, 10, 8); return byte(i) }

	// String2Int8
	// 10 -> int8(10)
	String2Int8 = func(str string) int8 { i, _ := strconv.ParseInt(str, 10, 8); return int8(i) }

	// String2Int16
	// 10 -> int16(10)
	String2Int16 = func(str string) int16 { i, _ := strconv.ParseInt(str, 10, 16); return int16(i) }

	// String2Int32
	// 10 -> int32(10)
	String2Int32 = func(str string) int32 { i, _ := strconv.ParseInt(str, 10, 32); return int32(i) }

	// String2Int64
	// 10 -> int64(10)
	String2Int64 = func(str string) (i int64) { i, _ = strconv.ParseInt(str, 10, 64); return }

	// String2UInt
	// 10 -> uint(10)
	String2UInt = func(str string) uint { i, _ := strconv.ParseUint(str, 10, 32); return uint(i) }

	// String2UInt8
	// 10 -> uint8(10)
	String2UInt8 = func(str string) uint8 { i, _ := strconv.ParseUint(str, 10, 8); return uint8(i) }

	// String2UInt16
	// 10 -> uint16(10)
	String2UInt16 = func(str string) uint16 { i, _ := strconv.ParseUint(str, 10, 16); return uint16(i) }

	// String2UInt32
	// 10 -> uint32(10)
	String2UInt32 = func(str string) uint32 { i, _ := strconv.ParseUint(str, 10, 32); return uint32(i) }

	// String2UInt64
	// 10 -> uint64(10)
	String2UInt64 = func(str string) (i uint64) { i, _ = strconv.ParseUint(str, 10, 64); return }

	// String2Float32
	// 10.25 -> float32(10.25)
	String2Float32 = func(str string) float32 { f, _ := strconv.ParseFloat(str, 32); return float32(f) }

	// String2Float64
	// 10.25 -> 10.25
	String2Float64 = func(str string) (f float64) { f, _ = strconv.ParseFloat(str, 64); return }
)
