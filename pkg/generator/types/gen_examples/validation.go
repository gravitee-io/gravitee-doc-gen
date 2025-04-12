package gen_examples

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"strings"
)

type exampleValidation struct {
	errors []string
	path   []string
	skip   bool
}

func (e *exampleValidation) AddErr(err string) {
	e.errors = append(e.errors, err)
}

func (e *exampleValidation) OnAttribute(ctx *schema.VisitContext, property string, attribute *jsonschema.Schema, parent *jsonschema.Schema) {
	if e.skip {
		return
	}
	defaultValue := schema.GetDefaultOrFirstExample(attribute, ctx)
	if schema.IsRequired(property, parent) && defaultValue == nil {
		path := e.path
		path = append(path, property)
		e.AddErr(fmt.Sprintf(
			"property %s of type %s is required but do not have any examples, it must be set",
			strings.Join(path, "."),
			schema.GetType(attribute)))
	}
}
func (e *exampleValidation) OnObjectStart(ctx *schema.VisitContext, property string, object *jsonschema.Schema) {
	e.path = append(e.path, property)
}
func (e *exampleValidation) OnObjectEnd(*schema.VisitContext) {
	e.path = e.path[:len(e.path)-1]
}
func (e *exampleValidation) OnArrayStart(ctx *schema.VisitContext, property string, array *jsonschema.Schema, itemTypeIsObject bool) {
	// no op
}
func (e *exampleValidation) OnArrayEnd(ctx *schema.VisitContext, itemTypeIsObject bool) {
	// no op
}
func (e *exampleValidation) OnOneOfStart(visitCtx *schema.VisitContext, oneOf *jsonschema.Schema, parent *jsonschema.Schema) {
	e.skip = true
}
func (e *exampleValidation) OnOneOfEnd(*schema.VisitContext) {
	e.skip = false
}
