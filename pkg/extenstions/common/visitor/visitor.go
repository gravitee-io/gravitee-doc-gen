package visitor

import (
	"fmt"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
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
		if IsAttribute(attribute) {
			fmt.Println("attribute:", name, ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
			attribute := visitor.OnAttribute(ctx, name, attribute, current)
			if ctx.nodeStack != nil && attribute != nil {
				if _, alreadyAdded := ctx.NodeStack().Peek().(*Object).Fields[name]; ctx.currentOneOf.Present && alreadyAdded {
					continue
				}
				ctx.NodeStack().add(ctx, attribute)
			}
		}

		if ctx.IsQueueNodes() && !IsAttribute(attribute) {
			fmt.Println("attribute (queue):", name, ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
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
			fmt.Println("oneof:", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
			visitor.OnOneOf(ctx, schema, current)
			Visit(ctx, visitor, schema)
		}
		fmt.Println("end oneof:", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
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

func visitNode(ctx *VisitContext, prop schemaProperty, visitor Visitor) {
	if isObject(prop.schema) {
		visitObject(ctx, prop, visitor)
	}
	if isArray(prop.schema) {
		visitArray(ctx, prop, visitor)
	}
}

func visitArray(ctx *VisitContext, prop schemaProperty, visitor Visitor) {
	var items *jsonschema.Schema
	var itemTypeIsObject bool
	if prop.schema.Items != nil {
		items = prop.schema.Items.(*jsonschema.Schema)
		// no support of multiple types
		itemTypeIsObject = !IsAttribute(items)
	}
	fmt.Println("start array:", prop.name, "@", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
	array, values := visitor.OnArrayStart(ctx, prop.name, prop.schema, itemTypeIsObject)
	if ctx.NodeStack() != nil {
		addArrayToStack(ctx, array, prop, itemTypeIsObject, values)
	}
	Visit(ctx, visitor, prop.schema.Items.(*jsonschema.Schema))
	fmt.Println("end array:", prop.name, "@", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
	visitor.OnArrayEnd(ctx, itemTypeIsObject)
	if itemTypeIsObject {
		ctx.NodeStack().pop()
	}
	ctx.NodeStack().pop()
}

func addArrayToStack(ctx *VisitContext, array *Array, prop schemaProperty, itemTypeIsObject bool, values []Value) {
	fmt.Println("add array:", prop.name, "@", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
	if array == nil {
		array = NewArray(prop.name)
	}
	ctx.NodeStack().add(ctx, array)
	if itemTypeIsObject {
		fmt.Println("add object:", prop.name, "@", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
		ctx.NodeStack().add(ctx, NewObject(""))
	} else if values != nil {
		for _, v := range values {
			ctx.NodeStack().add(ctx, v)
		}
	}
}

func visitObject(ctx *VisitContext, prop schemaProperty, visitor Visitor) {
	if ContainsOneOfs(prop.schema) {
		ctx.SetCurrentOneOf(findDiscriminators(prop.schema))
	}
	fmt.Println("start object:", prop.name, "@", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
	object := visitor.OnObjectStart(ctx, prop.name, prop.schema)
	if ctx.NodeStack() != nil {
		fmt.Println("add object:", prop.name, "@", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
		if object == nil {
			object = NewObject(prop.name)
		}
		ctx.NodeStack().add(ctx, object)
	}
	Visit(ctx, visitor, prop.schema)
	ctx.SetCurrentOneOf(OneOf{})
	fmt.Println("end object:", prop.name, "@", ctx.nodeStack.Peek().Name(), ctx.nodeStack.Peek().Kind().String())
	visitor.OnObjectEnd(ctx)
	ctx.NodeStack().pop()
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

func isObject(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "object")
}

func isArray(schema *jsonschema.Schema) bool {
	return slices.Contains(schema.Types, "array")
}
