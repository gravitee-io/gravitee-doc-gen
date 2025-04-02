package cmd

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg"
	"github.com/gravitee-io-labs/readme-gen/pkg/chunks"
	"github.com/gravitee-io-labs/readme-gen/pkg/generator"
	"github.com/gravitee-io-labs/readme-gen/pkg/output"
	"github.com/spf13/cobra"
	"os"
)

func MainCommand() *cobra.Command {
	var main = &cobra.Command{
		Use:   "readme-gen [OPTIONS]",
		Short: "Yield README.md for Gravitee plugins",
		Run:   run,
	}

	main.Flags().BoolVarP(&optionsData.Validate, "validate", "v", false, "Run validation only.")
	main.Flags().BoolVarP(&optionsData.DryRun, "dry-run", "d", true, "Run generation, write result to console")
	main.Flags().BoolVarP(&optionsData.Write, "write", "w", false, "Run generation, write result to console")
	main.MarkFlagsMutuallyExclusive("dry-run", "write")

	return main
}

var optionsData struct {
	Validate bool
	DryRun   bool
	Write    bool
}

func run(_ *cobra.Command, _ []string) {

	readyChunks, pl, cfg, loadErr := pkg.Load()
	var genError error
	var generated []chunks.Generated
	if loadErr == nil {
		generated, genError = generator.Generate(readyChunks)
	}

	if optionsData.Validate {
		if loadErr != nil || genError != nil {
			fmt.Println("Validation failed")
			failIf(loadErr)
			failIf(genError)
		} else {
			fmt.Println("Validation OK")
		}
	} else {
		failIf(loadErr)
		failIf(genError)
		fmt.Println("README Generated... writing")
		err := output.Yield(cfg, pl, generated, optionsData.Write)
		failIf(err)
	}

}

func fail(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
func failIf(err error) {
	if err != nil {
		fail(err)
	}
}
