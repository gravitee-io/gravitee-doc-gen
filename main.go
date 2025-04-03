package main

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/cmd"
	"github.com/gravitee-io-labs/readme-gen/pkg/config"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/code"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/common"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/options"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator/types/table"
	"os"
)

func main() {

	generator.Registry.Register(config.UnknownDataType, common.NoopTypeHandler, common.TemplateExistsTypeValidator)
	generator.Registry.Register(config.TableDataType, table.TypeHandler, table.TypeValidator)
	generator.Registry.Register(config.CodeDataType, code.TypeHandler, code.TypeValidator)
	generator.Registry.Register(config.Options, options.TypeHandler, options.TypeValidator)

	err := cmd.MainCommand().Execute()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
