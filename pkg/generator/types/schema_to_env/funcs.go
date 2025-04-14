package schema_to_env

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
	schemaFile := common.GetString(chunk, "schema")
	schemaFileExists := common.FileExists(schemaFile)

	if chunk.Required && !schemaFileExists {
		return false, errors.New("schema file not found")
	}

	return tmplExists && schemaFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	schemaFile := common.GetString(chunk, "schema")

	compiled, err := schema.CompileWithExtensions(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	indexPlaceholder := common.GetDataOrDefault[string](chunk, "indexPlaceholder", "X")
	prefix := common.GetDataOrDefault[string](chunk, "prefix", "")
	schemaVisitor := newSchemaVisitor(indexPlaceholder, prefix)
	schema.Visit(schema.NewVisitContextWithStack(schema.NewObject(""), true, true), &schemaVisitor, compiled)

	processed := chunks.Processed{Data: schemaVisitor}
	return processed, nil
}
