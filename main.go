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

package main

import (
	"fmt"
	"os"

	"github.com/gravitee-io/gravitee-doc-gen/cmd"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/generator"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/filehandlers"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/code"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/genexamples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/options"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/rawexamples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/schematoenv"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/schematoyaml"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/table"
)

const TableDataType = config.DataType("table")
const CodeDataType = config.DataType("code")
const Options = config.DataType("options")
const GenExamples = config.DataType("gen-examples")
const RawExamples = config.DataType("raw-examples")
const SchemaToYaml = config.DataType("schema-to-yaml")
const SchemaToEnv = config.DataType("schema-to-env")

func main() {
	generator.Registry.Register(config.UnknownDataType, chunks.NoopTypeHandler, chunks.TemplateExistsTypeValidator)
	generator.Registry.Register(TableDataType, table.TypeHandler, table.TypeValidator)
	generator.Registry.Register(CodeDataType, code.TypeHandler, code.TypeValidator)
	generator.Registry.Register(Options, options.TypeHandler, options.TypeValidator)
	generator.Registry.Register(GenExamples, genexamples.TypeHandler, genexamples.TypeValidator)
	generator.Registry.Register(RawExamples, rawexamples.TypeHandler, rawexamples.TypeValidator)
	generator.Registry.Register(SchemaToYaml, schematoyaml.TypeHandler, schematoyaml.TypeValidator)
	generator.Registry.Register(SchemaToEnv, schematoenv.TypeHandler, schematoenv.TypeValidator)

	bootstrap.Register(filehandlers.PropertiesFileHandler, filehandlers.PropertiesExt)
	bootstrap.Register(filehandlers.YamlFileHandler, filehandlers.YamlExt, filehandlers.YmlExt)
	bootstrap.Register(filehandlers.JSONFileHandler, filehandlers.JSONExt)

	bootstrap.RegisterPostProcessor("plugin", plugin.PostProcessor)
	bootstrap.RegisterPostProcessor("default-examples", examples.GenExamplePostProcessor)

	config.RegisterConfigResolver("plugin", func(string, string) (string, error) {
		return plugin.RelativeFile("default.yaml")
	})

	err := cmd.MainCommand().Execute()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
