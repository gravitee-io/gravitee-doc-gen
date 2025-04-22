package table

import (
	"errors"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
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
	tmplExists, err := chunks.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}
	tableFile := chunks.GetDataTypeFile(chunk)
	tableFileExists := util.FileExists(tableFile)

	if chunk.Required && !tableFileExists {
		return false, errors.New("table file not found")
	}

	return tmplExists && tableFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	rows, err := os.ReadFile(chunks.GetDataTypeFile(chunk))
	if err != nil {
		return chunks.Processed{}, err
	}

	table := Table{}
	processed := chunks.Processed{Data: &table}
	err = yaml.Unmarshal(rows, processed.Data)
	if err != nil {
		return chunks.Processed{}, err
	}
	table.Columns = getColumns(chunk)

	return processed, nil
}

func getColumns(chunk config.Chunk) []Columns {
	if cols, ok := chunk.Data["columns"]; ok && cols != nil {
		if colsMaps, ok := cols.([]interface{}); ok && colsMaps != nil {
			return getColumnsFromMap(chunk, colsMaps)
		}
		panic("columns should be a list: " + chunk.String())
	}
	panic("no columns definition for chunk: " + chunk.String())
}

func getColumnsFromMap(chunk config.Chunk, colsMaps []interface{}) []Columns {
	columns := make([]Columns, 0)
	for _, colsMap := range colsMaps {
		if colMap, ok := colsMap.(map[string]interface{}); ok && len(colMap) == 1 {
			for id, label := range colMap {
				columns = append(columns, Columns{Id: id, Label: label.(string)})
			}
		} else {
			panic("one columns spec should a single key/value: " + chunk.String())
		}
	}
	return columns
}
