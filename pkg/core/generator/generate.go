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
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
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
