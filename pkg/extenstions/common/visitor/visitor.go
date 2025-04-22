// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package visitor

import (
	"sort"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
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
		for _, s := range current.OneOf {
			s = orRef(s)
			visitor.OnOneOf(ctx, s, current)
			Visit(ctx, visitor, s)
		}
		visitor.OnOneOfEnd(ctx)
	}
}

func orderedAndResolved(parent *jsonschema.Schema) []schemaProperty {
	ordered := make([]schemaProperty, 0, len(parent.Properties))
	for name, s := range parent.Properties {
		ordered = append(ordered, schemaProperty{name: name, schema: orRef(s)})
	}

	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].name < ordered[j].name
	})
	return ordered
}

func visitAttribute(
	ctx *VisitContext,
	visitor Visitor,
	property string,
	schema *jsonschema.Schema,
	parent *jsonschema.Schema) bool {
	attribute := visitor.OnAttribute(ctx, property, schema, parent)
	if ctx.nodeStack != nil && attribute != nil {
		if object, ok := ctx.NodeStack().Peek().(*Object); ok {
			if _, alreadyAdded := object.Fields[property]; ctx.currentOneOf.Present && alreadyAdded {
				return false
			}
			ctx.NodeStack().add(ctx, attribute)
		} else {
			return false
		}
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
	if len(prop.schema.OneOf) > 0 {
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
		items = schema.Items(prop.schema)
		// no support of multiple types
		itemTypeIsObject = !schema.IsAttribute(items)
	}
	array, values := visitor.OnArrayStart(ctx, prop.name, prop.schema, itemTypeIsObject)
	if ctx.NodeStack() != nil {
		addArrayToStack(ctx, array, prop, itemTypeIsObject, values)
	}
	Visit(ctx, visitor, items)
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
	} else {
		for _, v := range values {
			ctx.NodeStack().add(ctx, v)
		}
	}
}

func findDiscriminators(parent *jsonschema.Schema) OneOf {
	found := make(map[string]int)
	expected := len(parent.OneOf)
	values := make(map[string][]any)

	for _, oneOf := range parent.OneOf {
		for name, prop := range oneOf.Properties {
			count := found[name]
			if len(prop.Constant) > 0 {
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
