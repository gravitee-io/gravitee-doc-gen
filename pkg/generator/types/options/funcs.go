package options

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	ext "github.com/gravitee-io-labs/readme-gen/pkg/schema/extensions"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"math/big"
	"strconv"
	"strings"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := common.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}

	schemaFile := common.GetFile(chunk, "schema")
	schemaFileExists := common.FileExists(schemaFile)
	if chunk.Required && !schemaFileExists {
		return false, errors.New("schema file not found")
	}

	compiler := jsonschema.NewCompiler()
	_, err = compiler.Compile(schemaFile)
	if err != nil {
		return false, err
	}

	return tmplExists && schemaFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {

	schemaFile := common.GetFile(chunk, "schema")

	root, err := schema.CompileWithExtensions(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	options := Options{Sections: []Section{{
		Attributes: make([]Attribute, 0),
	}}}

	ctx := schema.NewVisitContext(true, true)
	schema.Visit(ctx, &options, root)

	return chunks.Processed{Data: options}, err
}

func (options *Options) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) func() any {
	att := Attribute{
		Property:    property,
		Name:        attribute.Title,
		Type:        schema.GetType(attribute),
		TypeItem:    schema.GetTypeItem(attribute),
		Constraint:  getConstraint(attribute),
		Required:    schema.IsRequired(property, parent),
		Default:     schema.GetConstantOrDefault(attribute, ctx),
		IsConstant:  isConstant(attribute),
		EL:          isEL(attribute),
		Secret:      isSecret(attribute),
		Description: attribute.Description,
		Enums:       attribute.Enum,
	}
	options.AddAttribute(att)
	return nil
}

func (options *Options) OnObjectStart(ctx *schema.VisitContext, property string, object *jsonschema.Schema) {

	objectType := "object"
	if ctx.CurrentOneOf().Present {
		objectType = "oneOf"
	}
	options.Add(Section{
		Title: object.Title,
		Type:  objectType,
	})

	if ctx.CurrentOneOf().Present {
		specs := ctx.CurrentOneOf().Specs
		for _, spec := range specs {
			options.AddAttribute(Attribute{
				Name:     util.TitleCaseToTitle(util.Title(spec.Property)),
				Property: spec.Property,
				Type:     spec.Type,
				Required: true,
				Enums:    spec.Values,
				OneOf:    ctx.CurrentOneOf(),
			})
		}

	}
}

func (options *Options) OnArrayStart(ctx *schema.VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) func() any {
	if !schema.IsAttribute(array.Items.(*jsonschema.Schema)) {
		options.Add(Section{
			Title: array.Title,
			Type:  "array",
		})
	}
	return nil
}

func (options *Options) OnOneOfStart(ctx *schema.VisitContext, oneOf *jsonschema.Schema, parent *jsonschema.Schema) {
	specs := ctx.CurrentOneOf().Specs
	discriminatedBy := make(map[string]any)
	for _, spec := range specs {
		value := schema.GetConstantOrDefault(oneOf.Properties[spec.Property], ctx)
		discriminatedBy[spec.Property] = value
	}

	options.Add(Section{Title: oneOf.Title, OneOf: ctx.CurrentOneOf(), DiscriminatedBy: discriminatedBy})
}

func (options *Options) OnObjectEnd(*schema.VisitContext) {
	//no op
}

func (options *Options) OnArrayEnd(*schema.VisitContext, bool) {
	// no op
}

func (options *Options) OnOneOfEnd(*schema.VisitContext) {
	// no op
}

func (options *Options) Add(section Section) {
	section.Attributes = make([]Attribute, 0)
	options.Sections = append(options.Sections, section)
	options.current += 1
}

func (options *Options) AddAttribute(attribute Attribute) {
	options.Sections[options.current].Add(attribute)
}

func (s *Section) Add(attribute Attribute) {
	s.Attributes = append(s.Attributes, attribute)
}

func getConstraint(att *jsonschema.Schema) string {
	constraints := make([]string, 0)

	switch {
	case att.Pattern != nil:
		constraints = append(constraints, att.Pattern.String())
	case att.MinItems >= 0 || att.MaxItems >= 0:
		constraints = append(constraints, "["+valueOrZero(att.MinItems))
		constraints = append(constraints, valueOrInfinity(att.MaxItems)+"]")
		fallthrough
	case att.UniqueItems:
		constraints = append(constraints, "unique")
	case att.Minimum != nil || att.Maximum != nil || att.ExclusiveMinimum != nil || att.ExclusiveMaximum != nil:
		constraints = append(constraints, startBound(att.Minimum, att.ExclusiveMinimum))
		constraints = append(constraints, endBound(att.Maximum, att.ExclusiveMaximum))
	case att.MinLength >= 0 || att.MaxLength >= 0:
		constraints = append(constraints, "["+valueOrZero(att.MinLength))
		constraints = append(constraints, valueOrInfinity(att.MaxLength)+"]")
	}

	return strings.Join(constraints, ", ")
}

func startBound(inclusive *big.Rat, exclusive *big.Rat) string {
	if inclusive == nil && exclusive == nil {
		return "[-Inf"
	} else if exclusive != nil {
		return "(" + ratToString(exclusive)
	} else {
		return "[" + ratToString(inclusive)
	}
}

func endBound(inclusive *big.Rat, exclusive *big.Rat) string {
	if inclusive == nil && exclusive == nil {
		return "+Inf]"
	} else if exclusive != nil {
		return ratToString(exclusive) + ")"
	} else {
		return ratToString(inclusive) + "]"
	}
}

func ratToString(number *big.Rat) string {
	if number.IsInt() {
		return number.Num().String()
	}
	n, _ := number.FloatPrec()
	return number.FloatString(n)
}

func valueOrInfinity(value int) string {
	if value < 0 {
		return "+Inf"
	}
	return strconv.Itoa(value)
}

func valueOrZero(value int) string {
	if value < 0 {
		return "0"
	}
	return strconv.Itoa(value)
}

func isConstant(att *jsonschema.Schema) bool {
	return att.Constant != nil
}

func isEL(att *jsonschema.Schema) bool {
	return getGioConfig(att).El
}

func isSecret(att *jsonschema.Schema) bool {
	return getGioConfig(att).Secret
}

func getGioConfig(att *jsonschema.Schema) *ext.GioConfigSchema {
	if gioConfig, ok := att.Extensions[ext.GioConfigExtension]; ok {
		return gioConfig.(*ext.GioConfigSchema)
	}
	return &ext.GioConfigSchema{}
}
