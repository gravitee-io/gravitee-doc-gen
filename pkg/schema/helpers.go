package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

func IsRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}
func GetConstantOrDefault(att *jsonschema.Schema, defaultBoolean bool) any {
	if att.Constant != nil {
		return att.Constant[0]
	}
	def := att.Default
	if def == nil && GetType(att) == "boolean" && defaultBoolean {
		return false
	}
	return def
}

func GetTypeItem(attribute *jsonschema.Schema) string {
	if GetType(attribute) == "array" {
		return GetType(attribute.Items.(*jsonschema.Schema))
	}
	return ""
}
