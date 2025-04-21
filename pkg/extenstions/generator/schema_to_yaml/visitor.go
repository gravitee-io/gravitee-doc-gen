package schema_to_yaml

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/visitor"
	"gopkg.in/yaml.v3"
)

type yamlLine struct {
	visitor.Attribute
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

func (v *toYamlVisitor) OnObjectStart(object visitor.Object, level int) {
	if level == 0 || v.inArray {
		return
	}
	attribute := visitor.NewAttribute(object.Name(), nil)
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

func (v *toYamlVisitor) OnObjectEnd(visitor.Object, int) {
	// no op
}

func (v *toYamlVisitor) OnArrayStart(array visitor.Array, level int) {
	attribute := visitor.NewAttribute(array.Name(), nil)
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

func (v *toYamlVisitor) OnArrayItem(parent visitor.Array, value visitor.Value, level int) {
	attribute := visitor.NewAttribute("", nil)
	attribute.Value = encode(value.Value, attribute.Type == "string")
	v.Lines = append(v.Lines, yamlLine{
		Pad:        v.pad(level) + 2, // to have array padded in regard to their container name (not compulsory)
		ArrayStart: v.arrayFirstItem,
		Attribute:  *attribute,
	})
}

func (v *toYamlVisitor) OnArrayEnd(visitor.Array, int) {
	v.inArray = false
	v.arrayFirstItem = false
}

func (v *toYamlVisitor) OnAttribute(attribute visitor.Attribute, level int) {
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
