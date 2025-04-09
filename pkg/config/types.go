package config

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
)

const UnknownDataType = DataType("")

type Output struct {
	Template        string `yaml:"template"`
	Target          string `yaml:"target"`
	ProcessExisting bool   `yaml:"processExisting"`
}

type Config struct {
	Outputs []Output
	Chunks  []Chunk `yaml:"chunks"`
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

type RawData struct {
	Key     string `yaml:"key"`
	Content string `yaml:"content"`
}

type Schema struct {
	Main   string `yaml:"main"`
	Shared string `yaml:"shared"`
}
