package filehandlers

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"gopkg.in/yaml.v3"
	"os"
)

const YamlExt = ".yaml"
const YmlExt = ".yml"

func YamlFileHandler(yamlFile string) (any, error) {
	bytes, err := os.ReadFile(yamlFile)
	if err != nil {
		return nil, err
	}

	data := &util.Unstructured{}
	err = yaml.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
