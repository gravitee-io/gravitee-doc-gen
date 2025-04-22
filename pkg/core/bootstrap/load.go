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

package bootstrap

import (
	"os"
	"path/filepath"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"gopkg.in/yaml.v3"
)

const RootDirDataKey = "rootDir"

type data struct {
	Filename   string `yaml:"file"`
	ExportedAs string `yaml:"exportedAs"`
}

type fileContent struct {
	Data []data `yaml:"data"`
}

func Load(rootDir string) error {
	bootstrapFile := filepath.Join(rootDir, "bootstrap.yaml")
	_, err := os.Stat(bootstrapFile)
	if err != nil {
		return err
	}

	// add this here so any one can use it
	registry.data["rootDir"] = rootDir
	registry.exports["rootDir"] = "RootDir"

	content, err := util.RenderTemplateFromFile(bootstrapFile, GetExported())
	if err != nil {
		return err
	}

	bootstrap := &fileContent{}
	err = yaml.Unmarshal(content, bootstrap)
	if err != nil {
		return err
	}

	for _, data := range bootstrap.Data {
		_, err := load(data.Filename, data.ExportedAs)
		if err != nil {
			return err
		}
	}

	return nil
}
