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

package schema

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema/extensions"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func CompileWithExtensions(schemaFile string) (*jsonschema.Schema, error) {
	compiled, err := CompilerWithExtensions().Compile(schemaFile)
	if err == nil && compiled.Draft == nil {
		// panic("Schema version must set to 'Draft 7' as follow: \"$schema\": \"http://json-schema.org/draft-07/schema#\"")
		panic("Schema version must set: \"$schema\": \"http://json-schema.org/draft-07/schema#\" for instance")
	}
	return compiled, err
}

func CompilerWithExtensions() *jsonschema.Compiler {
	compiler := jsonschema.NewCompiler()
	compiler.ExtractAnnotations = true
	compiler.RegisterExtension(extensions.GioConfigExtension, nil, &extensions.GioConfigCompiler{})
	return compiler
}
