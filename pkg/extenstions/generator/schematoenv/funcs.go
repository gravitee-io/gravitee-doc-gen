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

package schematoenv

import (
	"errors"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	common2 "github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	visitor2 "github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
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
