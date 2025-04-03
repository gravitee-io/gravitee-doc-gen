package schema

import "github.com/santhosh-tekuri/jsonschema/v5"

type Visitor interface {
	OnAttribute(name string, schema *jsonschema.Schema, parent *jsonschema.Schema)
	OnObject(name string, schema *jsonschema.Schema)
	OnArray(name string, schema *jsonschema.Schema)
	OnOneOf(schema *jsonschema.Schema)
}
