package examples

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
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

		pl := bootstrap.GetData("plugin")
		codeSnippet, err := util.RenderTemplate(template, Code{
			Plugin:     pl.(plugin.Plugin),
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
