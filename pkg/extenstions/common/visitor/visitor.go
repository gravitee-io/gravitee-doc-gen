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
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func Visit(ctx *VisitContext, visitor Visitor, current *jsonschema.Schema) {
	queue := make([]SchemaProperty, 0)

	schemaPropertyList := NewSchemaPropertyList(current)
	for _, property := range schemaPropertyList {
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

	oneOfs := GetOneOfs(current)
	if len(oneOfs) > 0 {
		ctx.PushOneOf(newOneOfDescriptor(current))
		visitor.OnOneOfStart(ctx, current)
		for _, s := range oneOfs {
			s = schema.OrRef(s)
			visitor.OnOneOf(ctx, s, current)
			Visit(ctx, visitor, s)
		}
		visitor.OnOneOfEnd(ctx)
		ctx.PopOneOf()
	}
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
			if _, alreadyAdded := object.Fields[property]; ctx.PeekOneOf().Present && alreadyAdded {
				return false
			}
			ctx.NodeStack().add(ctx, attribute)
		} else {
			return false
		}
	}
	return true
}

func visitNode(ctx *VisitContext, prop SchemaProperty, visitor Visitor) {
	if schema.IsObject(prop.schema) {
		visitObject(ctx, prop, visitor)
	}
	if schema.IsArray(prop.schema) {
		visitArray(ctx, prop, visitor)
	}
}

func visitObject(ctx *VisitContext, prop SchemaProperty, visitor Visitor) {
	object := visitor.OnObjectStart(ctx, prop.name, prop.schema)
	if ctx.NodeStack() != nil {
		if object == nil {
			object = NewObject(prop.name)
		}
		ctx.NodeStack().add(ctx, object)
	}
	Visit(ctx, visitor, prop.schema)
	visitor.OnObjectEnd(ctx)
	ctx.NodeStack().pop()
}

func visitArray(ctx *VisitContext, prop SchemaProperty, visitor Visitor) {
	var items *jsonschema.Schema
	var itemTypeIsObject bool
	var oneOfAdded bool
	if prop.schema.Items != nil {
		items = schema.Items(prop.schema)
		itemTypeIsObject = !schema.IsAttribute(items)
		if itemTypeIsObject && len(GetOneOfs(items)) > 0 {
			ctx.PushOneOf(newOneOfDescriptor(items))
			oneOfAdded = true
		}
	}
	array, values := visitor.OnArrayStart(ctx, prop.name, prop.schema, itemTypeIsObject)
	if ctx.NodeStack() != nil {
		addArrayToStack(ctx, array, prop, itemTypeIsObject, values)
	}
	if itemTypeIsObject && len(values) == 0 {
		Visit(ctx, visitor, items)
	}

	if oneOfAdded {
		ctx.PopOneOf()
	}
	visitor.OnArrayEnd(ctx, itemTypeIsObject)
	if itemTypeIsObject && len(values) == 0 {
		ctx.NodeStack().pop()
	}
	ctx.NodeStack().pop()
}

func addArrayToStack(ctx *VisitContext, array *Array, prop SchemaProperty, itemTypeIsObject bool, values []Value) {
	if array == nil {
		array = NewArray(prop.name)
	}
	ctx.NodeStack().add(ctx, array)
	if itemTypeIsObject && len(values) == 0 {
		ctx.NodeStack().add(ctx, NewObject(""))
	} else {
		for _, v := range values {
			ctx.NodeStack().add(ctx, v)
		}
	}
}

func newOneOfDescriptor(parent *jsonschema.Schema) OneOfDescriptor {
	oneOfs := GetOneOfs(parent)
	found := make(map[string]int)
	expected := len(oneOfs)
	values := make(map[string]util.Set)

	for _, oneOf := range oneOfs {
		for _, property := range NewSchemaPropertyList(oneOf) {
			name, s := property.name, property.schema
			count := found[name]
			if len(s.Constant) > 0 {
				count += 1
				found[name] = count
				set := values[name]
				if set == nil {
					set = util.Set{}
				}
				set.Add(s.Constant[0])
				values[name] = set
			}
		}
	}

	result := make([]DiscriminatorSpec, 0)
	for name, count := range found {
		if count == expected {
			spec := DiscriminatorSpec{
				Values:   values[name].ToSlice(),
				Type:     schema.GetType(parent),
				Property: name,
			}
			result = append(result, spec)
		}
	}
	oneOf := OneOfDescriptor{}
	if len(result) > 0 {
		oneOf.ParentTitle = parent.Title
		oneOf.Present = true
		oneOf.Specs = result
	}
	return oneOf
}
