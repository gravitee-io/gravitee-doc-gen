package cmd

import (
	"fmt"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/output"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/extenstions/bootstrap/plugin"
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
	main.Flags().StringVarP(&optionsData.RootDir, "config-dir", "c", "", "Configuration directory, contains bootstrap and per plugin directory")
	main.MarkFlagsMutuallyExclusive("dry-run", "write")

	return main
}

var optionsData struct {
	Validate bool
	DryRun   bool
	Write    bool
	RootDir  string
}

const rootEnvVar = "RMG_ROOT"

func run(_ *cobra.Command, _ []string) {

	rootDir := optionsData.RootDir
	if rootDir == "" {
		rootDir = os.Getenv(rootEnvVar)
		if rootDir == "" {
			fail(fmt.Errorf("env variable %s must be set when %s is not", rootEnvVar, "--config-dir or -c"))
		}
	}

	generated, cfg, genError := core.Load(rootDir, func(_ string) (string, error) {
		return plugin.PluginRelatedFile("default.yaml")
	})

	if optionsData.Validate {
		if genError != nil {
			fmt.Println("Validation failed")
			failIf(genError)
			failIf(genError)
		} else {
			fmt.Println("Validation OK")
		}
	} else {
		failIf(genError)
		fmt.Println("Chunks generated... writing outputs")
		for _, out := range cfg.Outputs {
			err := output.Yield(out, generated, optionsData.Write)
			failIf(err)
		}
	}

}

func failIf(err error) {
	if err != nil {
		fail(err)
	}
}

func fail(err error) {
	fmt.Println("Error:", err)
	os.Exit(1)
}
