package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Visitor interface {
	OnAttribute(ctx *VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema)
	OnObjectStart(ctx *VisitContext, property string, object *jsonschema.Schema)
	OnObjectEnd(ctx *VisitContext)
	OnArrayStart(ctx *VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool)
	OnArrayEnd(ctx *VisitContext, itemTypeIsObject bool)
	OnOneOfStart(visitCtx *VisitContext, oneOf *jsonschema.Schema, parent *jsonschema.Schema)
	OnOneOfEnd(*VisitContext)
}

type OneOf struct {
	Parent  string
	Present bool
	Specs   []DiscriminatorSpec
}

func (o OneOf) IsZero() bool {
	return o.Parent == "" && o.Present == false && len(o.Specs) == 0
}

type DiscriminatorSpec struct {
	Values   []any
	Property string
	Type     string
}
