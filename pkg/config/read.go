package config

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"

	"gopkg.in/yaml.v3"
	"os"
	"path"
)

const rootEnvVar = "RMG_ROOT"
const defaultMainTemplateFile = "README.tmpl"

func Read(plugin Plugin) (Config, error) {

	root := os.Getenv(rootEnvVar)
	if root == "" {
		return Config{}, errors.New(rootEnvVar + " is not set")
	}

	specificConfig := path.Join(root, plugin.Type, plugin.Id+".yaml")
	defaultConfig := path.Join(root, plugin.Type, "/default.yaml")
	var configFile string
	if stat, err := os.Stat(specificConfig); err == nil && !stat.IsDir() {
		configFile = specificConfig
	} else if stat, err := os.Stat(defaultConfig); err == nil && !stat.IsDir() {
		configFile = defaultConfig
	} else {
		return Config{}, errors.New(fmt.Sprintf("Cannot find %s or %s ", specificConfig, defaultConfig))
	}

	fmt.Println("README Generation config: ", configFile)
	rendered, err := util.RenderTemplateFromFile(configFile, map[string]interface{}{rootEnvVar: root, PluginChunkId: plugin})

	defaultMainTemplate := path.Join(root, plugin.Type, defaultMainTemplateFile)
	if err != nil {
		return Config{
			MainTemplate: defaultMainTemplate,
		}, err
	}

	// read the config
	config := Config{}
	err = yaml.Unmarshal(rendered, &config)
	return config, err

}
