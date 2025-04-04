package util

import "encoding/json"

func AnyArrayToStructArray[I any](object interface{}) ([]I, error) {
	s := make([]I, 0)
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
