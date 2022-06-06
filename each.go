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

// Each function
//
// E: element type
//
// eachFunc: the each function
func Each[E any](eachFunc func(i int, e E), es ...E) {
	for i, e := range es {
		eachFunc(i, e)
	}
}

// EachPtr function
//
// E: element type
//
// eachFunc: the each function
func EachPtr[E any](eachFunc func(i int, e *E), es ...E) {
	for i := range es {
		eachFunc(i, &es[i])
	}
}
