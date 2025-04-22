package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

func IsRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}

func GetTypeItem(attribute *jsonschema.Schema) string {
	if GetType(attribute) == "array" {
		return GetType(attribute.Items.(*jsonschema.Schema))
	}
	return ""
}

func GetType(prop *jsonschema.Schema) string {
	if len(prop.Types) == 0 {
		return ""
	}
	t := prop.Types[0]
	if prop.Enum != nil && len(prop.Enum) > 0 {
		return "enum (" + t + ")"
	}
	return t
}

func IsArray(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "array")
}

func IsObject(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "object")
}

func IsAttribute(schema *jsonschema.Schema) bool {
	return !(IsObject(schema) || IsArray(schema))
}
