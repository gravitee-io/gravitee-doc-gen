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
	"fmt"
	"os"
	"path/filepath"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
)

var registry = struct {
	data           map[string]interface{}
	handlers       map[string]FileHandler
	exports        map[string]string
	postProcessors map[string]PostProcessor
}{
	data:           make(map[string]interface{}),
	handlers:       make(map[string]FileHandler),
	exports:        make(map[string]string),
	postProcessors: make(map[string]PostProcessor),
}

type FileHandler func(file string) (any, error)

type PostProcessor func(any) (any, error)

func Register(handler FileHandler, ext ...string) {
	if len(ext) == 0 {
		panic("Register handler must have at least one extension")
	}
	for _, ext := range ext {
		registry.handlers[ext] = handler
	}
}

func RegisterPostProcessor(key string, processor PostProcessor) {
	registry.postProcessors[key] = processor
}

func GetData(name string) any {
	if data, ok := registry.data[name]; ok {
		return data
	}
	panic(fmt.Sprintf("'%s' bootstrap data does not exist", name))
}

func GetExported() map[string]any {
	exported := make(map[string]any)
	for k, v := range registry.exports {
		exported[v] = GetData(k)
	}
	return exported
}

func load(file string, export string) (any, error) {
	stat, err := os.Stat(file)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, fmt.Errorf("%s is a directory, should be a file", file)
	}

	var val any
	var key string
	if handle, ok := registry.handlers[filepath.Ext(file)]; ok {
		val, err = handle(file)
		if err != nil {
			return nil, err
		}
		key = filepath.Base(util.BaseFileNoExt(file))
		registry.data[key] = val
	} else {
		panic(fmt.Sprintf("no '%s' handler for bootstrap file: %s ", filepath.Ext(file), file))
	}

	if postProcessor, ok := registry.postProcessors[key]; ok {
		updated, err := postProcessor(val)
		if err != nil {
			return nil, err
		}
		registry.data[key] = updated
	}
	registry.exports[key] = export
	return val, nil
}
