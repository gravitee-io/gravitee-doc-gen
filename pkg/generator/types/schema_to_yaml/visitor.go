package schema_to_yaml

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"gopkg.in/yaml.v3"
)

type yamlLine struct {
	schema.Attribute
	Pad        int
	ArrayStart bool
	Property   string
}

type toYamlVisitor struct {
	Lines          []yamlLine
	padding        int
	inArray        bool
	arrayFirstItem bool
}

func (v *toYamlVisitor) OnObjectStart(object schema.Object, level int) {
	if level == 0 || v.inArray {
		return
	}
	attribute := schema.NewAttribute(object.Name(), nil)
	attribute.Title = object.Title
	attribute.Description = object.Description
	v.Lines = append(v.Lines, yamlLine{
		Pad:       v.pad(level),
		Property:  object.Name(),
		Attribute: *attribute,
	})
}

func (v *toYamlVisitor) pad(level int) int {
	// level 0 is just the root level, all the rest is level>0 so to compute padding we need to lower the level by 1
	return v.padding * (level - 1)
}

func (v *toYamlVisitor) OnObjectEnd(schema.Object, int) {
	// no op
}

func (v *toYamlVisitor) OnArrayStart(array schema.Array, level int) {
	attribute := schema.NewAttribute(array.Name(), nil)
	attribute.Title = array.Title
	attribute.Description = array.Description
	v.Lines = append(v.Lines, yamlLine{
		Pad:       v.pad(level),
		Property:  array.Name(),
		Attribute: *attribute,
	})
	v.inArray = true
	v.arrayFirstItem = true
}

func (v *toYamlVisitor) OnArrayItem(parent schema.Array, value schema.Value, level int) {
	attribute := schema.NewAttribute("", nil)
	attribute.Value = encode(value.Value, attribute.Type == "string")
	v.Lines = append(v.Lines, yamlLine{
		Pad:        v.pad(level) + 2, // to have array padded in regard to their container name (not compulsory)
		ArrayStart: v.arrayFirstItem,
		Attribute:  *attribute,
	})
}

func (v *toYamlVisitor) OnArrayEnd(schema.Array, int) {
	v.inArray = false
	v.arrayFirstItem = false
}

func (v *toYamlVisitor) OnAttribute(attribute schema.Attribute, level int) {
	pad := v.pad(level)
	// OnAttribute() means we are in an object, we remove the object level padding if already in an array
	// Avoid extra padding (not compulsory but looks better)
	if v.inArray {
		pad -= v.padding
	}
	// mandatory padding for object second and following attributes attribute in an array
	if v.inArray && !v.arrayFirstItem {
		pad += 2
	}
	attribute.Value = encode(attribute.Value, attribute.Type == "string")
	v.Lines = append(v.Lines, yamlLine{
		Pad:        pad,
		Property:   attribute.Name(),
		ArrayStart: v.arrayFirstItem,
		Attribute:  attribute,
	})
	v.arrayFirstItem = false
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
