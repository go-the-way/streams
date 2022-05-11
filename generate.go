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
// E: element type
//
// genFunc: the generate function
func Generate[E any](start int, end int, genFunc func(index, step int) E) []E {
	length := end - start + 1
	es := make([]E, length, length)
	index := 0
	for i := start; i <= end; i++ {
		es[index] = genFunc(index, i)
		index++
	}
	return es
}