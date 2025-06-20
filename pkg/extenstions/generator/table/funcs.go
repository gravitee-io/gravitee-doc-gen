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

package table

import (
	"fmt"
	"os"
	"slices"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"gopkg.in/yaml.v3"
)

type Columns struct {
	ID    string `yaml:"id"`
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

func (t *Table) removeUnusedColumns() {
	ids := util.NewSet()
	for _, r := range t.Rows {
		for k := range r.Data {
			ids.Add(k)
		}
	}
	cols := make([]Columns, 0)
	usedCols := util.ToSlice[string](ids)
	for _, col := range t.Columns {
		if slices.Contains(usedCols, col.ID) {
			cols = append(cols, col)
		}
	}
	t.Columns = cols
}

func TypeValidator(chunk config.Chunk) (bool, error) {
	tmplExists, err := chunks.TemplateExistsTypeValidator(chunk)
	if err != nil || chunk.Required && !tmplExists {
		return false, err
	}
	tableFile := chunks.GetDataTypeFile(chunk)
	tableFileExists := util.FileExists(tableFile)

	if chunk.Required && !tableFileExists {
		return false, fmt.Errorf("table file not found: %s", tableFile)
	}

	if tableFileExists {
		table, err := parseConfig(chunk)
		if err != nil {
			return false, err
		}

		if len(table.Rows) == 0 {
			return false, fmt.Errorf("no rows configured for table in file: %s", tableFile)
		}
	}

	return tmplExists && tableFileExists, nil
}

func TypeHandler(chunk config.Chunk) (chunks.Processed, error) {
	table, err := parseConfig(chunk)
	if err != nil {
		return chunks.Processed{}, err
	}
	table.Columns = getColumns(chunk)
	table.removeUnusedColumns()
	return chunks.Processed{Data: table}, nil
}

func parseConfig(chunk config.Chunk) (Table, error) {
	bytes, err := os.ReadFile(chunks.GetDataTypeFile(chunk))
	if err != nil {
		return Table{}, err
	}

	table := Table{}
	if err = yaml.Unmarshal(bytes, &table); err != nil {
		return Table{}, err
	}
	return table, nil
}

func getColumns(chunk config.Chunk) []Columns {
	if cols, exists := chunk.Data["columns"]; exists && cols != nil {
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
				columns = append(columns, Columns{ID: id, Label: util.AnyToString(label)})
			}
		} else {
			panic("one columns spec should a single key/value: " + chunk.String())
		}
	}
	return columns
}
