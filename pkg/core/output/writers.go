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

package output

import (
	"fmt"
	"os"
	"path/filepath"
)

type Console struct {
	To string
}
type File struct {
	To string
}

func (f File) Write(generated []byte) (int, error) {
	stat, err := os.Stat(f.To)

	if err != nil && os.IsNotExist(err) {
		dir := filepath.Dir(f.To)
		err := os.MkdirAll(dir, 0744)
		if err != nil {
			return 0, err
		}
		newFile, err := os.Create(f.To)
		if err != nil {
			return 0, err
		}

		stat, err = newFile.Stat()
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	err = os.WriteFile(f.To, generated, stat.Mode().Perm())
	if err != nil {
		return 0, err
	}

	return len(generated), nil
}

func (c Console) Write(generated []byte) (int, error) {
	fmt.Println("--- Dry Run (" + c.To + ") ---")
	fmt.Println(string(generated))
	fmt.Println("-----------------------")
	fmt.Println()
	return len(generated), nil
}
