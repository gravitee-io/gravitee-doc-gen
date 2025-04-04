package extensions

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

type ExamplesSettings struct {
	Examples []ReadmeExample `json:",inline"`
}
type ReadmeExample struct {
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Language    string         `json:"language"`
	Filename    string         `json:"filename"`
	UseDefaults bool           `json:"useDefaults"`
	Template    string         `json:"template"`
	OneOfFilter OneOfFilter    `json:"oneOfFilter"`
	Properties  map[string]any `json:"properties"`
}

type OneOfFilter struct {
	Path           []string       `json:"path"`
	Discriminators map[string]any `json:"discriminators"`
	Index          int            `json:"index"`
}

type DiscriminatorValue struct {
	Property string `json:"property"`
	Value    string `json:"value"`
}
type Compiler struct {
}

const (
	ReadmeExamplesExtension = "x-readme-examples"
)

func (c Compiler) Compile(_ jsonschema.CompilerContext, m map[string]interface{}) (jsonschema.ExtSchema, error) {
	if e, ok := m[ReadmeExamplesExtension]; ok {
		if object, ok := e.([]interface{}); ok {
			extSchemaItem, err := util.AnyArrayToStructArray[ReadmeExample](object)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("failed to parse %s json: %v", ReadmeExamplesExtension, err))
			}
			return ExamplesSettings{Examples: extSchemaItem}, nil
		}
	}
	return nil, nil
}

func (c ExamplesSettings) Validate(_ jsonschema.ValidationContext, _ interface{}) error {
	// we don't validate payloads, no implementation required
	return nil
}
