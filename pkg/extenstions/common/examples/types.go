package examples

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/core/util"
	"github.com/gravitee-io-labs/readme-gen/pkg/extenstions/bootstrap/examples"
	"github.com/gravitee-io-labs/readme-gen/pkg/extenstions/bootstrap/plugin"
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
	data := bootstrap.GetData("gen-examples").(*examples.GenExamples)
	return data.FromRef(s.TemplateRef)
}

func (s BaseExampleSpec) getTemplateFile() string {
	ref, _ := s.templateFromRef()
	filename := ref.TemplateFilename()
	templateFile, err := plugin.PluginRelatedFile(filename)
	if err != nil {
		panic(err.Error())
	}
	return templateFile
}
