package examples

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
	"os"
)

type ExampleSpecProvider interface {
	ExampleSpecs() []ExampleSpec
	SetConfigData(ConfigData)
	GetConfigData() ConfigData
}

func LoadConfig(chunk config.Chunk, provider ExampleSpecProvider) error {
	examplesFile := common.GetFile(chunk, "examples")
	bytes, err := os.ReadFile(examplesFile)
	if err != nil {
		return err
	}
	cfg := &ConfigData{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return err
	}

	provider.SetConfigData(*cfg)
	return nil
}

func ValidateSpecs(specs []ExampleSpec, chunk config.Chunk) (bool, error) {
	for _, spec := range specs {

		err := spec.Validate()
		if err != nil {
			return false, err
		}

		schemaFile, _, err := CompileSchema(spec, chunk)
		if err != nil {
			return false, errors.New(fmt.Sprintf("failed to compile schema %s, %v", schemaFile, err))
		}
	}
	return false, nil
}

func CompileSchema(spec ExampleSpec, chunk config.Chunk) (*jsonschema.Schema, string, error) {
	schemaFile := common.GetFile(chunk, "schema")
	if spec.GetOverrideSchema() != "" {
		schemaFile = spec.GetOverrideSchema()
	}
	compiled, err := schema.CompileWithExtensions(schemaFile)
	return compiled, schemaFile, err
}
