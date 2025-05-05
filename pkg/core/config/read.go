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
	"path"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"gopkg.in/yaml.v3"
)

type ConfigResolver func(string, string) (string, error)

func DefaultConfigResolver(rootDir string, configFile string) (string, error) {
	file := "default.yaml"
	if configFile != "" {
		file = configFile
	}
	return path.Join(rootDir, file), nil
}

var resolvers = map[string]ConfigResolver{
	"default": DefaultConfigResolver,
}

func RegisterConfigResolver(name string, resolver ConfigResolver) {
	if name == "default" {
		panic("default resolver name 'default' is reserved")
	}
	resolvers[name] = resolver
}

func Read(rootDir string, configFile string) (Config, error) {
	s, _ := bootstrap.GetData(bootstrap.ConfigResolver).(string)
	resolver := resolvers[s]
	if resolver == nil {
		panic("unknown chunk config resolver: " + s)
	}

	configFile, err := resolver(rootDir, configFile)
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
