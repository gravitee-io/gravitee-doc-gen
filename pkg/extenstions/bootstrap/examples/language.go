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

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v3"
	"strings"
)

const Unknown Language = iota
const JSON Language = 1
const YAML Language = 2

type Language int

func From(language string) Language {

	switch strings.ToLower(language) {
	case "json":
		return JSON
	case "yaml":
		return YAML
	default:
		return Unknown
	}
}

func (l *Language) String() string {
	switch *l {
	case JSON:
		return "json"
	case YAML:
		return "yaml"
	case Unknown:
		return ""
	}
	panic("unknown language")
}

func (l *Language) Serialize(object any) (string, error) {
	switch *l {
	case JSON:
		buffer := bytes.Buffer{}
		e := json.NewEncoder(&buffer)
		e.SetEscapeHTML(false)
		e.SetIndent("", "  ")
		err := e.Encode(object)
		return buffer.String(), err
	case YAML:
		asJson, err := json.Marshal(object)
		if err != nil {
			return "", err
		}
		unstructured := make(map[string]interface{})
		err = json.Unmarshal(asJson, &unstructured)
		if err != nil {
			return "", err
		}
		buffer := bytes.Buffer{}
		encoder := yaml.NewEncoder(&buffer)
		encoder.SetIndent(2)
		err = encoder.Encode(unstructured)
		return buffer.String(), err
	}
	panic("unknown language")
}
func (l *Language) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	*l = From(str)
	return nil
}
func (l *Language) UnmarshalYAML(value *yaml.Node) error {
	var str string
	if err := value.Decode(&str); err != nil {
		return err
	}
	*l = From(str)
	return nil
}
