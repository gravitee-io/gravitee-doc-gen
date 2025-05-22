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

	unique := util.NewSet()

	for i, chunk := range configChunks {
		validate, err := Registry.getTypeValidator(chunk.Type)
		if err != nil {
			return nil, fmt.Errorf("cannot load validator for %s [index: %d]: %s", chunk.Template, i, err.Error())
		}
		exists, err := validate(chunk)
		if err != nil {
			return nil,
				fmt.Errorf("cannot validate chunk of type '%s' for template %s [index: %d]:\n %s",
					chunk.Type,
					chunk.Template,
					i,
					err.Error())
		}

		if !exists {
			result = append(result, chunks.Ready{
				Consumable: chunks.Consumable{
					ID:     chunk.ID(),
					Exists: exists,
				},
				CompiledTemplate: nil,
				Processed:        chunks.Processed{},
			})
			unique.Add(chunk.ID())
		} else {
			ready, err := generateChunk(chunk, i, unique)
			if err != nil {
				return nil, err
			}
			result = append(result, ready)
		}
	}

	if len(unique.Items()) != len(result) {
		return nil,
			errors.New("some chunks are using the same template filename, " +
				"for those set 'exportedAs' with a name to use in the template")
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
		return chunks.Ready{},
			fmt.Errorf("cannot load type handler data %s [index: %d]: %s", chunk.Template, index, err.Error())
	}

	var done chunks.Processed
	if processed, err := handle(chunk); err != nil {
		return chunks.Ready{}, err
	} else {
		done = processed
	}

	ready := chunks.Ready{
		Consumable: chunks.Consumable{
			ID:     chunk.ID(),
			Exists: true,
		},
		CompiledTemplate: tpl,
		Processed:        done,
	}
	unique.Add(chunk.ID())
	return ready, nil
}
