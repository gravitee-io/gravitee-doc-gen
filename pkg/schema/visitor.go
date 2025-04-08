package schema

import (
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
	"sort"
)

type VisitContext struct {
	CurrentOneOf        OneOf
	QueueNodes          bool
	AutoDefaultBooleans bool
}
type property struct {
	name   string
	schema *jsonschema.Schema
}

// TODO move to util and use anywhere !
type Unstructured map[string]interface{}

func Visit(parent *jsonschema.Schema, visitor Visitor, visitCtx *VisitContext) {

	queue := make([]property, 0)

	ordered := orderedAndResolved(parent)

	for _, property := range ordered {
		name, schema := property.name, property.schema
		if IsAttribute(schema) {
			visitor.OnAttribute(name, schema, parent, visitCtx)
		}

		if visitCtx.QueueNodes && (isObject(schema) || isArray(schema)) {
			visitor.OnAttribute(name, schema, parent, visitCtx)
			queue = append(queue, property)
		} else {
			visitNode(property, visitor, visitCtx)
		}
	}

	for _, pair := range queue {
		visitNode(pair, visitor, visitCtx)
	}

	for _, schema := range parent.OneOf {
		schema = orRef(schema)
		visitor.OnOneOfStart(schema, parent, visitCtx)
		Visit(schema, visitor, visitCtx)
		visitor.OnOneOfEnd()
	}

}

func orderedAndResolved(parent *jsonschema.Schema) []property {
	ordered := make([]property, 0, len(parent.Properties))
	for name, schema := range parent.Properties {
		ordered = append(ordered, property{name: name, schema: orRef(schema)})
	}

	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].name < ordered[j].name
	})
	return ordered

}

func visitNode(prop property, visitor Visitor, visitCtx *VisitContext) {
	if isObject(prop.schema) {
		if ContainsOneOfs(prop.schema) {
			visitCtx.CurrentOneOf = findDiscriminators(prop.schema)
		}
		visitor.OnObjectStart(prop.name, prop.schema, visitCtx)
		Visit(prop.schema, visitor, visitCtx)
		visitor.OnObjectEnd()
		visitCtx.CurrentOneOf = OneOf{}
	}
	if isArray(prop.schema) {
		var items *jsonschema.Schema
		var itemTypeIsObject bool
		if prop.schema.Items != nil {
			items = prop.schema.Items.(*jsonschema.Schema)
			// no support of multiple types
			itemTypeIsObject = !IsAttribute(items)
		}
		visitor.OnArrayStart(prop.name, prop.schema, itemTypeIsObject, visitCtx)
		Visit(prop.schema.Items.(*jsonschema.Schema), visitor, visitCtx)
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
