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

import "reflect"

// Ternary
// cond ? trueFunc() : falseFunc()
func Ternary(cond bool, trueFunc, falseFunc func()) {
	if cond {
		if fn := trueFunc; fn != nil {
			fn()
		}
	} else {
		if fn := falseFunc; fn != nil {
			fn()
		}
	}
}

// TernaryV
// cond ? trueFunc() T : falseFunc() T
func TernaryV[T any](cond bool, trueVal, falseVal T) T {
	if cond {
		return trueVal
	} else {
		return falseVal
	}
}

// If function
func If(condFunc func() bool, ifFunc func()) {
	if fn := condFunc; fn != nil {
		if fn() {
			if fn := ifFunc; fn != nil {
				fn()
			}
		}
	}
}

// IfV function
func IfV(cond bool, ifFunc func()) {
	if cond {
		if fn := ifFunc; fn != nil {
			fn()
		}
	}
}

// IfElse function
func IfElse(condFunc ...func() (bool, func())) {
	if condFunc != nil && len(condFunc) > 0 {
		for _, cF := range condFunc {
			if cfC, cfF := cF(); cfC && cfF != nil {
				cfF()
				return
			}
		}
	}
}

// SwitchCase function
//
// SW: switch type
//
// defaultFunc: the default function
//
// caseFunc: the case function
func SwitchCase[SW any](sw SW, defaultFunc func(), caseFunc ...func() (SW, func())) {
	if caseFunc != nil && len(caseFunc) > 0 {
		for _, cF := range caseFunc {
			if cfw, cfF := cF(); reflect.DeepEqual(sw, cfw) && cfF != nil {
				cfF()
				return
			}
		}
	}
	if fn := defaultFunc; fn != nil {
		fn()
	}
}

// SwitchCaseV function
//
// SW: switch type
//
// T: the val type
//
// defaultT: the default T
//
// caseFunc: the case function
func SwitchCaseV[SW any, T any](sw SW, defaultT T, caseFunc ...func() (SW, T)) T {
	if caseFunc != nil && len(caseFunc) > 0 {
		for _, cF := range caseFunc {
			if cfw, cfT := cF(); reflect.DeepEqual(sw, cfw) {
				return cfT
			}
		}
	}
	return defaultT
}
