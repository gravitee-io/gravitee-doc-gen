package examples

type Display struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Filename    string `yaml:"filename"`
}

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
