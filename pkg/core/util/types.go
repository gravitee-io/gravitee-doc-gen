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

import "slices"

type Unstructured map[string]interface{}

// Set a slices with unique values.
type set struct {
	items []any
}

type Set interface {
	// Add a value to the set if not in it already
	Add(v any)
	// Items turn a Set into a slice by copying values
	Items() []any
	Contains(v any) bool
}

func NewSet() Set {
	return &set{}
}

func (s *set) Add(v any) {
	if !slices.Contains(s.items, v) {
		s.items = append(s.items, v)
	}
}

func (s *set) Items() []any {
	slice := make([]any, len(s.items))
	copy(slice, s.items)
	return slice
}

func (s *set) Contains(v any) bool {
	return slices.Contains(s.items, v)
}

// ToSlice convert an untyped Set into a typed slice.
func ToSlice[T any](s Set) []T {
	impl, ok := s.(*set)
	if !ok {
		panic("not a set")
	}
	slice := make([]T, 0, len(impl.items))
	for _, v := range impl.items {
		slice = append(slice, As[T](v))
	}
	return slice
}

// As types an untyped value or panics.
func As[T any](x any) T {
	it, ok := x.(T)
	if !ok {
		panic("invalid type assertion")
	}
	return it
}
