// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package schematoyaml

import (
	"errors"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
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
	object := visitor.NewObject("")
	ctx := visitor.NewVisitContext(false, true).WithStack(object)
	schemaVisitor := &common.SchemaToNodeTreeVisitor{KeepAllOneOfAttributes: true}
	visitor.Visit(ctx, schemaVisitor, compiled)

	yamlVisitor := toYamlVisitor{
		Lines:   make([]yamlLine, 0),
		padding: 3,
	}
	common.Visit(ctx.NodeStack(), &yamlVisitor)

	processed := chunks.Processed{Data: yamlVisitor}
	return processed, nil
}
