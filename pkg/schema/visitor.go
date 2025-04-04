package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

type VisitContext struct {
	CurrentOneOf OneOf
	QueueNodes   bool
}
type pair struct {
	name   string
	schema *jsonschema.Schema
}

func Visit(parent *jsonschema.Schema, visitor Visitor, visitCtx *VisitContext) {

	queue := make([]pair, 0)
	for name, schema := range parent.Properties {
		schema = orRef(schema)
		if IsAttribute(schema) {
			visitor.OnAttribute(name, schema, parent, visitCtx)
		}
		pair := pair{name, schema}
		if visitCtx.QueueNodes && (isObject(schema) || isArray(schema)) {
			visitor.OnAttribute(name, schema, parent, visitCtx)
			queue = append(queue, pair)
		} else {
			visitNode(pair, visitor, visitCtx)
		}
	}

	for _, pair := range queue {
		visitNode(pair, visitor, visitCtx)
	}
	for i, schema := range parent.OneOf {
		schema = orRef(schema)
		visitor.OnOneOfStart(schema, parent, visitCtx, i)
		Visit(schema, visitor, visitCtx)
		visitor.OnOneOfEnd()
	}

}

func visitNode(pair pair, visitor Visitor, visitCtx *VisitContext) {
	if isObject(pair.schema) {
		if ContainsOneOfs(pair.schema) {
			visitCtx.CurrentOneOf = findDiscriminators(pair.schema)
		}
		visitor.OnObjectStart(pair.name, pair.schema, visitCtx)
		Visit(pair.schema, visitor, visitCtx)
		visitor.OnObjectEnd()
		visitCtx.CurrentOneOf = OneOf{}
	}
	if isArray(pair.schema) {
		var items *jsonschema.Schema
		var itemTypeIsObject bool
		if pair.schema.Items != nil {
			items = pair.schema.Items.(*jsonschema.Schema)
			// no support of multiple types
			itemTypeIsObject = !IsAttribute(items)
		}
		visitor.OnArrayStart(pair.name, pair.schema, itemTypeIsObject)
		Visit(pair.schema.Items.(*jsonschema.Schema), visitor, visitCtx)
		visitor.OnArrayEnd(itemTypeIsObject)
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

func ContainsOneOfs(schema *jsonschema.Schema) bool {
	return schema.OneOf != nil && len(schema.OneOf) > 0
}

func IsAttribute(schema *jsonschema.Schema) bool {
	return !(isObject(schema) || isArray(schema))
}

func findDiscriminators(parent *jsonschema.Schema) OneOf {
	found := make(map[string]int)
	expected := len(parent.OneOf)
	values := make(map[string][]any)

	for _, oneOf := range parent.OneOf {
		for name, prop := range oneOf.Properties {
			count := found[name]
			if prop.Constant != nil && len(prop.Constant) > 0 {
				count += 1
				found[name] = count
				array := values[name]
				if array == nil {
					array = make([]any, 0)
				}
				array = append(array, prop.Constant[0])
				values[name] = array
			}
		}
	}

	result := make([]DiscriminatorSpec, 0)
	for name, count := range found {
		if count == expected {
			spec := DiscriminatorSpec{
				Values:   values[name],
				Type:     GetType(parent),
				Property: name,
			}
			result = append(result, spec)
		}
	}
	oneOf := OneOf{}
	if len(result) > 0 {
		oneOf.Parent = parent.Title
		oneOf.Present = true
		oneOf.Specs = result

	}
	return oneOf
}

func orRef(schema *jsonschema.Schema) *jsonschema.Schema {
	if schema.Ref != nil {
		return schema.Ref
	}
	return schema
}

func isObject(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "object")
}

func isArray(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "array")
}
