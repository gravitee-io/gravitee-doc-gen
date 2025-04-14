package main

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/cmd"
	"github.com/gravitee-io-labs/readme-gen/pkg"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap/handlers"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/examples"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/code"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/gen_examples"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/options"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/raw_examples"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/schema_to_env"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/schema_to_yaml"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/table"
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
	generator.Registry.Register(config.UnknownDataType, common.NoopTypeHandler, common.TemplateExistsTypeValidator)
	generator.Registry.Register(TableDataType, table.TypeHandler, table.TypeValidator)
	generator.Registry.Register(CodeDataType, code.TypeHandler, code.TypeValidator)
	generator.Registry.Register(Options, options.TypeHandler, options.TypeValidator)
	generator.Registry.Register(GenExamples, gen_examples.TypeHandler, gen_examples.TypeValidator)
	generator.Registry.Register(RawExamples, raw_examples.TypeHandler, raw_examples.TypeValidator)
	generator.Registry.Register(SchemaToYaml, schema_to_yaml.TypeHandler, schema_to_yaml.TypeValidator)
	generator.Registry.Register(SchemaToEnv, schema_to_env.TypeHandler, schema_to_env.TypeValidator)

	bootstrap.Register(handlers.PropertiesFileHandler, handlers.PropertiesExt)
	bootstrap.Register(handlers.YamlFileHandler, handlers.YamlExt, handlers.YmlExt)
	bootstrap.Register(handlers.JsonFileHandler, handlers.JsonExt)

	bootstrap.RegisterPostProcessor("plugin", pkg.PluginPostProcessor)
	bootstrap.RegisterPostProcessor("gen-examples", examples.GenExamplePostProcessor)

	err := cmd.MainCommand().Execute()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
