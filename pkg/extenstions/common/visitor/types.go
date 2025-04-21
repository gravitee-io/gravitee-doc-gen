package visitor

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Visitor interface {
	OnAttribute(ctx *VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *Attribute
	OnObjectStart(ctx *VisitContext, property string, object *jsonschema.Schema) *Object
	OnObjectEnd(ctx *VisitContext)
	OnArrayStart(ctx *VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) (*Array, []Value)
	OnArrayEnd(ctx *VisitContext, itemTypeIsObject bool)
	OnOneOf(visitCtx *VisitContext, oneOf *jsonschema.Schema, parent *jsonschema.Schema)
	OnOneOfEnd(*VisitContext)
}

type DiscriminatorSpec struct {
	Values   []any
	Property string
	Type     string
}
