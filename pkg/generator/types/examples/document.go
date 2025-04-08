package examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
)

type DocumentBuilder struct {
	language               Language
	root                   *Object
	example                ExampleSpec
	stack                  []Stackable
	skipAttributes         bool
	oneOfIndex             int
	lastDiscriminatorValue map[string]interface{}
}

func NewDocumentBuilder(example ExampleSpec) *DocumentBuilder {
	root := NewObject("")
	return &DocumentBuilder{
		language:               From(example.Language),
		stack:                  []Stackable{root},
		root:                   root,
		example:                example,
		lastDiscriminatorValue: make(map[string]interface{}),
	}
}

func (b *DocumentBuilder) OnAttribute(property string, attribute *jsonschema.Schema, _ *jsonschema.Schema, ctx *schema.VisitContext) {
	if !b.skipAttributes && schema.IsAttribute(attribute) && !attribute.Deprecated {
		if value := b.getExampleValue(attribute, ctx); value != nil {
			b.Add(property, value)
		}
	}
}

func (b *DocumentBuilder) OnObjectStart(property string, _ *jsonschema.Schema, _ *schema.VisitContext) {
	b.Add(property, NewObject(property))
}

func (b *DocumentBuilder) OnArrayStart(property string, array *jsonschema.Schema, itemTypeIsObject bool, ctx *schema.VisitContext) {
	b.Add(property, NewArray(property))
	if itemTypeIsObject {
		b.Add("", NewObject(property))
	} else {
		value := b.getExampleValue(array, ctx)
		if items, ok := value.([]interface{}); ok {
			for _, i := range items {
				b.Add("", i)
			}
		} else {
			b.Add("", value)
		}
	}
}

func (b *DocumentBuilder) OnObjectEnd() {
	b.Pop()
}

func (b *DocumentBuilder) OnArrayEnd(itemTypeIsObject bool) {
	if itemTypeIsObject {
		b.Pop()
	}
	b.Pop()
}

func (b *DocumentBuilder) OnOneOfStart(oneOf *jsonschema.Schema, _ *jsonschema.Schema, ctx *schema.VisitContext) {

	if !slices.Equal(b.getAncestors(), b.example.OneOfFilter.Path) {
		b.skipAttributes = true
		return
	}

	discriminators := b.example.OneOfFilter.Discriminators
	for key, expectedValue := range discriminators {
		if discriminatorProperty, ok := oneOf.Properties[key]; ok {
			actualValue := schema.GetConstantOrDefault(discriminatorProperty, ctx.AutoDefaultBooleans)
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

func (b *DocumentBuilder) OnOneOfEnd() {
	b.skipAttributes = false
}

func (b *DocumentBuilder) Marshall() (string, error) {
	return b.language.Serialize(b.root.Fields)
}
func (b *DocumentBuilder) MarshallWithLanguage(language Language) (string, error) {
	return language.Serialize(b.root.Fields)
}

func (b *DocumentBuilder) Add(name string, value interface{}) {
	stackable := b.Peek()
	if stackable.IsArray() {
		array := stackable.(*Array)
		array.Items = append(array.Items, value)
	} else {
		stackable.(*Object).Fields[name] = value
	}

	if val, ok := value.(Stackable); ok {
		b.Push(val)
	}
}

func (b *DocumentBuilder) Push(value Stackable) {
	b.stack = append(b.stack, value)
}

func (b *DocumentBuilder) Peek() Stackable {
	if len(b.stack) == 0 {
		return nil
	}
	return b.stack[len(b.stack)-1]
}

func (b *DocumentBuilder) Pop() {
	// check if current needs to be removed
	var property string
	var remove bool
	stackable := b.Peek()
	if stackable != nil && stackable.IsEmpty() {
		property = stackable.Name()
		remove = true
	}

	b.stack = removeLast[Stackable](b.stack)

	if remove {
		if last := b.Peek(); last.IsArray() {
			array := last.(*Array)
			array.Items = removeLast(array.Items)
		} else {
			object := last.(*Object)
			delete(object.Fields, property)
		}
	}
}

func removeLast[T any](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}
	return slice[:len(slice)-1]
}

func (b *DocumentBuilder) getAncestors() []string {
	ancestors := make([]string, len(b.stack)-1)
	// skip root
	for i := 1; i < len(b.stack); i++ {
		ancestors[i-1] = b.stack[i].Name()
	}
	return ancestors
}

func (b *DocumentBuilder) getExampleValue(att *jsonschema.Schema, ctx *schema.VisitContext) any {
	var value any
	def := schema.GetConstantOrDefault(att, ctx.AutoDefaultBooleans)
	if b.example.UseSchemaDefaults {
		value = def
		if value == nil && len(att.Examples) > 0 {
			value = att.Examples[0]
		}
	} else {
		if len(att.Examples) > 0 {
			value = att.Examples[0]
		}
		if value == nil {
			value = def
		}
	}

	return value
}
