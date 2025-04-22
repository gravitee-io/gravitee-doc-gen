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
