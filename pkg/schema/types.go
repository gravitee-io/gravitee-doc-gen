package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Visitor interface {
	OnAttribute(name string, schema *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *VisitContext)
	OnObject(name string, schema *jsonschema.Schema, visitCtx *VisitContext)
	OnArray(name string, schema *jsonschema.Schema)
	OnOneOf(schema *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *VisitContext)
}

type OneOf struct {
	Values   []string
	Parent   string
	Property string
	Type     string
	Present  bool
}
