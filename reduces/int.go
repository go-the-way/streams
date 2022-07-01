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
	Int         = func(e int, sum *int) { *sum += e }                             // Int reduce
	Int8        = func(e int8, sum *int8) { *sum += e }                           // Int8 reduce
	Int16       = func(e int16, sum *int16) { *sum += e }                         // Int16 reduce
	Int32       = func(e int32, sum *int32) { *sum += e }                         // Int32 reduce
	Int64       = func(e int64, sum *int64) { *sum += e }                         // Int64 reduce
	Uint        = func(e uint, sum *uint) { *sum += e }                           // Uint reduce
	Uint8       = func(e uint8, sum *uint8) { *sum += e }                         // Uint8 reduce
	Uint16      = func(e uint16, sum *uint16) { *sum += e }                       // Uint16 reduce
	Uint32      = func(e uint32, sum *uint32) { *sum += e }                       // Uint32 reduce
	Uint64      = func(e uint64, sum *uint64) { *sum += e }                       // Uint64 reduce
	SliceInt    = func(ts []int, sum *[]int) { *sum = append(*sum, ts...) }       // SliceInt reduce
	SliceInt8   = func(ts []int8, sum *[]int8) { *sum = append(*sum, ts...) }     // SliceInt8 reduce
	SliceInt16  = func(ts []int16, sum *[]int16) { *sum = append(*sum, ts...) }   // SliceInt16 reduce
	SliceInt32  = func(ts []int32, sum *[]int32) { *sum = append(*sum, ts...) }   // SliceInt32 reduce
	SliceInt64  = func(ts []int64, sum *[]int64) { *sum = append(*sum, ts...) }   // SliceInt64 reduce
	SliceUint   = func(ts []uint, sum *[]uint) { *sum = append(*sum, ts...) }     // SliceUint reduce
	SliceUint8  = func(ts []uint8, sum *[]uint8) { *sum = append(*sum, ts...) }   // SliceUint8 reduce
	SliceUint16 = func(ts []uint16, sum *[]uint16) { *sum = append(*sum, ts...) } // SliceUint16 reduce
	SliceUint32 = func(ts []uint32, sum *[]uint32) { *sum = append(*sum, ts...) } // SliceUint32 reduce
	SliceUint64 = func(ts []uint64, sum *[]uint64) { *sum = append(*sum, ts...) } // SliceUint64 reduce
)
