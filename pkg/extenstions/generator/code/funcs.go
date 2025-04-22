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

package code

import (
	"errors"
	"os"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"gopkg.in/yaml.v3"
)

type Snippet struct {
	Title    string `yaml:"title"`
	Language string `yaml:"language"`
	Content  string `yaml:"content"`
}

type Code struct {
	Before   string    `yaml:"before"`
	After    string    `yaml:"after"`
	Snippets []Snippet `yaml:"snippets"`
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	codeFile := chunks.GetDataTypeFile(chunk)
	content, err := os.ReadFile(codeFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	code := Code{}
	err = yaml.Unmarshal(content, &code)
	if err != nil {
		return chunks.Processed{}, err
	}
	if util.FileExists(code.Before) {
		beforeContent, err := os.ReadFile(code.Before)
		if err != nil {
			return chunks.Processed{}, err
		}
		code.Before = string(beforeContent)
	} else {
		code.Before = ""
	}
	if util.FileExists(code.After) {
		afterContent, err := os.ReadFile(code.After)
		if err != nil {
			return chunks.Processed{}, err
		}
		code.After = string(afterContent)
	} else {
		code.After = ""
	}

	return chunks.Processed{Data: code}, nil
}

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := chunks.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}
	codeFile := chunks.GetDataTypeFile(chunk)
	codeFileExists := util.FileExists(codeFile)

	if chunk.Required && !codeFileExists {
		return false, errors.New("code file not found")
	}

	return tmplExists && codeFileExists, nil
}
