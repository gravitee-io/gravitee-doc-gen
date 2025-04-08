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

	ctx := schema.VisitContext{QueueNodes: true}
	schema.Visit(root, &options, &ctx)

	return chunks.Processed{Data: options}, err
}

func (options *Options) OnAttribute(property string, attribute *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *schema.VisitContext) {
	att := Attribute{
		Property:    property,
		Name:        attribute.Title,
		Type:        schema.GetType(attribute),
		Constraint:  getConstraint(attribute),
		Required:    schema.IsRequired(property, parent),
		Default:     schema.GetConstantOrDefault(attribute),
		IsConstant:  isConstant(attribute),
		EL:          isEL(attribute.Extensions),
		Secret:      isSecret(attribute.Extensions),
		Description: attribute.Description,
		Enums:       attribute.Enum,
	}
	options.AddAttribute(att)
}

func (options *Options) OnObjectStart(_ string, object *jsonschema.Schema, visitCtx *schema.VisitContext) {

	objectType := "object"
	if visitCtx.CurrentOneOf.Present {
		objectType = "oneOf"
	}
	options.Add(Section{
		Title: object.Title,
		Type:  objectType,
	})

	if visitCtx.CurrentOneOf.Present {
		specs := visitCtx.CurrentOneOf.Specs
		for _, spec := range specs {
			options.AddAttribute(Attribute{
				Name:     util.TitleCaseToTitle(util.Title(spec.Property)),
				Property: spec.Property,
				Type:     spec.Type,
				Required: true,
				Enums:    spec.Values,
				OneOf:    visitCtx.CurrentOneOf,
			})
		}

	}
}

func (options *Options) OnArrayStart(_ string, array *jsonschema.Schema, _ bool) {
	options.Add(Section{
		Title: array.Title,
		Type:  "array",
	})
}

func (options *Options) OnOneOfStart(schema *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *schema.VisitContext) {
	specs := visitCtx.CurrentOneOf.Specs
	discriminatedBy := make(map[string]any)
	for _, spec := range specs {
		value := schema.GetConstantOrDefault(oneOf.Properties[spec.Property])
		discriminatedBy[spec.Property] = value
	}

	options.Add(Section{Title: oneOf.Title, OneOf: visitCtx.CurrentOneOf, DiscriminatedBy: discriminatedBy})
}

func (options *Options) OnObjectEnd() {
	//no op
}

func (options *Options) OnArrayEnd(bool) {
	// no op
}

func (options *Options) OnOneOfEnd() {
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
	switch {
	case att.Pattern != nil:
		return att.Pattern.String()
	case att.Minimum != nil:
		return att.Minimum.String()
	case att.Maximum != nil:
		return att.Maximum.String()
	}
	return ""
}

func isConstant(att *jsonschema.Schema) bool {
	return att.Constant != nil
}

func isEL(extensions map[string]jsonschema.ExtSchema) bool {
	return isTrue(extensions, ext.SecretExtension)
}

func isSecret(extensions map[string]jsonschema.ExtSchema) bool {
	return isTrue(extensions, ext.SecretExtension)
}

func isTrue(extensions map[string]jsonschema.ExtSchema, name string) bool {
	if s, ok := extensions[name]; ok {
		return bool(s.(ext.BoolValueSchema))
	}
	return false
}
