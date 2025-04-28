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

package genexamples

import (
	"fmt"
	"strings"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/schema"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type exampleValidation struct {
	errors []string
	skip   bool
}

func (e *exampleValidation) AddErr(err string) {
	e.errors = append(e.errors, err)
}

func (e *exampleValidation) OnAttribute(
	ctx *visitor.VisitContext,
	property string,
	attribute *jsonschema.Schema,
	parent *jsonschema.Schema) *visitor.Attribute {
	if e.skip {
		return nil
	}
	defaultValue := visitor.GetValueOrFirstExample(attribute, ctx)
	if schema.IsRequired(property, parent) && defaultValue == nil {
		e.AddErr(fmt.Sprintf(
			"property %s of type %s is required but do not have any examples, it must be set",
			strings.Join(ctx.NodeStack().Properties(), "."),
			schema.GetType(attribute)))
	}
	return nil
}
func (e *exampleValidation) OnObjectStart(*visitor.VisitContext, string, *jsonschema.Schema) *visitor.Object {
	// no op
	return nil
}
func (e *exampleValidation) OnObjectEnd(*visitor.VisitContext) {
	// no op
}
func (e *exampleValidation) OnArrayStart(
	*visitor.VisitContext,
	string,
	*jsonschema.Schema,
	bool) (*visitor.Array, []visitor.Value) {
	// no op
	return nil, nil
}

func (e *exampleValidation) OnArrayEnd(*visitor.VisitContext, bool) {
	// no op
}
func (e *exampleValidation) OnOneOfStart(*visitor.VisitContext, *jsonschema.Schema) {
	// no op
}
func (e *exampleValidation) OnOneOf(*visitor.VisitContext, *jsonschema.Schema, *jsonschema.Schema) {
	e.skip = true
}
func (e *exampleValidation) OnOneOfEnd(*visitor.VisitContext) {
	e.skip = false
}
