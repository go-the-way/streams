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

package reduces

var (
	Int         = func(e int, sum *int) { *sum += e }
	Int8        = func(e int8, sum *int8) { *sum += e }
	Int16       = func(e int16, sum *int16) { *sum += e }
	Int32       = func(e int32, sum *int32) { *sum += e }
	Int64       = func(e int64, sum *int64) { *sum += e }
	Uint        = func(e uint, sum *uint) { *sum += e }
	Uint8       = func(e uint8, sum *uint8) { *sum += e }
	Uint16      = func(e uint16, sum *uint16) { *sum += e }
	Uint32      = func(e uint32, sum *uint32) { *sum += e }
	Uint64      = func(e uint64, sum *uint64) { *sum += e }
	SliceInt    = func(es []int, sum *[]int) { *sum = append(*sum, es...) }
	SliceInt8   = func(es []int8, sum *[]int8) { *sum = append(*sum, es...) }
	SliceInt16  = func(es []int16, sum *[]int16) { *sum = append(*sum, es...) }
	SliceInt32  = func(es []int32, sum *[]int32) { *sum = append(*sum, es...) }
	SliceInt64  = func(es []int64, sum *[]int64) { *sum = append(*sum, es...) }
	SliceUint   = func(es []uint, sum *[]uint) { *sum = append(*sum, es...) }
	SliceUint8  = func(es []uint8, sum *[]uint8) { *sum = append(*sum, es...) }
	SliceUint16 = func(es []uint16, sum *[]uint16) { *sum = append(*sum, es...) }
	SliceUint32 = func(es []uint32, sum *[]uint32) { *sum = append(*sum, es...) }
	SliceUint64 = func(es []uint64, sum *[]uint64) { *sum = append(*sum, es...) }
)
