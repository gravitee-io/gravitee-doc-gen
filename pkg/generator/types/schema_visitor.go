package types

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

type SchemaVisitor struct {
	KeepAllOneOfAttributes bool
	skipAttributes         bool
	oneOfCount             int
	lastDiscriminatorValue map[string]any
	oneOfIndex             int
}

func (v *SchemaVisitor) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *schema.Attribute {

	if v.skipAttributes {
		return nil
	}

	value := schema.GetDefaultOrFirstExample(attribute, ctx)

	if value != nil {
		nodeAttribute := schema.NewAttribute(property, parent)
		nodeAttribute.Value = value
		nodeAttribute.Title = attribute.Title
		nodeAttribute.Description = attribute.Description
		nodeAttribute.Type = schema.GetType(attribute)
		nodeAttribute.IsOneOfProperty = ctx.CurrentOneOf().Present
		nodeAttribute.IsOneOfDiscriminator = ctx.CurrentOneOf().IsDiscriminator(property)
		nodeAttribute.Enums = getEnums(attribute, property, ctx.CurrentOneOf())
		nodeAttribute.Default = schema.GetDefaultOrFirstExample(attribute, ctx)
		return nodeAttribute
	}
	return nil
}

func (v *SchemaVisitor) OnObjectStart(*schema.VisitContext, string, *jsonschema.Schema) *schema.Object {

	return nil
}

func (v *SchemaVisitor) OnObjectEnd(_ *schema.VisitContext) {
	// no op
}

func (v *SchemaVisitor) OnArrayStart(_ *schema.VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) (*schema.Array, []schema.Value) {
	newArray := schema.NewArray(property)
	newArray.Title = array.Title
	newArray.Description = array.Description
	newArray.ItemType = schema.GetTypeItem(array)
	if !itemTypeIsObject {
		values := make([]schema.Value, 0)
		for _, val := range array.Default.([]interface{}) {
			values = append(values, schema.NewValue(val))
		}
		return newArray, values
	}
	return newArray, nil
}

func (v *SchemaVisitor) OnArrayEnd(*schema.VisitContext, bool) {
	// no op
}

func (v *SchemaVisitor) OnOneOf(ctx *schema.VisitContext, oneOf *jsonschema.Schema, _ *jsonschema.Schema) {

	filter := ctx.OneOfFilter()

	// no filter,just check if we can keep all properties or not
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

	// for each configured discriminator values
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

func (v *SchemaVisitor) updateOneOfLatestDiscriminatorValue(ctx *schema.VisitContext, discriminatorProperty *jsonschema.Schema, key string) any {
	actualValue := schema.GetConstantOrDefault(discriminatorProperty, ctx)
	// reset if new value
	if actualValue != v.lastDiscriminatorValue[key] {
		v.oneOfIndex = 0
	} else {
		v.oneOfIndex++
	}
	v.lastDiscriminatorValue[key] = actualValue
	return actualValue
}

func (v *SchemaVisitor) OnOneOfEnd(*schema.VisitContext) {
	v.oneOfCount = 0
	v.skipAttributes = false
}

func getEnums(attribute *jsonschema.Schema, property string, oneOf schema.OneOf) []any {
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
