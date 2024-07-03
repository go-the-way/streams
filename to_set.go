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

// ToSet function
//
// T: element type
//
// R: result type
//
// toSetFunc: the to set function
func ToSet[T any, R comparable](ts []T, toSetFunc func(e T) R) Set[R] {
	s := NewSet[R]()
	for _, t := range ts {
		s.Add(toSetFunc(t))
	}
	return s
}

// ToSetThenMap function
//
// T: element type
//
// R: result type
//
// RR: map type
//
// toSetFunc: the to map function
//
// mapFunc: the map function
func ToSetThenMap[T any, R comparable, RR any](ts []T, toSetFunc func(t T) R, mapFunc func(r R) RR) []RR {
	rs := make([]RR, 0)
	s := NewSet[R]()
	for _, t := range ts {
		s.Add(toSetFunc(t))
	}
	s.Iterate(func(r R) {
		rs = append(rs, mapFunc(r))
	})
	return rs
}

// ToSet2 function
//
// T: element type
//
// R1: the set1 type
//
// R2: the set2 type
//
// toSetFunc1: the to set1 function
//
// toSetFunc2: the to set2 function
func ToSet2[T any, R1, R2 comparable](ts []T, toSetFunc1 func(e T) R1, toSetFunc2 func(e T) R2) (Set[R1], Set[R2]) {
	s1 := NewSet[R1]()
	s2 := NewSet[R2]()
	for _, t := range ts {
		s1.Add(toSetFunc1(t))
		s2.Add(toSetFunc2(t))
	}
	return s1, s2
}

// ToSet3 function
//
// T: element type
//
// R1: the set1 type
//
// R2: the set2 type
//
// R3: the set3 type
//
// toSetFunc1: the to set1 function
//
// toSetFunc2: the to set2 function
//
// toSetFunc3: the to set3 function
func ToSet3[T any, R1, R2, R3 comparable](ts []T, toSetFunc1 func(t T) R1, toSetFunc2 func(t T) R2, toSetFunc3 func(t T) R3) (Set[R1], Set[R2], Set[R3]) {
	s1 := NewSet[R1]()
	s2 := NewSet[R2]()
	s3 := NewSet[R3]()
	for _, t := range ts {
		s1.Add(toSetFunc1(t))
		s2.Add(toSetFunc2(t))
		s3.Add(toSetFunc3(t))
	}
	return s1, s2, s3
}
