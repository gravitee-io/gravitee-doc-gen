package examples

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func TypeValidator(chunk config.Chunk, provider ExampleSpecProvider) (bool, error) {
	tmplExists, err := common.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return tmplExists, err
	}

	examplesFile := common.GetString(chunk, "examples")
	examplesFileExists := common.FileExists(examplesFile)

	if chunk.Required && !examplesFileExists {
		return examplesFileExists, errors.New(fmt.Sprintf("example file not found: %s", examplesFile))
	}

	schemaFile := common.GetString(chunk, "schema")
	schemaFileExists := common.FileExists(schemaFile)
	if chunk.Required && !schemaFileExists {
		return schemaFileExists, errors.New(fmt.Sprintf("schema file not found: %s", schemaFile))
	}

	err = LoadConfig(chunk, provider)
	if err != nil {
		return true, err
	}

	b, err := ValidateSpecs(provider.ExampleSpecs(), chunk)
	if err != nil {
		return b, err
	}

	return tmplExists && schemaFileExists, nil
}

func ValidateJson(jsonToValidate string, validationSchema *jsonschema.Schema, origin string) error {
	payload := make(map[string]any)
	err := json.Unmarshal([]byte(jsonToValidate), &payload)
	if err != nil {
		return err
	}

	err = validationSchema.Validate(payload)
	if err != nil {
		return errors.New(fmt.Sprintf("schema validation error: [%s] could not be validated:\n\t%s\n%s", origin, err, jsonToValidate))
	}
	return nil

}
