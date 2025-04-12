package util

import (
	"encoding/json"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/core"
)

func AnyArrayToStructArray[I any](array []any) ([]I, error) {
	a := make([]I, 0)
	bytes, err := json.Marshal(array)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func AnyMapToStruct[T any](object *core.Unstructured) (*T, error) {
	s := new(T)
	bytes, err := json.Marshal(object)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func AnyToString(v any) string {
	return fmt.Sprintf("%v", v)
}
