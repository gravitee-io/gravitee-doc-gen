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

func (e *exampleValidation) OnAttribute(property string, attribute *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *schema.VisitContext) {
	if e.skip {
		return
	}
	defaultValue := schema.GetConstantOrDefault(attribute, visitCtx.AutoDefaultBooleans)
	var example any
	if len(attribute.Examples) > 0 {
		example = attribute.Examples[0]
	}
	if schema.IsRequired(property, parent) && defaultValue == nil && example == nil {
		path := e.path
		path = append(path, property)
		e.AddErr(fmt.Sprintf(
			"property %s of type %s is required but do not have any examples, it must be set",
			strings.Join(path, "."),
			schema.GetType(attribute)))
	}
}
func (e *exampleValidation) OnObjectStart(property string, object *jsonschema.Schema, visitCtx *schema.VisitContext) {
	e.path = append(e.path, property)
}
func (e *exampleValidation) OnObjectEnd(*schema.VisitContext) {
	e.path = e.path[:len(e.path)-1]
}
func (e *exampleValidation) OnArrayStart(property string, array *jsonschema.Schema, itemTypeIsObject bool, ctx *schema.VisitContext) {
	// no op
}
func (e *exampleValidation) OnArrayEnd(itemTypeIsObject bool) {
	// no op
}
func (e *exampleValidation) OnOneOfStart(oneOf *jsonschema.Schema, parent *jsonschema.Schema, visitCtx *schema.VisitContext) {
	e.skip = true
}
func (e *exampleValidation) OnOneOfEnd() {
	e.skip = false
}
