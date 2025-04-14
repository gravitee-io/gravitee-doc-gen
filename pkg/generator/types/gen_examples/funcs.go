package gen_examples

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/examples"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"strings"
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
		schema.Visit(schema.NewVisitContext(false, true).WithStack(schema.NewObject("")), ev, s)
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

	genSpec := spec.(examples.GenExampleSpec)

	validationSchema, _, _ := examples.CompileSchema(genSpec, chunk)

	object := schema.NewObject("")

	ctx := schema.NewVisitContext(true, true).
		WithStack(object).
		WithOneOfFilter(genSpec.OneOfFilter)

	schema.Visit(ctx, &types.SchemaVisitor{}, validationSchema)

	ref, _ := genSpec.TemplateFromRef()

	codeToEmbed, err := ref.Language.Serialize(object.Fields)
	if err != nil {
		panic(err)
	}

	var jsonToValidate string
	json := examples.JSON
	if t, _ := genSpec.TemplateFromRef(); t.Language != json {
		inJson, err := json.Serialize(object.Fields)
		if err != nil {
			panic(err)
		}
		jsonToValidate = inJson
	} else {
		jsonToValidate = codeToEmbed
	}

	if err := examples.ValidateJson(jsonToValidate, validationSchema, "generated example from schema"); err != nil {
		return "", err
	}

	return codeToEmbed, nil

}
