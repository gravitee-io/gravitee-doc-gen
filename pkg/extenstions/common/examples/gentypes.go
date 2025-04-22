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
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
)

type GenExampleSpec struct {
	BaseExampleSpec `yaml:",inline"`
	OneOfFilter     visitor.OneOfFilter `yaml:"oneOfFilter"`
}

func (s GenExampleSpec) GetTemplateFile() string {
	return s.getTemplateFile()
}

func (s GenExampleSpec) Validate() error {
	return s.BaseExampleSpec.Validate()
}

func (s GenExampleSpec) GetProperties() util.Unstructured {
	return s.Properties
}

func (s GenExampleSpec) GetOverrideSchema() string {
	return s.BaseExampleSpec.OverrideSchema
}

func (s GenExampleSpec) TemplateFromRef() (examples.GenTemplate, bool) {
	return s.BaseExampleSpec.templateFromRef()
}

func (s GenExampleSpec) GetDisplay() examples.Display {
	if ref, ok := s.TemplateFromRef(); ok {
		return ref.Display
	}
	return examples.Display{}
}

func (s GenExampleSpec) GetLanguage() examples.Language {
	if ref, ok := s.TemplateFromRef(); ok {
		return ref.Language
	}
	return examples.Unknown
}

type GenExampleProvider struct {
	ConfigData
}

func (p *GenExampleProvider) SetConfigData(d ConfigData) {
	p.ConfigData = d
}

func (p *GenExampleProvider) GetConfigData() ConfigData {
	return p.ConfigData
}

func (p *GenExampleProvider) ExampleSpecs() []ExampleSpec {
	ex := make([]ExampleSpec, 0)
	for _, e := range p.ConfigData.GenExamples {
		ex = append(ex, e)
	}
	return ex
}
