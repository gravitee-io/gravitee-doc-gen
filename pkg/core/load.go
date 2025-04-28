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

package core

import (
	"fmt"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/generator"
)

func Load(rootDir string, overrideConfigFile string) ([]chunks.Generated, config.Config, error) {
	err := bootstrap.Load(rootDir)
	if err != nil {
		return nil, config.Config{},
			fmt.Errorf("failed to load bootstrap.yaml: %s", err.Error())
	}

	if overrideConfigFile != "" {
		bootstrap.OverrideData(bootstrap.ConfigResolver, "default")
	}

	cfg, err := config.Read(rootDir, overrideConfigFile)
	if err != nil {
		return nil, cfg, err
	}

	ready, err := generator.GetReady(cfg.Chunks)
	if err != nil {
		return nil, config.Config{}, err
	}

	generated, err := generator.Generate(ready)
	if err != nil {
		return nil, config.Config{}, err
	}

	return generated, cfg, nil
}
