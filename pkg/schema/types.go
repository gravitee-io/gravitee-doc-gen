package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Visitor interface {
	OnAttribute(ctx *VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *Attribute
	OnObjectStart(ctx *VisitContext, property string, object *jsonschema.Schema)
	OnObjectEnd(ctx *VisitContext)
	OnArrayStart(ctx *VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) (values []Value)
	OnArrayEnd(ctx *VisitContext, itemTypeIsObject bool)
	OnOneOfStart(visitCtx *VisitContext, oneOf *jsonschema.Schema, parent *jsonschema.Schema)
	OnOneOfEnd(*VisitContext)
}

type OneOf struct {
	ParentTitle string
	Present     bool
	Specs       []DiscriminatorSpec
}

func (o OneOf) IsZero() bool {
	return o.ParentTitle == "" && o.Present == false && len(o.Specs) == 0
}

func (o OneOf) IsDiscriminator(property string) bool {
	for _, spec := range o.Specs {
		if spec.Property == property {
			return true
		}
	}
	return false
}

type DiscriminatorSpec struct {
	Values   []any
	Property string
	Type     string
}
