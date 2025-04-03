package table

import (
	"errors"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"gopkg.in/yaml.v3"
	"os"
)

type Columns struct {
	Id    string `yaml:"id"`
	Label string `yaml:"label"`
}

type Row struct {
	Deprecated bool           `yaml:"deprecated"`
	Data       map[string]any `yaml:"data"`
}

type Table struct {
	Columns []Columns
	Rows    []Row `yaml:"rows"`
}

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := common.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}
	tableFile := common.GetDataFile(chunk)
	tableFileExists := common.FileExists(tableFile)

	if chunk.Required && !tableFileExists {
		return false, errors.New("table file not found")
	}

	return tmplExists && tableFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	matrix, err := os.ReadFile(common.GetDataFile(chunk))
	if err != nil {
		return chunks.Processed{}, err
	}

	table := Table{}
	processed := chunks.Processed{Data: &table}
	err = yaml.Unmarshal(matrix, processed.Data)
	if err != nil {
		return chunks.Processed{}, err
	}
	table.Columns = getColumns(chunk)

	return processed, nil
}

func getColumns(chunk config.Chunk) []Columns {
	if cols, ok := chunk.Data["columns"]; ok && cols != nil {
		if colsMaps, ok := cols.(map[string]interface{}); ok && colsMaps != nil {
			columns := make([]Columns, 0)
			for id, label := range colsMaps {
				columns = append(columns, Columns{Id: id, Label: label.(string)})
			}
			return columns
		}
	}
	panic("no columns definition for chunk: " + chunk.String())
}
