// Copyright (C) 2015 The Gravitee team (http://gravitee.io)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package options

import (
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/common/visitor"
)

type Options struct {
	current  int
	Sections []Section
}

type Section struct {
	Title           string
	Type            string
	OneOf           visitor.OneOfDescriptor
	DiscriminatedBy map[string]any
	Attributes      []Attribute
}

func (s *Section) ELPresent() bool {
	for _, attribute := range s.Attributes {
		if attribute.EL {
			return true
		}
	}
	return false
}

func (s *Section) SecretPresent() bool {
	for _, attribute := range s.Attributes {
		if attribute.Secret {
			return true
		}
	}
	return false
}

func (s *Section) DefaultPresent() bool {
	for _, attribute := range s.Attributes {
		if attribute.Default != nil {
			return true
		}
	}
	return false
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
	OneOf       visitor.OneOfDescriptor
}
