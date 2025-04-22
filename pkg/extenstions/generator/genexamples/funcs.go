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
	"errors"
	"strings"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	bexamples "github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/examples"
	visitor2 "github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	provider := &examples.GenExampleProvider{}
	ok, err := examples.TypeValidator(chunk, provider)
	if err != nil {
		return false, err
	}
	if ok {
		return validateExamples(chunk, provider)
	}
	return ok, nil
}

func validateExamples(chunk config.Chunk, provider *examples.GenExampleProvider) (bool, error) {
	err := examples.LoadConfig(chunk, provider)
	if err != nil {
		return true, err
	}

	for _, spec := range provider.ExampleSpecs() {
		s, _, _ := examples.CompileSchema(spec, chunk)
		ev := &exampleValidation{
			errors: make([]string, 0),
		}
		visitor2.Visit(visitor2.NewVisitContext(false, true).WithStack(visitor2.NewObject("")), ev, s)
		if len(ev.errors) > 0 {
			return false, errors.New(strings.Join(ev.errors, "\n"))
		}
	}

	return true, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	return examples.ProcessAllExamples(chunk, &examples.GenExampleProvider{}, yieldCodeExampleAndValidate)
}

func yieldCodeExampleAndValidate(chunk config.Chunk, spec examples.ExampleSpec) (string, error) {
	genSpec := util.As[examples.GenExampleSpec](spec)

	validationSchema, _, _ := examples.CompileSchema(genSpec, chunk)

	object := visitor2.NewObject("")

	ctx := visitor2.NewVisitContext(true, true).
		WithStack(object).
		WithOneOfFilter(genSpec.OneOfFilter)

	visitor2.Visit(ctx, &common.SchemaToNodeTreeVisitor{}, validationSchema)

	ref, _ := genSpec.TemplateFromRef()

	codeToEmbed, err := ref.Language.Serialize(object.Fields)
	if err != nil {
		panic(err)
	}

	var jsonToValidate string
	json := bexamples.JSON
	if t, _ := genSpec.TemplateFromRef(); t.Language != json {
		inJSON, err := json.Serialize(object.Fields)
		if err != nil {
			panic(err)
		}
		jsonToValidate = inJSON
	} else {
		jsonToValidate = codeToEmbed
	}

	if err := examples.ValidateJson(jsonToValidate, validationSchema, "generated example from schema"); err != nil {
		return "", err
	}

	return codeToEmbed, nil
}
