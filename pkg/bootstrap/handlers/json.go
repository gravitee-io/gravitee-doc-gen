package handlers

import (
	"encoding/json"
	"github.com/gravitee-io-labs/readme-gen/pkg/core"
	"os"
)

const JsonExt = ".jon"

func JsonFileHandler(jsonFile string) (any, error) {
	bytes, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	data := &core.Unstructured{}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
