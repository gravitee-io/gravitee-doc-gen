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

func (l Language) String() string {
	switch l {
	case JSON:
		return "json"
	case YAML:
		return "yaml"
	case Unknown:
		return ""
	}
	panic("unknown language")
}

func (l Language) Serialize(object any) (string, error) {
	switch l {
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
