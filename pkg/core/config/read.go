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

package config

import (
	"fmt"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"gopkg.in/yaml.v3"
)

type FileResolver func(string) (string, error)

func Read(rootDir string, resolver FileResolver) (Config, error) {
	configFile, err := resolver(rootDir)
	if err != nil {
		return Config{}, err
	}
	fmt.Println("Generation config: ", configFile)

	rendered, err := util.RenderTemplateFromFile(configFile, bootstrap.GetExported())
	if err != nil {
		return Config{}, err
	}

	// read the config
	config := Config{}

	err = yaml.Unmarshal(rendered, &config)
	return config, err
}
