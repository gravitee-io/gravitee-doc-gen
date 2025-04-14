package gen_examples

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"strings"
)

type exampleValidation struct {
	errors []string
	skip   bool
}

func (e *exampleValidation) AddErr(err string) {
	e.errors = append(e.errors, err)
}

func (e *exampleValidation) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *schema.Attribute {
	if e.skip {
		return nil
	}
	defaultValue := schema.GetDefaultOrFirstExample(attribute, ctx)
	if schema.IsRequired(property, parent) && defaultValue == nil {
		e.AddErr(fmt.Sprintf(
			"property %s of type %s is required but do not have any examples, it must be set",
			strings.Join(ctx.NodeStack().Properties(), "."),
			schema.GetType(attribute)))
	}
	return nil
}
func (e *exampleValidation) OnObjectStart(*schema.VisitContext, string, *jsonschema.Schema) {
	//no op
}
func (e *exampleValidation) OnObjectEnd(*schema.VisitContext) {
	// no op
}
func (e *exampleValidation) OnArrayStart(*schema.VisitContext, string, *jsonschema.Schema, bool) []schema.Value {
	// no op
	return nil
}

func (e *exampleValidation) OnArrayEnd(*schema.VisitContext, bool) {
	// no op
}
func (e *exampleValidation) OnOneOfStart(*schema.VisitContext, *jsonschema.Schema, *jsonschema.Schema) {
	e.skip = true
}
func (e *exampleValidation) OnOneOfEnd(*schema.VisitContext) {
	e.skip = false
}
