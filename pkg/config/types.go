package config

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
)

const UnknownDataType = DataType("")
const TableDataType = DataType("table")
const CodeDataType = DataType("code")
const Options = DataType("options")
const Examples = DataType("examples")

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
	Template string         `yaml:"template"`
	Type     DataType       `yaml:"type"`
	Data     map[string]any `yaml:"data"`
	Required bool           `yaml:"required"`
}

func (c Chunk) Id() string {
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

// TODO move this to base package
type Plugin struct {
	Id    string
	Type  string
	Title string
}
