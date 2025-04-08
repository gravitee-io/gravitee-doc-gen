package examples

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
	"os"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := common.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return tmplExists, err
	}

	schemaFile := common.GetFile(chunk, "schema")
	schemaFileExists := common.FileExists(schemaFile)
	if chunk.Required && !schemaFileExists {
		return schemaFileExists, errors.New("schema file not found")
	}

	examplesFile := common.GetDataFile(chunk)
	examplesFileExists := common.FileExists(examplesFile)

	if chunk.Required && !examplesFileExists {
		return examplesFileExists, errors.New("example file not found")
	}

	cfg, err := loadConfig(chunk)
	if err != nil {
		return true, err
	}

	b, err := validateSpecs(cfg, chunk)
	if err != nil {
		return b, err
	}

	return tmplExists && schemaFileExists, nil
}

func validateSpecs(cfg Config, chunk config.Chunk) (bool, error) {
	for _, spec := range cfg.Specs {

		err := spec.Check()
		if err != nil {
			return false, err
		}

		err = checkCodeTemplateFile(spec, chunk)
		if err != nil {
			return false, err
		}

		schemaFile, _, err := compileSchema(spec, chunk)
		if err != nil {
			return false, errors.New(fmt.Sprintf("failed to compile schema %s, %v", schemaFile, err))
		}
	}
	return false, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {

	cfg, err := loadConfig(chunk)
	if err != nil {
		return chunks.Processed{}, err
	}

	snippets := make([]Snippet, 0)

	for _, spec := range cfg.Specs {

		codeToEmbed, err := yieldCodeExampleAndValidate(chunk, spec)
		if err != nil {
			return chunks.Processed{}, err
		}

		templateFile := getCodeTemplateFile(spec, chunk)

		template, err := util.TemplateWithFunctions(templateFile)
		if err != nil {
			return chunks.Processed{}, err
		}

		plugin := bootstrap.Registry.GetData("plugin")
		codeSnippet, err := util.RenderTemplate(template, Code{
			Plugin:     plugin.(config.Plugin),
			Properties: spec.Properties,
			Node:       codeToEmbed,
		})
		if err != nil {
			return chunks.Processed{}, err
		}

		snippets = append(snippets, Snippet{
			Display:  spec.Display,
			Language: From(spec.Language),
			Code:     string(codeSnippet),
		})
	}

	return chunks.Processed{Data: &Examples{
		Snippets: snippets,
	}}, nil
}

func loadConfig(chunk config.Chunk) (Config, error) {
	examplesFile := common.GetDataFile(chunk)
	bytes, err := os.ReadFile(examplesFile)
	if err != nil {
		return Config{}, err
	}
	cfg := &Config{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return Config{}, err
	}
	return *cfg, nil
}

func checkCodeTemplateFile(spec ExampleSpec, chunk config.Chunk) error {
	templateFile := getCodeTemplateFile(spec, chunk)
	if templateFile == "" {
		return errors.New(fmt.Sprintf("template with key '%s' does not exists for example: %v", spec.TemplateKey, chunk.Data))
	}
	if _, err := os.Stat(templateFile); err != nil {
		panic(fmt.Sprintf("can't stat file '%s' for template with key '%s': %v", templateFile, spec.TemplateKey, err))
	}
	return nil
}

func getCodeTemplateFile(spec ExampleSpec, chunk config.Chunk) string {
	var templateFile string
	if t, ok := chunk.Data[spec.TemplateKey]; ok {
		templateFile = t.(string)
	}
	return templateFile
}

func yieldCodeExampleAndValidate(chunk config.Chunk, spec ExampleSpec) (string, error) {
	var jsonToValidate string
	var codeToEmbed string
	var origin string

	_, validationSchema, _ := compileSchema(spec, chunk)
	if spec.UseSchemaDefaults {
		object := NewDocumentBuilder(spec)

		schema.Visit(validationSchema, object, &schema.VisitContext{})

		code, err := object.Marshall()
		jsonToValidate, err = object.MarshallWithLanguage(JSON)
		if err != nil {
			return "", err
		}
		codeToEmbed = code
		origin = "generated example from schema"
	} else {
		jsonBytes, err := os.ReadFile(spec.File)
		if err != nil {
			return "", errors.New(fmt.Sprintf("failed to read code example file %s: %v", spec.File, err))
		}
		jsonToValidate = string(jsonBytes)
		if spec.Language == YAML.String() {
			if yml, err := jsonToYaml(jsonToValidate); err != nil {
				panic(fmt.Sprintf("cannot json to yaml with example %v: %v", spec, err))
			} else {
				codeToEmbed = yml
			}
		} else {
			codeToEmbed = jsonToValidate
		}
		origin = spec.File
	}

	payload := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonToValidate), &payload)
	if err != nil {
		return "", err
	}

	err = validationSchema.Validate(payload)
	if err != nil {
		return "", errors.New(fmt.Sprintf("schema validation error: [%s] could not be validated:\n\t%s\n%s", origin, err, jsonToValidate))
	}
	return codeToEmbed, nil
}

func jsonToYaml(jsonToValidate string) (string, error) {
	y := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonToValidate), &y)
	if err != nil {
		return "", err
	}
	b, err := yaml.Marshal(y)
	return string(b), nil
}

func compileSchema(spec ExampleSpec, chunk config.Chunk) (string, *jsonschema.Schema, error) {
	schemaFile := common.GetFile(chunk, "schema")
	if spec.OverrideSchema != "" {
		schemaFile = spec.OverrideSchema
	}
	compiled, err := schema.CompileWithExtensions(schemaFile)
	return schemaFile, compiled, err
}
