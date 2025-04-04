package options

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
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

	root, err := schema.Compile(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	options := Options{Sections: []Section{{
		Attributes: make([]Attribute, 0),
	}}}

	ctx := schema.VisitContext{}
	schema.Visit(root, &options, &ctx)

	return chunks.Processed{Data: options}, err
}

func (options *Options) OnAttribute(name string, att *jsonschema.Schema, parent *jsonschema.Schema, _ *schema.VisitContext) {

	attribute := Attribute{
		Property:    name,
		Name:        att.Title,
		Type:        schema.GetType(att),
		Constraint:  getConstraint(att),
		Required:    schema.IsRequired(name, parent),
		Default:     schema.GetConstantOrDefault(att),
		IsConstant:  isConstant(att),
		EL:          isEL(att.Extensions),
		Secret:      isSecret(att.Extensions),
		Description: att.Description,
		Enums:       getEnums(att.Enum),
	}
	options.AddAttribute(attribute)
}

func (options *Options) OnObject(_ string, object *jsonschema.Schema, visitCtx *schema.VisitContext) {

	objectType := "object"
	if visitCtx.CurrentOneOf.Present {
		objectType = "oneOf"
	}
	options.Add(Section{
		Title: object.Title,
		Type:  objectType,
	})
	if visitCtx.CurrentOneOf.Present {
		oneOf := visitCtx.CurrentOneOf
		options.AddAttribute(Attribute{
			Name:     util.Title(oneOf.Property),
			Property: oneOf.Property,
			Type:     oneOf.Type,
			Required: true,
			Enums:    oneOf.Values,
			OneOf:    oneOf,
		})
	}
}

func (options *Options) OnArray(_ string, array *jsonschema.Schema) {
	options.Add(Section{
		Title: array.Title,
		Type:  "array",
	})
}

func (options *Options) OnOneOf(object *jsonschema.Schema, _ *jsonschema.Schema, visitCtx *schema.VisitContext) {
	discriminator := schema.GetConstantOrDefault(object.Properties[visitCtx.CurrentOneOf.Property])
	options.Add(Section{Title: object.Title, OneOf: visitCtx.CurrentOneOf, DiscriminatedBy: discriminator})
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
	return isTrue(extensions, schema.SecretExtension)
}

func isSecret(extensions map[string]jsonschema.ExtSchema) bool {
	return isTrue(extensions, schema.SecretExtension)
}

func isTrue(extensions map[string]jsonschema.ExtSchema, ext string) bool {
	if s, ok := extensions[ext]; ok {
		return bool(s.(schema.BoolValueSchema))
	}
	return false
}

func getEnums(enum []interface{}) []string {
	if enum == nil {
		return nil
	}
	return iToa(enum)
}

func iToa(enum []interface{}) []string {
	result := make([]string, len(enum))
	for i := range enum {
		result[i] = fmt.Sprintf("%v", enum[i])
	}
	return result
}
