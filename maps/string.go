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
	string2Bool    = func(str string) (b bool) { b, _ = strconv.ParseBool(str); return }
	string2Int     = func(str string) int { i, _ := strconv.ParseInt(str, 10, 32); return int(i) }
	string2Byte    = func(str string) byte { i, _ := strconv.ParseInt(str, 10, 8); return byte(i) }
	string2Int8    = func(str string) int8 { i, _ := strconv.ParseInt(str, 10, 8); return int8(i) }
	string2Int16   = func(str string) int16 { i, _ := strconv.ParseInt(str, 10, 16); return int16(i) }
	string2Int32   = func(str string) int32 { i, _ := strconv.ParseInt(str, 10, 32); return int32(i) }
	string2Int64   = func(str string) (i int64) { i, _ = strconv.ParseInt(str, 10, 64); return }
	string2UInt    = func(str string) uint { i, _ := strconv.ParseUint(str, 10, 32); return uint(i) }
	string2UInt8   = func(str string) uint8 { i, _ := strconv.ParseUint(str, 10, 8); return uint8(i) }
	string2UInt16  = func(str string) uint16 { i, _ := strconv.ParseUint(str, 10, 16); return uint16(i) }
	string2UInt32  = func(str string) uint32 { i, _ := strconv.ParseUint(str, 10, 32); return uint32(i) }
	string2UInt64  = func(str string) (i uint64) { i, _ = strconv.ParseUint(str, 10, 64); return }
	string2Float32 = func(str string) float32 { f, _ := strconv.ParseFloat(str, 32); return float32(f) }
	string2Float64 = func(str string) (f float64) { f, _ = strconv.ParseFloat(str, 64); return }
)

// String2Bool
// T -> true, F -> false
func String2Bool(str string) (b bool) { return Any2Any(str, string2Bool) }

// String2Int
// 10 -> 10
func String2Int(str string) int { return Any2Any(str, string2Int) }

// String2Byte
// 10 -> byte(10)
func String2Byte(str string) byte { return Any2Any(str, string2Byte) }

// String2Int8
// 10 -> int8(10)
func String2Int8(str string) int8 { return Any2Any(str, string2Int8) }

// String2Int16
// 10 -> int16(10)
func String2Int16(str string) int16 { return Any2Any(str, string2Int16) }

// String2Int32
// 10 -> int32(10)
func String2Int32(str string) int32 { return Any2Any(str, string2Int32) }

// String2Int64
// 10 -> int64(10)
func String2Int64(str string) (i int64) { return Any2Any(str, string2Int64) }

// String2UInt
// 10 -> uint(10)
func String2UInt(str string) uint { return Any2Any(str, string2UInt) }

// String2UInt8
// 10 -> uint8(10)
func String2UInt8(str string) uint8 { return Any2Any(str, string2UInt8) }

// String2UInt16
// 10 -> uint16(10)
func String2UInt16(str string) uint16 { return Any2Any(str, string2UInt16) }

// String2UInt32
// 10 -> uint32(10)
func String2UInt32(str string) uint32 { return Any2Any(str, string2UInt32) }

// String2UInt64
// 10 -> uint64(10)
func String2UInt64(str string) (i uint64) { return Any2Any(str, string2UInt64) }

// String2Float32
// 10.25 -> float32(10.25)
func String2Float32(str string) float32 { return Any2Any(str, string2Float32) }

// String2Float64
// 10.25 -> 10.25
func String2Float64(str string) (f float64) { return Any2Any(str, string2Float64) }
