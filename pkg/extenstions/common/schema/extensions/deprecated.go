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

package extensions

import (
	"errors"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

const (
	Deprecated = "deprecated"
)

type DeprecatedCompiler struct {
}
type DeprecatedSchema bool

func (c DeprecatedCompiler) Compile(
	_ jsonschema.CompilerContext,
	schema map[string]interface{}) (jsonschema.ExtSchema, error) {
	if e, exists := schema[Deprecated]; exists {
		if b, ok := e.(bool); !ok {
			return nil, errors.New(Deprecated + " must be a boolean")
		} else {
			return DeprecatedSchema(b), nil
		}
	}
	return DeprecatedSchema(false), nil
}

func (c DeprecatedSchema) Validate(_ jsonschema.ValidationContext, _ interface{}) error {
	// we don't validate payloads, no implementation required
	return nil
}
