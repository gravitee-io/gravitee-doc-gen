package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

func IsRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}
func GetConstantOrDefault(att *jsonschema.Schema) any {
	if att.Constant != nil {
		return att.Constant[0]
	}
	return att.Default
}
