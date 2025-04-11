package schema_to_yaml

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
	"slices"
)

type set map[any]bool

func fromSlice(slice []any) set {
	set := set{}
	for _, v := range slice {
		set[v] = true
	}
	return set
}

func (s set) toSlice() []any {
	slice := make([]any, 0, len(s))
	for v, _ := range s {
		slice = append(slice, v)
	}
	return slice
}

type baseLine struct {
	Title       string
	Description string
	Type        string
	Value       any
	When        map[string]set
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
	oneOfStarted        bool
	oneOfDiscriminators []string
}

func (s *schemaVisitor) OnAttribute(property string, attribute *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *schema.VisitContext) {

	if s.oneOfStarted && !s.isOneOfProperty(property, visitCtx.CurrentOneOf) {
		s.addOneOfProperty(property, attribute, parent, visitCtx)
		return
	}

	if s.oneOfStarted && slices.Contains(s.oneOfDiscriminators, property) {
		return
	}

	if s.oneOfStarted {
		s.oneOfDiscriminators = append(s.oneOfDiscriminators, property)
	}

	s.Lines = append(s.Lines, line{
		baseLine: baseLine{
			Title:       attribute.Title,
			Description: attribute.Description,
			Type:        schema.GetType(attribute),
			Value:       encode(getValue(attribute, visitCtx), schema.GetType(attribute) == "string"),
			Enums:       getEnums(attribute, property, visitCtx.CurrentOneOf),
		},
		Pad:        s.pad,
		Property:   property,
		ArrayStart: s.firstArrayItem,
	})
	if s.firstArrayItem {
		s.pad += 2
		s.firstArrayItem = false
	}
}

func (s *schemaVisitor) OnObjectStart(property string, object *jsonschema.Schema, _ *schema.VisitContext) {
	if s.inArray {
		return
	}
	s.Lines = append(s.Lines, line{
		Pad:      s.pad,
		Property: property,
		baseLine: baseLine{
			Title:       object.Title,
			Description: object.Description,
		},
	})
	s.pad += s.padding
}

func (s *schemaVisitor) OnObjectEnd(visitCtx *schema.VisitContext) {
	if s.oneOfStarted && visitCtx.CurrentOneOf.IsZero() {
		for property, oneOf := range s.oneOfProperties {
			s.Lines = append(s.Lines, oneOf.toLine(property, s.pad, s.firstArrayItem))
		}
		s.oneOfStarted = false
		s.oneOfDiscriminators = make([]string, 0)
		s.oneOfProperties = make(map[string]oneOfProperty)
	}
	if s.inArray {
		s.pad -= 2
		s.firstArrayItem = true
	} else {
		s.pad -= s.padding
	}
}

func (s *schemaVisitor) OnArrayStart(property string, array *jsonschema.Schema, itemTypeIsObject bool, _ *schema.VisitContext) {
	s.Lines = append(s.Lines, line{
		baseLine: baseLine{
			Title:       array.Title,
			Description: array.Description,
		},
		Pad:      s.pad,
		Property: property,
	})
	s.inArray = true
	s.firstArrayItem = true
	s.pad += 2
	if !itemTypeIsObject {
		def := array.Default
		array := def.([]interface{})
		for _, v := range array {
			s.Lines = append(s.Lines, line{
				Pad:        s.pad,
				ArrayStart: true,
				baseLine: baseLine{
					Value: v,
				},
			})
		}
	}
}

func (s *schemaVisitor) OnArrayEnd(itemTypeIsObject bool) {
	s.firstArrayItem = false
	s.inArray = false
	if itemTypeIsObject {
		s.pad -= 4
	} else {
		s.pad -= 2
	}
}

func (s *schemaVisitor) OnOneOfStart(_ *jsonschema.Schema, _ *jsonschema.Schema, _ *schema.VisitContext) {
	s.oneOfStarted = true
}

func (s *schemaVisitor) OnOneOfEnd() {
	// no op
}

func getValue(attribute *jsonschema.Schema, visitCtx *schema.VisitContext) any {
	value := schema.GetConstantOrDefault(attribute, visitCtx.AutoDefaultBooleans)
	if value == nil && len(attribute.Examples) > 0 {
		return attribute.Examples[0]
	}
	return value
}

func getEnums(attribute *jsonschema.Schema, property string, oneOf schema.OneOf) []any {
	if oneOf.IsZero() {
		return fromSlice(attribute.Enum).toSlice()
	}
	enums := make([]any, 0)
	for _, spec := range oneOf.Specs {
		if spec.Property == property {
			enums = append(enums, spec.Values...)
		}
	}
	return fromSlice(enums).toSlice()
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

func (s *schemaVisitor) addOneOfProperty(property string, attribute *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *schema.VisitContext) {
	if s.oneOfProperties == nil {
		s.oneOfProperties = make(map[string]oneOfProperty)
	}
	var update oneOfProperty
	if oneOfProp, ok := s.oneOfProperties[property]; ok {
		s.updateWhen(visitCtx, parent, &oneOfProp)
		update = oneOfProp
	} else {
		oneOfProp = oneOfProperty{
			baseLine: baseLine{
				When: make(map[string]set),
			},
		}
		oneOfProp.Value = getValue(attribute, visitCtx)
		oneOfProp.Title = attribute.Title
		oneOfProp.Enums = fromSlice(attribute.Enum).toSlice()
		s.updateWhen(visitCtx, parent, &oneOfProp)
		update = oneOfProp
	}
	s.oneOfProperties[property] = update

}

func (s *schemaVisitor) updateWhen(visitCtx *schema.VisitContext, parent *jsonschema.Schema, oneOfProperty *oneOfProperty) {
	for _, spec := range visitCtx.CurrentOneOf.Specs {
		value := getValue(parent.Properties[spec.Property], visitCtx)
		if s, ok := oneOfProperty.When[spec.Property]; ok {
			s[value] = true
			oneOfProperty.When[spec.Property] = s
		} else {
			s = set{}
			s[value] = true
			oneOfProperty.When[spec.Property] = s
		}
	}
}

func (s *schemaVisitor) isOneOfProperty(property string, oneOf schema.OneOf) bool {
	for _, spec := range oneOf.Specs {
		if spec.Property == property {
			return true
		}
	}
	return false
}
