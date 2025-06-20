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

package common

import (
	"slices"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type SchemaToNodeTreeVisitor struct {
	KeepAllOneOfAttributes bool
	skipAttributes         bool
	oneOfCount             int
	lastDiscriminatorValue map[string]any
	oneOfIndex             int
}

func (v *SchemaToNodeTreeVisitor) OnAttribute(
	ctx *visitor.VisitContext,
	property string,
	attribute *jsonschema.Schema,
	parent *jsonschema.Schema) *visitor.Attribute {
	if v.skipAttributes {
		return nil
	}

	value := visitor.GetValueOrFirstExample(attribute, ctx)

	if value != nil {
		nodeAttribute := visitor.NewAttribute(property, parent)
		nodeAttribute.Value = value
		nodeAttribute.Title = attribute.Title
		nodeAttribute.Description = attribute.Description
		nodeAttribute.Type = schema.GetType(attribute)
		nodeAttribute.IsOneOfProperty = ctx.PeekOneOf().Present
		nodeAttribute.IsOneOfDiscriminator = ctx.PeekOneOf().IsDiscriminator(property)
		nodeAttribute.Enums = getEnums(attribute, property, ctx.PeekOneOf())
		nodeAttribute.Default = visitor.GetValueOrFirstExample(attribute, ctx)
		return nodeAttribute
	}
	return nil
}

func (v *SchemaToNodeTreeVisitor) OnObjectStart(
	_ *visitor.VisitContext,
	property string,
	object *jsonschema.Schema) *visitor.Object {
	o := visitor.NewObject(property)
	o.Description = object.Description
	o.Title = object.Title
	return o
}

func (v *SchemaToNodeTreeVisitor) OnObjectEnd(_ *visitor.VisitContext) {
	// no op
}

func (v *SchemaToNodeTreeVisitor) OnArrayStart(
	ctx *visitor.VisitContext,
	property string,
	array *jsonschema.Schema,
	itemTypeIsObject bool) (*visitor.Array, []visitor.Value) {
	newArray := visitor.NewArray(property)
	newArray.Title = array.Title
	newArray.Description = array.Description
	newArray.ItemType = schema.GetTypeItem(array)
	if !itemTypeIsObject {
		values := make([]visitor.Value, 0)
		value := visitor.GetValueOrFirstExample(array, ctx)
		if defaults, ok := value.([]interface{}); ok {
			for _, val := range defaults {
				values = append(values, visitor.NewValue(val))
			}
		}
		return newArray, values
	} else if len(array.Examples) > 0 {
		values := make([]visitor.Value, 0)
		for _, example := range array.Examples {
			values = append(values, visitor.NewValue(example))
		}
		return newArray, values
	}
	return newArray, nil
}

func (v *SchemaToNodeTreeVisitor) OnArrayEnd(*visitor.VisitContext, bool) {
	// no op
}

func (v *SchemaToNodeTreeVisitor) OnOneOfStart(*visitor.VisitContext, *jsonschema.Schema) {
	// no op
}

func (v *SchemaToNodeTreeVisitor) OnOneOf(ctx *visitor.VisitContext, oneOf *jsonschema.Schema, _ *jsonschema.Schema) {
	filter := ctx.OneOfFilter()

	// no filter, just check if we can keep all properties or not
	if filter.IsZero() {
		v.oneOfCount++
		// case where the first oneOf is kept
		if !v.KeepAllOneOfAttributes && v.oneOfCount > 1 {
			v.skipAttributes = true
		}
		return
	}

	// if the required path is not the current one
	if len(filter.Path) > 0 && !slices.Equal(ctx.NodeStack().GetAncestorProperty(), filter.Path) {
		v.skipAttributes = true
		return
	}

	// initialize map of discriminator values
	if v.lastDiscriminatorValue == nil {
		v.lastDiscriminatorValue = make(map[string]any)
	}

	// for each configured discriminator value
	for property, expectedValue := range filter.Discriminators {
		// if the one of contains the discriminator property
		if discriminatorSchema, ok := oneOf.Properties[property]; ok {
			// get the discriminator value
			actualValue := v.updateOneOfLatestDiscriminatorValue(ctx, discriminatorSchema, property)
			// skip attributes if no match
			if v.skipAttributes = actualValue != expectedValue; v.skipAttributes {
				return
			}
		} else {
			// can't tell if it matches => skip it
			v.oneOfIndex = 0
			v.skipAttributes = true
			return
		}
	}

	// oneOf match but there can be several, check it is the one we really want
	if filter.Index != v.oneOfIndex {
		v.skipAttributes = true
		return
	}
}

func (v *SchemaToNodeTreeVisitor) updateOneOfLatestDiscriminatorValue(
	ctx *visitor.VisitContext,
	discriminatorProperty *jsonschema.Schema,
	key string) any {
	actualValue := visitor.GetValue(discriminatorProperty, ctx)
	// reset if new value
	if actualValue != v.lastDiscriminatorValue[key] {
		v.oneOfIndex = 0
	} else {
		v.oneOfIndex++
	}
	v.lastDiscriminatorValue[key] = actualValue
	return actualValue
}

func (v *SchemaToNodeTreeVisitor) OnOneOfEnd(*visitor.VisitContext) {
	v.oneOfCount = 0
	v.skipAttributes = false
}

func getEnums(attribute *jsonschema.Schema, property string, oneOf visitor.OneOfDescriptor) []any {
	if oneOf.IsZero() {
		return attribute.Enum
	}
	enums := make([]any, 0)
	for _, spec := range oneOf.Specs {
		if spec.Property == property {
			for _, v := range spec.Values {
				if !slices.Contains(enums, v) {
					enums = append(enums, v)
				}
			}
		}
	}
	return enums
}
