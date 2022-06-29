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
	"github.com/go-the-way/streams"
	"github.com/go-the-way/streams/reduces"
	"strings"
)

var (
	reduceString = func(str []string, sep string) string { return strings.Join(str, sep) }
	// ReduceString
	// []String{1, 2, 3} => "123"
	ReduceString = func(str []string) string { return reduceString(str, "") }
	// ReduceStringWithComma
	// []String{1, 2, 3} => "1,2,3"
	ReduceStringWithComma = func(str []string) string { return reduceString(str, ",") }
	// ReduceStringWithDot
	// []String{1, 2, 3} => "1.2.3"
	ReduceStringWithDot = func(str []string) string { return reduceString(str, ".") }
	// ReduceStringWithSpace
	// []String{1, 2, 3} => "1 2 3"
	ReduceStringWithSpace = func(str []string) string { return reduceString(str, " ") }
	// ReduceStringWithLine
	// []String{1, 2, 3} => "1\n2\n3"
	ReduceStringWithLine = func(str []string) string { return reduceString(str, "\n") }

	// ReduceInt8
	// []int8{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceInt8 = func(i []int8) int8 { return streams.Reduce(i, 0, reduces.Int8) }
	// ReduceInt16
	// []int8{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceInt16 = func(i []int16) int16 { return streams.Reduce(i, 0, reduces.Int16) }
	// ReduceInt32
	// []int32{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceInt32 = func(i []int32) int32 { return streams.Reduce(i, 0, reduces.Int32) }
	// ReduceInt
	// []int{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceInt = func(i []int) int { return streams.Reduce(i, 0, reduces.Int) }
	// ReduceInt64
	// []int64{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceInt64 = func(i []int64) int64 { return streams.Reduce(i, 0, reduces.Int64) }

	// ReduceUint8
	// []uint8{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceUint8 = func(i []uint8) uint8 { return streams.Reduce(i, 0, reduces.Uint8) }
	// ReduceUint16
	// []uint8{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceUint16 = func(i []uint16) uint16 { return streams.Reduce(i, 0, reduces.Uint16) }
	// ReduceUint32
	// []uint32{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceUint32 = func(i []uint32) uint32 { return streams.Reduce(i, 0, reduces.Uint32) }
	// ReduceUint
	// []uint{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceUint = func(i []uint) uint { return streams.Reduce(i, 0, reduces.Uint) }
	// ReduceUint64
	// []uint64{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceUint64 = func(i []uint64) uint64 { return streams.Reduce(i, 0, reduces.Uint64) }

	// ReduceFloat32
	// []float32{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceFloat32 = func(f []float32) float32 { return streams.Reduce(f, 0, reduces.Float32) }
	// ReduceFloat64
	// []float64{1, 2, 3} => 1 + 2 + 3 => 6
	ReduceFloat64 = func(f []float64) float64 { return streams.Reduce(f, 0, reduces.Float64) }
)
