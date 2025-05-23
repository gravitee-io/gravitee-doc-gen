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

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func IsRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}

func GetTypeItem(attribute *jsonschema.Schema) string {
	if GetType(attribute) == "array" {
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
	if t == "array" && IsAttribute(Items(prop)) {
		return "array (" + GetType(Items(prop)) + ")"
	}
	return t
}

func IsArray(schema *jsonschema.Schema) bool {
	return GetType(schema) == "array"
}

func IsObject(schema *jsonschema.Schema) bool {
	return GetType(schema) == "object"
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
		return OrRef(array.Items.(*jsonschema.Schema))
	}
	panic("array.Items is nil or an array of types (Draft < 2020), this is not supported.")
}

func OrRef(schema *jsonschema.Schema) *jsonschema.Schema {
	if schema.Ref != nil {
		ref := schema.Ref
		if defaultIsEmpty(ref) {
			ref.Default = schema.Default
		}
		if !ref.ReadOnly && schema.ReadOnly {
			ref.ReadOnly = true
		}
		if !ref.WriteOnly && schema.WriteOnly {
			ref.WriteOnly = true
		}
		if !ref.Deprecated && schema.Deprecated {
			ref.Deprecated = true
		}
		if ref.Description == "" {
			ref.Description = schema.Description
		}
		return ref
	}
	return schema
}

func defaultIsEmpty(ref *jsonschema.Schema) bool {
	switch ref.Default.(type) {
	case nil:
		return true
	case string:
		return ref.Default.(string) == ""
	default:
		return false
	}
}
