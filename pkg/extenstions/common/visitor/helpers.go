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

package visitor

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func GetDefaultOrFirstExample(att *jsonschema.Schema, ctx *VisitContext) any {
	value := GetConstantOrDefault(att, ctx)
	if value == nil && len(att.Examples) > 0 {
		value = att.Examples[0]
	}
	return value
}

func GetConstantOrDefault(att *jsonschema.Schema, ctx *VisitContext) any {
	if att.Constant != nil {
		return att.Constant[0]
	}
	def := att.Default
	if def == nil && schema.GetType(att) == "boolean" && ctx.IsAutoDefaultBooleans() {
		return false
	}
	return def
}
