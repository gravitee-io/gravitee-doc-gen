package schema

import (
	"fmt"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

func IsRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}
func GetConstantOrDefault(att *jsonschema.Schema) string {
	if att.Constant != nil {
		return fmt.Sprintf("%v", att.Constant[0])
	}
	if att.Default == nil {
		return ""
	}
	return fmt.Sprintf("%v", att.Default)
}
