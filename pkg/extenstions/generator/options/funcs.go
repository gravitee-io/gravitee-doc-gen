// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"errors"
	"math/big"
	"strconv"
	"strings"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema/extensions"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := chunks.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}

	schemaFile := chunks.GetString(chunk, "schema")
	schemaFileExists := util.FileExists(schemaFile)
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
	schemaFile := chunks.GetString(chunk, "schema")

	root, err := schema.CompileWithExtensions(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	options := Options{Sections: []Section{{
		Attributes: make([]Attribute, 0),
	}}}

	ctx := visitor.NewVisitContext(true, true).WithStack(visitor.NewObject(""))
	visitor.Visit(ctx, &options, root)

	return chunks.Processed{Data: options}, err
}

func (o *Options) OnAttribute(
	ctx *visitor.VisitContext,
	property string,
	attribute *jsonschema.Schema,
	parent *jsonschema.Schema) *visitor.Attribute {
	att := Attribute{
		Property:    property,
		Name:        attribute.Title,
		Type:        schema.GetType(attribute),
		TypeItem:    schema.GetTypeItem(attribute),
		Constraint:  getConstraint(attribute),
		Required:    schema.IsRequired(property, parent),
		Default:     visitor.GetValue(attribute, ctx),
		IsConstant:  isConstant(attribute),
		EL:          isEL(attribute),
		Secret:      isSecret(attribute),
		Description: attribute.Description,
		Enums:       enums(attribute),
	}
	o.AddAttribute(att)
	return nil
}

func enums(attribute *jsonschema.Schema) []any {
	if len(attribute.Enum) > 0 {
		return attribute.Enum
	}
	if schema.IsArray(attribute) && schema.IsAttribute(schema.Items(attribute)) && len(schema.Items(attribute).Enum) > 0 {
		return schema.Items(attribute).Enum
	}
	return nil
}

func (o *Options) OnObjectStart(ctx *visitor.VisitContext, _ string, object *jsonschema.Schema) *visitor.Object {
	objectType := "object"
	o.Add(Section{
		Title: object.Title,
		Type:  objectType,
	})

	return nil
}

func (o *Options) OnArrayStart(
	_ *visitor.VisitContext,
	_ string,
	array *jsonschema.Schema,
	itemTypeIsObject bool) (*visitor.Array, []visitor.Value) {
	if itemTypeIsObject {
		o.Add(Section{
			Title: array.Title,
			Type:  "array",
		})
	}
	return nil, nil
}

func (o *Options) OnOneOfStart(ctx *visitor.VisitContext, _ *jsonschema.Schema) {
	if ctx.PeekOneOf().Present {
		specs := ctx.PeekOneOf().Specs
		for _, spec := range specs {
			o.AddAttribute(Attribute{
				Name:     util.TitleCaseToTitle(util.Title(spec.Property)),
				Property: spec.Property,
				Type:     spec.Type,
				Required: true,
				Enums:    spec.Values,
				OneOf:    ctx.PeekOneOf(),
			})
		}
	}
}

func (o *Options) OnOneOf(ctx *visitor.VisitContext, oneOf *jsonschema.Schema, _ *jsonschema.Schema) {
	specs := ctx.PeekOneOf().Specs
	discriminatedBy := make(map[string]any)
	list := visitor.NewSchemaPropertyList(oneOf)
	for _, spec := range specs {
		value := visitor.GetValue(list.Get(spec.Property), ctx)
		discriminatedBy[spec.Property] = value
	}

	o.Add(Section{Title: oneOf.Title, OneOf: ctx.PeekOneOf(), DiscriminatedBy: discriminatedBy})
}

func (o *Options) OnObjectEnd(*visitor.VisitContext) {
	// no op
}

func (o *Options) OnArrayEnd(*visitor.VisitContext, bool) {
	// no op
}

func (o *Options) OnOneOfEnd(*visitor.VisitContext) {
	// no op
}

func (o *Options) Add(section Section) {
	section.Attributes = make([]Attribute, 0)
	o.Sections = append(o.Sections, section)
	o.current += 1
}

func (o *Options) AddAttribute(attribute Attribute) {
	o.Sections[o.current].Add(attribute)
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
	case att.ReadOnly:
		constraints = append(constraints, "read-only")
	case att.WriteOnly:
		constraints = append(constraints, "write-only")
	}

	return strings.Join(constraints, ", ")
}

func startBound(inclusive *big.Rat, exclusive *big.Rat) string {
	switch {
	case inclusive == nil && exclusive == nil:
		return "[-Inf"
	case exclusive != nil:
		return "(" + ratToString(exclusive)
	default:
		return "[" + ratToString(inclusive)
	}
}

func endBound(inclusive *big.Rat, exclusive *big.Rat) string {
	switch {
	case inclusive == nil && exclusive == nil:
		return "+Inf]"
	case exclusive != nil:
		return ratToString(exclusive) + ")"
	default:
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

func getGioConfig(att *jsonschema.Schema) *extensions.GioConfigSchema {
	if gioConfig, exists := att.Extensions[extensions.GioConfigExtension]; exists {
		if ext, ok := gioConfig.(*extensions.GioConfigSchema); ok {
			return ext
		}
	}
	return &extensions.GioConfigSchema{}
}
