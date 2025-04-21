package generator

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/core/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/util"
)

func Generate(readyChunks []chunks.Ready) ([]chunks.Generated, error) {

	generated := make([]chunks.Generated, 0, len(readyChunks))

	for _, chunk := range readyChunks {
		if !chunk.Exists {
			generated = append(generated, chunks.Generated{
				Consumable: chunk.Consumable,
				Content:    "",
			})
			continue
		}
		if rendered, err := util.RenderTemplate(chunk.CompiledTemplate, chunk.Data); err == nil {
			generated = append(generated, chunks.Generated{
				Consumable: chunk.Consumable,
				Content:    string(rendered),
			})
		} else {
			return nil, err
		}
	}

	return generated, nil
}
