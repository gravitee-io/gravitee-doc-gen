package options

import (
	"github.com/gravitee-io-labs/readme-gen/pkg/schema"
)

type Options struct {
	current  int
	Sections []Section
}

type Section struct {
	Title           string
	Type            string
	OneOf           schema.OneOf
	DiscriminatedBy map[string]any
	Attributes      []Attribute
}

func (s *Section) IsOneOfProperty(property string) bool {
	discriminated := s.DiscriminatedBy != nil
	if s.OneOf.Present && discriminated {
		_, hasProperty := s.DiscriminatedBy[property]
		return hasProperty
	}
	return false
}

type Attribute struct {
	Name        string
	Property    string
	Type        string
	TypeItem    string
	Constraint  string
	Required    bool
	Default     any
	IsConstant  bool
	EL          bool
	Secret      bool
	Description string
	Enums       []any
	OneOf       schema.OneOf
}
