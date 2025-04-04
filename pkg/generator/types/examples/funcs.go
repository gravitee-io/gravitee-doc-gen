package examples

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
	ext "github.com/gravitee-io-labs/readme-gen/pkg/schema/extensions"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"github.com/santhosh-tekuri/jsonschema/v5"
	"os"
)

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := common.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}

	schemaFile := common.GetFile(chunk, "schema")
	schemaFileExists := common.FileExists(schemaFile)
	if chunk.Required && !schemaFileExists {
		return false, errors.New("schema file not found")
	}

	_, err = schema.CompileWithExtensions(schemaFile)
	if err != nil {
		return false, err
	}

	return tmplExists && schemaFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {

	schemaFile := common.GetFile(chunk, "schema")

	root, err := schema.CompileWithExtensions(schemaFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	snippets := make([]Snippet, 0)
	readmeExamples := getExampleExtension(root)
	for i, readmeExample := range readmeExamples.Examples {

		templateFile := getTemplateFile(readmeExample.Template, chunk)
		if templateFile == "" {
			return chunks.Processed{}, errors.New(fmt.Sprintf("template '%s' not found for %s and title %s at index %d", readmeExample.Template, ext.ReadmeExamplesExtension, readmeExample.Title, i))
		}
		if _, err := os.Stat(templateFile); os.IsNotExist(err) {
			return chunks.Processed{}, errors.New(fmt.Sprintf("template '%s' not found", readmeExample.Template))
		}
		snippet := Snippet{
			Title:       readmeExample.Title,
			Description: readmeExample.Description,
			Filename:    readmeExample.Filename,
			Language:    From(readmeExample.Language),
		}
		object := NewDocumentBuilder(readmeExample, i)
		schema.Visit(root, object, &schema.VisitContext{QueueNodes: false})

		code, err := object.Marshall()
		if err != nil {
			return chunks.Processed{}, err
		}

		// TODO check schema is validated

		template, err := util.CompileTemplateWithFunctions(templateFile)
		if err != nil {
			return chunks.Processed{}, err
		}

		plugin := bootstrap.Registry.GetData("plugin")
		codeSnippet, err := util.RenderTemplate(template, Code{
			Plugin:        plugin.(config.Plugin),
			Properties:    readmeExample.Properties,
			Configuration: code,
		})
		if err != nil {
			return chunks.Processed{}, err
		}

		snippet.Code = string(codeSnippet)

		snippets = append(snippets, snippet)
	}

	return chunks.Processed{Data: &Examples{
		Snippets: snippets,
	}}, nil
}

func getExampleExtension(root *jsonschema.Schema) ext.ExamplesSettings {
	if raw, ok := root.Extensions[ext.ReadmeExamplesExtension]; ok {
		return raw.(ext.ExamplesSettings)
	}
	return ext.ExamplesSettings{Examples: make([]ext.ReadmeExample, 0)}
}

func getTemplateFile(name string, chunk config.Chunk) string {
	if t, ok := chunk.Data[name]; ok {
		return t.(string)
	}
	return ""
}
