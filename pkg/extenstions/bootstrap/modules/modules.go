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

package modules

import (
	"fmt"
	"path/filepath"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/util"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/filehandlers"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
	"github.com/rickar/props"
)

// DataKey is the bootstrap data the modules are loaded as, from .docgen/modules.yaml.
const DataKey = "modules"

// pluginFile is where a module declares the plugin it builds, relative to the module path.
const pluginFile = "src/main/resources/plugin.properties"

type Module struct {
	ID         string `json:"id"`
	Path       string `json:"path"`
	Name       string `json:"name"`
	ExportedAs string `json:"exportedAs"`
	// Plugin is the plugin the module builds, read from its own plugin.properties; nil when the module has none,
	// a library bundled in the plugins for instance.
	Plugin *plugin.Plugin `json:"-"`
}

type modulesFile struct {
	Modules []Module `json:"modules"`
}

func PostProcessor(data any) (any, error) {
	raw := util.As[*util.Unstructured](data)
	if raw == nil {
		return nil, fmt.Errorf("modules data is nil")
	}

	mf, err := util.AnyMapToStruct[modulesFile](raw)
	if err != nil {
		return nil, fmt.Errorf("cannot parse modules data: %w", err)
	}

	if len(mf.Modules) == 0 {
		return nil, fmt.Errorf("at least one module must be declared")
	}

	for i, m := range mf.Modules {
		if m.Path == "" {
			return nil, fmt.Errorf("module at index %d: path is required", i)
		}
		if m.ExportedAs == "" {
			return nil, fmt.Errorf("module at index %d: exportedAs is required", i)
		}
		file := filepath.Join(m.Path, pluginFile)
		if !util.FileExists(file) {
			continue
		}
		p, err := loadPlugin(file)
		if err != nil {
			return nil, fmt.Errorf("module at index %d: %w", i, err)
		}
		mf.Modules[i].Plugin = &p
	}

	return mf.Modules, nil
}

func loadPlugin(file string) (plugin.Plugin, error) {
	read, err := filehandlers.PropertiesFileHandler(file)
	if err != nil {
		return plugin.Plugin{}, fmt.Errorf("cannot read %s: %w", file, err)
	}
	properties, ok := read.(*props.Properties)
	if !ok {
		return plugin.Plugin{}, fmt.Errorf("cannot read %s as properties", file)
	}
	processed, err := plugin.PostProcessor(properties)
	if err != nil {
		return plugin.Plugin{}, fmt.Errorf("invalid %s: %w", file, err)
	}
	loaded, ok := processed.(plugin.Plugin)
	if !ok {
		return plugin.Plugin{}, fmt.Errorf("cannot read %s as a plugin", file)
	}
	return loaded, nil
}

// Find returns the declared module of the given id, false when there is none or when no module is declared.
func Find(id string) (Module, bool) {
	declared, ok := bootstrap.FindData(DataKey)
	if !ok {
		return Module{}, false
	}
	all, ok := declared.([]Module)
	if !ok {
		return Module{}, false
	}
	for _, m := range all {
		if m.ID == id {
			return m, true
		}
	}
	return Module{}, false
}
