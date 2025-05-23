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

package rawexamples

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	bexamples "github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/examples"
	"gopkg.in/yaml.v3"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	return examples.TypeValidator(chunk, &examples.RawExampleProvider{})
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	return examples.ProcessAllExamples(chunk, &examples.RawExampleProvider{}, readCodeExampleAndValidate)
}

func readCodeExampleAndValidate(chunk config.Chunk, spec examples.ExampleSpec) (string, error) {
	rawSpec := util.As[examples.RawExampleSpec](spec)

	bytes, err := os.ReadFile(rawSpec.File)
	if err != nil {
		return "", fmt.Errorf("failed to read code example file %s: %w", rawSpec.File, err)
	}

	codeToEmbed := string(bytes)
	var jsonToValidate = codeToEmbed
	if rawSpec.Language == bexamples.YAML {
		if converted, err := yamlToJSON(codeToEmbed); err == nil {
			jsonToValidate = converted
		} else {
			panic(fmt.Sprintf("cannot convert yaml to json with example %v: %v", rawSpec, err))
		}
	}
	validationSchema, _, err := examples.CompileSchema(rawSpec, chunk)
	if err != nil {
		return "", err
	}
	if err := examples.ValidateJson(jsonToValidate, validationSchema, rawSpec.File); err != nil {
		return "", err
	}

	tmpl, _ := rawSpec.TemplateFromRef()

	if rawSpec.Language == bexamples.JSON && tmpl.Language == bexamples.YAML {
		codeToEmbed, err = jsonToYaml(codeToEmbed)
		if err != nil {
			return "", err
		}
	} else if rawSpec.Language == bexamples.YAML && tmpl.Language == bexamples.JSON {
		codeToEmbed, err = yamlToJSON(codeToEmbed)
		if err != nil {
			return "", err
		}
	}
	return codeToEmbed, nil
}

func jsonToYaml(jsonBytes string) (string, error) {
	j := util.Unstructured{}
	err := json.Unmarshal([]byte(jsonBytes), &j)
	if err != nil {
		return "", err
	}
	converted, err := yaml.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(converted), nil
}

func yamlToJSON(yamlBytes string) (string, error) {
	y := util.Unstructured{}
	err := yaml.Unmarshal([]byte(yamlBytes), &y)
	if err != nil {
		return "", err
	}
	converted, err := json.Marshal(y)
	if err != nil {
		return "", err
	}
	return string(converted), nil
}
