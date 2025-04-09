package config

import (
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/bootstrap"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"

	"gopkg.in/yaml.v3"
)

type FileResolver func(string) (string, error)

func Read(rootDir string, resolver FileResolver) (Config, error) {

	configFile, err := resolver(rootDir)
	if err != nil {
		return Config{}, err
	}
	fmt.Println("Generation config: ", configFile)

	rendered, err := util.RenderTemplateFromFile(configFile, bootstrap.GetExported())

	// read the config
	config := Config{}

	err = yaml.Unmarshal(rendered, &config)
	return config, err
}
