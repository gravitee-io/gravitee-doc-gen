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
