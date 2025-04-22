package examples

import (
	"errors"
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/examples"
	"os"
)

type RawExampleSpec struct {
	examples.Display `yaml:",inline"`
	BaseExampleSpec  `yaml:",inline"`
	Language         examples.Language `yaml:"language"`
	File             string            `yaml:"file"`
}

func (s RawExampleSpec) GetTemplateFile() string {
	return s.getTemplateFile()
}

func (s RawExampleSpec) TemplateFromRef() (examples.GenTemplate, bool) {
	return s.BaseExampleSpec.templateFromRef()
}

func (s RawExampleSpec) GetOverrideSchema() string {
	return s.BaseExampleSpec.OverrideSchema
}

func (s RawExampleSpec) Validate() error {
	err := s.BaseExampleSpec.Validate()
	if err != nil {
		return err
	}
	if s.Language == examples.Unknown {
		return errors.New(fmt.Sprintf("unknown or unset language for spec: %v", s))
	}
	if s.File == "" {
		return errors.New(fmt.Sprintf("file must be set for spec: %v", s))
	}
	if stat, err := os.Stat(s.File); err != nil || stat.IsDir() {
		return errors.New(fmt.Sprintf("file cannot be loaded for spec %v: %s", s, err))
	}

	return nil
}

func (s RawExampleSpec) GetProperties() util.Unstructured {
	return s.Properties
}

func (s RawExampleSpec) GetDisplay() examples.Display {
	return s.Display
}

func (s RawExampleSpec) GetLanguage() examples.Language {
	return s.Language
}

type RawExampleProvider struct {
	ConfigData
}

func (p *RawExampleProvider) ExampleSpecs() []ExampleSpec {
	ex := make([]ExampleSpec, 0)
	for _, e := range p.ConfigData.RawExamples {
		ex = append(ex, e)
	}
	return ex
}

func (p *RawExampleProvider) GetConfigData() ConfigData {
	return p.ConfigData
}

func (p *RawExampleProvider) SetConfigData(d ConfigData) {
	p.ConfigData = d
}
