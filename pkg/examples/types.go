package examples

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/core"
	"os"
)

type Display struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Filename    string `yaml:"filename"`
}

type Examples struct {
	current  int
	Snippets []Snippet
}

type Snippet struct {
	Display
	Language Language
	Code     string
}

type Code struct {
	pkg.Plugin
	Properties map[string]interface{}
	Node       string
}

type ConfigData struct {
	GenExamples []GenExampleSpec `yaml:"genExamples"`
	RawExamples []RawExampleSpec `yaml:"rawExamples"`
}

type ExampleSpec interface {
	GetTemplateFile() string
	GetProperties() core.Unstructured
	TemplateFromRef() (GenTemplate, bool)
	Validate() error
	GetOverrideSchema() string
	GetDisplay() Display
	GetLanguage() Language
}

type BaseExampleSpec struct {
	TemplateRef    string            `yaml:"templateRef"`
	Properties     core.Unstructured `yaml:"properties"`
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

func (s BaseExampleSpec) templateFromRef() (GenTemplate, bool) {
	data := bootstrap.GetData("gen-examples").(*GenExamples)
	return data.FromRef(s.TemplateRef)
}

func (s BaseExampleSpec) getTemplateFile() string {
	ref, _ := s.templateFromRef()
	filename := ref.TemplateFilename()
	templateFile, err := pkg.PluginRelatedFile(filename)
	if err != nil {
		panic(err.Error())
	}
	return templateFile
}
