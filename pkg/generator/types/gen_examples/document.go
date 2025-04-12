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

func (b *DocumentBuilder) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) {
	if !b.skipAttributes && schema.IsAttribute(attribute) && !attribute.Deprecated {
		if value := schema.GetDefaultOrFirstExample(attribute, ctx); value != nil {
			b.Add(property, value, ctx)
		}
	}
}

func (b *DocumentBuilder) OnObjectStart(ctx *schema.VisitContext, property string, _ *jsonschema.Schema) {
	b.Add(property, schema.NewObject(property), ctx)
}

func (b *DocumentBuilder) OnArrayStart(ctx *schema.VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) {
	b.Add(property, schema.NewArray(property), ctx)
	if itemTypeIsObject {
		b.Add("", schema.NewObject(property), ctx)
	} else {
		value := schema.GetDefaultOrFirstExample(array, ctx)
		if items, ok := value.([]interface{}); ok {
			for _, v := range items {
				b.Add("", v, ctx)
			}
		} else {
			b.Add("", value, ctx)
		}
	}
}

func (b *DocumentBuilder) OnObjectEnd(ctx *schema.VisitContext) {
	ctx.NodeStack().Pop()
}

func (b *DocumentBuilder) OnArrayEnd(ctx *schema.VisitContext, itemTypeIsObject bool) {
	if itemTypeIsObject {
		ctx.NodeStack().Pop()
	}
	ctx.NodeStack().Pop()
}

func (b *DocumentBuilder) OnOneOfStart(ctx *schema.VisitContext, oneOf *jsonschema.Schema, _ *jsonschema.Schema) {

	if !slices.Equal(ctx.NodeStack().GetAncestorNames(), b.example.OneOfFilter.Path) {
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

func (b *DocumentBuilder) Add(name string, value interface{}, ctx *schema.VisitContext) {
	ctx.NodeStack().Add(ctx, name, value)
}
