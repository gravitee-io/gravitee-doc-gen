package examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
)

type Yielder func(config.Chunk, ExampleSpec) (string, error)

func ProcessAllExamples(chunk config.Chunk, provider ExampleSpecProvider, yield Yielder) (chunks.Processed, error) {

	err := LoadConfig(chunk, provider)
	if err != nil {
		return chunks.Processed{}, err
	}

	snippets := make([]Snippet, 0)

	for _, spec := range provider.ExampleSpecs() {

		codeToEmbed, err := yield(chunk, spec)
		if err != nil {
			return chunks.Processed{}, err
		}

		templateFile := spec.GetTemplateFile()

		template, err := util.TemplateWithFunctions(templateFile)
		if err != nil {
			return chunks.Processed{}, err
		}

		plugin := bootstrap.GetData("plugin")
		codeSnippet, err := util.RenderTemplate(template, Code{
			Plugin:     plugin.(pkg.Plugin),
			Properties: spec.GetProperties(),
			Node:       codeToEmbed,
		})
		if err != nil {
			return chunks.Processed{}, err
		}

		snippets = append(snippets, Snippet{
			Display:  spec.GetDisplay(),
			Language: spec.GetLanguage(),
			Code:     string(codeSnippet),
		})
	}

	return chunks.Processed{Data: &Examples{
		Snippets: snippets,
	}}, nil
}
