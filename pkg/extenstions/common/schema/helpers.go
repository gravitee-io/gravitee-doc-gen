package schema

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/visitor"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

func IsRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}

func GetTypeItem(attribute *jsonschema.Schema) string {
	if visitor.GetType(attribute) == "array" {
		return visitor.GetType(attribute.Items.(*jsonschema.Schema))
	}
	return ""
}
