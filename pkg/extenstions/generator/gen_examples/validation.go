package gen_examples

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/schema"
	visitor2 "github.com/gravitee-io-labs/readme-gen/pkg/extenstions/common/visitor"
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

func (e *exampleValidation) OnAttribute(ctx *visitor2.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) *visitor2.Attribute {
	if e.skip {
		return nil
	}
	defaultValue := visitor2.GetDefaultOrFirstExample(attribute, ctx)
	if schema.IsRequired(property, parent) && defaultValue == nil {
		e.AddErr(fmt.Sprintf(
			"property %s of type %s is required but do not have any examples, it must be set",
			strings.Join(ctx.NodeStack().Properties(), "."),
			visitor2.GetType(attribute)))
	}
	return nil
}
func (e *exampleValidation) OnObjectStart(*visitor2.VisitContext, string, *jsonschema.Schema) *visitor2.Object {
	//no op
	return nil
}
func (e *exampleValidation) OnObjectEnd(*visitor2.VisitContext) {
	// no op
}
func (e *exampleValidation) OnArrayStart(*visitor2.VisitContext, string, *jsonschema.Schema, bool) (*visitor2.Array, []visitor2.Value) {
	// no op
	return nil, nil
}

func (e *exampleValidation) OnArrayEnd(*visitor2.VisitContext, bool) {
	// no op
}
func (e *exampleValidation) OnOneOf(*visitor2.VisitContext, *jsonschema.Schema, *jsonschema.Schema) {
	e.skip = true
}
func (e *exampleValidation) OnOneOfEnd(*visitor2.VisitContext) {
	e.skip = false
}
