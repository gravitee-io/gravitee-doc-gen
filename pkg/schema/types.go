package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Visitor interface {
	OnAttribute(property string, attribute *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *VisitContext)
	OnObjectStart(property string, object *jsonschema.Schema, visitCtx *VisitContext)
	OnObjectEnd()
	OnArrayStart(property string, array *jsonschema.Schema, itemTypeIsObject bool)
	OnArrayEnd(itemTypeIsObject bool)
	OnOneOfStart(oneOf *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *VisitContext)
	OnOneOfEnd()
}

type OneOf struct {
	Parent  string
	Present bool
	Specs   []DiscriminatorSpec
}

type DiscriminatorSpec struct {
	Values   []any
	Property string
	Type     string
}
