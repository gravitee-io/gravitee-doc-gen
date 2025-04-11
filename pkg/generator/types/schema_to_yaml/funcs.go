package schema_to_yaml

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := common.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}
	schemaFile := common.GetFile(chunk, "schema")
	schemaFileExists := common.FileExists(schemaFile)

	if chunk.Required && !schemaFileExists {
		return false, errors.New("schema file not found")
	}

	return tmplExists && schemaFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	schemaFile := common.GetFile(chunk, "schema")

	compiled, err := schema.CompileWithExtensions(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}
	schemaVisitor := newSchemaVisitor()
	schemaVisitor.padding = 2
	schema.Visit(compiled, &schemaVisitor, &schema.VisitContext{AutoDefaultBooleans: true})

	processed := chunks.Processed{Data: schemaVisitor}
	return processed, nil
}
