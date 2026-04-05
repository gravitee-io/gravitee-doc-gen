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

package modules

import (
	"fmt"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
)

type Module struct {
	ID         string `json:"id"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	ExportedAs string `json:"exportedAs"`
}

type modulesFile struct {
	Modules []Module `json:"modules"`
}

func PostProcessor(data any) (any, error) {
	raw := util.As[*util.Unstructured](data)
	if raw == nil {
		return nil, fmt.Errorf("modules data is nil")
	}

	mf, err := util.AnyMapToStruct[modulesFile](raw)
	if err != nil {
		return nil, fmt.Errorf("cannot parse modules data: %w", err)
	}

	if len(mf.Modules) == 0 {
		return nil, fmt.Errorf("at least one module must be declared")
	}

	for i, m := range mf.Modules {
		if m.Path == "" {
			return nil, fmt.Errorf("module at index %d: path is required", i)
		}
		if m.ExportedAs == "" {
			return nil, fmt.Errorf("module at index %d: exportedAs is required", i)
		}
	}

	return mf.Modules, nil
}
