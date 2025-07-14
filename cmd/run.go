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

package cmd

import (
	"fmt"
	"os"

	"github.com/gravitee-io/gravitee-doc-gen/pkg/core"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/bootstrap"
	"github.com/gravitee-io/gravitee-doc-gen/pkg/core/output"
	"github.com/spf13/cobra"
)

func MainCommand() *cobra.Command {
	var main = &cobra.Command{
		Use:   "readme-gen [OPTIONS]",
		Short: "Yield README.md for Gravitee plugins",
		Run:   run,
	}

	main.Flags().StringVarP(
		&optionsData.Init,
		"scaffold",
		"s",
		"",
		"Use a scaffolder to initialize a project with typical files. e.g 'plugin'")
	main.Flags().BoolVarP(&optionsData.Validate, "validate", "v", false, "Run validation only.")
	main.Flags().BoolVarP(&optionsData.DryRun, "dry-run", "d", true, "Run generation, write result to console")
	main.Flags().BoolVarP(&optionsData.Write, "write", "w", false, "Run generation, write result to console")
	main.Flags().StringVarP(
		&optionsData.ConfigFile,
		"config-file",
		"f",
		"",
		"Forces a config file (relative to config-dir) "+
			"instead of using $DOCGEN_ROOT/default.yaml or the bootstrap config loader")
	main.Flags().StringVarP(
		&optionsData.RootDir,
		"config-dir",
		"c",
		"",
		"Configuration directory, contains bootstrap and per plugin directory")
	main.MarkFlagsMutuallyExclusive("dry-run", "write")

	return main
}

var optionsData struct {
	Init       string
	Validate   bool
	DryRun     bool
	Write      bool
	RootDir    string
	ConfigFile string
}

const rootEnvVar = "DOCGEN_ROOT"

func run(_ *cobra.Command, _ []string) {
	rootDir := optionsData.RootDir
	if rootDir == "" {
		rootDir = os.Getenv(rootEnvVar)
		if rootDir == "" {
			fail(fmt.Errorf("env variable %s must be set when %s is not", rootEnvVar, "--config-dir or -c"))
		}
	}

	if optionsData.Init != "" {
		if err := bootstrap.Load(rootDir); err != nil {
			fail(fmt.Errorf("error initializing project: %w", err))
		}
		if err := bootstrap.Scaffold(optionsData.Init); err != nil {
			fail(err)
		}
		return
	}

	generated, cfg, genError := core.Load(rootDir, optionsData.ConfigFile)

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
