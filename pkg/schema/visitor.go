package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

func Visit(parent *jsonschema.Schema, visitor Visitor) {

	queue := make(map[string]*jsonschema.Schema)

	for name, schema := range parent.Properties {
		schema = orRef(schema)
		if isAttribute(schema) {
			visitor.OnAttribute(name, schema, parent)
		}
		// process at the end to avoid mixing root level attribute and subsections
		if isObject(schema) || isArray(schema) {
			visitor.OnAttribute(name, schema, parent)
			queue[name] = schema
		}
	}

	for _, schema := range parent.OneOf {
		schema = orRef(schema)
		visitor.OnOneOf(schema)
		Visit(schema, visitor)
	}

	for name, schema := range queue {
		if isObject(schema) {
			visitor.OnObject(name, schema)
			Visit(schema, visitor)
		}
		if isArray(schema) {
			visitor.OnArray(name, schema)
			if schema.Items != nil {
				// no support of multiple types
				Visit(schema.Items.(*jsonschema.Schema), visitor)
			}
		}
	}
}

func orRef(schema *jsonschema.Schema) *jsonschema.Schema {
	if schema.Ref != nil {
		return schema.Ref
	}
	return schema
}

func isAttribute(schema *jsonschema.Schema) bool {
	return !(isObject(schema) || isArray(schema))
}

func isObject(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "object")
}

func isArray(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "array")
}
