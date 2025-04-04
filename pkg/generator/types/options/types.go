package options

import "github.com/gravitee-io-labs/readme-gen/pkg/schema"

type Options struct {
	current  int
	Sections []Section
}

type Section struct {
	Title           string
	Type            string
	OneOf           schema.OneOf
	DiscriminatedBy string
	Attributes      []Attribute
}

type Attribute struct {
	Name        string
	Property    string
	Type        string
	Constraint  string
	Required    bool
	Default     string
	IsConstant  bool
	EL          bool
	Secret      bool
	Description string
	Enums       []string
	OneOf       schema.OneOf
}
