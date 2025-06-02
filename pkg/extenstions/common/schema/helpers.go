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
	"slices"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema/extensions"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

const ArrayType = "array"
const ObjectType = "object"

func IsRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}

func GetTypeItem(attribute *jsonschema.Schema) string {
	if GetType(attribute) == ArrayType {
		return GetType(Items(attribute))
	}
	return ""
}

func GetType(prop *jsonschema.Schema) string {
	if len(prop.Types) == 0 {
		return ""
	}
	t := prop.Types[0]
	if len(prop.Enum) > 0 {
		return "enum (" + t + ")"
	}
	if t == ArrayType && IsAttribute(Items(prop)) {
		return ArrayType + " (" + GetType(Items(prop)) + ")"
	}
	return t
}

func IsArray(schema *jsonschema.Schema) bool {
	return GetType(schema) == ArrayType
}

func IsObject(schema *jsonschema.Schema) bool {
	return GetType(schema) == ObjectType
}

func IsAttribute(schema *jsonschema.Schema) bool {
	return !(IsObject(schema) || IsArray(schema))
}

func Items(array *jsonschema.Schema) *jsonschema.Schema {
	if array.Items != nil {
		if item, ok := array.Items.(*jsonschema.Schema); ok {
			return OrRef(item)
		}
	}
	if array.Items2020 != nil {
		return OrRef(array.Items2020)
	}
	panic("array.Items is nil or an array of types (Draft < 2020), this is not supported.")
}

func OrRef(schema *jsonschema.Schema) *jsonschema.Schema {
	if schema.Ref != nil {
		ref := schema.Ref
		switch {
		case defaultIsEmpty(ref):

			ref.Default = schema.Default
		case !ref.ReadOnly && schema.ReadOnly:
			ref.ReadOnly = true
		case !ref.WriteOnly && schema.WriteOnly:
			ref.WriteOnly = true
		case !ref.Deprecated && schema.Deprecated:
			ref.Deprecated = true
		case ref.Description == "":
			ref.Description = schema.Description
		case len(ref.Extensions) == 0 && len(schema.Extensions) > 0:
			ref.Extensions = schema.Extensions
		}
		return ref
	}
	return schema
}

func defaultIsEmpty(ref *jsonschema.Schema) bool {
	switch t := ref.Default.(type) {
	case nil:
		return true
	case string:
		return t == ""
	default:
		return false
	}
}

func IsDeprecated(schema *jsonschema.Schema) bool {
	if schema.Deprecated {
		return true
	}
	ext := GetExtension[extensions.DeprecatedSchema](schema, extensions.Deprecated)
	return bool(ext)
}

func GetExtension[S jsonschema.ExtSchema](att *jsonschema.Schema, name string) S {
	if extension, exists := att.Extensions[name]; exists {
		if ext, ok := extension.(S); ok {
			return ext
		}
	}
	return *new(S)
}
