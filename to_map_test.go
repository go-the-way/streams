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

import "testing"

func TestToMap(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	t.Log(ToMap(func(e *toMap) (int, string) {
		return e.id, e.name
	}, &toMap{10, "Apple"}, &toMap{20, "Pear"}))
}

func TestToMap2(t *testing.T) {
	type toMap struct {
		id   int
		name string
	}
	t.Log(ToMap2(func(e *toMap) (int, string) {
		return e.id, e.name
	}, func(e *toMap) (string, int) {
		return e.name, e.id
	}, &toMap{10, "Apple"}, &toMap{20, "Pear"}))
}
