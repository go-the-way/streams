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

// Reduce function
//
// ID: identity value type
//
// E: element type
//
// accumulatorFunc: the accumulator function
func Reduce[ID, E any](accumulatorFunc func(e E, sum *ID), id ID, es ...E) ID {
	for _, e := range es {
		accumulatorFunc(e, &id)
	}
	return id
}
