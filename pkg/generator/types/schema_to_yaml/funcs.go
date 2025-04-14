package schema_to_yaml

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types"
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
	object := schema.NewObject("")
	ctx := schema.NewVisitContext(false, true).WithStack(object)
	schemaVisitor := &types.SchemaVisitor{KeepAllOneOfAttributes: true}
	schema.Visit(ctx, schemaVisitor, compiled)

	visitor := toYamlVisitor{
		Lines:   make([]yamlLine, 0),
		padding: 3,
	}
	types.Visit(ctx, &visitor)

	processed := chunks.Processed{Data: visitor}
	return processed, nil
}
