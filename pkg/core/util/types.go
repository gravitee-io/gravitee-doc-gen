// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

type Unstructured map[string]interface{}
type Set map[any]bool

func (s Set) Add(v any) {
	s[v] = true
}

func (s Set) ToSlice() []any {
	slice := make([]any, 0, len(s))
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}

func ToSlice[T any](s Set) []T {
	slice := make([]T, 0, len(s))
	for v := range s {
		slice = append(slice, As[T](v))
	}
	return slice
}

func As[T any](x any) T {
	it, ok := x.(T)
	if !ok {
		panic("invalid type assertion")
	}
	return it
}
