package options

type Options struct {
	current  int
	Sections []Section
}

type Section struct {
	Title      string
	Comment    string
	Attributes []Attribute
}

type Attribute struct {
	Name        string
	Property    string
	Type        string
	Constraint  string
	Required    bool
	Default     string
	EL          bool
	Secret      bool
	Description string
	Enums       []string
}
