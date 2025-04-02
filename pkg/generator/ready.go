package generator

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
)

func GetReady(configChunks []config.Chunk) ([]chunks.Ready, error) {

	result := make([]chunks.Ready, 0, len(configChunks))

	unique := make(map[string]bool)

	for i, chunk := range configChunks {

		validate, err := Registry.GetTypeValidator(chunk.Type)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("cannot load validator for %s [index: %d]: %s", chunk.Template, i, err.Error()))
		}
		exists, err := validate(chunk)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("cannot validate data of type '%s' for template %s [index: %d]: %s", chunk.Type, chunk.Template, i, err.Error()))
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
			unique[chunk.Id()] = true
		} else {

			tpl, err := util.CompileTemplateWithFunctions(chunk.Template)
			if err != nil {
				return nil, err
			}

			handle, err := Registry.GetTypeHandler(chunk.Type)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("cannot load type handler data %s [index: %d]: %s", chunk.Template, i, err.Error()))
			}

			if processed, err := handle(chunk); err == nil {
				result = append(result, chunks.Ready{
					Consumable: chunks.Consumable{
						Id:     chunk.Id(),
						Exists: exists,
					},
					CompiledTemplate: tpl,
					Processed:        processed,
				})
				unique[chunk.Id()] = true
			} else {
				return nil, err
			}
		}

	}

	if len(unique) != len(result) {
		return nil, errors.New("some chunks are using the same template filename, this is not allowed")
	}

	return result, nil

}
