package examples

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"os"
)

type Examples struct {
	current  int
	Snippets []Snippet
}

type Display struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Filename    string `yaml:"filename"`
}

type Snippet struct {
	Display
	Language Language
	Code     string
}

type Code struct {
	config.Plugin
	Properties map[string]interface{}
	Node       string
}

type Config struct {
	Specs []ExampleSpec `yaml:"specs"`
}

type ExampleSpec struct {
	Display           `yaml:",inline"`
	Language          string         `yaml:"language"`
	TemplateKey       string         `yaml:"templateKey"`
	UseSchemaDefaults bool           `yaml:"useSchemaDefaults"`
	OneOfFilter       OneOfFilter    `yaml:"oneOfFilter"` // to keep ?
	Properties        map[string]any `yaml:"properties"`
	File              string         `yaml:"file"`
	OverrideSchema    string         `yaml:"overrideSchema"`
}

func (s ExampleSpec) Check() error {
	if s.Language == "" {
		return errors.New(fmt.Sprintf("language must be set for spec: %v", s))
	}
	if s.TemplateKey == "" {
		return errors.New(fmt.Sprintf("templateKey must be set for spec: %v", s))
	}
	if !s.UseSchemaDefaults && s.File == "" {
		return errors.New(fmt.Sprintf("file must be set for spec: %v", s))
	}
	if s.File != "" {
		if stat, err := os.Stat(s.File); err != nil || stat.IsDir() {
			return errors.New(fmt.Sprintf("file cannot be loaded for spec %v: %s", s, err))
		}
	}
	return nil
}

type OneOfFilter struct {
	Path           []string       `json:"path"`
	Discriminators map[string]any `json:"discriminators"`
	Index          int            `json:"index"`
}
