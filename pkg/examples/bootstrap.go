package examples

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/core"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
)

type GenExamples struct {
	Templates []GenTemplate `yaml:"templates"`
}

func (e GenExamples) FromRef(id string) (GenTemplate, bool) {
	for _, t := range e.Templates {
		if t.Id == id {
			return t, true
		}
	}
	return GenTemplate{}, false
}

type GenTemplate struct {
	Id       string   `yaml:"id"`
	Language Language `yaml:"language"`
	Display  `yaml:",inline"`
}

func (t GenTemplate) TemplateFilename() string {
	return t.Id + "." + t.Language.String() + ".tmpl"
}

func GenExamplePostProcessor(data any) (any, error) {
	object := data.(*core.Unstructured)
	return util.AnyMapToStruct[GenExamples](object)
}
