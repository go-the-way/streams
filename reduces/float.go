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
	Float32      = func(e float32, sum *float32) { *sum += e }                       // Float32 reduce
	Float64      = func(e float64, sum *float64) { *sum += e }                       // Float64 reduce
	SliceFloat32 = func(ts []float32, sum *[]float32) { *sum = append(*sum, ts...) } // SliceFloat32 reduce
	SliceFloat64 = func(ts []float64, sum *[]float64) { *sum = append(*sum, ts...) } // SliceFloat64 reduce
)
