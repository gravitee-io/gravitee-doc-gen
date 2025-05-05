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
	"errors"
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
	"os"
)

type Examples struct {
	current  int
	Snippets []Snippet
}

type Snippet struct {
	examples.Display
	Language examples.Language
	Code     string
}

type Code struct {
	plugin.Plugin
	Properties map[string]interface{}
	Node       string
}

type ExampleSpec interface {
	GetTemplateFile() string
	GetProperties() util.Unstructured
	TemplateFromRef() (examples.GenTemplate, bool)
	Validate() error
	GetOverrideSchema() string
	GetDisplay() examples.Display
	GetLanguage() examples.Language
}

type BaseExampleSpec struct {
	TemplateRef    string            `yaml:"templateRef"`
	Properties     util.Unstructured `yaml:"properties"`
	OverrideSchema string            `yaml:"overrideSchema"`
}

func (s BaseExampleSpec) Validate() error {
	if s.TemplateRef == "" {
		return errors.New(fmt.Sprintf("templateRef must be set for spec: %v", s))
	}
	_, ok := s.templateFromRef()
	if !ok {
		return errors.New(fmt.Sprintf("No template code example for ref '%s' in spec %v", s.TemplateRef, s))
	}

	if s.OverrideSchema != "" {
		if stat, err := os.Stat(s.OverrideSchema); err != nil || stat.IsDir() {
			return errors.New(fmt.Sprintf("overrideSchema cannot be loaded for spec %v: %s", s, err))
		}
	}
	return nil
}

func (s BaseExampleSpec) templateFromRef() (examples.GenTemplate, bool) {
	data := bootstrap.GetData("default-examples").(*examples.GenExamples)
	return data.FromRef(s.TemplateRef)
}

func (s BaseExampleSpec) getTemplateFile() string {
	ref, _ := s.templateFromRef()
	filename := ref.TemplateFilename()
	templateFile, err := plugin.RelativeFile(filename)
	if err != nil {
		panic(err.Error())
	}
	return templateFile
}
