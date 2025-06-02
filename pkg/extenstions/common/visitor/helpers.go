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
	"slices"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func GetValueOrFirstExample(att *jsonschema.Schema, ctx *VisitContext) any {
	value := GetValue(att, ctx)
	if value == nil && len(att.Examples) > 0 {
		value = att.Examples[0]
	}
	return value
}

func GetValue(att *jsonschema.Schema, ctx *VisitContext) any {
	if att.Constant != nil {
		return att.Constant[0]
	}

	def := att.Default

	if def == nil && len(att.Enum) == 1 {
		return att.Enum[0]
	}

	if def == nil && schema.GetType(att) == "boolean" && ctx.IsAutoDefaultBooleans() {
		return false
	}

	return def
}

func GetOneOfs(current *jsonschema.Schema) []*jsonschema.Schema {
	oneOfs := getDependentOneOfs(current)
	if len(current.OneOf) > 0 {
		oneOfs = append(oneOfs, current.OneOf...)
	}
	return oneOfs
}

func NewSchemaPropertyList(current *jsonschema.Schema) SchemaPropertyList {
	dependentProperties := getDependentProperties(current)

	ordered := SchemaPropertyList{}
	for name, s := range current.Properties {
		ordered.Add(name, schema.OrRef(s))
	}

	if len(dependentProperties) > 0 {
		for name, s := range dependentProperties {
			ordered.Add(name, schema.OrRef(s))
		}
	}
	ordered = slices.DeleteFunc(ordered, func(s SchemaProperty) bool {
		return schema.IsDeprecated(s.schema)
	})
	ordered.Sort()
	return ordered
}

func getDependentProperties(current *jsonschema.Schema) map[string]*jsonschema.Schema {
	result := map[string]*jsonschema.Schema{}

	if len(current.Dependencies) > 0 {
		for _, dep := range current.Dependencies {
			pushProperties(dep, result)
		}
	}
	if len(current.DependentSchemas) > 0 {
		for _, dep := range current.DependentSchemas {
			pushProperties(dep, result)
		}
	}

	if len(current.AllOf) > 0 {
		for _, dep := range current.AllOf {
			pushProperties(dep, result)
		}
	}

	return result
}

func pushProperties(dep any, out map[string]*jsonschema.Schema) {
	if d, ok := dep.(*jsonschema.Schema); ok {
		d = schema.OrRef(d)
		if len(d.Properties) > 0 {
			for name, s := range d.Properties {
				out[name] = s
			}
		}
	}
}

func getDependentOneOfs(current *jsonschema.Schema) []*jsonschema.Schema {
	result := make([]*jsonschema.Schema, 0)
	if len(current.Dependencies) > 0 {
		for _, dep := range current.Dependencies {
			result = append(result, collectOneOfs(dep)...)
		}
	}
	if len(current.DependentSchemas) > 0 {
		for _, dep := range current.DependentSchemas {
			result = append(result, collectOneOfs(dep)...)
		}
	}

	if len(current.AllOf) > 0 {
		for _, dep := range current.AllOf {
			result = append(result, collectOneOfs(dep)...)
		}
	}
	return result
}

func collectOneOfs(dep any) []*jsonschema.Schema {
	result := make([]*jsonschema.Schema, 0)
	if d, ok := dep.(*jsonschema.Schema); ok {
		d = schema.OrRef(d)
		if len(d.OneOf) > 0 {
			for _, oneOf := range d.OneOf {
				result = append(result, schema.OrRef(oneOf))
			}
		}
	}
	return result
}
