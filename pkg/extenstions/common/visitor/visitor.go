package visitor

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"sort"
)

type schemaProperty struct {
	name   string
	schema *jsonschema.Schema
}

func Visit(ctx *VisitContext, visitor Visitor, current *jsonschema.Schema) {

	queue := make([]schemaProperty, 0)

	ordered := orderedAndResolved(current)

	for _, property := range ordered {
		name, attribute := property.name, property.schema
		if schema.IsAttribute(attribute) {
			added := visitAttribute(ctx, visitor, name, attribute, current)
			// skip the rest as it was already added before
			if !added {
				continue
			}
		}

		if ctx.IsQueueNodes() && !schema.IsAttribute(attribute) {
			visitor.OnAttribute(ctx, name, attribute, current)
			queue = append(queue, property)
		} else {
			visitNode(ctx, property, visitor)
		}
	}

	for _, pair := range queue {
		visitNode(ctx, pair, visitor)
	}

	if len(current.OneOf) > 0 {
		for _, schema := range current.OneOf {
			schema = orRef(schema)
			visitor.OnOneOf(ctx, schema, current)
			Visit(ctx, visitor, schema)
		}
		visitor.OnOneOfEnd(ctx)
	}

}

func orderedAndResolved(parent *jsonschema.Schema) []schemaProperty {
	ordered := make([]schemaProperty, 0, len(parent.Properties))
	for name, schema := range parent.Properties {
		ordered = append(ordered, schemaProperty{name: name, schema: orRef(schema)})
	}

	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].name < ordered[j].name
	})
	return ordered

}

func visitAttribute(ctx *VisitContext, visitor Visitor, property string, schema *jsonschema.Schema, parent *jsonschema.Schema) bool {
	attribute := visitor.OnAttribute(ctx, property, schema, parent)
	if ctx.nodeStack != nil && attribute != nil {
		if _, alreadyAdded := ctx.NodeStack().Peek().(*Object).Fields[property]; ctx.currentOneOf.Present && alreadyAdded {
			return false
		}
		ctx.NodeStack().add(ctx, attribute)
	}
	return true
}

func visitNode(ctx *VisitContext, prop schemaProperty, visitor Visitor) {
	if schema.IsObject(prop.schema) {
		visitObject(ctx, prop, visitor)
	}
	if schema.IsArray(prop.schema) {
		visitArray(ctx, prop, visitor)
	}
}

func visitObject(ctx *VisitContext, prop schemaProperty, visitor Visitor) {
	if containsOneOfs(prop.schema) {
		ctx.SetCurrentOneOf(findDiscriminators(prop.schema))
	}
	object := visitor.OnObjectStart(ctx, prop.name, prop.schema)
	if ctx.NodeStack() != nil {
		if object == nil {
			object = NewObject(prop.name)
		}
		ctx.NodeStack().add(ctx, object)
	}
	Visit(ctx, visitor, prop.schema)
	ctx.SetCurrentOneOf(OneOf{})
	visitor.OnObjectEnd(ctx)
	ctx.NodeStack().pop()
}

func visitArray(ctx *VisitContext, prop schemaProperty, visitor Visitor) {
	var items *jsonschema.Schema
	var itemTypeIsObject bool
	if prop.schema.Items != nil {
		items = prop.schema.Items.(*jsonschema.Schema)
		// no support of multiple types
		itemTypeIsObject = !schema.IsAttribute(items)
	}
	array, values := visitor.OnArrayStart(ctx, prop.name, prop.schema, itemTypeIsObject)
	if ctx.NodeStack() != nil {
		addArrayToStack(ctx, array, prop, itemTypeIsObject, values)
	}
	Visit(ctx, visitor, prop.schema.Items.(*jsonschema.Schema))
	visitor.OnArrayEnd(ctx, itemTypeIsObject)
	if itemTypeIsObject {
		ctx.NodeStack().pop()
	}
	ctx.NodeStack().pop()
}

func addArrayToStack(ctx *VisitContext, array *Array, prop schemaProperty, itemTypeIsObject bool, values []Value) {
	if array == nil {
		array = NewArray(prop.name)
	}
	ctx.NodeStack().add(ctx, array)
	if itemTypeIsObject {
		ctx.NodeStack().add(ctx, NewObject(""))
	} else if values != nil {
		for _, v := range values {
			ctx.NodeStack().add(ctx, v)
		}
	}
}

func containsOneOfs(schema *jsonschema.Schema) bool {
	return schema.OneOf != nil && len(schema.OneOf) > 0
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
				Type:     schema.GetType(parent),
				Property: name,
			}
			result = append(result, spec)
		}
	}
	oneOf := OneOf{}
	if len(result) > 0 {
		oneOf.ParentTitle = parent.Title
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
