package schema_to_env

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

	ctx := schema.NewVisitContext(true, true).WithStack(schema.NewObject(""))
	schema.Visit(ctx, &types.SchemaVisitor{KeepAllOneOfAttributes: true}, compiled)

	indexPlaceholder := common.GetDataOrDefault[string](chunk, "indexPlaceholder", "X")
	prefix := common.GetDataOrDefault[string](chunk, "prefix", "")

	envVisitor := toEnvVisitor{
		Sections:         make([]*envSection, 0),
		jvmPaths:         make([]string, 0),
		envPaths:         make([]string, 0),
		indexPlaceholder: indexPlaceholder,
		prefix:           prefix,
	}

	types.Visit(ctx, &envVisitor)

	processed := chunks.Processed{Data: envVisitor}
	return processed, nil
}
