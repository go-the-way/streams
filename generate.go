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

package streams

// Generate function
//
// T: element type
//
// genFunc: the generate function
func Generate[T any](start, end int, genFunc func(index, step int) T) []T {
	length := end - start + 1
	ts := make([]T, length)
	index := 0
	for i := start; i <= end; i++ {
		ts[index] = genFunc(index, i)
		index++
	}
	return ts
}
