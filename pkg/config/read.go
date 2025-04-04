package config

import (
	"errors"
	"fmt"
	"github.com/gravitee-io-labs/readme-gen/pkg/util"

	"gopkg.in/yaml.v3"
	"os"
	"path"
)

const defaultMainTemplateFile = "README.tmpl"

func Read(rootDir string, plugin Plugin) (Config, error) {

	// TODO make a ConfigFileResolver
	specificConfig := path.Join(rootDir, plugin.Type, plugin.Id+".yaml")
	defaultConfig := path.Join(rootDir, plugin.Type, "/default.yaml")
	var configFile string
	if stat, err := os.Stat(specificConfig); err == nil && !stat.IsDir() {
		configFile = specificConfig
	} else if stat, err := os.Stat(defaultConfig); err == nil && !stat.IsDir() {
		configFile = defaultConfig
	} else {
		return Config{}, errors.New(fmt.Sprintf("Cannot find %s or %s ", specificConfig, defaultConfig))
	}

	fmt.Println("README Generation config: ", configFile)
	rendered, err := util.RenderTemplateFromFile(configFile, map[string]interface{}{"RootDir": rootDir, PluginChunkId: plugin})

	defaultMainTemplate := path.Join(rootDir, plugin.Type, defaultMainTemplateFile)
	if err != nil {
		return Config{}, err
	}

	// read the config
	config := Config{
		MainTemplate: defaultMainTemplate,
	}
	err = yaml.Unmarshal(rendered, &config)
	return config, err

}
