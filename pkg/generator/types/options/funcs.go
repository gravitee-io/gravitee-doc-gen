package options

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"slices"
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

	schema.Visit(root, &options)

	return chunks.Processed{Data: options}, err
}

func (options *Options) OnAttribute(name string, att *jsonschema.Schema, parent *jsonschema.Schema) {
	attribute := Attribute{
		Property:    name,
		Name:        att.Title,
		Type:        getType(att),
		Constraint:  getConstraint(att),
		Required:    isRequired(name, parent),
		Default:     getDefault(att),
		EL:          isEL(att.Extensions),
		Secret:      isSecret(att.Extensions),
		Description: att.Description,
		Enums:       getEnums(att.Enum),
	}
	options.AddAttribute(attribute)
}

func (options *Options) OnObject(_ string, object *jsonschema.Schema) {
	options.Add(Section{
		Title: object.Title,
	})
}

func (options *Options) OnArray(_ string, array *jsonschema.Schema) {
	options.Add(Section{
		Title:   array.Title,
		Comment: "array",
	})
}

func (options *Options) OnOneOf(object *jsonschema.Schema) {
	options.Add(Section{Title: object.Title})
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

func getType(att *jsonschema.Schema) string {
	if len(att.Types) == 0 {
		return ""
	}
	t := att.Types[0]
	if att.Enum != nil && len(att.Enum) > 0 {
		return "enum (" + t + ")"
	}
	return t
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

func isRequired(name string, parent *jsonschema.Schema) bool {
	required := parent.Required
	return required != nil && slices.Contains(required, name)
}

func getDefault(att *jsonschema.Schema) string {
	if att.Constant != nil {
		return fmt.Sprintf("%v (const)", att.Constant)
	}
	if att.Default == nil {
		return ""
	}
	return fmt.Sprintf("%v", att.Default)
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
