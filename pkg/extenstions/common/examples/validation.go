package examples

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func TypeValidator(chunk config.Chunk, provider ExampleSpecProvider) (bool, error) {
	tmplExists, err := chunks.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return tmplExists, err
	}

	examplesFile := chunks.GetString(chunk, "examples")
	examplesFileExists := util.FileExists(examplesFile)

	if chunk.Required && !examplesFileExists {
		return examplesFileExists, errors.New(fmt.Sprintf("example file not found: %s", examplesFile))
	}

	schemaFile := chunks.GetString(chunk, "schema")
	schemaFileExists := util.FileExists(schemaFile)
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
	schemaFile := chunks.GetString(chunk, "schema")
	if spec.GetOverrideSchema() != "" {
		schemaFile = spec.GetOverrideSchema()
	}
	compiled, err := schema.CompileWithExtensions(schemaFile)
	return compiled, schemaFile, err
}
