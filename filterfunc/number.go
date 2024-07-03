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

package filterfunc

import "github.com/go-the-way/streams/constraint"

// Eq filter
// true if v == eq
// false else
func Eq[T constraint.Number](eq T) func(v T) bool { return func(v T) bool { return v == eq } }

// Gt filter
// true if v > gt
// false else
func Gt[T constraint.Number](gt T) func(v T) bool { return func(v T) bool { return v > gt } }

// GtEq filter
// true if v >= gtEq
// false else
func GtEq[T constraint.Number](gtEq T) func(v T) bool { return func(v T) bool { return v >= gtEq } }

// Lt filter
// true if v < lt
// false else
func Lt[T constraint.Number](lt T) func(v T) bool { return func(v T) bool { return v < lt } }

// LtEq filter
// true if v <= ltEq
// false else
func LtEq[T constraint.Number](ltEq T) func(v T) bool {
	return func(v T) bool { return v <= ltEq }
}
