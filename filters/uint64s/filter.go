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

package ints

var (
	// Even filter
	// true if e % 2 == 0
	// false else
	Even = func(e uint64) bool { return e%2 == 0 }
	// Odd filter
	// true if e % 2 != 0
	// false else
	Odd = func(e uint64) bool { return !Even(e) }

	// Eq filter
	// true if e == eq
	// false else
	Eq = func(eq uint64) func(e uint64) bool { return func(e uint64) bool { return e == eq } }
	// Gt filter
	// true if e > gt
	// false else
	Gt = func(gt uint64) func(e uint64) bool { return func(e uint64) bool { return e > gt } }
	// GtEq filter
	// true if e >= gtEq
	// false else
	GtEq = func(gtEq uint64) func(e uint64) bool { return func(e uint64) bool { return e >= gtEq } }
	// Lt filter
	// true if e < lt
	// false else
	Lt = func(lt uint64) func(e uint64) bool { return func(e uint64) bool { return e < lt } }
	// LtEq filter
	// true if e <= ltEq
	// false else
	LtEq = func(ltEq uint64) func(e uint64) bool { return func(e uint64) bool { return e <= ltEq } }

	// Eq0 filter
	// true if e == 0
	// false else
	Eq0 = func(e uint64) bool { return Eq(0)(e) }
	// NotEq0 filter
	// true if e != 0
	// false else
	NotEq0 = func(e uint64) bool { return !Eq0(e) }
	// Gt0 filter
	// true if e > 0
	// false else
	Gt0 = func(e uint64) bool { return Gt(0)(e) }
	// GtEq0 filter
	// true if e >= 0
	// false else
	GtEq0 = func(e uint64) bool { return GtEq(0)(e) }
	// Lt0 filter
	// true if e < 0
	// false else
	Lt0 = func(e uint64) bool { return Lt(0)(e) }
	// LtEq0 filter
	// true if e <= 0
	// false else
	LtEq0 = func(e uint64) bool { return LtEq(0)(e) }
)
