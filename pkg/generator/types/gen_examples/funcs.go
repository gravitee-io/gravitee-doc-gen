package gen_examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/examples"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	return examples.TypeValidator(chunk, &examples.GenExampleProvider{})
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	return examples.ProcessAllExamples(chunk, &examples.GenExampleProvider{}, yieldCodeExampleAndValidate)
}

func yieldCodeExampleAndValidate(chunk config.Chunk, spec examples.ExampleSpec) (string, error) {

	genSpec := spec.(examples.GenExampleSpec)

	validationSchema, _, _ := examples.CompileSchema(genSpec, chunk)
	object := NewDocumentBuilder(genSpec)

	schema.Visit(validationSchema, object, &schema.VisitContext{})

	codeToEmbed, err := object.Marshall()
	if err != nil {
		panic(err)
	}

	var jsonToValidate string
	if t, _ := genSpec.TemplateFromRef(); t.Language != examples.JSON {
		inJson, err := object.MarshallWithLanguage(examples.JSON)
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
