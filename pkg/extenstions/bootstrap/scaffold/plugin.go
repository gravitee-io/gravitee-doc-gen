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

package scaffold

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
)

func PluginScaffolder() error {
	data := bootstrap.GetData("plugin")
	pl, _ := data.(plugin.Plugin)
	rootDir, _ := bootstrap.GetData(bootstrap.RootDirDataKey).(string)
	destinationDir := "."
	sourceDir := path.Clean(path.Join(rootDir, pl.Type, "__scaffold"))
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// skip dirs
		if info.IsDir() {
			return nil
		}

		in, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		destinationFile := filepath.Join(destinationDir, rel)
		err = os.MkdirAll(filepath.Dir(destinationFile), 0755)
		if err != nil {
			return err
		}

		if _, err := os.Stat(destinationFile); os.IsNotExist(err) {
			return os.WriteFile(destinationFile, in, info.Mode())
		} else if err != nil {
			return err
		}
		fmt.Println("Scaffold:", rel, "already exists")
		return nil
	})
}
