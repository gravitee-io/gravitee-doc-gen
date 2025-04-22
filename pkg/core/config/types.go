package config

import (
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
)

const UnknownDataType = DataType("")

type Config struct {
	Outputs []Output
	Chunks  []Chunk `yaml:"chunks"`
}

type Output struct {
	Template        string `yaml:"template"`
	Target          string `yaml:"target"`
	ProcessExisting bool   `yaml:"processExisting"`
}

type Chunk struct {
	ExportedAs string         `yaml:"exportedAs"`
	Template   string         `yaml:"template"`
	Type       DataType       `yaml:"type"`
	Data       map[string]any `yaml:"data"`
	Required   bool           `yaml:"required"`
}

func (c Chunk) Id() string {
	if c.ExportedAs != "" {
		return c.ExportedAs
	}
	return util.Title(util.BaseFileNoExt(c.Template))
}

func (c Chunk) String() string {
	return fmt.Sprintf("template:%s, type:%s", c.Template, c.Type)
}

type DataType string
