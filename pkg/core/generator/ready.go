package generator

import (
	"errors"
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
)

func GetReady(configChunks []config.Chunk) ([]chunks.Ready, error) {

	result := make([]chunks.Ready, 0, len(configChunks))

	unique := util.Set{}

	for i, chunk := range configChunks {

		validate, err := Registry.getTypeValidator(chunk.Type)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("cannot load validator for %s [index: %d]: %s", chunk.Template, i, err.Error()))
		}
		exists, err := validate(chunk)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("cannot validate chunk of type '%s' for template %s [index: %d]:\n %s", chunk.Type, chunk.Template, i, err.Error()))
		}

		if !exists {
			result = append(result, chunks.Ready{
				Consumable: chunks.Consumable{
					Id:     chunk.Id(),
					Exists: exists,
				},
				CompiledTemplate: nil,
				Processed:        chunks.Processed{},
			})
			unique.Add(chunk.Id())
		} else {
			ready, err := generateChunk(chunk, i, unique)
			if err != nil {
				return nil, err
			}
			result = append(result, ready)
		}

	}

	if len(unique) != len(result) {
		return nil, errors.New("some chunks are using the same template filename, for those set 'exportedAs' with a name to use in the template")
	}

	return result, nil

}

func generateChunk(chunk config.Chunk, index int, unique util.Set) (chunks.Ready, error) {
	tpl, err := util.TemplateWithFunctions(chunk.Template)
	if err != nil {
		return chunks.Ready{}, err
	}

	handle, err := Registry.getTypeHandler(chunk.Type)
	if err != nil {
		return chunks.Ready{}, errors.New(fmt.Sprintf("cannot load type handler data %s [index: %d]: %s", chunk.Template, index, err.Error()))
	}

	var done chunks.Processed
	if processed, err := handle(chunk); err != nil {
		return chunks.Ready{}, err
	} else {
		done = processed
	}

	ready := chunks.Ready{
		Consumable: chunks.Consumable{
			Id:     chunk.Id(),
			Exists: true,
		},
		CompiledTemplate: tpl,
		Processed:        done,
	}
	unique.Add(chunk.Id())
	return ready, nil
}
