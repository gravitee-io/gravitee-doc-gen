package raw_examples

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/core"
	"github.com/gravitee-io-labs/readme-gen/pkg/examples"
	"gopkg.in/yaml.v3"
	"os"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	return examples.TypeValidator(chunk, &examples.RawExampleProvider{})
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	return examples.ProcessAllExamples(chunk, &examples.RawExampleProvider{}, readCodeExampleAndValidate)
}

func readCodeExampleAndValidate(chunk config.Chunk, spec examples.ExampleSpec) (string, error) {

	rawSpec := spec.(examples.RawExampleSpec)

	bytes, err := os.ReadFile(rawSpec.File)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to read code example file %s: %v", rawSpec.File, err))
	}

	codeToEmbed := string(bytes)
	var jsonToValidate string
	if rawSpec.Language == examples.YAML {
		if converted, err := yamlToJson(codeToEmbed); err == nil {
			jsonToValidate = converted
		} else {
			panic(fmt.Sprintf("cannot yaml to json with example %v: %v", rawSpec, err))
		}
	} else {
		jsonToValidate = codeToEmbed
	}
	validationSchema, _, err := examples.CompileSchema(rawSpec, chunk)
	if err != nil {
		return "", err
	}
	if err := examples.ValidateJson(jsonToValidate, validationSchema, rawSpec.File); err != nil {
		return "", err
	}

	return codeToEmbed, nil
}

func yamlToJson(jsonToValidate string) (string, error) {
	y := core.Unstructured{}
	err := yaml.Unmarshal([]byte(jsonToValidate), &y)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(y)
	return string(b), nil
}
