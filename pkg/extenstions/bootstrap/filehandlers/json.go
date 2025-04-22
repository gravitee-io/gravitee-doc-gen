package filehandlers

import (
	"encoding/json"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"os"
)

const JsonExt = ".json"

func JsonFileHandler(jsonFile string) (any, error) {
	bytes, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	data := &util.Unstructured{}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
