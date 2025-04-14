package schema_to_yaml

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
	"slices"
)

type Set map[any]bool

func fromSlice(slice []any) Set {
	set := Set{}
	for _, v := range slice {
		set[v] = true
	}
	return set
}

func (s Set) toSlice() []any {
	slice := make([]any, 0, len(s))
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}

type baseLine struct {
	Title       string
	Description string
	Type        string
	Value       any
	When        map[string]Set
	Enums       []any
}

type line struct {
	baseLine
	Pad        int
	ArrayStart bool
	Property   string
}

type oneOfProperty struct {
	baseLine
}

func (o oneOfProperty) toLine(property string, pad int, arrayStart bool) line {
	return line{
		baseLine:   o.baseLine,
		Pad:        pad,
		ArrayStart: arrayStart,
		Property:   property,
	}
}

func newSchemaVisitor() schemaVisitor {
	return schemaVisitor{
		Lines:               make([]line, 0),
		oneOfProperties:     make(map[string]oneOfProperty),
		oneOfDiscriminators: make([]string, 0),
	}
}

type schemaVisitor struct {
	Lines               []line
	padding             int
	pad                 int
	inArray             bool
	firstArrayItem      bool
	oneOfProperties     map[string]oneOfProperty
	oneOfDiscriminators []string
}

func (v *schemaVisitor) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *schema.Attribute {

	if ctx.CurrentOneOf().Present && !ctx.CurrentOneOf().IsDiscriminator(property) {
		v.addOneOfProperty(ctx, property, attribute, parent)
		return nil
	}

	if ctx.CurrentOneOf().Present && slices.Contains(v.oneOfDiscriminators, property) {
		return nil
	}

	if ctx.CurrentOneOf().Present {
		v.oneOfDiscriminators = append(v.oneOfDiscriminators, property)
	}

	v.Lines = append(v.Lines, line{
		baseLine: baseLine{
			Title:       attribute.Title,
			Description: attribute.Description,
			Type:        schema.GetType(attribute),
			Value:       encode(schema.GetDefaultOrFirstExample(attribute, ctx), schema.GetType(attribute) == "string"),
			Enums:       getEnums(attribute, property, ctx.CurrentOneOf()),
		},
		Pad:        v.pad,
		Property:   property,
		ArrayStart: v.firstArrayItem,
	})
	if v.firstArrayItem {
		v.pad += 2
		v.firstArrayItem = false
	}
	return nil
}

func (v *schemaVisitor) OnObjectStart(_ *schema.VisitContext, property string, object *jsonschema.Schema) {
	if v.inArray {
		return
	}
	v.Lines = append(v.Lines, line{
		Pad:      v.pad,
		Property: property,
		baseLine: baseLine{
			Title:       object.Title,
			Description: object.Description,
		},
	})
	v.pad += v.padding
}

func (v *schemaVisitor) OnObjectEnd(ctx *schema.VisitContext) {
	if !ctx.CurrentOneOf().Present {
		for property, oneOf := range v.oneOfProperties {
			v.Lines = append(v.Lines, oneOf.toLine(property, v.pad, v.firstArrayItem))
		}
		v.oneOfDiscriminators = make([]string, 0)
		v.oneOfProperties = make(map[string]oneOfProperty)
	}
	if v.inArray {
		v.pad -= 2
		v.firstArrayItem = true
	} else {
		v.pad -= v.padding
	}
}

func (v *schemaVisitor) OnArrayStart(ctx *schema.VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) []schema.Value {
	v.Lines = append(v.Lines, line{
		baseLine: baseLine{
			Title:       array.Title,
			Description: array.Description,
		},
		Pad:      v.pad,
		Property: property,
	})
	v.inArray = true
	v.firstArrayItem = true
	v.pad += 2
	if !itemTypeIsObject {
		def := array.Default
		array := def.([]interface{})
		for _, i := range array {
			v.Lines = append(v.Lines, line{
				Pad:        v.pad,
				ArrayStart: true,
				baseLine: baseLine{
					Value: i,
				},
			})
		}
	}
	return nil
}

func (v *schemaVisitor) OnArrayEnd(_ *schema.VisitContext, itemTypeIsObject bool) {
	v.firstArrayItem = false
	v.inArray = false
	if itemTypeIsObject {
		v.pad -= 4
	} else {
		v.pad -= 2
	}
}

func (v *schemaVisitor) OnOneOfStart(*schema.VisitContext, *jsonschema.Schema, *jsonschema.Schema) {
	// no op
}

func (v *schemaVisitor) OnOneOfEnd(*schema.VisitContext) {
	// no op
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

func encode(value any, isString bool) any {
	if isString && value != nil {
		node := &yaml.Node{Kind: yaml.ScalarNode}
		_ = node.Encode(value)
		if node.Style == yaml.SingleQuotedStyle {
			node.Style = yaml.DoubleQuotedStyle
		}
		s, _ := yaml.Marshal(node)
		return string(s)
	}
	return value
}

func (v *schemaVisitor) addOneOfProperty(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) {
	if v.oneOfProperties == nil {
		v.oneOfProperties = make(map[string]oneOfProperty)
	}
	var update oneOfProperty
	if oneOfProp, ok := v.oneOfProperties[property]; ok {
		v.updateWhen(ctx, parent, &oneOfProp)
		update = oneOfProp
	} else {
		oneOfProp = oneOfProperty{
			baseLine: baseLine{
				When: make(map[string]Set),
			},
		}
		oneOfProp.Value = encode(schema.GetDefaultOrFirstExample(attribute, ctx), schema.GetType(attribute) == "string")
		oneOfProp.Title = attribute.Title
		oneOfProp.Enums = fromSlice(attribute.Enum).toSlice()
		v.updateWhen(ctx, parent, &oneOfProp)
		update = oneOfProp
	}
	v.oneOfProperties[property] = update

}

func (v *schemaVisitor) updateWhen(ctx *schema.VisitContext, parent *jsonschema.Schema, oneOfProperty *oneOfProperty) {
	for _, spec := range ctx.CurrentOneOf().Specs {
		value := schema.GetDefaultOrFirstExample(parent.Properties[spec.Property], ctx)
		if s, ok := oneOfProperty.When[spec.Property]; ok {
			s[value] = true
			oneOfProperty.When[spec.Property] = s
		} else {
			s = Set{}
			s[value] = true
			oneOfProperty.When[spec.Property] = s
		}
	}
}
