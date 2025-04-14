package gen_examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/examples"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

type DocumentBuilder struct {
	language               examples.Language
	root                   *schema.Object
	example                examples.GenExampleSpec
	skipAttributes         bool
	oneOfIndex             int
	lastDiscriminatorValue map[string]interface{}
}

func NewDocumentBuilder(example examples.GenExampleSpec) *DocumentBuilder {
	root := schema.NewObject("")
	ref, _ := example.TemplateFromRef()
	return &DocumentBuilder{
		language:               ref.Language,
		root:                   root,
		example:                example,
		lastDiscriminatorValue: make(map[string]interface{}),
	}
}

func (b *DocumentBuilder) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *schema.Attribute {
	if !b.skipAttributes && schema.IsAttribute(attribute) && !attribute.Deprecated {
		if value := schema.GetDefaultOrFirstExample(attribute, ctx); value != nil {
			return schema.NewAttribute(property, value)
		}
	}
	return nil
}

func (b *DocumentBuilder) OnObjectStart(*schema.VisitContext, string, *jsonschema.Schema) {
	// no op
}

func (b *DocumentBuilder) OnArrayStart(ctx *schema.VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) []schema.Value {
	if !itemTypeIsObject {
		example := schema.GetDefaultOrFirstExample(array, ctx)
		if example != nil {
			attributes := make([]schema.Value, 0)
			if items, isArray := example.([]any); isArray {
				for _, item := range items {
					attributes = append(attributes, schema.NewValue(item))
				}
			} else {
				attributes = append(attributes, schema.NewValue(example))
			}
			return attributes
		}
	}
	return nil
}

func (b *DocumentBuilder) OnObjectEnd(*schema.VisitContext) {
	// no op
}

func (b *DocumentBuilder) OnArrayEnd(*schema.VisitContext, bool) {
	// no op
}

func (b *DocumentBuilder) OnOneOfStart(ctx *schema.VisitContext, oneOf *jsonschema.Schema, _ *jsonschema.Schema) {

	if !slices.Equal(ctx.NodeStack().GetAncestorProperty(), b.example.OneOfFilter.Path) {
		b.skipAttributes = true
		return
	}

	discriminators := b.example.OneOfFilter.Discriminators
	for key, expectedValue := range discriminators {
		if discriminatorProperty, ok := oneOf.Properties[key]; ok {
			actualValue := schema.GetConstantOrDefault(discriminatorProperty, ctx)
			if actualValue != b.lastDiscriminatorValue[key] {
				b.oneOfIndex = 0
			} else {
				b.oneOfIndex++
			}
			b.lastDiscriminatorValue[key] = actualValue
			if b.skipAttributes = actualValue != expectedValue; b.skipAttributes {
				return
			}
		} else {
			b.oneOfIndex = 0
			b.skipAttributes = true
			return
		}
	}

	if b.example.OneOfFilter.Index != b.oneOfIndex {
		b.skipAttributes = true
		return
	}

}

func (b *DocumentBuilder) OnOneOfEnd(*schema.VisitContext) {
	b.skipAttributes = false
}

func (b *DocumentBuilder) Marshall() (string, error) {
	return b.language.Serialize(b.root.Fields)
}

func (b *DocumentBuilder) MarshallWithLanguage(language examples.Language) (string, error) {
	return language.Serialize(b.root.Fields)
}
