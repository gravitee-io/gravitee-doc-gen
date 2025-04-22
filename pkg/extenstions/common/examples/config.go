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

package examples

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"gopkg.in/yaml.v3"
	"os"
)

type ExampleSpecProvider interface {
	ExampleSpecs() []ExampleSpec
	SetConfigData(ConfigData)
	GetConfigData() ConfigData
}

type ConfigData struct {
	GenExamples []GenExampleSpec `yaml:"genExamples"`
	RawExamples []RawExampleSpec `yaml:"rawExamples"`
}

func LoadConfig(chunk config.Chunk, provider ExampleSpecProvider) error {
	examplesFile := chunks.GetString(chunk, "examples")
	bytes, err := os.ReadFile(examplesFile)
	if err != nil {
		return err
	}
	cfg := &ConfigData{}
	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return err
	}

	provider.SetConfigData(*cfg)
	return nil
}
