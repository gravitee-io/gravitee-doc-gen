package config

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"
	"path/filepath"
	"strings"
)

const UnknownDataType = DataType("")
const TableDataType = DataType("table")
const CodeDataType = DataType("code")
const Options = DataType("options")

type Config struct {
	MainTemplate string  `yaml:"mainTemplate"`
	Chunks       []Chunk `yaml:"chunks"`
}

type Chunk struct {
	Template string         `yaml:"template"`
	Type     DataType       `yaml:"type"`
	Data     map[string]any `yaml:"data"`
	Required bool           `yaml:"required"`
}

func (c Chunk) Id() string {
	id := strings.TrimSuffix(filepath.Base(c.Template), filepath.Ext(c.Template))
	return util.Title(id)
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

type Plugin struct {
	Id    string
	Type  string
	Title string
}
