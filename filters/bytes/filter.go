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

package bytes

var (
	// Even filter
	// true if v % 2 == 0
	// false else
	Even = func(v byte) bool { return v%2 == 0 }
	// Odd filter
	// true if v % 2 != 0
	// false else
	Odd = func(v byte) bool { return !Even(v) }

	// Eq filter
	// true if v == eq
	// false else
	Eq = func(eq byte) func(v byte) bool { return func(v byte) bool { return v == eq } }
	// Gt filter
	// true if v > gt
	// false else
	Gt = func(gt byte) func(v byte) bool { return func(v byte) bool { return v > gt } }
	// GtEq filter
	// true if v >= gtEq
	// false else
	GtEq = func(gtEq byte) func(v byte) bool { return func(v byte) bool { return v >= gtEq } }
	// Lt filter
	// true if v < lt
	// false else
	Lt = func(lt byte) func(v byte) bool { return func(v byte) bool { return v < lt } }
	// LtEq filter
	// true if v <= ltEq
	// false else
	LtEq = func(ltEq byte) func(v byte) bool { return func(v byte) bool { return v <= ltEq } }

	// Eq0 filter
	// true if v == 0
	// false else
	Eq0 = func(v byte) bool { return Eq(0)(v) }
	// NotEq0 filter
	// true if v != 0
	// false else
	NotEq0 = func(v byte) bool { return !Eq0(v) }
	// Gt0 filter
	// true if v > 0
	// false else
	Gt0 = func(v byte) bool { return Gt(0)(v) }
	// GtEq0 filter
	// true if v >= 0
	// false else
	GtEq0 = func(v byte) bool { return GtEq(0)(v) }
	// Lt0 filter
	// true if v < 0
	// false else
	Lt0 = func(v byte) bool { return Lt(0)(v) }
	// LtEq0 filter
	// true if v <= 0
	// false else
	LtEq0 = func(v byte) bool { return LtEq(0)(v) }
)
