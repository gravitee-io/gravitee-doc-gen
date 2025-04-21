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
		slice = append(slice, v.(T))
	}
	return slice
}
