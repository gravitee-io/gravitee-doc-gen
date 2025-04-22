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

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
)

const UnknownDataType = DataType("")

type Config struct {
	Outputs []Output
	Chunks  []Chunk `yaml:"chunks"`
}

type Output struct {
	Template        string `yaml:"template"`
	Target          string `yaml:"target"`
	ProcessExisting bool   `yaml:"processExisting"`
}

type Chunk struct {
	ExportedAs string         `yaml:"exportedAs"`
	Template   string         `yaml:"template"`
	Type       DataType       `yaml:"type"`
	Data       map[string]any `yaml:"data"`
	Required   bool           `yaml:"required"`
}

func (c Chunk) ID() string {
	if c.ExportedAs != "" {
		return c.ExportedAs
	}
	return util.Title(util.BaseFileNoExt(c.Template))
}

func (c Chunk) String() string {
	return fmt.Sprintf("template:%s, type:%s", c.Template, c.Type)
}

type DataType string
