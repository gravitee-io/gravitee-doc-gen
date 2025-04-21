package schema_to_yaml

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/util"
	common2 "github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common"
	schema2 "github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/schema"
	visitor3 "github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/visitor"
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

	compiled, err := schema2.CompileWithExtensions(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}
	object := visitor3.NewObject("")
	ctx := visitor3.NewVisitContext(false, true).WithStack(object)
	schemaVisitor := &common2.SchemaToNodeTreeVisitor{KeepAllOneOfAttributes: true}
	visitor3.Visit(ctx, schemaVisitor, compiled)

	visitor := toYamlVisitor{
		Lines:   make([]yamlLine, 0),
		padding: 3,
	}
	common2.Visit(ctx.NodeStack(), &visitor)

	processed := chunks.Processed{Data: visitor}
	return processed, nil
}
