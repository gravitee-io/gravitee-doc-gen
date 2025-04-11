package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Visitor interface {
	OnAttribute(property string, attribute *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *VisitContext)
	OnObjectStart(property string, object *jsonschema.Schema, visitCtx *VisitContext)
	OnObjectEnd(visitCtx *VisitContext)
	OnArrayStart(property string, array *jsonschema.Schema, itemTypeIsObject bool, ctx *VisitContext)
	OnArrayEnd(itemTypeIsObject bool)
	OnOneOfStart(oneOf *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *VisitContext)
	OnOneOfEnd()
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
