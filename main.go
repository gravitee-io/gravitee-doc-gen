package main

import (
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/cmd"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/chunks"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/config"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/generator"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/filehandlers"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/code"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/gen_examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/options"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/raw_examples"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/schema_to_env"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/schema_to_yaml"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/generator/table"
	"os"
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
	generator.Registry.Register(GenExamples, gen_examples.TypeHandler, gen_examples.TypeValidator)
	generator.Registry.Register(RawExamples, raw_examples.TypeHandler, raw_examples.TypeValidator)
	generator.Registry.Register(SchemaToYaml, schema_to_yaml.TypeHandler, schema_to_yaml.TypeValidator)
	generator.Registry.Register(SchemaToEnv, schema_to_env.TypeHandler, schema_to_env.TypeValidator)

	bootstrap.Register(filehandlers.PropertiesFileHandler, filehandlers.PropertiesExt)
	bootstrap.Register(filehandlers.YamlFileHandler, filehandlers.YamlExt, filehandlers.YmlExt)
	bootstrap.Register(filehandlers.JsonFileHandler, filehandlers.JsonExt)

	bootstrap.RegisterPostProcessor("plugin", plugin.PluginPostProcessor)
	bootstrap.RegisterPostProcessor("gen-examples", examples.GenExamplePostProcessor)

	err := cmd.MainCommand().Execute()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
