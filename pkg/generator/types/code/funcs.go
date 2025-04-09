package code

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"gopkg.in/yaml.v3"
	"os"
)

type Snippet struct {
	Title    string `yaml:"title"`
	Language string `yaml:"language"`
	Content  string `yaml:"content"`
}

type Code struct {
	Before   string    `yaml:"before"`
	After    string    `yaml:"after"`
	Snippets []Snippet `yaml:"snippets"`
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	codeFile := common.GetDataTypeFile(chunk)
	content, err := os.ReadFile(codeFile)
	if err != nil {
		return chunks.Processed{}, err
	}

	code := Code{}
	err = yaml.Unmarshal(content, &code)
	if err != nil {
		return chunks.Processed{}, err
	}
	if common.FileExists(code.Before) {
		beforeContent, err := os.ReadFile(code.Before)
		if err != nil {
			return chunks.Processed{}, err
		}
		code.Before = string(beforeContent)
	} else {
		code.Before = ""
	}
	if common.FileExists(code.After) {
		afterContent, err := os.ReadFile(code.After)
		if err != nil {
			return chunks.Processed{}, err
		}
		code.After = string(afterContent)
	} else {
		code.After = ""
	}

	return chunks.Processed{Data: code}, nil
}

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := common.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}
	codeFile := common.GetDataTypeFile(chunk)
	codeFileExists := common.FileExists(codeFile)

	if chunk.Required && !codeFileExists {
		return false, errors.New("code file not found")
	}

	return tmplExists && codeFileExists, nil
}
