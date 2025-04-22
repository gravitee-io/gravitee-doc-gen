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
