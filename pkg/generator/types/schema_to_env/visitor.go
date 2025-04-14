package schema_to_env

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
	"strings"
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
	for v := range s {
		slice = append(slice, v)
	}
	return slice
}

type section struct {
	Title       string
	Description string
	Variables   []variable
}

func (s *section) AddVariable(v variable) {
	s.Variables = append(s.Variables, v)
}

type variable struct {
	Title       string
	Description string
	Env         string
	JVM         string
	Type        string
	Default     any
	When        map[string]set
	Enums       []any
}

type oneOfProperty struct {
	variable
}

func newSchemaVisitor(indexPlaceholder string) schemaVisitor {
	first := &section{
		Variables: make([]variable, 0),
	}
	return schemaVisitor{
		Sections:            []*section{first},
		indexPlaceholder:    indexPlaceholder,
		currentSection:      first,
		oneOfProperties:     make(map[string]oneOfProperty),
		oneOfDiscriminators: make([]string, 0),
	}
}

type schemaVisitor struct {
	inArray             bool
	firstArrayItem      bool
	oneOfProperties     map[string]oneOfProperty
	oneOfStarted        bool
	oneOfDiscriminators []string
	currentSection      *section
	Sections            []*section
	indexPlaceholder    string
}

func (v *schemaVisitor) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *schema.Attribute {

	if v.oneOfStarted && !v.isOneOfDiscriminator(property, ctx.CurrentOneOf()) {
		v.addOneOfProperty(ctx, property, attribute, parent)
		return nil
	}
	if v.oneOfStarted && slices.Contains(v.oneOfDiscriminators, property) {
		return nil
	}

	if v.oneOfStarted {
		v.oneOfDiscriminators = append(v.oneOfDiscriminators, property)
	}

	v.currentSection.AddVariable(variable{
		Title:       attribute.Title,
		Description: attribute.Description,
		Env:         v.getEnv(ctx, property),
		JVM:         v.getJvm(ctx, property),
		Type:        schema.GetType(attribute),
		Default:     schema.GetDefaultOrFirstExample(attribute, ctx),
		Enums:       getEnums(attribute, property, ctx.CurrentOneOf()),
	})

	if v.firstArrayItem {
		v.firstArrayItem = false
	}

	return nil
}

func (v *schemaVisitor) OnObjectStart(ctx *schema.VisitContext, _ string, object *jsonschema.Schema) {
	if v.inArray {
		return
	}
	if len(ctx.NodeStack().GetAncestorProperty()) == 0 {
		section := &section{
			Title:       object.Title,
			Description: object.Description,
		}
		v.Sections = append(v.Sections, section)
		v.currentSection = section
	}
}

func (v *schemaVisitor) OnObjectEnd(ctx *schema.VisitContext) {
	if v.oneOfStarted && ctx.CurrentOneOf().IsZero() {
		for _, oneOf := range v.oneOfProperties {
			v.currentSection.AddVariable(oneOf.variable)
		}
		v.oneOfStarted = false
		v.oneOfDiscriminators = make([]string, 0)
		v.oneOfProperties = make(map[string]oneOfProperty)
	}
	if v.inArray {
		v.firstArrayItem = true
	}
}

func (v *schemaVisitor) OnArrayStart(ctx *schema.VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) []schema.Attribute {
	if len(ctx.NodeStack().GetAncestorProperty()) == 0 && itemTypeIsObject {
		section := &section{
			Title:       array.Title,
			Description: array.Description,
		}
		v.Sections = append(v.Sections, section)
		v.currentSection = section
	}
	v.inArray = true
	v.firstArrayItem = true

	if !itemTypeIsObject {
		v.currentSection.AddVariable(variable{
			Title:       array.Title,
			Description: array.Description,
			Env:         v.getEnv(ctx, property),
			JVM:         v.getJvm(ctx, property),
			Type:        schema.GetTypeItem(array),
			Default:     schema.GetDefaultOrFirstExample(array, ctx),
		})
	}

	return nil
}

func (v *schemaVisitor) OnArrayEnd(*schema.VisitContext, bool) {
	v.firstArrayItem = false
	v.inArray = false
}

func (v *schemaVisitor) OnOneOfStart(*schema.VisitContext, *jsonschema.Schema, *jsonschema.Schema) {
	v.oneOfStarted = true
}

func (v *schemaVisitor) OnOneOfEnd(*schema.VisitContext) {
	// no op
}

func (v *schemaVisitor) getJvm(ctx *schema.VisitContext, property string) string {
	return v.joinStackNodes(ctx, property, fmt.Sprintf("[%s]", v.indexPlaceholder), ".", strings.ToLower)
}
func (v *schemaVisitor) getEnv(ctx *schema.VisitContext, property string) string {
	return v.joinStackNodes(ctx, property, v.indexPlaceholder, "_", strings.ToUpper)
}

func (v *schemaVisitor) joinStackNodes(ctx *schema.VisitContext, property string, arraySyntax string, sep string, format func(string) string) string {
	nodes := ctx.NodeStack().Nodes()
	elements := make([]string, 0)
	elements = append(elements, format("gravitee"))
	for i := 1; i < len(nodes); i++ {
		switch nodes[i].Type() {
		case schema.ArrayNode:
			elements = append(elements, format(nodes[i].Name()))
			elements = append(elements, arraySyntax)
		default:
			elements = append(elements, format(nodes[i].Name()))
		}
	}
	elements = append(elements, format(property))
	return strings.Join(elements, sep)
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
			variable: variable{
				When: make(map[string]set),
			},
		}
		oneOfProp.Default = schema.GetDefaultOrFirstExample(attribute, ctx)
		oneOfProp.Title = attribute.Title
		oneOfProp.Description = attribute.Description
		oneOfProp.Type = schema.GetType(attribute)
		oneOfProp.Env = v.getEnv(ctx, property)
		oneOfProp.JVM = v.getJvm(ctx, property)
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
			s = set{}
			s[value] = true
			oneOfProperty.When[spec.Property] = s
		}
	}
}

func (v *schemaVisitor) isOneOfDiscriminator(property string, oneOf schema.OneOf) bool {
	for _, spec := range oneOf.Specs {
		if spec.Property == property {
			return true
		}
	}
	return false
}
