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

	data := make(map[string]interface{})
	for key, export := range bootstrap.Registry.GetExports() {
		data[export] = bootstrap.Registry.GetData(key)
	}
	data["RootDir"] = rootDir
	rendered, err2 := util.RenderTemplateFromFile(configFile, data)

	// read the config
	config := Config{}

	err2 = yaml.Unmarshal(rendered, &config)
	return config, err2
}
