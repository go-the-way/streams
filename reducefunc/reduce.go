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

package reducefunc

import "github.com/go-the-way/streams/constraint"

type Nr = constraint.Number

func Number[T Nr](num T, sum *T)              { *sum += num }
func NumberFunc[T Nr]() func(num T, sum *T)   { return func(num T, sum *T) { Number(num, sum) } }
func Slice[T Nr](ns []T, sum *[]T)            { *sum = append(*sum, ns...) }
func SliceFunc[T Nr]() func(ns []T, sum *[]T) { return func(ns []T, sum *[]T) { Slice(ns, sum) } }
