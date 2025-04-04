package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

type VisitContext struct {
	CurrentOneOf OneOf
}

func Visit(parent *jsonschema.Schema, visitor Visitor, visitCtx *VisitContext) {

	queue := make(map[string]*jsonschema.Schema)

	for name, schema := range parent.Properties {
		schema = orRef(schema)
		if isAttribute(schema) {
			visitor.OnAttribute(name, schema, parent, visitCtx)
		}
		// process at the end to avoid mixing root level attribute and subsections
		if isObject(schema) || isArray(schema) {
			visitor.OnAttribute(name, schema, parent, visitCtx)
			queue[name] = schema
		}
	}

	for name, schema := range queue {
		if isObject(schema) {
			if IsOneOf(schema) {
				visitCtx.CurrentOneOf = findDiscriminator(schema)
			}
			visitor.OnObject(name, schema, visitCtx)
			Visit(schema, visitor, visitCtx)
			visitCtx.CurrentOneOf = OneOf{}
		}
		if isArray(schema) {
			visitor.OnArray(name, schema)
			if schema.Items != nil {
				// no support of multiple types
				Visit(schema.Items.(*jsonschema.Schema), visitor, visitCtx)
			}
		}
	}
	for _, schema := range parent.OneOf {
		schema = orRef(schema)
		visitor.OnOneOf(schema, parent, visitCtx)
		Visit(schema, visitor, visitCtx)
	}

}

func GetType(prop *jsonschema.Schema) string {
	if len(prop.Types) == 0 {
		return ""
	}
	t := prop.Types[0]
	if prop.Enum != nil && len(prop.Enum) > 0 {
		return "enum (" + t + ")"
	}
	return t
}

func IsOneOf(schema *jsonschema.Schema) bool {
	return schema.OneOf != nil && len(schema.OneOf) > 0
}

func findDiscriminator(parent *jsonschema.Schema) OneOf {
	found := make(map[string]int)
	expected := len(parent.OneOf)
	values := make([]string, 0, expected)
	var maxCount int
	var property string

	for _, oneOf := range parent.OneOf {
		for name, prop := range oneOf.Properties {
			count := found[name]
			if GetType(prop) == "string" || GetType(prop) == "" && (prop.Constant != nil && len(prop.Constant) > 0) {
				count += 1
				found[name] = count
				values = append(values, prop.Constant[0].(string))
				if maxCount = max(maxCount, count); count == maxCount {
					property = name
				}
			}
		}
	}
	if maxCount == expected {
		return OneOf{
			Property: property,
			Values:   values,
			Type:     "string",
			Parent:   parent.Title,
			Present:  true,
		}
	}
	return OneOf{}
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
