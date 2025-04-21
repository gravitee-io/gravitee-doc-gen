package schema_to_env

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/util"
	common2 "github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/schema"
	visitor2 "github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/visitor"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := chunks.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}
	schemaFile := chunks.GetString(chunk, "schema")
	schemaFileExists := util.FileExists(schemaFile)

	if chunk.Required && !schemaFileExists {
		return false, errors.New("schema file not found")
	}

	return tmplExists && schemaFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	schemaFile := chunks.GetString(chunk, "schema")

	compiled, err := schema.CompileWithExtensions(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	ctx := visitor2.NewVisitContext(true, true).
		WithStack(visitor2.NewObject(""))
	visitor2.Visit(ctx, &common2.SchemaToNodeTreeVisitor{KeepAllOneOfAttributes: true}, compiled)

	indexPlaceholder := chunks.GetDataOrDefault[string](chunk, "indexPlaceholder", "X")
	prefix := chunks.GetDataOrDefault[string](chunk, "prefix", "")

	envVisitor := toEnvVisitor{
		Sections:         make([]*envSection, 0),
		jvmPaths:         make([]string, 0),
		envPaths:         make([]string, 0),
		indexPlaceholder: indexPlaceholder,
		prefix:           prefix,
	}

	common2.Visit(ctx.NodeStack(), &envVisitor)

	processed := chunks.Processed{Data: envVisitor}
	return processed, nil
}
