package examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/core"
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
)

type GenExampleSpec struct {
	BaseExampleSpec `yaml:",inline"`
	OneOfFilter     schema.OneOfFilter `yaml:"oneOfFilter"`
}

func (s GenExampleSpec) GetTemplateFile() string {
	return s.getTemplateFile()
}

func (s GenExampleSpec) Validate() error {
	return s.BaseExampleSpec.Validate()
}

func (s GenExampleSpec) GetProperties() core.Unstructured {
	return s.Properties
}

func (s GenExampleSpec) GetOverrideSchema() string {
	return s.BaseExampleSpec.OverrideSchema
}

func (s GenExampleSpec) TemplateFromRef() (GenTemplate, bool) {
	return s.BaseExampleSpec.templateFromRef()
}

func (s GenExampleSpec) GetDisplay() Display {
	if ref, ok := s.TemplateFromRef(); ok {
		return ref.Display
	}
	return Display{}
}

func (s GenExampleSpec) GetLanguage() Language {
	if ref, ok := s.TemplateFromRef(); ok {
		return ref.Language
	}
	return Unknown
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
